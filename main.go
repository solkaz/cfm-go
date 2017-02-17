package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/solkaz/cfm-go/filehandler"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("cfm", "A config file manager")

	// List
	list        = app.Command("list", "list multiple (or all) alias' mappings").Alias("ls")
	listAliases = list.Arg("aliases", "aliases to list").Strings()

	// Search
	search      = app.Command("search", "search all aliases in the .cfm file")
	searchAlias = search.Arg("alias", "alias to search for").Required().String()

	// Add
	add         = app.Command("add", "add an alias to the .cfm file")
	addAlias    = add.Arg("alias", "alias for the config file to add").Required().String()
	addFilePath = add.Arg("filepath", "file path of the config file").Required().String()

	// Remove
	remove      = app.Command("remove", "remove an alias").Alias("rm")
	removeAlias = remove.Arg("alias", "alias to remove").Required().String()
	removeForce = remove.Flag("force", "do not prompt for removal").Bool()

	// Remap
	remap         = app.Command("remap", "remap the file an alias points to")
	remapAlias    = remap.Arg("alias", "alias to remap").Required().String()
	remapFilePath = remap.Arg("filepath", "new file path for inputted alias").Required().String()

	// Check
	check      = app.Command("check", "check that the file an alias points to exists")
	checkAlias = check.Arg("alias", "alias to check").Required().String()

	// Rename
	rename         = app.Command("rename", "rename an alias").Alias("mv")
	renameOldAlias = rename.Arg("old_alias", "alias to change").Required().String()
	renameNewAlias = rename.Arg("new_name", "new alias").Required().String()

	// Edit
	edit      = app.Command("edit", "edit a config file (by referring to its alias)")
	editAlias = edit.Arg("alias", "alias of the file to edit").Required().String()
)

func main() {
	currentUser, err := user.Current()
	if err != nil {
		panic(err)
	}
	cfmFilePath := currentUser.HomeDir + "/.cfm"

	c, err := filehandler.LoadDataFile(cfmFilePath)
	if err != nil {
		panic(err)
	}

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case list.FullCommand():
		c.ListAliases(listAliases)

	case search.FullCommand():
		c.SearchAliases(*searchAlias)

	case add.FullCommand():
		c.AddAlias(*addAlias, *addFilePath)
		err = filehandler.SaveDataFile(cfmFilePath, c)
		if err != nil {
			fmt.Println(err)
		}

	case remove.FullCommand():
		if !(*removeForce) {
			fmt.Printf("Remove %q?\n", *removeAlias)
		}
		fmt.Printf("Removed %q\n", *removeAlias)

	case remap.FullCommand():
		fmt.Printf("Remapping %q to %q\n", *remapAlias, *remapFilePath)

	case check.FullCommand():
		fmt.Printf("Checking %q\n", *checkAlias)

	case rename.FullCommand():
		fmt.Printf("Renaming %q to %q\n", *renameOldAlias, *renameNewAlias)

	case edit.FullCommand():
		fmt.Printf("Editing %q\n", *editAlias)

	}
}
