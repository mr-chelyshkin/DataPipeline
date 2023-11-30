package file

import (
	"fmt"

	"github.com/mr-chelyshkin/DataPipeline/input/common"
)

type internalLogger struct {
	stream chan common.SourceLog
}

// Fatal ...
func (l *internalLogger) Fatal(v ...interface{}) {
	l.stream <- common.SourceLog{
		Body:  fmt.Sprint(v),
		Level: 5,
	}
}

// Fatalf ...
func (l *internalLogger) Fatalf(format string, v ...interface{}) {
	l.stream <- common.SourceLog{
		Body:  fmt.Sprintf(format, v...),
		Level: 5,
	}
}

// Fatalln ...
func (l *internalLogger) Fatalln(v ...interface{}) {
	l.stream <- common.SourceLog{
		Body:  fmt.Sprint(v),
		Level: 5,
	}
}

// Panic ...
func (l *internalLogger) Panic(v ...interface{}) {
	l.stream <- common.SourceLog{
		Body:  fmt.Sprint(v),
		Level: 6,
	}
}

// Panicf ...
func (l *internalLogger) Panicf(format string, v ...interface{}) {
	l.stream <- common.SourceLog{
		Body:  fmt.Sprintf(format, v...),
		Level: 5,
	}
}

// Panicln ...
func (l *internalLogger) Panicln(v ...interface{}) {
	l.stream <- common.SourceLog{
		Body:  fmt.Sprint(v),
		Level: 6,
	}
}

// Print ...
func (l *internalLogger) Print(v ...interface{}) {
	l.stream <- common.SourceLog{
		Body:  fmt.Sprint(v),
		Level: 1,
	}
}

// Printf ...
func (l *internalLogger) Printf(format string, v ...interface{}) {
	l.stream <- common.SourceLog{
		Body:  fmt.Sprintf(format, v...),
		Level: 1,
	}
}

// Println ...
func (l *internalLogger) Println(v ...interface{}) {
	l.stream <- common.SourceLog{
		Body:  fmt.Sprint(v),
		Level: 1,
	}
}
