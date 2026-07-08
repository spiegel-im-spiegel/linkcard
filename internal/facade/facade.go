package facade

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"runtime"

	"github.com/goark/errs"
	"github.com/goark/gocli/exitcode"
	"github.com/goark/gocli/rwi"
	"github.com/goark/struct2pflag"
	"github.com/spf13/pflag"
	"github.com/spiegel-im-spiegel/linkcard/internal/config"
	"github.com/spiegel-im-spiegel/linkcard/internal/linkcard"
)

// Execute is the main function of the application, which executes the command-line interface
func Execute(ui *rwi.RWI, appVersion string, args []string) (exit exitcode.ExitCode) {
	// defer function to catch panic and print stack trace
	defer func() {
		if r := recover(); r != nil {
			_ = ui.OutputErrln("Panic:", r)
			for depth := 0; ; depth++ {
				pc, _, line, ok := runtime.Caller(depth)
				if !ok {
					break
				}
				_ = ui.OutputErrln(" ->", depth, ":", runtime.FuncForPC(pc).Name(), ": line", line)
			}
			exit = exitcode.Abnormal
		}
	}()

	exit = exitcode.Normal
	if err := run(ui, args, NewVersionString(appVersion)); err != nil {
		var jerr *jsonOutputError
		if errors.As(err, &jerr) {
			_ = ui.OutputErrln(jerr.Payload)
			exit = exitcode.Abnormal
			return
		}
		_ = ui.OutputErrln(fmt.Sprintf("error: %v", err))
		exit = exitcode.Abnormal
	}
	return
}

func run(ui *rwi.RWI, args []string, ver versionString) error {
	// Import configuration from file (if not found, use default configuration)
	cfg, err := config.ImportConfigFromFile()
	if err != nil {
		return err
	}

	// Parse command-line arguments and bind them to the configuration struct
	fs := pflag.NewFlagSet(usageString(), pflag.ContinueOnError) // Create a new flag set for command-line options
	fs.SetOutput(ui.Writer())
	struct2pflag.Bind(fs, cfg)             // Bind command-line flags to the configuration struct
	if err := fs.Parse(args); err != nil { // Parse command-line arguments
		return err
	}

	// If the version flag is set, print version information and exit
	if cfg.VersionFlag {
		return ui.Outputln(ver.String())
	}

	// Create a context that listens for OS interrupt signals (e.g., Ctrl+C)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Generate link cards for the provided URLs and save them to a file
	cards, err := linkcard.GenerateLinkCard(ctx, cfg, fs.Args())
	if err != nil {
		return err
	}

	// Output the generated link cards as JSON to the standard output
	enc := json.NewEncoder(ui.Writer())
	enc.SetIndent("", "  ")
	if err := enc.Encode(cards); err != nil {
		return errs.Wrap(err, errs.WithContext("output", "stdout"))
	}
	return nil
}
