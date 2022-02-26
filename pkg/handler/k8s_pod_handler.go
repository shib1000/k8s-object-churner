package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/shib1000/k8s-object-churner/pkg/k8s"
	"go.uber.org/zap"
	"net/http"
)

type PodsDetailsHandler struct {
	k8sClient *k8s.K8sClient
}

func NewPodsDetailsHandler(k8sClient *k8s.K8sClient) Handler {
	return &PodsDetailsHandler{k8sClient: k8sClient}
}

func (hh *PodsDetailsHandler) HandleGET(c *gin.Context) {
	log := GetLoggerrFromContext(c)
	nsname := c.Param("nsname")
	podname := c.Param("podname")
	pod, err := hh.k8sClient.GetPodDetails(nsname, podname)
	if err != nil {
		log.Error("Encountered Error", zap.String("error=", err.Error()))
		c.Error(err)
		return
	}
	log.Info("Inside Get Pod Handler")
	if pod == nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "No Pod Found"})
		return
	} else {
		c.IndentedJSON(http.StatusOK, pod)
	}

}

func (hh *PodsDetailsHandler) HandlePOST(c *gin.Context) {
	log := GetLoggerrFromContext(c)
	log.Info("UnimplementedMethod")
	c.IndentedJSON(http.StatusNotImplemented, gin.H{"msg": "Method Not Implemented"})
}

func (hh *PodsDetailsHandler) HandleDELETE(c *gin.Context) {
	log := GetLoggerrFromContext(c)
	log.Info("UnimplementedMethod")
	c.IndentedJSON(http.StatusNotImplemented, gin.H{"msg": "Method Not Implemented"})
}
