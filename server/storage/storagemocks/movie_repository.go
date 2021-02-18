package storagemocks

import (
	"context"
	"github.com/yodra/awesome-golang-formation/server"
)

type MockRepository struct{}

func (repo *MockRepository) Save(ctx context.Context, movie server.Movie) error {
	return nil
}
