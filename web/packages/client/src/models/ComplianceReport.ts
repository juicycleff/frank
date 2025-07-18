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
import type { ComplianceFinding } from './ComplianceFinding';
import {
    ComplianceFindingFromJSON,
    ComplianceFindingFromJSONTyped,
    ComplianceFindingToJSON,
    ComplianceFindingToJSONTyped,
} from './ComplianceFinding';
import type { ComplianceControl } from './ComplianceControl';
import {
    ComplianceControlFromJSON,
    ComplianceControlFromJSONTyped,
    ComplianceControlToJSON,
    ComplianceControlToJSONTyped,
} from './ComplianceControl';
import type { ComplianceEvidence } from './ComplianceEvidence';
import {
    ComplianceEvidenceFromJSON,
    ComplianceEvidenceFromJSONTyped,
    ComplianceEvidenceToJSON,
    ComplianceEvidenceToJSONTyped,
} from './ComplianceEvidence';
import type { ComplianceRecommendation } from './ComplianceRecommendation';
import {
    ComplianceRecommendationFromJSON,
    ComplianceRecommendationFromJSONTyped,
    ComplianceRecommendationToJSON,
    ComplianceRecommendationToJSONTyped,
} from './ComplianceRecommendation';
import type { PlatformComplianceOverview } from './PlatformComplianceOverview';
import {
    PlatformComplianceOverviewFromJSON,
    PlatformComplianceOverviewFromJSONTyped,
    PlatformComplianceOverviewToJSON,
    PlatformComplianceOverviewToJSONTyped,
} from './PlatformComplianceOverview';
import type { ComplianceAttestation } from './ComplianceAttestation';
import {
    ComplianceAttestationFromJSON,
    ComplianceAttestationFromJSONTyped,
    ComplianceAttestationToJSON,
    ComplianceAttestationToJSONTyped,
} from './ComplianceAttestation';

/**
 * 
 * @export
 * @interface ComplianceReport
 */
export interface ComplianceReport {
    /**
     * A URL to the JSON Schema for this object.
     * @type {string}
     * @memberof ComplianceReport
     */
    readonly $schema?: string;
    /**
     * 
     * @type {ComplianceAttestation}
     * @memberof ComplianceReport
     */
    attestation?: ComplianceAttestation;
    /**
     * 
     * @type {Array<ComplianceControl>}
     * @memberof ComplianceReport
     */
    controls: Array<ComplianceControl> | null;
    /**
     * 
     * @type {Array<ComplianceEvidence>}
     * @memberof ComplianceReport
     */
    evidence: Array<ComplianceEvidence> | null;
    /**
     * 
     * @type {Array<ComplianceFinding>}
     * @memberof ComplianceReport
     */
    findings: Array<ComplianceFinding> | null;
    /**
     * 
     * @type {string}
     * @memberof ComplianceReport
     */
    framework: string;
    /**
     * 
     * @type {Date}
     * @memberof ComplianceReport
     */
    generatedAt: Date;
    /**
     * 
     * @type {PlatformComplianceOverview}
     * @memberof ComplianceReport
     */
    overview: PlatformComplianceOverview;
    /**
     * 
     * @type {string}
     * @memberof ComplianceReport
     */
    period: string;
    /**
     * 
     * @type {Array<ComplianceRecommendation>}
     * @memberof ComplianceReport
     */
    recommendations: Array<ComplianceRecommendation> | null;
    /**
     * 
     * @type {string}
     * @memberof ComplianceReport
     */
    reportType: string;
    /**
     * 
     * @type {number}
     * @memberof ComplianceReport
     */
    score: number;
    /**
     * 
     * @type {string}
     * @memberof ComplianceReport
     */
    status: string;
}

/**
 * Check if a given object implements the ComplianceReport interface.
 */
export function instanceOfComplianceReport(value: object): value is ComplianceReport {
    if (!('controls' in value) || value['controls'] === undefined) return false;
    if (!('evidence' in value) || value['evidence'] === undefined) return false;
    if (!('findings' in value) || value['findings'] === undefined) return false;
    if (!('framework' in value) || value['framework'] === undefined) return false;
    if (!('generatedAt' in value) || value['generatedAt'] === undefined) return false;
    if (!('overview' in value) || value['overview'] === undefined) return false;
    if (!('period' in value) || value['period'] === undefined) return false;
    if (!('recommendations' in value) || value['recommendations'] === undefined) return false;
    if (!('reportType' in value) || value['reportType'] === undefined) return false;
    if (!('score' in value) || value['score'] === undefined) return false;
    if (!('status' in value) || value['status'] === undefined) return false;
    return true;
}

export function ComplianceReportFromJSON(json: any): ComplianceReport {
    return ComplianceReportFromJSONTyped(json, false);
}

export function ComplianceReportFromJSONTyped(json: any, ignoreDiscriminator: boolean): ComplianceReport {
    if (json == null) {
        return json;
    }
    return {
        
        '$schema': json['$schema'] == null ? undefined : json['$schema'],
        'attestation': json['attestation'] == null ? undefined : ComplianceAttestationFromJSON(json['attestation']),
        'controls': (json['controls'] == null ? null : (json['controls'] as Array<any>).map(ComplianceControlFromJSON)),
        'evidence': (json['evidence'] == null ? null : (json['evidence'] as Array<any>).map(ComplianceEvidenceFromJSON)),
        'findings': (json['findings'] == null ? null : (json['findings'] as Array<any>).map(ComplianceFindingFromJSON)),
        'framework': json['framework'],
        'generatedAt': (new Date(json['generated_at'])),
        'overview': PlatformComplianceOverviewFromJSON(json['overview']),
        'period': json['period'],
        'recommendations': (json['recommendations'] == null ? null : (json['recommendations'] as Array<any>).map(ComplianceRecommendationFromJSON)),
        'reportType': json['report_type'],
        'score': json['score'],
        'status': json['status'],
    };
}

export function ComplianceReportToJSON(json: any): ComplianceReport {
    return ComplianceReportToJSONTyped(json, false);
}

export function ComplianceReportToJSONTyped(value?: Omit<ComplianceReport, '$schema'> | null, ignoreDiscriminator = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'attestation': ComplianceAttestationToJSON(value['attestation']),
        'controls': (value['controls'] == null ? null : (value['controls'] as Array<any>).map(ComplianceControlToJSON)),
        'evidence': (value['evidence'] == null ? null : (value['evidence'] as Array<any>).map(ComplianceEvidenceToJSON)),
        'findings': (value['findings'] == null ? null : (value['findings'] as Array<any>).map(ComplianceFindingToJSON)),
        'framework': value['framework'],
        'generated_at': ((value['generatedAt']).toISOString()),
        'overview': PlatformComplianceOverviewToJSON(value['overview']),
        'period': value['period'],
        'recommendations': (value['recommendations'] == null ? null : (value['recommendations'] as Array<any>).map(ComplianceRecommendationToJSON)),
        'report_type': value['reportType'],
        'score': value['score'],
        'status': value['status'],
    };
}

