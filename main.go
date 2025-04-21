package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type LinkRequest struct {
	Destination string `json:"destination"`
	Slashtag    string `json:"slashtag"`
	Title       string `json:"title"`
}

func main() {
	// MCP Server
	s := server.NewMCPServer(
		"Rebrandly API",
		"1.0.0",
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
		server.WithRecovery(),
	)

	// Add a shorten link tool
	linkShortenerTool := mcp.NewTool("create_short_link",
		mcp.WithDescription("Generate a short link using Rebrandly API"),
		mcp.WithString("destination_url", mcp.Required(), mcp.Description("Destination URL")),
		mcp.WithString("workspace", mcp.Description("Optional Rebrandly workspace ID")),
		mcp.WithString("slashtag", mcp.Description("Optional custom slashtag")),
		mcp.WithString("title", mcp.Description("Optional title for the link")),
	)

	// Add the shorten link handler
	s.AddTool(linkShortenerTool, shortenLinkHandler)

	// Start the server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

func shortenLinkHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	apiKey := os.Getenv("REBRANDLY_API_KEY")
	if apiKey == "" {
		return mcp.NewToolResultError("Missing Rebrandly API Key environment variable"), nil
	}

	args := req.Params.Arguments

	dest, ok := args["destination_url"].(string)
	if !ok || strings.TrimSpace(dest) == "" {
		return mcp.NewToolResultError("Missing or invalid destination URL"), nil
	}

	// Building request payload
	linkReq := LinkRequest{
		Destination: dest,
	}

	if val, ok := args["slashtag"].(string); ok {
		linkReq.Slashtag = val
	}

	if val, ok := args["title"].(string); ok {
		linkReq.Title = val
	}

	body, _ := json.Marshal(linkReq)

	httpReq, err := http.NewRequest("POST", "https://api.rebrandly.com/v1/links", strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}

	// Required Headers
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Accept", "application/json")
	httpReq.Header.Set("apikey", apiKey)

	// Optional workspace header
	if workspace, ok := args["workspace"].(string); ok && workspace != "" {
		httpReq.Header.Set("workspace", workspace)
	}

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return mcp.NewToolResultErrorFromErr("Request to Rebrandly failed", err), nil
	}
	defer resp.Body.Close()

	// Response from Rebrandly was not ok or created
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return mcp.NewToolResultError(fmt.Sprintf("Rebrandly returned status %d", resp.StatusCode)), nil
	}

	// Decoding the response
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return mcp.NewToolResultErrorFromErr("Failed to decode Rebrandly response", err), nil
	}

	shortUrl, ok := result["shortUrl"].(string)
	if !ok {
		return nil, errors.New("short url not found in response")
	}

	return mcp.NewToolResultText(fmt.Sprintf("Shortened URL: https://%s", shortUrl)), nil
}
