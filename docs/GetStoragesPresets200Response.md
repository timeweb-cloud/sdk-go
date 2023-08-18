# GetStoragesPresets200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StoragesPresets** | [**[]PresetsStorage**](PresetsStorage.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 

## Methods

### NewGetStoragesPresets200Response

`func NewGetStoragesPresets200Response(storagesPresets []PresetsStorage, meta Meta, responseId string, ) *GetStoragesPresets200Response`

NewGetStoragesPresets200Response instantiates a new GetStoragesPresets200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStoragesPresets200ResponseWithDefaults

`func NewGetStoragesPresets200ResponseWithDefaults() *GetStoragesPresets200Response`

NewGetStoragesPresets200ResponseWithDefaults instantiates a new GetStoragesPresets200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStoragesPresets

`func (o *GetStoragesPresets200Response) GetStoragesPresets() []PresetsStorage`

GetStoragesPresets returns the StoragesPresets field if non-nil, zero value otherwise.

### GetStoragesPresetsOk

`func (o *GetStoragesPresets200Response) GetStoragesPresetsOk() (*[]PresetsStorage, bool)`

GetStoragesPresetsOk returns a tuple with the StoragesPresets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragesPresets

`func (o *GetStoragesPresets200Response) SetStoragesPresets(v []PresetsStorage)`

SetStoragesPresets sets StoragesPresets field to given value.


### GetMeta

`func (o *GetStoragesPresets200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetStoragesPresets200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetStoragesPresets200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetResponseId

`func (o *GetStoragesPresets200Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *GetStoragesPresets200Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *GetStoragesPresets200Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


