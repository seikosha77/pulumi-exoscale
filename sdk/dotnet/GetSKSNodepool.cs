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
    public static class GetSKSNodepool
    {
        public static Task<GetSKSNodepoolResult> InvokeAsync(GetSKSNodepoolArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSKSNodepoolResult>("exoscale:index/getSKSNodepool:getSKSNodepool", args ?? new GetSKSNodepoolArgs(), options.WithDefaults());

        public static Output<GetSKSNodepoolResult> Invoke(GetSKSNodepoolInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSKSNodepoolResult>("exoscale:index/getSKSNodepool:getSKSNodepool", args ?? new GetSKSNodepoolInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSKSNodepoolArgs : global::Pulumi.InvokeArgs
    {
        [Input("antiAffinityGroupIds")]
        private List<string>? _antiAffinityGroupIds;

        /// <summary>
        /// A list of exoscale*anti*affinity_group (IDs) to be attached to the managed instances.
        /// </summary>
        public List<string> AntiAffinityGroupIds
        {
            get => _antiAffinityGroupIds ?? (_antiAffinityGroupIds = new List<string>());
            set => _antiAffinityGroupIds = value;
        }

        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        /// <summary>
        /// The pool creation date.
        /// </summary>
        [Input("createdAt")]
        public string? CreatedAt { get; set; }

        /// <summary>
        /// A deploy target ID.
        /// </summary>
        [Input("deployTargetId")]
        public string? DeployTargetId { get; set; }

        /// <summary>
        /// A free-form text describing the pool.
        /// </summary>
        [Input("description")]
        public string? Description { get; set; }

        /// <summary>
        /// The managed instances disk size (GiB; default: `50`).
        /// </summary>
        [Input("diskSize")]
        public int? DiskSize { get; set; }

        /// <summary>
        /// The ID of this resource.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// The underlying exoscale*instance*pool ID.
        /// </summary>
        [Input("instancePoolId")]
        public string? InstancePoolId { get; set; }

        /// <summary>
        /// The string used to prefix the managed instances name (default `pool`).
        /// </summary>
        [Input("instancePrefix")]
        public string? InstancePrefix { get; set; }

        /// <summary>
        /// The managed compute instances type (`&lt;family&gt;.&lt;size&gt;`, e.g. `standard.medium`; use the [Exoscale CLI](https://github.com/exoscale/cli/) - `exo compute instance-type list` - for the list of available types).
        /// </summary>
        [Input("instanceType")]
        public string? InstanceType { get; set; }

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

        [Input("name")]
        public string? Name { get; set; }

        [Input("privateNetworkIds")]
        private List<string>? _privateNetworkIds;

        /// <summary>
        /// A list of exoscale*private*network (IDs) to be attached to the managed instances.
        /// </summary>
        public List<string> PrivateNetworkIds
        {
            get => _privateNetworkIds ?? (_privateNetworkIds = new List<string>());
            set => _privateNetworkIds = value;
        }

        [Input("securityGroupIds")]
        private List<string>? _securityGroupIds;

        /// <summary>
        /// A list of exoscale*security*group (IDs) to be attached to the managed instances.
        /// </summary>
        public List<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new List<string>());
            set => _securityGroupIds = value;
        }

        [Input("size")]
        public int? Size { get; set; }

        /// <summary>
        /// The current pool state.
        /// </summary>
        [Input("state")]
        public string? State { get; set; }

        [Input("taints")]
        private Dictionary<string, string>? _taints;

        /// <summary>
        /// A map of key/value Kubernetes [taints](https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/) (`&lt;value&gt;:&lt;effect&gt;`).
        /// </summary>
        public Dictionary<string, string> Taints
        {
            get => _taints ?? (_taints = new Dictionary<string, string>());
            set => _taints = value;
        }

        /// <summary>
        /// The managed instances template ID.
        /// </summary>
        [Input("templateId")]
        public string? TemplateId { get; set; }

        /// <summary>
        /// The managed instances version.
        /// </summary>
        [Input("version")]
        public string? Version { get; set; }

        [Input("zone", required: true)]
        public string Zone { get; set; } = null!;

        public GetSKSNodepoolArgs()
        {
        }
        public static new GetSKSNodepoolArgs Empty => new GetSKSNodepoolArgs();
    }

    public sealed class GetSKSNodepoolInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("antiAffinityGroupIds")]
        private InputList<string>? _antiAffinityGroupIds;

        /// <summary>
        /// A list of exoscale*anti*affinity_group (IDs) to be attached to the managed instances.
        /// </summary>
        public InputList<string> AntiAffinityGroupIds
        {
            get => _antiAffinityGroupIds ?? (_antiAffinityGroupIds = new InputList<string>());
            set => _antiAffinityGroupIds = value;
        }

        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// The pool creation date.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// A deploy target ID.
        /// </summary>
        [Input("deployTargetId")]
        public Input<string>? DeployTargetId { get; set; }

        /// <summary>
        /// A free-form text describing the pool.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The managed instances disk size (GiB; default: `50`).
        /// </summary>
        [Input("diskSize")]
        public Input<int>? DiskSize { get; set; }

        /// <summary>
        /// The ID of this resource.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The underlying exoscale*instance*pool ID.
        /// </summary>
        [Input("instancePoolId")]
        public Input<string>? InstancePoolId { get; set; }

        /// <summary>
        /// The string used to prefix the managed instances name (default `pool`).
        /// </summary>
        [Input("instancePrefix")]
        public Input<string>? InstancePrefix { get; set; }

        /// <summary>
        /// The managed compute instances type (`&lt;family&gt;.&lt;size&gt;`, e.g. `standard.medium`; use the [Exoscale CLI](https://github.com/exoscale/cli/) - `exo compute instance-type list` - for the list of available types).
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

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

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("privateNetworkIds")]
        private InputList<string>? _privateNetworkIds;

        /// <summary>
        /// A list of exoscale*private*network (IDs) to be attached to the managed instances.
        /// </summary>
        public InputList<string> PrivateNetworkIds
        {
            get => _privateNetworkIds ?? (_privateNetworkIds = new InputList<string>());
            set => _privateNetworkIds = value;
        }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// A list of exoscale*security*group (IDs) to be attached to the managed instances.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// The current pool state.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        [Input("taints")]
        private InputMap<string>? _taints;

        /// <summary>
        /// A map of key/value Kubernetes [taints](https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/) (`&lt;value&gt;:&lt;effect&gt;`).
        /// </summary>
        public InputMap<string> Taints
        {
            get => _taints ?? (_taints = new InputMap<string>());
            set => _taints = value;
        }

        /// <summary>
        /// The managed instances template ID.
        /// </summary>
        [Input("templateId")]
        public Input<string>? TemplateId { get; set; }

        /// <summary>
        /// The managed instances version.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public GetSKSNodepoolInvokeArgs()
        {
        }
        public static new GetSKSNodepoolInvokeArgs Empty => new GetSKSNodepoolInvokeArgs();
    }


    [OutputType]
    public sealed class GetSKSNodepoolResult
    {
        /// <summary>
        /// A list of exoscale*anti*affinity_group (IDs) to be attached to the managed instances.
        /// </summary>
        public readonly ImmutableArray<string> AntiAffinityGroupIds;
        public readonly string ClusterId;
        /// <summary>
        /// The pool creation date.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// A deploy target ID.
        /// </summary>
        public readonly string? DeployTargetId;
        /// <summary>
        /// A free-form text describing the pool.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The managed instances disk size (GiB; default: `50`).
        /// </summary>
        public readonly int? DiskSize;
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The underlying exoscale*instance*pool ID.
        /// </summary>
        public readonly string InstancePoolId;
        /// <summary>
        /// The string used to prefix the managed instances name (default `pool`).
        /// </summary>
        public readonly string? InstancePrefix;
        /// <summary>
        /// The managed compute instances type (`&lt;family&gt;.&lt;size&gt;`, e.g. `standard.medium`; use the [Exoscale CLI](https://github.com/exoscale/cli/) - `exo compute instance-type list` - for the list of available types).
        /// </summary>
        public readonly string? InstanceType;
        /// <summary>
        /// A map of key/value labels.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Labels;
        public readonly string? Name;
        /// <summary>
        /// A list of exoscale*private*network (IDs) to be attached to the managed instances.
        /// </summary>
        public readonly ImmutableArray<string> PrivateNetworkIds;
        /// <summary>
        /// A list of exoscale*security*group (IDs) to be attached to the managed instances.
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroupIds;
        public readonly int? Size;
        /// <summary>
        /// The current pool state.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// A map of key/value Kubernetes [taints](https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/) (`&lt;value&gt;:&lt;effect&gt;`).
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Taints;
        /// <summary>
        /// The managed instances template ID.
        /// </summary>
        public readonly string TemplateId;
        /// <summary>
        /// The managed instances version.
        /// </summary>
        public readonly string Version;
        public readonly string Zone;

        [OutputConstructor]
        private GetSKSNodepoolResult(
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
            Taints = taints;
            TemplateId = templateId;
            Version = version;
            Zone = zone;
        }
    }
}
