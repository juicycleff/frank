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
 * @interface SeasonalPattern
 */
export interface SeasonalPattern {
    /**
     * 
     * @type {number}
     * @memberof SeasonalPattern
     */
    confidence: number;
    /**
     * 
     * @type {string}
     * @memberof SeasonalPattern
     */
    description: string;
    /**
     * 
     * @type {string}
     * @memberof SeasonalPattern
     */
    pattern: string;
}

/**
 * Check if a given object implements the SeasonalPattern interface.
 */
export function instanceOfSeasonalPattern(value: object): value is SeasonalPattern {
    if (!('confidence' in value) || value['confidence'] === undefined) return false;
    if (!('description' in value) || value['description'] === undefined) return false;
    if (!('pattern' in value) || value['pattern'] === undefined) return false;
    return true;
}

export function SeasonalPatternFromJSON(json: any): SeasonalPattern {
    return SeasonalPatternFromJSONTyped(json, false);
}

export function SeasonalPatternFromJSONTyped(json: any, ignoreDiscriminator: boolean): SeasonalPattern {
    if (json == null) {
        return json;
    }
    return {
        
        'confidence': json['confidence'],
        'description': json['description'],
        'pattern': json['pattern'],
    };
}

export function SeasonalPatternToJSON(json: any): SeasonalPattern {
    return SeasonalPatternToJSONTyped(json, false);
}

export function SeasonalPatternToJSONTyped(value?: SeasonalPattern | null, ignoreDiscriminator = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'confidence': value['confidence'],
        'description': value['description'],
        'pattern': value['pattern'],
    };
}

