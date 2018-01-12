package fix

import (
	"github.com/quickfixgo/quickfix"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	HandlerConfig
}

type HandlerConfig struct {
	settings *quickfix.Settings
	log      *logrus.Logger
}

func NewHandler(cfg HandlerConfig) *Handler {
	return &Handler{cfg}
}

func (h *Handler) OnCreate(sessionID quickfix.SessionID) {
	h.sLog(sessionID).Infoln("OnCreate")
}

func (h *Handler) OnLogon(sessionID quickfix.SessionID) {
	h.sLog(sessionID).Infoln("OnLogon")
}

func (h *Handler) OnLogout(sessionID quickfix.SessionID) {
	h.sLog(sessionID).Infoln("OnLogout")
}

func (h *Handler) FromAdmin(message *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
	h.sLog(sessionID).Infof("FromAdmin: %s\n", message.String())
	return nil
}

func (h *Handler) ToAdmin(message *quickfix.Message, sessionID quickfix.SessionID) {
	h.sLog(sessionID).Infof("ToAdmin: %s\n", message.String())
}

func (h *Handler) FromApp(message *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
	h.sLog(sessionID).Infof("FromApp: %s\n", message.String())
	return nil
}

func (h *Handler) ToApp(message *quickfix.Message, sessionID quickfix.SessionID) error {
	h.sLog(sessionID).Infof("ToApp: %s\n", message.String())
	return nil
}

func (h *Handler) sLog(sessionID quickfix.SessionID) *logrus.Entry {
	return h.log.WithField("sessionID", sessionID.String())
}
