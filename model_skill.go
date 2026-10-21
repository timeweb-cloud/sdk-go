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

// checks if the Skill type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Skill{}

// Skill Навык
type Skill struct {
	// ID навыка
	Id int32 `json:"id"`
	// Название навыка, уникальное в аккаунте
	Name string `json:"name"`
	// Когда применять навык. Это поле читает модель, когда выбирает навык
	WhenToUse string `json:"when_to_use"`
	// Инструкция навыка. Попадает в контекст модели только при срабатывании
	Instruction string `json:"instruction"`
	// Номер текущей версии
	CurrentVersion int32 `json:"current_version"`
	// У скольких агентов навык включён
	AgentsCount int32 `json:"agents_count"`
	// Имя агента, если навык включён ровно у одного агента
	FirstAgentName NullableString `json:"first_agent_name"`
	// ID навыка в каталоге, из которого установлен этот навык
	CatalogItemId NullableInt32 `json:"catalog_item_id"`
	// Версия каталожного навыка, из которой установлен или обновлён этот навык
	CatalogVersion NullableInt32 `json:"catalog_version"`
	// Навык установлен из каталога
	IsFromCatalog bool `json:"is_from_catalog"`
	// Установленный из каталога навык изменялся владельцем
	IsModified bool `json:"is_modified"`
	// Дата и время создания
	CreatedAt time.Time `json:"created_at"`
	// Дата и время последнего изменения
	UpdatedAt time.Time `json:"updated_at"`
}

// NewSkill instantiates a new Skill object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkill(id int32, name string, whenToUse string, instruction string, currentVersion int32, agentsCount int32, firstAgentName NullableString, catalogItemId NullableInt32, catalogVersion NullableInt32, isFromCatalog bool, isModified bool, createdAt time.Time, updatedAt time.Time) *Skill {
	this := Skill{}
	this.Id = id
	this.Name = name
	this.WhenToUse = whenToUse
	this.Instruction = instruction
	this.CurrentVersion = currentVersion
	this.AgentsCount = agentsCount
	this.FirstAgentName = firstAgentName
	this.CatalogItemId = catalogItemId
	this.CatalogVersion = catalogVersion
	this.IsFromCatalog = isFromCatalog
	this.IsModified = isModified
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewSkillWithDefaults instantiates a new Skill object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkillWithDefaults() *Skill {
	this := Skill{}
	return &this
}

// GetId returns the Id field value
func (o *Skill) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Skill) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Skill) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Skill) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Skill) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Skill) SetName(v string) {
	o.Name = v
}

// GetWhenToUse returns the WhenToUse field value
func (o *Skill) GetWhenToUse() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WhenToUse
}

// GetWhenToUseOk returns a tuple with the WhenToUse field value
// and a boolean to check if the value has been set.
func (o *Skill) GetWhenToUseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WhenToUse, true
}

// SetWhenToUse sets field value
func (o *Skill) SetWhenToUse(v string) {
	o.WhenToUse = v
}

// GetInstruction returns the Instruction field value
func (o *Skill) GetInstruction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Instruction
}

// GetInstructionOk returns a tuple with the Instruction field value
// and a boolean to check if the value has been set.
func (o *Skill) GetInstructionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Instruction, true
}

// SetInstruction sets field value
func (o *Skill) SetInstruction(v string) {
	o.Instruction = v
}

// GetCurrentVersion returns the CurrentVersion field value
func (o *Skill) GetCurrentVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CurrentVersion
}

// GetCurrentVersionOk returns a tuple with the CurrentVersion field value
// and a boolean to check if the value has been set.
func (o *Skill) GetCurrentVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentVersion, true
}

// SetCurrentVersion sets field value
func (o *Skill) SetCurrentVersion(v int32) {
	o.CurrentVersion = v
}

// GetAgentsCount returns the AgentsCount field value
func (o *Skill) GetAgentsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AgentsCount
}

// GetAgentsCountOk returns a tuple with the AgentsCount field value
// and a boolean to check if the value has been set.
func (o *Skill) GetAgentsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentsCount, true
}

// SetAgentsCount sets field value
func (o *Skill) SetAgentsCount(v int32) {
	o.AgentsCount = v
}

// GetFirstAgentName returns the FirstAgentName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Skill) GetFirstAgentName() string {
	if o == nil || o.FirstAgentName.Get() == nil {
		var ret string
		return ret
	}

	return *o.FirstAgentName.Get()
}

// GetFirstAgentNameOk returns a tuple with the FirstAgentName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Skill) GetFirstAgentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstAgentName.Get(), o.FirstAgentName.IsSet()
}

// SetFirstAgentName sets field value
func (o *Skill) SetFirstAgentName(v string) {
	o.FirstAgentName.Set(&v)
}

// GetCatalogItemId returns the CatalogItemId field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *Skill) GetCatalogItemId() int32 {
	if o == nil || o.CatalogItemId.Get() == nil {
		var ret int32
		return ret
	}

	return *o.CatalogItemId.Get()
}

// GetCatalogItemIdOk returns a tuple with the CatalogItemId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Skill) GetCatalogItemIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CatalogItemId.Get(), o.CatalogItemId.IsSet()
}

// SetCatalogItemId sets field value
func (o *Skill) SetCatalogItemId(v int32) {
	o.CatalogItemId.Set(&v)
}

// GetCatalogVersion returns the CatalogVersion field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *Skill) GetCatalogVersion() int32 {
	if o == nil || o.CatalogVersion.Get() == nil {
		var ret int32
		return ret
	}

	return *o.CatalogVersion.Get()
}

// GetCatalogVersionOk returns a tuple with the CatalogVersion field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Skill) GetCatalogVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CatalogVersion.Get(), o.CatalogVersion.IsSet()
}

// SetCatalogVersion sets field value
func (o *Skill) SetCatalogVersion(v int32) {
	o.CatalogVersion.Set(&v)
}

// GetIsFromCatalog returns the IsFromCatalog field value
func (o *Skill) GetIsFromCatalog() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsFromCatalog
}

// GetIsFromCatalogOk returns a tuple with the IsFromCatalog field value
// and a boolean to check if the value has been set.
func (o *Skill) GetIsFromCatalogOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsFromCatalog, true
}

// SetIsFromCatalog sets field value
func (o *Skill) SetIsFromCatalog(v bool) {
	o.IsFromCatalog = v
}

// GetIsModified returns the IsModified field value
func (o *Skill) GetIsModified() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsModified
}

// GetIsModifiedOk returns a tuple with the IsModified field value
// and a boolean to check if the value has been set.
func (o *Skill) GetIsModifiedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsModified, true
}

// SetIsModified sets field value
func (o *Skill) SetIsModified(v bool) {
	o.IsModified = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Skill) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Skill) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Skill) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Skill) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Skill) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Skill) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o Skill) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Skill) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["when_to_use"] = o.WhenToUse
	toSerialize["instruction"] = o.Instruction
	toSerialize["current_version"] = o.CurrentVersion
	toSerialize["agents_count"] = o.AgentsCount
	toSerialize["first_agent_name"] = o.FirstAgentName.Get()
	toSerialize["catalog_item_id"] = o.CatalogItemId.Get()
	toSerialize["catalog_version"] = o.CatalogVersion.Get()
	toSerialize["is_from_catalog"] = o.IsFromCatalog
	toSerialize["is_modified"] = o.IsModified
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

type NullableSkill struct {
	value *Skill
	isSet bool
}

func (v NullableSkill) Get() *Skill {
	return v.value
}

func (v *NullableSkill) Set(val *Skill) {
	v.value = val
	v.isSet = true
}

func (v NullableSkill) IsSet() bool {
	return v.isSet
}

func (v *NullableSkill) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkill(val *Skill) *NullableSkill {
	return &NullableSkill{value: val, isSet: true}
}

func (v NullableSkill) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkill) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


