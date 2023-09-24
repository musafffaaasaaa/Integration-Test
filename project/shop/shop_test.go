package shop

import (
	"testing"
)

func TestShopMethods(t *testing.T) {
	t.Run("AddItem", func(t *testing.T) {
		s := NewShop()
		s.AddItem("item1", 5)

		quantity, err := s.GetItem("item1")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if quantity != 5 {
			t.Errorf("Expected quantity 5, got %d", quantity)
		}
	})

	t.Run("GetItem", func(t *testing.T) {
		s := NewShop()
		s.AddItem("item1", 10)

		quantity, err := s.GetItem("item1")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if quantity != 10 {
			t.Errorf("Expected quantity 10, got %d", quantity)
		}

		_, err = s.GetItem("nonexistent")
		if err == nil {
			t.Errorf("Expected an error for a nonexistent item")
		}
	})

	t.Run("DeleteItem", func(t *testing.T) {
		s := NewShop()
		s.AddItem("item1", 5)

		err := s.DeleteItem("item1")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		_, err = s.GetItem("item1")
		if err == nil {
			t.Errorf("Expected an error for a deleted item")
		}

		err = s.DeleteItem("nonexistent")
		if err == nil {
			t.Errorf("Expected an error for deleting a nonexistent item")
		}
	})

	t.Run("DropInventory", func(t *testing.T) {
		s := NewShop()
		s.AddItem("item1", 5)

		s.DropInventory()

		_, err := s.GetItem("item1")
		if err == nil {
			t.Errorf("Expected an error after dropping the inventory")
		}

		_, err = s.GetItem("item2")
		if err == nil {
			t.Errorf("Expected an error after dropping the inventory")
		}
	})
}
