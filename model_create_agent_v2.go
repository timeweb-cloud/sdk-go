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

// checks if the CreateAgentV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAgentV2{}

// CreateAgentV2 Данные для создания AI агента через API v2
type CreateAgentV2 struct {
	// Название агента
	Name string `json:"name"`
	// Описание агента
	Description *string `json:"description,omitempty"`
	// Тип доступа к агенту
	AccessType string `json:"access_type"`
	// ID основной модели
	ModelId float32 `json:"model_id"`
	// Дневной лимит токенов для агента (0 — без лимита)
	TokenLimit *float32 `json:"token_limit,omitempty"`
	Settings AgentSettings `json:"settings"`
	// ID проекта
	ProjectId *float32 `json:"project_id,omitempty"`
	// Список ID дополнительных моделей агента
	AdditionalModelIds []float32 `json:"additional_model_ids,omitempty"`
	// Признак использования веб-поиска агентом
	IsWebSearchEnabled *bool `json:"is_web_search_enabled,omitempty"`
}

// NewCreateAgentV2 instantiates a new CreateAgentV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAgentV2(name string, accessType string, modelId float32, settings AgentSettings) *CreateAgentV2 {
	this := CreateAgentV2{}
	this.Name = name
	this.AccessType = accessType
	this.ModelId = modelId
	this.Settings = settings
	return &this
}

// NewCreateAgentV2WithDefaults instantiates a new CreateAgentV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAgentV2WithDefaults() *CreateAgentV2 {
	this := CreateAgentV2{}
	return &this
}

// GetName returns the Name field value
func (o *CreateAgentV2) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateAgentV2) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateAgentV2) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateAgentV2) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAgentV2) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateAgentV2) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateAgentV2) SetDescription(v string) {
	o.Description = &v
}

// GetAccessType returns the AccessType field value
func (o *CreateAgentV2) GetAccessType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value
// and a boolean to check if the value has been set.
func (o *CreateAgentV2) GetAccessTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessType, true
}

// SetAccessType sets field value
func (o *CreateAgentV2) SetAccessType(v string) {
	o.AccessType = v
}

// GetModelId returns the ModelId field value
func (o *CreateAgentV2) GetModelId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ModelId
}

// GetModelIdOk returns a tuple with the ModelId field value
// and a boolean to check if the value has been set.
func (o *CreateAgentV2) GetModelIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelId, true
}

// SetModelId sets field value
func (o *CreateAgentV2) SetModelId(v float32) {
	o.ModelId = v
}

// GetTokenLimit returns the TokenLimit field value if set, zero value otherwise.
func (o *CreateAgentV2) GetTokenLimit() float32 {
	if o == nil || IsNil(o.TokenLimit) {
		var ret float32
		return ret
	}
	return *o.TokenLimit
}

// GetTokenLimitOk returns a tuple with the TokenLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAgentV2) GetTokenLimitOk() (*float32, bool) {
	if o == nil || IsNil(o.TokenLimit) {
		return nil, false
	}
	return o.TokenLimit, true
}

// HasTokenLimit returns a boolean if a field has been set.
func (o *CreateAgentV2) HasTokenLimit() bool {
	if o != nil && !IsNil(o.TokenLimit) {
		return true
	}

	return false
}

// SetTokenLimit gets a reference to the given float32 and assigns it to the TokenLimit field.
func (o *CreateAgentV2) SetTokenLimit(v float32) {
	o.TokenLimit = &v
}

// GetSettings returns the Settings field value
func (o *CreateAgentV2) GetSettings() AgentSettings {
	if o == nil {
		var ret AgentSettings
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *CreateAgentV2) GetSettingsOk() (*AgentSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value
func (o *CreateAgentV2) SetSettings(v AgentSettings) {
	o.Settings = v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *CreateAgentV2) GetProjectId() float32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret float32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAgentV2) GetProjectIdOk() (*float32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *CreateAgentV2) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given float32 and assigns it to the ProjectId field.
func (o *CreateAgentV2) SetProjectId(v float32) {
	o.ProjectId = &v
}

// GetAdditionalModelIds returns the AdditionalModelIds field value if set, zero value otherwise.
func (o *CreateAgentV2) GetAdditionalModelIds() []float32 {
	if o == nil || IsNil(o.AdditionalModelIds) {
		var ret []float32
		return ret
	}
	return o.AdditionalModelIds
}

// GetAdditionalModelIdsOk returns a tuple with the AdditionalModelIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAgentV2) GetAdditionalModelIdsOk() ([]float32, bool) {
	if o == nil || IsNil(o.AdditionalModelIds) {
		return nil, false
	}
	return o.AdditionalModelIds, true
}

// HasAdditionalModelIds returns a boolean if a field has been set.
func (o *CreateAgentV2) HasAdditionalModelIds() bool {
	if o != nil && !IsNil(o.AdditionalModelIds) {
		return true
	}

	return false
}

// SetAdditionalModelIds gets a reference to the given []float32 and assigns it to the AdditionalModelIds field.
func (o *CreateAgentV2) SetAdditionalModelIds(v []float32) {
	o.AdditionalModelIds = v
}

// GetIsWebSearchEnabled returns the IsWebSearchEnabled field value if set, zero value otherwise.
func (o *CreateAgentV2) GetIsWebSearchEnabled() bool {
	if o == nil || IsNil(o.IsWebSearchEnabled) {
		var ret bool
		return ret
	}
	return *o.IsWebSearchEnabled
}

// GetIsWebSearchEnabledOk returns a tuple with the IsWebSearchEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAgentV2) GetIsWebSearchEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsWebSearchEnabled) {
		return nil, false
	}
	return o.IsWebSearchEnabled, true
}

// HasIsWebSearchEnabled returns a boolean if a field has been set.
func (o *CreateAgentV2) HasIsWebSearchEnabled() bool {
	if o != nil && !IsNil(o.IsWebSearchEnabled) {
		return true
	}

	return false
}

// SetIsWebSearchEnabled gets a reference to the given bool and assigns it to the IsWebSearchEnabled field.
func (o *CreateAgentV2) SetIsWebSearchEnabled(v bool) {
	o.IsWebSearchEnabled = &v
}

func (o CreateAgentV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAgentV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["access_type"] = o.AccessType
	toSerialize["model_id"] = o.ModelId
	if !IsNil(o.TokenLimit) {
		toSerialize["token_limit"] = o.TokenLimit
	}
	toSerialize["settings"] = o.Settings
	if !IsNil(o.ProjectId) {
		toSerialize["project_id"] = o.ProjectId
	}
	if !IsNil(o.AdditionalModelIds) {
		toSerialize["additional_model_ids"] = o.AdditionalModelIds
	}
	if !IsNil(o.IsWebSearchEnabled) {
		toSerialize["is_web_search_enabled"] = o.IsWebSearchEnabled
	}
	return toSerialize, nil
}

type NullableCreateAgentV2 struct {
	value *CreateAgentV2
	isSet bool
}

func (v NullableCreateAgentV2) Get() *CreateAgentV2 {
	return v.value
}

func (v *NullableCreateAgentV2) Set(val *CreateAgentV2) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAgentV2) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAgentV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAgentV2(val *CreateAgentV2) *NullableCreateAgentV2 {
	return &NullableCreateAgentV2{value: val, isSet: true}
}

func (v NullableCreateAgentV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAgentV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


