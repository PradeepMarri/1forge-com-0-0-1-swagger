package main

import (
	"github.com/1forge-finance-apis/mcp-server/config"
	"github.com/1forge-finance-apis/mcp-server/models"
	tools_forex "github.com/1forge-finance-apis/mcp-server/tools/forex"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_forex.CreateGet_quotesTool(cfg),
		tools_forex.CreateGet_symbolsTool(cfg),
	}
}
