package model

import (
	"fmt"
	"strconv"
)

type Pokemon struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

func (p *Pokemon) ToStringArr() []string {
	return []string{strconv.Itoa(int(p.ID)), p.Name}
}

func (p *Pokemon) ToStr() string {
	return fmt.Sprintf("%v,%v", p.ID, p.Name)
}
