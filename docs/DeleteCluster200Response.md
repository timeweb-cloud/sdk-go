# DeleteCluster200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterDelete** | [**DeleteServiceResponse**](DeleteServiceResponse.md) |  | 
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 

## Methods

### NewDeleteCluster200Response

`func NewDeleteCluster200Response(clusterDelete DeleteServiceResponse, responseId string, ) *DeleteCluster200Response`

NewDeleteCluster200Response instantiates a new DeleteCluster200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteCluster200ResponseWithDefaults

`func NewDeleteCluster200ResponseWithDefaults() *DeleteCluster200Response`

NewDeleteCluster200ResponseWithDefaults instantiates a new DeleteCluster200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterDelete

`func (o *DeleteCluster200Response) GetClusterDelete() DeleteServiceResponse`

GetClusterDelete returns the ClusterDelete field if non-nil, zero value otherwise.

### GetClusterDeleteOk

`func (o *DeleteCluster200Response) GetClusterDeleteOk() (*DeleteServiceResponse, bool)`

GetClusterDeleteOk returns a tuple with the ClusterDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterDelete

`func (o *DeleteCluster200Response) SetClusterDelete(v DeleteServiceResponse)`

SetClusterDelete sets ClusterDelete field to given value.


### GetResponseId

`func (o *DeleteCluster200Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *DeleteCluster200Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *DeleteCluster200Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


