package slack

import (
	"net/url"
)

type AuthService struct {
	api *SlackClient
}

type Auth struct {
	UserId 	 string  `json:"user_id"`
	Username string	 `json:"user"`
	Team 	 string  `json:"team"`
	TeamId 	 string  `json:"team_id"`
	TeamUrl  url.URL
}

func (s *AuthService) Test() (*Auth, error) {

	req, _ := s.api.NewRequest(_GET, "auth.test", nil)

	type authResp struct {
		Auth
		Url string
	}

	auth := new(authResp)

	_, err := s.api.Do(req, auth)

	if err != nil {
		return nil, err
	}

	u, _ := url.Parse(auth.Url)
	auth.TeamUrl = *u;

	return &auth.Auth, nil
}
