package rest

import (
	"fmt"
	"testing"
)

// databse with no erros
type noErrorDB struct{}
func (r noErrorDB) Get(key string) (string, error) {
	return "test", nil
}
func (r noErrorDB) Set(key, value string) error {
	return nil
}
func (r noErrorDB) Close() {}

// database with undefined error
type someErrorDB struct{}
func (e someErrorDB) Get(key string) (string, error) {
	err := fmt.Errorf("some error")
	return "", err
}
func (e someErrorDB) Set(key, value string) error {
	return fmt.Errorf("some error")
}
func (e someErrorDB) Close() {}



func TestCreateToken(t *testing.T) {
	//todo

	t.Run("no errors", func(t *testing.T) {
		if _, err := CreateToken("test", noErrorDB{}); err != nil {
			t.Errorf("GenerateToken() = %v, want %v", err, nil)
		}
	})

	t.Run("some error", func(t *testing.T) {
		if _, err := CreateToken("test", someErrorDB{}); err == nil {
			t.Errorf("GenerateToken() = %v, want %v", err, fmt.Errorf("some error"))
		}
	})

}

func TestRetrieveLink(t *testing.T) {

	t.Run("no errors", func(t *testing.T) {
		if _, err := RetrieveLink("test", noErrorDB{}); err != nil {
			t.Errorf("GenerateToken() = %v, want %v", err, nil)
		}
	})

	t.Run("some error", func(t *testing.T) {
		if _, err := RetrieveLink("test", someErrorDB{}); err == nil {
			t.Errorf("GenerateToken() = %v, want %v", err, fmt.Errorf("some error"))
		}
	})
}
