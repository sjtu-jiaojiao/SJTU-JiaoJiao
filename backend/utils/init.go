package utils

import (
	"time"
)

func init() {
	LoadLocalConfig()
	LoadConsulConfig()
	ticker := time.NewTicker(time.Duration(localConf.ConfigTTL) * time.Second)
	go func() {
		for {
			LoadConsulConfig()
			<-ticker.C
		}
	}()

	LoadLog()
}
