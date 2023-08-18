# CreateCluster201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 
**Cluster** | [**ClusterOut**](ClusterOut.md) |  | 

## Methods

### NewCreateCluster201Response

`func NewCreateCluster201Response(responseId string, cluster ClusterOut, ) *CreateCluster201Response`

NewCreateCluster201Response instantiates a new CreateCluster201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCluster201ResponseWithDefaults

`func NewCreateCluster201ResponseWithDefaults() *CreateCluster201Response`

NewCreateCluster201ResponseWithDefaults instantiates a new CreateCluster201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *CreateCluster201Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *CreateCluster201Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *CreateCluster201Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.


### GetCluster

`func (o *CreateCluster201Response) GetCluster() ClusterOut`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *CreateCluster201Response) GetClusterOk() (*ClusterOut, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *CreateCluster201Response) SetCluster(v ClusterOut)`

SetCluster sets Cluster field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


