package saragenda

import (
	"testing"
)

func TestChambre_String(t *testing.T) {
	tests := []struct {
		name    string
		chambre Chambre
		want    string
	}{
		{
			"ChambreEmpty",
			Chambre{},
			`+-------------------+
|                   |
|   () Chambre ''   |
|                   |
+-------------------+
| Réservations (0)  |
+-------------------+`,
		},
		{
			"ChambreValid",
			Chambre{
				ID:           "pharaon",
				Name:         "Suite Pharaon",
				Reservations: []*Reservation{},
				ToCheck:      nil,
			},
			`+---------------------------------------+
|                                       |
|   (pharaon) Chambre 'Suite Pharaon'   |
|                                       |
+---------------------------------------+
|           Réservations (0)            |
+---------------------------------------+`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.chambre.String(); got != tt.want {
				t.Errorf("Chambre.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChambres_String(t *testing.T) {
	tests := []struct {
		name string
		ch   Chambres
		want string
	}{
		{
			"ChambresEmpty",
			Chambres{},
			"",
		},
		{
			"ChambresValid",
			Chambres{"pharaon": &Chambre{"pharaon", "Suite Pharaon", []*Reservation{}, nil}},
			`+---------------------------------------+
|                                       |
|   (pharaon) Chambre 'Suite Pharaon'   |
|                                       |
+---------------------------------------+
|           Réservations (0)            |
+---------------------------------------+

`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ch.String(); got != tt.want {
				t.Errorf("Chambres.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetManager(t *testing.T) {
	type args struct {
		conf Config
		st   Storage
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"NillSet",
			args{nil, nil},
			true,
		},
		{
			"ValidSet",
			args{NewViperWrapper(), NewMemoryStore()},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SetManager(tt.args.conf, tt.args.st); (err != nil) != tt.wantErr {
				t.Errorf("SetManager() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

