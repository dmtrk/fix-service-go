package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/fmovlex/fix-service-go/pkg/fix"
	"github.com/quickfixgo/quickfix"
	"github.com/sirupsen/logrus"
)

const (
	usage       = "usage: fix /path/to/config.cfg"
	initiator   = "initiator"
	acceptor    = "acceptor"
	defaultRole = acceptor
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println(usage)
		os.Exit(1)
	}

	var log = logrus.New()

	// use json formatter in prod
	log.Out = os.Stdout
	log.Formatter = &logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	}

	settings, err := parseSettings(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	storeFactory := quickfix.NewFileStoreFactory(settings)
	logFactory, err := quickfix.NewFileLogFactory(settings)
	if err != nil {
		log.Fatalf("failed to initialize log factory: %v", err)
	}

	connectorCfg := fix.ConnectorConfig{
		Settings:     settings,
		StoreFactory: storeFactory,
		LogFactory:   logFactory,
		Log:          log,
	}

	role := parseRole(settings)

	var connector fix.Connector
	switch role {
	case acceptor:
		connector, err = fix.NewAcceptor(connectorCfg)
		if err != nil {
			log.Fatalf("failed to initialize fix acceptor: %v", err)
		}
	case initiator:
		connector, err = fix.NewInitiator(connectorCfg)
		if err != nil {
			log.Fatalf("failed to initialize fix initiator: %v", err)
		}
	}

	log.Infoln("starting connector...")

	if err := connector.Start(); err != nil {
		log.Fatalf("failed to start connector: %v", err)
	}

	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt, os.Kill)
	rec := <-interrupt

	log.Infof("got os signal %v, stopping connector", rec)
	connector.Stop()
}

func parseSettings(path string) (*quickfix.Settings, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}
	defer file.Close()

	settings, err := quickfix.ParseSettings(file)
	if err != nil {
		return nil, fmt.Errorf("failed to parse quickfix settings: %v", err)
	}

	return settings, nil
}

func parseRole(settings *quickfix.Settings) string {
	v, err := settings.GlobalSettings().Setting("ConnectionType")
	if err != nil {
		return defaultRole
	}

	switch strings.ToLower(v) {
	case initiator:
		return initiator
	case acceptor:
		return acceptor
	default:
		return defaultRole
	}
}
