package runner

import (
	"fmt"
	"log/slog"
	"math/rand/v2"
	"time"

	"github.com/elastic/go-ucfg"

	"github.com/taylor-swanson/gensample/internal/context"
	"github.com/taylor-swanson/gensample/internal/emitter"
	"github.com/taylor-swanson/gensample/internal/field"
	"github.com/taylor-swanson/gensample/internal/output"
)

type config struct {
	Emitter  *ucfg.Config   `config:"emitter" validate:"required"`
	Output   *ucfg.Config   `config:"output" validate:"required"`
	Fields   []*ucfg.Config `config:"fields" validate:"required"`
	Records  int            `config:"records"`
	Limit    int            `config:"limit"`
	Interval time.Duration  `config:"interval"`
}

func defaultConfig() config {
	return config{
		Records: 1,
	}
}

type Runner struct {
	rng     rand.Rand
	emitter emitter.Emitter
	output  output.Output
	fields  []*field.Field
	seed    uint64

	config config
}

func (r *Runner) Run() error {
	var ticker *time.Ticker
	if r.config.Interval > 0 {
		ticker = time.NewTicker(r.config.Interval)
		defer ticker.Stop()
	}

	var err error
	var iteration uint64
	var total int
	for {
		ctx := context.Context{
			Total: r.config.Records,
			Rand:  rand.New(rand.NewPCG(r.seed, r.seed^iteration)),
		}

		if r.config.Limit > 0 && total >= r.config.Limit {
			return r.output.Close()
		}

		for i := 0; i < r.config.Records; i++ {
			ctx.Current = i
			ctx.Progress = float64(ctx.Current) / float64(ctx.Total-1)

			if r.config.Limit > 0 && total >= r.config.Limit {
				slog.Info("Limit reached", "interval", iteration, "records", i)
				return r.output.Close()
			}
			if _, err = r.output.Write([]byte(r.emitter.Emit(&ctx))); err != nil {
				_ = r.output.Close()
				return fmt.Errorf("runner: unable to emit to output: %w", err)
			}

			total++
		}
		slog.Info("Interval completed", "interval", iteration, "records", r.config.Records)

		if r.config.Interval == 0 {
			return r.output.Close()
		}
		if err = r.output.NewInterval(); err != nil {
			_ = r.output.Close()
			return fmt.Errorf("runner: unable to create new interval: %w", err)
		}
		if ticker != nil {
			<-ticker.C
		}

		iteration++
	}
}

func New(cfg *ucfg.Config, seed uint64) (Runner, error) {
	var err error

	if seed == 0 {
		seed = uint64(time.Now().UnixNano())
	}

	slog.Info("Starting runner", "seed", seed)

	r := Runner{
		seed: seed,
	}
	c := defaultConfig()
	if err = cfg.Unpack(&c); err != nil {
		return r, fmt.Errorf("runner: failed to unpack config: %w", err)
	}

	r.config = c

	var fields []*field.Field
	for _, v := range c.Fields {
		f, err := field.New(v)
		if err != nil {
			return r, fmt.Errorf("runner: failed to create field: %w", err)
		}
		fields = append(fields, f)
	}

	if r.emitter, err = emitter.New(c.Emitter, fields); err != nil {
		return r, fmt.Errorf("runner: failed to create emitter: %w", err)
	}

	if r.output, err = output.New(c.Output); err != nil {
		return r, fmt.Errorf("runner: failed to create output: %w", err)
	}

	return r, nil
}
