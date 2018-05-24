package md380

import (
	"errors"

	// ToDo: implement
	_ "github.com/DaleFarnsworth/codeplug/codeplug"
	"github.com/markuslindenberg/gencp/codeplug"
	"github.com/markuslindenberg/gencp/model"
)

func init() {
	model.Register(newMD380())
}

type md380 struct {
	names   []*model.Name
	formats []*model.Format
}

func newMD380() *md380 {
	return &md380{
		names: []*model.Name{
			&model.Name{
				ID:    "md380",
				Brand: "TYT",
				Model: "MD380",
			},
			&model.Name{
				ID:    "md390",
				Brand: "TYT",
				Model: "MD390",
			},
			&model.Name{
				ID:    "rt3",
				Brand: "Retevis",
				Model: "RT-3",
			},
		},
		formats: []*model.Format{
			&model.Format{
				ID:          "rdt",
				Description: ".rdt file for TYT CPS",
				Extension:   "rdt",
				Mimetype:    "application/octet-stream",
				Preferred:   true,
			},
			&model.Format{
				ID:          "json",
				Description: "JSON",
				Extension:   "json",
				Mimetype:    "application/json",
				Preferred:   false,
			},
			&model.Format{
				ID:          "xlsx",
				Description: "XLSX spreadsheet",
				Extension:   "xlsx",
				Mimetype:    "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
				Preferred:   false,
			},
			&model.Format{
				ID:          "text",
				Description: "Text file",
				Extension:   "txt",
				Mimetype:    "text/plain",
				Preferred:   false,
			},
		},
	}
}

func (m *md380) GetID() string {
	return "md380"
}

func (m *md380) GetNames() []*model.Name {
	return m.names
}

func (m *md380) GetFormats() []*model.Format {
	return m.formats
}

func (m *md380) Generate(name string, format string, dmrid uint, callsign string, codeplug *codeplug.Codeplug) (data []byte, err error) {
	return nil, errors.New("not implemented")
}

func (m *md380) Flash(name string, dmrid uint, callsign string, codeplug *codeplug.Codeplug) error {
	return errors.New("not implemented")
}
