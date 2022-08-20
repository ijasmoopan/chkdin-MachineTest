package user

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/ijasmoopan/chkdin-MachineTest/models"
)

type Model struct {
	DB *sql.DB
}
func (user Model) DBAddUser(usr models.User) error {
	sqlString := `INSERT INTO users(
						email, 
						password) 
					VALUES (
						$1,
						$2);`

	_, err := user.DB.Exec(sqlString, usr.Email, usr.Password)
	if err != nil {
		log.Println("Error happened while adding new user", err)
		return err
	}
	return nil
}

func (user Model) DBValidateEmail(email string) (bool, error) {
	var dbUsername string
	sqlString := `SELECT email
					 FROM users
					 WHERE email = $1;`

	row := user.DB.QueryRow(sqlString, email)
	err := row.Scan(&dbUsername)
	if err == sql.ErrNoRows {
		return true, nil
	} else {
		return false, err
	}
}

func (user Model) DBValidateUser(userLogin models.User)(error){
	var userValid models.User
	sqlString := `SELECT id, 
						email,
						password
					FROM users
						WHERE email = $1
							AND password = $2;`

	row := user.DB.QueryRow(sqlString, userLogin.Email, userLogin.Password)
	err := row.Scan(&userValid.ID,
					&userValid.Email,
					&userValid.Password)
	if err != nil {
		return err
	}
	return  nil
}

func (user Model) DBGetUser(email string)(models.UserResponse, error){
	var usr models.UserResponse
	sqlString := `SELECT id,
						email
					FROM users
						WHERE email = $1;`

	row := user.DB.QueryRow(sqlString, email)
	err := row.Scan(&usr.ID,
					&usr.Email)
	if err == sql.ErrNoRows {
		return usr, errors.New("E-mail is not valid")
	} else if err != nil {
		return usr, err
	}
	return usr, nil
}

func (user Model) DBPatchUser(patchUser models.User)(error){
	sqlString := `UPDATE users SET `
	var count int = 1
	var param []interface{}

	if patchUser.Email != "" && patchUser.Password != "" {
		param = append(param, patchUser.Email, patchUser.Password)
		sqlString = fmt.Sprint(sqlString, ` email = $`, count, ` , password = $`, count+1)
		count += 2
	} else if patchUser.Email != "" {
		param = append(param, patchUser.Email)
		sqlString = fmt.Sprint(sqlString, ` email = $`, count)
		count++
	} else if patchUser.Password != "" {
		param = append(param, patchUser.Password)
		sqlString = fmt.Sprint(sqlString, ` password = $`, count)
		count++
	}
	param = append(param, patchUser.ID)
	sqlString = fmt.Sprint(sqlString, ` WHERE id = $`, count)

	_, err := user.DB.Exec(sqlString, param...)
	if err != nil {
		log.Println("Error:", err)
		return err
	}
	return nil
}

func (user Model) DBDeleteUser(email string)(error){
	sqlString := `DELETE 
					FROM users
					WHERE email = $1;`
	_, err := user.DB.Exec(sqlString, email)
	if err != nil {
		return err
	}
	return nil
}