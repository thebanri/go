package monster

import (
	"fmt"
	"math/rand/v2"
)

type Monster struct {
	ID   int
	Name string
}

type Slice struct {
	Monster []Monster
}

func CreateMonster(names ...string) *Monster {
	monster := Monster{ID: rand.IntN(1000)}
	name := "Noname Monster"
	if len(names) > 0 && names[0] != "" {
		name = names[0]
	}
	monster.Name = name
	return &monster
}

type UpdateRequest struct {
	ID      int
	NewName string
}

func (s *Slice) UpdateMonster(req UpdateRequest) {

	if req.NewName != "" && req.ID != 0 {
		if req.ID != 0 {
			for i := range s.Monster {
				if s.Monster[i].ID == req.ID {
					s.Monster[i].Name = req.NewName
					fmt.Printf("ID %d olan canavarın ismi %s olarak güncellendi\n", req.ID, req.NewName)
					return
				}
			}
		}
	} else {
		fmt.Println("Lütfen verileri eksiksiz girin!")
	}
}
