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

// checks if the HttpResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HttpResource{}

// HttpResource CDN-ресурс
type HttpResource struct {
	// ID CDN-ресурса. Генерируется автоматически при создании.
	Id int32 `json:"id"`
	// Название CDN-ресурса.
	Name string `json:"name"`
	// Описание CDN-ресурса.
	Description *string `json:"description,omitempty"`
	// Источник контента: имя S3-бакета, если ресурс создан поверх хранилища, либо хост origin-сервера.
	Source string `json:"source"`
	TrafficUsage TrafficUsage `json:"traffic_usage"`
	// Статус CDN-ресурса. - `created` — ресурс создан и раздает контент; - `processing` — конфигурация применяется на стороне CDN; - `stopped` — раздача приостановлена; - `failed` — настройка ресурса завершилась с ошибкой; - `no_paid` — ресурс не оплачен; - `blocked` — ресурс заблокирован; - `traffic_limit_exceeded` — раздача остановлена автоматически из-за достижения лимита трафика `traffic_limit_bytes`.
	Status string `json:"status"`
	// Технический домен, выданный ресурсу. Доступен сразу после создания и всегда остается в списке доменов ресурса.
	CdnDomain string `json:"cdn_domain"`
	// ID тарифа CDN. Список доступных тарифов можно получить в `/api/v1/cdn/presets`.
	PresetId int32 `json:"preset_id"`
	// ID проекта, к которому привязан ресурс.
	ProjectId *int32 `json:"project_id,omitempty"`
	// Ссылка на аватар ресурса.
	AvatarLink NullableString `json:"avatar_link,omitempty"`
	// ID S3-хранилища, которое используется в качестве источника контента. `null`, если источником является origin-сервер.
	StorageId NullableInt32 `json:"storage_id,omitempty"`
	// Лимит исходящего трафика на расчетный месяц, в байтах. При достижении лимита раздача останавливается, а ресурс переходит в статус `traffic_limit_exceeded`. `null` — лимит не задан.
	TrafficLimitBytes NullableInt32 `json:"traffic_limit_bytes,omitempty"`
}

// NewHttpResource instantiates a new HttpResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHttpResource(id int32, name string, source string, trafficUsage TrafficUsage, status string, cdnDomain string, presetId int32) *HttpResource {
	this := HttpResource{}
	this.Id = id
	this.Name = name
	this.Source = source
	this.TrafficUsage = trafficUsage
	this.Status = status
	this.CdnDomain = cdnDomain
	this.PresetId = presetId
	return &this
}

// NewHttpResourceWithDefaults instantiates a new HttpResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpResourceWithDefaults() *HttpResource {
	this := HttpResource{}
	return &this
}

// GetId returns the Id field value
func (o *HttpResource) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *HttpResource) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *HttpResource) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *HttpResource) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *HttpResource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *HttpResource) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *HttpResource) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpResource) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *HttpResource) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *HttpResource) SetDescription(v string) {
	o.Description = &v
}

// GetSource returns the Source field value
func (o *HttpResource) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *HttpResource) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *HttpResource) SetSource(v string) {
	o.Source = v
}

// GetTrafficUsage returns the TrafficUsage field value
func (o *HttpResource) GetTrafficUsage() TrafficUsage {
	if o == nil {
		var ret TrafficUsage
		return ret
	}

	return o.TrafficUsage
}

// GetTrafficUsageOk returns a tuple with the TrafficUsage field value
// and a boolean to check if the value has been set.
func (o *HttpResource) GetTrafficUsageOk() (*TrafficUsage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrafficUsage, true
}

// SetTrafficUsage sets field value
func (o *HttpResource) SetTrafficUsage(v TrafficUsage) {
	o.TrafficUsage = v
}

// GetStatus returns the Status field value
func (o *HttpResource) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *HttpResource) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *HttpResource) SetStatus(v string) {
	o.Status = v
}

// GetCdnDomain returns the CdnDomain field value
func (o *HttpResource) GetCdnDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CdnDomain
}

// GetCdnDomainOk returns a tuple with the CdnDomain field value
// and a boolean to check if the value has been set.
func (o *HttpResource) GetCdnDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CdnDomain, true
}

// SetCdnDomain sets field value
func (o *HttpResource) SetCdnDomain(v string) {
	o.CdnDomain = v
}

// GetPresetId returns the PresetId field value
func (o *HttpResource) GetPresetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PresetId
}

// GetPresetIdOk returns a tuple with the PresetId field value
// and a boolean to check if the value has been set.
func (o *HttpResource) GetPresetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PresetId, true
}

// SetPresetId sets field value
func (o *HttpResource) SetPresetId(v int32) {
	o.PresetId = v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *HttpResource) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpResource) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *HttpResource) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *HttpResource) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetAvatarLink returns the AvatarLink field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HttpResource) GetAvatarLink() string {
	if o == nil || IsNil(o.AvatarLink.Get()) {
		var ret string
		return ret
	}
	return *o.AvatarLink.Get()
}

// GetAvatarLinkOk returns a tuple with the AvatarLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HttpResource) GetAvatarLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvatarLink.Get(), o.AvatarLink.IsSet()
}

// HasAvatarLink returns a boolean if a field has been set.
func (o *HttpResource) HasAvatarLink() bool {
	if o != nil && o.AvatarLink.IsSet() {
		return true
	}

	return false
}

// SetAvatarLink gets a reference to the given NullableString and assigns it to the AvatarLink field.
func (o *HttpResource) SetAvatarLink(v string) {
	o.AvatarLink.Set(&v)
}
// SetAvatarLinkNil sets the value for AvatarLink to be an explicit nil
func (o *HttpResource) SetAvatarLinkNil() {
	o.AvatarLink.Set(nil)
}

// UnsetAvatarLink ensures that no value is present for AvatarLink, not even an explicit nil
func (o *HttpResource) UnsetAvatarLink() {
	o.AvatarLink.Unset()
}

// GetStorageId returns the StorageId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HttpResource) GetStorageId() int32 {
	if o == nil || IsNil(o.StorageId.Get()) {
		var ret int32
		return ret
	}
	return *o.StorageId.Get()
}

// GetStorageIdOk returns a tuple with the StorageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HttpResource) GetStorageIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.StorageId.Get(), o.StorageId.IsSet()
}

// HasStorageId returns a boolean if a field has been set.
func (o *HttpResource) HasStorageId() bool {
	if o != nil && o.StorageId.IsSet() {
		return true
	}

	return false
}

// SetStorageId gets a reference to the given NullableInt32 and assigns it to the StorageId field.
func (o *HttpResource) SetStorageId(v int32) {
	o.StorageId.Set(&v)
}
// SetStorageIdNil sets the value for StorageId to be an explicit nil
func (o *HttpResource) SetStorageIdNil() {
	o.StorageId.Set(nil)
}

// UnsetStorageId ensures that no value is present for StorageId, not even an explicit nil
func (o *HttpResource) UnsetStorageId() {
	o.StorageId.Unset()
}

// GetTrafficLimitBytes returns the TrafficLimitBytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HttpResource) GetTrafficLimitBytes() int32 {
	if o == nil || IsNil(o.TrafficLimitBytes.Get()) {
		var ret int32
		return ret
	}
	return *o.TrafficLimitBytes.Get()
}

// GetTrafficLimitBytesOk returns a tuple with the TrafficLimitBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HttpResource) GetTrafficLimitBytesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrafficLimitBytes.Get(), o.TrafficLimitBytes.IsSet()
}

// HasTrafficLimitBytes returns a boolean if a field has been set.
func (o *HttpResource) HasTrafficLimitBytes() bool {
	if o != nil && o.TrafficLimitBytes.IsSet() {
		return true
	}

	return false
}

// SetTrafficLimitBytes gets a reference to the given NullableInt32 and assigns it to the TrafficLimitBytes field.
func (o *HttpResource) SetTrafficLimitBytes(v int32) {
	o.TrafficLimitBytes.Set(&v)
}
// SetTrafficLimitBytesNil sets the value for TrafficLimitBytes to be an explicit nil
func (o *HttpResource) SetTrafficLimitBytesNil() {
	o.TrafficLimitBytes.Set(nil)
}

// UnsetTrafficLimitBytes ensures that no value is present for TrafficLimitBytes, not even an explicit nil
func (o *HttpResource) UnsetTrafficLimitBytes() {
	o.TrafficLimitBytes.Unset()
}

func (o HttpResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HttpResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["source"] = o.Source
	toSerialize["traffic_usage"] = o.TrafficUsage
	toSerialize["status"] = o.Status
	toSerialize["cdn_domain"] = o.CdnDomain
	toSerialize["preset_id"] = o.PresetId
	if !IsNil(o.ProjectId) {
		toSerialize["project_id"] = o.ProjectId
	}
	if o.AvatarLink.IsSet() {
		toSerialize["avatar_link"] = o.AvatarLink.Get()
	}
	if o.StorageId.IsSet() {
		toSerialize["storage_id"] = o.StorageId.Get()
	}
	if o.TrafficLimitBytes.IsSet() {
		toSerialize["traffic_limit_bytes"] = o.TrafficLimitBytes.Get()
	}
	return toSerialize, nil
}

type NullableHttpResource struct {
	value *HttpResource
	isSet bool
}

func (v NullableHttpResource) Get() *HttpResource {
	return v.value
}

func (v *NullableHttpResource) Set(val *HttpResource) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpResource) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpResource(val *HttpResource) *NullableHttpResource {
	return &NullableHttpResource{value: val, isSet: true}
}

func (v NullableHttpResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


