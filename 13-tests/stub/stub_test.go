package stub_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	. "github.com/junhyuk0801/learning-go/13-tests/stub"
)

type GetPetNamesStub struct {
	Entities
}

func (ps GetPetNamesStub) GetPets(userId string) ([]Pet, error) {
	switch userId {
	case "1":
		return []Pet{{Name: "Bubbles"}}, nil
	case "2":
		return []Pet{{Name: "Stampy"}, {Name: "Snowball"}}, nil
	default:
		return nil, fmt.Errorf("invalid id: %s", userId)
	}
}

func TestLogicGetPetNames(t *testing.T) {
	data := []struct {
		name     string
		userID   string
		petNames []string
	}{
		{"case1", "1", []string{"Bubbles"}},
		{"case2", "2", []string{"Stampy", "Snowball"}},
		{"case3", "3", nil},
	}

	l := Logic{GetPetNamesStub{}}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			petNames, err := l.GetPetNames(d.userID)
			if err != nil {
				t.Error(err)
			}
			if diff := cmp.Diff(d.petNames, petNames); diff != "" {
				t.Error(diff)
			}
		})
	}
}

type EntitiesStub struct {
	getUser     func(id string) (User, error)
	getPets     func(userId string) ([]Pet, error)
	getChildren func(userId string) ([]Person, error)
	getFriends  func(userId string) ([]Person, error)
	saveUser    func(user User) error
}

func (es EntitiesStub) GetUser(id string) (User, error) {
	return es.getUser(id)
}

func (es EntitiesStub) GetPets(userId string) ([]Pet, error) {
	return es.getPets(userId)
}

func (es EntitiesStub) GetChildren(userId string) ([]Person, error) {
	return es.getChildren(userId)
}

func (es EntitiesStub) GetFriends(userId string) ([]Person, error) {
	return es.getFriends(userId)
}

func (es EntitiesStub) SaveUser(user User) error {
	return es.saveUser(user)
}

func TestLogicGetPetNames2(t *testing.T) {
	data := []struct {
		name     string
		getPets  func(string) ([]Pet, error)
		userId   string
		petNames []string
		errMsg   string
	}{
		{"case1", func(userId string) ([]Pet, error) { return []Pet{{Name: "Bubbles"}}, nil }, "1", []string{"Bubbles"}, ""},
		{"case2", func(userId string) ([]Pet, error) { return nil, errors.New("invalid id: 3") }, "3", nil, "invalid id: 3"},
	}

	l := Logic{}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			l.Entities = EntitiesStub{getPets: d.getPets}
			petNames, err := l.GetPetNames(d.userId)
			if diff := cmp.Diff(petNames, d.petNames); diff != "" {
				t.Error(diff)
			}

			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != d.errMsg {
				t.Errorf("Expected error `%s`, got `%s`", d.errMsg, errMsg)
			}
		})
	}
}
