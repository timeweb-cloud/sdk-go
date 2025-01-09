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

// checks if the VdsNetworksInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VdsNetworksInner{}

// VdsNetworksInner struct for VdsNetworksInner
type VdsNetworksInner struct {
	// ID сети. Есть только у приватных сетей.
	Id *string `json:"id,omitempty"`
	// Тип сети.
	Type string `json:"type"`
	// Тип преобразования сетевых адресов.
	NatMode *string `json:"nat_mode,omitempty"`
	// Пропускная способность сети.
	Bandwidth NullableFloat32 `json:"bandwidth,omitempty"`
	// Список IP-адресов сети.
	Ips []VdsNetworksInnerIpsInner `json:"ips"`
	// Это логическое значение, которое показывает, подключена ли DDoS-защита. Только для публичных сетей.
	IsDdosGuard *bool `json:"is_ddos_guard,omitempty"`
	// Это логическое значение, которое показывает, примонтирован ли образ к серверу.
	IsImageMounted *bool `json:"is_image_mounted,omitempty"`
	// Список заблокированных портов на сервере.
	BlockedPorts []int32 `json:"blocked_ports,omitempty"`
}

// NewVdsNetworksInner instantiates a new VdsNetworksInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVdsNetworksInner(type_ string, ips []VdsNetworksInnerIpsInner) *VdsNetworksInner {
	this := VdsNetworksInner{}
	this.Type = type_
	this.Ips = ips
	return &this
}

// NewVdsNetworksInnerWithDefaults instantiates a new VdsNetworksInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVdsNetworksInnerWithDefaults() *VdsNetworksInner {
	this := VdsNetworksInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VdsNetworksInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VdsNetworksInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VdsNetworksInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VdsNetworksInner) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value
func (o *VdsNetworksInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VdsNetworksInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *VdsNetworksInner) SetType(v string) {
	o.Type = v
}

// GetNatMode returns the NatMode field value if set, zero value otherwise.
func (o *VdsNetworksInner) GetNatMode() string {
	if o == nil || IsNil(o.NatMode) {
		var ret string
		return ret
	}
	return *o.NatMode
}

// GetNatModeOk returns a tuple with the NatMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VdsNetworksInner) GetNatModeOk() (*string, bool) {
	if o == nil || IsNil(o.NatMode) {
		return nil, false
	}
	return o.NatMode, true
}

// HasNatMode returns a boolean if a field has been set.
func (o *VdsNetworksInner) HasNatMode() bool {
	if o != nil && !IsNil(o.NatMode) {
		return true
	}

	return false
}

// SetNatMode gets a reference to the given string and assigns it to the NatMode field.
func (o *VdsNetworksInner) SetNatMode(v string) {
	o.NatMode = &v
}

// GetBandwidth returns the Bandwidth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VdsNetworksInner) GetBandwidth() float32 {
	if o == nil || IsNil(o.Bandwidth.Get()) {
		var ret float32
		return ret
	}
	return *o.Bandwidth.Get()
}

// GetBandwidthOk returns a tuple with the Bandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VdsNetworksInner) GetBandwidthOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bandwidth.Get(), o.Bandwidth.IsSet()
}

// HasBandwidth returns a boolean if a field has been set.
func (o *VdsNetworksInner) HasBandwidth() bool {
	if o != nil && o.Bandwidth.IsSet() {
		return true
	}

	return false
}

// SetBandwidth gets a reference to the given NullableFloat32 and assigns it to the Bandwidth field.
func (o *VdsNetworksInner) SetBandwidth(v float32) {
	o.Bandwidth.Set(&v)
}
// SetBandwidthNil sets the value for Bandwidth to be an explicit nil
func (o *VdsNetworksInner) SetBandwidthNil() {
	o.Bandwidth.Set(nil)
}

// UnsetBandwidth ensures that no value is present for Bandwidth, not even an explicit nil
func (o *VdsNetworksInner) UnsetBandwidth() {
	o.Bandwidth.Unset()
}

// GetIps returns the Ips field value
// If the value is explicit nil, the zero value for []VdsNetworksInnerIpsInner will be returned
func (o *VdsNetworksInner) GetIps() []VdsNetworksInnerIpsInner {
	if o == nil {
		var ret []VdsNetworksInnerIpsInner
		return ret
	}

	return o.Ips
}

// GetIpsOk returns a tuple with the Ips field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VdsNetworksInner) GetIpsOk() ([]VdsNetworksInnerIpsInner, bool) {
	if o == nil || IsNil(o.Ips) {
		return nil, false
	}
	return o.Ips, true
}

// SetIps sets field value
func (o *VdsNetworksInner) SetIps(v []VdsNetworksInnerIpsInner) {
	o.Ips = v
}

// GetIsDdosGuard returns the IsDdosGuard field value if set, zero value otherwise.
func (o *VdsNetworksInner) GetIsDdosGuard() bool {
	if o == nil || IsNil(o.IsDdosGuard) {
		var ret bool
		return ret
	}
	return *o.IsDdosGuard
}

// GetIsDdosGuardOk returns a tuple with the IsDdosGuard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VdsNetworksInner) GetIsDdosGuardOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDdosGuard) {
		return nil, false
	}
	return o.IsDdosGuard, true
}

// HasIsDdosGuard returns a boolean if a field has been set.
func (o *VdsNetworksInner) HasIsDdosGuard() bool {
	if o != nil && !IsNil(o.IsDdosGuard) {
		return true
	}

	return false
}

// SetIsDdosGuard gets a reference to the given bool and assigns it to the IsDdosGuard field.
func (o *VdsNetworksInner) SetIsDdosGuard(v bool) {
	o.IsDdosGuard = &v
}

// GetIsImageMounted returns the IsImageMounted field value if set, zero value otherwise.
func (o *VdsNetworksInner) GetIsImageMounted() bool {
	if o == nil || IsNil(o.IsImageMounted) {
		var ret bool
		return ret
	}
	return *o.IsImageMounted
}

// GetIsImageMountedOk returns a tuple with the IsImageMounted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VdsNetworksInner) GetIsImageMountedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsImageMounted) {
		return nil, false
	}
	return o.IsImageMounted, true
}

// HasIsImageMounted returns a boolean if a field has been set.
func (o *VdsNetworksInner) HasIsImageMounted() bool {
	if o != nil && !IsNil(o.IsImageMounted) {
		return true
	}

	return false
}

// SetIsImageMounted gets a reference to the given bool and assigns it to the IsImageMounted field.
func (o *VdsNetworksInner) SetIsImageMounted(v bool) {
	o.IsImageMounted = &v
}

// GetBlockedPorts returns the BlockedPorts field value if set, zero value otherwise.
func (o *VdsNetworksInner) GetBlockedPorts() []int32 {
	if o == nil || IsNil(o.BlockedPorts) {
		var ret []int32
		return ret
	}
	return o.BlockedPorts
}

// GetBlockedPortsOk returns a tuple with the BlockedPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VdsNetworksInner) GetBlockedPortsOk() ([]int32, bool) {
	if o == nil || IsNil(o.BlockedPorts) {
		return nil, false
	}
	return o.BlockedPorts, true
}

// HasBlockedPorts returns a boolean if a field has been set.
func (o *VdsNetworksInner) HasBlockedPorts() bool {
	if o != nil && !IsNil(o.BlockedPorts) {
		return true
	}

	return false
}

// SetBlockedPorts gets a reference to the given []int32 and assigns it to the BlockedPorts field.
func (o *VdsNetworksInner) SetBlockedPorts(v []int32) {
	o.BlockedPorts = v
}

func (o VdsNetworksInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VdsNetworksInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.NatMode) {
		toSerialize["nat_mode"] = o.NatMode
	}
	if o.Bandwidth.IsSet() {
		toSerialize["bandwidth"] = o.Bandwidth.Get()
	}
	if o.Ips != nil {
		toSerialize["ips"] = o.Ips
	}
	if !IsNil(o.IsDdosGuard) {
		toSerialize["is_ddos_guard"] = o.IsDdosGuard
	}
	if !IsNil(o.IsImageMounted) {
		toSerialize["is_image_mounted"] = o.IsImageMounted
	}
	if !IsNil(o.BlockedPorts) {
		toSerialize["blocked_ports"] = o.BlockedPorts
	}
	return toSerialize, nil
}

type NullableVdsNetworksInner struct {
	value *VdsNetworksInner
	isSet bool
}

func (v NullableVdsNetworksInner) Get() *VdsNetworksInner {
	return v.value
}

func (v *NullableVdsNetworksInner) Set(val *VdsNetworksInner) {
	v.value = val
	v.isSet = true
}

func (v NullableVdsNetworksInner) IsSet() bool {
	return v.isSet
}

func (v *NullableVdsNetworksInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVdsNetworksInner(val *VdsNetworksInner) *NullableVdsNetworksInner {
	return &NullableVdsNetworksInner{value: val, isSet: true}
}

func (v NullableVdsNetworksInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVdsNetworksInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


