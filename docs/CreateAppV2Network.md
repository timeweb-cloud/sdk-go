# CreateAppV2Network

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID VPC, к которой необходимо подключить приложение. | [optional] 
**LocalIp** | Pointer to **string** | Локальный IP-адрес приложения в выбранной VPC. | [optional] 
**FloatingIp** | Pointer to **string** | Публичный IPv4-адрес свободного плавающего IP аккаунта или create_ip для создания нового IP (оплачивается отдельно). IP остаётся на аккаунте и привязывается к приложению. Доступно для аккаунтов, которым включена эта возможность; для остальных поле игнорируется. Если поле не передано, публичный IP выделяется автоматически. | [optional] 

## Methods

### NewCreateAppV2Network

`func NewCreateAppV2Network() *CreateAppV2Network`

NewCreateAppV2Network instantiates a new CreateAppV2Network object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAppV2NetworkWithDefaults

`func NewCreateAppV2NetworkWithDefaults() *CreateAppV2Network`

NewCreateAppV2NetworkWithDefaults instantiates a new CreateAppV2Network object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateAppV2Network) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateAppV2Network) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateAppV2Network) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateAppV2Network) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocalIp

`func (o *CreateAppV2Network) GetLocalIp() string`

GetLocalIp returns the LocalIp field if non-nil, zero value otherwise.

### GetLocalIpOk

`func (o *CreateAppV2Network) GetLocalIpOk() (*string, bool)`

GetLocalIpOk returns a tuple with the LocalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIp

`func (o *CreateAppV2Network) SetLocalIp(v string)`

SetLocalIp sets LocalIp field to given value.

### HasLocalIp

`func (o *CreateAppV2Network) HasLocalIp() bool`

HasLocalIp returns a boolean if a field has been set.

### GetFloatingIp

`func (o *CreateAppV2Network) GetFloatingIp() string`

GetFloatingIp returns the FloatingIp field if non-nil, zero value otherwise.

### GetFloatingIpOk

`func (o *CreateAppV2Network) GetFloatingIpOk() (*string, bool)`

GetFloatingIpOk returns a tuple with the FloatingIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatingIp

`func (o *CreateAppV2Network) SetFloatingIp(v string)`

SetFloatingIp sets FloatingIp field to given value.

### HasFloatingIp

`func (o *CreateAppV2Network) HasFloatingIp() bool`

HasFloatingIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


