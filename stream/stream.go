package stream

import (
	"context"
	"time"

	"github.com/mr-chelyshkin/DataPipeline/input"
	"github.com/mr-chelyshkin/DataPipeline/logger"
)

type Producer interface {
	Start()
}

type Stream struct {
	logger *logger.Logger

	producers []Producer
}

func NewStream(ctx context.Context, options ...OptionFunc) (*Stream, error) {
	zeroLogger := logger.NewLogger()
	return &Stream{
		logger: zeroLogger,

		producers: []Producer{input.NewFileProducer(zeroLogger)},
	}, nil
}

func (s *Stream) Start() {
	s.logger.Info("addrr", "asd", "streamer started")

	for _, producer := range s.producers {
		producer.Start()
	}
	time.Sleep(time.Second * 300)
	for {
		select {}
	}
}
