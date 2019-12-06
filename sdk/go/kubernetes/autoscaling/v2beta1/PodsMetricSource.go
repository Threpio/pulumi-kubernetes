// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	metaV1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
)

// PodsMetricSource indicates how to scale on a metric describing each pod in the current scale
// target (for example, transactions-processed-per-second). The values will be averaged together
// before being compared to the target value.
type PodsMetricSource struct {
	// metricName is the name of the metric in question
	MetricName string `pulumi:"metricName"`

	// selector is the string-encoded form of a standard kubernetes label selector for the given metric
	// When set, it is passed as an additional parameter to the metrics server for more specific
	// metrics scoping When unset, just the metricName will be used to gather metrics.
	Selector *metaV1.LabelSelector `pulumi:"selector"`

	// targetAverageValue is the target value of the average of the metric across all relevant pods (as
	// a quantity)
	TargetAverageValue string `pulumi:"targetAverageValue"`

}

var _PodsMetricSourceType = reflect.TypeOf((*PodsMetricSource)(nil)).Elem()

// PodsMetricSourceInput represents an input type that resolves to a PodsMetricSource.
type PodsMetricSourceInput interface {
	ElementType() reflect.Type

	ToPodsMetricSourceOutput() PodsMetricSourceOutput
	ToPodsMetricSourceOutputWithContext(ctx context.Context) PodsMetricSourceOutput
}

// PodsMetricSourceArgs is a PodsMetricSourceInput whose fields are all Input types.
type PodsMetricSourceArgs struct {
	// metricName is the name of the metric in question
	MetricName pulumi.StringInput `pulumi:"metricName"`

	// targetAverageValue is the target value of the average of the metric across all relevant pods (as
	// a quantity)
	TargetAverageValue pulumi.StringInput `pulumi:"targetAverageValue"`

	// selector is the string-encoded form of a standard kubernetes label selector for the given metric
	// When set, it is passed as an additional parameter to the metrics server for more specific
	// metrics scoping When unset, just the metricName will be used to gather metrics.
	Selector metaV1.LabelSelectorInput `pulumi:"selector"`

}

func (a PodsMetricSourceArgs) ElementType() reflect.Type {
	return _PodsMetricSourceType
}

func (a PodsMetricSourceArgs) ToPodsMetricSourceOutput() PodsMetricSourceOutput {
	return pulumi.ToOutput(a).(PodsMetricSourceOutput)
}

func (a PodsMetricSourceArgs) ToPodsMetricSourceOutputWithContext(ctx context.Context) PodsMetricSourceOutput {
	return pulumi.ToOutputWithContext(ctx, a).(PodsMetricSourceOutput)
}

// PodsMetricSourceOutput is an output type that resolves to a Input.
type PodsMetricSourceOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(PodsMetricSourceOutput{}) }

func (PodsMetricSourceOutput) ElementType() reflect.Type {
	return _PodsMetricSourceType
}

func (o PodsMetricSourceOutput) MetricName() pulumi.StringOutput {
	return o.Apply(func(v PodsMetricSource) string {
		return v.MetricName
	}).(pulumi.StringOutput)
}

func (o PodsMetricSourceOutput) Selector() metaV1.LabelSelectorOutput {
	return o.Apply(func(v PodsMetricSource) *metaV1.LabelSelector {
		return v.Selector
	}).(metaV1.LabelSelectorOutput)
}

func (o PodsMetricSourceOutput) TargetAverageValue() pulumi.StringOutput {
	return o.Apply(func(v PodsMetricSource) string {
		return v.TargetAverageValue
	}).(pulumi.StringOutput)
}
