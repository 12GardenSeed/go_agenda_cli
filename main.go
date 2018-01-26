package main

import (
	"os"

	"github.com/12GardenSeed/Agenda/cmd"
	"github.com/12GardenSeed/Agenda/entity"
)

func main() {
	cmd.Execute()
}

func init() {
	// Create data files if needed
	os.Mkdir("data", 0777)
	os.OpenFile(entity.USER_PATH, os.O_CREATE, os.ModePerm)
	os.OpenFile(entity.MEETING_PATH, os.O_CREATE, os.ModePerm)
	os.OpenFile(cmd.CURUSER_PATH, os.O_CREATE, os.ModePerm)
}
