# HttpResourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delivery** | Pointer to [**ConfigDelivery**](ConfigDelivery.md) |  | [optional] 
**Cache** | Pointer to [**ConfigCache**](ConfigCache.md) |  | [optional] 
**Origin** | Pointer to [**ConfigOrigin**](ConfigOrigin.md) |  | [optional] 
**Security** | Pointer to [**ConfigSecurity**](ConfigSecurity.md) |  | [optional] 
**HttpHeaders** | Pointer to [**ConfigHttpHeaders**](ConfigHttpHeaders.md) |  | [optional] 
**Access** | Pointer to [**ConfigAccess**](ConfigAccess.md) |  | [optional] 
**Domains** | Pointer to [**ConfigDomains**](ConfigDomains.md) |  | [optional] 
**Robots** | Pointer to [**ConfigRobots**](ConfigRobots.md) |  | [optional] 

## Methods

### NewHttpResourceConfig

`func NewHttpResourceConfig() *HttpResourceConfig`

NewHttpResourceConfig instantiates a new HttpResourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpResourceConfigWithDefaults

`func NewHttpResourceConfigWithDefaults() *HttpResourceConfig`

NewHttpResourceConfigWithDefaults instantiates a new HttpResourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelivery

`func (o *HttpResourceConfig) GetDelivery() ConfigDelivery`

GetDelivery returns the Delivery field if non-nil, zero value otherwise.

### GetDeliveryOk

`func (o *HttpResourceConfig) GetDeliveryOk() (*ConfigDelivery, bool)`

GetDeliveryOk returns a tuple with the Delivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivery

`func (o *HttpResourceConfig) SetDelivery(v ConfigDelivery)`

SetDelivery sets Delivery field to given value.

### HasDelivery

`func (o *HttpResourceConfig) HasDelivery() bool`

HasDelivery returns a boolean if a field has been set.

### GetCache

`func (o *HttpResourceConfig) GetCache() ConfigCache`

GetCache returns the Cache field if non-nil, zero value otherwise.

### GetCacheOk

`func (o *HttpResourceConfig) GetCacheOk() (*ConfigCache, bool)`

GetCacheOk returns a tuple with the Cache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache

`func (o *HttpResourceConfig) SetCache(v ConfigCache)`

SetCache sets Cache field to given value.

### HasCache

`func (o *HttpResourceConfig) HasCache() bool`

HasCache returns a boolean if a field has been set.

### GetOrigin

`func (o *HttpResourceConfig) GetOrigin() ConfigOrigin`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *HttpResourceConfig) GetOriginOk() (*ConfigOrigin, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *HttpResourceConfig) SetOrigin(v ConfigOrigin)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *HttpResourceConfig) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetSecurity

`func (o *HttpResourceConfig) GetSecurity() ConfigSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *HttpResourceConfig) GetSecurityOk() (*ConfigSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *HttpResourceConfig) SetSecurity(v ConfigSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *HttpResourceConfig) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetHttpHeaders

`func (o *HttpResourceConfig) GetHttpHeaders() ConfigHttpHeaders`

GetHttpHeaders returns the HttpHeaders field if non-nil, zero value otherwise.

### GetHttpHeadersOk

`func (o *HttpResourceConfig) GetHttpHeadersOk() (*ConfigHttpHeaders, bool)`

GetHttpHeadersOk returns a tuple with the HttpHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpHeaders

`func (o *HttpResourceConfig) SetHttpHeaders(v ConfigHttpHeaders)`

SetHttpHeaders sets HttpHeaders field to given value.

### HasHttpHeaders

`func (o *HttpResourceConfig) HasHttpHeaders() bool`

HasHttpHeaders returns a boolean if a field has been set.

### GetAccess

`func (o *HttpResourceConfig) GetAccess() ConfigAccess`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *HttpResourceConfig) GetAccessOk() (*ConfigAccess, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *HttpResourceConfig) SetAccess(v ConfigAccess)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *HttpResourceConfig) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetDomains

`func (o *HttpResourceConfig) GetDomains() ConfigDomains`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *HttpResourceConfig) GetDomainsOk() (*ConfigDomains, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *HttpResourceConfig) SetDomains(v ConfigDomains)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *HttpResourceConfig) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetRobots

`func (o *HttpResourceConfig) GetRobots() ConfigRobots`

GetRobots returns the Robots field if non-nil, zero value otherwise.

### GetRobotsOk

`func (o *HttpResourceConfig) GetRobotsOk() (*ConfigRobots, bool)`

GetRobotsOk returns a tuple with the Robots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRobots

`func (o *HttpResourceConfig) SetRobots(v ConfigRobots)`

SetRobots sets Robots field to given value.

### HasRobots

`func (o *HttpResourceConfig) HasRobots() bool`

HasRobots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


