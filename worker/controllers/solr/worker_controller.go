package worker

import (
	"worker/config"
	"worker/services"
	client "worker/services/repositories"
	con "worker/utils/connections"
)

var (
	Worker = services.NewWorker(
		(*client.QueueClient)(con.NewQueueClient(config.RABBITUSER, config.RABBITPASSWORD, config.RABBITHOST, config.RABBITPORT)),
	)
)

func StartWorker() {

	Worker.QueueWorker("solr")

}
