package main

import (
	"os"

	"github.com/sirupsen/logrus"

	"github.com/j0suetm-com/jtm_svc/api"
	"github.com/j0suetm-com/jtm_svc/util"
)

func main() {
	cfg, err := util.LoadCfg("config.json")
	if err != nil {
		logrus.Error(err)

		os.Exit(1)
	}

	rtr, err := api.New(*cfg)
	if err != nil {
		logrus.Error(err)

		os.Exit(1)
	}

	rtr.Run(":" + cfg.Server.Port)
}
