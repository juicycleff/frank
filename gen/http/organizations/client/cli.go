// Code generated by goa v3.20.0, DO NOT EDIT.
//
// organizations HTTP client CLI support package
//
// Command:
// $ goa gen github.com/juicycleff/frank/design -o .

package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	organizations "github.com/juicycleff/frank/gen/organizations"
	goa "goa.design/goa/v3/pkg"
)

// BuildListPayload builds the payload for the organizations list endpoint from
// CLI flags.
func BuildListPayload(organizationsListOffset string, organizationsListLimit string, organizationsListSearch string, organizationsListJWT string) (*organizations.ListPayload, error) {
	var err error
	var offset int
	{
		if organizationsListOffset != "" {
			var v int64
			v, err = strconv.ParseInt(organizationsListOffset, 10, strconv.IntSize)
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
		if organizationsListLimit != "" {
			var v int64
			v, err = strconv.ParseInt(organizationsListLimit, 10, strconv.IntSize)
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
	var search *string
	{
		if organizationsListSearch != "" {
			search = &organizationsListSearch
		}
	}
	var jwt *string
	{
		if organizationsListJWT != "" {
			jwt = &organizationsListJWT
		}
	}
	v := &organizations.ListPayload{}
	v.Offset = offset
	v.Limit = limit
	v.Search = search
	v.JWT = jwt

	return v, nil
}

// BuildCreatePayload builds the payload for the organizations create endpoint
// from CLI flags.
func BuildCreatePayload(organizationsCreateBody string, organizationsCreateJWT string) (*organizations.CreatePayload, error) {
	var err error
	var body CreateRequestBody
	{
		err = json.Unmarshal([]byte(organizationsCreateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"organization\": {\n         \"domain\": \"acme.com\",\n         \"features\": [\n            \"sso\",\n            \"webhooks\"\n         ],\n         \"logo_url\": \"Ut maiores.\",\n         \"metadata\": {\n            \"Consectetur sunt eius.\": \"Sunt natus dolorem ut deleniti odio eligendi.\",\n            \"Ducimus itaque.\": \"Voluptatibus eum delectus aut unde.\",\n            \"Impedit nihil quibusdam esse sed.\": \"Tempore magnam delectus asperiores non.\"\n         },\n         \"name\": \"Acme Inc.\",\n         \"plan\": \"enterprise\",\n         \"slug\": \"acme\",\n         \"trial_days\": 30\n      }\n   }'")
		}
		if body.Organization == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("organization", "body"))
		}
		if body.Organization != nil {
			if err2 := ValidateCreateOrganizationRequestRequestBody(body.Organization); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if err != nil {
			return nil, err
		}
	}
	var jwt *string
	{
		if organizationsCreateJWT != "" {
			jwt = &organizationsCreateJWT
		}
	}
	v := &organizations.CreatePayload{}
	if body.Organization != nil {
		v.Organization = marshalCreateOrganizationRequestRequestBodyToOrganizationsCreateOrganizationRequest(body.Organization)
	}
	v.JWT = jwt

	return v, nil
}

// BuildGetPayload builds the payload for the organizations get endpoint from
// CLI flags.
func BuildGetPayload(organizationsGetID string, organizationsGetJWT string) (*organizations.GetPayload, error) {
	var id string
	{
		id = organizationsGetID
	}
	var jwt *string
	{
		if organizationsGetJWT != "" {
			jwt = &organizationsGetJWT
		}
	}
	v := &organizations.GetPayload{}
	v.ID = id
	v.JWT = jwt

	return v, nil
}

// BuildUpdatePayload builds the payload for the organizations update endpoint
// from CLI flags.
func BuildUpdatePayload(organizationsUpdateBody string, organizationsUpdateID string, organizationsUpdateJWT string) (*organizations.UpdatePayload, error) {
	var err error
	var body UpdateRequestBody
	{
		err = json.Unmarshal([]byte(organizationsUpdateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"organization\": {\n         \"active\": false,\n         \"domain\": \"Voluptas magnam reprehenderit.\",\n         \"logo_url\": \"Nihil architecto inventore id ea laudantium excepturi.\",\n         \"metadata\": {\n            \"Numquam fuga suscipit quis ea sapiente.\": \"Sed aut molestiae totam aut.\",\n            \"Qui deserunt maxime quam.\": \"Non autem a repellat aut doloremque.\",\n            \"Sunt nostrum et quae voluptatum eligendi.\": \"Temporibus non.\"\n         },\n         \"name\": \"Omnis delectus maiores odit magni dolor magnam.\",\n         \"plan\": \"Alias doloremque explicabo perspiciatis natus.\"\n      }\n   }'")
		}
		if body.Organization == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("organization", "body"))
		}
		if err != nil {
			return nil, err
		}
	}
	var id string
	{
		id = organizationsUpdateID
	}
	var jwt *string
	{
		if organizationsUpdateJWT != "" {
			jwt = &organizationsUpdateJWT
		}
	}
	v := &organizations.UpdatePayload{}
	if body.Organization != nil {
		v.Organization = marshalUpdateOrganizationRequestRequestBodyToOrganizationsUpdateOrganizationRequest(body.Organization)
	}
	v.ID = id
	v.JWT = jwt

	return v, nil
}

// BuildDeletePayload builds the payload for the organizations delete endpoint
// from CLI flags.
func BuildDeletePayload(organizationsDeleteID string, organizationsDeleteJWT string) (*organizations.DeletePayload, error) {
	var id string
	{
		id = organizationsDeleteID
	}
	var jwt *string
	{
		if organizationsDeleteJWT != "" {
			jwt = &organizationsDeleteJWT
		}
	}
	v := &organizations.DeletePayload{}
	v.ID = id
	v.JWT = jwt

	return v, nil
}

// BuildListMembersPayload builds the payload for the organizations
// list_members endpoint from CLI flags.
func BuildListMembersPayload(organizationsListMembersID string, organizationsListMembersOffset string, organizationsListMembersLimit string, organizationsListMembersSearch string, organizationsListMembersJWT string) (*organizations.ListMembersPayload, error) {
	var err error
	var id string
	{
		id = organizationsListMembersID
	}
	var offset int
	{
		if organizationsListMembersOffset != "" {
			var v int64
			v, err = strconv.ParseInt(organizationsListMembersOffset, 10, strconv.IntSize)
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
		if organizationsListMembersLimit != "" {
			var v int64
			v, err = strconv.ParseInt(organizationsListMembersLimit, 10, strconv.IntSize)
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
	var search *string
	{
		if organizationsListMembersSearch != "" {
			search = &organizationsListMembersSearch
		}
	}
	var jwt *string
	{
		if organizationsListMembersJWT != "" {
			jwt = &organizationsListMembersJWT
		}
	}
	v := &organizations.ListMembersPayload{}
	v.ID = id
	v.Offset = offset
	v.Limit = limit
	v.Search = search
	v.JWT = jwt

	return v, nil
}

// BuildAddMemberPayload builds the payload for the organizations add_member
// endpoint from CLI flags.
func BuildAddMemberPayload(organizationsAddMemberBody string, organizationsAddMemberID string, organizationsAddMemberJWT string) (*organizations.AddMemberPayload, error) {
	var err error
	var body AddMemberRequestBody
	{
		err = json.Unmarshal([]byte(organizationsAddMemberBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"member\": {\n         \"roles\": [\n            \"member\"\n         ],\n         \"user_id\": \"Enim accusamus voluptatum non rerum maxime numquam.\"\n      }\n   }'")
		}
		if body.Member == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("member", "body"))
		}
		if body.Member != nil {
			if err2 := ValidateAddOrganizationMemberRequestRequestBody(body.Member); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if err != nil {
			return nil, err
		}
	}
	var id string
	{
		id = organizationsAddMemberID
	}
	var jwt *string
	{
		if organizationsAddMemberJWT != "" {
			jwt = &organizationsAddMemberJWT
		}
	}
	v := &organizations.AddMemberPayload{}
	if body.Member != nil {
		v.Member = marshalAddOrganizationMemberRequestRequestBodyToOrganizationsAddOrganizationMemberRequest(body.Member)
	}
	v.ID = id
	v.JWT = jwt

	return v, nil
}

// BuildUpdateMemberPayload builds the payload for the organizations
// update_member endpoint from CLI flags.
func BuildUpdateMemberPayload(organizationsUpdateMemberBody string, organizationsUpdateMemberID string, organizationsUpdateMemberUserID string, organizationsUpdateMemberJWT string) (*organizations.UpdateMemberPayload, error) {
	var err error
	var body UpdateMemberRequestBody
	{
		err = json.Unmarshal([]byte(organizationsUpdateMemberBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"member\": {\n         \"roles\": [\n            \"Deserunt at quia pariatur eos ea.\",\n            \"In ex.\"\n         ]\n      }\n   }'")
		}
		if body.Member == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("member", "body"))
		}
		if body.Member != nil {
			if err2 := ValidateUpdateOrganizationMemberRequestRequestBody(body.Member); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if err != nil {
			return nil, err
		}
	}
	var id string
	{
		id = organizationsUpdateMemberID
	}
	var userID string
	{
		userID = organizationsUpdateMemberUserID
	}
	var jwt *string
	{
		if organizationsUpdateMemberJWT != "" {
			jwt = &organizationsUpdateMemberJWT
		}
	}
	v := &organizations.UpdateMemberPayload{}
	if body.Member != nil {
		v.Member = marshalUpdateOrganizationMemberRequestRequestBodyToOrganizationsUpdateOrganizationMemberRequest(body.Member)
	}
	v.ID = id
	v.UserID = userID
	v.JWT = jwt

	return v, nil
}

// BuildRemoveMemberPayload builds the payload for the organizations
// remove_member endpoint from CLI flags.
func BuildRemoveMemberPayload(organizationsRemoveMemberID string, organizationsRemoveMemberUserID string, organizationsRemoveMemberJWT string) (*organizations.RemoveMemberPayload, error) {
	var id string
	{
		id = organizationsRemoveMemberID
	}
	var userID string
	{
		userID = organizationsRemoveMemberUserID
	}
	var jwt *string
	{
		if organizationsRemoveMemberJWT != "" {
			jwt = &organizationsRemoveMemberJWT
		}
	}
	v := &organizations.RemoveMemberPayload{}
	v.ID = id
	v.UserID = userID
	v.JWT = jwt

	return v, nil
}

// BuildListFeaturesPayload builds the payload for the organizations
// list_features endpoint from CLI flags.
func BuildListFeaturesPayload(organizationsListFeaturesID string, organizationsListFeaturesJWT string) (*organizations.ListFeaturesPayload, error) {
	var id string
	{
		id = organizationsListFeaturesID
	}
	var jwt *string
	{
		if organizationsListFeaturesJWT != "" {
			jwt = &organizationsListFeaturesJWT
		}
	}
	v := &organizations.ListFeaturesPayload{}
	v.ID = id
	v.JWT = jwt

	return v, nil
}

// BuildEnableFeaturePayload builds the payload for the organizations
// enable_feature endpoint from CLI flags.
func BuildEnableFeaturePayload(organizationsEnableFeatureBody string, organizationsEnableFeatureID string, organizationsEnableFeatureJWT string) (*organizations.EnableFeaturePayload, error) {
	var err error
	var body EnableFeatureRequestBody
	{
		err = json.Unmarshal([]byte(organizationsEnableFeatureBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"feature\": {\n         \"feature_key\": \"Laborum repellat repudiandae voluptatem laboriosam voluptate velit.\",\n         \"settings\": {\n            \"Odio est tempore hic ut sit.\": \"Neque necessitatibus sed.\"\n         }\n      }\n   }'")
		}
		if body.Feature == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("feature", "body"))
		}
		if err != nil {
			return nil, err
		}
	}
	var id string
	{
		id = organizationsEnableFeatureID
	}
	var jwt *string
	{
		if organizationsEnableFeatureJWT != "" {
			jwt = &organizationsEnableFeatureJWT
		}
	}
	v := &organizations.EnableFeaturePayload{}
	if body.Feature != nil {
		v.Feature = marshalEnableFeatureRequestRequestBodyToOrganizationsEnableFeatureRequest(body.Feature)
	}
	v.ID = id
	v.JWT = jwt

	return v, nil
}

// BuildDisableFeaturePayload builds the payload for the organizations
// disable_feature endpoint from CLI flags.
func BuildDisableFeaturePayload(organizationsDisableFeatureID string, organizationsDisableFeatureFeatureKey string, organizationsDisableFeatureJWT string) (*organizations.DisableFeaturePayload, error) {
	var id string
	{
		id = organizationsDisableFeatureID
	}
	var featureKey string
	{
		featureKey = organizationsDisableFeatureFeatureKey
	}
	var jwt *string
	{
		if organizationsDisableFeatureJWT != "" {
			jwt = &organizationsDisableFeatureJWT
		}
	}
	v := &organizations.DisableFeaturePayload{}
	v.ID = id
	v.FeatureKey = featureKey
	v.JWT = jwt

	return v, nil
}
