# CreateServerNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID сети | [optional] 
**FloatingIp** | Pointer to **string** | Публичный IP | [optional] 
**LocalIp** | Pointer to **string** | Приватный IP | [optional] 
**Ip** | Pointer to **string** | Приватный IP | [optional] 
**NetworkDriveIds** | Pointer to **[]string** | Массив ID сетевых дисков | [optional] 

## Methods

### NewCreateServerNetwork

`func NewCreateServerNetwork() *CreateServerNetwork`

NewCreateServerNetwork instantiates a new CreateServerNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateServerNetworkWithDefaults

`func NewCreateServerNetworkWithDefaults() *CreateServerNetwork`

NewCreateServerNetworkWithDefaults instantiates a new CreateServerNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateServerNetwork) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateServerNetwork) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateServerNetwork) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateServerNetwork) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFloatingIp

`func (o *CreateServerNetwork) GetFloatingIp() string`

GetFloatingIp returns the FloatingIp field if non-nil, zero value otherwise.

### GetFloatingIpOk

`func (o *CreateServerNetwork) GetFloatingIpOk() (*string, bool)`

GetFloatingIpOk returns a tuple with the FloatingIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatingIp

`func (o *CreateServerNetwork) SetFloatingIp(v string)`

SetFloatingIp sets FloatingIp field to given value.

### HasFloatingIp

`func (o *CreateServerNetwork) HasFloatingIp() bool`

HasFloatingIp returns a boolean if a field has been set.

### GetLocalIp

`func (o *CreateServerNetwork) GetLocalIp() string`

GetLocalIp returns the LocalIp field if non-nil, zero value otherwise.

### GetLocalIpOk

`func (o *CreateServerNetwork) GetLocalIpOk() (*string, bool)`

GetLocalIpOk returns a tuple with the LocalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIp

`func (o *CreateServerNetwork) SetLocalIp(v string)`

SetLocalIp sets LocalIp field to given value.

### HasLocalIp

`func (o *CreateServerNetwork) HasLocalIp() bool`

HasLocalIp returns a boolean if a field has been set.

### GetIp

`func (o *CreateServerNetwork) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *CreateServerNetwork) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *CreateServerNetwork) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *CreateServerNetwork) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetNetworkDriveIds

`func (o *CreateServerNetwork) GetNetworkDriveIds() []string`

GetNetworkDriveIds returns the NetworkDriveIds field if non-nil, zero value otherwise.

### GetNetworkDriveIdsOk

`func (o *CreateServerNetwork) GetNetworkDriveIdsOk() (*[]string, bool)`

GetNetworkDriveIdsOk returns a tuple with the NetworkDriveIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDriveIds

`func (o *CreateServerNetwork) SetNetworkDriveIds(v []string)`

SetNetworkDriveIds sets NetworkDriveIds field to given value.

### HasNetworkDriveIds

`func (o *CreateServerNetwork) HasNetworkDriveIds() bool`

HasNetworkDriveIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


