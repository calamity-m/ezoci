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
	log    *log.Logger
	pretty bool
}

func (c *InterceptHandler) Handle(ctx context.Context, r slog.Record) error {

	if c.pretty {
		var out = fmt.Sprintf("%s | %s", r.Level, r.Message)

		switch r.Level {
		case slog.LevelDebug:
			out = colorize(darkGray, out)
		case slog.LevelWarn:
			out = colorize(lightYellow, out)
		case slog.LevelError:
			out = colorize(lightRed, out)
		}

		c.log.Println(
			colorize(lightGray, r.Time.String()),
			out,
		)
	} else {
		c.log.Println(r.Message)
	}

	return nil
}

func newInterceptHandler(w io.Writer, opts *slog.HandlerOptions, structured bool, pretty bool, debug bool) *InterceptHandler {
	if opts == nil {
		opts = &slog.HandlerOptions{}
	}

	if debug {
		opts.Level = slog.LevelDebug
	}

	var h slog.Handler
	if structured {
		h = slog.NewJSONHandler(w, opts)
	} else {
		h = slog.NewTextHandler(w, opts)
	}

	return &InterceptHandler{
		log:     log.New(w, "", 0),
		Handler: h,
		pretty:  pretty,
	}
}

func NewInterceptedTextHandler(w io.Writer, pretty bool, debug bool, opts *slog.HandlerOptions) *InterceptHandler {
	return newInterceptHandler(w, opts, false, pretty, debug)
}

func NewInterceptedJSONHandler(w io.Writer, pretty bool, debug bool, opts *slog.HandlerOptions) *InterceptHandler {
	return newInterceptHandler(w, opts, true, pretty, debug)
}
