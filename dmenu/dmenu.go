package dmenu

import (
	"fmt"
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

	argArray = add(argArray, "-sep", cfg.Sep)
	argArray = add(argArray, "-p", cfg.Prompt)
	argArray = add(argArray, "-l", cfg.Lines)
	argArray = add(argArray, "-i", cfg.InsensitiveCase)
	argArray = add(argArray, "-a", cfg.Active)
	argArray = add(argArray, "-u", cfg.Urgent)
	argArray = add(argArray, "-only-match", cfg.OnlyMatch)
	argArray = add(argArray, "-no-custom", cfg.NoCustom)
	return argArray

}

func add(argList []string, arg string, value interface{}) []string {
	bp, ok := value.(bool)
	if ok {
		if bp {
			return append(argList, arg)
		}
		return argList
	}

	st, ok := value.(string)
	argList = append(argList, arg)
	if ok && st != "" {
		argList = append(argList, fmt.Sprintf("%s", st))
		return argList
	}

	argList = append(argList, fmt.Sprintf("%v", value))

	return argList
}
