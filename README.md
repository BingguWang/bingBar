# bingBar

### 项目运行

#### Redis运行

1. mkdir -p /home/wangbing/redis/conf && mkdir -p /home/wangbing/redis/data

2. cd /home/wangbing/redis/conf sudo wget http://download.redis.io/redis-stable/redis.conf


3. 在redis.conf里先开启Redis AOF:

- appendonly yes
- appendfsync everysec

4. docker run 命令运行容器:
  
   ```
   docker run --name redis -p 6379:6379 --privileged=true \
      -v /home/wangbing/reids/data:/data \
      -v /home/wangbing/reids/conf:/etc/redis/redis.conf \
      -e TZ=Asia/Shanghai \
      --restart=always \
      --appendonly yes \
      -d redis:7.0 /etc/redis/redis.conf --requirepass "123456"
   
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
   
3. 创建数据库user
   ```
   CREATE DATABASE `d_user` CHARACTER SET 'utf8mb4' COLLATE 'utf8mb4_unicode_ci';
   ```
   
