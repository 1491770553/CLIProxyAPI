# CLI Proxy API

为 CLI 提供 OpenAI / Gemini / Claude / Codex 兼容 API 的代理服务器，支持 OAuth 登录多账户轮询。

## 快速开始

### 下载运行

**Windows**
```bash
# 下载最新 Release 双击运行或命令行启动
.\cli-proxy-api.exe
```

**Linux**
```bash
chmod +x cli-proxy-api-linux-amd64
./cli-proxy-api-linux-amd64
```

### 默认端口

服务启动后访问 `http://localhost:8317/management.html` 打开管理面板。

## 功能特性

- **多 Provider 支持** - OpenAI (含 Codex)、Gemini、Claude Code、Qwen Code、iFlow、Antigravity
- **OAuth 登录** - 无需 API Key，通过浏览器授权即可使用官方订阅
- **多账户轮询** - 支持同 Provider 多账户负载均衡
- **兼容协议** - 提供 OpenAI / Gemini / Claude 兼容 API，适配各类 CLI 工具和 SDK
- **流式响应** - 完整支持 SSE 流式输出
- **模型映射** - 支持模型别名和强制前缀
- **多模态路由** - 检测图片内容时自动切换到视觉模型（需配置 `vision: true`）
- **管理面板** - 内置 Web UI，支持配置管理、Usage 统计、密钥管理
- **Amp CLI 集成** - 内置支持 Amp CLI IDE 扩展

## API 端点

| 端点 | 协议 |
|------|------|
| `POST /v1/chat/completions` | OpenAI |
| `POST /v1/completions` | OpenAI |
| `POST /v1/messages` | Claude |
| `POST /v1beta/models/*action` | Gemini |
| `GET /v1/models` | 统一模型列表 |

## 配置说明

配置文件位于运行目录或 `$HOME/.config/cli-proxy-api/config.yaml`：

```yaml
host: ""
port: 8317
remoteManagement:
  secretKey: your-secret-key  # 管理面板访问密钥
```

完整配置项请参考 [docs](docs/) 目录或内置管理面板。

## 认证与供应商切换

### OAuth 凭证存储

OAuth 凭证存储在 `auth-dir` 目录（默认 `~/.cli-proxy-api/`）下的 JSON 文件中：

```
~/.cli-proxy-api/
├── anthropic_xxx.json    # Claude OAuth 凭证
├── google_xxx.json        # Gemini OAuth 凭证
├── codex_xxx.json        # Codex OAuth 凭证
└── ...
```

### OAuth 登录流程

```bash
1. 调用管理 API 获取 OAuth URL
   GET /v0/management/anthropic-auth-url

2. 用户在浏览器中完成授权

3. 回调保存凭证到 auth-dir
   GET /anthropic/callback?code=xxx&state=xxx
```

### 切换供应商的方式

**方式一：客户端请求时指定**
- 客户端请求时在请求体中指定 `model` 名称
- 系统根据模型名匹配到对应的供应商凭证

**方式二：模型别名路由**
```yaml
claude-api-key:
  - api-key: sk-xxx
    models:
      - name: MiniMax-M2.7
        alias: kimi-k2.6   # 客户端用 kimi-k2.6 实际调用 MiniMax-M2.7
```

**方式三：前缀路由**
```yaml
openai-compatibility:
  - name: "openrouter"
    prefix: "test"        # 请求 "test/kimi-k2" 路由到这个 provider
```

### 多账户轮询

```yaml
routing:
  strategy: 'round-robin'  # 轮询 或 'fill-first' # 用完一个再用下一个
```

```yaml
quota-exceeded:
  switch-project: true      # 配额用完自动切换
  switch-preview-model: true # 切换到预览模型
```

### 多模态（图片理解）

当模型支持视觉识别时，可在配置中声明 `vision: true`，系统将自动完成跨 Provider 路由：

```yaml
claude-api-key:
  - api-key: sk-xxx
    models:
      - name: MiniMax-M2.7
        alias: kimi-k2.6

openai-compatibility:
  - name: "智谱"
    api-key-entries:
      - api-key: xxx
    models:
      - name: glm-4.6V
        alias: kimi-k2.6
        vision: true   # 带图片的 kimi-k2.6 请求自动切换到 glm-4.6V
```

客户端只需使用标准接口传图片，系统自动检测并路由到视觉模型，无需修改调用方式。

### 管理 API 管理凭证

| 接口 | 说明 |
|------|------|
| `GET /v0/management/auth-files` | 列出所有 OAuth 凭证 |
| `POST /v0/management/auth-files` | 上传 OAuth 凭证文件 |
| `DELETE /v0/management/auth-files` | 删除凭证 |
| `PATCH /v0/management/auth-files/status` | 启用/禁用凭证 |

## 开发构建

**编译 Windows**
```bash
Remove-Item Env:GOOS;Remove-Item Env:GOARCH; go build -o cli-proxy-api.exe ./cmd/server
```

**编译 Linux amd64**
```bash
$env:GOOS="linux"; $env:GOARCH="amd64"; go build -o cli-proxy-api-linux-amd64 ./cmd/server
```

**通用方式**
```bash
# 编译 Windows
go build -o cli-proxy-api.exe ./cmd/server

# 编译 Linux amd64
GOOS=linux GOARCH=amd64 go build -o cli-proxy-api-linux-amd64 ./cmd/server
```

## 文档

- [用户手册](https://help.router-for.me/)
- [Management API](https://help.router-for.me/cn/management/api)
- [SDK 使用指南](docs/sdk-usage_CN.md)

## 谁在使用

基于 CLIProxyAPI 构建的项目：[vibeproxy](https://github.com/automazeio/vibeproxy)、[CCS](https://github.com/kaitranntt/ccs)、[CodMate](https://github.com/loocor/CodMate)、[霖君](https://github.com/wangdabaoqq/LinJun) 等。

## 许可证

MIT License
