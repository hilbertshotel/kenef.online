package dep

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func Load() (*Dependencies, error) {
	log := initLogger()

	cfg, err := loadConfig("./config.json")
	if err != nil {
		return nil, err
	}

	tmp, err := initTemplates(cfg)
	if err != nil {
		return nil, err
	}

	return &Dependencies{
		Log: log,
		Cfg: cfg,
		Tmp: tmp,
	}, nil
}

// Logger
func initLogger() *Logger {
	okLog := log.New(os.Stdout, "ok ", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "error ", log.Ldate|log.Ltime)

	return &Logger{
		Error: func(err error) {
			_, file, line, _ := runtime.Caller(1)
			msg := fmt.Sprintf("(%v %v) %v", filepath.Base(file), line, err)
			errLog.Println(msg)
		},

		Ok: func(msg string) {
			okLog.Println(msg)
		},
	}
}

// Config
func loadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// Templates
func initTemplates(cfg *Config) (*template.Template, error) {
	tmp, err := template.ParseGlob(cfg.TmpDir)
	if err != nil {
		return nil, err
	}

	return tmp, nil
}
