package kiss

import "github.com/bearchit/kiss/log"

type logger log.Logger

type Kiss struct {
	Logger *logger
}

type Config struct {
	Logger *logger
}

func New() *Kiss {
	return &Kiss{
		Logger: (*logger)(log.New()),
	}
}

func (k *Kiss) Handler(hf handlerFunc) *handler {
	return &handler{
		Handler: hf,
		Logger:  k.Logger,
	}
}
