package vectorstore

import (
	"context"
	"log"

	"github.com/google/uuid"
	qdrant "github.com/qdrant/go-client/qdrant"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client qdrant.PointsClient

func InitQdrant() {
	conn, err := grpc.Dial("localhost:6334", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("No se pudo conectar a Qdrant: %v", err)
	}

	client = qdrant.NewPointsClient(conn)
	log.Println("✅ Qdrant conectado")
}

func InsertEmbedding(texto string, vector []float32) error {
	_, err := client.Upsert(context.Background(), &qdrant.UpsertPoints{
		CollectionName: "embeddings",
		Points: []*qdrant.PointStruct{
			{
				Id: &qdrant.PointId{
					PointIdOptions: &qdrant.PointId_Uuid{
						Uuid: uuid.New().String(),
					},
				},
				Vectors: &qdrant.Vectors{
					VectorsOptions: &qdrant.Vectors_Vector{
						Vector: &qdrant.Vector{Data: vector},
					},
				},
				Payload: map[string]*qdrant.Value{
					"texto": {Kind: &qdrant.Value_StringValue{StringValue: texto}},
				},
			},
		},
	})
	return err
}
