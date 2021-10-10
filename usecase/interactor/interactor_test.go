package interactor

import (
	"errors"
	"fmt"
	"testing"

	"github.com/AlexisDragneel/academy-go-q3202/domain/model"
	"github.com/AlexisDragneel/academy-go-q3202/usecase/repository"
)

type mockRepository struct {
}

type mockErrorRepository struct {
}

func (m *mockRepository) FindAll(p []*model.Pokemon) ([]*model.Pokemon, error) {
	return []*model.Pokemon{
		{},
	}, nil
}

func (m *mockRepository) FindAllAsync(p []*model.Pokemon, t string, items, itemsWorker int64) ([]*model.Pokemon, error) {
	return []*model.Pokemon{
		{},
	}, nil
}

func (m *mockRepository) FindById(p *model.Pokemon) (*model.Pokemon, error) {
	return &model.Pokemon{}, nil
}

func (m *mockRepository) PostPokemons(p []*model.Pokemon) (int, error) {
	return len(p), nil
}

func (m *mockErrorRepository) FindAll(p []*model.Pokemon) ([]*model.Pokemon, error) {
	return nil, errors.New("")
}

func (m *mockErrorRepository) FindAllAsync(p []*model.Pokemon, t string, items, itemsWorker int64) ([]*model.Pokemon, error) {
	return nil, errors.New("")
}

func (m *mockErrorRepository) FindById(p *model.Pokemon) (*model.Pokemon, error) {
	return nil, errors.New("")
}

func (m *mockErrorRepository) PostPokemons(p []*model.Pokemon) (int, error) {
	return 0, errors.New("")
}

type (
	testCase struct {
		message       string
		repo          repository.PokemonRepository
		expectedError error
	}

	testCaseGet struct {
		expectedResult []*model.Pokemon
		testCase
	}

	testCaseFindById struct {
		expectedResult *model.Pokemon
		testCase
	}

	testCasePostPokemons struct {
		expectedResult int
		testCase
	}
)

var (
	getTestCases = []testCaseGet{
		{
			testCase: testCase{
				message: "should return an array of pokemons",
				repo:    &mockRepository{},
			},
			expectedResult: []*model.Pokemon{
				{},
			},
		},
		{
			testCase: testCase{
				message:       "should return an error",
				repo:          &mockErrorRepository{},
				expectedError: errors.New(""),
			},
		},
	}

	findByTestCase = []testCaseFindById{
		{
			expectedResult: &model.Pokemon{},
			testCase: testCase{
				message: "should return a pokemon",
				repo:    &mockRepository{},
			},
		},
		{
			testCase: testCase{
				message:       "Should return an error",
				repo:          &mockErrorRepository{},
				expectedError: errors.New(""),
			},
		},
	}

	postTestCase = []testCasePostPokemons{
		{
			expectedResult: 1,
			testCase: testCase{
				message: "successful post pokemons",
				repo:    &mockRepository{},
			},
		},
		{
			testCase: testCase{
				message:       "unable to post pokemons",
				repo:          &mockErrorRepository{},
				expectedError: errors.New(""),
			},
		},
	}
)

func TestGet(t *testing.T) {
	for _, test := range getTestCases {
		fmt.Println(test.message)
		mockInteractor := NewPokemonInteractor(test.repo)
		var p []*model.Pokemon
		p, err := mockInteractor.Get(p)

		if p != nil {
			if len(p) != len(test.expectedResult) {
				t.Errorf("length must be %v and is %v", len(test.expectedResult), len(p))
			}
		}

		if err != nil {
			if err.Error() != test.expectedError.Error() {
				t.Error("unexpected error happens")
			}
		}

	}
}

func TestGetById(t *testing.T) {
	for _, test := range findByTestCase {
		fmt.Println(test.message)
		mockInteractor := NewPokemonInteractor(test.repo)
		var p *model.Pokemon
		p, err := mockInteractor.GetById(p)

		if p != nil {
			if p.ID != test.expectedResult.ID {
				t.Errorf("Id must be %v and is %v", test.expectedResult.ID, p.ID)
			}
		}

		if err != nil {
			if err.Error() != test.expectedError.Error() {
				t.Error("unexpected error happens")
			}
		}
	}
}

func TestPostPokemons(t *testing.T) {
	for _, test := range postTestCase {
		fmt.Println(test.message)
		mockInteractor := NewPokemonInteractor(test.repo)
		var p = []*model.Pokemon{{}}
		res, err := mockInteractor.PostPokemons(p)

		if p != nil {
			if res != test.expectedResult {
				t.Errorf("length must be %v and is %v", test.expectedResult, p)
			}
		}

		if err != nil {
			if err.Error() != test.expectedError.Error() {
				t.Error("unexpected error happens")
			}
		}
	}
}
