# GetRestorePoints200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RestorePoints** | [**[]RestorePoint**](RestorePoint.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 

## Methods

### NewGetRestorePoints200Response

`func NewGetRestorePoints200Response(restorePoints []RestorePoint, meta Meta, ) *GetRestorePoints200Response`

NewGetRestorePoints200Response instantiates a new GetRestorePoints200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRestorePoints200ResponseWithDefaults

`func NewGetRestorePoints200ResponseWithDefaults() *GetRestorePoints200Response`

NewGetRestorePoints200ResponseWithDefaults instantiates a new GetRestorePoints200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRestorePoints

`func (o *GetRestorePoints200Response) GetRestorePoints() []RestorePoint`

GetRestorePoints returns the RestorePoints field if non-nil, zero value otherwise.

### GetRestorePointsOk

`func (o *GetRestorePoints200Response) GetRestorePointsOk() (*[]RestorePoint, bool)`

GetRestorePointsOk returns a tuple with the RestorePoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestorePoints

`func (o *GetRestorePoints200Response) SetRestorePoints(v []RestorePoint)`

SetRestorePoints sets RestorePoints field to given value.


### GetMeta

`func (o *GetRestorePoints200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetRestorePoints200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetRestorePoints200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


