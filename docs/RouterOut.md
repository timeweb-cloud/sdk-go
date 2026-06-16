# RouterOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID роутера | 
**AccountId** | **string** | ID аккаунта | 
**AvatarLink** | **NullableString** | Ссылка на аватар роутера | 
**Name** | **string** | Имя роутера | 
**Comment** | **NullableString** | Описание роутера | 
**Status** | **string** | Статус роутера | 
**Zone** | **string** | Зона доступности | 
**Ips** | [**[]RouterOutIpsInner**](RouterOutIpsInner.md) | IP-адреса | 
**PresetId** | **int32** | ID тарифа | 
**Preset** | [**RouterPreset**](RouterPreset.md) |  | 
**Nodes** | [**[]RouterOutNodesInner**](RouterOutNodesInner.md) | Ноды | 
**Networks** | [**[]RouterNetworkMeta**](RouterNetworkMeta.md) | Сети | 
**CreatedAt** | **time.Time** | Дата и время создания роутера в формате ISO8601 | 
**ProjectId** | Pointer to **int32** | ID проекта | [optional] 
**ParentServices** | [**[]RouterOutParentServicesInner**](RouterOutParentServicesInner.md) | Родительские сервисы | 

## Methods

### NewRouterOut

`func NewRouterOut(id string, accountId string, avatarLink NullableString, name string, comment NullableString, status string, zone string, ips []RouterOutIpsInner, presetId int32, preset RouterPreset, nodes []RouterOutNodesInner, networks []RouterNetworkMeta, createdAt time.Time, parentServices []RouterOutParentServicesInner, ) *RouterOut`

NewRouterOut instantiates a new RouterOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouterOutWithDefaults

`func NewRouterOutWithDefaults() *RouterOut`

NewRouterOutWithDefaults instantiates a new RouterOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RouterOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RouterOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RouterOut) SetId(v string)`

SetId sets Id field to given value.


### GetAccountId

`func (o *RouterOut) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *RouterOut) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *RouterOut) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAvatarLink

`func (o *RouterOut) GetAvatarLink() string`

GetAvatarLink returns the AvatarLink field if non-nil, zero value otherwise.

### GetAvatarLinkOk

`func (o *RouterOut) GetAvatarLinkOk() (*string, bool)`

GetAvatarLinkOk returns a tuple with the AvatarLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarLink

`func (o *RouterOut) SetAvatarLink(v string)`

SetAvatarLink sets AvatarLink field to given value.


### SetAvatarLinkNil

`func (o *RouterOut) SetAvatarLinkNil(b bool)`

 SetAvatarLinkNil sets the value for AvatarLink to be an explicit nil

### UnsetAvatarLink
`func (o *RouterOut) UnsetAvatarLink()`

UnsetAvatarLink ensures that no value is present for AvatarLink, not even an explicit nil
### GetName

`func (o *RouterOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RouterOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RouterOut) SetName(v string)`

SetName sets Name field to given value.


### GetComment

`func (o *RouterOut) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RouterOut) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RouterOut) SetComment(v string)`

SetComment sets Comment field to given value.


### SetCommentNil

`func (o *RouterOut) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *RouterOut) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetStatus

`func (o *RouterOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RouterOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RouterOut) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetZone

`func (o *RouterOut) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *RouterOut) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *RouterOut) SetZone(v string)`

SetZone sets Zone field to given value.


### GetIps

`func (o *RouterOut) GetIps() []RouterOutIpsInner`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *RouterOut) GetIpsOk() (*[]RouterOutIpsInner, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *RouterOut) SetIps(v []RouterOutIpsInner)`

SetIps sets Ips field to given value.


### GetPresetId

`func (o *RouterOut) GetPresetId() int32`

GetPresetId returns the PresetId field if non-nil, zero value otherwise.

### GetPresetIdOk

`func (o *RouterOut) GetPresetIdOk() (*int32, bool)`

GetPresetIdOk returns a tuple with the PresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetId

`func (o *RouterOut) SetPresetId(v int32)`

SetPresetId sets PresetId field to given value.


### GetPreset

`func (o *RouterOut) GetPreset() RouterPreset`

GetPreset returns the Preset field if non-nil, zero value otherwise.

### GetPresetOk

`func (o *RouterOut) GetPresetOk() (*RouterPreset, bool)`

GetPresetOk returns a tuple with the Preset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreset

`func (o *RouterOut) SetPreset(v RouterPreset)`

SetPreset sets Preset field to given value.


### GetNodes

`func (o *RouterOut) GetNodes() []RouterOutNodesInner`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *RouterOut) GetNodesOk() (*[]RouterOutNodesInner, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *RouterOut) SetNodes(v []RouterOutNodesInner)`

SetNodes sets Nodes field to given value.


### GetNetworks

`func (o *RouterOut) GetNetworks() []RouterNetworkMeta`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *RouterOut) GetNetworksOk() (*[]RouterNetworkMeta, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *RouterOut) SetNetworks(v []RouterNetworkMeta)`

SetNetworks sets Networks field to given value.


### GetCreatedAt

`func (o *RouterOut) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RouterOut) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RouterOut) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetProjectId

`func (o *RouterOut) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *RouterOut) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *RouterOut) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *RouterOut) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetParentServices

`func (o *RouterOut) GetParentServices() []RouterOutParentServicesInner`

GetParentServices returns the ParentServices field if non-nil, zero value otherwise.

### GetParentServicesOk

`func (o *RouterOut) GetParentServicesOk() (*[]RouterOutParentServicesInner, bool)`

GetParentServicesOk returns a tuple with the ParentServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentServices

`func (o *RouterOut) SetParentServices(v []RouterOutParentServicesInner)`

SetParentServices sets ParentServices field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


