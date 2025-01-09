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

// checks if the Bucket type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Bucket{}

// Bucket Хранилище S3
type Bucket struct {
	// ID для каждого экземпляра хранилища. Автоматически генерируется при создании.
	Id float32 `json:"id"`
	// Удобочитаемое имя, установленное для хранилища.
	Name string `json:"name"`
	DiskStats BucketDiskStats `json:"disk_stats"`
	// Тип хранилища.
	Type string `json:"type"`
	// ID тарифа хранилища.
	PresetId NullableFloat32 `json:"preset_id"`
	// Статус хранилища.
	Status string `json:"status"`
	// Количество файлов в хранилище.
	ObjectAmount float32 `json:"object_amount"`
	// Регион хранилища.
	Location string `json:"location"`
	// Адрес хранилища для подключения.
	Hostname string `json:"hostname"`
	// Ключ доступа от хранилища.
	AccessKey string `json:"access_key"`
	// Секретный ключ доступа от хранилища.
	SecretKey string `json:"secret_key"`
}

// NewBucket instantiates a new Bucket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBucket(id float32, name string, diskStats BucketDiskStats, type_ string, presetId NullableFloat32, status string, objectAmount float32, location string, hostname string, accessKey string, secretKey string) *Bucket {
	this := Bucket{}
	this.Id = id
	this.Name = name
	this.DiskStats = diskStats
	this.Type = type_
	this.PresetId = presetId
	this.Status = status
	this.ObjectAmount = objectAmount
	this.Location = location
	this.Hostname = hostname
	this.AccessKey = accessKey
	this.SecretKey = secretKey
	return &this
}

// NewBucketWithDefaults instantiates a new Bucket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBucketWithDefaults() *Bucket {
	this := Bucket{}
	return &this
}

// GetId returns the Id field value
func (o *Bucket) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Bucket) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Bucket) SetId(v float32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Bucket) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Bucket) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Bucket) SetName(v string) {
	o.Name = v
}

// GetDiskStats returns the DiskStats field value
func (o *Bucket) GetDiskStats() BucketDiskStats {
	if o == nil {
		var ret BucketDiskStats
		return ret
	}

	return o.DiskStats
}

// GetDiskStatsOk returns a tuple with the DiskStats field value
// and a boolean to check if the value has been set.
func (o *Bucket) GetDiskStatsOk() (*BucketDiskStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiskStats, true
}

// SetDiskStats sets field value
func (o *Bucket) SetDiskStats(v BucketDiskStats) {
	o.DiskStats = v
}

// GetType returns the Type field value
func (o *Bucket) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Bucket) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Bucket) SetType(v string) {
	o.Type = v
}

// GetPresetId returns the PresetId field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *Bucket) GetPresetId() float32 {
	if o == nil || o.PresetId.Get() == nil {
		var ret float32
		return ret
	}

	return *o.PresetId.Get()
}

// GetPresetIdOk returns a tuple with the PresetId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Bucket) GetPresetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PresetId.Get(), o.PresetId.IsSet()
}

// SetPresetId sets field value
func (o *Bucket) SetPresetId(v float32) {
	o.PresetId.Set(&v)
}

// GetStatus returns the Status field value
func (o *Bucket) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Bucket) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Bucket) SetStatus(v string) {
	o.Status = v
}

// GetObjectAmount returns the ObjectAmount field value
func (o *Bucket) GetObjectAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ObjectAmount
}

// GetObjectAmountOk returns a tuple with the ObjectAmount field value
// and a boolean to check if the value has been set.
func (o *Bucket) GetObjectAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectAmount, true
}

// SetObjectAmount sets field value
func (o *Bucket) SetObjectAmount(v float32) {
	o.ObjectAmount = v
}

// GetLocation returns the Location field value
func (o *Bucket) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *Bucket) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *Bucket) SetLocation(v string) {
	o.Location = v
}

// GetHostname returns the Hostname field value
func (o *Bucket) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *Bucket) GetHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *Bucket) SetHostname(v string) {
	o.Hostname = v
}

// GetAccessKey returns the AccessKey field value
func (o *Bucket) GetAccessKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value
// and a boolean to check if the value has been set.
func (o *Bucket) GetAccessKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessKey, true
}

// SetAccessKey sets field value
func (o *Bucket) SetAccessKey(v string) {
	o.AccessKey = v
}

// GetSecretKey returns the SecretKey field value
func (o *Bucket) GetSecretKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value
// and a boolean to check if the value has been set.
func (o *Bucket) GetSecretKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretKey, true
}

// SetSecretKey sets field value
func (o *Bucket) SetSecretKey(v string) {
	o.SecretKey = v
}

func (o Bucket) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Bucket) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["disk_stats"] = o.DiskStats
	toSerialize["type"] = o.Type
	toSerialize["preset_id"] = o.PresetId.Get()
	toSerialize["status"] = o.Status
	toSerialize["object_amount"] = o.ObjectAmount
	toSerialize["location"] = o.Location
	toSerialize["hostname"] = o.Hostname
	toSerialize["access_key"] = o.AccessKey
	toSerialize["secret_key"] = o.SecretKey
	return toSerialize, nil
}

type NullableBucket struct {
	value *Bucket
	isSet bool
}

func (v NullableBucket) Get() *Bucket {
	return v.value
}

func (v *NullableBucket) Set(val *Bucket) {
	v.value = val
	v.isSet = true
}

func (v NullableBucket) IsSet() bool {
	return v.isSet
}

func (v *NullableBucket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBucket(val *Bucket) *NullableBucket {
	return &NullableBucket{value: val, isSet: true}
}

func (v NullableBucket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBucket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


