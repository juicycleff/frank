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
import type { PlatformPerformanceStats } from './PlatformPerformanceStats';
import {
    PlatformPerformanceStatsFromJSON,
    PlatformPerformanceStatsFromJSONTyped,
    PlatformPerformanceStatsToJSON,
    PlatformPerformanceStatsToJSONTyped,
} from './PlatformPerformanceStats';
import type { PlatformSecurityStats } from './PlatformSecurityStats';
import {
    PlatformSecurityStatsFromJSON,
    PlatformSecurityStatsFromJSONTyped,
    PlatformSecurityStatsToJSON,
    PlatformSecurityStatsToJSONTyped,
} from './PlatformSecurityStats';
import type { PlatformAPIStats } from './PlatformAPIStats';
import {
    PlatformAPIStatsFromJSON,
    PlatformAPIStatsFromJSONTyped,
    PlatformAPIStatsToJSON,
    PlatformAPIStatsToJSONTyped,
} from './PlatformAPIStats';
import type { PlatformGrowthStats } from './PlatformGrowthStats';
import {
    PlatformGrowthStatsFromJSON,
    PlatformGrowthStatsFromJSONTyped,
    PlatformGrowthStatsToJSON,
    PlatformGrowthStatsToJSONTyped,
} from './PlatformGrowthStats';
import type { PlatformOrgStats } from './PlatformOrgStats';
import {
    PlatformOrgStatsFromJSON,
    PlatformOrgStatsFromJSONTyped,
    PlatformOrgStatsToJSON,
    PlatformOrgStatsToJSONTyped,
} from './PlatformOrgStats';
import type { PlatformUserStats } from './PlatformUserStats';
import {
    PlatformUserStatsFromJSON,
    PlatformUserStatsFromJSONTyped,
    PlatformUserStatsToJSON,
    PlatformUserStatsToJSONTyped,
} from './PlatformUserStats';
import type { PlatformRevenueStats } from './PlatformRevenueStats';
import {
    PlatformRevenueStatsFromJSON,
    PlatformRevenueStatsFromJSONTyped,
    PlatformRevenueStatsToJSON,
    PlatformRevenueStatsToJSONTyped,
} from './PlatformRevenueStats';

/**
 * 
 * @export
 * @interface PlatformStats
 */
export interface PlatformStats {
    /**
     * A URL to the JSON Schema for this object.
     * @type {string}
     * @memberof PlatformStats
     */
    readonly $schema?: string;
    /**
     * 
     * @type {PlatformAPIStats}
     * @memberof PlatformStats
     */
    api: PlatformAPIStats;
    /**
     * 
     * @type {Date}
     * @memberof PlatformStats
     */
    generatedAt: Date;
    /**
     * 
     * @type {PlatformGrowthStats}
     * @memberof PlatformStats
     */
    growth: PlatformGrowthStats;
    /**
     * 
     * @type {PlatformOrgStats}
     * @memberof PlatformStats
     */
    organizations: PlatformOrgStats;
    /**
     * 
     * @type {PlatformPerformanceStats}
     * @memberof PlatformStats
     */
    performance: PlatformPerformanceStats;
    /**
     * 
     * @type {string}
     * @memberof PlatformStats
     */
    period: string;
    /**
     * 
     * @type {PlatformRevenueStats}
     * @memberof PlatformStats
     */
    revenue: PlatformRevenueStats;
    /**
     * 
     * @type {PlatformSecurityStats}
     * @memberof PlatformStats
     */
    security: PlatformSecurityStats;
    /**
     * 
     * @type {PlatformUserStats}
     * @memberof PlatformStats
     */
    users: PlatformUserStats;
}

/**
 * Check if a given object implements the PlatformStats interface.
 */
export function instanceOfPlatformStats(value: object): value is PlatformStats {
    if (!('api' in value) || value['api'] === undefined) return false;
    if (!('generatedAt' in value) || value['generatedAt'] === undefined) return false;
    if (!('growth' in value) || value['growth'] === undefined) return false;
    if (!('organizations' in value) || value['organizations'] === undefined) return false;
    if (!('performance' in value) || value['performance'] === undefined) return false;
    if (!('period' in value) || value['period'] === undefined) return false;
    if (!('revenue' in value) || value['revenue'] === undefined) return false;
    if (!('security' in value) || value['security'] === undefined) return false;
    if (!('users' in value) || value['users'] === undefined) return false;
    return true;
}

export function PlatformStatsFromJSON(json: any): PlatformStats {
    return PlatformStatsFromJSONTyped(json, false);
}

export function PlatformStatsFromJSONTyped(json: any, ignoreDiscriminator: boolean): PlatformStats {
    if (json == null) {
        return json;
    }
    return {
        
        '$schema': json['$schema'] == null ? undefined : json['$schema'],
        'api': PlatformAPIStatsFromJSON(json['api']),
        'generatedAt': (new Date(json['generated_at'])),
        'growth': PlatformGrowthStatsFromJSON(json['growth']),
        'organizations': PlatformOrgStatsFromJSON(json['organizations']),
        'performance': PlatformPerformanceStatsFromJSON(json['performance']),
        'period': json['period'],
        'revenue': PlatformRevenueStatsFromJSON(json['revenue']),
        'security': PlatformSecurityStatsFromJSON(json['security']),
        'users': PlatformUserStatsFromJSON(json['users']),
    };
}

export function PlatformStatsToJSON(json: any): PlatformStats {
    return PlatformStatsToJSONTyped(json, false);
}

export function PlatformStatsToJSONTyped(value?: Omit<PlatformStats, '$schema'> | null, ignoreDiscriminator = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'api': PlatformAPIStatsToJSON(value['api']),
        'generated_at': ((value['generatedAt']).toISOString()),
        'growth': PlatformGrowthStatsToJSON(value['growth']),
        'organizations': PlatformOrgStatsToJSON(value['organizations']),
        'performance': PlatformPerformanceStatsToJSON(value['performance']),
        'period': value['period'],
        'revenue': PlatformRevenueStatsToJSON(value['revenue']),
        'security': PlatformSecurityStatsToJSON(value['security']),
        'users': PlatformUserStatsToJSON(value['users']),
    };
}

