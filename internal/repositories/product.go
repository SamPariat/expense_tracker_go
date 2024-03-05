package repositories

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	constants "github.com/SamPariat/expenses-tracker/common"
	"github.com/SamPariat/expenses-tracker/internal/core/domain"
	"github.com/SamPariat/expenses-tracker/internal/core/ports"
)

type ProductRepositoryImpl struct {
	mongoClient *mongo.Client
}

func NewProductRepository(mongoClient *mongo.Client) ports.ProductRepository {
	return &ProductRepositoryImpl{
		mongoClient: mongoClient,
	}
}

func (productRepository *ProductRepositoryImpl) GetProducts(page int64, limit int64) ([]domain.Product, error) {
	findOptions := options.Find()
	findOptions.SetLimit(limit)
	findOptions.SetSkip(limit * (page - 1))

	cursor, err := productRepository.mongoClient.Database(constants.Database).Collection(constants.ProductsCollection).Find(context.TODO(), bson.D{}, findOptions)
	if err != nil {
		return nil, err
	}

	var products []domain.Product
	err = cursor.All(context.Background(), &products)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (productRepository *ProductRepositoryImpl) GetProduct(id primitive.ObjectID) (domain.Product, error) {
	var product domain.Product

	filter := bson.D{{Key: "_id", Value: id}}

	err := productRepository.mongoClient.Database(constants.Database).Collection(constants.ProductsCollection).FindOne(context.TODO(), filter).Decode(&product)
	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}

func (productRepository *ProductRepositoryImpl) CreateProduct(product *domain.Product) (domain.Product, error) {
	now := primitive.NewDateTimeFromTime(time.Now())

	product.CreatedAt = now
	product.UpdatedAt = now

	_, err := productRepository.mongoClient.Database(constants.Database).Collection(constants.ProductsCollection).InsertOne(context.TODO(), product)
	if err != nil {
		return domain.Product{}, err
	}

	return *product, nil
}

func (productRepository *ProductRepositoryImpl) UpdateProduct(id primitive.ObjectID, product *domain.Product) (int64, error) {
	now := primitive.NewDateTimeFromTime(time.Now())

	product.UpdatedAt = now

	filter := bson.M{"_id": id}
	saveData := bson.M{"$set": product}

	updateResult, err := productRepository.mongoClient.Database(constants.Database).Collection(constants.ProductsCollection).UpdateOne(context.TODO(), filter, saveData)
	if err != nil {
		return 0, err
	}
	if updateResult.ModifiedCount == 0 {
		return 0, errors.New("no product updated")
	}

	return updateResult.ModifiedCount, nil
}

func (productRepository *ProductRepositoryImpl) DeleteProduct(id primitive.ObjectID) (int64, error) {
	filter := bson.M{"_id": id}

	deleteResult, err := productRepository.mongoClient.Database(constants.Database).Collection(constants.ProductsCollection).DeleteOne(context.TODO(), filter)
	if err != nil {
		return 0, err
	}
	if deleteResult.DeletedCount == 0 {
		return 0, errors.New("no product deleted")
	}

	return deleteResult.DeletedCount, nil
}
