package cmd

import "flag"

var (
	DirFlag *string

	TypeNameFlag *string
	//FileFlag     = flag.String("f", "", "set the file flag to locate the Type that need be parse")
)

func init() {
	DirFlag = flag.String("d", ".",
		"set the dir flag to locate the Type that need be parse;"+
			"if the destination dir is {your current path}/internal/dir,"+
			"that your input dir should be ./internal/dir")

	TypeNameFlag = flag.String("t", "", "set the type name that the Type need be parse")

}
