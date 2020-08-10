package server

import (
	"context"
	"fmt"
	"hash/fnv"
	"io/ioutil"
	"net"
	"os"
	"time"

	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/go-utils/hashutils"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	auth "github.com/envoyproxy/go-control-plane/envoy/api/v2/auth"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	sds "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v2"
	cache_types "github.com/envoyproxy/go-control-plane/pkg/cache/types"
	cache "github.com/envoyproxy/go-control-plane/pkg/cache/v2"
	server "github.com/envoyproxy/go-control-plane/pkg/server/v2"
)

// These values must match the values in the envoy sidecar's common_tls_context
const (
	// sdsClient         = "sds_client"         // node ID
	serverCert        = "server_cert"        // name of a tls_certificate_sds_secret_config
	validationContext = "validation_context" // name of the validation_context_sds_secret_config
)

var (
	sdsClient   = os.Getenv("GATEWAY_PROXY_POD_NAME") + "." + os.Getenv("GATEWAY_PROXY_POD_NAMESPACE") // node ID
	grpcOptions = []grpc.ServerOption{grpc.MaxConcurrentStreams(10000)}
)

type EnvoyKey struct{}

func (h *EnvoyKey) ID(_ *core.Node) string {
	return sdsClient
}

func SetupEnvoySDS() (*grpc.Server, cache.SnapshotCache) {
	grpcServer := grpc.NewServer(grpcOptions...)
	hasher := &EnvoyKey{}
	snapshotCache := cache.NewSnapshotCache(false, hasher, nil)
	svr := server.NewServer(context.Background(), snapshotCache, nil)

	// register services
	sds.RegisterSecretDiscoveryServiceServer(grpcServer, svr)
	return grpcServer, snapshotCache
}

func RunSDSServer(ctx context.Context, grpcServer *grpc.Server, serverAddress string) (<-chan struct{}, error) {
	lis, err := net.Listen("tcp", serverAddress)
	if err != nil {
		return nil, err
	}
	contextutils.LoggerFrom(ctx).Infof("sds server listening on %s", serverAddress)
	go func() {
		if err = grpcServer.Serve(lis); err != nil {
			contextutils.LoggerFrom(ctx).Fatalw("fatal error in gRPC server", zap.String("address", serverAddress), zap.Error(err))
		}
	}()
	serverStopped := make(chan struct{})
	go func() {
		<-ctx.Done()
		contextutils.LoggerFrom(ctx).Infof("stopping sds server on %s\n", serverAddress)
		grpcServer.GracefulStop()
		serverStopped <- struct{}{}
	}()
	return serverStopped, nil
}

func getSnapshotVersion(key, cert, ca []byte) (string, error) {
	hash, err := hashutils.HashAllSafe(fnv.New64(), key, cert, ca)
	return fmt.Sprintf("%d", hash), err
}

func UpdateSDSConfig(ctx context.Context, sslKeyFile, sslCertFile, sslCaFile string, snapshotCache cache.SnapshotCache) error {

	key, err := readAndVerifyCert(sslKeyFile)
	if err != nil {
		return err
	}
	cert, err := readAndVerifyCert(sslCertFile)
	if err != nil {
		return err
	}
	ca, err := readAndVerifyCert(sslCaFile)
	if err != nil {
		return err
	}
	snapshotVersion, err := getSnapshotVersion(key, cert, ca)
	if err != nil {
		contextutils.LoggerFrom(ctx).Info("Error getting snapshot version", zap.Error(err))
		return err
	}
	contextutils.LoggerFrom(ctx).Infof("Updating SDS config. sdsClient is %s. Snapshot version is %s", sdsClient, snapshotVersion)

	items := []cache_types.Resource{
		serverCertSecret(key, cert),
		validationContextSecret(ca),
	}
	secretSnapshot := cache.Snapshot{}
	secretSnapshot.Resources[cache_types.Secret] = cache.NewResources(snapshotVersion, items)
	return snapshotCache.SetSnapshot(sdsClient, secretSnapshot)
}

// readAndVerifyCert will read the file from the given
// path, then check every 100ms until the file length stops
// changing. This is needed because the filesystem watcher
// that gets triggered by a WRITE doesn't have a guarantee
// that the write has finished yet.
func readAndVerifyCert(certFilePath string) ([]byte, error) {
	var err error
	firstRead, err := ioutil.ReadFile(certFilePath)
	if err != nil {
		return nil, err
	}
	time.Sleep(time.Millisecond * 100)
	secondRead, err := ioutil.ReadFile(certFilePath)
	if err != nil {
		return nil, err
	}

	if len(firstRead) != len(secondRead) {
		// The file is still being written to,
		// try again.
		return readAndVerifyCert(certFilePath)
	}
	return secondRead, nil
}

func serverCertSecret(privateKey, certChain []byte) cache_types.Resource {
	return &auth.Secret{
		Name: serverCert,
		Type: &auth.Secret_TlsCertificate{
			TlsCertificate: &auth.TlsCertificate{
				CertificateChain: &core.DataSource{
					Specifier: &core.DataSource_InlineBytes{
						InlineBytes: certChain,
					},
				},
				PrivateKey: &core.DataSource{
					Specifier: &core.DataSource_InlineBytes{
						InlineBytes: privateKey,
					},
				},
			},
		},
	}
}

func validationContextSecret(caCert []byte) cache_types.Resource {
	return &auth.Secret{
		Name: validationContext,
		Type: &auth.Secret_ValidationContext{
			ValidationContext: &auth.CertificateValidationContext{
				TrustedCa: &core.DataSource{
					Specifier: &core.DataSource_InlineBytes{
						InlineBytes: caCert,
					},
				},
			},
		},
	}
}
