package cmd

import (
	"strings"
)

type Flag struct {
	Name      string
	Shorthand string
	Default   string
	Usage     string
	Variable  string
}

var flags = map[string]*Flag{
	"root": {
		Name:      "root",
		Shorthand: "r",
		Default:   ".",
		Usage:     "The root of the project",
	},
	"directory": {
		Name:      "directory",
		Shorthand: "d",
		Default:   ".",
		Usage:     "Terraform working directory relative to root directory",
	},
	"version": {
		Name:      "version",
		Shorthand: "v",
		Default:   "",
		Usage:     "Terraform version to use",
	},
	"interactive": {
		Name:      "interactive",
		Shorthand: "i",
		Default:   "true",
		Usage:     "Run terraform in interactive TTY mode (set to false for non-interactive)",
	},
}
var shorthandToNameMap map[string]string

func initialize() {
	shorthandToNameMap = make(map[string]string)
	for k, v := range flags {
		shorthandToNameMap[v.Shorthand] = k
	}
}

func parsArgs(args []string) []string {
	var tfArgs []string

	index := 0
	for index < len(args) {
		arg := args[index]

		// Not a flag: pass straight through to terraform
		if !strings.HasPrefix(arg, "-") {
			tfArgs = append(tfArgs, arg)
			index++
			continue
		}

		// Long form: --name or --name=value
		if strings.HasPrefix(arg, "--") {
			name := strings.TrimPrefix(arg, "--")
			value := ""
			if eq := strings.IndexRune(name, '='); eq >= 0 {
				value = name[eq+1:]
				name = name[:eq]
			} else if index+1 < len(args) {
				value = args[index+1]
			}

			if flag, exists := flags[name]; exists {
				flag.Variable = value
				if strings.Contains(arg, "=") {
					index++
				} else {
					index += 2
				}
				continue
			}

			// Unknown long flag: forward to terraform
			tfArgs = append(tfArgs, arg)
			index++
			continue
		}

		// Short form: -x
		name := strings.TrimPrefix(arg, "-")
		if flagName, exists := shorthandToNameMap[name]; exists {
			value := ""
			if index+1 < len(args) {
				value = args[index+1]
			}
			flags[flagName].Variable = value
			index += 2
			continue
		}

		// Unknown short flag: forward to terraform
		tfArgs = append(tfArgs, arg)
		index++
	}

	return tfArgs
}

func countLeadingDashes(arg string) int {
	leadingDashes := 0
	for _, char := range arg {
		if char == '-' {
			leadingDashes++
		} else {
			break
		}
	}
	return leadingDashes
}
