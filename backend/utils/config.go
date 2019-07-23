package utils

import (
	"os"
	"path"
	"runtime"

	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/consul"
)

// Config map the main config
type Config struct {
	Deploy    string            `json:"deploy"`
	Hosts     map[string]string `json:"hosts"`
	ConfigTTL int               `json:"config_ttl"`
	Test      string            `json:"test"`
}

// LocalConf is local config
var LocalConf Config

// ConsulConf is consul config
var ConsulConf config.Config

// LoadLocalConfig load config from local
func LoadLocalConfig() {
	_, filename, _, ok := runtime.Caller(0)
	LogPanic(ok, "No caller information")
	p := os.Getenv("JJ_CONFIGPATH")
	if p != "" {
		LogPanic(config.LoadFile(p))
	} else {
		LogPanic(config.LoadFile(path.Dir(filename) + "/../config.json"))
	}
	LogPanic(config.Scan(&LocalConf))
	Info("Local config loaded")
}

// GetDeployHost return current deployed host
func GetDeployHost(name string) string {
	return LocalConf.Hosts[name+"_"+LocalConf.Deploy]
}

// LoadConsulConfig load config from consul
func LoadConsulConfig() {
	consulSource := consul.NewSource(
		consul.WithAddress(GetDeployHost("consul")),
		consul.WithPrefix(""),
	)

	// we need tmp var to prevent race
	tmp := config.NewConfig()
	LogPanic(tmp.Load(consulSource))
	ConsulConf = tmp
	Info("Consul config loaded")
}

// GetStringConfig get string config from consul
func GetStringConfig(path ...string) string {
	return ConsulConf.Get(path...).String("")
}
