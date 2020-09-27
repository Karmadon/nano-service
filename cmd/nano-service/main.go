package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/getsentry/sentry-go"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/karmadon/nano-service/internal/app/nano-service/controllers/application"
)

const Version = "0.0.1"

func init() {
	readConfigs()

	lvl, ok := os.LookupEnv("LOG_LEVEL")
	if !ok {
		lvl = "debug"
	}

	ll, err := log.ParseLevel(lvl)
	if err != nil {
		ll = log.DebugLevel
	}

	log.SetLevel(ll)
}

func main() {
	options := &application.Options{}

	app, err := application.NewApplication(options)
	if err != nil {
		log.Fatal(err)
	}

	err = app.Assemble()
	if err != nil {
		log.Fatal(err)
	}

	err = app.Prepare()
	if err != nil {
		log.Fatal(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		defer sentry.Recover()
		oscall := <-c
		log.Errorf("system call:%+v", oscall)
		app.Stop()
		cancel()
	}()

	app.Start(ctx)
}

func readConfigs() {
	viper.SetConfigName("config.prod")         // name of config file (without extension)
	viper.SetConfigType("yaml")                // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/nano-service/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.nano-service") // call multiple times to add many search paths
	viper.AddConfigPath("./config")            // optionally look for config in the working directory
	viper.AddConfigPath(".")                   // optionally look for config in the working directory

	err := viper.ReadInConfig() // Find and read the config file

	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
