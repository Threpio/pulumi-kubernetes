// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// StatusCause provides more information about an api.Status failure, including cases when multiple
// errors are encountered.
type StatusCause struct {
	// The field of the resource that has caused this error, as named by its JSON serialization. May
	// include dot and postfix notation for nested attributes. Arrays are zero-indexed.  Fields may
	// appear more than once in an array of causes due to fields having multiple errors. Optional.
	// 
	// Examples:
	//   "name" - the field "name" on the current resource
	//   "items[0].name" - the field "name" on the first array entry in "items"
	Field *string `pulumi:"field"`

	// A human-readable description of the cause of the error.  This field may be presented as-is to a
	// reader.
	Message *string `pulumi:"message"`

	// A machine-readable description of the cause of the error. If this value is empty there is no
	// information available.
	Reason *string `pulumi:"reason"`

}

var _StatusCauseType = reflect.TypeOf((*StatusCause)(nil)).Elem()

// StatusCauseInput represents an input type that resolves to a StatusCause.
type StatusCauseInput interface {
	ElementType() reflect.Type

	ToStatusCauseOutput() StatusCauseOutput
	ToStatusCauseOutputWithContext(ctx context.Context) StatusCauseOutput
}

// StatusCauseArgs is a StatusCauseInput whose fields are all Input types.
type StatusCauseArgs struct {
	// The field of the resource that has caused this error, as named by its JSON serialization. May
	// include dot and postfix notation for nested attributes. Arrays are zero-indexed.  Fields may
	// appear more than once in an array of causes due to fields having multiple errors. Optional.
	// 
	// Examples:
	//   "name" - the field "name" on the current resource
	//   "items[0].name" - the field "name" on the first array entry in "items"
	Field pulumi.StringInput `pulumi:"field"`

	// A human-readable description of the cause of the error.  This field may be presented as-is to a
	// reader.
	Message pulumi.StringInput `pulumi:"message"`

	// A machine-readable description of the cause of the error. If this value is empty there is no
	// information available.
	Reason pulumi.StringInput `pulumi:"reason"`

}

func (a StatusCauseArgs) ElementType() reflect.Type {
	return _StatusCauseType
}

func (a StatusCauseArgs) ToStatusCauseOutput() StatusCauseOutput {
	return pulumi.ToOutput(a).(StatusCauseOutput)
}

func (a StatusCauseArgs) ToStatusCauseOutputWithContext(ctx context.Context) StatusCauseOutput {
	return pulumi.ToOutputWithContext(ctx, a).(StatusCauseOutput)
}

// StatusCauseOutput is an output type that resolves to a Input.
type StatusCauseOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(StatusCauseOutput{}) }

func (StatusCauseOutput) ElementType() reflect.Type {
	return _StatusCauseType
}

func (o StatusCauseOutput) Field() pulumi.StringOutput {
	return o.Apply(func(v StatusCause) *string {
		return v.Field
	}).(pulumi.StringOutput)
}

func (o StatusCauseOutput) Message() pulumi.StringOutput {
	return o.Apply(func(v StatusCause) *string {
		return v.Message
	}).(pulumi.StringOutput)
}

func (o StatusCauseOutput) Reason() pulumi.StringOutput {
	return o.Apply(func(v StatusCause) *string {
		return v.Reason
	}).(pulumi.StringOutput)
}

var _StatusCauseArrayType = reflect.TypeOf((*[]StatusCause)(nil)).Elem()

type StatusCauseArrayInput interface {
	ElementType() reflect.Type

	ToStatusCauseArrayOutput() StatusCauseArrayOutput
	ToStatusCauseArrayOutputWithContext(ctx context.Context) StatusCauseArrayOutput
}

type StatusCauseArray []StatusCauseInput

func (a StatusCauseArray) ElementType() reflect.Type {
	return _StatusCauseArrayType
}

func (a StatusCauseArray) ToStatusCauseArrayOutput() StatusCauseArrayOutput {
	return pulumi.ToOutput(a).(StatusCauseArrayOutput)
}

func (a StatusCauseArray) ToStatusCauseArrayOutputWithContext(ctx context.Context) StatusCauseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(StatusCauseArrayOutput)
}

type StatusCauseArrayOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(StatusCauseArrayOutput{}) }

func (StatusCauseArrayOutput) ElementType() reflect.Type {
	return _StatusCauseArrayType
}

func (o StatusCauseArrayOutput) Index(i pulumi.IntInput) StatusCauseOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) StatusCause {
		return vs[0].([]StatusCause)[vs[1].(int)]
	}).(StatusCauseOutput)
}