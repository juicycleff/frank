/**
 * Generated by orval v7.7.0 🍺
 * Do not edit manually.
 * Frank Authentication Server
 * A comprehensive authentication server with OAuth2, MFA, Passkeys, SSO, and more
 * OpenAPI spec version: 1.0.0
 */

export interface UpdateScopeRequestBody {
  /** Whether this scope is included by default */
  default_scope?: boolean;
  /** Scope description */
  description?: string;
  /** Whether this scope can be requested by any client */
  public?: boolean;
}
