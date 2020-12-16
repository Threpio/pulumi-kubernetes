// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2Beta2
{

    [OutputType]
    public sealed class ContainerResourceMetricStatus
    {
        /// <summary>
        /// Container is the name of the container in the pods of the scaling target
        /// </summary>
        public readonly string Container;
        /// <summary>
        /// current contains the current value for the given metric
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2Beta2.MetricValueStatus Current;
        /// <summary>
        /// Name is the name of the resource in question.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private ContainerResourceMetricStatus(
            string container,

            Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2Beta2.MetricValueStatus current,

            string name)
        {
            Container = container;
            Current = current;
            Name = name;
        }
    }
}