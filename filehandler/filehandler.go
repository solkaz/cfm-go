package filehandler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/solkaz/cfm-go/utils"
)

// EditorConfig holds the command and flags to execute when editing a file
type EditorConfig struct {
	Command string   `json:"command"`
	Flags   []string `json:"flags"`
}

// AliasMap maps an
type AliasMap map[string]string

func (a AliasMap) IsValidAlias(alias string) bool {
	_, ok := a[alias]
	return ok
}

// CfmConfig represents a .cfm file. It has an AliasMap and an EditorConfig
type CfmConfig struct {
	Aliases AliasMap     `json:"aliases"`
	E       EditorConfig `json:"editor"`
}

// ListAliases ...
func (c *CfmConfig) ListAliases(aliases *[]string) {
	// Print all aliases if there were no specified aliases
	if len(*aliases) == 0 {
		for alias, src := range c.Aliases {
			fmt.Println(utils.MapAliasString(alias, src))
		}
	} else {
		// Check that each alias exists within the AliasMapx
		for _, alias := range *aliases {
			if c.Aliases.IsValidAlias(alias) {
				fmt.Println(utils.MapAliasString(alias, c.Aliases[alias]))
			} else {
				// TODO: Allow user to add a new alias
				fmt.Printf("Alias %s not valid\n", alias)
			}
		}
	}
}

// SearchAliases ...
func (c *CfmConfig) SearchAliases(phrase string) {
	var matchedAliases []string
	//
	for alias := range c.Aliases {
		if strings.Contains(alias, phrase) {
			matchedAliases = append(matchedAliases, alias)
		}
	}
	if len(matchedAliases) != 0 {
		for _, alias := range matchedAliases {
			fmt.Println(utils.MapAliasString(alias, c.Aliases[alias]))
		}
	} else {
		fmt.Printf("No aliases found for %s\n", phrase)
	}
}

// AddAlias ...
func (c *CfmConfig) AddAlias(alias, filepath string) {
	// Check that the file is not already in the CFM configuration;
	if c.Aliases.IsValidAlias(alias) {
		fmt.Printf("Alias %s is already mapped to %s\n", alias, c.Aliases[alias])
		return
	}
	c.Aliases[alias] = filepath
}

// MakeEditorCommand returns a string that will invoke the user's preferred
// editor to edit their config files
func (c *CfmConfig) MakeEditorCommand() string {
	return c.E.Command + strings.Join(c.E.Flags, " ")
}

// LoadDataFile ...
func LoadDataFile(filepath string) (c CfmConfig, e error) {
	b, e := ioutil.ReadFile(filepath)
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	e = json.Unmarshal(b, &c)
	if e != nil {
		fmt.Println("error: ", e.Error())
		return
	}
	return
}

// SaveDataFile writes the data in c to the .cfm file
func SaveDataFile(filepath string, c CfmConfig) error {
	b, err := json.Marshal(c)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filepath, b, 0644)
}
