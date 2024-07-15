package main

import (
	"flag"
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/elastic/go-ucfg"
	"github.com/elastic/go-ucfg/yaml"

	"gensample/internal/context"
	"gensample/internal/emitter"
)

var (
	count      int
	seed       uint64
	configFile string
)

type config struct {
	Emitter *ucfg.Config `config:"emitter" validate:"required"`
}

func main() {
	flag.IntVar(&count, "n", 1, "number of logs to generate")
	flag.Uint64Var(&seed, "s", 0, "seed for random number generator (zero for random seed)")
	flag.StringVar(&configFile, "c", "config.yml", "path to config file")

	flag.Parse()

	var c config

	cfg, err := yaml.NewConfigWithFile(configFile)
	if err != nil {
		panic(err)
	}
	if err = cfg.Unpack(&c); err != nil {
		panic(err)
	}

	em, err := emitter.New(c.Emitter)
	if err != nil {
		panic(err)
	}

	ctx := context.Context{
		Total: count,
	}

	if seed == 0 {
		seed = uint64(time.Now().UnixNano())
	}
	ctx.Rand = rand.New(rand.NewPCG(seed, seed))

	for i := 0; i < ctx.Total; i++ {
		ctx.Current = i
		ctx.Progress = float64(ctx.Current) / float64(ctx.Total-1)

		fmt.Println(em.Emit(&ctx))
	}
}
