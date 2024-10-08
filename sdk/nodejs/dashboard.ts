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
 * $ pulumi import splight:index/dashboard:Dashboard [options] splight_dashboard.<name> <dashboard_id>
 * ```
 */
export class Dashboard extends pulumi.CustomResource {
    /**
     * Get an existing Dashboard resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DashboardState, opts?: pulumi.CustomResourceOptions): Dashboard {
        return new Dashboard(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'splight:index/dashboard:Dashboard';

    /**
     * Returns true if the given object is an instance of Dashboard.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Dashboard {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Dashboard.__pulumiType;
    }

    /**
     * dashboard description
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * dashboard name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * related assets of the resource
     */
    public readonly relatedAssets!: pulumi.Output<outputs.DashboardRelatedAsset[] | undefined>;
    /**
     * tags of the resource
     */
    public readonly tags!: pulumi.Output<outputs.DashboardTag[] | undefined>;

    /**
     * Create a Dashboard resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DashboardArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DashboardArgs | DashboardState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DashboardState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["relatedAssets"] = state ? state.relatedAssets : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as DashboardArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["relatedAssets"] = args ? args.relatedAssets : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Dashboard.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Dashboard resources.
 */
export interface DashboardState {
    /**
     * dashboard description
     */
    description?: pulumi.Input<string>;
    /**
     * dashboard name
     */
    name?: pulumi.Input<string>;
    /**
     * related assets of the resource
     */
    relatedAssets?: pulumi.Input<pulumi.Input<inputs.DashboardRelatedAsset>[]>;
    /**
     * tags of the resource
     */
    tags?: pulumi.Input<pulumi.Input<inputs.DashboardTag>[]>;
}

/**
 * The set of arguments for constructing a Dashboard resource.
 */
export interface DashboardArgs {
    /**
     * dashboard description
     */
    description?: pulumi.Input<string>;
    /**
     * dashboard name
     */
    name?: pulumi.Input<string>;
    /**
     * related assets of the resource
     */
    relatedAssets?: pulumi.Input<pulumi.Input<inputs.DashboardRelatedAsset>[]>;
    /**
     * tags of the resource
     */
    tags?: pulumi.Input<pulumi.Input<inputs.DashboardTag>[]>;
}
