# ToolConfigField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | Ключ поля в &#x60;config_values&#x60; | 
**Label** | **string** | Название поля для пользователя | 
**Type** | **string** | Тип значения | 
**Required** | **bool** | Поле обязательно | 
**Description** | Pointer to **string** | Подсказка к полю | [optional] 
**Placeholder** | Pointer to **string** | Пример значения | [optional] 
**HeaderName** | Pointer to **string** | HTTP-заголовок, в котором значение передаётся MCP-серверу | [optional] 
**Prefix** | Pointer to **string** | Префикс, который добавляется перед значением в заголовке | [optional] 
**Options** | Pointer to **[]string** | Допустимые значения (для &#x60;select&#x60;) | [optional] 
**Default** | Pointer to **string** | Значение по умолчанию | [optional] 

## Methods

### NewToolConfigField

`func NewToolConfigField(key string, label string, type_ string, required bool, ) *ToolConfigField`

NewToolConfigField instantiates a new ToolConfigField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolConfigFieldWithDefaults

`func NewToolConfigFieldWithDefaults() *ToolConfigField`

NewToolConfigFieldWithDefaults instantiates a new ToolConfigField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ToolConfigField) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ToolConfigField) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ToolConfigField) SetKey(v string)`

SetKey sets Key field to given value.


### GetLabel

`func (o *ToolConfigField) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ToolConfigField) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ToolConfigField) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetType

`func (o *ToolConfigField) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ToolConfigField) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ToolConfigField) SetType(v string)`

SetType sets Type field to given value.


### GetRequired

`func (o *ToolConfigField) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ToolConfigField) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ToolConfigField) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetDescription

`func (o *ToolConfigField) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ToolConfigField) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ToolConfigField) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ToolConfigField) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPlaceholder

`func (o *ToolConfigField) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *ToolConfigField) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *ToolConfigField) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *ToolConfigField) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### GetHeaderName

`func (o *ToolConfigField) GetHeaderName() string`

GetHeaderName returns the HeaderName field if non-nil, zero value otherwise.

### GetHeaderNameOk

`func (o *ToolConfigField) GetHeaderNameOk() (*string, bool)`

GetHeaderNameOk returns a tuple with the HeaderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderName

`func (o *ToolConfigField) SetHeaderName(v string)`

SetHeaderName sets HeaderName field to given value.

### HasHeaderName

`func (o *ToolConfigField) HasHeaderName() bool`

HasHeaderName returns a boolean if a field has been set.

### GetPrefix

`func (o *ToolConfigField) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *ToolConfigField) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *ToolConfigField) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *ToolConfigField) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetOptions

`func (o *ToolConfigField) GetOptions() []string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ToolConfigField) GetOptionsOk() (*[]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ToolConfigField) SetOptions(v []string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ToolConfigField) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetDefault

`func (o *ToolConfigField) GetDefault() string`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *ToolConfigField) GetDefaultOk() (*string, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *ToolConfigField) SetDefault(v string)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *ToolConfigField) HasDefault() bool`

HasDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


