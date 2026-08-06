# UpdateHttpResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Название CDN-ресурса. | [optional] 
**Description** | Pointer to **string** | Описание CDN-ресурса. | [optional] 
**PresetId** | Pointer to **int32** | ID тарифа CDN. Список доступных тарифов можно получить в &#x60;/api/v1/cdn/presets&#x60;. | [optional] 
**StorageId** | Pointer to **int32** | ID S3-хранилища, которое будет источником контента. Нельзя передавать вместе с &#x60;config.origin.servers&#x60;. | [optional] 
**TrafficLimitBytes** | Pointer to **NullableInt32** | Лимит исходящего трафика на расчетный месяц, в байтах. &#x60;null&#x60; — снять лимит. Если ресурс был остановлен по лимиту, при снятии или увеличении лимита раздача возобновится. | [optional] 
**Config** | Pointer to [**HttpResourceConfig**](HttpResourceConfig.md) |  | [optional] 

## Methods

### NewUpdateHttpResource

`func NewUpdateHttpResource() *UpdateHttpResource`

NewUpdateHttpResource instantiates a new UpdateHttpResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateHttpResourceWithDefaults

`func NewUpdateHttpResourceWithDefaults() *UpdateHttpResource`

NewUpdateHttpResourceWithDefaults instantiates a new UpdateHttpResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateHttpResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateHttpResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateHttpResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateHttpResource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateHttpResource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateHttpResource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateHttpResource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateHttpResource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPresetId

`func (o *UpdateHttpResource) GetPresetId() int32`

GetPresetId returns the PresetId field if non-nil, zero value otherwise.

### GetPresetIdOk

`func (o *UpdateHttpResource) GetPresetIdOk() (*int32, bool)`

GetPresetIdOk returns a tuple with the PresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetId

`func (o *UpdateHttpResource) SetPresetId(v int32)`

SetPresetId sets PresetId field to given value.

### HasPresetId

`func (o *UpdateHttpResource) HasPresetId() bool`

HasPresetId returns a boolean if a field has been set.

### GetStorageId

`func (o *UpdateHttpResource) GetStorageId() int32`

GetStorageId returns the StorageId field if non-nil, zero value otherwise.

### GetStorageIdOk

`func (o *UpdateHttpResource) GetStorageIdOk() (*int32, bool)`

GetStorageIdOk returns a tuple with the StorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageId

`func (o *UpdateHttpResource) SetStorageId(v int32)`

SetStorageId sets StorageId field to given value.

### HasStorageId

`func (o *UpdateHttpResource) HasStorageId() bool`

HasStorageId returns a boolean if a field has been set.

### GetTrafficLimitBytes

`func (o *UpdateHttpResource) GetTrafficLimitBytes() int32`

GetTrafficLimitBytes returns the TrafficLimitBytes field if non-nil, zero value otherwise.

### GetTrafficLimitBytesOk

`func (o *UpdateHttpResource) GetTrafficLimitBytesOk() (*int32, bool)`

GetTrafficLimitBytesOk returns a tuple with the TrafficLimitBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficLimitBytes

`func (o *UpdateHttpResource) SetTrafficLimitBytes(v int32)`

SetTrafficLimitBytes sets TrafficLimitBytes field to given value.

### HasTrafficLimitBytes

`func (o *UpdateHttpResource) HasTrafficLimitBytes() bool`

HasTrafficLimitBytes returns a boolean if a field has been set.

### SetTrafficLimitBytesNil

`func (o *UpdateHttpResource) SetTrafficLimitBytesNil(b bool)`

 SetTrafficLimitBytesNil sets the value for TrafficLimitBytes to be an explicit nil

### UnsetTrafficLimitBytes
`func (o *UpdateHttpResource) UnsetTrafficLimitBytes()`

UnsetTrafficLimitBytes ensures that no value is present for TrafficLimitBytes, not even an explicit nil
### GetConfig

`func (o *UpdateHttpResource) GetConfig() HttpResourceConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateHttpResource) GetConfigOk() (*HttpResourceConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateHttpResource) SetConfig(v HttpResourceConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateHttpResource) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


