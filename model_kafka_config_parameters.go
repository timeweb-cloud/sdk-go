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

// checks if the KafkaConfigParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KafkaConfigParameters{}

// KafkaConfigParameters Настройки топика Kafka. Все значения возвращаются в виде строк. Не заданные явно параметры возвращаются со значениями по умолчанию.
type KafkaConfigParameters struct {
	// Количество партиций топика.
	Partitions *string `json:"partitions,omitempty"`
	// Политика очистки старых сегментов лога: `delete` — удалять, `compact` — уплотнять.
	CleanupPolicy *string `json:"cleanup_policy,omitempty"`
	// Тип сжатия сообщений в топике.
	CompressionType *string `json:"compression_type,omitempty"`
	// Время (в мс) хранения меток удаления для уплотняемых топиков.
	DeleteRetentionMs *string `json:"delete_retention_ms,omitempty"`
	// Задержка (в мс) перед удалением файла из файловой системы.
	FileDeleteDelayMs *string `json:"file_delete_delay_ms,omitempty"`
	// Количество сообщений, после которого данные принудительно сбрасываются на диск.
	FlushMessages *string `json:"flush_messages,omitempty"`
	// Интервал (в мс), после которого данные принудительно сбрасываются на диск.
	FlushMs *string `json:"flush_ms,omitempty"`
	// Интервал (в байтах), с которым Kafka добавляет запись в индекс смещений.
	IndexIntervalBytes *string `json:"index_interval_bytes,omitempty"`
	// Минимальное время (в мс), в течение которого сообщение остается неуплотненным.
	MinCompactionLagMs *string `json:"min_compaction_lag_ms,omitempty"`
	// Максимальное время (в мс), в течение которого сообщение может оставаться неуплотненным.
	MaxCompactionLagMs *string `json:"max_compaction_lag_ms,omitempty"`
	// Максимальный размер (в байтах) пакета сообщений.
	MaxMessageBytes *string `json:"max_message_bytes,omitempty"`
	// Версия формата сообщений, в котором Kafka добавляет сообщения в лог.
	MessageFormatVersion *string `json:"message_format_version,omitempty"`
	// Максимально допустимая разница (в мс) между временной меткой сообщения и временем его получения брокером.
	MessageTimestampDifferenceMaxMs *string `json:"message_timestamp_difference_max_ms,omitempty"`
	// Понижение версии формата сообщений для старых клиентов.
	MessageDownconversionEnable *string `json:"message_downconversion_enable,omitempty"`
	// Источник временной метки сообщения: `CreateTime` — время создания сообщения клиентом, `LogAppendTime` — время добавления сообщения в лог брокером.
	MessageTimestampType *string `json:"message_timestamp_type,omitempty"`
	// Доля неуплотненных данных в логе, при которой запускается уплотнение.
	MinCleanableDirtyRatio *string `json:"min_cleanable_dirty_ratio,omitempty"`
	// Минимальное количество синхронизированных реплик, необходимое для подтверждения записи.
	MinInsyncReplicas *string `json:"min_insync_replicas,omitempty"`
	// Предварительное выделение места на диске при создании нового сегмента лога.
	Preallocate *string `json:"preallocate,omitempty"`
	// Максимальный размер (в байтах) партиции топика, после которого старые сегменты удаляются. `-1` — без ограничения.
	RetentionBytes *string `json:"retention_bytes,omitempty"`
	// Время (в мс) хранения сообщений в топике. `-1` — хранить бессрочно.
	RetentionMs *string `json:"retention_ms,omitempty"`
	// Максимальный размер (в байтах) одного сегмента лога.
	SegmentBytes *string `json:"segment_bytes,omitempty"`
	// Максимальный размер (в байтах) индексного файла сегмента лога.
	SegmentIndexBytes *string `json:"segment_index_bytes,omitempty"`
	// Максимальное случайное отклонение (в мс) от времени ротации сегмента.
	SegmentJitterMs *string `json:"segment_jitter_ms,omitempty"`
	// Период (в мс), после которого Kafka создает новый сегмент лога.
	SegmentMs *string `json:"segment_ms,omitempty"`
	// Возможность выбрать лидером партиции реплику, которая не входит в число синхронизированных.
	UncleanLeaderElectionEnable *string `json:"unclean_leader_election_enable,omitempty"`
}

// NewKafkaConfigParameters instantiates a new KafkaConfigParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKafkaConfigParameters() *KafkaConfigParameters {
	this := KafkaConfigParameters{}
	return &this
}

// NewKafkaConfigParametersWithDefaults instantiates a new KafkaConfigParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKafkaConfigParametersWithDefaults() *KafkaConfigParameters {
	this := KafkaConfigParameters{}
	return &this
}

// GetPartitions returns the Partitions field value if set, zero value otherwise.
func (o *KafkaConfigParameters) GetPartitions() string {
	if o == nil || IsNil(o.Partitions) {
		var ret string
		return ret
	}
	return *o.Partitions
}

// GetPartitionsOk returns a tuple with the Partitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaConfigParameters) GetPartitionsOk() (*string, bool) {
	if o == nil || IsNil(o.Partitions) {
		return nil, false
	}
	return o.Partitions, true
}

// HasPartitions returns a boolean if a field has been set.
func (o *KafkaConfigParameters) HasPartitions() bool {
	if o != nil && !IsNil(o.Partitions) {
		return true
	}

	return false
}

// SetPartitions gets a reference to the given string and assigns it to the Partitions field.
func (o *KafkaConfigParameters) SetPartitions(v string) {
	o.Partitions = &v
}

// GetCleanupPolicy returns the CleanupPolicy field value if set, zero value otherwise.
func (o *KafkaConfigParameters) GetCleanupPolicy() string {
	if o == nil || IsNil(o.CleanupPolicy) {
		var ret string
		return ret
	}
	return *o.CleanupPolicy
}

// GetCleanupPolicyOk returns a tuple with the CleanupPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaConfigParameters) GetCleanupPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.CleanupPolicy) {
		return nil, false
	}
	return o.CleanupPolicy, true
}

// HasCleanupPolicy returns a boolean if a field has been set.
func (o *KafkaConfigParameters) HasCleanupPolicy() bool {
	if o != nil && !IsNil(o.CleanupPolicy) {
		return true
	}

	return false
}

// SetCleanupPolicy gets a reference to the given string and assigns it to the CleanupPolicy field.
func (o *KafkaConfigParameters) SetCleanupPolicy(v string) {
	o.CleanupPolicy = &v
}

// GetCompressionType returns the CompressionType field value if set, zero value otherwise.
func (o *KafkaConfigParameters) GetCompressionType() string {
	if o == nil || IsNil(o.CompressionType) {
		var ret string
		return ret
	}
	return *o.CompressionType
}

// GetCompressionTypeOk returns a tuple with the CompressionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaConfigParameters) GetCompressionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CompressionType) {
		return nil, false
	}
	return o.CompressionType, true
}

// HasCompressionType returns a boolean if a field has been set.
func (o *KafkaConfigParameters) HasCompressionType() bool {
	if o != nil && !IsNil(o.CompressionType) {
		return true
	}

	return false
}

// SetCompressionType gets a reference to the given string and assigns it to the CompressionType field.
func (o *KafkaConfigParameters) SetCompressionType(v string) {
	o.CompressionType = &v
}

// GetDeleteRetentionMs returns the DeleteRetentionMs field value if set, zero value otherwise.
func (o *KafkaConfigParameters) GetDeleteRetentionMs() string {
	if o == nil || IsNil(o.DeleteRetentionMs) {
		var ret string
		return ret
	}
	return *o.DeleteRetentionMs
}

// GetDeleteRetentionMsOk returns a tuple with the DeleteRetentionMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaConfigParameters) GetDeleteRetentionMsOk() (*string, bool) {
	if o == nil || IsNil(o.DeleteRetentionMs) {
		return nil, false
	}
	return o.DeleteRetentionMs, true
}

// HasDeleteRetentionMs returns a boolean if a field has been set.
func (o *KafkaConfigParameters) HasDeleteRetentionMs() bool {
	if o != nil && !IsNil(o.DeleteRetentionMs) {
		return true
	}

	return false
}

// SetDeleteRetentionMs gets a reference to the given string and assigns it to the DeleteRetentionMs field.
func (o *KafkaConfigParameters) SetDeleteRetentionMs(v string) {
	o.DeleteRetentionMs = &v
}

// GetFileDeleteDelayMs returns the FileDeleteDelayMs field value if set, zero value otherwise.
func (o *KafkaConfigParameters) GetFileDeleteDelayMs() string {
	if o == nil || IsNil(o.FileDeleteDelayMs) {
		var ret string
		return ret
	}
	return *o.FileDeleteDelayMs
}

// GetFileDeleteDelayMsOk returns a tuple with the FileDeleteDelayMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaConfigParameters) GetFileDeleteDelayMsOk() (*string, bool) {
	if o == nil || IsNil(o.FileDeleteDelayMs) {
		return nil, false
	}
	return o.FileDeleteDelayMs, true
}

// HasFileDeleteDelayMs returns a boolean if a field has been set.
func (o *KafkaConfigParameters) HasFileDeleteDelayMs() bool {
	if o != nil && !IsNil(o.FileDeleteDelayMs) {
		return true
	}

	return false
}

// SetFileDeleteDelayMs gets a reference to the given string and assigns it to the FileDeleteDelayMs field.
func (o *KafkaConfigParameters) SetFileDeleteDelayMs(v string) {
	o.FileDeleteDelayMs = &v
}

// GetFlushMessages returns the FlushMessages field value if set, zero value otherwise.
func (o *KafkaConfigParameters) GetFlushMessages() string {
	if o == nil || IsNil(o.FlushMessages) {
		var ret string
		return ret
	}
	return *o.FlushMessages
}

// GetFlushMessagesOk returns a tuple with the FlushMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaConfigParameters) GetFlushMessagesOk() (*string, bool) {
	if o == nil || IsNil(o.FlushMessages) {
		return nil, false
	}
	return o.FlushMessages, true
}

// HasFlushMessages returns a boolean if a field has been set.
func (o *KafkaConfigParameters) HasFlushMessages() bool {
	if o != nil && !IsNil(o.FlushMessages) {
		return true
	}

	return false
}

// SetFlushMessages gets a reference to the given string and assigns it to the FlushMessages field.
func (o *KafkaConfigParameters) SetFlushMessages(v string) {
	o.FlushMessages = &v
}

// GetFlushMs returns the FlushMs field value if set, zero value otherwise.
func (o *KafkaConfigParameters) GetFlushMs() string {
	if o == nil || IsNil(o.FlushMs) {
		var ret string
		return ret
	}
	return *o.FlushMs
}

// GetFlushMsOk returns a tuple with the FlushMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaConfigParameters) GetFlushMsOk() (*string, bool) {
	if o == nil || IsNil(o.FlushMs) {
		return nil, false
	}
	return o.FlushMs, true
}

// HasFlushMs returns a boolean if a field has been set.
func (o *KafkaConfigParameters) HasFlushMs() bool {
	if o != nil && !IsNil(o.FlushMs) {
		return true
	}

	return false
}

// SetFlushMs gets a reference to the given string and assigns it to the FlushMs field.
func (o *KafkaConfigParameters) SetFlushMs(v string) {
	o.FlushMs = &v
}

// GetIndexIntervalBytes returns the IndexIntervalBytes field value if set, zero value otherwise.
func (o *KafkaConfigParameters) GetIndexIntervalBytes() string {
	if o == nil || IsNil(o.IndexIntervalBytes) {
		var ret string
		return ret
	}
	return *o.IndexIntervalBytes
}

// GetIndexIntervalBytesOk returns a tuple with the IndexIntervalBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaConfigParameters) GetIndexIntervalBytesOk() (*string, bool) {
	if o == nil || IsNil(o.IndexIntervalBytes) {
		return nil, false
	}
	return o.IndexIntervalBytes, true
}

// HasIndexIntervalBytes returns a boolean if a field has been set.
func (o *KafkaConfigParameters) HasIndexIntervalBytes() bool {
	if o != nil && !IsNil(o.IndexIntervalBytes) {
		return true
	}

	return false
}

// SetIndexIntervalBytes gets a reference to the given string and assigns it to the IndexIntervalBytes field.
func (o *KafkaConfigParameters) SetIndexIntervalBytes(v string) {
	o.IndexIntervalBytes = &v
}

// GetMinCompactionLagMs returns the MinCompactionLagMs field value if set, zero value otherwise.
func (o *KafkaConfigParameters) GetMinCompactionLagMs() string {
	if o == nil || IsNil(o.MinCompactionLagMs) {
		var ret string
		return ret
	}
	return *o.MinCompactionLagMs
}

// GetMinCompactionLagMsOk returns a tuple with the MinCompactionLagMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaConfigParameters) GetMinCompactionLagMsOk() (*string, bool) {
	if o == nil || IsNil(o.MinCompactionLagMs) {
		return nil, false
	}
	return o.MinCompactionLagMs, true
}

// HasMinCompactionLagMs returns a boolean if a field has been set.
func (o *KafkaConfigParameters) HasMinCompactionLagMs() bool {
	if o != nil && !IsNil(o.MinCompactionLagMs) {
		return true
	}

	return false
}

// SetMinCompactionLagMs gets a reference to the given string and assigns it to the MinCompactionLagMs field.
func (o *KafkaConfigParameters) SetMinCompactionLagMs(v string) {
	o.MinCompactionLagMs = &v
}

// GetMaxCompactionLagMs returns the MaxCompactionLagMs field value if set, zero value otherwise.
func (o *KafkaConfigParameters) GetMaxCompactionLagMs() string {
	if o == nil || IsNil(o.MaxCompactionLagMs) {
		var ret string
		return ret
	}
	return *o.MaxCompactionLagMs
}

// GetMaxCompactionLagMsOk returns a tuple with the MaxCompactionLagMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaConfigParameters) GetMaxCompactionLagMsOk() (*string, bool) {
	if o == nil || IsNil(o.MaxCompactionLagMs) {
		return nil, false
	}
	return o.MaxCompactionLagMs, true
}

// HasMaxCompactionLagMs returns a boolean if a field has been set.
func (o *KafkaConfigParameters) HasMaxCompactionLagMs() bool {
	if o != nil && !IsNil(o.MaxCompactionLagMs) {
		return true
	}

	return false
}

// SetMaxCompactionLagMs gets a reference to the given string and assigns it to the MaxCompactionLagMs field.
func (o *KafkaConfigParameters) SetMaxCompactionLagMs(v string) {
	o.MaxCompactionLagMs = &v
}

// GetMaxMessageBytes returns the MaxMessageBytes field value if set, zero value otherwise.
func (o *KafkaConfigParameters) GetMaxMessageBytes() string {
	if o == nil || IsNil(o.MaxMessageBytes) {
		var ret string
		return ret
	}
	return *o.MaxMessageBytes
}

// GetMaxMessageBytesOk returns a tuple with the MaxMessageBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaConfigParameters) GetMaxMessageBytesOk() (*string, bool) {
	if o == nil || IsNil(o.MaxMessageBytes) {
		return nil, false
	}
	return o.MaxMessageBytes, true
}

// HasMaxMessageBytes returns a boolean if a field has been set.
func (o *KafkaConfigParameters) HasMaxMessageBytes() bool {
	if o != nil && !IsNil(o.MaxMessageBytes) {
		return true
	}

	return false
}

// SetMaxMessageBytes gets a reference to the given string and assigns it to the MaxMessageBytes field.
func (o *KafkaConfigParameters) SetMaxMessageBytes(v string) {
	o.MaxMessageBytes = &v
}

// GetMessageFormatVersion returns the MessageFormatVersion field value if set, zero value otherwise.
func (o *KafkaConfigParameters) GetMessageFormatVersion() string {
	if o == nil || IsNil(o.MessageFormatVersion) {
		var ret string
		return ret
	}
	return *o.MessageFormatVersion
}

// GetMessageFormatVersionOk returns a tuple with the MessageFormatVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaConfigParameters) GetMessageFormatVersionOk() (*string, bool) {
	if o == nil || IsNil(o.MessageFormatVersion) {
		return nil, false
	}
	return o.MessageFormatVersion, true
}

// HasMessageFormatVersion returns a boolean if a field has been set.
func (o *KafkaConfigParameters) HasMessageFormatVersion() bool {
	if o != nil && !IsNil(o.MessageFormatVersion) {
		return true
	}

	return false
}

// SetMessageFormatVersion gets a reference to the given string and assigns it to the MessageFormatVersion field.
func (o *KafkaConfigParameters) SetMessageFormatVersion(v string) {
	o.MessageFormatVersion = &v
}

// GetMessageTimestampDifferenceMaxMs returns the MessageTimestampDifferenceMaxMs field value if set, zero value otherwise.
func (o *KafkaConfigParameters) GetMessageTimestampDifferenceMaxMs() string {
	if o == nil || IsNil(o.MessageTimestampDifferenceMaxMs) {
		var ret string
		return ret
	}
	return *o.MessageTimestampDifferenceMaxMs
}

// GetMessageTimestampDifferenceMaxMsOk returns a tuple with the MessageTimestampDifferenceMaxMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaConfigParameters) GetMessageTimestampDifferenceMaxMsOk() (*string, bool) {
	if o == nil || IsNil(o.MessageTimestampDifferenceMaxMs) {
		return nil, false
	}
	return o.MessageTimestampDifferenceMaxMs, true
}

// HasMessageTimestampDifferenceMaxMs returns a boolean if a field has been set.
func (o *KafkaConfigParameters) HasMessageTimestampDifferenceMaxMs() bool {
	if o != nil && !IsNil(o.MessageTimestampDifferenceMaxMs) {
		return true
	}

	return false
}

// SetMessageTimestampDifferenceMaxMs gets a reference to the given string and assigns it to the MessageTimestampDifferenceMaxMs field.
func (o *KafkaConfigParameters) SetMessageTimestampDifferenceMaxMs(v string) {
	o.MessageTimestampDifferenceMaxMs = &v
}

// GetMessageDownconversionEnable returns the MessageDownconversionEnable field value if set, zero value otherwise.
func (o *KafkaConfigParameters) GetMessageDownconversionEnable() string {
	if o == nil || IsNil(o.MessageDownconversionEnable) {
		var ret string
		return ret
	}
	return *o.MessageDownconversionEnable
}

// GetMessageDownconversionEnableOk returns a tuple with the MessageDownconversionEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaConfigParameters) GetMessageDownconversionEnableOk() (*string, bool) {
	if o == nil || IsNil(o.MessageDownconversionEnable) {
		return nil, false
	}
	return o.MessageDownconversionEnable, true
}

// HasMessageDownconversionEnable returns a boolean if a field has been set.
func (o *KafkaConfigParameters) HasMessageDownconversionEnable() bool {
	if o != nil && !IsNil(o.MessageDownconversionEnable) {
		return true
	}

	return false
}

// SetMessageDownconversionEnable gets a reference to the given string and assigns it to the MessageDownconversionEnable field.
func (o *KafkaConfigParameters) SetMessageDownconversionEnable(v string) {
	o.MessageDownconversionEnable = &v
}

// GetMessageTimestampType returns the MessageTimestampType field value if set, zero value otherwise.
func (o *KafkaConfigParameters) GetMessageTimestampType() string {
	if o == nil || IsNil(o.MessageTimestampType) {
		var ret string
		return ret
	}
	return *o.MessageTimestampType
}

// GetMessageTimestampTypeOk returns a tuple with the MessageTimestampType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaConfigParameters) GetMessageTimestampTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MessageTimestampType) {
		return nil, false
	}
	return o.MessageTimestampType, true
}

// HasMessageTimestampType returns a boolean if a field has been set.
func (o *KafkaConfigParameters) HasMessageTimestampType() bool {
	if o != nil && !IsNil(o.MessageTimestampType) {
		return true
	}

	return false
}

// SetMessageTimestampType gets a reference to the given string and assigns it to the MessageTimestampType field.
func (o *KafkaConfigParameters) SetMessageTimestampType(v string) {
	o.MessageTimestampType = &v
}

// GetMinCleanableDirtyRatio returns the MinCleanableDirtyRatio field value if set, zero value otherwise.
func (o *KafkaConfigParameters) GetMinCleanableDirtyRatio() string {
	if o == nil || IsNil(o.MinCleanableDirtyRatio) {
		var ret string
		return ret
	}
	return *o.MinCleanableDirtyRatio
}

// GetMinCleanableDirtyRatioOk returns a tuple with the MinCleanableDirtyRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaConfigParameters) GetMinCleanableDirtyRatioOk() (*string, bool) {
	if o == nil || IsNil(o.MinCleanableDirtyRatio) {
		return nil, false
	}
	return o.MinCleanableDirtyRatio, true
}

// HasMinCleanableDirtyRatio returns a boolean if a field has been set.
func (o *KafkaConfigParameters) HasMinCleanableDirtyRatio() bool {
	if o != nil && !IsNil(o.MinCleanableDirtyRatio) {
		return true
	}

	return false
}

// SetMinCleanableDirtyRatio gets a reference to the given string and assigns it to the MinCleanableDirtyRatio field.
func (o *KafkaConfigParameters) SetMinCleanableDirtyRatio(v string) {
	o.MinCleanableDirtyRatio = &v
}

// GetMinInsyncReplicas returns the MinInsyncReplicas field value if set, zero value otherwise.
func (o *KafkaConfigParameters) GetMinInsyncReplicas() string {
	if o == nil || IsNil(o.MinInsyncReplicas) {
		var ret string
		return ret
	}
	return *o.MinInsyncReplicas
}

// GetMinInsyncReplicasOk returns a tuple with the MinInsyncReplicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaConfigParameters) GetMinInsyncReplicasOk() (*string, bool) {
	if o == nil || IsNil(o.MinInsyncReplicas) {
		return nil, false
	}
	return o.MinInsyncReplicas, true
}

// HasMinInsyncReplicas returns a boolean if a field has been set.
func (o *KafkaConfigParameters) HasMinInsyncReplicas() bool {
	if o != nil && !IsNil(o.MinInsyncReplicas) {
		return true
	}

	return false
}

// SetMinInsyncReplicas gets a reference to the given string and assigns it to the MinInsyncReplicas field.
func (o *KafkaConfigParameters) SetMinInsyncReplicas(v string) {
	o.MinInsyncReplicas = &v
}

// GetPreallocate returns the Preallocate field value if set, zero value otherwise.
func (o *KafkaConfigParameters) GetPreallocate() string {
	if o == nil || IsNil(o.Preallocate) {
		var ret string
		return ret
	}
	return *o.Preallocate
}

// GetPreallocateOk returns a tuple with the Preallocate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaConfigParameters) GetPreallocateOk() (*string, bool) {
	if o == nil || IsNil(o.Preallocate) {
		return nil, false
	}
	return o.Preallocate, true
}

// HasPreallocate returns a boolean if a field has been set.
func (o *KafkaConfigParameters) HasPreallocate() bool {
	if o != nil && !IsNil(o.Preallocate) {
		return true
	}

	return false
}

// SetPreallocate gets a reference to the given string and assigns it to the Preallocate field.
func (o *KafkaConfigParameters) SetPreallocate(v string) {
	o.Preallocate = &v
}

// GetRetentionBytes returns the RetentionBytes field value if set, zero value otherwise.
func (o *KafkaConfigParameters) GetRetentionBytes() string {
	if o == nil || IsNil(o.RetentionBytes) {
		var ret string
		return ret
	}
	return *o.RetentionBytes
}

// GetRetentionBytesOk returns a tuple with the RetentionBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaConfigParameters) GetRetentionBytesOk() (*string, bool) {
	if o == nil || IsNil(o.RetentionBytes) {
		return nil, false
	}
	return o.RetentionBytes, true
}

// HasRetentionBytes returns a boolean if a field has been set.
func (o *KafkaConfigParameters) HasRetentionBytes() bool {
	if o != nil && !IsNil(o.RetentionBytes) {
		return true
	}

	return false
}

// SetRetentionBytes gets a reference to the given string and assigns it to the RetentionBytes field.
func (o *KafkaConfigParameters) SetRetentionBytes(v string) {
	o.RetentionBytes = &v
}

// GetRetentionMs returns the RetentionMs field value if set, zero value otherwise.
func (o *KafkaConfigParameters) GetRetentionMs() string {
	if o == nil || IsNil(o.RetentionMs) {
		var ret string
		return ret
	}
	return *o.RetentionMs
}

// GetRetentionMsOk returns a tuple with the RetentionMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaConfigParameters) GetRetentionMsOk() (*string, bool) {
	if o == nil || IsNil(o.RetentionMs) {
		return nil, false
	}
	return o.RetentionMs, true
}

// HasRetentionMs returns a boolean if a field has been set.
func (o *KafkaConfigParameters) HasRetentionMs() bool {
	if o != nil && !IsNil(o.RetentionMs) {
		return true
	}

	return false
}

// SetRetentionMs gets a reference to the given string and assigns it to the RetentionMs field.
func (o *KafkaConfigParameters) SetRetentionMs(v string) {
	o.RetentionMs = &v
}

// GetSegmentBytes returns the SegmentBytes field value if set, zero value otherwise.
func (o *KafkaConfigParameters) GetSegmentBytes() string {
	if o == nil || IsNil(o.SegmentBytes) {
		var ret string
		return ret
	}
	return *o.SegmentBytes
}

// GetSegmentBytesOk returns a tuple with the SegmentBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaConfigParameters) GetSegmentBytesOk() (*string, bool) {
	if o == nil || IsNil(o.SegmentBytes) {
		return nil, false
	}
	return o.SegmentBytes, true
}

// HasSegmentBytes returns a boolean if a field has been set.
func (o *KafkaConfigParameters) HasSegmentBytes() bool {
	if o != nil && !IsNil(o.SegmentBytes) {
		return true
	}

	return false
}

// SetSegmentBytes gets a reference to the given string and assigns it to the SegmentBytes field.
func (o *KafkaConfigParameters) SetSegmentBytes(v string) {
	o.SegmentBytes = &v
}

// GetSegmentIndexBytes returns the SegmentIndexBytes field value if set, zero value otherwise.
func (o *KafkaConfigParameters) GetSegmentIndexBytes() string {
	if o == nil || IsNil(o.SegmentIndexBytes) {
		var ret string
		return ret
	}
	return *o.SegmentIndexBytes
}

// GetSegmentIndexBytesOk returns a tuple with the SegmentIndexBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaConfigParameters) GetSegmentIndexBytesOk() (*string, bool) {
	if o == nil || IsNil(o.SegmentIndexBytes) {
		return nil, false
	}
	return o.SegmentIndexBytes, true
}

// HasSegmentIndexBytes returns a boolean if a field has been set.
func (o *KafkaConfigParameters) HasSegmentIndexBytes() bool {
	if o != nil && !IsNil(o.SegmentIndexBytes) {
		return true
	}

	return false
}

// SetSegmentIndexBytes gets a reference to the given string and assigns it to the SegmentIndexBytes field.
func (o *KafkaConfigParameters) SetSegmentIndexBytes(v string) {
	o.SegmentIndexBytes = &v
}

// GetSegmentJitterMs returns the SegmentJitterMs field value if set, zero value otherwise.
func (o *KafkaConfigParameters) GetSegmentJitterMs() string {
	if o == nil || IsNil(o.SegmentJitterMs) {
		var ret string
		return ret
	}
	return *o.SegmentJitterMs
}

// GetSegmentJitterMsOk returns a tuple with the SegmentJitterMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaConfigParameters) GetSegmentJitterMsOk() (*string, bool) {
	if o == nil || IsNil(o.SegmentJitterMs) {
		return nil, false
	}
	return o.SegmentJitterMs, true
}

// HasSegmentJitterMs returns a boolean if a field has been set.
func (o *KafkaConfigParameters) HasSegmentJitterMs() bool {
	if o != nil && !IsNil(o.SegmentJitterMs) {
		return true
	}

	return false
}

// SetSegmentJitterMs gets a reference to the given string and assigns it to the SegmentJitterMs field.
func (o *KafkaConfigParameters) SetSegmentJitterMs(v string) {
	o.SegmentJitterMs = &v
}

// GetSegmentMs returns the SegmentMs field value if set, zero value otherwise.
func (o *KafkaConfigParameters) GetSegmentMs() string {
	if o == nil || IsNil(o.SegmentMs) {
		var ret string
		return ret
	}
	return *o.SegmentMs
}

// GetSegmentMsOk returns a tuple with the SegmentMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaConfigParameters) GetSegmentMsOk() (*string, bool) {
	if o == nil || IsNil(o.SegmentMs) {
		return nil, false
	}
	return o.SegmentMs, true
}

// HasSegmentMs returns a boolean if a field has been set.
func (o *KafkaConfigParameters) HasSegmentMs() bool {
	if o != nil && !IsNil(o.SegmentMs) {
		return true
	}

	return false
}

// SetSegmentMs gets a reference to the given string and assigns it to the SegmentMs field.
func (o *KafkaConfigParameters) SetSegmentMs(v string) {
	o.SegmentMs = &v
}

// GetUncleanLeaderElectionEnable returns the UncleanLeaderElectionEnable field value if set, zero value otherwise.
func (o *KafkaConfigParameters) GetUncleanLeaderElectionEnable() string {
	if o == nil || IsNil(o.UncleanLeaderElectionEnable) {
		var ret string
		return ret
	}
	return *o.UncleanLeaderElectionEnable
}

// GetUncleanLeaderElectionEnableOk returns a tuple with the UncleanLeaderElectionEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaConfigParameters) GetUncleanLeaderElectionEnableOk() (*string, bool) {
	if o == nil || IsNil(o.UncleanLeaderElectionEnable) {
		return nil, false
	}
	return o.UncleanLeaderElectionEnable, true
}

// HasUncleanLeaderElectionEnable returns a boolean if a field has been set.
func (o *KafkaConfigParameters) HasUncleanLeaderElectionEnable() bool {
	if o != nil && !IsNil(o.UncleanLeaderElectionEnable) {
		return true
	}

	return false
}

// SetUncleanLeaderElectionEnable gets a reference to the given string and assigns it to the UncleanLeaderElectionEnable field.
func (o *KafkaConfigParameters) SetUncleanLeaderElectionEnable(v string) {
	o.UncleanLeaderElectionEnable = &v
}

func (o KafkaConfigParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KafkaConfigParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Partitions) {
		toSerialize["partitions"] = o.Partitions
	}
	if !IsNil(o.CleanupPolicy) {
		toSerialize["cleanup_policy"] = o.CleanupPolicy
	}
	if !IsNil(o.CompressionType) {
		toSerialize["compression_type"] = o.CompressionType
	}
	if !IsNil(o.DeleteRetentionMs) {
		toSerialize["delete_retention_ms"] = o.DeleteRetentionMs
	}
	if !IsNil(o.FileDeleteDelayMs) {
		toSerialize["file_delete_delay_ms"] = o.FileDeleteDelayMs
	}
	if !IsNil(o.FlushMessages) {
		toSerialize["flush_messages"] = o.FlushMessages
	}
	if !IsNil(o.FlushMs) {
		toSerialize["flush_ms"] = o.FlushMs
	}
	if !IsNil(o.IndexIntervalBytes) {
		toSerialize["index_interval_bytes"] = o.IndexIntervalBytes
	}
	if !IsNil(o.MinCompactionLagMs) {
		toSerialize["min_compaction_lag_ms"] = o.MinCompactionLagMs
	}
	if !IsNil(o.MaxCompactionLagMs) {
		toSerialize["max_compaction_lag_ms"] = o.MaxCompactionLagMs
	}
	if !IsNil(o.MaxMessageBytes) {
		toSerialize["max_message_bytes"] = o.MaxMessageBytes
	}
	if !IsNil(o.MessageFormatVersion) {
		toSerialize["message_format_version"] = o.MessageFormatVersion
	}
	if !IsNil(o.MessageTimestampDifferenceMaxMs) {
		toSerialize["message_timestamp_difference_max_ms"] = o.MessageTimestampDifferenceMaxMs
	}
	if !IsNil(o.MessageDownconversionEnable) {
		toSerialize["message_downconversion_enable"] = o.MessageDownconversionEnable
	}
	if !IsNil(o.MessageTimestampType) {
		toSerialize["message_timestamp_type"] = o.MessageTimestampType
	}
	if !IsNil(o.MinCleanableDirtyRatio) {
		toSerialize["min_cleanable_dirty_ratio"] = o.MinCleanableDirtyRatio
	}
	if !IsNil(o.MinInsyncReplicas) {
		toSerialize["min_insync_replicas"] = o.MinInsyncReplicas
	}
	if !IsNil(o.Preallocate) {
		toSerialize["preallocate"] = o.Preallocate
	}
	if !IsNil(o.RetentionBytes) {
		toSerialize["retention_bytes"] = o.RetentionBytes
	}
	if !IsNil(o.RetentionMs) {
		toSerialize["retention_ms"] = o.RetentionMs
	}
	if !IsNil(o.SegmentBytes) {
		toSerialize["segment_bytes"] = o.SegmentBytes
	}
	if !IsNil(o.SegmentIndexBytes) {
		toSerialize["segment_index_bytes"] = o.SegmentIndexBytes
	}
	if !IsNil(o.SegmentJitterMs) {
		toSerialize["segment_jitter_ms"] = o.SegmentJitterMs
	}
	if !IsNil(o.SegmentMs) {
		toSerialize["segment_ms"] = o.SegmentMs
	}
	if !IsNil(o.UncleanLeaderElectionEnable) {
		toSerialize["unclean_leader_election_enable"] = o.UncleanLeaderElectionEnable
	}
	return toSerialize, nil
}

type NullableKafkaConfigParameters struct {
	value *KafkaConfigParameters
	isSet bool
}

func (v NullableKafkaConfigParameters) Get() *KafkaConfigParameters {
	return v.value
}

func (v *NullableKafkaConfigParameters) Set(val *KafkaConfigParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableKafkaConfigParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableKafkaConfigParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKafkaConfigParameters(val *KafkaConfigParameters) *NullableKafkaConfigParameters {
	return &NullableKafkaConfigParameters{value: val, isSet: true}
}

func (v NullableKafkaConfigParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKafkaConfigParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


