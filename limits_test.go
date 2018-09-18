package yaxml

import (
	"fmt"
	"testing"
	"time"
	"strings"
	"reflect"
)

var limitListTest []Limit = []Limit{
	{"2014-07-22 20:00:00 +0000", "2014-07-22 21:00:00 +0000", 500},
	{"2014-07-22 21:00:00 +0000", "2014-07-22 22:00:00 +0000", 450}, 
	{"2014-07-22 22:00:00 +0000", "2014-07-22 23:00:00 +0000", 590}, 
	{"2014-07-22 23:00:00 +0000", "2014-07-23 00:00:00 +0000", 600}, 
	{"2014-07-23 00:00:00 +0000", "2014-07-23 01:00:00 +0000", 300}, 
	{"2014-07-23 01:00:00 +0000", "2014-07-23 02:00:00 +0000", 200}, 
	{"2014-07-23 02:00:00 +0000", "2014-07-23 03:00:00 +0000", 500}, 
	{"2014-07-23 03:00:00 +0000", "2014-07-23 04:00:00 +0000", 500}, 
	{"2014-07-23 04:00:00 +0000", "2014-07-23 05:00:00 +0000", 500}, 
	{"2014-07-23 05:00:00 +0000", "2014-07-23 06:00:00 +0000", 100}, 
	{"2014-07-23 06:00:00 +0000", "2014-07-23 07:00:00 +0000", 100}, 
	{"2014-07-23 07:00:00 +0000", "2014-07-23 08:00:00 +0000", 100}, 
	{"2014-07-23 08:00:00 +0000", "2014-07-23 09:00:00 +0000", 100}, 
	{"2014-07-23 09:00:00 +0000", "2014-07-23 10:00:00 +0000", 200}, 
	{"2014-07-23 10:00:00 +0000", "2014-07-23 11:00:00 +0000", 300}, 
	{"2014-07-23 11:00:00 +0000", "2014-07-23 12:00:00 +0000", 300}, 
	{"2014-07-23 12:00:00 +0000", "2014-07-23 13:00:00 +0000", 300}, 
	{"2014-07-23 13:00:00 +0000", "2014-07-23 14:00:00 +0000", 300}, 
	{"2014-07-23 14:00:00 +0000", "2014-07-23 15:00:00 +0000", 300}, 
	{"2014-07-23 15:00:00 +0000", "2014-07-23 16:00:00 +0000", 300}, 
	{"2014-07-23 16:00:00 +0000", "2014-07-23 17:00:00 +0000", 400}, 
	{"2014-07-23 17:00:00 +0000", "2014-07-23 18:00:00 +0000", 500}, 
	{"2014-07-23 18:00:00 +0000", "2014-07-23 19:00:00 +0000", 500}, 
	{"2014-07-23 19:00:00 +0000", "2014-07-23 20:00:00 +0000", 600},	
}

func getLimitsExample() *Limits {
	ls := NewLimits()
	from := time.Now().UTC()
	to := from.Add(60 * time.Minute)
	ls.Response.Limits = make([]Limit, 24)
	for i := 0; i < 24; i++ {
		ls.Response.Limits[i] = Limit{
			TimeIntervalFromString: strings.TrimSuffix(from.String(), " UTC"),
			TimeIntervalToString: strings.TrimSuffix(to.String(), " UTC"),
			Value: (i+1)*100,
		}
		from, to = to, to.Add(60 * time.Minute)
	}
	err := ls.createIndex()
	if err != nil {
		fmt.Println(err)
	}
	return ls
}

func TestLimits_createIndex(t *testing.T) {
	ls := NewLimits()
	ls.Response.Limits = limitListTest
	err := ls.createIndex()
	if err != nil {
		fmt.Println(err)
	}
	expected := map[int]int{20:0, 21:1, 22:2, 23:3, 0:4, 1:5, 2:6, 3:7, 4:8, 5:9, 6:10, 7:11, 8:12, 9:13, 10:14, 11:15, 12:16, 13:17, 14:18, 15:19, 16:20, 17:21, 18:22, 19:23}
	if !reflect.DeepEqual(ls.index, expected) {
		t.Errorf("(*Limits).createIndex():\n got \t = %#v\n want \t = %#v", ls.index, expected)
	}
}

func TestLimits_All(t *testing.T) {
	ls := NewLimits()
	ls.Response.Limits = limitListTest
	got := *ls.All()
	if !reflect.DeepEqual(got, limitListTest) {
		t.Errorf("(*Limits).All(): \n got \t = %#v \n want \t = %#v", got, limitListTest)
	}
}

func TestLimits_Now(t *testing.T) {
	ls := getLimitsExample()
	l := ls.Now()
	if l.Value != 100 {
		t.Errorf("(*Limits).Now(): \n got \t = %#v \n want \t = 100", l.Value)
	}
}

func TestLimits_RPS(t *testing.T) {
	ls := getLimitsExample()
	actual := ls.RPS()
	var expected float32 = 100.0 / 2000.0
	if actual != expected {
		t.Errorf("(*Limits).RPS(): \n got \t = %#v \n want \t = %#v", actual, expected)
	}
}



