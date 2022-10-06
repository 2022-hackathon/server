package model

import "time"

type User struct {
	Id       string `gorm:"size:32; primaryKey" json:"id"`
	Pw       string `gorm:"size:128; NOT_NULL" json:"pw"`
	NickName string `gorm:"size:32" json:"nickname"`
	Money    int    `json:"money"`

	Categorie []Categorie `gorm:"foreignKey:UserId"`
}

type Categorie struct {
	UserCategorie string `gorm:"primaryKey; size:32" json:"categorie"`
	UserId        string `gorm:"size:32" json:"id"`

	Boards []Board `gorm:"foreignKey:BoardCategorie; references:UserCategorie"`
}

type Board struct {
	Title          string `gorm:"size:32" json:"title"`
	Content        string `gorm:"size:512" json:"content"`
	UserId         string `gorm:"size:32" json:"id"`
	Date           string `gorm:"size:32" json:"date"`
	BoardCategorie string `gorm:"size:32" json:"categorie"`
	Boardscol      int    `gorm:"primaryKey; autoIncrement"`

	CreatedAt time.Time
}

type StockInfo struct {
	Item_name        string `gorm:"size:32" json:"name"`
	Current_price    string `gorm:"size:32" json:"current"`     //현제가
	Then_yesterday   string `gorm:"size:32" json:"yesterday"`   //전일비
	Fluctuation_rate string `gorm:"size:32" json:"fluctuation"` //등략률
	Trading_volume   string `gorm:"size:32" json:"tranding"`    //거래량
	Market_cap       string `gorm:"size:32" json:"cap"`         //시가총액
	Per              string `gorm:"size:32"`                    // per(회사의 수익애 비해 주가가 낮은지)-> 낮을수록 안정성이 높음
}
