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
	"time"
)

// checks if the Tool type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Tool{}

// Tool MCP-инструмент аккаунта
type Tool struct {
	// ID инструмента
	Id int32 `json:"id"`
	// Логин аккаунта-владельца
	AccountId string `json:"account_id"`
	// ID шаблона маркетплейса, из которого установлен инструмент. `null` — собственный MCP-сервер
	TemplateId NullableInt32 `json:"template_id"`
	// Ключ шаблона маркетплейса. Сейчас в ответах не возвращается
	Key NullableString `json:"key,omitempty"`
	// Название инструмента
	Name NullableString `json:"name"`
	// Описание инструмента
	Description NullableString `json:"description"`
	// Тип инструмента. У инструментов из шаблонов может быть `null`
	Type NullableString `json:"type"`
	// Транспорт MCP: `http` — Streamable HTTP, `sse` — Server-Sent Events. У инструментов из шаблонов всегда `null`
	Protocol NullableString `json:"protocol"`
	// Адрес MCP-сервера. У инструментов из шаблонов всегда `null`
	Url NullableString `json:"url"`
	// Методы MCP-сервера. `null` или пустой массив — методы ещё не получены, например до завершения OAuth2-авторизации
	Methods []ToolMethod `json:"methods"`
	AuthData ToolAuthData `json:"auth_data"`
	// Статус инструмента: `active` — работает, `inactive` — выключен вручную, `pending_auth` — ожидает OAuth2-авторизации, `error` — последнее обновление методов завершилось ошибкой (причина в `status_info`). Агент вызывает только инструменты в статусе `active`
	Status string `json:"status"`
	// Причина статуса `error`
	StatusInfo NullableString `json:"status_info"`
	// Дата и время создания
	CreatedAt time.Time `json:"created_at"`
	// Дата и время последнего изменения
	UpdatedAt time.Time `json:"updated_at"`
}

// NewTool instantiates a new Tool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTool(id int32, accountId string, templateId NullableInt32, name NullableString, description NullableString, type_ NullableString, protocol NullableString, url NullableString, methods []ToolMethod, authData ToolAuthData, status string, statusInfo NullableString, createdAt time.Time, updatedAt time.Time) *Tool {
	this := Tool{}
	this.Id = id
	this.AccountId = accountId
	this.TemplateId = templateId
	this.Name = name
	this.Description = description
	this.Type = type_
	this.Protocol = protocol
	this.Url = url
	this.Methods = methods
	this.AuthData = authData
	this.Status = status
	this.StatusInfo = statusInfo
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewToolWithDefaults instantiates a new Tool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewToolWithDefaults() *Tool {
	this := Tool{}
	return &this
}

// GetId returns the Id field value
func (o *Tool) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Tool) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Tool) SetId(v int32) {
	o.Id = v
}

// GetAccountId returns the AccountId field value
func (o *Tool) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *Tool) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *Tool) SetAccountId(v string) {
	o.AccountId = v
}

// GetTemplateId returns the TemplateId field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *Tool) GetTemplateId() int32 {
	if o == nil || o.TemplateId.Get() == nil {
		var ret int32
		return ret
	}

	return *o.TemplateId.Get()
}

// GetTemplateIdOk returns a tuple with the TemplateId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tool) GetTemplateIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TemplateId.Get(), o.TemplateId.IsSet()
}

// SetTemplateId sets field value
func (o *Tool) SetTemplateId(v int32) {
	o.TemplateId.Set(&v)
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tool) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tool) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *Tool) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *Tool) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *Tool) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *Tool) UnsetKey() {
	o.Key.Unset()
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Tool) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tool) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *Tool) SetName(v string) {
	o.Name.Set(&v)
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Tool) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tool) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *Tool) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Tool) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}

	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tool) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// SetType sets field value
func (o *Tool) SetType(v string) {
	o.Type.Set(&v)
}

// GetProtocol returns the Protocol field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Tool) GetProtocol() string {
	if o == nil || o.Protocol.Get() == nil {
		var ret string
		return ret
	}

	return *o.Protocol.Get()
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tool) GetProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Protocol.Get(), o.Protocol.IsSet()
}

// SetProtocol sets field value
func (o *Tool) SetProtocol(v string) {
	o.Protocol.Set(&v)
}

// GetUrl returns the Url field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Tool) GetUrl() string {
	if o == nil || o.Url.Get() == nil {
		var ret string
		return ret
	}

	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tool) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// SetUrl sets field value
func (o *Tool) SetUrl(v string) {
	o.Url.Set(&v)
}

// GetMethods returns the Methods field value
// If the value is explicit nil, the zero value for []ToolMethod will be returned
func (o *Tool) GetMethods() []ToolMethod {
	if o == nil {
		var ret []ToolMethod
		return ret
	}

	return o.Methods
}

// GetMethodsOk returns a tuple with the Methods field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tool) GetMethodsOk() ([]ToolMethod, bool) {
	if o == nil || IsNil(o.Methods) {
		return nil, false
	}
	return o.Methods, true
}

// SetMethods sets field value
func (o *Tool) SetMethods(v []ToolMethod) {
	o.Methods = v
}

// GetAuthData returns the AuthData field value
func (o *Tool) GetAuthData() ToolAuthData {
	if o == nil {
		var ret ToolAuthData
		return ret
	}

	return o.AuthData
}

// GetAuthDataOk returns a tuple with the AuthData field value
// and a boolean to check if the value has been set.
func (o *Tool) GetAuthDataOk() (*ToolAuthData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthData, true
}

// SetAuthData sets field value
func (o *Tool) SetAuthData(v ToolAuthData) {
	o.AuthData = v
}

// GetStatus returns the Status field value
func (o *Tool) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Tool) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Tool) SetStatus(v string) {
	o.Status = v
}

// GetStatusInfo returns the StatusInfo field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Tool) GetStatusInfo() string {
	if o == nil || o.StatusInfo.Get() == nil {
		var ret string
		return ret
	}

	return *o.StatusInfo.Get()
}

// GetStatusInfoOk returns a tuple with the StatusInfo field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tool) GetStatusInfoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StatusInfo.Get(), o.StatusInfo.IsSet()
}

// SetStatusInfo sets field value
func (o *Tool) SetStatusInfo(v string) {
	o.StatusInfo.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *Tool) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Tool) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Tool) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Tool) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Tool) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Tool) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o Tool) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Tool) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["account_id"] = o.AccountId
	toSerialize["template_id"] = o.TemplateId.Get()
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	toSerialize["name"] = o.Name.Get()
	toSerialize["description"] = o.Description.Get()
	toSerialize["type"] = o.Type.Get()
	toSerialize["protocol"] = o.Protocol.Get()
	toSerialize["url"] = o.Url.Get()
	if o.Methods != nil {
		toSerialize["methods"] = o.Methods
	}
	toSerialize["auth_data"] = o.AuthData
	toSerialize["status"] = o.Status
	toSerialize["status_info"] = o.StatusInfo.Get()
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

type NullableTool struct {
	value *Tool
	isSet bool
}

func (v NullableTool) Get() *Tool {
	return v.value
}

func (v *NullableTool) Set(val *Tool) {
	v.value = val
	v.isSet = true
}

func (v NullableTool) IsSet() bool {
	return v.isSet
}

func (v *NullableTool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTool(val *Tool) *NullableTool {
	return &NullableTool{value: val, isSet: true}
}

func (v NullableTool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


