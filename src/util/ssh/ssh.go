package ssh

import (
	"bytes"

	"golang.org/x/crypto/ssh"
)

func ExecuteRemoteCommand(host, port, user, password, command string) (string, error) {
	// 配置 SSH 客户端
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// 连接到远程主机
	client, err := ssh.Dial("tcp", host+":"+port, config)
	if err != nil {
		return "", err
	}
	defer client.Close()

	// 创建 SSH 会话
	session, err := client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	// 执行命令
	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	err = session.Run(command)
	if err != nil {
		return "", err
	}
	return stdoutBuf.String(), nil
}
