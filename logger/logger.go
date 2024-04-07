package logger

import (
	"log"
	"os"
)


// colors
func makeTextColorized(message string, color string) string {
	return color + message + "\033[0m"
}


// log types
type logger struct {
	color  string
	prefix func(string) string
}

var Error = logger {
	color: "\033[31m",
	prefix: func(message string) string {return "ERROR " + message + " "},
}

var Warning = logger {
	color: "\033[33m",
	prefix: func(message string) string {return "WARNING " + message + " "},
}

var Info = logger {
	color: "\033[32m",
	prefix: func(message string) string {return "INFO " + message + " "},
}

var Custom = logger {
	color: "\033[36m",
	prefix: func(message string) string {
		if message != "" {
			return message + " "
		}
		return "LOG "
	},
}


// logger fuction
func (lg logger) Log(message string, prefix string) {
	logger := log.New(os.Stdout, makeTextColorized(lg.prefix(prefix), lg.color), log.Ldate|log.Ltime)
	logger.Printf(makeTextColorized(message, lg.color))
}