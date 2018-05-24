package codeplug

import "strings"

// Contact is a DMR contact
type Contact struct {
	Name        string
	ID          string
	PrivateCall bool
}

type ContactSlice []*Contact

func (c ContactSlice) Len() int           { return len(c) }
func (c ContactSlice) Less(i, j int) bool { return strings.Compare(c[i].ID, c[j].ID) == -1 }
func (c ContactSlice) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }

// Channel is a DMR channel
type Channel struct {
	Name        string
	RxFrequency uint64
	TxFrequency uint64
	Slot        uint8
	ColorCode   uint8
	Contact     *Contact
	GroupList   *ContactList
	ScanList    *ScanList
}

// ContactList is a named list of contacts
type ContactList struct {
	Name     string
	Contacts ContactSlice
}

// ChannelList is a named list of channels
type ChannelList struct {
	Name     string
	Channels []*Channel
}

// ScanList is a list of channels to scan
type ScanList ChannelList

// Zone is a grouped list of channels
type Zone ChannelList

// Codeplug contains the channel/zone/contact configuration
type Codeplug struct {
	Contacts   ContactSlice
	Channels   []*Channel
	GroupLists []*ContactList
	ScanLists  []*ScanList
	Zones      []*Zone
}

// NewCodeplug returns an allocated Codeplug
func NewCodeplug() *Codeplug {
	return &Codeplug{
		Contacts:   ContactSlice{},
		Channels:   []*Channel{},
		GroupLists: []*ContactList{},
		ScanLists:  []*ScanList{},
		Zones:      []*Zone{},
	}
}
