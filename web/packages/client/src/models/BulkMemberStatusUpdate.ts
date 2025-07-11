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
import type { MembershipStatus } from './MembershipStatus';
import {
    MembershipStatusFromJSON,
    MembershipStatusFromJSONTyped,
    MembershipStatusToJSON,
    MembershipStatusToJSONTyped,
} from './MembershipStatus';

/**
 * 
 * @export
 * @interface BulkMemberStatusUpdate
 */
export interface BulkMemberStatusUpdate {
    /**
     * Whether to notify affected users
     * @type {boolean}
     * @memberof BulkMemberStatusUpdate
     */
    notifyUsers: boolean;
    /**
     * Reason for operation
     * @type {string}
     * @memberof BulkMemberStatusUpdate
     */
    reason?: string;
    /**
     * Status of the membership
     * @type {MembershipStatus}
     * @memberof BulkMemberStatusUpdate
     */
    status: MembershipStatus;
    /**
     * User ID
     * @type {string}
     * @memberof BulkMemberStatusUpdate
     */
    userId: string;
}



/**
 * Check if a given object implements the BulkMemberStatusUpdate interface.
 */
export function instanceOfBulkMemberStatusUpdate(value: object): value is BulkMemberStatusUpdate {
    if (!('notifyUsers' in value) || value['notifyUsers'] === undefined) return false;
    if (!('status' in value) || value['status'] === undefined) return false;
    if (!('userId' in value) || value['userId'] === undefined) return false;
    return true;
}

export function BulkMemberStatusUpdateFromJSON(json: any): BulkMemberStatusUpdate {
    return BulkMemberStatusUpdateFromJSONTyped(json, false);
}

export function BulkMemberStatusUpdateFromJSONTyped(json: any, ignoreDiscriminator: boolean): BulkMemberStatusUpdate {
    if (json == null) {
        return json;
    }
    return {
        
        'notifyUsers': json['notifyUsers'],
        'reason': json['reason'] == null ? undefined : json['reason'],
        'status': MembershipStatusFromJSON(json['status']),
        'userId': json['userId'],
    };
}

export function BulkMemberStatusUpdateToJSON(json: any): BulkMemberStatusUpdate {
    return BulkMemberStatusUpdateToJSONTyped(json, false);
}

export function BulkMemberStatusUpdateToJSONTyped(value?: BulkMemberStatusUpdate | null, ignoreDiscriminator = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'notifyUsers': value['notifyUsers'],
        'reason': value['reason'],
        'status': MembershipStatusToJSON(value['status']),
        'userId': value['userId'],
    };
}

