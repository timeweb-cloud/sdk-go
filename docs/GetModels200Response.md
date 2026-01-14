# GetModels200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Models** | [**[]Model**](Model.md) |  | 
**Meta** | [**GetModels200ResponseMeta**](GetModels200ResponseMeta.md) |  | 

## Methods

### NewGetModels200Response

`func NewGetModels200Response(models []Model, meta GetModels200ResponseMeta, ) *GetModels200Response`

NewGetModels200Response instantiates a new GetModels200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetModels200ResponseWithDefaults

`func NewGetModels200ResponseWithDefaults() *GetModels200Response`

NewGetModels200ResponseWithDefaults instantiates a new GetModels200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModels

`func (o *GetModels200Response) GetModels() []Model`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *GetModels200Response) GetModelsOk() (*[]Model, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *GetModels200Response) SetModels(v []Model)`

SetModels sets Models field to given value.


### GetMeta

`func (o *GetModels200Response) GetMeta() GetModels200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetModels200Response) GetMetaOk() (*GetModels200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetModels200Response) SetMeta(v GetModels200ResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


