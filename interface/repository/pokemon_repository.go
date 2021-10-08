package repository

import (
	"encoding/csv"
	"github.com/AlexisDragneel/academy-go-q3202/domain/model"
	"github.com/AlexisDragneel/academy-go-q3202/usecase/repository"
	"os"
	"strconv"
)

type pokemonRepository struct {
}

// NewPokemonRepository expose the creation of the repository to handle the dependency injection
func NewPokemonRepository() repository.PokemonRepository {
	return &pokemonRepository{}
}

// FindAll opens a file and read the whole file and convert the information to and array of pokemons
func (pr *pokemonRepository) FindAll(p []*model.Pokemon) ([]*model.Pokemon, error) {
	records, err := readData("db.csv")

	if err != nil {
		return nil, err
	}

	for _, record := range records {
		id, err := strconv.ParseUint(record[0], 10, 32)

		if err != nil {
			return nil, err
		}

		pokemon := &model.Pokemon{
			ID:   id,
			Name: record[1],
		}

		p = append(p, pokemon)
	}

	return p, nil

}

// FindById opens a file a read the data until finds the expected pokemon by id
func (pr *pokemonRepository) FindById(p *model.Pokemon) (*model.Pokemon, error) {
	records, err := readData("db.csv")

	if err != nil {
		return nil, err
	}

	for _, record := range records {
		csvId, err := strconv.ParseUint(record[0], 10, 32)

		if err != nil {
			return nil, err
		}

		if csvId != p.ID {
			continue
		}

		p.Name = record[1]

		break
	}

	return p, nil
}

// PostPokemons receive an array of pokemons and append this information to the corresponding file
func (pr *pokemonRepository) PostPokemons(p []*model.Pokemon) (int, error) {

	f, err := openFile("db.csv")

	if err != nil {
		return 0, err
	}

	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()
	for _, pokemon := range p {
		if err := w.Write(pokemon.ToStringArr()); err != nil {
			return 0, err
		}
	}
	return len(p), nil
}

func readData(fileName string) ([][]string, error) {
	f, err := openFile(fileName)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	r := csv.NewReader(f)

	// skip first line
	if _, err := r.Read(); err != nil {
		return nil, err
	}

	records, err := r.ReadAll()

	if err != nil {
		return nil, err
	}

	return records, nil
}

func openFile(fileName string) (*os.File, error) {
	return os.OpenFile(fileName, os.O_RDWR|os.O_APPEND, 0600)
}
