/**
 * Generated by orval v7.7.0 🍺
 * Do not edit manually.
 * Frank Authentication Server
 * A comprehensive authentication server with OAuth2, MFA, Passkeys, SSO, and more
 * OpenAPI spec version: 1.0.0
 */
import type { EmailTemplateResponseMetadata } from "./emailTemplateResponseMetadata";

/**
 * Email template information
 */
export interface EmailTemplateResponse {
  /** Whether template is active */
  active: boolean;
  /** Created At */
  created_at: string;
  /** HTML content */
  html_content: string;
  /** Created At */
  id: string;
  /** Template locale */
  locale: string;
  /** Template metadata */
  metadata?: EmailTemplateResponseMetadata;
  /** Template name */
  name: string;
  /** Organization ID */
  organization_id?: string;
  /** Email subject */
  subject: string;
  /** Whether this is a system template */
  system: boolean;
  /** Text content */
  text_content?: string;
  /** Template type */
  type: string;
  /** Updated At */
  updated_at: string;
}
