package shop

import (
	"errors"
)

type Shop struct {
	inventory map[string]int
}

func NewShop() *Shop {
	return &Shop{
		inventory: make(map[string]int),
	}
}

func (s *Shop) AddItem(item string, quantity int) {
	s.inventory[item] += quantity
}

func (s *Shop) GetItem(item string) (int, error) {
	quantity, ok := s.inventory[item]
	if !ok {
		return 0, errors.New("item not found")
	}
	return quantity, nil
}

func (s *Shop) DeleteItem(item string) error {
	_, ok := s.inventory[item]
	if !ok {
		return errors.New("item not found")
	}
	delete(s.inventory, item)
	return nil
}

func (s *Shop) DropInventory() {
	s.inventory = make(map[string]int)
}
