// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Exoscale.Inputs
{

    public sealed class InstancePoolInstanceGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("ipv6Address")]
        public Input<string>? Ipv6Address { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("publicIpAddress")]
        public Input<string>? PublicIpAddress { get; set; }

        public InstancePoolInstanceGetArgs()
        {
        }
        public static new InstancePoolInstanceGetArgs Empty => new InstancePoolInstanceGetArgs();
    }
}
