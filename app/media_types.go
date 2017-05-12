// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "amc": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/citrusleaf/amc/api_design
// --out=$(GOPATH)/src/github.com/citrusleaf/amc/temp
// --version=v1.2.0-dirty

package app

import (
	"github.com/goadesign/goa"
)

// Cluster Modules (default view)
//
// Identifier: application/vnd.aerospike.amc.connection.modules.response+json; view=default
type AerospikeAmcConnectionModulesResponse struct {
	// Module's Hash
	Hash *string `form:"hash,omitempty" json:"hash,omitempty" xml:"hash,omitempty"`
	// Module's Name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Nodes from which the module is absent
	NodesAbsent []string `form:"nodesAbsent,omitempty" json:"nodesAbsent,omitempty" xml:"nodesAbsent,omitempty"`
	// Nodes in which the module is present
	NodesPresent []string `form:"nodesPresent,omitempty" json:"nodesPresent,omitempty" xml:"nodesPresent,omitempty"`
	// Is Module present in all nodes?
	Synced *bool `form:"synced,omitempty" json:"synced,omitempty" xml:"synced,omitempty"`
	// Module's Source Type
	Type *string `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
}

// Cluster Modules (full view)
//
// Identifier: application/vnd.aerospike.amc.connection.modules.response+json; view=full
type AerospikeAmcConnectionModulesResponseFull struct {
	// Module's Name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Module's Source Code
	Source *string `form:"source,omitempty" json:"source,omitempty" xml:"source,omitempty"`
	// Module's Source Type
	Type *string `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
}

// User Connection (default view)
//
// Identifier: application/vnd.aerospike.amc.connection.query.response+json; view=default
type AerospikeAmcConnectionQueryResponse struct {
	// UI should connect to this connection automatically after AMC login
	ConnectOnLogin bool `form:"connectOnLogin" json:"connectOnLogin" xml:"connectOnLogin"`
	// If AMC is already connected to this cluster for the current user.
	Connected bool `form:"connected" json:"connected" xml:"connected"`
	// Connection Id
	ID string `form:"id" json:"id" xml:"id"`
	// Connection Name
	Name string `form:"name" json:"name" xml:"name"`
	// Seeds
	Seeds []*NodeSeed `form:"seeds,omitempty" json:"seeds,omitempty" xml:"seeds,omitempty"`
}

// Validate validates the AerospikeAmcConnectionQueryResponse media type instance.
func (mt *AerospikeAmcConnectionQueryResponse) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}

	if ok := goa.ValidatePattern(`[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`, mt.ID); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.id`, mt.ID, `[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`))
	}
	for _, e := range mt.Seeds {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// User Connection Entity Tree (default view)
//
// Identifier: application/vnd.aerospike.amc.connection.tree.response+json; view=default
type AerospikeAmcConnectionTreeResponse struct {
	// Entity Type
	EntityType string `form:"entityType" json:"entityType" xml:"entityType"`
	// Connection Id
	ID string `form:"id" json:"id" xml:"id"`
	// Last Update Of This Entity in Unix Seconds
	LastUpdate int `form:"lastUpdate" json:"lastUpdate" xml:"lastUpdate"`
	// modules
	Modules []*AerospikeAmcEntityModuleResponse `form:"modules,omitempty" json:"modules,omitempty" xml:"modules,omitempty"`
	// Nodes
	Nodes []*AerospikeAmcEntityNodeResponse `form:"nodes,omitempty" json:"nodes,omitempty" xml:"nodes,omitempty"`
	// Cureent connection status.
	Status string `form:"status" json:"status" xml:"status"`
}

// Validate validates the AerospikeAmcConnectionTreeResponse media type instance.
func (mt *AerospikeAmcConnectionTreeResponse) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.EntityType == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "entityType"))
	}

	if mt.Status == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "status"))
	}
	if ok := goa.ValidatePattern(`[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`, mt.ID); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.id`, mt.ID, `[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`))
	}
	for _, e := range mt.Modules {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	for _, e := range mt.Nodes {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// Index Entity (default view)
//
// Identifier: application/vnd.aerospike.amc.entity.index.response+json; view=default
type AerospikeAmcEntityIndexResponse struct {
	// Bin Name
	BinName string `form:"binName" json:"binName" xml:"binName"`
	// Type
	EntityType string `form:"entityType" json:"entityType" xml:"entityType"`
	// Last Update Of This Entity in Unix Seconds
	LastUpdate int `form:"lastUpdate" json:"lastUpdate" xml:"lastUpdate"`
	// Set Name
	Name string `form:"name" json:"name" xml:"name"`
	// Index Type
	Type string `form:"type" json:"type" xml:"type"`
}

// Validate validates the AerospikeAmcEntityIndexResponse media type instance.
func (mt *AerospikeAmcEntityIndexResponse) Validate() (err error) {
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.BinName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "binName"))
	}
	if mt.Type == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "type"))
	}
	if mt.EntityType == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "entityType"))
	}

	return
}

// Module Entity (default view)
//
// Identifier: application/vnd.aerospike.amc.entity.module.response+json; view=default
type AerospikeAmcEntityModuleResponse struct {
	// Type
	EntityType string `form:"entityType" json:"entityType" xml:"entityType"`
	// Module Hash
	Hash string `form:"hash" json:"hash" xml:"hash"`
	// Last Update Of This Entity in Unix Seconds
	LastUpdate int `form:"lastUpdate" json:"lastUpdate" xml:"lastUpdate"`
	// Module Name
	Name string `form:"name" json:"name" xml:"name"`
	// Module Type
	Type string `form:"type" json:"type" xml:"type"`
}

// Validate validates the AerospikeAmcEntityModuleResponse media type instance.
func (mt *AerospikeAmcEntityModuleResponse) Validate() (err error) {
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Hash == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "hash"))
	}
	if mt.Type == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "type"))
	}
	if mt.EntityType == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "entityType"))
	}

	return
}

// Namespace Entity (default view)
//
// Identifier: application/vnd.aerospike.amc.entity.namespace.response+json; view=default
type AerospikeAmcEntityNamespaceResponse struct {
	// Type
	EntityType string `form:"entityType" json:"entityType" xml:"entityType"`
	// Last Update Of This Entity in Unix Seconds
	LastUpdate int `form:"lastUpdate" json:"lastUpdate" xml:"lastUpdate"`
	// Namespace Name
	Name string `form:"name" json:"name" xml:"name"`
	// Namespaces
	Sets []*AerospikeAmcEntitySetResponse `form:"sets,omitempty" json:"sets,omitempty" xml:"sets,omitempty"`
}

// Validate validates the AerospikeAmcEntityNamespaceResponse media type instance.
func (mt *AerospikeAmcEntityNamespaceResponse) Validate() (err error) {
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.EntityType == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "entityType"))
	}

	for _, e := range mt.Sets {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// Node Entity (default view)
//
// Identifier: application/vnd.aerospike.amc.entity.node.response+json; view=default
type AerospikeAmcEntityNodeResponse struct {
	// Type
	EntityType string `form:"entityType" json:"entityType" xml:"entityType"`
	// Network Host Address
	Host string `form:"host" json:"host" xml:"host"`
	// Node Id
	ID string `form:"id" json:"id" xml:"id"`
	// Last Update Of This Entity in Unix Seconds
	LastUpdate int `form:"lastUpdate" json:"lastUpdate" xml:"lastUpdate"`
	// Namespaces
	Namespaces []*AerospikeAmcEntityNamespaceResponse `form:"namespaces,omitempty" json:"namespaces,omitempty" xml:"namespaces,omitempty"`
}

// Validate validates the AerospikeAmcEntityNodeResponse media type instance.
func (mt *AerospikeAmcEntityNodeResponse) Validate() (err error) {
	if mt.Host == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "host"))
	}
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.EntityType == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "entityType"))
	}

	for _, e := range mt.Namespaces {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// Set Entity (default view)
//
// Identifier: application/vnd.aerospike.amc.entity.set.response+json; view=default
type AerospikeAmcEntitySetResponse struct {
	// Type
	EntityType string `form:"entityType" json:"entityType" xml:"entityType"`
	// Indexes
	Indexes []*AerospikeAmcEntityIndexResponse `form:"indexes,omitempty" json:"indexes,omitempty" xml:"indexes,omitempty"`
	// Last Update Of This Entity in Unix Seconds
	LastUpdate int `form:"lastUpdate" json:"lastUpdate" xml:"lastUpdate"`
	// Set Name
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the AerospikeAmcEntitySetResponse media type instance.
func (mt *AerospikeAmcEntitySetResponse) Validate() (err error) {
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.EntityType == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "entityType"))
	}

	for _, e := range mt.Indexes {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// AMC Server System information (default view)
//
// Identifier: application/vnd.aerospike.amc.system.response+json; view=default
type AerospikeAmcSystemResponse struct {
	// AMC Version
	Version *string `form:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
}

// User (default view)
//
// Identifier: application/vnd.aerospike.amc.user.query.response+json; view=default
type AerospikeAmcUserQueryResponse struct {
	// User account is active
	Active *bool `form:"active,omitempty" json:"active,omitempty" xml:"active,omitempty"`
	// User's fullname
	FullName *string `form:"fullName,omitempty" json:"fullName,omitempty" xml:"fullName,omitempty"`
	// Additional Notes
	Notes *string `form:"notes,omitempty" json:"notes,omitempty" xml:"notes,omitempty"`
	// AMC User Roles
	Roles []string `form:"roles,omitempty" json:"roles,omitempty" xml:"roles,omitempty"`
	// User Id
	Username string `form:"username" json:"username" xml:"username"`
}

// Validate validates the AerospikeAmcUserQueryResponse media type instance.
func (mt *AerospikeAmcUserQueryResponse) Validate() (err error) {
	if mt.Username == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "username"))
	}
	return
}
