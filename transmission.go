package usrp

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"
)

type Transmission struct {
	Group     string       // name of a group this transmission belongs to
	StartTime time.Time    // start of Transmission
	EndTime   time.Time    // end of Transmission
	Callsign  string       // who's transmitting
	Audio     bytes.Buffer // raw audio payload
}

func NewTransmission(group, callsign string) Transmission {
	log.Printf("BEGIN TX - %s - %s", group, callsign)
	return Transmission{
		Group:     group,
		Callsign:  callsign,
		StartTime: time.Now(),
	}
}

func (t *Transmission) WriteAudioToFile(OutputPath string) (int64, error) {
	timestampID := time.Now().UTC().Unix()
	filename := fmt.Sprintf("%s-%d-%s.pcm", t.Group, timestampID, t.Callsign)
	err := os.WriteFile(OutputPath+string(os.PathSeparator)+filename, t.Audio.Bytes(), 0600)
	if err != nil {
		return timestampID, err
	}
	return timestampID, nil
}

func (t *Transmission) EndTransmission() {
	t.EndTime = time.Now()
	timeElapsed := t.EndTime.Sub(t.StartTime)
	log.Printf("END TX - %s - %s - %v - %d bytes", t.Group, t.Callsign, timeElapsed, t.Audio.Len())
	if timeElapsed.Seconds() < 5 || timeElapsed.Seconds() > 200 {
		log.Println("Not writing to disk .. elapsed time is too short or too long")
		return
	}
	id, err := t.WriteAudioToFile("./audio")
	if err != nil {
		log.Printf("didn't write audio: %v", err)
		return
	}
	log.Printf("Wrote %d to disk", id)
}
