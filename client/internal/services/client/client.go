package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Client struct {
	addr string
}

func New(addr string) *Client {
	return &Client{addr: addr}
}

type RegisterRequest struct {
	Avatar   string `json:"avatar"`
	Birthday string `json:"birthday"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (c *Client) Register(rr RegisterRequest) error {
	jsonData, err := json.Marshal(rr)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		"POST", c.addr+"/auth/registration",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("registration failed with status code %d", resp.StatusCode)
	}

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

			rr := RegisterRequest{
				Name:     name,
				Email:    email,
				Birthday: birthday,
				Password: password,
				Avatar:   avatar,
			}

			err := client.Register(rr)
			if err != nil {
				log.Println(err)
			} else {
				log.Println("Registration successful")
			}
		case "2":
			fmt.Print("Enter token: ")
			scanner.Scan()
			token := scanner.Text()

			err := client.GetInfo(token)
			if err != nil {
				log.Println(err)
			} else {
				log.Println("Got info successfully")
			}
		case "3":
			return
		default:
			log.Println("Invalid option")
		}
	}
}
