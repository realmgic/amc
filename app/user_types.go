// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "amc": Application User Types
//
// Command:
// $ goagen
// --design=github.com/citrusleaf/amc/api_design
// --out=$(GOPATH)/src/github.com/citrusleaf/amc/temp
// --version=v1.3.0

package app

import (
	"github.com/goadesign/goa"
)

// nodeSeed user type.
type nodeSeed struct {
	// Node Host. Valid DNS, IPv4 or IPv6 value
	Host *string `form:"host,omitempty" json:"host,omitempty" xml:"host,omitempty"`
	// Node's port
	Port *int `form:"port,omitempty" json:"port,omitempty" xml:"port,omitempty"`
	// Node's TLS name
	TLSName *string `form:"tlsName,omitempty" json:"tlsName,omitempty" xml:"tlsName,omitempty"`
}

// Validate validates the nodeSeed type instance.
func (ut *nodeSeed) Validate() (err error) {
	if ut.Host == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "host"))
	}
	if ut.Port == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "port"))
	}
	if ut.Port != nil {
		if *ut.Port < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`request.port`, *ut.Port, 0, true))
		}
	}
	if ut.Port != nil {
		if *ut.Port > 65535 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`request.port`, *ut.Port, 65535, false))
		}
	}
	return
}

// Publicize creates NodeSeed from nodeSeed
func (ut *nodeSeed) Publicize() *NodeSeed {
	var pub NodeSeed
	if ut.Host != nil {
		pub.Host = *ut.Host
	}
	if ut.Port != nil {
		pub.Port = *ut.Port
	}
	if ut.TLSName != nil {
		pub.TLSName = ut.TLSName
	}
	return &pub
}

// NodeSeed user type.
type NodeSeed struct {
	// Node Host. Valid DNS, IPv4 or IPv6 value
	Host string `form:"host" json:"host" xml:"host"`
	// Node's port
	Port int `form:"port" json:"port" xml:"port"`
	// Node's TLS name
	TLSName *string `form:"tlsName,omitempty" json:"tlsName,omitempty" xml:"tlsName,omitempty"`
}

// Validate validates the NodeSeed type instance.
func (ut *NodeSeed) Validate() (err error) {
	if ut.Host == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "host"))
	}

	if ut.Port < 0 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`type.port`, ut.Port, 0, true))
	}
	if ut.Port > 65535 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`type.port`, ut.Port, 65535, false))
	}
	return
}

// Notification Data type
type notification struct {
	// Cluster ID
	ConnID *string `form:"connId,omitempty" json:"connId,omitempty" xml:"connId,omitempty"`
	// Notification Description
	Desc *string `form:"desc,omitempty" json:"desc,omitempty" xml:"desc,omitempty"`
	// Notification ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Last Occured Time
	LastOccured *int `form:"lastOccured,omitempty" json:"lastOccured,omitempty" xml:"lastOccured,omitempty"`
	// Status
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Notification Type
	Type *string `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
}

// Validate validates the notification type instance.
func (ut *notification) Validate() (err error) {
	if ut.ID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "id"))
	}
	if ut.ConnID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "connId"))
	}
	if ut.Desc == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "desc"))
	}
	if ut.Status == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "status"))
	}
	if ut.Type == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "type"))
	}
	if ut.LastOccured == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "lastOccured"))
	}
	return
}

// Publicize creates Notification from notification
func (ut *notification) Publicize() *Notification {
	var pub Notification
	if ut.ConnID != nil {
		pub.ConnID = *ut.ConnID
	}
	if ut.Desc != nil {
		pub.Desc = *ut.Desc
	}
	if ut.ID != nil {
		pub.ID = *ut.ID
	}
	if ut.LastOccured != nil {
		pub.LastOccured = *ut.LastOccured
	}
	if ut.Status != nil {
		pub.Status = *ut.Status
	}
	if ut.Type != nil {
		pub.Type = *ut.Type
	}
	return &pub
}

// Notification Data type
type Notification struct {
	// Cluster ID
	ConnID string `form:"connId" json:"connId" xml:"connId"`
	// Notification Description
	Desc string `form:"desc" json:"desc" xml:"desc"`
	// Notification ID
	ID string `form:"id" json:"id" xml:"id"`
	// Last Occured Time
	LastOccured int `form:"lastOccured" json:"lastOccured" xml:"lastOccured"`
	// Status
	Status string `form:"status" json:"status" xml:"status"`
	// Notification Type
	Type string `form:"type" json:"type" xml:"type"`
}

// Validate validates the Notification type instance.
func (ut *Notification) Validate() (err error) {
	if ut.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "id"))
	}
	if ut.ConnID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "connId"))
	}
	if ut.Desc == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "desc"))
	}
	if ut.Status == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "status"))
	}
	if ut.Type == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "type"))
	}

	return
}

// privilege user type.
type privilege struct {
	// Namespace
	Namespace *string `form:"namespace,omitempty" json:"namespace,omitempty" xml:"namespace,omitempty"`
	// Database Privileges
	Privilege *string `form:"privilege,omitempty" json:"privilege,omitempty" xml:"privilege,omitempty"`
	// Set
	Set *string `form:"set,omitempty" json:"set,omitempty" xml:"set,omitempty"`
}

// Validate validates the privilege type instance.
func (ut *privilege) Validate() (err error) {
	if ut.Privilege == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "privilege"))
	}
	return
}

// Publicize creates Privilege from privilege
func (ut *privilege) Publicize() *Privilege {
	var pub Privilege
	if ut.Namespace != nil {
		pub.Namespace = ut.Namespace
	}
	if ut.Privilege != nil {
		pub.Privilege = *ut.Privilege
	}
	if ut.Set != nil {
		pub.Set = ut.Set
	}
	return &pub
}

// Privilege user type.
type Privilege struct {
	// Namespace
	Namespace *string `form:"namespace,omitempty" json:"namespace,omitempty" xml:"namespace,omitempty"`
	// Database Privileges
	Privilege string `form:"privilege" json:"privilege" xml:"privilege"`
	// Set
	Set *string `form:"set,omitempty" json:"set,omitempty" xml:"set,omitempty"`
}

// Validate validates the Privilege type instance.
func (ut *Privilege) Validate() (err error) {
	if ut.Privilege == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "privilege"))
	}
	return
}
