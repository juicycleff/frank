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
import type { Passkey } from './Passkey';
import {
    PasskeyFromJSON,
    PasskeyFromJSONTyped,
    PasskeyToJSON,
    PasskeyToJSONTyped,
} from './Passkey';

/**
 * 
 * @export
 * @interface PasskeyRegistrationFinishResponse
 */
export interface PasskeyRegistrationFinishResponse {
    [key: string]: any | any;
    /**
     * A URL to the JSON Schema for this object.
     * @type {string}
     * @memberof PasskeyRegistrationFinishResponse
     */
    readonly $schema?: string;
    /**
     * Success message
     * @type {string}
     * @memberof PasskeyRegistrationFinishResponse
     */
    message: string;
    /**
     * Created passkey
     * @type {Passkey}
     * @memberof PasskeyRegistrationFinishResponse
     */
    passkey: Passkey;
    /**
     * Whether registration was successful
     * @type {boolean}
     * @memberof PasskeyRegistrationFinishResponse
     */
    success: boolean;
}

/**
 * Check if a given object implements the PasskeyRegistrationFinishResponse interface.
 */
export function instanceOfPasskeyRegistrationFinishResponse(value: object): value is PasskeyRegistrationFinishResponse {
    if (!('message' in value) || value['message'] === undefined) return false;
    if (!('passkey' in value) || value['passkey'] === undefined) return false;
    if (!('success' in value) || value['success'] === undefined) return false;
    return true;
}

export function PasskeyRegistrationFinishResponseFromJSON(json: any): PasskeyRegistrationFinishResponse {
    return PasskeyRegistrationFinishResponseFromJSONTyped(json, false);
}

export function PasskeyRegistrationFinishResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): PasskeyRegistrationFinishResponse {
    if (json == null) {
        return json;
    }
    return {
        
            ...json,
        '$schema': json['$schema'] == null ? undefined : json['$schema'],
        'message': json['message'],
        'passkey': PasskeyFromJSON(json['passkey']),
        'success': json['success'],
    };
}

export function PasskeyRegistrationFinishResponseToJSON(json: any): PasskeyRegistrationFinishResponse {
    return PasskeyRegistrationFinishResponseToJSONTyped(json, false);
}

export function PasskeyRegistrationFinishResponseToJSONTyped(value?: Omit<PasskeyRegistrationFinishResponse, '$schema'> | null, ignoreDiscriminator = false): any {
    if (value == null) {
        return value;
    }

    return {
        
            ...value,
        'message': value['message'],
        'passkey': PasskeyToJSON(value['passkey']),
        'success': value['success'],
    };
}

