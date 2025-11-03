package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/openshift-online/rh-trex/pkg/api"
	"github.com/openshift-online/rh-trex/pkg/api/openapi"
	"github.com/openshift-online/rh-trex/pkg/api/presenters"
	"github.com/openshift-online/rh-trex/pkg/errors"
	"github.com/openshift-online/rh-trex/pkg/services"
)

var _ RestHandler = clusterHandler{}

type clusterHandler struct {
	cluster services.ClusterService
	generic services.GenericService
}

func NewClusterHandler(cluster services.ClusterService, generic services.GenericService) *clusterHandler {
	return &clusterHandler{
		cluster: cluster,
		generic: generic,
	}
}

func (h clusterHandler) Create(w http.ResponseWriter, r *http.Request) {
	var clusterReq openapi.ClusterCreateRequest
	cfg := &handlerConfig{
		&clusterReq,
		[]validate{
			validateEmpty(&clusterReq, "Id", "id"),
			validateNotEmpty(&clusterReq, "Species", "species"),
		},
		func() (interface{}, *errors.ServiceError) {
			ctx := r.Context()
			cluster := presenters.ConvertCluster(clusterReq)
			cluster, err := h.cluster.Create(ctx, cluster)
			if err != nil {
				return nil, err
			}
			return presenters.PresentCluster(cluster), nil
		},
		handleError,
	}

	handle(w, r, cfg, http.StatusCreated)
}

func (h clusterHandler) List(w http.ResponseWriter, r *http.Request) {
	cfg := &handlerConfig{
		Action: func() (interface{}, *errors.ServiceError) {
			ctx := r.Context()

			listArgs := services.NewListArguments(r.URL.Query())
			var clusters []api.Cluster
			paging, err := h.generic.List(ctx, "username", listArgs, &clusters)
			if err != nil {
				return nil, err
			}
			dinoList := openapi.ClusterList{
				Kind:  "ClusterList",
				Page:  int32(paging.Page),
				Size:  int32(paging.Size),
				Total: int32(paging.Total),
				Items: []openapi.Cluster{},
			}

			for _, dino := range clusters {
				converted := presenters.PresentCluster(&dino)
				dinoList.Items = append(dinoList.Items, converted)
			}
			if listArgs.Fields != nil {
				filteredItems, err := presenters.SliceFilter(listArgs.Fields, dinoList.Items)
				if err != nil {
					return nil, err
				}
				return filteredItems, nil
			}
			return dinoList, nil
		},
	}

	handleList(w, r, cfg)
}

func (h clusterHandler) Get(w http.ResponseWriter, r *http.Request) {
	cfg := &handlerConfig{
		Action: func() (interface{}, *errors.ServiceError) {
			id := mux.Vars(r)["id"]
			ctx := r.Context()
			cluster, err := h.cluster.Get(ctx, id)
			if err != nil {
				return nil, err
			}

			return presenters.PresentCluster(cluster), nil
		},
	}

	handleGet(w, r, cfg)
}

func (h clusterHandler) Delete(w http.ResponseWriter, r *http.Request) {
	cfg := &handlerConfig{
		Action: func() (interface{}, *errors.ServiceError) {
			id := mux.Vars(r)["id"]
			ctx := r.Context()
			err := h.cluster.Delete(ctx, id)
			if err != nil {
				return nil, err
			}
			return nil, nil
		},
	}
	handleDelete(w, r, cfg, http.StatusNoContent)
}

func (h clusterHandler) Patch(w http.ResponseWriter, r *http.Request) {
}
