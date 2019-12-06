// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// A topology selector requirement is a selector that matches given label. This is an alpha feature
// and may change in the future.
type TopologySelectorLabelRequirement struct {
	// The label key that the selector applies to.
	Key string `pulumi:"key"`

	// An array of string values. One value must match the label to be selected. Each entry in Values
	// is ORed.
	Values []string `pulumi:"values"`

}

var _TopologySelectorLabelRequirementType = reflect.TypeOf((*TopologySelectorLabelRequirement)(nil)).Elem()

// TopologySelectorLabelRequirementInput represents an input type that resolves to a TopologySelectorLabelRequirement.
type TopologySelectorLabelRequirementInput interface {
	ElementType() reflect.Type

	ToTopologySelectorLabelRequirementOutput() TopologySelectorLabelRequirementOutput
	ToTopologySelectorLabelRequirementOutputWithContext(ctx context.Context) TopologySelectorLabelRequirementOutput
}

// TopologySelectorLabelRequirementArgs is a TopologySelectorLabelRequirementInput whose fields are all Input types.
type TopologySelectorLabelRequirementArgs struct {
	// The label key that the selector applies to.
	Key pulumi.StringInput `pulumi:"key"`

	// An array of string values. One value must match the label to be selected. Each entry in Values
	// is ORed.
	Values pulumi.StringArrayInput `pulumi:"values"`

}

func (a TopologySelectorLabelRequirementArgs) ElementType() reflect.Type {
	return _TopologySelectorLabelRequirementType
}

func (a TopologySelectorLabelRequirementArgs) ToTopologySelectorLabelRequirementOutput() TopologySelectorLabelRequirementOutput {
	return pulumi.ToOutput(a).(TopologySelectorLabelRequirementOutput)
}

func (a TopologySelectorLabelRequirementArgs) ToTopologySelectorLabelRequirementOutputWithContext(ctx context.Context) TopologySelectorLabelRequirementOutput {
	return pulumi.ToOutputWithContext(ctx, a).(TopologySelectorLabelRequirementOutput)
}

// TopologySelectorLabelRequirementOutput is an output type that resolves to a Input.
type TopologySelectorLabelRequirementOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(TopologySelectorLabelRequirementOutput{}) }

func (TopologySelectorLabelRequirementOutput) ElementType() reflect.Type {
	return _TopologySelectorLabelRequirementType
}

func (o TopologySelectorLabelRequirementOutput) Key() pulumi.StringOutput {
	return o.Apply(func(v TopologySelectorLabelRequirement) string {
		return v.Key
	}).(pulumi.StringOutput)
}

func (o TopologySelectorLabelRequirementOutput) Values() pulumi.StringArrayOutput {
	return o.Apply(func(v TopologySelectorLabelRequirement) []string {
		return v.Values
	}).(pulumi.StringArrayOutput)
}

var _TopologySelectorLabelRequirementArrayType = reflect.TypeOf((*[]TopologySelectorLabelRequirement)(nil)).Elem()

type TopologySelectorLabelRequirementArrayInput interface {
	ElementType() reflect.Type

	ToTopologySelectorLabelRequirementArrayOutput() TopologySelectorLabelRequirementArrayOutput
	ToTopologySelectorLabelRequirementArrayOutputWithContext(ctx context.Context) TopologySelectorLabelRequirementArrayOutput
}

type TopologySelectorLabelRequirementArray []TopologySelectorLabelRequirementInput

func (a TopologySelectorLabelRequirementArray) ElementType() reflect.Type {
	return _TopologySelectorLabelRequirementArrayType
}

func (a TopologySelectorLabelRequirementArray) ToTopologySelectorLabelRequirementArrayOutput() TopologySelectorLabelRequirementArrayOutput {
	return pulumi.ToOutput(a).(TopologySelectorLabelRequirementArrayOutput)
}

func (a TopologySelectorLabelRequirementArray) ToTopologySelectorLabelRequirementArrayOutputWithContext(ctx context.Context) TopologySelectorLabelRequirementArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(TopologySelectorLabelRequirementArrayOutput)
}

type TopologySelectorLabelRequirementArrayOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(TopologySelectorLabelRequirementArrayOutput{}) }

func (TopologySelectorLabelRequirementArrayOutput) ElementType() reflect.Type {
	return _TopologySelectorLabelRequirementArrayType
}

func (o TopologySelectorLabelRequirementArrayOutput) Index(i pulumi.IntInput) TopologySelectorLabelRequirementOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) TopologySelectorLabelRequirement {
		return vs[0].([]TopologySelectorLabelRequirement)[vs[1].(int)]
	}).(TopologySelectorLabelRequirementOutput)
}