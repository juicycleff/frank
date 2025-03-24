/**
 * Generated by orval v7.7.0 🍺
 * Do not edit manually.
 * Frank Authentication Server
 * A comprehensive authentication server with OAuth2, MFA, Passkeys, SSO, and more
 * OpenAPI spec version: 1.0.0
 */
import type { APIKeyResponseMetadata } from "./aPIKeyResponseMetadata";

/**
 * API key information without the actual key
 */
export interface APIKeyResponse {
  /** Whether API key is active */
  active: boolean;
  /** Creation timestamp */
  created_at: string;
  /** Expiry timestamp */
  expires_at?: string;
  /** API key ID */
  id: string;
  /** Last used timestamp */
  last_used?: string;
  /** Key metadata */
  metadata?: APIKeyResponseMetadata;
  /** API key name */
  name: string;
  /** Organization ID */
  organization_id?: string;
  /** Key permissions */
  permissions?: string[];
  /** Key scopes */
  scopes?: string[];
  /** API key type (client/server) */
  type: string;
  /** Last update timestamp */
  updated_at?: string;
  /** User ID who owns the key */
  user_id?: string;
}
