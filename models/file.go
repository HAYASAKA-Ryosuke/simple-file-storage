package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/HAYASAKA-Ryosuke/simple-file-storage/database"
)

type File struct {
	Id        int32  `json:"Id"`
	Title     string `json:"title"`
	CreatedAt string `json:"createdat"`
	UpdatedAt string `json:"updatedat"`
}

func FetchFileMany(search, sort string, page, limit int) ([]*File, error) {
	whiteList := map[string]string{"Id": "id", "title": "Title", "createdAt": "createdat", "updatedAt": "updatedat"}
	db := database.GetDatabase()
	orderBy := ""
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
			id,
			title,
			createdat
			updatedat
		FROM
			file
		WHERE title LIKE %?%
		ORDER BY ? ?
		LIMIT ?, ?
		`
	} else {
		sqlStr = `
	        SELECT
			id,
			title,
			createdat
			updatedat
		FROM
			file
		ORDER BY ? ?
		LIMIT ?, ?
	        `
	}

	var rows *sql.Rows
	var err error

	if search != "" {
		rows, err = db.Query(sqlStr, search, orderBy, ascOrDesc, page*limit, page*limit+limit)
	} else {
		rows, err = db.Query(sqlStr, orderBy, ascOrDesc, page*limit, page*limit+limit)
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
			return nil, err
		}
		files = append(files, l)
	}
	return files, nil
}

func FetchFileCount(search string) (int, error) {
	db := database.GetDatabase()
	searchColumn := ""
	if search != "" {
		searchColumn = "WHERE title LIKE %?%"
	}

	sqlStr := fmt.Sprintf(`
	SELECT
		COUNT(id)
	FROM
		file
	%s
	`, searchColumn)

	totalCount := 0
	row := db.QueryRow(sqlStr, search)
	err := row.Scan(&totalCount)
	if err != nil {
		return -1, err
	}
	return totalCount, nil
}

func CreateFile(fileName string) (int64, error) {
	var ctx context.Context
	db := database.GetDatabase()
	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return -1, err
	}
	sqlStr := `INSERT INTO file(title) VALUES(?)`

	result, execErr := db.ExecContext(ctx, sqlStr, fileName)
	if execErr != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return -1, rollbackErr
		}
		return -1, execErr
	}
	if err := tx.Commit(); err != nil {
		return -1, err
	}
	fileId, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return fileId, nil
}
