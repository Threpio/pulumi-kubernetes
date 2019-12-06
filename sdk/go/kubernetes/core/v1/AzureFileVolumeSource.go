// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// AzureFile represents an Azure File Service mount on the host and bind mount to the pod.
type AzureFileVolumeSource struct {
	// Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	ReadOnly *bool `pulumi:"readOnly"`

	// the name of secret that contains Azure Storage Account Name and Key
	SecretName string `pulumi:"secretName"`

	// Share Name
	ShareName string `pulumi:"shareName"`

}

var _AzureFileVolumeSourceType = reflect.TypeOf((*AzureFileVolumeSource)(nil)).Elem()

// AzureFileVolumeSourceInput represents an input type that resolves to a AzureFileVolumeSource.
type AzureFileVolumeSourceInput interface {
	ElementType() reflect.Type

	ToAzureFileVolumeSourceOutput() AzureFileVolumeSourceOutput
	ToAzureFileVolumeSourceOutputWithContext(ctx context.Context) AzureFileVolumeSourceOutput
}

// AzureFileVolumeSourceArgs is a AzureFileVolumeSourceInput whose fields are all Input types.
type AzureFileVolumeSourceArgs struct {
	// the name of secret that contains Azure Storage Account Name and Key
	SecretName pulumi.StringInput `pulumi:"secretName"`

	// Share Name
	ShareName pulumi.StringInput `pulumi:"shareName"`

	// Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	ReadOnly pulumi.BoolInput `pulumi:"readOnly"`

}

func (a AzureFileVolumeSourceArgs) ElementType() reflect.Type {
	return _AzureFileVolumeSourceType
}

func (a AzureFileVolumeSourceArgs) ToAzureFileVolumeSourceOutput() AzureFileVolumeSourceOutput {
	return pulumi.ToOutput(a).(AzureFileVolumeSourceOutput)
}

func (a AzureFileVolumeSourceArgs) ToAzureFileVolumeSourceOutputWithContext(ctx context.Context) AzureFileVolumeSourceOutput {
	return pulumi.ToOutputWithContext(ctx, a).(AzureFileVolumeSourceOutput)
}

// AzureFileVolumeSourceOutput is an output type that resolves to a Input.
type AzureFileVolumeSourceOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(AzureFileVolumeSourceOutput{}) }

func (AzureFileVolumeSourceOutput) ElementType() reflect.Type {
	return _AzureFileVolumeSourceType
}

func (o AzureFileVolumeSourceOutput) ReadOnly() pulumi.BoolOutput {
	return o.Apply(func(v AzureFileVolumeSource) *bool {
		return v.ReadOnly
	}).(pulumi.BoolOutput)
}

func (o AzureFileVolumeSourceOutput) SecretName() pulumi.StringOutput {
	return o.Apply(func(v AzureFileVolumeSource) string {
		return v.SecretName
	}).(pulumi.StringOutput)
}

func (o AzureFileVolumeSourceOutput) ShareName() pulumi.StringOutput {
	return o.Apply(func(v AzureFileVolumeSource) string {
		return v.ShareName
	}).(pulumi.StringOutput)
}
