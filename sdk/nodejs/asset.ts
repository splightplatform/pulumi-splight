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
 * $ pulumi import splight:index/asset:Asset [options] splight_asset.<name> <asset_id>
 * ```
 */
export class Asset extends pulumi.CustomResource {
    /**
     * Get an existing Asset resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AssetState, opts?: pulumi.CustomResourceOptions): Asset {
        return new Asset(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'splight:index/asset:Asset';

    /**
     * Returns true if the given object is an instance of Asset.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Asset {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Asset.__pulumiType;
    }

    /**
     * description of the resource
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * GeoJSON GeomtryCollection
     */
    public readonly geometry!: pulumi.Output<string>;
    /**
     * kind of the resource
     */
    public readonly kind!: pulumi.Output<outputs.AssetKind | undefined>;
    /**
     * name of the resource
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * tags of the resource
     */
    public readonly tags!: pulumi.Output<outputs.AssetTag[] | undefined>;

    /**
     * Create a Asset resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AssetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AssetArgs | AssetState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AssetState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["geometry"] = state ? state.geometry : undefined;
            resourceInputs["kind"] = state ? state.kind : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as AssetArgs | undefined;
            if ((!args || args.geometry === undefined) && !opts.urn) {
                throw new Error("Missing required property 'geometry'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["geometry"] = args ? args.geometry : undefined;
            resourceInputs["kind"] = args ? args.kind : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Asset.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Asset resources.
 */
export interface AssetState {
    /**
     * description of the resource
     */
    description?: pulumi.Input<string>;
    /**
     * GeoJSON GeomtryCollection
     */
    geometry?: pulumi.Input<string>;
    /**
     * kind of the resource
     */
    kind?: pulumi.Input<inputs.AssetKind>;
    /**
     * name of the resource
     */
    name?: pulumi.Input<string>;
    /**
     * tags of the resource
     */
    tags?: pulumi.Input<pulumi.Input<inputs.AssetTag>[]>;
}

/**
 * The set of arguments for constructing a Asset resource.
 */
export interface AssetArgs {
    /**
     * description of the resource
     */
    description?: pulumi.Input<string>;
    /**
     * GeoJSON GeomtryCollection
     */
    geometry: pulumi.Input<string>;
    /**
     * kind of the resource
     */
    kind?: pulumi.Input<inputs.AssetKind>;
    /**
     * name of the resource
     */
    name?: pulumi.Input<string>;
    /**
     * tags of the resource
     */
    tags?: pulumi.Input<pulumi.Input<inputs.AssetTag>[]>;
}
