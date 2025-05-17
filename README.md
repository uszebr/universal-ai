# AIService - Go Client for AI API Requests

This package provides a simple and extensible Go client for making structured AI requests (e.g., to OpenAI's Chat API). It handles sending messages, tools, tool choices, and parsing the response.

## Features

- Easy-to-use interface for interacting with AI APIs
- Supports messages with system, user, assistant, and tool roles
- Tool calling and tool choice support
- Context-aware HTTP requests
- Uses [resty](https://github.com/go-resty/resty) for HTTP requests

## Installation

Make sure you have Go installed (1.18+ recommended).

```bash
go get github.com/go-resty/resty/v2
```
