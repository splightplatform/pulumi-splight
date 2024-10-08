// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import splight:index/alert:Alert [options] splight_alert.<name> <alert_id>
 * ```
 */
export class Alert extends pulumi.CustomResource {
    /**
     * Get an existing Alert resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AlertState, opts?: pulumi.CustomResourceOptions): Alert {
        return new Alert(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'splight:index/alert:Alert';

    /**
     * Returns true if the given object is an instance of Alert.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Alert {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Alert.__pulumiType;
    }

    /**
     * aggregation to be applied to reads before comparisson
     */
    public readonly aggregation!: pulumi.Output<string>;
    /**
     * traces to be used to compute the results
     */
    public readonly alertItems!: pulumi.Output<outputs.AlertAlertItem[]>;
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
     * The name of the resource
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * operator to be used to compare the read value with the threshold value
     */
    public readonly operator!: pulumi.Output<string>;
    /**
     * [day|hour|minute] schedule unit
     */
    public readonly rateUnit!: pulumi.Output<string>;
    /**
     * schedule value
     */
    public readonly rateValue!: pulumi.Output<number>;
    /**
     * related assets of the resource
     */
    public readonly relatedAssets!: pulumi.Output<outputs.AlertRelatedAsset[] | undefined>;
    /**
     * [sev1,...,sev8] severity for the alert
     */
    public readonly severity!: pulumi.Output<string>;
    /**
     * tags of the resource
     */
    public readonly tags!: pulumi.Output<outputs.AlertTag[] | undefined>;
    /**
     * variable to be used to compare with thresholds
     */
    public readonly targetVariable!: pulumi.Output<string>;
    public readonly thresholds!: pulumi.Output<outputs.AlertThreshold[]>;
    /**
     * window to fetch data from. Data out of that window will not be considered for evaluation
     */
    public readonly timeWindow!: pulumi.Output<number>;
    /**
     * [cron|rate] type for the cron
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a Alert resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AlertArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AlertArgs | AlertState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AlertState | undefined;
            resourceInputs["aggregation"] = state ? state.aggregation : undefined;
            resourceInputs["alertItems"] = state ? state.alertItems : undefined;
            resourceInputs["cronDom"] = state ? state.cronDom : undefined;
            resourceInputs["cronDow"] = state ? state.cronDow : undefined;
            resourceInputs["cronHours"] = state ? state.cronHours : undefined;
            resourceInputs["cronMinutes"] = state ? state.cronMinutes : undefined;
            resourceInputs["cronMonth"] = state ? state.cronMonth : undefined;
            resourceInputs["cronYear"] = state ? state.cronYear : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["operator"] = state ? state.operator : undefined;
            resourceInputs["rateUnit"] = state ? state.rateUnit : undefined;
            resourceInputs["rateValue"] = state ? state.rateValue : undefined;
            resourceInputs["relatedAssets"] = state ? state.relatedAssets : undefined;
            resourceInputs["severity"] = state ? state.severity : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["targetVariable"] = state ? state.targetVariable : undefined;
            resourceInputs["thresholds"] = state ? state.thresholds : undefined;
            resourceInputs["timeWindow"] = state ? state.timeWindow : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as AlertArgs | undefined;
            if ((!args || args.aggregation === undefined) && !opts.urn) {
                throw new Error("Missing required property 'aggregation'");
            }
            if ((!args || args.alertItems === undefined) && !opts.urn) {
                throw new Error("Missing required property 'alertItems'");
            }
            if ((!args || args.description === undefined) && !opts.urn) {
                throw new Error("Missing required property 'description'");
            }
            if ((!args || args.operator === undefined) && !opts.urn) {
                throw new Error("Missing required property 'operator'");
            }
            if ((!args || args.severity === undefined) && !opts.urn) {
                throw new Error("Missing required property 'severity'");
            }
            if ((!args || args.targetVariable === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetVariable'");
            }
            if ((!args || args.thresholds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'thresholds'");
            }
            if ((!args || args.timeWindow === undefined) && !opts.urn) {
                throw new Error("Missing required property 'timeWindow'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["aggregation"] = args ? args.aggregation : undefined;
            resourceInputs["alertItems"] = args ? args.alertItems : undefined;
            resourceInputs["cronDom"] = args ? args.cronDom : undefined;
            resourceInputs["cronDow"] = args ? args.cronDow : undefined;
            resourceInputs["cronHours"] = args ? args.cronHours : undefined;
            resourceInputs["cronMinutes"] = args ? args.cronMinutes : undefined;
            resourceInputs["cronMonth"] = args ? args.cronMonth : undefined;
            resourceInputs["cronYear"] = args ? args.cronYear : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["operator"] = args ? args.operator : undefined;
            resourceInputs["rateUnit"] = args ? args.rateUnit : undefined;
            resourceInputs["rateValue"] = args ? args.rateValue : undefined;
            resourceInputs["relatedAssets"] = args ? args.relatedAssets : undefined;
            resourceInputs["severity"] = args ? args.severity : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["targetVariable"] = args ? args.targetVariable : undefined;
            resourceInputs["thresholds"] = args ? args.thresholds : undefined;
            resourceInputs["timeWindow"] = args ? args.timeWindow : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Alert.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Alert resources.
 */
export interface AlertState {
    /**
     * aggregation to be applied to reads before comparisson
     */
    aggregation?: pulumi.Input<string>;
    /**
     * traces to be used to compute the results
     */
    alertItems?: pulumi.Input<pulumi.Input<inputs.AlertAlertItem>[]>;
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
     * The name of the resource
     */
    name?: pulumi.Input<string>;
    /**
     * operator to be used to compare the read value with the threshold value
     */
    operator?: pulumi.Input<string>;
    /**
     * [day|hour|minute] schedule unit
     */
    rateUnit?: pulumi.Input<string>;
    /**
     * schedule value
     */
    rateValue?: pulumi.Input<number>;
    /**
     * related assets of the resource
     */
    relatedAssets?: pulumi.Input<pulumi.Input<inputs.AlertRelatedAsset>[]>;
    /**
     * [sev1,...,sev8] severity for the alert
     */
    severity?: pulumi.Input<string>;
    /**
     * tags of the resource
     */
    tags?: pulumi.Input<pulumi.Input<inputs.AlertTag>[]>;
    /**
     * variable to be used to compare with thresholds
     */
    targetVariable?: pulumi.Input<string>;
    thresholds?: pulumi.Input<pulumi.Input<inputs.AlertThreshold>[]>;
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
 * The set of arguments for constructing a Alert resource.
 */
export interface AlertArgs {
    /**
     * aggregation to be applied to reads before comparisson
     */
    aggregation: pulumi.Input<string>;
    /**
     * traces to be used to compute the results
     */
    alertItems: pulumi.Input<pulumi.Input<inputs.AlertAlertItem>[]>;
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
     * The name of the resource
     */
    name?: pulumi.Input<string>;
    /**
     * operator to be used to compare the read value with the threshold value
     */
    operator: pulumi.Input<string>;
    /**
     * [day|hour|minute] schedule unit
     */
    rateUnit?: pulumi.Input<string>;
    /**
     * schedule value
     */
    rateValue?: pulumi.Input<number>;
    /**
     * related assets of the resource
     */
    relatedAssets?: pulumi.Input<pulumi.Input<inputs.AlertRelatedAsset>[]>;
    /**
     * [sev1,...,sev8] severity for the alert
     */
    severity: pulumi.Input<string>;
    /**
     * tags of the resource
     */
    tags?: pulumi.Input<pulumi.Input<inputs.AlertTag>[]>;
    /**
     * variable to be used to compare with thresholds
     */
    targetVariable: pulumi.Input<string>;
    thresholds: pulumi.Input<pulumi.Input<inputs.AlertThreshold>[]>;
    /**
     * window to fetch data from. Data out of that window will not be considered for evaluation
     */
    timeWindow: pulumi.Input<number>;
    /**
     * [cron|rate] type for the cron
     */
    type: pulumi.Input<string>;
}
