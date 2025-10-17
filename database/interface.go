package database

import (
	sqlite "github.com/Lazy-Parser/Server/database/sqlite"
	"github.com/Lazy-Parser/Server/entity"
	"gorm.io/gorm"
)

type UserRepo interface {
	FindById(id uint) (entity.User, error)
	FindByUsername(username string) (entity.User, error)
	// Update - find user by ID and update all info from [update].
	//
	// Returns: (Error, RowsEffected)
	Update(id uint, update map[string]interface{}) (error, int64)
	Create(entity.User) error
}

func CreateUserRepo(db *gorm.DB) UserRepo {
	return sqlite.CreateUserRepo(db)
}

type RoleRepo interface {
	FindById(id uint) (entity.Role, error)
	FindByName(name string) (entity.Role, error)
	Create(entity.Role) error
}
