## 安装 go-ethereum 相关tools

```shell
git config --global http.proxy http://localhost:8001
git config --global https.proxy http://localhost:8001
git clone https://github.com/figoxu/go-ethereum.git

git config --global --unset http.proxy
git config --global --unset https.proxy
 git config  --list
```

备注: 8001 是 佛跳墙的http端口

git config --global http.proxy http://localhost:8001
516 git config --global https.proxy http://localhost:8001
517 git clone https://github.com/figoxu/go-ethereum.git
518 cd go-ethereum
519 make all
520 cd build/bin

git config --global --unset http.proxy
git config --global --unset https.proxy

把下面地址加入path
/Users/zhangxiazhen/figoxu/github.com/go-ethereum/build/bin



# solc绿色安装

https://github.com/ethereum/solidity/releases
下载 solc-macos 放在path里
mv /Users/zhangxiazhen/Downloads/solc-macos /Users/zhangxiazhen/figoxu/github.com/go-ethereum/build/bin/solc

