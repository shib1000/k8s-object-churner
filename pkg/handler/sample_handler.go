package handler

import (
	"github.com/gin-gonic/gin"
	koclog "github.com/shib1000/k8s-object-churner/pkg/log"
	"net/http"
)

type SampleHandler struct{}

func NewSampleHandler() Handler {
	return &SampleHandler{}
}

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func (hh *SampleHandler) HandleGET(c *gin.Context) {
	log := koclog.GetKocLoggerInstance()
	log.Info("Inside Ping API")
	c.IndentedJSON(http.StatusOK, albums)
}

func (hh *SampleHandler) HandlePOST(c *gin.Context) {
	log := GetLoggerrFromContext(c)
	log.Info("UnimplementedMethod")
	c.IndentedJSON(http.StatusNotImplemented, gin.H{"msg": "Method Not Implemented"})
}

func (hh *SampleHandler) HandleDELETE(c *gin.Context) {
	log := GetLoggerrFromContext(c)
	log.Info("UnimplementedMethod")
	c.IndentedJSON(http.StatusNotImplemented, gin.H{"msg": "Method Not Implemented"})
}
