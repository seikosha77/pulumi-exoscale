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
    public sealed class GetInstancePoolListPoolResult
    {
        public readonly ImmutableArray<string> AffinityGroupIds;
        public readonly string DeployTargetId;
        public readonly string Description;
        public readonly int DiskSize;
        public readonly ImmutableArray<string> ElasticIpIds;
        public readonly string? Id;
        public readonly string InstancePrefix;
        public readonly string InstanceType;
        public readonly ImmutableArray<Outputs.GetInstancePoolListPoolInstanceResult> Instances;
        public readonly bool Ipv6;
        public readonly string KeyPair;
        public readonly ImmutableDictionary<string, string>? Labels;
        public readonly string? Name;
        public readonly ImmutableArray<string> NetworkIds;
        public readonly ImmutableArray<string> SecurityGroupIds;
        public readonly int Size;
        public readonly string State;
        public readonly string TemplateId;
        public readonly string UserData;
        public readonly string Zone;

        [OutputConstructor]
        private GetInstancePoolListPoolResult(
            ImmutableArray<string> affinityGroupIds,

            string deployTargetId,

            string description,

            int diskSize,

            ImmutableArray<string> elasticIpIds,

            string? id,

            string instancePrefix,

            string instanceType,

            ImmutableArray<Outputs.GetInstancePoolListPoolInstanceResult> instances,

            bool ipv6,

            string keyPair,

            ImmutableDictionary<string, string>? labels,

            string? name,

            ImmutableArray<string> networkIds,

            ImmutableArray<string> securityGroupIds,

            int size,

            string state,

            string templateId,

            string userData,

            string zone)
        {
            AffinityGroupIds = affinityGroupIds;
            DeployTargetId = deployTargetId;
            Description = description;
            DiskSize = diskSize;
            ElasticIpIds = elasticIpIds;
            Id = id;
            InstancePrefix = instancePrefix;
            InstanceType = instanceType;
            Instances = instances;
            Ipv6 = ipv6;
            KeyPair = keyPair;
            Labels = labels;
            Name = name;
            NetworkIds = networkIds;
            SecurityGroupIds = securityGroupIds;
            Size = size;
            State = state;
            TemplateId = templateId;
            UserData = userData;
            Zone = zone;
        }
    }
}
