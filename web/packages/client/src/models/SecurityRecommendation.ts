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
 * @interface SecurityRecommendation
 */
export interface SecurityRecommendation {
    /**
     * 
     * @type {string}
     * @memberof SecurityRecommendation
     */
    category: string;
    /**
     * 
     * @type {string}
     * @memberof SecurityRecommendation
     */
    description: string;
    /**
     * 
     * @type {Date}
     * @memberof SecurityRecommendation
     */
    dueDate?: Date;
    /**
     * 
     * @type {string}
     * @memberof SecurityRecommendation
     */
    effort: string;
    /**
     * 
     * @type {string}
     * @memberof SecurityRecommendation
     */
    id: string;
    /**
     * 
     * @type {string}
     * @memberof SecurityRecommendation
     */
    impact: string;
    /**
     * 
     * @type {string}
     * @memberof SecurityRecommendation
     */
    priority: string;
    /**
     * 
     * @type {string}
     * @memberof SecurityRecommendation
     */
    status: string;
    /**
     * 
     * @type {string}
     * @memberof SecurityRecommendation
     */
    title: string;
}

/**
 * Check if a given object implements the SecurityRecommendation interface.
 */
export function instanceOfSecurityRecommendation(value: object): value is SecurityRecommendation {
    if (!('category' in value) || value['category'] === undefined) return false;
    if (!('description' in value) || value['description'] === undefined) return false;
    if (!('effort' in value) || value['effort'] === undefined) return false;
    if (!('id' in value) || value['id'] === undefined) return false;
    if (!('impact' in value) || value['impact'] === undefined) return false;
    if (!('priority' in value) || value['priority'] === undefined) return false;
    if (!('status' in value) || value['status'] === undefined) return false;
    if (!('title' in value) || value['title'] === undefined) return false;
    return true;
}

export function SecurityRecommendationFromJSON(json: any): SecurityRecommendation {
    return SecurityRecommendationFromJSONTyped(json, false);
}

export function SecurityRecommendationFromJSONTyped(json: any, ignoreDiscriminator: boolean): SecurityRecommendation {
    if (json == null) {
        return json;
    }
    return {
        
        'category': json['category'],
        'description': json['description'],
        'dueDate': json['due_date'] == null ? undefined : (new Date(json['due_date'])),
        'effort': json['effort'],
        'id': json['id'],
        'impact': json['impact'],
        'priority': json['priority'],
        'status': json['status'],
        'title': json['title'],
    };
}

export function SecurityRecommendationToJSON(json: any): SecurityRecommendation {
    return SecurityRecommendationToJSONTyped(json, false);
}

export function SecurityRecommendationToJSONTyped(value?: SecurityRecommendation | null, ignoreDiscriminator = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'category': value['category'],
        'description': value['description'],
        'due_date': value['dueDate'] == null ? undefined : ((value['dueDate']).toISOString()),
        'effort': value['effort'],
        'id': value['id'],
        'impact': value['impact'],
        'priority': value['priority'],
        'status': value['status'],
        'title': value['title'],
    };
}

