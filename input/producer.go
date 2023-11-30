package input

import (
	"context"
	"fmt"

	"github.com/mr-chelyshkin/DataPipeline"
	"github.com/mr-chelyshkin/DataPipeline/input/common"
)

type logger interface {
	Info(k, v, msg string)
}

type source interface {
	GetMessageChanel() chan common.SourceMessage
	GetLogChannel() chan common.SourceLog
	Serve(ctx context.Context)
}

type Producer struct {
	logger  logger
	sources []source
}

func (p *Producer) Start() {
	p.logger.Info("asd", "asd", "start producer")

	ctx := context.Background()
	for _, source := range p.sources {
		go p.sourceListen(ctx, source)
	}
}

func (p *Producer) sourceListen(ctx context.Context, source source) {
	go source.Serve(ctx)
	for {
		select {
		case log := <-source.GetLogChannel():
			switch log.Level {
			case DataPipeline.DebugLevel:
				p.logger.Info("test", "test", log.Body)
			case DataPipeline.InfoLevel:
				p.logger.Info("test", "test", log.Body)
			case DataPipeline.WarnLevel:
				p.logger.Info("test", "test", log.Body)
			}
		case msg := <-source.GetMessageChanel():
			fmt.Println("AAAAA ", msg.Body)
		}
	}
}
