// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// NodeStatus is information about the current status of a node.
type NodeStatus struct {
	// List of addresses reachable to the node. Queried from cloud provider, if available. More info:
	// https://kubernetes.io/docs/concepts/nodes/node/#addresses Note: This field is declared as
	// mergeable, but the merge key is not sufficiently unique, which can cause data corruption when it
	// is merged. Callers should instead use a full-replacement patch. See http://pr.k8s.io/79391 for
	// an example.
	Addresses []NodeAddress `pulumi:"addresses"`

	// Allocatable represents the resources of a node that are available for scheduling. Defaults to
	// Capacity.
	Allocatable map[string]string `pulumi:"allocatable"`

	// Capacity represents the total resources of a node. More info:
	// https://kubernetes.io/docs/concepts/storage/persistent-volumes#capacity
	Capacity map[string]string `pulumi:"capacity"`

	// Conditions is an array of current observed node conditions. More info:
	// https://kubernetes.io/docs/concepts/nodes/node/#condition
	Conditions []NodeCondition `pulumi:"conditions"`

	// Status of the config assigned to the node via the dynamic Kubelet config feature.
	Config *NodeConfigStatus `pulumi:"config"`

	// Endpoints of daemons running on the Node.
	DaemonEndpoints *NodeDaemonEndpoints `pulumi:"daemonEndpoints"`

	// List of container images on this node
	Images []ContainerImage `pulumi:"images"`

	// Set of ids/uuids to uniquely identify the node. More info:
	// https://kubernetes.io/docs/concepts/nodes/node/#info
	NodeInfo *NodeSystemInfo `pulumi:"nodeInfo"`

	// NodePhase is the recently observed lifecycle phase of the node. More info:
	// https://kubernetes.io/docs/concepts/nodes/node/#phase The field is never populated, and now is
	// deprecated.
	Phase *string `pulumi:"phase"`

	// List of volumes that are attached to the node.
	VolumesAttached []AttachedVolume `pulumi:"volumesAttached"`

	// List of attachable volumes in use (mounted) by the node.
	VolumesInUse []string `pulumi:"volumesInUse"`

}

var _NodeStatusType = reflect.TypeOf((*NodeStatus)(nil)).Elem()

// NodeStatusInput represents an input type that resolves to a NodeStatus.
type NodeStatusInput interface {
	ElementType() reflect.Type

	ToNodeStatusOutput() NodeStatusOutput
	ToNodeStatusOutputWithContext(ctx context.Context) NodeStatusOutput
}

// NodeStatusArgs is a NodeStatusInput whose fields are all Input types.
type NodeStatusArgs struct {
	// List of addresses reachable to the node. Queried from cloud provider, if available. More info:
	// https://kubernetes.io/docs/concepts/nodes/node/#addresses Note: This field is declared as
	// mergeable, but the merge key is not sufficiently unique, which can cause data corruption when it
	// is merged. Callers should instead use a full-replacement patch. See http://pr.k8s.io/79391 for
	// an example.
	Addresses NodeAddressArrayInput `pulumi:"addresses"`

	// Allocatable represents the resources of a node that are available for scheduling. Defaults to
	// Capacity.
	Allocatable pulumi.StringMapInput `pulumi:"allocatable"`

	// Capacity represents the total resources of a node. More info:
	// https://kubernetes.io/docs/concepts/storage/persistent-volumes#capacity
	Capacity pulumi.StringMapInput `pulumi:"capacity"`

	// Conditions is an array of current observed node conditions. More info:
	// https://kubernetes.io/docs/concepts/nodes/node/#condition
	Conditions NodeConditionArrayInput `pulumi:"conditions"`

	// Status of the config assigned to the node via the dynamic Kubelet config feature.
	Config NodeConfigStatusInput `pulumi:"config"`

	// Endpoints of daemons running on the Node.
	DaemonEndpoints NodeDaemonEndpointsInput `pulumi:"daemonEndpoints"`

	// List of container images on this node
	Images ContainerImageArrayInput `pulumi:"images"`

	// Set of ids/uuids to uniquely identify the node. More info:
	// https://kubernetes.io/docs/concepts/nodes/node/#info
	NodeInfo NodeSystemInfoInput `pulumi:"nodeInfo"`

	// NodePhase is the recently observed lifecycle phase of the node. More info:
	// https://kubernetes.io/docs/concepts/nodes/node/#phase The field is never populated, and now is
	// deprecated.
	Phase pulumi.StringInput `pulumi:"phase"`

	// List of volumes that are attached to the node.
	VolumesAttached AttachedVolumeArrayInput `pulumi:"volumesAttached"`

	// List of attachable volumes in use (mounted) by the node.
	VolumesInUse pulumi.StringArrayInput `pulumi:"volumesInUse"`

}

func (a NodeStatusArgs) ElementType() reflect.Type {
	return _NodeStatusType
}

func (a NodeStatusArgs) ToNodeStatusOutput() NodeStatusOutput {
	return pulumi.ToOutput(a).(NodeStatusOutput)
}

func (a NodeStatusArgs) ToNodeStatusOutputWithContext(ctx context.Context) NodeStatusOutput {
	return pulumi.ToOutputWithContext(ctx, a).(NodeStatusOutput)
}

// NodeStatusOutput is an output type that resolves to a Input.
type NodeStatusOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(NodeStatusOutput{}) }

func (NodeStatusOutput) ElementType() reflect.Type {
	return _NodeStatusType
}

func (o NodeStatusOutput) Addresses() NodeAddressArrayOutput {
	return o.Apply(func(v NodeStatus) []NodeAddress {
		return v.Addresses
	}).(NodeAddressArrayOutput)
}

func (o NodeStatusOutput) Allocatable() pulumi.StringMapOutput {
	return o.Apply(func(v NodeStatus) map[string]string {
		return v.Allocatable
	}).(pulumi.StringMapOutput)
}

func (o NodeStatusOutput) Capacity() pulumi.StringMapOutput {
	return o.Apply(func(v NodeStatus) map[string]string {
		return v.Capacity
	}).(pulumi.StringMapOutput)
}

func (o NodeStatusOutput) Conditions() NodeConditionArrayOutput {
	return o.Apply(func(v NodeStatus) []NodeCondition {
		return v.Conditions
	}).(NodeConditionArrayOutput)
}

func (o NodeStatusOutput) Config() NodeConfigStatusOutput {
	return o.Apply(func(v NodeStatus) *NodeConfigStatus {
		return v.Config
	}).(NodeConfigStatusOutput)
}

func (o NodeStatusOutput) DaemonEndpoints() NodeDaemonEndpointsOutput {
	return o.Apply(func(v NodeStatus) *NodeDaemonEndpoints {
		return v.DaemonEndpoints
	}).(NodeDaemonEndpointsOutput)
}

func (o NodeStatusOutput) Images() ContainerImageArrayOutput {
	return o.Apply(func(v NodeStatus) []ContainerImage {
		return v.Images
	}).(ContainerImageArrayOutput)
}

func (o NodeStatusOutput) NodeInfo() NodeSystemInfoOutput {
	return o.Apply(func(v NodeStatus) *NodeSystemInfo {
		return v.NodeInfo
	}).(NodeSystemInfoOutput)
}

func (o NodeStatusOutput) Phase() pulumi.StringOutput {
	return o.Apply(func(v NodeStatus) *string {
		return v.Phase
	}).(pulumi.StringOutput)
}

func (o NodeStatusOutput) VolumesAttached() AttachedVolumeArrayOutput {
	return o.Apply(func(v NodeStatus) []AttachedVolume {
		return v.VolumesAttached
	}).(AttachedVolumeArrayOutput)
}

func (o NodeStatusOutput) VolumesInUse() pulumi.StringArrayOutput {
	return o.Apply(func(v NodeStatus) []string {
		return v.VolumesInUse
	}).(pulumi.StringArrayOutput)
}
