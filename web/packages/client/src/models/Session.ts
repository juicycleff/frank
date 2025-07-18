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
 * @interface Session
 */
export interface Session {
    [key: string]: any | any;
    /**
     * A URL to the JSON Schema for this object.
     * @type {string}
     * @memberof Session
     */
    readonly $schema?: string;
    /**
     * Whether session is active
     * @type {boolean}
     * @memberof Session
     */
    active: boolean;
    /**
     * 
     * @type {Date}
     * @memberof Session
     */
    createdAt: Date;
    /**
     * Device ID
     * @type {string}
     * @memberof Session
     */
    deviceId?: string;
    /**
     * Session expiration time
     * @type {Date}
     * @memberof Session
     */
    expiresAt: Date;
    /**
     * 
     * @type {string}
     * @memberof Session
     */
    id: string;
    /**
     * IP address
     * @type {string}
     * @memberof Session
     */
    ipAddress?: string;
    /**
     * Last activity time
     * @type {Date}
     * @memberof Session
     */
    lastActiveAt: Date;
    /**
     * Location
     * @type {string}
     * @memberof Session
     */
    location?: string;
    /**
     * Additional session metadata
     * @type {object}
     * @memberof Session
     */
    metadata?: object;
    /**
     * Organization ID
     * @type {string}
     * @memberof Session
     */
    organizationId?: string;
    /**
     * Session token
     * @type {string}
     * @memberof Session
     */
    token?: string;
    /**
     * 
     * @type {Date}
     * @memberof Session
     */
    updatedAt: Date;
    /**
     * User agent
     * @type {string}
     * @memberof Session
     */
    userAgent?: string;
    /**
     * User ID
     * @type {string}
     * @memberof Session
     */
    userId: string;
}

/**
 * Check if a given object implements the Session interface.
 */
export function instanceOfSession(value: object): value is Session {
    if (!('active' in value) || value['active'] === undefined) return false;
    if (!('createdAt' in value) || value['createdAt'] === undefined) return false;
    if (!('expiresAt' in value) || value['expiresAt'] === undefined) return false;
    if (!('id' in value) || value['id'] === undefined) return false;
    if (!('lastActiveAt' in value) || value['lastActiveAt'] === undefined) return false;
    if (!('updatedAt' in value) || value['updatedAt'] === undefined) return false;
    if (!('userId' in value) || value['userId'] === undefined) return false;
    return true;
}

export function SessionFromJSON(json: any): Session {
    return SessionFromJSONTyped(json, false);
}

export function SessionFromJSONTyped(json: any, ignoreDiscriminator: boolean): Session {
    if (json == null) {
        return json;
    }
    return {
        
            ...json,
        '$schema': json['$schema'] == null ? undefined : json['$schema'],
        'active': json['active'],
        'createdAt': (new Date(json['createdAt'])),
        'deviceId': json['deviceId'] == null ? undefined : json['deviceId'],
        'expiresAt': (new Date(json['expiresAt'])),
        'id': json['id'],
        'ipAddress': json['ipAddress'] == null ? undefined : json['ipAddress'],
        'lastActiveAt': (new Date(json['lastActiveAt'])),
        'location': json['location'] == null ? undefined : json['location'],
        'metadata': json['metadata'] == null ? undefined : json['metadata'],
        'organizationId': json['organizationId'] == null ? undefined : json['organizationId'],
        'token': json['token'] == null ? undefined : json['token'],
        'updatedAt': (new Date(json['updatedAt'])),
        'userAgent': json['userAgent'] == null ? undefined : json['userAgent'],
        'userId': json['userId'],
    };
}

export function SessionToJSON(json: any): Session {
    return SessionToJSONTyped(json, false);
}

export function SessionToJSONTyped(value?: Omit<Session, '$schema'> | null, ignoreDiscriminator = false): any {
    if (value == null) {
        return value;
    }

    return {
        
            ...value,
        'active': value['active'],
        'createdAt': ((value['createdAt']).toISOString()),
        'deviceId': value['deviceId'],
        'expiresAt': ((value['expiresAt']).toISOString()),
        'id': value['id'],
        'ipAddress': value['ipAddress'],
        'lastActiveAt': ((value['lastActiveAt']).toISOString()),
        'location': value['location'],
        'metadata': value['metadata'],
        'organizationId': value['organizationId'],
        'token': value['token'],
        'updatedAt': ((value['updatedAt']).toISOString()),
        'userAgent': value['userAgent'],
        'userId': value['userId'],
    };
}

