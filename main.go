package main

import (
	"flag"
	"fmt"
	"os"
)

// type cmdArgs struct {
// 	workspace string
// 	variable  string
// 	value     string
// }

func create() {
	createCmd := flag.NewFlagSet("create", flag.ExitOnError)
	workspacePtr := createCmd.String("w", "", "TFC workspace name")
	namePtr := createCmd.String("n", "", "TFC workspace variable name")
	valuePtr := createCmd.String("v", "", "TFC workspace variable value")
	createCmd.Parse(os.Args[2:])
	output := "workspace: " + *workspacePtr + "\n" +
		"variable: " + *namePtr + "\n" +
		"value: " + *valuePtr + "\n"
	fmt.Print(output)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("missing command")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "create":
		create()
	default:
		fmt.Println("invalid command")
	}
}
