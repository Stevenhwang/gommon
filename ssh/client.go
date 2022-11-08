package ssh

import (
	"fmt"
	"time"

	gossh "golang.org/x/crypto/ssh"
)

// 创建 ssh client, password 认证
func GetSSHClientByPassword(host string, port uint, user string, password string) (*gossh.Client, error) {
	addr := fmt.Sprintf("%s:%d", host, port)
	config := &gossh.ClientConfig{
		Timeout:         5 * time.Second,
		User:            user,
		HostKeyCallback: gossh.InsecureIgnoreHostKey(),
		Auth:            []gossh.AuthMethod{gossh.Password(password)},
	}
	return gossh.Dial("tcp", addr, config)
}

// 创建 ssh client, key 认证
func GetSSHClientByKey(host string, port uint, user string, key []byte) (*gossh.Client, error) {
	addr := fmt.Sprintf("%s:%d", host, port)
	signer, err := gossh.ParsePrivateKey(key)
	if err != nil {
		return nil, fmt.Errorf("unable to parse private key: %v", err)
	}
	config := &gossh.ClientConfig{
		Timeout:         5 * time.Second,
		User:            user,
		HostKeyCallback: gossh.InsecureIgnoreHostKey(),
		Auth:            []gossh.AuthMethod{gossh.PublicKeys(signer)},
	}
	return gossh.Dial("tcp", addr, config)
}
