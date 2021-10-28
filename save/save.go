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
}
type saveStore struct {
	Saves []saveform
}

func Print(t time.Duration) string {
	hr := int(t.Hours())
	mn := int(t.Minutes()) - int(t.Hours())*60
	sc := int(t.Seconds()) - int(t.Minutes())*60
	return fmt.Sprintf("%.2d:%.2d:%.2d", hr, mn, sc)
}

func load() saveStore {
	f, err := ioutil.ReadFile("C:/Users/Gebruiker/Desktop/python/Go/timer")
	if err != nil {
		log.Fatal(err)
	}
	s := saveStore{}
	json.Unmarshal(f, &s)
	return s
}

func Save(start time.Time) {
	sf := saveform{start, Print(time.Since(start))}
	s := load()
	s.Saves = append(s.Saves, sf)
	js, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile("C:/Users/Gebruiker/Desktop/python/Go/timer", js, 0644)
}
