package server

import (
	"testing"

	cache "github.com/envoyproxy/go-control-plane/pkg/cache/v2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSDSServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SDS Server Suite")
}

// getSnapshotCache is a utility func to expose
// the snapshot cache for test assertions
func GetSnapshotCache(srv Server) cache.SnapshotCache {
	return srv.snapshotCache
}
