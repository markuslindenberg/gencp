package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"

	"github.com/markuslindenberg/gencp/brandmeister"
	"github.com/markuslindenberg/gencp/codeplug"
)

func generateCodeplug(repeaters []string, tgs []string, mcc string, tgLimit int) (*codeplug.Codeplug, error) {
	cp := codeplug.NewCodeplug()

	groups, err := brandmeister.GetTalkgroups()
	if err != nil {
		return nil, err
	}

	isMn := func(r rune) bool {
		return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
	}
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)

	for id, name := range groups {
		groups[id], _, _ = transform.String(t, name)
	}

	contacts1 := make(codeplug.ContactSlice, 0)
	contacts2 := make(codeplug.ContactSlice, 0)
	contacts := make(map[string]*codeplug.Contact)
	for id, name := range groups {
		contact := codeplug.Contact{
			Name: name,
			ID:   id,
		}
		if strings.HasPrefix(id, mcc) {
			contacts1 = append(contacts1, &contact)
			contacts[id] = &contact
		} else {
			if len(id) <= tgLimit {
				contacts2 = append(contacts2, &contact)
				contacts[id] = &contact
			}
		}
	}
	sort.Sort(contacts1)
	sort.Sort(contacts2)
	cp.Contacts = append(cp.Contacts, contacts1...)
	cp.Contacts = append(cp.Contacts, contacts2...)

	for _, id := range repeaters {
		repeater, err := brandmeister.GetRepeater(id)
		if err != nil {
			return nil, err
		}
		callsign := strings.SplitN(repeater.Callsign, " ", 2)[0]

		profile, err := brandmeister.GetProfile(id)
		if err != nil {
			return nil, err
		}

		zone := codeplug.Zone{
			Name:     callsign + " " + repeater.City,
			Channels: []*codeplug.Channel{},
		}
		ts1GroupList := codeplug.ContactList{
			Name:     fmt.Sprint(callsign, " TS1"),
			Contacts: codeplug.ContactSlice{},
		}
		ts2GroupList := codeplug.ContactList{
			Name:     fmt.Sprint(callsign, " TS2"),
			Contacts: codeplug.ContactSlice{},
		}
		scanList := codeplug.ScanList{
			Name:     callsign,
			Channels: []*codeplug.Channel{},
		}
		channel := codeplug.Channel{
			RxFrequency: repeater.Tx,
			TxFrequency: repeater.Rx,
			ColorCode:   repeater.Colorcode,
		}

		ts1Channels := []*codeplug.Channel{}
		ts2Channels := []*codeplug.Channel{}

		c := channel
		c.Contact = contacts["9"]
		if c.Contact == nil {
			contact := codeplug.Contact{
				Name: "TG9",
				ID:   "9",
			}
			cp.Contacts = append(cp.Contacts, &contact)
			contacts["9"] = &contact
			c.Contact = &contact
		}
		c.Name = "Local"
		c.Repeater = callsign
		c.Slot = 2
		c.GroupList = &ts2GroupList
		c.ScanList = &scanList
		ts2Channels = append(ts2Channels, &c)
		ts2GroupList.Contacts = append(ts2GroupList.Contacts, c.Contact)
		scanList.Channels = append(scanList.Channels, &c)

		for _, s := range profile.Clusters {
			tg := fmt.Sprint(s.Talkgroup)
			c := channel
			c.Contact = contacts[tg]
			if c.Contact == nil {
				contact := codeplug.Contact{
					Name: fmt.Sprint("TG", tg),
					ID:   tg,
				}
				cp.Contacts = append(cp.Contacts, &contact)
				contacts[tg] = &contact
				c.Contact = &contact
			}
			c.Name = groups[fmt.Sprint(s.ExtTalkgroup)]
			if c.Name == "" {
				c.Name = fmt.Sprint("TG", tg)
			}
			c.Repeater = callsign
			c.ScanList = &scanList
			if s.Slot == 1 {
				c.Slot = 1
				c.GroupList = &ts1GroupList
				ts1Channels = append(ts1Channels, &c)
				ts1GroupList.Contacts = append(ts1GroupList.Contacts, c.Contact)
			} else if s.Slot == 2 {
				c.Slot = 2
				c.GroupList = &ts2GroupList
				ts2Channels = append(ts2Channels, &c)
				ts2GroupList.Contacts = append(ts2GroupList.Contacts, c.Contact)
			}
			scanList.Channels = append(scanList.Channels, &c)
		}
		for _, s := range profile.StaticSubscriptions {
			tg := fmt.Sprint(s.Talkgroup)
			c := channel
			c.Contact = contacts[tg]
			if c.Contact == nil {
				contact := codeplug.Contact{
					Name: fmt.Sprint("TG", tg),
					ID:   tg,
				}
				cp.Contacts = append(cp.Contacts, &contact)
				contacts[tg] = &contact
				c.Contact = &contact
			}
			c.Name = groups[tg]
			c.Repeater = callsign
			if c.Name == "" {
				c.Name = fmt.Sprint("TG", tg)
			}
			c.ScanList = &scanList
			if s.Slot == 1 {
				c.Slot = 1
				c.GroupList = &ts1GroupList
				ts1Channels = append(ts1Channels, &c)
				ts1GroupList.Contacts = append(ts1GroupList.Contacts, c.Contact)
			} else if s.Slot == 2 {
				c.Slot = 2
				c.GroupList = &ts2GroupList
				ts2Channels = append(ts2Channels, &c)
				ts2GroupList.Contacts = append(ts2GroupList.Contacts, c.Contact)
			}
			scanList.Channels = append(scanList.Channels, &c)
		}
		for _, s := range profile.TimedSubscriptions {
			tg := fmt.Sprint(s.Talkgroup)
			c := channel
			c.Contact = contacts[tg]
			if c.Contact == nil {
				contact := codeplug.Contact{
					Name: fmt.Sprint("TG", tg),
					ID:   tg,
				}
				cp.Contacts = append(cp.Contacts, &contact)
				contacts[tg] = &contact
				c.Contact = &contact
			}
			c.Name = groups[tg]
			if c.Name == "" {
				c.Name = fmt.Sprint("TG", tg)
			}
			c.Repeater = callsign
			c.ScanList = &scanList
			if s.Slot == 1 {
				c.Slot = 1
				c.GroupList = &ts1GroupList
				ts1Channels = append(ts1Channels, &c)
				ts1GroupList.Contacts = append(ts1GroupList.Contacts, c.Contact)
			} else if s.Slot == 2 {
				c.Slot = 2
				c.GroupList = &ts2GroupList
				ts2Channels = append(ts2Channels, &c)
				ts2GroupList.Contacts = append(ts2GroupList.Contacts, c.Contact)
			}
			scanList.Channels = append(scanList.Channels, &c)
		}

		for _, tg := range tgs {
			c := channel
			c.Contact = contacts[tg]
			if c.Contact == nil {
				contact := codeplug.Contact{
					Name: fmt.Sprint("TG", tg),
					ID:   tg,
				}
				cp.Contacts = append(cp.Contacts, &contact)
				contacts[tg] = &contact
				c.Contact = &contact
			}
			c.Name = groups[tg]
			if c.Name == "" {
				c.Name = fmt.Sprint("TG", tg)
			}
			c.Repeater = callsign
			c.ScanList = &scanList
			c.Slot = 2
			c.GroupList = &ts2GroupList
			ts2Channels = append(ts2Channels, &c)
			ts2GroupList.Contacts = append(ts2GroupList.Contacts, c.Contact)
			scanList.Channels = append(scanList.Channels, &c)
		}

		cp.GroupLists = append(cp.GroupLists, &ts1GroupList)
		cp.GroupLists = append(cp.GroupLists, &ts2GroupList)
		cp.Channels = append(cp.Channels, ts2Channels...)
		zone.Channels = append(zone.Channels, ts2Channels...)
		cp.Channels = append(cp.Channels, ts1Channels...)
		zone.Channels = append(zone.Channels, ts1Channels...)
		cp.ScanLists = append(cp.ScanLists, &scanList)
		cp.Zones = append(cp.Zones, &zone)
	}

	return cp, nil
}
