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
 * $ pulumi import splight:index/componentRoutine:ComponentRoutine [options] splight_component_routine.<name> <component_routine_id>
 * ```
 */
export class ComponentRoutine extends pulumi.CustomResource {
    /**
     * Get an existing ComponentRoutine resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ComponentRoutineState, opts?: pulumi.CustomResourceOptions): ComponentRoutine {
        return new ComponentRoutine(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'splight:index/componentRoutine:ComponentRoutine';

    /**
     * Returns true if the given object is an instance of ComponentRoutine.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ComponentRoutine {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ComponentRoutine.__pulumiType;
    }

    /**
     * reference to component to be attached
     */
    public readonly componentId!: pulumi.Output<string>;
    /**
     * static config parameters of the routine
     */
    public readonly configs!: pulumi.Output<outputs.ComponentRoutineConfig[] | undefined>;
    /**
     * optional complementary information about the routine
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * asset attribute where to ingest data. Only valid for IncomingRoutine
     */
    public readonly inputs!: pulumi.Output<outputs.ComponentRoutineInput[] | undefined>;
    /**
     * name of the routine
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * asset attribute where to ingest data. Only valid for IncomingRoutine
     */
    public readonly outputs!: pulumi.Output<outputs.ComponentRoutineOutput[] | undefined>;
    /**
     * [IncomingRoutine|OutgoingRoutine] direction of the data flow (from device to system or from system to device)
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a ComponentRoutine resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ComponentRoutineArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ComponentRoutineArgs | ComponentRoutineState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ComponentRoutineState | undefined;
            resourceInputs["componentId"] = state ? state.componentId : undefined;
            resourceInputs["configs"] = state ? state.configs : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["inputs"] = state ? state.inputs : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["outputs"] = state ? state.outputs : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as ComponentRoutineArgs | undefined;
            if ((!args || args.componentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'componentId'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["componentId"] = args ? args.componentId : undefined;
            resourceInputs["configs"] = args ? args.configs : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["inputs"] = args ? args.inputs : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["outputs"] = args ? args.outputs : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ComponentRoutine.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ComponentRoutine resources.
 */
export interface ComponentRoutineState {
    /**
     * reference to component to be attached
     */
    componentId?: pulumi.Input<string>;
    /**
     * static config parameters of the routine
     */
    configs?: pulumi.Input<pulumi.Input<inputs.ComponentRoutineConfig>[]>;
    /**
     * optional complementary information about the routine
     */
    description?: pulumi.Input<string>;
    /**
     * asset attribute where to ingest data. Only valid for IncomingRoutine
     */
    inputs?: pulumi.Input<pulumi.Input<inputs.ComponentRoutineInput>[]>;
    /**
     * name of the routine
     */
    name?: pulumi.Input<string>;
    /**
     * asset attribute where to ingest data. Only valid for IncomingRoutine
     */
    outputs?: pulumi.Input<pulumi.Input<inputs.ComponentRoutineOutput>[]>;
    /**
     * [IncomingRoutine|OutgoingRoutine] direction of the data flow (from device to system or from system to device)
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ComponentRoutine resource.
 */
export interface ComponentRoutineArgs {
    /**
     * reference to component to be attached
     */
    componentId: pulumi.Input<string>;
    /**
     * static config parameters of the routine
     */
    configs?: pulumi.Input<pulumi.Input<inputs.ComponentRoutineConfig>[]>;
    /**
     * optional complementary information about the routine
     */
    description?: pulumi.Input<string>;
    /**
     * asset attribute where to ingest data. Only valid for IncomingRoutine
     */
    inputs?: pulumi.Input<pulumi.Input<inputs.ComponentRoutineInput>[]>;
    /**
     * name of the routine
     */
    name?: pulumi.Input<string>;
    /**
     * asset attribute where to ingest data. Only valid for IncomingRoutine
     */
    outputs?: pulumi.Input<pulumi.Input<inputs.ComponentRoutineOutput>[]>;
    /**
     * [IncomingRoutine|OutgoingRoutine] direction of the data flow (from device to system or from system to device)
     */
    type: pulumi.Input<string>;
}
