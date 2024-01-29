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
    public static class GetNlb
    {
        /// <summary>
        /// Fetch Exoscale [Network Load Balancers (NLB)](https://community.exoscale.com/documentation/compute/network-load-balancer/) data.
        /// 
        /// Corresponding resource: exoscale_nlb.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Exoscale = Pulumi.Exoscale;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myNlb = Exoscale.GetNlb.Invoke(new()
        ///     {
        ///         Zone = "ch-gva-2",
        ///         Name = "my-nlb",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["myNlbId"] = myNlb.Apply(getNlbResult =&gt; getNlbResult.Id),
        ///     };
        /// });
        /// ```
        /// 
        /// Please refer to the examples
        /// directory for complete configuration examples.
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetNlbResult> InvokeAsync(GetNlbArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNlbResult>("exoscale:index/getNlb:getNlb", args ?? new GetNlbArgs(), options.WithDefaults());

        /// <summary>
        /// Fetch Exoscale [Network Load Balancers (NLB)](https://community.exoscale.com/documentation/compute/network-load-balancer/) data.
        /// 
        /// Corresponding resource: exoscale_nlb.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Exoscale = Pulumi.Exoscale;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myNlb = Exoscale.GetNlb.Invoke(new()
        ///     {
        ///         Zone = "ch-gva-2",
        ///         Name = "my-nlb",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["myNlbId"] = myNlb.Apply(getNlbResult =&gt; getNlbResult.Id),
        ///     };
        /// });
        /// ```
        /// 
        /// Please refer to the examples
        /// directory for complete configuration examples.
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetNlbResult> Invoke(GetNlbInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNlbResult>("exoscale:index/getNlb:getNlb", args ?? new GetNlbInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNlbArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Network Load Balancers (NLB) ID to match (conflicts with `name`).
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// The NLB name to match (conflicts with `id`).
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        /// </summary>
        [Input("zone", required: true)]
        public string Zone { get; set; } = null!;

        public GetNlbArgs()
        {
        }
        public static new GetNlbArgs Empty => new GetNlbArgs();
    }

    public sealed class GetNlbInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Network Load Balancers (NLB) ID to match (conflicts with `name`).
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The NLB name to match (conflicts with `id`).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public GetNlbInvokeArgs()
        {
        }
        public static new GetNlbInvokeArgs Empty => new GetNlbInvokeArgs();
    }


    [OutputType]
    public sealed class GetNlbResult
    {
        /// <summary>
        /// The NLB creation date.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// The Network Load Balancers (NLB) description.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The Network Load Balancers (NLB) ID to match (conflicts with `name`).
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The NLB public IPv4 address.
        /// </summary>
        public readonly string IpAddress;
        /// <summary>
        /// The NLB name to match (conflicts with `id`).
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The current NLB state.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private GetNlbResult(
            string createdAt,

            string description,

            string? id,

            string ipAddress,

            string? name,

            string state,

            string zone)
        {
            CreatedAt = createdAt;
            Description = description;
            Id = id;
            IpAddress = ipAddress;
            Name = name;
            State = state;
            Zone = zone;
        }
    }
}