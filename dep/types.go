package dep

import (
	"html/template"
)

type Logger struct {
	Error func(error)
	Ok func(string)
}

type Config struct {
	HostAddr string `json:"host_addr"`
	FileServer string `json:"file_server"`
	TmpDir string	`json:"tmp_dir"`
}

type Dependencies struct {
	Log *Logger
	Cfg *Config
	Tmp *template.Template
}