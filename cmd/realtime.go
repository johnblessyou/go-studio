package cmd

import (
	"bufio"
	"fmt"
	"os/exec"
)

// 执行shell命令并获取实时输出结果
func execScript(command string) error {
	cmd := exec.Command("/bin/bash", command, ">&2")
	stderr, _ := cmd.StderrPipe()
	if err := cmd.Start(); err != nil {
		return err
	}
	scanner := bufio.NewScanner(stderr)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	return cmd.Wait()
}
