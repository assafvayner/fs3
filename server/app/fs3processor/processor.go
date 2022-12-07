package fs3processor

import (
	"log"
)

type Fs3RequestProcessor struct {
	Logger *log.Logger
}

func NewFs3RequestProcessor(logger *log.Logger) *Fs3RequestProcessor {
	return &Fs3RequestProcessor{
		Logger: logger,
	}
}
