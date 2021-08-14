package models

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/HAYASAKA-Ryosuke/simple-file-storage/database"
)

type File struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	CreatedAt string `json:"createdat"`
	UpdatedAt string `json:"updatedat"`
}

func FetchFileMany(search, sort string, page, limit int) ([]*File, error) {
	whiteList := map[string]string{"Id": "id", "Title": "title", "createdAt": "createdat", "updatedAt": "updatedat"}
	db := database.GetDatabase()
	orderBy := "Id"
	ascOrDesc := "ASC"
	if sort[0] == '-' {
		whiteListValue, ok := whiteList[sort[1:]]
		if ok == false {
			return nil, errors.New("invalid sort")
		}
		orderBy = whiteListValue
		ascOrDesc = "DESC"
	} else {
		whiteListValue, ok := whiteList[sort]
		if ok == false {
			return nil, errors.New("invalid sort")
		}
		orderBy = whiteListValue
	}
	var sqlStr string
	if search != "" {
		sqlStr = `
	        SELECT
		*
		FROM
			file
		WHERE title LIKE %?%
		ORDER BY ?
		LIMIT ?, ?
		`
	} else {
		sqlStr = `
	        SELECT
		*
		FROM
			file
		ORDER BY ?
		LIMIT ?, ?
	        `
	}
	fmt.Println(sqlStr)

	var rows *sql.Rows
	var err error

	if search != "" {
		rows, err = db.Query(sqlStr, search, orderBy+" "+ascOrDesc, (page-1)*limit, (page-1)*limit+limit)
	} else {
		rows, err = db.Query(sqlStr, orderBy+" "+ascOrDesc, (page-1)*limit, (page-1)*limit+limit)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var files []*File

	for rows.Next() {
		l := &File{}
		err := rows.Scan(
			&l.Id,
			&l.Title,
			&l.CreatedAt,
			&l.UpdatedAt,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		fmt.Println(l)
		files = append(files, l)
	}
	return files, nil
}

func FetchFileCount(search string) (int, error) {
	db := database.GetDatabase()
	searchColumn := ""
	var sqlStr string
	if search != "" {
		searchColumn = "WHERE title LIKE %?%"
		sqlStr = fmt.Sprintf(`
		SELECT
			COUNT(id)
		FROM
			file
		%s
		`, searchColumn)
	} else {
		sqlStr = fmt.Sprintf(`
		SELECT
			COUNT(id)
		FROM
			file
		`)
	}

	totalCount := 0
	row := db.QueryRow(sqlStr, search)
	err := row.Scan(&totalCount)
	if err != nil {
		return -1, err
	}
	return totalCount, nil
}

func CreateFile(fileName string) (int64, error) {
	db := database.GetDatabase()
	sqlStr := `INSERT INTO file(title) VALUES(?)`

	result, execErr := db.Exec(sqlStr, fileName)
	if execErr != nil {
		return -1, execErr
	}
	fileId, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return fileId, nil
}

func FetchFile(fileId int) (*File, error) {
	db := database.GetDatabase()
	sqlStr := `SELECT * FROM file WHERE id = ?`
	file := &File{}
	if err := db.QueryRow(sqlStr, "2").Scan(&file.Id, &file.Title, &file.CreatedAt, &file.UpdatedAt); err != nil {
		return nil, errors.New("file record not found")
	}
	return file, nil
}
