package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/solo-io/gloo/pkg/version"
	"github.com/solo-io/gloo/projects/sds/pkg/run"
	"github.com/solo-io/gloo/projects/sds/pkg/server"
	"github.com/solo-io/go-utils/contextutils"

	"go.uber.org/zap"
	v1 "k8s.io/api/core/v1"
)

var (
	glooMTLSEnabled      = false
	istioRotationEnabled = false

	// This must match the value of the sds_config target_uri in the envoy instance that it is providing
	// secrets to.
	sdsServerAddress   = "127.0.0.1:8234"
	startupFileRetries = 10 // Number of seconds to wait for certs to exist before crashing

	// The Node ID. Defaults to podname.namespace but can be overriden using env var SDS_CLIENT
	sdsClient = os.Getenv("GATEWAY_PROXY_POD_NAME") + "." + os.Getenv("GATEWAY_PROXY_POD_NAMESPACE")

	glooMtlsSecretDir     = "/etc/envoy/ssl/"
	glooServerCert        = "server_cert"
	glooValidationContext = "validation_context"

	istioCertDir           = "/etc/istio-certs/"
	istioServerCert        = "istio_server_cert"
	istioValidationContext = "istio_validation_context"
)

func init() {
	if enabled := os.Getenv("GLOO_MTLS_CERT_ROTATION_ENABLED"); enabled != "" && enabled != "false" {
		glooMTLSEnabled = true
	}
	if enabled := os.Getenv("ISTIO_CERT_ROTATION_ENABLED"); enabled != "" && enabled != "false" {
		istioRotationEnabled = true
	}
	if sdsClientOverride := os.Getenv("SDS_CLIENT"); sdsClientOverride != "" {
		sdsClient = sdsClientOverride
	}
	// At least one must be enabled, otherwise we have nothing to do.
	if !glooMTLSEnabled && !istioRotationEnabled {
		panic("at least one of Istio Cert rotation or Gloo Cert rotation must be enabled, using env vars ISTIO_CERT_ROTATION_ENABLED or GLOO_MTLS_CERT_ROTATION_ENABLED")
	}
}

func main() {
	secrets := []server.Secret{}
	if istioRotationEnabled {
		istioCertsSecret := server.Secret{
			ServerCert:        istioServerCert,
			ValidationContext: istioValidationContext,
			SslCaFile:         istioCertDir + "root-cert.pem",
			SslCertFile:       istioCertDir + "cert-chain.pem",
			SslKeyFile:        istioCertDir + "key.pem",
		}
		secrets = append(secrets, istioCertsSecret)
	}

	if glooMTLSEnabled {
		glooMtlsSecret := server.Secret{
			ServerCert:        glooServerCert,
			ValidationContext: glooValidationContext,
			SslCaFile:         glooMtlsSecretDir + v1.ServiceAccountRootCAKey,
			SslCertFile:       glooMtlsSecretDir + v1.TLSCertKey,
			SslKeyFile:        glooMtlsSecretDir + v1.TLSPrivateKeyKey,
		}
		secrets = append(secrets, glooMtlsSecret)
	}

	ctx := contextutils.WithLogger(context.Background(), "sds_server")
	ctx = contextutils.WithLoggerValues(ctx, "version", version.Version)

	for _, s := range secrets {
		// Check to see if files exist first to avoid crashloops
		if err := checkFilesExist([]string{s.SslKeyFile, s.SslCertFile, s.SslCaFile}); err != nil {
			contextutils.LoggerFrom(ctx).Fatal(err)
		}
	}
	contextutils.LoggerFrom(ctx).Infow(
		"starting up",
		zap.Bool("glooCertRotationEnabled", glooMTLSEnabled),
		zap.Bool("istioCertRotationEnabled", istioRotationEnabled),
	)

	if err := run.Run(ctx, secrets, sdsClient, sdsServerAddress); err != nil {
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
