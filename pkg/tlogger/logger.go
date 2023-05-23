package tlogger

import (
	"fmt"
	"gb-auth-gate/config"
	"gb-auth-gate/pkg/terrors/interface"
	"github.com/afiskon/promtail-client/promtail"
	"github.com/sarulabs/di"
	"github.com/sirupsen/logrus"
	"time"
)

type TLogger struct {
	Config *config.Config `di:"config"`
	loki   promtail.Client
	logger *logrus.Logger
}

func NewLogger() ILogger {

	return &TLogger{}
}

func BuildLogger(ctn di.Container) (interface{}, error) {
	cfg := ctn.Get("config").(*config.Config)

	logger := logrus.New()
	logger.Formatter = &logrus.TextFormatter{
		FullTimestamp: true,
	}

	loki, err := promtail.NewClientProto(promtail.ClientConfig{
		PushURL: fmt.Sprintf("http://%s:%s",
			cfg.LoggerConfig.Loki.Host, cfg.LoggerConfig.Loki.Port),
		SendLevel:          promtail.INFO,
		PrintLevel:         promtail.DISABLE,
		BatchEntriesNumber: cfg.LoggerConfig.Loki.Batch.Number,
		BatchWait:          time.Duration(cfg.LoggerConfig.Loki.Batch.Wait) * time.Second,
		Labels: fmt.Sprintf("{source=\"%s\"}",
			cfg.BaseConfig.Service.Name),
	})
	if err != nil {
		return nil, err
	}
	return &TLogger{
		Config: cfg,
		logger: logger,
		loki:   loki,
	}, nil
}

func (T *TLogger) log(level logrus.Level, msg string, args ...interface{}) {
	T.logger.Logf(level, msg, args...)
}

func (T *TLogger) sendLog(msg string, level LogLevel) {
	switch level {
	case ERROR:
		T.loki.Errorf(msg)
		break
	case WARN:
		T.loki.Warnf(msg)
		break
	case DEBUG:
		T.loki.Debugf(msg)
		break
	default:
		T.loki.Infof(msg)
	}
}

func (T *TLogger) Infof(msgf string, args ...interface{}) {
	T.log(logrus.InfoLevel, msgf, args...)
}

func (T *TLogger) Debugf(msgf string, args ...interface{}) {
	T.log(logrus.DebugLevel, msgf, args...)
}

func (T *TLogger) Warnf(msgf string, args ...interface{}) {
	T.log(logrus.WarnLevel, msgf, args...)
}

func (T *TLogger) Errorf(msgf string, args ...interface{}) {
	T.log(logrus.ErrorLevel, msgf, args...)
}

func (T *TLogger) ErrorFull(err error, requestId string) {
	if tErr, ok := err.(_interface.IError); ok {
		go T.sendLog(fmt.Sprintf("source=%s %s requestId=%s",
			T.Config.BaseConfig.Service.Name,
			tErr.LoggerMessage(), requestId),
			ERROR)
	} else {
		go T.sendLog(fmt.Sprintf(labels,
			T.Config.BaseConfig.Service.Name,
			500, err.Error(), requestId),
			ERROR)
	}
	T.Errorf("Error: %s", err.Error())
}

// Write implementing for putting it into fiber logger
func (T *TLogger) Write(p []byte) (n int, err error) {

	return len(p), nil
}
