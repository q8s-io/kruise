// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PortBinding PortBinding represents a binding between a host IP address and a host port
// swagger:model PortBinding
type PortBinding struct {

	// Host IP address that the container's port is mapped to.
	HostIP string `json:"HostIp,omitempty"`

	// Host port number that the container's port is mapped to. range (0,65535]
	// Pattern: ^([1-9]|[1-9]\d{1,3}|[1-5]\d{4}|6[0-4]\d{3}|65[0-4]\d{2}|655[0-2]\d|6553[0-5])$
	HostPort string `json:"HostPort,omitempty"`
}

// Validate validates this port binding
func (m *PortBinding) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHostPort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortBinding) validateHostPort(formats strfmt.Registry) error {

	if swag.IsZero(m.HostPort) { // not required
		return nil
	}

	if err := validate.Pattern("HostPort", "body", string(m.HostPort), `^([1-9]|[1-9]\d{1,3}|[1-5]\d{4}|6[0-4]\d{3}|65[0-4]\d{2}|655[0-2]\d|6553[0-5])$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PortBinding) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortBinding) UnmarshalBinary(b []byte) error {
	var res PortBinding
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
