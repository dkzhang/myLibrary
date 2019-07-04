package downloadFile

import (
	"io"
	"net/http"
	"os"
)

func HttpDownloadFile(url string, filePath string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}

	_, err = io.Copy(f, res.Body)
	return err
}
