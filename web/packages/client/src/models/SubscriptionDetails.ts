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
import type { SubscriptionEvent } from './SubscriptionEvent';
import {
    SubscriptionEventFromJSON,
    SubscriptionEventFromJSONTyped,
    SubscriptionEventToJSON,
    SubscriptionEventToJSONTyped,
} from './SubscriptionEvent';
import type { SubscriptionSummary } from './SubscriptionSummary';
import {
    SubscriptionSummaryFromJSON,
    SubscriptionSummaryFromJSONTyped,
    SubscriptionSummaryToJSON,
    SubscriptionSummaryToJSONTyped,
} from './SubscriptionSummary';
import type { OrganizationSummary } from './OrganizationSummary';
import {
    OrganizationSummaryFromJSON,
    OrganizationSummaryFromJSONTyped,
    OrganizationSummaryToJSON,
    OrganizationSummaryToJSONTyped,
} from './OrganizationSummary';
import type { SubscriptionBilling } from './SubscriptionBilling';
import {
    SubscriptionBillingFromJSON,
    SubscriptionBillingFromJSONTyped,
    SubscriptionBillingToJSON,
    SubscriptionBillingToJSONTyped,
} from './SubscriptionBilling';
import type { PaymentRecord } from './PaymentRecord';
import {
    PaymentRecordFromJSON,
    PaymentRecordFromJSONTyped,
    PaymentRecordToJSON,
    PaymentRecordToJSONTyped,
} from './PaymentRecord';
import type { SubscriptionUsage } from './SubscriptionUsage';
import {
    SubscriptionUsageFromJSON,
    SubscriptionUsageFromJSONTyped,
    SubscriptionUsageToJSON,
    SubscriptionUsageToJSONTyped,
} from './SubscriptionUsage';
import type { InvoiceSummary } from './InvoiceSummary';
import {
    InvoiceSummaryFromJSON,
    InvoiceSummaryFromJSONTyped,
    InvoiceSummaryToJSON,
    InvoiceSummaryToJSONTyped,
} from './InvoiceSummary';

/**
 * 
 * @export
 * @interface SubscriptionDetails
 */
export interface SubscriptionDetails {
    /**
     * A URL to the JSON Schema for this object.
     * @type {string}
     * @memberof SubscriptionDetails
     */
    readonly $schema?: string;
    /**
     * 
     * @type {SubscriptionBilling}
     * @memberof SubscriptionDetails
     */
    billing: SubscriptionBilling;
    /**
     * 
     * @type {Array<SubscriptionEvent>}
     * @memberof SubscriptionDetails
     */
    events: Array<SubscriptionEvent> | null;
    /**
     * 
     * @type {Array<InvoiceSummary>}
     * @memberof SubscriptionDetails
     */
    invoices: Array<InvoiceSummary> | null;
    /**
     * 
     * @type {OrganizationSummary}
     * @memberof SubscriptionDetails
     */
    organization: OrganizationSummary;
    /**
     * 
     * @type {Array<PaymentRecord>}
     * @memberof SubscriptionDetails
     */
    paymentHistory: Array<PaymentRecord> | null;
    /**
     * 
     * @type {SubscriptionSummary}
     * @memberof SubscriptionDetails
     */
    subscription: SubscriptionSummary;
    /**
     * 
     * @type {SubscriptionUsage}
     * @memberof SubscriptionDetails
     */
    usage: SubscriptionUsage;
}

/**
 * Check if a given object implements the SubscriptionDetails interface.
 */
export function instanceOfSubscriptionDetails(value: object): value is SubscriptionDetails {
    if (!('billing' in value) || value['billing'] === undefined) return false;
    if (!('events' in value) || value['events'] === undefined) return false;
    if (!('invoices' in value) || value['invoices'] === undefined) return false;
    if (!('organization' in value) || value['organization'] === undefined) return false;
    if (!('paymentHistory' in value) || value['paymentHistory'] === undefined) return false;
    if (!('subscription' in value) || value['subscription'] === undefined) return false;
    if (!('usage' in value) || value['usage'] === undefined) return false;
    return true;
}

export function SubscriptionDetailsFromJSON(json: any): SubscriptionDetails {
    return SubscriptionDetailsFromJSONTyped(json, false);
}

export function SubscriptionDetailsFromJSONTyped(json: any, ignoreDiscriminator: boolean): SubscriptionDetails {
    if (json == null) {
        return json;
    }
    return {
        
        '$schema': json['$schema'] == null ? undefined : json['$schema'],
        'billing': SubscriptionBillingFromJSON(json['billing']),
        'events': (json['events'] == null ? null : (json['events'] as Array<any>).map(SubscriptionEventFromJSON)),
        'invoices': (json['invoices'] == null ? null : (json['invoices'] as Array<any>).map(InvoiceSummaryFromJSON)),
        'organization': OrganizationSummaryFromJSON(json['organization']),
        'paymentHistory': (json['payment_history'] == null ? null : (json['payment_history'] as Array<any>).map(PaymentRecordFromJSON)),
        'subscription': SubscriptionSummaryFromJSON(json['subscription']),
        'usage': SubscriptionUsageFromJSON(json['usage']),
    };
}

export function SubscriptionDetailsToJSON(json: any): SubscriptionDetails {
    return SubscriptionDetailsToJSONTyped(json, false);
}

export function SubscriptionDetailsToJSONTyped(value?: Omit<SubscriptionDetails, '$schema'> | null, ignoreDiscriminator = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'billing': SubscriptionBillingToJSON(value['billing']),
        'events': (value['events'] == null ? null : (value['events'] as Array<any>).map(SubscriptionEventToJSON)),
        'invoices': (value['invoices'] == null ? null : (value['invoices'] as Array<any>).map(InvoiceSummaryToJSON)),
        'organization': OrganizationSummaryToJSON(value['organization']),
        'payment_history': (value['paymentHistory'] == null ? null : (value['paymentHistory'] as Array<any>).map(PaymentRecordToJSON)),
        'subscription': SubscriptionSummaryToJSON(value['subscription']),
        'usage': SubscriptionUsageToJSON(value['usage']),
    };
}

