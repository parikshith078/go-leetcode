package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"pari/leetcode/questions"
)

func main() {
	add := flag.Bool("add", false, "Add new leetcode question to the db with: title, link, discription in order")
	list := flag.Bool("list", false, "List all the leetcode questions in the db")
	flag.Parse()

	// filePath := os.Getenv("GO_DEV_DBS") + "leetcode-questions.json"
	filePath := "/Users/parikshith/Craft/projects/go-leetcode/db.json"

	data := questions.DatabaseModel{}
	if err := data.Load(filePath); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		args := flag.Args()
		if len(args) != 3 {
			fmt.Fprintln(os.Stderr, errors.New("Invalid args! 3 agrs are needed: title link discription").Error())
			os.Exit(1)
		}
		title, link, disc := args[0], args[1], args[2]

		if err := data.Add(title, link, disc); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		if err := data.Store(filePath); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	case *list:
		data.List()
	}
}
