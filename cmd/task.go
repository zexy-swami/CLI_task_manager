package main

import (
	"github.com/zexy-swami/CLI_task_manager/internal/commands"
	"github.com/zexy-swami/CLI_task_manager/internal/db"
	"log"
)

const dbPath = `..\db_path\task_list.db`

func main() {
	if err := db.InitDB(dbPath); err != nil {
		log.Fatalf("Error occurred while initing database. Error: %s\n", err.Error())
	}

	if err := commands.RootCommand.Execute(); err != nil {
		log.Fatalf("Error occurred while executing commands. Error: %s\n", err.Error())
	}
}
