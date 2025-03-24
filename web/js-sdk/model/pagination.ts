/**
 * Generated by orval v7.7.0 🍺
 * Do not edit manually.
 * Frank Authentication Server
 * A comprehensive authentication server with OAuth2, MFA, Passkeys, SSO, and more
 * OpenAPI spec version: 1.0.0
 */

/**
 * Pagination parameters
 */
export interface Pagination {
  /** Current page number */
  current_page: number;
  /** Has next page */
  has_next: boolean;
  /** Has previous page */
  has_previous: boolean;
  /**
   * Limit
   * @minimum 1
   * @maximum 100
   */
  limit: number;
  /**
   * Offset
   * @minimum 0
   */
  offset: number;
  /** Total number of items */
  total: number;
  /** Total number of pages */
  total_pages: number;
}
