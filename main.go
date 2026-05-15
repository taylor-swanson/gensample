// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/elastic/go-ucfg/yaml"

	"github.com/taylor-swanson/gensample/internal/runner"
	"github.com/taylor-swanson/gensample/internal/version"

	_ "github.com/taylor-swanson/gensample/internal/emitter/json"
	_ "github.com/taylor-swanson/gensample/internal/emitter/template"

	_ "github.com/taylor-swanson/gensample/internal/generator/date"
	_ "github.com/taylor-swanson/gensample/internal/generator/names"
	_ "github.com/taylor-swanson/gensample/internal/generator/net"
	_ "github.com/taylor-swanson/gensample/internal/generator/number"
	_ "github.com/taylor-swanson/gensample/internal/generator/strings"
	_ "github.com/taylor-swanson/gensample/internal/generator/uuid"
	_ "github.com/taylor-swanson/gensample/internal/generator/web"

	_ "github.com/taylor-swanson/gensample/internal/output/file"
	_ "github.com/taylor-swanson/gensample/internal/output/stdout"
	_ "github.com/taylor-swanson/gensample/internal/output/tcp"
	_ "github.com/taylor-swanson/gensample/internal/output/udp"
)

func usage() {
	_, _ = fmt.Fprintln(os.Stderr, version.Name, "[flags]", "CONFIG_FILE")
	_, _ = fmt.Fprintln(os.Stderr)
	_, _ = fmt.Fprintln(os.Stderr, "Args:")
	_, _ = fmt.Fprintln(os.Stderr, "  CONFIG_FILE    path to config file")
	_, _ = fmt.Fprintln(os.Stderr)
	_, _ = fmt.Fprintln(os.Stderr, "Flags:")
	flag.PrintDefaults()
}

func main() {
	seed := flag.Uint64("s", 0, "seed for random number generator (0 for current time)")
	showVersion := flag.Bool("v", false, "show version")
	debug := flag.Bool("d", false, "enable debug logging")

	flag.Usage = usage
	flag.Parse()

	if *showVersion {
		fmt.Printf("%s version %s [commit %v]\n", version.Name, version.Version, version.Commit)
		os.Exit(0)
	}
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(2)
	}
	configFile := flag.Arg(0)

	level := slog.LevelInfo
	if *debug {
		level = slog.LevelDebug
	}
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: level})))

	cfg, err := yaml.NewConfigWithFile(configFile)
	if err != nil {
		slog.Error("Failed to open config file", slog.String("filename", configFile), slog.String("error", err.Error()))
		os.Exit(1)
	}

	r, err := runner.New(cfg, *seed)
	if err != nil {
		slog.Error("Failed to create runner", slog.String("error", err.Error()))
		os.Exit(1)
	}

	if err = r.Run(); err != nil {
		slog.Error("Runner failed", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
