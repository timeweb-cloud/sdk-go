# ConfigRobots

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Режим отдачи &#x60;robots.txt&#x60;. - &#x60;deny&#x60; — CDN отдает &#x60;robots.txt&#x60;, запрещающий индексацию; - &#x60;cached&#x60; — &#x60;robots.txt&#x60; берется с источника контента; - &#x60;custom&#x60; — CDN отдает содержимое из поля &#x60;content&#x60;. | 
**Content** | Pointer to **string** | Содержимое &#x60;robots.txt&#x60;. Обязательно и учитывается только при &#x60;type&#x60; &#x3D; &#x60;custom&#x60;. | [optional] 

## Methods

### NewConfigRobots

`func NewConfigRobots(type_ string, ) *ConfigRobots`

NewConfigRobots instantiates a new ConfigRobots object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigRobotsWithDefaults

`func NewConfigRobotsWithDefaults() *ConfigRobots`

NewConfigRobotsWithDefaults instantiates a new ConfigRobots object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ConfigRobots) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConfigRobots) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConfigRobots) SetType(v string)`

SetType sets Type field to given value.


### GetContent

`func (o *ConfigRobots) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ConfigRobots) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ConfigRobots) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ConfigRobots) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


