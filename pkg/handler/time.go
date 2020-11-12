package handler

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func actualtime(t time.Time) float64 {
	h := float64(t.Hour())
	hh := float64(t.Hour())
	m := float64(t.Minute())
	s := float64(t.Second())
	fldata := []float64{h, m, s}
	fmt.Print(fldata)
	return hh
}

func (h *Handler) getNow(c *gin.Context) {
	t := time.Now()
	// time := t.Format("03:04:05")
	c.JSON(200, gin.H{"time": actualtime(t)})
}

func (h *Handler) getString(c *gin.Context) {

}

func (h *Handler) getAdd(c *gin.Context) {

}

func (h *Handler) postCorrect(c *gin.Context) {

}
