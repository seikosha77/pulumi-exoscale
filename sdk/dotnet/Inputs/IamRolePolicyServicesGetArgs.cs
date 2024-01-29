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

    public sealed class IamRolePolicyServicesGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("rules")]
        private InputList<Inputs.IamRolePolicyServicesRuleGetArgs>? _rules;

        /// <summary>
        /// List of IAM service rules (if type is `rules`).
        /// </summary>
        public InputList<Inputs.IamRolePolicyServicesRuleGetArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.IamRolePolicyServicesRuleGetArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// Service type (`rules`, `allow`, or `deny`).
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public IamRolePolicyServicesGetArgs()
        {
        }
        public static new IamRolePolicyServicesGetArgs Empty => new IamRolePolicyServicesGetArgs();
    }
}