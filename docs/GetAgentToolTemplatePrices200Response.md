# GetAgentToolTemplatePrices200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SystemToolPrices** | [**[]ToolPrice**](ToolPrice.md) |  | 
**Meta** | [**GetAgentTools200ResponseMeta**](GetAgentTools200ResponseMeta.md) |  | 

## Methods

### NewGetAgentToolTemplatePrices200Response

`func NewGetAgentToolTemplatePrices200Response(systemToolPrices []ToolPrice, meta GetAgentTools200ResponseMeta, ) *GetAgentToolTemplatePrices200Response`

NewGetAgentToolTemplatePrices200Response instantiates a new GetAgentToolTemplatePrices200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAgentToolTemplatePrices200ResponseWithDefaults

`func NewGetAgentToolTemplatePrices200ResponseWithDefaults() *GetAgentToolTemplatePrices200Response`

NewGetAgentToolTemplatePrices200ResponseWithDefaults instantiates a new GetAgentToolTemplatePrices200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSystemToolPrices

`func (o *GetAgentToolTemplatePrices200Response) GetSystemToolPrices() []ToolPrice`

GetSystemToolPrices returns the SystemToolPrices field if non-nil, zero value otherwise.

### GetSystemToolPricesOk

`func (o *GetAgentToolTemplatePrices200Response) GetSystemToolPricesOk() (*[]ToolPrice, bool)`

GetSystemToolPricesOk returns a tuple with the SystemToolPrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemToolPrices

`func (o *GetAgentToolTemplatePrices200Response) SetSystemToolPrices(v []ToolPrice)`

SetSystemToolPrices sets SystemToolPrices field to given value.


### GetMeta

`func (o *GetAgentToolTemplatePrices200Response) GetMeta() GetAgentTools200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetAgentToolTemplatePrices200Response) GetMetaOk() (*GetAgentTools200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetAgentToolTemplatePrices200Response) SetMeta(v GetAgentTools200ResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


