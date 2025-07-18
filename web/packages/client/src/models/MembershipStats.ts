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
 * @interface MembershipStats
 */
export interface MembershipStats {
    /**
     * A URL to the JSON Schema for this object.
     * @type {string}
     * @memberof MembershipStats
     */
    readonly $schema?: string;
    /**
     * Active members
     * @type {number}
     * @memberof MembershipStats
     */
    activeMembers: number;
    /**
     * Billing contacts
     * @type {number}
     * @memberof MembershipStats
     */
    billingContacts: number;
    /**
     * Expired invitations
     * @type {number}
     * @memberof MembershipStats
     */
    expiredInvitations: number;
    /**
     * Growth rate
     * @type {number}
     * @memberof MembershipStats
     */
    growthRate: number;
    /**
     * Inactive members
     * @type {number}
     * @memberof MembershipStats
     */
    inactiveMembers: number;
    /**
     * Members by role
     * @type {{ [key: string]: number; }}
     * @memberof MembershipStats
     */
    membersByRole: { [key: string]: number; };
    /**
     * Pending invitations
     * @type {number}
     * @memberof MembershipStats
     */
    pendingInvitations: number;
    /**
     * Primary contacts
     * @type {number}
     * @memberof MembershipStats
     */
    primaryContacts: number;
    /**
     * Invitations sent in last 30 days
     * @type {number}
     * @memberof MembershipStats
     */
    recentInvites: number;
    /**
     * Members joined in last 30 days
     * @type {number}
     * @memberof MembershipStats
     */
    recentJoins: number;
    /**
     * Suspended members
     * @type {number}
     * @memberof MembershipStats
     */
    suspendedMembers: number;
    /**
     * Total members
     * @type {number}
     * @memberof MembershipStats
     */
    totalMembers: number;
}

/**
 * Check if a given object implements the MembershipStats interface.
 */
export function instanceOfMembershipStats(value: object): value is MembershipStats {
    if (!('activeMembers' in value) || value['activeMembers'] === undefined) return false;
    if (!('billingContacts' in value) || value['billingContacts'] === undefined) return false;
    if (!('expiredInvitations' in value) || value['expiredInvitations'] === undefined) return false;
    if (!('growthRate' in value) || value['growthRate'] === undefined) return false;
    if (!('inactiveMembers' in value) || value['inactiveMembers'] === undefined) return false;
    if (!('membersByRole' in value) || value['membersByRole'] === undefined) return false;
    if (!('pendingInvitations' in value) || value['pendingInvitations'] === undefined) return false;
    if (!('primaryContacts' in value) || value['primaryContacts'] === undefined) return false;
    if (!('recentInvites' in value) || value['recentInvites'] === undefined) return false;
    if (!('recentJoins' in value) || value['recentJoins'] === undefined) return false;
    if (!('suspendedMembers' in value) || value['suspendedMembers'] === undefined) return false;
    if (!('totalMembers' in value) || value['totalMembers'] === undefined) return false;
    return true;
}

export function MembershipStatsFromJSON(json: any): MembershipStats {
    return MembershipStatsFromJSONTyped(json, false);
}

export function MembershipStatsFromJSONTyped(json: any, ignoreDiscriminator: boolean): MembershipStats {
    if (json == null) {
        return json;
    }
    return {
        
        '$schema': json['$schema'] == null ? undefined : json['$schema'],
        'activeMembers': json['activeMembers'],
        'billingContacts': json['billingContacts'],
        'expiredInvitations': json['expiredInvitations'],
        'growthRate': json['growthRate'],
        'inactiveMembers': json['inactiveMembers'],
        'membersByRole': json['membersByRole'],
        'pendingInvitations': json['pendingInvitations'],
        'primaryContacts': json['primaryContacts'],
        'recentInvites': json['recentInvites'],
        'recentJoins': json['recentJoins'],
        'suspendedMembers': json['suspendedMembers'],
        'totalMembers': json['totalMembers'],
    };
}

export function MembershipStatsToJSON(json: any): MembershipStats {
    return MembershipStatsToJSONTyped(json, false);
}

export function MembershipStatsToJSONTyped(value?: Omit<MembershipStats, '$schema'> | null, ignoreDiscriminator = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'activeMembers': value['activeMembers'],
        'billingContacts': value['billingContacts'],
        'expiredInvitations': value['expiredInvitations'],
        'growthRate': value['growthRate'],
        'inactiveMembers': value['inactiveMembers'],
        'membersByRole': value['membersByRole'],
        'pendingInvitations': value['pendingInvitations'],
        'primaryContacts': value['primaryContacts'],
        'recentInvites': value['recentInvites'],
        'recentJoins': value['recentJoins'],
        'suspendedMembers': value['suspendedMembers'],
        'totalMembers': value['totalMembers'],
    };
}

