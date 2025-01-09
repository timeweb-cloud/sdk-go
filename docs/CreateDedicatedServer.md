# CreateDedicatedServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanId** | **float32** | ID списка дополнительных услуг выделенного сервера. | 
**PresetId** | **float32** | ID тарифа выделенного сервера. | 
**OsId** | Pointer to **NullableFloat32** | ID операционной системы, которая будет установлена на выделенный сервер. | [optional] 
**CpId** | Pointer to **NullableFloat32** | ID панели управления, которая будет установлена на выделенный сервер. | [optional] 
**BandwidthId** | Pointer to **float32** | ID интернет-канала, который будет установлен на выделенный сервер. | [optional] 
**NetworkDriveId** | Pointer to **float32** | ID сетевого диска, который будет установлен на выделенный сервер. | [optional] 
**AdditionalIpAddrId** | Pointer to **NullableFloat32** | ID дополнительного IP-адреса, который будет установлен на выделенный сервер. | [optional] 
**PaymentPeriod** | **string** | Период оплаты. | 
**Name** | **string** | Удобочитаемое имя выделенного сервера. Максимальная длина — 255 символов, имя должно быть уникальным. | 
**Comment** | Pointer to **string** | Комментарий к выделенному серверу. Максимальная длина — 255 символов. | [optional] 

## Methods

### NewCreateDedicatedServer

`func NewCreateDedicatedServer(planId float32, presetId float32, paymentPeriod string, name string, ) *CreateDedicatedServer`

NewCreateDedicatedServer instantiates a new CreateDedicatedServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDedicatedServerWithDefaults

`func NewCreateDedicatedServerWithDefaults() *CreateDedicatedServer`

NewCreateDedicatedServerWithDefaults instantiates a new CreateDedicatedServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlanId

`func (o *CreateDedicatedServer) GetPlanId() float32`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *CreateDedicatedServer) GetPlanIdOk() (*float32, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *CreateDedicatedServer) SetPlanId(v float32)`

SetPlanId sets PlanId field to given value.


### GetPresetId

`func (o *CreateDedicatedServer) GetPresetId() float32`

GetPresetId returns the PresetId field if non-nil, zero value otherwise.

### GetPresetIdOk

`func (o *CreateDedicatedServer) GetPresetIdOk() (*float32, bool)`

GetPresetIdOk returns a tuple with the PresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetId

`func (o *CreateDedicatedServer) SetPresetId(v float32)`

SetPresetId sets PresetId field to given value.


### GetOsId

`func (o *CreateDedicatedServer) GetOsId() float32`

GetOsId returns the OsId field if non-nil, zero value otherwise.

### GetOsIdOk

`func (o *CreateDedicatedServer) GetOsIdOk() (*float32, bool)`

GetOsIdOk returns a tuple with the OsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsId

`func (o *CreateDedicatedServer) SetOsId(v float32)`

SetOsId sets OsId field to given value.

### HasOsId

`func (o *CreateDedicatedServer) HasOsId() bool`

HasOsId returns a boolean if a field has been set.

### SetOsIdNil

`func (o *CreateDedicatedServer) SetOsIdNil(b bool)`

 SetOsIdNil sets the value for OsId to be an explicit nil

### UnsetOsId
`func (o *CreateDedicatedServer) UnsetOsId()`

UnsetOsId ensures that no value is present for OsId, not even an explicit nil
### GetCpId

`func (o *CreateDedicatedServer) GetCpId() float32`

GetCpId returns the CpId field if non-nil, zero value otherwise.

### GetCpIdOk

`func (o *CreateDedicatedServer) GetCpIdOk() (*float32, bool)`

GetCpIdOk returns a tuple with the CpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpId

`func (o *CreateDedicatedServer) SetCpId(v float32)`

SetCpId sets CpId field to given value.

### HasCpId

`func (o *CreateDedicatedServer) HasCpId() bool`

HasCpId returns a boolean if a field has been set.

### SetCpIdNil

`func (o *CreateDedicatedServer) SetCpIdNil(b bool)`

 SetCpIdNil sets the value for CpId to be an explicit nil

### UnsetCpId
`func (o *CreateDedicatedServer) UnsetCpId()`

UnsetCpId ensures that no value is present for CpId, not even an explicit nil
### GetBandwidthId

`func (o *CreateDedicatedServer) GetBandwidthId() float32`

GetBandwidthId returns the BandwidthId field if non-nil, zero value otherwise.

### GetBandwidthIdOk

`func (o *CreateDedicatedServer) GetBandwidthIdOk() (*float32, bool)`

GetBandwidthIdOk returns a tuple with the BandwidthId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthId

`func (o *CreateDedicatedServer) SetBandwidthId(v float32)`

SetBandwidthId sets BandwidthId field to given value.

### HasBandwidthId

`func (o *CreateDedicatedServer) HasBandwidthId() bool`

HasBandwidthId returns a boolean if a field has been set.

### GetNetworkDriveId

`func (o *CreateDedicatedServer) GetNetworkDriveId() float32`

GetNetworkDriveId returns the NetworkDriveId field if non-nil, zero value otherwise.

### GetNetworkDriveIdOk

`func (o *CreateDedicatedServer) GetNetworkDriveIdOk() (*float32, bool)`

GetNetworkDriveIdOk returns a tuple with the NetworkDriveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDriveId

`func (o *CreateDedicatedServer) SetNetworkDriveId(v float32)`

SetNetworkDriveId sets NetworkDriveId field to given value.

### HasNetworkDriveId

`func (o *CreateDedicatedServer) HasNetworkDriveId() bool`

HasNetworkDriveId returns a boolean if a field has been set.

### GetAdditionalIpAddrId

`func (o *CreateDedicatedServer) GetAdditionalIpAddrId() float32`

GetAdditionalIpAddrId returns the AdditionalIpAddrId field if non-nil, zero value otherwise.

### GetAdditionalIpAddrIdOk

`func (o *CreateDedicatedServer) GetAdditionalIpAddrIdOk() (*float32, bool)`

GetAdditionalIpAddrIdOk returns a tuple with the AdditionalIpAddrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalIpAddrId

`func (o *CreateDedicatedServer) SetAdditionalIpAddrId(v float32)`

SetAdditionalIpAddrId sets AdditionalIpAddrId field to given value.

### HasAdditionalIpAddrId

`func (o *CreateDedicatedServer) HasAdditionalIpAddrId() bool`

HasAdditionalIpAddrId returns a boolean if a field has been set.

### SetAdditionalIpAddrIdNil

`func (o *CreateDedicatedServer) SetAdditionalIpAddrIdNil(b bool)`

 SetAdditionalIpAddrIdNil sets the value for AdditionalIpAddrId to be an explicit nil

### UnsetAdditionalIpAddrId
`func (o *CreateDedicatedServer) UnsetAdditionalIpAddrId()`

UnsetAdditionalIpAddrId ensures that no value is present for AdditionalIpAddrId, not even an explicit nil
### GetPaymentPeriod

`func (o *CreateDedicatedServer) GetPaymentPeriod() string`

GetPaymentPeriod returns the PaymentPeriod field if non-nil, zero value otherwise.

### GetPaymentPeriodOk

`func (o *CreateDedicatedServer) GetPaymentPeriodOk() (*string, bool)`

GetPaymentPeriodOk returns a tuple with the PaymentPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentPeriod

`func (o *CreateDedicatedServer) SetPaymentPeriod(v string)`

SetPaymentPeriod sets PaymentPeriod field to given value.


### GetName

`func (o *CreateDedicatedServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDedicatedServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDedicatedServer) SetName(v string)`

SetName sets Name field to given value.


### GetComment

`func (o *CreateDedicatedServer) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateDedicatedServer) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateDedicatedServer) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateDedicatedServer) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


