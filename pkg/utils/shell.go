package utils

import (
	"os"
	"os/exec"
)

// taken from kairos
func SH(c string) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", c)
	cmd.Env = os.Environ()
	o, err := cmd.CombinedOutput()
	return string(o), err
)