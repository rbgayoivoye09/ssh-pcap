# ssh-pcap


```bash

# 开始抓包
go run src/main/main.go -c config/base.yaml ssh -p
# 查看抓包进程
go run src/main/main.go -c config/base.yaml ssh --show
# 停止抓包
go run src/main/main.go -c config/base.yaml ssh -s
# 下载抓包
go run src/main/main.go -c config/base.yaml ssh -d

```