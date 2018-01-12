package fix

import (
	"fmt"

	"github.com/quickfixgo/quickfix"
	"github.com/sirupsen/logrus"
)

type Initiator struct {
	log      *logrus.Logger
	internal *quickfix.Initiator
}

func NewInitiator(cfg ConnectorConfig) (*Initiator, error) {
	handler := NewHandler(HandlerConfig{
		log:      cfg.Log,
		settings: cfg.Settings,
	})

	internal, err := quickfix.NewInitiator(handler, cfg.StoreFactory, cfg.Settings, cfg.LogFactory)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize internal initiator: %v", err)
	}

	return &Initiator{
		log:      cfg.Log,
		internal: internal,
	}, nil
}

func (i *Initiator) Start() error {
	return i.internal.Start()
}

func (i *Initiator) Stop() {
	i.internal.Stop()
}
