package common

import "websocket-tester/pkg/log"

var (
	Version   = "v1.8.0"
	GitHash   = ""
	BuildTime = ""
	GoVersion = ""
)

func DisplayVersionInfo() {
	log.Logger().Printf("version -> %s", Version)
	if GitHash != "" {
		log.Logger().Printf("git hash -> %s", GitHash)
	}
	if BuildTime != "" {
		log.Logger().Printf("build time -> %s", BuildTime)
	}
	if GoVersion != "" {
		log.Logger().Printf("go version -> %s", GoVersion)
	}
}
