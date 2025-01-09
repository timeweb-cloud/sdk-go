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

// checks if the Image type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Image{}

// Image struct for Image
type Image struct {
	// ID образа.
	Id string `json:"id"`
	Status ImageStatus `json:"status"`
	// Значение времени, указанное в комбинированном формате даты и времени ISO8601, которое представляет, когда был создан образ.
	CreatedAt string `json:"created_at"`
	// Значение времени, указанное в комбинированном формате даты и времени ISO8601, которое представляет, когда был удален образ.
	DeletedAt string `json:"deleted_at"`
	// Размер физического диска в мегабайтах.
	Size int32 `json:"size"`
	// Размер виртуального диска в мегабайтах.
	VirtualSize int32 `json:"virtual_size"`
	// Имя образа.
	Name string `json:"name"`
	// Описание образа.
	Description string `json:"description"`
	// ID связанного с образом диска.
	DiskId int32 `json:"disk_id"`
	// Локация образа.
	Location string `json:"location"`
	Os OS `json:"os"`
	// Процент создания образа.
	Progress int32 `json:"progress"`
	// Логическое значение, которое показывает, является ли образ кастомным.
	IsCustom bool `json:"is_custom"`
	// Тип образа.
	Type string `json:"type"`
}

// NewImage instantiates a new Image object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImage(id string, status ImageStatus, createdAt string, deletedAt string, size int32, virtualSize int32, name string, description string, diskId int32, location string, os OS, progress int32, isCustom bool, type_ string) *Image {
	this := Image{}
	this.Id = id
	this.Status = status
	this.CreatedAt = createdAt
	this.DeletedAt = deletedAt
	this.Size = size
	this.VirtualSize = virtualSize
	this.Name = name
	this.Description = description
	this.DiskId = diskId
	this.Location = location
	this.Os = os
	this.Progress = progress
	this.IsCustom = isCustom
	this.Type = type_
	return &this
}

// NewImageWithDefaults instantiates a new Image object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageWithDefaults() *Image {
	this := Image{}
	return &this
}

// GetId returns the Id field value
func (o *Image) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Image) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Image) SetId(v string) {
	o.Id = v
}

// GetStatus returns the Status field value
func (o *Image) GetStatus() ImageStatus {
	if o == nil {
		var ret ImageStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Image) GetStatusOk() (*ImageStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Image) SetStatus(v ImageStatus) {
	o.Status = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Image) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Image) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Image) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetDeletedAt returns the DeletedAt field value
func (o *Image) GetDeletedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value
// and a boolean to check if the value has been set.
func (o *Image) GetDeletedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeletedAt, true
}

// SetDeletedAt sets field value
func (o *Image) SetDeletedAt(v string) {
	o.DeletedAt = v
}

// GetSize returns the Size field value
func (o *Image) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *Image) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *Image) SetSize(v int32) {
	o.Size = v
}

// GetVirtualSize returns the VirtualSize field value
func (o *Image) GetVirtualSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VirtualSize
}

// GetVirtualSizeOk returns a tuple with the VirtualSize field value
// and a boolean to check if the value has been set.
func (o *Image) GetVirtualSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VirtualSize, true
}

// SetVirtualSize sets field value
func (o *Image) SetVirtualSize(v int32) {
	o.VirtualSize = v
}

// GetName returns the Name field value
func (o *Image) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Image) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Image) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *Image) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Image) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Image) SetDescription(v string) {
	o.Description = v
}

// GetDiskId returns the DiskId field value
func (o *Image) GetDiskId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DiskId
}

// GetDiskIdOk returns a tuple with the DiskId field value
// and a boolean to check if the value has been set.
func (o *Image) GetDiskIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiskId, true
}

// SetDiskId sets field value
func (o *Image) SetDiskId(v int32) {
	o.DiskId = v
}

// GetLocation returns the Location field value
func (o *Image) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *Image) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *Image) SetLocation(v string) {
	o.Location = v
}

// GetOs returns the Os field value
func (o *Image) GetOs() OS {
	if o == nil {
		var ret OS
		return ret
	}

	return o.Os
}

// GetOsOk returns a tuple with the Os field value
// and a boolean to check if the value has been set.
func (o *Image) GetOsOk() (*OS, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Os, true
}

// SetOs sets field value
func (o *Image) SetOs(v OS) {
	o.Os = v
}

// GetProgress returns the Progress field value
func (o *Image) GetProgress() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Progress
}

// GetProgressOk returns a tuple with the Progress field value
// and a boolean to check if the value has been set.
func (o *Image) GetProgressOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Progress, true
}

// SetProgress sets field value
func (o *Image) SetProgress(v int32) {
	o.Progress = v
}

// GetIsCustom returns the IsCustom field value
func (o *Image) GetIsCustom() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsCustom
}

// GetIsCustomOk returns a tuple with the IsCustom field value
// and a boolean to check if the value has been set.
func (o *Image) GetIsCustomOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsCustom, true
}

// SetIsCustom sets field value
func (o *Image) SetIsCustom(v bool) {
	o.IsCustom = v
}

// GetType returns the Type field value
func (o *Image) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Image) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Image) SetType(v string) {
	o.Type = v
}

func (o Image) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Image) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["status"] = o.Status
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["deleted_at"] = o.DeletedAt
	toSerialize["size"] = o.Size
	toSerialize["virtual_size"] = o.VirtualSize
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["disk_id"] = o.DiskId
	toSerialize["location"] = o.Location
	toSerialize["os"] = o.Os
	toSerialize["progress"] = o.Progress
	toSerialize["is_custom"] = o.IsCustom
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableImage struct {
	value *Image
	isSet bool
}

func (v NullableImage) Get() *Image {
	return v.value
}

func (v *NullableImage) Set(val *Image) {
	v.value = val
	v.isSet = true
}

func (v NullableImage) IsSet() bool {
	return v.isSet
}

func (v *NullableImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImage(val *Image) *NullableImage {
	return &NullableImage{value: val, isSet: true}
}

func (v NullableImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


