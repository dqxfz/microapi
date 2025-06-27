package feign

import (
	"context"
	"fmt"
	"github.com/goccy/go-json"
	"net/http"
)

// UserService 定义用户服务的接口
type UserService interface {
	GetUser(ctx context.Context, userID string) (*User, error)
}

// User 定义用户数据结构
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// NewUserService 创建一个新的 UserService 实例
func NewUserService(baseURL string) UserService {
	return &userService{
		baseURL: baseURL,
		client:  &http.Client{},
	}
}

type userService struct {
	baseURL string
	client  *http.Client
}

func (s *userService) GetUser(ctx context.Context, userID string) (*User, error) {
	url := s.baseURL + "/users/" + userID

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
