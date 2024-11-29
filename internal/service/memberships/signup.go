package memberships

import (
	"errors"

	"github.com/joshuatheokurniawansiregar/music-catalog/internal/model/memberships"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (s *service) SignUp(request memberships.SignupRequest) error {
	user, err := s.membershipsRepo.GetUser(request.Email, request.Username, 0)
	if err != nil && err != gorm.ErrRecordNotFound{
		log.Error().Err(err).Msg("error get user from database")
		return err
	}

	if user != nil{
		return errors.New("email or username exists")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil{
		return err
	}

	model  := memberships.User{
		Email: request.Email,
		Username: request.Username,
		Password: string(password),
		CreatedBy: request.Email,
		UpdatedBy: request.Email,
	}
	
	err = s.membershipsRepo.CreateUser(model)
	if err != nil{
		return err
	}
	return nil
}
