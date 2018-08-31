package saragenda

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/laurent22/ical-go"
	"errors"
)

type Reservation struct {
	ID uint
	Debut       time.Time
	Fin         time.Time
	Description string
	Errors []error
}

type Config interface {
	UnmarshalKey(key string, rawVal interface{}) error
}

type Reservations []*Reservation

type Chambre struct {
	ID string
	Name         string
	Reservations Reservations
	ToCheck []string
}

func (chambre Chambre) String() string {
	head := fmt.Sprintf("   (%s) Chambre '%s'   ", chambre.ID, chambre.Name)
	reservations := fmt.Sprintf("Réservations (%d)", len(chambre.Reservations))
	tirets := "+"
	emptySpaces := ""
	for range head {
		tirets += "-"
		emptySpaces += " "
	}
	tirets += "+"

	nbResSpaces := len(reservations)+(len(emptySpaces)-len(reservations))/2
	reservations = emptySpaces[nbResSpaces:] + reservations + emptySpaces[nbResSpaces:]
	if nbResSpaces%2 == 0 {
		reservations += " "
	}
	return tirets+"\n|"+emptySpaces+"|\n|"+head+"|\n|"+emptySpaces+"|\n"+tirets+"\n|"+reservations+"|\n"+tirets
}

type Chambres map[string]*Chambre

func (ch Chambres) String() string {
	st := ""
	for _, c := range ch {
		st += fmt.Sprintf("%v\n\n", c)
	}
	return st
}

var store Storage

var config Config

func SetManager(conf Config, st Storage) error {
	if conf == nil || st == nil {
		return errors.New("Nil parameters supplied to the Manager")
	}
	config = conf
	store = st
	err := initChambres()
	return err
}

func initChambres() error {
	var chambres Chambres
	err := config.UnmarshalKey("chambres", &chambres)
	if err != nil {
		fmt.Println("Can't read config")
		fmt.Println(err)
		return err
	}
	for _, chambre := range chambres {
		err = store.AddChambre(chambre)
		if err != nil {
			fmt.Println("Can't add chambre to storage")
			fmt.Println(err)
		}
	}
	return nil
}

func LoadChambres() error {
	chambres, err := store.GetChambres()
	if err != nil {
		return err
	}
	for key := range chambres {
		getReservations(key)
	}
	return nil
}

func GetChambres() (Chambres, error) {
	return store.GetChambres()
}

func getReservations(name string) (*Chambre, error) {
	chambre, err := store.GetChambre(name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for _, toCheck := range chambre.ToCheck {
		queryUrl(toCheck, chambre)
	}
	err = store.EditChambre(name, chambre)
	if err != nil {
		fmt.Println(err)
		return chambre, err
	}
	return chambre, nil
}

func queryUrl(url string, chambre *Chambre) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	calNodes, err := ical.ParseCalendar(string(body))
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, event := range calNodes.ChildrenByName("VEVENT") {
		debut, _ := time.Parse("20060102", event.ChildByName("DTSTART").Value)
		fin, _ := time.Parse("20060102", event.ChildByName("DTEND").Value)
		reservation :=  Reservation{0, debut, fin, event.ChildByName("SUMMARY").Value, []error{}}
		// checkDoubleTime(chambre, &reservation)
		store.AddReservation(chambre.ID, &reservation)
	}
	return
}

/*
func checkDoubleTime(chambre *Chambre, reservation *Reservation) {
	for _, res := range chambre.Reservations {
		if reservation.Debut.Sub(res.Debut) > 0 && res.Fin.Sub(reservation.Debut) > 0 { // reservation.Debut, res.Debut, reservation.Fin
			reservation.Errors = append(reservation.Errors, errors.New("Déjà réservé avec : " + res.Description))
			res.Errors = append(res.Errors, errors.New("Déjà réservé avec : " + reservation.Description))
			store.EditReservation(chambre.ID, res.ID, res)
		}
	}
} */