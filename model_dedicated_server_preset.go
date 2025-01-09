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

// checks if the DedicatedServerPreset type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DedicatedServerPreset{}

// DedicatedServerPreset Выделенный сервер
type DedicatedServerPreset struct {
	// ID тарифа выделенного сервера.
	Id float32 `json:"id"`
	// Описание характеристик тарифа выделенного сервера.
	Description string `json:"description"`
	// Это логическое значение, которое показывает, доступен ли IPMI у данного тарифа.
	IsIpmiEnabled bool `json:"is_ipmi_enabled"`
	// Это логическое значение, которое показывает, готов ли выделенный сервер к моментальной выдаче.
	IsPreInstalled bool `json:"is_pre_installed"`
	Cpu DedicatedServerPresetCpu `json:"cpu"`
	Disk DedicatedServerPresetDisk `json:"disk"`
	// Стоимость тарифа выделенного сервера
	Price *float32 `json:"price,omitempty"`
	Memory DedicatedServerPresetMemory `json:"memory"`
	// Локация.
	Location string `json:"location"`
}

// NewDedicatedServerPreset instantiates a new DedicatedServerPreset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDedicatedServerPreset(id float32, description string, isIpmiEnabled bool, isPreInstalled bool, cpu DedicatedServerPresetCpu, disk DedicatedServerPresetDisk, memory DedicatedServerPresetMemory, location string) *DedicatedServerPreset {
	this := DedicatedServerPreset{}
	this.Id = id
	this.Description = description
	this.IsIpmiEnabled = isIpmiEnabled
	this.IsPreInstalled = isPreInstalled
	this.Cpu = cpu
	this.Disk = disk
	this.Memory = memory
	this.Location = location
	return &this
}

// NewDedicatedServerPresetWithDefaults instantiates a new DedicatedServerPreset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDedicatedServerPresetWithDefaults() *DedicatedServerPreset {
	this := DedicatedServerPreset{}
	return &this
}

// GetId returns the Id field value
func (o *DedicatedServerPreset) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DedicatedServerPreset) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DedicatedServerPreset) SetId(v float32) {
	o.Id = v
}

// GetDescription returns the Description field value
func (o *DedicatedServerPreset) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *DedicatedServerPreset) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *DedicatedServerPreset) SetDescription(v string) {
	o.Description = v
}

// GetIsIpmiEnabled returns the IsIpmiEnabled field value
func (o *DedicatedServerPreset) GetIsIpmiEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsIpmiEnabled
}

// GetIsIpmiEnabledOk returns a tuple with the IsIpmiEnabled field value
// and a boolean to check if the value has been set.
func (o *DedicatedServerPreset) GetIsIpmiEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsIpmiEnabled, true
}

// SetIsIpmiEnabled sets field value
func (o *DedicatedServerPreset) SetIsIpmiEnabled(v bool) {
	o.IsIpmiEnabled = v
}

// GetIsPreInstalled returns the IsPreInstalled field value
func (o *DedicatedServerPreset) GetIsPreInstalled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPreInstalled
}

// GetIsPreInstalledOk returns a tuple with the IsPreInstalled field value
// and a boolean to check if the value has been set.
func (o *DedicatedServerPreset) GetIsPreInstalledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPreInstalled, true
}

// SetIsPreInstalled sets field value
func (o *DedicatedServerPreset) SetIsPreInstalled(v bool) {
	o.IsPreInstalled = v
}

// GetCpu returns the Cpu field value
func (o *DedicatedServerPreset) GetCpu() DedicatedServerPresetCpu {
	if o == nil {
		var ret DedicatedServerPresetCpu
		return ret
	}

	return o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value
// and a boolean to check if the value has been set.
func (o *DedicatedServerPreset) GetCpuOk() (*DedicatedServerPresetCpu, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cpu, true
}

// SetCpu sets field value
func (o *DedicatedServerPreset) SetCpu(v DedicatedServerPresetCpu) {
	o.Cpu = v
}

// GetDisk returns the Disk field value
func (o *DedicatedServerPreset) GetDisk() DedicatedServerPresetDisk {
	if o == nil {
		var ret DedicatedServerPresetDisk
		return ret
	}

	return o.Disk
}

// GetDiskOk returns a tuple with the Disk field value
// and a boolean to check if the value has been set.
func (o *DedicatedServerPreset) GetDiskOk() (*DedicatedServerPresetDisk, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Disk, true
}

// SetDisk sets field value
func (o *DedicatedServerPreset) SetDisk(v DedicatedServerPresetDisk) {
	o.Disk = v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *DedicatedServerPreset) GetPrice() float32 {
	if o == nil || IsNil(o.Price) {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DedicatedServerPreset) GetPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *DedicatedServerPreset) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *DedicatedServerPreset) SetPrice(v float32) {
	o.Price = &v
}

// GetMemory returns the Memory field value
func (o *DedicatedServerPreset) GetMemory() DedicatedServerPresetMemory {
	if o == nil {
		var ret DedicatedServerPresetMemory
		return ret
	}

	return o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value
// and a boolean to check if the value has been set.
func (o *DedicatedServerPreset) GetMemoryOk() (*DedicatedServerPresetMemory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Memory, true
}

// SetMemory sets field value
func (o *DedicatedServerPreset) SetMemory(v DedicatedServerPresetMemory) {
	o.Memory = v
}

// GetLocation returns the Location field value
func (o *DedicatedServerPreset) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *DedicatedServerPreset) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *DedicatedServerPreset) SetLocation(v string) {
	o.Location = v
}

func (o DedicatedServerPreset) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DedicatedServerPreset) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["description"] = o.Description
	toSerialize["is_ipmi_enabled"] = o.IsIpmiEnabled
	toSerialize["is_pre_installed"] = o.IsPreInstalled
	toSerialize["cpu"] = o.Cpu
	toSerialize["disk"] = o.Disk
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	toSerialize["memory"] = o.Memory
	toSerialize["location"] = o.Location
	return toSerialize, nil
}

type NullableDedicatedServerPreset struct {
	value *DedicatedServerPreset
	isSet bool
}

func (v NullableDedicatedServerPreset) Get() *DedicatedServerPreset {
	return v.value
}

func (v *NullableDedicatedServerPreset) Set(val *DedicatedServerPreset) {
	v.value = val
	v.isSet = true
}

func (v NullableDedicatedServerPreset) IsSet() bool {
	return v.isSet
}

func (v *NullableDedicatedServerPreset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDedicatedServerPreset(val *DedicatedServerPreset) *NullableDedicatedServerPreset {
	return &NullableDedicatedServerPreset{value: val, isSet: true}
}

func (v NullableDedicatedServerPreset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDedicatedServerPreset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


