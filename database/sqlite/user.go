package sqlite

import (
	"github.com/Lazy-Parser/Server/entity"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func CreateUserRepo(db *gorm.DB) *userRepo {
	return &userRepo{db: db}
}

func (repo *userRepo) FindById(id uint) (entity.User, error) {
	var user entity.User
	if err := repo.db.Where("id = ?", id).First(&user).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (repo *userRepo) FindByUsername(username string) (entity.User, error) {
	var user entity.User
	if err := repo.db.Where("username = ?", username).First(&user).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (repo *userRepo) Update(id uint, update map[string]interface{}) (error, int64) {
	res := repo.db.Model(&entity.User{}).Where("id = ?", id).Updates(update)
	return res.Error, res.RowsAffected
}

func (repo *userRepo) Create(user entity.User) error {
	return repo.db.Create(&user).Error
}
