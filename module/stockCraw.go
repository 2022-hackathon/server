package module

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"example.com/m/v2/db"
	"example.com/m/v2/model"
	"github.com/PuerkitoBio/goquery"
	"github.com/djimenez/iconv-go"
)

var url string = "https://finance.naver.com/sise/sise_quant.naver"

func Crawling() error {

	fmt.Println("크롤링 시작")

	err := scrape()
	if err != nil {
		return err
	}

	return nil
}

func scrape() error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		return errors.New("status code error : " + res.Status)
	}

	utfBody, err := iconv.NewReader(res.Body, "EUC-KR", "utf-8")
	if err != nil {
		return err
	}

	doc, err := goquery.NewDocumentFromReader(utfBody)
	if err != nil {
		return err
	}

	html, _ := doc.Find("table.type_2").Find("tbody").Html()

	html = strings.ReplaceAll(html, "\t", "")

	htmls := strings.Split(html, "\n")

	last := []string{}

	for i, v := range htmls {

		if strings.Contains(v, `</a>`) {
			last = append(last, v)
		} else if strings.Contains(v, `<span class="tah p11`) {
			v += htmls[i+1]
			v += htmls[i+2]
			last = append(last, v)
		} else if strings.Contains(v, `<td class="number">`) {
			last = append(last, v)
		}
	}

	lastString := ""
	lastStringSlice := []string{}
	flag := false

	for _, v := range last {
		temp := strings.Split(v, "")
		for _, v2 := range temp {

			if v2 == ">" {
				flag = true
			}
			if v2 == "<" {
				flag = false
				if lastString == "" {
					continue
				}
				lastStringSlice = append(lastStringSlice, lastString)
				lastString = ""
			}

			if flag {
				if v2 == ">" {
					continue
				} else if v2 == "" {
					continue
				}
				v2 = strings.TrimSpace(v2)
				lastString += v2
			}
		}
	}

	/*for _, v := range lastStringSlice {
		fmt.Printf("%s===", v)
		time.Sleep(time.Second)
	}*/

	err = insertStock(lastStringSlice)
	if err != nil {
		return err
	}

	return nil
}

const plusNum int = 11

func insertStock(data []string) error {

	stocks := []model.StockInfo{}

	start := 0
	end := plusNum

	cnt := 1

	for {
		stock := &model.StockInfo{}

		for _, v := range data[start:end] {

			v = strings.TrimSpace(v)
			if v == "" {
				continue
			}

			switch cnt {
			case 1:
				stock.Item_name = v
			case 2:
				stock.Current_price = v
			case 3:
				stock.Then_yesterday = v
			case 4:
				stock.Fluctuation_rate = v
			case 5:
				stock.Trading_volume = v
			case 9:
				stock.Market_cap = v
			case 10:
				stock.Per = v
			default:
			}
			cnt++
			if cnt == 11 {
				cnt = 0
			}
		}

		stocks = append(stocks, *stock)
		start += plusNum
		end += plusNum

		if end >= len(data) {
			break
		}
	}

	err := db.DB.Model(&model.StockInfo{}).Create(&stocks).Error
	if err != nil {
		return err
	}

	return nil
}
