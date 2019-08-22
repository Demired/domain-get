package pull

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

var (
	baseURL = "https://domain-pull.com/api/%s/%s/%s"
)

//Pull func
func Pull(token, version, date, path string) error {
	var url = fmt.Sprintf(baseURL, token, version, date)
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return errors.New("asd")
		// return errors.New(res.Body)
	}
	file, err := os.Create(fmt.Sprintf("%s/%s.txt", path, time.Now().Format("20060102")))
	if err != nil {
		return err
	}
	_, err = io.Copy(file, res.Body)
	if err != nil {
		return err
	}
	return nil
}
