# GetDatabasesPresets200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**DatabasesPresets** | [**[]PresetsDbs**](PresetsDbs.md) |  | 
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 

## Methods

### NewGetDatabasesPresets200Response

`func NewGetDatabasesPresets200Response(meta Meta, databasesPresets []PresetsDbs, responseId string, ) *GetDatabasesPresets200Response`

NewGetDatabasesPresets200Response instantiates a new GetDatabasesPresets200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDatabasesPresets200ResponseWithDefaults

`func NewGetDatabasesPresets200ResponseWithDefaults() *GetDatabasesPresets200Response`

NewGetDatabasesPresets200ResponseWithDefaults instantiates a new GetDatabasesPresets200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetDatabasesPresets200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetDatabasesPresets200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetDatabasesPresets200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetDatabasesPresets

`func (o *GetDatabasesPresets200Response) GetDatabasesPresets() []PresetsDbs`

GetDatabasesPresets returns the DatabasesPresets field if non-nil, zero value otherwise.

### GetDatabasesPresetsOk

`func (o *GetDatabasesPresets200Response) GetDatabasesPresetsOk() (*[]PresetsDbs, bool)`

GetDatabasesPresetsOk returns a tuple with the DatabasesPresets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabasesPresets

`func (o *GetDatabasesPresets200Response) SetDatabasesPresets(v []PresetsDbs)`

SetDatabasesPresets sets DatabasesPresets field to given value.


### GetResponseId

`func (o *GetDatabasesPresets200Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *GetDatabasesPresets200Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *GetDatabasesPresets200Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


