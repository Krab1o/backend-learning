package db

import (
	"encoding/json"
	"log"
	"os"
)

type data struct {
	ID	uint	`json:"id"`
}

var d data

const directoryPath = "data/"
const dataPath = directoryPath + "data.json"
const idPath = directoryPath + "id.json"
const directoryPermissions = 0777
const dataPermissions = 0644

func InitDB() {
	_, err := os.Stat(directoryPath)
	if os.IsNotExist(err) {
		log.Println("Data directory doesn't exist. Creating...")
		createDirectory()
	} else {
		log.Println("Data directory already exists")
	}

	_, err = os.Stat(dataPath)
	if os.IsNotExist(err) {
		log.Println("Data file doensn't exist. Creating...")
		createFile(dataPath)
	} else {
		log.Println("Data file already exists")
	}
	
	_, err = os.Stat(idPath)
	if os.IsNotExist(err) {
		log.Println("Data file doensn't exist. Creating...")
		d.ID = 1
		bytes, err := json.Marshal(d)
		if err != nil {
			log.Println(err)
		}
		log.Println(string(bytes))
		createFile(idPath)
		err = os.WriteFile(idPath, bytes, os.FileMode(dataPermissions))
		if err != nil {
			log.Printf("DB: failed to save ID to file, %v", err)
		}
	} else {
		log.Println("Data file already exists")
		bytes, err := os.ReadFile(idPath)
		if err != nil {
			log.Printf("DB: failed to open ID from file, %v", err)
		}
		json.Unmarshal(bytes, &d)
	}
	
}

func createDirectory() {
	err := os.Mkdir(directoryPath, os.FileMode(directoryPermissions))
	if (err != nil) {
		log.Println(err)
	}
}

func createFile(path string) {
	file, err := os.Create(path)
	if (err != nil) {
		log.Println(err)
	}
	defer file.Close()
}

