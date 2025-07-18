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
import type { ContextType } from './ContextType';
import {
    ContextTypeFromJSON,
    ContextTypeFromJSONTyped,
    ContextTypeToJSON,
    ContextTypeToJSONTyped,
} from './ContextType';
import type { UserType } from './UserType';
import {
    UserTypeFromJSON,
    UserTypeFromJSONTyped,
    UserTypeToJSON,
    UserTypeToJSONTyped,
} from './UserType';

/**
 * 
 * @export
 * @interface UpdateRoleRequest
 */
export interface UpdateRoleRequest {
    /**
     * A URL to the JSON Schema for this object.
     * @type {string}
     * @memberof UpdateRoleRequest
     */
    readonly $schema?: string;
    /**
     * Updated active status
     * @type {boolean}
     * @memberof UpdateRoleRequest
     */
    active?: boolean;
    /**
     * 
     * @type {Array<ContextType>}
     * @memberof UpdateRoleRequest
     */
    applicableContexts?: Array<ContextType> | null;
    /**
     * Updated applicable user types
     * @type {Array<UserType>}
     * @memberof UpdateRoleRequest
     */
    applicableUserTypes?: Array<UserType> | null;
    /**
     * Updated color
     * @type {string}
     * @memberof UpdateRoleRequest
     */
    color?: string;
    /**
     * 
     * @type {string}
     * @memberof UpdateRoleRequest
     */
    conditions?: string;
    /**
     * 
     * @type {boolean}
     * @memberof UpdateRoleRequest
     */
    dangerous?: boolean;
    /**
     * Updated description
     * @type {string}
     * @memberof UpdateRoleRequest
     */
    description?: string;
    /**
     * Updated display name
     * @type {string}
     * @memberof UpdateRoleRequest
     */
    displayName?: string;
    /**
     * Updated default status
     * @type {boolean}
     * @memberof UpdateRoleRequest
     */
    isDefault?: boolean;
    /**
     * Updated name
     * @type {string}
     * @memberof UpdateRoleRequest
     */
    name?: string;
    /**
     * Updated parent role ID
     * @type {string}
     * @memberof UpdateRoleRequest
     */
    parentId?: string;
    /**
     * 
     * @type {string}
     * @memberof UpdateRoleRequest
     */
    permissionGroup?: string;
    /**
     * Updated priority
     * @type {number}
     * @memberof UpdateRoleRequest
     */
    priority?: number;
    /**
     * 
     * @type {number}
     * @memberof UpdateRoleRequest
     */
    riskLevel?: number;
}

/**
 * Check if a given object implements the UpdateRoleRequest interface.
 */
export function instanceOfUpdateRoleRequest(value: object): value is UpdateRoleRequest {
    return true;
}

export function UpdateRoleRequestFromJSON(json: any): UpdateRoleRequest {
    return UpdateRoleRequestFromJSONTyped(json, false);
}

export function UpdateRoleRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UpdateRoleRequest {
    if (json == null) {
        return json;
    }
    return {
        
        '$schema': json['$schema'] == null ? undefined : json['$schema'],
        'active': json['active'] == null ? undefined : json['active'],
        'applicableContexts': json['applicableContexts'] == null ? undefined : ((json['applicableContexts'] as Array<any>).map(ContextTypeFromJSON)),
        'applicableUserTypes': json['applicableUserTypes'] == null ? undefined : ((json['applicableUserTypes'] as Array<any>).map(UserTypeFromJSON)),
        'color': json['color'] == null ? undefined : json['color'],
        'conditions': json['conditions'] == null ? undefined : json['conditions'],
        'dangerous': json['dangerous'] == null ? undefined : json['dangerous'],
        'description': json['description'] == null ? undefined : json['description'],
        'displayName': json['displayName'] == null ? undefined : json['displayName'],
        'isDefault': json['isDefault'] == null ? undefined : json['isDefault'],
        'name': json['name'] == null ? undefined : json['name'],
        'parentId': json['parentId'] == null ? undefined : json['parentId'],
        'permissionGroup': json['permissionGroup'] == null ? undefined : json['permissionGroup'],
        'priority': json['priority'] == null ? undefined : json['priority'],
        'riskLevel': json['riskLevel'] == null ? undefined : json['riskLevel'],
    };
}

export function UpdateRoleRequestToJSON(json: any): UpdateRoleRequest {
    return UpdateRoleRequestToJSONTyped(json, false);
}

export function UpdateRoleRequestToJSONTyped(value?: Omit<UpdateRoleRequest, '$schema'> | null, ignoreDiscriminator = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'active': value['active'],
        'applicableContexts': value['applicableContexts'] == null ? undefined : ((value['applicableContexts'] as Array<any>).map(ContextTypeToJSON)),
        'applicableUserTypes': value['applicableUserTypes'] == null ? undefined : ((value['applicableUserTypes'] as Array<any>).map(UserTypeToJSON)),
        'color': value['color'],
        'conditions': value['conditions'],
        'dangerous': value['dangerous'],
        'description': value['description'],
        'displayName': value['displayName'],
        'isDefault': value['isDefault'],
        'name': value['name'],
        'parentId': value['parentId'],
        'permissionGroup': value['permissionGroup'],
        'priority': value['priority'],
        'riskLevel': value['riskLevel'],
    };
}

