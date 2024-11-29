package memberships

import (
	"github.com/joshuatheokurniawansiregar/music-catalog/internal/model/memberships"
	"gorm.io/gorm"
)

func (r *repository) CreateUser(model memberships.User)error{
	return r.db.Create(&model).Error
}

func (r *repository) GetUser(email, username string, id uint64)(*memberships.User, error){
	user := memberships.User{}
	var tx *gorm.DB = r.db.Where("email = ?", email).Or("username = ?", username).Or("id = ?", id).First(&user)
	if tx.Error != nil{
		return nil, tx.Error
	}

	return &user, nil
}