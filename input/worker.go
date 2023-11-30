package input

import (
	"github.com/mr-chelyshkin/DataPipeline/input/file"
)

func NewFileProducer(logger logger) *Producer {
	return &Producer{
		logger: logger,
		sources: []source{
			file.NewSource("/tmp/test.log"),
		},
	}
}

func newProducer() {}
