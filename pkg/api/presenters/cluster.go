package presenters

import (
	"github.com/openshift-online/rh-trex/pkg/api"
	"github.com/openshift-online/rh-trex/pkg/api/openapi"
	"github.com/openshift-online/rh-trex/pkg/util"
)

func ConvertCluster(cluster openapi.ClusterCreateRequest) *api.Cluster {
	return &api.Cluster{
		Meta: api.Meta{
			ID: util.NilToEmptyString(cluster.Id),
		},
	}
}

func PresentCluster(cluster *api.Cluster) openapi.Cluster {
	reference := PresentReference(cluster.ID, cluster)
	return openapi.Cluster{
		Id:   reference.Id,
		Kind: *reference.Kind,
		Href: reference.Href,
	}
}
