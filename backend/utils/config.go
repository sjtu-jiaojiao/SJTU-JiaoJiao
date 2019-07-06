package utils

import (
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

var localConf Config
var consulConf config.Config

// LoadLocalConfig load config from local
func LoadLocalConfig() {
	_, filename, _, ok := runtime.Caller(0)
	LogPanic(ok, "No caller information")
	LogPanic(config.LoadFile(path.Dir(filename) + "/../config.json"))
	LogPanic(config.Scan(&localConf))
}

// GetDeployHost return current deployed host
func GetDeployHost(name string) string {
	return localConf.Hosts[name+"_"+localConf.Deploy]
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
	consulConf = tmp
}

// GetStringConfig get string config from consul
func GetStringConfig(path ...string) string {
	return consulConf.Get(path...).String("")
}
