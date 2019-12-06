// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// EndpointPort represents a Port used by an EndpointSlice
type EndpointPort struct {
	// The name of this port. All ports in an EndpointSlice must have a unique name. If the
	// EndpointSlice is dervied from a Kubernetes service, this corresponds to the
	// Service.ports[].name. Name must either be an empty string or pass IANA_SVC_NAME validation: *
	// must be no more than 15 characters long * may contain only [-a-z0-9] * must contain at least one
	// letter [a-z] * it must not start or end with a hyphen, nor contain adjacent hyphens Default is
	// empty string.
	Name *string `pulumi:"name"`

	// The port number of the endpoint. If this is not specified, ports are not restricted and must be
	// interpreted in the context of the specific consumer.
	Port *int `pulumi:"port"`

	// The IP protocol for this port. Must be UDP, TCP, or SCTP. Default is TCP.
	Protocol *string `pulumi:"protocol"`

}

var _EndpointPortType = reflect.TypeOf((*EndpointPort)(nil)).Elem()

// EndpointPortInput represents an input type that resolves to a EndpointPort.
type EndpointPortInput interface {
	ElementType() reflect.Type

	ToEndpointPortOutput() EndpointPortOutput
	ToEndpointPortOutputWithContext(ctx context.Context) EndpointPortOutput
}

// EndpointPortArgs is a EndpointPortInput whose fields are all Input types.
type EndpointPortArgs struct {
	// The name of this port. All ports in an EndpointSlice must have a unique name. If the
	// EndpointSlice is dervied from a Kubernetes service, this corresponds to the
	// Service.ports[].name. Name must either be an empty string or pass IANA_SVC_NAME validation: *
	// must be no more than 15 characters long * may contain only [-a-z0-9] * must contain at least one
	// letter [a-z] * it must not start or end with a hyphen, nor contain adjacent hyphens Default is
	// empty string.
	Name pulumi.StringInput `pulumi:"name"`

	// The port number of the endpoint. If this is not specified, ports are not restricted and must be
	// interpreted in the context of the specific consumer.
	Port pulumi.IntInput `pulumi:"port"`

	// The IP protocol for this port. Must be UDP, TCP, or SCTP. Default is TCP.
	Protocol pulumi.StringInput `pulumi:"protocol"`

}

func (a EndpointPortArgs) ElementType() reflect.Type {
	return _EndpointPortType
}

func (a EndpointPortArgs) ToEndpointPortOutput() EndpointPortOutput {
	return pulumi.ToOutput(a).(EndpointPortOutput)
}

func (a EndpointPortArgs) ToEndpointPortOutputWithContext(ctx context.Context) EndpointPortOutput {
	return pulumi.ToOutputWithContext(ctx, a).(EndpointPortOutput)
}

// EndpointPortOutput is an output type that resolves to a Input.
type EndpointPortOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(EndpointPortOutput{}) }

func (EndpointPortOutput) ElementType() reflect.Type {
	return _EndpointPortType
}

func (o EndpointPortOutput) Name() pulumi.StringOutput {
	return o.Apply(func(v EndpointPort) *string {
		return v.Name
	}).(pulumi.StringOutput)
}

func (o EndpointPortOutput) Port() pulumi.IntOutput {
	return o.Apply(func(v EndpointPort) *int {
		return v.Port
	}).(pulumi.IntOutput)
}

func (o EndpointPortOutput) Protocol() pulumi.StringOutput {
	return o.Apply(func(v EndpointPort) *string {
		return v.Protocol
	}).(pulumi.StringOutput)
}

var _EndpointPortArrayType = reflect.TypeOf((*[]EndpointPort)(nil)).Elem()

type EndpointPortArrayInput interface {
	ElementType() reflect.Type

	ToEndpointPortArrayOutput() EndpointPortArrayOutput
	ToEndpointPortArrayOutputWithContext(ctx context.Context) EndpointPortArrayOutput
}

type EndpointPortArray []EndpointPortInput

func (a EndpointPortArray) ElementType() reflect.Type {
	return _EndpointPortArrayType
}

func (a EndpointPortArray) ToEndpointPortArrayOutput() EndpointPortArrayOutput {
	return pulumi.ToOutput(a).(EndpointPortArrayOutput)
}

func (a EndpointPortArray) ToEndpointPortArrayOutputWithContext(ctx context.Context) EndpointPortArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(EndpointPortArrayOutput)
}

type EndpointPortArrayOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(EndpointPortArrayOutput{}) }

func (EndpointPortArrayOutput) ElementType() reflect.Type {
	return _EndpointPortArrayType
}

func (o EndpointPortArrayOutput) Index(i pulumi.IntInput) EndpointPortOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) EndpointPort {
		return vs[0].([]EndpointPort)[vs[1].(int)]
	}).(EndpointPortOutput)
}