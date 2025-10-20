# AgentSettingsWidget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WhitelistDomains** | **[]string** | Массив разрешенных доменов для виджета | 
**Name** | **string** | Отображаемое имя агента в виджете | 
**Signature** | Pointer to **string** | Подпись/подзаголовок, отображаемый под именем агента в виджете | [optional] 
**WelcomeMessage** | **string** | Приветственное сообщение, показываемое при открытии виджета | 
**PrimaryColor** | **string** | Основной цвет виджета (hex-код цвета в формате #RRGGBB) | 
**Font** | **string** | Семейство шрифтов для виджета | 
**IconUrl** | Pointer to **string** | URL иконки виджета | [optional] 
**ChatPosition** | **string** | Позиция виджета чата на странице | 

## Methods

### NewAgentSettingsWidget

`func NewAgentSettingsWidget(whitelistDomains []string, name string, welcomeMessage string, primaryColor string, font string, chatPosition string, ) *AgentSettingsWidget`

NewAgentSettingsWidget instantiates a new AgentSettingsWidget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentSettingsWidgetWithDefaults

`func NewAgentSettingsWidgetWithDefaults() *AgentSettingsWidget`

NewAgentSettingsWidgetWithDefaults instantiates a new AgentSettingsWidget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWhitelistDomains

`func (o *AgentSettingsWidget) GetWhitelistDomains() []string`

GetWhitelistDomains returns the WhitelistDomains field if non-nil, zero value otherwise.

### GetWhitelistDomainsOk

`func (o *AgentSettingsWidget) GetWhitelistDomainsOk() (*[]string, bool)`

GetWhitelistDomainsOk returns a tuple with the WhitelistDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelistDomains

`func (o *AgentSettingsWidget) SetWhitelistDomains(v []string)`

SetWhitelistDomains sets WhitelistDomains field to given value.


### GetName

`func (o *AgentSettingsWidget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentSettingsWidget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentSettingsWidget) SetName(v string)`

SetName sets Name field to given value.


### GetSignature

`func (o *AgentSettingsWidget) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *AgentSettingsWidget) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *AgentSettingsWidget) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *AgentSettingsWidget) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetWelcomeMessage

`func (o *AgentSettingsWidget) GetWelcomeMessage() string`

GetWelcomeMessage returns the WelcomeMessage field if non-nil, zero value otherwise.

### GetWelcomeMessageOk

`func (o *AgentSettingsWidget) GetWelcomeMessageOk() (*string, bool)`

GetWelcomeMessageOk returns a tuple with the WelcomeMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWelcomeMessage

`func (o *AgentSettingsWidget) SetWelcomeMessage(v string)`

SetWelcomeMessage sets WelcomeMessage field to given value.


### GetPrimaryColor

`func (o *AgentSettingsWidget) GetPrimaryColor() string`

GetPrimaryColor returns the PrimaryColor field if non-nil, zero value otherwise.

### GetPrimaryColorOk

`func (o *AgentSettingsWidget) GetPrimaryColorOk() (*string, bool)`

GetPrimaryColorOk returns a tuple with the PrimaryColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryColor

`func (o *AgentSettingsWidget) SetPrimaryColor(v string)`

SetPrimaryColor sets PrimaryColor field to given value.


### GetFont

`func (o *AgentSettingsWidget) GetFont() string`

GetFont returns the Font field if non-nil, zero value otherwise.

### GetFontOk

`func (o *AgentSettingsWidget) GetFontOk() (*string, bool)`

GetFontOk returns a tuple with the Font field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFont

`func (o *AgentSettingsWidget) SetFont(v string)`

SetFont sets Font field to given value.


### GetIconUrl

`func (o *AgentSettingsWidget) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *AgentSettingsWidget) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *AgentSettingsWidget) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *AgentSettingsWidget) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetChatPosition

`func (o *AgentSettingsWidget) GetChatPosition() string`

GetChatPosition returns the ChatPosition field if non-nil, zero value otherwise.

### GetChatPositionOk

`func (o *AgentSettingsWidget) GetChatPositionOk() (*string, bool)`

GetChatPositionOk returns a tuple with the ChatPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatPosition

`func (o *AgentSettingsWidget) SetChatPosition(v string)`

SetChatPosition sets ChatPosition field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


