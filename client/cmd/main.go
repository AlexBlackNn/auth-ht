package main

import (
	"bufio"
	"fmt"
	log "log/slog"
	"os"

	"github.com/AlexBlackNn/authloyalty/client/internal/dto"
	"github.com/AlexBlackNn/authloyalty/client/internal/services/client"
)

func main() {
	serviceClient := client.New("http://localhost:8000")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("1. Register")
		fmt.Println("2. Get Info")
		fmt.Println("3. Login")
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

			err := serviceClient.Register(rr)
			if err != nil {
				log.Error(err.Error())
			} else {
				log.Info("Registration successful")
			}
		case "2":
			err := serviceClient.GetInfo()
			if err != nil {
				log.Error(err.Error())
			} else {
				log.Info("Got info successfully")
			}
		case "3":

			fmt.Print("Enter email: ")
			scanner.Scan()
			email := scanner.Text()

			fmt.Print("Enter password: ")
			scanner.Scan()
			password := scanner.Text()

			lr := dto.LoginRequest{
				Email:    email,
				Password: password,
			}

			err := serviceClient.Login(lr)
			if err != nil {
				log.Error(err.Error())
			} else {
				log.Info("Login successful")
			}
		case "4":
			return
		default:
			log.Info("Invalid option")
		}
	}
}
