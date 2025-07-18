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
 * @interface UserStats
 */
export interface UserStats {
    /**
     * A URL to the JSON Schema for this object.
     * @type {string}
     * @memberof UserStats
     */
    readonly $schema?: string;
    /**
     * Number of active users
     * @type {number}
     * @memberof UserStats
     */
    activeUsers: number;
    /**
     * Number of end users
     * @type {number}
     * @memberof UserStats
     */
    endUsers: number;
    /**
     * Number of external users
     * @type {number}
     * @memberof UserStats
     */
    externalUsers: number;
    /**
     * Number of internal users
     * @type {number}
     * @memberof UserStats
     */
    internalUsers: number;
    /**
     * Number of users with MFA enabled
     * @type {number}
     * @memberof UserStats
     */
    mfaEnabled: number;
    /**
     * Number of users logged in recently
     * @type {number}
     * @memberof UserStats
     */
    recentLogins: number;
    /**
     * Total number of users
     * @type {number}
     * @memberof UserStats
     */
    totalUsers: number;
    /**
     * Number of users with verified emails
     * @type {number}
     * @memberof UserStats
     */
    verifiedEmails: number;
    /**
     * Number of users with verified phones
     * @type {number}
     * @memberof UserStats
     */
    verifiedPhones: number;
}

/**
 * Check if a given object implements the UserStats interface.
 */
export function instanceOfUserStats(value: object): value is UserStats {
    if (!('activeUsers' in value) || value['activeUsers'] === undefined) return false;
    if (!('endUsers' in value) || value['endUsers'] === undefined) return false;
    if (!('externalUsers' in value) || value['externalUsers'] === undefined) return false;
    if (!('internalUsers' in value) || value['internalUsers'] === undefined) return false;
    if (!('mfaEnabled' in value) || value['mfaEnabled'] === undefined) return false;
    if (!('recentLogins' in value) || value['recentLogins'] === undefined) return false;
    if (!('totalUsers' in value) || value['totalUsers'] === undefined) return false;
    if (!('verifiedEmails' in value) || value['verifiedEmails'] === undefined) return false;
    if (!('verifiedPhones' in value) || value['verifiedPhones'] === undefined) return false;
    return true;
}

export function UserStatsFromJSON(json: any): UserStats {
    return UserStatsFromJSONTyped(json, false);
}

export function UserStatsFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserStats {
    if (json == null) {
        return json;
    }
    return {
        
        '$schema': json['$schema'] == null ? undefined : json['$schema'],
        'activeUsers': json['activeUsers'],
        'endUsers': json['endUsers'],
        'externalUsers': json['externalUsers'],
        'internalUsers': json['internalUsers'],
        'mfaEnabled': json['mfaEnabled'],
        'recentLogins': json['recentLogins'],
        'totalUsers': json['totalUsers'],
        'verifiedEmails': json['verifiedEmails'],
        'verifiedPhones': json['verifiedPhones'],
    };
}

export function UserStatsToJSON(json: any): UserStats {
    return UserStatsToJSONTyped(json, false);
}

export function UserStatsToJSONTyped(value?: Omit<UserStats, '$schema'> | null, ignoreDiscriminator = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'activeUsers': value['activeUsers'],
        'endUsers': value['endUsers'],
        'externalUsers': value['externalUsers'],
        'internalUsers': value['internalUsers'],
        'mfaEnabled': value['mfaEnabled'],
        'recentLogins': value['recentLogins'],
        'totalUsers': value['totalUsers'],
        'verifiedEmails': value['verifiedEmails'],
        'verifiedPhones': value['verifiedPhones'],
    };
}

