// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Core.V1
{

    /// <summary>
    /// Represents a Quobyte mount that lasts the lifetime of a pod. Quobyte volumes do not support ownership management or SELinux relabeling.
    /// </summary>
    [OutputType]
    public sealed class QuobyteVolumeSource
    {
        /// <summary>
        /// Group to map volume access to Default is no group
        /// </summary>
        public readonly string Group;
        /// <summary>
        /// ReadOnly here will force the Quobyte volume to be mounted with read-only permissions. Defaults to false.
        /// </summary>
        public readonly bool ReadOnly;
        /// <summary>
        /// Registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes
        /// </summary>
        public readonly string Registry;
        /// <summary>
        /// Tenant owning the given Quobyte volume in the Backend Used with dynamically provisioned Quobyte volumes, value is set by the plugin
        /// </summary>
        public readonly string Tenant;
        /// <summary>
        /// User to map volume access to Defaults to serivceaccount user
        /// </summary>
        public readonly string User;
        /// <summary>
        /// Volume is a string that references an already created Quobyte volume by name.
        /// </summary>
        public readonly string Volume;

        [OutputConstructor]
        private QuobyteVolumeSource(
            string group,

            bool readOnly,

            string registry,

            string tenant,

            string user,

            string volume)
        {
            Group = group;
            ReadOnly = readOnly;
            Registry = registry;
            Tenant = tenant;
            User = user;
            Volume = volume;
        }
    }
}
