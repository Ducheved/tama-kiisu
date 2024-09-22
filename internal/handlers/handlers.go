package handlers

import (
	"net/http"

	"github.com/Ducheved/tama-kiisu/internal/services"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	catService *services.CatService
}

func NewHandlers(catService *services.CatService) *Handlers {
	return &Handlers{catService: catService}
}

func (h *Handlers) IndexHandler(c *gin.Context) {
	cat, err := h.catService.GetCat()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "index.html", gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"cat": cat,
	})
}

func (h *Handlers) ActionHandler(c *gin.Context) {
	action := c.PostForm("action")
	ip := c.ClientIP()

	cat, err := h.catService.PerformAction(action, ip)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cat)
}
