package services

import (
	"errors"
	"fmt"

	"github.com/HAYASAKA-Ryosuke/simple-file-storage/models"
)

func FetchFiles(search string, sort string, page, limit int) ([]*models.File, int, error) {
	fileTotalCount, err := models.FetchFileCount(search)
	if err != nil {
		fmt.Println("error", err)
		return nil, -1, errors.New("internal error")
	}
	files, err := models.FetchFileMany(search, sort, page, limit)
	if err != nil {
		fmt.Println("error", err)
		return nil, -1, errors.New("internal error")
	}
	return files, fileTotalCount, nil
}
