package service

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	log "log/slog"
	"os"

	"github.com/AlexBlackNn/authloyalty/client/internal/client"
	"github.com/AlexBlackNn/authloyalty/client/internal/config"
	"github.com/AlexBlackNn/authloyalty/client/internal/dto"
	"github.com/looplab/fsm"
)

type Service struct {
	serviceClient *client.Client
	log           *log.Logger
	cfg           *config.Config
}

func New(cfg *config.Config, log *log.Logger) *Service {
	return &Service{
		log:           log,
		cfg:           cfg,
		serviceClient: &client.Client{Addr: "http://localhost:8000"}}
}

func (s *Service) Start() error {
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

				err := s.serviceClient.Register(rr)
				if err != nil {
					s.log.Error(err.Error())
					return err
				} else {
					s.log.Info("Registration successful")
				}
			}

			err := fsm.Event(context.Background(), "info")
			if err != nil {
				s.log.Error(err.Error())
				return err
			}

			if fsm.Current() == "info" {
				err := s.serviceClient.GetInfo()
				if err != nil {
					s.log.Error(err.Error())
					return err
				} else {
					s.log.Info("Got info successfully")
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

				err := s.serviceClient.Login(lr)
				if err != nil {
					s.log.Error(err.Error())
					return err
				} else {
					s.log.Info("Login successful")
				}
			}

			err := fsm.Event(context.Background(), "info")
			if err != nil {
				s.log.Error(err.Error())
			}

			if fsm.Current() == "info" {
				err := s.serviceClient.GetInfo()
				if err != nil {
					s.log.Error(err.Error())
					return err
				} else {
					s.log.Info("Got info successfully")
				}
			}

		case "3":
			return nil
		default:
			s.log.Info("Invalid option")
			return errors.New("invalid option")
		}
	}
}
