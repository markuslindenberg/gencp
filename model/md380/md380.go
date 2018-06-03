package md380

import (
	"errors"
	"fmt"
	"strings"

	"gopkg.in/square/go-jose.v2/json"

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

func truncate(input string, length int) string {
	result := input
	chars := 0
	for i := range input {
		if chars >= length {
			result = input[:i]
			break
		}
		chars++
	}
	return result
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
				ID:          "json",
				Description: "JSON",
				Extension:   "json",
				Mimetype:    "application/json",
				Preferred:   true,
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

func (m *md380) Generate(name string, format string, dmrid string, callsign string, codeplug *codeplug.Codeplug) ([]byte, error) {
	// Truncate all names to 16 characters
	for _, c := range codeplug.Channels {
		maxLen := 16 - (len(c.Repeater) + 1)
		c.Name = truncate(c.Name, maxLen) + " " + c.Repeater
	}
	for _, c := range codeplug.Contacts {
		c.Name = truncate(c.Name, 16)
	}
	for _, l := range codeplug.GroupLists {
		l.Name = truncate(l.Name, 16)
	}
	for _, l := range codeplug.ScanLists {
		l.Name = truncate(l.Name, 16)
	}
	for _, l := range codeplug.Zones {
		l.Name = truncate(l.Name, 16)
	}

	md380cp := codeplugTemplate
	md380cp.Contacts = []Contact{
		Contact{
			CallID:          "262999",
			CallReceiveTone: "No",
			CallType:        "Private",
			Name:            "BM-GPS",
		},
	}
	md380cp.Channels = []Channel{}
	md380cp.GroupLists = []GroupList{}
	md380cp.ScanLists = []ScanList{}
	md380cp.Zones = []Zone{}

	md380cp.GeneralSettings.IntroScreenLine1 = callsign
	md380cp.GeneralSettings.IntroScreenLine2 = "BM"
	md380cp.GeneralSettings.RadioName = callsign
	md380cp.GeneralSettings.RadioID = dmrid

	for _, c := range codeplug.Channels {
		channel := channelTemplate
		channel.ColorCode = fmt.Sprint(c.ColorCode)
		channel.ContactName = c.Contact.Name
		channel.GroupList = c.GroupList.Name
		channel.Name = c.Name
		channel.RepeaterSlot = fmt.Sprint(c.Slot)
		channel.RxFrequency = c.RxFrequency
		channel.TxFrequency = c.TxFrequency
		channel.ScanList = c.ScanList.Name
		if strings.HasPrefix(c.Contact.ID, "91") {
			channel.RxOnly = "On"
		}
		md380cp.Channels = append(md380cp.Channels, channel)
	}

	for _, c := range codeplug.Contacts {
		contact := Contact{
			CallID:          c.ID,
			CallReceiveTone: "No",
			CallType:        "Group",
			Name:            c.Name,
		}
		md380cp.Contacts = append(md380cp.Contacts, contact)
	}

	for _, g := range codeplug.GroupLists {
		grouplist := GroupList{
			Contacts: make([]string, len(g.Contacts)),
			Name:     g.Name,
		}
		for i, c := range g.Contacts {
			grouplist.Contacts[i] = c.Name
		}
		md380cp.GroupLists = append(md380cp.GroupLists, grouplist)
	}

	for _, z := range codeplug.Zones {
		zone := Zone{
			Channels: make([]string, len(z.Channels)),
			Name:     z.Name,
		}
		for i, c := range z.Channels {
			zone.Channels[i] = c.Name
		}
		md380cp.Zones = append(md380cp.Zones, zone)
	}

	for _, s := range codeplug.ScanLists {
		scanlist := scanlistTemplate
		scanlist.Channels = make([]string, len(s.Channels))
		scanlist.Name = s.Name
		for i, c := range s.Channels {
			scanlist.Channels[i] = c.Name
		}
		md380cp.ScanLists = append(md380cp.ScanLists, scanlist)
	}

	data, err := json.MarshalIndent(md380cp, "", "\t")
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (m *md380) Flash(name string, dmrid string, callsign string, codeplug *codeplug.Codeplug) error {
	return errors.New("not implemented")
}
