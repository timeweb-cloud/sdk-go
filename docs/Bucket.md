# Bucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | ID для каждого экземпляра хранилища. Автоматически генерируется при создании. | 
**Name** | **string** | Удобочитаемое имя, установленное для хранилища. | 
**Description** | **string** | Комментарий к хранилищу. | 
**DiskStats** | [**BucketDiskStats**](BucketDiskStats.md) |  | 
**Type** | **string** | Тип хранилища. | 
**PresetId** | **NullableFloat32** | ID тарифа хранилища. | 
**ConfiguratorId** | **NullableFloat32** | ID конфигуратора хранилища. | 
**AvatarLink** | **NullableString** | Ссылка на аватар хранилища. | 
**Status** | **string** | Статус хранилища. | 
**ObjectAmount** | **float32** | Количество файлов в хранилище. | 
**Location** | **string** | Регион хранилища. | 
**Hostname** | **string** | Адрес хранилища для подключения. | 
**AccessKey** | **string** | Ключ доступа от хранилища. | 
**SecretKey** | **string** | Секретный ключ доступа от хранилища. | 
**MovedInQuarantineAt** | **NullableTime** | Дата перемещения в карантин. | 
**StorageClass** | **string** | Класс хранилища. | 
**ProjectId** | **float32** | ID проекта. | 
**RateId** | **float32** | ID тарифа. | 
**WebsiteConfig** | [**NullableBucketWebsiteConfig**](BucketWebsiteConfig.md) |  | 
**IsAllowAutoUpgrade** | **bool** | Разрешено ли автоматическое повышение тарифа. | 

## Methods

### NewBucket

`func NewBucket(id float32, name string, description string, diskStats BucketDiskStats, type_ string, presetId NullableFloat32, configuratorId NullableFloat32, avatarLink NullableString, status string, objectAmount float32, location string, hostname string, accessKey string, secretKey string, movedInQuarantineAt NullableTime, storageClass string, projectId float32, rateId float32, websiteConfig NullableBucketWebsiteConfig, isAllowAutoUpgrade bool, ) *Bucket`

NewBucket instantiates a new Bucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketWithDefaults

`func NewBucketWithDefaults() *Bucket`

NewBucketWithDefaults instantiates a new Bucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Bucket) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Bucket) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Bucket) SetId(v float32)`

SetId sets Id field to given value.


### GetName

`func (o *Bucket) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Bucket) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Bucket) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Bucket) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Bucket) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Bucket) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDiskStats

`func (o *Bucket) GetDiskStats() BucketDiskStats`

GetDiskStats returns the DiskStats field if non-nil, zero value otherwise.

### GetDiskStatsOk

`func (o *Bucket) GetDiskStatsOk() (*BucketDiskStats, bool)`

GetDiskStatsOk returns a tuple with the DiskStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskStats

`func (o *Bucket) SetDiskStats(v BucketDiskStats)`

SetDiskStats sets DiskStats field to given value.


### GetType

`func (o *Bucket) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Bucket) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Bucket) SetType(v string)`

SetType sets Type field to given value.


### GetPresetId

`func (o *Bucket) GetPresetId() float32`

GetPresetId returns the PresetId field if non-nil, zero value otherwise.

### GetPresetIdOk

`func (o *Bucket) GetPresetIdOk() (*float32, bool)`

GetPresetIdOk returns a tuple with the PresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetId

`func (o *Bucket) SetPresetId(v float32)`

SetPresetId sets PresetId field to given value.


### SetPresetIdNil

`func (o *Bucket) SetPresetIdNil(b bool)`

 SetPresetIdNil sets the value for PresetId to be an explicit nil

### UnsetPresetId
`func (o *Bucket) UnsetPresetId()`

UnsetPresetId ensures that no value is present for PresetId, not even an explicit nil
### GetConfiguratorId

`func (o *Bucket) GetConfiguratorId() float32`

GetConfiguratorId returns the ConfiguratorId field if non-nil, zero value otherwise.

### GetConfiguratorIdOk

`func (o *Bucket) GetConfiguratorIdOk() (*float32, bool)`

GetConfiguratorIdOk returns a tuple with the ConfiguratorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguratorId

`func (o *Bucket) SetConfiguratorId(v float32)`

SetConfiguratorId sets ConfiguratorId field to given value.


### SetConfiguratorIdNil

`func (o *Bucket) SetConfiguratorIdNil(b bool)`

 SetConfiguratorIdNil sets the value for ConfiguratorId to be an explicit nil

### UnsetConfiguratorId
`func (o *Bucket) UnsetConfiguratorId()`

UnsetConfiguratorId ensures that no value is present for ConfiguratorId, not even an explicit nil
### GetAvatarLink

`func (o *Bucket) GetAvatarLink() string`

GetAvatarLink returns the AvatarLink field if non-nil, zero value otherwise.

### GetAvatarLinkOk

`func (o *Bucket) GetAvatarLinkOk() (*string, bool)`

GetAvatarLinkOk returns a tuple with the AvatarLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarLink

`func (o *Bucket) SetAvatarLink(v string)`

SetAvatarLink sets AvatarLink field to given value.


### SetAvatarLinkNil

`func (o *Bucket) SetAvatarLinkNil(b bool)`

 SetAvatarLinkNil sets the value for AvatarLink to be an explicit nil

### UnsetAvatarLink
`func (o *Bucket) UnsetAvatarLink()`

UnsetAvatarLink ensures that no value is present for AvatarLink, not even an explicit nil
### GetStatus

`func (o *Bucket) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Bucket) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Bucket) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetObjectAmount

`func (o *Bucket) GetObjectAmount() float32`

GetObjectAmount returns the ObjectAmount field if non-nil, zero value otherwise.

### GetObjectAmountOk

`func (o *Bucket) GetObjectAmountOk() (*float32, bool)`

GetObjectAmountOk returns a tuple with the ObjectAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectAmount

`func (o *Bucket) SetObjectAmount(v float32)`

SetObjectAmount sets ObjectAmount field to given value.


### GetLocation

`func (o *Bucket) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Bucket) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Bucket) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetHostname

`func (o *Bucket) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *Bucket) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *Bucket) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetAccessKey

`func (o *Bucket) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *Bucket) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *Bucket) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.


### GetSecretKey

`func (o *Bucket) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *Bucket) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *Bucket) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.


### GetMovedInQuarantineAt

`func (o *Bucket) GetMovedInQuarantineAt() time.Time`

GetMovedInQuarantineAt returns the MovedInQuarantineAt field if non-nil, zero value otherwise.

### GetMovedInQuarantineAtOk

`func (o *Bucket) GetMovedInQuarantineAtOk() (*time.Time, bool)`

GetMovedInQuarantineAtOk returns a tuple with the MovedInQuarantineAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovedInQuarantineAt

`func (o *Bucket) SetMovedInQuarantineAt(v time.Time)`

SetMovedInQuarantineAt sets MovedInQuarantineAt field to given value.


### SetMovedInQuarantineAtNil

`func (o *Bucket) SetMovedInQuarantineAtNil(b bool)`

 SetMovedInQuarantineAtNil sets the value for MovedInQuarantineAt to be an explicit nil

### UnsetMovedInQuarantineAt
`func (o *Bucket) UnsetMovedInQuarantineAt()`

UnsetMovedInQuarantineAt ensures that no value is present for MovedInQuarantineAt, not even an explicit nil
### GetStorageClass

`func (o *Bucket) GetStorageClass() string`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *Bucket) GetStorageClassOk() (*string, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *Bucket) SetStorageClass(v string)`

SetStorageClass sets StorageClass field to given value.


### GetProjectId

`func (o *Bucket) GetProjectId() float32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Bucket) GetProjectIdOk() (*float32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Bucket) SetProjectId(v float32)`

SetProjectId sets ProjectId field to given value.


### GetRateId

`func (o *Bucket) GetRateId() float32`

GetRateId returns the RateId field if non-nil, zero value otherwise.

### GetRateIdOk

`func (o *Bucket) GetRateIdOk() (*float32, bool)`

GetRateIdOk returns a tuple with the RateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateId

`func (o *Bucket) SetRateId(v float32)`

SetRateId sets RateId field to given value.


### GetWebsiteConfig

`func (o *Bucket) GetWebsiteConfig() BucketWebsiteConfig`

GetWebsiteConfig returns the WebsiteConfig field if non-nil, zero value otherwise.

### GetWebsiteConfigOk

`func (o *Bucket) GetWebsiteConfigOk() (*BucketWebsiteConfig, bool)`

GetWebsiteConfigOk returns a tuple with the WebsiteConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteConfig

`func (o *Bucket) SetWebsiteConfig(v BucketWebsiteConfig)`

SetWebsiteConfig sets WebsiteConfig field to given value.


### SetWebsiteConfigNil

`func (o *Bucket) SetWebsiteConfigNil(b bool)`

 SetWebsiteConfigNil sets the value for WebsiteConfig to be an explicit nil

### UnsetWebsiteConfig
`func (o *Bucket) UnsetWebsiteConfig()`

UnsetWebsiteConfig ensures that no value is present for WebsiteConfig, not even an explicit nil
### GetIsAllowAutoUpgrade

`func (o *Bucket) GetIsAllowAutoUpgrade() bool`

GetIsAllowAutoUpgrade returns the IsAllowAutoUpgrade field if non-nil, zero value otherwise.

### GetIsAllowAutoUpgradeOk

`func (o *Bucket) GetIsAllowAutoUpgradeOk() (*bool, bool)`

GetIsAllowAutoUpgradeOk returns a tuple with the IsAllowAutoUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllowAutoUpgrade

`func (o *Bucket) SetIsAllowAutoUpgrade(v bool)`

SetIsAllowAutoUpgrade sets IsAllowAutoUpgrade field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


