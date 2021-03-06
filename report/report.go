package report

import (
	"github.com/dcso/spyre/config"
	"github.com/dcso/spyre/log"

	"github.com/spf13/afero"
)

var targets []target

func Init() error {
	for _, spec := range config.ReportTargets {
		tgt, err := mkTarget(spec)
		if err != nil {
			return err
		}
		targets = append(targets, tgt)
	}
	log.Noticef("Writing report to %s", config.ReportTargets)
	return nil
}

// AddStringf adds a single message with fmt.Printf-style parameters.
func AddStringf(f string, v ...interface{}) {
	for _, t := range targets {
		t.formatMessage(t.writer, f, v...)
	}
}

func AddFileInfo(file afero.File, description, message string, extra ...string) {
	for _, t := range targets {
		t.formatFileEntry(t.writer, file, description, message, extra...)
	}
}

// Close shuts down all reporting targets
func Close() {
	for _, t := range targets {
		t.finish(t.writer)
		t.writer.Close()
	}
}
