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
 * @interface WebhookStats
 */
export interface WebhookStats {
    /**
     * A URL to the JSON Schema for this object.
     * @type {string}
     * @memberof WebhookStats
     */
    readonly $schema?: string;
    /**
     * Average response time in milliseconds
     * @type {number}
     * @memberof WebhookStats
     */
    averageResponseTime: number;
    /**
     * Errors by status code
     * @type {{ [key: string]: number; }}
     * @memberof WebhookStats
     */
    errorsByCode: { [key: string]: number; };
    /**
     * Events by type
     * @type {{ [key: string]: number; }}
     * @memberof WebhookStats
     */
    eventsByType: { [key: string]: number; };
    /**
     * Events sent this month
     * @type {number}
     * @memberof WebhookStats
     */
    eventsMonth: number;
    /**
     * Events sent today
     * @type {number}
     * @memberof WebhookStats
     */
    eventsToday: number;
    /**
     * Events sent this week
     * @type {number}
     * @memberof WebhookStats
     */
    eventsWeek: number;
    /**
     * Failed deliveries
     * @type {number}
     * @memberof WebhookStats
     */
    failedEvents: number;
    /**
     * Last successful delivery
     * @type {Date}
     * @memberof WebhookStats
     */
    lastDelivery?: Date;
    /**
     * Last failed delivery
     * @type {Date}
     * @memberof WebhookStats
     */
    lastFailure?: Date;
    /**
     * Recent response times
     * @type {Array<number>}
     * @memberof WebhookStats
     */
    responseTimes?: Array<number> | null;
    /**
     * Success rate percentage
     * @type {number}
     * @memberof WebhookStats
     */
    successRate: number;
    /**
     * Successful deliveries
     * @type {number}
     * @memberof WebhookStats
     */
    successfulEvents: number;
    /**
     * Total events sent
     * @type {number}
     * @memberof WebhookStats
     */
    totalEvents: number;
    /**
     * Webhook ID
     * @type {string}
     * @memberof WebhookStats
     */
    webhookId: string;
}

/**
 * Check if a given object implements the WebhookStats interface.
 */
export function instanceOfWebhookStats(value: object): value is WebhookStats {
    if (!('averageResponseTime' in value) || value['averageResponseTime'] === undefined) return false;
    if (!('errorsByCode' in value) || value['errorsByCode'] === undefined) return false;
    if (!('eventsByType' in value) || value['eventsByType'] === undefined) return false;
    if (!('eventsMonth' in value) || value['eventsMonth'] === undefined) return false;
    if (!('eventsToday' in value) || value['eventsToday'] === undefined) return false;
    if (!('eventsWeek' in value) || value['eventsWeek'] === undefined) return false;
    if (!('failedEvents' in value) || value['failedEvents'] === undefined) return false;
    if (!('successRate' in value) || value['successRate'] === undefined) return false;
    if (!('successfulEvents' in value) || value['successfulEvents'] === undefined) return false;
    if (!('totalEvents' in value) || value['totalEvents'] === undefined) return false;
    if (!('webhookId' in value) || value['webhookId'] === undefined) return false;
    return true;
}

export function WebhookStatsFromJSON(json: any): WebhookStats {
    return WebhookStatsFromJSONTyped(json, false);
}

export function WebhookStatsFromJSONTyped(json: any, ignoreDiscriminator: boolean): WebhookStats {
    if (json == null) {
        return json;
    }
    return {
        
        '$schema': json['$schema'] == null ? undefined : json['$schema'],
        'averageResponseTime': json['averageResponseTime'],
        'errorsByCode': json['errorsByCode'],
        'eventsByType': json['eventsByType'],
        'eventsMonth': json['eventsMonth'],
        'eventsToday': json['eventsToday'],
        'eventsWeek': json['eventsWeek'],
        'failedEvents': json['failedEvents'],
        'lastDelivery': json['lastDelivery'] == null ? undefined : (new Date(json['lastDelivery'])),
        'lastFailure': json['lastFailure'] == null ? undefined : (new Date(json['lastFailure'])),
        'responseTimes': json['responseTimes'] == null ? undefined : json['responseTimes'],
        'successRate': json['successRate'],
        'successfulEvents': json['successfulEvents'],
        'totalEvents': json['totalEvents'],
        'webhookId': json['webhookId'],
    };
}

export function WebhookStatsToJSON(json: any): WebhookStats {
    return WebhookStatsToJSONTyped(json, false);
}

export function WebhookStatsToJSONTyped(value?: Omit<WebhookStats, '$schema'> | null, ignoreDiscriminator = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'averageResponseTime': value['averageResponseTime'],
        'errorsByCode': value['errorsByCode'],
        'eventsByType': value['eventsByType'],
        'eventsMonth': value['eventsMonth'],
        'eventsToday': value['eventsToday'],
        'eventsWeek': value['eventsWeek'],
        'failedEvents': value['failedEvents'],
        'lastDelivery': value['lastDelivery'] == null ? undefined : ((value['lastDelivery']).toISOString()),
        'lastFailure': value['lastFailure'] == null ? undefined : ((value['lastFailure']).toISOString()),
        'responseTimes': value['responseTimes'],
        'successRate': value['successRate'],
        'successfulEvents': value['successfulEvents'],
        'totalEvents': value['totalEvents'],
        'webhookId': value['webhookId'],
    };
}

