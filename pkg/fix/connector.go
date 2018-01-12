package fix

import (
	"github.com/quickfixgo/quickfix"
	"github.com/sirupsen/logrus"
)

type Connector interface {
	Start() error
	Stop()
}

type ConnectorConfig struct {
	Settings     *quickfix.Settings
	StoreFactory quickfix.MessageStoreFactory
	LogFactory   quickfix.LogFactory
	Log          *logrus.Logger
}
