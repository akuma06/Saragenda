package saragenda

import (
	"net/url"
	"time"
	"github.com/laurent22/ical-go"
	"strings"
	"net/http"
	"fmt"
	"io/ioutil"
)

type AirbnbParser struct {
	URL *url.URL
	events []EventParsed
}

func (ap *AirbnbParser) Parse(icalUrl *url.URL) error {
	ap.URL = icalUrl
	events, err := ap.LoadIcal(icalUrl)
	if err != nil {
		return err
	}
	ap.events = make([]EventParsed, len(events))
	for i := range events {
		ap.events[i] = NewAirbnbEvent(events[i])
	}
	return nil
}

func (ap AirbnbParser) Events() []EventParsed {
	return ap.events
}

func (ap AirbnbParser) LoadIcal(icalUrl *url.URL) ([]*ical.Node, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", icalUrl.String(), nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	calNodes, err := ical.ParseCalendar(ap.SanitizeDesc(string(body)))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return calNodes.ChildrenByName("VEVENT"), nil
}

func (ap AirbnbParser) SanitizeDesc(body string) string {
	partOneDesc := strings.Split(body, "DESCRIPTION:")
	result := make([]string, len(partOneDesc))
		for i, part := range partOneDesc {
			parts := strings.Split(part, "SUMMARY")
			if len(parts) >= 2 {
				parts[0] = strings.Replace(parts[0], " ", "", -1)
				parts[0] = strings.Replace(parts[0], "\n", "", -1)
				newLines := strings.Split(parts[0], "\\n")
				parts[0] = strings.Join(newLines, "\n") + "END:DESCRIPTION\n"
			}
			result[i] = strings.Join(parts, "SUMMARY")
		}
	return strings.Join(result, "BEGIN:DESCRIPTION\n")
}

type AirbnbEvent struct {
	event *ical.Node
}

func NewAirbnbEvent(event *ical.Node) *AirbnbEvent {
	return &AirbnbEvent{event}
}

func (a AirbnbEvent) Firstname() string {
	if a.event == nil {
		return ""
	}
	summary := a.event.ChildByName("SUMMARY")
	if summary == nil || summary.Value == "Not available" {
		return ""
	}
	names := strings.Split(summary.Value, " ")
	if len(names) < 1 {
		return ""
	}
	return names[0]
}

func (a AirbnbEvent) Lastname() string {
	if a.event == nil {
		return ""
	}
	summary := a.event.ChildByName("SUMMARY")
	if summary == nil || summary.Value == "Not available" {
		return ""
	}
	names := strings.Split(summary.Value, " ")
	if len(names) < 2 {
		return ""
	}
	if names[1][0] == '(' { // In case there is no lastname, we don't want the airbnb ID
		return ""
	}
	return names[1]
}

func (a AirbnbEvent) Debut() time.Time {
	if a.event == nil || a.event.ChildByName("DTSTART") == nil {
		return time.Time{}
	}
	debut, err := time.Parse("20060102", a.event.ChildByName("DTSTART").Value)
	if err != nil {
		return time.Time{}
	}
	return debut
}

func (a AirbnbEvent) Fin() time.Time {
	if a.event == nil || a.event.ChildByName("DTEND") == nil {
		return time.Time{}
	}
	fin, err := time.Parse("20060102", a.event.ChildByName("DTEND").Value)
	if err != nil {
		return time.Time{}
	}
	return fin
}

func (a AirbnbEvent) Email() string {
	if a.event == nil {
		return ""
	}
	description := a.event.ChildrenByName("DESCRIPTION")
	if len(description) == 0{
		return ""
	}
	email := description[0].ChildByName("EMAIL")
	if email == nil {
		return ""
	}
	if email.Value == "(aucunaliasd'e-maildisponible)" {
		return ""
	}
	return email.Value
}

func (a AirbnbEvent) Phone() string {
	if a.event == nil {
		return ""
	}
	description := a.event.ChildByName("DESCRIPTION")
	if description == nil {
		return ""
	}
	phone := description.ChildByName("PHONE")
	if phone == nil {
		return ""
	}
	return phone.Value
}

func (a AirbnbEvent) TransactionId() string {
	if a.event == nil {
		return ""
	}
	summary := a.event.ChildByName("SUMMARY")
	if summary == nil || summary.Value == "Not available" || summary.Value == "Airbnb (Not available)" {
		return ""
	}
	names := strings.Split(summary.Value, " ")
	if len(names) < 2 {
		return ""
	}
	if len(names) == 2 {
		if len(names[1]) > 2 && names[1][0] == '(' {
			return names[1][1:(len(names[1]) - 1)]
		}
		return ""
	}
	if len(names[2]) < 3 { // if it's just ()
		return ""
	}
	return names[2][1:(len(names[2])-1)]
}

func (a AirbnbEvent) Type() string {
	return "airbnb"
}

func (a AirbnbEvent) UID() string {
	if a.event == nil || a.event.ChildByName("UID") == nil {
		return ""
	}
	return a.event.ChildByName("UID").Value
}
