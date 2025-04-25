package cmd

import "flag"

var (
	DirFlag *string

	TypeNameFlag *string
)

func init() {
	DirFlag = flag.String("d", ".",
		"set the dir flag to locate the Type that need to be parse;"+
			"the default value is '.';"+
			"if the destination dir is {your current path}/internal/dir,"+
			"that your input dir should be ./internal/dir")

	TypeNameFlag = flag.String("t", "", "set the type name that the Type need to be parse")

}
