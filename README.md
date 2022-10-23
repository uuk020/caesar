<h1 align="center"> Caesar </h1>

<p align="center">echo 框架构建的简单密码管理 Web 应用后端 API 接口</p>

### 环境需求
* go 版本不低于 1.18
* Mysql 版本不低于 5.7 或 sqlite3 版本 3.39
* Redis 版本不低于 5.0

### 安装使用
* 下载安装
  ```shell
    git clone git@github.com:uuk020/caesar.git 
  ```
* 初始化
  ```text
  把 settings-dev.yaml 修改为 settings.yaml
  ```
* 编译 go
  ```shell
  go build main.go
  ```   
* 部署Go应用
- [Nginx 部署](https://eddycjy.gitbook.io/golang/di-3-ke-gin/nginx)
- [Caddy 部署](https://caddyserver.com/docs/)

## 接口文档
- [接口文档说明](./caesar-api.md)
## 最后
欢迎提出 issue 和 pull request

## License
Apache 2.0