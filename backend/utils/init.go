package utils

import (
	"time"
)

func init() {
	LoadLocalConfig()
	LoadConsulConfig()
	ticker := time.NewTicker(time.Duration(LocalConf.ConfigTTL) * time.Second)
	go func() {
		for {
			<-ticker.C
			LoadConsulConfig()
		}
	}()

	LoadLog()
}
