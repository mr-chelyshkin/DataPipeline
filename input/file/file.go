package file

import (
	"context"

	"github.com/hpcloud/tail"
	"github.com/mr-chelyshkin/DataPipeline/input/common"
)

// Source ...
type Source struct {
	tailLogger *internalLogger
	filepath   string
	message    chan common.SourceMessage
}

// NewSource ...
func NewSource(path string) *Source {
	return &Source{
		tailLogger: &internalLogger{stream: make(chan common.SourceLog, 1)},
		message:    make(chan common.SourceMessage, 1),

		filepath: path,
	}
}

// GetMessageChanel return channel with input data.
func (s *Source) GetMessageChanel() chan common.SourceMessage {
	return s.message
}

// GetLogChannel return channel with log messages.
func (s *Source) GetLogChannel() chan common.SourceLog {
	return s.tailLogger.stream
}

// Serve ...
func (s *Source) Serve(ctx context.Context) {
	t, _ := tail.TailFile(s.filepath, tail.Config{
		Logger: s.tailLogger,
		Follow: true,
		ReOpen: true,
		Poll:   true,
	})
	for line := range t.Lines {
		s.message <- common.SourceMessage{Body: line.Text}
	}
}
