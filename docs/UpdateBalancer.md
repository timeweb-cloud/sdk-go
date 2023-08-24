# UpdateBalancer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Удобочитаемое имя, установленное для балансировщика. | [optional] 
**Algo** | Pointer to **string** | Алгоритм переключений балансировщика. | [optional] 
**IsSticky** | Pointer to **bool** | Это логическое значение, которое показывает, сохраняется ли сессия. | [optional] 
**IsUseProxy** | Pointer to **bool** | Это логическое значение, которое показывает, выступает ли балансировщик в качестве прокси. | [optional] 
**IsSsl** | Pointer to **bool** | Это логическое значение, которое показывает, требуется ли перенаправление на SSL. | [optional] 
**IsKeepalive** | Pointer to **bool** | Это логическое значение, которое показывает, выдает ли балансировщик сигнал о проверке жизнеспособности. | [optional] 
**Proto** | Pointer to **string** | Протокол. | [optional] 
**Port** | Pointer to **float32** | Порт балансировщика. | [optional] 
**Path** | Pointer to **string** | Адрес балансировщика. | [optional] 
**Inter** | Pointer to **float32** | Интервал проверки. | [optional] 
**Timeout** | Pointer to **float32** | Таймаут ответа балансировщика. | [optional] 
**Fall** | Pointer to **float32** | Порог количества ошибок. | [optional] 
**Rise** | Pointer to **float32** | Порог количества успешных ответов. | [optional] 

## Methods

### NewUpdateBalancer

`func NewUpdateBalancer() *UpdateBalancer`

NewUpdateBalancer instantiates a new UpdateBalancer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBalancerWithDefaults

`func NewUpdateBalancerWithDefaults() *UpdateBalancer`

NewUpdateBalancerWithDefaults instantiates a new UpdateBalancer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateBalancer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateBalancer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateBalancer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateBalancer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAlgo

`func (o *UpdateBalancer) GetAlgo() string`

GetAlgo returns the Algo field if non-nil, zero value otherwise.

### GetAlgoOk

`func (o *UpdateBalancer) GetAlgoOk() (*string, bool)`

GetAlgoOk returns a tuple with the Algo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgo

`func (o *UpdateBalancer) SetAlgo(v string)`

SetAlgo sets Algo field to given value.

### HasAlgo

`func (o *UpdateBalancer) HasAlgo() bool`

HasAlgo returns a boolean if a field has been set.

### GetIsSticky

`func (o *UpdateBalancer) GetIsSticky() bool`

GetIsSticky returns the IsSticky field if non-nil, zero value otherwise.

### GetIsStickyOk

`func (o *UpdateBalancer) GetIsStickyOk() (*bool, bool)`

GetIsStickyOk returns a tuple with the IsSticky field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSticky

`func (o *UpdateBalancer) SetIsSticky(v bool)`

SetIsSticky sets IsSticky field to given value.

### HasIsSticky

`func (o *UpdateBalancer) HasIsSticky() bool`

HasIsSticky returns a boolean if a field has been set.

### GetIsUseProxy

`func (o *UpdateBalancer) GetIsUseProxy() bool`

GetIsUseProxy returns the IsUseProxy field if non-nil, zero value otherwise.

### GetIsUseProxyOk

`func (o *UpdateBalancer) GetIsUseProxyOk() (*bool, bool)`

GetIsUseProxyOk returns a tuple with the IsUseProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUseProxy

`func (o *UpdateBalancer) SetIsUseProxy(v bool)`

SetIsUseProxy sets IsUseProxy field to given value.

### HasIsUseProxy

`func (o *UpdateBalancer) HasIsUseProxy() bool`

HasIsUseProxy returns a boolean if a field has been set.

### GetIsSsl

`func (o *UpdateBalancer) GetIsSsl() bool`

GetIsSsl returns the IsSsl field if non-nil, zero value otherwise.

### GetIsSslOk

`func (o *UpdateBalancer) GetIsSslOk() (*bool, bool)`

GetIsSslOk returns a tuple with the IsSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSsl

`func (o *UpdateBalancer) SetIsSsl(v bool)`

SetIsSsl sets IsSsl field to given value.

### HasIsSsl

`func (o *UpdateBalancer) HasIsSsl() bool`

HasIsSsl returns a boolean if a field has been set.

### GetIsKeepalive

`func (o *UpdateBalancer) GetIsKeepalive() bool`

GetIsKeepalive returns the IsKeepalive field if non-nil, zero value otherwise.

### GetIsKeepaliveOk

`func (o *UpdateBalancer) GetIsKeepaliveOk() (*bool, bool)`

GetIsKeepaliveOk returns a tuple with the IsKeepalive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKeepalive

`func (o *UpdateBalancer) SetIsKeepalive(v bool)`

SetIsKeepalive sets IsKeepalive field to given value.

### HasIsKeepalive

`func (o *UpdateBalancer) HasIsKeepalive() bool`

HasIsKeepalive returns a boolean if a field has been set.

### GetProto

`func (o *UpdateBalancer) GetProto() string`

GetProto returns the Proto field if non-nil, zero value otherwise.

### GetProtoOk

`func (o *UpdateBalancer) GetProtoOk() (*string, bool)`

GetProtoOk returns a tuple with the Proto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProto

`func (o *UpdateBalancer) SetProto(v string)`

SetProto sets Proto field to given value.

### HasProto

`func (o *UpdateBalancer) HasProto() bool`

HasProto returns a boolean if a field has been set.

### GetPort

`func (o *UpdateBalancer) GetPort() float32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *UpdateBalancer) GetPortOk() (*float32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *UpdateBalancer) SetPort(v float32)`

SetPort sets Port field to given value.

### HasPort

`func (o *UpdateBalancer) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPath

`func (o *UpdateBalancer) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *UpdateBalancer) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *UpdateBalancer) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *UpdateBalancer) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetInter

`func (o *UpdateBalancer) GetInter() float32`

GetInter returns the Inter field if non-nil, zero value otherwise.

### GetInterOk

`func (o *UpdateBalancer) GetInterOk() (*float32, bool)`

GetInterOk returns a tuple with the Inter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInter

`func (o *UpdateBalancer) SetInter(v float32)`

SetInter sets Inter field to given value.

### HasInter

`func (o *UpdateBalancer) HasInter() bool`

HasInter returns a boolean if a field has been set.

### GetTimeout

`func (o *UpdateBalancer) GetTimeout() float32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *UpdateBalancer) GetTimeoutOk() (*float32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *UpdateBalancer) SetTimeout(v float32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *UpdateBalancer) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetFall

`func (o *UpdateBalancer) GetFall() float32`

GetFall returns the Fall field if non-nil, zero value otherwise.

### GetFallOk

`func (o *UpdateBalancer) GetFallOk() (*float32, bool)`

GetFallOk returns a tuple with the Fall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFall

`func (o *UpdateBalancer) SetFall(v float32)`

SetFall sets Fall field to given value.

### HasFall

`func (o *UpdateBalancer) HasFall() bool`

HasFall returns a boolean if a field has been set.

### GetRise

`func (o *UpdateBalancer) GetRise() float32`

GetRise returns the Rise field if non-nil, zero value otherwise.

### GetRiseOk

`func (o *UpdateBalancer) GetRiseOk() (*float32, bool)`

GetRiseOk returns a tuple with the Rise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRise

`func (o *UpdateBalancer) SetRise(v float32)`

SetRise sets Rise field to given value.

### HasRise

`func (o *UpdateBalancer) HasRise() bool`

HasRise returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


