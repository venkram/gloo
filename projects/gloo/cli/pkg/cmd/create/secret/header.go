package secret

import (
	"context"

	"github.com/solo-io/gloo/projects/gloo/cli/pkg/printers"

	errors "github.com/rotisserie/eris"
	"github.com/solo-io/gloo/pkg/cliutil"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/argsutils"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/helpers"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/spf13/cobra"
)

func headerCmd(opts *options.Options) *cobra.Command {
	input := &opts.Create.InputSecret.HeaderSecret
	cmd := &cobra.Command{
		Use:   "header",
		Short: `Create a header secret with the given name`,
		Long: "Create a header secret with the given name. The format of the secret data is: " +
			"`{\"header-name\" : [header-name string] , \"value\" : [header-value string]}`",
		RunE: func(c *cobra.Command, args []string) error {
			if err := argsutils.MetadataArgsParse(opts, args); err != nil {
				return err
			}
			if opts.Top.Interactive {
				// gather any missing args that are available through interactive mode
				if err := HeaderSecretArgsInteractive(input); err != nil {
					return err
				}
			}
			// create the secret
			if err := createHeaderSecret(opts.Top.Ctx, opts.Metadata, *input, opts.Create.DryRun, opts.Top.Output); err != nil {
				return err
			}
			return nil
		},
	}

	flags := cmd.Flags()
	flags.StringVar(&input.HeaderName, "header-name", "", "header name")
	flags.StringVar(&input.Value, "value", "", "header value")

	return cmd
}

const (
	headerPromptName  = "Enter header name: "
	headerPromptValue = "Enter header value: "
)

func HeaderSecretArgsInteractive(input *options.HeaderSecret) error {
	if err := cliutil.GetStringInput(headerPromptName, &input.HeaderName); err != nil {
		return err
	}
	if err := cliutil.GetStringInput(headerPromptValue, &input.Value); err != nil {
		return err
	}

	return nil
}

func createHeaderSecret(ctx context.Context, meta core.Metadata, input options.HeaderSecret, dryRun bool, outputType printers.OutputType) error {
	if input.HeaderName == "" {
		return errors.Errorf("must provide header name")
	}
	secret := &gloov1.Secret{
		Metadata: meta,
		Kind: &gloov1.Secret_Header{
			Header: &gloov1.HeaderSecret{
				HeaderName: input.HeaderName,
				Value:      input.Value,
			},
		},
	}

	if !dryRun {
		var err error
		secretClient := helpers.MustSecretClientWithOptions(0, []string{meta.Namespace})
		if secret, err = secretClient.Write(secret, clients.WriteOpts{Ctx: ctx}); err != nil {
			return err
		}
	}

	_ = printers.PrintSecrets(gloov1.SecretList{secret}, outputType)
	return nil
}
