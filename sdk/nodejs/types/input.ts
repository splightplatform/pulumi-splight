// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface AlertAlertItem {
    expressionPlain: pulumi.Input<string>;
    queryPlain: pulumi.Input<string>;
    refId: pulumi.Input<string>;
    type: pulumi.Input<string>;
}

export interface AlertThreshold {
    /**
     * [alert|warning|no_alert] status value for the threshold
     */
    status: pulumi.Input<string>;
    /**
     * optional custom value to be displayed in the platform.
     */
    statusText?: pulumi.Input<string>;
    /**
     * value to be considered to compare
     */
    value: pulumi.Input<number>;
}

export interface AssetKind {
    /**
     * kind id
     */
    id: pulumi.Input<string>;
    /**
     * kind name
     */
    name: pulumi.Input<string>;
}

export interface ComponentInput {
    description?: pulumi.Input<string>;
    multiple?: pulumi.Input<boolean>;
    name: pulumi.Input<string>;
    required?: pulumi.Input<boolean>;
    sensitive?: pulumi.Input<boolean>;
    type: pulumi.Input<string>;
    value?: pulumi.Input<string>;
}

export interface ComponentRoutineConfig {
    description?: pulumi.Input<string>;
    multiple?: pulumi.Input<boolean>;
    name: pulumi.Input<string>;
    required?: pulumi.Input<boolean>;
    sensitive?: pulumi.Input<boolean>;
    type: pulumi.Input<string>;
    value?: pulumi.Input<string>;
}

export interface ComponentRoutineInput {
    description?: pulumi.Input<string>;
    multiple?: pulumi.Input<boolean>;
    name: pulumi.Input<string>;
    required?: pulumi.Input<boolean>;
    sensitive?: pulumi.Input<boolean>;
    type?: pulumi.Input<string>;
    valueType: pulumi.Input<string>;
    values?: pulumi.Input<pulumi.Input<inputs.ComponentRoutineInputValue>[]>;
}

export interface ComponentRoutineInputValue {
    asset: pulumi.Input<string>;
    attribute: pulumi.Input<string>;
}

export interface ComponentRoutineOutput {
    description?: pulumi.Input<string>;
    multiple?: pulumi.Input<boolean>;
    name: pulumi.Input<string>;
    required?: pulumi.Input<boolean>;
    sensitive?: pulumi.Input<boolean>;
    type?: pulumi.Input<string>;
    valueType: pulumi.Input<string>;
    values?: pulumi.Input<pulumi.Input<inputs.ComponentRoutineOutputValue>[]>;
}

export interface ComponentRoutineOutputValue {
    asset: pulumi.Input<string>;
    attribute: pulumi.Input<string>;
}

export interface DashboardChartChartItem {
    color: pulumi.Input<string>;
    expressionPlain: pulumi.Input<string>;
    hidden?: pulumi.Input<boolean>;
    label?: pulumi.Input<string>;
    /**
     * asset filter
     */
    queryFilterAsset?: pulumi.Input<inputs.DashboardChartChartItemQueryFilterAsset>;
    /**
     * Attribute filter
     */
    queryFilterAttribute?: pulumi.Input<inputs.DashboardChartChartItemQueryFilterAttribute>;
    queryGroupFunction?: pulumi.Input<string>;
    queryGroupUnit?: pulumi.Input<string>;
    queryLimit?: pulumi.Input<number>;
    queryPlain: pulumi.Input<string>;
    querySortDirection?: pulumi.Input<number>;
    refId: pulumi.Input<string>;
    type: pulumi.Input<string>;
}

export interface DashboardChartChartItemQueryFilterAsset {
    /**
     * The ID of this resource.
     */
    id: pulumi.Input<string>;
    name: pulumi.Input<string>;
}

export interface DashboardChartChartItemQueryFilterAttribute {
    /**
     * The ID of this resource.
     */
    id: pulumi.Input<string>;
    name: pulumi.Input<string>;
}

export interface DashboardChartThreshold {
    color: pulumi.Input<string>;
    displayText: pulumi.Input<string>;
    value: pulumi.Input<number>;
}

export interface DashboardChartValueMapping {
    displayText: pulumi.Input<string>;
    matchValue: pulumi.Input<string>;
    order: pulumi.Input<number>;
    type: pulumi.Input<string>;
}

export interface FunctionFunctionItem {
    expressionPlain: pulumi.Input<string>;
    queryPlain: pulumi.Input<string>;
    refId: pulumi.Input<string>;
    type: pulumi.Input<string>;
}

