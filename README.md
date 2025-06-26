# Rebrandly MCP Tool (Go)

This project implements a simple MCP server in Go that exposes a single tool (`create_short_link`) to generate short URLs using the Rebrandly API.

## Requirements

- Go 1.20 or higher
- A [Rebrandly](https://www.rebrandly.com) account and an API Key

## Build Instructions

1. Clone this repository or save the `main.go` file.
2. Build the binary using:

```bash
go build -o rebrandly-tool main.go
```

This will generate an executable named `rebrandly-tool`.

## Usage with Claude Desktop

Claude Desktop supports custom MCP servers. To connect this tool:

1. Open your Claude Desktop configuration.
2. Add the following entry under `mcpServers`:

```json
{
  "mcpServers": {
    "Rebrandly": {
      "command": "PATH/TO/rebrandly-tool",
      "args": [],
      "env": {
        "REBRANDLY_API_KEY": "YOUR_KEY"
      }
    }
  }
}
```

> 🔧 Replace `PATH/TO/rebrandly-tool` with the actual path to the binary, and `YOUR_KEY` with your Rebrandly API Key.

## Tool Available

### `create_short_link`

This tool generates a short URL via Rebrandly.

#### Parameters

| Name        | Required | Description                                          |
| ----------- | -------- | ---------------------------------------------------- |
| destination | ✅       | The original long URL to shorten                     |
| workspace   | ❌       | Optional workspace ID (for multi-workspace accounts) |
| slashtag    | ❌       | Optional custom slug                                 |
| title       | ❌       | Optional title for the short link                    |

## License

Licensed under MIT - see [LICENSE](./LICENSE) file.

## Rebrandly in MCP Registries

- [https://mcpreview.com/mcp-servers/rebrandly/rebrandly-mcp](https://mcpreview.com/mcp-servers/rebrandly/rebrandly-mcp)
- [https://mcp.so/server/rebrandly-mcp/rebrandly](https://mcp.so/server/rebrandly-mcp/rebrandly)
- [https://glama.ai/mcp/servers/@rebrandly/rebrandly-mcp](https://glama.ai/mcp/servers/@rebrandly/rebrandly-mcp)
