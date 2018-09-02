package saragenda

import (
	"reflect"
	"testing"
	"time"

	"github.com/laurent22/ical-go"
)

func getNode(name string) *ical.Node {
	nodeStrings := map[string]string{
		"empty": ``,
		"valid": `BEGIN:VEVENT
DTSTART;VALUE=DATE:20180828
DTEND;VALUE=DATE:20180830
UID:49cbd9bb67c47200868a38bf20dd5e54@booking.com
SUMMARY: CLOSED - Laure Blanche
END:VEVENT`,
	}
	node, err := ical.ParseCalendar(nodeStrings[name])
	if err != nil {
		return nil
	}

	return node
}

func TestBookingEvent_Firstname(t *testing.T) {
	tests := []struct {
		name string
		b    BookingEvent
		want string
	}{
		{"empty", NewBookingEvent(getNode("empty")), ""},
		{"valid", NewBookingEvent(getNode("valid")), "Laure"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Firstname(); got != tt.want {
				t.Errorf("BookingEvent.Firstname() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBookingEvent_Lastname(t *testing.T) {
	tests := []struct {
		name string
		b    BookingEvent
		want string
	}{
		{"empty", NewBookingEvent(getNode("empty")), ""},
		{"valid", NewBookingEvent(getNode("valid")), "Blanche"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Lastname(); got != tt.want {
				t.Errorf("BookingEvent.Lastname() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBookingEvent_Debut(t *testing.T) {
	debutTest, _ := time.Parse("02/01/2006", "28/08/2018")
	tests := []struct {
		name string
		b    BookingEvent
		want time.Time
	}{
		{"empty", NewBookingEvent(getNode("empty")), time.Time{}},
		{"valid", NewBookingEvent(getNode("valid")), debutTest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Debut(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BookingEvent.Debut() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBookingEvent_Fin(t *testing.T) {
	finTest, _ := time.Parse("02/01/2006", "30/08/2018")
	tests := []struct {
		name string
		b    BookingEvent
		want time.Time
	}{
		{"empty", NewBookingEvent(getNode("empty")), time.Time{}},
		{"valid", NewBookingEvent(getNode("valid")), finTest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Fin(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BookingEvent.Fin() = %v, want %v", got, tt.want)
			}
		})
	}
}
