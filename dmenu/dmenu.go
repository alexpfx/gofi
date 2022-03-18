package dmenu

import (
	"fmt"
	"strings"
)

func New(config Config) DMenu {
	return &dmenu{
		config: config,
	}
}

type DMenu interface {
	Build() string
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

func (d *dmenu) Build() string {
	return fmt.Sprint("-dmenu ", buildParams(d.config))
}

func buildParams(cfg Config) string {
	sb := &strings.Builder{}

	sb = add(sb, "-sep", cfg.Sep)
	sb = add(sb, "-p", cfg.Prompt)
	sb = add(sb, "-l", cfg.Lines)
	sb = add(sb, "-i", cfg.InsensitiveCase)
	sb = add(sb, "-a", cfg.Active)
	sb = add(sb, "-u", cfg.Urgent)
	sb = add(sb, "-only-match", cfg.OnlyMatch)
	sb = add(sb, "-no-custom", cfg.NoCustom)

	return sb.String()
}

func add(sb *strings.Builder, arg string, value interface{}) *strings.Builder {
	if value == nil {
		return sb
	}
	_, ok := value.(bool)
	if ok {
		sb.WriteString(arg + " ")
		return sb
	}
	_, ok = value.(string)
	if ok {
		sb.WriteString(fmt.Sprintf("%s '%s' ", arg, value))
		return sb
	}
	sb.WriteString(fmt.Sprintf("%s %v ", arg, value))
	return sb
}
