package logger

import "log"

type MigrationLogger struct {
	verbose bool
}

func (ml *MigrationLogger) Printf(format string, v ...interface{}) {
	log.Println(format, v)
}

func (ml *MigrationLogger) Verbose() bool {
	return ml.verbose
}
