package main

import (
	"context"
	"fmt"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	log "github.com/sirupsen/logrus"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	// List all containers on the Compose network
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	// Print the health and run status of each cont
	for _, cont := range containers {
		health, err := cli.ContainerInspect(context.Background(), cont.ID)
		if err != nil {
			log.Fatal(err)
		}
		if health.State.Health != nil {
			if health.State.Health.Status == "healthy" {
				fmt.Printf("Container %s is healthy\n", cont.ID)
			} else {
				fmt.Printf("Container %s is not healthy\n", cont.ID)
			}
		}

		if cont.State == "running" {
			fmt.Printf("Container %s is running\n", cont.ID)
		} else {
			fmt.Printf("Container %s is not running\n", cont.ID)
		}
	}

	// Check if the "items" cont is running
	itemsRunning := false
	for _, cont := range containers {
		if cont.Names[0] == "items" {
			itemsRunning = true
			break
		}
	}

	// If the "items" cont is not running, start it
	if !itemsRunning {
		fmt.Println("Starting the items cont...")

		resp, err := cli.ContainerCreate(context.Background(), &container.Config{
			Image: "asw2-parcial2-items",
		}, &container.HostConfig{
			PortBindings: nat.PortMap{
				"8090/tcp": []nat.PortBinding{
					{
						HostIP:   "0.0.0.0",
						HostPort: "8090",
					},
				},
			},
		}, nil, &v1.Platform{}, "items")
		if err != nil {
			log.Fatal(err)
		}

		if err := cli.ContainerStart(context.Background(), resp.ID, types.ContainerStartOptions{}); err != nil {
			log.Fatal(err)
		}

		fmt.Println("The items cont has been started.")
	}

	// Stop the "items" cont after 5 seconds
	time.Sleep(5 * time.Second)

	for _, cont := range containers {
		if cont.Names[0] == "items" {
			if err := cli.ContainerStop(context.Background(), cont.ID, nil); err != nil {
				log.Fatal(err)
			}

			fmt.Println("The items cont has been stopped.")
			break
		}
	}

	// Remove the "items" cont
	for _, cont := range containers {
		if cont.Names[0] == "items" {
			if err := cli.ContainerRemove(context.Background(), cont.ID, types.ContainerRemoveOptions{}); err != nil {
				log.Fatal(err)
			}

			fmt.Println("The items cont has been removed.")
			break
		}
	}
}
