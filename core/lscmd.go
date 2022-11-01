package core

import (
	"encoding/json"
	"fmt"
	"github.com/scmtble/gomodctl/parse"
	"github.com/spf13/cobra"
	"golang.org/x/mod/modfile"
	"golang.org/x/mod/module"
)

type CusRequire struct {
	Path     string `json:"path"`
	Version  string `json:"version"`
	Indirect bool   `json:"indirect"`
}

func NewCusRequire(version module.Version, indirect bool) *CusRequire {
	return &CusRequire{
		Path:     version.Path,
		Version:  version.Version,
		Indirect: indirect,
	}
}

func splitReqs(reqs []*modfile.Require) ([]*modfile.Require, []*modfile.Require) {
	var indirect []*modfile.Require
	var direct []*modfile.Require

	for _, v := range reqs {
		if v.Indirect {
			indirect = append(indirect, v)
			continue
		}
		direct = append(direct, v)
	}

	return direct, indirect
}

func formatPrint(reqs []*modfile.Require, lsAllFlag, jsonFlag bool) {
	direct, indirect := splitReqs(reqs)
	if jsonFlag {
		if lsAllFlag {
			formatJsonPrint(reqs)
			return
		}
		formatJsonPrint(direct)
		return
	}
	if len(direct) > 0 {
		fmt.Println("Direct: ")
		formatTextPrint(direct)
	}
	if lsAllFlag {
		if len(direct) > 0 {
			fmt.Println("Indirect: ")
			formatTextPrint(indirect)
		}
	}
}

func formatTextPrint(reqs []*modfile.Require) {
	for _, v := range reqs {
		fmt.Printf("\t%s@%s\n", v.Mod.Path, v.Mod.Version)
	}
}

func formatJsonPrint(reqs []*modfile.Require) {
	var arr []*CusRequire

	for _, v := range reqs {
		arr = append(arr, NewCusRequire(v.Mod, v.Indirect))
	}
	data, err := json.Marshal(arr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}

func NewLsCmd() *cobra.Command {
	var lsAllFlag bool
	var jsonFlag bool
	lsCmd := &cobra.Command{
		Use: "ls",
		Run: func(cmd *cobra.Command, args []string) {
			parse.MustParse()
			formatPrint(parse.ModFile.Require, lsAllFlag, jsonFlag)
		},
		Short: "Print dependencies",
	}
	flags := lsCmd.Flags()
	flags.BoolVarP(&lsAllFlag, "all", "a", false, "Print format all dependencies (include indirect dependencies)")
	flags.BoolVarP(&jsonFlag, "json", "j", false, "Print in json format All dependencies")
	return lsCmd
}
