package ssh

import (
	"fmt"
	"time"

	gossh "golang.org/x/crypto/ssh"
)

var defaultPort uint = 22
var defaultUser = "root"
var defaultTimeout uint = 5

type SSHOptions struct {
	Port    uint
	User    string
	Timeout uint
}

// 创建 ssh client, password 认证
func GetSSHClientByPassword(host string, password string, opts SSHOptions) (*gossh.Client, error) {
	var port = defaultPort
	var user = defaultUser
	var timeout = defaultTimeout
	if opts.Port != 0 {
		port = opts.Port
	}
	if len(opts.User) != 0 {
		user = opts.User
	}
	if opts.Timeout != 0 {
		timeout = opts.Timeout
	}
	addr := fmt.Sprintf("%s:%d", host, port)
	config := &gossh.ClientConfig{
		Timeout:         time.Duration(timeout) * time.Second,
		User:            user,
		HostKeyCallback: gossh.InsecureIgnoreHostKey(),
		Auth:            []gossh.AuthMethod{gossh.Password(password)},
	}
	return gossh.Dial("tcp", addr, config)
}

// 创建 ssh client, key 认证
func GetSSHClientByKey(host string, key []byte, opts SSHOptions) (*gossh.Client, error) {
	var port = defaultPort
	var user = defaultUser
	var timeout = defaultTimeout
	if opts.Port != 0 {
		port = opts.Port
	}
	if len(opts.User) != 0 {
		user = opts.User
	}
	if opts.Timeout != 0 {
		timeout = opts.Timeout
	}
	addr := fmt.Sprintf("%s:%d", host, port)
	signer, err := gossh.ParsePrivateKey(key)
	if err != nil {
		return nil, fmt.Errorf("unable to parse private key: %v", err)
	}
	config := &gossh.ClientConfig{
		Timeout:         time.Duration(timeout) * time.Second,
		User:            user,
		HostKeyCallback: gossh.InsecureIgnoreHostKey(),
		Auth:            []gossh.AuthMethod{gossh.PublicKeys(signer)},
	}
	return gossh.Dial("tcp", addr, config)
}
