# NetworkDrivePreset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | ID тарифа. | 
**CostPerGb** | **float32** | Стоимость тарифа сетевого диска. | 
**Min** | **float32** | Минимальный размер диска (в Гб). | 
**Max** | **float32** | Максимальный размер диска (в Гб). | 
**Step** | **float32** | Размер шага диска | 
**Type** | **string** | Тип сетевого диска. | 
**Read** | [**NetworkDrivePresetRead**](NetworkDrivePresetRead.md) |  | 
**Write** | [**NetworkDrivePresetWrite**](NetworkDrivePresetWrite.md) |  | 

## Methods

### NewNetworkDrivePreset

`func NewNetworkDrivePreset(id float32, costPerGb float32, min float32, max float32, step float32, type_ string, read NetworkDrivePresetRead, write NetworkDrivePresetWrite, ) *NetworkDrivePreset`

NewNetworkDrivePreset instantiates a new NetworkDrivePreset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDrivePresetWithDefaults

`func NewNetworkDrivePresetWithDefaults() *NetworkDrivePreset`

NewNetworkDrivePresetWithDefaults instantiates a new NetworkDrivePreset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkDrivePreset) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkDrivePreset) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkDrivePreset) SetId(v float32)`

SetId sets Id field to given value.


### GetCostPerGb

`func (o *NetworkDrivePreset) GetCostPerGb() float32`

GetCostPerGb returns the CostPerGb field if non-nil, zero value otherwise.

### GetCostPerGbOk

`func (o *NetworkDrivePreset) GetCostPerGbOk() (*float32, bool)`

GetCostPerGbOk returns a tuple with the CostPerGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerGb

`func (o *NetworkDrivePreset) SetCostPerGb(v float32)`

SetCostPerGb sets CostPerGb field to given value.


### GetMin

`func (o *NetworkDrivePreset) GetMin() float32`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *NetworkDrivePreset) GetMinOk() (*float32, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *NetworkDrivePreset) SetMin(v float32)`

SetMin sets Min field to given value.


### GetMax

`func (o *NetworkDrivePreset) GetMax() float32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *NetworkDrivePreset) GetMaxOk() (*float32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *NetworkDrivePreset) SetMax(v float32)`

SetMax sets Max field to given value.


### GetStep

`func (o *NetworkDrivePreset) GetStep() float32`

GetStep returns the Step field if non-nil, zero value otherwise.

### GetStepOk

`func (o *NetworkDrivePreset) GetStepOk() (*float32, bool)`

GetStepOk returns a tuple with the Step field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStep

`func (o *NetworkDrivePreset) SetStep(v float32)`

SetStep sets Step field to given value.


### GetType

`func (o *NetworkDrivePreset) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkDrivePreset) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkDrivePreset) SetType(v string)`

SetType sets Type field to given value.


### GetRead

`func (o *NetworkDrivePreset) GetRead() NetworkDrivePresetRead`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *NetworkDrivePreset) GetReadOk() (*NetworkDrivePresetRead, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *NetworkDrivePreset) SetRead(v NetworkDrivePresetRead)`

SetRead sets Read field to given value.


### GetWrite

`func (o *NetworkDrivePreset) GetWrite() NetworkDrivePresetWrite`

GetWrite returns the Write field if non-nil, zero value otherwise.

### GetWriteOk

`func (o *NetworkDrivePreset) GetWriteOk() (*NetworkDrivePresetWrite, bool)`

GetWriteOk returns a tuple with the Write field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrite

`func (o *NetworkDrivePreset) SetWrite(v NetworkDrivePresetWrite)`

SetWrite sets Write field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


