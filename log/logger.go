package log

import (
	"os"
	"strings"

	"errors"

	"github.com/Sirupsen/logrus"
)

const (
	File   = "file"
	Stdout = "stdout"
	Stderr = "stderr"

	JSON = "json"
	Text = "text"
)

var (
	ErrInvalidOutput = errors.New("Invalid output")
	ErrInvalidFormat = errors.New("Invalid format")
	ErrInvalidLevel  = errors.New("Invalid level")
	ErrPathNotGiven  = errors.New("File path must given when using file output")
)

type Logger struct {
	*logrus.Logger
}

func New() *Logger {
	return &Logger{Logger: logrus.New()}
}

func (l *Logger) SetOutput(out string, path string) error {
	switch strings.ToLower(out) {
	case File:
		if len(strings.TrimSpace(path)) <= 0 {
			return ErrPathNotGiven
		}

		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
		if err != nil {
			return err
		}
		l.Out = f

	case Stdout:
		l.Out = os.Stdout

	case Stderr:
		l.Out = os.Stderr

	default:
		return ErrInvalidOutput
	}

	return nil
}

func (l *Logger) SetLevel(level string) error {
	lv, err := logrus.ParseLevel(level)
	if err != nil {
		return ErrInvalidLevel
	}

	l.Level = lv
	return nil
}

func (l *Logger) SetFormat(format string) error {
	switch format {
	case JSON:
		l.Formatter = new(logrus.JSONFormatter)

	case Text:
		l.Formatter = new(logrus.TextFormatter)

	default:
		return ErrInvalidFormat
	}

	return nil
}

func (l *Logger) SetErrorHook() {
	l.Hooks.Add(&ErrorHook{})
}
