package noDB

import (
	"testing"
	"url-shortener/internal/database"	
)

func TestSet(t *testing.T) {

	db, _ := New()
	t.Run("no error", func(t *testing.T) {
		if err := db.Set("test", "test"); err != nil {
			t.Errorf("Set() = %v, want %v", err, nil)
		}
	})

	t.Run("some error", func(t *testing.T) {
		if err := db.Set("test", "test"); err == nil {
			t.Errorf("Set() = %v, want %v", err, database.UniqueError{})
		}
	})
}

func TestGet(t *testing.T) {
	db, _ := New()

	t.Run("no error", func(t *testing.T) {
		if _, err := db.Get("test"); err == nil {
			t.Errorf("Get() = %v, want %v", err, database.NotFoundError{})
		}
	})

	db.Set("test", "test")

	t.Run("no error", func(t *testing.T) {
		if _, err := db.Get("test"); err != nil {
			t.Errorf("Get() = %v, want %v", err, nil)
		}
	})
	
}

func TestNew(t *testing.T) {
	if _, err := New(); err != nil {
		t.Errorf("New() = %v, want %v", err, nil)
	}
}
