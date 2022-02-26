package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/shib1000/k8s-object-churner/pkg/config"
	error2 "github.com/shib1000/k8s-object-churner/pkg/error"
	"go.uber.org/zap"
)

type Handler interface {
	HandleGET(c *gin.Context)
	HandlePOST(c *gin.Context)
	HandleDELETE(c *gin.Context)
}

func SetConfigMgrInContext(context *gin.Context, cfgmgr *config.ConfigMgr) bool {
	context.Set("APP_CONFIG", cfgmgr)
	return true
}

func GetConfigMgrFromContext(context *gin.Context) *config.ConfigMgr {
	val, found := context.Get("APP_CONFIG")
	if !found {
		return nil
	}
	mgr, ok := val.(*config.ConfigMgr)
	if !ok {
		return nil
	} else {
		return mgr
	}
}

func SetLoggerInContext(context *gin.Context, log *zap.Logger) bool {
	context.Set("LOGGER", log)
	return true
}

func GetLoggerrFromContext(context *gin.Context) *zap.Logger {
	val, found := context.Get("LOGGER")
	if !found {
		return nil
	}
	logger, ok := val.(*zap.Logger)
	if !ok {
		return nil
	} else {
		return logger
	}
}

//Middleware Error Handler in server package
func JSONAppErrorReporter(c *gin.Context) {
	c.Next()
	detectedErrors := c.Errors.ByType(gin.ErrorTypeAny)
	log := GetLoggerrFromContext(c)

	if len(detectedErrors) > 0 {
		log.Debug("Handling APP error")
		err := detectedErrors[0].Err
		var parsedError *error2.HttpError
		switch err.(type) {
		case *error2.HttpError:
			parsedError = err.(*error2.HttpError)
		default:
			log.Error("Error Encounter", zap.Error(err))
			parsedError = error2.NewHttpError()
		}
		// Put the error into response
		c.JSON(parsedError.Code(), parsedError)
		c.Abort()
	}
}
