# GetTokens200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**ApiKeys** | [**[]ApiKey**](ApiKey.md) |  | 

## Methods

### NewGetTokens200Response

`func NewGetTokens200Response(meta Meta, apiKeys []ApiKey, ) *GetTokens200Response`

NewGetTokens200Response instantiates a new GetTokens200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTokens200ResponseWithDefaults

`func NewGetTokens200ResponseWithDefaults() *GetTokens200Response`

NewGetTokens200ResponseWithDefaults instantiates a new GetTokens200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTokens200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTokens200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTokens200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetApiKeys

`func (o *GetTokens200Response) GetApiKeys() []ApiKey`

GetApiKeys returns the ApiKeys field if non-nil, zero value otherwise.

### GetApiKeysOk

`func (o *GetTokens200Response) GetApiKeysOk() (*[]ApiKey, bool)`

GetApiKeysOk returns a tuple with the ApiKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeys

`func (o *GetTokens200Response) SetApiKeys(v []ApiKey)`

SetApiKeys sets ApiKeys field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


