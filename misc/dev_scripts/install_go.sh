#! /bin/sh

PACKAGE=go1.16.6.linux-amd64.tar.gz
# 1.下载安装包
wget https://studygolang.com/dl/golang/$PACKAGE

# 2.解压到 /usr/local（建议）
tar xzvf $PACKAGE -C /usr/local

# 3.设置环境变量
WORKSPACE=/vagrant_data/gocar # 任何工作目录都行
mkdir -p $WORKSPACE/bin $WORKSPACE/src $WORKSPACE/pkg
echo "
#! /bin/sh
export GOROOT=/usr/local/go
export GOPATH=$WORKSPACE
export GO111MODULE=on                   	# 开启go mod包管理支持
export GOPROXY=https://goproxy.cn,direct        # 设置包下载代理 可选
export GOSUMDB=off
export PATH=\$PATH:\$GOROOT/bin:\$GOPATH/bin
" > /etc/profile.d/goenv.sh
source /etc/profile.d/goenv.sh

# verify
echo installed `go version`
