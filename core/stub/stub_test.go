package stub_test

import (
	"fmt"
	"testing"

	"github.com/aliyev-vladislav/learning-go/core/stub"
	"github.com/google/go-cmp/cmp"
)

type GetPetNamesStub struct {
	stub.Entities
}

func (ps GetPetNamesStub) GetPets(userID string) ([]stub.Pet, error) {
	switch userID {
	case "1":
		return []stub.Pet{{Name: "Bubbles"}}, nil
	case "2":
		return []stub.Pet{{Name: "Stampy"}, {Name: "Shoball II"}}, nil
	default:
		return nil, fmt.Errorf("invalid id: %s", userID)
	}
}

func TestLogicGetPerNames(t *testing.T) {
	data := []struct {
		name     string
		userID   string
		petNames []string
	}{
		{"case1", "1", []string{"Bubbles"}},
		{"case2", "2", []string{"Stampy", "Showball II"}},
		{"case3", "3", nil},
	}
	l := stub.Logic{GetPetNamesStub{}}
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
