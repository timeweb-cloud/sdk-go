/*
Timeweb Cloud API

# Введение API Timeweb Cloud позволяет вам управлять ресурсами в облаке программным способом с использованием обычных HTTP-запросов.  Множество функций, которые доступны в панели управления Timeweb Cloud, также доступны через API, что позволяет вам автоматизировать ваши собственные сценарии.  В этой документации сперва будет описан общий дизайн и принципы работы API, а после этого конкретные конечные точки. Также будут приведены примеры запросов к ним.   ## Запросы Запросы должны выполняться по протоколу `HTTPS`, чтобы гарантировать шифрование транзакций. Поддерживаются следующие методы запроса: |Метод|Применение| |--- |--- | |GET|Извлекает данные о коллекциях и отдельных ресурсах.| |POST|Для коллекций создает новый ресурс этого типа. Также используется для выполнения действий с конкретным ресурсом.| |PUT|Обновляет существующий ресурс.| |PATCH|Некоторые ресурсы поддерживают частичное обновление, то есть обновление только части атрибутов ресурса, в этом случае вместо метода PUT будет использован PATCH.| |DELETE|Удаляет ресурс.|  Методы `POST`, `PUT` и `PATCH` могут включать объект в тело запроса с типом содержимого `application/json`.  ### Параметры в запросах Некоторые коллекции поддерживают пагинацию, поиск или сортировку в запросах. В параметрах запроса требуется передать: - `limit` — обозначает количество записей, которое необходимо вернуть  - `offset` — указывает на смещение, относительно начала списка  - `search` — позволяет указать набор символов для поиска  - `sort` — можно задать правило сортировки коллекции  ## Ответы Запросы вернут один из следующих кодов состояния ответа HTTP:  |Статус|Описание| |--- |--- | |200 OK|Действие с ресурсом было выполнено успешно.| |201 Created|Ресурс был успешно создан. При этом ресурс может быть как уже готовым к использованию, так и находиться в процессе запуска.| |204 No Content|Действие с ресурсом было выполнено успешно, и ответ не содержит дополнительной информации в теле.| |400 Bad Request|Был отправлен неверный запрос, например, в нем отсутствуют обязательные параметры и т. д. Тело ответа будет содержать дополнительную информацию об ошибке.| |401 Unauthorized|Ошибка аутентификации.| |403 Forbidden|Аутентификация прошла успешно, но недостаточно прав для выполнения действия.| |404 Not Found|Запрашиваемый ресурс не найден.| |409 Conflict|Запрос конфликтует с текущим состоянием.| |423 Locked|Ресурс из запроса заблокирован от применения к нему указанного метода.| |429 Too Many Requests|Был достигнут лимит по количеству запросов в единицу времени.| |500 Internal Server Error|При выполнении запроса произошла какая-то внутренняя ошибка. Чтобы решить эту проблему, лучше всего создать тикет в панели управления.|  ### Структура успешного ответа Все конечные точки будут возвращать данные в формате `JSON`. Ответы на `GET`-запросы будут иметь на верхнем уровне следующую структуру атрибутов:  |Название поля|Тип|Описание| |--- |--- |--- | |[entity_name]|object, object[], string[], number[], boolean|Динамическое поле, которое будет меняться в зависимости от запрашиваемого ресурса и будет содержать все атрибуты, необходимые для описания этого ресурса. Например, при запросе списка баз данных будет возвращаться поле `dbs`, а при запросе конкретного облачного сервера `server`. Для некоторых конечных точек в ответе может возвращаться сразу несколько ресурсов.| |meta|object|Опционально. Объект, который содержит вспомогательную информацию о ресурсе. Чаще всего будет встречаться при запросе коллекций и содержать поле `total`, которое будет указывать на количество элементов в коллекции.| |response_id|string|Опционально. В большинстве случаев в ответе будет содержаться ID ответа в формате UUIDv4, который однозначно указывает на ваш запрос внутри нашей системы. Если вам потребуется задать вопрос нашей поддержке, приложите к вопросу этот ID— так мы сможем найти ответ на него намного быстрее. Также вы можете использовать этот ID, чтобы убедиться, что это новый ответ на запрос и результат не был получен из кэша.|  Пример запроса на получение списка SSH-ключей: ```     HTTP/2.0 200 OK     {       \"ssh_keys\":[           {             \"body\":\"ssh-rsa AAAAB3NzaC1sdfghjkOAsBwWhs= example@device.local\",             \"created_at\":\"2021-09-15T19:52:27Z\",             \"expired_at\":null,             \"id\":5297,             \"is_default\":false,             \"name\":\"example@device.local\",             \"used_at\":null,             \"used_by\":[]           }       ],       \"meta\":{           \"total\":1       },       \"response_id\":\"94608d15-8672-4eed-8ab6-28bd6fa3cdf7\"     } ```  ### Структура ответа с ошибкой |Название поля|Тип|Описание| |--- |--- |--- | |status_code|number|Короткий числовой идентификатор ошибки.| |error_code|string|Короткий текстовый идентификатор ошибки, который уточняет числовой идентификатор и удобен для программной обработки. Самый простой пример — это код `not_found` для ошибки 404.| |message|string, string[]|Опционально. В большинстве случаев в ответе будет содержаться человекочитаемое подробное описание ошибки или ошибок, которые помогут понять, что нужно исправить.| |response_id|string|Опционально. В большинстве случае в ответе будет содержаться ID ответа в формате UUIDv4, который однозначно указывает на ваш запрос внутри нашей системы. Если вам потребуется задать вопрос нашей поддержке, приложите к вопросу этот ID — так мы сможем найти ответ на него намного быстрее.|  Пример: ```     HTTP/2.0 403 Forbidden     {       \"status_code\": 403,       \"error_code\":  \"forbidden\",       \"message\":     \"You do not have access for the attempted action\",       \"response_id\": \"94608d15-8672-4eed-8ab6-28bd6fa3cdf7\"     } ```  ## Статусы ресурсов Важно учесть, что при создании большинства ресурсов внутри платформы вам будет сразу возвращен ответ от сервера со статусом `200 OK` или `201 Created` и ID созданного ресурса в теле ответа, но при этом этот ресурс может быть ещё в *состоянии запуска*.  Для того чтобы понять, в каком состоянии сейчас находится ваш ресурс, мы добавили поле `status` в ответ на получение информации о ресурсе.  Список статусов будет отличаться в зависимости от типа ресурса. Увидеть поддерживаемый список статусов вы сможете в описании каждого конкретного ресурса.     ## Ограничение скорости запросов (Rate Limiting) Чтобы обеспечить стабильность для всех пользователей, Timeweb Cloud защищает API от всплесков входящего трафика, анализируя количество запросов c каждого аккаунта к каждой конечной точке.  Если ваше приложение отправляет более 20 запросов в секунду на одну конечную точку, то для этого запроса API может вернуть код состояния HTTP `429 Too Many Requests`.   ## Аутентификация Доступ к API осуществляется с помощью JWT-токена. Токенами можно управлять внутри панели управления Timeweb Cloud в разделе *API и Terraform*.  Токен необходимо передавать в заголовке каждого запроса в формате: ```   Authorization: Bearer $TIMEWEB_CLOUD_TOKEN ```  ## Формат примеров API Примеры в этой документации описаны с помощью `curl`, HTTP-клиента командной строки. На компьютерах `Linux` и `macOS` обычно по умолчанию установлен `curl`, и он доступен для загрузки на всех популярных платформах, включая `Windows`.  Каждый пример разделен на несколько строк символом `\\`, который совместим с `bash`. Типичный пример выглядит так: ```   curl -X PATCH      -H \"Content-Type: application/json\"      -H \"Authorization: Bearer $TIMEWEB_CLOUD_TOKEN\"      -d '{\"name\":\"Cute Corvus\",\"comment\":\"Development Server\"}'      \"https://api.timeweb.cloud/api/v1/dedicated/1051\" ``` - Параметр `-X` задает метод запроса. Для согласованности метод будет указан во всех примерах, даже если он явно не требуется для методов `GET`. - Строки `-H` задают требуемые HTTP-заголовки. - Примеры, для которых требуется объект JSON в теле запроса, передают требуемые данные через параметр `-d`.  Чтобы использовать приведенные примеры, не подставляя каждый раз в них свой токен, вы можете добавить токен один раз в переменные окружения в вашей консоли. Например, на `Linux` это можно сделать с помощью команды:  ``` TIMEWEB_CLOUD_TOKEN=\"token\" ```  После этого токен будет автоматически подставляться в ваши запросы.  Обратите внимание, что все значения в этой документации являются примерами. Не полагайтесь на IDы операционных систем, тарифов и т.д., используемые в примерах. Используйте соответствующую конечную точку для получения значений перед созданием ресурсов.   ## Версионирование API построено согласно принципам [семантического версионирования](https://semver.org/lang/ru). Это значит, что мы гарантируем обратную совместимость всех изменений в пределах одной мажорной версии.  Мажорная версия каждой конечной точки обозначается в пути запроса, например, запрос `/api/v1/servers` указывает, что этот метод имеет версию 1.

API version: 1.0.0
Contact: info@timeweb.cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the AgentWidgetSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentWidgetSettings{}

// AgentWidgetSettings Настройки виджета чата агента для встраивания на сайт
type AgentWidgetSettings struct {
	// Домены, на которых разрешено показывать виджет. Пустой массив — без ограничений.
	WhitelistDomains []string `json:"whitelist_domains"`
	// Отображаемое имя агента в виджете
	Name string `json:"name"`
	// Подпись под именем агента в виджете
	Signature *string `json:"signature,omitempty"`
	// Приветственное сообщение при открытии виджета
	WelcomeMessage string `json:"welcome_message"`
	// Основной цвет виджета в формате `#RRGGBB`
	PrimaryColor string `json:"primary_color"`
	// Цвет текста в формате `#RRGGBB`
	TextColor *string `json:"text_color,omitempty"`
	// Цвет фона в формате `#RRGGBB`
	BackgroundColor *string `json:"background_color,omitempty"`
	// Цвет шапки и подвала виджета в формате `#RRGGBB`
	HeaderFooterColor *string `json:"header_footer_color,omitempty"`
	// Шрифт виджета. Пустая строка — шрифт сайта.
	Font string `json:"font"`
	// URL иконки агента в виджете
	IconUrl *string `json:"icon_url,omitempty"`
	// Использовать загруженный в панели аватар как иконку агента
	UseAvatar *bool `json:"use_avatar,omitempty"`
	// Показывать кнопку открытия виджета
	ShowButton *bool `json:"show_button,omitempty"`
	// Показывать в виджете статус выполнения инструментов агента
	ShowToolProcessing *bool `json:"show_tool_processing,omitempty"`
	// Показывать кнопки подтверждения и отмены операций, которые агент выполняет через инструменты
	ShowConfirmationButtons *bool `json:"show_confirmation_buttons,omitempty"`
	// Положение виджета на странице
	ChatPosition string `json:"chat_position"`
	// Смещение виджета по горизонтали в пикселях
	ChatPositionAdjustmentX *float32 `json:"chat_position_adjustment_x,omitempty"`
	// Смещение виджета по вертикали в пикселях
	ChatPositionAdjustmentY *float32 `json:"chat_position_adjustment_y,omitempty"`
}

// NewAgentWidgetSettings instantiates a new AgentWidgetSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentWidgetSettings(whitelistDomains []string, name string, welcomeMessage string, primaryColor string, font string, chatPosition string) *AgentWidgetSettings {
	this := AgentWidgetSettings{}
	this.WhitelistDomains = whitelistDomains
	this.Name = name
	this.WelcomeMessage = welcomeMessage
	this.PrimaryColor = primaryColor
	this.Font = font
	var useAvatar bool = false
	this.UseAvatar = &useAvatar
	var showButton bool = true
	this.ShowButton = &showButton
	var showToolProcessing bool = false
	this.ShowToolProcessing = &showToolProcessing
	var showConfirmationButtons bool = true
	this.ShowConfirmationButtons = &showConfirmationButtons
	this.ChatPosition = chatPosition
	return &this
}

// NewAgentWidgetSettingsWithDefaults instantiates a new AgentWidgetSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentWidgetSettingsWithDefaults() *AgentWidgetSettings {
	this := AgentWidgetSettings{}
	var useAvatar bool = false
	this.UseAvatar = &useAvatar
	var showButton bool = true
	this.ShowButton = &showButton
	var showToolProcessing bool = false
	this.ShowToolProcessing = &showToolProcessing
	var showConfirmationButtons bool = true
	this.ShowConfirmationButtons = &showConfirmationButtons
	return &this
}

// GetWhitelistDomains returns the WhitelistDomains field value
func (o *AgentWidgetSettings) GetWhitelistDomains() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.WhitelistDomains
}

// GetWhitelistDomainsOk returns a tuple with the WhitelistDomains field value
// and a boolean to check if the value has been set.
func (o *AgentWidgetSettings) GetWhitelistDomainsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WhitelistDomains, true
}

// SetWhitelistDomains sets field value
func (o *AgentWidgetSettings) SetWhitelistDomains(v []string) {
	o.WhitelistDomains = v
}

// GetName returns the Name field value
func (o *AgentWidgetSettings) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AgentWidgetSettings) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AgentWidgetSettings) SetName(v string) {
	o.Name = v
}

// GetSignature returns the Signature field value if set, zero value otherwise.
func (o *AgentWidgetSettings) GetSignature() string {
	if o == nil || IsNil(o.Signature) {
		var ret string
		return ret
	}
	return *o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentWidgetSettings) GetSignatureOk() (*string, bool) {
	if o == nil || IsNil(o.Signature) {
		return nil, false
	}
	return o.Signature, true
}

// HasSignature returns a boolean if a field has been set.
func (o *AgentWidgetSettings) HasSignature() bool {
	if o != nil && !IsNil(o.Signature) {
		return true
	}

	return false
}

// SetSignature gets a reference to the given string and assigns it to the Signature field.
func (o *AgentWidgetSettings) SetSignature(v string) {
	o.Signature = &v
}

// GetWelcomeMessage returns the WelcomeMessage field value
func (o *AgentWidgetSettings) GetWelcomeMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WelcomeMessage
}

// GetWelcomeMessageOk returns a tuple with the WelcomeMessage field value
// and a boolean to check if the value has been set.
func (o *AgentWidgetSettings) GetWelcomeMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WelcomeMessage, true
}

// SetWelcomeMessage sets field value
func (o *AgentWidgetSettings) SetWelcomeMessage(v string) {
	o.WelcomeMessage = v
}

// GetPrimaryColor returns the PrimaryColor field value
func (o *AgentWidgetSettings) GetPrimaryColor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrimaryColor
}

// GetPrimaryColorOk returns a tuple with the PrimaryColor field value
// and a boolean to check if the value has been set.
func (o *AgentWidgetSettings) GetPrimaryColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrimaryColor, true
}

// SetPrimaryColor sets field value
func (o *AgentWidgetSettings) SetPrimaryColor(v string) {
	o.PrimaryColor = v
}

// GetTextColor returns the TextColor field value if set, zero value otherwise.
func (o *AgentWidgetSettings) GetTextColor() string {
	if o == nil || IsNil(o.TextColor) {
		var ret string
		return ret
	}
	return *o.TextColor
}

// GetTextColorOk returns a tuple with the TextColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentWidgetSettings) GetTextColorOk() (*string, bool) {
	if o == nil || IsNil(o.TextColor) {
		return nil, false
	}
	return o.TextColor, true
}

// HasTextColor returns a boolean if a field has been set.
func (o *AgentWidgetSettings) HasTextColor() bool {
	if o != nil && !IsNil(o.TextColor) {
		return true
	}

	return false
}

// SetTextColor gets a reference to the given string and assigns it to the TextColor field.
func (o *AgentWidgetSettings) SetTextColor(v string) {
	o.TextColor = &v
}

// GetBackgroundColor returns the BackgroundColor field value if set, zero value otherwise.
func (o *AgentWidgetSettings) GetBackgroundColor() string {
	if o == nil || IsNil(o.BackgroundColor) {
		var ret string
		return ret
	}
	return *o.BackgroundColor
}

// GetBackgroundColorOk returns a tuple with the BackgroundColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentWidgetSettings) GetBackgroundColorOk() (*string, bool) {
	if o == nil || IsNil(o.BackgroundColor) {
		return nil, false
	}
	return o.BackgroundColor, true
}

// HasBackgroundColor returns a boolean if a field has been set.
func (o *AgentWidgetSettings) HasBackgroundColor() bool {
	if o != nil && !IsNil(o.BackgroundColor) {
		return true
	}

	return false
}

// SetBackgroundColor gets a reference to the given string and assigns it to the BackgroundColor field.
func (o *AgentWidgetSettings) SetBackgroundColor(v string) {
	o.BackgroundColor = &v
}

// GetHeaderFooterColor returns the HeaderFooterColor field value if set, zero value otherwise.
func (o *AgentWidgetSettings) GetHeaderFooterColor() string {
	if o == nil || IsNil(o.HeaderFooterColor) {
		var ret string
		return ret
	}
	return *o.HeaderFooterColor
}

// GetHeaderFooterColorOk returns a tuple with the HeaderFooterColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentWidgetSettings) GetHeaderFooterColorOk() (*string, bool) {
	if o == nil || IsNil(o.HeaderFooterColor) {
		return nil, false
	}
	return o.HeaderFooterColor, true
}

// HasHeaderFooterColor returns a boolean if a field has been set.
func (o *AgentWidgetSettings) HasHeaderFooterColor() bool {
	if o != nil && !IsNil(o.HeaderFooterColor) {
		return true
	}

	return false
}

// SetHeaderFooterColor gets a reference to the given string and assigns it to the HeaderFooterColor field.
func (o *AgentWidgetSettings) SetHeaderFooterColor(v string) {
	o.HeaderFooterColor = &v
}

// GetFont returns the Font field value
func (o *AgentWidgetSettings) GetFont() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Font
}

// GetFontOk returns a tuple with the Font field value
// and a boolean to check if the value has been set.
func (o *AgentWidgetSettings) GetFontOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Font, true
}

// SetFont sets field value
func (o *AgentWidgetSettings) SetFont(v string) {
	o.Font = v
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise.
func (o *AgentWidgetSettings) GetIconUrl() string {
	if o == nil || IsNil(o.IconUrl) {
		var ret string
		return ret
	}
	return *o.IconUrl
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentWidgetSettings) GetIconUrlOk() (*string, bool) {
	if o == nil || IsNil(o.IconUrl) {
		return nil, false
	}
	return o.IconUrl, true
}

// HasIconUrl returns a boolean if a field has been set.
func (o *AgentWidgetSettings) HasIconUrl() bool {
	if o != nil && !IsNil(o.IconUrl) {
		return true
	}

	return false
}

// SetIconUrl gets a reference to the given string and assigns it to the IconUrl field.
func (o *AgentWidgetSettings) SetIconUrl(v string) {
	o.IconUrl = &v
}

// GetUseAvatar returns the UseAvatar field value if set, zero value otherwise.
func (o *AgentWidgetSettings) GetUseAvatar() bool {
	if o == nil || IsNil(o.UseAvatar) {
		var ret bool
		return ret
	}
	return *o.UseAvatar
}

// GetUseAvatarOk returns a tuple with the UseAvatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentWidgetSettings) GetUseAvatarOk() (*bool, bool) {
	if o == nil || IsNil(o.UseAvatar) {
		return nil, false
	}
	return o.UseAvatar, true
}

// HasUseAvatar returns a boolean if a field has been set.
func (o *AgentWidgetSettings) HasUseAvatar() bool {
	if o != nil && !IsNil(o.UseAvatar) {
		return true
	}

	return false
}

// SetUseAvatar gets a reference to the given bool and assigns it to the UseAvatar field.
func (o *AgentWidgetSettings) SetUseAvatar(v bool) {
	o.UseAvatar = &v
}

// GetShowButton returns the ShowButton field value if set, zero value otherwise.
func (o *AgentWidgetSettings) GetShowButton() bool {
	if o == nil || IsNil(o.ShowButton) {
		var ret bool
		return ret
	}
	return *o.ShowButton
}

// GetShowButtonOk returns a tuple with the ShowButton field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentWidgetSettings) GetShowButtonOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowButton) {
		return nil, false
	}
	return o.ShowButton, true
}

// HasShowButton returns a boolean if a field has been set.
func (o *AgentWidgetSettings) HasShowButton() bool {
	if o != nil && !IsNil(o.ShowButton) {
		return true
	}

	return false
}

// SetShowButton gets a reference to the given bool and assigns it to the ShowButton field.
func (o *AgentWidgetSettings) SetShowButton(v bool) {
	o.ShowButton = &v
}

// GetShowToolProcessing returns the ShowToolProcessing field value if set, zero value otherwise.
func (o *AgentWidgetSettings) GetShowToolProcessing() bool {
	if o == nil || IsNil(o.ShowToolProcessing) {
		var ret bool
		return ret
	}
	return *o.ShowToolProcessing
}

// GetShowToolProcessingOk returns a tuple with the ShowToolProcessing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentWidgetSettings) GetShowToolProcessingOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowToolProcessing) {
		return nil, false
	}
	return o.ShowToolProcessing, true
}

// HasShowToolProcessing returns a boolean if a field has been set.
func (o *AgentWidgetSettings) HasShowToolProcessing() bool {
	if o != nil && !IsNil(o.ShowToolProcessing) {
		return true
	}

	return false
}

// SetShowToolProcessing gets a reference to the given bool and assigns it to the ShowToolProcessing field.
func (o *AgentWidgetSettings) SetShowToolProcessing(v bool) {
	o.ShowToolProcessing = &v
}

// GetShowConfirmationButtons returns the ShowConfirmationButtons field value if set, zero value otherwise.
func (o *AgentWidgetSettings) GetShowConfirmationButtons() bool {
	if o == nil || IsNil(o.ShowConfirmationButtons) {
		var ret bool
		return ret
	}
	return *o.ShowConfirmationButtons
}

// GetShowConfirmationButtonsOk returns a tuple with the ShowConfirmationButtons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentWidgetSettings) GetShowConfirmationButtonsOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowConfirmationButtons) {
		return nil, false
	}
	return o.ShowConfirmationButtons, true
}

// HasShowConfirmationButtons returns a boolean if a field has been set.
func (o *AgentWidgetSettings) HasShowConfirmationButtons() bool {
	if o != nil && !IsNil(o.ShowConfirmationButtons) {
		return true
	}

	return false
}

// SetShowConfirmationButtons gets a reference to the given bool and assigns it to the ShowConfirmationButtons field.
func (o *AgentWidgetSettings) SetShowConfirmationButtons(v bool) {
	o.ShowConfirmationButtons = &v
}

// GetChatPosition returns the ChatPosition field value
func (o *AgentWidgetSettings) GetChatPosition() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChatPosition
}

// GetChatPositionOk returns a tuple with the ChatPosition field value
// and a boolean to check if the value has been set.
func (o *AgentWidgetSettings) GetChatPositionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChatPosition, true
}

// SetChatPosition sets field value
func (o *AgentWidgetSettings) SetChatPosition(v string) {
	o.ChatPosition = v
}

// GetChatPositionAdjustmentX returns the ChatPositionAdjustmentX field value if set, zero value otherwise.
func (o *AgentWidgetSettings) GetChatPositionAdjustmentX() float32 {
	if o == nil || IsNil(o.ChatPositionAdjustmentX) {
		var ret float32
		return ret
	}
	return *o.ChatPositionAdjustmentX
}

// GetChatPositionAdjustmentXOk returns a tuple with the ChatPositionAdjustmentX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentWidgetSettings) GetChatPositionAdjustmentXOk() (*float32, bool) {
	if o == nil || IsNil(o.ChatPositionAdjustmentX) {
		return nil, false
	}
	return o.ChatPositionAdjustmentX, true
}

// HasChatPositionAdjustmentX returns a boolean if a field has been set.
func (o *AgentWidgetSettings) HasChatPositionAdjustmentX() bool {
	if o != nil && !IsNil(o.ChatPositionAdjustmentX) {
		return true
	}

	return false
}

// SetChatPositionAdjustmentX gets a reference to the given float32 and assigns it to the ChatPositionAdjustmentX field.
func (o *AgentWidgetSettings) SetChatPositionAdjustmentX(v float32) {
	o.ChatPositionAdjustmentX = &v
}

// GetChatPositionAdjustmentY returns the ChatPositionAdjustmentY field value if set, zero value otherwise.
func (o *AgentWidgetSettings) GetChatPositionAdjustmentY() float32 {
	if o == nil || IsNil(o.ChatPositionAdjustmentY) {
		var ret float32
		return ret
	}
	return *o.ChatPositionAdjustmentY
}

// GetChatPositionAdjustmentYOk returns a tuple with the ChatPositionAdjustmentY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentWidgetSettings) GetChatPositionAdjustmentYOk() (*float32, bool) {
	if o == nil || IsNil(o.ChatPositionAdjustmentY) {
		return nil, false
	}
	return o.ChatPositionAdjustmentY, true
}

// HasChatPositionAdjustmentY returns a boolean if a field has been set.
func (o *AgentWidgetSettings) HasChatPositionAdjustmentY() bool {
	if o != nil && !IsNil(o.ChatPositionAdjustmentY) {
		return true
	}

	return false
}

// SetChatPositionAdjustmentY gets a reference to the given float32 and assigns it to the ChatPositionAdjustmentY field.
func (o *AgentWidgetSettings) SetChatPositionAdjustmentY(v float32) {
	o.ChatPositionAdjustmentY = &v
}

func (o AgentWidgetSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentWidgetSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["whitelist_domains"] = o.WhitelistDomains
	toSerialize["name"] = o.Name
	if !IsNil(o.Signature) {
		toSerialize["signature"] = o.Signature
	}
	toSerialize["welcome_message"] = o.WelcomeMessage
	toSerialize["primary_color"] = o.PrimaryColor
	if !IsNil(o.TextColor) {
		toSerialize["text_color"] = o.TextColor
	}
	if !IsNil(o.BackgroundColor) {
		toSerialize["background_color"] = o.BackgroundColor
	}
	if !IsNil(o.HeaderFooterColor) {
		toSerialize["header_footer_color"] = o.HeaderFooterColor
	}
	toSerialize["font"] = o.Font
	if !IsNil(o.IconUrl) {
		toSerialize["icon_url"] = o.IconUrl
	}
	if !IsNil(o.UseAvatar) {
		toSerialize["use_avatar"] = o.UseAvatar
	}
	if !IsNil(o.ShowButton) {
		toSerialize["show_button"] = o.ShowButton
	}
	if !IsNil(o.ShowToolProcessing) {
		toSerialize["show_tool_processing"] = o.ShowToolProcessing
	}
	if !IsNil(o.ShowConfirmationButtons) {
		toSerialize["show_confirmation_buttons"] = o.ShowConfirmationButtons
	}
	toSerialize["chat_position"] = o.ChatPosition
	if !IsNil(o.ChatPositionAdjustmentX) {
		toSerialize["chat_position_adjustment_x"] = o.ChatPositionAdjustmentX
	}
	if !IsNil(o.ChatPositionAdjustmentY) {
		toSerialize["chat_position_adjustment_y"] = o.ChatPositionAdjustmentY
	}
	return toSerialize, nil
}

type NullableAgentWidgetSettings struct {
	value *AgentWidgetSettings
	isSet bool
}

func (v NullableAgentWidgetSettings) Get() *AgentWidgetSettings {
	return v.value
}

func (v *NullableAgentWidgetSettings) Set(val *AgentWidgetSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentWidgetSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentWidgetSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentWidgetSettings(val *AgentWidgetSettings) *NullableAgentWidgetSettings {
	return &NullableAgentWidgetSettings{value: val, isSet: true}
}

func (v NullableAgentWidgetSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentWidgetSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


