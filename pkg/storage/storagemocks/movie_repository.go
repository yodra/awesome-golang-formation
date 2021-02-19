package storagemocks

import (
	"context"
	"github.com/yodra/awesome-golang-formation/pkg/domain"
)

type MockRepository struct{}

func (repo *MockRepository) Save(ctx context.Context, movie domain.Movie) error {
	return nil
}
