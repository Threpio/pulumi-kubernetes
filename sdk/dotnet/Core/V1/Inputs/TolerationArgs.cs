// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Core.V1
{

    /// <summary>
    /// The pod this Toleration is attached to tolerates any taint that matches the triple &lt;key,value,effect&gt; using the matching operator &lt;operator&gt;.
    /// </summary>
    public class TolerationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.
        /// 
        /// Possible enum values:
        ///  - `"NoExecute"` Evict any already-running pods that do not tolerate the taint. Currently enforced by NodeController.
        ///  - `"NoSchedule"` Do not allow new pods to schedule onto the node unless they tolerate the taint, but allow all pods submitted to Kubelet without going through the scheduler to start, and allow all already-running pods to continue running. Enforced by the scheduler.
        ///  - `"PreferNoSchedule"` Like TaintEffectNoSchedule, but the scheduler tries not to schedule new pods onto the node, rather than prohibiting new pods from scheduling onto the node entirely. Enforced by the scheduler.
        /// </summary>
        [Input("effect")]
        public Input<string>? Effect { get; set; }

        /// <summary>
        /// Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.
        /// 
        /// Possible enum values:
        ///  - `"Equal"`
        ///  - `"Exists"`
        /// </summary>
        [Input("operator")]
        public Input<string>? Operator { get; set; }

        /// <summary>
        /// TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.
        /// </summary>
        [Input("tolerationSeconds")]
        public Input<int>? TolerationSeconds { get; set; }

        /// <summary>
        /// Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public TolerationArgs()
        {
        }
    }
}
