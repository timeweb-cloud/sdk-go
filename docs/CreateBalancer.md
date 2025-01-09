# CreateBalancer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Удобочитаемое имя, установленное для балансировщика. | 
**Algo** | **string** | Алгоритм переключений балансировщика. | 
**IsSticky** | **bool** | Это логическое значение, которое показывает, сохраняется ли сессия. | 
**IsUseProxy** | **bool** | Это логическое значение, которое показывает, выступает ли балансировщик в качестве прокси. | 
**IsSsl** | **bool** | Это логическое значение, которое показывает, требуется ли перенаправление на SSL. | 
**IsKeepalive** | **bool** | Это логическое значение, которое показывает, выдает ли балансировщик сигнал о проверке жизнеспособности. | 
**Proto** | **string** | Протокол. | 
**Port** | **float32** | Порт балансировщика. | 
**Path** | **string** | Адрес балансировщика. | 
**Inter** | **float32** | Интервал проверки. | 
**Timeout** | **float32** | Таймаут ответа балансировщика. | 
**Fall** | **float32** | Порог количества ошибок. | 
**Rise** | **float32** | Порог количества успешных ответов. | 
**PresetId** | **float32** | ID тарифа. | 
**Network** | Pointer to [**Network**](Network.md) |  | [optional] 
**AvailabilityZone** | Pointer to [**AvailabilityZone**](AvailabilityZone.md) |  | [optional] 

## Methods

### NewCreateBalancer

`func NewCreateBalancer(name string, algo string, isSticky bool, isUseProxy bool, isSsl bool, isKeepalive bool, proto string, port float32, path string, inter float32, timeout float32, fall float32, rise float32, presetId float32, ) *CreateBalancer`

NewCreateBalancer instantiates a new CreateBalancer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBalancerWithDefaults

`func NewCreateBalancerWithDefaults() *CreateBalancer`

NewCreateBalancerWithDefaults instantiates a new CreateBalancer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateBalancer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateBalancer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateBalancer) SetName(v string)`

SetName sets Name field to given value.


### GetAlgo

`func (o *CreateBalancer) GetAlgo() string`

GetAlgo returns the Algo field if non-nil, zero value otherwise.

### GetAlgoOk

`func (o *CreateBalancer) GetAlgoOk() (*string, bool)`

GetAlgoOk returns a tuple with the Algo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgo

`func (o *CreateBalancer) SetAlgo(v string)`

SetAlgo sets Algo field to given value.


### GetIsSticky

`func (o *CreateBalancer) GetIsSticky() bool`

GetIsSticky returns the IsSticky field if non-nil, zero value otherwise.

### GetIsStickyOk

`func (o *CreateBalancer) GetIsStickyOk() (*bool, bool)`

GetIsStickyOk returns a tuple with the IsSticky field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSticky

`func (o *CreateBalancer) SetIsSticky(v bool)`

SetIsSticky sets IsSticky field to given value.


### GetIsUseProxy

`func (o *CreateBalancer) GetIsUseProxy() bool`

GetIsUseProxy returns the IsUseProxy field if non-nil, zero value otherwise.

### GetIsUseProxyOk

`func (o *CreateBalancer) GetIsUseProxyOk() (*bool, bool)`

GetIsUseProxyOk returns a tuple with the IsUseProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUseProxy

`func (o *CreateBalancer) SetIsUseProxy(v bool)`

SetIsUseProxy sets IsUseProxy field to given value.


### GetIsSsl

`func (o *CreateBalancer) GetIsSsl() bool`

GetIsSsl returns the IsSsl field if non-nil, zero value otherwise.

### GetIsSslOk

`func (o *CreateBalancer) GetIsSslOk() (*bool, bool)`

GetIsSslOk returns a tuple with the IsSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSsl

`func (o *CreateBalancer) SetIsSsl(v bool)`

SetIsSsl sets IsSsl field to given value.


### GetIsKeepalive

`func (o *CreateBalancer) GetIsKeepalive() bool`

GetIsKeepalive returns the IsKeepalive field if non-nil, zero value otherwise.

### GetIsKeepaliveOk

`func (o *CreateBalancer) GetIsKeepaliveOk() (*bool, bool)`

GetIsKeepaliveOk returns a tuple with the IsKeepalive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKeepalive

`func (o *CreateBalancer) SetIsKeepalive(v bool)`

SetIsKeepalive sets IsKeepalive field to given value.


### GetProto

`func (o *CreateBalancer) GetProto() string`

GetProto returns the Proto field if non-nil, zero value otherwise.

### GetProtoOk

`func (o *CreateBalancer) GetProtoOk() (*string, bool)`

GetProtoOk returns a tuple with the Proto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProto

`func (o *CreateBalancer) SetProto(v string)`

SetProto sets Proto field to given value.


### GetPort

`func (o *CreateBalancer) GetPort() float32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CreateBalancer) GetPortOk() (*float32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CreateBalancer) SetPort(v float32)`

SetPort sets Port field to given value.


### GetPath

`func (o *CreateBalancer) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CreateBalancer) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CreateBalancer) SetPath(v string)`

SetPath sets Path field to given value.


### GetInter

`func (o *CreateBalancer) GetInter() float32`

GetInter returns the Inter field if non-nil, zero value otherwise.

### GetInterOk

`func (o *CreateBalancer) GetInterOk() (*float32, bool)`

GetInterOk returns a tuple with the Inter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInter

`func (o *CreateBalancer) SetInter(v float32)`

SetInter sets Inter field to given value.


### GetTimeout

`func (o *CreateBalancer) GetTimeout() float32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *CreateBalancer) GetTimeoutOk() (*float32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *CreateBalancer) SetTimeout(v float32)`

SetTimeout sets Timeout field to given value.


### GetFall

`func (o *CreateBalancer) GetFall() float32`

GetFall returns the Fall field if non-nil, zero value otherwise.

### GetFallOk

`func (o *CreateBalancer) GetFallOk() (*float32, bool)`

GetFallOk returns a tuple with the Fall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFall

`func (o *CreateBalancer) SetFall(v float32)`

SetFall sets Fall field to given value.


### GetRise

`func (o *CreateBalancer) GetRise() float32`

GetRise returns the Rise field if non-nil, zero value otherwise.

### GetRiseOk

`func (o *CreateBalancer) GetRiseOk() (*float32, bool)`

GetRiseOk returns a tuple with the Rise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRise

`func (o *CreateBalancer) SetRise(v float32)`

SetRise sets Rise field to given value.


### GetPresetId

`func (o *CreateBalancer) GetPresetId() float32`

GetPresetId returns the PresetId field if non-nil, zero value otherwise.

### GetPresetIdOk

`func (o *CreateBalancer) GetPresetIdOk() (*float32, bool)`

GetPresetIdOk returns a tuple with the PresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetId

`func (o *CreateBalancer) SetPresetId(v float32)`

SetPresetId sets PresetId field to given value.


### GetNetwork

`func (o *CreateBalancer) GetNetwork() Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *CreateBalancer) GetNetworkOk() (*Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *CreateBalancer) SetNetwork(v Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *CreateBalancer) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetAvailabilityZone

`func (o *CreateBalancer) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *CreateBalancer) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *CreateBalancer) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *CreateBalancer) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


