package data

import (
	"context"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/logging"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/utilities"
	"github.com/JECSand/go-cqrs-service-boilerplate/query_service/config"
	"github.com/JECSand/go-cqrs-service-boilerplate/query_service/core/models"
	"github.com/gofrs/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoRepository struct {
	log logging.Logger
	cfg *config.Config
	db  *mongo.Client
}

func NewMongoRepository(log logging.Logger, cfg *config.Config, db *mongo.Client) *mongoRepository {
	return &mongoRepository{
		log: log,
		cfg: cfg,
		db:  db,
	}
}

func (p *mongoRepository) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "mongoRepository.CreateUser")
	defer span.Finish()
	collection := p.db.Database(p.cfg.Mongo.DB).Collection(p.cfg.MongoCollections.Users)
	_, err := collection.InsertOne(ctx, user, &options.InsertOneOptions{})
	if err != nil {
		p.traceErr(span, err)
		return nil, errors.Wrap(err, "InsertOne")
	}
	return user, nil
}

func (p *mongoRepository) UpdateUser(ctx context.Context, user *models.User) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "mongoRepository.UpdateUser")
	defer span.Finish()
	collection := p.db.Database(p.cfg.Mongo.DB).Collection(p.cfg.MongoCollections.Users)
	ops := options.FindOneAndUpdate()
	ops.SetReturnDocument(options.After)
	ops.SetUpsert(true)
	var updated models.User
	if err := collection.FindOneAndUpdate(ctx, bson.M{"_id": user.ID}, bson.M{"$set": user}, ops).Decode(&updated); err != nil {
		p.traceErr(span, err)
		return nil, errors.Wrap(err, "Decode")
	}
	return &updated, nil
}

func (p *mongoRepository) GetUserById(ctx context.Context, id uuid.UUID) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "mongoRepository.GetUserById")
	defer span.Finish()
	collection := p.db.Database(p.cfg.Mongo.DB).Collection(p.cfg.MongoCollections.Users)
	var user models.User
	if err := collection.FindOne(ctx, bson.M{"_id": id.String()}).Decode(&user); err != nil {
		p.traceErr(span, err)
		return nil, errors.Wrap(err, "Decode")
	}
	return &user, nil
}

func (p *mongoRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "mongoRepository.DeleteUser")
	defer span.Finish()
	collection := p.db.Database(p.cfg.Mongo.DB).Collection(p.cfg.MongoCollections.Users)
	return collection.FindOneAndDelete(ctx, bson.M{"_id": id.String()}).Err()
}

func (p *mongoRepository) Search(ctx context.Context, search string, pagination *utilities.Pagination) (*models.UsersList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "mongoRepository.Search")
	defer span.Finish()
	collection := p.db.Database(p.cfg.Mongo.DB).Collection(p.cfg.MongoCollections.Users)
	filter := bson.D{
		{Key: "$or", Value: bson.A{
			bson.D{{Key: "name", Value: primitive.Regex{Pattern: search, Options: "gi"}}},
			bson.D{{Key: "description", Value: primitive.Regex{Pattern: search, Options: "gi"}}},
		}},
	}
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		p.traceErr(span, err)
		return nil, errors.Wrap(err, "CountDocuments")
	}
	if count == 0 {
		return &models.UsersList{Users: make([]*models.User, 0)}, nil
	}
	limit := int64(pagination.GetLimit())
	skip := int64(pagination.GetOffset())
	cursor, err := collection.Find(ctx, filter, &options.FindOptions{
		Limit: &limit,
		Skip:  &skip,
	})
	if err != nil {
		p.traceErr(span, err)
		return nil, errors.Wrap(err, "Find")
	}
	defer cursor.Close(ctx) // nolint: errcheck
	users := make([]*models.User, 0, pagination.GetSize())
	for cursor.Next(ctx) {
		var prod models.User
		if err = cursor.Decode(&prod); err != nil {
			p.traceErr(span, err)
			return nil, errors.Wrap(err, "Find")
		}
		users = append(users, &prod)
	}
	if err = cursor.Err(); err != nil {
		span.SetTag("error", true)
		span.LogKV("error_code", err.Error())
		return nil, errors.Wrap(err, "cursor.Err")
	}
	return models.NewUserListWithPagination(users, count, pagination), nil
}

func (p *mongoRepository) traceErr(span opentracing.Span, err error) {
	span.SetTag("error", true)
	span.LogKV("error_code", err.Error())
}
