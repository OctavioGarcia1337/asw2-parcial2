package services

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"worker/config"
	client "worker/services/repositories"
)

type WorkerService struct {
	queue *client.QueueClient
}

func NewWorker(
	queue *client.QueueClient,
) *WorkerService {
	return &WorkerService{
		queue: queue,
	}
}

func (s *WorkerService) QueueWorker(qname string) {
	s.queue.ProcessMessages(qname, func(id string) {
		resp, err := http.Get(fmt.Sprintf("http://%s:%d/items/%s", config.BUSQUEDAHOST, config.BUSQUEDAPORT, id))
		log.Debug("Item sent " + id)
		if err != nil {
			log.Debug("error in get request")
			log.Debug(resp)
		}
	})
}
