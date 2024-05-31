// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as splight from "@pulumi/splight";
 *
 * const fileInnerTest = new splight.File("fileInnerTest", {
 *     description: "Sample file for testing inner file",
 *     file: "./variables.tf",
 *     parent: "1234-1234-1234-1234",
 * });
 * ```
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
     * the path for the file resource in your system
     */
    public readonly file!: pulumi.Output<string>;
    /**
     * the id reference for a folder to be placed in
     */
    public readonly parent!: pulumi.Output<string | undefined>;
    /**
     * assets to be linked
     */
    public readonly relatedAssets!: pulumi.Output<string[] | undefined>;

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
            resourceInputs["file"] = state ? state.file : undefined;
            resourceInputs["parent"] = state ? state.parent : undefined;
            resourceInputs["relatedAssets"] = state ? state.relatedAssets : undefined;
        } else {
            const args = argsOrState as FileArgs | undefined;
            if ((!args || args.file === undefined) && !opts.urn) {
                throw new Error("Missing required property 'file'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["file"] = args ? args.file : undefined;
            resourceInputs["parent"] = args ? args.parent : undefined;
            resourceInputs["relatedAssets"] = args ? args.relatedAssets : undefined;
            resourceInputs["checksum"] = undefined /*out*/;
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
     * the path for the file resource in your system
     */
    file?: pulumi.Input<string>;
    /**
     * the id reference for a folder to be placed in
     */
    parent?: pulumi.Input<string>;
    /**
     * assets to be linked
     */
    relatedAssets?: pulumi.Input<pulumi.Input<string>[]>;
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
     * the path for the file resource in your system
     */
    file: pulumi.Input<string>;
    /**
     * the id reference for a folder to be placed in
     */
    parent?: pulumi.Input<string>;
    /**
     * assets to be linked
     */
    relatedAssets?: pulumi.Input<pulumi.Input<string>[]>;
}