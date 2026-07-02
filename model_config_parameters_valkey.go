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

// checks if the ConfigParametersValkey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigParametersValkey{}

// ConfigParametersValkey Параметры Valkey (`valkey` | `valkey7` | `valkey8_1` | `valkey9_1`)
type ConfigParametersValkey struct {
	// Ограничение буфера вывода для обычных клиентских подключений. Формат: `hard-limit soft-limit soft-seconds` (`valkey` | `valkey7` | `valkey8_1` | `valkey9_1`).
	ClientOutputBufferLimitNormal *string `json:"client-output-buffer-limit normal,omitempty"`
	// Ограничение буфера вывода для клиентов pub/sub. Формат: `hard-limit soft-limit soft-seconds` (`valkey` | `valkey7` | `valkey8_1` | `valkey9_1`).
	ClientOutputBufferLimitPubsub *string `json:"client-output-buffer-limit pubsub,omitempty"`
	// Количество логических баз данных на сервере (`valkey` | `valkey7` | `valkey8_1` | `valkey9_1`).
	Databases *string `json:"databases,omitempty"`
	// Время ожидания в секундах перед закрытием неактивного клиентского соединения. `0` — отключено (`valkey` | `valkey7` | `valkey8_1` | `valkey9_1`).
	Timeout *string `json:"timeout,omitempty"`
	// Политика вытеснения ключей при достижении лимита памяти (`valkey` | `valkey7` | `valkey8_1` | `valkey9_1`).
	MaxmemoryPolicy *string `json:"maxmemory-policy,omitempty"`
	// Минимальное время выполнения команды в микросекундах для записи в журнал медленных команд (`valkey` | `valkey7` | `valkey8_1` | `valkey9_1`).
	SlowlogLogSlowerThan *string `json:"slowlog-log-slower-than,omitempty"`
	// Максимальное количество записей, хранящихся в журнале медленных команд (`valkey` | `valkey7` | `valkey8_1` | `valkey9_1`).
	SlowlogMaxLen *string `json:"slowlog-max-len,omitempty"`
	// Условие создания снимка RDB на диск. Формат: `seconds changes` — сохранение выполняется, если за указанное время было сделано не менее указанного количества изменений (`valkey` | `valkey7` | `valkey8_1` | `valkey9_1`).
	Save *string `json:"save,omitempty"`
	// Включение режима AOF (Append Only File) для персистентного хранения данных (`valkey` | `valkey7` | `valkey8_1` | `valkey9_1`).
	Appendonly *string `json:"appendonly,omitempty"`
	// Режим синхронизации AOF-файла с диском: `always` — при каждой записи, `everysec` — раз в секунду, `no` — управление передаётся ОС (`valkey` | `valkey7` | `valkey8_1` | `valkey9_1`).
	Appendfsync *string `json:"appendfsync,omitempty"`
	// Интервал проверки активности TCP-соединения в секундах. `0` — отключено (`valkey` | `valkey7` | `valkey8_1` | `valkey9_1`).
	TcpKeepalive *string `json:"tcp-keepalive,omitempty"`
}

// NewConfigParametersValkey instantiates a new ConfigParametersValkey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigParametersValkey() *ConfigParametersValkey {
	this := ConfigParametersValkey{}
	return &this
}

// NewConfigParametersValkeyWithDefaults instantiates a new ConfigParametersValkey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigParametersValkeyWithDefaults() *ConfigParametersValkey {
	this := ConfigParametersValkey{}
	return &this
}

// GetClientOutputBufferLimitNormal returns the ClientOutputBufferLimitNormal field value if set, zero value otherwise.
func (o *ConfigParametersValkey) GetClientOutputBufferLimitNormal() string {
	if o == nil || IsNil(o.ClientOutputBufferLimitNormal) {
		var ret string
		return ret
	}
	return *o.ClientOutputBufferLimitNormal
}

// GetClientOutputBufferLimitNormalOk returns a tuple with the ClientOutputBufferLimitNormal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersValkey) GetClientOutputBufferLimitNormalOk() (*string, bool) {
	if o == nil || IsNil(o.ClientOutputBufferLimitNormal) {
		return nil, false
	}
	return o.ClientOutputBufferLimitNormal, true
}

// HasClientOutputBufferLimitNormal returns a boolean if a field has been set.
func (o *ConfigParametersValkey) HasClientOutputBufferLimitNormal() bool {
	if o != nil && !IsNil(o.ClientOutputBufferLimitNormal) {
		return true
	}

	return false
}

// SetClientOutputBufferLimitNormal gets a reference to the given string and assigns it to the ClientOutputBufferLimitNormal field.
func (o *ConfigParametersValkey) SetClientOutputBufferLimitNormal(v string) {
	o.ClientOutputBufferLimitNormal = &v
}

// GetClientOutputBufferLimitPubsub returns the ClientOutputBufferLimitPubsub field value if set, zero value otherwise.
func (o *ConfigParametersValkey) GetClientOutputBufferLimitPubsub() string {
	if o == nil || IsNil(o.ClientOutputBufferLimitPubsub) {
		var ret string
		return ret
	}
	return *o.ClientOutputBufferLimitPubsub
}

// GetClientOutputBufferLimitPubsubOk returns a tuple with the ClientOutputBufferLimitPubsub field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersValkey) GetClientOutputBufferLimitPubsubOk() (*string, bool) {
	if o == nil || IsNil(o.ClientOutputBufferLimitPubsub) {
		return nil, false
	}
	return o.ClientOutputBufferLimitPubsub, true
}

// HasClientOutputBufferLimitPubsub returns a boolean if a field has been set.
func (o *ConfigParametersValkey) HasClientOutputBufferLimitPubsub() bool {
	if o != nil && !IsNil(o.ClientOutputBufferLimitPubsub) {
		return true
	}

	return false
}

// SetClientOutputBufferLimitPubsub gets a reference to the given string and assigns it to the ClientOutputBufferLimitPubsub field.
func (o *ConfigParametersValkey) SetClientOutputBufferLimitPubsub(v string) {
	o.ClientOutputBufferLimitPubsub = &v
}

// GetDatabases returns the Databases field value if set, zero value otherwise.
func (o *ConfigParametersValkey) GetDatabases() string {
	if o == nil || IsNil(o.Databases) {
		var ret string
		return ret
	}
	return *o.Databases
}

// GetDatabasesOk returns a tuple with the Databases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersValkey) GetDatabasesOk() (*string, bool) {
	if o == nil || IsNil(o.Databases) {
		return nil, false
	}
	return o.Databases, true
}

// HasDatabases returns a boolean if a field has been set.
func (o *ConfigParametersValkey) HasDatabases() bool {
	if o != nil && !IsNil(o.Databases) {
		return true
	}

	return false
}

// SetDatabases gets a reference to the given string and assigns it to the Databases field.
func (o *ConfigParametersValkey) SetDatabases(v string) {
	o.Databases = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *ConfigParametersValkey) GetTimeout() string {
	if o == nil || IsNil(o.Timeout) {
		var ret string
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersValkey) GetTimeoutOk() (*string, bool) {
	if o == nil || IsNil(o.Timeout) {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *ConfigParametersValkey) HasTimeout() bool {
	if o != nil && !IsNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given string and assigns it to the Timeout field.
func (o *ConfigParametersValkey) SetTimeout(v string) {
	o.Timeout = &v
}

// GetMaxmemoryPolicy returns the MaxmemoryPolicy field value if set, zero value otherwise.
func (o *ConfigParametersValkey) GetMaxmemoryPolicy() string {
	if o == nil || IsNil(o.MaxmemoryPolicy) {
		var ret string
		return ret
	}
	return *o.MaxmemoryPolicy
}

// GetMaxmemoryPolicyOk returns a tuple with the MaxmemoryPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersValkey) GetMaxmemoryPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.MaxmemoryPolicy) {
		return nil, false
	}
	return o.MaxmemoryPolicy, true
}

// HasMaxmemoryPolicy returns a boolean if a field has been set.
func (o *ConfigParametersValkey) HasMaxmemoryPolicy() bool {
	if o != nil && !IsNil(o.MaxmemoryPolicy) {
		return true
	}

	return false
}

// SetMaxmemoryPolicy gets a reference to the given string and assigns it to the MaxmemoryPolicy field.
func (o *ConfigParametersValkey) SetMaxmemoryPolicy(v string) {
	o.MaxmemoryPolicy = &v
}

// GetSlowlogLogSlowerThan returns the SlowlogLogSlowerThan field value if set, zero value otherwise.
func (o *ConfigParametersValkey) GetSlowlogLogSlowerThan() string {
	if o == nil || IsNil(o.SlowlogLogSlowerThan) {
		var ret string
		return ret
	}
	return *o.SlowlogLogSlowerThan
}

// GetSlowlogLogSlowerThanOk returns a tuple with the SlowlogLogSlowerThan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersValkey) GetSlowlogLogSlowerThanOk() (*string, bool) {
	if o == nil || IsNil(o.SlowlogLogSlowerThan) {
		return nil, false
	}
	return o.SlowlogLogSlowerThan, true
}

// HasSlowlogLogSlowerThan returns a boolean if a field has been set.
func (o *ConfigParametersValkey) HasSlowlogLogSlowerThan() bool {
	if o != nil && !IsNil(o.SlowlogLogSlowerThan) {
		return true
	}

	return false
}

// SetSlowlogLogSlowerThan gets a reference to the given string and assigns it to the SlowlogLogSlowerThan field.
func (o *ConfigParametersValkey) SetSlowlogLogSlowerThan(v string) {
	o.SlowlogLogSlowerThan = &v
}

// GetSlowlogMaxLen returns the SlowlogMaxLen field value if set, zero value otherwise.
func (o *ConfigParametersValkey) GetSlowlogMaxLen() string {
	if o == nil || IsNil(o.SlowlogMaxLen) {
		var ret string
		return ret
	}
	return *o.SlowlogMaxLen
}

// GetSlowlogMaxLenOk returns a tuple with the SlowlogMaxLen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersValkey) GetSlowlogMaxLenOk() (*string, bool) {
	if o == nil || IsNil(o.SlowlogMaxLen) {
		return nil, false
	}
	return o.SlowlogMaxLen, true
}

// HasSlowlogMaxLen returns a boolean if a field has been set.
func (o *ConfigParametersValkey) HasSlowlogMaxLen() bool {
	if o != nil && !IsNil(o.SlowlogMaxLen) {
		return true
	}

	return false
}

// SetSlowlogMaxLen gets a reference to the given string and assigns it to the SlowlogMaxLen field.
func (o *ConfigParametersValkey) SetSlowlogMaxLen(v string) {
	o.SlowlogMaxLen = &v
}

// GetSave returns the Save field value if set, zero value otherwise.
func (o *ConfigParametersValkey) GetSave() string {
	if o == nil || IsNil(o.Save) {
		var ret string
		return ret
	}
	return *o.Save
}

// GetSaveOk returns a tuple with the Save field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersValkey) GetSaveOk() (*string, bool) {
	if o == nil || IsNil(o.Save) {
		return nil, false
	}
	return o.Save, true
}

// HasSave returns a boolean if a field has been set.
func (o *ConfigParametersValkey) HasSave() bool {
	if o != nil && !IsNil(o.Save) {
		return true
	}

	return false
}

// SetSave gets a reference to the given string and assigns it to the Save field.
func (o *ConfigParametersValkey) SetSave(v string) {
	o.Save = &v
}

// GetAppendonly returns the Appendonly field value if set, zero value otherwise.
func (o *ConfigParametersValkey) GetAppendonly() string {
	if o == nil || IsNil(o.Appendonly) {
		var ret string
		return ret
	}
	return *o.Appendonly
}

// GetAppendonlyOk returns a tuple with the Appendonly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersValkey) GetAppendonlyOk() (*string, bool) {
	if o == nil || IsNil(o.Appendonly) {
		return nil, false
	}
	return o.Appendonly, true
}

// HasAppendonly returns a boolean if a field has been set.
func (o *ConfigParametersValkey) HasAppendonly() bool {
	if o != nil && !IsNil(o.Appendonly) {
		return true
	}

	return false
}

// SetAppendonly gets a reference to the given string and assigns it to the Appendonly field.
func (o *ConfigParametersValkey) SetAppendonly(v string) {
	o.Appendonly = &v
}

// GetAppendfsync returns the Appendfsync field value if set, zero value otherwise.
func (o *ConfigParametersValkey) GetAppendfsync() string {
	if o == nil || IsNil(o.Appendfsync) {
		var ret string
		return ret
	}
	return *o.Appendfsync
}

// GetAppendfsyncOk returns a tuple with the Appendfsync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersValkey) GetAppendfsyncOk() (*string, bool) {
	if o == nil || IsNil(o.Appendfsync) {
		return nil, false
	}
	return o.Appendfsync, true
}

// HasAppendfsync returns a boolean if a field has been set.
func (o *ConfigParametersValkey) HasAppendfsync() bool {
	if o != nil && !IsNil(o.Appendfsync) {
		return true
	}

	return false
}

// SetAppendfsync gets a reference to the given string and assigns it to the Appendfsync field.
func (o *ConfigParametersValkey) SetAppendfsync(v string) {
	o.Appendfsync = &v
}

// GetTcpKeepalive returns the TcpKeepalive field value if set, zero value otherwise.
func (o *ConfigParametersValkey) GetTcpKeepalive() string {
	if o == nil || IsNil(o.TcpKeepalive) {
		var ret string
		return ret
	}
	return *o.TcpKeepalive
}

// GetTcpKeepaliveOk returns a tuple with the TcpKeepalive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersValkey) GetTcpKeepaliveOk() (*string, bool) {
	if o == nil || IsNil(o.TcpKeepalive) {
		return nil, false
	}
	return o.TcpKeepalive, true
}

// HasTcpKeepalive returns a boolean if a field has been set.
func (o *ConfigParametersValkey) HasTcpKeepalive() bool {
	if o != nil && !IsNil(o.TcpKeepalive) {
		return true
	}

	return false
}

// SetTcpKeepalive gets a reference to the given string and assigns it to the TcpKeepalive field.
func (o *ConfigParametersValkey) SetTcpKeepalive(v string) {
	o.TcpKeepalive = &v
}

func (o ConfigParametersValkey) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigParametersValkey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientOutputBufferLimitNormal) {
		toSerialize["client-output-buffer-limit normal"] = o.ClientOutputBufferLimitNormal
	}
	if !IsNil(o.ClientOutputBufferLimitPubsub) {
		toSerialize["client-output-buffer-limit pubsub"] = o.ClientOutputBufferLimitPubsub
	}
	if !IsNil(o.Databases) {
		toSerialize["databases"] = o.Databases
	}
	if !IsNil(o.Timeout) {
		toSerialize["timeout"] = o.Timeout
	}
	if !IsNil(o.MaxmemoryPolicy) {
		toSerialize["maxmemory-policy"] = o.MaxmemoryPolicy
	}
	if !IsNil(o.SlowlogLogSlowerThan) {
		toSerialize["slowlog-log-slower-than"] = o.SlowlogLogSlowerThan
	}
	if !IsNil(o.SlowlogMaxLen) {
		toSerialize["slowlog-max-len"] = o.SlowlogMaxLen
	}
	if !IsNil(o.Save) {
		toSerialize["save"] = o.Save
	}
	if !IsNil(o.Appendonly) {
		toSerialize["appendonly"] = o.Appendonly
	}
	if !IsNil(o.Appendfsync) {
		toSerialize["appendfsync"] = o.Appendfsync
	}
	if !IsNil(o.TcpKeepalive) {
		toSerialize["tcp-keepalive"] = o.TcpKeepalive
	}
	return toSerialize, nil
}

type NullableConfigParametersValkey struct {
	value *ConfigParametersValkey
	isSet bool
}

func (v NullableConfigParametersValkey) Get() *ConfigParametersValkey {
	return v.value
}

func (v *NullableConfigParametersValkey) Set(val *ConfigParametersValkey) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigParametersValkey) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigParametersValkey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigParametersValkey(val *ConfigParametersValkey) *NullableConfigParametersValkey {
	return &NullableConfigParametersValkey{value: val, isSet: true}
}

func (v NullableConfigParametersValkey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigParametersValkey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


