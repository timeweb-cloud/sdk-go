# UpdateServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configurator** | Pointer to [**UpdateServerConfigurator**](UpdateServerConfigurator.md) |  | [optional] 
**OsId** | Pointer to **float32** | ID операционной системы, которая будет установлена на облачный сервер. | [optional] 
**SoftwareId** | Pointer to **float32** | ID программного обеспечения сервера. | [optional] 
**PresetId** | Pointer to **float32** | ID тарифа сервера. Нельзя передавать вместе с ключом &#x60;configurator&#x60;. | [optional] 
**Bandwidth** | Pointer to **float32** | Пропускная способность тарифа. Доступные значения от 100 до 1000 с шагом 100. | [optional] 
**Name** | Pointer to **string** | Имя облачного сервера. Максимальная длина — 255 символов, имя должно быть уникальным. | [optional] 
**AvatarId** | Pointer to **string** | ID аватара сервера. Описание методов работы с аватарами появится позднее. | [optional] 
**Comment** | Pointer to **string** | Комментарий к облачному серверу. Максимальная длина — 255 символов. | [optional] 
**ImageId** | Pointer to **string** | ID образа, который будет установлен на облачный сервер. Нельзя передавать вместе с &#x60;os_id&#x60;. | [optional] 
**CloudInit** | Pointer to **string** | Cloud-init скрипт | [optional] 

## Methods

### NewUpdateServer

`func NewUpdateServer() *UpdateServer`

NewUpdateServer instantiates a new UpdateServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateServerWithDefaults

`func NewUpdateServerWithDefaults() *UpdateServer`

NewUpdateServerWithDefaults instantiates a new UpdateServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurator

`func (o *UpdateServer) GetConfigurator() UpdateServerConfigurator`

GetConfigurator returns the Configurator field if non-nil, zero value otherwise.

### GetConfiguratorOk

`func (o *UpdateServer) GetConfiguratorOk() (*UpdateServerConfigurator, bool)`

GetConfiguratorOk returns a tuple with the Configurator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurator

`func (o *UpdateServer) SetConfigurator(v UpdateServerConfigurator)`

SetConfigurator sets Configurator field to given value.

### HasConfigurator

`func (o *UpdateServer) HasConfigurator() bool`

HasConfigurator returns a boolean if a field has been set.

### GetOsId

`func (o *UpdateServer) GetOsId() float32`

GetOsId returns the OsId field if non-nil, zero value otherwise.

### GetOsIdOk

`func (o *UpdateServer) GetOsIdOk() (*float32, bool)`

GetOsIdOk returns a tuple with the OsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsId

`func (o *UpdateServer) SetOsId(v float32)`

SetOsId sets OsId field to given value.

### HasOsId

`func (o *UpdateServer) HasOsId() bool`

HasOsId returns a boolean if a field has been set.

### GetSoftwareId

`func (o *UpdateServer) GetSoftwareId() float32`

GetSoftwareId returns the SoftwareId field if non-nil, zero value otherwise.

### GetSoftwareIdOk

`func (o *UpdateServer) GetSoftwareIdOk() (*float32, bool)`

GetSoftwareIdOk returns a tuple with the SoftwareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareId

`func (o *UpdateServer) SetSoftwareId(v float32)`

SetSoftwareId sets SoftwareId field to given value.

### HasSoftwareId

`func (o *UpdateServer) HasSoftwareId() bool`

HasSoftwareId returns a boolean if a field has been set.

### GetPresetId

`func (o *UpdateServer) GetPresetId() float32`

GetPresetId returns the PresetId field if non-nil, zero value otherwise.

### GetPresetIdOk

`func (o *UpdateServer) GetPresetIdOk() (*float32, bool)`

GetPresetIdOk returns a tuple with the PresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetId

`func (o *UpdateServer) SetPresetId(v float32)`

SetPresetId sets PresetId field to given value.

### HasPresetId

`func (o *UpdateServer) HasPresetId() bool`

HasPresetId returns a boolean if a field has been set.

### GetBandwidth

`func (o *UpdateServer) GetBandwidth() float32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *UpdateServer) GetBandwidthOk() (*float32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *UpdateServer) SetBandwidth(v float32)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *UpdateServer) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetName

`func (o *UpdateServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateServer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAvatarId

`func (o *UpdateServer) GetAvatarId() string`

GetAvatarId returns the AvatarId field if non-nil, zero value otherwise.

### GetAvatarIdOk

`func (o *UpdateServer) GetAvatarIdOk() (*string, bool)`

GetAvatarIdOk returns a tuple with the AvatarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarId

`func (o *UpdateServer) SetAvatarId(v string)`

SetAvatarId sets AvatarId field to given value.

### HasAvatarId

`func (o *UpdateServer) HasAvatarId() bool`

HasAvatarId returns a boolean if a field has been set.

### GetComment

`func (o *UpdateServer) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *UpdateServer) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *UpdateServer) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *UpdateServer) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetImageId

`func (o *UpdateServer) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *UpdateServer) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *UpdateServer) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *UpdateServer) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetCloudInit

`func (o *UpdateServer) GetCloudInit() string`

GetCloudInit returns the CloudInit field if non-nil, zero value otherwise.

### GetCloudInitOk

`func (o *UpdateServer) GetCloudInitOk() (*string, bool)`

GetCloudInitOk returns a tuple with the CloudInit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInit

`func (o *UpdateServer) SetCloudInit(v string)`

SetCloudInit sets CloudInit field to given value.

### HasCloudInit

`func (o *UpdateServer) HasCloudInit() bool`

HasCloudInit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


