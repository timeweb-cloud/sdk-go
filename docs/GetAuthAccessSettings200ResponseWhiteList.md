# GetAuthAccessSettings200ResponseWhiteList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ips** | **[]string** | Список разрешенных IP-адресов. | 
**Countries** | **[]string** | Список разрешенных стран. | 

## Methods

### NewGetAuthAccessSettings200ResponseWhiteList

`func NewGetAuthAccessSettings200ResponseWhiteList(ips []string, countries []string, ) *GetAuthAccessSettings200ResponseWhiteList`

NewGetAuthAccessSettings200ResponseWhiteList instantiates a new GetAuthAccessSettings200ResponseWhiteList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAuthAccessSettings200ResponseWhiteListWithDefaults

`func NewGetAuthAccessSettings200ResponseWhiteListWithDefaults() *GetAuthAccessSettings200ResponseWhiteList`

NewGetAuthAccessSettings200ResponseWhiteListWithDefaults instantiates a new GetAuthAccessSettings200ResponseWhiteList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIps

`func (o *GetAuthAccessSettings200ResponseWhiteList) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *GetAuthAccessSettings200ResponseWhiteList) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *GetAuthAccessSettings200ResponseWhiteList) SetIps(v []string)`

SetIps sets Ips field to given value.


### GetCountries

`func (o *GetAuthAccessSettings200ResponseWhiteList) GetCountries() []string`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *GetAuthAccessSettings200ResponseWhiteList) GetCountriesOk() (*[]string, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *GetAuthAccessSettings200ResponseWhiteList) SetCountries(v []string)`

SetCountries sets Countries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


