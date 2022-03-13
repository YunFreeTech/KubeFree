<p align="center"><a href="https://YunFreeTech.io"><img src="https://YunFreeTech.oss-cn-beijing.aliyuncs.com/kubefree/img/logo-red.png" alt="kubefree" width="300" /></a></p>
<P align="center"><b>KubeFree</b> [kubəˈpaɪ]，一款简单易用的开源 Kubernetes 可视化管理面板。</P>
<p align="center">
  <a href="http://www.apache.org/licenses/LICENSE-2.0"><img src="https://img.shields.io/github/license/YunFreeTech/kubefree?color=%231890FF&style=flat-square" alt="License: Apache License v2"></a>
  <a href="https://app.codacy.com/gh/YunFreeTech/kubefree?utm_source=github.com&utm_medium=referral&utm_content=YunFreeTech/kubefree&utm_campaign=Badge_Grade_Dashboard"><img src="https://app.codacy.com/project/badge/Grade/da67574fd82b473992781d1386b937ef" alt="Codacy"></a>
  <a href="https://hub.docker.com/r/YunFreeTech/kubefree-server"><img src="https://img.shields.io/docker/pulls/YunFreeTech/kubefree-server" alt="Docker Pulls"></a>
  <a href="https://github.com/YunFreeTech/KubeFree"><img src="https://img.shields.io/github/stars/YunFreeTech/KubeFree" alt="Stars"></a>
</p>
<hr />

KubeFree 是一款简单易用的开源 Kubernetes 可视化管理面板。

KubeFree 允许管理员导入多个 Kubernetes 集群，并且通过权限控制，将不同 cluster、namespace 的权限分配给指定用户。它允许开发人员管理 Kubernetes 集群中运行的应用程序并对其进行故障排查，供开发人员更好地处理 Kubernetes 集群中的复杂性。


### 项目结构

.
├── Dockerfile                                      # 构建容器镜像使用的 dockerfile
├── Makefile                                        # 编译文件
├── LICENSE
├── README.md
├── go.mod
├── conf                                            # 配置文件
├── pkg                                             # 主目录
├── interanl                                        # 私有业务代码
├── migrate                                         # 数据库变更文件目录
├── web
│   ├── dashboard                                   # 前端 dashboard 模块
│   ├── kubefree                                      # 前端管理模块
│   └── terminal                                    # terminal 模块   
### 配置文件
KubeFree 会默认加载该路径下的配置文件 /etc/kubefree/app.yaml，请参考下列配置创建对应目录及配置文件


apiVersion: v1
kind: AppConfig
spec:
  server:
    bind:
      host: 0.0.0.0
      port: 80
    ssl:
      enable: false
      certificate:
      certificateKey:
  db:
    path: /var/lib/kubefree/db/kubefree.db
### 数据库文件
KubeFree 使用 BoltDB 作为底层数据存储，使用 storm 作为上层 ORM 框架


数据库文件地址 /var/lib/kubefree/db/kubefree.db
### 构建二进制文件
build


make build_web
### 启动后端服务

在 cmd/server 目录运行   
``` 
go run main.go
```
#### 启动前端服务
KubeFree 的前端主要包括两部分，需要分别启动

#### 管理模块


在 web/kubefree 目录运行 
```
npm install
npm run serve
dashboard 模块
```

在 web/dashboard 目录运行
```
npm install    
npm run serve
```

### UI 展示

![UI展示](https://YunFreeTech.oss-cn-beijing.aliyuncs.com/kubefree/img/kubefree-demo.gif)

### 快速开始

    sudo docker run --privileged -d --restart=unless-stopped -p 80:80 YunFreeTech/kubefree-server

- 打开浏览器访问：http://localhost:80/
- 用户名：admin
- 密码：kubefree

### 在线体验

- 环境地址：<https://demo.YunFreeTech.io:8080/>
- 用户名：demo
- 密码：Password@123

| :warning: 注意                 |
| :--------------------------- |
| 该环境仅作体验目的使用，我们会定时清理、重置数据！    |
| 请勿修改体验环境用户的密码！               |
| 请勿在环境中添加业务生产环境地址、用户名密码等敏感信息！ |

### 微信交流群

![wechat-group](https://YunFreeTech.io/docs/img/wechat-group.png)

### 致谢

- [Vue](https://cn.vuejs.org) 前端框架
- [FIT2CLOUD UI](https://github.com/fit2cloud-ui/fit2cloud-ui/) FIT2CLOUD UI 组件库
- [Vue-element-admin](https://github.com/PanJiaChen/vue-element-admin) 项目脚手架

### License & Copyright

Copyright (c) 2014-2022 FIT2CLOUD 飞致云

[https://www.fit2cloud.com](https://www.fit2cloud.com)<br>

KubeFree is licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
