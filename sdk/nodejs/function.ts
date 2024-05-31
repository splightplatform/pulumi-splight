// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as splight from "@pulumi/splight";
 *
 * const functionTest = new splight.Function("functionTest", {
 *     description: "Created with Terraform",
 *     type: "rate",
 *     timeWindow: 600,
 *     rateValue: 10,
 *     rateUnit: "minute",
 *     targetVariable: "A",
 *     targetAsset: {
 *         id: "49551a15-d79b-40dc-9434-1b33d6b2fcb2",
 *         name: "An asset",
 *     },
 *     targetAttribute: {
 *         id: "49551a15-d79b-40dc-9434-1b33d6b2fcb2",
 *         name: "An attribute",
 *     },
 *     functionItems: [
 *         {
 *             refId: "A",
 *             type: "QUERY",
 *             expressionPlain: "",
 *             queryPlain: JSON.stringify([{
 *                 $match: {
 *                     asset: "49551a15-d79b-40dc-9434-1b33d6b2fcb2",
 *                     attribute: "c1d0d94b-5feb-4ebb-a527-0b0a34196252",
 *                 },
 *             }]),
 *         },
 *         {
 *             refId: "B",
 *             type: "QUERY",
 *             expressionPlain: "",
 *             queryPlain: JSON.stringify([{
 *                 $match: {
 *                     asset: "49551a15-d79b-40dc-9434-1b33d6b2fcb2",
 *                     attribute: "c1d0d94b-5feb-4ebb-a527-0b0a34196252",
 *                 },
 *             }]),
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import splight:index/function:Function [options] splight_function.<name> <function_id>
 * ```
 */
export class Function extends pulumi.CustomResource {
    /**
     * Get an existing Function resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FunctionState, opts?: pulumi.CustomResourceOptions): Function {
        return new Function(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'splight:index/function:Function';

    /**
     * Returns true if the given object is an instance of Function.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Function {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Function.__pulumiType;
    }

    /**
     * schedule value for cron
     */
    public readonly cronDom!: pulumi.Output<number>;
    /**
     * schedule value for cron
     */
    public readonly cronDow!: pulumi.Output<number>;
    /**
     * schedule value for cron
     */
    public readonly cronHours!: pulumi.Output<number>;
    /**
     * schedule value for cron
     */
    public readonly cronMinutes!: pulumi.Output<number>;
    /**
     * schedule value for cron
     */
    public readonly cronMonth!: pulumi.Output<number>;
    /**
     * schedule value for cron
     */
    public readonly cronYear!: pulumi.Output<number>;
    /**
     * The description of the resource
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * traces to be used to compute the results
     */
    public readonly functionItems!: pulumi.Output<outputs.FunctionFunctionItem[]>;
    /**
     * The name of the resource
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * [day|hour|minute] schedule unit
     */
    public readonly rateUnit!: pulumi.Output<string>;
    /**
     * schedule value
     */
    public readonly rateValue!: pulumi.Output<number>;
    /**
     * asset where to ingest results
     */
    public readonly targetAsset!: pulumi.Output<{[key: string]: string}>;
    /**
     * attribute where to ingest results
     */
    public readonly targetAttribute!: pulumi.Output<{[key: string]: string}>;
    /**
     * variable to be considered to be ingested
     */
    public readonly targetVariable!: pulumi.Output<string>;
    /**
     * window to fetch data from. Data out of that window will not be considered for evaluation
     */
    public readonly timeWindow!: pulumi.Output<number>;
    /**
     * [cron|rate] type for the cron
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a Function resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FunctionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FunctionArgs | FunctionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FunctionState | undefined;
            resourceInputs["cronDom"] = state ? state.cronDom : undefined;
            resourceInputs["cronDow"] = state ? state.cronDow : undefined;
            resourceInputs["cronHours"] = state ? state.cronHours : undefined;
            resourceInputs["cronMinutes"] = state ? state.cronMinutes : undefined;
            resourceInputs["cronMonth"] = state ? state.cronMonth : undefined;
            resourceInputs["cronYear"] = state ? state.cronYear : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["functionItems"] = state ? state.functionItems : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["rateUnit"] = state ? state.rateUnit : undefined;
            resourceInputs["rateValue"] = state ? state.rateValue : undefined;
            resourceInputs["targetAsset"] = state ? state.targetAsset : undefined;
            resourceInputs["targetAttribute"] = state ? state.targetAttribute : undefined;
            resourceInputs["targetVariable"] = state ? state.targetVariable : undefined;
            resourceInputs["timeWindow"] = state ? state.timeWindow : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as FunctionArgs | undefined;
            if ((!args || args.description === undefined) && !opts.urn) {
                throw new Error("Missing required property 'description'");
            }
            if ((!args || args.functionItems === undefined) && !opts.urn) {
                throw new Error("Missing required property 'functionItems'");
            }
            if ((!args || args.targetAsset === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetAsset'");
            }
            if ((!args || args.targetAttribute === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetAttribute'");
            }
            if ((!args || args.targetVariable === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetVariable'");
            }
            if ((!args || args.timeWindow === undefined) && !opts.urn) {
                throw new Error("Missing required property 'timeWindow'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["cronDom"] = args ? args.cronDom : undefined;
            resourceInputs["cronDow"] = args ? args.cronDow : undefined;
            resourceInputs["cronHours"] = args ? args.cronHours : undefined;
            resourceInputs["cronMinutes"] = args ? args.cronMinutes : undefined;
            resourceInputs["cronMonth"] = args ? args.cronMonth : undefined;
            resourceInputs["cronYear"] = args ? args.cronYear : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["functionItems"] = args ? args.functionItems : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["rateUnit"] = args ? args.rateUnit : undefined;
            resourceInputs["rateValue"] = args ? args.rateValue : undefined;
            resourceInputs["targetAsset"] = args ? args.targetAsset : undefined;
            resourceInputs["targetAttribute"] = args ? args.targetAttribute : undefined;
            resourceInputs["targetVariable"] = args ? args.targetVariable : undefined;
            resourceInputs["timeWindow"] = args ? args.timeWindow : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Function.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Function resources.
 */
export interface FunctionState {
    /**
     * schedule value for cron
     */
    cronDom?: pulumi.Input<number>;
    /**
     * schedule value for cron
     */
    cronDow?: pulumi.Input<number>;
    /**
     * schedule value for cron
     */
    cronHours?: pulumi.Input<number>;
    /**
     * schedule value for cron
     */
    cronMinutes?: pulumi.Input<number>;
    /**
     * schedule value for cron
     */
    cronMonth?: pulumi.Input<number>;
    /**
     * schedule value for cron
     */
    cronYear?: pulumi.Input<number>;
    /**
     * The description of the resource
     */
    description?: pulumi.Input<string>;
    /**
     * traces to be used to compute the results
     */
    functionItems?: pulumi.Input<pulumi.Input<inputs.FunctionFunctionItem>[]>;
    /**
     * The name of the resource
     */
    name?: pulumi.Input<string>;
    /**
     * [day|hour|minute] schedule unit
     */
    rateUnit?: pulumi.Input<string>;
    /**
     * schedule value
     */
    rateValue?: pulumi.Input<number>;
    /**
     * asset where to ingest results
     */
    targetAsset?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * attribute where to ingest results
     */
    targetAttribute?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * variable to be considered to be ingested
     */
    targetVariable?: pulumi.Input<string>;
    /**
     * window to fetch data from. Data out of that window will not be considered for evaluation
     */
    timeWindow?: pulumi.Input<number>;
    /**
     * [cron|rate] type for the cron
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Function resource.
 */
export interface FunctionArgs {
    /**
     * schedule value for cron
     */
    cronDom?: pulumi.Input<number>;
    /**
     * schedule value for cron
     */
    cronDow?: pulumi.Input<number>;
    /**
     * schedule value for cron
     */
    cronHours?: pulumi.Input<number>;
    /**
     * schedule value for cron
     */
    cronMinutes?: pulumi.Input<number>;
    /**
     * schedule value for cron
     */
    cronMonth?: pulumi.Input<number>;
    /**
     * schedule value for cron
     */
    cronYear?: pulumi.Input<number>;
    /**
     * The description of the resource
     */
    description: pulumi.Input<string>;
    /**
     * traces to be used to compute the results
     */
    functionItems: pulumi.Input<pulumi.Input<inputs.FunctionFunctionItem>[]>;
    /**
     * The name of the resource
     */
    name?: pulumi.Input<string>;
    /**
     * [day|hour|minute] schedule unit
     */
    rateUnit?: pulumi.Input<string>;
    /**
     * schedule value
     */
    rateValue?: pulumi.Input<number>;
    /**
     * asset where to ingest results
     */
    targetAsset: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * attribute where to ingest results
     */
    targetAttribute: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * variable to be considered to be ingested
     */
    targetVariable: pulumi.Input<string>;
    /**
     * window to fetch data from. Data out of that window will not be considered for evaluation
     */
    timeWindow: pulumi.Input<number>;
    /**
     * [cron|rate] type for the cron
     */
    type: pulumi.Input<string>;
}
