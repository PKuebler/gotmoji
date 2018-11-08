package main

import (
	"github.com/pkuebler/gotmoji/cmd"
	flags "github.com/spf13/pflag"
)

type Menu struct {
	Items   []*MenuItem
	Current *MenuItem
}

type MenuItem struct {
	Flag   string
	Short  string
	Desc   string
	Active bool
	Run    func()
}

func (menu *Menu) init() {
	for _, item := range menu.Items {
		flags.BoolVarP(&item.Active, item.Flag, item.Short, false, item.Desc)
	}

	flags.Parse()

	for _, item := range menu.Items {
		if item.Active {
			menu.Current = item
			return
		}
	}
}

func main() {
	menu := Menu{
		Items: []*MenuItem{
			&MenuItem{Flag: "init", Short: "i", Desc: "Initialize gitmoji as a commit hook"},
			&MenuItem{Flag: "remove", Short: "r", Desc: "Remove a previously initialized commit hook"},
			&MenuItem{Flag: "commit", Short: "c", Desc: "Interactively commit using the prompts", Run: cmd.Commit},
			&MenuItem{Flag: "list", Short: "l", Desc: "List all the available gitmojis", Run: cmd.List},
			&MenuItem{Flag: "search", Short: "s", Desc: "Search gitmojis"},
			&MenuItem{Flag: "version", Short: "v", Desc: "Print gitmoji-cli installed version", Run: cmd.Version},
		},
	}

	menu.init()

	menu.Current.Run()
}
