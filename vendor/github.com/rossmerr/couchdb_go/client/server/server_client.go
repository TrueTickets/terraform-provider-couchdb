// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new server API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for server API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DBsInfo(params *DBsInfoParams) (*DBsInfoOK, error)

	ActiveTasks(params *ActiveTasksParams) (*ActiveTasksOK, error)

	AllDBs(params *AllDBsParams) (*AllDBsOK, error)

	ClusterSetupGet(params *ClusterSetupGetParams) (*ClusterSetupGetOK, error)

	ClusterSetupPost(params *ClusterSetupPostParams) (*ClusterSetupPostOK, error)

	Membership(params *MembershipParams) (*MembershipOK, error)

	MetaInformation(params *MetaInformationParams) (*MetaInformationOK, error)

	Replication(params *ReplicationParams) (*ReplicationOK, *ReplicationAccepted, error)

	SearchAnalyze(params *SearchAnalyzeParams) (*SearchAnalyzeOK, error)

	Up(params *UpParams) (*UpOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DBsInfo returns information of a list of the specified databases in the couch d b instance

  This enables you to request information about multiple databases in a single request, in place of multiple GET /{db} requests.

*/
func (a *Client) DBsInfo(params *DBsInfoParams) (*DBsInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDBsInfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DBsInfo",
		Method:             "POST",
		PathPattern:        "/_dbs_info",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DBsInfoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DBsInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DBsInfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ActiveTasks lists of running tasks including the task type name status and process ID

  The result is a JSON array of the currently running tasks, with each task being described with a single object. Depending on operation type set of response object fields might be different.

*/
func (a *Client) ActiveTasks(params *ActiveTasksParams) (*ActiveTasksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewActiveTasksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "activeTasks",
		Method:             "GET",
		PathPattern:        "/_active_tasks",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ActiveTasksReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ActiveTasksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for activeTasks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AllDBs returns a list of all the databases in the couch d b instance
*/
func (a *Client) AllDBs(params *AllDBsParams) (*AllDBsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAllDBsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "allDBs",
		Method:             "GET",
		PathPattern:        "/_all_dbs",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AllDBsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AllDBsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for allDBs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ClusterSetupGet returns the status of the node or cluster per the cluster setup wizard
*/
func (a *Client) ClusterSetupGet(params *ClusterSetupGetParams) (*ClusterSetupGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewClusterSetupGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "clusterSetupGet",
		Method:             "GET",
		PathPattern:        "/_cluster_setup",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ClusterSetupGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ClusterSetupGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for clusterSetupGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ClusterSetupPost configures a node as a single standalone node as part of a cluster or finalise a cluster
*/
func (a *Client) ClusterSetupPost(params *ClusterSetupPostParams) (*ClusterSetupPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewClusterSetupPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "clusterSetupPost",
		Method:             "POST",
		PathPattern:        "/_cluster_setup",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ClusterSetupPostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ClusterSetupPostOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for clusterSetupPost: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Membership displays the nodes that are part of the cluster as cluster nodes

  The field all_nodes displays all nodes this node knows about, including the ones that are part of the cluster. The endpoint is useful when setting up a cluster, see Node Management

*/
func (a *Client) Membership(params *MembershipParams) (*MembershipOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMembershipParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "membership",
		Method:             "GET",
		PathPattern:        "/_membership",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &MembershipReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*MembershipOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for membership: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  MetaInformation accessings the root of a couch d b instance returns meta information about the instance

  The response is a JSON structure containing information about the server, including a welcome message and the version of the server.

*/
func (a *Client) MetaInformation(params *MetaInformationParams) (*MetaInformationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMetaInformationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "metaInformation",
		Method:             "GET",
		PathPattern:        "/",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &MetaInformationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*MetaInformationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for metaInformation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Replication requests configure or stop a replication operation
*/
func (a *Client) Replication(params *ReplicationParams) (*ReplicationOK, *ReplicationAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplicationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "replication",
		Method:             "POST",
		PathPattern:        "/_replicate",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ReplicationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ReplicationOK:
		return value, nil, nil
	case *ReplicationAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for server: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SearchAnalyze tests the results of lucene analyzer tokenization on sample text

  *Warning*
Search endpoints require a running search plugin connected to each cluster node. See Search Plugin Installation for details.

*/
func (a *Client) SearchAnalyze(params *SearchAnalyzeParams) (*SearchAnalyzeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchAnalyzeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "searchAnalyze",
		Method:             "POST",
		PathPattern:        "/_search_analyze",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SearchAnalyzeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SearchAnalyzeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for searchAnalyze: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Up confirms that the server is up running and ready to respond to requests

  If maintenance_mode is true or nolb, the endpoint will return a 404 response.

*/
func (a *Client) Up(params *UpParams) (*UpOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "up",
		Method:             "GET",
		PathPattern:        "/_up",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for up: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
