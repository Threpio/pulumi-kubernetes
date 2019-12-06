// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// NodeSystemInfo is a set of ids/uuids to uniquely identify the node.
type NodeSystemInfo struct {
	// The Architecture reported by the node
	Architecture string `pulumi:"architecture"`

	// Boot ID reported by the node.
	BootID string `pulumi:"bootID"`

	// ContainerRuntime Version reported by the node through runtime remote API (e.g. docker://1.5.0).
	ContainerRuntimeVersion string `pulumi:"containerRuntimeVersion"`

	// Kernel Version reported by the node from 'uname -r' (e.g. 3.16.0-0.bpo.4-amd64).
	KernelVersion string `pulumi:"kernelVersion"`

	// KubeProxy Version reported by the node.
	KubeProxyVersion string `pulumi:"kubeProxyVersion"`

	// Kubelet Version reported by the node.
	KubeletVersion string `pulumi:"kubeletVersion"`

	// MachineID reported by the node. For unique machine identification in the cluster this field is
	// preferred. Learn more from man(5) machine-id:
	// http://man7.org/linux/man-pages/man5/machine-id.5.html
	MachineID string `pulumi:"machineID"`

	// The Operating System reported by the node
	OperatingSystem string `pulumi:"operatingSystem"`

	// OS Image reported by the node from /etc/os-release (e.g. Debian GNU/Linux 7 (wheezy)).
	OsImage string `pulumi:"osImage"`

	// SystemUUID reported by the node. For unique machine identification MachineID is preferred. This
	// field is specific to Red Hat hosts
	// https://access.redhat.com/documentation/en-US/Red_Hat_Subscription_Management/1/html/RHSM/getting-system-uuid.html
	SystemUUID string `pulumi:"systemUUID"`

}

var _NodeSystemInfoType = reflect.TypeOf((*NodeSystemInfo)(nil)).Elem()

// NodeSystemInfoInput represents an input type that resolves to a NodeSystemInfo.
type NodeSystemInfoInput interface {
	ElementType() reflect.Type

	ToNodeSystemInfoOutput() NodeSystemInfoOutput
	ToNodeSystemInfoOutputWithContext(ctx context.Context) NodeSystemInfoOutput
}

// NodeSystemInfoArgs is a NodeSystemInfoInput whose fields are all Input types.
type NodeSystemInfoArgs struct {
	// The Architecture reported by the node
	Architecture pulumi.StringInput `pulumi:"architecture"`

	// Boot ID reported by the node.
	BootID pulumi.StringInput `pulumi:"bootID"`

	// ContainerRuntime Version reported by the node through runtime remote API (e.g. docker://1.5.0).
	ContainerRuntimeVersion pulumi.StringInput `pulumi:"containerRuntimeVersion"`

	// Kernel Version reported by the node from 'uname -r' (e.g. 3.16.0-0.bpo.4-amd64).
	KernelVersion pulumi.StringInput `pulumi:"kernelVersion"`

	// KubeProxy Version reported by the node.
	KubeProxyVersion pulumi.StringInput `pulumi:"kubeProxyVersion"`

	// Kubelet Version reported by the node.
	KubeletVersion pulumi.StringInput `pulumi:"kubeletVersion"`

	// MachineID reported by the node. For unique machine identification in the cluster this field is
	// preferred. Learn more from man(5) machine-id:
	// http://man7.org/linux/man-pages/man5/machine-id.5.html
	MachineID pulumi.StringInput `pulumi:"machineID"`

	// The Operating System reported by the node
	OperatingSystem pulumi.StringInput `pulumi:"operatingSystem"`

	// OS Image reported by the node from /etc/os-release (e.g. Debian GNU/Linux 7 (wheezy)).
	OsImage pulumi.StringInput `pulumi:"osImage"`

	// SystemUUID reported by the node. For unique machine identification MachineID is preferred. This
	// field is specific to Red Hat hosts
	// https://access.redhat.com/documentation/en-US/Red_Hat_Subscription_Management/1/html/RHSM/getting-system-uuid.html
	SystemUUID pulumi.StringInput `pulumi:"systemUUID"`

}

func (a NodeSystemInfoArgs) ElementType() reflect.Type {
	return _NodeSystemInfoType
}

func (a NodeSystemInfoArgs) ToNodeSystemInfoOutput() NodeSystemInfoOutput {
	return pulumi.ToOutput(a).(NodeSystemInfoOutput)
}

func (a NodeSystemInfoArgs) ToNodeSystemInfoOutputWithContext(ctx context.Context) NodeSystemInfoOutput {
	return pulumi.ToOutputWithContext(ctx, a).(NodeSystemInfoOutput)
}

// NodeSystemInfoOutput is an output type that resolves to a Input.
type NodeSystemInfoOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(NodeSystemInfoOutput{}) }

func (NodeSystemInfoOutput) ElementType() reflect.Type {
	return _NodeSystemInfoType
}

func (o NodeSystemInfoOutput) Architecture() pulumi.StringOutput {
	return o.Apply(func(v NodeSystemInfo) string {
		return v.Architecture
	}).(pulumi.StringOutput)
}

func (o NodeSystemInfoOutput) BootID() pulumi.StringOutput {
	return o.Apply(func(v NodeSystemInfo) string {
		return v.BootID
	}).(pulumi.StringOutput)
}

func (o NodeSystemInfoOutput) ContainerRuntimeVersion() pulumi.StringOutput {
	return o.Apply(func(v NodeSystemInfo) string {
		return v.ContainerRuntimeVersion
	}).(pulumi.StringOutput)
}

func (o NodeSystemInfoOutput) KernelVersion() pulumi.StringOutput {
	return o.Apply(func(v NodeSystemInfo) string {
		return v.KernelVersion
	}).(pulumi.StringOutput)
}

func (o NodeSystemInfoOutput) KubeProxyVersion() pulumi.StringOutput {
	return o.Apply(func(v NodeSystemInfo) string {
		return v.KubeProxyVersion
	}).(pulumi.StringOutput)
}

func (o NodeSystemInfoOutput) KubeletVersion() pulumi.StringOutput {
	return o.Apply(func(v NodeSystemInfo) string {
		return v.KubeletVersion
	}).(pulumi.StringOutput)
}

func (o NodeSystemInfoOutput) MachineID() pulumi.StringOutput {
	return o.Apply(func(v NodeSystemInfo) string {
		return v.MachineID
	}).(pulumi.StringOutput)
}

func (o NodeSystemInfoOutput) OperatingSystem() pulumi.StringOutput {
	return o.Apply(func(v NodeSystemInfo) string {
		return v.OperatingSystem
	}).(pulumi.StringOutput)
}

func (o NodeSystemInfoOutput) OsImage() pulumi.StringOutput {
	return o.Apply(func(v NodeSystemInfo) string {
		return v.OsImage
	}).(pulumi.StringOutput)
}

func (o NodeSystemInfoOutput) SystemUUID() pulumi.StringOutput {
	return o.Apply(func(v NodeSystemInfo) string {
		return v.SystemUUID
	}).(pulumi.StringOutput)
}
