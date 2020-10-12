package services

import (
	"bytes"
	"log"
	"text/template"
)

// TemplatingServiceGolang implements TemplatingService
type TemplatingServiceGolang struct {
}

// ResolveTemplate combines template with data into string
func (service *TemplatingServiceGolang) ResolveTemplate(templateStr string, payload map[string]interface{}) (string, error) {
	log.Printf("TemplatingServiceGolang.ResolveTemplate '%s' with payload: %+v", templateStr, payload)
	tmpl, err := template.New("TemplatingServiceGolangTemplate").Parse(templateStr)
	if err != nil {
		return "", err
	}

	var out bytes.Buffer
	if err := tmpl.Execute(&out, payload); err != nil {
		return "", err
	}
	return out.String(), nil
}
