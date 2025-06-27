package main

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 声明接口
type UserClient interface {
	GetUser(id string) (string, error)
}

// 实现接口的结构体
type userClient struct {
	baseURL string
	client  *http.Client
}

func (u *userClient) GetUser(id string) (string, error) {
	req, err := http.NewRequest("GET", u.baseURL+"/user?id="+id, nil)
	if err != nil {
		return "", err
	}
	resp, err := u.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// 构造函数
func NewUserClient(baseURL string) UserClient {
	return &userClient{
		baseURL: baseURL,
		client:  &http.Client{},
	}
}

func main() {
	// 启动 Gin 服务，添加 /user 路由
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		id := c.Query("id")
		c.String(http.StatusOK, "User ID: %s", id)
	})
	r.Run(":8080")
}
