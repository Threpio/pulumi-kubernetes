// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// ContainerStateTerminated is a terminated state of a container.
type ContainerStateTerminated struct {
	// Container's ID in the format 'docker://<container_id>'
	ContainerID *string `pulumi:"containerID"`

	// Exit status from the last termination of the container
	ExitCode int `pulumi:"exitCode"`

	// Time at which the container last terminated
	FinishedAt *string `pulumi:"finishedAt"`

	// Message regarding the last termination of the container
	Message *string `pulumi:"message"`

	// (brief) reason from the last termination of the container
	Reason *string `pulumi:"reason"`

	// Signal from the last termination of the container
	Signal *int `pulumi:"signal"`

	// Time at which previous execution of the container started
	StartedAt *string `pulumi:"startedAt"`

}

var _ContainerStateTerminatedType = reflect.TypeOf((*ContainerStateTerminated)(nil)).Elem()

// ContainerStateTerminatedInput represents an input type that resolves to a ContainerStateTerminated.
type ContainerStateTerminatedInput interface {
	ElementType() reflect.Type

	ToContainerStateTerminatedOutput() ContainerStateTerminatedOutput
	ToContainerStateTerminatedOutputWithContext(ctx context.Context) ContainerStateTerminatedOutput
}

// ContainerStateTerminatedArgs is a ContainerStateTerminatedInput whose fields are all Input types.
type ContainerStateTerminatedArgs struct {
	// Exit status from the last termination of the container
	ExitCode pulumi.IntInput `pulumi:"exitCode"`

	// Container's ID in the format 'docker://<container_id>'
	ContainerID pulumi.StringInput `pulumi:"containerID"`

	// Time at which the container last terminated
	FinishedAt pulumi.StringInput `pulumi:"finishedAt"`

	// Message regarding the last termination of the container
	Message pulumi.StringInput `pulumi:"message"`

	// (brief) reason from the last termination of the container
	Reason pulumi.StringInput `pulumi:"reason"`

	// Signal from the last termination of the container
	Signal pulumi.IntInput `pulumi:"signal"`

	// Time at which previous execution of the container started
	StartedAt pulumi.StringInput `pulumi:"startedAt"`

}

func (a ContainerStateTerminatedArgs) ElementType() reflect.Type {
	return _ContainerStateTerminatedType
}

func (a ContainerStateTerminatedArgs) ToContainerStateTerminatedOutput() ContainerStateTerminatedOutput {
	return pulumi.ToOutput(a).(ContainerStateTerminatedOutput)
}

func (a ContainerStateTerminatedArgs) ToContainerStateTerminatedOutputWithContext(ctx context.Context) ContainerStateTerminatedOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ContainerStateTerminatedOutput)
}

// ContainerStateTerminatedOutput is an output type that resolves to a Input.
type ContainerStateTerminatedOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(ContainerStateTerminatedOutput{}) }

func (ContainerStateTerminatedOutput) ElementType() reflect.Type {
	return _ContainerStateTerminatedType
}

func (o ContainerStateTerminatedOutput) ContainerID() pulumi.StringOutput {
	return o.Apply(func(v ContainerStateTerminated) *string {
		return v.ContainerID
	}).(pulumi.StringOutput)
}

func (o ContainerStateTerminatedOutput) ExitCode() pulumi.IntOutput {
	return o.Apply(func(v ContainerStateTerminated) int {
		return v.ExitCode
	}).(pulumi.IntOutput)
}

func (o ContainerStateTerminatedOutput) FinishedAt() pulumi.StringOutput {
	return o.Apply(func(v ContainerStateTerminated) *string {
		return v.FinishedAt
	}).(pulumi.StringOutput)
}

func (o ContainerStateTerminatedOutput) Message() pulumi.StringOutput {
	return o.Apply(func(v ContainerStateTerminated) *string {
		return v.Message
	}).(pulumi.StringOutput)
}

func (o ContainerStateTerminatedOutput) Reason() pulumi.StringOutput {
	return o.Apply(func(v ContainerStateTerminated) *string {
		return v.Reason
	}).(pulumi.StringOutput)
}

func (o ContainerStateTerminatedOutput) Signal() pulumi.IntOutput {
	return o.Apply(func(v ContainerStateTerminated) *int {
		return v.Signal
	}).(pulumi.IntOutput)
}

func (o ContainerStateTerminatedOutput) StartedAt() pulumi.StringOutput {
	return o.Apply(func(v ContainerStateTerminated) *string {
		return v.StartedAt
	}).(pulumi.StringOutput)
}
