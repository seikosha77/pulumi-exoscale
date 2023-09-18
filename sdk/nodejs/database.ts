// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## Import
 *
 * An existing database service may be imported by `<name>@<zone>`
 *
 * ```sh
 *  $ pulumi import exoscale:index/database:Database \
 * ```
 *
 *  exoscale_database.my_database \
 *
 *  my-database@ch-gva-2
 */
export class Database extends pulumi.CustomResource {
    /**
     * Get an existing Database resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatabaseState, opts?: pulumi.CustomResourceOptions): Database {
        return new Database(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'exoscale:index/database:Database';

    /**
     * Returns true if the given object is an instance of Database.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Database {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Database.__pulumiType;
    }

    /**
     * CA Certificate required to reach a DBaaS service through a TLS-protected connection.
     */
    public /*out*/ readonly caCertificate!: pulumi.Output<string>;
    /**
     * The creation date of the database service.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The disk size of the database service.
     */
    public /*out*/ readonly diskSize!: pulumi.Output<number>;
    /**
     * *grafana* database service type specific arguments. Structure is documented below.
     */
    public readonly grafana!: pulumi.Output<outputs.DatabaseGrafana | undefined>;
    /**
     * *kafka* database service type specific arguments. Structure is documented below.
     */
    public readonly kafka!: pulumi.Output<outputs.DatabaseKafka | undefined>;
    /**
     * The day of week to perform the automated database service maintenance (`never`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `sunday`).
     */
    public readonly maintenanceDow!: pulumi.Output<string>;
    /**
     * The time of day to perform the automated database service maintenance (`HH:MM:SS`)
     */
    public readonly maintenanceTime!: pulumi.Output<string>;
    /**
     * *mysql* database service type specific arguments. Structure is documented below.
     */
    public readonly mysql!: pulumi.Output<outputs.DatabaseMysql | undefined>;
    /**
     * ❗ The name of the database service.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The number of CPUs of the database service.
     */
    public /*out*/ readonly nodeCpus!: pulumi.Output<number>;
    /**
     * The amount of memory of the database service.
     */
    public /*out*/ readonly nodeMemory!: pulumi.Output<number>;
    /**
     * The number of nodes of the database service.
     */
    public /*out*/ readonly nodes!: pulumi.Output<number>;
    /**
     * *opensearch* database service type specific arguments. Structure is documented below.
     */
    public readonly opensearch!: pulumi.Output<outputs.DatabaseOpensearch | undefined>;
    /**
     * *pg* database service type specific arguments. Structure is documented below.
     */
    public readonly pg!: pulumi.Output<outputs.DatabasePg | undefined>;
    /**
     * The plan of the database service (use the [Exoscale CLI](https://github.com/exoscale/cli/) - `exo dbaas type show <TYPE> --plans` - for reference).
     */
    public readonly plan!: pulumi.Output<string>;
    /**
     * *redis* database service type specific arguments. Structure is documented below.
     */
    public readonly redis!: pulumi.Output<outputs.DatabaseRedis | undefined>;
    /**
     * The current state of the database service.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The database service protection boolean flag against termination/power-off.
     */
    public readonly terminationProtection!: pulumi.Output<boolean>;
    public readonly timeouts!: pulumi.Output<outputs.DatabaseTimeouts | undefined>;
    /**
     * ❗ The type of the database service (`kafka`, `mysql`, `opensearch`, `pg`, `redis`, `grafana`).
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * The date of the latest database service update.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a Database resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatabaseArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatabaseArgs | DatabaseState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DatabaseState | undefined;
            resourceInputs["caCertificate"] = state ? state.caCertificate : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["diskSize"] = state ? state.diskSize : undefined;
            resourceInputs["grafana"] = state ? state.grafana : undefined;
            resourceInputs["kafka"] = state ? state.kafka : undefined;
            resourceInputs["maintenanceDow"] = state ? state.maintenanceDow : undefined;
            resourceInputs["maintenanceTime"] = state ? state.maintenanceTime : undefined;
            resourceInputs["mysql"] = state ? state.mysql : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nodeCpus"] = state ? state.nodeCpus : undefined;
            resourceInputs["nodeMemory"] = state ? state.nodeMemory : undefined;
            resourceInputs["nodes"] = state ? state.nodes : undefined;
            resourceInputs["opensearch"] = state ? state.opensearch : undefined;
            resourceInputs["pg"] = state ? state.pg : undefined;
            resourceInputs["plan"] = state ? state.plan : undefined;
            resourceInputs["redis"] = state ? state.redis : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["terminationProtection"] = state ? state.terminationProtection : undefined;
            resourceInputs["timeouts"] = state ? state.timeouts : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as DatabaseArgs | undefined;
            if ((!args || args.plan === undefined) && !opts.urn) {
                throw new Error("Missing required property 'plan'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            if ((!args || args.zone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zone'");
            }
            resourceInputs["grafana"] = args ? args.grafana : undefined;
            resourceInputs["kafka"] = args ? args.kafka : undefined;
            resourceInputs["maintenanceDow"] = args ? args.maintenanceDow : undefined;
            resourceInputs["maintenanceTime"] = args ? args.maintenanceTime : undefined;
            resourceInputs["mysql"] = args ? args.mysql : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["opensearch"] = args ? args.opensearch : undefined;
            resourceInputs["pg"] = args ? args.pg : undefined;
            resourceInputs["plan"] = args ? args.plan : undefined;
            resourceInputs["redis"] = args ? args.redis : undefined;
            resourceInputs["terminationProtection"] = args ? args.terminationProtection : undefined;
            resourceInputs["timeouts"] = args ? args.timeouts : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["caCertificate"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["diskSize"] = undefined /*out*/;
            resourceInputs["nodeCpus"] = undefined /*out*/;
            resourceInputs["nodeMemory"] = undefined /*out*/;
            resourceInputs["nodes"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Database.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Database resources.
 */
export interface DatabaseState {
    /**
     * CA Certificate required to reach a DBaaS service through a TLS-protected connection.
     */
    caCertificate?: pulumi.Input<string>;
    /**
     * The creation date of the database service.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The disk size of the database service.
     */
    diskSize?: pulumi.Input<number>;
    /**
     * *grafana* database service type specific arguments. Structure is documented below.
     */
    grafana?: pulumi.Input<inputs.DatabaseGrafana>;
    /**
     * *kafka* database service type specific arguments. Structure is documented below.
     */
    kafka?: pulumi.Input<inputs.DatabaseKafka>;
    /**
     * The day of week to perform the automated database service maintenance (`never`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `sunday`).
     */
    maintenanceDow?: pulumi.Input<string>;
    /**
     * The time of day to perform the automated database service maintenance (`HH:MM:SS`)
     */
    maintenanceTime?: pulumi.Input<string>;
    /**
     * *mysql* database service type specific arguments. Structure is documented below.
     */
    mysql?: pulumi.Input<inputs.DatabaseMysql>;
    /**
     * ❗ The name of the database service.
     */
    name?: pulumi.Input<string>;
    /**
     * The number of CPUs of the database service.
     */
    nodeCpus?: pulumi.Input<number>;
    /**
     * The amount of memory of the database service.
     */
    nodeMemory?: pulumi.Input<number>;
    /**
     * The number of nodes of the database service.
     */
    nodes?: pulumi.Input<number>;
    /**
     * *opensearch* database service type specific arguments. Structure is documented below.
     */
    opensearch?: pulumi.Input<inputs.DatabaseOpensearch>;
    /**
     * *pg* database service type specific arguments. Structure is documented below.
     */
    pg?: pulumi.Input<inputs.DatabasePg>;
    /**
     * The plan of the database service (use the [Exoscale CLI](https://github.com/exoscale/cli/) - `exo dbaas type show <TYPE> --plans` - for reference).
     */
    plan?: pulumi.Input<string>;
    /**
     * *redis* database service type specific arguments. Structure is documented below.
     */
    redis?: pulumi.Input<inputs.DatabaseRedis>;
    /**
     * The current state of the database service.
     */
    state?: pulumi.Input<string>;
    /**
     * The database service protection boolean flag against termination/power-off.
     */
    terminationProtection?: pulumi.Input<boolean>;
    timeouts?: pulumi.Input<inputs.DatabaseTimeouts>;
    /**
     * ❗ The type of the database service (`kafka`, `mysql`, `opensearch`, `pg`, `redis`, `grafana`).
     */
    type?: pulumi.Input<string>;
    /**
     * The date of the latest database service update.
     */
    updatedAt?: pulumi.Input<string>;
    /**
     * ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Database resource.
 */
export interface DatabaseArgs {
    /**
     * *grafana* database service type specific arguments. Structure is documented below.
     */
    grafana?: pulumi.Input<inputs.DatabaseGrafana>;
    /**
     * *kafka* database service type specific arguments. Structure is documented below.
     */
    kafka?: pulumi.Input<inputs.DatabaseKafka>;
    /**
     * The day of week to perform the automated database service maintenance (`never`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `sunday`).
     */
    maintenanceDow?: pulumi.Input<string>;
    /**
     * The time of day to perform the automated database service maintenance (`HH:MM:SS`)
     */
    maintenanceTime?: pulumi.Input<string>;
    /**
     * *mysql* database service type specific arguments. Structure is documented below.
     */
    mysql?: pulumi.Input<inputs.DatabaseMysql>;
    /**
     * ❗ The name of the database service.
     */
    name?: pulumi.Input<string>;
    /**
     * *opensearch* database service type specific arguments. Structure is documented below.
     */
    opensearch?: pulumi.Input<inputs.DatabaseOpensearch>;
    /**
     * *pg* database service type specific arguments. Structure is documented below.
     */
    pg?: pulumi.Input<inputs.DatabasePg>;
    /**
     * The plan of the database service (use the [Exoscale CLI](https://github.com/exoscale/cli/) - `exo dbaas type show <TYPE> --plans` - for reference).
     */
    plan: pulumi.Input<string>;
    /**
     * *redis* database service type specific arguments. Structure is documented below.
     */
    redis?: pulumi.Input<inputs.DatabaseRedis>;
    /**
     * The database service protection boolean flag against termination/power-off.
     */
    terminationProtection?: pulumi.Input<boolean>;
    timeouts?: pulumi.Input<inputs.DatabaseTimeouts>;
    /**
     * ❗ The type of the database service (`kafka`, `mysql`, `opensearch`, `pg`, `redis`, `grafana`).
     */
    type: pulumi.Input<string>;
    /**
     * ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    zone: pulumi.Input<string>;
}
