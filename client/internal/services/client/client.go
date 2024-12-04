package client

import (
	"encoding/json"
	"fmt"
	log "log/slog"
	"net/http"
	"time"

	"github.com/AlexBlackNn/authloyalty/client/internal/domain"
	"github.com/AlexBlackNn/authloyalty/client/internal/dto"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	addr string
	domain.UserInfo
}

func New(addr string) *Client {
	return &Client{addr: addr}
}

func (c *Client) Register(rr dto.RegisterRequest) error {
	jsonData, err := json.Marshal(rr)
	if err != nil {
		return err
	}

	restyClient := resty.New()
	restyClient.
		SetRetryCount(10).
		SetRetryWaitTime(10 * time.Second).
		SetRetryMaxWaitTime(5 * time.Millisecond)

	resp, err := restyClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(jsonData).
		Post(c.addr + "/auth/registration")

	defer resp.Body()

	log.Info("http request finished successfully",
		"url", c.addr,
		"statusCode", resp.StatusCode(),
		"body", string(resp.Body()),
	)

	serviceResponse := dto.Response{}
	err = json.Unmarshal(resp.Body(), &serviceResponse)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	if resp.StatusCode() != http.StatusCreated {
		return fmt.Errorf("registration failed with status code %d", resp.StatusCode)
	}

	c.SetAccessToken(serviceResponse.AccessToken).
		SetRefreshToken(serviceResponse.RefreshToken).
		SetUserId(serviceResponse.UserID)

	return nil
}

func (c *Client) GetInfo() error {
	restyClient := resty.New()
	restyClient.
		SetRetryCount(10).
		SetRetryWaitTime(10 * time.Second).
		SetRetryMaxWaitTime(5 * time.Millisecond)

	resp, err := restyClient.R().
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", "Bearer "+c.AccessToken).
		Get(c.addr + "/auth/info")

	defer resp.Body()

	log.Info("http request finished successfully",
		"url", c.addr,
		"statusCode", resp.StatusCode(),
		"body", string(resp.Body()),
	)

	serviceResponse := dto.UserResponse{}
	err = json.Unmarshal(resp.Body(), &serviceResponse)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("get info failed with status code %d", resp.StatusCode())
	}

	c.SetAvatar(serviceResponse.Avatar).
		SetEmail(serviceResponse.Email).
		SetName(serviceResponse.Name)
	return nil
}
