package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"unicode"

	"github.com/markuslindenberg/gencp/model"
	_ "github.com/markuslindenberg/gencp/model/md380"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	var (
		dmridFlag     = flag.String("d", "0", "DMR ID")
		callsignFlag  = flag.String("c", "XX0XX", "callsign")
		mccFlag       = flag.String("mcc", "262", "MCC (country code)")
		tgLimitFlag   = flag.Uint("tglimit", 4, "Minimum ID length for TGs in other countries")
		repeatersFlag = flag.String("r", "262440,262477,262412,262436,262421", "repeater ID(s)")
		tgsFlag       = flag.String("t", "26223", "extra talkgroup(s) in all zones")
		modelFlag     = flag.String("m", "md380", "radio model")
		formatFlag    = flag.String("o", "json", "output format")
		listFlag      = flag.Bool("l", false, "list supported models and formats")
		// flashFlag  = flag.Bool("f", false, "flash codeplug to radio")
		serveFlag = flag.Bool("s", false, "run as webserver")
		// portFlag      = flag.String("p", "localhost:8080", "listen on host:port")
	)
	flag.Parse()

	if *listFlag {
		names := model.List()
		sort.Strings(names)
		for _, name := range names {
			m, err := model.Get(name)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%#v\n", m.GetNames())
			fmt.Printf("%#v\n", m.GetFormats())
		}
		return
	}

	if *serveFlag {
		log.Fatal("Not yet implemented")
	}

	m, err := model.Get(*modelFlag)
	if err != nil {
		log.Fatal(err)
	}

	formatFound := false
	for _, f := range m.GetFormats() {
		if f.ID == *formatFlag {
			formatFound = true
			break
		}
	}
	if !formatFound {
		log.Fatal("Unsupported output format")
	}

	for _, c := range *dmridFlag {
		if !unicode.IsDigit(c) {
			log.Fatal("DMR ID must be numeric")
		}
	}

	for _, c := range *mccFlag {
		if !unicode.IsDigit(c) {
			log.Fatal("MCC must be numeric")
		}
	}

	repeaters := strings.Split(*repeatersFlag, ",")
	for _, repeater := range repeaters {
		if repeater == "" {
			log.Fatal("Repeater ID cannot be empty")
		}
		for _, c := range repeater {
			if !unicode.IsDigit(c) {
				log.Fatal("Repeater ID must be numeric")
			}
		}
	}

	tgs := strings.Split(*tgsFlag, ",")
	for _, tg := range tgs {
		for _, c := range tg {
			if !unicode.IsDigit(c) {
				log.Fatal("Talkgroup ID must be numeric")
			}
		}
	}

	cp, err := generateCodeplug(repeaters, tgs, *mccFlag, int(*tgLimitFlag))
	if err != nil {
		log.Fatal(err)
	}

	data, err := m.Generate(*modelFlag, *formatFlag, *dmridFlag, *callsignFlag, cp)
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(data)
}
