package group

type GroupRepository interface {
	GetForChat(chatId int64) (*Group, error)
	GetForCurator(curatorId int64) (*Group, error)
	List() ([]Group, error)
}
