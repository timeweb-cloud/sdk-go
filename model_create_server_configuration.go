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

// checks if the CreateServerConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateServerConfiguration{}

// CreateServerConfiguration Параметры конфигурации сервера. Нельзя передавать вместе с `preset_id`.
type CreateServerConfiguration struct {
	// ID конфигуратора сервера.
	ConfiguratorId float32 `json:"configurator_id"`
	// Размер диска в МБ.
	Disk float32 `json:"disk"`
	// Количество ядер процессора.
	Cpu float32 `json:"cpu"`
	// Размер ОЗУ сервера в МБ.
	Ram float32 `json:"ram"`
	// Количество видеокарт.
	Gpu *float32 `json:"gpu,omitempty"`
}

// NewCreateServerConfiguration instantiates a new CreateServerConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateServerConfiguration(configuratorId float32, disk float32, cpu float32, ram float32) *CreateServerConfiguration {
	this := CreateServerConfiguration{}
	this.ConfiguratorId = configuratorId
	this.Disk = disk
	this.Cpu = cpu
	this.Ram = ram
	return &this
}

// NewCreateServerConfigurationWithDefaults instantiates a new CreateServerConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateServerConfigurationWithDefaults() *CreateServerConfiguration {
	this := CreateServerConfiguration{}
	return &this
}

// GetConfiguratorId returns the ConfiguratorId field value
func (o *CreateServerConfiguration) GetConfiguratorId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ConfiguratorId
}

// GetConfiguratorIdOk returns a tuple with the ConfiguratorId field value
// and a boolean to check if the value has been set.
func (o *CreateServerConfiguration) GetConfiguratorIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfiguratorId, true
}

// SetConfiguratorId sets field value
func (o *CreateServerConfiguration) SetConfiguratorId(v float32) {
	o.ConfiguratorId = v
}

// GetDisk returns the Disk field value
func (o *CreateServerConfiguration) GetDisk() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Disk
}

// GetDiskOk returns a tuple with the Disk field value
// and a boolean to check if the value has been set.
func (o *CreateServerConfiguration) GetDiskOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Disk, true
}

// SetDisk sets field value
func (o *CreateServerConfiguration) SetDisk(v float32) {
	o.Disk = v
}

// GetCpu returns the Cpu field value
func (o *CreateServerConfiguration) GetCpu() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value
// and a boolean to check if the value has been set.
func (o *CreateServerConfiguration) GetCpuOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cpu, true
}

// SetCpu sets field value
func (o *CreateServerConfiguration) SetCpu(v float32) {
	o.Cpu = v
}

// GetRam returns the Ram field value
func (o *CreateServerConfiguration) GetRam() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Ram
}

// GetRamOk returns a tuple with the Ram field value
// and a boolean to check if the value has been set.
func (o *CreateServerConfiguration) GetRamOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ram, true
}

// SetRam sets field value
func (o *CreateServerConfiguration) SetRam(v float32) {
	o.Ram = v
}

// GetGpu returns the Gpu field value if set, zero value otherwise.
func (o *CreateServerConfiguration) GetGpu() float32 {
	if o == nil || IsNil(o.Gpu) {
		var ret float32
		return ret
	}
	return *o.Gpu
}

// GetGpuOk returns a tuple with the Gpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateServerConfiguration) GetGpuOk() (*float32, bool) {
	if o == nil || IsNil(o.Gpu) {
		return nil, false
	}
	return o.Gpu, true
}

// HasGpu returns a boolean if a field has been set.
func (o *CreateServerConfiguration) HasGpu() bool {
	if o != nil && !IsNil(o.Gpu) {
		return true
	}

	return false
}

// SetGpu gets a reference to the given float32 and assigns it to the Gpu field.
func (o *CreateServerConfiguration) SetGpu(v float32) {
	o.Gpu = &v
}

func (o CreateServerConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateServerConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["configurator_id"] = o.ConfiguratorId
	toSerialize["disk"] = o.Disk
	toSerialize["cpu"] = o.Cpu
	toSerialize["ram"] = o.Ram
	if !IsNil(o.Gpu) {
		toSerialize["gpu"] = o.Gpu
	}
	return toSerialize, nil
}

type NullableCreateServerConfiguration struct {
	value *CreateServerConfiguration
	isSet bool
}

func (v NullableCreateServerConfiguration) Get() *CreateServerConfiguration {
	return v.value
}

func (v *NullableCreateServerConfiguration) Set(val *CreateServerConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateServerConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateServerConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateServerConfiguration(val *CreateServerConfiguration) *NullableCreateServerConfiguration {
	return &NullableCreateServerConfiguration{value: val, isSet: true}
}

func (v NullableCreateServerConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateServerConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


