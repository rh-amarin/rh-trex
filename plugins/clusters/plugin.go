package clusters

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/openshift-online/rh-trex/cmd/trex/environments"
	"github.com/openshift-online/rh-trex/cmd/trex/environments/registry"
	"github.com/openshift-online/rh-trex/cmd/trex/server"
	"github.com/openshift-online/rh-trex/pkg/api"
	"github.com/openshift-online/rh-trex/pkg/api/presenters"
	"github.com/openshift-online/rh-trex/pkg/auth"
	"github.com/openshift-online/rh-trex/pkg/controllers"
	"github.com/openshift-online/rh-trex/pkg/dao"
	"github.com/openshift-online/rh-trex/pkg/db"
	"github.com/openshift-online/rh-trex/pkg/handlers"
	"github.com/openshift-online/rh-trex/pkg/services"
	"github.com/openshift-online/rh-trex/plugins/events"
	"github.com/openshift-online/rh-trex/plugins/generic"
)

// ServiceLocator Service Locator
type ServiceLocator func() services.ClusterService

func NewServiceLocator(env *environments.Env) ServiceLocator {
	return func() services.ClusterService {
		return services.NewClusterService(
			db.NewAdvisoryLockFactory(env.Database.SessionFactory),
			dao.NewClusterDao(&env.Database.SessionFactory),
			events.Service(&env.Services),
		)
	}
}

// Service helper function to get the cluster service from the registry
func Service(s *environments.Services) services.ClusterService {
	if s == nil {
		return nil
	}
	if obj := s.GetService("Clusters"); obj != nil {
		locator := obj.(ServiceLocator)
		return locator()
	}
	return nil
}

func init() {
	// Service registration
	registry.RegisterService("Clusters", func(env interface{}) interface{} {
		return NewServiceLocator(env.(*environments.Env))
	})

	// Routes registration
	server.RegisterRoutes("clusters", func(apiV1Router *mux.Router, services server.ServicesInterface, authMiddleware auth.JWTMiddleware, authzMiddleware auth.AuthorizationMiddleware) {
		envServices := services.(*environments.Services)
		clusterHandler := handlers.NewClusterHandler(Service(envServices), generic.Service(envServices))

		clustersRouter := apiV1Router.PathPrefix("/clusters").Subrouter()
		clustersRouter.HandleFunc("", clusterHandler.List).Methods(http.MethodGet)
		clustersRouter.HandleFunc("/{id}", clusterHandler.Get).Methods(http.MethodGet)
		clustersRouter.HandleFunc("", clusterHandler.Create).Methods(http.MethodPost)
		clustersRouter.HandleFunc("/{id}", clusterHandler.Patch).Methods(http.MethodPatch)
		clustersRouter.HandleFunc("/{id}", clusterHandler.Delete).Methods(http.MethodDelete)
		clustersRouter.Use(authMiddleware.AuthenticateAccountJWT)
		clustersRouter.Use(authzMiddleware.AuthorizeApi)
	})

	// Controller registration
	server.RegisterController("Clusters", func(manager *controllers.KindControllerManager, services *environments.Services) {
		dinoServices := Service(services)

		manager.Add(&controllers.ControllerConfig{
			Source: "Clusters",
			Handlers: map[api.EventType][]controllers.ControllerHandlerFunc{
				api.CreateEventType: {dinoServices.OnUpsert},
				api.UpdateEventType: {dinoServices.OnUpsert},
				api.DeleteEventType: {dinoServices.OnDelete},
			},
		})
	})

	// Presenter registration
	presenters.RegisterPath(api.Cluster{}, "clusters")
	presenters.RegisterPath(&api.Cluster{}, "clusters")
	presenters.RegisterKind(api.Cluster{}, "Cluster")
	presenters.RegisterKind(&api.Cluster{}, "Cluster")
}
