/**
 * Generated by orval v7.7.0 🍺
 * Do not edit manually.
 * Frank Authentication Server
 * A comprehensive authentication server with OAuth2, MFA, Passkeys, SSO, and more
 * OpenAPI spec version: 1.0.0
 */

export interface MetricsResponseBody {
  /** Total error count */
  errors?: number;
  /** Number of goroutines */
  goroutines: number;
  /** Memory usage in bytes */
  memory_usage: number;
  /** Requests per second */
  request_rate?: number;
  /** Total request count */
  requests?: number;
  /** System uptime in seconds */
  uptime: number;
}
