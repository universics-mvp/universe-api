package challenge_infrastructure

type ChallengeSchema struct {
	ID          string   `json:"id" bson:"_id"`
	Title       string   `json:"title" bson:"title"`
	Description string   `json:"description" bson:"description"`
	CreatorID   string   `json:"creator_id" bson:"creator_id"`
	Groups      []string `json:"groups" bson:"groups"`
}
