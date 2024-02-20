文件管理系统demo

1. 本地启动，运行main.go
    访问http://127.0.0.1:8080/swagger/index.html，进入swagger
    接口包括文件上传，下载，删除，文件列表，以及pv统计

2. k8s部署

生成镜像
    docker build . -t mine:file-management

通过pvc挂载本地宿主机目录(/Users/mac/go/files)
   kubectl apply -f storage.yaml

部署deployment,service,nginx-ingress
    kubectl apply -f demo.yaml
    ingress对外暴露 www.file-management.com
    访问 https://www.file-management.com/swagger/index.html，进入swagger，调用接口

