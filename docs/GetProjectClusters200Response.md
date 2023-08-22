# GetProjectClusters200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clusters** | [**[]Clusterk8s**](Clusterk8s.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 

## Methods

### NewGetProjectClusters200Response

`func NewGetProjectClusters200Response(clusters []Clusterk8s, meta Meta, ) *GetProjectClusters200Response`

NewGetProjectClusters200Response instantiates a new GetProjectClusters200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProjectClusters200ResponseWithDefaults

`func NewGetProjectClusters200ResponseWithDefaults() *GetProjectClusters200Response`

NewGetProjectClusters200ResponseWithDefaults instantiates a new GetProjectClusters200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusters

`func (o *GetProjectClusters200Response) GetClusters() []Clusterk8s`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *GetProjectClusters200Response) GetClustersOk() (*[]Clusterk8s, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *GetProjectClusters200Response) SetClusters(v []Clusterk8s)`

SetClusters sets Clusters field to given value.


### GetMeta

`func (o *GetProjectClusters200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetProjectClusters200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetProjectClusters200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


