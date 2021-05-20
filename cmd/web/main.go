package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {
	StartApp()
}

func StartApp() {
	initLog()
	readConfiguration()
	startWebServer()
}

func initLog() {
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
}

func startWebServer() {
	r := gin.Default()
	Routes(r)
	r.Run(fmt.Sprintf(":%d", viper.GetInt("server.port")))
}

func readConfiguration() {
	env := flag.String("E", "dev", "Execution environment")
	flag.Parse()

	osValue := os.Getenv("ENV")
	if osValue != "" {
		env = &osValue
	}

	logrus.Infof("Starting wallet api services in %s environment ...", *env)

	viper.AddConfigPath("./cmd/web/config")
	viper.SetConfigName("env_" + *env)

	if err := viper.ReadInConfig(); err != nil {
		logrus.Errorf("error reading configuration from viper: %v", err)
		panic(err)
	}
}
