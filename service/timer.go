package service

import (
	"github.com/Lazy-Parser/Server/process"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

var counter = 1

func title() string {
	return "timer-" + strconv.Itoa(counter)
}

func titleWithId(id string) string {
	return "timer-" + id
}

func TimerStartHandler(c *gin.Context, pm *process.Manager) {
	pm.Append(process.NewTimer(title()))

	p, _ := pm.Get(title())
	c.JSON(200, gin.H{
		"process": p.GetID(),
		"status":  p.GetStatus().String(),
	})
	counter++
}

func TimerGetHandler(c *gin.Context, pm *process.Manager) {
	counter := c.Param("id")
	id := titleWithId(counter)
	p, ok := pm.Get(id)
	if !ok {
		c.JSON(200, gin.H{
			"Timer did not found": "",
		})
		return
	}

	timer, ok := p.(*process.Timer)
	if !ok {
		c.JSON(200, gin.H{
			"failed tp cast process to Timer": "",
		})
		return
	}

	c.JSON(200, gin.H{
		"process": timer.GetID(),
		"status": timer.GetTime(),
	})
}

type stats struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

func TimerGetAllHandler(c *gin.Context, pm *process.Manager) {
	res := make([]stats, 0, len(pm.GetList()))

	for title, p := range pm.GetList() {
		log.Printf("Name: %s", title)
		res = append(res, stats{ID: title, Status: p.GetStatus().String()})
	}

	c.JSON(200, gin.H{"timers": res})
}

func TimerStopHandler(c *gin.Context, pm *process.Manager) {
	counter := c.Param("id")
	id := titleWithId(counter)

	if ok := pm.Stop(id); ok {
		TimerGetAllHandler(c, pm)
	}
}
