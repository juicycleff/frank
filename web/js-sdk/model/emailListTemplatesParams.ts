/**
 * Generated by orval v7.7.0 🍺
 * Do not edit manually.
 * Frank Authentication Server
 * A comprehensive authentication server with OAuth2, MFA, Passkeys, SSO, and more
 * OpenAPI spec version: 1.0.0
 */

export type EmailListTemplatesParams = {
  /**
   * Pagination offset
   */
  offset?: number;
  /**
   * Number of items to return
   */
  limit?: number;
  /**
   * Filter by template type
   */
  type?: string;
  /**
   * Filter by organization ID
   */
  organization_id?: string;
  /**
   * Filter by locale
   */
  locale?: string;
};
