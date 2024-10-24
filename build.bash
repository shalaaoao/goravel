# bash

# 复制.env.prod文件
cp ./.env.prod /home/godeploy/goravel/.env

echo 1

# 编译
go run . artisan build --os=linux --name=main

echo 2

# 移动
mv main /home/godeploy/goravel

echo 3

# 重启spv
supervisorctl restart goravel-api-server

echo 4