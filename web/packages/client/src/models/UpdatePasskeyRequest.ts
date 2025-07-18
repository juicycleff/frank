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
 * @interface UpdatePasskeyRequest
 */
export interface UpdatePasskeyRequest {
    /**
     * A URL to the JSON Schema for this object.
     * @type {string}
     * @memberof UpdatePasskeyRequest
     */
    readonly $schema?: string;
    /**
     * Updated active status
     * @type {boolean}
     * @memberof UpdatePasskeyRequest
     */
    active?: boolean;
    /**
     * Updated name
     * @type {string}
     * @memberof UpdatePasskeyRequest
     */
    name?: string;
}

/**
 * Check if a given object implements the UpdatePasskeyRequest interface.
 */
export function instanceOfUpdatePasskeyRequest(value: object): value is UpdatePasskeyRequest {
    return true;
}

export function UpdatePasskeyRequestFromJSON(json: any): UpdatePasskeyRequest {
    return UpdatePasskeyRequestFromJSONTyped(json, false);
}

export function UpdatePasskeyRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UpdatePasskeyRequest {
    if (json == null) {
        return json;
    }
    return {
        
        '$schema': json['$schema'] == null ? undefined : json['$schema'],
        'active': json['active'] == null ? undefined : json['active'],
        'name': json['name'] == null ? undefined : json['name'],
    };
}

export function UpdatePasskeyRequestToJSON(json: any): UpdatePasskeyRequest {
    return UpdatePasskeyRequestToJSONTyped(json, false);
}

export function UpdatePasskeyRequestToJSONTyped(value?: Omit<UpdatePasskeyRequest, '$schema'> | null, ignoreDiscriminator = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'active': value['active'],
        'name': value['name'],
    };
}

