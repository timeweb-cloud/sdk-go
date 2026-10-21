# GetAiProviders200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Providers** | [**[]AiProvider**](AiProvider.md) |  | 
**Meta** | [**GetAiProviders200ResponseMeta**](GetAiProviders200ResponseMeta.md) |  | 

## Methods

### NewGetAiProviders200Response

`func NewGetAiProviders200Response(providers []AiProvider, meta GetAiProviders200ResponseMeta, ) *GetAiProviders200Response`

NewGetAiProviders200Response instantiates a new GetAiProviders200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAiProviders200ResponseWithDefaults

`func NewGetAiProviders200ResponseWithDefaults() *GetAiProviders200Response`

NewGetAiProviders200ResponseWithDefaults instantiates a new GetAiProviders200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviders

`func (o *GetAiProviders200Response) GetProviders() []AiProvider`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *GetAiProviders200Response) GetProvidersOk() (*[]AiProvider, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *GetAiProviders200Response) SetProviders(v []AiProvider)`

SetProviders sets Providers field to given value.


### GetMeta

`func (o *GetAiProviders200Response) GetMeta() GetAiProviders200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetAiProviders200Response) GetMetaOk() (*GetAiProviders200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetAiProviders200Response) SetMeta(v GetAiProviders200ResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


