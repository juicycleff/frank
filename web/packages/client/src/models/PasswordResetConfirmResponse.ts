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
 * @interface PasswordResetConfirmResponse
 */
export interface PasswordResetConfirmResponse {
    [key: string]: any | any;
    /**
     * A URL to the JSON Schema for this object.
     * @type {string}
     * @memberof PasswordResetConfirmResponse
     */
    readonly $schema?: string;
    /**
     * Response message
     * @type {string}
     * @memberof PasswordResetConfirmResponse
     */
    message: string;
    /**
     * Whether password was reset successfully
     * @type {boolean}
     * @memberof PasswordResetConfirmResponse
     */
    success: boolean;
}

/**
 * Check if a given object implements the PasswordResetConfirmResponse interface.
 */
export function instanceOfPasswordResetConfirmResponse(value: object): value is PasswordResetConfirmResponse {
    if (!('message' in value) || value['message'] === undefined) return false;
    if (!('success' in value) || value['success'] === undefined) return false;
    return true;
}

export function PasswordResetConfirmResponseFromJSON(json: any): PasswordResetConfirmResponse {
    return PasswordResetConfirmResponseFromJSONTyped(json, false);
}

export function PasswordResetConfirmResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): PasswordResetConfirmResponse {
    if (json == null) {
        return json;
    }
    return {
        
            ...json,
        '$schema': json['$schema'] == null ? undefined : json['$schema'],
        'message': json['message'],
        'success': json['success'],
    };
}

export function PasswordResetConfirmResponseToJSON(json: any): PasswordResetConfirmResponse {
    return PasswordResetConfirmResponseToJSONTyped(json, false);
}

export function PasswordResetConfirmResponseToJSONTyped(value?: Omit<PasswordResetConfirmResponse, '$schema'> | null, ignoreDiscriminator = false): any {
    if (value == null) {
        return value;
    }

    return {
        
            ...value,
        'message': value['message'],
        'success': value['success'],
    };
}

