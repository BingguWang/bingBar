# bingBar

## 介绍
之江星海存储云平台后端仓库。

## 软件架构
软件架构说明



### 项目运行

1. 项目构建


2. shell运行(构建容器并运行容器)


#### Redis运行

1. mkdir -p /home/wangbing/redis/conf
mkdir -p /home/wangbing/redis/data

2. cd /home/redis-test/conf 
sudo wget http://download.redis.io/redis-stable/redis.conf


3. 在redis.conf里先开启Redis AOF:
-   appendonly yes  
-   appendfsync everysec

4. docker run 命令运行容器:
   ```
   docker run -p 6379:6379 \
   --name redis \
   --privileged=true \
   -e TZ=Asia/Shanghai \
   -v /home/wangbing/redis/conf:/etc/redis/conf \
   -v /home/wangbing/redis/data:/data \
   -d redis:7.0 \
   redis-server /etc/redis/conf/redis.conf \
   --appendonly yes \
   --bind 0.0.0.0 \
   --requirepass "123456"
   ```   
#### Etcd运行
   docker 运行 etcd
   ```
   docker run -d --name etcd-server \
    --network app-tier \
    --publish 2379:2379 \
    --env ALLOW_NONE_AUTHENTICATION=yes \
    --env ETCD_ADVERTISE_CLIENT_URLS=http://etcd-server:2379 \
    bitnami/etcd:latest
   ```

#### Mysql运行
mysql 镜像位置: icmp.harbor/starocean/mysql:8.0

1. 在/home/mysql/conf下创建my.cnf文件,添加如下内容:
   ```
   [client]
   default-character-set=utf8mb4
   [mysqld]
   init_connect='SET collation_connection = utf8mb4_unicode_ci'
   init_connect='SET NAMES utf8mb4'
   character-set-server=utf8mb4
   collation-server=utf8mb4_unicode_ci
   skip-character-set-client-handshake
   [mysql]
   default-character-set = utf8mb4
   ```
   
2. docker run 命令运行容器:
   ```
   docker run \
   --name mysql \
   -p 3306:3306 \
   -v /home/mysql/data:/var/lib/mysql \
   -v /home/mysql/conf:/etc/mysql/conf.d \
   -v /home/mysql/log:/var/log/mysql \
   -e MYSQL_ROOT_PASSWORD=123456 \
   --privileged=true -d mysql:8.0
   ```
3. 创建数据库starocean
   ```
   CREATE DATABASE `d_user` CHARACTER SET 'utf8mb4' COLLATE 'utf8mb4_unicode_ci';
   ```
   
