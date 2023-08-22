# GetDatabaseClusters200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**Dbs** | [**[]DatabaseCluster**](DatabaseCluster.md) |  | 

## Methods

### NewGetDatabaseClusters200Response

`func NewGetDatabaseClusters200Response(meta Meta, dbs []DatabaseCluster, ) *GetDatabaseClusters200Response`

NewGetDatabaseClusters200Response instantiates a new GetDatabaseClusters200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDatabaseClusters200ResponseWithDefaults

`func NewGetDatabaseClusters200ResponseWithDefaults() *GetDatabaseClusters200Response`

NewGetDatabaseClusters200ResponseWithDefaults instantiates a new GetDatabaseClusters200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetDatabaseClusters200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetDatabaseClusters200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetDatabaseClusters200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetDbs

`func (o *GetDatabaseClusters200Response) GetDbs() []DatabaseCluster`

GetDbs returns the Dbs field if non-nil, zero value otherwise.

### GetDbsOk

`func (o *GetDatabaseClusters200Response) GetDbsOk() (*[]DatabaseCluster, bool)`

GetDbsOk returns a tuple with the Dbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbs

`func (o *GetDatabaseClusters200Response) SetDbs(v []DatabaseCluster)`

SetDbs sets Dbs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


