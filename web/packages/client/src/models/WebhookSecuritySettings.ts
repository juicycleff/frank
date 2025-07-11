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
 * @interface WebhookSecuritySettings
 */
export interface WebhookSecuritySettings {
    [key: string]: any | any;
    /**
     * A URL to the JSON Schema for this object.
     * @type {string}
     * @memberof WebhookSecuritySettings
     */
    readonly $schema?: string;
    /**
     * Custom security headers
     * @type {{ [key: string]: string; }}
     * @memberof WebhookSecuritySettings
     */
    customHeaders?: { [key: string]: string; };
    /**
     * Allowed source IP ranges
     * @type {Array<string>}
     * @memberof WebhookSecuritySettings
     */
    ipWhitelist?: Array<string>;
    /**
     * Whether HTTPS is required
     * @type {boolean}
     * @memberof WebhookSecuritySettings
     */
    requireHttps: boolean;
    /**
     * Signature algorithm
     * @type {string}
     * @memberof WebhookSecuritySettings
     */
    signatureAlgorithm: string;
    /**
     * Whether signature verification is enabled
     * @type {boolean}
     * @memberof WebhookSecuritySettings
     */
    signatureEnabled: boolean;
    /**
     * Whether to verify SSL certificates
     * @type {boolean}
     * @memberof WebhookSecuritySettings
     */
    verifySsl: boolean;
    /**
     * Webhook ID
     * @type {string}
     * @memberof WebhookSecuritySettings
     */
    webhookId: string;
}

/**
 * Check if a given object implements the WebhookSecuritySettings interface.
 */
export function instanceOfWebhookSecuritySettings(value: object): value is WebhookSecuritySettings {
    if (!('requireHttps' in value) || value['requireHttps'] === undefined) return false;
    if (!('signatureAlgorithm' in value) || value['signatureAlgorithm'] === undefined) return false;
    if (!('signatureEnabled' in value) || value['signatureEnabled'] === undefined) return false;
    if (!('verifySsl' in value) || value['verifySsl'] === undefined) return false;
    if (!('webhookId' in value) || value['webhookId'] === undefined) return false;
    return true;
}

export function WebhookSecuritySettingsFromJSON(json: any): WebhookSecuritySettings {
    return WebhookSecuritySettingsFromJSONTyped(json, false);
}

export function WebhookSecuritySettingsFromJSONTyped(json: any, ignoreDiscriminator: boolean): WebhookSecuritySettings {
    if (json == null) {
        return json;
    }
    return {
        
            ...json,
        '$schema': json['$schema'] == null ? undefined : json['$schema'],
        'customHeaders': json['customHeaders'] == null ? undefined : json['customHeaders'],
        'ipWhitelist': json['ipWhitelist'] == null ? undefined : json['ipWhitelist'],
        'requireHttps': json['requireHttps'],
        'signatureAlgorithm': json['signatureAlgorithm'],
        'signatureEnabled': json['signatureEnabled'],
        'verifySsl': json['verifySsl'],
        'webhookId': json['webhookId'],
    };
}

export function WebhookSecuritySettingsToJSON(json: any): WebhookSecuritySettings {
    return WebhookSecuritySettingsToJSONTyped(json, false);
}

export function WebhookSecuritySettingsToJSONTyped(value?: Omit<WebhookSecuritySettings, '$schema'> | null, ignoreDiscriminator = false): any {
    if (value == null) {
        return value;
    }

    return {
        
            ...value,
        'customHeaders': value['customHeaders'],
        'ipWhitelist': value['ipWhitelist'],
        'requireHttps': value['requireHttps'],
        'signatureAlgorithm': value['signatureAlgorithm'],
        'signatureEnabled': value['signatureEnabled'],
        'verifySsl': value['verifySsl'],
        'webhookId': value['webhookId'],
    };
}

