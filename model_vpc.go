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

// checks if the Vpc type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Vpc{}

// Vpc struct for Vpc
type Vpc struct {
	// ID сети.
	Id string `json:"id"`
	// Имя сети.
	Name string `json:"name"`
	// Маска подсети.
	SubnetV4 string `json:"subnet_v4"`
	// Локация сети.
	Location string `json:"location"`
	// Дата создания сети.
	CreatedAt time.Time `json:"created_at"`
	// Описание.
	Description string `json:"description"`
	AvailabilityZone AvailabilityZone `json:"availability_zone"`
	// Публичный IP-адрес сети.
	PublicIp NullableString `json:"public_ip"`
	// Тип сети.
	Type string `json:"type"`
	// Занятые адреса в сети
	BusyAddress []string `json:"busy_address"`
}

// NewVpc instantiates a new Vpc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVpc(id string, name string, subnetV4 string, location string, createdAt time.Time, description string, availabilityZone AvailabilityZone, publicIp NullableString, type_ string, busyAddress []string) *Vpc {
	this := Vpc{}
	this.Id = id
	this.Name = name
	this.SubnetV4 = subnetV4
	this.Location = location
	this.CreatedAt = createdAt
	this.Description = description
	this.AvailabilityZone = availabilityZone
	this.PublicIp = publicIp
	this.Type = type_
	this.BusyAddress = busyAddress
	return &this
}

// NewVpcWithDefaults instantiates a new Vpc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVpcWithDefaults() *Vpc {
	this := Vpc{}
	return &this
}

// GetId returns the Id field value
func (o *Vpc) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Vpc) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Vpc) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Vpc) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Vpc) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Vpc) SetName(v string) {
	o.Name = v
}

// GetSubnetV4 returns the SubnetV4 field value
func (o *Vpc) GetSubnetV4() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubnetV4
}

// GetSubnetV4Ok returns a tuple with the SubnetV4 field value
// and a boolean to check if the value has been set.
func (o *Vpc) GetSubnetV4Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubnetV4, true
}

// SetSubnetV4 sets field value
func (o *Vpc) SetSubnetV4(v string) {
	o.SubnetV4 = v
}

// GetLocation returns the Location field value
func (o *Vpc) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *Vpc) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *Vpc) SetLocation(v string) {
	o.Location = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Vpc) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Vpc) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Vpc) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDescription returns the Description field value
func (o *Vpc) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Vpc) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Vpc) SetDescription(v string) {
	o.Description = v
}

// GetAvailabilityZone returns the AvailabilityZone field value
func (o *Vpc) GetAvailabilityZone() AvailabilityZone {
	if o == nil {
		var ret AvailabilityZone
		return ret
	}

	return o.AvailabilityZone
}

// GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field value
// and a boolean to check if the value has been set.
func (o *Vpc) GetAvailabilityZoneOk() (*AvailabilityZone, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvailabilityZone, true
}

// SetAvailabilityZone sets field value
func (o *Vpc) SetAvailabilityZone(v AvailabilityZone) {
	o.AvailabilityZone = v
}

// GetPublicIp returns the PublicIp field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Vpc) GetPublicIp() string {
	if o == nil || o.PublicIp.Get() == nil {
		var ret string
		return ret
	}

	return *o.PublicIp.Get()
}

// GetPublicIpOk returns a tuple with the PublicIp field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Vpc) GetPublicIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PublicIp.Get(), o.PublicIp.IsSet()
}

// SetPublicIp sets field value
func (o *Vpc) SetPublicIp(v string) {
	o.PublicIp.Set(&v)
}

// GetType returns the Type field value
func (o *Vpc) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Vpc) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Vpc) SetType(v string) {
	o.Type = v
}

// GetBusyAddress returns the BusyAddress field value
func (o *Vpc) GetBusyAddress() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.BusyAddress
}

// GetBusyAddressOk returns a tuple with the BusyAddress field value
// and a boolean to check if the value has been set.
func (o *Vpc) GetBusyAddressOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BusyAddress, true
}

// SetBusyAddress sets field value
func (o *Vpc) SetBusyAddress(v []string) {
	o.BusyAddress = v
}

func (o Vpc) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Vpc) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["subnet_v4"] = o.SubnetV4
	toSerialize["location"] = o.Location
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["description"] = o.Description
	toSerialize["availability_zone"] = o.AvailabilityZone
	toSerialize["public_ip"] = o.PublicIp.Get()
	toSerialize["type"] = o.Type
	toSerialize["busy_address"] = o.BusyAddress
	return toSerialize, nil
}

type NullableVpc struct {
	value *Vpc
	isSet bool
}

func (v NullableVpc) Get() *Vpc {
	return v.value
}

func (v *NullableVpc) Set(val *Vpc) {
	v.value = val
	v.isSet = true
}

func (v NullableVpc) IsSet() bool {
	return v.isSet
}

func (v *NullableVpc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVpc(val *Vpc) *NullableVpc {
	return &NullableVpc{value: val, isSet: true}
}

func (v NullableVpc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVpc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


