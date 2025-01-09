# ClustersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | Pointer to **string** | ID запроса | [optional] 
**Meta** | [**SchemasMeta**](SchemasMeta.md) |  | 
**Clusters** | [**[]ClusterOut**](ClusterOut.md) | Массив объектов Кластер | 

## Methods

### NewClustersResponse

`func NewClustersResponse(meta SchemasMeta, clusters []ClusterOut, ) *ClustersResponse`

NewClustersResponse instantiates a new ClustersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClustersResponseWithDefaults

`func NewClustersResponseWithDefaults() *ClustersResponse`

NewClustersResponseWithDefaults instantiates a new ClustersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *ClustersResponse) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *ClustersResponse) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *ClustersResponse) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.

### HasResponseId

`func (o *ClustersResponse) HasResponseId() bool`

HasResponseId returns a boolean if a field has been set.

### GetMeta

`func (o *ClustersResponse) GetMeta() SchemasMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ClustersResponse) GetMetaOk() (*SchemasMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ClustersResponse) SetMeta(v SchemasMeta)`

SetMeta sets Meta field to given value.


### GetClusters

`func (o *ClustersResponse) GetClusters() []ClusterOut`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *ClustersResponse) GetClustersOk() (*[]ClusterOut, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *ClustersResponse) SetClusters(v []ClusterOut)`

SetClusters sets Clusters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


