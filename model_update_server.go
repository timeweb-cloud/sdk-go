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

// checks if the UpdateServer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateServer{}

// UpdateServer struct for UpdateServer
type UpdateServer struct {
	Configurator *UpdateServerConfigurator `json:"configurator,omitempty"`
	// ID операционной системы, которая будет установлена на облачный сервер.
	OsId *float32 `json:"os_id,omitempty"`
	// ID программного обеспечения сервера.
	SoftwareId *float32 `json:"software_id,omitempty"`
	// ID тарифа сервера. Нельзя передавать вместе с ключом `configurator`.
	PresetId *float32 `json:"preset_id,omitempty"`
	// Пропускная способность тарифа. Доступные значения от 100 до 1000 с шагом 100.
	Bandwidth *float32 `json:"bandwidth,omitempty"`
	// Имя облачного сервера. Максимальная длина — 255 символов, имя должно быть уникальным.
	Name *string `json:"name,omitempty"`
	// ID аватара сервера. Описание методов работы с аватарами появится позднее.
	AvatarId *string `json:"avatar_id,omitempty"`
	// Комментарий к облачному серверу. Максимальная длина — 255 символов.
	Comment *string `json:"comment,omitempty"`
	// ID образа, который будет установлен на облачный сервер. Нельзя передавать вместе с `os_id`.
	ImageId *string `json:"image_id,omitempty"`
	// Cloud-init скрипт
	CloudInit *string `json:"cloud_init,omitempty"`
}

// NewUpdateServer instantiates a new UpdateServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateServer() *UpdateServer {
	this := UpdateServer{}
	return &this
}

// NewUpdateServerWithDefaults instantiates a new UpdateServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateServerWithDefaults() *UpdateServer {
	this := UpdateServer{}
	return &this
}

// GetConfigurator returns the Configurator field value if set, zero value otherwise.
func (o *UpdateServer) GetConfigurator() UpdateServerConfigurator {
	if o == nil || IsNil(o.Configurator) {
		var ret UpdateServerConfigurator
		return ret
	}
	return *o.Configurator
}

// GetConfiguratorOk returns a tuple with the Configurator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServer) GetConfiguratorOk() (*UpdateServerConfigurator, bool) {
	if o == nil || IsNil(o.Configurator) {
		return nil, false
	}
	return o.Configurator, true
}

// HasConfigurator returns a boolean if a field has been set.
func (o *UpdateServer) HasConfigurator() bool {
	if o != nil && !IsNil(o.Configurator) {
		return true
	}

	return false
}

// SetConfigurator gets a reference to the given UpdateServerConfigurator and assigns it to the Configurator field.
func (o *UpdateServer) SetConfigurator(v UpdateServerConfigurator) {
	o.Configurator = &v
}

// GetOsId returns the OsId field value if set, zero value otherwise.
func (o *UpdateServer) GetOsId() float32 {
	if o == nil || IsNil(o.OsId) {
		var ret float32
		return ret
	}
	return *o.OsId
}

// GetOsIdOk returns a tuple with the OsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServer) GetOsIdOk() (*float32, bool) {
	if o == nil || IsNil(o.OsId) {
		return nil, false
	}
	return o.OsId, true
}

// HasOsId returns a boolean if a field has been set.
func (o *UpdateServer) HasOsId() bool {
	if o != nil && !IsNil(o.OsId) {
		return true
	}

	return false
}

// SetOsId gets a reference to the given float32 and assigns it to the OsId field.
func (o *UpdateServer) SetOsId(v float32) {
	o.OsId = &v
}

// GetSoftwareId returns the SoftwareId field value if set, zero value otherwise.
func (o *UpdateServer) GetSoftwareId() float32 {
	if o == nil || IsNil(o.SoftwareId) {
		var ret float32
		return ret
	}
	return *o.SoftwareId
}

// GetSoftwareIdOk returns a tuple with the SoftwareId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServer) GetSoftwareIdOk() (*float32, bool) {
	if o == nil || IsNil(o.SoftwareId) {
		return nil, false
	}
	return o.SoftwareId, true
}

// HasSoftwareId returns a boolean if a field has been set.
func (o *UpdateServer) HasSoftwareId() bool {
	if o != nil && !IsNil(o.SoftwareId) {
		return true
	}

	return false
}

// SetSoftwareId gets a reference to the given float32 and assigns it to the SoftwareId field.
func (o *UpdateServer) SetSoftwareId(v float32) {
	o.SoftwareId = &v
}

// GetPresetId returns the PresetId field value if set, zero value otherwise.
func (o *UpdateServer) GetPresetId() float32 {
	if o == nil || IsNil(o.PresetId) {
		var ret float32
		return ret
	}
	return *o.PresetId
}

// GetPresetIdOk returns a tuple with the PresetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServer) GetPresetIdOk() (*float32, bool) {
	if o == nil || IsNil(o.PresetId) {
		return nil, false
	}
	return o.PresetId, true
}

// HasPresetId returns a boolean if a field has been set.
func (o *UpdateServer) HasPresetId() bool {
	if o != nil && !IsNil(o.PresetId) {
		return true
	}

	return false
}

// SetPresetId gets a reference to the given float32 and assigns it to the PresetId field.
func (o *UpdateServer) SetPresetId(v float32) {
	o.PresetId = &v
}

// GetBandwidth returns the Bandwidth field value if set, zero value otherwise.
func (o *UpdateServer) GetBandwidth() float32 {
	if o == nil || IsNil(o.Bandwidth) {
		var ret float32
		return ret
	}
	return *o.Bandwidth
}

// GetBandwidthOk returns a tuple with the Bandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServer) GetBandwidthOk() (*float32, bool) {
	if o == nil || IsNil(o.Bandwidth) {
		return nil, false
	}
	return o.Bandwidth, true
}

// HasBandwidth returns a boolean if a field has been set.
func (o *UpdateServer) HasBandwidth() bool {
	if o != nil && !IsNil(o.Bandwidth) {
		return true
	}

	return false
}

// SetBandwidth gets a reference to the given float32 and assigns it to the Bandwidth field.
func (o *UpdateServer) SetBandwidth(v float32) {
	o.Bandwidth = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateServer) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServer) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateServer) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateServer) SetName(v string) {
	o.Name = &v
}

// GetAvatarId returns the AvatarId field value if set, zero value otherwise.
func (o *UpdateServer) GetAvatarId() string {
	if o == nil || IsNil(o.AvatarId) {
		var ret string
		return ret
	}
	return *o.AvatarId
}

// GetAvatarIdOk returns a tuple with the AvatarId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServer) GetAvatarIdOk() (*string, bool) {
	if o == nil || IsNil(o.AvatarId) {
		return nil, false
	}
	return o.AvatarId, true
}

// HasAvatarId returns a boolean if a field has been set.
func (o *UpdateServer) HasAvatarId() bool {
	if o != nil && !IsNil(o.AvatarId) {
		return true
	}

	return false
}

// SetAvatarId gets a reference to the given string and assigns it to the AvatarId field.
func (o *UpdateServer) SetAvatarId(v string) {
	o.AvatarId = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *UpdateServer) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServer) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *UpdateServer) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *UpdateServer) SetComment(v string) {
	o.Comment = &v
}

// GetImageId returns the ImageId field value if set, zero value otherwise.
func (o *UpdateServer) GetImageId() string {
	if o == nil || IsNil(o.ImageId) {
		var ret string
		return ret
	}
	return *o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServer) GetImageIdOk() (*string, bool) {
	if o == nil || IsNil(o.ImageId) {
		return nil, false
	}
	return o.ImageId, true
}

// HasImageId returns a boolean if a field has been set.
func (o *UpdateServer) HasImageId() bool {
	if o != nil && !IsNil(o.ImageId) {
		return true
	}

	return false
}

// SetImageId gets a reference to the given string and assigns it to the ImageId field.
func (o *UpdateServer) SetImageId(v string) {
	o.ImageId = &v
}

// GetCloudInit returns the CloudInit field value if set, zero value otherwise.
func (o *UpdateServer) GetCloudInit() string {
	if o == nil || IsNil(o.CloudInit) {
		var ret string
		return ret
	}
	return *o.CloudInit
}

// GetCloudInitOk returns a tuple with the CloudInit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServer) GetCloudInitOk() (*string, bool) {
	if o == nil || IsNil(o.CloudInit) {
		return nil, false
	}
	return o.CloudInit, true
}

// HasCloudInit returns a boolean if a field has been set.
func (o *UpdateServer) HasCloudInit() bool {
	if o != nil && !IsNil(o.CloudInit) {
		return true
	}

	return false
}

// SetCloudInit gets a reference to the given string and assigns it to the CloudInit field.
func (o *UpdateServer) SetCloudInit(v string) {
	o.CloudInit = &v
}

func (o UpdateServer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateServer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Configurator) {
		toSerialize["configurator"] = o.Configurator
	}
	if !IsNil(o.OsId) {
		toSerialize["os_id"] = o.OsId
	}
	if !IsNil(o.SoftwareId) {
		toSerialize["software_id"] = o.SoftwareId
	}
	if !IsNil(o.PresetId) {
		toSerialize["preset_id"] = o.PresetId
	}
	if !IsNil(o.Bandwidth) {
		toSerialize["bandwidth"] = o.Bandwidth
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.AvatarId) {
		toSerialize["avatar_id"] = o.AvatarId
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.ImageId) {
		toSerialize["image_id"] = o.ImageId
	}
	if !IsNil(o.CloudInit) {
		toSerialize["cloud_init"] = o.CloudInit
	}
	return toSerialize, nil
}

type NullableUpdateServer struct {
	value *UpdateServer
	isSet bool
}

func (v NullableUpdateServer) Get() *UpdateServer {
	return v.value
}

func (v *NullableUpdateServer) Set(val *UpdateServer) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateServer) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateServer(val *UpdateServer) *NullableUpdateServer {
	return &NullableUpdateServer{value: val, isSet: true}
}

func (v NullableUpdateServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


