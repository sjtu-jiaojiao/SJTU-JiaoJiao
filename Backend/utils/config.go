package utils

import (
	"path"
	"runtime"
	"strconv"

	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/consul"
)

// Host map the config "hosts"
type Host struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
}

// Config map the main config
type Config struct {
	Deploy string          `json:"deploy"`
	Hosts  map[string]Host `json:"hosts"`
}

var consulConf config.Config

func configLoad() {
	var fileConf Config
	_, filename, _, ok := runtime.Caller(0)
	LogPanic(ok, "No caller information")
	LogPanic(config.LoadFile(path.Dir(filename) + "/../config.json"))
	LogPanic(config.Scan(&fileConf))

	consulConfig := fileConf.Hosts["consul-"+fileConf.Deploy]
	consulSource := consul.NewSource(
		consul.WithAddress(consulConfig.Address+":"+strconv.Itoa(consulConfig.Port)),
		consul.WithPrefix(""),
	)

	consulConf = config.NewConfig()
	LogPanic(consulConf.Load(consulSource))
}

// GetConfig get config from consul
func GetConfig(path ...string) string {
	return consulConf.Get(path...).String("")
}
