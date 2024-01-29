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
    public sealed class IamRolePolicyServices
    {
        /// <summary>
        /// List of IAM service rules (if type is `rules`).
        /// </summary>
        public readonly ImmutableArray<Outputs.IamRolePolicyServicesRule> Rules;
        /// <summary>
        /// Service type (`rules`, `allow`, or `deny`).
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private IamRolePolicyServices(
            ImmutableArray<Outputs.IamRolePolicyServicesRule> rules,

            string? type)
        {
            Rules = rules;
            Type = type;
        }
    }
}