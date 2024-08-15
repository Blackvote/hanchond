package evmos

import (
	"fmt"
	"os/exec"
)

func (e *Evmos) GetTransaction(txhash string) (string, error) {
	command := exec.Command( //nolint:gosec
		e.BinaryPath,
		"q",
		"tx",
		"--type=hash",
		txhash,
		"--home",
		e.HomeDir,
		"--node",
		fmt.Sprintf("http://localhost:%d", e.Ports.P26657),
	)
	out, err := command.CombinedOutput()
	return string(out), err
}