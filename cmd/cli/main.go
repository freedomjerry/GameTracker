package main

import (
	"fmt"
	"github.com/Gametracker"
	"log"
	"os"
)

const dbFileName = "game.db.json"
func main()  {

	store, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")
	poker.NewCLI(store, os.Stdin).PlayPoker()
}
