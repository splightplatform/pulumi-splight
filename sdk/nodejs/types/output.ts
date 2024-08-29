// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface ActionAsset {
    /**
     * asset id
     */
    id: string;
    /**
     * asset name
     */
    name: string;
}

export interface ActionSetpoint {
    /**
     * the target attribute of the setpoint which should also be an attribute of the specified asset
     */
    attribute: outputs.ActionSetpointAttribute;
    /**
     * setpoint ID
     */
    id: string;
    /**
     * setpoint name
     */
    name: string;
    /**
     * JSON encoded scalar value
     */
    value: string;
}

export interface ActionSetpointAttribute {
    /**
     * attribute id
     */
    id: string;
    /**
     * attribute name
     */
    name: string;
}

export interface AlertAlertItem {
    /**
     * how the expression is shown (i.e 'A * 2')
     */
    expression: string;
    /**
     * actual mongo query containing the expression
     */
    expressionPlain: string;
    /**
     * ID of the function item
     */
    id: string;
    /**
     * Asset/Attribute filter
     */
    queryFilterAsset: outputs.AlertAlertItemQueryFilterAsset;
    /**
     * Asset/Attribute filter
     */
    queryFilterAttribute: outputs.AlertAlertItemQueryFilterAttribute;
    /**
     * function used to aggregate data
     */
    queryGroupFunction?: string;
    /**
     * time window to apply the aggregation
     */
    queryGroupUnit?: string;
    /**
     * actual mongo query
     */
    queryPlain: string;
    /**
     * identifier of the variable (i.e 'A')
     */
    refId: string;
    /**
     * either QUERY or EXPRESSION
     */
    type: string;
}

export interface AlertAlertItemQueryFilterAsset {
    /**
     * ID of the resource
     */
    id?: string;
    /**
     * name of the resource
     */
    name?: string;
}

export interface AlertAlertItemQueryFilterAttribute {
    /**
     * ID of the resource
     */
    id?: string;
    /**
     * name of the resource
     */
    name?: string;
}

export interface AlertThreshold {
    /**
     * [alert|warning|no_alert] status value for the threshold
     */
    status: string;
    /**
     * optional custom value to be displayed in the platform.
     */
    statusText?: string;
    /**
     * value to be considered to compare
     */
    value: number;
}

export interface AssetKind {
    /**
     * kind id
     */
    id: string;
    /**
     * kind name
     */
    name: string;
}

export interface AssetTag {
    /**
     * tag id
     */
    id: string;
    /**
     * tag name
     */
    name: string;
}

export interface CommandAction {
    /**
     * asset associated with the action (to be deprecated)
     */
    asset: outputs.CommandActionAsset;
    /**
     * action ID
     */
    id: string;
    /**
     * setpoint name
     */
    name: string;
}

export interface CommandActionAsset {
    /**
     * asset id
     */
    id: string;
    /**
     * asset name
     */
    name: string;
}

export interface ComponentInput {
    description?: string;
    multiple?: boolean;
    name: string;
    required?: boolean;
    sensitive?: boolean;
    type: string;
    value?: string;
}

export interface ComponentRoutineConfig {
    description?: string;
    multiple?: boolean;
    name: string;
    required?: boolean;
    sensitive?: boolean;
    type: string;
    value?: string;
}

export interface ComponentRoutineInput {
    description?: string;
    multiple?: boolean;
    name: string;
    required?: boolean;
    sensitive?: boolean;
    type?: string;
    valueType: string;
    values?: outputs.ComponentRoutineInputValue[];
}

export interface ComponentRoutineInputValue {
    asset: string;
    attribute: string;
}

export interface ComponentRoutineOutput {
    description?: string;
    multiple?: boolean;
    name: string;
    required?: boolean;
    sensitive?: boolean;
    type?: string;
    valueType: string;
    values?: outputs.ComponentRoutineOutputValue[];
}

export interface ComponentRoutineOutputValue {
    asset: string;
    attribute: string;
}

export interface ComponentTag {
    /**
     * tag id
     */
    id: string;
    /**
     * tag name
     */
    name: string;
}

export interface DashboardActionlistChartChartItem {
    color: string;
    expressionPlain: string;
    hidden?: boolean;
    label?: string;
    /**
     * Asset/Attribute filter
     */
    queryFilterAsset: outputs.DashboardActionlistChartChartItemQueryFilterAsset;
    /**
     * Asset/Attribute filter
     */
    queryFilterAttribute: outputs.DashboardActionlistChartChartItemQueryFilterAttribute;
    queryGroupFunction?: string;
    queryGroupUnit?: string;
    queryLimit?: number;
    queryPlain: string;
    querySortDirection?: number;
    refId: string;
    type: string;
}

export interface DashboardActionlistChartChartItemQueryFilterAsset {
    /**
     * ID of the resource
     */
    id?: string;
    /**
     * name of the resource
     */
    name?: string;
}

export interface DashboardActionlistChartChartItemQueryFilterAttribute {
    /**
     * ID of the resource
     */
    id?: string;
    /**
     * name of the resource
     */
    name?: string;
}

export interface DashboardActionlistChartThreshold {
    color: string;
    displayText: string;
    value: number;
}

export interface DashboardActionlistChartValueMapping {
    displayText: string;
    matchValue: string;
    order: number;
    type: string;
}

export interface DashboardAlerteventsChartChartItem {
    color: string;
    expressionPlain: string;
    hidden?: boolean;
    label?: string;
    /**
     * Asset/Attribute filter
     */
    queryFilterAsset: outputs.DashboardAlerteventsChartChartItemQueryFilterAsset;
    /**
     * Asset/Attribute filter
     */
    queryFilterAttribute: outputs.DashboardAlerteventsChartChartItemQueryFilterAttribute;
    queryGroupFunction?: string;
    queryGroupUnit?: string;
    queryLimit?: number;
    queryPlain: string;
    querySortDirection?: number;
    refId: string;
    type: string;
}

export interface DashboardAlerteventsChartChartItemQueryFilterAsset {
    /**
     * ID of the resource
     */
    id?: string;
    /**
     * name of the resource
     */
    name?: string;
}

export interface DashboardAlerteventsChartChartItemQueryFilterAttribute {
    /**
     * ID of the resource
     */
    id?: string;
    /**
     * name of the resource
     */
    name?: string;
}

export interface DashboardAlerteventsChartThreshold {
    color: string;
    displayText: string;
    value: number;
}

export interface DashboardAlerteventsChartValueMapping {
    displayText: string;
    matchValue: string;
    order: number;
    type: string;
}

export interface DashboardAlertlistChartChartItem {
    color: string;
    expressionPlain: string;
    hidden?: boolean;
    label?: string;
    /**
     * Asset/Attribute filter
     */
    queryFilterAsset: outputs.DashboardAlertlistChartChartItemQueryFilterAsset;
    /**
     * Asset/Attribute filter
     */
    queryFilterAttribute: outputs.DashboardAlertlistChartChartItemQueryFilterAttribute;
    queryGroupFunction?: string;
    queryGroupUnit?: string;
    queryLimit?: number;
    queryPlain: string;
    querySortDirection?: number;
    refId: string;
    type: string;
}

export interface DashboardAlertlistChartChartItemQueryFilterAsset {
    /**
     * ID of the resource
     */
    id?: string;
    /**
     * name of the resource
     */
    name?: string;
}

export interface DashboardAlertlistChartChartItemQueryFilterAttribute {
    /**
     * ID of the resource
     */
    id?: string;
    /**
     * name of the resource
     */
    name?: string;
}

export interface DashboardAlertlistChartThreshold {
    color: string;
    displayText: string;
    value: number;
}

export interface DashboardAlertlistChartValueMapping {
    displayText: string;
    matchValue: string;
    order: number;
    type: string;
}

export interface DashboardAssetlistChartChartItem {
    color: string;
    expressionPlain: string;
    hidden?: boolean;
    label?: string;
    /**
     * Asset/Attribute filter
     */
    queryFilterAsset: outputs.DashboardAssetlistChartChartItemQueryFilterAsset;
    /**
     * Asset/Attribute filter
     */
    queryFilterAttribute: outputs.DashboardAssetlistChartChartItemQueryFilterAttribute;
    queryGroupFunction?: string;
    queryGroupUnit?: string;
    queryLimit?: number;
    queryPlain: string;
    querySortDirection?: number;
    refId: string;
    type: string;
}

export interface DashboardAssetlistChartChartItemQueryFilterAsset {
    /**
     * ID of the resource
     */
    id?: string;
    /**
     * name of the resource
     */
    name?: string;
}

export interface DashboardAssetlistChartChartItemQueryFilterAttribute {
    /**
     * ID of the resource
     */
    id?: string;
    /**
     * name of the resource
     */
    name?: string;
}

export interface DashboardAssetlistChartThreshold {
    color: string;
    displayText: string;
    value: number;
}

export interface DashboardAssetlistChartValueMapping {
    displayText: string;
    matchValue: string;
    order: number;
    type: string;
}

export interface DashboardBarChartChartItem {
    color: string;
    expressionPlain: string;
    hidden?: boolean;
    label?: string;
    /**
     * Asset/Attribute filter
     */
    queryFilterAsset: outputs.DashboardBarChartChartItemQueryFilterAsset;
    /**
     * Asset/Attribute filter
     */
    queryFilterAttribute: outputs.DashboardBarChartChartItemQueryFilterAttribute;
    queryGroupFunction?: string;
    queryGroupUnit?: string;
    queryLimit?: number;
    queryPlain: string;
    querySortDirection?: number;
    refId: string;
    type: string;
}

export interface DashboardBarChartChartItemQueryFilterAsset {
    /**
     * ID of the resource
     */
    id?: string;
    /**
     * name of the resource
     */
    name?: string;
}

export interface DashboardBarChartChartItemQueryFilterAttribute {
    /**
     * ID of the resource
     */
    id?: string;
    /**
     * name of the resource
     */
    name?: string;
}

export interface DashboardBarChartThreshold {
    color: string;
    displayText: string;
    value: number;
}

export interface DashboardBarChartValueMapping {
    displayText: string;
    matchValue: string;
    order: number;
    type: string;
}

export interface DashboardBargaugeChartChartItem {
    color: string;
    expressionPlain: string;
    hidden?: boolean;
    label?: string;
    /**
     * Asset/Attribute filter
     */
    queryFilterAsset: outputs.DashboardBargaugeChartChartItemQueryFilterAsset;
    /**
     * Asset/Attribute filter
     */
    queryFilterAttribute: outputs.DashboardBargaugeChartChartItemQueryFilterAttribute;
    queryGroupFunction?: string;
    queryGroupUnit?: string;
    queryLimit?: number;
    queryPlain: string;
    querySortDirection?: number;
    refId: string;
    type: string;
}

export interface DashboardBargaugeChartChartItemQueryFilterAsset {
    /**
     * ID of the resource
     */
    id?: string;
    /**
     * name of the resource
     */
    name?: string;
}

export interface DashboardBargaugeChartChartItemQueryFilterAttribute {
    /**
     * ID of the resource
     */
    id?: string;
    /**
     * name of the resource
     */
    name?: string;
}

export interface DashboardBargaugeChartThreshold {
    color: string;
    displayText: string;
    value: number;
}

export interface DashboardBargaugeChartValueMapping {
    displayText: string;
    matchValue: string;
    order: number;
    type: string;
}

export interface DashboardCommandlistChartChartItem {
    color: string;
    expressionPlain: string;
    hidden?: boolean;
    label?: string;
    /**
     * Asset/Attribute filter
     */
    queryFilterAsset: outputs.DashboardCommandlistChartChartItemQueryFilterAsset;
    /**
     * Asset/Attribute filter
     */
    queryFilterAttribute: outputs.DashboardCommandlistChartChartItemQueryFilterAttribute;
    queryGroupFunction?: string;
    queryGroupUnit?: string;
    queryLimit?: number;
    queryPlain: string;
    querySortDirection?: number;
    refId: string;
    type: string;
}

export interface DashboardCommandlistChartChartItemQueryFilterAsset {
    /**
     * ID of the resource
     */
    id?: string;
    /**
     * name of the resource
     */
    name?: string;
}

export interface DashboardCommandlistChartChartItemQueryFilterAttribute {
    /**
     * ID of the resource
     */
    id?: string;
    /**
     * name of the resource
     */
    name?: string;
}

export interface DashboardCommandlistChartThreshold {
    color: string;
    displayText: string;
    value: number;
}

export interface DashboardCommandlistChartValueMapping {
    displayText: string;
    matchValue: string;
    order: number;
    type: string;
}

export interface DashboardGaugeChartChartItem {
    color: string;
    expressionPlain: string;
    hidden?: boolean;
    label?: string;
    /**
     * Asset/Attribute filter
     */
    queryFilterAsset: outputs.DashboardGaugeChartChartItemQueryFilterAsset;
    /**
     * Asset/Attribute filter
     */
    queryFilterAttribute: outputs.DashboardGaugeChartChartItemQueryFilterAttribute;
    queryGroupFunction?: string;
    queryGroupUnit?: string;
    queryLimit?: number;
    queryPlain: string;
    querySortDirection?: number;
    refId: string;
    type: string;
}

export interface DashboardGaugeChartChartItemQueryFilterAsset {
    /**
     * ID of the resource
     */
    id?: string;
    /**
     * name of the resource
     */
    name?: string;
}

export interface DashboardGaugeChartChartItemQueryFilterAttribute {
    /**
     * ID of the resource
     */
    id?: string;
    /**
     * name of the resource
     */
    name?: string;
}

export interface DashboardGaugeChartThreshold {
    color: string;
    displayText: string;
    value: number;
}

export interface DashboardGaugeChartValueMapping {
    displayText: string;
    matchValue: string;
    order: number;
    type: string;
}

export interface DashboardHistogramChartChartItem {
    color: string;
    expressionPlain: string;
    hidden?: boolean;
    label?: string;
    /**
     * Asset/Attribute filter
     */
    queryFilterAsset: outputs.DashboardHistogramChartChartItemQueryFilterAsset;
    /**
     * Asset/Attribute filter
     */
    queryFilterAttribute: outputs.DashboardHistogramChartChartItemQueryFilterAttribute;
    queryGroupFunction?: string;
    queryGroupUnit?: string;
    queryLimit?: number;
    queryPlain: string;
    querySortDirection?: number;
    refId: string;
    type: string;
}

export interface DashboardHistogramChartChartItemQueryFilterAsset {
    /**
     * ID of the resource
     */
    id?: string;
    /**
     * name of the resource
     */
    name?: string;
}

export interface DashboardHistogramChartChartItemQueryFilterAttribute {
    /**
     * ID of the resource
     */
    id?: string;
    /**
     * name of the resource
     */
    name?: string;
}

export interface DashboardHistogramChartThreshold {
    color: string;
    displayText: string;
    value: number;
}

export interface DashboardHistogramChartValueMapping {
    displayText: string;
    matchValue: string;
    order: number;
    type: string;
}

export interface DashboardImageChartChartItem {
    color: string;
    expressionPlain: string;
    hidden?: boolean;
    label?: string;
    /**
     * Asset/Attribute filter
     */
    queryFilterAsset: outputs.DashboardImageChartChartItemQueryFilterAsset;
    /**
     * Asset/Attribute filter
     */
    queryFilterAttribute: outputs.DashboardImageChartChartItemQueryFilterAttribute;
    queryGroupFunction?: string;
    queryGroupUnit?: string;
    queryLimit?: number;
    queryPlain: string;
    querySortDirection?: number;
    refId: string;
    type: string;
}

export interface DashboardImageChartChartItemQueryFilterAsset {
    /**
     * ID of the resource
     */
    id?: string;
    /**
     * name of the resource
     */
    name?: string;
}

export interface DashboardImageChartChartItemQueryFilterAttribute {
    /**
     * ID of the resource
     */
    id?: string;
    /**
     * name of the resource
     */
    name?: string;
}

export interface DashboardImageChartThreshold {
    color: string;
    displayText: string;
    value: number;
}

export interface DashboardImageChartValueMapping {
    displayText: string;
    matchValue: string;
    order: number;
    type: string;
}

export interface DashboardStatChartChartItem {
    color: string;
    expressionPlain: string;
    hidden?: boolean;
    label?: string;
    /**
     * Asset/Attribute filter
     */
    queryFilterAsset: outputs.DashboardStatChartChartItemQueryFilterAsset;
    /**
     * Asset/Attribute filter
     */
    queryFilterAttribute: outputs.DashboardStatChartChartItemQueryFilterAttribute;
    queryGroupFunction?: string;
    queryGroupUnit?: string;
    queryLimit?: number;
    queryPlain: string;
    querySortDirection?: number;
    refId: string;
    type: string;
}

export interface DashboardStatChartChartItemQueryFilterAsset {
    /**
     * ID of the resource
     */
    id?: string;
    /**
     * name of the resource
     */
    name?: string;
}

export interface DashboardStatChartChartItemQueryFilterAttribute {
    /**
     * ID of the resource
     */
    id?: string;
    /**
     * name of the resource
     */
    name?: string;
}

export interface DashboardStatChartThreshold {
    color: string;
    displayText: string;
    value: number;
}

export interface DashboardStatChartValueMapping {
    displayText: string;
    matchValue: string;
    order: number;
    type: string;
}

export interface DashboardTableChartChartItem {
    color: string;
    expressionPlain: string;
    hidden?: boolean;
    label?: string;
    /**
     * Asset/Attribute filter
     */
    queryFilterAsset: outputs.DashboardTableChartChartItemQueryFilterAsset;
    /**
     * Asset/Attribute filter
     */
    queryFilterAttribute: outputs.DashboardTableChartChartItemQueryFilterAttribute;
    queryGroupFunction?: string;
    queryGroupUnit?: string;
    queryLimit?: number;
    queryPlain: string;
    querySortDirection?: number;
    refId: string;
    type: string;
}

export interface DashboardTableChartChartItemQueryFilterAsset {
    /**
     * ID of the resource
     */
    id?: string;
    /**
     * name of the resource
     */
    name?: string;
}

export interface DashboardTableChartChartItemQueryFilterAttribute {
    /**
     * ID of the resource
     */
    id?: string;
    /**
     * name of the resource
     */
    name?: string;
}

export interface DashboardTableChartThreshold {
    color: string;
    displayText: string;
    value: number;
}

export interface DashboardTableChartValueMapping {
    displayText: string;
    matchValue: string;
    order: number;
    type: string;
}

export interface DashboardTextChartChartItem {
    color: string;
    expressionPlain: string;
    hidden?: boolean;
    label?: string;
    /**
     * Asset/Attribute filter
     */
    queryFilterAsset: outputs.DashboardTextChartChartItemQueryFilterAsset;
    /**
     * Asset/Attribute filter
     */
    queryFilterAttribute: outputs.DashboardTextChartChartItemQueryFilterAttribute;
    queryGroupFunction?: string;
    queryGroupUnit?: string;
    queryLimit?: number;
    queryPlain: string;
    querySortDirection?: number;
    refId: string;
    type: string;
}

export interface DashboardTextChartChartItemQueryFilterAsset {
    /**
     * ID of the resource
     */
    id?: string;
    /**
     * name of the resource
     */
    name?: string;
}

export interface DashboardTextChartChartItemQueryFilterAttribute {
    /**
     * ID of the resource
     */
    id?: string;
    /**
     * name of the resource
     */
    name?: string;
}

export interface DashboardTextChartThreshold {
    color: string;
    displayText: string;
    value: number;
}

export interface DashboardTextChartValueMapping {
    displayText: string;
    matchValue: string;
    order: number;
    type: string;
}

export interface DashboardTimeseriesChartChartItem {
    color: string;
    expressionPlain: string;
    hidden?: boolean;
    label?: string;
    /**
     * Asset/Attribute filter
     */
    queryFilterAsset: outputs.DashboardTimeseriesChartChartItemQueryFilterAsset;
    /**
     * Asset/Attribute filter
     */
    queryFilterAttribute: outputs.DashboardTimeseriesChartChartItemQueryFilterAttribute;
    queryGroupFunction?: string;
    queryGroupUnit?: string;
    queryLimit?: number;
    queryPlain: string;
    querySortDirection?: number;
    refId: string;
    type: string;
}

export interface DashboardTimeseriesChartChartItemQueryFilterAsset {
    /**
     * ID of the resource
     */
    id?: string;
    /**
     * name of the resource
     */
    name?: string;
}

export interface DashboardTimeseriesChartChartItemQueryFilterAttribute {
    /**
     * ID of the resource
     */
    id?: string;
    /**
     * name of the resource
     */
    name?: string;
}

export interface DashboardTimeseriesChartThreshold {
    color: string;
    displayText: string;
    value: number;
}

export interface DashboardTimeseriesChartValueMapping {
    displayText: string;
    matchValue: string;
    order: number;
    type: string;
}

export interface FunctionFunctionItem {
    /**
     * how the expression is shown (i.e 'A * 2')
     */
    expression: string;
    /**
     * actual mongo query containing the expression
     */
    expressionPlain: string;
    /**
     * ID of the function item
     */
    id: string;
    /**
     * Asset/Attribute filter
     */
    queryFilterAsset: outputs.FunctionFunctionItemQueryFilterAsset;
    /**
     * Asset/Attribute filter
     */
    queryFilterAttribute: outputs.FunctionFunctionItemQueryFilterAttribute;
    /**
     * function used to aggregate data
     */
    queryGroupFunction?: string;
    /**
     * time window to apply the aggregation
     */
    queryGroupUnit?: string;
    /**
     * actual mongo query
     */
    queryPlain: string;
    /**
     * identifier of the variable (i.e 'A')
     */
    refId: string;
    /**
     * either QUERY or EXPRESSION
     */
    type: string;
}

export interface FunctionFunctionItemQueryFilterAsset {
    /**
     * ID of the resource
     */
    id?: string;
    /**
     * name of the resource
     */
    name?: string;
}

export interface FunctionFunctionItemQueryFilterAttribute {
    /**
     * ID of the resource
     */
    id?: string;
    /**
     * name of the resource
     */
    name?: string;
}

export interface FunctionTargetAsset {
    /**
     * ID of the resource
     */
    id?: string;
    /**
     * name of the resource
     */
    name?: string;
}

export interface FunctionTargetAttribute {
    /**
     * ID of the resource
     */
    id?: string;
    /**
     * name of the resource
     */
    name?: string;
}

export interface GetAssetKindsKind {
    /**
     * ID of the resource
     */
    id: string;
    /**
     * name of the resource
     */
    name: string;
}

export interface GetTagsTag {
    /**
     * ID of the resource
     */
    id: string;
    /**
     * name of the resource
     */
    name: string;
}

