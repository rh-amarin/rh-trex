package integration

import (
	"context"
	"net/http"
	"testing"

	"github.com/openshift-online/rh-trex/pkg/api"

	. "github.com/onsi/gomega"

	"github.com/openshift-online/rh-trex/test"
)

func TestClusterGet(t *testing.T) {
	t.Log("step1")
	h, client := test.RegisterIntegration(t)

	account := h.NewRandAccount()
	ctx := h.NewAuthenticatedContext(account)

	t.Log("step2")
	// 401 using no JWT token
	_, _, err := client.DefaultAPI.GetClusterById(context.Background(), "foo").Execute()
	Expect(err).To(HaveOccurred(), "Expected 401 but got nil error")
	t.Log("step3")

	// GET responses per openapi spec: 200 and 404,
	_, resp, err := client.DefaultAPI.GetClusterById(ctx, "foo").Execute()
	t.Log("step3.1")
	Expect(err).To(HaveOccurred(), "Expected 404")
	if resp == nil {
		t.Log("resp is Nil")
	}
	Expect(resp.StatusCode).To(Equal(http.StatusNotFound))
	t.Log("step4")

	clusterNew, err := h.Factories.NewCluster(h.NewID())
	Expect(err).NotTo(HaveOccurred())
	t.Log("step5")

	cluster, resp, err := client.DefaultAPI.GetClusterById(ctx, clusterNew.ID).Execute()
	Expect(err).NotTo(HaveOccurred())
	Expect(resp.StatusCode).To(Equal(http.StatusOK))

	Expect(cluster.Cluster.GetId()).To(Equal(clusterNew.ID), "found object does not match test object")
}

func contains(et api.EventType, events api.EventList) bool {
	for _, e := range events {
		if e.EventType == et {
			return true
		}
	}
	return false
}
