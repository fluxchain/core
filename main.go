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
	n.Mine(50)
}
