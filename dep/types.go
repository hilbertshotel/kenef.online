package dep

import (
	"html/template"
)

type Logger struct {
	Error func(error)
	Ok func(string)
}

type Config struct {
	HostAddr string `json:"hostAddr"`
	TmpDir string	`json:"tmpDir"`
}

type Dependencies struct {
	Log *Logger
	Cfg *Config
	Tmp *template.Template
}