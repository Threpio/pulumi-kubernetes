// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
	metaV1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
)

// RoleBinding references a role, but does not contain it.  It can reference a Role in the same
// namespace or a ClusterRole in the global namespace. It adds who information via Subjects and
// namespace information by which namespace it exists in.  RoleBindings in a given namespace only
// have effect in that namespace.
type RoleBinding struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should
	// convert recognized schemas to the latest internal value, and may reject unrecognized values.
	// More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer
	// this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More
	// info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`

	// Standard object's metadata.
	Metadata metaV1.ObjectMetaOutput `pulumi:"metadata"`

	// RoleRef can reference a Role in the current namespace or a ClusterRole in the global namespace.
	// If the RoleRef cannot be resolved, the Authorizer must return an error.
	RoleRef RoleRefOutput `pulumi:"roleRef"`

	// Subjects holds references to the objects the role applies to.
	Subjects SubjectArrayOutput `pulumi:"subjects"`

}

// RoleBindingArgs is the set of arguments needed to create a RoleBinding resource.
type RoleBindingArgs struct {
	// RoleRef can reference a Role in the current namespace or a ClusterRole in the global namespace.
	// If the RoleRef cannot be resolved, the Authorizer must return an error.
	RoleRef RoleRefInput `pulumi:"roleRef"`

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

	// Standard object's metadata.
	Metadata metaV1.ObjectMetaInput `pulumi:"metadata"`

	// Subjects holds references to the objects the role applies to.
	Subjects SubjectArrayInput `pulumi:"subjects"`

}

// NewRoleBinding creates a RoleBinding resource with the given unique name, arguments, and options.
func NewRoleBinding(ctx *pulumi.Context, name string, args *RoleBindingArgs, opts ...pulumi.ResourceOption) (*RoleBinding, error) {
	inputs := map[string]pulumi.Input{}
	if args == nil || args.RoleRef == nil {
		return nil, errors.New("missing required argument 'RoleRef'")
	}
	if args != nil {
		args.ApiVersion = pulumi.String("rbac.authorization.k8s.io/v1beta1")
		args.Kind = pulumi.String("RoleBinding")
		inputs["roleRef"] = args.RoleRef.ToRoleRefOutput()
		if i := args.ApiVersion; i != nil {
			inputs["apiVersion"] = i.ToStringOutput()
		}
		if i := args.Kind; i != nil {
			inputs["kind"] = i.ToStringOutput()
		}
		if i := args.Metadata; i != nil {
			inputs["metadata"] = i.ToObjectMetaOutput()
		}
		if i := args.Subjects; i != nil {
			inputs["subjects"] = i.ToSubjectArrayOutput()
		}
	}
	var resource RoleBinding
	err := ctx.RegisterResource("kubernetes:rbac.authorization.k8s.io/v1beta1:RoleBinding", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoleBinding gets an existing RoleBinding resource's state with the given name and ID.
func GetRoleBinding(ctx *pulumi.Context, name string, id pulumi.IDInput, opts ...pulumi.ResourceOption) (*RoleBinding, error) {
	var resource RoleBinding
	err := ctx.ReadResource("kubernetes:rbac.authorization.k8s.io/v1beta1:RoleBinding", name, id, nil, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// RoleBinding references a role, but does not contain it.  It can reference a Role in the same
// namespace or a ClusterRole in the global namespace. It adds who information via Subjects and
// namespace information by which namespace it exists in.  RoleBindings in a given namespace only
// have effect in that namespace.
type RoleBindingProperty struct {
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

	// Standard object's metadata.
	Metadata *metaV1.ObjectMeta `pulumi:"metadata"`

	// RoleRef can reference a Role in the current namespace or a ClusterRole in the global namespace.
	// If the RoleRef cannot be resolved, the Authorizer must return an error.
	RoleRef RoleRef `pulumi:"roleRef"`

	// Subjects holds references to the objects the role applies to.
	Subjects []Subject `pulumi:"subjects"`

}

var _RoleBindingPropertyType = reflect.TypeOf((*RoleBindingProperty)(nil)).Elem()

// RoleBindingPropertyInput represents an input type that resolves to a RoleBindingProperty.
type RoleBindingPropertyInput interface {
	ElementType() reflect.Type

	ToRoleBindingPropertyOutput() RoleBindingPropertyOutput
	ToRoleBindingPropertyOutputWithContext(ctx context.Context) RoleBindingPropertyOutput
}

// RoleBindingPropertyArgs is a RoleBindingPropertyInput whose fields are all Input types.
type RoleBindingPropertyArgs struct {
	// RoleRef can reference a Role in the current namespace or a ClusterRole in the global namespace.
	// If the RoleRef cannot be resolved, the Authorizer must return an error.
	RoleRef RoleRefInput `pulumi:"roleRef"`

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

	// Standard object's metadata.
	Metadata metaV1.ObjectMetaInput `pulumi:"metadata"`

	// Subjects holds references to the objects the role applies to.
	Subjects SubjectArrayInput `pulumi:"subjects"`

}

func (a RoleBindingPropertyArgs) ElementType() reflect.Type {
	return _RoleBindingPropertyType
}

func (a RoleBindingPropertyArgs) ToRoleBindingPropertyOutput() RoleBindingPropertyOutput {
	return pulumi.ToOutput(a).(RoleBindingPropertyOutput)
}

func (a RoleBindingPropertyArgs) ToRoleBindingPropertyOutputWithContext(ctx context.Context) RoleBindingPropertyOutput {
	return pulumi.ToOutputWithContext(ctx, a).(RoleBindingPropertyOutput)
}

// RoleBindingPropertyOutput is an output type that resolves to a Input.
type RoleBindingPropertyOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(RoleBindingPropertyOutput{}) }

func (RoleBindingPropertyOutput) ElementType() reflect.Type {
	return _RoleBindingPropertyType
}

func (o RoleBindingPropertyOutput) ApiVersion() pulumi.StringOutput {
	return o.Apply(func(v RoleBindingProperty) *string {
		return v.ApiVersion
	}).(pulumi.StringOutput)
}

func (o RoleBindingPropertyOutput) Kind() pulumi.StringOutput {
	return o.Apply(func(v RoleBindingProperty) *string {
		return v.Kind
	}).(pulumi.StringOutput)
}

func (o RoleBindingPropertyOutput) Metadata() metaV1.ObjectMetaOutput {
	return o.Apply(func(v RoleBindingProperty) *metaV1.ObjectMeta {
		return v.Metadata
	}).(metaV1.ObjectMetaOutput)
}

func (o RoleBindingPropertyOutput) RoleRef() RoleRefOutput {
	return o.Apply(func(v RoleBindingProperty) RoleRef {
		return v.RoleRef
	}).(RoleRefOutput)
}

func (o RoleBindingPropertyOutput) Subjects() SubjectArrayOutput {
	return o.Apply(func(v RoleBindingProperty) []Subject {
		return v.Subjects
	}).(SubjectArrayOutput)
}

var _RoleBindingPropertyArrayType = reflect.TypeOf((*[]RoleBindingProperty)(nil)).Elem()

type RoleBindingPropertyArrayInput interface {
	ElementType() reflect.Type

	ToRoleBindingPropertyArrayOutput() RoleBindingPropertyArrayOutput
	ToRoleBindingPropertyArrayOutputWithContext(ctx context.Context) RoleBindingPropertyArrayOutput
}

type RoleBindingPropertyArray []RoleBindingPropertyInput

func (a RoleBindingPropertyArray) ElementType() reflect.Type {
	return _RoleBindingPropertyArrayType
}

func (a RoleBindingPropertyArray) ToRoleBindingPropertyArrayOutput() RoleBindingPropertyArrayOutput {
	return pulumi.ToOutput(a).(RoleBindingPropertyArrayOutput)
}

func (a RoleBindingPropertyArray) ToRoleBindingPropertyArrayOutputWithContext(ctx context.Context) RoleBindingPropertyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(RoleBindingPropertyArrayOutput)
}

type RoleBindingPropertyArrayOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(RoleBindingPropertyArrayOutput{}) }

func (RoleBindingPropertyArrayOutput) ElementType() reflect.Type {
	return _RoleBindingPropertyArrayType
}

func (o RoleBindingPropertyArrayOutput) Index(i pulumi.IntInput) RoleBindingPropertyOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) RoleBindingProperty {
		return vs[0].([]RoleBindingProperty)[vs[1].(int)]
	}).(RoleBindingPropertyOutput)
}