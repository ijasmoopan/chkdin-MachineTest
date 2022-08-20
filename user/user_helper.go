package user

import (
	"database/sql"

	"github.com/ijasmoopan/chkdin-MachineTest/models"
)

type Repo struct {
	user interface {
		DBAddUser(models.User)(error)
		DBValidateEmail(string)(bool, error)
		DBValidateUser(models.User)(error)
		DBGetUser(string)(models.UserResponse, error)
		DBPatchUser(models.User)(error)
		DBDeleteUser(string)(error)
	}
}

func InterfaceHandler(db *sql.DB) *Repo {
	return &Repo {
		user: Model{DB: db},
	}
}