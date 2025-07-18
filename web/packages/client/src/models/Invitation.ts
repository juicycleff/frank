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
import type { OrganizationSummary } from './OrganizationSummary';
import {
    OrganizationSummaryFromJSON,
    OrganizationSummaryFromJSONTyped,
    OrganizationSummaryToJSON,
    OrganizationSummaryToJSONTyped,
} from './OrganizationSummary';
import type { UserSummary } from './UserSummary';
import {
    UserSummaryFromJSON,
    UserSummaryFromJSONTyped,
    UserSummaryToJSON,
    UserSummaryToJSONTyped,
} from './UserSummary';
import type { RoleSummary } from './RoleSummary';
import {
    RoleSummaryFromJSON,
    RoleSummaryFromJSONTyped,
    RoleSummaryToJSON,
    RoleSummaryToJSONTyped,
} from './RoleSummary';

/**
 * 
 * @export
 * @interface Invitation
 */
export interface Invitation {
    [key: string]: any | any;
    /**
     * A URL to the JSON Schema for this object.
     * @type {string}
     * @memberof Invitation
     */
    readonly $schema?: string;
    /**
     * When invitation was accepted
     * @type {Date}
     * @memberof Invitation
     */
    acceptedAt?: Date;
    /**
     * User who accepted the invitation
     * @type {string}
     * @memberof Invitation
     */
    acceptedBy?: string;
    /**
     * 
     * @type {Date}
     * @memberof Invitation
     */
    createdAt: Date;
    /**
     * Custom invitation fields
     * @type {object}
     * @memberof Invitation
     */
    customFields?: object;
    /**
     * When invitation was declined
     * @type {Date}
     * @memberof Invitation
     */
    declinedAt?: Date;
    /**
     * Invited user email address
     * @type {string}
     * @memberof Invitation
     */
    email: string;
    /**
     * Invitation expiration time
     * @type {Date}
     * @memberof Invitation
     */
    expiresAt: Date;
    /**
     * 
     * @type {string}
     * @memberof Invitation
     */
    id: string;
    /**
     * User who sent the invitation
     * @type {string}
     * @memberof Invitation
     */
    invitedBy?: string;
    /**
     * User who sent the invitation
     * @type {UserSummary}
     * @memberof Invitation
     */
    inviter?: UserSummary;
    /**
     * When invitation was last sent
     * @type {Date}
     * @memberof Invitation
     */
    lastSentAt?: Date;
    /**
     * Personal message from inviter
     * @type {string}
     * @memberof Invitation
     */
    message?: string;
    /**
     * Organization information
     * @type {OrganizationSummary}
     * @memberof Invitation
     */
    organization?: OrganizationSummary;
    /**
     * Organization ID
     * @type {string}
     * @memberof Invitation
     */
    organizationId: string;
    /**
     * URL to redirect to after acceptance
     * @type {string}
     * @memberof Invitation
     */
    redirectUrl?: string;
    /**
     * Role information
     * @type {RoleSummary}
     * @memberof Invitation
     */
    role?: RoleSummary;
    /**
     * Role ID to assign
     * @type {string}
     * @memberof Invitation
     */
    roleId: string;
    /**
     * Number of times invitation was sent
     * @type {number}
     * @memberof Invitation
     */
    sendCount: number;
    /**
     * Invitation status
     * @type {string}
     * @memberof Invitation
     */
    status: InvitationStatusEnum;
    /**
     * Invitation token
     * @type {string}
     * @memberof Invitation
     */
    token?: string;
    /**
     * 
     * @type {Date}
     * @memberof Invitation
     */
    updatedAt: Date;
}


/**
 * @export
 */
export const InvitationStatusEnum = {
    Pending: 'pending',
    Accepted: 'accepted',
    Declined: 'declined',
    Expired: 'expired',
    Cancelled: 'cancelled'
} as const;
export type InvitationStatusEnum = typeof InvitationStatusEnum[keyof typeof InvitationStatusEnum];


/**
 * Check if a given object implements the Invitation interface.
 */
export function instanceOfInvitation(value: object): value is Invitation {
    if (!('createdAt' in value) || value['createdAt'] === undefined) return false;
    if (!('email' in value) || value['email'] === undefined) return false;
    if (!('expiresAt' in value) || value['expiresAt'] === undefined) return false;
    if (!('id' in value) || value['id'] === undefined) return false;
    if (!('organizationId' in value) || value['organizationId'] === undefined) return false;
    if (!('roleId' in value) || value['roleId'] === undefined) return false;
    if (!('sendCount' in value) || value['sendCount'] === undefined) return false;
    if (!('status' in value) || value['status'] === undefined) return false;
    if (!('updatedAt' in value) || value['updatedAt'] === undefined) return false;
    return true;
}

export function InvitationFromJSON(json: any): Invitation {
    return InvitationFromJSONTyped(json, false);
}

export function InvitationFromJSONTyped(json: any, ignoreDiscriminator: boolean): Invitation {
    if (json == null) {
        return json;
    }
    return {
        
            ...json,
        '$schema': json['$schema'] == null ? undefined : json['$schema'],
        'acceptedAt': json['acceptedAt'] == null ? undefined : (new Date(json['acceptedAt'])),
        'acceptedBy': json['acceptedBy'] == null ? undefined : json['acceptedBy'],
        'createdAt': (new Date(json['createdAt'])),
        'customFields': json['customFields'] == null ? undefined : json['customFields'],
        'declinedAt': json['declinedAt'] == null ? undefined : (new Date(json['declinedAt'])),
        'email': json['email'],
        'expiresAt': (new Date(json['expiresAt'])),
        'id': json['id'],
        'invitedBy': json['invitedBy'] == null ? undefined : json['invitedBy'],
        'inviter': json['inviter'] == null ? undefined : UserSummaryFromJSON(json['inviter']),
        'lastSentAt': json['lastSentAt'] == null ? undefined : (new Date(json['lastSentAt'])),
        'message': json['message'] == null ? undefined : json['message'],
        'organization': json['organization'] == null ? undefined : OrganizationSummaryFromJSON(json['organization']),
        'organizationId': json['organizationId'],
        'redirectUrl': json['redirectUrl'] == null ? undefined : json['redirectUrl'],
        'role': json['role'] == null ? undefined : RoleSummaryFromJSON(json['role']),
        'roleId': json['roleId'],
        'sendCount': json['sendCount'],
        'status': json['status'],
        'token': json['token'] == null ? undefined : json['token'],
        'updatedAt': (new Date(json['updatedAt'])),
    };
}

export function InvitationToJSON(json: any): Invitation {
    return InvitationToJSONTyped(json, false);
}

export function InvitationToJSONTyped(value?: Omit<Invitation, '$schema'> | null, ignoreDiscriminator = false): any {
    if (value == null) {
        return value;
    }

    return {
        
            ...value,
        'acceptedAt': value['acceptedAt'] == null ? undefined : ((value['acceptedAt']).toISOString()),
        'acceptedBy': value['acceptedBy'],
        'createdAt': ((value['createdAt']).toISOString()),
        'customFields': value['customFields'],
        'declinedAt': value['declinedAt'] == null ? undefined : ((value['declinedAt']).toISOString()),
        'email': value['email'],
        'expiresAt': ((value['expiresAt']).toISOString()),
        'id': value['id'],
        'invitedBy': value['invitedBy'],
        'inviter': UserSummaryToJSON(value['inviter']),
        'lastSentAt': value['lastSentAt'] == null ? undefined : ((value['lastSentAt']).toISOString()),
        'message': value['message'],
        'organization': OrganizationSummaryToJSON(value['organization']),
        'organizationId': value['organizationId'],
        'redirectUrl': value['redirectUrl'],
        'role': RoleSummaryToJSON(value['role']),
        'roleId': value['roleId'],
        'sendCount': value['sendCount'],
        'status': value['status'],
        'token': value['token'],
        'updatedAt': ((value['updatedAt']).toISOString()),
    };
}

