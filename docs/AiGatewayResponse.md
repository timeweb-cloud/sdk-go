# AiGatewayResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID ответа. Передавайте его как есть в остальные запросы &#x60;/v1/responses&#x60; | [optional] 
**Object** | Pointer to **string** | Тип объекта | [optional] 
**CreatedAt** | Pointer to **int32** | Время создания, Unix-время в секундах | [optional] 
**Status** | Pointer to **string** | Статус ответа: &#x60;queued&#x60;, &#x60;in_progress&#x60;, &#x60;completed&#x60;, &#x60;incomplete&#x60; (достигнут лимит выходных токенов), &#x60;failed&#x60;, &#x60;cancelling&#x60; (отмена выполняется), &#x60;cancelled&#x60; | [optional] 
**Model** | Pointer to **string** | Модель, которая ответила | [optional] 
**Output** | Pointer to **[]map[string]interface{}** | Выходные элементы: сообщения, вызовы функций, рассуждения | [optional] 
**IncompleteDetails** | Pointer to [**NullableAiGatewayResponseIncompleteDetails**](AiGatewayResponseIncompleteDetails.md) |  | [optional] 
**Usage** | Pointer to [**AiGatewayResponseUsage**](AiGatewayResponseUsage.md) |  | [optional] 

## Methods

### NewAiGatewayResponse

`func NewAiGatewayResponse() *AiGatewayResponse`

NewAiGatewayResponse instantiates a new AiGatewayResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiGatewayResponseWithDefaults

`func NewAiGatewayResponseWithDefaults() *AiGatewayResponse`

NewAiGatewayResponseWithDefaults instantiates a new AiGatewayResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AiGatewayResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiGatewayResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiGatewayResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AiGatewayResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObject

`func (o *AiGatewayResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *AiGatewayResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *AiGatewayResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *AiGatewayResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AiGatewayResponse) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AiGatewayResponse) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AiGatewayResponse) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AiGatewayResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetStatus

`func (o *AiGatewayResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AiGatewayResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AiGatewayResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AiGatewayResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetModel

`func (o *AiGatewayResponse) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AiGatewayResponse) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AiGatewayResponse) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AiGatewayResponse) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetOutput

`func (o *AiGatewayResponse) GetOutput() []map[string]interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *AiGatewayResponse) GetOutputOk() (*[]map[string]interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *AiGatewayResponse) SetOutput(v []map[string]interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *AiGatewayResponse) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetIncompleteDetails

`func (o *AiGatewayResponse) GetIncompleteDetails() AiGatewayResponseIncompleteDetails`

GetIncompleteDetails returns the IncompleteDetails field if non-nil, zero value otherwise.

### GetIncompleteDetailsOk

`func (o *AiGatewayResponse) GetIncompleteDetailsOk() (*AiGatewayResponseIncompleteDetails, bool)`

GetIncompleteDetailsOk returns a tuple with the IncompleteDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncompleteDetails

`func (o *AiGatewayResponse) SetIncompleteDetails(v AiGatewayResponseIncompleteDetails)`

SetIncompleteDetails sets IncompleteDetails field to given value.

### HasIncompleteDetails

`func (o *AiGatewayResponse) HasIncompleteDetails() bool`

HasIncompleteDetails returns a boolean if a field has been set.

### SetIncompleteDetailsNil

`func (o *AiGatewayResponse) SetIncompleteDetailsNil(b bool)`

 SetIncompleteDetailsNil sets the value for IncompleteDetails to be an explicit nil

### UnsetIncompleteDetails
`func (o *AiGatewayResponse) UnsetIncompleteDetails()`

UnsetIncompleteDetails ensures that no value is present for IncompleteDetails, not even an explicit nil
### GetUsage

`func (o *AiGatewayResponse) GetUsage() AiGatewayResponseUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *AiGatewayResponse) GetUsageOk() (*AiGatewayResponseUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *AiGatewayResponse) SetUsage(v AiGatewayResponseUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *AiGatewayResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


