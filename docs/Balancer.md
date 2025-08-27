# Balancer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | ID для каждого экземпляра балансировщика. Автоматически генерируется при создании. | 
**AccountId** | Pointer to **string** | ID пользователя. | [optional] 
**Algo** | **string** | Алгоритм переключений балансировщика. | 
**CreatedAt** | **time.Time** | Значение времени, указанное в комбинированном формате даты и времени ISO8601, которое представляет, когда был создан балансировщик. | 
**Fall** | **float32** | Порог количества ошибок. | 
**Inter** | **float32** | Интервал проверки. | 
**Ip** | **NullableString** | IP-адрес сетевого интерфейса IPv4. | 
**LocalIp** | **NullableString** | Локальный IP-адрес сетевого интерфейса IPv4. | 
**IsKeepalive** | **bool** | Это логическое значение, которое показывает, выдает ли балансировщик сигнал о проверке жизнеспособности. | 
**Name** | **string** | Удобочитаемое имя, установленное для балансировщика. | 
**Path** | **string** | Адрес балансировщика. | 
**Port** | **float32** | Порт балансировщика. | 
**Proto** | **string** | Протокол. | 
**Rise** | **float32** | Порог количества успешных ответов. | 
**Maxconn** | **float32** | Максимальное количество соединений. | 
**ConnectTimeout** | **float32** | Таймаут подключения. | 
**ClientTimeout** | **float32** | Таймаут клиента. | 
**ServerTimeout** | **float32** | Таймаут сервера. | 
**HttprequestTimeout** | **float32** | Таймаут HTTP запроса. | 
**PresetId** | **float32** | ID тарифа. | 
**IsSsl** | **bool** | Это логическое значение, которое показывает, требуется ли перенаправление на SSL. | 
**Status** | **string** | Статус балансировщика. | 
**IsSticky** | **bool** | Это логическое значение, которое показывает, сохраняется ли сессия. | 
**Timeout** | **float32** | Таймаут ответа балансировщика. | 
**AvatarLink** | **NullableString** | Ссылка на аватар балансировщика. | 
**IsUseProxy** | **bool** | Это логическое значение, которое показывает, выступает ли балансировщик в качестве прокси. | 
**Rules** | [**[]Rule**](Rule.md) |  | 
**Ips** | **[]string** | Список IP-адресов, привязанных к балансировщику | 
**Location** | **string** | Географическое расположение балансировщика | 
**AvailabilityZone** | [**AvailabilityZone**](AvailabilityZone.md) |  | 
**ProjectId** | **int32** | ID проекта | 
**Networks** | [**[]BalancerNetworksInner**](BalancerNetworksInner.md) | Список сетей сервера. | 

## Methods

### NewBalancer

`func NewBalancer(id float32, algo string, createdAt time.Time, fall float32, inter float32, ip NullableString, localIp NullableString, isKeepalive bool, name string, path string, port float32, proto string, rise float32, maxconn float32, connectTimeout float32, clientTimeout float32, serverTimeout float32, httprequestTimeout float32, presetId float32, isSsl bool, status string, isSticky bool, timeout float32, avatarLink NullableString, isUseProxy bool, rules []Rule, ips []string, location string, availabilityZone AvailabilityZone, projectId int32, networks []BalancerNetworksInner, ) *Balancer`

NewBalancer instantiates a new Balancer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalancerWithDefaults

`func NewBalancerWithDefaults() *Balancer`

NewBalancerWithDefaults instantiates a new Balancer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Balancer) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Balancer) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Balancer) SetId(v float32)`

SetId sets Id field to given value.


### GetAccountId

`func (o *Balancer) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Balancer) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Balancer) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *Balancer) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAlgo

`func (o *Balancer) GetAlgo() string`

GetAlgo returns the Algo field if non-nil, zero value otherwise.

### GetAlgoOk

`func (o *Balancer) GetAlgoOk() (*string, bool)`

GetAlgoOk returns a tuple with the Algo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgo

`func (o *Balancer) SetAlgo(v string)`

SetAlgo sets Algo field to given value.


### GetCreatedAt

`func (o *Balancer) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Balancer) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Balancer) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetFall

`func (o *Balancer) GetFall() float32`

GetFall returns the Fall field if non-nil, zero value otherwise.

### GetFallOk

`func (o *Balancer) GetFallOk() (*float32, bool)`

GetFallOk returns a tuple with the Fall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFall

`func (o *Balancer) SetFall(v float32)`

SetFall sets Fall field to given value.


### GetInter

`func (o *Balancer) GetInter() float32`

GetInter returns the Inter field if non-nil, zero value otherwise.

### GetInterOk

`func (o *Balancer) GetInterOk() (*float32, bool)`

GetInterOk returns a tuple with the Inter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInter

`func (o *Balancer) SetInter(v float32)`

SetInter sets Inter field to given value.


### GetIp

`func (o *Balancer) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *Balancer) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *Balancer) SetIp(v string)`

SetIp sets Ip field to given value.


### SetIpNil

`func (o *Balancer) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *Balancer) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetLocalIp

`func (o *Balancer) GetLocalIp() string`

GetLocalIp returns the LocalIp field if non-nil, zero value otherwise.

### GetLocalIpOk

`func (o *Balancer) GetLocalIpOk() (*string, bool)`

GetLocalIpOk returns a tuple with the LocalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIp

`func (o *Balancer) SetLocalIp(v string)`

SetLocalIp sets LocalIp field to given value.


### SetLocalIpNil

`func (o *Balancer) SetLocalIpNil(b bool)`

 SetLocalIpNil sets the value for LocalIp to be an explicit nil

### UnsetLocalIp
`func (o *Balancer) UnsetLocalIp()`

UnsetLocalIp ensures that no value is present for LocalIp, not even an explicit nil
### GetIsKeepalive

`func (o *Balancer) GetIsKeepalive() bool`

GetIsKeepalive returns the IsKeepalive field if non-nil, zero value otherwise.

### GetIsKeepaliveOk

`func (o *Balancer) GetIsKeepaliveOk() (*bool, bool)`

GetIsKeepaliveOk returns a tuple with the IsKeepalive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKeepalive

`func (o *Balancer) SetIsKeepalive(v bool)`

SetIsKeepalive sets IsKeepalive field to given value.


### GetName

`func (o *Balancer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Balancer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Balancer) SetName(v string)`

SetName sets Name field to given value.


### GetPath

`func (o *Balancer) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Balancer) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Balancer) SetPath(v string)`

SetPath sets Path field to given value.


### GetPort

`func (o *Balancer) GetPort() float32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Balancer) GetPortOk() (*float32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Balancer) SetPort(v float32)`

SetPort sets Port field to given value.


### GetProto

`func (o *Balancer) GetProto() string`

GetProto returns the Proto field if non-nil, zero value otherwise.

### GetProtoOk

`func (o *Balancer) GetProtoOk() (*string, bool)`

GetProtoOk returns a tuple with the Proto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProto

`func (o *Balancer) SetProto(v string)`

SetProto sets Proto field to given value.


### GetRise

`func (o *Balancer) GetRise() float32`

GetRise returns the Rise field if non-nil, zero value otherwise.

### GetRiseOk

`func (o *Balancer) GetRiseOk() (*float32, bool)`

GetRiseOk returns a tuple with the Rise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRise

`func (o *Balancer) SetRise(v float32)`

SetRise sets Rise field to given value.


### GetMaxconn

`func (o *Balancer) GetMaxconn() float32`

GetMaxconn returns the Maxconn field if non-nil, zero value otherwise.

### GetMaxconnOk

`func (o *Balancer) GetMaxconnOk() (*float32, bool)`

GetMaxconnOk returns a tuple with the Maxconn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxconn

`func (o *Balancer) SetMaxconn(v float32)`

SetMaxconn sets Maxconn field to given value.


### GetConnectTimeout

`func (o *Balancer) GetConnectTimeout() float32`

GetConnectTimeout returns the ConnectTimeout field if non-nil, zero value otherwise.

### GetConnectTimeoutOk

`func (o *Balancer) GetConnectTimeoutOk() (*float32, bool)`

GetConnectTimeoutOk returns a tuple with the ConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectTimeout

`func (o *Balancer) SetConnectTimeout(v float32)`

SetConnectTimeout sets ConnectTimeout field to given value.


### GetClientTimeout

`func (o *Balancer) GetClientTimeout() float32`

GetClientTimeout returns the ClientTimeout field if non-nil, zero value otherwise.

### GetClientTimeoutOk

`func (o *Balancer) GetClientTimeoutOk() (*float32, bool)`

GetClientTimeoutOk returns a tuple with the ClientTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTimeout

`func (o *Balancer) SetClientTimeout(v float32)`

SetClientTimeout sets ClientTimeout field to given value.


### GetServerTimeout

`func (o *Balancer) GetServerTimeout() float32`

GetServerTimeout returns the ServerTimeout field if non-nil, zero value otherwise.

### GetServerTimeoutOk

`func (o *Balancer) GetServerTimeoutOk() (*float32, bool)`

GetServerTimeoutOk returns a tuple with the ServerTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTimeout

`func (o *Balancer) SetServerTimeout(v float32)`

SetServerTimeout sets ServerTimeout field to given value.


### GetHttprequestTimeout

`func (o *Balancer) GetHttprequestTimeout() float32`

GetHttprequestTimeout returns the HttprequestTimeout field if non-nil, zero value otherwise.

### GetHttprequestTimeoutOk

`func (o *Balancer) GetHttprequestTimeoutOk() (*float32, bool)`

GetHttprequestTimeoutOk returns a tuple with the HttprequestTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttprequestTimeout

`func (o *Balancer) SetHttprequestTimeout(v float32)`

SetHttprequestTimeout sets HttprequestTimeout field to given value.


### GetPresetId

`func (o *Balancer) GetPresetId() float32`

GetPresetId returns the PresetId field if non-nil, zero value otherwise.

### GetPresetIdOk

`func (o *Balancer) GetPresetIdOk() (*float32, bool)`

GetPresetIdOk returns a tuple with the PresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetId

`func (o *Balancer) SetPresetId(v float32)`

SetPresetId sets PresetId field to given value.


### GetIsSsl

`func (o *Balancer) GetIsSsl() bool`

GetIsSsl returns the IsSsl field if non-nil, zero value otherwise.

### GetIsSslOk

`func (o *Balancer) GetIsSslOk() (*bool, bool)`

GetIsSslOk returns a tuple with the IsSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSsl

`func (o *Balancer) SetIsSsl(v bool)`

SetIsSsl sets IsSsl field to given value.


### GetStatus

`func (o *Balancer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Balancer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Balancer) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetIsSticky

`func (o *Balancer) GetIsSticky() bool`

GetIsSticky returns the IsSticky field if non-nil, zero value otherwise.

### GetIsStickyOk

`func (o *Balancer) GetIsStickyOk() (*bool, bool)`

GetIsStickyOk returns a tuple with the IsSticky field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSticky

`func (o *Balancer) SetIsSticky(v bool)`

SetIsSticky sets IsSticky field to given value.


### GetTimeout

`func (o *Balancer) GetTimeout() float32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *Balancer) GetTimeoutOk() (*float32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *Balancer) SetTimeout(v float32)`

SetTimeout sets Timeout field to given value.


### GetAvatarLink

`func (o *Balancer) GetAvatarLink() string`

GetAvatarLink returns the AvatarLink field if non-nil, zero value otherwise.

### GetAvatarLinkOk

`func (o *Balancer) GetAvatarLinkOk() (*string, bool)`

GetAvatarLinkOk returns a tuple with the AvatarLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarLink

`func (o *Balancer) SetAvatarLink(v string)`

SetAvatarLink sets AvatarLink field to given value.


### SetAvatarLinkNil

`func (o *Balancer) SetAvatarLinkNil(b bool)`

 SetAvatarLinkNil sets the value for AvatarLink to be an explicit nil

### UnsetAvatarLink
`func (o *Balancer) UnsetAvatarLink()`

UnsetAvatarLink ensures that no value is present for AvatarLink, not even an explicit nil
### GetIsUseProxy

`func (o *Balancer) GetIsUseProxy() bool`

GetIsUseProxy returns the IsUseProxy field if non-nil, zero value otherwise.

### GetIsUseProxyOk

`func (o *Balancer) GetIsUseProxyOk() (*bool, bool)`

GetIsUseProxyOk returns a tuple with the IsUseProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUseProxy

`func (o *Balancer) SetIsUseProxy(v bool)`

SetIsUseProxy sets IsUseProxy field to given value.


### GetRules

`func (o *Balancer) GetRules() []Rule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *Balancer) GetRulesOk() (*[]Rule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *Balancer) SetRules(v []Rule)`

SetRules sets Rules field to given value.


### GetIps

`func (o *Balancer) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *Balancer) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *Balancer) SetIps(v []string)`

SetIps sets Ips field to given value.


### GetLocation

`func (o *Balancer) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Balancer) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Balancer) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetAvailabilityZone

`func (o *Balancer) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *Balancer) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *Balancer) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.


### GetProjectId

`func (o *Balancer) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Balancer) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Balancer) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetNetworks

`func (o *Balancer) GetNetworks() []BalancerNetworksInner`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *Balancer) GetNetworksOk() (*[]BalancerNetworksInner, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *Balancer) SetNetworks(v []BalancerNetworksInner)`

SetNetworks sets Networks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


