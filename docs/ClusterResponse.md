# ClusterResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | Pointer to **string** | ID запроса | [optional] 
**Cluster** | [**ClusterOut**](ClusterOut.md) |  | 

## Methods

### NewClusterResponse

`func NewClusterResponse(cluster ClusterOut, ) *ClusterResponse`

NewClusterResponse instantiates a new ClusterResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterResponseWithDefaults

`func NewClusterResponseWithDefaults() *ClusterResponse`

NewClusterResponseWithDefaults instantiates a new ClusterResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *ClusterResponse) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *ClusterResponse) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *ClusterResponse) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.

### HasResponseId

`func (o *ClusterResponse) HasResponseId() bool`

HasResponseId returns a boolean if a field has been set.

### GetCluster

`func (o *ClusterResponse) GetCluster() ClusterOut`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ClusterResponse) GetClusterOk() (*ClusterOut, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ClusterResponse) SetCluster(v ClusterOut)`

SetCluster sets Cluster field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


