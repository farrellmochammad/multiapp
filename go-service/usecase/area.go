package usecase

import (
	models "go-service/model"
	"go-service/repository"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
)

func DeliveryArea(r *gin.RouterGroup) {
	store := persistence.NewInMemoryStore(time.Second)

	r.GET("/area", cache.CachePage(store, time.Minute, func(c *gin.Context) {
		areaList, err := repository.ReadArea()
		if err != nil {
			c.JSON(404, gin.H{
				"status":  "failed",
				"message": "No area data",
			})
		}

		converter, err := repository.ReadConverter()
		if err != nil {
			panic(err)
		}

		for i, _ := range areaList {
			float, err := strconv.ParseFloat(areaList[i].Price, 64)
			if err == nil {
				areaList[i].UsdPrice = strconv.FormatFloat(float/converter.UsdIdr, 'f', 10, 64)
			}
		}

		c.JSON(200, areaList)
	}))

	r.GET("/statistics", cache.CachePage(store, time.Minute, func(c *gin.Context) {
		role := c.GetString("role")
		if strings.ToLower(role) == "admin" {
			var fetch models.Fetch
			c.Bind(&fetch)
			if fetch.AreaProvinsi == "" || fetch.Week == "" {
				c.JSON(403, gin.H{
					"status":  "failed",
					"message": "empty Querystring ",
				})
				return
			}

			t, _ := time.Parse(time.RFC3339, strings.Replace(fetch.Week, " ", "+", 1))
			_, week := t.ISOWeek()

			areaList, err := repository.ReadArea()
			if err != nil {
				panic(err)
			}

			var numbers []float64
			for i, _ := range areaList {
				t, _ := time.Parse(time.RFC3339, strings.Replace(areaList[i].TglParsed, " ", "+", 1))
				_, weekdate := t.ISOWeek()
				if strings.ToLower(areaList[i].AreaProvinsi) == strings.ToLower(fetch.AreaProvinsi) && weekdate == week {
					number, _ := strconv.ParseFloat(areaList[i].Price, 64)
					numbers = append(numbers, number)
				}
			}

			if len(numbers) == 0 {
				c.JSON(404, gin.H{
					"status":  "success",
					"message": "No statistic data with name area_provinsi and week",
				})
			} else {
				min, max, avg, med := countValue(numbers)
				c.JSON(200, gin.H{
					"min":     min,
					"max":     max,
					"average": avg,
					"median":  med,
				})
			}

		} else {
			c.JSON(403, gin.H{
				"status":  "failed",
				"message": "No permission",
			})
		}

	}))
}

func countValue(arr []float64) (float64, float64, float64, float64) {
	min := 10000000.0
	max := -1.0
	sum := 0.0
	median := 0.0
	for _, num := range arr {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
		sum += num
	}

	if len(arr)%2 == 1 {
		median = arr[((len(arr)+1)/2)-1]
	} else {
		median = (arr[(len(arr)/2)-1] + arr[(len(arr)/2+1)-1]) / 2
	}

	return min, max, sum / float64(len(arr)), median
}
