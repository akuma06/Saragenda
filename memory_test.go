package saragenda

import (
	"reflect"
	"testing"
)

func TestNewMemoryStore(t *testing.T) {
	tests := []struct {
		name string
		want *MemoryStore
	}{
		{
			"StoreValid",
			&MemoryStore{
				Chambres: make(map[string]*Chambre),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMemoryStore(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMemoryStore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMemoryStore_GetChambres(t *testing.T) {
	type fields struct {
		Chambres map[string]*Chambre
	}
	tests := []struct {
		name    string
		fields  fields
		want    map[string]*Chambre
		wantErr bool
	}{
		{
			"NoChambres",
			fields{map[string]*Chambre{} },
			map[string]*Chambre{},
			false,
		},
		{
			"ChambresValid",
			fields{
					Chambres: map[string]*Chambre{"pharaon": {"pharaon", "Suite Pharaon", []*Reservation{}, nil}},
			},
			map[string]*Chambre{"pharaon": {"pharaon", "Suite Pharaon", []*Reservation{}, nil}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MemoryStore{
				Chambres: tt.fields.Chambres,
			}
			got, err := m.GetChambres()
			if (err != nil) != tt.wantErr {
				t.Errorf("MemoryStore.GetChambres() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MemoryStore.GetChambres() = %v, want %v", got, tt.want)
			}
		})
	}
}
