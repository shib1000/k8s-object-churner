package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shib1000/k8s-object-churner/pkg/config"
	"github.com/shib1000/k8s-object-churner/pkg/handler"
	"github.com/shib1000/k8s-object-churner/pkg/k8s"
	koclog "github.com/shib1000/k8s-object-churner/pkg/log"
	"github.com/shib1000/k8s-object-churner/pkg/middleware"
	"github.com/shib1000/k8s-object-churner/pkg/worker"
	"go.uber.org/zap"
)

func main() {
	appCfg, err := config.NewConfigMgr()
	if err != nil {
		panic(err.Error())
	}
	k8sClient, err := k8s.NewK8sClient(appCfg)
	if err != nil {
		panic(err.Error())
	}
	router := gin.New()
	router.Use(gin.Logger())
	appLogger := koclog.GetKocLoggerInstance()
	router.Use(func(context *gin.Context) {
		handler.SetConfigMgrInContext(context, appCfg)
		handler.SetLoggerInContext(context, appLogger)
		context.Next()
	})
	router.Use(handler.JSONAppErrorReporter)
	router.Use(gin.Recovery())

	router.Use(middleware.NewRequestIdMiddlerWare().Handle)
	router.GET("/sample", handler.NewSampleHandler().HandleGET)
	router.GET("/ping", handler.NewHealthHandler().HandleGET)
	router.GET("/k8s/pod/:nsname/:podname", handler.NewPodsDetailsHandler(k8sClient).HandleGET)

	router.GET("/k8s/cfmap/:nsname/:cmlabelkey/:cmlabelvalue", handler.NewConfigMapHandler(k8sClient).HandleGET)
	router.POST("/k8s/cfmap/:nsname", handler.NewConfigMapHandler(k8sClient).HandlePOST)
	router.DELETE("/k8s/cfmap/:nsname/:cfname", handler.NewConfigMapHandler(k8sClient).HandleDELETE)

	//start workers
	wrkerNS := appCfg.GetConfigString("WORKER_NS")
	appLogger.Info("Started ConfigMap Churner", zap.String("Namespace", wrkerNS))
	worker.NewConfigMapChurner(k8sClient, wrkerNS)

	//start inet
	router.Run(":8080")
}
