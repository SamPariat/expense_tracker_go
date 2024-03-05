package repositories

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	constants "github.com/SamPariat/expenses-tracker/common"
	"github.com/SamPariat/expenses-tracker/internal/core/domain"
	"github.com/SamPariat/expenses-tracker/internal/core/ports"
)

type UserRepositoryImpl struct {
	mongoClient *mongo.Client
}

func NewUserRepository(mongoClient *mongo.Client) ports.UserRepository {
	return &UserRepositoryImpl{
		mongoClient: mongoClient,
	}
}

func (userRepository *UserRepositoryImpl) GetUsers(page int64, limit int64) ([]domain.User, error) {
	cursor, err := userRepository.mongoClient.Database(constants.Database).Collection(constants.UsersCollection).Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	var users []domain.User

	err = cursor.All(context.Background(), &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (userRepository *UserRepositoryImpl) GetUser(id primitive.ObjectID) (domain.User, error) {
	filter := bson.M{"_id": id}

	var user domain.User

	err := userRepository.mongoClient.Database(constants.Database).Collection(constants.UsersCollection).FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (userRepository *UserRepositoryImpl) CreateUser(user *domain.User) (domain.User, error) {
	user.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	user.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	savedResult, err := userRepository.mongoClient.Database(constants.Database).Collection(constants.UsersCollection).InsertOne(context.TODO(), user)
	if err != nil {
		return domain.User{}, err
	}

	user.Id = savedResult.InsertedID.(primitive.ObjectID)

	return *user, nil
}

func (userRepository *UserRepositoryImpl) UpdateUser(id primitive.ObjectID, user *domain.User) (int64, error) {
	user.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	filter := bson.M{"_id": id}
	update := bson.M{"$set": user}

	updateResult, err := userRepository.mongoClient.Database(constants.Database).Collection(constants.UsersCollection).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return 0, err
	}

	return updateResult.ModifiedCount, nil
}

func (userRepository *UserRepositoryImpl) DeleteUser(id primitive.ObjectID) (int64, error) {
	filter := bson.M{"_id": id}

	deletedResult, err := userRepository.mongoClient.Database(constants.Database).Collection(constants.UsersCollection).DeleteOne(context.TODO(), filter)
	if err != nil {
		return 0, err
	}

	return deletedResult.DeletedCount, nil
}
