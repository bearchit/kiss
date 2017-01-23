package kiss

import "github.com/mangoplate/kiss/log"

type Kiss struct {
	Logger *log.Logger
}

type Config struct {
	Logger *log.Logger
}

func New() *Kiss {
	return &Kiss{
		Logger: log.New(),
	}
}

func (k *Kiss) Handler(hf handlerFunc) *handler {
	return &handler{
		Handler: hf,
		Logger:  k.Logger,
	}
}
