package filehandler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
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

// IsValidAlias ...
func (a AliasMap) IsValidAlias(alias string) bool {
	_, ok := a[alias]
	return ok
}

// CfmConfig represents a .cfm file. It has an AliasMap and an EditorConfig
type CfmConfig struct {
	Aliases AliasMap     `json:"aliases"`
	E       EditorConfig `json:"editor"`
}

// Check determines whether the file exists at the location the alias is mapped to
func (c *CfmConfig) Check(alias string) {
	if c.Aliases.IsValidAlias(alias) {
		if _, err := os.Stat(c.Aliases[alias]); os.IsNotExist(err) {
			fmt.Printf("File doesn't exist at %s for alias %s\n", c.Aliases[alias], alias)
			return
		}
		fmt.Printf("File exists at %s for alias %s\n", c.Aliases[alias], alias)
		return
	}
	fmt.Printf("Alias %s not saved to the configuration\n", alias)
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
func (c *CfmConfig) AddAlias(alias, filepath string, force bool) bool {
	// Check that the file is not already in the CFM configuration;
	if c.Aliases.IsValidAlias(alias) && !force {
		fmt.Printf("Alias %s is already mapped to %s\n", alias, c.Aliases[alias])
		return false
	}
	c.Aliases[alias] = filepath
	return true
}

// RemapAlias ...
func (c *CfmConfig) RemapAlias(alias, filepath string) bool {
	if !c.Aliases.IsValidAlias(alias) {
		fmt.Printf("Alias %s not saved to the configuration\n", alias)
		return false
		// TODO: Offer to add alias
	}
	c.Aliases[alias] = filepath
	return true
}

// RemoveAlias ...
func (c *CfmConfig) RemoveAlias(alias string, force bool) bool {
	if c.Aliases.IsValidAlias(alias) {
		if !force && !utils.ConfirmAction(fmt.Sprintf("Remove alias %s", alias)) {
			return false
		}
		delete(c.Aliases, alias)
		return true
	}
	if !force {
		fmt.Printf("%s does not exist\n", alias)
	}
	return false
}

// RenameAlias ...
func (c *CfmConfig) RenameAlias(oldAlias, newAlias string) bool {
	if c.Aliases.IsValidAlias(oldAlias) {
		filepath := c.Aliases[oldAlias]
		delete(c.Aliases, oldAlias)
		c.Aliases[newAlias] = filepath
		return true
	}
	fmt.Printf("Alias %s not saved to the configuration", oldAlias)
	return false
}

// MakeEditorCommand returns a string that will invoke the user's preferred
// editor to edit their config files

// EditConfigFile ...
func (c *CfmConfig) EditConfigFile(alias string) {
	// Check that the alias exists
	if c.Aliases.IsValidAlias(alias) {
		filepath := c.Aliases[alias]

		cmd := exec.Command(c.E.Command, (append([]string{filepath}, c.E.Flags...))...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
		return;
	}
	fmt.Printf("%s does not exist\n", alias)
}

// LoadDataFile ...
func LoadDataFile(filepath string) (c CfmConfig, e error) {
	b, e := ioutil.ReadFile(filepath)
	if e != nil {
		return
	}
	e = json.Unmarshal(b, &c)
	if e != nil {
		return
	}
	return
}

// SaveDataFile writes the data in c to the .cfm file
func SaveDataFile(filepath string, c CfmConfig) error {
	b, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filepath, b, 0644)
}
