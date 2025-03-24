/**
 * Generated by orval v7.7.0 🍺
 * Do not edit manually.
 * Frank Authentication Server
 * A comprehensive authentication server with OAuth2, MFA, Passkeys, SSO, and more
 * OpenAPI spec version: 1.0.0
 */
import type { WebhookEventResponseHeaders } from "./webhookEventResponseHeaders";

/**
 * Webhook event information
 */
export interface WebhookEventResponse {
  /** Number of delivery attempts */
  attempts: number;
  /** Creation timestamp */
  created_at: string;
  /** Whether event was delivered */
  delivered: boolean;
  /** Delivery timestamp */
  delivered_at?: string;
  /** Error from last attempt */
  error?: string;
  /** Event type */
  event_type: string;
  /** Event headers */
  headers?: WebhookEventResponseHeaders;
  /** Event ID */
  id: string;
  /** Next retry timestamp */
  next_retry?: string;
  /** Event payload */
  payload?: unknown;
  /** Response from last attempt */
  response_body?: string;
  /** HTTP status code from last attempt */
  status_code?: number;
  /** Last update timestamp */
  updated_at?: string;
  /** Webhook ID */
  webhook_id: string;
}
