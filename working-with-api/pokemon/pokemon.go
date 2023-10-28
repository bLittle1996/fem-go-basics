package pokemon

import "fmt"

type Pokemon struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (pokemon Pokemon) String() string {
	return fmt.Sprintf("I choose you, %s (%d)!", pokemon.Name, pokemon.Id)
}
