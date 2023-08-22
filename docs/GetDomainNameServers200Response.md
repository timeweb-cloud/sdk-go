# GetDomainNameServers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**NameServers** | [**[]DomainNameServer**](DomainNameServer.md) |  | 

## Methods

### NewGetDomainNameServers200Response

`func NewGetDomainNameServers200Response(meta Meta, nameServers []DomainNameServer, ) *GetDomainNameServers200Response`

NewGetDomainNameServers200Response instantiates a new GetDomainNameServers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDomainNameServers200ResponseWithDefaults

`func NewGetDomainNameServers200ResponseWithDefaults() *GetDomainNameServers200Response`

NewGetDomainNameServers200ResponseWithDefaults instantiates a new GetDomainNameServers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetDomainNameServers200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetDomainNameServers200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetDomainNameServers200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetNameServers

`func (o *GetDomainNameServers200Response) GetNameServers() []DomainNameServer`

GetNameServers returns the NameServers field if non-nil, zero value otherwise.

### GetNameServersOk

`func (o *GetDomainNameServers200Response) GetNameServersOk() (*[]DomainNameServer, bool)`

GetNameServersOk returns a tuple with the NameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameServers

`func (o *GetDomainNameServers200Response) SetNameServers(v []DomainNameServer)`

SetNameServers sets NameServers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


