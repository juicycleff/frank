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
 * @interface MonthlyGrowth
 */
export interface MonthlyGrowth {
    /**
     * Month name
     * @type {string}
     * @memberof MonthlyGrowth
     */
    month: string;
    /**
     * New organizations
     * @type {number}
     * @memberof MonthlyGrowth
     */
    organizations: number;
    /**
     * Revenue
     * @type {number}
     * @memberof MonthlyGrowth
     */
    revenue: number;
    /**
     * New users
     * @type {number}
     * @memberof MonthlyGrowth
     */
    users: number;
}

/**
 * Check if a given object implements the MonthlyGrowth interface.
 */
export function instanceOfMonthlyGrowth(value: object): value is MonthlyGrowth {
    if (!('month' in value) || value['month'] === undefined) return false;
    if (!('organizations' in value) || value['organizations'] === undefined) return false;
    if (!('revenue' in value) || value['revenue'] === undefined) return false;
    if (!('users' in value) || value['users'] === undefined) return false;
    return true;
}

export function MonthlyGrowthFromJSON(json: any): MonthlyGrowth {
    return MonthlyGrowthFromJSONTyped(json, false);
}

export function MonthlyGrowthFromJSONTyped(json: any, ignoreDiscriminator: boolean): MonthlyGrowth {
    if (json == null) {
        return json;
    }
    return {
        
        'month': json['month'],
        'organizations': json['organizations'],
        'revenue': json['revenue'],
        'users': json['users'],
    };
}

export function MonthlyGrowthToJSON(json: any): MonthlyGrowth {
    return MonthlyGrowthToJSONTyped(json, false);
}

export function MonthlyGrowthToJSONTyped(value?: MonthlyGrowth | null, ignoreDiscriminator = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'month': value['month'],
        'organizations': value['organizations'],
        'revenue': value['revenue'],
        'users': value['users'],
    };
}

