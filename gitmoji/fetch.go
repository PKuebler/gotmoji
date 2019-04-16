package gitmoji

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkuebler/gotmoji/models"
)

const cachePath string = "./gitmojis.json"

func FetchRemoteEmojis() (error, *models.Directory) {
	resp, err := http.Get("https://raw.githubusercontent.com/carloscuesta/gitmoji/master/src/data/gitmojis.json")
	if err != nil {
		return err, nil
	}
	defer resp.Body.Close()

	directory := &models.Directory{}
	err = json.NewDecoder(resp.Body).Decode(directory)

	return err, directory
}

func FetchCachedEmojis() (error, *models.Directory) {
	data, err := ioutil.ReadFile(cachePath)
	if err != nil {
		return err, nil
	}

	directory := &models.Directory{}
	err = json.Unmarshal(data, &directory)

	return err, directory
}

func CacheEmojis(directory *models.Directory) error {
	directoryJson, _ := json.Marshal(*directory)
	err := ioutil.WriteFile(cachePath, directoryJson, 0644)

	return err
}

func FetchEmojis() *models.Directory {
	err, directory := FetchCachedEmojis()

	if err == nil {
		fmt.Println("Use cached emoji directory")
		return directory
	}

	fmt.Println("Try to download new emoji list...")
	err, directory = FetchRemoteEmojis()
	if err != nil {
		return nil
	}

	fmt.Println("Save new emoji list...")
	err = CacheEmojis(directory)
	if err != nil {
		return nil
	}

	return directory
}
