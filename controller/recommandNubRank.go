package controller

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"example.com/m/v2/db"
	"example.com/m/v2/model"
	"example.com/m/v2/module"
	"github.com/gin-gonic/gin"
)

type Stock struct {
	Stock_name string
	Ranks      Stocks_rank
}

type Stocks_rank struct {
	Yester_rank      int
	Fluctuation_rank int
	Cap_rank         int
	Per_rank         int
}

type Last_Stock struct {
	Stock_name string
	rank       int
}

type Kv struct {
	Key   string
	Value int
}

func RecommandNubRank(c *gin.Context) {

	/*err := module.Crawling()
	if err != nil {
		log.Println(err)
		module.Response(c, http.StatusBadRequest, "get rank find error")
		return
	}*/

	orgStockInfos := []model.StockInfo{}
	err := db.DB.Model(&model.StockInfo{}).Select("item_name", "then_yesterday", "fluctuation_rate", "market_cap", "per").Find(&orgStockInfos).Error
	if err != nil {
		log.Println(err)
		module.Response(c, http.StatusBadRequest, "get rank find error")
		return
	}

	stocks := make([]Stock, len(orgStockInfos))
	last_stock := make([]Last_Stock, len(orgStockInfos))

	Yester_rank := map[string]int{}
	Fluctuation_rank := map[string]int{}
	Cap_rank := map[string]int{}
	Per_rank := map[string]int{}

	var yes_kv []Kv
	var flu_kv []Kv
	var cap_kv []Kv
	var per_kv []Kv

	for i, v := range orgStockInfos {
		stocks[i].Stock_name = v.Item_name
	}

	// 주식정보 가져와서 숫자로 변환
	for _, v := range orgStockInfos {

		Yester_rank[v.Item_name], _ = strconv.Atoi(v.Then_yesterday)

		Cap_rank[v.Item_name], _ = strconv.Atoi(v.Market_cap)

		v.Fluctuation_rate = strings.ReplaceAll(v.Fluctuation_rate, "%", "")
		v.Fluctuation_rate = module.EditDigit(v.Fluctuation_rate)
		Fluctuation_rank[v.Item_name], _ = strconv.Atoi(v.Fluctuation_rate)

		if v.Per == "N/A" {
			Per_rank[v.Item_name] = 0
		} else {
			v.Per = module.EditDigit(v.Per)
			Per_rank[v.Item_name], _ = strconv.Atoi(v.Per)
		}
	}

	// 숫자로 변환한 주식정보를 정렬
	for k, v := range Yester_rank {
		yes_kv = append(yes_kv, Kv{k, v})
	}

	for k, v := range Fluctuation_rank {
		flu_kv = append(flu_kv, Kv{k, v})
	}

	for k, v := range Cap_rank {
		cap_kv = append(cap_kv, Kv{k, v})
	}

	for k, v := range Per_rank {
		per_kv = append(per_kv, Kv{k, v})
	}

	sort.Slice(yes_kv, func(i, j int) bool {
		return yes_kv[i].Value > yes_kv[j].Value
	})

	sort.Slice(flu_kv, func(i, j int) bool {
		return flu_kv[i].Value > flu_kv[j].Value
	})

	sort.Slice(cap_kv, func(i, j int) bool {
		return cap_kv[i].Value > cap_kv[j].Value
	})

	sort.Slice(per_kv, func(i, j int) bool {
		return per_kv[i].Value > per_kv[j].Value
	})

	for i, v := range yes_kv {
		for j, v2 := range stocks {
			if v.Key == v2.Stock_name {
				stocks[j].Ranks.Yester_rank = i
			}
		}
	}
	for i, v := range flu_kv {
		for j, v2 := range stocks {
			if v.Key == v2.Stock_name {
				stocks[j].Ranks.Fluctuation_rank = i
			}
		}
	}
	for i, v := range cap_kv {
		for j, v2 := range stocks {
			if v.Key == v2.Stock_name {
				stocks[j].Ranks.Cap_rank = i
			}
		}
	}
	for i, v := range per_kv {
		for j, v2 := range stocks {
			if v.Key == v2.Stock_name {
				stocks[j].Ranks.Per_rank = i
			}
		}
	}

	for i, v := range stocks {
		last_stock[i].Stock_name = v.Stock_name
		last_stock[i].rank = module.Average(v.Ranks.Cap_rank, v.Ranks.Fluctuation_rank, v.Ranks.Yester_rank, v.Ranks.Per_rank)
	}

	sort.Slice(last_stock, func(i, j int) bool {
		return last_stock[i].rank > last_stock[j].rank
	})

	fmt.Println(last_stock)

	c.JSON(http.StatusAccepted, gin.H{
		"status": http.StatusAccepted,
		"data":   last_stock,
	})
}
