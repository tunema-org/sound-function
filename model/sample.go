package model

import "time"

type Sample struct {
	ID        int
	UserID    int
	Name      string
	BPM       int
	Key       SampleKey
	KeyScale  SampleKeyScale
	Time      int
	FileURL   string
	CoverUrl  string
	Price     float64
	CreatedAt time.Time
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
