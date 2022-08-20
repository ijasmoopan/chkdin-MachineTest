package user

import (
	"database/sql"

	"github.com/ijasmoopan/chkdin-MachineTest/models"
)

type Repo struct {
	user interface {
		DBAddUser(models.User)(error)
		DBValidateEmail(string)(bool, error)
		DBValidateUser(models.User)(models.UserResponse, error)
		DBGetUser(string)(models.UserResponse, error)
		DBPatchUser(models.User)(error)
		DBDeleteUser(string)(error)
		DBAuthUser(float64)(models.UserResponse, error)
	}
}

func InterfaceHandler(db *sql.DB) *Repo {
	return &Repo {
		user: Model{DB: db},
	}
}