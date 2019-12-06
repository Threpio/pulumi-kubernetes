// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// APIGroup contains the name, the supported versions, and the preferred version of a group.
type APIGroup struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should
	// convert recognized schemas to the latest internal value, and may reject unrecognized values.
	// More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer
	// this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More
	// info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`

	// name is the name of the group.
	Name string `pulumi:"name"`

	// preferredVersion is the version preferred by the API server, which probably is the storage
	// version.
	PreferredVersion *GroupVersionForDiscovery `pulumi:"preferredVersion"`

	// a map of client CIDR to server address that is serving this group. This is to help clients reach
	// servers in the most network-efficient way possible. Clients can use the appropriate server
	// address as per the CIDR that they match. In case of multiple matches, clients should use the
	// longest matching CIDR. The server returns only those CIDRs that it thinks that the client can
	// match. For example: the master will return an internal IP CIDR only, if the client reaches the
	// server using an internal IP. Server looks at X-Forwarded-For header or X-Real-Ip header or
	// request.RemoteAddr (in that order) to get the client IP.
	ServerAddressByClientCIDRs []ServerAddressByClientCIDR `pulumi:"serverAddressByClientCIDRs"`

	// versions are the versions supported in this group.
	Versions []GroupVersionForDiscovery `pulumi:"versions"`

}

var _APIGroupType = reflect.TypeOf((*APIGroup)(nil)).Elem()

// APIGroupInput represents an input type that resolves to a APIGroup.
type APIGroupInput interface {
	ElementType() reflect.Type

	ToAPIGroupOutput() APIGroupOutput
	ToAPIGroupOutputWithContext(ctx context.Context) APIGroupOutput
}

// APIGroupArgs is a APIGroupInput whose fields are all Input types.
type APIGroupArgs struct {
	// name is the name of the group.
	Name pulumi.StringInput `pulumi:"name"`

	// versions are the versions supported in this group.
	Versions GroupVersionForDiscoveryArrayInput `pulumi:"versions"`

	// APIVersion defines the versioned schema of this representation of an object. Servers should
	// convert recognized schemas to the latest internal value, and may reject unrecognized values.
	// More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringInput `pulumi:"apiVersion"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer
	// this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More
	// info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringInput `pulumi:"kind"`

	// preferredVersion is the version preferred by the API server, which probably is the storage
	// version.
	PreferredVersion GroupVersionForDiscoveryInput `pulumi:"preferredVersion"`

	// a map of client CIDR to server address that is serving this group. This is to help clients reach
	// servers in the most network-efficient way possible. Clients can use the appropriate server
	// address as per the CIDR that they match. In case of multiple matches, clients should use the
	// longest matching CIDR. The server returns only those CIDRs that it thinks that the client can
	// match. For example: the master will return an internal IP CIDR only, if the client reaches the
	// server using an internal IP. Server looks at X-Forwarded-For header or X-Real-Ip header or
	// request.RemoteAddr (in that order) to get the client IP.
	ServerAddressByClientCIDRs ServerAddressByClientCIDRArrayInput `pulumi:"serverAddressByClientCIDRs"`

}

func (a APIGroupArgs) ElementType() reflect.Type {
	return _APIGroupType
}

func (a APIGroupArgs) ToAPIGroupOutput() APIGroupOutput {
	return pulumi.ToOutput(a).(APIGroupOutput)
}

func (a APIGroupArgs) ToAPIGroupOutputWithContext(ctx context.Context) APIGroupOutput {
	return pulumi.ToOutputWithContext(ctx, a).(APIGroupOutput)
}

// APIGroupOutput is an output type that resolves to a Input.
type APIGroupOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(APIGroupOutput{}) }

func (APIGroupOutput) ElementType() reflect.Type {
	return _APIGroupType
}

func (o APIGroupOutput) ApiVersion() pulumi.StringOutput {
	return o.Apply(func(v APIGroup) *string {
		return v.ApiVersion
	}).(pulumi.StringOutput)
}

func (o APIGroupOutput) Kind() pulumi.StringOutput {
	return o.Apply(func(v APIGroup) *string {
		return v.Kind
	}).(pulumi.StringOutput)
}

func (o APIGroupOutput) Name() pulumi.StringOutput {
	return o.Apply(func(v APIGroup) string {
		return v.Name
	}).(pulumi.StringOutput)
}

func (o APIGroupOutput) PreferredVersion() GroupVersionForDiscoveryOutput {
	return o.Apply(func(v APIGroup) *GroupVersionForDiscovery {
		return v.PreferredVersion
	}).(GroupVersionForDiscoveryOutput)
}

func (o APIGroupOutput) ServerAddressByClientCIDRs() ServerAddressByClientCIDRArrayOutput {
	return o.Apply(func(v APIGroup) []ServerAddressByClientCIDR {
		return v.ServerAddressByClientCIDRs
	}).(ServerAddressByClientCIDRArrayOutput)
}

func (o APIGroupOutput) Versions() GroupVersionForDiscoveryArrayOutput {
	return o.Apply(func(v APIGroup) []GroupVersionForDiscovery {
		return v.Versions
	}).(GroupVersionForDiscoveryArrayOutput)
}

var _APIGroupArrayType = reflect.TypeOf((*[]APIGroup)(nil)).Elem()

type APIGroupArrayInput interface {
	ElementType() reflect.Type

	ToAPIGroupArrayOutput() APIGroupArrayOutput
	ToAPIGroupArrayOutputWithContext(ctx context.Context) APIGroupArrayOutput
}

type APIGroupArray []APIGroupInput

func (a APIGroupArray) ElementType() reflect.Type {
	return _APIGroupArrayType
}

func (a APIGroupArray) ToAPIGroupArrayOutput() APIGroupArrayOutput {
	return pulumi.ToOutput(a).(APIGroupArrayOutput)
}

func (a APIGroupArray) ToAPIGroupArrayOutputWithContext(ctx context.Context) APIGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(APIGroupArrayOutput)
}

type APIGroupArrayOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(APIGroupArrayOutput{}) }

func (APIGroupArrayOutput) ElementType() reflect.Type {
	return _APIGroupArrayType
}

func (o APIGroupArrayOutput) Index(i pulumi.IntInput) APIGroupOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) APIGroup {
		return vs[0].([]APIGroup)[vs[1].(int)]
	}).(APIGroupOutput)
}