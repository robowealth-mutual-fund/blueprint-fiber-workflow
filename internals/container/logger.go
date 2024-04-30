package container

import (
	log "github.com/robowealth-mutual-fund/stdlog"

	"github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/config"
)

func (c *Container) NewLogger(appConfig config.Config) {
	level := log.DEBUG_LEVEL

	if appConfig.Environment == "PRODUCTION" {
		level = log.INFO_LEVEL
	}

	log.SetGlobalLogLevel(level)
	log.SetGlobalApplicationName(appConfig.PlatformName)
}
