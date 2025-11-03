package factories

import (
	"context"
	"fmt"

	"github.com/openshift-online/rh-trex/cmd/trex/environments"
	"github.com/openshift-online/rh-trex/pkg/api"
	"github.com/openshift-online/rh-trex/plugins/clusters"
)

func (f *Factories) NewCluster(species string) (*api.Cluster, error) {
	dinoService := clusters.Service(&environments.Environment().Services)

	cluster := &api.Cluster{
		// TODO: angel
	}

	dino, err := dinoService.Create(context.Background(), cluster)
	if err != nil {
		return nil, err
	}

	return dino, nil
}

func (f *Factories) NewClusterList(namePrefix string, count int) ([]*api.Cluster, error) {
	var clusters []*api.Cluster
	for i := 1; i <= count; i++ {
		name := fmt.Sprintf("%s_%d", namePrefix, i)
		c, _ := f.NewCluster(name)
		clusters = append(clusters, c)
	}
	return clusters, nil
}
