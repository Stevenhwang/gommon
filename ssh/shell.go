package ssh

import (
	gossh "golang.org/x/crypto/ssh"
)

// SSHExec 执行 ssh shell 远程命令
func SSHExec(cmd string, client *gossh.Client) (string, error) {
	// defer client.Close()
	// 创建session
	session, err := client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()
	// 执行远程命令
	combo, err := session.CombinedOutput(cmd)
	if err != nil {
		return "", err
	}
	return string(combo), nil
}
