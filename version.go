package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

var autoprint bool
var GitVersion = "Unknown"

func main() {
	var packagename string
	var help bool
	var setvar bool
	flag.StringVar(&packagename, "package", "main", "Name of package")
	flag.BoolVar(&help, "help", false, "Print some help info")
	flag.BoolVar(&autoprint, "autoprint", false, "Add code to print version at startup")
	flag.BoolVar(&setvar, "setvar", false, "Set external variable GitVersion to the latest version")
	flag.Parse()
	fmt.Printf("(c)Jan Kåre Vatne 2024 (jkvatne@online.no)\n")
	if help {
		fmt.Printf("\n")
		fmt.Printf("This program assumes that git is installed and on the path\n")
		fmt.Printf("It generates a file named 'gitversion.go' containing a line like this\n")
		fmt.Printf("  const Version=\"17e4a89b-dirty\"\n")
		fmt.Printf("\n")
		fmt.Printf("In your program you can just use the constant, f.ex:")
		fmt.Printf("  fmt.Println(Version)\n")
		fmt.Printf("\n")
		fmt.Printf("If the package is different from 'main', use the '-package' argument.")
		fmt.Printf("$version -package myprogram\n")
		fmt.Printf("\n")
		fmt.Printf("Installation:\n")
		fmt.Printf("$go install github.com/jkvatne/version@latest\n")
		fmt.Printf("\n")
		fmt.Printf("In the Goland IDE, to run it automatically before building the exe file:\n")
		fmt.Printf("  Ctrl-Alt-S, Tools, External Tools, Click '+' \n")
		fmt.Printf("  Set name and description. Set Program to 'version.exe', and click OK\n")
		fmt.Printf("  Set working directory to '$ProjectFileDir$' \n")
		fmt.Printf("  Edit the run configuration(s)\n")
		fmt.Printf("  At the bottom, in the 'Before launch' list, click '+'\n")
		fmt.Printf("  Select 'Run external tool' and select the tool you added (version).")
		fmt.Printf("\n")
		return
	}

	cmd := exec.Command("git", "describe", "--abbrev=8", "--dirty", "--always", "--tags")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal("The git command failed (not inside a repo?)")
	}
	fmt.Printf("Generated \"gitversion.go\" for git revision %s\n", out)
	outfile, err := os.Create("gitversion.go")
	if err != nil {
		log.Fatal("Could not open file \"gitversion.go\"")
	}

	defer outfile.Close()
	out = out[:len(out)-1]
	fmt.Fprintf(outfile, "package %s\n\n", packagename)
	if autoprint {
		fmt.Fprintf(outfile, "import \"fmt\"\n\n")
	}
	if !setvar {
		fmt.Fprintf(outfile, "const Version=\"%s\"\n\n", out)
	}
	if setvar {
		fmt.Fprintf(outfile, "func init() {\n")
		fmt.Fprintf(outfile, "   GitVersion=\"%s\"\n", out)
		fmt.Fprintf(outfile, "}\n")
	} else if autoprint {
		fmt.Fprintf(outfile, "func init() {\n")
		fmt.Fprintf(outfile, "	fmt.Printf(\"Version %%s\\n\", Version)\n")
		fmt.Fprintf(outfile, "}\n")
	}
	err = outfile.Close()
	if err != nil {
		log.Fatal("Could not close file \"gitversion.go\"")
	}
}
