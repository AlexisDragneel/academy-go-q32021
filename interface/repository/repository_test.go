package repository

import (
	"testing"

	"github.com/AlexisDragneel/academy-go-q3202/domain/model"
)

func TestPokemonRepository_FindAll(t *testing.T) {
	records, _ := readData("db.csv")
	var p []*model.Pokemon

	pr := NewPokemonRepository()

	if p, _ = pr.FindAll(p); len(p) != len(records) {
		t.Errorf("excpetd to be: %v but got: %v", len(records), len(p))
	}
}

func TestPokemonRepository_FindById(t *testing.T) {
	records, _ := readData("db.csv")
	expected := records[0]
	p := &model.Pokemon{
		ID: uint64(1),
	}
	pr := NewPokemonRepository()
	if p, _ = pr.FindById(p); p.Name != expected[1] {
		t.Errorf("id should be %v but is: %v", expected[1], p.Name)
	}
}
