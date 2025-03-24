/**
 * Generated by orval v7.7.0 🍺
 * Do not edit manually.
 * Frank Authentication Server
 * A comprehensive authentication server with OAuth2, MFA, Passkeys, SSO, and more
 * OpenAPI spec version: 1.0.0
 */

/**
 * Refresh token response
 */
export interface RefreshTokenResponse {
  /** Token expiry timestamp */
  expires_at: number;
  /** New refresh token */
  refresh_token: string;
  /** New JWTAuth access token */
  token: string;
}
