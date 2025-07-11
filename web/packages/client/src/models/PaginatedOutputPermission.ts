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
import type { Pagination } from './Pagination';
import {
    PaginationFromJSON,
    PaginationFromJSONTyped,
    PaginationToJSON,
    PaginationToJSONTyped,
} from './Pagination';
import type { Permission } from './Permission';
import {
    PermissionFromJSON,
    PermissionFromJSONTyped,
    PermissionToJSON,
    PermissionToJSONTyped,
} from './Permission';

/**
 * 
 * @export
 * @interface PaginatedOutputPermission
 */
export interface PaginatedOutputPermission {
    /**
     * A URL to the JSON Schema for this object.
     * @type {string}
     * @memberof PaginatedOutputPermission
     */
    readonly $schema?: string;
    /**
     * 
     * @type {Array<Permission>}
     * @memberof PaginatedOutputPermission
     */
    data: Array<Permission> | null;
    /**
     * 
     * @type {Pagination}
     * @memberof PaginatedOutputPermission
     */
    pagination: Pagination;
}

/**
 * Check if a given object implements the PaginatedOutputPermission interface.
 */
export function instanceOfPaginatedOutputPermission(value: object): value is PaginatedOutputPermission {
    if (!('data' in value) || value['data'] === undefined) return false;
    if (!('pagination' in value) || value['pagination'] === undefined) return false;
    return true;
}

export function PaginatedOutputPermissionFromJSON(json: any): PaginatedOutputPermission {
    return PaginatedOutputPermissionFromJSONTyped(json, false);
}

export function PaginatedOutputPermissionFromJSONTyped(json: any, ignoreDiscriminator: boolean): PaginatedOutputPermission {
    if (json == null) {
        return json;
    }
    return {
        
        '$schema': json['$schema'] == null ? undefined : json['$schema'],
        'data': (json['data'] == null ? null : (json['data'] as Array<any>).map(PermissionFromJSON)),
        'pagination': PaginationFromJSON(json['pagination']),
    };
}

export function PaginatedOutputPermissionToJSON(json: any): PaginatedOutputPermission {
    return PaginatedOutputPermissionToJSONTyped(json, false);
}

export function PaginatedOutputPermissionToJSONTyped(value?: Omit<PaginatedOutputPermission, '$schema'> | null, ignoreDiscriminator = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'data': (value['data'] == null ? null : (value['data'] as Array<any>).map(PermissionToJSON)),
        'pagination': PaginationToJSON(value['pagination']),
    };
}

