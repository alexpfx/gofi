package dmenu

import (
	"github.com/alexpfx/gofi"
)

func New(config Config) DMenu {
	return &dmenu{
		config: config,
	}
}

type DMenu interface {
	Build() []string
}

type Config struct {
	Sep             string
	Prompt          string
	Lines           int
	InsensitiveCase bool
	Active          string
	Urgent          string
	OnlyMatch       bool
	NoCustom        bool
}

type dmenu struct {
	config Config
}

func (d *dmenu) Build() []string {
	return append([]string{"-dmenu"}, buildParams(d.config)...)
}

func buildParams(cfg Config) []string {
	argArray := make([]string, 0)

	argArray = gofi.Add(argArray, "-sep", cfg.Sep)
	argArray = gofi.Add(argArray, "-p", cfg.Prompt)
	argArray = gofi.Add(argArray, "-l", cfg.Lines)
	argArray = gofi.Add(argArray, "-i", cfg.InsensitiveCase)
	argArray = gofi.Add(argArray, "-a", cfg.Active)
	argArray = gofi.Add(argArray, "-u", cfg.Urgent)
	argArray = gofi.Add(argArray, "-only-match", cfg.OnlyMatch)
	argArray = gofi.Add(argArray, "-no-custom", cfg.NoCustom)
	return argArray

}
