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
 * $ pulumi import splight:index/slackLine:SlackLine [options] splight_slack_line.<name> <slack_line_id>
 * ```
 */
export class SlackLine extends pulumi.CustomResource {
    /**
     * Get an existing SlackLine resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SlackLineState, opts?: pulumi.CustomResourceOptions): SlackLine {
        return new SlackLine(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'splight:index/slackLine:SlackLine';

    /**
     * Returns true if the given object is an instance of SlackLine.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SlackLine {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SlackLine.__pulumiType;
    }

    /**
     * description of the resource
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * geo position and shape of the resource
     */
    public readonly geometry!: pulumi.Output<string | undefined>;
    /**
     * kind of the resource
     */
    public /*out*/ readonly kinds!: pulumi.Output<outputs.SlackLineKind[]>;
    /**
     * name of the resource
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * attribute of the resource
     */
    public /*out*/ readonly switchStatusEnds!: pulumi.Output<outputs.SlackLineSwitchStatusEnd[]>;
    /**
     * attribute of the resource
     */
    public /*out*/ readonly switchStatusStarts!: pulumi.Output<outputs.SlackLineSwitchStatusStart[]>;
    /**
     * tags of the resource
     */
    public readonly tags!: pulumi.Output<outputs.SlackLineTag[] | undefined>;
    /**
     * timezone that overrides location-based timezone of the resource
     */
    public readonly timezone!: pulumi.Output<string>;

    /**
     * Create a SlackLine resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SlackLineArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SlackLineArgs | SlackLineState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SlackLineState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["geometry"] = state ? state.geometry : undefined;
            resourceInputs["kinds"] = state ? state.kinds : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["switchStatusEnds"] = state ? state.switchStatusEnds : undefined;
            resourceInputs["switchStatusStarts"] = state ? state.switchStatusStarts : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["timezone"] = state ? state.timezone : undefined;
        } else {
            const args = argsOrState as SlackLineArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["geometry"] = args ? args.geometry : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["timezone"] = args ? args.timezone : undefined;
            resourceInputs["kinds"] = undefined /*out*/;
            resourceInputs["switchStatusEnds"] = undefined /*out*/;
            resourceInputs["switchStatusStarts"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SlackLine.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SlackLine resources.
 */
export interface SlackLineState {
    /**
     * description of the resource
     */
    description?: pulumi.Input<string>;
    /**
     * geo position and shape of the resource
     */
    geometry?: pulumi.Input<string>;
    /**
     * kind of the resource
     */
    kinds?: pulumi.Input<pulumi.Input<inputs.SlackLineKind>[]>;
    /**
     * name of the resource
     */
    name?: pulumi.Input<string>;
    /**
     * attribute of the resource
     */
    switchStatusEnds?: pulumi.Input<pulumi.Input<inputs.SlackLineSwitchStatusEnd>[]>;
    /**
     * attribute of the resource
     */
    switchStatusStarts?: pulumi.Input<pulumi.Input<inputs.SlackLineSwitchStatusStart>[]>;
    /**
     * tags of the resource
     */
    tags?: pulumi.Input<pulumi.Input<inputs.SlackLineTag>[]>;
    /**
     * timezone that overrides location-based timezone of the resource
     */
    timezone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SlackLine resource.
 */
export interface SlackLineArgs {
    /**
     * description of the resource
     */
    description?: pulumi.Input<string>;
    /**
     * geo position and shape of the resource
     */
    geometry?: pulumi.Input<string>;
    /**
     * name of the resource
     */
    name?: pulumi.Input<string>;
    /**
     * tags of the resource
     */
    tags?: pulumi.Input<pulumi.Input<inputs.SlackLineTag>[]>;
    /**
     * timezone that overrides location-based timezone of the resource
     */
    timezone?: pulumi.Input<string>;
}
