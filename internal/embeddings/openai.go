package embeddings

import (
	"context"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func GenerarEmbeddings(textos []string) ([][]float32, error) {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY_RETAIL_BRAIN"))

	resp, err := client.CreateEmbeddings(context.Background(), openai.EmbeddingRequest{
		Model: openai.SmallEmbedding3, // text-embedding-3-small
		Input: textos,
	})
	if err != nil {
		return nil, err
	}

	var embeddings [][]float32
	for _, r := range resp.Data {
		embeddings = append(embeddings, r.Embedding)
	}
	return embeddings, nil
}
