package filehandler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

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
