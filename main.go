package main

import (
	"flag"
	"fmt"
	"log"
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
		dmridFlag     = flag.Uint("d", 0, "DMR ID")
		callsignFlag  = flag.String("c", "XX0XX", "callsign")
		repeatersFlag = flag.String("r", "262440,262477,262412,262436,262421", "repeater ID(s)")
		tgsFlag       = flag.String("t", "26223", "extra talkgroup(s) in all zones")
		modelFlag     = flag.String("m", "md380", "radio model")
		formatFlag    = flag.String("o", "rdt", "output format")
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

	repeaters := strings.Split(*repeatersFlag, ",")
	for _, repeater := range repeaters {
		for _, c := range repeater {
			if !unicode.IsDigit(c) {
				log.Fatal("Invalid character in repeater ID")
			}
		}
	}

	tgs := strings.Split(*tgsFlag, ",")
	for _, tg := range tgs {
		for _, c := range tg {
			if !unicode.IsDigit(c) {
				log.Fatal("Invalid character in talkgroup ID")
			}
		}
	}

	cp, err := generateCodeplug(repeaters, tgs)
	if err != nil {
		log.Fatal(err)
	}

	_, err = m.Generate(*modelFlag, *formatFlag, *dmridFlag, *callsignFlag, cp)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Success to /dev/null")
}
