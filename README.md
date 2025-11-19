# workflow

CI/CD with godeploy

## 核心 workflow yml 模版文件

```bash
# 打包/发布（golang）
_go-build-and-release-app.yml

# 打包/发布（react）
_react-build-and-release-web.yml

# 通用部署（golang+react）
_deploy-app.yml
```

## 使用方法

```bash
# ==== 准备工作 ====
# 1. ubuntu24、系统用户 leon、ssh 密钥
# 2. 远程环境已安装 godeploy 命令，工作目录为用户 leon 配置好读写权限
# 3. 配置文件：1. /srv/app/my-app/godeploy.env 2. /etc/systemd/system/app.service

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

https://github.com/terobox/godeploy

```bash
# 下载当前最新版本
curl -fsSL https://raw.githubusercontent.com/terobox/godeploy/main/install.sh | sudo bash

# 验证安装
godeploy --version
```

## build, release and deploy in one yml 拼接

```bash
# 合并步骤：
# 0. 合并 env
# 1. jobs.build-and-release.if 新增，手动只部署，不 build/release（可选）
# 2. jobs.deploy 新增
# 3. jobs.deploy.needs 新增，先构建再部署

# 参考：
# auto-backend.yml
```

## 一次打包/发布，多机部署

```bash
# 调整步骤：
# 0. 在 deploy 这个 job 里加 matrix 策略。deploy.strategy 新增，参数如下：
deploy:
    needs: build-and-release   # ★ 关键：先构建再部署
    runs-on: ubuntu-latest
    strategy:
        fail-fast: false # 某一台失败不会导致其它机器被取消
        matrix:
            host_secret_name: [ 'SSH_HOST_08', 'SSH_HOST_15', 'SSH_HOST_21' ]
# 说明：会并发跑 3 个 deploy job。fail-fast: false 保证其中一个失败，不会把剩下两个 job 停掉。

# 1. 修改 SSH 步骤里的 host 字段。改成用 matrix 动态取秘钥名：
# 旧：
host: ${{ secrets.SSH_HOST_08 }} # 5/5
# 新：
host: ${{ secrets[matrix.host_secret_name] }}

# 2. 如果你想在 log 里看到当前是哪个“逻辑”主机，可以顺手在 Show deploy info 里加一行（可选）：
echo "matrix.host_secret_name: ${{ matrix.host_secret_name }}"
```

## action-gh-release 同一个 repo 里多个 workflow 都用同一个 tag（比如 v0.0.5），存在冲突问题