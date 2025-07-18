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
 * @interface CreateInvitationRequest
 */
export interface CreateInvitationRequest {
    [key: string]: any | any;
    /**
     * A URL to the JSON Schema for this object.
     * @type {string}
     * @memberof CreateInvitationRequest
     */
    readonly $schema?: string;
    /**
     * Custom invitation fields
     * @type {object}
     * @memberof CreateInvitationRequest
     */
    customFields?: object;
    /**
     * Email address to invite
     * @type {string}
     * @memberof CreateInvitationRequest
     */
    email: string;
    /**
     * Custom expiration time
     * @type {Date}
     * @memberof CreateInvitationRequest
     */
    expiresAt?: Date;
    /**
     * Personal invitation message
     * @type {string}
     * @memberof CreateInvitationRequest
     */
    message?: string;
    /**
     * Post-acceptance redirect URL
     * @type {string}
     * @memberof CreateInvitationRequest
     */
    redirectUrl?: string;
    /**
     * Role ID to assign
     * @type {string}
     * @memberof CreateInvitationRequest
     */
    roleId: string;
    /**
     * Whether to send invitation email
     * @type {boolean}
     * @memberof CreateInvitationRequest
     */
    sendEmail: boolean;
}

/**
 * Check if a given object implements the CreateInvitationRequest interface.
 */
export function instanceOfCreateInvitationRequest(value: object): value is CreateInvitationRequest {
    if (!('email' in value) || value['email'] === undefined) return false;
    if (!('roleId' in value) || value['roleId'] === undefined) return false;
    if (!('sendEmail' in value) || value['sendEmail'] === undefined) return false;
    return true;
}

export function CreateInvitationRequestFromJSON(json: any): CreateInvitationRequest {
    return CreateInvitationRequestFromJSONTyped(json, false);
}

export function CreateInvitationRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): CreateInvitationRequest {
    if (json == null) {
        return json;
    }
    return {
        
            ...json,
        '$schema': json['$schema'] == null ? undefined : json['$schema'],
        'customFields': json['customFields'] == null ? undefined : json['customFields'],
        'email': json['email'],
        'expiresAt': json['expiresAt'] == null ? undefined : (new Date(json['expiresAt'])),
        'message': json['message'] == null ? undefined : json['message'],
        'redirectUrl': json['redirectUrl'] == null ? undefined : json['redirectUrl'],
        'roleId': json['roleId'],
        'sendEmail': json['sendEmail'],
    };
}

export function CreateInvitationRequestToJSON(json: any): CreateInvitationRequest {
    return CreateInvitationRequestToJSONTyped(json, false);
}

export function CreateInvitationRequestToJSONTyped(value?: Omit<CreateInvitationRequest, '$schema'> | null, ignoreDiscriminator = false): any {
    if (value == null) {
        return value;
    }

    return {
        
            ...value,
        'customFields': value['customFields'],
        'email': value['email'],
        'expiresAt': value['expiresAt'] == null ? undefined : ((value['expiresAt']).toISOString()),
        'message': value['message'],
        'redirectUrl': value['redirectUrl'],
        'roleId': value['roleId'],
        'sendEmail': value['sendEmail'],
    };
}

