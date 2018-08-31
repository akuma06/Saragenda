package saragenda

import (
	"errors"
	"github.com/gosimple/slug"
)

type MemoryStore struct {
	Chambres map[string]*Chambre
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		Chambres: make(map[string]*Chambre),
	}
}

func (m *MemoryStore) GetChambres() (map[string]*Chambre, error) {
	return m.Chambres, nil
}

func (m *MemoryStore) GetChambre(keyChambre string) (*Chambre, error) {
	if chambre, ok := m.Chambres[keyChambre]; ok {
		return chambre, nil
	}
	return nil, errors.New("impossible de trouver la chambre")
}

func (m *MemoryStore) AddChambre(chambre *Chambre) error {
	key := slug.Make(chambre.Name)
	if _, ok := m.Chambres[key]; ok {
		return errors.New("la Chambre existe déjà")
	}
	chambre.ID = key
	m.Chambres[key] = chambre
	return nil
}

func (m *MemoryStore) EditChambre(keyChambre string, chambre *Chambre) error  {
	if _, ok := m.Chambres[keyChambre]; ok {
		m.Chambres[keyChambre] = chambre
		return nil
	}
	return errors.New("la Chambre n'existe pas")
}

func (m *MemoryStore) DeleteChambre(keyChambre string) error  {
	if _, ok := m.Chambres[keyChambre]; ok {
		delete(m.Chambres, keyChambre)
		return nil
	}
	return errors.New("la Chambre n'existe pas")
}

func (m *MemoryStore) AddReservation(keyChambre string, reservation *Reservation) error  {
	chambre, err := m.GetChambre(keyChambre)
	if err != nil {
		return err
	}
	reservation.ID = uint(len(chambre.Reservations))
	chambre.Reservations = append(m.Chambres[keyChambre].Reservations, reservation)
	return nil
}

func (m *MemoryStore) GetReservation(keyChambre string, id uint) (*Reservation, error)  {
	chambre, err := m.GetChambre(keyChambre)
	if err != nil {
		return nil, err
	}
	if int(id) < len(chambre.Reservations) {
		return chambre.Reservations[id], nil
	}
	return nil, errors.New("impossible de trouver la réservation")
}


func (m *MemoryStore) GetReservations(keyChambre string) (Reservations, error)  {
	chambre, err := m.GetChambre(keyChambre)
	if err != nil {
		return nil, err
	}
	return chambre.Reservations, nil
}


func (m *MemoryStore) EditReservation(keyChambre string, id uint, reservation *Reservation) error  {
	_, err := m.GetReservation(keyChambre, id)
	if err != nil {
		return err
	}
	chambre, _ := m.GetChambre(keyChambre)
	chambre.Reservations[id] = reservation
	return nil
}

func (m *MemoryStore) DeleteReservation(keyChambre string, id uint) error {
	_, err := m.GetReservation(keyChambre, id)
	if err != nil {
		return err
	}
	chambre, _ := m.GetChambre(keyChambre)
	chambre.Reservations =  append(chambre.Reservations[:id], chambre.Reservations[id+1:]...)
	for i, res := range chambre.Reservations {
		res.ID = uint(i)
	}
	return nil
}