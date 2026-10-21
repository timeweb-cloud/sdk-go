# TextContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Тип блока | 
**Text** | **string** | Результат инструмента — как правило, JSON, сериализованный в строку. Для изменяющих инструментов — объект подтверждения со статусом (см. «Подтверждение изменений» в описании раздела) | 

## Methods

### NewTextContent

`func NewTextContent(type_ string, text string, ) *TextContent`

NewTextContent instantiates a new TextContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTextContentWithDefaults

`func NewTextContentWithDefaults() *TextContent`

NewTextContentWithDefaults instantiates a new TextContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TextContent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TextContent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TextContent) SetType(v string)`

SetType sets Type field to given value.


### GetText

`func (o *TextContent) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *TextContent) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *TextContent) SetText(v string)`

SetText sets Text field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


