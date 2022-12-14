package services

import (
	client "admin/services/repositories"
	e "admin/utils/errors"
)

type DockerServiceImpl struct {
	docker *client.DockerClient
}

func NewDockerServiceImpl(
	docker *client.DockerClient,
) *DockerServiceImpl {
	return &DockerServiceImpl{
		docker: docker,
	}
}

func (s *DockerServiceImpl) CreateContainer(image string, name string) (string, e.ApiError) {
	id, err := s.docker.CreateNewContainer(image, name)
	if err != nil {
		return id, e.NewInternalServerApiError("Error creating container", err)
	}
	return id, nil
}
