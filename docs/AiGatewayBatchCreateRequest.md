# AiGatewayBatchCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputFileId** | **string** | ID файла, загруженного с &#x60;purpose&#x3D;batch&#x60; | 
**Endpoint** | **string** | Эндпоинт запросов батча — должен совпадать с &#x60;url&#x60; каждой строки входного файла | 
**CompletionWindow** | **string** | Срок выполнения | 
**Metadata** | Pointer to **map[string]string** | Ваши метки батча: пары «строка — строка» | [optional] 

## Methods

### NewAiGatewayBatchCreateRequest

`func NewAiGatewayBatchCreateRequest(inputFileId string, endpoint string, completionWindow string, ) *AiGatewayBatchCreateRequest`

NewAiGatewayBatchCreateRequest instantiates a new AiGatewayBatchCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiGatewayBatchCreateRequestWithDefaults

`func NewAiGatewayBatchCreateRequestWithDefaults() *AiGatewayBatchCreateRequest`

NewAiGatewayBatchCreateRequestWithDefaults instantiates a new AiGatewayBatchCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputFileId

`func (o *AiGatewayBatchCreateRequest) GetInputFileId() string`

GetInputFileId returns the InputFileId field if non-nil, zero value otherwise.

### GetInputFileIdOk

`func (o *AiGatewayBatchCreateRequest) GetInputFileIdOk() (*string, bool)`

GetInputFileIdOk returns a tuple with the InputFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFileId

`func (o *AiGatewayBatchCreateRequest) SetInputFileId(v string)`

SetInputFileId sets InputFileId field to given value.


### GetEndpoint

`func (o *AiGatewayBatchCreateRequest) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *AiGatewayBatchCreateRequest) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *AiGatewayBatchCreateRequest) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetCompletionWindow

`func (o *AiGatewayBatchCreateRequest) GetCompletionWindow() string`

GetCompletionWindow returns the CompletionWindow field if non-nil, zero value otherwise.

### GetCompletionWindowOk

`func (o *AiGatewayBatchCreateRequest) GetCompletionWindowOk() (*string, bool)`

GetCompletionWindowOk returns a tuple with the CompletionWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionWindow

`func (o *AiGatewayBatchCreateRequest) SetCompletionWindow(v string)`

SetCompletionWindow sets CompletionWindow field to given value.


### GetMetadata

`func (o *AiGatewayBatchCreateRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AiGatewayBatchCreateRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AiGatewayBatchCreateRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AiGatewayBatchCreateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


