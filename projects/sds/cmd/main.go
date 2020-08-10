package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/solo-io/gloo/pkg/version"
	"github.com/solo-io/gloo/projects/sds/pkg/run"
	"github.com/solo-io/go-utils/contextutils"
	// v1 "k8s.io/api/core/v1"
)

var (
	// secretDir   = "/etc/envoy/ssl/"
	// sslKeyFile  = secretDir + v1.TLSPrivateKeyKey        //tls.key
	// sslCertFile = secretDir + v1.TLSCertKey              //tls.crt
	// sslCaFile   = secretDir + v1.ServiceAccountRootCAKey //ca.crt

	// TODO: Remove Istio test code, make configurable.
	secretDir          = "/etc/istio-certs/"
	sslCaFile          = secretDir + "root-cert.pem"
	sslKeyFile         = secretDir + "key.pem"
	sslCertFile        = secretDir + "cert-chain.pem"
	startupFileRetries = 10 // Number of seconds to wait for certs to exist before crashing

	// This must match the value of the sds_config target_uri in the envoy instance that it is providing
	// secrets to.
	sdsServerAddress = "127.0.0.1:8234"
)

func main() {
	if envKeyDir := os.Getenv("KEY_DIR"); envKeyDir != "" {
		secretDir = envKeyDir
	}
	if envSDSAddr := os.Getenv("SDS_SERVER_ADDRESS"); envSDSAddr != "" {
		sdsServerAddress = envSDSAddr
	}

	ctx := contextutils.WithLogger(context.Background(), "sds_server")
	ctx = contextutils.WithLoggerValues(ctx, "version", version.Version)

	// Check to see if files exist first to avoid crashloops
	if err := checkFilesExist([]string{sslKeyFile, sslCertFile, sslCaFile}); err != nil {
		contextutils.LoggerFrom(ctx).Fatal(err)
	}

	if err := run.Run(ctx, sslKeyFile, sslCertFile, sslCaFile, sdsServerAddress); err != nil {
		contextutils.LoggerFrom(ctx).Fatal(err)
	}
}

// checkFilesExist returns an err if any of the
// given filePaths do not exist.
func checkFilesExist(filePaths []string) error {
	for _, filePath := range filePaths {
		if !fileExists(filePath, startupFileRetries) {
			return fmt.Errorf("could not find file '%v'", filePath)
		}
	}
	return nil
}

// fileExists checks to see if a file exists every
// second for `retries` seconds. Returns false after
// all retries are used if it still can't find it
func fileExists(filePath string, retries int) bool {
	if _, err := os.Stat(filePath); err == nil {
		return true
	}
	if retries > 0 {
		time.Sleep(time.Second * 1)
		return fileExists(filePath, retries-1)
	}
	return false
}
