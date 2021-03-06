package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Node20Node Get a node into RackHD
// swagger:model node.2.0_Node
type Node20Node struct {

	// auto discover
	AutoDiscover string `json:"autoDiscover,omitempty"`

	// boot settings
	BootSettings interface{} `json:"bootSettings,omitempty"`

	// ibms
	Ibms []interface{} `json:"ibms"`

	// id
	ID string `json:"id,omitempty"`

	// identifiers
	Identifiers []string `json:"identifiers"`

	// Name of the node
	Name string `json:"name,omitempty"`

	// The list of obm settings
	Obms []*NodesPostObmByID `json:"obms"`

	// relations
	Relations []*RelationsObj `json:"relations"`

	// tags
	Tags string `json:"tags,omitempty"`

	// Type of node
	Type string `json:"type,omitempty"`

	// workflows
	Workflows string `json:"workflows,omitempty"`
}

// Validate validates this node 2 0 node
func (m *Node20Node) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIbms(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIdentifiers(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateObms(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRelations(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Node20Node) validateIbms(formats strfmt.Registry) error {

	if swag.IsZero(m.Ibms) { // not required
		return nil
	}

	for i := 0; i < len(m.Ibms); i++ {

	}

	return nil
}

func (m *Node20Node) validateIdentifiers(formats strfmt.Registry) error {

	if swag.IsZero(m.Identifiers) { // not required
		return nil
	}

	for i := 0; i < len(m.Identifiers); i++ {

	}

	return nil
}

func (m *Node20Node) validateObms(formats strfmt.Registry) error {

	if swag.IsZero(m.Obms) { // not required
		return nil
	}

	for i := 0; i < len(m.Obms); i++ {

		if swag.IsZero(m.Obms[i]) { // not required
			continue
		}

		if m.Obms[i] != nil {

			if err := m.Obms[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("obms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Node20Node) validateRelations(formats strfmt.Registry) error {

	if swag.IsZero(m.Relations) { // not required
		return nil
	}

	for i := 0; i < len(m.Relations); i++ {

		if swag.IsZero(m.Relations[i]) { // not required
			continue
		}

		if m.Relations[i] != nil {

			if err := m.Relations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("relations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var node20NodeTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["compute","compute-container","switch","dae","pdu","mgmt","enclosure","rack"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		node20NodeTypeTypePropEnum = append(node20NodeTypeTypePropEnum, v)
	}
}

const (
	// Node20NodeTypeCompute captures enum value "compute"
	Node20NodeTypeCompute string = "compute"
	// Node20NodeTypeComputeContainer captures enum value "compute-container"
	Node20NodeTypeComputeContainer string = "compute-container"
	// Node20NodeTypeSwitch captures enum value "switch"
	Node20NodeTypeSwitch string = "switch"
	// Node20NodeTypeDae captures enum value "dae"
	Node20NodeTypeDae string = "dae"
	// Node20NodeTypePdu captures enum value "pdu"
	Node20NodeTypePdu string = "pdu"
	// Node20NodeTypeMgmt captures enum value "mgmt"
	Node20NodeTypeMgmt string = "mgmt"
	// Node20NodeTypeEnclosure captures enum value "enclosure"
	Node20NodeTypeEnclosure string = "enclosure"
	// Node20NodeTypeRack captures enum value "rack"
	Node20NodeTypeRack string = "rack"
)

// prop value enum
func (m *Node20Node) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, node20NodeTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Node20Node) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}
