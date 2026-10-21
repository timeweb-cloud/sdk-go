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

// checks if the ToolTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ToolTemplate{}

// ToolTemplate Шаблон MCP-инструмента из маркетплейса
type ToolTemplate struct {
	// ID шаблона. Передаётся в `template_id` при установке
	Id int32 `json:"id"`
	// Название
	Name string `json:"name"`
	// Уникальный ключ шаблона
	Key NullableString `json:"key,omitempty"`
	// Описание
	Description string `json:"description"`
	// Тип инструмента
	Type string `json:"type"`
	// Версия шаблона
	Version string `json:"version"`
	// Шаблон одобрен для маркетплейса: `1` — да, `0` — нет
	IsApproved int32 `json:"is_approved"`
	// Автор шаблона. `null` — шаблон Timeweb Cloud
	AuthorAccountId NullableString `json:"author_account_id,omitempty"`
	// Системный шаблон платформы: `1` — да, `0` — нет
	IsSystem int32 `json:"is_system"`
	// Сколько раз шаблон установлен
	InstallCount int32 `json:"install_count"`
	// Категория шаблона
	Category NullableString `json:"category,omitempty"`
	ConfigSchema NullableToolTemplateConfigSchema `json:"config_schema,omitempty"`
	// Способ авторизации на MCP-сервере
	AuthType string `json:"auth_type"`
	// Где работает MCP-сервер: `managed` — на стороне Timeweb Cloud, `external` — у стороннего сервиса
	DeploymentType string `json:"deployment_type"`
	// Дата и время создания
	CreatedAt time.Time `json:"created_at"`
	// Дата и время последнего изменения
	UpdatedAt time.Time `json:"updated_at"`
}

// NewToolTemplate instantiates a new ToolTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToolTemplate(id int32, name string, description string, type_ string, version string, isApproved int32, isSystem int32, installCount int32, authType string, deploymentType string, createdAt time.Time, updatedAt time.Time) *ToolTemplate {
	this := ToolTemplate{}
	this.Id = id
	this.Name = name
	this.Description = description
	this.Type = type_
	this.Version = version
	this.IsApproved = isApproved
	this.IsSystem = isSystem
	this.InstallCount = installCount
	this.AuthType = authType
	this.DeploymentType = deploymentType
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewToolTemplateWithDefaults instantiates a new ToolTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewToolTemplateWithDefaults() *ToolTemplate {
	this := ToolTemplate{}
	return &this
}

// GetId returns the Id field value
func (o *ToolTemplate) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ToolTemplate) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ToolTemplate) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ToolTemplate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ToolTemplate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ToolTemplate) SetName(v string) {
	o.Name = v
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ToolTemplate) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ToolTemplate) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *ToolTemplate) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *ToolTemplate) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *ToolTemplate) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *ToolTemplate) UnsetKey() {
	o.Key.Unset()
}

// GetDescription returns the Description field value
func (o *ToolTemplate) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ToolTemplate) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ToolTemplate) SetDescription(v string) {
	o.Description = v
}

// GetType returns the Type field value
func (o *ToolTemplate) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ToolTemplate) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ToolTemplate) SetType(v string) {
	o.Type = v
}

// GetVersion returns the Version field value
func (o *ToolTemplate) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ToolTemplate) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ToolTemplate) SetVersion(v string) {
	o.Version = v
}

// GetIsApproved returns the IsApproved field value
func (o *ToolTemplate) GetIsApproved() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IsApproved
}

// GetIsApprovedOk returns a tuple with the IsApproved field value
// and a boolean to check if the value has been set.
func (o *ToolTemplate) GetIsApprovedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsApproved, true
}

// SetIsApproved sets field value
func (o *ToolTemplate) SetIsApproved(v int32) {
	o.IsApproved = v
}

// GetAuthorAccountId returns the AuthorAccountId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ToolTemplate) GetAuthorAccountId() string {
	if o == nil || IsNil(o.AuthorAccountId.Get()) {
		var ret string
		return ret
	}
	return *o.AuthorAccountId.Get()
}

// GetAuthorAccountIdOk returns a tuple with the AuthorAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ToolTemplate) GetAuthorAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthorAccountId.Get(), o.AuthorAccountId.IsSet()
}

// HasAuthorAccountId returns a boolean if a field has been set.
func (o *ToolTemplate) HasAuthorAccountId() bool {
	if o != nil && o.AuthorAccountId.IsSet() {
		return true
	}

	return false
}

// SetAuthorAccountId gets a reference to the given NullableString and assigns it to the AuthorAccountId field.
func (o *ToolTemplate) SetAuthorAccountId(v string) {
	o.AuthorAccountId.Set(&v)
}
// SetAuthorAccountIdNil sets the value for AuthorAccountId to be an explicit nil
func (o *ToolTemplate) SetAuthorAccountIdNil() {
	o.AuthorAccountId.Set(nil)
}

// UnsetAuthorAccountId ensures that no value is present for AuthorAccountId, not even an explicit nil
func (o *ToolTemplate) UnsetAuthorAccountId() {
	o.AuthorAccountId.Unset()
}

// GetIsSystem returns the IsSystem field value
func (o *ToolTemplate) GetIsSystem() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IsSystem
}

// GetIsSystemOk returns a tuple with the IsSystem field value
// and a boolean to check if the value has been set.
func (o *ToolTemplate) GetIsSystemOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSystem, true
}

// SetIsSystem sets field value
func (o *ToolTemplate) SetIsSystem(v int32) {
	o.IsSystem = v
}

// GetInstallCount returns the InstallCount field value
func (o *ToolTemplate) GetInstallCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.InstallCount
}

// GetInstallCountOk returns a tuple with the InstallCount field value
// and a boolean to check if the value has been set.
func (o *ToolTemplate) GetInstallCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstallCount, true
}

// SetInstallCount sets field value
func (o *ToolTemplate) SetInstallCount(v int32) {
	o.InstallCount = v
}

// GetCategory returns the Category field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ToolTemplate) GetCategory() string {
	if o == nil || IsNil(o.Category.Get()) {
		var ret string
		return ret
	}
	return *o.Category.Get()
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ToolTemplate) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Category.Get(), o.Category.IsSet()
}

// HasCategory returns a boolean if a field has been set.
func (o *ToolTemplate) HasCategory() bool {
	if o != nil && o.Category.IsSet() {
		return true
	}

	return false
}

// SetCategory gets a reference to the given NullableString and assigns it to the Category field.
func (o *ToolTemplate) SetCategory(v string) {
	o.Category.Set(&v)
}
// SetCategoryNil sets the value for Category to be an explicit nil
func (o *ToolTemplate) SetCategoryNil() {
	o.Category.Set(nil)
}

// UnsetCategory ensures that no value is present for Category, not even an explicit nil
func (o *ToolTemplate) UnsetCategory() {
	o.Category.Unset()
}

// GetConfigSchema returns the ConfigSchema field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ToolTemplate) GetConfigSchema() ToolTemplateConfigSchema {
	if o == nil || IsNil(o.ConfigSchema.Get()) {
		var ret ToolTemplateConfigSchema
		return ret
	}
	return *o.ConfigSchema.Get()
}

// GetConfigSchemaOk returns a tuple with the ConfigSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ToolTemplate) GetConfigSchemaOk() (*ToolTemplateConfigSchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigSchema.Get(), o.ConfigSchema.IsSet()
}

// HasConfigSchema returns a boolean if a field has been set.
func (o *ToolTemplate) HasConfigSchema() bool {
	if o != nil && o.ConfigSchema.IsSet() {
		return true
	}

	return false
}

// SetConfigSchema gets a reference to the given NullableToolTemplateConfigSchema and assigns it to the ConfigSchema field.
func (o *ToolTemplate) SetConfigSchema(v ToolTemplateConfigSchema) {
	o.ConfigSchema.Set(&v)
}
// SetConfigSchemaNil sets the value for ConfigSchema to be an explicit nil
func (o *ToolTemplate) SetConfigSchemaNil() {
	o.ConfigSchema.Set(nil)
}

// UnsetConfigSchema ensures that no value is present for ConfigSchema, not even an explicit nil
func (o *ToolTemplate) UnsetConfigSchema() {
	o.ConfigSchema.Unset()
}

// GetAuthType returns the AuthType field value
func (o *ToolTemplate) GetAuthType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value
// and a boolean to check if the value has been set.
func (o *ToolTemplate) GetAuthTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthType, true
}

// SetAuthType sets field value
func (o *ToolTemplate) SetAuthType(v string) {
	o.AuthType = v
}

// GetDeploymentType returns the DeploymentType field value
func (o *ToolTemplate) GetDeploymentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeploymentType
}

// GetDeploymentTypeOk returns a tuple with the DeploymentType field value
// and a boolean to check if the value has been set.
func (o *ToolTemplate) GetDeploymentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeploymentType, true
}

// SetDeploymentType sets field value
func (o *ToolTemplate) SetDeploymentType(v string) {
	o.DeploymentType = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ToolTemplate) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ToolTemplate) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ToolTemplate) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ToolTemplate) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ToolTemplate) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ToolTemplate) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o ToolTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ToolTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	toSerialize["description"] = o.Description
	toSerialize["type"] = o.Type
	toSerialize["version"] = o.Version
	toSerialize["is_approved"] = o.IsApproved
	if o.AuthorAccountId.IsSet() {
		toSerialize["author_account_id"] = o.AuthorAccountId.Get()
	}
	toSerialize["is_system"] = o.IsSystem
	toSerialize["install_count"] = o.InstallCount
	if o.Category.IsSet() {
		toSerialize["category"] = o.Category.Get()
	}
	if o.ConfigSchema.IsSet() {
		toSerialize["config_schema"] = o.ConfigSchema.Get()
	}
	toSerialize["auth_type"] = o.AuthType
	toSerialize["deployment_type"] = o.DeploymentType
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

type NullableToolTemplate struct {
	value *ToolTemplate
	isSet bool
}

func (v NullableToolTemplate) Get() *ToolTemplate {
	return v.value
}

func (v *NullableToolTemplate) Set(val *ToolTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableToolTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableToolTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToolTemplate(val *ToolTemplate) *NullableToolTemplate {
	return &NullableToolTemplate{value: val, isSet: true}
}

func (v NullableToolTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToolTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


