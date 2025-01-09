# ServerDisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | ID диска. | 
**Size** | **float32** | Размер диска (в Мб). | 
**Used** | **float32** | Количество использованной памяти диска (в Мб). | 
**Type** | **string** | Тип диска. | 
**IsMounted** | **bool** | Является ли диск примонтированным. | 
**IsSystem** | **bool** | Является ли диск системным. | 
**SystemName** | **string** | Системное название диска. | 
**Status** | **string** | Статус диска. | 

## Methods

### NewServerDisk

`func NewServerDisk(id float32, size float32, used float32, type_ string, isMounted bool, isSystem bool, systemName string, status string, ) *ServerDisk`

NewServerDisk instantiates a new ServerDisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerDiskWithDefaults

`func NewServerDiskWithDefaults() *ServerDisk`

NewServerDiskWithDefaults instantiates a new ServerDisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServerDisk) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerDisk) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerDisk) SetId(v float32)`

SetId sets Id field to given value.


### GetSize

`func (o *ServerDisk) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ServerDisk) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ServerDisk) SetSize(v float32)`

SetSize sets Size field to given value.


### GetUsed

`func (o *ServerDisk) GetUsed() float32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *ServerDisk) GetUsedOk() (*float32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *ServerDisk) SetUsed(v float32)`

SetUsed sets Used field to given value.


### GetType

`func (o *ServerDisk) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServerDisk) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServerDisk) SetType(v string)`

SetType sets Type field to given value.


### GetIsMounted

`func (o *ServerDisk) GetIsMounted() bool`

GetIsMounted returns the IsMounted field if non-nil, zero value otherwise.

### GetIsMountedOk

`func (o *ServerDisk) GetIsMountedOk() (*bool, bool)`

GetIsMountedOk returns a tuple with the IsMounted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMounted

`func (o *ServerDisk) SetIsMounted(v bool)`

SetIsMounted sets IsMounted field to given value.


### GetIsSystem

`func (o *ServerDisk) GetIsSystem() bool`

GetIsSystem returns the IsSystem field if non-nil, zero value otherwise.

### GetIsSystemOk

`func (o *ServerDisk) GetIsSystemOk() (*bool, bool)`

GetIsSystemOk returns a tuple with the IsSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystem

`func (o *ServerDisk) SetIsSystem(v bool)`

SetIsSystem sets IsSystem field to given value.


### GetSystemName

`func (o *ServerDisk) GetSystemName() string`

GetSystemName returns the SystemName field if non-nil, zero value otherwise.

### GetSystemNameOk

`func (o *ServerDisk) GetSystemNameOk() (*string, bool)`

GetSystemNameOk returns a tuple with the SystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemName

`func (o *ServerDisk) SetSystemName(v string)`

SetSystemName sets SystemName field to given value.


### GetStatus

`func (o *ServerDisk) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServerDisk) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServerDisk) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


