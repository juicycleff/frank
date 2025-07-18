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
 * @interface PlatformPerformanceStats
 */
export interface PlatformPerformanceStats {
    /**
     * 
     * @type {number}
     * @memberof PlatformPerformanceStats
     */
    avgResponseTimeMs: number;
    /**
     * 
     * @type {number}
     * @memberof PlatformPerformanceStats
     */
    cacheHitRatePercent: number;
    /**
     * 
     * @type {number}
     * @memberof PlatformPerformanceStats
     */
    databaseLatencyMs: number;
    /**
     * 
     * @type {number}
     * @memberof PlatformPerformanceStats
     */
    p95ResponseTimeMs: number;
    /**
     * 
     * @type {number}
     * @memberof PlatformPerformanceStats
     */
    p99ResponseTimeMs: number;
    /**
     * 
     * @type {number}
     * @memberof PlatformPerformanceStats
     */
    uptimePercent: number;
}

/**
 * Check if a given object implements the PlatformPerformanceStats interface.
 */
export function instanceOfPlatformPerformanceStats(value: object): value is PlatformPerformanceStats {
    if (!('avgResponseTimeMs' in value) || value['avgResponseTimeMs'] === undefined) return false;
    if (!('cacheHitRatePercent' in value) || value['cacheHitRatePercent'] === undefined) return false;
    if (!('databaseLatencyMs' in value) || value['databaseLatencyMs'] === undefined) return false;
    if (!('p95ResponseTimeMs' in value) || value['p95ResponseTimeMs'] === undefined) return false;
    if (!('p99ResponseTimeMs' in value) || value['p99ResponseTimeMs'] === undefined) return false;
    if (!('uptimePercent' in value) || value['uptimePercent'] === undefined) return false;
    return true;
}

export function PlatformPerformanceStatsFromJSON(json: any): PlatformPerformanceStats {
    return PlatformPerformanceStatsFromJSONTyped(json, false);
}

export function PlatformPerformanceStatsFromJSONTyped(json: any, ignoreDiscriminator: boolean): PlatformPerformanceStats {
    if (json == null) {
        return json;
    }
    return {
        
        'avgResponseTimeMs': json['avg_response_time_ms'],
        'cacheHitRatePercent': json['cache_hit_rate_percent'],
        'databaseLatencyMs': json['database_latency_ms'],
        'p95ResponseTimeMs': json['p95_response_time_ms'],
        'p99ResponseTimeMs': json['p99_response_time_ms'],
        'uptimePercent': json['uptime_percent'],
    };
}

export function PlatformPerformanceStatsToJSON(json: any): PlatformPerformanceStats {
    return PlatformPerformanceStatsToJSONTyped(json, false);
}

export function PlatformPerformanceStatsToJSONTyped(value?: PlatformPerformanceStats | null, ignoreDiscriminator = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'avg_response_time_ms': value['avgResponseTimeMs'],
        'cache_hit_rate_percent': value['cacheHitRatePercent'],
        'database_latency_ms': value['databaseLatencyMs'],
        'p95_response_time_ms': value['p95ResponseTimeMs'],
        'p99_response_time_ms': value['p99ResponseTimeMs'],
        'uptime_percent': value['uptimePercent'],
    };
}

