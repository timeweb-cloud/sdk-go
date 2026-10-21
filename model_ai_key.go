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

// checks if the AiKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiKey{}

// AiKey Ключ AI Gateway
type AiKey struct {
	// ID ключа (UUID)
	Id string `json:"id"`
	// Логин аккаунта
	AccountId string `json:"account_id"`
	// Название ключа
	Name string `json:"name"`
	// Описание ключа
	Description NullableString `json:"description"`
	// Тип ключа: `aikey` — ключ, созданный пользователем (в том числе ключ AI-решения), `k8s` — ключ, выпущенный платформой для кластера Kubernetes
	EntityType string `json:"entity_type"`
	// ID ресурса, для которого платформа выпустила ключ (например, кластера Kubernetes). `null` — у ключей, созданных пользователем
	EntityId NullableString `json:"entity_id"`
	// ID проекта, к которому привязан ключ
	ProjectId NullableInt32 `json:"project_id"`
	// Дата и время окончания действия ключа. `null` — ключ бессрочный
	ExpiresAt NullableTime `json:"expires_at"`
	// Ключ управляется платформой (выпущен для AI-решения или кластера Kubernetes): изменить, перевыпустить и удалить его нельзя — такие запросы вернут 403
	IsReadonly bool `json:"is_readonly"`
	// Запасные модели в порядке приоритета. `[]` — не настроены, `null` — настройки временно не удалось прочитать (повторите запрос позже). Запасные модели применяются к любой модели, запрошенной этим ключом, и только в `/v1/chat/completions` и `/v1/responses`: если запрошенная модель не ответила, запрос по порядку повторяется на запасных. Запасная модель, совпадающая с запрошенной, пропускается. Резерв средств и проверка лимитов делаются по цене запрошенной модели, списание и учёт расхода — по цене ответившей: если запасная модель дороже, итоговое списание может превысить резерв и денежный лимит ключа.
	FallbackModels []string `json:"fallback_models"`
	// Дата и время создания ключа
	CreatedAt time.Time `json:"created_at"`
	// Дата и время изменения ключа
	UpdatedAt time.Time `json:"updated_at"`
}

// NewAiKey instantiates a new AiKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiKey(id string, accountId string, name string, description NullableString, entityType string, entityId NullableString, projectId NullableInt32, expiresAt NullableTime, isReadonly bool, fallbackModels []string, createdAt time.Time, updatedAt time.Time) *AiKey {
	this := AiKey{}
	this.Id = id
	this.AccountId = accountId
	this.Name = name
	this.Description = description
	this.EntityType = entityType
	this.EntityId = entityId
	this.ProjectId = projectId
	this.ExpiresAt = expiresAt
	this.IsReadonly = isReadonly
	this.FallbackModels = fallbackModels
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewAiKeyWithDefaults instantiates a new AiKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiKeyWithDefaults() *AiKey {
	this := AiKey{}
	return &this
}

// GetId returns the Id field value
func (o *AiKey) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AiKey) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AiKey) SetId(v string) {
	o.Id = v
}

// GetAccountId returns the AccountId field value
func (o *AiKey) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *AiKey) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *AiKey) SetAccountId(v string) {
	o.AccountId = v
}

// GetName returns the Name field value
func (o *AiKey) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiKey) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AiKey) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AiKey) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiKey) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *AiKey) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetEntityType returns the EntityType field value
func (o *AiKey) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *AiKey) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *AiKey) SetEntityType(v string) {
	o.EntityType = v
}

// GetEntityId returns the EntityId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AiKey) GetEntityId() string {
	if o == nil || o.EntityId.Get() == nil {
		var ret string
		return ret
	}

	return *o.EntityId.Get()
}

// GetEntityIdOk returns a tuple with the EntityId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiKey) GetEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EntityId.Get(), o.EntityId.IsSet()
}

// SetEntityId sets field value
func (o *AiKey) SetEntityId(v string) {
	o.EntityId.Set(&v)
}

// GetProjectId returns the ProjectId field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *AiKey) GetProjectId() int32 {
	if o == nil || o.ProjectId.Get() == nil {
		var ret int32
		return ret
	}

	return *o.ProjectId.Get()
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiKey) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectId.Get(), o.ProjectId.IsSet()
}

// SetProjectId sets field value
func (o *AiKey) SetProjectId(v int32) {
	o.ProjectId.Set(&v)
}

// GetExpiresAt returns the ExpiresAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *AiKey) GetExpiresAt() time.Time {
	if o == nil || o.ExpiresAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.ExpiresAt.Get()
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiKey) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiresAt.Get(), o.ExpiresAt.IsSet()
}

// SetExpiresAt sets field value
func (o *AiKey) SetExpiresAt(v time.Time) {
	o.ExpiresAt.Set(&v)
}

// GetIsReadonly returns the IsReadonly field value
func (o *AiKey) GetIsReadonly() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsReadonly
}

// GetIsReadonlyOk returns a tuple with the IsReadonly field value
// and a boolean to check if the value has been set.
func (o *AiKey) GetIsReadonlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsReadonly, true
}

// SetIsReadonly sets field value
func (o *AiKey) SetIsReadonly(v bool) {
	o.IsReadonly = v
}

// GetFallbackModels returns the FallbackModels field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *AiKey) GetFallbackModels() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.FallbackModels
}

// GetFallbackModelsOk returns a tuple with the FallbackModels field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiKey) GetFallbackModelsOk() ([]string, bool) {
	if o == nil || IsNil(o.FallbackModels) {
		return nil, false
	}
	return o.FallbackModels, true
}

// SetFallbackModels sets field value
func (o *AiKey) SetFallbackModels(v []string) {
	o.FallbackModels = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *AiKey) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AiKey) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AiKey) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *AiKey) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *AiKey) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *AiKey) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o AiKey) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["account_id"] = o.AccountId
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description.Get()
	toSerialize["entity_type"] = o.EntityType
	toSerialize["entity_id"] = o.EntityId.Get()
	toSerialize["project_id"] = o.ProjectId.Get()
	toSerialize["expires_at"] = o.ExpiresAt.Get()
	toSerialize["is_readonly"] = o.IsReadonly
	if o.FallbackModels != nil {
		toSerialize["fallback_models"] = o.FallbackModels
	}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

type NullableAiKey struct {
	value *AiKey
	isSet bool
}

func (v NullableAiKey) Get() *AiKey {
	return v.value
}

func (v *NullableAiKey) Set(val *AiKey) {
	v.value = val
	v.isSet = true
}

func (v NullableAiKey) IsSet() bool {
	return v.isSet
}

func (v *NullableAiKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiKey(val *AiKey) *NullableAiKey {
	return &NullableAiKey{value: val, isSet: true}
}

func (v NullableAiKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


