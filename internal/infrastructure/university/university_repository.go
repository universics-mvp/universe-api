package university_infrastructure

import (
	"encoding/json"
	"io"
	"net/http"

	"main/internal/config"
)

type UniversityRepository struct {
	env config.Env
}

func NewUniversityRepository(env config.Env) UniversityRepository {
	return UniversityRepository{
		env: env,
	}
}

func (u UniversityRepository) GetUsers() ([]*StudentDTO, error) {
	request, err := http.NewRequest(http.MethodGet, u.env.ApiURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	respBody, err := io.ReadAll(resp.Body)

	var studentResp []*StudentDTO

	err = json.Unmarshal(respBody, &studentResp)
	if err != nil {
		return nil, err
	}

	return studentResp, nil
}
