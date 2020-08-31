package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Hello world")

	port := os.Getenv("PORT")
	if port == "" {
		logrus.Fatal("Port is not set")
	}
}
