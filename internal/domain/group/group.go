package group

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Group struct {
	ID        *primitive.ObjectID
	ChatID    int64
	CuratorID int64
}

func CreateGroup(ChatID int64, CuratorID int64) Group {
	return Group{ChatID: ChatID, CuratorID: CuratorID}
}
