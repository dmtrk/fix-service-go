package fix

import (
	"fmt"

	"github.com/quickfixgo/quickfix"
	"github.com/sirupsen/logrus"
)

type Acceptor struct {
	log      *logrus.Logger
	internal *quickfix.Acceptor
}

func NewAcceptor(cfg ConnectorConfig) (*Acceptor, error) {
	handler := NewHandler(HandlerConfig{
		log:      cfg.Log,
		settings: cfg.Settings,
	})

	internal, err := quickfix.NewAcceptor(handler, cfg.StoreFactory, cfg.Settings, cfg.LogFactory)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize internal acceptor: %v", err)
	}
	return &Acceptor{
		log:      cfg.Log,
		internal: internal,
	}, nil
}

func (a *Acceptor) Start() error {
	return a.internal.Start()
}

func (a *Acceptor) Stop() {
	a.internal.Stop()
}
