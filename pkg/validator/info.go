package validator

import (
	"fmt"
	"net/url"
)

const (
	MAX_SHORT_FIELD_LENGTH = 80
	MAX_LONG_FIELD_LENGTH  = 300
	MAX_VALIDATOR_INFO     = 576
)

type Info struct {
	Name        string  `pulumi:"name"`
	Website     *string `pulumi:"website,optional"`
	Logo        *string `pulumi:"logo,optional"`
	Description *string `pulumi:"description,optional"`
}

func (i *Info) ToEnv() (map[string]string, error) {
	if len(i.Name) > MAX_SHORT_FIELD_LENGTH {
		return nil, fmt.Errorf("name exceeds maximum length of %d", MAX_SHORT_FIELD_LENGTH)
	}

	env := map[string]string{
		"INFO_NAME": i.Name,
	}

	if i.Website != nil {
		if len(*i.Website) > MAX_SHORT_FIELD_LENGTH {
			return nil, fmt.Errorf("website exceeds maximum length of %d", MAX_SHORT_FIELD_LENGTH)
		}
		if _, err := url.ParseRequestURI(*i.Website); err != nil {
			return nil, fmt.Errorf("invalid website URL: %w", err)
		}
		env["INFO_WEB"] = *i.Website
	}

	if i.Logo != nil {
		if len(*i.Logo) > MAX_SHORT_FIELD_LENGTH {
			return nil, fmt.Errorf("logo exceeds maximum length of %d", MAX_SHORT_FIELD_LENGTH)
		}
		if _, err := url.ParseRequestURI(*i.Logo); err != nil {
			return nil, fmt.Errorf("invalid logo URL: %w", err)
		}
		env["INFO_LOGO"] = *i.Logo
	}

	if i.Description != nil {
		if len(*i.Description) > MAX_LONG_FIELD_LENGTH {
			return nil, fmt.Errorf("description exceeds maximum length of %d", MAX_LONG_FIELD_LENGTH)
		}
		env["INFO_DESCRIPTION"] = *i.Description
	}

	totalLength := len(i.Name)
	if i.Website != nil {
		totalLength += len(*i.Website)
	}
	if i.Logo != nil {
		totalLength += len(*i.Logo)
	}
	if i.Description != nil {
		totalLength += len(*i.Description)
	}

	if totalLength > MAX_VALIDATOR_INFO {
		return nil, fmt.Errorf("total length of fields exceeds maximum allowed length of %d", MAX_VALIDATOR_INFO)
	}

	return env, nil
}
