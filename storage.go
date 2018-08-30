package saragenda

type Storage interface {
	GetChambres() (map[string]*Chambre, error)
	AddChambre(chambre *Chambre) error
	GetChambre(keyChambre string) (*Chambre, error)
	EditChambre(keyChambre string, chambre *Chambre) error
	DeleteChambre(keyChambre string) error
	GetReservations(keyChambre string) (Reservations, error)
	AddReservation(keyChambre string, reservation *Reservation) error
	GetReservation(keyChambre string, id uint) (*Reservation, error)
	EditReservation(keyChambre string, id uint, reservation *Reservation) error
	DeleteReservation(keyChambre string, id uint) error
}