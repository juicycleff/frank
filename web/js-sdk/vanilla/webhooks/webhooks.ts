/**
 * Generated by orval v7.7.0 🍺
 * Do not edit manually.
 * Frank Authentication Server
 * A comprehensive authentication server with OAuth2, MFA, Passkeys, SSO, and more
 * OpenAPI spec version: 1.0.0
 */
import type {
  BadRequestError,
  CreateRequestBody3,
  ForbiddenError,
  InternalServerError,
  ListEventsResponseBody,
  ListResponseBody5,
  NotFoundError,
  SendResponseBody,
  TriggerEventRequestBody,
  UnauthorizedError,
  UpdateRequestBody5,
  WebhookEventResponse,
  WebhookResponse,
  WebhookSecretResponse,
  WebhooksListEventsParams,
  WebhooksListParams,
} from "../../model";

import { customInstance } from "../../mutator/custom-instance";

/**
 * List webhooks
 * @summary list webhooks
 */
export type webhooksListResponse200 = {
  data: ListResponseBody5;
  status: 200;
};

export type webhooksListResponse400 = {
  data: BadRequestError;
  status: 400;
};

export type webhooksListResponse401 = {
  data: UnauthorizedError;
  status: 401;
};

export type webhooksListResponse403 = {
  data: ForbiddenError;
  status: 403;
};

export type webhooksListResponse404 = {
  data: NotFoundError;
  status: 404;
};

export type webhooksListResponse500 = {
  data: InternalServerError;
  status: 500;
};

export type webhooksListResponseComposite =
  | webhooksListResponse200
  | webhooksListResponse400
  | webhooksListResponse401
  | webhooksListResponse403
  | webhooksListResponse404
  | webhooksListResponse500;

export type webhooksListResponse = webhooksListResponseComposite & {
  headers: Headers;
};

export const getWebhooksListUrl = (params?: WebhooksListParams) => {
  const normalizedParams = new URLSearchParams();

  Object.entries(params || {}).forEach(([key, value]) => {
    if (value !== undefined) {
      normalizedParams.append(key, value === null ? "null" : value.toString());
    }
  });

  const stringifiedParams = normalizedParams.toString();

  return stringifiedParams.length > 0
    ? `/v1/webhooks?${stringifiedParams}`
    : `/v1/webhooks`;
};

export const webhooksList = async (
  params?: WebhooksListParams,
  options?: RequestInit,
): Promise<webhooksListResponse> => {
  return customInstance<webhooksListResponse>(getWebhooksListUrl(params), {
    ...options,
    method: "GET",
  });
};

/**
 * Create a new webhook
 * @summary create webhooks
 */
export type webhooksCreateResponse201 = {
  data: WebhookSecretResponse;
  status: 201;
};

export type webhooksCreateResponse400 = {
  data: BadRequestError;
  status: 400;
};

export type webhooksCreateResponse401 = {
  data: UnauthorizedError;
  status: 401;
};

export type webhooksCreateResponse403 = {
  data: ForbiddenError;
  status: 403;
};

export type webhooksCreateResponse404 = {
  data: NotFoundError;
  status: 404;
};

export type webhooksCreateResponse500 = {
  data: InternalServerError;
  status: 500;
};

export type webhooksCreateResponseComposite =
  | webhooksCreateResponse201
  | webhooksCreateResponse400
  | webhooksCreateResponse401
  | webhooksCreateResponse403
  | webhooksCreateResponse404
  | webhooksCreateResponse500;

export type webhooksCreateResponse = webhooksCreateResponseComposite & {
  headers: Headers;
};

export const getWebhooksCreateUrl = () => {
  return `/v1/webhooks`;
};

export const webhooksCreate = async (
  createRequestBody3: CreateRequestBody3,
  options?: RequestInit,
): Promise<webhooksCreateResponse> => {
  return customInstance<webhooksCreateResponse>(getWebhooksCreateUrl(), {
    ...options,
    method: "POST",
    headers: { "Content-Type": "application/json", ...options?.headers },
    body: JSON.stringify(createRequestBody3),
  });
};

/**
 * Receive webhook callbacks from external sources
 * @summary receive webhooks
 */
export type webhooksReceiveResponse200 = {
  data: SendResponseBody;
  status: 200;
};

export type webhooksReceiveResponse400 = {
  data: BadRequestError;
  status: 400;
};

export type webhooksReceiveResponse401 = {
  data: UnauthorizedError;
  status: 401;
};

export type webhooksReceiveResponse403 = {
  data: ForbiddenError;
  status: 403;
};

export type webhooksReceiveResponse404 = {
  data: NotFoundError;
  status: 404;
};

export type webhooksReceiveResponse500 = {
  data: InternalServerError;
  status: 500;
};

export type webhooksReceiveResponseComposite =
  | webhooksReceiveResponse200
  | webhooksReceiveResponse400
  | webhooksReceiveResponse401
  | webhooksReceiveResponse403
  | webhooksReceiveResponse404
  | webhooksReceiveResponse500;

export type webhooksReceiveResponse = webhooksReceiveResponseComposite & {
  headers: Headers;
};

export const getWebhooksReceiveUrl = (id: string) => {
  return `/v1/webhooks/external/receive/${id}`;
};

export const webhooksReceive = async (
  id: string,
  options?: RequestInit,
): Promise<webhooksReceiveResponse> => {
  return customInstance<webhooksReceiveResponse>(getWebhooksReceiveUrl(id), {
    ...options,
    method: "POST",
  });
};

/**
 * Manually trigger a webhook event
 * @summary trigger_event webhooks
 */
export type webhooksTriggerEventResponse200 = {
  data: WebhookEventResponse;
  status: 200;
};

export type webhooksTriggerEventResponse400 = {
  data: BadRequestError;
  status: 400;
};

export type webhooksTriggerEventResponse401 = {
  data: UnauthorizedError;
  status: 401;
};

export type webhooksTriggerEventResponse403 = {
  data: ForbiddenError;
  status: 403;
};

export type webhooksTriggerEventResponse404 = {
  data: NotFoundError;
  status: 404;
};

export type webhooksTriggerEventResponse500 = {
  data: InternalServerError;
  status: 500;
};

export type webhooksTriggerEventResponseComposite =
  | webhooksTriggerEventResponse200
  | webhooksTriggerEventResponse400
  | webhooksTriggerEventResponse401
  | webhooksTriggerEventResponse403
  | webhooksTriggerEventResponse404
  | webhooksTriggerEventResponse500;

export type webhooksTriggerEventResponse =
  webhooksTriggerEventResponseComposite & {
    headers: Headers;
  };

export const getWebhooksTriggerEventUrl = () => {
  return `/v1/webhooks/trigger`;
};

export const webhooksTriggerEvent = async (
  triggerEventRequestBody: TriggerEventRequestBody,
  options?: RequestInit,
): Promise<webhooksTriggerEventResponse> => {
  return customInstance<webhooksTriggerEventResponse>(
    getWebhooksTriggerEventUrl(),
    {
      ...options,
      method: "POST",
      headers: { "Content-Type": "application/json", ...options?.headers },
      body: JSON.stringify(triggerEventRequestBody),
    },
  );
};

/**
 * Delete webhook
 * @summary delete webhooks
 */
export type webhooksDeleteResponse204 = {
  data: void;
  status: 204;
};

export type webhooksDeleteResponse400 = {
  data: BadRequestError;
  status: 400;
};

export type webhooksDeleteResponse401 = {
  data: UnauthorizedError;
  status: 401;
};

export type webhooksDeleteResponse403 = {
  data: ForbiddenError;
  status: 403;
};

export type webhooksDeleteResponse404 = {
  data: NotFoundError;
  status: 404;
};

export type webhooksDeleteResponse500 = {
  data: InternalServerError;
  status: 500;
};

export type webhooksDeleteResponseComposite =
  | webhooksDeleteResponse204
  | webhooksDeleteResponse400
  | webhooksDeleteResponse401
  | webhooksDeleteResponse403
  | webhooksDeleteResponse404
  | webhooksDeleteResponse500;

export type webhooksDeleteResponse = webhooksDeleteResponseComposite & {
  headers: Headers;
};

export const getWebhooksDeleteUrl = (id: string) => {
  return `/v1/webhooks/${id}`;
};

export const webhooksDelete = async (
  id: string,
  options?: RequestInit,
): Promise<webhooksDeleteResponse> => {
  return customInstance<webhooksDeleteResponse>(getWebhooksDeleteUrl(id), {
    ...options,
    method: "DELETE",
  });
};

/**
 * Get webhook by ID
 * @summary get webhooks
 */
export type webhooksGetResponse200 = {
  data: WebhookResponse;
  status: 200;
};

export type webhooksGetResponse400 = {
  data: BadRequestError;
  status: 400;
};

export type webhooksGetResponse401 = {
  data: UnauthorizedError;
  status: 401;
};

export type webhooksGetResponse403 = {
  data: ForbiddenError;
  status: 403;
};

export type webhooksGetResponse404 = {
  data: NotFoundError;
  status: 404;
};

export type webhooksGetResponse500 = {
  data: InternalServerError;
  status: 500;
};

export type webhooksGetResponseComposite =
  | webhooksGetResponse200
  | webhooksGetResponse400
  | webhooksGetResponse401
  | webhooksGetResponse403
  | webhooksGetResponse404
  | webhooksGetResponse500;

export type webhooksGetResponse = webhooksGetResponseComposite & {
  headers: Headers;
};

export const getWebhooksGetUrl = (id: string) => {
  return `/v1/webhooks/${id}`;
};

export const webhooksGet = async (
  id: string,
  options?: RequestInit,
): Promise<webhooksGetResponse> => {
  return customInstance<webhooksGetResponse>(getWebhooksGetUrl(id), {
    ...options,
    method: "GET",
  });
};

/**
 * Update webhook
 * @summary update webhooks
 */
export type webhooksUpdateResponse200 = {
  data: WebhookResponse;
  status: 200;
};

export type webhooksUpdateResponse400 = {
  data: BadRequestError;
  status: 400;
};

export type webhooksUpdateResponse401 = {
  data: UnauthorizedError;
  status: 401;
};

export type webhooksUpdateResponse403 = {
  data: ForbiddenError;
  status: 403;
};

export type webhooksUpdateResponse404 = {
  data: NotFoundError;
  status: 404;
};

export type webhooksUpdateResponse500 = {
  data: InternalServerError;
  status: 500;
};

export type webhooksUpdateResponseComposite =
  | webhooksUpdateResponse200
  | webhooksUpdateResponse400
  | webhooksUpdateResponse401
  | webhooksUpdateResponse403
  | webhooksUpdateResponse404
  | webhooksUpdateResponse500;

export type webhooksUpdateResponse = webhooksUpdateResponseComposite & {
  headers: Headers;
};

export const getWebhooksUpdateUrl = (id: string) => {
  return `/v1/webhooks/${id}`;
};

export const webhooksUpdate = async (
  id: string,
  updateRequestBody5: UpdateRequestBody5,
  options?: RequestInit,
): Promise<webhooksUpdateResponse> => {
  return customInstance<webhooksUpdateResponse>(getWebhooksUpdateUrl(id), {
    ...options,
    method: "PUT",
    headers: { "Content-Type": "application/json", ...options?.headers },
    body: JSON.stringify(updateRequestBody5),
  });
};

/**
 * List webhook events
 * @summary list_events webhooks
 */
export type webhooksListEventsResponse200 = {
  data: ListEventsResponseBody;
  status: 200;
};

export type webhooksListEventsResponse400 = {
  data: BadRequestError;
  status: 400;
};

export type webhooksListEventsResponse401 = {
  data: UnauthorizedError;
  status: 401;
};

export type webhooksListEventsResponse403 = {
  data: ForbiddenError;
  status: 403;
};

export type webhooksListEventsResponse404 = {
  data: NotFoundError;
  status: 404;
};

export type webhooksListEventsResponse500 = {
  data: InternalServerError;
  status: 500;
};

export type webhooksListEventsResponseComposite =
  | webhooksListEventsResponse200
  | webhooksListEventsResponse400
  | webhooksListEventsResponse401
  | webhooksListEventsResponse403
  | webhooksListEventsResponse404
  | webhooksListEventsResponse500;

export type webhooksListEventsResponse = webhooksListEventsResponseComposite & {
  headers: Headers;
};

export const getWebhooksListEventsUrl = (
  id: string,
  params?: WebhooksListEventsParams,
) => {
  const normalizedParams = new URLSearchParams();

  Object.entries(params || {}).forEach(([key, value]) => {
    if (value !== undefined) {
      normalizedParams.append(key, value === null ? "null" : value.toString());
    }
  });

  const stringifiedParams = normalizedParams.toString();

  return stringifiedParams.length > 0
    ? `/v1/webhooks/${id}/events?${stringifiedParams}`
    : `/v1/webhooks/${id}/events`;
};

export const webhooksListEvents = async (
  id: string,
  params?: WebhooksListEventsParams,
  options?: RequestInit,
): Promise<webhooksListEventsResponse> => {
  return customInstance<webhooksListEventsResponse>(
    getWebhooksListEventsUrl(id, params),
    {
      ...options,
      method: "GET",
    },
  );
};

/**
 * Replay a webhook event
 * @summary replay_event webhooks
 */
export type webhooksReplayEventResponse200 = {
  data: WebhookEventResponse;
  status: 200;
};

export type webhooksReplayEventResponse400 = {
  data: BadRequestError;
  status: 400;
};

export type webhooksReplayEventResponse401 = {
  data: UnauthorizedError;
  status: 401;
};

export type webhooksReplayEventResponse403 = {
  data: ForbiddenError;
  status: 403;
};

export type webhooksReplayEventResponse404 = {
  data: NotFoundError;
  status: 404;
};

export type webhooksReplayEventResponse500 = {
  data: InternalServerError;
  status: 500;
};

export type webhooksReplayEventResponseComposite =
  | webhooksReplayEventResponse200
  | webhooksReplayEventResponse400
  | webhooksReplayEventResponse401
  | webhooksReplayEventResponse403
  | webhooksReplayEventResponse404
  | webhooksReplayEventResponse500;

export type webhooksReplayEventResponse =
  webhooksReplayEventResponseComposite & {
    headers: Headers;
  };

export const getWebhooksReplayEventUrl = (id: string, eventId: string) => {
  return `/v1/webhooks/${id}/events/${eventId}/replay`;
};

export const webhooksReplayEvent = async (
  id: string,
  eventId: string,
  options?: RequestInit,
): Promise<webhooksReplayEventResponse> => {
  return customInstance<webhooksReplayEventResponse>(
    getWebhooksReplayEventUrl(id, eventId),
    {
      ...options,
      method: "POST",
    },
  );
};
