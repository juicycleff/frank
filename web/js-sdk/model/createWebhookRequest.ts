/**
 * Generated by orval v7.7.0 🍺
 * Do not edit manually.
 * Frank Authentication Server
 * A comprehensive authentication server with OAuth2, MFA, Passkeys, SSO, and more
 * OpenAPI spec version: 1.0.0
 */
import type { CreateWebhookRequestFormat } from "./createWebhookRequestFormat";
import type { CreateWebhookRequestMetadata } from "./createWebhookRequestMetadata";

/**
 * Create webhook request
 */
export interface CreateWebhookRequest {
  /** Event types to subscribe to */
  event_types: string[];
  /** Payload format */
  format?: CreateWebhookRequestFormat;
  /** Webhook metadata */
  metadata?: CreateWebhookRequestMetadata;
  /** Webhook name */
  name: string;
  /**
   * Number of retries on failure
   * @minimum 0
   * @maximum 10
   */
  retry_count?: number;
  /**
   * Timeout in milliseconds
   * @minimum 1000
   * @maximum 30000
   */
  timeout_ms?: number;
  /** Webhook URL */
  url: string;
}
