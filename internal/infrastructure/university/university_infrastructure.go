package university_infrastructure

import (
	"encoding/json"
	"io"
	"main/internal/config"
	"net/http"
)

type UniversityInfrastructure struct {
	env config.Env
}

func NewUniversityInfrastructure (env config.Env) UniversityInfrastructure {
	return UniversityInfrastructure{
		env: env,
	}
}

func (u UniversityInfrastructure) GetUsers () ([]*StudentDTO, error) {
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