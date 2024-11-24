package llm

import (
	"context"

	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/googleai"
	"github.com/tmc/langchaingo/tools/sqldatabase"
	_ "github.com/tmc/langchaingo/tools/sqldatabase/sqlite3"
)

func InitSQLDatabaseChain(ctx context.Context, apikey string, filepath string, topk int) (*chains.SQLDatabaseChain, error) {
	llm, err := googleai.New(ctx, googleai.WithAPIKey(apikey))

	if err != nil {
		return nil, err
	}

	db, err := sqldatabase.NewSQLDatabaseWithDSN("sqlite3", filepath, nil)

	if err != nil {
		return nil, err
	}

	sqlDatabaseChain := chains.NewSQLDatabaseChain(llm, topk, db)

	return sqlDatabaseChain, nil
}

func Run(ctx context.Context, chain *chains.SQLDatabaseChain, prompt string) (string, error) {
	return chains.Run(ctx, chain, prompt)
}
