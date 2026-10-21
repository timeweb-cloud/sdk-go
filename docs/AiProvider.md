# AiProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID провайдера | 
**Name** | **string** | Имя провайдера | 
**Type** | **string** | &#x60;external&#x60; — внешний провайдер, &#x60;internal&#x60; — модели, размещённые в Timeweb Cloud | 

## Methods

### NewAiProvider

`func NewAiProvider(id int32, name string, type_ string, ) *AiProvider`

NewAiProvider instantiates a new AiProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiProviderWithDefaults

`func NewAiProviderWithDefaults() *AiProvider`

NewAiProviderWithDefaults instantiates a new AiProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AiProvider) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiProvider) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiProvider) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *AiProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AiProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AiProvider) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *AiProvider) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AiProvider) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AiProvider) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


