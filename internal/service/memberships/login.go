package memberships

import (
	"errors"

	"github.com/joshuatheokurniawansiregar/music-catalog/internal/model/memberships"
	"github.com/joshuatheokurniawansiregar/music-catalog/pkg/jwt"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(request memberships.LoginRequest)(string, error){
	userDetail, err := s.membershipsRepo.GetUser(request.Email, "",0)
	if err != nil{
		log.Error().Err(err).Msg("error get user from database")
		return "",err
	}

	if userDetail == nil{
		return "", errors.New("email does not exists")
	}
	// hashedPassword, err:= bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	// if err != nil{
	// 	return "", err
	// }
	err = bcrypt.CompareHashAndPassword([]byte(userDetail.Password), []byte(request.Password))
	if err != nil{
		return "", err
	}

	accessToken, err :=  jwt.CreateToken(int64(userDetail.ID), userDetail.Username, s.cfg.Service.SecretJWT)
	if err != nil{
		log.Error().Err(err).Msg("failed to create JWT Token")
		return "", nil
	}
	return accessToken,nil
}