package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// be sure config doesn't import anything more than it requires in order to avoid an import cycle
// uses standard library only

// AppConfig holds the application config
type AppConfig struct {
	Port          string
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	Secure        bool
	SessionManager *scs.SessionManager
}
