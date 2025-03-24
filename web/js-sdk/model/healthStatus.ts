/**
 * Generated by orval v7.7.0 🍺
 * Do not edit manually.
 * Frank Authentication Server
 * A comprehensive authentication server with OAuth2, MFA, Passkeys, SSO, and more
 * OpenAPI spec version: 1.0.0
 */
import type { HealthStatusStatus } from "./healthStatusStatus";

/**
 * Service health status
 */
export interface HealthStatus {
  /** Additional message */
  message?: string;
  /** Service name */
  service: string;
  /** Service status */
  status: HealthStatusStatus;
}
