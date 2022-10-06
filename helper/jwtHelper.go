package helper

import (
	"context"
	"fmt"
	"time"

	"example.com/m/v2/redis"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TokenDetails struct {
	AccessToken string
	AccessUuid  string // redis에 저장할 때 key값
	AtExpires   int64
}

type AccessDetails struct { // redis에 접근하기 위한 구조체
	AccessUuid string
	UserId     string
}

// token 추출
func ExtractToken(c *gin.Context) string {

	token := c.Request.Header.Get("token")

	if token == "" {
		return ""
	}

	return token
}

// token의 파싱 방법을 찾아 비교하고 token형태로 반환
func VerifyToken(c *gin.Context) (*jwt.Token, error) {

	tokenString := ExtractToken(c)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected singing method : %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// token 유효성 검사
func TokenVaild(c *gin.Context) error {

	token, err := VerifyToken(c)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

// token에서 redis에 접근하기위 한 access uuid와 유저의 id를 추출
func ExtractTokenMetadata(c *gin.Context) (*AccessDetails, error) {

	token, err := VerifyToken(c)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userId := fmt.Sprintf("%s", claims["id"])

		if err != nil {
			return nil, err
		}
		return &AccessDetails{
			AccessUuid: accessUuid,
			UserId:     userId,
		}, nil
	}
	return nil, err
}

var accessString = []byte("secret") // access token의 key

// 토큰 제작
func CreateJWT(id string) (*TokenDetails, error) {

	var err error

	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Hour * 1).Unix() // 엑세스 토큰 15분
	td.AccessUuid = uuid.NewString()

	// 엑세스 토큰
	atClaims := jwt.MapClaims{}
	atClaims["id"] = id
	atClaims["exp"] = td.AtExpires
	atClaims["access_uuid"] = td.AccessUuid

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims) // claims를 담아 토큰 만들기
	td.AccessToken, err = at.SignedString([]byte(accessString))
	if err != nil {
		return nil, err
	}

	return td, nil
}

func StoreAuth(td *TokenDetails) error {

	at := time.Unix(td.AtExpires, 0)
	now := time.Now()

	errAccess := redis.Client.Set(context.Background(), td.AccessUuid, td.AccessToken, at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}

	return nil
}

func DeleteAuth(dtk string) (int64, error) {

	dresult, err := redis.Client.Del(context.Background(), dtk).Result()
	if err != nil {
		return 0, err
	}
	return dresult, nil
}

func CheckAuth(authD *AccessDetails) (string, error) {

	userId, err := redis.Client.Get(context.Background(), authD.AccessUuid).Result()
	if err != nil {
		return "", nil
	}

	return userId, nil
}
