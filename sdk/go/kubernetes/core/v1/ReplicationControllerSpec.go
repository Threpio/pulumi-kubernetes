// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// ReplicationControllerSpec is the specification of a replication controller.
type ReplicationControllerSpec struct {
	// Minimum number of seconds for which a newly created pod should be ready without any of its
	// container crashing, for it to be considered available. Defaults to 0 (pod will be considered
	// available as soon as it is ready)
	MinReadySeconds *int `pulumi:"minReadySeconds"`

	// Replicas is the number of desired replicas. This is a pointer to distinguish between explicit
	// zero and unspecified. Defaults to 1. More info:
	// https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller#what-is-a-replicationcontroller
	Replicas *int `pulumi:"replicas"`

	// Selector is a label query over pods that should match the Replicas count. If Selector is empty,
	// it is defaulted to the labels present on the Pod template. Label keys and values that must match
	// in order to be controlled by this replication controller, if empty defaulted to labels on Pod
	// template. More info:
	// https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
	Selector map[string]string `pulumi:"selector"`

	// Template is the object that describes the pod that will be created if insufficient replicas are
	// detected. This takes precedence over a TemplateRef. More info:
	// https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller#pod-template
	Template *PodTemplateSpec `pulumi:"template"`

}

var _ReplicationControllerSpecType = reflect.TypeOf((*ReplicationControllerSpec)(nil)).Elem()

// ReplicationControllerSpecInput represents an input type that resolves to a ReplicationControllerSpec.
type ReplicationControllerSpecInput interface {
	ElementType() reflect.Type

	ToReplicationControllerSpecOutput() ReplicationControllerSpecOutput
	ToReplicationControllerSpecOutputWithContext(ctx context.Context) ReplicationControllerSpecOutput
}

// ReplicationControllerSpecArgs is a ReplicationControllerSpecInput whose fields are all Input types.
type ReplicationControllerSpecArgs struct {
	// Minimum number of seconds for which a newly created pod should be ready without any of its
	// container crashing, for it to be considered available. Defaults to 0 (pod will be considered
	// available as soon as it is ready)
	MinReadySeconds pulumi.IntInput `pulumi:"minReadySeconds"`

	// Replicas is the number of desired replicas. This is a pointer to distinguish between explicit
	// zero and unspecified. Defaults to 1. More info:
	// https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller#what-is-a-replicationcontroller
	Replicas pulumi.IntInput `pulumi:"replicas"`

	// Selector is a label query over pods that should match the Replicas count. If Selector is empty,
	// it is defaulted to the labels present on the Pod template. Label keys and values that must match
	// in order to be controlled by this replication controller, if empty defaulted to labels on Pod
	// template. More info:
	// https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
	Selector pulumi.StringMapInput `pulumi:"selector"`

	// Template is the object that describes the pod that will be created if insufficient replicas are
	// detected. This takes precedence over a TemplateRef. More info:
	// https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller#pod-template
	Template PodTemplateSpecInput `pulumi:"template"`

}

func (a ReplicationControllerSpecArgs) ElementType() reflect.Type {
	return _ReplicationControllerSpecType
}

func (a ReplicationControllerSpecArgs) ToReplicationControllerSpecOutput() ReplicationControllerSpecOutput {
	return pulumi.ToOutput(a).(ReplicationControllerSpecOutput)
}

func (a ReplicationControllerSpecArgs) ToReplicationControllerSpecOutputWithContext(ctx context.Context) ReplicationControllerSpecOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ReplicationControllerSpecOutput)
}

// ReplicationControllerSpecOutput is an output type that resolves to a Input.
type ReplicationControllerSpecOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(ReplicationControllerSpecOutput{}) }

func (ReplicationControllerSpecOutput) ElementType() reflect.Type {
	return _ReplicationControllerSpecType
}

func (o ReplicationControllerSpecOutput) MinReadySeconds() pulumi.IntOutput {
	return o.Apply(func(v ReplicationControllerSpec) *int {
		return v.MinReadySeconds
	}).(pulumi.IntOutput)
}

func (o ReplicationControllerSpecOutput) Replicas() pulumi.IntOutput {
	return o.Apply(func(v ReplicationControllerSpec) *int {
		return v.Replicas
	}).(pulumi.IntOutput)
}

func (o ReplicationControllerSpecOutput) Selector() pulumi.StringMapOutput {
	return o.Apply(func(v ReplicationControllerSpec) map[string]string {
		return v.Selector
	}).(pulumi.StringMapOutput)
}

func (o ReplicationControllerSpecOutput) Template() PodTemplateSpecOutput {
	return o.Apply(func(v ReplicationControllerSpec) *PodTemplateSpec {
		return v.Template
	}).(PodTemplateSpecOutput)
}
