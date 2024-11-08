package main

import (
	"fmt"
	"log"
	"os"
	"flag"
	"os/exec"
)

func main() {
	var packagename string
	flag.StringVar(&packagename,"package", "main", "Name of package")
	flag.Parse()
	fmt.Printf("Generating a gitversion.go file for package %s\n",packagename)
	fmt.Printf("(c)Avanti Engineering 2018\n")
	cmd := exec.Command("git", "describe", "--abbrev=8", "--dirty", "--always", "--tags")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Generated \"gitversion.go\" for git revision %s", out)
	outfile, err := os.Create("gitversion.go")
	if err != nil {
		log.Fatal("Could not open file \"gitversion.go\"")
	}
	defer outfile.Close()
	out = out[:len(out)-1]
	fmt.Fprintf(outfile, "package %s\n\n", packagename)
	fmt.Fprintf(outfile, "// Version is extracted from the latest tag in the git repository\n")
	fmt.Fprintf(outfile, "const Version=\"%s\"\n", out)
	outfile.Close()
	fmt.Printf("Ok, exit now")
	os.Exit(0)
}
