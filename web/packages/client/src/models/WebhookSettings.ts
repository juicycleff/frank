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
 * @interface WebhookSettings
 */
export interface WebhookSettings {
    /**
     * Allowed webhook events
     * @type {Array<string>}
     * @memberof WebhookSettings
     */
    allowedEvents: Array<string> | null;
    /**
     * Whether webhooks are enabled
     * @type {boolean}
     * @memberof WebhookSettings
     */
    enabled: boolean;
    /**
     * Number of retry attempts
     * @type {number}
     * @memberof WebhookSettings
     */
    retryAttempts: number;
    /**
     * Webhook timeout in seconds
     * @type {number}
     * @memberof WebhookSettings
     */
    timeoutSeconds: number;
}

/**
 * Check if a given object implements the WebhookSettings interface.
 */
export function instanceOfWebhookSettings(value: object): value is WebhookSettings {
    if (!('allowedEvents' in value) || value['allowedEvents'] === undefined) return false;
    if (!('enabled' in value) || value['enabled'] === undefined) return false;
    if (!('retryAttempts' in value) || value['retryAttempts'] === undefined) return false;
    if (!('timeoutSeconds' in value) || value['timeoutSeconds'] === undefined) return false;
    return true;
}

export function WebhookSettingsFromJSON(json: any): WebhookSettings {
    return WebhookSettingsFromJSONTyped(json, false);
}

export function WebhookSettingsFromJSONTyped(json: any, ignoreDiscriminator: boolean): WebhookSettings {
    if (json == null) {
        return json;
    }
    return {
        
        'allowedEvents': json['allowedEvents'] == null ? null : json['allowedEvents'],
        'enabled': json['enabled'],
        'retryAttempts': json['retryAttempts'],
        'timeoutSeconds': json['timeoutSeconds'],
    };
}

export function WebhookSettingsToJSON(json: any): WebhookSettings {
    return WebhookSettingsToJSONTyped(json, false);
}

export function WebhookSettingsToJSONTyped(value?: WebhookSettings | null, ignoreDiscriminator = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'allowedEvents': value['allowedEvents'],
        'enabled': value['enabled'],
        'retryAttempts': value['retryAttempts'],
        'timeoutSeconds': value['timeoutSeconds'],
    };
}

