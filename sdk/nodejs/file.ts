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
 * $ pulumi import splight:index/file:File [options] splight_file.<name> <file_id>
 * ```
 */
export class File extends pulumi.CustomResource {
    /**
     * Get an existing File resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FileState, opts?: pulumi.CustomResourceOptions): File {
        return new File(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'splight:index/file:File';

    /**
     * Returns true if the given object is an instance of File.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is File {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === File.__pulumiType;
    }

    public /*out*/ readonly checksum!: pulumi.Output<string>;
    /**
     * complementary information to describe the file
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * the id reference for a folder to be placed in
     */
    public readonly parent!: pulumi.Output<string | undefined>;
    /**
     * the path for the file resource in your system
     */
    public readonly path!: pulumi.Output<string>;
    /**
     * related assets of the resource
     */
    public readonly relatedAssets!: pulumi.Output<outputs.FileRelatedAsset[] | undefined>;
    /**
     * tags of the resource
     */
    public readonly tags!: pulumi.Output<outputs.FileTag[] | undefined>;
    public /*out*/ readonly uploaded!: pulumi.Output<boolean>;

    /**
     * Create a File resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FileArgs | FileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FileState | undefined;
            resourceInputs["checksum"] = state ? state.checksum : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["parent"] = state ? state.parent : undefined;
            resourceInputs["path"] = state ? state.path : undefined;
            resourceInputs["relatedAssets"] = state ? state.relatedAssets : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["uploaded"] = state ? state.uploaded : undefined;
        } else {
            const args = argsOrState as FileArgs | undefined;
            if ((!args || args.path === undefined) && !opts.urn) {
                throw new Error("Missing required property 'path'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["parent"] = args ? args.parent : undefined;
            resourceInputs["path"] = args ? args.path : undefined;
            resourceInputs["relatedAssets"] = args ? args.relatedAssets : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["checksum"] = undefined /*out*/;
            resourceInputs["uploaded"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(File.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering File resources.
 */
export interface FileState {
    checksum?: pulumi.Input<string>;
    /**
     * complementary information to describe the file
     */
    description?: pulumi.Input<string>;
    /**
     * the id reference for a folder to be placed in
     */
    parent?: pulumi.Input<string>;
    /**
     * the path for the file resource in your system
     */
    path?: pulumi.Input<string>;
    /**
     * related assets of the resource
     */
    relatedAssets?: pulumi.Input<pulumi.Input<inputs.FileRelatedAsset>[]>;
    /**
     * tags of the resource
     */
    tags?: pulumi.Input<pulumi.Input<inputs.FileTag>[]>;
    uploaded?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a File resource.
 */
export interface FileArgs {
    /**
     * complementary information to describe the file
     */
    description?: pulumi.Input<string>;
    /**
     * the id reference for a folder to be placed in
     */
    parent?: pulumi.Input<string>;
    /**
     * the path for the file resource in your system
     */
    path: pulumi.Input<string>;
    /**
     * related assets of the resource
     */
    relatedAssets?: pulumi.Input<pulumi.Input<inputs.FileRelatedAsset>[]>;
    /**
     * tags of the resource
     */
    tags?: pulumi.Input<pulumi.Input<inputs.FileTag>[]>;
}
