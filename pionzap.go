package pionzap

import (
	"github.com/ckb20110916/logusezap"
	"github.com/pion/logging"
	"go.uber.org/zap"
	"sync"
)

type logger struct {
	sugared *zap.SugaredLogger
}

func (l *logger) Trace(msg string) {
	l.sugared.Debug(msg)
}

func (l *logger) Tracef(format string, args ...any) {
	l.sugared.Debugf(format, args...)
}

func (l *logger) Debug(msg string) {
	l.sugared.Debug(msg)
}

func (l *logger) Debugf(format string, args ...any) {
	l.sugared.Debugf(format, args...)
}

func (l *logger) Info(msg string) {
	l.sugared.Info(msg)
}

func (l *logger) Infof(format string, args ...any) {
	l.sugared.Infof(format, args...)
}

func (l *logger) Warn(msg string) {
	l.sugared.Warn(msg)
}

func (l *logger) Warnf(format string, args ...any) {
	l.sugared.Warnf(format, args...)
}

func (l *logger) Error(msg string) {
	l.sugared.Error(msg)
}

func (l *logger) Errorf(format string, args ...any) {
	l.sugared.Errorf(format, args...)
}

var (
	Factory = &factory{}
)

type factory struct {
	mutex   sync.Mutex
	loggers []*logger
}

func (f *factory) NewLogger(scope string) logging.LeveledLogger {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	named := logusezap.Logger.Named(scope)
	l := &logger{
		sugared: named.Sugar(),
	}
	f.loggers = append(f.loggers, l)
	return l
}

func (f *factory) SyncAll() {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	for i := range f.loggers {
		_ = f.loggers[i].sugared.Sync()
	}
}
