package worker

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/shib1000/k8s-object-churner/pkg/k8s"
	koclog "github.com/shib1000/k8s-object-churner/pkg/log"
	"go.uber.org/zap"
	"time"
)

type ConfigMapChurner struct {
	nsname string
	client *k8s.K8sClient
	values map[string]string
	labels map[string]string
	log    *zap.Logger
}

const (
	Label_Key   = "koc_cf_churned"
	Label_Value = "true"
)

func NewConfigMapChurner(client *k8s.K8sClient, nsname string) Worker {
	churner := ConfigMapChurner{}
	churner.client = client
	churner.nsname = nsname
	churner.labels = map[string]string{Label_Key: Label_Value}
	churner.values = map[string]string{"dummy": "dummy"}
	churner.log = koclog.GetKocLoggerInstance()
	s := gocron.NewScheduler(time.UTC)
	s.Every(5).Seconds().Do(churner.Churn)
	s.Every(60).Seconds().Do(churner.Delete)
	s.StartAsync()
	return &churner
}

func (wrker *ConfigMapChurner) Churn() {
	cfname := fmt.Sprintf("%s-%s", "koc-cf-obj", RandomString(5))
	_, err := wrker.client.CreateConfigMap(wrker.nsname,
		cfname, wrker.values, wrker.labels)
	if err != nil {
		wrker.log.Error("Encountered Error in Worker at create time", zap.Error(err))
	} else {
		wrker.log.Info("Created New Config Map", zap.String("Name", cfname))
	}
}

func (wrker *ConfigMapChurner) Delete() {
	cflist, err := wrker.client.GetConfigMaps(wrker.nsname, Label_Key, Label_Value)
	if err != nil {
		wrker.log.Error("Encountered Error in Worker at delete time", zap.Error(err))
	} else if len(cflist.Items) == 0 {
		wrker.log.Info("NO Objects Found for Deletion")
	} else {
		currentTime := time.Now()
		for _, cf := range cflist.Items {
			if currentTime.Sub(cf.ObjectMeta.CreationTimestamp.Time).Seconds() > 50 {
				wrker.client.DeleteConfigMap(wrker.nsname, cf.Name)
				if err != nil {
					wrker.log.Error("Encountered Error in Worker at delete time", zap.Error(err))
				}
				wrker.log.Info("Object Deleted Successfully", zap.String("name", cf.Name))
			}

		}
	}

}
