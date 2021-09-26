package repository

import (
	"alexis.zapata-github.com/capstone-project/domain/model"
	"alexis.zapata-github.com/capstone-project/usecase/repository"
	"encoding/csv"
	"os"
	"strconv"
)

type pokemonRepository struct {
}

func NewPokemonRepository() repository.PokemonRepository {
	return &pokemonRepository{}
}

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

func (pr *pokemonRepository) FindById(p *model.Pokemon, id string) (*model.Pokemon, error) {
	records, err := readData("db.csv")

	if err != nil {
		return nil, err
	}

	pokemonId, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		return nil, err
	}

	for _, record := range records {
		csvId, err := strconv.ParseUint(record[0], 10, 32)

		if err != nil {
			return nil, err
		}

		if csvId != pokemonId {
			continue
		}

		p = &model.Pokemon{
			ID:   csvId,
			Name: record[1],
		}

		break
	}

	return p, nil
}

func (pr *pokemonRepository) PostPokemons(p []*model.Pokemon) (int, error) {

	f, err := openFile("db.csv", os.O_APPEND)

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
	f, err := openFile(fileName, os.O_WRONLY)

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

func openFile(fileName string, access int) (*os.File, error) {
	switch access {
	case os.O_APPEND:
		return os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0600)
	default:
		return os.Open(fileName)
	}
}
