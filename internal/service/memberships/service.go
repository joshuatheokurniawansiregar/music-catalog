package memberships

import (
	"github.com/joshuatheokurniawansiregar/music-catalog/internal/configs"
	"github.com/joshuatheokurniawansiregar/music-catalog/internal/model/memberships"
)

//go:generate mockgen -source=service.go -destination=service_mock_test.go -package=memberships
type repository interface {
	CreateUser(model memberships.User) error
	GetUser(email, username string, id uint64)(*memberships.User, error)
}

type service struct{
	membershipsRepo repository
	cfg *configs.Config
}

func NewService(membershipsRepo repository, cfg *configs.Config)*service{
	return &service{
		membershipsRepo: membershipsRepo,
		cfg: cfg,
	}
}
