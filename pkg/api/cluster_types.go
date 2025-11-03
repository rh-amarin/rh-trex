package api

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Cluster struct {
	Meta
	Kind string
	// Cluster name (unique)
	Name string
	// Cluster specification CLM doesn't know how to unmarshall the spec, it only stores and forwards to adapters to do their job But CLM will validate the schema before accepting the request
	Spec datatypes.JSON `gorm:"type:jsonb"`
	// Resource URI
	Href string
	// labels for the API resource as pairs of name:value strings
	Labels datatypes.JSON `gorm:"type:jsonb"`
	// Generation field is updated on customer updates, reflecting the version of the \"intent\" of the customer
	Generation int32
	// Status stores the cluster status aggregation from all adapters
	Status    ClusterStatus
	CreatedBy string
	UpdatedBy string
}

type (
	ClusterList  []*Cluster
	ClusterIndex map[string]*Cluster
)

func (l ClusterList) Index() ClusterIndex {
	index := ClusterIndex{}
	for _, o := range l {
		index[o.ID] = o
	}
	return index
}

func (d *Cluster) BeforeCreate(tx *gorm.DB) error {
	d.ID = NewID()
	return nil
}
