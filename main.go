package main

import (
	"github.com/fluxchain/core/node"
	"github.com/fluxchain/core/parameters"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.SetLevel(logrus.DebugLevel)

	logrus.Info("starting flux...")

	n := node.New()
	n.Bootstrap("database.db", parameters.Main)
	defer n.Teardown()

	go n.Mine(50)
	n.RegisterRPC()
	n.ServeRPC(":3030")
}
