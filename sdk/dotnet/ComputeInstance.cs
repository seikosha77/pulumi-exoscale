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
    /// <summary>
    /// ## Import
    /// 
    /// An existing compute instance may be imported by `&lt;ID&gt;@&lt;zone&gt;`
    /// 
    /// ```sh
    ///  $ pulumi import exoscale:index/computeInstance:ComputeInstance \
    /// ```
    /// 
    ///  exoscale_compute_instance.my_instance \
    /// 
    ///  f81d4fae-7dec-11d0-a765-00a0c91e6bf6@ch-gva-2
    /// </summary>
    [ExoscaleResourceType("exoscale:index/computeInstance:ComputeInstance")]
    public partial class ComputeInstance : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A list of [exoscale_anti_affinity_group](./anti_affinity_group.md) (IDs) to attach to the instance (may only be set at
        /// creation time).
        /// </summary>
        [Output("antiAffinityGroupIds")]
        public Output<ImmutableArray<string>> AntiAffinityGroupIds { get; private set; } = null!;

        /// <summary>
        /// The instance creation date.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// A deploy target ID.
        /// </summary>
        [Output("deployTargetId")]
        public Output<string?> DeployTargetId { get; private set; } = null!;

        /// <summary>
        /// The instance disk size (GiB; at least `10`). **WARNING**: updating this attribute stops/restarts the instance.
        /// </summary>
        [Output("diskSize")]
        public Output<int> DiskSize { get; private set; } = null!;

        /// <summary>
        /// A list of [exoscale_elastic_ip](./elastic_ip.md) (IDs) to attach to the instance.
        /// </summary>
        [Output("elasticIpIds")]
        public Output<ImmutableArray<string>> ElasticIpIds { get; private set; } = null!;

        /// <summary>
        /// Enable IPv6 on the instance (boolean; default: `false`).
        /// </summary>
        [Output("ipv6")]
        public Output<bool?> Ipv6 { get; private set; } = null!;

        /// <summary>
        /// The instance (main network interface) IPv6 address (if enabled).
        /// </summary>
        [Output("ipv6Address")]
        public Output<string> Ipv6Address { get; private set; } = null!;

        /// <summary>
        /// A map of key/value labels.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// The compute instance name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Private network interfaces (may be specified multiple times). Structure is documented below.
        /// </summary>
        [Output("networkInterfaces")]
        public Output<ImmutableArray<Outputs.ComputeInstanceNetworkInterface>> NetworkInterfaces { get; private set; } = null!;

        /// <summary>
        /// A list of private networks (IDs) attached to the instance. Please use the `network_interface.*.network_id` argument
        /// instead.
        /// </summary>
        [Output("privateNetworkIds")]
        public Output<ImmutableArray<string>> PrivateNetworkIds { get; private set; } = null!;

        /// <summary>
        /// The instance (main network interface) IPv4 address.
        /// </summary>
        [Output("publicIpAddress")]
        public Output<string> PublicIpAddress { get; private set; } = null!;

        /// <summary>
        /// Domain name for reverse DNS record.
        /// </summary>
        [Output("reverseDns")]
        public Output<string?> ReverseDns { get; private set; } = null!;

        /// <summary>
        /// A list of [exoscale_security_group](./security_group.md) (IDs) to attach to the instance.
        /// </summary>
        [Output("securityGroupIds")]
        public Output<ImmutableArray<string>> SecurityGroupIds { get; private set; } = null!;

        /// <summary>
        /// The [exoscale_ssh_key](./ssh_key.md) (name) to authorize in the instance (may only be set at creation time).
        /// </summary>
        [Output("sshKey")]
        public Output<string?> SshKey { get; private set; } = null!;

        /// <summary>
        /// The instance state (`running` or `stopped`; default: `running`).
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The [exoscale_compute_template](../data-sources/compute_template.md) (ID) to use when creating the instance.
        /// </summary>
        [Output("templateId")]
        public Output<string> TemplateId { get; private set; } = null!;

        /// <summary>
        /// The instance type (`&lt;family&gt;.&lt;size&gt;`, e.g. `standard.medium`; use the [Exoscale CLI](https://github.com/exoscale/cli/) -
        /// `exo compute instance-type list` - for the list of available types). **WARNING**: updating this attribute stops/restarts
        /// the instance.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// [cloud-init](https://cloudinit.readthedocs.io/) configuration (no need to base64-encode or gzip it as the provider will
        /// take care of it).
        /// </summary>
        [Output("userData")]
        public Output<string?> UserData { get; private set; } = null!;

        /// <summary>
        /// The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a ComputeInstance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ComputeInstance(string name, ComputeInstanceArgs args, CustomResourceOptions? options = null)
            : base("exoscale:index/computeInstance:ComputeInstance", name, args ?? new ComputeInstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ComputeInstance(string name, Input<string> id, ComputeInstanceState? state = null, CustomResourceOptions? options = null)
            : base("exoscale:index/computeInstance:ComputeInstance", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse/pulumi-exoscale",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ComputeInstance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ComputeInstance Get(string name, Input<string> id, ComputeInstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new ComputeInstance(name, id, state, options);
        }
    }

    public sealed class ComputeInstanceArgs : global::Pulumi.ResourceArgs
    {
        [Input("antiAffinityGroupIds")]
        private InputList<string>? _antiAffinityGroupIds;

        /// <summary>
        /// A list of [exoscale_anti_affinity_group](./anti_affinity_group.md) (IDs) to attach to the instance (may only be set at
        /// creation time).
        /// </summary>
        public InputList<string> AntiAffinityGroupIds
        {
            get => _antiAffinityGroupIds ?? (_antiAffinityGroupIds = new InputList<string>());
            set => _antiAffinityGroupIds = value;
        }

        /// <summary>
        /// A deploy target ID.
        /// </summary>
        [Input("deployTargetId")]
        public Input<string>? DeployTargetId { get; set; }

        /// <summary>
        /// The instance disk size (GiB; at least `10`). **WARNING**: updating this attribute stops/restarts the instance.
        /// </summary>
        [Input("diskSize")]
        public Input<int>? DiskSize { get; set; }

        [Input("elasticIpIds")]
        private InputList<string>? _elasticIpIds;

        /// <summary>
        /// A list of [exoscale_elastic_ip](./elastic_ip.md) (IDs) to attach to the instance.
        /// </summary>
        public InputList<string> ElasticIpIds
        {
            get => _elasticIpIds ?? (_elasticIpIds = new InputList<string>());
            set => _elasticIpIds = value;
        }

        /// <summary>
        /// Enable IPv6 on the instance (boolean; default: `false`).
        /// </summary>
        [Input("ipv6")]
        public Input<bool>? Ipv6 { get; set; }

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
        /// The compute instance name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("networkInterfaces")]
        private InputList<Inputs.ComputeInstanceNetworkInterfaceArgs>? _networkInterfaces;

        /// <summary>
        /// Private network interfaces (may be specified multiple times). Structure is documented below.
        /// </summary>
        public InputList<Inputs.ComputeInstanceNetworkInterfaceArgs> NetworkInterfaces
        {
            get => _networkInterfaces ?? (_networkInterfaces = new InputList<Inputs.ComputeInstanceNetworkInterfaceArgs>());
            set => _networkInterfaces = value;
        }

        /// <summary>
        /// Domain name for reverse DNS record.
        /// </summary>
        [Input("reverseDns")]
        public Input<string>? ReverseDns { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// A list of [exoscale_security_group](./security_group.md) (IDs) to attach to the instance.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// The [exoscale_ssh_key](./ssh_key.md) (name) to authorize in the instance (may only be set at creation time).
        /// </summary>
        [Input("sshKey")]
        public Input<string>? SshKey { get; set; }

        /// <summary>
        /// The instance state (`running` or `stopped`; default: `running`).
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// The [exoscale_compute_template](../data-sources/compute_template.md) (ID) to use when creating the instance.
        /// </summary>
        [Input("templateId", required: true)]
        public Input<string> TemplateId { get; set; } = null!;

        /// <summary>
        /// The instance type (`&lt;family&gt;.&lt;size&gt;`, e.g. `standard.medium`; use the [Exoscale CLI](https://github.com/exoscale/cli/) -
        /// `exo compute instance-type list` - for the list of available types). **WARNING**: updating this attribute stops/restarts
        /// the instance.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// [cloud-init](https://cloudinit.readthedocs.io/) configuration (no need to base64-encode or gzip it as the provider will
        /// take care of it).
        /// </summary>
        [Input("userData")]
        public Input<string>? UserData { get; set; }

        /// <summary>
        /// The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public ComputeInstanceArgs()
        {
        }
        public static new ComputeInstanceArgs Empty => new ComputeInstanceArgs();
    }

    public sealed class ComputeInstanceState : global::Pulumi.ResourceArgs
    {
        [Input("antiAffinityGroupIds")]
        private InputList<string>? _antiAffinityGroupIds;

        /// <summary>
        /// A list of [exoscale_anti_affinity_group](./anti_affinity_group.md) (IDs) to attach to the instance (may only be set at
        /// creation time).
        /// </summary>
        public InputList<string> AntiAffinityGroupIds
        {
            get => _antiAffinityGroupIds ?? (_antiAffinityGroupIds = new InputList<string>());
            set => _antiAffinityGroupIds = value;
        }

        /// <summary>
        /// The instance creation date.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// A deploy target ID.
        /// </summary>
        [Input("deployTargetId")]
        public Input<string>? DeployTargetId { get; set; }

        /// <summary>
        /// The instance disk size (GiB; at least `10`). **WARNING**: updating this attribute stops/restarts the instance.
        /// </summary>
        [Input("diskSize")]
        public Input<int>? DiskSize { get; set; }

        [Input("elasticIpIds")]
        private InputList<string>? _elasticIpIds;

        /// <summary>
        /// A list of [exoscale_elastic_ip](./elastic_ip.md) (IDs) to attach to the instance.
        /// </summary>
        public InputList<string> ElasticIpIds
        {
            get => _elasticIpIds ?? (_elasticIpIds = new InputList<string>());
            set => _elasticIpIds = value;
        }

        /// <summary>
        /// Enable IPv6 on the instance (boolean; default: `false`).
        /// </summary>
        [Input("ipv6")]
        public Input<bool>? Ipv6 { get; set; }

        /// <summary>
        /// The instance (main network interface) IPv6 address (if enabled).
        /// </summary>
        [Input("ipv6Address")]
        public Input<string>? Ipv6Address { get; set; }

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
        /// The compute instance name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("networkInterfaces")]
        private InputList<Inputs.ComputeInstanceNetworkInterfaceGetArgs>? _networkInterfaces;

        /// <summary>
        /// Private network interfaces (may be specified multiple times). Structure is documented below.
        /// </summary>
        public InputList<Inputs.ComputeInstanceNetworkInterfaceGetArgs> NetworkInterfaces
        {
            get => _networkInterfaces ?? (_networkInterfaces = new InputList<Inputs.ComputeInstanceNetworkInterfaceGetArgs>());
            set => _networkInterfaces = value;
        }

        [Input("privateNetworkIds")]
        private InputList<string>? _privateNetworkIds;

        /// <summary>
        /// A list of private networks (IDs) attached to the instance. Please use the `network_interface.*.network_id` argument
        /// instead.
        /// </summary>
        [Obsolete(@"Use the network_interface block instead.")]
        public InputList<string> PrivateNetworkIds
        {
            get => _privateNetworkIds ?? (_privateNetworkIds = new InputList<string>());
            set => _privateNetworkIds = value;
        }

        /// <summary>
        /// The instance (main network interface) IPv4 address.
        /// </summary>
        [Input("publicIpAddress")]
        public Input<string>? PublicIpAddress { get; set; }

        /// <summary>
        /// Domain name for reverse DNS record.
        /// </summary>
        [Input("reverseDns")]
        public Input<string>? ReverseDns { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// A list of [exoscale_security_group](./security_group.md) (IDs) to attach to the instance.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// The [exoscale_ssh_key](./ssh_key.md) (name) to authorize in the instance (may only be set at creation time).
        /// </summary>
        [Input("sshKey")]
        public Input<string>? SshKey { get; set; }

        /// <summary>
        /// The instance state (`running` or `stopped`; default: `running`).
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// The [exoscale_compute_template](../data-sources/compute_template.md) (ID) to use when creating the instance.
        /// </summary>
        [Input("templateId")]
        public Input<string>? TemplateId { get; set; }

        /// <summary>
        /// The instance type (`&lt;family&gt;.&lt;size&gt;`, e.g. `standard.medium`; use the [Exoscale CLI](https://github.com/exoscale/cli/) -
        /// `exo compute instance-type list` - for the list of available types). **WARNING**: updating this attribute stops/restarts
        /// the instance.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// [cloud-init](https://cloudinit.readthedocs.io/) configuration (no need to base64-encode or gzip it as the provider will
        /// take care of it).
        /// </summary>
        [Input("userData")]
        public Input<string>? UserData { get; set; }

        /// <summary>
        /// The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public ComputeInstanceState()
        {
        }
        public static new ComputeInstanceState Empty => new ComputeInstanceState();
    }
}
