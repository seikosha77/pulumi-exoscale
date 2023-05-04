// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * !> **WARNING:** This data source is **DEPRECATED** and will be removed in the next major version. Please use exoscale.ComputeInstance instead.
 */
export function getCompute(args?: GetComputeArgs, opts?: pulumi.InvokeOptions): Promise<GetComputeResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("exoscale:index/getCompute:getCompute", {
        "hostname": args.hostname,
        "id": args.id,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getCompute.
 */
export interface GetComputeArgs {
    /**
     * The instance hostname to match.
     */
    hostname?: string;
    /**
     * The compute instance ID to match.
     */
    id?: string;
    /**
     * The instance tags to match (map of key/value).
     */
    tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getCompute.
 */
export interface GetComputeResult {
    /**
     * Number of cpu the Compute instance is running with
     */
    readonly cpu: number;
    /**
     * Date when the Compute instance was created
     */
    readonly created: string;
    /**
     * Size of the Compute instance disk
     */
    readonly diskSize: number;
    /**
     * The instance hostname to match.
     */
    readonly hostname?: string;
    /**
     * The compute instance ID to match.
     */
    readonly id?: string;
    /**
     * Compute instance public ipv6 address (if ipv6 is enabled)
     */
    readonly ip6Address: string;
    /**
     * Compute instance public ipv4 address
     */
    readonly ipAddress: string;
    /**
     * Memory allocated for the Compute instance
     */
    readonly memory: number;
    /**
     * List of Compute instance private IP addresses (in managed Private Networks only)
     */
    readonly privateNetworkIpAddresses: string[];
    /**
     * Current size of the Compute instance
     */
    readonly size: string;
    /**
     * State of the Compute instance
     */
    readonly state: string;
    /**
     * The instance tags to match (map of key/value).
     */
    readonly tags?: {[key: string]: string};
    /**
     * Name of the template for the Compute instance
     */
    readonly template: string;
    /**
     * Name of the availability zone for the Compute instance
     */
    readonly zone: string;
}
/**
 * !> **WARNING:** This data source is **DEPRECATED** and will be removed in the next major version. Please use exoscale.ComputeInstance instead.
 */
export function getComputeOutput(args?: GetComputeOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetComputeResult> {
    return pulumi.output(args).apply((a: any) => getCompute(a, opts))
}

/**
 * A collection of arguments for invoking getCompute.
 */
export interface GetComputeOutputArgs {
    /**
     * The instance hostname to match.
     */
    hostname?: pulumi.Input<string>;
    /**
     * The compute instance ID to match.
     */
    id?: pulumi.Input<string>;
    /**
     * The instance tags to match (map of key/value).
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
