// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Exoscale.Outputs
{

    [OutputType]
    public sealed class GetSksNodepoolListNodepoolResult
    {
        public readonly ImmutableArray<string> AntiAffinityGroupIds;
        public readonly string ClusterId;
        public readonly string CreatedAt;
        public readonly string? DeployTargetId;
        public readonly string? Description;
        public readonly int? DiskSize;
        public readonly string? Id;
        public readonly string InstancePoolId;
        public readonly string? InstancePrefix;
        public readonly string? InstanceType;
        public readonly ImmutableDictionary<string, string>? Labels;
        public readonly string? Name;
        public readonly ImmutableArray<string> PrivateNetworkIds;
        public readonly ImmutableArray<string> SecurityGroupIds;
        public readonly int? Size;
        public readonly string State;
        public readonly bool? StorageLvm;
        public readonly ImmutableDictionary<string, string>? Taints;
        public readonly string TemplateId;
        public readonly string Version;
        public readonly string Zone;

        [OutputConstructor]
        private GetSksNodepoolListNodepoolResult(
            ImmutableArray<string> antiAffinityGroupIds,

            string clusterId,

            string createdAt,

            string? deployTargetId,

            string? description,

            int? diskSize,

            string? id,

            string instancePoolId,

            string? instancePrefix,

            string? instanceType,

            ImmutableDictionary<string, string>? labels,

            string? name,

            ImmutableArray<string> privateNetworkIds,

            ImmutableArray<string> securityGroupIds,

            int? size,

            string state,

            bool? storageLvm,

            ImmutableDictionary<string, string>? taints,

            string templateId,

            string version,

            string zone)
        {
            AntiAffinityGroupIds = antiAffinityGroupIds;
            ClusterId = clusterId;
            CreatedAt = createdAt;
            DeployTargetId = deployTargetId;
            Description = description;
            DiskSize = diskSize;
            Id = id;
            InstancePoolId = instancePoolId;
            InstancePrefix = instancePrefix;
            InstanceType = instanceType;
            Labels = labels;
            Name = name;
            PrivateNetworkIds = privateNetworkIds;
            SecurityGroupIds = securityGroupIds;
            Size = size;
            State = state;
            StorageLvm = storageLvm;
            Taints = taints;
            TemplateId = templateId;
            Version = version;
            Zone = zone;
        }
    }
}
