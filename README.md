
# Ginblog（项目已完成，欢迎使用)

## 重要更新

- 上传七牛服务器区域配置，现在在 `config/config.ini` 下可以配置七牛对象存储的服务器区域，配置更方便。

- 取消后台页面新增文章中上传图片为“必填”的限制，避免在配置七牛云上传不成功时报错。

- 修改静态资源托管路径，前端项目下直接打包，不用再移动到static目录下，更加方便。

- 增加了展示页面的移动端适配


## 介绍

gin+gorm+vue 全栈制作一个博客。

## 目录结构

```shell
├─  .gitignore
│  go.mod // 项目依赖
│  go.sum
│  LICENSE
│  main.go //主程序
│  README.md
│  tree.txt
│          
├─api     // 一般接口会和model中的数据模型一一对应，便于项目管理  
├─config // 项目配置入口   
├─database  // 数据库备份文件（初始化）
├─log  // 项目日志
├─middleware  // 中间件
├─model // 数据模型层
├─routes
│      router.go // 路由入口    
├─static // 打包静态文件
│  ├─admin  // 后台管理页面 (已废弃，打包静态文件在web/admin/dist下)         
│  └─front  // 前端展示页面 (已废弃，打包静态文件在web/front/dist下) 
├─upload   
├─utils // 项目公用工具库
│  │  setting.go 
│  ├─errmsg   
│  └─validator         
└─web // 前端开发源码（VUECLI项目源文件）
    ├─admin             
    └─front

本项目中，api与数据模型一一对应，在api中会有对应的函数调用模型中的同名函数；模型中的函数操作数据库，api中的函数负责处理前端过来的url请求。

.prettierrc : ESLint代码格式配置
babel.config.js : 在这里配置本项目需要使用的antui(按需引入)

本项目下载的式源码，自行封装的插件
富文本编辑器：http://tinymce.ax-z.cn/      https://www.tiny.cloud/docs/tinymce/6/vue-zip/

前端管理使用的组件库：https://ant-design.antgroup.com/components/overview-cn
前端展示使用的组件库：https://vuetifyjs.com/en/getting-started/installation/#nuxt-install
```

## 运行&&部署

1. 克隆项目

```shell
git clone https://github.com/CUMTB-wangru/ginblog
```

2. 转到下面文件夹下

```shell
cd yourPath/ginbolg
```

3. 安装依赖

```
go mod tidy
```

4. 初始化项目配置config.ini

```ini
./config/config.ini

[server]
AppMode = debug # debug 开发模式，release 生产模式
HttpPort = :3000 # 项目端口
JwtKey = 89js82js72 #JWT密钥，随机字符串即可

[database]
Db = mysql #数据库类型，不能变更为其他形式
DbHost = 127.0.0.1 # 数据库地址
DbPort = 3306 # 数据库端口
DbUser = ginblog # 数据库用户名
DbPassWord = admin123 # 数据库用户密码
DbName = ginblog # 数据库名

[qiniu]
# 七牛储存信息
Zone = 1 # 1:华东;2:华北;3:华南,不填默认华北。境外服务器特殊使用环境自行配置
AccessKey =
SecretKey =
Bucket =
QiniuSever =
```

5. 在database中将sql文件导入数据库

   推荐navicat或者其他sql管理工具导入

6. 启动项目

```shell
 go run main.go
```

此时，项目启动，你可以访问页面

```shell
首页
http://localhost:3000
后台管理页面
http://localhost:3000/admin

默认管理员:admin  密码:123456
```

enjoy~~~~

#### ==使用、二开过程中，发现问题或者有功能需求欢迎提交 `Iusse` 或者直接 `PR`==

## 实现功能

1. 简单的用户管理权限设置
2. 用户密码加密存储
3. 文章分类自定义
4. 列表分页
5. 图片上传七牛云
6. JWT 认证
7. 自定义日志功能
8. 跨域 cors 设置
9. 文章评论功能

## 技术栈
```
- golang
    - Gin web framework
    - gorm(v1 && v2)
    - jwt-go
    - scrypt
    - logrus
    - gin-contrib/cors
    - go-playground/validator/v10
    - go-ini
- JavaScript
    - vue
    - vue cli
    - vue router
    - ant design vue
    - vuetify
    - axios
    - tinymce
    - moment
- MySQL version:8.0.21

## 项目预览

- 前端展示页面
  ![](https://gitee.com/wejectchan/ginblog/raw/master/upload/front1.png)

- 前端展示页面
  ![](https://gitee.com/wejectchan/ginblog/raw/master/upload/front2.png)

- 后台登录页面

  ![](https://gitee.com/wejectchan/ginblog/raw/master/upload/admin.jpg
  )

- 后台管理页面

  ![](https://gitee.com/wejectchan/ginblog/raw/master/upload/admin2.jpg)


## Docker部署

### 一、如何安装docker

官方文档：[Get Docker | Docker Documentation](https://docs.docker.com/get-docker/)

选择对应的系统进行查看，以ubuntu 18.04 LTS为例

- 卸载旧版本

```shell
sudo apt-get remove docker docker-engine docker.io containerd runc
Reading package lists... Done
Building dependency tree    
Reading state information... Done
Package 'docker-engine' is not installed, so not removed
Package 'docker' is not installed, so not removed
Package 'containerd' is not installed, so not removed
Package 'docker.io' is not installed, so not removed
Package 'runc' is not installed, so not removed
0 upgraded, 0 newly installed, 0 to remove and 3 not upgraded.
```

- 添加新版本仓库

```shell
$ sudo apt-get update

$ sudo apt-get install \
  apt-transport-https \
  ca-certificates \
  curl \
  gnupg-agent \
  software-properties-common
```

- 获取官方GPG key

```shell
$ curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
```

- 验证key,如果输出的是下列内容，则说明正确

```shell
$ sudo apt-key fingerprint 0EBFCD88

pub  rsa4096 2017-02-22 [SCEA]
   9DC8 5822 9FC7 DD38 854A E2D8 8D81 803C 0EBF CD88
uid      [ unknown] Docker Release (CE deb) <docker@docker.com>
sub  rsa4096 2017-02-22 [S]
```

- 添加仓库地址

```shell
$ sudo add-apt-repository \
"deb [arch=amd64] https://download.docker.com/linux/ubuntu \
$(lsb_release -cs) \
stable"

#也可以用国内的仓库下载，速度较快，推荐

$ sudo add-apt-repository \
  "deb [arch=amd64] https://mirrors.ustc.edu.cn/docker-ce/linux/ubuntu/ \
  $(lsb_release -cs) \
  stable"
```

- 更新仓库和安装

```shell
 $ sudo apt-get update

 $ sudo apt-get install docker-ce docker-ce-cli containerd.io
```

- 进行验证，运行hello-world

```shell
$ docker pull hello-world
$ docker run hello-world
#出现以下信息，表示docker安装成功，已经可以正常运行
Hello from Docker!
This message shows that your installation appears to be working correctly.
To generate this message, Docker took the following steps:

  1. The Docker client contacted the Docker daemon.
  2. The Docker daemon pulled the "hello-world" image from the Docker Hub.(amd64)
  3. The Docker daemon created a new container from that image which runs the executable that produces the output you are currently reading.
  4. The Docker daemon streamed that output to the Docker client, which sent it to your terminal.
	To try something more ambitious, you can run an Ubuntu container with:
	$ docker run -it ubuntu bash
	Share images, automate workflows, and more with a free Docker ID:
	https://hub.docker.com/
	For more examples and ideas, visit:
	https://docs.docker.com/get-started/
```

- 额外：配置国内仓储镜像，保证拉取速度

```shell
# 这里以阿里云为例，阿里云为用户提供了私人仓库，配置有所不同。

# 具体需要登录阿里云管理，进入控制台，找到 `镜像加速服务` -> `镜像工具` -> `镜像加速器`

sudo mkdir -p /etc/docker
sudo tee /etc/docker/daemon.json <<-'EOF'
{
 "registry-mirrors": ["https://XXX你的id.mirror.aliyuncs.com"]
}
EOF
sudo systemctl daemon-reload
sudo systemctl restart docker
```

到此，我们的docker环境就正式配置好了。

### 二、拉取镜像和创建镜像和容器编排

接下来，我们就要制作运行项目需要的镜像了

### Mysql服务器的镜像

**首先，个人非常不建议mysql用docker来部署，有几个原因：**

1. **必须做数据卷的映射，`千万不能`
   将数据库数据放在docker容器中运行，否则一但删除容器数据将全部清空，所以一定要做数据持久化！！；**
2. **不利于io，数据读写在容器中读写一次，在绑定的卷中还要读写一次，两倍读写压力，性能上要打折扣。**

如果非要在docker上部署mysql，可以这么做

```shell
#首先确定mysql是否能被搜素到，这步可以跳过，也可以在dockerhub.com中搜索
$ docker search mysql

#拉取镜像
docker pull mysql  #这里默认是拉取的最新版本，如果需要特定版本可以在镜像后面添加tag，具体版本信息可以在dockerhub.com查询

#特定版本拉取,比如要拉取8.0.22(版本号一定要是官方放出的版本号，否则是查找不到的)
docker pull mysql:8.0.22

#这时可以查看下拉取的镜像
docker images

#运行镜像
docker run -d -p 3306:3306 -v /my/own/datadir:/var/lib/mysql --name ginblog-mysql -e MYSQL_ROOT_PASSWORD=admin123  mysql

# -d 表示后台运行，并返回容器id
# -p 3006:3306 表示端口映射，具体为 -p 主机端口：容器端口
# --name 给容器取个名字
# -e MYSQL_ROOT_PASSWORD=password 给mysql root管理员设置密码
# -v /my/own/datadir:/var/lib/mysql 添加数据卷/my/own/datadir是主机的数据库路径 /var/lib/mysql是容器中的数据库路径，这一步非常重要

#进入容器配置
docker exec -it ginblog-mysql bash

root@ed9345077e02:/# mysql -u root -p
Enter password:
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 8
Server version: 8.0.22 MySQL Community Server - GPL
Copyright (c) 2000, 2020, Oracle and/or its affiliates. All rights reserved.
Oracle is a registered trademark of Oracle Corporation and/or its affiliates.
Other names may be trademarks of their respective owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql>

# 之后就和一般情况下mysql的操作一样了。
```

### 制作ginblog项目镜像

- 首相要拉取我们的ginblog项目

```shell
# 新建一个项目文件夹，在你认为任何适合的地方都可以

$ cd /
$ mkdir app

# 我们这里利用git来远程同步

$ git clone 项目地址
```

- 编写Dockerfile

```dockerfile
FROM golang:latest
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct

WORKDIR /app
COPY . /app

RUN go build .

EXPOSE 3000

ENTRYPOINT ["./ginblog"]
```

- 配置ginblog的config

```ini
# config/config.ini

# DbHost = ginblog-mysql 是为了后面容器互通做准备，对应的是mysql容器的name

Db = mysql
DbHost = ginblog-mysql
DbPort = 3306
DbUser = ginblog
DbPassWord = admin123
DbName = ginblog
```

这样的话，我们大概就配置好docker的build镜像文件了。

- 最后我们在确定下前端 `web` 文件夹下的axios请求地址。`前端推荐使用 Nginx 部署`

```js
// 在 web/admin/src/plugin/http.js 和 web/front/src/plugin/http.js 两个文件夹中,将 baseURL地址改为部署的服务器线上地址

axios.defaults.baseURL = 'http://localhost:3000/api/v1'

// 修改为

axios.defaults.baseURL = 'http://线上服务器ip或域名:3000/api/v1'
```

**别忘了，修改完地址后，重新打包下**

```shell
$ yarn build

$ npm run build
```

### 生成镜像

最后一步，就是生成我们的ginblog docker image了，这部很简单，运行下列命令

```shell
$ docker build -t ginblog .
$ docker run -d -p 3000:3000 --name ginblog ginblog

#这样访问服务器IP:3000 就可以访问网站了
```
