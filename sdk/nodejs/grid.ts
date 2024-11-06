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
 * $ pulumi import splight:index/grid:Grid [options] splight_grid.<name> <grid_id>
 * ```
 */
export class Grid extends pulumi.CustomResource {
    /**
     * Get an existing Grid resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GridState, opts?: pulumi.CustomResourceOptions): Grid {
        return new Grid(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'splight:index/grid:Grid';

    /**
     * Returns true if the given object is an instance of Grid.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Grid {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Grid.__pulumiType;
    }

    /**
     * timezone that overrides location-based timezone of the resource
     */
    public readonly customTimezone!: pulumi.Output<string | undefined>;
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
    public /*out*/ readonly kinds!: pulumi.Output<outputs.GridKind[]>;
    /**
     * name of the resource
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * tags of the resource
     */
    public readonly tags!: pulumi.Output<outputs.GridTag[] | undefined>;

    /**
     * Create a Grid resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: GridArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GridArgs | GridState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GridState | undefined;
            resourceInputs["customTimezone"] = state ? state.customTimezone : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["geometry"] = state ? state.geometry : undefined;
            resourceInputs["kinds"] = state ? state.kinds : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as GridArgs | undefined;
            resourceInputs["customTimezone"] = args ? args.customTimezone : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["geometry"] = args ? args.geometry : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["kinds"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Grid.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Grid resources.
 */
export interface GridState {
    /**
     * timezone that overrides location-based timezone of the resource
     */
    customTimezone?: pulumi.Input<string>;
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
    kinds?: pulumi.Input<pulumi.Input<inputs.GridKind>[]>;
    /**
     * name of the resource
     */
    name?: pulumi.Input<string>;
    /**
     * tags of the resource
     */
    tags?: pulumi.Input<pulumi.Input<inputs.GridTag>[]>;
}

/**
 * The set of arguments for constructing a Grid resource.
 */
export interface GridArgs {
    /**
     * timezone that overrides location-based timezone of the resource
     */
    customTimezone?: pulumi.Input<string>;
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
    tags?: pulumi.Input<pulumi.Input<inputs.GridTag>[]>;
}
