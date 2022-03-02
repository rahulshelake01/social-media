package repository

import (
	"database/sql"
	"fmt"
	"social-media/config"
	"social-media/models"
	"strings"
)

type SocialMediaRepositoryInterface interface {
	AddPost(models.SocialMediaAddPostApiRequest) (int64, error)
}

type SocialMediaRepositoryStruct struct {
	DB     *sql.DB
	Config config.AppConfig
}

func SocialMediaRepository(db *sql.DB, Config config.AppConfig) SocialMediaRepositoryInterface {
	return SocialMediaRepositoryStruct{DB: db, Config: Config}
}

func (smrp SocialMediaRepositoryStruct) AddPost(data models.SocialMediaAddPostApiRequest) (int64, error) {

	var columns []string
	var values []interface{}

	if data.Title != "" {
		columns = append(columns, "title")
		values = append(values, data.Title)
	}

	if data.Description != "" {
		columns = append(columns, "description")
		values = append(values, data.Description)
	}

	if data.Category != 0 {
		columns = append(columns, "category")
		values = append(values, data.Category)
	}

	if data.UserID != 0 {
		columns = append(columns, "userid")
		values = append(values, data.UserID)
	}

	statement, err := smrp.DB.Prepare(`INSERT INTO post ( ` + strings.Join(columns, ",") + `) VALUES (` + strings.Repeat("?", len(columns)) + `)`)

	if err != nil {
		fmt.Println("Error while prepare statement : ", err)
		return 0, err
	}

	defer statement.Close()

	res, err := statement.Exec(values...)

	if err != nil {
		fmt.Println("Error while execute statement : ", err)
		return 0, err
	}

	lastInsertedId, _ := res.LastInsertId()

	return lastInsertedId, nil
}
