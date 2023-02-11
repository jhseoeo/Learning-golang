package stub

type User struct{}
type Pet struct {
	Name string
}
type Person struct{}

type Entities interface {
	GetUser(id string) (User, error)
	GetPets(userId string) ([]Pet, error)
	GetChildren(userId string) ([]Person, error)
	GetFriends(userId string) ([]Person, error)
	SaveUser(user User) error
}

type Logic struct {
	Entities Entities
}

func (l Logic) GetPetNames(userId string) ([]string, error) {
	pets, err := l.Entities.GetPets(userId)
	if err != nil {
		return nil, err
	}

	out := make([]string, 0, len(pets))
	for _, p := range pets {
		out = append(out, p.Name)
	}

	return out, nil
}
