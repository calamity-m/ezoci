package logging

import (
	"context"
	"fmt"
	"io"
	"log"
	"log/slog"
	"strconv"
)

const (
	reset = "\033[0m"

	black        = 30
	red          = 31
	green        = 32
	yellow       = 33
	blue         = 34
	magenta      = 35
	cyan         = 36
	lightGray    = 37
	darkGray     = 90
	lightRed     = 91
	lightGreen   = 92
	lightYellow  = 93
	lightBlue    = 94
	lightMagenta = 95
	lightCyan    = 96
	white        = 97
)

func colorize(colorCode int, v string) string {
	return fmt.Sprintf("\033[%sm%s%s", strconv.Itoa(colorCode), v, reset)
}

type InterceptHandler struct {
	slog.Handler
	log     *log.Logger
	pretty  bool
	verbose bool
}

func (c *InterceptHandler) Handle(ctx context.Context, r slog.Record) error {

	format := ""
	if c.verbose {
		format = fmt.Sprintf("%s | %s | ", r.Time.String(), r.Level)
	}
	format = fmt.Sprintf("%s%s", format, r.Message)

	if c.pretty {
		switch r.Level {
		case slog.LevelDebug:
			format = colorize(darkGray, format)
		case slog.LevelWarn:
			format = colorize(lightYellow, format)
		case slog.LevelError:
			format = colorize(lightRed, format)
		}

	}

	c.log.Println(format)
	return nil
}

func NewInterceptHandler(w io.Writer, opts *slog.HandlerOptions, pretty bool, verbose bool) *InterceptHandler {
	if opts == nil {
		opts = &slog.HandlerOptions{}
	}

	if verbose {
		opts.Level = slog.LevelDebug
	}

	return &InterceptHandler{
		log:     log.New(w, "", 0),
		Handler: slog.NewTextHandler(w, opts),
		pretty:  pretty,
		verbose: verbose,
	}
}
