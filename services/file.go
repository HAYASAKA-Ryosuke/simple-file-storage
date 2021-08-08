package services

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strconv"

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

func CreateFile(fileForm multipart.File, fileName string) (bool, error) {
	fileId, err := models.CreateFile(fileName)
	if err != nil {
		fmt.Println("error", err)
		return false, errors.New("internal error")
	}
	writeFile, err := os.Create(strconv.FormatInt(fileId, 10))
	defer writeFile.Close()
	if err != nil {
		return false, errors.New("failed create file")
	}
	io.Copy(writeFile, fileForm)
	return true, nil
}
