# HttpResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID CDN-ресурса. Генерируется автоматически при создании. | 
**Name** | **string** | Название CDN-ресурса. | 
**Description** | Pointer to **string** | Описание CDN-ресурса. | [optional] 
**Source** | **string** | Источник контента: имя S3-бакета, если ресурс создан поверх хранилища, либо хост origin-сервера. | 
**TrafficUsage** | [**TrafficUsage**](TrafficUsage.md) |  | 
**Status** | **string** | Статус CDN-ресурса. - &#x60;created&#x60; — ресурс создан и раздает контент; - &#x60;processing&#x60; — конфигурация применяется на стороне CDN; - &#x60;stopped&#x60; — раздача приостановлена; - &#x60;failed&#x60; — настройка ресурса завершилась с ошибкой; - &#x60;no_paid&#x60; — ресурс не оплачен; - &#x60;blocked&#x60; — ресурс заблокирован; - &#x60;traffic_limit_exceeded&#x60; — раздача остановлена автоматически из-за достижения лимита трафика &#x60;traffic_limit_bytes&#x60;. | 
**CdnDomain** | **string** | Технический домен, выданный ресурсу. Доступен сразу после создания и всегда остается в списке доменов ресурса. | 
**PresetId** | **int32** | ID тарифа CDN. Список доступных тарифов можно получить в &#x60;/api/v1/cdn/presets&#x60;. | 
**ProjectId** | Pointer to **int32** | ID проекта, к которому привязан ресурс. | [optional] 
**AvatarLink** | Pointer to **NullableString** | Ссылка на аватар ресурса. | [optional] 
**StorageId** | Pointer to **NullableInt32** | ID S3-хранилища, которое используется в качестве источника контента. &#x60;null&#x60;, если источником является origin-сервер. | [optional] 
**TrafficLimitBytes** | Pointer to **NullableInt32** | Лимит исходящего трафика на расчетный месяц, в байтах. При достижении лимита раздача останавливается, а ресурс переходит в статус &#x60;traffic_limit_exceeded&#x60;. &#x60;null&#x60; — лимит не задан. | [optional] 

## Methods

### NewHttpResource

`func NewHttpResource(id int32, name string, source string, trafficUsage TrafficUsage, status string, cdnDomain string, presetId int32, ) *HttpResource`

NewHttpResource instantiates a new HttpResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpResourceWithDefaults

`func NewHttpResourceWithDefaults() *HttpResource`

NewHttpResourceWithDefaults instantiates a new HttpResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HttpResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HttpResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HttpResource) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *HttpResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HttpResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HttpResource) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *HttpResource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HttpResource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HttpResource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HttpResource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSource

`func (o *HttpResource) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *HttpResource) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *HttpResource) SetSource(v string)`

SetSource sets Source field to given value.


### GetTrafficUsage

`func (o *HttpResource) GetTrafficUsage() TrafficUsage`

GetTrafficUsage returns the TrafficUsage field if non-nil, zero value otherwise.

### GetTrafficUsageOk

`func (o *HttpResource) GetTrafficUsageOk() (*TrafficUsage, bool)`

GetTrafficUsageOk returns a tuple with the TrafficUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficUsage

`func (o *HttpResource) SetTrafficUsage(v TrafficUsage)`

SetTrafficUsage sets TrafficUsage field to given value.


### GetStatus

`func (o *HttpResource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HttpResource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HttpResource) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCdnDomain

`func (o *HttpResource) GetCdnDomain() string`

GetCdnDomain returns the CdnDomain field if non-nil, zero value otherwise.

### GetCdnDomainOk

`func (o *HttpResource) GetCdnDomainOk() (*string, bool)`

GetCdnDomainOk returns a tuple with the CdnDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdnDomain

`func (o *HttpResource) SetCdnDomain(v string)`

SetCdnDomain sets CdnDomain field to given value.


### GetPresetId

`func (o *HttpResource) GetPresetId() int32`

GetPresetId returns the PresetId field if non-nil, zero value otherwise.

### GetPresetIdOk

`func (o *HttpResource) GetPresetIdOk() (*int32, bool)`

GetPresetIdOk returns a tuple with the PresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetId

`func (o *HttpResource) SetPresetId(v int32)`

SetPresetId sets PresetId field to given value.


### GetProjectId

`func (o *HttpResource) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *HttpResource) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *HttpResource) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *HttpResource) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetAvatarLink

`func (o *HttpResource) GetAvatarLink() string`

GetAvatarLink returns the AvatarLink field if non-nil, zero value otherwise.

### GetAvatarLinkOk

`func (o *HttpResource) GetAvatarLinkOk() (*string, bool)`

GetAvatarLinkOk returns a tuple with the AvatarLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarLink

`func (o *HttpResource) SetAvatarLink(v string)`

SetAvatarLink sets AvatarLink field to given value.

### HasAvatarLink

`func (o *HttpResource) HasAvatarLink() bool`

HasAvatarLink returns a boolean if a field has been set.

### SetAvatarLinkNil

`func (o *HttpResource) SetAvatarLinkNil(b bool)`

 SetAvatarLinkNil sets the value for AvatarLink to be an explicit nil

### UnsetAvatarLink
`func (o *HttpResource) UnsetAvatarLink()`

UnsetAvatarLink ensures that no value is present for AvatarLink, not even an explicit nil
### GetStorageId

`func (o *HttpResource) GetStorageId() int32`

GetStorageId returns the StorageId field if non-nil, zero value otherwise.

### GetStorageIdOk

`func (o *HttpResource) GetStorageIdOk() (*int32, bool)`

GetStorageIdOk returns a tuple with the StorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageId

`func (o *HttpResource) SetStorageId(v int32)`

SetStorageId sets StorageId field to given value.

### HasStorageId

`func (o *HttpResource) HasStorageId() bool`

HasStorageId returns a boolean if a field has been set.

### SetStorageIdNil

`func (o *HttpResource) SetStorageIdNil(b bool)`

 SetStorageIdNil sets the value for StorageId to be an explicit nil

### UnsetStorageId
`func (o *HttpResource) UnsetStorageId()`

UnsetStorageId ensures that no value is present for StorageId, not even an explicit nil
### GetTrafficLimitBytes

`func (o *HttpResource) GetTrafficLimitBytes() int32`

GetTrafficLimitBytes returns the TrafficLimitBytes field if non-nil, zero value otherwise.

### GetTrafficLimitBytesOk

`func (o *HttpResource) GetTrafficLimitBytesOk() (*int32, bool)`

GetTrafficLimitBytesOk returns a tuple with the TrafficLimitBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficLimitBytes

`func (o *HttpResource) SetTrafficLimitBytes(v int32)`

SetTrafficLimitBytes sets TrafficLimitBytes field to given value.

### HasTrafficLimitBytes

`func (o *HttpResource) HasTrafficLimitBytes() bool`

HasTrafficLimitBytes returns a boolean if a field has been set.

### SetTrafficLimitBytesNil

`func (o *HttpResource) SetTrafficLimitBytesNil(b bool)`

 SetTrafficLimitBytesNil sets the value for TrafficLimitBytes to be an explicit nil

### UnsetTrafficLimitBytes
`func (o *HttpResource) UnsetTrafficLimitBytes()`

UnsetTrafficLimitBytes ensures that no value is present for TrafficLimitBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


