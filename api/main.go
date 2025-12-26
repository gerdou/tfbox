package api

import "github.com/gerdou/tfbox/internal"

func Run(rootDirectory, workingDirectory, tfVersion string, tfArgs []string, showLogs bool) error {
	interactive := showLogs
	return internal.Run(rootDirectory, workingDirectory, tfVersion, tfArgs, interactive, showLogs)
}
