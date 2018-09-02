package saragenda

import (
	"github.com/laurent22/ical-go"
	"time"
	"strings"
	"net/url"
)

type BookingParser struct {
	URL *url.URL
	events []EventParsed
}

func (bp BookingParser) Events() []EventParsed {
	return bp.events
}

func (bp BookingParser) Parse(icalUrl *url.URL) error {
	bp.URL = icalUrl
	events, err := getVEventsFromIcal(icalUrl)
	if err != nil {
		return err
	}
	bp.events = make([]EventParsed, len(events))
	for i := range events {
		bp.events[i] = NewBookingEvent(events[i])
	}
	return nil
}

type BookingEvent struct {
	event *ical.Node
}

func NewBookingEvent(event *ical.Node) BookingEvent {
	return BookingEvent{event}
}

func (b BookingEvent) Firstname() string  {
	if b.event == nil || b.event.ChildByName("SUMMARY") == nil {
		return ""
	}
	desc := strings.Split(b.event.ChildByName("SUMMARY").Value, " - ")
	if len(desc) == 2 {
		if desc[1] == "Not Available" {
			return ""
		}
		names := strings.Split(desc[1], " ")
		return names[0]
	}
	return ""
}

func (b BookingEvent) Lastname() string {
	if b.event == nil || b.event.ChildByName("SUMMARY") == nil {
		return ""
	}
	desc := strings.Split(b.event.ChildByName("SUMMARY").Value, " - ")
	if len(desc) == 2 {
		if desc[1] == "Not Available" {
			return ""
		}
		names := strings.Split(desc[1], " ")
		return names[1]
	}
	return ""
}
func (b BookingEvent) Debut() time.Time {
	if b.event == nil || b.event.ChildByName("DTSTART") == nil {
		return time.Time{}
	}
	debut, err := time.Parse("20060102", b.event.ChildByName("DTSTART").Value)
	if err != nil {
		return time.Time{}
	}
	return debut
}
func (b BookingEvent) Fin() time.Time {
	if b.event == nil || b.event.ChildByName("DTEND") == nil {
		return time.Time{}
	}
	fin, err := time.Parse("20060102", b.event.ChildByName("DTEND").Value)
	if err != nil {
		return time.Time{}
	}
	return fin
}

func (b BookingEvent) Email() string {
	return ""
}
func (b BookingEvent) Phone() string {
	return ""
}
func (b BookingEvent) TransactionId() string {
	return ""
}