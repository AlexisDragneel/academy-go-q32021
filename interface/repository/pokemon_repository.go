package repository

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/AlexisDragneel/academy-go-q3202/domain/model"
	"github.com/AlexisDragneel/academy-go-q3202/usecase/repository"
	"github.com/AlexisDragneel/academy-go-q3202/utils"
)

type pokemonRepository struct {
}

var wg sync.WaitGroup

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

	return parsePokemons(p, records)
}

func (pr *pokemonRepository) FindAllAsync(p []*model.Pokemon, t string, items, itemsWorker int64) ([]*model.Pokemon, error) {
	return readDataAsync("db.csv", p, t, items, itemsWorker)
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

func worker(id int64, r *csv.Reader, t string, jobs <-chan int64, results chan<- []string) {
	defer wg.Done()
	for j := range jobs {
		log.Printf("worker %v starting job %v", id, j)

		for {
			line, err := r.Read()
			if err == io.EOF {
				break
			}

			if len(line) != 2 {
				continue
			}

			pid, err := strconv.ParseUint(line[0], 10, 32)
			if shouldBeAdded(t, pid) {
				results <- line
				break
			}
		}
	}
}

func readDataAsync(fileName string, p []*model.Pokemon, t string, items, itemsWorker int64) ([]*model.Pokemon, error) {

	var result [][]string

	jobs := make(chan int64, items)
	lines := make(chan []string, items)

	workers := items / itemsWorker

	f, err := openFile(fileName)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	r, err := openReader(f)

	if err != nil {
		return nil, err
	}

	for w := int64(0); w < workers; w++ {
		wg.Add(1)
		go worker(w, r, t, jobs, lines)
	}

	go func() {
		for j := int64(0); j < items; j++ {
			jobs <- j
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(lines)
	}()

	for line := range lines {
		result = append(result, line)
	}

	return parsePokemons(p, result)
}

func readData(fileName string) ([][]string, error) {
	f, err := openFile(fileName)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	r, err := openReader(f)

	if err != nil {
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

func openReader(f *os.File) (*csv.Reader, error) {
	r := csv.NewReader(f)
	r.FieldsPerRecord = -1
	// skip first line
	if _, err := r.Read(); err != nil {
		return nil, err
	}
	return r, nil
}

func parsePokemons(p []*model.Pokemon, records [][]string) ([]*model.Pokemon, error) {
	for _, record := range records {
		if len(record) != 2 {
			continue
		}

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

func shouldBeAdded(t string, id uint64) bool {
	switch t {
	case utils.Odd:
		return int(id)%2 != 0
	case utils.Even:
		return int(id)%2 == 0
	default:
		return true
	}
}
