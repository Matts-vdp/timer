package save

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

const FILE = "save.json"

// used to store 1 item
type saveform struct {
	Start    time.Time
	Duration time.Duration
	Info     string
}

func (sav saveform) String() string {
	str := sav.Start.Format("02/01/2006")
	str += " : " + Print(sav.Duration)
	str += " : " + sav.Info + "\n"
	return str
}

// all saved items
type saveStore struct {
	Saves []saveform
}

// adds new item
// the new duration gets added to the previous if the day is the same
func (sav *saveStore) add(sf saveform) {
	for i := range sav.Saves {
		if sav.Saves[i].Start.YearDay() == sf.Start.YearDay() && sav.Saves[i].Start.Year() == sf.Start.Year() {
			sav.Saves[i].Duration += sf.Duration
			sav.Saves[i].Info += " & " + sf.Info
			return
		}
	}
	sav.Saves = append(sav.Saves, sf)
}

func (sav saveStore) String() string {
	var str string
	for _, s := range sav.Saves {
		str += s.String()
	}
	return str
}

func Print(t time.Duration) string {
	hr := int(t.Hours())
	mn := int(t.Minutes()) - int(t.Hours())*60
	sc := int(t.Seconds()) - int(t.Minutes())*60
	return fmt.Sprintf("%.2d:%.2d:%.2d", hr, mn, sc)
}

// reads the saved items from a file
func Load() saveStore {
	f, err := ioutil.ReadFile(FILE)
	if err != nil {
		log.Fatal(err)
	}
	s := saveStore{}
	json.Unmarshal(f, &s)
	return s
}

// saves the items to a file
func Save(start time.Time, info string) {
	sf := saveform{start, time.Since(start), info}
	s := Load()
	s.add(sf)
	js, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile(FILE, js, 0644)
}
