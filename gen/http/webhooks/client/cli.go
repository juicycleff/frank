// Code generated by goa v3.20.0, DO NOT EDIT.
//
// webhooks HTTP client CLI support package
//
// Command:
// $ goa gen github.com/juicycleff/frank/design -o .

package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	webhooks "github.com/juicycleff/frank/gen/webhooks"
	goa "goa.design/goa/v3/pkg"
)

// BuildListPayload builds the payload for the webhooks list endpoint from CLI
// flags.
func BuildListPayload(webhooksListOffset string, webhooksListLimit string, webhooksListOrganizationID string, webhooksListEventTypes string, webhooksListJWT string) (*webhooks.ListPayload, error) {
	var err error
	var offset int
	{
		if webhooksListOffset != "" {
			var v int64
			v, err = strconv.ParseInt(webhooksListOffset, 10, strconv.IntSize)
			offset = int(v)
			if err != nil {
				return nil, fmt.Errorf("invalid value for offset, must be INT")
			}
			if offset < 0 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("offset", offset, 0, true))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var limit int
	{
		if webhooksListLimit != "" {
			var v int64
			v, err = strconv.ParseInt(webhooksListLimit, 10, strconv.IntSize)
			limit = int(v)
			if err != nil {
				return nil, fmt.Errorf("invalid value for limit, must be INT")
			}
			if limit < 1 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("limit", limit, 1, true))
			}
			if limit > 100 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("limit", limit, 100, false))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var organizationID *string
	{
		if webhooksListOrganizationID != "" {
			organizationID = &webhooksListOrganizationID
		}
	}
	var eventTypes []string
	{
		if webhooksListEventTypes != "" {
			err = json.Unmarshal([]byte(webhooksListEventTypes), &eventTypes)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for eventTypes, \nerror: %s, \nexample of valid JSON:\n%s", err, "'[\n      \"Quaerat et.\",\n      \"Sed nam qui provident ullam.\",\n      \"Impedit veniam molestias.\"\n   ]'")
			}
		}
	}
	var jwt *string
	{
		if webhooksListJWT != "" {
			jwt = &webhooksListJWT
		}
	}
	v := &webhooks.ListPayload{}
	v.Offset = offset
	v.Limit = limit
	v.OrganizationID = organizationID
	v.EventTypes = eventTypes
	v.JWT = jwt

	return v, nil
}

// BuildCreatePayload builds the payload for the webhooks create endpoint from
// CLI flags.
func BuildCreatePayload(webhooksCreateBody string, webhooksCreateJWT string) (*webhooks.CreatePayload, error) {
	var err error
	var body CreateRequestBody
	{
		err = json.Unmarshal([]byte(webhooksCreateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"organization_id\": \"Et eius deleniti amet.\",\n      \"webhook\": {\n         \"event_types\": [\n            \"user.created\",\n            \"user.updated\"\n         ],\n         \"format\": \"json\",\n         \"metadata\": {\n            \"Aspernatur sint veritatis.\": \"Quis est vel aut voluptatum.\",\n            \"Soluta quos velit corrupti qui minima.\": \"Non accusamus.\"\n         },\n         \"name\": \"User Events\",\n         \"retry_count\": 3,\n         \"timeout_ms\": 13431,\n         \"url\": \"https://example.com/webhooks/receive\"\n      }\n   }'")
		}
		if body.Webhook == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("webhook", "body"))
		}
		if body.Webhook != nil {
			if err2 := ValidateCreateWebhookRequestRequestBody(body.Webhook); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if err != nil {
			return nil, err
		}
	}
	var jwt *string
	{
		if webhooksCreateJWT != "" {
			jwt = &webhooksCreateJWT
		}
	}
	v := &webhooks.CreatePayload{
		OrganizationID: body.OrganizationID,
	}
	if body.Webhook != nil {
		v.Webhook = marshalCreateWebhookRequestRequestBodyToWebhooksCreateWebhookRequest(body.Webhook)
	}
	v.JWT = jwt

	return v, nil
}

// BuildGetPayload builds the payload for the webhooks get endpoint from CLI
// flags.
func BuildGetPayload(webhooksGetID string, webhooksGetJWT string) (*webhooks.GetPayload, error) {
	var id string
	{
		id = webhooksGetID
	}
	var jwt *string
	{
		if webhooksGetJWT != "" {
			jwt = &webhooksGetJWT
		}
	}
	v := &webhooks.GetPayload{}
	v.ID = id
	v.JWT = jwt

	return v, nil
}

// BuildUpdatePayload builds the payload for the webhooks update endpoint from
// CLI flags.
func BuildUpdatePayload(webhooksUpdateBody string, webhooksUpdateID string, webhooksUpdateJWT string) (*webhooks.UpdatePayload, error) {
	var err error
	var body UpdateRequestBody
	{
		err = json.Unmarshal([]byte(webhooksUpdateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"webhook\": {\n         \"active\": false,\n         \"event_types\": [\n            \"Commodi repellendus quia est illum eligendi.\",\n            \"Nesciunt quia quis.\",\n            \"Beatae quis et recusandae est explicabo dolorem.\",\n            \"Impedit non tempora cumque.\"\n         ],\n         \"format\": \"form\",\n         \"metadata\": {\n            \"Doloremque explicabo asperiores sunt in quasi.\": \"Molestiae hic nesciunt voluptas soluta.\",\n            \"Sequi eos voluptatem et soluta vitae.\": \"Et et adipisci ab.\"\n         },\n         \"name\": \"Inventore consequatur velit consequatur consequatur saepe unde.\",\n         \"retry_count\": 7,\n         \"timeout_ms\": 25434,\n         \"url\": \"http://lesch.info/katrina\"\n      }\n   }'")
		}
		if body.Webhook == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("webhook", "body"))
		}
		if body.Webhook != nil {
			if err2 := ValidateUpdateWebhookRequestRequestBody(body.Webhook); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if err != nil {
			return nil, err
		}
	}
	var id string
	{
		id = webhooksUpdateID
	}
	var jwt *string
	{
		if webhooksUpdateJWT != "" {
			jwt = &webhooksUpdateJWT
		}
	}
	v := &webhooks.UpdatePayload{}
	if body.Webhook != nil {
		v.Webhook = marshalUpdateWebhookRequestRequestBodyToWebhooksUpdateWebhookRequest(body.Webhook)
	}
	v.ID = id
	v.JWT = jwt

	return v, nil
}

// BuildDeletePayload builds the payload for the webhooks delete endpoint from
// CLI flags.
func BuildDeletePayload(webhooksDeleteID string, webhooksDeleteJWT string) (*webhooks.DeletePayload, error) {
	var id string
	{
		id = webhooksDeleteID
	}
	var jwt *string
	{
		if webhooksDeleteJWT != "" {
			jwt = &webhooksDeleteJWT
		}
	}
	v := &webhooks.DeletePayload{}
	v.ID = id
	v.JWT = jwt

	return v, nil
}

// BuildTriggerEventPayload builds the payload for the webhooks trigger_event
// endpoint from CLI flags.
func BuildTriggerEventPayload(webhooksTriggerEventBody string, webhooksTriggerEventJWT string) (*webhooks.TriggerEventPayload, error) {
	var err error
	var body TriggerEventRequestBody
	{
		err = json.Unmarshal([]byte(webhooksTriggerEventBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"event\": {\n         \"event_type\": \"user.created\",\n         \"headers\": {\n            \"Et voluptate.\": \"In placeat sed et.\",\n            \"Perspiciatis sint aperiam quia similique.\": \"Voluptatibus id autem vero.\"\n         },\n         \"payload\": {\n            \"email\": \"user@example.com\",\n            \"user_id\": \"123\"\n         }\n      },\n      \"organization_id\": \"Quis animi adipisci unde.\"\n   }'")
		}
		if body.Event == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("event", "body"))
		}
		if body.Event != nil {
			if err2 := ValidateTriggerEventRequestRequestBody(body.Event); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if err != nil {
			return nil, err
		}
	}
	var jwt *string
	{
		if webhooksTriggerEventJWT != "" {
			jwt = &webhooksTriggerEventJWT
		}
	}
	v := &webhooks.TriggerEventPayload{
		OrganizationID: body.OrganizationID,
	}
	if body.Event != nil {
		v.Event = marshalTriggerEventRequestRequestBodyToWebhooksTriggerEventRequest(body.Event)
	}
	v.JWT = jwt

	return v, nil
}

// BuildListEventsPayload builds the payload for the webhooks list_events
// endpoint from CLI flags.
func BuildListEventsPayload(webhooksListEventsID string, webhooksListEventsOffset string, webhooksListEventsLimit string, webhooksListEventsEventType string, webhooksListEventsDelivered string, webhooksListEventsJWT string) (*webhooks.ListEventsPayload, error) {
	var err error
	var id string
	{
		id = webhooksListEventsID
	}
	var offset int
	{
		if webhooksListEventsOffset != "" {
			var v int64
			v, err = strconv.ParseInt(webhooksListEventsOffset, 10, strconv.IntSize)
			offset = int(v)
			if err != nil {
				return nil, fmt.Errorf("invalid value for offset, must be INT")
			}
			if offset < 0 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("offset", offset, 0, true))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var limit int
	{
		if webhooksListEventsLimit != "" {
			var v int64
			v, err = strconv.ParseInt(webhooksListEventsLimit, 10, strconv.IntSize)
			limit = int(v)
			if err != nil {
				return nil, fmt.Errorf("invalid value for limit, must be INT")
			}
			if limit < 1 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("limit", limit, 1, true))
			}
			if limit > 100 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("limit", limit, 100, false))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var eventType *string
	{
		if webhooksListEventsEventType != "" {
			eventType = &webhooksListEventsEventType
		}
	}
	var delivered *bool
	{
		if webhooksListEventsDelivered != "" {
			var val bool
			val, err = strconv.ParseBool(webhooksListEventsDelivered)
			delivered = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for delivered, must be BOOL")
			}
		}
	}
	var jwt *string
	{
		if webhooksListEventsJWT != "" {
			jwt = &webhooksListEventsJWT
		}
	}
	v := &webhooks.ListEventsPayload{}
	v.ID = id
	v.Offset = offset
	v.Limit = limit
	v.EventType = eventType
	v.Delivered = delivered
	v.JWT = jwt

	return v, nil
}

// BuildReplayEventPayload builds the payload for the webhooks replay_event
// endpoint from CLI flags.
func BuildReplayEventPayload(webhooksReplayEventID string, webhooksReplayEventEventID string, webhooksReplayEventJWT string) (*webhooks.ReplayEventPayload, error) {
	var id string
	{
		id = webhooksReplayEventID
	}
	var eventID string
	{
		eventID = webhooksReplayEventEventID
	}
	var jwt *string
	{
		if webhooksReplayEventJWT != "" {
			jwt = &webhooksReplayEventJWT
		}
	}
	v := &webhooks.ReplayEventPayload{}
	v.ID = id
	v.EventID = eventID
	v.JWT = jwt

	return v, nil
}

// BuildReceivePayload builds the payload for the webhooks receive endpoint
// from CLI flags.
func BuildReceivePayload(webhooksReceiveID string) (*webhooks.ReceivePayload, error) {
	var id string
	{
		id = webhooksReceiveID
	}
	v := &webhooks.ReceivePayload{}
	v.ID = id

	return v, nil
}
