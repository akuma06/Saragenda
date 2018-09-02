package saragenda

import "testing"

func Test_correctIcal(t *testing.T) {
	testBody := `BEGIN:VCALENDAR
PRODID;X-RICAL-TZSOURCE=TZINFO:-//Airbnb Inc//Hosting Calendar 0.8.8//EN
CALSCALE:GREGORIAN
VERSION:2.0
BEGIN:VEVENT
DTEND;VALUE=DATE:20180402
DTSTART;VALUE=DATE:20180401
UID:1418fb94e984-caff16e26e074a77391d4d089dc79ce3@airbnb.com
DESCRIPTION:CHECKIN: 01/04/2018\nCHECKOUT: 02/04/2018\nNIGHTS: 1\nPHONE: 
 +33 6 88 88 77 77\nEMAIL: (aucun alias d'e-mail disponible)\nPROPERTY: L
 es chambres du soleil levant\n
SUMMARY:Paul Salmon (HMTGHRDQDK)
LOCATION:Les chambres du soleil levant
END:VEVENT
BEGIN:VEVENT
DTEND;VALUE=DATE:20180919
DTSTART;VALUE=DATE:20180918
UID:6fec1092d3fa-ccf07d1e706802211644237885e736ee@airbnb.com
SUMMARY:Not available
END:VEVENT
END:VCALENDAR`
	validBody := `BEGIN:VCALENDAR
PRODID;X-RICAL-TZSOURCE=TZINFO:-//Airbnb Inc//Hosting Calendar 0.8.8//EN
CALSCALE:GREGORIAN
VERSION:2.0
BEGIN:VEVENT
DTEND;VALUE=DATE:20180402
DTSTART;VALUE=DATE:20180401
UID:1418fb94e984-caff16e26e074a77391d4d089dc79ce3@airbnb.com
BEGIN:DESCRIPTION
CHECKIN:01/04/2018
CHECKOUT:02/04/2018
NIGHTS:1
PHONE:+33688887777
EMAIL:(aucunaliasd'e-maildisponible)
PROPERTY:Leschambresdusoleillevant
END:DESCRIPTION
SUMMARY:Paul Salmon (HMTGHRDQDK)
LOCATION:Les chambres du soleil levant
END:VEVENT
BEGIN:VEVENT
DTEND;VALUE=DATE:20180919
DTSTART;VALUE=DATE:20180918
UID:6fec1092d3fa-ccf07d1e706802211644237885e736ee@airbnb.com
SUMMARY:Not available
END:VEVENT
END:VCALENDAR`
	type args struct {
		body string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"valid", args{testBody}, validBody},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := correctIcal(tt.args.body); got != tt.want {
				t.Errorf("correctIcal() = %v, want %v", got, tt.want)
			}
		})
	}
}
