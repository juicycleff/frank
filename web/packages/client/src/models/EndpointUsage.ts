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
 * @interface EndpointUsage
 */
export interface EndpointUsage {
    /**
     * Average response time in milliseconds
     * @type {number}
     * @memberof EndpointUsage
     */
    avgResponseTime: number;
    /**
     * API endpoint
     * @type {string}
     * @memberof EndpointUsage
     */
    endpoint: string;
    /**
     * HTTP method
     * @type {string}
     * @memberof EndpointUsage
     */
    method: string;
    /**
     * Request count
     * @type {number}
     * @memberof EndpointUsage
     */
    requestCount: number;
    /**
     * Success rate percentage
     * @type {number}
     * @memberof EndpointUsage
     */
    successRate: number;
}

/**
 * Check if a given object implements the EndpointUsage interface.
 */
export function instanceOfEndpointUsage(value: object): value is EndpointUsage {
    if (!('avgResponseTime' in value) || value['avgResponseTime'] === undefined) return false;
    if (!('endpoint' in value) || value['endpoint'] === undefined) return false;
    if (!('method' in value) || value['method'] === undefined) return false;
    if (!('requestCount' in value) || value['requestCount'] === undefined) return false;
    if (!('successRate' in value) || value['successRate'] === undefined) return false;
    return true;
}

export function EndpointUsageFromJSON(json: any): EndpointUsage {
    return EndpointUsageFromJSONTyped(json, false);
}

export function EndpointUsageFromJSONTyped(json: any, ignoreDiscriminator: boolean): EndpointUsage {
    if (json == null) {
        return json;
    }
    return {
        
        'avgResponseTime': json['avgResponseTime'],
        'endpoint': json['endpoint'],
        'method': json['method'],
        'requestCount': json['requestCount'],
        'successRate': json['successRate'],
    };
}

export function EndpointUsageToJSON(json: any): EndpointUsage {
    return EndpointUsageToJSONTyped(json, false);
}

export function EndpointUsageToJSONTyped(value?: EndpointUsage | null, ignoreDiscriminator = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'avgResponseTime': value['avgResponseTime'],
        'endpoint': value['endpoint'],
        'method': value['method'],
        'requestCount': value['requestCount'],
        'successRate': value['successRate'],
    };
}

