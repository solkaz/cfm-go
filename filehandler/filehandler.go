package filehandler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// EditorConfig holds the command and flags to execute when editing a file
type EditorConfig struct {
	Command string   `json:"command"`
	Flags   []string `json:"flags"`
}

// AliasMap maps an
type AliasMap map[string]string

type CfmConfig struct {
	Aliases AliasMap     `json:"aliases"`
	E       EditorConfig `json:"editor"`
}

// LoadDataFile ...
func LoadDataFile(filepath string) (c CfmConfig, e error) {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = json.Unmarshal(b, &c)
	if err != nil {
		fmt.Println("error: ", err.Error())
		return
	}
	return
}
