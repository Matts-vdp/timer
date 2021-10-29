package save

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type saveform struct {
	Start    time.Time
	Duration string
	Info     string
}

func (sav saveform) String() string {
	str := sav.Start.Format("02/01/2006 15:04")
	str += " : " + sav.Duration
	str += " : " + sav.Info + "\n"
	return str
}

type saveStore struct {
	Saves []saveform
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

func Load() saveStore {
	f, err := ioutil.ReadFile("C:/Users/Gebruiker/Desktop/python/Go/timer/save.json")
	if err != nil {
		log.Fatal(err)
	}
	s := saveStore{}
	json.Unmarshal(f, &s)
	return s
}

func Save(start time.Time, info string) {
	sf := saveform{start, Print(time.Since(start)), info}
	s := Load()
	s.Saves = append(s.Saves, sf)
	js, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile("C:/Users/Gebruiker/Desktop/python/Go/timer/save.json", js, 0644)
}
