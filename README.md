# Rebrandly MCP Tool (Go) - ARCHIVED

> **This repository has been archived.** Rebrandly now provides an official MCP server with expanded capabilities. Please use the official solution instead.

## Official Rebrandly MCP Server

Rebrandly offers an official MCP server that connects AI assistants like Claude, Gemini, and GitHub Copilot to your Rebrandly account with full-featured support for:

- **Link Management** - Create, update, search, and organize branded short links
- **Analytics** - Access click statistics, geographic data, and device information
- **Account Management** - Workspace administration, teammate management, and domain operations

### Quick Start

Add this configuration to your Claude Desktop `claude_desktop_config.json`:

```json
{
  "mcpServers": {
    "rebrandly": {
      "command": "npx",
      "args": [
        "-y",
        "supergateway",
        "--streamableHttp",
        "https://mcp.rebrandly.com/v1/mcp",
        "--header",
        "REBRANDLY_API_KEY:YOUR_API_KEY"
      ]
    }
  }
}
```

Replace `YOUR_API_KEY` with your Rebrandly API key from your account settings.

### Documentation

- [MCP Server Overview](https://developers.rebrandly.com/docs/mcp-server)
- [Quick Start Guide](https://developers.rebrandly.com/docs/quick-start-for-mcp)

---

## About This Repository

This was an early implementation of a Rebrandly MCP server written in Go. It exposed a single `create_short_link` tool for generating short URLs. The official Rebrandly MCP server supersedes this implementation with a more complete feature set.

## License

Licensed under MIT - see [LICENSE](./LICENSE) file.