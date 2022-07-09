package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/gommon/log"
	"net/http"
	"time"
)

type GHService struct {
	baseUrl string
	client  *http.Client
}

func NewGHService(client *http.Client, baseUrl string) *GHService {
	return &GHService{
		client:  client,
		baseUrl: baseUrl,
	}
}

func (s *GHService) GetUser(user string) (*User, error) {
	url := fmt.Sprintf("%s/users/%s", s.baseUrl, user)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(req.Context(), 3*time.Second)
	defer cancel()

	req = req.WithContext(ctx)
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	res, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var resp User
	err = json.NewDecoder(res.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	log.Info(resp)
	return &resp, nil
}

type User struct {
	AvatarUrl         string      `json:"avatar_url"`
	Bio               string      `json:"bio"`
	Blog              string      `json:"blog"`
	Company           string      `json:"company"`
	CreatedAt         time.Time   `json:"created_at"`
	Email             interface{} `json:"email"`
	EventsUrl         string      `json:"events_url"`
	Followers         int         `json:"followers"`
	FollowersUrl      string      `json:"followers_url"`
	Following         int         `json:"following"`
	FollowingUrl      string      `json:"following_url"`
	GistsUrl          string      `json:"gists_url"`
	GravatarId        string      `json:"gravatar_id"`
	Hireable          bool        `json:"hireable"`
	HtmlUrl           string      `json:"html_url"`
	Id                int         `json:"id"`
	Location          string      `json:"location"`
	Login             string      `json:"login"`
	Name              string      `json:"name"`
	NodeId            string      `json:"node_id"`
	OrganizationsUrl  string      `json:"organizations_url"`
	PublicGists       int         `json:"public_gists"`
	PublicRepos       int         `json:"public_repos"`
	ReceivedEventsUrl string      `json:"received_events_url"`
	ReposUrl          string      `json:"repos_url"`
	SiteAdmin         bool        `json:"site_admin"`
	StarredUrl        string      `json:"starred_url"`
	SubscriptionsUrl  string      `json:"subscriptions_url"`
	TwitterUsername   interface{} `json:"twitter_username"`
	Type              string      `json:"type"`
	UpdatedAt         time.Time   `json:"updated_at"`
	Url               string      `json:"url"`
}
