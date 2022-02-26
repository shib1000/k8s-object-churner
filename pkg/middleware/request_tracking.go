package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strings"
)

const (
	REQUEST_ID = "x-request-id"
)

type RequestIdMiddlerWare struct {
}

func NewRequestIdMiddlerWare() *RequestIdMiddlerWare {
	return &RequestIdMiddlerWare{}
}
func (hh *RequestIdMiddlerWare) Handle(c *gin.Context) {
	requestId := c.GetHeader(REQUEST_ID)
	if strings.TrimSpace(requestId) == "" {
		requestId = uuid.New().String()
	}
	c.Set(REQUEST_ID, requestId)
	c.Next()
}
