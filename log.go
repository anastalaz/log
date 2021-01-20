package log

import (
	"fmt"
	"log"
	"os"
)

var (
	Debugger     = false
	colorDefault = "\033[0m"
	colorError   = "\033[31m"
	colorInfo    = "\033[32m"
	colorDebug   = "\033[36m"

	logDebug = log.New(os.Stdout, colorDebug+"[DEBUG]\t", log.Lshortfile)
	logInfo  = log.New(os.Stdout, colorInfo+"[INFO]\t", log.Lshortfile)
	logError = log.New(os.Stderr, colorError+"[ERROR]\t", log.Lshortfile)
)

// Debug prints the messages to stdout in cyan color only if the Debugger
// flag is set to true
func Debug(messages ...interface{}) {
	if Debugger {
		logDebug.Output(2, fmt.Sprintf("\t%v"+colorDefault, messages))
	}
}

// Info prints the messages to stdout in green color
func Info(messages ...interface{}) {
	logInfo.Output(2, fmt.Sprintf("\t%v"+colorDefault, messages))
}

// Info prints the messages to stderr in red color
func Error(messages ...interface{}) {
	logError.Output(2, fmt.Sprintf("\t%v"+colorDefault, messages))
}
