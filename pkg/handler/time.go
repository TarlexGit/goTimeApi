package handler

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func actualtime(t time.Time) float64 {
	fmt.Println(t)

	// месяц день год поменять на день месяц год по заданию
	fdt := t.Format("010206.150405")
	fmt.Println(fdt)
	// Перевод строки во float64
	var x float64 = 0
	if fdt, err := strconv.ParseFloat(fdt, 64); err == nil {
		x = fdt
	}
	return x
}

func (h *Handler) getNow(c *gin.Context) {
	t := time.Now()
	c.JSON(200, gin.H{"time": actualtime(t)})
}

func stringTime(t time.Time) string {
	// 20 октября 2018 года,19 часов,35 минут, 21 секунда.
	// русифицировать месяца
	dateStr := fmt.Sprintf(
		"%v %v %v года, %v часов %v минут %v секунда",
		t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute(), t.Second(),
	)
	return string(dateStr)
}

func (h *Handler) getString(c *gin.Context) {
	r := c.DefaultQuery("time", "Guest")
	tr, err := time.Parse("010206.150405", r)

	if err != nil {
		panic(err)
	}

	date := stringTime(tr)
	c.JSON(200, gin.H{"str": date})
}

// func (h *Handler) getAdd(c *gin.Context) {
// 	// req*  = request from handler
// 	reqTime := c.DefaultQuery("time", "010206.150405")
// 	reqDelta := c.DefaultQuery("delta", "010206.152255")

// 	timeT, err := time.Parse("010206.150405", reqTime)
// 	deltaT, err := time.Parse("010206.150405", reqDelta)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Разница между датами
// 	firstTime := timeT.Unix()
// 	secondTime := deltaT.Unix()
// 	deltaMinute := (secondTime - firstTime) / 60
// 	fmt.Println(deltaMinute)

// 	newT := timeT.Add(time.Minute * time.Duration(deltaMinute))
// 	fmt.Printf("Adding 1 hour\n: %s\n", newT)

// 	c.JSON(200, gin.H{"time": deltaMinute})
// }

func (h *Handler) getAdd(c *gin.Context) {
	// req*  = request from handler
	reqTime := c.DefaultQuery("time", "010206.150405")
	reqDelta := c.DefaultQuery("delta", "010206.000000")
	startT, err := time.Parse("010206.150405", reqTime)
	if err != nil {
		panic(err)
	}
	fmt.Print(startT)

	data := strings.Split(reqDelta, ".")
	boxdate := data[0]

	// Parse time from string
	boxtime := data[1]
	boxTimestr := strings.Split(boxtime, "")
	// t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute(), t.Second()
	btHour := boxTimestr[0] + boxTimestr[1]
	btMinute := boxTimestr[2] + boxTimestr[3]
	btSecond := boxTimestr[4] + boxTimestr[5]

	hourInt, err := strconv.Atoi(btHour)
	minuteInt, err := strconv.Atoi(btMinute)
	secondInt, err := strconv.Atoi(btSecond)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	timeT := [3]int{hourInt, minuteInt, secondInt}

	// var x0 time.Duration
	// if fdt, err := strconv.ParseInt(btHour,0); err == nil {
	// 	x = fdt
	// }

	endTime := startT.Add(time.Hour*time.Duration(timeT[0]) + time.Minute*time.Duration(timeT[1]) + time.Second*time.Duration(timeT[2]))

	c.JSON(200, gin.H{"full": timeT, "boxdate": boxdate, "startT": startT, "endTime": endTime})
}

func (h *Handler) postCorrect(c *gin.Context) {

}
