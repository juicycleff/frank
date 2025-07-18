/* tslint:disable */
/* eslint-disable */
/**
 * Frank Authentication API
 * Multi-tenant authentication SaaS platform API with Clerk.js compatibility
 *
 * The version of the OpenAPI document: 1.0.0
 * Contact: support@frankauth.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface FeatureCondition
 */
export interface FeatureCondition {
    /**
     * 
     * @type {string}
     * @memberof FeatureCondition
     */
    attribute: string;
    /**
     * 
     * @type {string}
     * @memberof FeatureCondition
     */
    operator: string;
    /**
     * 
     * @type {any}
     * @memberof FeatureCondition
     */
    value: any | null;
}

/**
 * Check if a given object implements the FeatureCondition interface.
 */
export function instanceOfFeatureCondition(value: object): value is FeatureCondition {
    if (!('attribute' in value) || value['attribute'] === undefined) return false;
    if (!('operator' in value) || value['operator'] === undefined) return false;
    if (!('value' in value) || value['value'] === undefined) return false;
    return true;
}

export function FeatureConditionFromJSON(json: any): FeatureCondition {
    return FeatureConditionFromJSONTyped(json, false);
}

export function FeatureConditionFromJSONTyped(json: any, ignoreDiscriminator: boolean): FeatureCondition {
    if (json == null) {
        return json;
    }
    return {
        
        'attribute': json['attribute'],
        'operator': json['operator'],
        'value': json['value'],
    };
}

export function FeatureConditionToJSON(json: any): FeatureCondition {
    return FeatureConditionToJSONTyped(json, false);
}

export function FeatureConditionToJSONTyped(value?: FeatureCondition | null, ignoreDiscriminator = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'attribute': value['attribute'],
        'operator': value['operator'],
        'value': value['value'],
    };
}

