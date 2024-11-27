package ssh

import (
	"bytes"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"

	"github.com/rbgayoivoye09/ssh-pcap/src/util/flag"
	. "github.com/rbgayoivoye09/ssh-pcap/src/util/log"
)

func ExecuteRemoteCommand(host, port, user, password, cmd_type string, command []string, local_file_path string) (string, error) {
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

	// range command slice combine to a string by ";"
	commandStr := ""
	for _, v := range command {
		commandStr += v + " ; "
	}
	commandStr = commandStr[:len(commandStr)-1]

	Logger.Sugar().Infof("commandStr: %v", commandStr)

	if cmd_type == flag.DOWN_FLAG {
		err = session.Run(commandStr)
		if err != nil {
			return "", err
		} else {
			// 打开SFTP会话
			sftpClient, err := sftp.NewClient(client)
			if err != nil {
				Logger.Sugar().Infof("Failed to create sftp client: ", err)
				return "", err
			}
			defer sftpClient.Close()

			// 远程目录
			remoteDir := "/tmp/" // 修改为你的远程目录

			// 列出远程目录下的所有文件
			files, err := sftpClient.ReadDir(remoteDir)
			if err != nil {
				Logger.Sugar().Infof("Failed to read directory: ", err)
				return "", err
			}

			// 筛选出以 .tar.gz 结尾的文件
			var tarFiles []os.FileInfo
			for _, file := range files {
				if strings.HasSuffix(file.Name(), ".tar.gz") {
					tarFiles = append(tarFiles, file)
				}
				if strings.HasSuffix(file.Name(), ".gz") {
					tarFiles = append(tarFiles, file)
				}
			}

			// 按修改时间对文件排序
			sort.Slice(tarFiles, func(i, j int) bool {
				return tarFiles[i].ModTime().After(tarFiles[j].ModTime())
			})

			// 获取最新的 .tar.gz 文件
			if len(tarFiles) > 0 {
				latestFile := tarFiles[0]
				fmt.Printf("latestFile: %v\n", latestFile)
				remoteFilePath := fmt.Sprintf("%s%s", remoteDir, latestFile.Name())
				fmt.Printf("remoteFilePath: %v\n", remoteFilePath)

				// 下载最新文件到本地
				localFilePath := local_file_path + latestFile.Name() // 修改为你想要保存的本地路径
				fmt.Printf("localFilePath: %v\n", localFilePath)
				err = downloadFile(sftpClient, remoteFilePath, localFilePath)
				if err != nil {
					Logger.Sugar().Infof("Failed to download file: ", err)
					return "", err
				}

				Logger.Sugar().Infof("File downloaded successfully:", localFilePath)
			} else {
				Logger.Sugar().Infof("No files found in the remote directory.")
			}

		}
	} else if cmd_type == flag.SHOW_FLAG {

		err = session.Run(commandStr)
		if err != nil {
			return "", err
		}

	} else {
		err = session.Start(commandStr)
		if err != nil {
			return "", err
		}

	}
	return stdoutBuf.String(), nil
}

// 下载文件
func downloadFile(sftpClient *sftp.Client, remoteFilePath, localFilePath string) error {
	srcFile, err := sftpClient.Open(remoteFilePath)
	if err != nil {
		return err
	} else {
		fmt.Printf("srcFile: %v\n", srcFile)
	}
	defer srcFile.Close()

	dstFile, err := os.Create(localFilePath)
	if err != nil {
		return err
	} else {
		fmt.Printf("dstFile: %v\n", dstFile)
	}
	defer dstFile.Close()

	if _, err := srcFile.WriteTo(dstFile); err != nil {
		return err
	}

	return nil
}
