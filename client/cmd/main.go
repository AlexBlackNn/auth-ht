package main

import (
	"bufio"
	"context"
	"fmt"
	log "log/slog"
	"os"

	"github.com/AlexBlackNn/authloyalty/client/internal/client"
	"github.com/AlexBlackNn/authloyalty/client/internal/dto"
	"github.com/looplab/fsm"
)

func main() {
	serviceClient := client.New("http://localhost:8000")

	fsm := fsm.NewFSM(
		"start",
		fsm.Events{
			{Name: "start", Src: []string{"start"}, Dst: "info"},
			{Name: "info", Src: []string{"start"}, Dst: "info"},
		},
		fsm.Callbacks{},
	)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("1. Registration")
		fmt.Println("2. Login")
		fmt.Println("3. Exit")

		fmt.Print("Choose an option: ")
		scanner.Scan()
		option := scanner.Text()

		switch option {
		case "1":
			if fsm.Current() == "start" {
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
			}

			err := fsm.Event(context.Background(), "info")
			if err != nil {
				fmt.Println(err)
			}

			if fsm.Current() == "info" {
				err := serviceClient.GetInfo()
				if err != nil {
					log.Error(err.Error())
				} else {
					log.Info("Got info successfully")
				}
			}

		case "2":
			if fsm.Current() == "start" {
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
			}

			err := fsm.Event(context.Background(), "info")
			if err != nil {
				fmt.Println(err)
			}
			if fsm.Current() == "info" {
				err := serviceClient.GetInfo()
				if err != nil {
					log.Error(err.Error())
				} else {
					log.Info("Got info successfully")
				}
			}
		case "3":
			return
		default:
			log.Info("Invalid option")
		}
	}
}
