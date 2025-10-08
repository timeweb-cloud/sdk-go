# GetServicePrices200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServicesCosts** | [**[]ServicePrice**](ServicePrice.md) | Список сервисов с информацией о стоимости | 
**Meta** | [**GetServicePrices200ResponseMeta**](GetServicePrices200ResponseMeta.md) |  | 

## Methods

### NewGetServicePrices200Response

`func NewGetServicePrices200Response(servicesCosts []ServicePrice, meta GetServicePrices200ResponseMeta, ) *GetServicePrices200Response`

NewGetServicePrices200Response instantiates a new GetServicePrices200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServicePrices200ResponseWithDefaults

`func NewGetServicePrices200ResponseWithDefaults() *GetServicePrices200Response`

NewGetServicePrices200ResponseWithDefaults instantiates a new GetServicePrices200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServicesCosts

`func (o *GetServicePrices200Response) GetServicesCosts() []ServicePrice`

GetServicesCosts returns the ServicesCosts field if non-nil, zero value otherwise.

### GetServicesCostsOk

`func (o *GetServicePrices200Response) GetServicesCostsOk() (*[]ServicePrice, bool)`

GetServicesCostsOk returns a tuple with the ServicesCosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicesCosts

`func (o *GetServicePrices200Response) SetServicesCosts(v []ServicePrice)`

SetServicesCosts sets ServicesCosts field to given value.


### GetMeta

`func (o *GetServicePrices200Response) GetMeta() GetServicePrices200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetServicePrices200Response) GetMetaOk() (*GetServicePrices200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetServicePrices200Response) SetMeta(v GetServicePrices200ResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


