package common

import (
	"github.com/mr-chelyshkin/DataPipeline"
)

type SourceMessage struct {
	Body string
}

type SourceLog struct {
	Level DataPipeline.LogLevel
	Body  string
}
