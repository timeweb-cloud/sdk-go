# GetBalancersPresets200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**BalancersPresets** | [**[]PresetsBalancer**](PresetsBalancer.md) |  | 
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 

## Methods

### NewGetBalancersPresets200Response

`func NewGetBalancersPresets200Response(meta Meta, balancersPresets []PresetsBalancer, responseId string, ) *GetBalancersPresets200Response`

NewGetBalancersPresets200Response instantiates a new GetBalancersPresets200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBalancersPresets200ResponseWithDefaults

`func NewGetBalancersPresets200ResponseWithDefaults() *GetBalancersPresets200Response`

NewGetBalancersPresets200ResponseWithDefaults instantiates a new GetBalancersPresets200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetBalancersPresets200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetBalancersPresets200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetBalancersPresets200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetBalancersPresets

`func (o *GetBalancersPresets200Response) GetBalancersPresets() []PresetsBalancer`

GetBalancersPresets returns the BalancersPresets field if non-nil, zero value otherwise.

### GetBalancersPresetsOk

`func (o *GetBalancersPresets200Response) GetBalancersPresetsOk() (*[]PresetsBalancer, bool)`

GetBalancersPresetsOk returns a tuple with the BalancersPresets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalancersPresets

`func (o *GetBalancersPresets200Response) SetBalancersPresets(v []PresetsBalancer)`

SetBalancersPresets sets BalancersPresets field to given value.


### GetResponseId

`func (o *GetBalancersPresets200Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *GetBalancersPresets200Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *GetBalancersPresets200Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


