package group_infrastructure

import (
	"context"

	"main/internal/domain/group"
	"main/pkg"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type GroupRepository struct {
	collection *mongo.Collection
	logger     pkg.Logger
}

func NewGroupRepository(db pkg.MongoDatabase, logger pkg.Logger) group.GroupRepository {
	return GroupRepository{
		collection: db.Collection("groups"),
		logger:     logger,
	}
}

func (repo GroupRepository) Save(group *group.Group) (*group.Group, error) {
	if group.ID == nil {
		return repo.create(group)
	}
	return repo.update(group)
}

func (repo GroupRepository) create(group *group.Group) (*group.Group, error) {
	newId := primitive.NewObjectID()
	group.ID = &newId
	_, err := repo.collection.InsertOne(context.Background(), mapToSchema(*group))
	if err != nil {
		return nil, err
	}

	return group, nil
}

func (repo GroupRepository) update(sess *group.Group) (*group.Group, error) {
	_, err := repo.collection.ReplaceOne(context.Background(), bson.M{"_id": sess.ID}, mapToSchema(*sess))
	if err != nil {
		return nil, err
	}

	return sess, nil
}

func (c GroupRepository) GetForChat(chatId int64) (*group.Group, error) {
	var schema GroupSchema
	err := c.collection.FindOne(context.Background(), bson.M{"chat_id": chatId}).Decode(&schema)
	if err != nil {
		return nil, err
	}
	entity := mapShema(schema)
	return &entity, nil
}

func (c GroupRepository) GetForCurator(curatorId int64) (*group.Group, error) {
	var schema GroupSchema
	err := c.collection.FindOne(context.Background(), bson.M{"curator_id": curatorId}).Decode(&schema)
	if err != nil {
		return nil, err
	}
	entity := mapShema(schema)
	return &entity, nil
}

func (gr GroupRepository) List() ([]group.Group, error) {
	cur, err := gr.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var result []GroupSchema

	err = cur.All(context.Background(), &result)
	if err != nil {
		return nil, err
	}
	var groups = make([]group.Group, 0)
	for _, m := range result {
		groups = append(groups, mapShema(m))
	}
	return groups, nil
}

func mapShema(schema GroupSchema) group.Group {
	return group.Group{
		ID:        &schema.ID,
		ChatID:    schema.ChatID,
		CuratorID: schema.CuratorID,
		Title:     schema.Title,
	}
}

func mapToSchema(group group.Group) GroupSchema {
	return GroupSchema{
		ID:        *group.ID,
		ChatID:    group.ChatID,
		CuratorID: group.CuratorID,
		Title:     group.Title,
	}
}
