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

// checks if the CreateServerNetwork type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateServerNetwork{}

// CreateServerNetwork Параметры конфигурации приватной сети сервера
type CreateServerNetwork struct {
	// ID сети
	Id *string `json:"id,omitempty"`
	// Публичный IP
	FloatingIp *string `json:"floating_ip,omitempty"`
	// Приватный IP
	LocalIp *string `json:"local_ip,omitempty"`
	// Приватный IP
	// Deprecated
	Ip *string `json:"ip,omitempty"`
	// Массив ID сетевых дисков
	NetworkDriveIds []string `json:"network_drive_ids,omitempty"`
}

// NewCreateServerNetwork instantiates a new CreateServerNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateServerNetwork() *CreateServerNetwork {
	this := CreateServerNetwork{}
	return &this
}

// NewCreateServerNetworkWithDefaults instantiates a new CreateServerNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateServerNetworkWithDefaults() *CreateServerNetwork {
	this := CreateServerNetwork{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateServerNetwork) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateServerNetwork) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateServerNetwork) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateServerNetwork) SetId(v string) {
	o.Id = &v
}

// GetFloatingIp returns the FloatingIp field value if set, zero value otherwise.
func (o *CreateServerNetwork) GetFloatingIp() string {
	if o == nil || IsNil(o.FloatingIp) {
		var ret string
		return ret
	}
	return *o.FloatingIp
}

// GetFloatingIpOk returns a tuple with the FloatingIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateServerNetwork) GetFloatingIpOk() (*string, bool) {
	if o == nil || IsNil(o.FloatingIp) {
		return nil, false
	}
	return o.FloatingIp, true
}

// HasFloatingIp returns a boolean if a field has been set.
func (o *CreateServerNetwork) HasFloatingIp() bool {
	if o != nil && !IsNil(o.FloatingIp) {
		return true
	}

	return false
}

// SetFloatingIp gets a reference to the given string and assigns it to the FloatingIp field.
func (o *CreateServerNetwork) SetFloatingIp(v string) {
	o.FloatingIp = &v
}

// GetLocalIp returns the LocalIp field value if set, zero value otherwise.
func (o *CreateServerNetwork) GetLocalIp() string {
	if o == nil || IsNil(o.LocalIp) {
		var ret string
		return ret
	}
	return *o.LocalIp
}

// GetLocalIpOk returns a tuple with the LocalIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateServerNetwork) GetLocalIpOk() (*string, bool) {
	if o == nil || IsNil(o.LocalIp) {
		return nil, false
	}
	return o.LocalIp, true
}

// HasLocalIp returns a boolean if a field has been set.
func (o *CreateServerNetwork) HasLocalIp() bool {
	if o != nil && !IsNil(o.LocalIp) {
		return true
	}

	return false
}

// SetLocalIp gets a reference to the given string and assigns it to the LocalIp field.
func (o *CreateServerNetwork) SetLocalIp(v string) {
	o.LocalIp = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
// Deprecated
func (o *CreateServerNetwork) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CreateServerNetwork) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *CreateServerNetwork) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
// Deprecated
func (o *CreateServerNetwork) SetIp(v string) {
	o.Ip = &v
}

// GetNetworkDriveIds returns the NetworkDriveIds field value if set, zero value otherwise.
func (o *CreateServerNetwork) GetNetworkDriveIds() []string {
	if o == nil || IsNil(o.NetworkDriveIds) {
		var ret []string
		return ret
	}
	return o.NetworkDriveIds
}

// GetNetworkDriveIdsOk returns a tuple with the NetworkDriveIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateServerNetwork) GetNetworkDriveIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.NetworkDriveIds) {
		return nil, false
	}
	return o.NetworkDriveIds, true
}

// HasNetworkDriveIds returns a boolean if a field has been set.
func (o *CreateServerNetwork) HasNetworkDriveIds() bool {
	if o != nil && !IsNil(o.NetworkDriveIds) {
		return true
	}

	return false
}

// SetNetworkDriveIds gets a reference to the given []string and assigns it to the NetworkDriveIds field.
func (o *CreateServerNetwork) SetNetworkDriveIds(v []string) {
	o.NetworkDriveIds = v
}

func (o CreateServerNetwork) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateServerNetwork) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.FloatingIp) {
		toSerialize["floating_ip"] = o.FloatingIp
	}
	if !IsNil(o.LocalIp) {
		toSerialize["local_ip"] = o.LocalIp
	}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.NetworkDriveIds) {
		toSerialize["network_drive_ids"] = o.NetworkDriveIds
	}
	return toSerialize, nil
}

type NullableCreateServerNetwork struct {
	value *CreateServerNetwork
	isSet bool
}

func (v NullableCreateServerNetwork) Get() *CreateServerNetwork {
	return v.value
}

func (v *NullableCreateServerNetwork) Set(val *CreateServerNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateServerNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateServerNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateServerNetwork(val *CreateServerNetwork) *NullableCreateServerNetwork {
	return &NullableCreateServerNetwork{value: val, isSet: true}
}

func (v NullableCreateServerNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateServerNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


