package handler

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func convertStrFloat64(data string) float64 {
	var x float64 = 0
	if data, err := strconv.ParseFloat(data, 64); err == nil {
		x = data
	}
	return x
}

func parseStringData(data []string, startT time.Time) float64 {
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
	// endDate := startT.AddDate(dateT[0], dateT[1], dateT[2])

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

	endDateAndTime := startT.AddDate(dateT[0], dateT[1], dateT[2]).Add(time.Hour*time.Duration(timeT[0]) + time.Minute*time.Duration(timeT[1]) + time.Second*time.Duration(timeT[2]))
	formatDateTime := endDateAndTime.Format("010206.150405")

	x := convertStrFloat64(formatDateTime)
	return x
}

func actualtime(t time.Time) float64 {
	// !!! месяц день год поменять на день месяц год по заданию
	fdt := t.Format("010206.150405")
	startT, err := time.Parse("010206.150405", fdt)
	delta := string(deltaFromFile())
	data := strings.Split(delta, ".")
	if err != nil {
		panic(err)
	}
	endTime := parseStringData(data, startT)
	fmt.Print(endTime)
	return endTime
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

func deltaToFile(t string) {
	// Каждый раз перезаписывает файл
	mydata := []byte(t)
	err := ioutil.WriteFile("pkg/handler/temp/file.txt", mydata, 0777)
	if err != nil {
		fmt.Println(err)
	}
}

func deltaFromFile() string {
	data, err := ioutil.ReadFile("pkg/handler/temp/file.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(data))
	return string(data)
}
