package main

import (
	"github.com/fluxchain/core/node"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.SetLevel(logrus.DebugLevel)

	logrus.Info("starting flux...")

	n := node.New()
	n.Bootstrap("database.db")
	defer n.Teardown()
	n.Mine(50)
}
