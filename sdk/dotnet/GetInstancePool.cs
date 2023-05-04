// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Exoscale
{
    public static class GetInstancePool
    {
        public static Task<GetInstancePoolResult> InvokeAsync(GetInstancePoolArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInstancePoolResult>("exoscale:index/getInstancePool:getInstancePool", args ?? new GetInstancePoolArgs(), options.WithDefaults());

        public static Output<GetInstancePoolResult> Invoke(GetInstancePoolInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInstancePoolResult>("exoscale:index/getInstancePool:getInstancePool", args ?? new GetInstancePoolInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstancePoolArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The instance pool ID to match (conflicts with `name`).
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        [Input("labels")]
        private Dictionary<string, string>? _labels;

        /// <summary>
        /// A map of key/value labels.
        /// </summary>
        public Dictionary<string, string> Labels
        {
            get => _labels ?? (_labels = new Dictionary<string, string>());
            set => _labels = value;
        }

        /// <summary>
        /// The pool name to match (conflicts with `id`).
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        /// </summary>
        [Input("zone", required: true)]
        public string Zone { get; set; } = null!;

        public GetInstancePoolArgs()
        {
        }
        public static new GetInstancePoolArgs Empty => new GetInstancePoolArgs();
    }

    public sealed class GetInstancePoolInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The instance pool ID to match (conflicts with `name`).
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// A map of key/value labels.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The pool name to match (conflicts with `id`).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public GetInstancePoolInvokeArgs()
        {
        }
        public static new GetInstancePoolInvokeArgs Empty => new GetInstancePoolInvokeArgs();
    }


    [OutputType]
    public sealed class GetInstancePoolResult
    {
        /// <summary>
        /// The list of attached exoscale*anti*affinity_group (IDs).
        /// </summary>
        public readonly ImmutableArray<string> AffinityGroupIds;
        /// <summary>
        /// The deploy target ID.
        /// </summary>
        public readonly string DeployTargetId;
        /// <summary>
        /// The instance pool description.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The managed instances disk size.
        /// </summary>
        public readonly int DiskSize;
        /// <summary>
        /// The list of attached exoscale*elastic*ip (IDs).
        /// </summary>
        public readonly ImmutableArray<string> ElasticIpIds;
        /// <summary>
        /// The instance pool ID to match (conflicts with `name`).
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The string used to prefix the managed instances name.
        /// </summary>
        public readonly string InstancePrefix;
        /// <summary>
        /// The managed instances type.
        /// </summary>
        public readonly string InstanceType;
        /// <summary>
        /// The list of managed instances. Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstancePoolInstanceResult> Instances;
        /// <summary>
        /// Whether IPv6 is enabled on managed instances.
        /// </summary>
        public readonly bool Ipv6;
        /// <summary>
        /// The exoscale*ssh*key (name) authorized on the managed instances.
        /// </summary>
        public readonly string KeyPair;
        /// <summary>
        /// A map of key/value labels.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Labels;
        /// <summary>
        /// The pool name to match (conflicts with `id`).
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The list of attached exoscale*private*network (IDs).
        /// </summary>
        public readonly ImmutableArray<string> NetworkIds;
        /// <summary>
        /// The list of attached exoscale*security*group (IDs).
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroupIds;
        /// <summary>
        /// The number managed instances.
        /// </summary>
        public readonly int Size;
        /// <summary>
        /// The pool state.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The managed instances exoscale*compute*template ID.
        /// </summary>
        public readonly string TemplateId;
        /// <summary>
        /// [cloud-init](http://cloudinit.readthedocs.io/en/latest/) configuration.
        /// </summary>
        public readonly string UserData;
        /// <summary>
        /// The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private GetInstancePoolResult(
            ImmutableArray<string> affinityGroupIds,

            string deployTargetId,

            string description,

            int diskSize,

            ImmutableArray<string> elasticIpIds,

            string? id,

            string instancePrefix,

            string instanceType,

            ImmutableArray<Outputs.GetInstancePoolInstanceResult> instances,

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
