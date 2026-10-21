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

// checks if the SkillCatalogItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkillCatalogItem{}

// SkillCatalogItem Навык из каталога Timeweb Cloud
type SkillCatalogItem struct {
	// ID навыка в каталоге
	Id int32 `json:"id"`
	// Уникальный ключ навыка в каталоге
	Key string `json:"key"`
	// Назначение: `communication` — манера общения, `support` — регламенты обслуживания, `sales` — продажи, `content` — документы и база знаний, `safety` — ограничения и безопасность
	Category string `json:"category"`
	// Название навыка
	Name string `json:"name"`
	// Краткое описание навыка
	Description NullableString `json:"description"`
	// Когда применять навык
	WhenToUse string `json:"when_to_use"`
	// Полная инструкция навыка
	Instruction string `json:"instruction"`
	// Версия навыка в каталоге
	Version int32 `json:"version"`
	// Навык уже установлен в аккаунте
	IsInstalled bool `json:"is_installed"`
	// ID установленного навыка в аккаунте
	InstalledSkillId NullableInt32 `json:"installed_skill_id"`
	// Установленный навык изменялся владельцем
	IsModified bool `json:"is_modified"`
	// В каталоге есть версия новее установленной
	HasCatalogUpdate bool `json:"has_catalog_update"`
}

// NewSkillCatalogItem instantiates a new SkillCatalogItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkillCatalogItem(id int32, key string, category string, name string, description NullableString, whenToUse string, instruction string, version int32, isInstalled bool, installedSkillId NullableInt32, isModified bool, hasCatalogUpdate bool) *SkillCatalogItem {
	this := SkillCatalogItem{}
	this.Id = id
	this.Key = key
	this.Category = category
	this.Name = name
	this.Description = description
	this.WhenToUse = whenToUse
	this.Instruction = instruction
	this.Version = version
	this.IsInstalled = isInstalled
	this.InstalledSkillId = installedSkillId
	this.IsModified = isModified
	this.HasCatalogUpdate = hasCatalogUpdate
	return &this
}

// NewSkillCatalogItemWithDefaults instantiates a new SkillCatalogItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkillCatalogItemWithDefaults() *SkillCatalogItem {
	this := SkillCatalogItem{}
	return &this
}

// GetId returns the Id field value
func (o *SkillCatalogItem) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SkillCatalogItem) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SkillCatalogItem) SetId(v int32) {
	o.Id = v
}

// GetKey returns the Key field value
func (o *SkillCatalogItem) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *SkillCatalogItem) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *SkillCatalogItem) SetKey(v string) {
	o.Key = v
}

// GetCategory returns the Category field value
func (o *SkillCatalogItem) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *SkillCatalogItem) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *SkillCatalogItem) SetCategory(v string) {
	o.Category = v
}

// GetName returns the Name field value
func (o *SkillCatalogItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SkillCatalogItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SkillCatalogItem) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SkillCatalogItem) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SkillCatalogItem) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *SkillCatalogItem) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetWhenToUse returns the WhenToUse field value
func (o *SkillCatalogItem) GetWhenToUse() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WhenToUse
}

// GetWhenToUseOk returns a tuple with the WhenToUse field value
// and a boolean to check if the value has been set.
func (o *SkillCatalogItem) GetWhenToUseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WhenToUse, true
}

// SetWhenToUse sets field value
func (o *SkillCatalogItem) SetWhenToUse(v string) {
	o.WhenToUse = v
}

// GetInstruction returns the Instruction field value
func (o *SkillCatalogItem) GetInstruction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Instruction
}

// GetInstructionOk returns a tuple with the Instruction field value
// and a boolean to check if the value has been set.
func (o *SkillCatalogItem) GetInstructionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Instruction, true
}

// SetInstruction sets field value
func (o *SkillCatalogItem) SetInstruction(v string) {
	o.Instruction = v
}

// GetVersion returns the Version field value
func (o *SkillCatalogItem) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *SkillCatalogItem) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *SkillCatalogItem) SetVersion(v int32) {
	o.Version = v
}

// GetIsInstalled returns the IsInstalled field value
func (o *SkillCatalogItem) GetIsInstalled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsInstalled
}

// GetIsInstalledOk returns a tuple with the IsInstalled field value
// and a boolean to check if the value has been set.
func (o *SkillCatalogItem) GetIsInstalledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsInstalled, true
}

// SetIsInstalled sets field value
func (o *SkillCatalogItem) SetIsInstalled(v bool) {
	o.IsInstalled = v
}

// GetInstalledSkillId returns the InstalledSkillId field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *SkillCatalogItem) GetInstalledSkillId() int32 {
	if o == nil || o.InstalledSkillId.Get() == nil {
		var ret int32
		return ret
	}

	return *o.InstalledSkillId.Get()
}

// GetInstalledSkillIdOk returns a tuple with the InstalledSkillId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SkillCatalogItem) GetInstalledSkillIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.InstalledSkillId.Get(), o.InstalledSkillId.IsSet()
}

// SetInstalledSkillId sets field value
func (o *SkillCatalogItem) SetInstalledSkillId(v int32) {
	o.InstalledSkillId.Set(&v)
}

// GetIsModified returns the IsModified field value
func (o *SkillCatalogItem) GetIsModified() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsModified
}

// GetIsModifiedOk returns a tuple with the IsModified field value
// and a boolean to check if the value has been set.
func (o *SkillCatalogItem) GetIsModifiedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsModified, true
}

// SetIsModified sets field value
func (o *SkillCatalogItem) SetIsModified(v bool) {
	o.IsModified = v
}

// GetHasCatalogUpdate returns the HasCatalogUpdate field value
func (o *SkillCatalogItem) GetHasCatalogUpdate() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasCatalogUpdate
}

// GetHasCatalogUpdateOk returns a tuple with the HasCatalogUpdate field value
// and a boolean to check if the value has been set.
func (o *SkillCatalogItem) GetHasCatalogUpdateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasCatalogUpdate, true
}

// SetHasCatalogUpdate sets field value
func (o *SkillCatalogItem) SetHasCatalogUpdate(v bool) {
	o.HasCatalogUpdate = v
}

func (o SkillCatalogItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkillCatalogItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["key"] = o.Key
	toSerialize["category"] = o.Category
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description.Get()
	toSerialize["when_to_use"] = o.WhenToUse
	toSerialize["instruction"] = o.Instruction
	toSerialize["version"] = o.Version
	toSerialize["is_installed"] = o.IsInstalled
	toSerialize["installed_skill_id"] = o.InstalledSkillId.Get()
	toSerialize["is_modified"] = o.IsModified
	toSerialize["has_catalog_update"] = o.HasCatalogUpdate
	return toSerialize, nil
}

type NullableSkillCatalogItem struct {
	value *SkillCatalogItem
	isSet bool
}

func (v NullableSkillCatalogItem) Get() *SkillCatalogItem {
	return v.value
}

func (v *NullableSkillCatalogItem) Set(val *SkillCatalogItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSkillCatalogItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSkillCatalogItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkillCatalogItem(val *SkillCatalogItem) *NullableSkillCatalogItem {
	return &NullableSkillCatalogItem{value: val, isSet: true}
}

func (v NullableSkillCatalogItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkillCatalogItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


