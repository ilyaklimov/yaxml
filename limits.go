package yaxml

import (
	"time"
)

// Index is map, where key=hour(TimeIntervalFrom), value=key(Limits).
type Limits struct {
	index 		map[int]int 	`xml:"-" json:"-"`
	Response 	LimitsResponse 	`xml:"response" json:"response"`
}

type LimitsResponse struct {
	Limits 	[]Limit 		`xml:"limits>time-interval" json:"limits"`
}

type Limit struct {
    TimeIntervalFromString 	string 	`xml:"from,attr" json:"timeIntervalFrom"`
    TimeIntervalToString 	string 	`xml:"to,attr" json:"timeIntervalTo"`
    Value 					int 	`xml:",chardata" json:"value"`
}

func NewLimits() *Limits {
	ls := new(Limits)
	ls.index = make(map[int]int, 24)
	return ls
}

func (ls *Limits) createIndex() error {
	for k, l := range ls.Response.Limits {
		ls.index[l.TimeIntervalFrom().Hour()] = k
	}
	return nil
}

func (ls *Limits) All() *[]Limit {
	return &ls.Response.Limits
}

func (ls *Limits) Now() *Limit {
	h := time.Now().UTC().Hour()
	k := ls.index[h]
	return &ls.Response.Limits[k]
}

// For sleep, use time.Sleep((*Limits).Next().Sub(time.Now().UTC()))
func (ls *Limits) Next() *Limit {
	h := time.Now().UTC().Hour()
	k := 0
	if ls.index[h] < 23 {
		k = ls.index[h]+1
	}
	return &ls.Response.Limits[k]
}

// RPS (Request Per Second) - number of requests to the service.
// RPS = hourly_limits / 2000
// For function time.Sleep(), use time.Duration(result * 1000) * time.Millisecond,
// because RPS can be < 0.
func (ls *Limits) RPS() float32 {
	return float32(ls.Now().Value) / 2000
}

const limitLayout string = "2006-01-02 15:04:05 +0000"

func (l *Limit) TimeIntervalFrom() time.Time {
	t, _ := time.Parse(limitLayout, l.TimeIntervalFromString)
	return t
}

func (l *Limit) TimeIntervalTo() time.Time {
	t, _ := time.Parse(limitLayout, l.TimeIntervalToString)
	return t
}

