package saragenda

import (
	"reflect"
	"testing"
	"time"

	"github.com/laurent22/ical-go"
)

func getAirbnbNode(name string) *ical.Node {
	nodeStrings := map[string]string{
		"empty": ``,
		"valid": `BEGIN:VEVENT
DTEND;VALUE=DATE:20180904
DTSTART;VALUE=DATE:20180903
UID:1418fb94e984-4e06061a747a4e4680112c2ed2f01919@airbnb.com
DESCRIPTION:CHECKIN: 03/09/2018\nCHECKOUT: 04/09/2018\nNIGHTS: 1\nPHONE: 
 +33 6 77 22 33 09\nEMAIL: paul-0z7aghtrxyw80cec@guest.airbnb.com\nPRO
 PERTY: Les chambres du soleil levant\n
SUMMARY:Paul Salmon (HDMTRFM9VC)
LOCATION:Les chambres du soleil levant
END:VEVENT`,
		"notvalid": `BEGIN:VEVENT
DTEND;VALUE=DATE:20180826
DTSTART;VALUE=DATE:20180824
UID:6fec1092d3fa-c9e130257954d2b42ed9e1995131731a@airbnb.com
SUMMARY:Not available
END:VEVENT`,
	}
	node, err := ical.ParseCalendar(correctIcal(nodeStrings[name]))
	if err != nil {
		return nil
	}

	return node
}

func TestAirbnbEvent_Firstname(t *testing.T) {
	tests := []struct {
		name string
		b    AirbnbEvent
		want string
	}{
		{"empty", NewAirbnbEvent(getAirbnbNode("empty")), ""},
		{"valid", NewAirbnbEvent(getAirbnbNode("valid")), "Paul"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Firstname(); got != tt.want {
				t.Errorf("AirbnbEvent.Firstname() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAirbnbEvent_Lastname(t *testing.T) {
	tests := []struct {
		name string
		b    AirbnbEvent
		want string
	}{
		{"empty", NewAirbnbEvent(getAirbnbNode("empty")), ""},
		{"valid", NewAirbnbEvent(getAirbnbNode("valid")), "Salmon"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Lastname(); got != tt.want {
				t.Errorf("AirbnbEvent.Lastname() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAirbnbEvent_Debut(t *testing.T) {
	debutTest, _ := time.Parse("02/01/2006", "03/09/2018")
	tests := []struct {
		name string
		b    AirbnbEvent
		want time.Time
	}{
		{"empty", NewAirbnbEvent(getAirbnbNode("empty")), time.Time{}},
		{"valid", NewAirbnbEvent(getAirbnbNode("valid")), debutTest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Debut(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AirbnbEvent.Debut() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAirbnbEvent_Fin(t *testing.T) {
	finTest, _ := time.Parse("02/01/2006", "04/09/2018")
	tests := []struct {
		name string
		b    AirbnbEvent
		want time.Time
	}{
		{"empty", NewAirbnbEvent(getAirbnbNode("empty")), time.Time{}},
		{"valid", NewAirbnbEvent(getAirbnbNode("valid")), finTest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Fin(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AirbnbEvent.Fin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAirbnbEvent_Email(t *testing.T) {
	type fields struct {
		event *ical.Node
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"empty", fields{getAirbnbNode("empty")}, ""},
		{"valid", fields{getAirbnbNode("valid")}, "paul-0z7aghtrxyw80cec@guest.airbnb.com"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := AirbnbEvent{
				event: tt.fields.event,
			}
			if got := a.Email(); got != tt.want {
				t.Errorf("AirbnbEvent.Email() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAirbnbEvent_Phone(t *testing.T) {
	type fields struct {
		event *ical.Node
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"empty", fields{getAirbnbNode("empty")}, ""},
		{"valid", fields{getAirbnbNode("valid")}, "+33677223309"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := AirbnbEvent{
				event: tt.fields.event,
			}
			if got := a.Phone(); got != tt.want {
				t.Errorf("AirbnbEvent.Phone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAirbnbEvent_TransactionId(t *testing.T) {
	type fields struct {
		event *ical.Node
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"empty", fields{getAirbnbNode("empty")}, ""},
		{"valid", fields{getAirbnbNode("valid")}, "HDMTRFM9VC"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := AirbnbEvent{
				event: tt.fields.event,
			}
			if got := a.TransactionId(); got != tt.want {
				t.Errorf("AirbnbEvent.TransactionId() = %v, want %v", got, tt.want)
			}
		})
	}
}
