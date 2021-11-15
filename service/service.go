package service

import (
	"fmt"
	"os"
	"parts/repository"
)

type Service struct {
	repo *repository.Repository
}

func Create() (*Service, error) {
	repo, err := repository.Create()

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return nil, err
	}

	fmt.Println("Repository acquired")

	return &Service{repo: repo}, err
}
