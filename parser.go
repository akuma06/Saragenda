package saragenda

import (
	"time"
	"net/url"
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
	Type() string
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