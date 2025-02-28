# CreateServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to [**CreateServerConfiguration**](CreateServerConfiguration.md) |  | [optional] 
**IsDdosGuard** | Pointer to **bool** | Защита от DDoS. Серверу выдается защищенный IP-адрес с защитой уровня L3 / L4. Для включения защиты уровня L7 необходимо создать тикет в техническую поддержку. | [optional] 
**OsId** | Pointer to **float32** | ID операционной системы, которая будет установлена на облачный сервер. Нельзя передавать вместе с &#x60;image_id&#x60;. | [optional] 
**ImageId** | Pointer to **string** | ID образа, который будет установлен на облачный сервер. Нельзя передавать вместе с &#x60;os_id&#x60;. | [optional] 
**SoftwareId** | Pointer to **float32** | ID программного обеспечения сервера. | [optional] 
**PresetId** | Pointer to **float32** | ID тарифа сервера. Нельзя передавать вместе с ключом &#x60;configurator&#x60;. | [optional] 
**Bandwidth** | Pointer to **float32** | Пропускная способность тарифа. Доступные значения от 100 до 1000 с шагом 100. | [optional] 
**Name** | **string** | Имя облачного сервера. Максимальная длина — 255 символов, имя должно быть уникальным. | 
**AvatarId** | Pointer to **string** | ID аватара сервера. Описание методов работы с аватарами появится позднее. | [optional] 
**Comment** | Pointer to **string** | Комментарий к облачному серверу. Максимальная длина — 255 символов. | [optional] 
**SshKeysIds** | Pointer to **[]float32** | Список SSH-ключей. | [optional] 
**IsLocalNetwork** | Pointer to **bool** | Локальная сеть. | [optional] 
**Network** | Pointer to [**Network**](Network.md) |  | [optional] 
**CloudInit** | Pointer to **string** | Cloud-init скрипт | [optional] 
**AvailabilityZone** | Pointer to [**AvailabilityZone**](AvailabilityZone.md) |  | [optional] 

## Methods

### NewCreateServer

`func NewCreateServer(name string, ) *CreateServer`

NewCreateServer instantiates a new CreateServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateServerWithDefaults

`func NewCreateServerWithDefaults() *CreateServer`

NewCreateServerWithDefaults instantiates a new CreateServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *CreateServer) GetConfiguration() CreateServerConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *CreateServer) GetConfigurationOk() (*CreateServerConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *CreateServer) SetConfiguration(v CreateServerConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *CreateServer) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetIsDdosGuard

`func (o *CreateServer) GetIsDdosGuard() bool`

GetIsDdosGuard returns the IsDdosGuard field if non-nil, zero value otherwise.

### GetIsDdosGuardOk

`func (o *CreateServer) GetIsDdosGuardOk() (*bool, bool)`

GetIsDdosGuardOk returns a tuple with the IsDdosGuard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDdosGuard

`func (o *CreateServer) SetIsDdosGuard(v bool)`

SetIsDdosGuard sets IsDdosGuard field to given value.

### HasIsDdosGuard

`func (o *CreateServer) HasIsDdosGuard() bool`

HasIsDdosGuard returns a boolean if a field has been set.

### GetOsId

`func (o *CreateServer) GetOsId() float32`

GetOsId returns the OsId field if non-nil, zero value otherwise.

### GetOsIdOk

`func (o *CreateServer) GetOsIdOk() (*float32, bool)`

GetOsIdOk returns a tuple with the OsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsId

`func (o *CreateServer) SetOsId(v float32)`

SetOsId sets OsId field to given value.

### HasOsId

`func (o *CreateServer) HasOsId() bool`

HasOsId returns a boolean if a field has been set.

### GetImageId

`func (o *CreateServer) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *CreateServer) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *CreateServer) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *CreateServer) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetSoftwareId

`func (o *CreateServer) GetSoftwareId() float32`

GetSoftwareId returns the SoftwareId field if non-nil, zero value otherwise.

### GetSoftwareIdOk

`func (o *CreateServer) GetSoftwareIdOk() (*float32, bool)`

GetSoftwareIdOk returns a tuple with the SoftwareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareId

`func (o *CreateServer) SetSoftwareId(v float32)`

SetSoftwareId sets SoftwareId field to given value.

### HasSoftwareId

`func (o *CreateServer) HasSoftwareId() bool`

HasSoftwareId returns a boolean if a field has been set.

### GetPresetId

`func (o *CreateServer) GetPresetId() float32`

GetPresetId returns the PresetId field if non-nil, zero value otherwise.

### GetPresetIdOk

`func (o *CreateServer) GetPresetIdOk() (*float32, bool)`

GetPresetIdOk returns a tuple with the PresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetId

`func (o *CreateServer) SetPresetId(v float32)`

SetPresetId sets PresetId field to given value.

### HasPresetId

`func (o *CreateServer) HasPresetId() bool`

HasPresetId returns a boolean if a field has been set.

### GetBandwidth

`func (o *CreateServer) GetBandwidth() float32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *CreateServer) GetBandwidthOk() (*float32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *CreateServer) SetBandwidth(v float32)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *CreateServer) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetName

`func (o *CreateServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateServer) SetName(v string)`

SetName sets Name field to given value.


### GetAvatarId

`func (o *CreateServer) GetAvatarId() string`

GetAvatarId returns the AvatarId field if non-nil, zero value otherwise.

### GetAvatarIdOk

`func (o *CreateServer) GetAvatarIdOk() (*string, bool)`

GetAvatarIdOk returns a tuple with the AvatarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarId

`func (o *CreateServer) SetAvatarId(v string)`

SetAvatarId sets AvatarId field to given value.

### HasAvatarId

`func (o *CreateServer) HasAvatarId() bool`

HasAvatarId returns a boolean if a field has been set.

### GetComment

`func (o *CreateServer) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateServer) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateServer) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateServer) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetSshKeysIds

`func (o *CreateServer) GetSshKeysIds() []float32`

GetSshKeysIds returns the SshKeysIds field if non-nil, zero value otherwise.

### GetSshKeysIdsOk

`func (o *CreateServer) GetSshKeysIdsOk() (*[]float32, bool)`

GetSshKeysIdsOk returns a tuple with the SshKeysIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeysIds

`func (o *CreateServer) SetSshKeysIds(v []float32)`

SetSshKeysIds sets SshKeysIds field to given value.

### HasSshKeysIds

`func (o *CreateServer) HasSshKeysIds() bool`

HasSshKeysIds returns a boolean if a field has been set.

### GetIsLocalNetwork

`func (o *CreateServer) GetIsLocalNetwork() bool`

GetIsLocalNetwork returns the IsLocalNetwork field if non-nil, zero value otherwise.

### GetIsLocalNetworkOk

`func (o *CreateServer) GetIsLocalNetworkOk() (*bool, bool)`

GetIsLocalNetworkOk returns a tuple with the IsLocalNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocalNetwork

`func (o *CreateServer) SetIsLocalNetwork(v bool)`

SetIsLocalNetwork sets IsLocalNetwork field to given value.

### HasIsLocalNetwork

`func (o *CreateServer) HasIsLocalNetwork() bool`

HasIsLocalNetwork returns a boolean if a field has been set.

### GetNetwork

`func (o *CreateServer) GetNetwork() Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *CreateServer) GetNetworkOk() (*Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *CreateServer) SetNetwork(v Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *CreateServer) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetCloudInit

`func (o *CreateServer) GetCloudInit() string`

GetCloudInit returns the CloudInit field if non-nil, zero value otherwise.

### GetCloudInitOk

`func (o *CreateServer) GetCloudInitOk() (*string, bool)`

GetCloudInitOk returns a tuple with the CloudInit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInit

`func (o *CreateServer) SetCloudInit(v string)`

SetCloudInit sets CloudInit field to given value.

### HasCloudInit

`func (o *CreateServer) HasCloudInit() bool`

HasCloudInit returns a boolean if a field has been set.

### GetAvailabilityZone

`func (o *CreateServer) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *CreateServer) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *CreateServer) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *CreateServer) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


