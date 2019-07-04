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
	if !ok {
		panic("No caller information")
	}
	err := config.LoadFile(path.Dir(filename) + "/../config.json")
	if err != nil {
		panic(err)
	}
	err = config.Scan(&fileConf)
	if err != nil {
		panic(err)
	}

	consulConfig := fileConf.Hosts["consul-"+fileConf.Deploy]
	consulSource := consul.NewSource(
		consul.WithAddress(consulConfig.Address+":"+strconv.Itoa(consulConfig.Port)),
		consul.WithPrefix(""),
	)

	consulConf = config.NewConfig()
	err = consulConf.Load(consulSource)
	if err != nil {
		panic(err)
	}
}

// GetConfig get config from consul
func GetConfig(path ...string) string {
	return consulConf.Get(path...).String("")
}
