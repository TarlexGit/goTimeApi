package handler

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func actualtime(t time.Time) float64 {
	dt := time.Now()
	fmt.Println(dt)

	// месяц день год поменять на день месяц год по заданию
	fdt := dt.Format("010206.150405")

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

}

func (h *Handler) postCorrect(c *gin.Context) {

}
