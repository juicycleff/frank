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
import type { ResourceMetric } from './ResourceMetric';
import {
    ResourceMetricFromJSON,
    ResourceMetricFromJSONTyped,
    ResourceMetricToJSON,
    ResourceMetricToJSONTyped,
} from './ResourceMetric';

/**
 * 
 * @export
 * @interface ResourceUtilizationMetrics
 */
export interface ResourceUtilizationMetrics {
    /**
     * 
     * @type {ResourceMetric}
     * @memberof ResourceUtilizationMetrics
     */
    cpu: ResourceMetric;
    /**
     * 
     * @type {ResourceMetric}
     * @memberof ResourceUtilizationMetrics
     */
    disk: ResourceMetric;
    /**
     * 
     * @type {ResourceMetric}
     * @memberof ResourceUtilizationMetrics
     */
    memory: ResourceMetric;
    /**
     * 
     * @type {ResourceMetric}
     * @memberof ResourceUtilizationMetrics
     */
    network: ResourceMetric;
}

/**
 * Check if a given object implements the ResourceUtilizationMetrics interface.
 */
export function instanceOfResourceUtilizationMetrics(value: object): value is ResourceUtilizationMetrics {
    if (!('cpu' in value) || value['cpu'] === undefined) return false;
    if (!('disk' in value) || value['disk'] === undefined) return false;
    if (!('memory' in value) || value['memory'] === undefined) return false;
    if (!('network' in value) || value['network'] === undefined) return false;
    return true;
}

export function ResourceUtilizationMetricsFromJSON(json: any): ResourceUtilizationMetrics {
    return ResourceUtilizationMetricsFromJSONTyped(json, false);
}

export function ResourceUtilizationMetricsFromJSONTyped(json: any, ignoreDiscriminator: boolean): ResourceUtilizationMetrics {
    if (json == null) {
        return json;
    }
    return {
        
        'cpu': ResourceMetricFromJSON(json['cpu']),
        'disk': ResourceMetricFromJSON(json['disk']),
        'memory': ResourceMetricFromJSON(json['memory']),
        'network': ResourceMetricFromJSON(json['network']),
    };
}

export function ResourceUtilizationMetricsToJSON(json: any): ResourceUtilizationMetrics {
    return ResourceUtilizationMetricsToJSONTyped(json, false);
}

export function ResourceUtilizationMetricsToJSONTyped(value?: ResourceUtilizationMetrics | null, ignoreDiscriminator = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'cpu': ResourceMetricToJSON(value['cpu']),
        'disk': ResourceMetricToJSON(value['disk']),
        'memory': ResourceMetricToJSON(value['memory']),
        'network': ResourceMetricToJSON(value['network']),
    };
}

