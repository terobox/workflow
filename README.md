# workflow
workflwo templates

## 使用方法

```bash
# ==== 准备工作 ====
# 1. ubuntu24、系统用户 leon、ssh 密钥
# 2. 远程环境已安装 godeploy 命令，工作目录为用户 leon 配置好读写权限
# 3. 配置文件：1. /srv/app/my-app/godeploy.env 2. /etc/systemd/system/app.service
# 4. 

# ==== 提交 tag ====
git tag v0.0.1
git push origin v0.0.1
# 提交 tag 后：
# 1. _go-build-and-release-app.yml 将自动为 golang 项目打包、发布
# 2. _react-build-and-release-web.yml 将自动为 react 项目打包，发布

# ==== 手动部署 ====
# 等待 golang、react 打包/发布完成，手动执行 _deploy-app.yml，选择对应的 tag
# 完成 CI/CD
```

## 安装 godeploy

```bash
# 下载当前最新版本
curl -fsSL https://raw.githubusercontent.com/terobox/godeploy/main/install.sh | sudo bash

# 验证安装
godeploy --version
```