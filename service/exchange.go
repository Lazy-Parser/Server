package service

import (
	"slices"
	"strings"

	"github.com/Lazy-Parser/Server/process"
	"github.com/gin-gonic/gin"
)

var list = []string{
	"mexc",
}

func IsInList(name string) bool {
	return slices.Contains(list, name)
}

func FireErrorIsNotInList(c *gin.Context, providedName string) {
	c.AbortWithStatusJSON(400, gin.H{"error": "Provided name is not in list. Allowed exchanges: " + strings.Join(list, ",")})
}

func ExList(c *gin.Context) {
	c.JSON(200, gin.H{"exchanges": list})
}

func ExStart(c *gin.Context, pManager *process.Manager) {
	exchangeName := c.Param("name")
	if !IsInList(exchangeName) {
		FireErrorIsNotInList(c, exchangeName)
		return
	}

	title := "mexc"
	if err := pManager.Append(process.NewMexcProcess(title)); err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Error occurred while starting MEXC exchange: " + err.Error()})
	}

	p, _ := pManager.Get(title)
	c.JSON(200, gin.H{"status": "success", "exchange": p.GetStatus()})
}
