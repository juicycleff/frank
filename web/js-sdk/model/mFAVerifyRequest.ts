/**
 * Generated by orval v7.7.0 🍺
 * Do not edit manually.
 * Frank Authentication Server
 * A comprehensive authentication server with OAuth2, MFA, Passkeys, SSO, and more
 * OpenAPI spec version: 1.0.0
 */
import type { MFAVerifyRequestMethod } from "./mFAVerifyRequestMethod";

/**
 * MFA verification request
 */
export interface MFAVerifyRequest {
  /** Verification code */
  code: string;
  /** MFA method to verify */
  method: MFAVerifyRequestMethod;
  /** Phone number for SMS verification */
  phone_number?: string;
}
