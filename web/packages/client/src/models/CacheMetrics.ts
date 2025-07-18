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
 * @interface CacheMetrics
 */
export interface CacheMetrics {
    /**
     * 
     * @type {number}
     * @memberof CacheMetrics
     */
    evictionRate: number;
    /**
     * 
     * @type {number}
     * @memberof CacheMetrics
     */
    expiredKeys: number;
    /**
     * 
     * @type {number}
     * @memberof CacheMetrics
     */
    hitRatePercent: number;
    /**
     * 
     * @type {number}
     * @memberof CacheMetrics
     */
    keyCount: number;
    /**
     * 
     * @type {number}
     * @memberof CacheMetrics
     */
    memoryUsagePercent: number;
    /**
     * 
     * @type {number}
     * @memberof CacheMetrics
     */
    missRatePercent: number;
}

/**
 * Check if a given object implements the CacheMetrics interface.
 */
export function instanceOfCacheMetrics(value: object): value is CacheMetrics {
    if (!('evictionRate' in value) || value['evictionRate'] === undefined) return false;
    if (!('expiredKeys' in value) || value['expiredKeys'] === undefined) return false;
    if (!('hitRatePercent' in value) || value['hitRatePercent'] === undefined) return false;
    if (!('keyCount' in value) || value['keyCount'] === undefined) return false;
    if (!('memoryUsagePercent' in value) || value['memoryUsagePercent'] === undefined) return false;
    if (!('missRatePercent' in value) || value['missRatePercent'] === undefined) return false;
    return true;
}

export function CacheMetricsFromJSON(json: any): CacheMetrics {
    return CacheMetricsFromJSONTyped(json, false);
}

export function CacheMetricsFromJSONTyped(json: any, ignoreDiscriminator: boolean): CacheMetrics {
    if (json == null) {
        return json;
    }
    return {
        
        'evictionRate': json['eviction_rate'],
        'expiredKeys': json['expired_keys'],
        'hitRatePercent': json['hit_rate_percent'],
        'keyCount': json['key_count'],
        'memoryUsagePercent': json['memory_usage_percent'],
        'missRatePercent': json['miss_rate_percent'],
    };
}

export function CacheMetricsToJSON(json: any): CacheMetrics {
    return CacheMetricsToJSONTyped(json, false);
}

export function CacheMetricsToJSONTyped(value?: CacheMetrics | null, ignoreDiscriminator = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'eviction_rate': value['evictionRate'],
        'expired_keys': value['expiredKeys'],
        'hit_rate_percent': value['hitRatePercent'],
        'key_count': value['keyCount'],
        'memory_usage_percent': value['memoryUsagePercent'],
        'miss_rate_percent': value['missRatePercent'],
    };
}

