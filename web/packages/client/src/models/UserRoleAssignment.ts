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
 * @interface UserRoleAssignment
 */
export interface UserRoleAssignment {
    /**
     * Whether assignment is active
     * @type {boolean}
     * @memberof UserRoleAssignment
     */
    active: boolean;
    /**
     * When role was assigned
     * @type {Date}
     * @memberof UserRoleAssignment
     */
    assignedAt: Date;
    /**
     * Who assigned this role
     * @type {string}
     * @memberof UserRoleAssignment
     */
    assignedBy?: string;
    /**
     * Context ID
     * @type {string}
     * @memberof UserRoleAssignment
     */
    contextId?: string;
    /**
     * Assignment context type
     * @type {string}
     * @memberof UserRoleAssignment
     */
    contextType: string;
    /**
     * Role display name
     * @type {string}
     * @memberof UserRoleAssignment
     */
    displayName: string;
    /**
     * When assignment expires
     * @type {Date}
     * @memberof UserRoleAssignment
     */
    expiresAt?: Date;
    /**
     * Assignment ID
     * @type {string}
     * @memberof UserRoleAssignment
     */
    id: string;
    /**
     * Role ID
     * @type {string}
     * @memberof UserRoleAssignment
     */
    roleId: string;
    /**
     * Role name
     * @type {string}
     * @memberof UserRoleAssignment
     */
    roleName: string;
}

/**
 * Check if a given object implements the UserRoleAssignment interface.
 */
export function instanceOfUserRoleAssignment(value: object): value is UserRoleAssignment {
    if (!('active' in value) || value['active'] === undefined) return false;
    if (!('assignedAt' in value) || value['assignedAt'] === undefined) return false;
    if (!('contextType' in value) || value['contextType'] === undefined) return false;
    if (!('displayName' in value) || value['displayName'] === undefined) return false;
    if (!('id' in value) || value['id'] === undefined) return false;
    if (!('roleId' in value) || value['roleId'] === undefined) return false;
    if (!('roleName' in value) || value['roleName'] === undefined) return false;
    return true;
}

export function UserRoleAssignmentFromJSON(json: any): UserRoleAssignment {
    return UserRoleAssignmentFromJSONTyped(json, false);
}

export function UserRoleAssignmentFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserRoleAssignment {
    if (json == null) {
        return json;
    }
    return {
        
        'active': json['active'],
        'assignedAt': (new Date(json['assignedAt'])),
        'assignedBy': json['assignedBy'] == null ? undefined : json['assignedBy'],
        'contextId': json['contextId'] == null ? undefined : json['contextId'],
        'contextType': json['contextType'],
        'displayName': json['displayName'],
        'expiresAt': json['expiresAt'] == null ? undefined : (new Date(json['expiresAt'])),
        'id': json['id'],
        'roleId': json['roleId'],
        'roleName': json['roleName'],
    };
}

export function UserRoleAssignmentToJSON(json: any): UserRoleAssignment {
    return UserRoleAssignmentToJSONTyped(json, false);
}

export function UserRoleAssignmentToJSONTyped(value?: UserRoleAssignment | null, ignoreDiscriminator = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'active': value['active'],
        'assignedAt': ((value['assignedAt']).toISOString()),
        'assignedBy': value['assignedBy'],
        'contextId': value['contextId'],
        'contextType': value['contextType'],
        'displayName': value['displayName'],
        'expiresAt': value['expiresAt'] == null ? undefined : ((value['expiresAt']).toISOString()),
        'id': value['id'],
        'roleId': value['roleId'],
        'roleName': value['roleName'],
    };
}

