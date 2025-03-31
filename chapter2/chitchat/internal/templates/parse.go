package templates

import (
	"errors"
	"html/template"
	"io"

	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/concurrent"
)

type TemplateConfig struct {
	Writer io.Writer
	Name   string
	Data   any
	Path   []string
}

type ParseTemplate struct {
	writer io.Writer
	name   string
	data   any
	path   []string
}

func NewParseTemplate(cfg TemplateConfig) *ParseTemplate {
	return &ParseTemplate{
		writer: cfg.Writer,
		name:   cfg.Name,
		data:   cfg.Data,
		path:   cfg.Path,
	}
}

func (p *ParseTemplate) ParseAndRender() error {
	files, err := concurrent.PathFile(p.path...)
	if err != nil {
		return err
	}

	if len(files) == 0 {
		return errors.New("no template files found")
	}

	templates, err := template.ParseFiles(files...)
	if err != nil {
		return err
	}

	if err = templates.ExecuteTemplate(p.writer, p.name, p.data); err != nil {
		return err
	}

	return nil
}
