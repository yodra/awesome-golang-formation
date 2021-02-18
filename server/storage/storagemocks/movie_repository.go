package storagemocks

import "github.com/yodra/awesome-golang-formation/server"

type MockRepository struct{}

func (repo *MockRepository) Save(_ server.Movie) error {
	return nil
}
