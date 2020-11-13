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

	// !!! месяц день год поменять на день месяц год по заданию
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
	// !!! русифицировать месяца
	// 20 октября 2018 года,19 часов,35 минут, 21 секунда.
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

func (h *Handler) getAdd(c *gin.Context) {
	// !!! парсинги вывести в отдельную функцию ->float64
	// req*  = request from handler
	reqTime := c.DefaultQuery("time", "010206.150405")
	reqDelta := c.DefaultQuery("delta", "010206.000000")
	startT, err := time.Parse("010206.150405", reqTime)
	if err != nil {
		panic(err)
	}
	fmt.Print(startT)

	data := strings.Split(reqDelta, ".")

	// Parse date from string
	boxdate := data[0]
	boxDatestr := strings.Split(boxdate, "")

	btMonth := boxDatestr[0] + boxDatestr[1]
	btDay := boxDatestr[2] + boxDatestr[3]
	btYear := boxDatestr[4] + boxDatestr[5]

	monthInt, err := strconv.Atoi(btMonth)
	dayInt, err := strconv.Atoi(btDay)
	yearInt, err := strconv.Atoi(btYear)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	dateT := [3]int{monthInt, dayInt, yearInt}
	endDate := startT.AddDate(dateT[0], dateT[1], dateT[2])

	// Parse time from string
	boxtime := data[1]
	boxTimestr := strings.Split(boxtime, "")

	btHour := boxTimestr[0] + boxTimestr[1]
	btMinute := boxTimestr[2] + boxTimestr[3]
	btSecond := boxTimestr[4] + boxTimestr[5]

	hourInt, err := strconv.Atoi(btHour)
	minuteInt, err := strconv.Atoi(btMinute)
	secondInt, err := strconv.Atoi(btSecond)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	timeT := [3]int{hourInt, minuteInt, secondInt}

	endTime := startT.Add(time.Hour*time.Duration(timeT[0]) + time.Minute*time.Duration(timeT[1]) + time.Second*time.Duration(timeT[2]))

	endDateAndTime := startT.AddDate(dateT[0], dateT[1], dateT[2]).Add(time.Hour*time.Duration(timeT[0]) + time.Minute*time.Duration(timeT[1]) + time.Second*time.Duration(timeT[2]))
	formatDateTime := actualtime(endDateAndTime)
	c.JSON(200, gin.H{"time": formatDateTime, "endDate": endDate, "endTime": endTime})
}

func (h *Handler) postCorrect(c *gin.Context) {

}
