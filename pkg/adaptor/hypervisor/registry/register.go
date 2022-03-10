package registry

import (
	"github.com/confidential-containers/cloud-api-adapter/pkg/podnetwork"
	"github.com/confidential-containers/cloud-api-adapter/pkg/adaptor/hypervisor"
)


func NewServer(cfg interface{}, workerNode podnetwork.WorkerNode, daemonPort string) hypervisor.Server {
	return newServer(cfg, workerNode, daemonPort)
}