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

// checks if the ConfigParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigParameters{}

// ConfigParameters struct for ConfigParameters
type ConfigParameters struct {
	// Интервал между значениями столбцов с атрибутом `AUTO_INCREMENT` (`mysql5` | `mysql`).
	AutoIncrementIncrement *string `json:"auto_increment_increment,omitempty"`
	// Начальное значение для столбцов с атрибутом `AUTO_INCREMENT` (`mysql5` | `mysql`).
	AutoIncrementOffset *string `json:"auto_increment_offset,omitempty"`
	// Количество операций ввода-вывода в секунду `IOPS` (`mysql5` | `mysql`).
	InnodbIoCapacity *string `json:"innodb_io_capacity,omitempty"`
	// Количество потоков ввода-вывода, используемых для операций очистки (`mysql5` | `mysql`).
	InnodbPurgeThreads *string `json:"innodb_purge_threads,omitempty"`
	// Количество потоков ввода-вывода, используемых для операций чтения (`mysql5` | `mysql`).
	InnodbReadIoThreads *string `json:"innodb_read_io_threads,omitempty"`
	// Максимальное число потоков, которые могут исполняться (`mysql5` | `mysql`).
	InnodbThreadConcurrency *string `json:"innodb_thread_concurrency,omitempty"`
	// Количество потоков ввода-вывода, используемых для операций записи (`mysql5` | `mysql`).
	InnodbWriteIoThreads *string `json:"innodb_write_io_threads,omitempty"`
	// Минимальный размер буфера (`mysql5` | `mysql`).
	JoinBufferSize *string `json:"join_buffer_size,omitempty"`
	// Максимальный размер одного пакета, строки или параметра, отправляемого функцией `mysql_stmt_send_long_data()` (`mysql5` | `mysql`).
	MaxAllowedPacket *string `json:"max_allowed_packet,omitempty"`
	// Максимальный размер пользовательских MEMORY-таблиц (`mysql5` | `mysql`).
	MaxHeapTableSize *string `json:"max_heap_table_size,omitempty"`
	// Доля измененных или удаленных записей в таблице, при которой процесс автоочистки выполнит команду `ANALYZE` (`postgres` | `postgres14`| `postgres15`).
	AutovacuumAnalyzeScaleFactor *string `json:"autovacuum_analyze_scale_factor,omitempty"`
	// Задержка между запусками процесса фоновой записи (`postgres` | `postgres14`| `postgres15`).
	BgwriterDelay *string `json:"bgwriter_delay,omitempty"`
	// Максимальное число элементов буферного кеша (`postgres` | `postgres14`| `postgres15`).
	BgwriterLruMaxpages *string `json:"bgwriter_lru_maxpages,omitempty"`
	// Время ожидания, по истечении которого будет выполняться проверка состояния перекрестной блокировки (`postgres` | `postgres14`| `postgres15`).
	DeadlockTimeout *string `json:"deadlock_timeout,omitempty"`
	// Максимальный размер очереди записей индекса `GIN` (`postgres` | `postgres14`| `postgres15`).
	GinPendingListLimit *string `json:"gin_pending_list_limit,omitempty"`
	// Время простоя открытой транзакции, при превышении которого будет завершена сессия с этой транзакцией (`postgres` | `postgres14`| `postgres15`).
	IdleInTransactionSessionTimeout *string `json:"idle_in_transaction_session_timeout,omitempty"`
	// Время простоя не открытой транзакции, при превышении которого будет завершена сессия с этой транзакцией (`postgres` | `postgres14`| `postgres15`).
	IdleSessionTimeout *string `json:"idle_session_timeout,omitempty"`
	// Значение количества элементов в списке `FROM` при превышении которого, планировщик будет переносить в список явные инструкции `JOIN` (`postgres` | `postgres14`| `postgres15`).
	JoinCollapseLimit *string `json:"join_collapse_limit,omitempty"`
	// Время ожидания освобождения блокировки (`postgres` | `postgres14`| `postgres15`).
	LockTimeout *string `json:"lock_timeout,omitempty"`
	// Максимальное число транзакций, которые могут одновременно находиться в подготовленном состоянии (`postgres` | `postgres14`| `postgres15`).
	MaxPreparedTransactions *string `json:"max_prepared_transactions,omitempty"`
	// Допустимое количество соединений (`postgres` | `postgres14`| `postgres15` | `mysql`).
	MaxConnections *string `json:"max_connections,omitempty"`
	// Устанавливает количество буферов общей памяти, используемых сервером (`postgres` | `postgres14`| `postgres15`).
	SharedBuffers *string `json:"shared_buffers,omitempty"`
	// Устанавливает количество буферов дисковых страниц в общей памяти для WAL (`postgres` | `postgres14`| `postgres15`).
	WalBuffers *string `json:"wal_buffers,omitempty"`
	// Устанавливает максимальное количество временных буферов, используемых каждой сессией (`postgres` | `postgres14`| `postgres15`).
	TempBuffers *string `json:"temp_buffers,omitempty"`
	// Устанавливает максимальное количество памяти, используемое для рабочих пространств запросов (`postgres` | `postgres14`| `postgres15`).
	WorkMem *string `json:"work_mem,omitempty"`
	// Устанавливает режим SQL. Можно задать несколько режимов, разделяя их запятой. (`mysql`).
	SqlMode *string `json:"sql_mode,omitempty"`
	// Параметр включает или отключает работу MySQL Query Cache (`mysql`).
	QueryCacheType *string `json:"query_cache_type,omitempty"`
	// Размер в байтах, доступный для кэша запросов (`mysql`).
	QueryCacheSize *string `json:"query_cache_size,omitempty"`
}

// NewConfigParameters instantiates a new ConfigParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigParameters() *ConfigParameters {
	this := ConfigParameters{}
	return &this
}

// NewConfigParametersWithDefaults instantiates a new ConfigParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigParametersWithDefaults() *ConfigParameters {
	this := ConfigParameters{}
	return &this
}

// GetAutoIncrementIncrement returns the AutoIncrementIncrement field value if set, zero value otherwise.
func (o *ConfigParameters) GetAutoIncrementIncrement() string {
	if o == nil || IsNil(o.AutoIncrementIncrement) {
		var ret string
		return ret
	}
	return *o.AutoIncrementIncrement
}

// GetAutoIncrementIncrementOk returns a tuple with the AutoIncrementIncrement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParameters) GetAutoIncrementIncrementOk() (*string, bool) {
	if o == nil || IsNil(o.AutoIncrementIncrement) {
		return nil, false
	}
	return o.AutoIncrementIncrement, true
}

// HasAutoIncrementIncrement returns a boolean if a field has been set.
func (o *ConfigParameters) HasAutoIncrementIncrement() bool {
	if o != nil && !IsNil(o.AutoIncrementIncrement) {
		return true
	}

	return false
}

// SetAutoIncrementIncrement gets a reference to the given string and assigns it to the AutoIncrementIncrement field.
func (o *ConfigParameters) SetAutoIncrementIncrement(v string) {
	o.AutoIncrementIncrement = &v
}

// GetAutoIncrementOffset returns the AutoIncrementOffset field value if set, zero value otherwise.
func (o *ConfigParameters) GetAutoIncrementOffset() string {
	if o == nil || IsNil(o.AutoIncrementOffset) {
		var ret string
		return ret
	}
	return *o.AutoIncrementOffset
}

// GetAutoIncrementOffsetOk returns a tuple with the AutoIncrementOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParameters) GetAutoIncrementOffsetOk() (*string, bool) {
	if o == nil || IsNil(o.AutoIncrementOffset) {
		return nil, false
	}
	return o.AutoIncrementOffset, true
}

// HasAutoIncrementOffset returns a boolean if a field has been set.
func (o *ConfigParameters) HasAutoIncrementOffset() bool {
	if o != nil && !IsNil(o.AutoIncrementOffset) {
		return true
	}

	return false
}

// SetAutoIncrementOffset gets a reference to the given string and assigns it to the AutoIncrementOffset field.
func (o *ConfigParameters) SetAutoIncrementOffset(v string) {
	o.AutoIncrementOffset = &v
}

// GetInnodbIoCapacity returns the InnodbIoCapacity field value if set, zero value otherwise.
func (o *ConfigParameters) GetInnodbIoCapacity() string {
	if o == nil || IsNil(o.InnodbIoCapacity) {
		var ret string
		return ret
	}
	return *o.InnodbIoCapacity
}

// GetInnodbIoCapacityOk returns a tuple with the InnodbIoCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParameters) GetInnodbIoCapacityOk() (*string, bool) {
	if o == nil || IsNil(o.InnodbIoCapacity) {
		return nil, false
	}
	return o.InnodbIoCapacity, true
}

// HasInnodbIoCapacity returns a boolean if a field has been set.
func (o *ConfigParameters) HasInnodbIoCapacity() bool {
	if o != nil && !IsNil(o.InnodbIoCapacity) {
		return true
	}

	return false
}

// SetInnodbIoCapacity gets a reference to the given string and assigns it to the InnodbIoCapacity field.
func (o *ConfigParameters) SetInnodbIoCapacity(v string) {
	o.InnodbIoCapacity = &v
}

// GetInnodbPurgeThreads returns the InnodbPurgeThreads field value if set, zero value otherwise.
func (o *ConfigParameters) GetInnodbPurgeThreads() string {
	if o == nil || IsNil(o.InnodbPurgeThreads) {
		var ret string
		return ret
	}
	return *o.InnodbPurgeThreads
}

// GetInnodbPurgeThreadsOk returns a tuple with the InnodbPurgeThreads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParameters) GetInnodbPurgeThreadsOk() (*string, bool) {
	if o == nil || IsNil(o.InnodbPurgeThreads) {
		return nil, false
	}
	return o.InnodbPurgeThreads, true
}

// HasInnodbPurgeThreads returns a boolean if a field has been set.
func (o *ConfigParameters) HasInnodbPurgeThreads() bool {
	if o != nil && !IsNil(o.InnodbPurgeThreads) {
		return true
	}

	return false
}

// SetInnodbPurgeThreads gets a reference to the given string and assigns it to the InnodbPurgeThreads field.
func (o *ConfigParameters) SetInnodbPurgeThreads(v string) {
	o.InnodbPurgeThreads = &v
}

// GetInnodbReadIoThreads returns the InnodbReadIoThreads field value if set, zero value otherwise.
func (o *ConfigParameters) GetInnodbReadIoThreads() string {
	if o == nil || IsNil(o.InnodbReadIoThreads) {
		var ret string
		return ret
	}
	return *o.InnodbReadIoThreads
}

// GetInnodbReadIoThreadsOk returns a tuple with the InnodbReadIoThreads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParameters) GetInnodbReadIoThreadsOk() (*string, bool) {
	if o == nil || IsNil(o.InnodbReadIoThreads) {
		return nil, false
	}
	return o.InnodbReadIoThreads, true
}

// HasInnodbReadIoThreads returns a boolean if a field has been set.
func (o *ConfigParameters) HasInnodbReadIoThreads() bool {
	if o != nil && !IsNil(o.InnodbReadIoThreads) {
		return true
	}

	return false
}

// SetInnodbReadIoThreads gets a reference to the given string and assigns it to the InnodbReadIoThreads field.
func (o *ConfigParameters) SetInnodbReadIoThreads(v string) {
	o.InnodbReadIoThreads = &v
}

// GetInnodbThreadConcurrency returns the InnodbThreadConcurrency field value if set, zero value otherwise.
func (o *ConfigParameters) GetInnodbThreadConcurrency() string {
	if o == nil || IsNil(o.InnodbThreadConcurrency) {
		var ret string
		return ret
	}
	return *o.InnodbThreadConcurrency
}

// GetInnodbThreadConcurrencyOk returns a tuple with the InnodbThreadConcurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParameters) GetInnodbThreadConcurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.InnodbThreadConcurrency) {
		return nil, false
	}
	return o.InnodbThreadConcurrency, true
}

// HasInnodbThreadConcurrency returns a boolean if a field has been set.
func (o *ConfigParameters) HasInnodbThreadConcurrency() bool {
	if o != nil && !IsNil(o.InnodbThreadConcurrency) {
		return true
	}

	return false
}

// SetInnodbThreadConcurrency gets a reference to the given string and assigns it to the InnodbThreadConcurrency field.
func (o *ConfigParameters) SetInnodbThreadConcurrency(v string) {
	o.InnodbThreadConcurrency = &v
}

// GetInnodbWriteIoThreads returns the InnodbWriteIoThreads field value if set, zero value otherwise.
func (o *ConfigParameters) GetInnodbWriteIoThreads() string {
	if o == nil || IsNil(o.InnodbWriteIoThreads) {
		var ret string
		return ret
	}
	return *o.InnodbWriteIoThreads
}

// GetInnodbWriteIoThreadsOk returns a tuple with the InnodbWriteIoThreads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParameters) GetInnodbWriteIoThreadsOk() (*string, bool) {
	if o == nil || IsNil(o.InnodbWriteIoThreads) {
		return nil, false
	}
	return o.InnodbWriteIoThreads, true
}

// HasInnodbWriteIoThreads returns a boolean if a field has been set.
func (o *ConfigParameters) HasInnodbWriteIoThreads() bool {
	if o != nil && !IsNil(o.InnodbWriteIoThreads) {
		return true
	}

	return false
}

// SetInnodbWriteIoThreads gets a reference to the given string and assigns it to the InnodbWriteIoThreads field.
func (o *ConfigParameters) SetInnodbWriteIoThreads(v string) {
	o.InnodbWriteIoThreads = &v
}

// GetJoinBufferSize returns the JoinBufferSize field value if set, zero value otherwise.
func (o *ConfigParameters) GetJoinBufferSize() string {
	if o == nil || IsNil(o.JoinBufferSize) {
		var ret string
		return ret
	}
	return *o.JoinBufferSize
}

// GetJoinBufferSizeOk returns a tuple with the JoinBufferSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParameters) GetJoinBufferSizeOk() (*string, bool) {
	if o == nil || IsNil(o.JoinBufferSize) {
		return nil, false
	}
	return o.JoinBufferSize, true
}

// HasJoinBufferSize returns a boolean if a field has been set.
func (o *ConfigParameters) HasJoinBufferSize() bool {
	if o != nil && !IsNil(o.JoinBufferSize) {
		return true
	}

	return false
}

// SetJoinBufferSize gets a reference to the given string and assigns it to the JoinBufferSize field.
func (o *ConfigParameters) SetJoinBufferSize(v string) {
	o.JoinBufferSize = &v
}

// GetMaxAllowedPacket returns the MaxAllowedPacket field value if set, zero value otherwise.
func (o *ConfigParameters) GetMaxAllowedPacket() string {
	if o == nil || IsNil(o.MaxAllowedPacket) {
		var ret string
		return ret
	}
	return *o.MaxAllowedPacket
}

// GetMaxAllowedPacketOk returns a tuple with the MaxAllowedPacket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParameters) GetMaxAllowedPacketOk() (*string, bool) {
	if o == nil || IsNil(o.MaxAllowedPacket) {
		return nil, false
	}
	return o.MaxAllowedPacket, true
}

// HasMaxAllowedPacket returns a boolean if a field has been set.
func (o *ConfigParameters) HasMaxAllowedPacket() bool {
	if o != nil && !IsNil(o.MaxAllowedPacket) {
		return true
	}

	return false
}

// SetMaxAllowedPacket gets a reference to the given string and assigns it to the MaxAllowedPacket field.
func (o *ConfigParameters) SetMaxAllowedPacket(v string) {
	o.MaxAllowedPacket = &v
}

// GetMaxHeapTableSize returns the MaxHeapTableSize field value if set, zero value otherwise.
func (o *ConfigParameters) GetMaxHeapTableSize() string {
	if o == nil || IsNil(o.MaxHeapTableSize) {
		var ret string
		return ret
	}
	return *o.MaxHeapTableSize
}

// GetMaxHeapTableSizeOk returns a tuple with the MaxHeapTableSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParameters) GetMaxHeapTableSizeOk() (*string, bool) {
	if o == nil || IsNil(o.MaxHeapTableSize) {
		return nil, false
	}
	return o.MaxHeapTableSize, true
}

// HasMaxHeapTableSize returns a boolean if a field has been set.
func (o *ConfigParameters) HasMaxHeapTableSize() bool {
	if o != nil && !IsNil(o.MaxHeapTableSize) {
		return true
	}

	return false
}

// SetMaxHeapTableSize gets a reference to the given string and assigns it to the MaxHeapTableSize field.
func (o *ConfigParameters) SetMaxHeapTableSize(v string) {
	o.MaxHeapTableSize = &v
}

// GetAutovacuumAnalyzeScaleFactor returns the AutovacuumAnalyzeScaleFactor field value if set, zero value otherwise.
func (o *ConfigParameters) GetAutovacuumAnalyzeScaleFactor() string {
	if o == nil || IsNil(o.AutovacuumAnalyzeScaleFactor) {
		var ret string
		return ret
	}
	return *o.AutovacuumAnalyzeScaleFactor
}

// GetAutovacuumAnalyzeScaleFactorOk returns a tuple with the AutovacuumAnalyzeScaleFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParameters) GetAutovacuumAnalyzeScaleFactorOk() (*string, bool) {
	if o == nil || IsNil(o.AutovacuumAnalyzeScaleFactor) {
		return nil, false
	}
	return o.AutovacuumAnalyzeScaleFactor, true
}

// HasAutovacuumAnalyzeScaleFactor returns a boolean if a field has been set.
func (o *ConfigParameters) HasAutovacuumAnalyzeScaleFactor() bool {
	if o != nil && !IsNil(o.AutovacuumAnalyzeScaleFactor) {
		return true
	}

	return false
}

// SetAutovacuumAnalyzeScaleFactor gets a reference to the given string and assigns it to the AutovacuumAnalyzeScaleFactor field.
func (o *ConfigParameters) SetAutovacuumAnalyzeScaleFactor(v string) {
	o.AutovacuumAnalyzeScaleFactor = &v
}

// GetBgwriterDelay returns the BgwriterDelay field value if set, zero value otherwise.
func (o *ConfigParameters) GetBgwriterDelay() string {
	if o == nil || IsNil(o.BgwriterDelay) {
		var ret string
		return ret
	}
	return *o.BgwriterDelay
}

// GetBgwriterDelayOk returns a tuple with the BgwriterDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParameters) GetBgwriterDelayOk() (*string, bool) {
	if o == nil || IsNil(o.BgwriterDelay) {
		return nil, false
	}
	return o.BgwriterDelay, true
}

// HasBgwriterDelay returns a boolean if a field has been set.
func (o *ConfigParameters) HasBgwriterDelay() bool {
	if o != nil && !IsNil(o.BgwriterDelay) {
		return true
	}

	return false
}

// SetBgwriterDelay gets a reference to the given string and assigns it to the BgwriterDelay field.
func (o *ConfigParameters) SetBgwriterDelay(v string) {
	o.BgwriterDelay = &v
}

// GetBgwriterLruMaxpages returns the BgwriterLruMaxpages field value if set, zero value otherwise.
func (o *ConfigParameters) GetBgwriterLruMaxpages() string {
	if o == nil || IsNil(o.BgwriterLruMaxpages) {
		var ret string
		return ret
	}
	return *o.BgwriterLruMaxpages
}

// GetBgwriterLruMaxpagesOk returns a tuple with the BgwriterLruMaxpages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParameters) GetBgwriterLruMaxpagesOk() (*string, bool) {
	if o == nil || IsNil(o.BgwriterLruMaxpages) {
		return nil, false
	}
	return o.BgwriterLruMaxpages, true
}

// HasBgwriterLruMaxpages returns a boolean if a field has been set.
func (o *ConfigParameters) HasBgwriterLruMaxpages() bool {
	if o != nil && !IsNil(o.BgwriterLruMaxpages) {
		return true
	}

	return false
}

// SetBgwriterLruMaxpages gets a reference to the given string and assigns it to the BgwriterLruMaxpages field.
func (o *ConfigParameters) SetBgwriterLruMaxpages(v string) {
	o.BgwriterLruMaxpages = &v
}

// GetDeadlockTimeout returns the DeadlockTimeout field value if set, zero value otherwise.
func (o *ConfigParameters) GetDeadlockTimeout() string {
	if o == nil || IsNil(o.DeadlockTimeout) {
		var ret string
		return ret
	}
	return *o.DeadlockTimeout
}

// GetDeadlockTimeoutOk returns a tuple with the DeadlockTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParameters) GetDeadlockTimeoutOk() (*string, bool) {
	if o == nil || IsNil(o.DeadlockTimeout) {
		return nil, false
	}
	return o.DeadlockTimeout, true
}

// HasDeadlockTimeout returns a boolean if a field has been set.
func (o *ConfigParameters) HasDeadlockTimeout() bool {
	if o != nil && !IsNil(o.DeadlockTimeout) {
		return true
	}

	return false
}

// SetDeadlockTimeout gets a reference to the given string and assigns it to the DeadlockTimeout field.
func (o *ConfigParameters) SetDeadlockTimeout(v string) {
	o.DeadlockTimeout = &v
}

// GetGinPendingListLimit returns the GinPendingListLimit field value if set, zero value otherwise.
func (o *ConfigParameters) GetGinPendingListLimit() string {
	if o == nil || IsNil(o.GinPendingListLimit) {
		var ret string
		return ret
	}
	return *o.GinPendingListLimit
}

// GetGinPendingListLimitOk returns a tuple with the GinPendingListLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParameters) GetGinPendingListLimitOk() (*string, bool) {
	if o == nil || IsNil(o.GinPendingListLimit) {
		return nil, false
	}
	return o.GinPendingListLimit, true
}

// HasGinPendingListLimit returns a boolean if a field has been set.
func (o *ConfigParameters) HasGinPendingListLimit() bool {
	if o != nil && !IsNil(o.GinPendingListLimit) {
		return true
	}

	return false
}

// SetGinPendingListLimit gets a reference to the given string and assigns it to the GinPendingListLimit field.
func (o *ConfigParameters) SetGinPendingListLimit(v string) {
	o.GinPendingListLimit = &v
}

// GetIdleInTransactionSessionTimeout returns the IdleInTransactionSessionTimeout field value if set, zero value otherwise.
func (o *ConfigParameters) GetIdleInTransactionSessionTimeout() string {
	if o == nil || IsNil(o.IdleInTransactionSessionTimeout) {
		var ret string
		return ret
	}
	return *o.IdleInTransactionSessionTimeout
}

// GetIdleInTransactionSessionTimeoutOk returns a tuple with the IdleInTransactionSessionTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParameters) GetIdleInTransactionSessionTimeoutOk() (*string, bool) {
	if o == nil || IsNil(o.IdleInTransactionSessionTimeout) {
		return nil, false
	}
	return o.IdleInTransactionSessionTimeout, true
}

// HasIdleInTransactionSessionTimeout returns a boolean if a field has been set.
func (o *ConfigParameters) HasIdleInTransactionSessionTimeout() bool {
	if o != nil && !IsNil(o.IdleInTransactionSessionTimeout) {
		return true
	}

	return false
}

// SetIdleInTransactionSessionTimeout gets a reference to the given string and assigns it to the IdleInTransactionSessionTimeout field.
func (o *ConfigParameters) SetIdleInTransactionSessionTimeout(v string) {
	o.IdleInTransactionSessionTimeout = &v
}

// GetIdleSessionTimeout returns the IdleSessionTimeout field value if set, zero value otherwise.
func (o *ConfigParameters) GetIdleSessionTimeout() string {
	if o == nil || IsNil(o.IdleSessionTimeout) {
		var ret string
		return ret
	}
	return *o.IdleSessionTimeout
}

// GetIdleSessionTimeoutOk returns a tuple with the IdleSessionTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParameters) GetIdleSessionTimeoutOk() (*string, bool) {
	if o == nil || IsNil(o.IdleSessionTimeout) {
		return nil, false
	}
	return o.IdleSessionTimeout, true
}

// HasIdleSessionTimeout returns a boolean if a field has been set.
func (o *ConfigParameters) HasIdleSessionTimeout() bool {
	if o != nil && !IsNil(o.IdleSessionTimeout) {
		return true
	}

	return false
}

// SetIdleSessionTimeout gets a reference to the given string and assigns it to the IdleSessionTimeout field.
func (o *ConfigParameters) SetIdleSessionTimeout(v string) {
	o.IdleSessionTimeout = &v
}

// GetJoinCollapseLimit returns the JoinCollapseLimit field value if set, zero value otherwise.
func (o *ConfigParameters) GetJoinCollapseLimit() string {
	if o == nil || IsNil(o.JoinCollapseLimit) {
		var ret string
		return ret
	}
	return *o.JoinCollapseLimit
}

// GetJoinCollapseLimitOk returns a tuple with the JoinCollapseLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParameters) GetJoinCollapseLimitOk() (*string, bool) {
	if o == nil || IsNil(o.JoinCollapseLimit) {
		return nil, false
	}
	return o.JoinCollapseLimit, true
}

// HasJoinCollapseLimit returns a boolean if a field has been set.
func (o *ConfigParameters) HasJoinCollapseLimit() bool {
	if o != nil && !IsNil(o.JoinCollapseLimit) {
		return true
	}

	return false
}

// SetJoinCollapseLimit gets a reference to the given string and assigns it to the JoinCollapseLimit field.
func (o *ConfigParameters) SetJoinCollapseLimit(v string) {
	o.JoinCollapseLimit = &v
}

// GetLockTimeout returns the LockTimeout field value if set, zero value otherwise.
func (o *ConfigParameters) GetLockTimeout() string {
	if o == nil || IsNil(o.LockTimeout) {
		var ret string
		return ret
	}
	return *o.LockTimeout
}

// GetLockTimeoutOk returns a tuple with the LockTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParameters) GetLockTimeoutOk() (*string, bool) {
	if o == nil || IsNil(o.LockTimeout) {
		return nil, false
	}
	return o.LockTimeout, true
}

// HasLockTimeout returns a boolean if a field has been set.
func (o *ConfigParameters) HasLockTimeout() bool {
	if o != nil && !IsNil(o.LockTimeout) {
		return true
	}

	return false
}

// SetLockTimeout gets a reference to the given string and assigns it to the LockTimeout field.
func (o *ConfigParameters) SetLockTimeout(v string) {
	o.LockTimeout = &v
}

// GetMaxPreparedTransactions returns the MaxPreparedTransactions field value if set, zero value otherwise.
func (o *ConfigParameters) GetMaxPreparedTransactions() string {
	if o == nil || IsNil(o.MaxPreparedTransactions) {
		var ret string
		return ret
	}
	return *o.MaxPreparedTransactions
}

// GetMaxPreparedTransactionsOk returns a tuple with the MaxPreparedTransactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParameters) GetMaxPreparedTransactionsOk() (*string, bool) {
	if o == nil || IsNil(o.MaxPreparedTransactions) {
		return nil, false
	}
	return o.MaxPreparedTransactions, true
}

// HasMaxPreparedTransactions returns a boolean if a field has been set.
func (o *ConfigParameters) HasMaxPreparedTransactions() bool {
	if o != nil && !IsNil(o.MaxPreparedTransactions) {
		return true
	}

	return false
}

// SetMaxPreparedTransactions gets a reference to the given string and assigns it to the MaxPreparedTransactions field.
func (o *ConfigParameters) SetMaxPreparedTransactions(v string) {
	o.MaxPreparedTransactions = &v
}

// GetMaxConnections returns the MaxConnections field value if set, zero value otherwise.
func (o *ConfigParameters) GetMaxConnections() string {
	if o == nil || IsNil(o.MaxConnections) {
		var ret string
		return ret
	}
	return *o.MaxConnections
}

// GetMaxConnectionsOk returns a tuple with the MaxConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParameters) GetMaxConnectionsOk() (*string, bool) {
	if o == nil || IsNil(o.MaxConnections) {
		return nil, false
	}
	return o.MaxConnections, true
}

// HasMaxConnections returns a boolean if a field has been set.
func (o *ConfigParameters) HasMaxConnections() bool {
	if o != nil && !IsNil(o.MaxConnections) {
		return true
	}

	return false
}

// SetMaxConnections gets a reference to the given string and assigns it to the MaxConnections field.
func (o *ConfigParameters) SetMaxConnections(v string) {
	o.MaxConnections = &v
}

// GetSharedBuffers returns the SharedBuffers field value if set, zero value otherwise.
func (o *ConfigParameters) GetSharedBuffers() string {
	if o == nil || IsNil(o.SharedBuffers) {
		var ret string
		return ret
	}
	return *o.SharedBuffers
}

// GetSharedBuffersOk returns a tuple with the SharedBuffers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParameters) GetSharedBuffersOk() (*string, bool) {
	if o == nil || IsNil(o.SharedBuffers) {
		return nil, false
	}
	return o.SharedBuffers, true
}

// HasSharedBuffers returns a boolean if a field has been set.
func (o *ConfigParameters) HasSharedBuffers() bool {
	if o != nil && !IsNil(o.SharedBuffers) {
		return true
	}

	return false
}

// SetSharedBuffers gets a reference to the given string and assigns it to the SharedBuffers field.
func (o *ConfigParameters) SetSharedBuffers(v string) {
	o.SharedBuffers = &v
}

// GetWalBuffers returns the WalBuffers field value if set, zero value otherwise.
func (o *ConfigParameters) GetWalBuffers() string {
	if o == nil || IsNil(o.WalBuffers) {
		var ret string
		return ret
	}
	return *o.WalBuffers
}

// GetWalBuffersOk returns a tuple with the WalBuffers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParameters) GetWalBuffersOk() (*string, bool) {
	if o == nil || IsNil(o.WalBuffers) {
		return nil, false
	}
	return o.WalBuffers, true
}

// HasWalBuffers returns a boolean if a field has been set.
func (o *ConfigParameters) HasWalBuffers() bool {
	if o != nil && !IsNil(o.WalBuffers) {
		return true
	}

	return false
}

// SetWalBuffers gets a reference to the given string and assigns it to the WalBuffers field.
func (o *ConfigParameters) SetWalBuffers(v string) {
	o.WalBuffers = &v
}

// GetTempBuffers returns the TempBuffers field value if set, zero value otherwise.
func (o *ConfigParameters) GetTempBuffers() string {
	if o == nil || IsNil(o.TempBuffers) {
		var ret string
		return ret
	}
	return *o.TempBuffers
}

// GetTempBuffersOk returns a tuple with the TempBuffers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParameters) GetTempBuffersOk() (*string, bool) {
	if o == nil || IsNil(o.TempBuffers) {
		return nil, false
	}
	return o.TempBuffers, true
}

// HasTempBuffers returns a boolean if a field has been set.
func (o *ConfigParameters) HasTempBuffers() bool {
	if o != nil && !IsNil(o.TempBuffers) {
		return true
	}

	return false
}

// SetTempBuffers gets a reference to the given string and assigns it to the TempBuffers field.
func (o *ConfigParameters) SetTempBuffers(v string) {
	o.TempBuffers = &v
}

// GetWorkMem returns the WorkMem field value if set, zero value otherwise.
func (o *ConfigParameters) GetWorkMem() string {
	if o == nil || IsNil(o.WorkMem) {
		var ret string
		return ret
	}
	return *o.WorkMem
}

// GetWorkMemOk returns a tuple with the WorkMem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParameters) GetWorkMemOk() (*string, bool) {
	if o == nil || IsNil(o.WorkMem) {
		return nil, false
	}
	return o.WorkMem, true
}

// HasWorkMem returns a boolean if a field has been set.
func (o *ConfigParameters) HasWorkMem() bool {
	if o != nil && !IsNil(o.WorkMem) {
		return true
	}

	return false
}

// SetWorkMem gets a reference to the given string and assigns it to the WorkMem field.
func (o *ConfigParameters) SetWorkMem(v string) {
	o.WorkMem = &v
}

// GetSqlMode returns the SqlMode field value if set, zero value otherwise.
func (o *ConfigParameters) GetSqlMode() string {
	if o == nil || IsNil(o.SqlMode) {
		var ret string
		return ret
	}
	return *o.SqlMode
}

// GetSqlModeOk returns a tuple with the SqlMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParameters) GetSqlModeOk() (*string, bool) {
	if o == nil || IsNil(o.SqlMode) {
		return nil, false
	}
	return o.SqlMode, true
}

// HasSqlMode returns a boolean if a field has been set.
func (o *ConfigParameters) HasSqlMode() bool {
	if o != nil && !IsNil(o.SqlMode) {
		return true
	}

	return false
}

// SetSqlMode gets a reference to the given string and assigns it to the SqlMode field.
func (o *ConfigParameters) SetSqlMode(v string) {
	o.SqlMode = &v
}

// GetQueryCacheType returns the QueryCacheType field value if set, zero value otherwise.
func (o *ConfigParameters) GetQueryCacheType() string {
	if o == nil || IsNil(o.QueryCacheType) {
		var ret string
		return ret
	}
	return *o.QueryCacheType
}

// GetQueryCacheTypeOk returns a tuple with the QueryCacheType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParameters) GetQueryCacheTypeOk() (*string, bool) {
	if o == nil || IsNil(o.QueryCacheType) {
		return nil, false
	}
	return o.QueryCacheType, true
}

// HasQueryCacheType returns a boolean if a field has been set.
func (o *ConfigParameters) HasQueryCacheType() bool {
	if o != nil && !IsNil(o.QueryCacheType) {
		return true
	}

	return false
}

// SetQueryCacheType gets a reference to the given string and assigns it to the QueryCacheType field.
func (o *ConfigParameters) SetQueryCacheType(v string) {
	o.QueryCacheType = &v
}

// GetQueryCacheSize returns the QueryCacheSize field value if set, zero value otherwise.
func (o *ConfigParameters) GetQueryCacheSize() string {
	if o == nil || IsNil(o.QueryCacheSize) {
		var ret string
		return ret
	}
	return *o.QueryCacheSize
}

// GetQueryCacheSizeOk returns a tuple with the QueryCacheSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParameters) GetQueryCacheSizeOk() (*string, bool) {
	if o == nil || IsNil(o.QueryCacheSize) {
		return nil, false
	}
	return o.QueryCacheSize, true
}

// HasQueryCacheSize returns a boolean if a field has been set.
func (o *ConfigParameters) HasQueryCacheSize() bool {
	if o != nil && !IsNil(o.QueryCacheSize) {
		return true
	}

	return false
}

// SetQueryCacheSize gets a reference to the given string and assigns it to the QueryCacheSize field.
func (o *ConfigParameters) SetQueryCacheSize(v string) {
	o.QueryCacheSize = &v
}

func (o ConfigParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoIncrementIncrement) {
		toSerialize["auto_increment_increment"] = o.AutoIncrementIncrement
	}
	if !IsNil(o.AutoIncrementOffset) {
		toSerialize["auto_increment_offset"] = o.AutoIncrementOffset
	}
	if !IsNil(o.InnodbIoCapacity) {
		toSerialize["innodb_io_capacity"] = o.InnodbIoCapacity
	}
	if !IsNil(o.InnodbPurgeThreads) {
		toSerialize["innodb_purge_threads"] = o.InnodbPurgeThreads
	}
	if !IsNil(o.InnodbReadIoThreads) {
		toSerialize["innodb_read_io_threads"] = o.InnodbReadIoThreads
	}
	if !IsNil(o.InnodbThreadConcurrency) {
		toSerialize["innodb_thread_concurrency"] = o.InnodbThreadConcurrency
	}
	if !IsNil(o.InnodbWriteIoThreads) {
		toSerialize["innodb_write_io_threads"] = o.InnodbWriteIoThreads
	}
	if !IsNil(o.JoinBufferSize) {
		toSerialize["join_buffer_size"] = o.JoinBufferSize
	}
	if !IsNil(o.MaxAllowedPacket) {
		toSerialize["max_allowed_packet"] = o.MaxAllowedPacket
	}
	if !IsNil(o.MaxHeapTableSize) {
		toSerialize["max_heap_table_size"] = o.MaxHeapTableSize
	}
	if !IsNil(o.AutovacuumAnalyzeScaleFactor) {
		toSerialize["autovacuum_analyze_scale_factor"] = o.AutovacuumAnalyzeScaleFactor
	}
	if !IsNil(o.BgwriterDelay) {
		toSerialize["bgwriter_delay"] = o.BgwriterDelay
	}
	if !IsNil(o.BgwriterLruMaxpages) {
		toSerialize["bgwriter_lru_maxpages"] = o.BgwriterLruMaxpages
	}
	if !IsNil(o.DeadlockTimeout) {
		toSerialize["deadlock_timeout"] = o.DeadlockTimeout
	}
	if !IsNil(o.GinPendingListLimit) {
		toSerialize["gin_pending_list_limit"] = o.GinPendingListLimit
	}
	if !IsNil(o.IdleInTransactionSessionTimeout) {
		toSerialize["idle_in_transaction_session_timeout"] = o.IdleInTransactionSessionTimeout
	}
	if !IsNil(o.IdleSessionTimeout) {
		toSerialize["idle_session_timeout"] = o.IdleSessionTimeout
	}
	if !IsNil(o.JoinCollapseLimit) {
		toSerialize["join_collapse_limit"] = o.JoinCollapseLimit
	}
	if !IsNil(o.LockTimeout) {
		toSerialize["lock_timeout"] = o.LockTimeout
	}
	if !IsNil(o.MaxPreparedTransactions) {
		toSerialize["max_prepared_transactions"] = o.MaxPreparedTransactions
	}
	if !IsNil(o.MaxConnections) {
		toSerialize["max_connections"] = o.MaxConnections
	}
	if !IsNil(o.SharedBuffers) {
		toSerialize["shared_buffers"] = o.SharedBuffers
	}
	if !IsNil(o.WalBuffers) {
		toSerialize["wal_buffers"] = o.WalBuffers
	}
	if !IsNil(o.TempBuffers) {
		toSerialize["temp_buffers"] = o.TempBuffers
	}
	if !IsNil(o.WorkMem) {
		toSerialize["work_mem"] = o.WorkMem
	}
	if !IsNil(o.SqlMode) {
		toSerialize["sql_mode"] = o.SqlMode
	}
	if !IsNil(o.QueryCacheType) {
		toSerialize["query_cache_type"] = o.QueryCacheType
	}
	if !IsNil(o.QueryCacheSize) {
		toSerialize["query_cache_size"] = o.QueryCacheSize
	}
	return toSerialize, nil
}

type NullableConfigParameters struct {
	value *ConfigParameters
	isSet bool
}

func (v NullableConfigParameters) Get() *ConfigParameters {
	return v.value
}

func (v *NullableConfigParameters) Set(val *ConfigParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigParameters(val *ConfigParameters) *NullableConfigParameters {
	return &NullableConfigParameters{value: val, isSet: true}
}

func (v NullableConfigParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


