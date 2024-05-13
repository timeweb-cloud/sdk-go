# GetDatabaseClusterTypes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**Types** | [**[]DatabaseType**](DatabaseType.md) |  | 

## Methods

### NewGetDatabaseClusterTypes200Response

`func NewGetDatabaseClusterTypes200Response(meta Meta, types []DatabaseType, ) *GetDatabaseClusterTypes200Response`

NewGetDatabaseClusterTypes200Response instantiates a new GetDatabaseClusterTypes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDatabaseClusterTypes200ResponseWithDefaults

`func NewGetDatabaseClusterTypes200ResponseWithDefaults() *GetDatabaseClusterTypes200Response`

NewGetDatabaseClusterTypes200ResponseWithDefaults instantiates a new GetDatabaseClusterTypes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetDatabaseClusterTypes200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetDatabaseClusterTypes200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetDatabaseClusterTypes200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetTypes

`func (o *GetDatabaseClusterTypes200Response) GetTypes() []DatabaseType`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *GetDatabaseClusterTypes200Response) GetTypesOk() (*[]DatabaseType, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *GetDatabaseClusterTypes200Response) SetTypes(v []DatabaseType)`

SetTypes sets Types field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


