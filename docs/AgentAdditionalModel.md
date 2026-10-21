# AgentAdditionalModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModelId** | **int32** | ID модели | 
**Width** | Pointer to **int32** | Ширина генерируемого изображения в пикселях (для моделей генерации изображений). Если не передана — значение по умолчанию для модели. | [optional] 
**Height** | Pointer to **int32** | Высота генерируемого изображения в пикселях (для моделей генерации изображений). Если не передана — значение по умолчанию для модели. | [optional] 

## Methods

### NewAgentAdditionalModel

`func NewAgentAdditionalModel(modelId int32, ) *AgentAdditionalModel`

NewAgentAdditionalModel instantiates a new AgentAdditionalModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentAdditionalModelWithDefaults

`func NewAgentAdditionalModelWithDefaults() *AgentAdditionalModel`

NewAgentAdditionalModelWithDefaults instantiates a new AgentAdditionalModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModelId

`func (o *AgentAdditionalModel) GetModelId() int32`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *AgentAdditionalModel) GetModelIdOk() (*int32, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *AgentAdditionalModel) SetModelId(v int32)`

SetModelId sets ModelId field to given value.


### GetWidth

`func (o *AgentAdditionalModel) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *AgentAdditionalModel) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *AgentAdditionalModel) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *AgentAdditionalModel) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *AgentAdditionalModel) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *AgentAdditionalModel) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *AgentAdditionalModel) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *AgentAdditionalModel) HasHeight() bool`

HasHeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


