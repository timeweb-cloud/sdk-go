# CreateHttpResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Название CDN-ресурса. | 
**PresetId** | **int32** | ID тарифа CDN. Список доступных тарифов можно получить в &#x60;/api/v1/cdn/presets&#x60;. | 
**Description** | Pointer to **string** | Описание CDN-ресурса. | [optional] 
**StorageId** | Pointer to **int32** | ID S3-хранилища, которое будет источником контента. Нельзя передавать вместе с &#x60;server&#x60;. | [optional] 
**Server** | Pointer to [**OriginServer**](OriginServer.md) |  | [optional] 
**UseHttps** | Pointer to **bool** | Обращаться к источнику контента по HTTPS. | [optional] 
**DeliveryDomain** | Pointer to **string** | Собственный домен, с которого будет раздаваться контент. Домен добавляется к техническому домену ресурса, для его работы нужно направить CNAME-запись на &#x60;cdn_domain&#x60;. | [optional] 
**ProjectId** | Pointer to **int32** | ID проекта, в который нужно поместить ресурс. Если не указан, ресурс попадет в проект по умолчанию. | [optional] 
**TrafficLimitBytes** | Pointer to **NullableInt32** | Лимит исходящего трафика на расчетный месяц, в байтах. При достижении лимита раздача останавливается автоматически. | [optional] 

## Methods

### NewCreateHttpResource

`func NewCreateHttpResource(name string, presetId int32, ) *CreateHttpResource`

NewCreateHttpResource instantiates a new CreateHttpResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateHttpResourceWithDefaults

`func NewCreateHttpResourceWithDefaults() *CreateHttpResource`

NewCreateHttpResourceWithDefaults instantiates a new CreateHttpResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateHttpResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateHttpResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateHttpResource) SetName(v string)`

SetName sets Name field to given value.


### GetPresetId

`func (o *CreateHttpResource) GetPresetId() int32`

GetPresetId returns the PresetId field if non-nil, zero value otherwise.

### GetPresetIdOk

`func (o *CreateHttpResource) GetPresetIdOk() (*int32, bool)`

GetPresetIdOk returns a tuple with the PresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetId

`func (o *CreateHttpResource) SetPresetId(v int32)`

SetPresetId sets PresetId field to given value.


### GetDescription

`func (o *CreateHttpResource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateHttpResource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateHttpResource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateHttpResource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStorageId

`func (o *CreateHttpResource) GetStorageId() int32`

GetStorageId returns the StorageId field if non-nil, zero value otherwise.

### GetStorageIdOk

`func (o *CreateHttpResource) GetStorageIdOk() (*int32, bool)`

GetStorageIdOk returns a tuple with the StorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageId

`func (o *CreateHttpResource) SetStorageId(v int32)`

SetStorageId sets StorageId field to given value.

### HasStorageId

`func (o *CreateHttpResource) HasStorageId() bool`

HasStorageId returns a boolean if a field has been set.

### GetServer

`func (o *CreateHttpResource) GetServer() OriginServer`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *CreateHttpResource) GetServerOk() (*OriginServer, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *CreateHttpResource) SetServer(v OriginServer)`

SetServer sets Server field to given value.

### HasServer

`func (o *CreateHttpResource) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetUseHttps

`func (o *CreateHttpResource) GetUseHttps() bool`

GetUseHttps returns the UseHttps field if non-nil, zero value otherwise.

### GetUseHttpsOk

`func (o *CreateHttpResource) GetUseHttpsOk() (*bool, bool)`

GetUseHttpsOk returns a tuple with the UseHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseHttps

`func (o *CreateHttpResource) SetUseHttps(v bool)`

SetUseHttps sets UseHttps field to given value.

### HasUseHttps

`func (o *CreateHttpResource) HasUseHttps() bool`

HasUseHttps returns a boolean if a field has been set.

### GetDeliveryDomain

`func (o *CreateHttpResource) GetDeliveryDomain() string`

GetDeliveryDomain returns the DeliveryDomain field if non-nil, zero value otherwise.

### GetDeliveryDomainOk

`func (o *CreateHttpResource) GetDeliveryDomainOk() (*string, bool)`

GetDeliveryDomainOk returns a tuple with the DeliveryDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDomain

`func (o *CreateHttpResource) SetDeliveryDomain(v string)`

SetDeliveryDomain sets DeliveryDomain field to given value.

### HasDeliveryDomain

`func (o *CreateHttpResource) HasDeliveryDomain() bool`

HasDeliveryDomain returns a boolean if a field has been set.

### GetProjectId

`func (o *CreateHttpResource) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateHttpResource) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateHttpResource) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *CreateHttpResource) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetTrafficLimitBytes

`func (o *CreateHttpResource) GetTrafficLimitBytes() int32`

GetTrafficLimitBytes returns the TrafficLimitBytes field if non-nil, zero value otherwise.

### GetTrafficLimitBytesOk

`func (o *CreateHttpResource) GetTrafficLimitBytesOk() (*int32, bool)`

GetTrafficLimitBytesOk returns a tuple with the TrafficLimitBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficLimitBytes

`func (o *CreateHttpResource) SetTrafficLimitBytes(v int32)`

SetTrafficLimitBytes sets TrafficLimitBytes field to given value.

### HasTrafficLimitBytes

`func (o *CreateHttpResource) HasTrafficLimitBytes() bool`

HasTrafficLimitBytes returns a boolean if a field has been set.

### SetTrafficLimitBytesNil

`func (o *CreateHttpResource) SetTrafficLimitBytesNil(b bool)`

 SetTrafficLimitBytesNil sets the value for TrafficLimitBytes to be an explicit nil

### UnsetTrafficLimitBytes
`func (o *CreateHttpResource) UnsetTrafficLimitBytes()`

UnsetTrafficLimitBytes ensures that no value is present for TrafficLimitBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


