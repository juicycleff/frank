/**
 * Generated by orval v7.7.0 🍺
 * Do not edit manually.
 * Frank Authentication Server
 * A comprehensive authentication server with OAuth2, MFA, Passkeys, SSO, and more
 * OpenAPI spec version: 1.0.0
 */
import type { OrganizationResponseMetadata } from "./organizationResponseMetadata";

/**
 * Organization information
 */
export interface OrganizationResponse {
  /** Whether organization is active */
  active: boolean;
  /** Creation timestamp */
  created_at: string;
  /** Organization domain */
  domain?: string;
  /** Organization ID */
  id: string;
  /** Organization logo URL */
  logo_url?: string;
  /** Organization metadata */
  metadata?: OrganizationResponseMetadata;
  /** Organization name */
  name: string;
  /** Organization plan */
  plan?: string;
  /** Organization slug */
  slug: string;
  /** Trial end date */
  trial_ends_at?: string;
  /** Whether trial has been used */
  trial_used?: boolean;
  /** Last update timestamp */
  updated_at: string;
}
