# GetDatabaseClusterReplicas200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**Replicas** | [**[]Replica**](Replica.md) |  | 

## Methods

### NewGetDatabaseClusterReplicas200Response

`func NewGetDatabaseClusterReplicas200Response(meta Meta, replicas []Replica, ) *GetDatabaseClusterReplicas200Response`

NewGetDatabaseClusterReplicas200Response instantiates a new GetDatabaseClusterReplicas200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDatabaseClusterReplicas200ResponseWithDefaults

`func NewGetDatabaseClusterReplicas200ResponseWithDefaults() *GetDatabaseClusterReplicas200Response`

NewGetDatabaseClusterReplicas200ResponseWithDefaults instantiates a new GetDatabaseClusterReplicas200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetDatabaseClusterReplicas200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetDatabaseClusterReplicas200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetDatabaseClusterReplicas200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetReplicas

`func (o *GetDatabaseClusterReplicas200Response) GetReplicas() []Replica`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *GetDatabaseClusterReplicas200Response) GetReplicasOk() (*[]Replica, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *GetDatabaseClusterReplicas200Response) SetReplicas(v []Replica)`

SetReplicas sets Replicas field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


