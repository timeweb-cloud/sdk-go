# GetAiKeys200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AiKeys** | [**[]AiKey**](AiKey.md) |  | 
**Meta** | [**GetAiKeys200ResponseMeta**](GetAiKeys200ResponseMeta.md) |  | 

## Methods

### NewGetAiKeys200Response

`func NewGetAiKeys200Response(aiKeys []AiKey, meta GetAiKeys200ResponseMeta, ) *GetAiKeys200Response`

NewGetAiKeys200Response instantiates a new GetAiKeys200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAiKeys200ResponseWithDefaults

`func NewGetAiKeys200ResponseWithDefaults() *GetAiKeys200Response`

NewGetAiKeys200ResponseWithDefaults instantiates a new GetAiKeys200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAiKeys

`func (o *GetAiKeys200Response) GetAiKeys() []AiKey`

GetAiKeys returns the AiKeys field if non-nil, zero value otherwise.

### GetAiKeysOk

`func (o *GetAiKeys200Response) GetAiKeysOk() (*[]AiKey, bool)`

GetAiKeysOk returns a tuple with the AiKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiKeys

`func (o *GetAiKeys200Response) SetAiKeys(v []AiKey)`

SetAiKeys sets AiKeys field to given value.


### GetMeta

`func (o *GetAiKeys200Response) GetMeta() GetAiKeys200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetAiKeys200Response) GetMetaOk() (*GetAiKeys200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetAiKeys200Response) SetMeta(v GetAiKeys200ResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


