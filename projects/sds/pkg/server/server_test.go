package server_test

import (
	"context"
	"io/ioutil"
	"time"

	"github.com/solo-io/gloo/projects/sds/pkg/server"

	envoy_api_v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	envoy_service_discovery_v2 "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v2"
	"github.com/spf13/afero"
	"google.golang.org/grpc"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SDS Server", func() {

	var (
		fs                        afero.Fs
		dir                       string
		keyFile, certFile, caFile afero.File
		err                       error
		serverAddr                = "127.0.0.1:8888"
		sdsClient                 = "test-client"
		srv                       *server.Server
	)

	BeforeEach(func() {
		fs = afero.NewOsFs()
		dir, err = afero.TempDir(fs, "", "")
		Expect(err).To(BeNil())
		fileString := `test`
		keyFile, err = afero.TempFile(fs, dir, "")
		Expect(err).To(BeNil())
		_, err = keyFile.WriteString(fileString)
		Expect(err).To(BeNil())
		certFile, err = afero.TempFile(fs, dir, "")
		Expect(err).To(BeNil())
		_, err = certFile.WriteString(fileString)
		Expect(err).To(BeNil())
		caFile, err = afero.TempFile(fs, dir, "")
		Expect(err).To(BeNil())
		_, err = caFile.WriteString(fileString)
		Expect(err).To(BeNil())
		secrets := []server.Secret{
			{
				ServerCert:        "test-server",
				SslCaFile:         caFile.Name(),
				SslCertFile:       certFile.Name(),
				SslKeyFile:        keyFile.Name(),
				ValidationContext: "test-validation",
			},
		}
		srv = server.SetupEnvoySDS(secrets, sdsClient, serverAddr)
	})

	AfterEach(func() {
		_ = fs.RemoveAll(dir)
	})

	It("correctly reads tls secrets from files to generate snapshot version", func() {
		certs := filesToBytes(keyFile.Name(), certFile.Name(), caFile.Name())
		snapshotVersion, err := server.GetSnapshotVersion(certs)
		Expect(err).To(BeNil())
		Expect(snapshotVersion).To(Equal("6730780456972595554"))

		// Test that the snapshot version changes if the contents of the file changes
		_, err = keyFile.WriteString(`newFileString`)
		Expect(err).To(BeNil())
		certs = filesToBytes(keyFile.Name(), certFile.Name(), caFile.Name())
		snapshotVersion, err = server.GetSnapshotVersion(certs)
		Expect(err).To(BeNil())
		Expect(snapshotVersion).To(Equal("4234248347190811569"))
	})

	It("correctly updates SDSConfig", func() {
		ctx, _ := context.WithCancel(context.Background())
		snapshotCache := server.GetSnapshotCache(*srv)
		err := srv.UpdateSDSConfig(ctx)
		Expect(err).NotTo(HaveOccurred())
		_, err = snapshotCache.GetSnapshot(srv.ID(nil))
		Expect(err).To(BeNil())
	})

	Context("Test gRPC Server", func() {
		var (
			ctx    context.Context
			cancel context.CancelFunc
		)

		BeforeEach(func() {
			ctx, cancel = context.WithCancel(context.Background())
			_, err = srv.Run(ctx)
			// Give it a second to come up + read the certs
			time.Sleep(time.Second * 1)
			Expect(err).To(BeNil())
		})

		AfterEach(func() {
			cancel()
		})

		It("accepts client connections", func() {
			// Check that it's answering
			var conn *grpc.ClientConn

			// Initiate a connection with the server
			conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
			Expect(err).To(BeNil())
			defer conn.Close()

			client := envoy_service_discovery_v2.NewSecretDiscoveryServiceClient(conn)

			// Before any snapshot is set, expect an error when fetching secrets
			resp, err := client.FetchSecrets(ctx, &envoy_api_v2.DiscoveryRequest{})
			Expect(err).NotTo(BeNil())

			// After snapshot is set, expect to see the secrets
			srv.UpdateSDSConfig(ctx)
			resp, err = client.FetchSecrets(ctx, &envoy_api_v2.DiscoveryRequest{})
			Expect(err).To(BeNil())
			Expect(len(resp.GetResources())).To(Equal(2))
			Expect(resp.Validate()).To(BeNil())
		})
	})
})

func filesToBytes(keyName, certName, caName string) [][]byte {
	certs := [][]byte{}
	keyBytes, err := ioutil.ReadFile(keyName)
	Expect(err).NotTo(HaveOccurred())
	certs = append(certs, keyBytes)
	certBytes, err := ioutil.ReadFile(certName)
	Expect(err).NotTo(HaveOccurred())
	certs = append(certs, certBytes)
	caBytes, err := ioutil.ReadFile(caName)
	Expect(err).NotTo(HaveOccurred())
	certs = append(certs, caBytes)
	return certs
}
