// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as splight from "@splightplatform/pulumi-splight";
 *
 * const myAsset = new splight.Asset("myAsset", {
 *     description: "My Asset Description",
 *     geometry: JSON.stringify({
 *         type: "GeometryCollection",
 *         geometries: [{
 *             type: "Point",
 *             coordinates: [
 *                 0,
 *                 0,
 *             ],
 *         }],
 *     }),
 * });
 * const myAttribute = new splight.AssetAttribute("myAttribute", {
 *     type: "Number",
 *     unit: "meters",
 *     asset: myAsset.id,
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import splight:index/assetAttribute:AssetAttribute [options] splight_asset_attribute.<name> <asset_attribute_id>
 * ```
 */
export class AssetAttribute extends pulumi.CustomResource {
    /**
     * Get an existing AssetAttribute resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AssetAttributeState, opts?: pulumi.CustomResourceOptions): AssetAttribute {
        return new AssetAttribute(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'splight:index/assetAttribute:AssetAttribute';

    /**
     * Returns true if the given object is an instance of AssetAttribute.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AssetAttribute {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AssetAttribute.__pulumiType;
    }

    /**
     * reference to the asset to be linked to
     */
    public readonly asset!: pulumi.Output<string>;
    /**
     * name of the resource
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * [String|Boolean|Number] type of the data to be ingested in this attribute
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * optional reference to the unit of the measure
     */
    public readonly unit!: pulumi.Output<string | undefined>;

    /**
     * Create a AssetAttribute resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AssetAttributeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AssetAttributeArgs | AssetAttributeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AssetAttributeState | undefined;
            resourceInputs["asset"] = state ? state.asset : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["unit"] = state ? state.unit : undefined;
        } else {
            const args = argsOrState as AssetAttributeArgs | undefined;
            if ((!args || args.asset === undefined) && !opts.urn) {
                throw new Error("Missing required property 'asset'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["asset"] = args ? args.asset : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["unit"] = args ? args.unit : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AssetAttribute.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AssetAttribute resources.
 */
export interface AssetAttributeState {
    /**
     * reference to the asset to be linked to
     */
    asset?: pulumi.Input<string>;
    /**
     * name of the resource
     */
    name?: pulumi.Input<string>;
    /**
     * [String|Boolean|Number] type of the data to be ingested in this attribute
     */
    type?: pulumi.Input<string>;
    /**
     * optional reference to the unit of the measure
     */
    unit?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AssetAttribute resource.
 */
export interface AssetAttributeArgs {
    /**
     * reference to the asset to be linked to
     */
    asset: pulumi.Input<string>;
    /**
     * name of the resource
     */
    name?: pulumi.Input<string>;
    /**
     * [String|Boolean|Number] type of the data to be ingested in this attribute
     */
    type: pulumi.Input<string>;
    /**
     * optional reference to the unit of the measure
     */
    unit?: pulumi.Input<string>;
}
