// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as exoscale from "@pulumi/exoscale";
 * import * as exoscale from "@pulumiverse/exoscale";
 *
 * const myDatabaseDatabase = new exoscale.Database("myDatabaseDatabase", {
 *     zone: "ch-gva-2",
 *     type: "pg",
 *     plan: "startup-4",
 *     maintenanceDow: "sunday",
 *     maintenanceTime: "23:00:00",
 *     terminationProtection: true,
 *     pg: {
 *         version: "13",
 *         backupSchedule: "04:00",
 *         ipFilters: [
 *             "1.2.3.4/32",
 *             "5.6.7.8/32",
 *         ],
 *         pgSettings: JSON.stringify({
 *             timezone: "Europe/Zurich",
 *         }),
 *     },
 * });
 * const myDatabaseDatabaseURI = exoscale.getDatabaseURI({
 *     name: "my-database",
 *     type: "pg",
 * });
 * export const myDatabaseUri = myDatabaseDatabaseURI.then(myDatabaseDatabaseURI => myDatabaseDatabaseURI.uri);
 * ```
 */
export function getDatabaseURI(args: GetDatabaseURIArgs, opts?: pulumi.InvokeOptions): Promise<GetDatabaseURIResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("exoscale:index/getDatabaseURI:getDatabaseURI", {
        "name": args.name,
        "type": args.type,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getDatabaseURI.
 */
export interface GetDatabaseURIArgs {
    /**
     * The database name to match.
     */
    name: string;
    /**
     * The type of the database service (`kafka`, `mysql`, `opensearch`, `pg`, `redis`).
     */
    type: string;
    /**
     * (Required) The Exoscale Zone name.
     */
    zone: string;
}

/**
 * A collection of values returned by getDatabaseURI.
 */
export interface GetDatabaseURIResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The database name to match.
     */
    readonly name: string;
    /**
     * The type of the database service (`kafka`, `mysql`, `opensearch`, `pg`, `redis`).
     */
    readonly type: string;
    /**
     * The database service connection URI.
     */
    readonly uri: string;
    /**
     * (Required) The Exoscale Zone name.
     */
    readonly zone: string;
}
/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as exoscale from "@pulumi/exoscale";
 * import * as exoscale from "@pulumiverse/exoscale";
 *
 * const myDatabaseDatabase = new exoscale.Database("myDatabaseDatabase", {
 *     zone: "ch-gva-2",
 *     type: "pg",
 *     plan: "startup-4",
 *     maintenanceDow: "sunday",
 *     maintenanceTime: "23:00:00",
 *     terminationProtection: true,
 *     pg: {
 *         version: "13",
 *         backupSchedule: "04:00",
 *         ipFilters: [
 *             "1.2.3.4/32",
 *             "5.6.7.8/32",
 *         ],
 *         pgSettings: JSON.stringify({
 *             timezone: "Europe/Zurich",
 *         }),
 *     },
 * });
 * const myDatabaseDatabaseURI = exoscale.getDatabaseURI({
 *     name: "my-database",
 *     type: "pg",
 * });
 * export const myDatabaseUri = myDatabaseDatabaseURI.then(myDatabaseDatabaseURI => myDatabaseDatabaseURI.uri);
 * ```
 */
export function getDatabaseURIOutput(args: GetDatabaseURIOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDatabaseURIResult> {
    return pulumi.output(args).apply((a: any) => getDatabaseURI(a, opts))
}

/**
 * A collection of arguments for invoking getDatabaseURI.
 */
export interface GetDatabaseURIOutputArgs {
    /**
     * The database name to match.
     */
    name: pulumi.Input<string>;
    /**
     * The type of the database service (`kafka`, `mysql`, `opensearch`, `pg`, `redis`).
     */
    type: pulumi.Input<string>;
    /**
     * (Required) The Exoscale Zone name.
     */
    zone: pulumi.Input<string>;
}