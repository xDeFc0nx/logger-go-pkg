package logger

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/sanity-io/slog"
)

var decoration = `%c $: -------------------------------------------------------------------`

// Info logs informational messages with formatted string.
func Info(format string, args ...interface{}) {
	logger := slog.New(os.Stdout, slog.ColorDecoratedOutput)
	logger.Info().Msg(fmt.Sprintf(decoration+"\nINFO: "+format, append([]interface{}{color.GreenString()}, args...)...))
}

// Warning logs warning messages with formatted string.
func Warning(format string, args ...interface{}) {
	logger := slog.New(os.Stdout, slog.ColorDecoratedOutput)
	logger.Warn().Msg(fmt.Sprintf(decoration+"\nWARNING: "+format, append([]interface{}{color.YellowString()}, args...)...))
}

// Error logs error messages with formatted string.
func Error(format string, args ...interface{}) {
	logger := slog.New(os.Stderr, slog.ColorDecoratedOutput)
	logger.Error().Msg(fmt.Sprintf(decoration+"\nERROR: "+format, append([]interface{}{color.RedString()}, args...)...))
}

// Success logs success messages with formatted string.
func Success(format string, args ...interface{}) {
	logger := slog.New(os.Stdout, slog.ColorDecoratedOutput)
	logger.Info().Msg(fmt.Sprintf(decoration+"\nSUCCESS: "+format, append([]interface{}{color.MagentaString()}, args...)...))
}

// Debug logs debug messages with formatted string.
func Debug(format string, args ...interface{}) {
	logger := slog.New(os.Stdout, slog.ColorDecoratedOutput)
	logger.Debug().Msg(fmt.Sprintf(decoration+"\nDEBUG: "+format, append([]interface{}{color.BlueString()}, args...)...))
}
