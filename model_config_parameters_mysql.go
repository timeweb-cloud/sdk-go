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

// checks if the ConfigParametersMysql type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigParametersMysql{}

// ConfigParametersMysql Параметры MySQL (`mysql5` | `mysql` | `mysql8_4`)
type ConfigParametersMysql struct {
	// Размер буфера, используемого при соединениях таблиц без индексов (`mysql5` | `mysql` | `mysql8_4`).
	JoinBufferSize *string `json:"join_buffer_size,omitempty"`
	// Максимальное количество одновременных подключений к серверу (`mysql5` | `mysql` | `mysql8_4` | `postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	MaxConnections *string `json:"max_connections,omitempty"`
	// Размер буфера сортировки для операций ORDER BY и GROUP BY (`mysql5` | `mysql` | `mysql8_4`).
	SortBufferSize *string `json:"sort_buffer_size,omitempty"`
	// Количество потоков, которые сервер сохраняет для повторного использования (`mysql5` | `mysql` | `mysql8_4`).
	ThreadCacheSize *string `json:"thread_cache_size,omitempty"`
	// Размер буферного пула InnoDB для хранения данных и индексов в памяти (`mysql5` | `mysql` | `mysql8_4`).
	InnodbBufferPoolSize *string `json:"innodb_buffer_pool_size,omitempty"`
	// Интервал между значениями столбцов с атрибутом `AUTO_INCREMENT` (`mysql5` | `mysql` | `mysql8_4`).
	AutoIncrementIncrement *string `json:"auto_increment_increment,omitempty"`
	// Начальное значение для столбцов с атрибутом `AUTO_INCREMENT` (`mysql5` | `mysql` | `mysql8_4`).
	AutoIncrementOffset *string `json:"auto_increment_offset,omitempty"`
	// Количество операций ввода-вывода в секунду `IOPS`, используемых InnoDB (`mysql5` | `mysql` | `mysql8_4`).
	InnodbIoCapacity *string `json:"innodb_io_capacity,omitempty"`
	// Количество потоков, используемых для фоновой очистки undo-записей InnoDB (`mysql5` | `mysql` | `mysql8_4`).
	InnodbPurgeThreads *string `json:"innodb_purge_threads,omitempty"`
	// Количество потоков ввода-вывода для операций чтения InnoDB (`mysql5` | `mysql` | `mysql8_4`).
	InnodbReadIoThreads *string `json:"innodb_read_io_threads,omitempty"`
	// Ограничение количества одновременно выполняющихся потоков InnoDB (`mysql5` | `mysql` | `mysql8_4`).
	InnodbThreadConcurrency *string `json:"innodb_thread_concurrency,omitempty"`
	// Количество потоков ввода-вывода для операций записи InnoDB (`mysql5` | `mysql` | `mysql8_4`).
	InnodbWriteIoThreads *string `json:"innodb_write_io_threads,omitempty"`
	// Размер файла журнала транзакций InnoDB redo log (`mysql5` | `mysql` | `mysql8_4`).
	InnodbLogFileSize *string `json:"innodb_log_file_size,omitempty"`
	// Максимальный размер пакета данных, который может передаваться между клиентом и сервером (`mysql5` | `mysql` | `mysql8_4`).
	MaxAllowedPacket *string `json:"max_allowed_packet,omitempty"`
	// Максимальный размер таблиц типа MEMORY (`mysql5` | `mysql` | `mysql8_4`).
	MaxHeapTableSize *string `json:"max_heap_table_size,omitempty"`
	// Режим работы SQL сервера, определяющий поведение обработки запросов (`mysql5` | `mysql` | `mysql8_4`).
	SqlMode *string `json:"sql_mode,omitempty"`
	// Тип кэша запросов (`mysql5` | `mysql` | `mysql8_4`).
	QueryCacheType *string `json:"query_cache_type,omitempty"`
	// Объем памяти, выделяемый для кэширования результатов запросов (`mysql5` | `mysql` | `mysql8_4`).
	QueryCacheSize *string `json:"query_cache_size,omitempty"`
	// Режим записи журнала InnoDB при фиксации транзакций (`mysql5` | `mysql` | `mysql8_4`).
	InnodbFlushLogAtTrxCommit *string `json:"innodb_flush_log_at_trx_commit,omitempty"`
	// Уровень изоляции транзакций по умолчанию (`mysql5` | `mysql` | `mysql8_4`).
	TransactionIsolation *string `json:"transaction_isolation,omitempty"`
	// Время выполнения запроса, после которого он считается долгим и может попасть в slow query log (`mysql5` | `mysql` | `mysql8_4`).
	LongQueryTime *string `json:"long_query_time,omitempty"`
	// Максимальный размер временных таблиц в памяти (`mysql5` | `mysql` | `mysql8_4`).
	TmpTableSize *string `json:"tmp_table_size,omitempty"`
	// Количество открытых таблиц, которые сервер может хранить в кэше (`mysql5` | `mysql` | `mysql8_4`).
	TableOpenCache *string `json:"table_open_cache,omitempty"`
	// Количество экземпляров кэша открытых таблиц для снижения конкуренции между потоками (`mysql5` | `mysql` | `mysql8_4`).
	TableOpenCacheInstances *string `json:"table_open_cache_instances,omitempty"`
	// Метод выполнения операций записи и синхронизации файлов InnoDB (`mysql5` | `mysql` | `mysql8_4`).
	InnodbFlushMethod *string `json:"innodb_flush_method,omitempty"`
	// Включение строгой проверки операций InnoDB (`mysql5` | `mysql` | `mysql8_4`).
	InnodbStrictMode *string `json:"innodb_strict_mode,omitempty"`
	// Включение журнала медленных запросов (`mysql5` | `mysql` | `mysql8_4`).
	SlowQueryLog *string `json:"slow_query_log,omitempty"`
	// Размер кэша бинарного журнала для транзакций (`mysql5` | `mysql` | `mysql8_4`).
	BinlogCacheSize *string `json:"binlog_cache_size,omitempty"`
	// Задержка синхронизации групповой фиксации бинарного журнала в микросекундах (`mysql5` | `mysql` | `mysql8_4`).
	BinlogGroupCommitSyncDelay *string `json:"binlog_group_commit_sync_delay,omitempty"`
	// Количество информации, записываемой в бинарный журнал при row-based репликации (`mysql5` | `mysql` | `mysql8_4`).
	BinlogRowImage *string `json:"binlog_row_image,omitempty"`
	// Включение записи SQL-запросов в бинарный журнал при row-based репликации (`mysql5` | `mysql` | `mysql8_4`).
	BinlogRowsQueryLogEvents *string `json:"binlog_rows_query_log_events,omitempty"`
	// Кодировка по умолчанию для сервера MySQL (`mysql5` | `mysql` | `mysql8_4`).
	CharacterSetServer *string `json:"character_set_server,omitempty"`
	// Определяет автоматическое поведение TIMESTAMP без явных значений по умолчанию (`mysql5` | `mysql` | `mysql8_4`).
	ExplicitDefaultsForTimestamp *string `json:"explicit_defaults_for_timestamp,omitempty"`
	// Максимальная длина результата функции GROUP_CONCAT (`mysql5` | `mysql` | `mysql8_4`).
	GroupConcatMaxLen *string `json:"group_concat_max_len,omitempty"`
	// Включение или отключение адаптивного хэш-индекса InnoDB для ускорения поиска по индексам (`mysql5` | `mysql` | `mysql8_4`).
	InnodbAdaptiveHashIndex *string `json:"innodb_adaptive_hash_index,omitempty"`
	// Время ожидания блокировки InnoDB перед завершением транзакции с ошибкой (`mysql5` | `mysql` | `mysql8_4`).
	InnodbLockWaitTimeout *string `json:"innodb_lock_wait_timeout,omitempty"`
	// Включение распределения памяти InnoDB между NUMA-узлами (`mysql5` | `mysql` | `mysql8_4`).
	InnodbNumaInterleave *string `json:"innodb_numa_interleave,omitempty"`
	// Время ожидания данных от клиента при чтении сетевого соединения (`mysql5` | `mysql` | `mysql8_4`).
	NetReadTimeout *string `json:"net_read_timeout,omitempty"`
	// Время ожидания записи данных клиенту через сетевое соединение (`mysql5` | `mysql` | `mysql8_4`).
	NetWriteTimeout *string `json:"net_write_timeout,omitempty"`
	// Максимальное время выполнения регулярных выражений (`mysql` | `mysql8_4`).
	RegexpTimeLimit *string `json:"regexp_time_limit,omitempty"`
	// Количество операций записи бинарного журнала перед принудительной синхронизацией на диск (`mysql5` | `mysql` | `mysql8_4`).
	SyncBinlog *string `json:"sync_binlog,omitempty"`
	// Количество определений таблиц, хранящихся в кэше (`mysql5` | `mysql` | `mysql8_4`).
	TableDefinitionCache *string `json:"table_definition_cache,omitempty"`
	// Разрешение создания хранимых функций без проверки бинарной регистрации (`mysql5` | `mysql` | `mysql8_4`).
	LogBinTrustFunctionCreators *string `json:"log_bin_trust_function_creators,omitempty"`
	// Отключение DNS-разрешения имен клиентов при подключении к серверу (`mysql5` | `mysql` | `mysql8_4`).
	SkipNameResolve *string `json:"skip_name_resolve,omitempty"`
	// Общий размер redo log InnoDB для хранения журнала восстановления (`mysql8_4`).
	InnodbRedoLogCapacity *string `json:"innodb_redo_log_capacity,omitempty"`
	// Время ожидания неактивного клиентского соединения перед закрытием (`mysql5` | `mysql` | `mysql8_4`).
	WaitTimeout *string `json:"wait_timeout,omitempty"`
	// Время ожидания неактивного интерактивного соединения перед закрытием (`mysql5` | `mysql` | `mysql8_4`).
	InteractiveTimeout *string `json:"interactive_timeout,omitempty"`
	// Часовой пояс сервера MySQL по умолчанию (`mysql5` | `mysql` | `mysql8_4`).
	DefaultTimeZone *string `json:"default-time-zone,omitempty"`
	// Режим строгой проверки операций в Percona XtraDB Cluster (`mysql5` | `mysql` | `mysql8_4`).
	PxcStrictMode *string `json:"pxc_strict_mode,omitempty"`
}

// NewConfigParametersMysql instantiates a new ConfigParametersMysql object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigParametersMysql() *ConfigParametersMysql {
	this := ConfigParametersMysql{}
	return &this
}

// NewConfigParametersMysqlWithDefaults instantiates a new ConfigParametersMysql object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigParametersMysqlWithDefaults() *ConfigParametersMysql {
	this := ConfigParametersMysql{}
	return &this
}

// GetJoinBufferSize returns the JoinBufferSize field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetJoinBufferSize() string {
	if o == nil || IsNil(o.JoinBufferSize) {
		var ret string
		return ret
	}
	return *o.JoinBufferSize
}

// GetJoinBufferSizeOk returns a tuple with the JoinBufferSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetJoinBufferSizeOk() (*string, bool) {
	if o == nil || IsNil(o.JoinBufferSize) {
		return nil, false
	}
	return o.JoinBufferSize, true
}

// HasJoinBufferSize returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasJoinBufferSize() bool {
	if o != nil && !IsNil(o.JoinBufferSize) {
		return true
	}

	return false
}

// SetJoinBufferSize gets a reference to the given string and assigns it to the JoinBufferSize field.
func (o *ConfigParametersMysql) SetJoinBufferSize(v string) {
	o.JoinBufferSize = &v
}

// GetMaxConnections returns the MaxConnections field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetMaxConnections() string {
	if o == nil || IsNil(o.MaxConnections) {
		var ret string
		return ret
	}
	return *o.MaxConnections
}

// GetMaxConnectionsOk returns a tuple with the MaxConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetMaxConnectionsOk() (*string, bool) {
	if o == nil || IsNil(o.MaxConnections) {
		return nil, false
	}
	return o.MaxConnections, true
}

// HasMaxConnections returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasMaxConnections() bool {
	if o != nil && !IsNil(o.MaxConnections) {
		return true
	}

	return false
}

// SetMaxConnections gets a reference to the given string and assigns it to the MaxConnections field.
func (o *ConfigParametersMysql) SetMaxConnections(v string) {
	o.MaxConnections = &v
}

// GetSortBufferSize returns the SortBufferSize field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetSortBufferSize() string {
	if o == nil || IsNil(o.SortBufferSize) {
		var ret string
		return ret
	}
	return *o.SortBufferSize
}

// GetSortBufferSizeOk returns a tuple with the SortBufferSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetSortBufferSizeOk() (*string, bool) {
	if o == nil || IsNil(o.SortBufferSize) {
		return nil, false
	}
	return o.SortBufferSize, true
}

// HasSortBufferSize returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasSortBufferSize() bool {
	if o != nil && !IsNil(o.SortBufferSize) {
		return true
	}

	return false
}

// SetSortBufferSize gets a reference to the given string and assigns it to the SortBufferSize field.
func (o *ConfigParametersMysql) SetSortBufferSize(v string) {
	o.SortBufferSize = &v
}

// GetThreadCacheSize returns the ThreadCacheSize field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetThreadCacheSize() string {
	if o == nil || IsNil(o.ThreadCacheSize) {
		var ret string
		return ret
	}
	return *o.ThreadCacheSize
}

// GetThreadCacheSizeOk returns a tuple with the ThreadCacheSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetThreadCacheSizeOk() (*string, bool) {
	if o == nil || IsNil(o.ThreadCacheSize) {
		return nil, false
	}
	return o.ThreadCacheSize, true
}

// HasThreadCacheSize returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasThreadCacheSize() bool {
	if o != nil && !IsNil(o.ThreadCacheSize) {
		return true
	}

	return false
}

// SetThreadCacheSize gets a reference to the given string and assigns it to the ThreadCacheSize field.
func (o *ConfigParametersMysql) SetThreadCacheSize(v string) {
	o.ThreadCacheSize = &v
}

// GetInnodbBufferPoolSize returns the InnodbBufferPoolSize field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetInnodbBufferPoolSize() string {
	if o == nil || IsNil(o.InnodbBufferPoolSize) {
		var ret string
		return ret
	}
	return *o.InnodbBufferPoolSize
}

// GetInnodbBufferPoolSizeOk returns a tuple with the InnodbBufferPoolSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetInnodbBufferPoolSizeOk() (*string, bool) {
	if o == nil || IsNil(o.InnodbBufferPoolSize) {
		return nil, false
	}
	return o.InnodbBufferPoolSize, true
}

// HasInnodbBufferPoolSize returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasInnodbBufferPoolSize() bool {
	if o != nil && !IsNil(o.InnodbBufferPoolSize) {
		return true
	}

	return false
}

// SetInnodbBufferPoolSize gets a reference to the given string and assigns it to the InnodbBufferPoolSize field.
func (o *ConfigParametersMysql) SetInnodbBufferPoolSize(v string) {
	o.InnodbBufferPoolSize = &v
}

// GetAutoIncrementIncrement returns the AutoIncrementIncrement field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetAutoIncrementIncrement() string {
	if o == nil || IsNil(o.AutoIncrementIncrement) {
		var ret string
		return ret
	}
	return *o.AutoIncrementIncrement
}

// GetAutoIncrementIncrementOk returns a tuple with the AutoIncrementIncrement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetAutoIncrementIncrementOk() (*string, bool) {
	if o == nil || IsNil(o.AutoIncrementIncrement) {
		return nil, false
	}
	return o.AutoIncrementIncrement, true
}

// HasAutoIncrementIncrement returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasAutoIncrementIncrement() bool {
	if o != nil && !IsNil(o.AutoIncrementIncrement) {
		return true
	}

	return false
}

// SetAutoIncrementIncrement gets a reference to the given string and assigns it to the AutoIncrementIncrement field.
func (o *ConfigParametersMysql) SetAutoIncrementIncrement(v string) {
	o.AutoIncrementIncrement = &v
}

// GetAutoIncrementOffset returns the AutoIncrementOffset field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetAutoIncrementOffset() string {
	if o == nil || IsNil(o.AutoIncrementOffset) {
		var ret string
		return ret
	}
	return *o.AutoIncrementOffset
}

// GetAutoIncrementOffsetOk returns a tuple with the AutoIncrementOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetAutoIncrementOffsetOk() (*string, bool) {
	if o == nil || IsNil(o.AutoIncrementOffset) {
		return nil, false
	}
	return o.AutoIncrementOffset, true
}

// HasAutoIncrementOffset returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasAutoIncrementOffset() bool {
	if o != nil && !IsNil(o.AutoIncrementOffset) {
		return true
	}

	return false
}

// SetAutoIncrementOffset gets a reference to the given string and assigns it to the AutoIncrementOffset field.
func (o *ConfigParametersMysql) SetAutoIncrementOffset(v string) {
	o.AutoIncrementOffset = &v
}

// GetInnodbIoCapacity returns the InnodbIoCapacity field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetInnodbIoCapacity() string {
	if o == nil || IsNil(o.InnodbIoCapacity) {
		var ret string
		return ret
	}
	return *o.InnodbIoCapacity
}

// GetInnodbIoCapacityOk returns a tuple with the InnodbIoCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetInnodbIoCapacityOk() (*string, bool) {
	if o == nil || IsNil(o.InnodbIoCapacity) {
		return nil, false
	}
	return o.InnodbIoCapacity, true
}

// HasInnodbIoCapacity returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasInnodbIoCapacity() bool {
	if o != nil && !IsNil(o.InnodbIoCapacity) {
		return true
	}

	return false
}

// SetInnodbIoCapacity gets a reference to the given string and assigns it to the InnodbIoCapacity field.
func (o *ConfigParametersMysql) SetInnodbIoCapacity(v string) {
	o.InnodbIoCapacity = &v
}

// GetInnodbPurgeThreads returns the InnodbPurgeThreads field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetInnodbPurgeThreads() string {
	if o == nil || IsNil(o.InnodbPurgeThreads) {
		var ret string
		return ret
	}
	return *o.InnodbPurgeThreads
}

// GetInnodbPurgeThreadsOk returns a tuple with the InnodbPurgeThreads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetInnodbPurgeThreadsOk() (*string, bool) {
	if o == nil || IsNil(o.InnodbPurgeThreads) {
		return nil, false
	}
	return o.InnodbPurgeThreads, true
}

// HasInnodbPurgeThreads returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasInnodbPurgeThreads() bool {
	if o != nil && !IsNil(o.InnodbPurgeThreads) {
		return true
	}

	return false
}

// SetInnodbPurgeThreads gets a reference to the given string and assigns it to the InnodbPurgeThreads field.
func (o *ConfigParametersMysql) SetInnodbPurgeThreads(v string) {
	o.InnodbPurgeThreads = &v
}

// GetInnodbReadIoThreads returns the InnodbReadIoThreads field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetInnodbReadIoThreads() string {
	if o == nil || IsNil(o.InnodbReadIoThreads) {
		var ret string
		return ret
	}
	return *o.InnodbReadIoThreads
}

// GetInnodbReadIoThreadsOk returns a tuple with the InnodbReadIoThreads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetInnodbReadIoThreadsOk() (*string, bool) {
	if o == nil || IsNil(o.InnodbReadIoThreads) {
		return nil, false
	}
	return o.InnodbReadIoThreads, true
}

// HasInnodbReadIoThreads returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasInnodbReadIoThreads() bool {
	if o != nil && !IsNil(o.InnodbReadIoThreads) {
		return true
	}

	return false
}

// SetInnodbReadIoThreads gets a reference to the given string and assigns it to the InnodbReadIoThreads field.
func (o *ConfigParametersMysql) SetInnodbReadIoThreads(v string) {
	o.InnodbReadIoThreads = &v
}

// GetInnodbThreadConcurrency returns the InnodbThreadConcurrency field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetInnodbThreadConcurrency() string {
	if o == nil || IsNil(o.InnodbThreadConcurrency) {
		var ret string
		return ret
	}
	return *o.InnodbThreadConcurrency
}

// GetInnodbThreadConcurrencyOk returns a tuple with the InnodbThreadConcurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetInnodbThreadConcurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.InnodbThreadConcurrency) {
		return nil, false
	}
	return o.InnodbThreadConcurrency, true
}

// HasInnodbThreadConcurrency returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasInnodbThreadConcurrency() bool {
	if o != nil && !IsNil(o.InnodbThreadConcurrency) {
		return true
	}

	return false
}

// SetInnodbThreadConcurrency gets a reference to the given string and assigns it to the InnodbThreadConcurrency field.
func (o *ConfigParametersMysql) SetInnodbThreadConcurrency(v string) {
	o.InnodbThreadConcurrency = &v
}

// GetInnodbWriteIoThreads returns the InnodbWriteIoThreads field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetInnodbWriteIoThreads() string {
	if o == nil || IsNil(o.InnodbWriteIoThreads) {
		var ret string
		return ret
	}
	return *o.InnodbWriteIoThreads
}

// GetInnodbWriteIoThreadsOk returns a tuple with the InnodbWriteIoThreads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetInnodbWriteIoThreadsOk() (*string, bool) {
	if o == nil || IsNil(o.InnodbWriteIoThreads) {
		return nil, false
	}
	return o.InnodbWriteIoThreads, true
}

// HasInnodbWriteIoThreads returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasInnodbWriteIoThreads() bool {
	if o != nil && !IsNil(o.InnodbWriteIoThreads) {
		return true
	}

	return false
}

// SetInnodbWriteIoThreads gets a reference to the given string and assigns it to the InnodbWriteIoThreads field.
func (o *ConfigParametersMysql) SetInnodbWriteIoThreads(v string) {
	o.InnodbWriteIoThreads = &v
}

// GetInnodbLogFileSize returns the InnodbLogFileSize field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetInnodbLogFileSize() string {
	if o == nil || IsNil(o.InnodbLogFileSize) {
		var ret string
		return ret
	}
	return *o.InnodbLogFileSize
}

// GetInnodbLogFileSizeOk returns a tuple with the InnodbLogFileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetInnodbLogFileSizeOk() (*string, bool) {
	if o == nil || IsNil(o.InnodbLogFileSize) {
		return nil, false
	}
	return o.InnodbLogFileSize, true
}

// HasInnodbLogFileSize returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasInnodbLogFileSize() bool {
	if o != nil && !IsNil(o.InnodbLogFileSize) {
		return true
	}

	return false
}

// SetInnodbLogFileSize gets a reference to the given string and assigns it to the InnodbLogFileSize field.
func (o *ConfigParametersMysql) SetInnodbLogFileSize(v string) {
	o.InnodbLogFileSize = &v
}

// GetMaxAllowedPacket returns the MaxAllowedPacket field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetMaxAllowedPacket() string {
	if o == nil || IsNil(o.MaxAllowedPacket) {
		var ret string
		return ret
	}
	return *o.MaxAllowedPacket
}

// GetMaxAllowedPacketOk returns a tuple with the MaxAllowedPacket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetMaxAllowedPacketOk() (*string, bool) {
	if o == nil || IsNil(o.MaxAllowedPacket) {
		return nil, false
	}
	return o.MaxAllowedPacket, true
}

// HasMaxAllowedPacket returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasMaxAllowedPacket() bool {
	if o != nil && !IsNil(o.MaxAllowedPacket) {
		return true
	}

	return false
}

// SetMaxAllowedPacket gets a reference to the given string and assigns it to the MaxAllowedPacket field.
func (o *ConfigParametersMysql) SetMaxAllowedPacket(v string) {
	o.MaxAllowedPacket = &v
}

// GetMaxHeapTableSize returns the MaxHeapTableSize field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetMaxHeapTableSize() string {
	if o == nil || IsNil(o.MaxHeapTableSize) {
		var ret string
		return ret
	}
	return *o.MaxHeapTableSize
}

// GetMaxHeapTableSizeOk returns a tuple with the MaxHeapTableSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetMaxHeapTableSizeOk() (*string, bool) {
	if o == nil || IsNil(o.MaxHeapTableSize) {
		return nil, false
	}
	return o.MaxHeapTableSize, true
}

// HasMaxHeapTableSize returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasMaxHeapTableSize() bool {
	if o != nil && !IsNil(o.MaxHeapTableSize) {
		return true
	}

	return false
}

// SetMaxHeapTableSize gets a reference to the given string and assigns it to the MaxHeapTableSize field.
func (o *ConfigParametersMysql) SetMaxHeapTableSize(v string) {
	o.MaxHeapTableSize = &v
}

// GetSqlMode returns the SqlMode field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetSqlMode() string {
	if o == nil || IsNil(o.SqlMode) {
		var ret string
		return ret
	}
	return *o.SqlMode
}

// GetSqlModeOk returns a tuple with the SqlMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetSqlModeOk() (*string, bool) {
	if o == nil || IsNil(o.SqlMode) {
		return nil, false
	}
	return o.SqlMode, true
}

// HasSqlMode returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasSqlMode() bool {
	if o != nil && !IsNil(o.SqlMode) {
		return true
	}

	return false
}

// SetSqlMode gets a reference to the given string and assigns it to the SqlMode field.
func (o *ConfigParametersMysql) SetSqlMode(v string) {
	o.SqlMode = &v
}

// GetQueryCacheType returns the QueryCacheType field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetQueryCacheType() string {
	if o == nil || IsNil(o.QueryCacheType) {
		var ret string
		return ret
	}
	return *o.QueryCacheType
}

// GetQueryCacheTypeOk returns a tuple with the QueryCacheType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetQueryCacheTypeOk() (*string, bool) {
	if o == nil || IsNil(o.QueryCacheType) {
		return nil, false
	}
	return o.QueryCacheType, true
}

// HasQueryCacheType returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasQueryCacheType() bool {
	if o != nil && !IsNil(o.QueryCacheType) {
		return true
	}

	return false
}

// SetQueryCacheType gets a reference to the given string and assigns it to the QueryCacheType field.
func (o *ConfigParametersMysql) SetQueryCacheType(v string) {
	o.QueryCacheType = &v
}

// GetQueryCacheSize returns the QueryCacheSize field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetQueryCacheSize() string {
	if o == nil || IsNil(o.QueryCacheSize) {
		var ret string
		return ret
	}
	return *o.QueryCacheSize
}

// GetQueryCacheSizeOk returns a tuple with the QueryCacheSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetQueryCacheSizeOk() (*string, bool) {
	if o == nil || IsNil(o.QueryCacheSize) {
		return nil, false
	}
	return o.QueryCacheSize, true
}

// HasQueryCacheSize returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasQueryCacheSize() bool {
	if o != nil && !IsNil(o.QueryCacheSize) {
		return true
	}

	return false
}

// SetQueryCacheSize gets a reference to the given string and assigns it to the QueryCacheSize field.
func (o *ConfigParametersMysql) SetQueryCacheSize(v string) {
	o.QueryCacheSize = &v
}

// GetInnodbFlushLogAtTrxCommit returns the InnodbFlushLogAtTrxCommit field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetInnodbFlushLogAtTrxCommit() string {
	if o == nil || IsNil(o.InnodbFlushLogAtTrxCommit) {
		var ret string
		return ret
	}
	return *o.InnodbFlushLogAtTrxCommit
}

// GetInnodbFlushLogAtTrxCommitOk returns a tuple with the InnodbFlushLogAtTrxCommit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetInnodbFlushLogAtTrxCommitOk() (*string, bool) {
	if o == nil || IsNil(o.InnodbFlushLogAtTrxCommit) {
		return nil, false
	}
	return o.InnodbFlushLogAtTrxCommit, true
}

// HasInnodbFlushLogAtTrxCommit returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasInnodbFlushLogAtTrxCommit() bool {
	if o != nil && !IsNil(o.InnodbFlushLogAtTrxCommit) {
		return true
	}

	return false
}

// SetInnodbFlushLogAtTrxCommit gets a reference to the given string and assigns it to the InnodbFlushLogAtTrxCommit field.
func (o *ConfigParametersMysql) SetInnodbFlushLogAtTrxCommit(v string) {
	o.InnodbFlushLogAtTrxCommit = &v
}

// GetTransactionIsolation returns the TransactionIsolation field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetTransactionIsolation() string {
	if o == nil || IsNil(o.TransactionIsolation) {
		var ret string
		return ret
	}
	return *o.TransactionIsolation
}

// GetTransactionIsolationOk returns a tuple with the TransactionIsolation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetTransactionIsolationOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionIsolation) {
		return nil, false
	}
	return o.TransactionIsolation, true
}

// HasTransactionIsolation returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasTransactionIsolation() bool {
	if o != nil && !IsNil(o.TransactionIsolation) {
		return true
	}

	return false
}

// SetTransactionIsolation gets a reference to the given string and assigns it to the TransactionIsolation field.
func (o *ConfigParametersMysql) SetTransactionIsolation(v string) {
	o.TransactionIsolation = &v
}

// GetLongQueryTime returns the LongQueryTime field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetLongQueryTime() string {
	if o == nil || IsNil(o.LongQueryTime) {
		var ret string
		return ret
	}
	return *o.LongQueryTime
}

// GetLongQueryTimeOk returns a tuple with the LongQueryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetLongQueryTimeOk() (*string, bool) {
	if o == nil || IsNil(o.LongQueryTime) {
		return nil, false
	}
	return o.LongQueryTime, true
}

// HasLongQueryTime returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasLongQueryTime() bool {
	if o != nil && !IsNil(o.LongQueryTime) {
		return true
	}

	return false
}

// SetLongQueryTime gets a reference to the given string and assigns it to the LongQueryTime field.
func (o *ConfigParametersMysql) SetLongQueryTime(v string) {
	o.LongQueryTime = &v
}

// GetTmpTableSize returns the TmpTableSize field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetTmpTableSize() string {
	if o == nil || IsNil(o.TmpTableSize) {
		var ret string
		return ret
	}
	return *o.TmpTableSize
}

// GetTmpTableSizeOk returns a tuple with the TmpTableSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetTmpTableSizeOk() (*string, bool) {
	if o == nil || IsNil(o.TmpTableSize) {
		return nil, false
	}
	return o.TmpTableSize, true
}

// HasTmpTableSize returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasTmpTableSize() bool {
	if o != nil && !IsNil(o.TmpTableSize) {
		return true
	}

	return false
}

// SetTmpTableSize gets a reference to the given string and assigns it to the TmpTableSize field.
func (o *ConfigParametersMysql) SetTmpTableSize(v string) {
	o.TmpTableSize = &v
}

// GetTableOpenCache returns the TableOpenCache field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetTableOpenCache() string {
	if o == nil || IsNil(o.TableOpenCache) {
		var ret string
		return ret
	}
	return *o.TableOpenCache
}

// GetTableOpenCacheOk returns a tuple with the TableOpenCache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetTableOpenCacheOk() (*string, bool) {
	if o == nil || IsNil(o.TableOpenCache) {
		return nil, false
	}
	return o.TableOpenCache, true
}

// HasTableOpenCache returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasTableOpenCache() bool {
	if o != nil && !IsNil(o.TableOpenCache) {
		return true
	}

	return false
}

// SetTableOpenCache gets a reference to the given string and assigns it to the TableOpenCache field.
func (o *ConfigParametersMysql) SetTableOpenCache(v string) {
	o.TableOpenCache = &v
}

// GetTableOpenCacheInstances returns the TableOpenCacheInstances field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetTableOpenCacheInstances() string {
	if o == nil || IsNil(o.TableOpenCacheInstances) {
		var ret string
		return ret
	}
	return *o.TableOpenCacheInstances
}

// GetTableOpenCacheInstancesOk returns a tuple with the TableOpenCacheInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetTableOpenCacheInstancesOk() (*string, bool) {
	if o == nil || IsNil(o.TableOpenCacheInstances) {
		return nil, false
	}
	return o.TableOpenCacheInstances, true
}

// HasTableOpenCacheInstances returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasTableOpenCacheInstances() bool {
	if o != nil && !IsNil(o.TableOpenCacheInstances) {
		return true
	}

	return false
}

// SetTableOpenCacheInstances gets a reference to the given string and assigns it to the TableOpenCacheInstances field.
func (o *ConfigParametersMysql) SetTableOpenCacheInstances(v string) {
	o.TableOpenCacheInstances = &v
}

// GetInnodbFlushMethod returns the InnodbFlushMethod field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetInnodbFlushMethod() string {
	if o == nil || IsNil(o.InnodbFlushMethod) {
		var ret string
		return ret
	}
	return *o.InnodbFlushMethod
}

// GetInnodbFlushMethodOk returns a tuple with the InnodbFlushMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetInnodbFlushMethodOk() (*string, bool) {
	if o == nil || IsNil(o.InnodbFlushMethod) {
		return nil, false
	}
	return o.InnodbFlushMethod, true
}

// HasInnodbFlushMethod returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasInnodbFlushMethod() bool {
	if o != nil && !IsNil(o.InnodbFlushMethod) {
		return true
	}

	return false
}

// SetInnodbFlushMethod gets a reference to the given string and assigns it to the InnodbFlushMethod field.
func (o *ConfigParametersMysql) SetInnodbFlushMethod(v string) {
	o.InnodbFlushMethod = &v
}

// GetInnodbStrictMode returns the InnodbStrictMode field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetInnodbStrictMode() string {
	if o == nil || IsNil(o.InnodbStrictMode) {
		var ret string
		return ret
	}
	return *o.InnodbStrictMode
}

// GetInnodbStrictModeOk returns a tuple with the InnodbStrictMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetInnodbStrictModeOk() (*string, bool) {
	if o == nil || IsNil(o.InnodbStrictMode) {
		return nil, false
	}
	return o.InnodbStrictMode, true
}

// HasInnodbStrictMode returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasInnodbStrictMode() bool {
	if o != nil && !IsNil(o.InnodbStrictMode) {
		return true
	}

	return false
}

// SetInnodbStrictMode gets a reference to the given string and assigns it to the InnodbStrictMode field.
func (o *ConfigParametersMysql) SetInnodbStrictMode(v string) {
	o.InnodbStrictMode = &v
}

// GetSlowQueryLog returns the SlowQueryLog field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetSlowQueryLog() string {
	if o == nil || IsNil(o.SlowQueryLog) {
		var ret string
		return ret
	}
	return *o.SlowQueryLog
}

// GetSlowQueryLogOk returns a tuple with the SlowQueryLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetSlowQueryLogOk() (*string, bool) {
	if o == nil || IsNil(o.SlowQueryLog) {
		return nil, false
	}
	return o.SlowQueryLog, true
}

// HasSlowQueryLog returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasSlowQueryLog() bool {
	if o != nil && !IsNil(o.SlowQueryLog) {
		return true
	}

	return false
}

// SetSlowQueryLog gets a reference to the given string and assigns it to the SlowQueryLog field.
func (o *ConfigParametersMysql) SetSlowQueryLog(v string) {
	o.SlowQueryLog = &v
}

// GetBinlogCacheSize returns the BinlogCacheSize field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetBinlogCacheSize() string {
	if o == nil || IsNil(o.BinlogCacheSize) {
		var ret string
		return ret
	}
	return *o.BinlogCacheSize
}

// GetBinlogCacheSizeOk returns a tuple with the BinlogCacheSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetBinlogCacheSizeOk() (*string, bool) {
	if o == nil || IsNil(o.BinlogCacheSize) {
		return nil, false
	}
	return o.BinlogCacheSize, true
}

// HasBinlogCacheSize returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasBinlogCacheSize() bool {
	if o != nil && !IsNil(o.BinlogCacheSize) {
		return true
	}

	return false
}

// SetBinlogCacheSize gets a reference to the given string and assigns it to the BinlogCacheSize field.
func (o *ConfigParametersMysql) SetBinlogCacheSize(v string) {
	o.BinlogCacheSize = &v
}

// GetBinlogGroupCommitSyncDelay returns the BinlogGroupCommitSyncDelay field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetBinlogGroupCommitSyncDelay() string {
	if o == nil || IsNil(o.BinlogGroupCommitSyncDelay) {
		var ret string
		return ret
	}
	return *o.BinlogGroupCommitSyncDelay
}

// GetBinlogGroupCommitSyncDelayOk returns a tuple with the BinlogGroupCommitSyncDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetBinlogGroupCommitSyncDelayOk() (*string, bool) {
	if o == nil || IsNil(o.BinlogGroupCommitSyncDelay) {
		return nil, false
	}
	return o.BinlogGroupCommitSyncDelay, true
}

// HasBinlogGroupCommitSyncDelay returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasBinlogGroupCommitSyncDelay() bool {
	if o != nil && !IsNil(o.BinlogGroupCommitSyncDelay) {
		return true
	}

	return false
}

// SetBinlogGroupCommitSyncDelay gets a reference to the given string and assigns it to the BinlogGroupCommitSyncDelay field.
func (o *ConfigParametersMysql) SetBinlogGroupCommitSyncDelay(v string) {
	o.BinlogGroupCommitSyncDelay = &v
}

// GetBinlogRowImage returns the BinlogRowImage field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetBinlogRowImage() string {
	if o == nil || IsNil(o.BinlogRowImage) {
		var ret string
		return ret
	}
	return *o.BinlogRowImage
}

// GetBinlogRowImageOk returns a tuple with the BinlogRowImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetBinlogRowImageOk() (*string, bool) {
	if o == nil || IsNil(o.BinlogRowImage) {
		return nil, false
	}
	return o.BinlogRowImage, true
}

// HasBinlogRowImage returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasBinlogRowImage() bool {
	if o != nil && !IsNil(o.BinlogRowImage) {
		return true
	}

	return false
}

// SetBinlogRowImage gets a reference to the given string and assigns it to the BinlogRowImage field.
func (o *ConfigParametersMysql) SetBinlogRowImage(v string) {
	o.BinlogRowImage = &v
}

// GetBinlogRowsQueryLogEvents returns the BinlogRowsQueryLogEvents field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetBinlogRowsQueryLogEvents() string {
	if o == nil || IsNil(o.BinlogRowsQueryLogEvents) {
		var ret string
		return ret
	}
	return *o.BinlogRowsQueryLogEvents
}

// GetBinlogRowsQueryLogEventsOk returns a tuple with the BinlogRowsQueryLogEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetBinlogRowsQueryLogEventsOk() (*string, bool) {
	if o == nil || IsNil(o.BinlogRowsQueryLogEvents) {
		return nil, false
	}
	return o.BinlogRowsQueryLogEvents, true
}

// HasBinlogRowsQueryLogEvents returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasBinlogRowsQueryLogEvents() bool {
	if o != nil && !IsNil(o.BinlogRowsQueryLogEvents) {
		return true
	}

	return false
}

// SetBinlogRowsQueryLogEvents gets a reference to the given string and assigns it to the BinlogRowsQueryLogEvents field.
func (o *ConfigParametersMysql) SetBinlogRowsQueryLogEvents(v string) {
	o.BinlogRowsQueryLogEvents = &v
}

// GetCharacterSetServer returns the CharacterSetServer field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetCharacterSetServer() string {
	if o == nil || IsNil(o.CharacterSetServer) {
		var ret string
		return ret
	}
	return *o.CharacterSetServer
}

// GetCharacterSetServerOk returns a tuple with the CharacterSetServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetCharacterSetServerOk() (*string, bool) {
	if o == nil || IsNil(o.CharacterSetServer) {
		return nil, false
	}
	return o.CharacterSetServer, true
}

// HasCharacterSetServer returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasCharacterSetServer() bool {
	if o != nil && !IsNil(o.CharacterSetServer) {
		return true
	}

	return false
}

// SetCharacterSetServer gets a reference to the given string and assigns it to the CharacterSetServer field.
func (o *ConfigParametersMysql) SetCharacterSetServer(v string) {
	o.CharacterSetServer = &v
}

// GetExplicitDefaultsForTimestamp returns the ExplicitDefaultsForTimestamp field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetExplicitDefaultsForTimestamp() string {
	if o == nil || IsNil(o.ExplicitDefaultsForTimestamp) {
		var ret string
		return ret
	}
	return *o.ExplicitDefaultsForTimestamp
}

// GetExplicitDefaultsForTimestampOk returns a tuple with the ExplicitDefaultsForTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetExplicitDefaultsForTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.ExplicitDefaultsForTimestamp) {
		return nil, false
	}
	return o.ExplicitDefaultsForTimestamp, true
}

// HasExplicitDefaultsForTimestamp returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasExplicitDefaultsForTimestamp() bool {
	if o != nil && !IsNil(o.ExplicitDefaultsForTimestamp) {
		return true
	}

	return false
}

// SetExplicitDefaultsForTimestamp gets a reference to the given string and assigns it to the ExplicitDefaultsForTimestamp field.
func (o *ConfigParametersMysql) SetExplicitDefaultsForTimestamp(v string) {
	o.ExplicitDefaultsForTimestamp = &v
}

// GetGroupConcatMaxLen returns the GroupConcatMaxLen field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetGroupConcatMaxLen() string {
	if o == nil || IsNil(o.GroupConcatMaxLen) {
		var ret string
		return ret
	}
	return *o.GroupConcatMaxLen
}

// GetGroupConcatMaxLenOk returns a tuple with the GroupConcatMaxLen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetGroupConcatMaxLenOk() (*string, bool) {
	if o == nil || IsNil(o.GroupConcatMaxLen) {
		return nil, false
	}
	return o.GroupConcatMaxLen, true
}

// HasGroupConcatMaxLen returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasGroupConcatMaxLen() bool {
	if o != nil && !IsNil(o.GroupConcatMaxLen) {
		return true
	}

	return false
}

// SetGroupConcatMaxLen gets a reference to the given string and assigns it to the GroupConcatMaxLen field.
func (o *ConfigParametersMysql) SetGroupConcatMaxLen(v string) {
	o.GroupConcatMaxLen = &v
}

// GetInnodbAdaptiveHashIndex returns the InnodbAdaptiveHashIndex field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetInnodbAdaptiveHashIndex() string {
	if o == nil || IsNil(o.InnodbAdaptiveHashIndex) {
		var ret string
		return ret
	}
	return *o.InnodbAdaptiveHashIndex
}

// GetInnodbAdaptiveHashIndexOk returns a tuple with the InnodbAdaptiveHashIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetInnodbAdaptiveHashIndexOk() (*string, bool) {
	if o == nil || IsNil(o.InnodbAdaptiveHashIndex) {
		return nil, false
	}
	return o.InnodbAdaptiveHashIndex, true
}

// HasInnodbAdaptiveHashIndex returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasInnodbAdaptiveHashIndex() bool {
	if o != nil && !IsNil(o.InnodbAdaptiveHashIndex) {
		return true
	}

	return false
}

// SetInnodbAdaptiveHashIndex gets a reference to the given string and assigns it to the InnodbAdaptiveHashIndex field.
func (o *ConfigParametersMysql) SetInnodbAdaptiveHashIndex(v string) {
	o.InnodbAdaptiveHashIndex = &v
}

// GetInnodbLockWaitTimeout returns the InnodbLockWaitTimeout field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetInnodbLockWaitTimeout() string {
	if o == nil || IsNil(o.InnodbLockWaitTimeout) {
		var ret string
		return ret
	}
	return *o.InnodbLockWaitTimeout
}

// GetInnodbLockWaitTimeoutOk returns a tuple with the InnodbLockWaitTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetInnodbLockWaitTimeoutOk() (*string, bool) {
	if o == nil || IsNil(o.InnodbLockWaitTimeout) {
		return nil, false
	}
	return o.InnodbLockWaitTimeout, true
}

// HasInnodbLockWaitTimeout returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasInnodbLockWaitTimeout() bool {
	if o != nil && !IsNil(o.InnodbLockWaitTimeout) {
		return true
	}

	return false
}

// SetInnodbLockWaitTimeout gets a reference to the given string and assigns it to the InnodbLockWaitTimeout field.
func (o *ConfigParametersMysql) SetInnodbLockWaitTimeout(v string) {
	o.InnodbLockWaitTimeout = &v
}

// GetInnodbNumaInterleave returns the InnodbNumaInterleave field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetInnodbNumaInterleave() string {
	if o == nil || IsNil(o.InnodbNumaInterleave) {
		var ret string
		return ret
	}
	return *o.InnodbNumaInterleave
}

// GetInnodbNumaInterleaveOk returns a tuple with the InnodbNumaInterleave field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetInnodbNumaInterleaveOk() (*string, bool) {
	if o == nil || IsNil(o.InnodbNumaInterleave) {
		return nil, false
	}
	return o.InnodbNumaInterleave, true
}

// HasInnodbNumaInterleave returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasInnodbNumaInterleave() bool {
	if o != nil && !IsNil(o.InnodbNumaInterleave) {
		return true
	}

	return false
}

// SetInnodbNumaInterleave gets a reference to the given string and assigns it to the InnodbNumaInterleave field.
func (o *ConfigParametersMysql) SetInnodbNumaInterleave(v string) {
	o.InnodbNumaInterleave = &v
}

// GetNetReadTimeout returns the NetReadTimeout field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetNetReadTimeout() string {
	if o == nil || IsNil(o.NetReadTimeout) {
		var ret string
		return ret
	}
	return *o.NetReadTimeout
}

// GetNetReadTimeoutOk returns a tuple with the NetReadTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetNetReadTimeoutOk() (*string, bool) {
	if o == nil || IsNil(o.NetReadTimeout) {
		return nil, false
	}
	return o.NetReadTimeout, true
}

// HasNetReadTimeout returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasNetReadTimeout() bool {
	if o != nil && !IsNil(o.NetReadTimeout) {
		return true
	}

	return false
}

// SetNetReadTimeout gets a reference to the given string and assigns it to the NetReadTimeout field.
func (o *ConfigParametersMysql) SetNetReadTimeout(v string) {
	o.NetReadTimeout = &v
}

// GetNetWriteTimeout returns the NetWriteTimeout field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetNetWriteTimeout() string {
	if o == nil || IsNil(o.NetWriteTimeout) {
		var ret string
		return ret
	}
	return *o.NetWriteTimeout
}

// GetNetWriteTimeoutOk returns a tuple with the NetWriteTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetNetWriteTimeoutOk() (*string, bool) {
	if o == nil || IsNil(o.NetWriteTimeout) {
		return nil, false
	}
	return o.NetWriteTimeout, true
}

// HasNetWriteTimeout returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasNetWriteTimeout() bool {
	if o != nil && !IsNil(o.NetWriteTimeout) {
		return true
	}

	return false
}

// SetNetWriteTimeout gets a reference to the given string and assigns it to the NetWriteTimeout field.
func (o *ConfigParametersMysql) SetNetWriteTimeout(v string) {
	o.NetWriteTimeout = &v
}

// GetRegexpTimeLimit returns the RegexpTimeLimit field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetRegexpTimeLimit() string {
	if o == nil || IsNil(o.RegexpTimeLimit) {
		var ret string
		return ret
	}
	return *o.RegexpTimeLimit
}

// GetRegexpTimeLimitOk returns a tuple with the RegexpTimeLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetRegexpTimeLimitOk() (*string, bool) {
	if o == nil || IsNil(o.RegexpTimeLimit) {
		return nil, false
	}
	return o.RegexpTimeLimit, true
}

// HasRegexpTimeLimit returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasRegexpTimeLimit() bool {
	if o != nil && !IsNil(o.RegexpTimeLimit) {
		return true
	}

	return false
}

// SetRegexpTimeLimit gets a reference to the given string and assigns it to the RegexpTimeLimit field.
func (o *ConfigParametersMysql) SetRegexpTimeLimit(v string) {
	o.RegexpTimeLimit = &v
}

// GetSyncBinlog returns the SyncBinlog field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetSyncBinlog() string {
	if o == nil || IsNil(o.SyncBinlog) {
		var ret string
		return ret
	}
	return *o.SyncBinlog
}

// GetSyncBinlogOk returns a tuple with the SyncBinlog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetSyncBinlogOk() (*string, bool) {
	if o == nil || IsNil(o.SyncBinlog) {
		return nil, false
	}
	return o.SyncBinlog, true
}

// HasSyncBinlog returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasSyncBinlog() bool {
	if o != nil && !IsNil(o.SyncBinlog) {
		return true
	}

	return false
}

// SetSyncBinlog gets a reference to the given string and assigns it to the SyncBinlog field.
func (o *ConfigParametersMysql) SetSyncBinlog(v string) {
	o.SyncBinlog = &v
}

// GetTableDefinitionCache returns the TableDefinitionCache field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetTableDefinitionCache() string {
	if o == nil || IsNil(o.TableDefinitionCache) {
		var ret string
		return ret
	}
	return *o.TableDefinitionCache
}

// GetTableDefinitionCacheOk returns a tuple with the TableDefinitionCache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetTableDefinitionCacheOk() (*string, bool) {
	if o == nil || IsNil(o.TableDefinitionCache) {
		return nil, false
	}
	return o.TableDefinitionCache, true
}

// HasTableDefinitionCache returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasTableDefinitionCache() bool {
	if o != nil && !IsNil(o.TableDefinitionCache) {
		return true
	}

	return false
}

// SetTableDefinitionCache gets a reference to the given string and assigns it to the TableDefinitionCache field.
func (o *ConfigParametersMysql) SetTableDefinitionCache(v string) {
	o.TableDefinitionCache = &v
}

// GetLogBinTrustFunctionCreators returns the LogBinTrustFunctionCreators field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetLogBinTrustFunctionCreators() string {
	if o == nil || IsNil(o.LogBinTrustFunctionCreators) {
		var ret string
		return ret
	}
	return *o.LogBinTrustFunctionCreators
}

// GetLogBinTrustFunctionCreatorsOk returns a tuple with the LogBinTrustFunctionCreators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetLogBinTrustFunctionCreatorsOk() (*string, bool) {
	if o == nil || IsNil(o.LogBinTrustFunctionCreators) {
		return nil, false
	}
	return o.LogBinTrustFunctionCreators, true
}

// HasLogBinTrustFunctionCreators returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasLogBinTrustFunctionCreators() bool {
	if o != nil && !IsNil(o.LogBinTrustFunctionCreators) {
		return true
	}

	return false
}

// SetLogBinTrustFunctionCreators gets a reference to the given string and assigns it to the LogBinTrustFunctionCreators field.
func (o *ConfigParametersMysql) SetLogBinTrustFunctionCreators(v string) {
	o.LogBinTrustFunctionCreators = &v
}

// GetSkipNameResolve returns the SkipNameResolve field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetSkipNameResolve() string {
	if o == nil || IsNil(o.SkipNameResolve) {
		var ret string
		return ret
	}
	return *o.SkipNameResolve
}

// GetSkipNameResolveOk returns a tuple with the SkipNameResolve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetSkipNameResolveOk() (*string, bool) {
	if o == nil || IsNil(o.SkipNameResolve) {
		return nil, false
	}
	return o.SkipNameResolve, true
}

// HasSkipNameResolve returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasSkipNameResolve() bool {
	if o != nil && !IsNil(o.SkipNameResolve) {
		return true
	}

	return false
}

// SetSkipNameResolve gets a reference to the given string and assigns it to the SkipNameResolve field.
func (o *ConfigParametersMysql) SetSkipNameResolve(v string) {
	o.SkipNameResolve = &v
}

// GetInnodbRedoLogCapacity returns the InnodbRedoLogCapacity field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetInnodbRedoLogCapacity() string {
	if o == nil || IsNil(o.InnodbRedoLogCapacity) {
		var ret string
		return ret
	}
	return *o.InnodbRedoLogCapacity
}

// GetInnodbRedoLogCapacityOk returns a tuple with the InnodbRedoLogCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetInnodbRedoLogCapacityOk() (*string, bool) {
	if o == nil || IsNil(o.InnodbRedoLogCapacity) {
		return nil, false
	}
	return o.InnodbRedoLogCapacity, true
}

// HasInnodbRedoLogCapacity returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasInnodbRedoLogCapacity() bool {
	if o != nil && !IsNil(o.InnodbRedoLogCapacity) {
		return true
	}

	return false
}

// SetInnodbRedoLogCapacity gets a reference to the given string and assigns it to the InnodbRedoLogCapacity field.
func (o *ConfigParametersMysql) SetInnodbRedoLogCapacity(v string) {
	o.InnodbRedoLogCapacity = &v
}

// GetWaitTimeout returns the WaitTimeout field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetWaitTimeout() string {
	if o == nil || IsNil(o.WaitTimeout) {
		var ret string
		return ret
	}
	return *o.WaitTimeout
}

// GetWaitTimeoutOk returns a tuple with the WaitTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetWaitTimeoutOk() (*string, bool) {
	if o == nil || IsNil(o.WaitTimeout) {
		return nil, false
	}
	return o.WaitTimeout, true
}

// HasWaitTimeout returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasWaitTimeout() bool {
	if o != nil && !IsNil(o.WaitTimeout) {
		return true
	}

	return false
}

// SetWaitTimeout gets a reference to the given string and assigns it to the WaitTimeout field.
func (o *ConfigParametersMysql) SetWaitTimeout(v string) {
	o.WaitTimeout = &v
}

// GetInteractiveTimeout returns the InteractiveTimeout field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetInteractiveTimeout() string {
	if o == nil || IsNil(o.InteractiveTimeout) {
		var ret string
		return ret
	}
	return *o.InteractiveTimeout
}

// GetInteractiveTimeoutOk returns a tuple with the InteractiveTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetInteractiveTimeoutOk() (*string, bool) {
	if o == nil || IsNil(o.InteractiveTimeout) {
		return nil, false
	}
	return o.InteractiveTimeout, true
}

// HasInteractiveTimeout returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasInteractiveTimeout() bool {
	if o != nil && !IsNil(o.InteractiveTimeout) {
		return true
	}

	return false
}

// SetInteractiveTimeout gets a reference to the given string and assigns it to the InteractiveTimeout field.
func (o *ConfigParametersMysql) SetInteractiveTimeout(v string) {
	o.InteractiveTimeout = &v
}

// GetDefaultTimeZone returns the DefaultTimeZone field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetDefaultTimeZone() string {
	if o == nil || IsNil(o.DefaultTimeZone) {
		var ret string
		return ret
	}
	return *o.DefaultTimeZone
}

// GetDefaultTimeZoneOk returns a tuple with the DefaultTimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetDefaultTimeZoneOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultTimeZone) {
		return nil, false
	}
	return o.DefaultTimeZone, true
}

// HasDefaultTimeZone returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasDefaultTimeZone() bool {
	if o != nil && !IsNil(o.DefaultTimeZone) {
		return true
	}

	return false
}

// SetDefaultTimeZone gets a reference to the given string and assigns it to the DefaultTimeZone field.
func (o *ConfigParametersMysql) SetDefaultTimeZone(v string) {
	o.DefaultTimeZone = &v
}

// GetPxcStrictMode returns the PxcStrictMode field value if set, zero value otherwise.
func (o *ConfigParametersMysql) GetPxcStrictMode() string {
	if o == nil || IsNil(o.PxcStrictMode) {
		var ret string
		return ret
	}
	return *o.PxcStrictMode
}

// GetPxcStrictModeOk returns a tuple with the PxcStrictMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigParametersMysql) GetPxcStrictModeOk() (*string, bool) {
	if o == nil || IsNil(o.PxcStrictMode) {
		return nil, false
	}
	return o.PxcStrictMode, true
}

// HasPxcStrictMode returns a boolean if a field has been set.
func (o *ConfigParametersMysql) HasPxcStrictMode() bool {
	if o != nil && !IsNil(o.PxcStrictMode) {
		return true
	}

	return false
}

// SetPxcStrictMode gets a reference to the given string and assigns it to the PxcStrictMode field.
func (o *ConfigParametersMysql) SetPxcStrictMode(v string) {
	o.PxcStrictMode = &v
}

func (o ConfigParametersMysql) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigParametersMysql) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.JoinBufferSize) {
		toSerialize["join_buffer_size"] = o.JoinBufferSize
	}
	if !IsNil(o.MaxConnections) {
		toSerialize["max_connections"] = o.MaxConnections
	}
	if !IsNil(o.SortBufferSize) {
		toSerialize["sort_buffer_size"] = o.SortBufferSize
	}
	if !IsNil(o.ThreadCacheSize) {
		toSerialize["thread_cache_size"] = o.ThreadCacheSize
	}
	if !IsNil(o.InnodbBufferPoolSize) {
		toSerialize["innodb_buffer_pool_size"] = o.InnodbBufferPoolSize
	}
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
	if !IsNil(o.InnodbLogFileSize) {
		toSerialize["innodb_log_file_size"] = o.InnodbLogFileSize
	}
	if !IsNil(o.MaxAllowedPacket) {
		toSerialize["max_allowed_packet"] = o.MaxAllowedPacket
	}
	if !IsNil(o.MaxHeapTableSize) {
		toSerialize["max_heap_table_size"] = o.MaxHeapTableSize
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
	if !IsNil(o.InnodbFlushLogAtTrxCommit) {
		toSerialize["innodb_flush_log_at_trx_commit"] = o.InnodbFlushLogAtTrxCommit
	}
	if !IsNil(o.TransactionIsolation) {
		toSerialize["transaction_isolation"] = o.TransactionIsolation
	}
	if !IsNil(o.LongQueryTime) {
		toSerialize["long_query_time"] = o.LongQueryTime
	}
	if !IsNil(o.TmpTableSize) {
		toSerialize["tmp_table_size"] = o.TmpTableSize
	}
	if !IsNil(o.TableOpenCache) {
		toSerialize["table_open_cache"] = o.TableOpenCache
	}
	if !IsNil(o.TableOpenCacheInstances) {
		toSerialize["table_open_cache_instances"] = o.TableOpenCacheInstances
	}
	if !IsNil(o.InnodbFlushMethod) {
		toSerialize["innodb_flush_method"] = o.InnodbFlushMethod
	}
	if !IsNil(o.InnodbStrictMode) {
		toSerialize["innodb_strict_mode"] = o.InnodbStrictMode
	}
	if !IsNil(o.SlowQueryLog) {
		toSerialize["slow_query_log"] = o.SlowQueryLog
	}
	if !IsNil(o.BinlogCacheSize) {
		toSerialize["binlog_cache_size"] = o.BinlogCacheSize
	}
	if !IsNil(o.BinlogGroupCommitSyncDelay) {
		toSerialize["binlog_group_commit_sync_delay"] = o.BinlogGroupCommitSyncDelay
	}
	if !IsNil(o.BinlogRowImage) {
		toSerialize["binlog_row_image"] = o.BinlogRowImage
	}
	if !IsNil(o.BinlogRowsQueryLogEvents) {
		toSerialize["binlog_rows_query_log_events"] = o.BinlogRowsQueryLogEvents
	}
	if !IsNil(o.CharacterSetServer) {
		toSerialize["character_set_server"] = o.CharacterSetServer
	}
	if !IsNil(o.ExplicitDefaultsForTimestamp) {
		toSerialize["explicit_defaults_for_timestamp"] = o.ExplicitDefaultsForTimestamp
	}
	if !IsNil(o.GroupConcatMaxLen) {
		toSerialize["group_concat_max_len"] = o.GroupConcatMaxLen
	}
	if !IsNil(o.InnodbAdaptiveHashIndex) {
		toSerialize["innodb_adaptive_hash_index"] = o.InnodbAdaptiveHashIndex
	}
	if !IsNil(o.InnodbLockWaitTimeout) {
		toSerialize["innodb_lock_wait_timeout"] = o.InnodbLockWaitTimeout
	}
	if !IsNil(o.InnodbNumaInterleave) {
		toSerialize["innodb_numa_interleave"] = o.InnodbNumaInterleave
	}
	if !IsNil(o.NetReadTimeout) {
		toSerialize["net_read_timeout"] = o.NetReadTimeout
	}
	if !IsNil(o.NetWriteTimeout) {
		toSerialize["net_write_timeout"] = o.NetWriteTimeout
	}
	if !IsNil(o.RegexpTimeLimit) {
		toSerialize["regexp_time_limit"] = o.RegexpTimeLimit
	}
	if !IsNil(o.SyncBinlog) {
		toSerialize["sync_binlog"] = o.SyncBinlog
	}
	if !IsNil(o.TableDefinitionCache) {
		toSerialize["table_definition_cache"] = o.TableDefinitionCache
	}
	if !IsNil(o.LogBinTrustFunctionCreators) {
		toSerialize["log_bin_trust_function_creators"] = o.LogBinTrustFunctionCreators
	}
	if !IsNil(o.SkipNameResolve) {
		toSerialize["skip_name_resolve"] = o.SkipNameResolve
	}
	if !IsNil(o.InnodbRedoLogCapacity) {
		toSerialize["innodb_redo_log_capacity"] = o.InnodbRedoLogCapacity
	}
	if !IsNil(o.WaitTimeout) {
		toSerialize["wait_timeout"] = o.WaitTimeout
	}
	if !IsNil(o.InteractiveTimeout) {
		toSerialize["interactive_timeout"] = o.InteractiveTimeout
	}
	if !IsNil(o.DefaultTimeZone) {
		toSerialize["default-time-zone"] = o.DefaultTimeZone
	}
	if !IsNil(o.PxcStrictMode) {
		toSerialize["pxc_strict_mode"] = o.PxcStrictMode
	}
	return toSerialize, nil
}

type NullableConfigParametersMysql struct {
	value *ConfigParametersMysql
	isSet bool
}

func (v NullableConfigParametersMysql) Get() *ConfigParametersMysql {
	return v.value
}

func (v *NullableConfigParametersMysql) Set(val *ConfigParametersMysql) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigParametersMysql) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigParametersMysql) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigParametersMysql(val *ConfigParametersMysql) *NullableConfigParametersMysql {
	return &NullableConfigParametersMysql{value: val, isSet: true}
}

func (v NullableConfigParametersMysql) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigParametersMysql) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


