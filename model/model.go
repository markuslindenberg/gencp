package model

import (
	"errors"

	"github.com/markuslindenberg/gencp/codeplug"
)

var (
	models = map[string]Model{}
)

// Name is a name/brand supported by this model
type Name struct {
	ID    string
	Brand string
	Model string
}

// Format is a output format for the generated codeplug
type Format struct {
	ID          string
	Description string
	Extension   string
	Mimetype    string
	Preferred   bool
}

// Model is a supported DMR radio
type Model interface {
	GetID() string
	GetNames() []*Name
	GetFormats() []*Format
	Generate(name string, format string, dmrid uint, callsign string, codeplug *codeplug.Codeplug) (data []byte, err error)
	Flash(name string, dmrid uint, callsign string, codeplug *codeplug.Codeplug) error
}

// Register makes a radio model available to gencp.
// Register is not goroutine safe.
func Register(model Model) {
	models[model.GetID()] = model
}

// List lists all registered models
func List() []string {
	names := make([]string, 0, len(models))
	for name := range models {
		names = append(names, name)
	}
	return names
}

// Get returns all available models
func Get(name string) (Model, error) {
	model := models[name]
	if model == nil {
		return nil, errors.New("Model not found")
	}
	return model, nil
}
