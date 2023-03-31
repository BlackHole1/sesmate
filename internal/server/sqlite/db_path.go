package sqlite

import (
	"log"
	"os"
	"path"
)

func dbPath() string {
	config, err := os.UserConfigDir()
	if err != nil {
		log.Fatalln(err.Error())
	}

	dir := path.Join(config, "sesmate")
	if err = os.MkdirAll(dir, os.ModePerm); err != nil {
		log.Fatalln(err.Error())
	}

	file := path.Join(dir, "data.db")

	fileInfo, err := os.Stat(file)
	if os.IsNotExist(err) {
		if _, err = os.Create(file); err != nil {
			log.Fatalln(err.Error())
		}
	} else if err != nil {
		log.Fatalln(err.Error())
	} else if fileInfo.IsDir() {
		log.Fatalln("data.db is a directory")
	}

	return file
}
