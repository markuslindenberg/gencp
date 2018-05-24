package main

import (
	"errors"
	"fmt"
	"sort"

	"github.com/markuslindenberg/gencp/brandmeister"
	"github.com/markuslindenberg/gencp/codeplug"
)

func generateCodeplug(repeaters []string, tgs []string) (*codeplug.Codeplug, error) {
	cp := codeplug.NewCodeplug()

	groups, err := brandmeister.GetTalkgroups()
	if err != nil {
		return nil, err
	}

	contacts := make(map[string]*codeplug.Contact)
	for id, name := range groups {
		contact := codeplug.Contact{
			Name: name,
			ID:   id,
		}
		cp.Contacts = append(cp.Contacts, &contact)
		contacts[id] = &contact
	}
	sort.Sort(cp.Contacts)

	for _, id := range repeaters {
		repeater, err := brandmeister.GetRepeater(id)
		if err != nil {
			return nil, err
		}
		profile, err := brandmeister.GetProfile(id)
		if err != nil {
			return nil, err
		}

		// Group lists
		zone := codeplug.Zone{
			Name:     repeater.Callsign,
			Channels: []*codeplug.Channel{},
		}
		ts1GroupList := codeplug.ContactList{
			Name:     fmt.Sprint(repeater.Callsign, " TS1"),
			Contacts: codeplug.ContactSlice{},
		}
		ts2GroupList := codeplug.ContactList{
			Name:     fmt.Sprint(repeater.Callsign, " TS2"),
			Contacts: codeplug.ContactSlice{},
		}
		channel := codeplug.Channel{
			RxFrequency: 0,
			TxFrequency: 0,
			ColorCode:   repeater.Colorcode,
		}

		ts1Channels := []*codeplug.Channel{}
		ts2Channels := []*codeplug.Channel{}

		c := channel
		c.Contact = contacts["9"]
		if c.Contact == nil {
			return nil, errors.New("Talkgroup 9 not found")
		}
		c.Name = groups["9"]
		if c.Name == "" {
			c.Name = "TG9"
		}
		c.Slot = 2
		c.GroupList = &ts2GroupList
		ts2Channels = append(ts2Channels, &c)
		ts2GroupList.Contacts = append(ts2GroupList.Contacts, c.Contact)

		for _, s := range profile.StaticSubscriptions {
			tg := fmt.Sprint(s.Talkgroup)
			c := channel
			c.Contact = contacts[tg]
			if c.Contact == nil {
				return nil, fmt.Errorf("Talkgroup %s not found", tg)
			}
			c.Name = groups[tg]
			if c.Name == "" {
				c.Name = fmt.Sprint("TG", tg)
			}
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
		}
		for _, s := range profile.Clusters {
			tg := fmt.Sprint(s.Talkgroup)
			c := channel
			c.Contact = contacts[tg]
			if c.Contact == nil {
				return nil, fmt.Errorf("Talkgroup %s not found", tg)
			}
			c.Name = groups[fmt.Sprint(s.ExtTalkgroup)]
			if c.Name == "" {
				c.Name = fmt.Sprint("TG", tg)
			}
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
		}
		for _, s := range profile.TimedSubscriptions {
			tg := fmt.Sprint(s.Talkgroup)
			c := channel
			c.Contact = contacts[tg]
			if c.Contact == nil {
				return nil, fmt.Errorf("Talkgroup %s not found", tg)
			}
			c.Name = groups[tg]
			if c.Name == "" {
				c.Name = fmt.Sprint("TG", tg)
			}
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
		}

		cp.GroupLists = append(cp.GroupLists, &ts1GroupList)
		cp.GroupLists = append(cp.GroupLists, &ts2GroupList)
		cp.Channels = append(cp.Channels, ts2Channels...)
		zone.Channels = append(zone.Channels, ts2Channels...)
		cp.Channels = append(cp.Channels, ts1Channels...)
		zone.Channels = append(zone.Channels, ts1Channels...)

		cp.Zones = append(cp.Zones, &zone)
	}

	return cp, nil
}
