# 构建本地镜像
docker build -t yexy0123/go-server:1.0.0 -f ./DockerFile .

# 推送镜像至 docker 官方镜像仓库
docker push yexy0123/go-server:1.0.0

# 启动httpserver
docker run -itd --name=goserv -p8888:8888 yexy0123/go-server:1.0.0


#获取goserv pid
PID=$(docker inspect --format {{.State.Pid}} goserv)

#nsenter 进入容器的网络ms
nsenter -n -t$PID

#获取网络配置
ip addr

