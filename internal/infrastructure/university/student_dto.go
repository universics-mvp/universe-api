package university_infrastructure

type StudentDTO struct {
	Id       string `json:"id"`
	GroupId  string `json:"group_id"`
	FullName string `json:"full_name"`
}
