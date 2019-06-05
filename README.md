# findlinks
Deploying a Go Application in a Docker Container

### 正常部署流程
1.您应该知道的第一件事是Docker允许您孤立地运行应用程序并且关注点很大，但允许它们与外部世界进行通信和交互。

2.把simple-app放到github上去:https://github.com/zhengjianfeng1103/findlinks

3.运行命令:docker run golang go get -v github.com/zhengjianfeng1103/findlinks
> $ docker run golang go get -v github.com/zhengjianfeng1103/findlinks  
> golang.org/x/net (download)  
> golang.org/x/net/html/atom  
> golang.org/x/net/html
> github.com/flaviocopes/findlinks

4.运行命令:docker ps -l 列出最新的容器列表
> CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS                     PORTS               NAMES
343d96441f16        golang              "go get -v github...."   3 minutes ago       Exited (0) 2 minutes ago  mystifying_swanson
> 容器将会推出，直到go get命令完成执行

5.运行命令:docker commit $(docker ps -lq) findlinks
> docker ps -lq 获取最后的容器ID
> 使用findlinks作为仓库名称，打包成镜像
> REPOSITORY          TAG                 IMAGE ID 
>            CREATED             SIZE
findlinks           latest              4e7ebb87d02e        11 seconds ago      720MB

5.运行命令:docker run -p 8000:8000 findlinks findlinks
> 使用findlinks命令，执行finlinks镜像

### 精简构建流程 缩小Docker Image包  
1.为什么这个包那么大，达到720m？
> 因为内置的go的编译包

2.构建mini版本的Docker Image容器

2.1 安装本地的golang:1.8.3版本包 /gopath  
2.2 本地创建项目文件夹 /app  
2.3 去除 CGO模块  
2.4 创建Dockerfile,告诉Docker使用 iron/base这个非常轻量级的镜像  

> docker run --rm -it -v "$GOPATH":/gopath -v "$(pwd)":/app -e "GOPATH=/gopath" -w /app golang:1.8.3 sh -c 'CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o findlinks'
> 
