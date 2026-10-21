# AgentWidgetSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WhitelistDomains** | **[]string** | Домены, на которых разрешено показывать виджет. Пустой массив — без ограничений. | 
**Name** | **string** | Отображаемое имя агента в виджете | 
**Signature** | Pointer to **string** | Подпись под именем агента в виджете | [optional] 
**WelcomeMessage** | **string** | Приветственное сообщение при открытии виджета | 
**PrimaryColor** | **string** | Основной цвет виджета в формате &#x60;#RRGGBB&#x60; | 
**TextColor** | Pointer to **string** | Цвет текста в формате &#x60;#RRGGBB&#x60; | [optional] 
**BackgroundColor** | Pointer to **string** | Цвет фона в формате &#x60;#RRGGBB&#x60; | [optional] 
**HeaderFooterColor** | Pointer to **string** | Цвет шапки и подвала виджета в формате &#x60;#RRGGBB&#x60; | [optional] 
**Font** | **string** | Шрифт виджета. Пустая строка — шрифт сайта. | 
**IconUrl** | Pointer to **string** | URL иконки агента в виджете | [optional] 
**UseAvatar** | Pointer to **bool** | Использовать загруженный в панели аватар как иконку агента | [optional] [default to false]
**ShowButton** | Pointer to **bool** | Показывать кнопку открытия виджета | [optional] [default to true]
**ShowToolProcessing** | Pointer to **bool** | Показывать в виджете статус выполнения инструментов агента | [optional] [default to false]
**ShowConfirmationButtons** | Pointer to **bool** | Показывать кнопки подтверждения и отмены операций, которые агент выполняет через инструменты | [optional] [default to true]
**ChatPosition** | **string** | Положение виджета на странице | 
**ChatPositionAdjustmentX** | Pointer to **float32** | Смещение виджета по горизонтали в пикселях | [optional] 
**ChatPositionAdjustmentY** | Pointer to **float32** | Смещение виджета по вертикали в пикселях | [optional] 

## Methods

### NewAgentWidgetSettings

`func NewAgentWidgetSettings(whitelistDomains []string, name string, welcomeMessage string, primaryColor string, font string, chatPosition string, ) *AgentWidgetSettings`

NewAgentWidgetSettings instantiates a new AgentWidgetSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentWidgetSettingsWithDefaults

`func NewAgentWidgetSettingsWithDefaults() *AgentWidgetSettings`

NewAgentWidgetSettingsWithDefaults instantiates a new AgentWidgetSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWhitelistDomains

`func (o *AgentWidgetSettings) GetWhitelistDomains() []string`

GetWhitelistDomains returns the WhitelistDomains field if non-nil, zero value otherwise.

### GetWhitelistDomainsOk

`func (o *AgentWidgetSettings) GetWhitelistDomainsOk() (*[]string, bool)`

GetWhitelistDomainsOk returns a tuple with the WhitelistDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelistDomains

`func (o *AgentWidgetSettings) SetWhitelistDomains(v []string)`

SetWhitelistDomains sets WhitelistDomains field to given value.


### GetName

`func (o *AgentWidgetSettings) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentWidgetSettings) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentWidgetSettings) SetName(v string)`

SetName sets Name field to given value.


### GetSignature

`func (o *AgentWidgetSettings) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *AgentWidgetSettings) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *AgentWidgetSettings) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *AgentWidgetSettings) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetWelcomeMessage

`func (o *AgentWidgetSettings) GetWelcomeMessage() string`

GetWelcomeMessage returns the WelcomeMessage field if non-nil, zero value otherwise.

### GetWelcomeMessageOk

`func (o *AgentWidgetSettings) GetWelcomeMessageOk() (*string, bool)`

GetWelcomeMessageOk returns a tuple with the WelcomeMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWelcomeMessage

`func (o *AgentWidgetSettings) SetWelcomeMessage(v string)`

SetWelcomeMessage sets WelcomeMessage field to given value.


### GetPrimaryColor

`func (o *AgentWidgetSettings) GetPrimaryColor() string`

GetPrimaryColor returns the PrimaryColor field if non-nil, zero value otherwise.

### GetPrimaryColorOk

`func (o *AgentWidgetSettings) GetPrimaryColorOk() (*string, bool)`

GetPrimaryColorOk returns a tuple with the PrimaryColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryColor

`func (o *AgentWidgetSettings) SetPrimaryColor(v string)`

SetPrimaryColor sets PrimaryColor field to given value.


### GetTextColor

`func (o *AgentWidgetSettings) GetTextColor() string`

GetTextColor returns the TextColor field if non-nil, zero value otherwise.

### GetTextColorOk

`func (o *AgentWidgetSettings) GetTextColorOk() (*string, bool)`

GetTextColorOk returns a tuple with the TextColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextColor

`func (o *AgentWidgetSettings) SetTextColor(v string)`

SetTextColor sets TextColor field to given value.

### HasTextColor

`func (o *AgentWidgetSettings) HasTextColor() bool`

HasTextColor returns a boolean if a field has been set.

### GetBackgroundColor

`func (o *AgentWidgetSettings) GetBackgroundColor() string`

GetBackgroundColor returns the BackgroundColor field if non-nil, zero value otherwise.

### GetBackgroundColorOk

`func (o *AgentWidgetSettings) GetBackgroundColorOk() (*string, bool)`

GetBackgroundColorOk returns a tuple with the BackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundColor

`func (o *AgentWidgetSettings) SetBackgroundColor(v string)`

SetBackgroundColor sets BackgroundColor field to given value.

### HasBackgroundColor

`func (o *AgentWidgetSettings) HasBackgroundColor() bool`

HasBackgroundColor returns a boolean if a field has been set.

### GetHeaderFooterColor

`func (o *AgentWidgetSettings) GetHeaderFooterColor() string`

GetHeaderFooterColor returns the HeaderFooterColor field if non-nil, zero value otherwise.

### GetHeaderFooterColorOk

`func (o *AgentWidgetSettings) GetHeaderFooterColorOk() (*string, bool)`

GetHeaderFooterColorOk returns a tuple with the HeaderFooterColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderFooterColor

`func (o *AgentWidgetSettings) SetHeaderFooterColor(v string)`

SetHeaderFooterColor sets HeaderFooterColor field to given value.

### HasHeaderFooterColor

`func (o *AgentWidgetSettings) HasHeaderFooterColor() bool`

HasHeaderFooterColor returns a boolean if a field has been set.

### GetFont

`func (o *AgentWidgetSettings) GetFont() string`

GetFont returns the Font field if non-nil, zero value otherwise.

### GetFontOk

`func (o *AgentWidgetSettings) GetFontOk() (*string, bool)`

GetFontOk returns a tuple with the Font field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFont

`func (o *AgentWidgetSettings) SetFont(v string)`

SetFont sets Font field to given value.


### GetIconUrl

`func (o *AgentWidgetSettings) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *AgentWidgetSettings) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *AgentWidgetSettings) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *AgentWidgetSettings) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetUseAvatar

`func (o *AgentWidgetSettings) GetUseAvatar() bool`

GetUseAvatar returns the UseAvatar field if non-nil, zero value otherwise.

### GetUseAvatarOk

`func (o *AgentWidgetSettings) GetUseAvatarOk() (*bool, bool)`

GetUseAvatarOk returns a tuple with the UseAvatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAvatar

`func (o *AgentWidgetSettings) SetUseAvatar(v bool)`

SetUseAvatar sets UseAvatar field to given value.

### HasUseAvatar

`func (o *AgentWidgetSettings) HasUseAvatar() bool`

HasUseAvatar returns a boolean if a field has been set.

### GetShowButton

`func (o *AgentWidgetSettings) GetShowButton() bool`

GetShowButton returns the ShowButton field if non-nil, zero value otherwise.

### GetShowButtonOk

`func (o *AgentWidgetSettings) GetShowButtonOk() (*bool, bool)`

GetShowButtonOk returns a tuple with the ShowButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowButton

`func (o *AgentWidgetSettings) SetShowButton(v bool)`

SetShowButton sets ShowButton field to given value.

### HasShowButton

`func (o *AgentWidgetSettings) HasShowButton() bool`

HasShowButton returns a boolean if a field has been set.

### GetShowToolProcessing

`func (o *AgentWidgetSettings) GetShowToolProcessing() bool`

GetShowToolProcessing returns the ShowToolProcessing field if non-nil, zero value otherwise.

### GetShowToolProcessingOk

`func (o *AgentWidgetSettings) GetShowToolProcessingOk() (*bool, bool)`

GetShowToolProcessingOk returns a tuple with the ShowToolProcessing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowToolProcessing

`func (o *AgentWidgetSettings) SetShowToolProcessing(v bool)`

SetShowToolProcessing sets ShowToolProcessing field to given value.

### HasShowToolProcessing

`func (o *AgentWidgetSettings) HasShowToolProcessing() bool`

HasShowToolProcessing returns a boolean if a field has been set.

### GetShowConfirmationButtons

`func (o *AgentWidgetSettings) GetShowConfirmationButtons() bool`

GetShowConfirmationButtons returns the ShowConfirmationButtons field if non-nil, zero value otherwise.

### GetShowConfirmationButtonsOk

`func (o *AgentWidgetSettings) GetShowConfirmationButtonsOk() (*bool, bool)`

GetShowConfirmationButtonsOk returns a tuple with the ShowConfirmationButtons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowConfirmationButtons

`func (o *AgentWidgetSettings) SetShowConfirmationButtons(v bool)`

SetShowConfirmationButtons sets ShowConfirmationButtons field to given value.

### HasShowConfirmationButtons

`func (o *AgentWidgetSettings) HasShowConfirmationButtons() bool`

HasShowConfirmationButtons returns a boolean if a field has been set.

### GetChatPosition

`func (o *AgentWidgetSettings) GetChatPosition() string`

GetChatPosition returns the ChatPosition field if non-nil, zero value otherwise.

### GetChatPositionOk

`func (o *AgentWidgetSettings) GetChatPositionOk() (*string, bool)`

GetChatPositionOk returns a tuple with the ChatPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatPosition

`func (o *AgentWidgetSettings) SetChatPosition(v string)`

SetChatPosition sets ChatPosition field to given value.


### GetChatPositionAdjustmentX

`func (o *AgentWidgetSettings) GetChatPositionAdjustmentX() float32`

GetChatPositionAdjustmentX returns the ChatPositionAdjustmentX field if non-nil, zero value otherwise.

### GetChatPositionAdjustmentXOk

`func (o *AgentWidgetSettings) GetChatPositionAdjustmentXOk() (*float32, bool)`

GetChatPositionAdjustmentXOk returns a tuple with the ChatPositionAdjustmentX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatPositionAdjustmentX

`func (o *AgentWidgetSettings) SetChatPositionAdjustmentX(v float32)`

SetChatPositionAdjustmentX sets ChatPositionAdjustmentX field to given value.

### HasChatPositionAdjustmentX

`func (o *AgentWidgetSettings) HasChatPositionAdjustmentX() bool`

HasChatPositionAdjustmentX returns a boolean if a field has been set.

### GetChatPositionAdjustmentY

`func (o *AgentWidgetSettings) GetChatPositionAdjustmentY() float32`

GetChatPositionAdjustmentY returns the ChatPositionAdjustmentY field if non-nil, zero value otherwise.

### GetChatPositionAdjustmentYOk

`func (o *AgentWidgetSettings) GetChatPositionAdjustmentYOk() (*float32, bool)`

GetChatPositionAdjustmentYOk returns a tuple with the ChatPositionAdjustmentY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatPositionAdjustmentY

`func (o *AgentWidgetSettings) SetChatPositionAdjustmentY(v float32)`

SetChatPositionAdjustmentY sets ChatPositionAdjustmentY field to given value.

### HasChatPositionAdjustmentY

`func (o *AgentWidgetSettings) HasChatPositionAdjustmentY() bool`

HasChatPositionAdjustmentY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


