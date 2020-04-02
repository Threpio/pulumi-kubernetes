// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package providers

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The provider type for the kubernetes package.
type Provider struct {
	pulumi.ProviderResourceState
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		args = &ProviderArgs{}
	}
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:kubernetes", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// If present, the name of the kubeconfig cluster to use.
	Cluster *string `pulumi:"cluster"`
	// If present, the name of the kubeconfig context to use.
	Context *string `pulumi:"context"`
	// BETA FEATURE - If present and set to true, enable server-side diff calculations.
	// This feature is in developer preview, and is disabled by default.
	//
	// This config can be specified in the following ways, using this precedence:
	// 1. This `enableDryRun` parameter.
	// 2. The `PULUMI_K8S_ENABLE_DRY_RUN` environment variable.
	EnableDryRun *bool `pulumi:"enableDryRun"`
	// The contents of a kubeconfig file. If this is set, this config will be used instead of $KUBECONFIG.
	Kubeconfig *string `pulumi:"kubeconfig"`
	// If present, the default namespace to use. This flag is ignored for cluster-scoped resources.
	//
	// A namespace can be specified in multiple places, and the precedence is as follows:
	// 1. `.metadata.namespace` set on the resource.
	// 2. This `namespace` parameter.
	// 3. `namespace` set for the active context in the kubeconfig.
	Namespace *string `pulumi:"namespace"`
	// If present and set to true, suppress apiVersion deprecation warnings from the CLI.
	//
	// This config can be specified in the following ways, using this precedence:
	// 1. This `suppressDeprecationWarnings` parameter.
	// 2. The `PULUMI_K8S_SUPPRESS_DEPRECATION_WARNINGS` environment variable.
	SuppressDeprecationWarnings *bool `pulumi:"suppressDeprecationWarnings"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// If present, the name of the kubeconfig cluster to use.
	Cluster pulumi.StringPtrInput
	// If present, the name of the kubeconfig context to use.
	Context pulumi.StringPtrInput
	// BETA FEATURE - If present and set to true, enable server-side diff calculations.
	// This feature is in developer preview, and is disabled by default.
	//
	// This config can be specified in the following ways, using this precedence:
	// 1. This `enableDryRun` parameter.
	// 2. The `PULUMI_K8S_ENABLE_DRY_RUN` environment variable.
	EnableDryRun pulumi.BoolPtrInput
	// The contents of a kubeconfig file. If this is set, this config will be used instead of $KUBECONFIG.
	Kubeconfig pulumi.StringPtrInput
	// If present, the default namespace to use. This flag is ignored for cluster-scoped resources.
	//
	// A namespace can be specified in multiple places, and the precedence is as follows:
	// 1. `.metadata.namespace` set on the resource.
	// 2. This `namespace` parameter.
	// 3. `namespace` set for the active context in the kubeconfig.
	Namespace pulumi.StringPtrInput
	// If present and set to true, suppress apiVersion deprecation warnings from the CLI.
	//
	// This config can be specified in the following ways, using this precedence:
	// 1. This `suppressDeprecationWarnings` parameter.
	// 2. The `PULUMI_K8S_SUPPRESS_DEPRECATION_WARNINGS` environment variable.
	SuppressDeprecationWarnings pulumi.BoolPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}
