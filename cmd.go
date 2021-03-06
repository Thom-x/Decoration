package main

import (
	"os/exec"
	"fmt"
	"github.com/fatih/color"
	"io"
	"os"
	"bufio"
	"github.com/mgutz/str"
)

func runCommand() error {
	return decorate(os.Stdin, os.Stdout)
}


func decorate(r io.Reader, w io.Writer) error {
	args := str.ToArgv(flags.args)
	cmd := exec.Command(flags.program, args...)
	
	stdout, e := cmd.StdoutPipe()
	if e != nil {
		return e
	}

	stderr, e := cmd.StderrPipe()
	if e != nil {
       	return e
    }	

	e = cmd.Start()
	if e != nil {
		return e
	}	

	merged := io.MultiReader(stdout, stderr)
	scanner := bufio.NewScanner(merged)

	for scanner.Scan() {
		if(flags.color != "") {
			switch(flags.color) {
				case "black":
				color.Set(color.FgBlack)
				case "red":
				color.Set(color.FgRed)
				case "green":
				color.Set(color.FgGreen)
				case "yellow":
				color.Set(color.FgYellow)
				case "blue":
				color.Set(color.FgBlue)
				case "magenta":
				color.Set(color.FgMagenta)
				case "cyan":
				color.Set(color.FgCyan)
				case "white":
				color.Set(color.FgWhite)
			}
		}
		_, e := fmt.Fprintln(w, flags.prefix + scanner.Text() + flags.suffix)
		if(flags.color != "") {
			color.Unset()
		}
		if e != nil {
			return e
		}		
	}
	return nil
}
