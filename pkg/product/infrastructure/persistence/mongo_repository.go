package persistence

import (
	"context"
	"time"

	"github.com/jamsxd/product-crud-example/pkg/product/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type basicMongoRepository struct {
	collection *mongo.Collection
}

func NewMongoRepository(client *mongo.Client, dbName, collection string) domain.ProductRepository {
	return &basicMongoRepository{collection: client.Database(dbName).Collection(collection)}
}

func (r *basicMongoRepository) FindAll(ctx context.Context) ([]domain.Product, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	cursor, err := r.collection.Find(ctxTimeout, bson.D{})
	if err != nil {
		return nil, err
	}

	var products []domain.Product
	if err := cursor.All(ctx, &products); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *basicMongoRepository) FindBySku(ctx context.Context, sku string) (*domain.Product, error) {
	var product domain.Product
	ctxTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	err := r.collection.FindOne(ctxTimeout, bson.D{{
		Key:   "sku",
		Value: sku,
	}}).Decode(&product)
	if err != nil && err == mongo.ErrNoDocuments {
		return nil, domain.ErrProductNotFound
	}
	return &product, err
}

func (r *basicMongoRepository) Upsert(ctx context.Context, product *domain.Product) error {
	upsert := true
	ctxTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	_, err := r.collection.UpdateOne(ctxTimeout, map[string]interface{}{"sku": product.Sku}, map[string]interface{}{"$set": product}, &options.UpdateOptions{
		Upsert: &upsert,
	})
	return err
}

func (r *basicMongoRepository) Delete(ctx context.Context, sku string) error {
	ctxTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	_, err := r.collection.DeleteOne(ctxTimeout, map[string]interface{}{"sku": sku})
	return err
}
