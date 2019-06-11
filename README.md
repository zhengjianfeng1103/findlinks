# findlinks
## Go run app
go run findlinks.go
http://localhost:8000/?q=flaviocopes.com


## Deploying a Go Application in a Docker Container
### 正常部署流程
1.您应该知道的第一件事是Docker允许您孤立地运行应用程序并且关注点很大，但允许它们与外部世界进行通信和交互。

2.把simple-app放到github上去:https://github.com/zhengjianfeng1103/findlinks

3.运行命令:docker run golang go get -v github.com/zhengjianfeng1103/findlinks
> 首先下载golang Docker镜像，如果你还没有它，那么它将获取存储库并将扫描标准库中未包含的其他依赖项。
```
> $ docker run -v golang go get -v github.com/zhengjianfeng1103/findlinks  
> golang.org/x/net (download)  
> golang.org/x/net/html/atom  
> golang.org/x/net/html
> github.com/flaviocopes/findlinks
```

4.运行命令:docker ps -l 列出最新的容器列表
> CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS                     PORTS               NAMES
343d96441f16        golang              "go get -v github...."   3 minutes ago       Exited (0) 2 minutes ago  mystifying_swanson
> 容器将会退出, 直到go get命令完成执行

5.运行命令:docker commit $(docker ps -lq) findlinks
> docker ps -lq 获取最后的容器ID
> 使用findlinks作为仓库名称，把运行的container容器中的打包到名为findlinks的repository仓库中, tag默认为latest  
>> REPOSITORY          TAG                 IMAGE ID 
>            CREATED             SIZE  
findlinks           latest              4e7ebb87d02e        11 seconds ago      720MB

5.运行命令:docker run -p 8000:8000 findlinks findlinks
> 使用findlinks命令，执行finlinks镜像

### 精简构建流程 缩小Docker Image包  
1. 为什么这个包那么大，达到720m？
> 因为内置的go的编译包  
> 我们告诉Docker运行golang：1.8.3 映像并静态编译我们的应用程序，禁用CGO，这意味着镜像不需要动态链接时通常需要访问的C库  
> 使用docker容器编译打包出filelinks二进制可执行文件   
> > docker run --rm -it -v "$GOPATH":/gopath -v "$(pwd)":/app -e "GOPATH=/gopath" -w /app golang:1.8.3 sh -c 'CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o findlinks'  
 
``` Unable to find image 'golang:1.8.3' 
    locally 1.8.3: Pulling from library/golang
```
>> 1. --rm 容器退出时就能够自动清理容器内部的文件系统  
>> 2. -it Docker分配连接到容器的stdin的伪TTY; 在容器中创建一个交互式的bash shell  
>> 3. -v 将当前工作目录装载到容器中  
>> 4. -w 设置工作目录
>> 5. -e 设置任何环境变量  


2. 构建mini版本的Docker Image容器  
docker build -t [TAG] [sourceDir]  
2.1 使用前面编译出的可执行文件,构建镜像. Dockerfile如下:
```
    FROM iron/base
    WORKDIR /app
    COPY findlinks /app/
    ENTRYPOINT ["./findlinks"]
```
>docker build -t flaviocopes/golang-docker-example-findlinks .

```
Step 1/4 : FROM iron/base
latest: Pulling from iron/base
Step 2/4 : WORKDIR /app
Step 3/4 : COPY findlinks /app/
Step 4/4 : ENTRYPOINT ["./findlinks"]

REPOSITORY TAG  IMAGE ID  CREATED SIZE  
flaviocopes/golang-docker-example-findlinks   latest              d27681ab8465        18 minutes ago      8.91MB
```
3.运行mini镜像
docker run -x -xx -x [REPOSITORY:TAG]
>docker run --rm -it -p 8000:8000 flaviocopes/golang-docker-example-findlinks