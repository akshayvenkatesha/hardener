// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Ruleresult ruleresult
// swagger:model ruleresult
type Ruleresult struct {

	// result
	Result string `json:"result,omitempty"`

	// rulename
	Rulename string `json:"rulename,omitempty"`
}

// Validate validates this ruleresult
func (m *Ruleresult) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Ruleresult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Ruleresult) UnmarshalBinary(b []byte) error {
	var res Ruleresult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
