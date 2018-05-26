# gencp - Generate Brandmeister codeplug for MD380 DMR radios

**Warning: This is experimental stuff, use responsibly.**

This tool generates a codeplug for MD380 (and RT3 etc.) DMR radios. 

The codeplug is created using repeater and talkgroup data from the Brandmeister API.

* All talkgroups from the given country (-mcc) plus all talkgroups with IDs up to 4 digits (-tglimit) length will be added to the contacts list.
* A zone will be created for every given repeater containing all subscribed talkgroups as channels.
* All additional talkgroups (-t) will be created as channel in every zone.


## Installation

See [Go's installation instructions](https://golang.org/doc/install), then use `go get`:

```bash
go get github.com/markuslindenberg/gencp
```

## Usage

This example creates a codeplug for my MD-380 containing contacts for DL (262) and channels for some repeaters near my QTH.
**Please use your own callsign and DMR id!**

```bash
./gencp -c DO8ML -mcc 262 -d 2644025 -r 262440,262477,262421,262490,262411,262623 -t 26223 > codeplug.json
```

The JSON codeplug file can be imported and programmed using the excellent [Editcp codeplug editor](https://www.farnsworth.org/dale/codeplug/editcp/).
