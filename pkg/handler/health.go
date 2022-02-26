package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthHandler struct{}

func NewHealthHandler() Handler {
	return &HealthHandler{}
}

func (hh *HealthHandler) HandleGET(c *gin.Context) {
	log := GetLoggerrFromContext(c)
	log.Info("Inside Ping API")
	c.IndentedJSON(http.StatusOK, "pong")
}

func (hh *HealthHandler) HandlePOST(c *gin.Context) {
	log := GetLoggerrFromContext(c)
	log.Info("UnimplementedMethod")
	c.IndentedJSON(http.StatusNotImplemented, gin.H{"msg": "Method Not Implemented"})
}

func (hh *HealthHandler) HandleDELETE(c *gin.Context) {
	log := GetLoggerrFromContext(c)
	log.Info("UnimplementedMethod")
	c.IndentedJSON(http.StatusNotImplemented, gin.H{"msg": "Method Not Implemented"})
}
