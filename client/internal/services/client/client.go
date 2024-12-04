package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	log "log/slog"
	"net/http"
	"os"
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

	fmt.Println(c)
	return nil
}

func (c *Client) GetInfo(token string) error {
	req, err := http.NewRequest("GET", c.addr+"/auth/info", nil)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("get info failed with status code %d", resp.StatusCode)
	}
	return nil
}

func main() {
	client := New("http://localhost:8000")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("1. Register")
		fmt.Println("2. Get Info")
		fmt.Println("3. Exit")
		fmt.Print("Choose an option: ")
		scanner.Scan()
		option := scanner.Text()

		switch option {
		case "1":
			fmt.Print("Enter username: ")
			scanner.Scan()
			name := scanner.Text()

			fmt.Print("Enter email: ")
			scanner.Scan()
			email := scanner.Text()

			fmt.Print("Enter birthday: ")
			scanner.Scan()
			birthday := scanner.Text()

			fmt.Print("Enter password: ")
			scanner.Scan()
			password := scanner.Text()

			fmt.Print("Enter avatar (base64 encoded): ")
			scanner.Scan()
			avatar := scanner.Text()

			rr := dto.RegisterRequest{
				Name:     name,
				Email:    email,
				Birthday: birthday,
				Password: password,
				Avatar:   avatar,
			}

			err := client.Register(rr)
			if err != nil {
				log.Error(err.Error())
			} else {
				log.Info("Registration successful")
			}
		case "2":
			fmt.Print("Enter token: ")
			scanner.Scan()
			token := scanner.Text()

			err := client.GetInfo(token)
			if err != nil {
				log.Error(err.Error())
			} else {
				log.Info("Got info successfully")
			}
		case "3":
			return
		default:
			log.Info("Invalid option")
		}
	}
}
