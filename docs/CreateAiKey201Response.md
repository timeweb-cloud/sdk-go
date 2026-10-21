# CreateAiKey201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AiKey** | [**AiKey**](AiKey.md) |  | 
**Key** | **string** | Значение ключа для запросов к AI Gateway API. Показывается только в этом ответе — сохраните его | 

## Methods

### NewCreateAiKey201Response

`func NewCreateAiKey201Response(aiKey AiKey, key string, ) *CreateAiKey201Response`

NewCreateAiKey201Response instantiates a new CreateAiKey201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAiKey201ResponseWithDefaults

`func NewCreateAiKey201ResponseWithDefaults() *CreateAiKey201Response`

NewCreateAiKey201ResponseWithDefaults instantiates a new CreateAiKey201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAiKey

`func (o *CreateAiKey201Response) GetAiKey() AiKey`

GetAiKey returns the AiKey field if non-nil, zero value otherwise.

### GetAiKeyOk

`func (o *CreateAiKey201Response) GetAiKeyOk() (*AiKey, bool)`

GetAiKeyOk returns a tuple with the AiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiKey

`func (o *CreateAiKey201Response) SetAiKey(v AiKey)`

SetAiKey sets AiKey field to given value.


### GetKey

`func (o *CreateAiKey201Response) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateAiKey201Response) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateAiKey201Response) SetKey(v string)`

SetKey sets Key field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


