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

	n := node.New()
	n.Bootstrap()
	n.Mine(50)
}
