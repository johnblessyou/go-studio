package cmd

import (
	"io"
	"os"
	"os/exec"

	"github.com/creack/pty"
)

// 在pty中执行shell命令并获取实时输出结果
func execScriptToPty(command string) (err error) {
	cmd := exec.Command("/bin/bash", command)
	ptmx, err := pty.Start(cmd)
	if err != nil {
		return err
	}
	_, err = io.Copy(os.Stdout, ptmx)
	return err
}
