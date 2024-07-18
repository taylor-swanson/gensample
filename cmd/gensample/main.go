package main

import (
	"flag"
	"log"

	"github.com/elastic/go-ucfg/yaml"

	"github.com/taylor-swanson/gensample/internal/runner"
)

var (
	seed       uint64
	configFile string
)

func main() {
	flag.Uint64Var(&seed, "s", 0, "seed for random number generator (default is current time)")
	flag.StringVar(&configFile, "c", "config.yml", "path to config file")

	flag.Parse()

	cfg, err := yaml.NewConfigWithFile(configFile)
	if err != nil {
		panic(err)
	}

	r, err := runner.New(cfg, seed)
	if err != nil {
		log.Fatalf("Failed to create runner: %v", err)
	}

	if err = r.Run(); err != nil {
		log.Fatalf("Runner failed: %v", err)
	}
}
