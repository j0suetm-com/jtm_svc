package main

import (
	"github.com/sirupsen/logrus"

	"github.com/j0suetm-com/jtm_svc/api"
	"github.com/j0suetm-com/jtm_svc/util"
)

func main() {
	cfg, err := util.LoadCfg("config.json")
	if err != nil {
		logrus.Error(err)
	}

	rtr, err := api.New(*cfg)
	if err != nil {
		logrus.Error(err)
	}

	rtr.Run(":" + cfg.Server.Port)
}
