package repository

import (
	"database/sql"
	"fmt"
	"social-media/config"
	"social-media/models"
	"strings"
)

type UserRepositoryInterface interface {
	Register(request *models.UserRegisterRequest) error
	FindUserByEmail(email string) (*models.UserDetails, bool, error)
}

type UserRepositoryStruct struct {
	DB     *sql.DB
	Config config.AppConfig
}

func UserRepository(DB *sql.DB, Config config.AppConfig) UserRepositoryInterface {
	return UserRepositoryStruct{DB, Config}
}

func (userRepo UserRepositoryStruct) Register(request *models.UserRegisterRequest) error {

	var (
		base        = " INSERT users SET "
		set         []string
		values      []interface{}
		insertQuery string
	)

	set = append(set, "email = ?")
	values = append(values, request.Email)

	set = append(set, "pass = ?")
	values = append(values, request.Password)

	insertQuery = strings.Trim(base+strings.Join(set, ","), ",")

	stmt, err := userRepo.DB.Prepare(insertQuery)

	if err != nil {
		fmt.Println("Failed to prepare insert register user query --> ", err)
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(values...)

	if err != nil {
		fmt.Println("Failed to execute insert register user query --> ", err)
		return err
	}

	return nil
}

func (userRepo UserRepositoryStruct) FindUserByEmail(email string) (*models.UserDetails, bool, error) {

	fmt.Println("In repository.FindUserByEmail..... : ", email)

	var (
		userDetails = new(models.UserDetails)
		userQuery   = "SELECT id, email, pass from users where email = ?"
		isFound     = false
		err         error
	)

	err = userRepo.DB.QueryRow(userQuery, email).Scan(&userDetails.UID, &userDetails.Email, &userDetails.Password)

	if err == nil {
		isFound = true
	} else {
		if err == sql.ErrNoRows {
			err = nil
		} else {
			fmt.Println("Error to fetch user by email", err)
		}
	}

	return userDetails, isFound, err
}
