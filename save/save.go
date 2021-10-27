package save

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"log"
	"time"
)

type saveform struct {
	start    time.Time
	duration time.Duration
}

func Save(start time.Time, duration time.Duration) {
	sf := saveform{start, duration}
	js, err := json.MarshalIndent(sf, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile("save.json", js, fs.ModeAppend)
}
