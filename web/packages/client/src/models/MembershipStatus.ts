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


/**
 * 
 * @export
 */
export const MembershipStatus = {
    Pending: 'pending',
    Active: 'active',
    Inactive: 'inactive',
    Suspended: 'suspended'
} as const;
export type MembershipStatus = typeof MembershipStatus[keyof typeof MembershipStatus];


export function instanceOfMembershipStatus(value: any): boolean {
    for (const key in MembershipStatus) {
        if (Object.prototype.hasOwnProperty.call(MembershipStatus, key)) {
            if (MembershipStatus[key as keyof typeof MembershipStatus] === value) {
                return true;
            }
        }
    }
    return false;
}

export function MembershipStatusFromJSON(json: any): MembershipStatus {
    return MembershipStatusFromJSONTyped(json, false);
}

export function MembershipStatusFromJSONTyped(json: any, ignoreDiscriminator: boolean): MembershipStatus {
    return json as MembershipStatus;
}

export function MembershipStatusToJSON(value?: MembershipStatus | null): any {
    return value as any;
}

export function MembershipStatusToJSONTyped(value: any, ignoreDiscriminator: boolean): MembershipStatus {
    return value as MembershipStatus;
}

