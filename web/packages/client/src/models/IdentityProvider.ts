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
import type { SSOProviderConfig } from './SSOProviderConfig';
import {
    SSOProviderConfigFromJSON,
    SSOProviderConfigFromJSONTyped,
    SSOProviderConfigToJSON,
    SSOProviderConfigToJSONTyped,
} from './SSOProviderConfig';
import type { SSSProviderStats } from './SSSProviderStats';
import {
    SSSProviderStatsFromJSON,
    SSSProviderStatsFromJSONTyped,
    SSSProviderStatsToJSON,
    SSSProviderStatsToJSONTyped,
} from './SSSProviderStats';

/**
 * 
 * @export
 * @interface IdentityProvider
 */
export interface IdentityProvider {
    /**
     * A URL to the JSON Schema for this object.
     * @type {string}
     * @memberof IdentityProvider
     */
    readonly $schema?: string;
    /**
     * Whether provider is active
     * @type {boolean}
     * @memberof IdentityProvider
     */
    active: boolean;
    /**
     * Attribute mapping configuration
     * @type {{ [key: string]: string; }}
     * @memberof IdentityProvider
     */
    attributeMapping?: { [key: string]: string; };
    /**
     * Whether to auto-create users
     * @type {boolean}
     * @memberof IdentityProvider
     */
    autoProvision: boolean;
    /**
     * Login button text
     * @type {string}
     * @memberof IdentityProvider
     */
    buttonText?: string;
    /**
     * Provider-specific configuration
     * @type {SSOProviderConfig}
     * @memberof IdentityProvider
     */
    config?: SSOProviderConfig;
    /**
     * 
     * @type {Date}
     * @memberof IdentityProvider
     */
    createdAt: Date;
    /**
     * 
     * @type {string}
     * @memberof IdentityProvider
     */
    createdBy: string;
    /**
     * Default role for new users
     * @type {string}
     * @memberof IdentityProvider
     */
    defaultRole?: string;
    /**
     * Email domain for auto-provisioning
     * @type {string}
     * @memberof IdentityProvider
     */
    domain?: string;
    /**
     * Whether provider is enabled
     * @type {boolean}
     * @memberof IdentityProvider
     */
    enabled: boolean;
    /**
     * Provider icon URL
     * @type {string}
     * @memberof IdentityProvider
     */
    iconUrl?: string;
    /**
     * 
     * @type {string}
     * @memberof IdentityProvider
     */
    id: string;
    /**
     * Identity provider name
     * @type {string}
     * @memberof IdentityProvider
     */
    name: string;
    /**
     * Organization information
     * @type {OrganizationSummary}
     * @memberof IdentityProvider
     */
    organization?: OrganizationSummary;
    /**
     * Organization ID
     * @type {string}
     * @memberof IdentityProvider
     */
    organizationId: string;
    /**
     * Authentication protocol
     * @type {string}
     * @memberof IdentityProvider
     */
    protocol: string;
    /**
     * Usage statistics
     * @type {SSSProviderStats}
     * @memberof IdentityProvider
     */
    stats?: SSSProviderStats;
    /**
     * Provider type (oidc, saml, oauth2)
     * @type {string}
     * @memberof IdentityProvider
     */
    type: string;
    /**
     * 
     * @type {Date}
     * @memberof IdentityProvider
     */
    updatedAt: Date;
    /**
     * 
     * @type {string}
     * @memberof IdentityProvider
     */
    updatedBy: string;
}

/**
 * Check if a given object implements the IdentityProvider interface.
 */
export function instanceOfIdentityProvider(value: object): value is IdentityProvider {
    if (!('active' in value) || value['active'] === undefined) return false;
    if (!('autoProvision' in value) || value['autoProvision'] === undefined) return false;
    if (!('createdAt' in value) || value['createdAt'] === undefined) return false;
    if (!('createdBy' in value) || value['createdBy'] === undefined) return false;
    if (!('enabled' in value) || value['enabled'] === undefined) return false;
    if (!('id' in value) || value['id'] === undefined) return false;
    if (!('name' in value) || value['name'] === undefined) return false;
    if (!('organizationId' in value) || value['organizationId'] === undefined) return false;
    if (!('protocol' in value) || value['protocol'] === undefined) return false;
    if (!('type' in value) || value['type'] === undefined) return false;
    if (!('updatedAt' in value) || value['updatedAt'] === undefined) return false;
    if (!('updatedBy' in value) || value['updatedBy'] === undefined) return false;
    return true;
}

export function IdentityProviderFromJSON(json: any): IdentityProvider {
    return IdentityProviderFromJSONTyped(json, false);
}

export function IdentityProviderFromJSONTyped(json: any, ignoreDiscriminator: boolean): IdentityProvider {
    if (json == null) {
        return json;
    }
    return {
        
        '$schema': json['$schema'] == null ? undefined : json['$schema'],
        'active': json['active'],
        'attributeMapping': json['attributeMapping'] == null ? undefined : json['attributeMapping'],
        'autoProvision': json['autoProvision'],
        'buttonText': json['buttonText'] == null ? undefined : json['buttonText'],
        'config': json['config'] == null ? undefined : SSOProviderConfigFromJSON(json['config']),
        'createdAt': (new Date(json['createdAt'])),
        'createdBy': json['createdBy'],
        'defaultRole': json['defaultRole'] == null ? undefined : json['defaultRole'],
        'domain': json['domain'] == null ? undefined : json['domain'],
        'enabled': json['enabled'],
        'iconUrl': json['iconUrl'] == null ? undefined : json['iconUrl'],
        'id': json['id'],
        'name': json['name'],
        'organization': json['organization'] == null ? undefined : OrganizationSummaryFromJSON(json['organization']),
        'organizationId': json['organizationId'],
        'protocol': json['protocol'],
        'stats': json['stats'] == null ? undefined : SSSProviderStatsFromJSON(json['stats']),
        'type': json['type'],
        'updatedAt': (new Date(json['updatedAt'])),
        'updatedBy': json['updatedBy'],
    };
}

export function IdentityProviderToJSON(json: any): IdentityProvider {
    return IdentityProviderToJSONTyped(json, false);
}

export function IdentityProviderToJSONTyped(value?: Omit<IdentityProvider, '$schema'> | null, ignoreDiscriminator = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'active': value['active'],
        'attributeMapping': value['attributeMapping'],
        'autoProvision': value['autoProvision'],
        'buttonText': value['buttonText'],
        'config': SSOProviderConfigToJSON(value['config']),
        'createdAt': ((value['createdAt']).toISOString()),
        'createdBy': value['createdBy'],
        'defaultRole': value['defaultRole'],
        'domain': value['domain'],
        'enabled': value['enabled'],
        'iconUrl': value['iconUrl'],
        'id': value['id'],
        'name': value['name'],
        'organization': OrganizationSummaryToJSON(value['organization']),
        'organizationId': value['organizationId'],
        'protocol': value['protocol'],
        'stats': SSSProviderStatsToJSON(value['stats']),
        'type': value['type'],
        'updatedAt': ((value['updatedAt']).toISOString()),
        'updatedBy': value['updatedBy'],
    };
}

