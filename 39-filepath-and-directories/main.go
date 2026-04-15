package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	path1 := filepath.Join("C", "Document", "Music")

	fmt.Println(path1)

	path2 := filepath.Join("C:", "config", "config.yaml")
	fmt.Println(path2)

	fmt.Println(filepath.Base(path2))
	fmt.Println(filepath.Dir(path2))
	fmt.Println(filepath.Ext(path2))

	badPath := "c:/./.bas/config/,/webhook.php"

	fmt.Println(filepath.Clean(badPath))

	fmt.Println("===========MAKING DIRECTORRIES==============")

	// dir := "downloads/assets"

	// if err := os.MkdirAll(filepath.Clean(dir), 0755); err != nil {
	// 	log.Fatal(err)
	// }

	//this is to remove all folder in the directory but u myst call the parent directory

	//Remove all remove folder and all its children
	//remove only the current folder

	if err := os.RemoveAll("downloads"); err != nil {
		log.Fatal(err)
	}

}
