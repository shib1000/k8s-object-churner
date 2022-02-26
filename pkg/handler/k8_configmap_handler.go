package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/shib1000/k8s-object-churner/pkg/k8s"
	"go.uber.org/zap"
	"net/http"
)

type ConfigMapHandler struct {
	k8sClient *k8s.K8sClient
}

type PostBody struct {
	Name   string `json:"name" binding:"required"`
	Values []struct {
		Key   string `json:"key" binding:"required"`
		Value string `json:"value" binding:"required"`
	} `json:"values" binding:"required"`
}

func NewConfigMapHandler(k8sClient *k8s.K8sClient) Handler {
	return &ConfigMapHandler{k8sClient: k8sClient}
}

func (hh *ConfigMapHandler) HandleGET(c *gin.Context) {
	log := GetLoggerrFromContext(c)
	nsname := c.Param("nsname")
	lname := c.Param("cmlabelkey")
	lval := c.Param("cmlabelvalue")
	cfmpls, err := hh.k8sClient.GetConfigMaps(nsname, lname, lval)
	if err != nil {
		log.Error("Encountered Error", zap.String("error=", err.Error()))
		c.Error(err)
		return
	}
	log.Info("Inside Get ConfigMap Handler")
	if cfmpls == nil || len(cfmpls.Items) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"msg": "No cfmaps Found"})
		return
	} else {
		c.IndentedJSON(http.StatusOK, cfmpls)
	}
}

func (hh *ConfigMapHandler) HandlePOST(c *gin.Context) {
	log := GetLoggerrFromContext(c)
	nsname := c.Param("nsname")
	var bdy PostBody
	c.BindJSON(&bdy)
	valuesMap := make(map[string]string)
	for _, element := range bdy.Values {
		valuesMap[element.Key] = element.Value
	}
	cfmap, err := hh.k8sClient.CreateConfigMap(nsname, bdy.Name, valuesMap, nil)
	if err != nil {
		log.Error("Encountered Error", zap.String("error=", err.Error()))
		c.Error(err)
		return
	}

	c.IndentedJSON(http.StatusOK, cfmap)
}

func (hh *ConfigMapHandler) HandleDELETE(c *gin.Context) {
	log := GetLoggerrFromContext(c)
	nsname := c.Param("nsname")
	cfname := c.Param("cfname")
	err := hh.k8sClient.DeleteConfigMap(nsname, cfname)
	if err != nil {
		log.Error("Encountered Error", zap.String("error=", err.Error()))
		c.Error(err)
		return
	}
	c.Status(http.StatusOK)
}
