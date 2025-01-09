# VpcService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | ID сервисв. | 
**Name** | **string** | Имя сервиса. | 
**PublicIp** | Pointer to **string** | Публичный IP-адрес сервиса | [optional] 
**LocalIp** | Pointer to **string** | Приватный IP-адрес сервиса | [optional] 
**Type** | Pointer to **string** | Тип сервиса. | [optional] 

## Methods

### NewVpcService

`func NewVpcService(id float32, name string, ) *VpcService`

NewVpcService instantiates a new VpcService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpcServiceWithDefaults

`func NewVpcServiceWithDefaults() *VpcService`

NewVpcServiceWithDefaults instantiates a new VpcService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VpcService) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VpcService) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VpcService) SetId(v float32)`

SetId sets Id field to given value.


### GetName

`func (o *VpcService) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VpcService) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VpcService) SetName(v string)`

SetName sets Name field to given value.


### GetPublicIp

`func (o *VpcService) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *VpcService) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *VpcService) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *VpcService) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetLocalIp

`func (o *VpcService) GetLocalIp() string`

GetLocalIp returns the LocalIp field if non-nil, zero value otherwise.

### GetLocalIpOk

`func (o *VpcService) GetLocalIpOk() (*string, bool)`

GetLocalIpOk returns a tuple with the LocalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIp

`func (o *VpcService) SetLocalIp(v string)`

SetLocalIp sets LocalIp field to given value.

### HasLocalIp

`func (o *VpcService) HasLocalIp() bool`

HasLocalIp returns a boolean if a field has been set.

### GetType

`func (o *VpcService) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VpcService) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VpcService) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VpcService) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


