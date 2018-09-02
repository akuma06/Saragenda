package saragenda

import (
	"time"
	"github.com/laurent22/ical-go"
	"net/url"
	"net/http"
	"fmt"
	"io/ioutil"
	"strings"
)

type Parser interface {
	Parse(urlToParse *url.URL) error
	Events() []EventParsed
}

type EventParsed interface {
	Firstname() string
	Lastname() string
	Debut() time.Time
	Fin() time.Time
	Email() string
	Phone() string
	TransactionId() string
}

func getVEventsFromIcal(icalUrl *url.URL) ([]*ical.Node, error) {
	resp, err := http.Get(icalUrl.String())
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

	calNodes, err := ical.ParseCalendar(correctIcal(string(body)))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return calNodes.ChildrenByName("VEVENT"), nil
}

func correctIcal(body string) string {
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