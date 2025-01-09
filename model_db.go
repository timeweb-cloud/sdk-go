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

// checks if the Db type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Db{}

// Db База данных
type Db struct {
	// ID для каждого экземпляра базы данных. Автоматически генерируется при создании.
	Id float32 `json:"id"`
	// Значение времени, указанное в комбинированном формате даты и времени ISO8601, которое представляет, когда была создана база данных.
	CreatedAt string `json:"created_at"`
	// ID пользователя.
	AccountId string `json:"account_id"`
	// Логин для подключения к базе данных.
	Login string `json:"login"`
	// Локация сервера.
	Location *string `json:"location,omitempty"`
	// Пароль для подключения к базе данных.
	Password string `json:"password"`
	// Название базы данных.
	Name string `json:"name"`
	// Хост.
	Host NullableString `json:"host"`
	Type DbType `json:"type"`
	// Тип хеширования базы данных (mysql5 | mysql | postgres).
	HashType NullableString `json:"hash_type"`
	// Порт
	Port int32 `json:"port"`
	// IP-адрес сетевого интерфейса IPv4.
	Ip NullableString `json:"ip"`
	// IP-адрес локального сетевого интерфейса IPv4.
	LocalIp NullableString `json:"local_ip"`
	// Текущий статус базы данных.
	Status string `json:"status"`
	// ID тарифа.
	PresetId int32 `json:"preset_id"`
	DiskStats NullableDbDiskStats `json:"disk_stats"`
	ConfigParameters ConfigParameters `json:"config_parameters"`
	// Это логическое значение, которое показывает, доступна ли база данных только по локальному IP адресу.
	IsOnlyLocalIpAccess bool `json:"is_only_local_ip_access"`
	AvailabilityZone AvailabilityZone `json:"availability_zone"`
}

// NewDb instantiates a new Db object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDb(id float32, createdAt string, accountId string, login string, password string, name string, host NullableString, type_ DbType, hashType NullableString, port int32, ip NullableString, localIp NullableString, status string, presetId int32, diskStats NullableDbDiskStats, configParameters ConfigParameters, isOnlyLocalIpAccess bool, availabilityZone AvailabilityZone) *Db {
	this := Db{}
	this.Id = id
	this.CreatedAt = createdAt
	this.AccountId = accountId
	this.Login = login
	this.Password = password
	this.Name = name
	this.Host = host
	this.Type = type_
	this.HashType = hashType
	this.Port = port
	this.Ip = ip
	this.LocalIp = localIp
	this.Status = status
	this.PresetId = presetId
	this.DiskStats = diskStats
	this.ConfigParameters = configParameters
	this.IsOnlyLocalIpAccess = isOnlyLocalIpAccess
	this.AvailabilityZone = availabilityZone
	return &this
}

// NewDbWithDefaults instantiates a new Db object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbWithDefaults() *Db {
	this := Db{}
	return &this
}

// GetId returns the Id field value
func (o *Db) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Db) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Db) SetId(v float32) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Db) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Db) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Db) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetAccountId returns the AccountId field value
func (o *Db) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *Db) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *Db) SetAccountId(v string) {
	o.AccountId = v
}

// GetLogin returns the Login field value
func (o *Db) GetLogin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Login
}

// GetLoginOk returns a tuple with the Login field value
// and a boolean to check if the value has been set.
func (o *Db) GetLoginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Login, true
}

// SetLogin sets field value
func (o *Db) SetLogin(v string) {
	o.Login = v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *Db) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Db) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *Db) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *Db) SetLocation(v string) {
	o.Location = &v
}

// GetPassword returns the Password field value
func (o *Db) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *Db) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *Db) SetPassword(v string) {
	o.Password = v
}

// GetName returns the Name field value
func (o *Db) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Db) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Db) SetName(v string) {
	o.Name = v
}

// GetHost returns the Host field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Db) GetHost() string {
	if o == nil || o.Host.Get() == nil {
		var ret string
		return ret
	}

	return *o.Host.Get()
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Db) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Host.Get(), o.Host.IsSet()
}

// SetHost sets field value
func (o *Db) SetHost(v string) {
	o.Host.Set(&v)
}

// GetType returns the Type field value
func (o *Db) GetType() DbType {
	if o == nil {
		var ret DbType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Db) GetTypeOk() (*DbType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Db) SetType(v DbType) {
	o.Type = v
}

// GetHashType returns the HashType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Db) GetHashType() string {
	if o == nil || o.HashType.Get() == nil {
		var ret string
		return ret
	}

	return *o.HashType.Get()
}

// GetHashTypeOk returns a tuple with the HashType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Db) GetHashTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HashType.Get(), o.HashType.IsSet()
}

// SetHashType sets field value
func (o *Db) SetHashType(v string) {
	o.HashType.Set(&v)
}

// GetPort returns the Port field value
func (o *Db) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *Db) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *Db) SetPort(v int32) {
	o.Port = v
}

// GetIp returns the Ip field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Db) GetIp() string {
	if o == nil || o.Ip.Get() == nil {
		var ret string
		return ret
	}

	return *o.Ip.Get()
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Db) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ip.Get(), o.Ip.IsSet()
}

// SetIp sets field value
func (o *Db) SetIp(v string) {
	o.Ip.Set(&v)
}

// GetLocalIp returns the LocalIp field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Db) GetLocalIp() string {
	if o == nil || o.LocalIp.Get() == nil {
		var ret string
		return ret
	}

	return *o.LocalIp.Get()
}

// GetLocalIpOk returns a tuple with the LocalIp field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Db) GetLocalIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LocalIp.Get(), o.LocalIp.IsSet()
}

// SetLocalIp sets field value
func (o *Db) SetLocalIp(v string) {
	o.LocalIp.Set(&v)
}

// GetStatus returns the Status field value
func (o *Db) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Db) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Db) SetStatus(v string) {
	o.Status = v
}

// GetPresetId returns the PresetId field value
func (o *Db) GetPresetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PresetId
}

// GetPresetIdOk returns a tuple with the PresetId field value
// and a boolean to check if the value has been set.
func (o *Db) GetPresetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PresetId, true
}

// SetPresetId sets field value
func (o *Db) SetPresetId(v int32) {
	o.PresetId = v
}

// GetDiskStats returns the DiskStats field value
// If the value is explicit nil, the zero value for DbDiskStats will be returned
func (o *Db) GetDiskStats() DbDiskStats {
	if o == nil || o.DiskStats.Get() == nil {
		var ret DbDiskStats
		return ret
	}

	return *o.DiskStats.Get()
}

// GetDiskStatsOk returns a tuple with the DiskStats field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Db) GetDiskStatsOk() (*DbDiskStats, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiskStats.Get(), o.DiskStats.IsSet()
}

// SetDiskStats sets field value
func (o *Db) SetDiskStats(v DbDiskStats) {
	o.DiskStats.Set(&v)
}

// GetConfigParameters returns the ConfigParameters field value
func (o *Db) GetConfigParameters() ConfigParameters {
	if o == nil {
		var ret ConfigParameters
		return ret
	}

	return o.ConfigParameters
}

// GetConfigParametersOk returns a tuple with the ConfigParameters field value
// and a boolean to check if the value has been set.
func (o *Db) GetConfigParametersOk() (*ConfigParameters, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigParameters, true
}

// SetConfigParameters sets field value
func (o *Db) SetConfigParameters(v ConfigParameters) {
	o.ConfigParameters = v
}

// GetIsOnlyLocalIpAccess returns the IsOnlyLocalIpAccess field value
func (o *Db) GetIsOnlyLocalIpAccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsOnlyLocalIpAccess
}

// GetIsOnlyLocalIpAccessOk returns a tuple with the IsOnlyLocalIpAccess field value
// and a boolean to check if the value has been set.
func (o *Db) GetIsOnlyLocalIpAccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsOnlyLocalIpAccess, true
}

// SetIsOnlyLocalIpAccess sets field value
func (o *Db) SetIsOnlyLocalIpAccess(v bool) {
	o.IsOnlyLocalIpAccess = v
}

// GetAvailabilityZone returns the AvailabilityZone field value
func (o *Db) GetAvailabilityZone() AvailabilityZone {
	if o == nil {
		var ret AvailabilityZone
		return ret
	}

	return o.AvailabilityZone
}

// GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field value
// and a boolean to check if the value has been set.
func (o *Db) GetAvailabilityZoneOk() (*AvailabilityZone, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvailabilityZone, true
}

// SetAvailabilityZone sets field value
func (o *Db) SetAvailabilityZone(v AvailabilityZone) {
	o.AvailabilityZone = v
}

func (o Db) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Db) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["account_id"] = o.AccountId
	toSerialize["login"] = o.Login
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	toSerialize["password"] = o.Password
	toSerialize["name"] = o.Name
	toSerialize["host"] = o.Host.Get()
	toSerialize["type"] = o.Type
	toSerialize["hash_type"] = o.HashType.Get()
	toSerialize["port"] = o.Port
	toSerialize["ip"] = o.Ip.Get()
	toSerialize["local_ip"] = o.LocalIp.Get()
	toSerialize["status"] = o.Status
	toSerialize["preset_id"] = o.PresetId
	toSerialize["disk_stats"] = o.DiskStats.Get()
	toSerialize["config_parameters"] = o.ConfigParameters
	toSerialize["is_only_local_ip_access"] = o.IsOnlyLocalIpAccess
	toSerialize["availability_zone"] = o.AvailabilityZone
	return toSerialize, nil
}

type NullableDb struct {
	value *Db
	isSet bool
}

func (v NullableDb) Get() *Db {
	return v.value
}

func (v *NullableDb) Set(val *Db) {
	v.value = val
	v.isSet = true
}

func (v NullableDb) IsSet() bool {
	return v.isSet
}

func (v *NullableDb) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDb(val *Db) *NullableDb {
	return &NullableDb{value: val, isSet: true}
}

func (v NullableDb) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDb) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


