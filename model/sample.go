package model

import "time"

type Sample struct {
	ID        int            `json:"id"`
	UserID    int            `json:"user_id"`
	Name      string         `json:"name"`
	BPM       int            `json:"bpm"`
	Key       SampleKey      `json:"key"`
	KeyScale  SampleKeyScale `json:"key_scale"`
	Time      int            `json:"time"`
	FileURL   string         `json:"file_url"`
	CoverURL  string         `json:"cover_url"`
	Price     float64        `json:"price"`
	CreatedAt time.Time      `json:"created_at"`
}

type SampleKey string

const (
	SampleKeyC  SampleKey = "C"
	SampleKeyCs SampleKey = "C#"
	SampleKeyDb SampleKey = "Db"
	SampleKeyD  SampleKey = "D"
	SampleKeyDs SampleKey = "D#"
	SampleKeyEb SampleKey = "Eb"
	SampleKeyE  SampleKey = "E"
	SampleKeyF  SampleKey = "F"
	SampleKeyFs SampleKey = "F#"
	SampleKeyGb SampleKey = "Gb"
	SampleKeyG  SampleKey = "G"
	SampleKeyGs SampleKey = "G#"
	SampleKeyAb SampleKey = "Ab"
	SampleKeyA  SampleKey = "A"
	SampleKeyAs SampleKey = "A#"
	SampleKeyBb SampleKey = "Bb"
	SampleKeyB  SampleKey = "B"
)

var SampleValidKeys = []SampleKey{
	SampleKeyC,
	SampleKeyCs,
	SampleKeyDb,
	SampleKeyD,
	SampleKeyDs,
	SampleKeyEb,
	SampleKeyE,
	SampleKeyF,
	SampleKeyFs,
	SampleKeyGb,
	SampleKeyG,
	SampleKeyGs,
	SampleKeyAb,
	SampleKeyA,
	SampleKeyAs,
	SampleKeyBb,
	SampleKeyB,
}

type SampleKeyScale string

const (
	SampleKeyScaleMajor SampleKeyScale = "major"
	SampleKeyScaleMinor SampleKeyScale = "minor"
)

var SampleValidKeyScales = []SampleKeyScale{
	SampleKeyScaleMajor,
	SampleKeyScaleMinor,
}
