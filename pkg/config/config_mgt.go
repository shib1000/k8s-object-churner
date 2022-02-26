package config

import (
	"github.com/spf13/viper"
)

type ConfigMgr struct {
	cfg *viper.Viper
}

func NewConfigMgr() (*ConfigMgr, error) {
	cfg := viper.New()
	cfg.SetConfigName("default") // name of config file (without extension)
	cfg.SetConfigType("yaml")    // REQUIRED if the config file does not have the extension in the name
	cfg.AddConfigPath("configs")
	cfg.AddConfigPath("/etc/k8s-event-churner/")
	cfg.AddConfigPath(".") // optionally look for config in the working directory
	cfg.AutomaticEnv()
	err := cfg.ReadInConfig() // Find and read the config file

	if err != nil { // Handle errors reading the config file
		return nil, err
	}

	cfg.SetConfigName("config")
	cfg.AddConfigPath(".")
	cfg.MergeInConfig()
	//fmt.Println(fmt.Sprintf("config=%", cfg.Get("KUBECONFIG")))
	return &ConfigMgr{cfg: cfg}, nil
}

func (cm *ConfigMgr) GetConfig(key string) interface{} {
	return cm.cfg.Get(key)
}

func (cm *ConfigMgr) GetConfigString(key string) string {
	s, success := cm.cfg.Get(key).(string)
	if !success {
		return ""
	}
	return s
}
