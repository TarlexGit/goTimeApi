package handler

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getNow(c *gin.Context) {
	t := time.Now()
	c.JSON(200, gin.H{"time": actualtime(t)})
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
	// req*  = request from handler
	reqTime := c.DefaultQuery("time", "010206.150405")
	reqDelta := c.DefaultQuery("delta", "010206.000000")
	startT, err := time.Parse("010206.150405", reqTime)
	if err != nil {
		panic(err)
	}
	fmt.Print(startT)

	data := strings.Split(reqDelta, ".")
	formatDateTime := parseStringData(data, startT)

	c.JSON(200, gin.H{"time": formatDateTime})
}

func (h *Handler) postCorrect(c *gin.Context) {
	r := c.DefaultQuery("time", "Guest")
	deltaToFile(r)
	c.JSON(200, gin.H{"time": deltaFromFile()})
}
