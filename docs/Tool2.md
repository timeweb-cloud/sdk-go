# Tool2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Имя инструмента | 
**Title** | Pointer to **string** | Короткое название инструмента | [optional] 
**Description** | Pointer to **string** | Что делает инструмент. Описания изменяющих инструментов, как правило, начинаются с &#x60;[WRITE]&#x60;, у платных добавляется &#x60;[BILLABLE]&#x60; | [optional] 
**InputSchema** | **map[string]interface{}** | JSON Schema аргументов инструмента | 
**Annotations** | Pointer to [**Tool2Annotations**](Tool2Annotations.md) |  | [optional] 

## Methods

### NewTool2

`func NewTool2(name string, inputSchema map[string]interface{}, ) *Tool2`

NewTool2 instantiates a new Tool2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTool2WithDefaults

`func NewTool2WithDefaults() *Tool2`

NewTool2WithDefaults instantiates a new Tool2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Tool2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Tool2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Tool2) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *Tool2) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Tool2) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Tool2) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Tool2) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *Tool2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Tool2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Tool2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Tool2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInputSchema

`func (o *Tool2) GetInputSchema() map[string]interface{}`

GetInputSchema returns the InputSchema field if non-nil, zero value otherwise.

### GetInputSchemaOk

`func (o *Tool2) GetInputSchemaOk() (*map[string]interface{}, bool)`

GetInputSchemaOk returns a tuple with the InputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSchema

`func (o *Tool2) SetInputSchema(v map[string]interface{})`

SetInputSchema sets InputSchema field to given value.


### GetAnnotations

`func (o *Tool2) GetAnnotations() Tool2Annotations`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *Tool2) GetAnnotationsOk() (*Tool2Annotations, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *Tool2) SetAnnotations(v Tool2Annotations)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *Tool2) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


