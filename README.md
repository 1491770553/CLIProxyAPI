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
