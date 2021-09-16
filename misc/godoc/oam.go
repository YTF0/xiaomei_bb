package godoc

import "github.com/lovego/cmd"

func shell() error {
	_, err := cmd.Run(cmd.O{}, `docker`, `exec`, `-it`,
		`--detach-keys=ctrl-@`, `workspace-godoc`, `bash`,
	)
	return err
}

func ps() error {
	_, err := cmd.Run(cmd.O{}, `docker`, `ps`, `-f`, `name=workspace-godoc`)
	return err
}
