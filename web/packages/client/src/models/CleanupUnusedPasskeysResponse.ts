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
 * @interface CleanupUnusedPasskeysResponse
 */
export interface CleanupUnusedPasskeysResponse {
    /**
     * A URL to the JSON Schema for this object.
     * @type {string}
     * @memberof CleanupUnusedPasskeysResponse
     */
    readonly $schema?: string;
    /**
     * Number of passkeys deleted
     * @type {number}
     * @memberof CleanupUnusedPasskeysResponse
     */
    deletedCount: number;
}

/**
 * Check if a given object implements the CleanupUnusedPasskeysResponse interface.
 */
export function instanceOfCleanupUnusedPasskeysResponse(value: object): value is CleanupUnusedPasskeysResponse {
    if (!('deletedCount' in value) || value['deletedCount'] === undefined) return false;
    return true;
}

export function CleanupUnusedPasskeysResponseFromJSON(json: any): CleanupUnusedPasskeysResponse {
    return CleanupUnusedPasskeysResponseFromJSONTyped(json, false);
}

export function CleanupUnusedPasskeysResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): CleanupUnusedPasskeysResponse {
    if (json == null) {
        return json;
    }
    return {
        
        '$schema': json['$schema'] == null ? undefined : json['$schema'],
        'deletedCount': json['deletedCount'],
    };
}

export function CleanupUnusedPasskeysResponseToJSON(json: any): CleanupUnusedPasskeysResponse {
    return CleanupUnusedPasskeysResponseToJSONTyped(json, false);
}

export function CleanupUnusedPasskeysResponseToJSONTyped(value?: Omit<CleanupUnusedPasskeysResponse, '$schema'> | null, ignoreDiscriminator = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'deletedCount': value['deletedCount'],
    };
}

