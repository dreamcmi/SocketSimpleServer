package log

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"strings"
	"time"
)

var Log zerolog.Logger

func init() {
	output := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.DateTime}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("[ %-6s]", i))
	}

	Log = log.Output(output).With().Timestamp().Logger()
}
