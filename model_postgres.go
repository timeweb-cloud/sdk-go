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

// checks if the Postgres type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Postgres{}

// Postgres Параметры PostgreSQL (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`)
type Postgres struct {
	// Максимальное количество одновременных подключений к серверу (`mysql5` | `mysql` | `mysql8_4` | `postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	MaxConnections *string `json:"max_connections,omitempty"`
	// Доля изменения строк таблицы перед запуском автоматического анализа (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	AutovacuumAnalyzeScaleFactor *string `json:"autovacuum_analyze_scale_factor,omitempty"`
	// Максимальное количество процессов autovacuum, которые могут работать одновременно (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	AutovacuumMaxWorkers *string `json:"autovacuum_max_workers,omitempty"`
	// Интервал между запусками процессов autovacuum (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	AutovacuumNaptime *string `json:"autovacuum_naptime,omitempty"`
	// Доля вставленных строк перед запуском vacuum для таблиц с большим количеством вставок (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	AutovacuumVacuumInsertScaleFactor *string `json:"autovacuum_vacuum_insert_scale_factor,omitempty"`
	// Доля измененных или удаленных строк перед запуском autovacuum (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	AutovacuumVacuumScaleFactor *string `json:"autovacuum_vacuum_scale_factor,omitempty"`
	// Объем памяти, используемый одним процессом autovacuum (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	AutovacuumWorkMem *string `json:"autovacuum_work_mem,omitempty"`
	// Интервал между циклами фонового процесса записи страниц (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	BgwriterDelay *string `json:"bgwriter_delay,omitempty"`
	// Максимальное количество страниц, записываемых background writer за один цикл (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	BgwriterLruMaxpages *string `json:"bgwriter_lru_maxpages,omitempty"`
	// Время ожидания блокировки перед проверкой взаимной блокировки (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	DeadlockTimeout *string `json:"deadlock_timeout,omitempty"`
	// Максимальный размер списка ожидающих вставок индекса GIN (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	GinPendingListLimit *string `json:"gin_pending_list_limit,omitempty"`
	// Время ожидания неактивной транзакционной сессии перед завершением соединения (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	IdleInTransactionSessionTimeout *string `json:"idle_in_transaction_session_timeout,omitempty"`
	// Максимальное количество таблиц в JOIN, которые планировщик может переупорядочить (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	JoinCollapseLimit *string `json:"join_collapse_limit,omitempty"`
	// Максимальное время ожидания блокировки перед отменой запроса (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	LockTimeout *string `json:"lock_timeout,omitempty"`
	// Максимальное количество подготовленных транзакций, которые могут существовать одновременно (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	MaxPreparedTransactions *string `json:"max_prepared_transactions,omitempty"`
	// Размер общей памяти, используемой PostgreSQL для буферного кэша (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	SharedBuffers *string `json:"shared_buffers,omitempty"`
	// Минимальное время выполнения запроса, после которого он записывается в журнал (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	LogMinDurationStatement *string `json:"log_min_duration_statement,omitempty"`
	// Размер памяти, используемой для буферизации WAL-записей (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	WalBuffers *string `json:"wal_buffers,omitempty"`
	// Максимальный объем памяти для временных таблиц каждой сессии (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	TempBuffers *string `json:"temp_buffers,omitempty"`
	// Объем памяти, используемый одной операцией сортировки или хеширования (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	WorkMem *string `json:"work_mem,omitempty"`
	// Уровень изоляции транзакций по умолчанию (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	DefaultTransactionIsolation *string `json:"default_transaction_isolation,omitempty"`
	// Оценка объема дискового кэша, доступного планировщику запросов (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	EffectiveCacheSize *string `json:"effective_cache_size,omitempty"`
	// Максимальный размер WAL перед запуском контрольной точки (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	MaxWalSize *string `json:"max_wal_size,omitempty"`
	// Минимальный размер WAL, который сохраняется между контрольными точками (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	MinWalSize *string `json:"min_wal_size,omitempty"`
	// Уровень детализации записи WAL для восстановления и репликации (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	WalLevel *string `json:"wal_level,omitempty"`
	// Максимальное количество слотов репликации, которые могут быть созданы (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	MaxReplicationSlots *string `json:"max_replication_slots,omitempty"`
	// Максимальное количество процессов отправки WAL для репликации (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	MaxWalSenders *string `json:"max_wal_senders,omitempty"`
	// Максимальное количество фоновых процессов PostgreSQL (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	MaxWorkerProcesses *string `json:"max_worker_processes,omitempty"`
	// Максимальное количество процессов логической репликации (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	MaxLogicalReplicationWorkers *string `json:"max_logical_replication_workers,omitempty"`
	// Максимальное количество параллельных процессов для операций обслуживания (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	MaxParallelMaintenanceWorkers *string `json:"max_parallel_maintenance_workers,omitempty"`
	// Максимальное количество параллельных рабочих процессов для запросов (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	MaxParallelWorkers *string `json:"max_parallel_workers,omitempty"`
	// Максимальное количество параллельных рабочих процессов на один Gather-узел (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	MaxParallelWorkersPerGather *string `json:"max_parallel_workers_per_gather,omitempty"`
	// Разрешение использования NULL в массивах PostgreSQL (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	ArrayNulls *string `json:"array_nulls,omitempty"`
	// Количество страниц, после записи которых выполняется принудительная очистка данных на диск серверным процессом (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	BackendFlushAfter *string `json:"backend_flush_after,omitempty"`
	// Управление использованием обратного слеша в строковых литералах (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	BackslashQuote *string `json:"backslash_quote,omitempty"`
	// Количество страниц, после которого background writer выполняет очистку данных на диск (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	BgwriterFlushAfter *string `json:"bgwriter_flush_after,omitempty"`
	// Множитель количества страниц, которые background writer пытается очистить (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	BgwriterLruMultiplier *string `json:"bgwriter_lru_multiplier,omitempty"`
	// Определяет режим транзакций только для чтения по умолчанию (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	DefaultTransactionReadOnly *string `json:"default_transaction_read_only,omitempty"`
	// Разрешение использования Hash Aggregate планировщиком запросов (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	EnableHashagg *string `json:"enable_hashagg,omitempty"`
	// Разрешение использования Hash Join планировщиком запросов (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	EnableHashjoin *string `json:"enable_hashjoin,omitempty"`
	// Разрешение использования инкрементальной сортировки планировщиком (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	EnableIncrementalSort *string `json:"enable_incremental_sort,omitempty"`
	// Разрешение использования обычного индексного сканирования (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	EnableIndexscan *string `json:"enable_indexscan,omitempty"`
	// Разрешение использования index-only scan (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	EnableIndexonlyscan *string `json:"enable_indexonlyscan,omitempty"`
	// Разрешение использования материализации промежуточных результатов запросов (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	EnableMaterial *string `json:"enable_material,omitempty"`
	// Разрешение использования Memoize узлов планировщиком запросов (`postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	EnableMemoize *string `json:"enable_memoize,omitempty"`
	// Разрешение использования Merge Join планировщиком запросов (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	EnableMergejoin *string `json:"enable_mergejoin,omitempty"`
	// Разрешение использования параллельного Append для запросов (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	EnableParallelAppend *string `json:"enable_parallel_append,omitempty"`
	// Разрешение использования параллельных Hash операций (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	EnableParallelHash *string `json:"enable_parallel_hash,omitempty"`
	// Разрешение удаления ненужных разделов таблицы при планировании запроса (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	EnablePartitionPruning *string `json:"enable_partition_pruning,omitempty"`
	// Разрешение выполнения соединений между секционированными таблицами с учетом секций (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	EnablePartitionwiseJoin *string `json:"enable_partitionwise_join,omitempty"`
	// Разрешение выполнения агрегатных операций отдельно для секций таблиц (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	EnablePartitionwiseAggregate *string `json:"enable_partitionwise_aggregate,omitempty"`
	// Разрешение использования последовательного сканирования таблиц планировщиком запросов (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	EnableSeqscan *string `json:"enable_seqscan,omitempty"`
	// Разрешение использования операций сортировки планировщиком запросов (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	EnableSort *string `json:"enable_sort,omitempty"`
	// Разрешение использования TID Scan для поиска строк по физическим идентификаторам (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	EnableTidscan *string `json:"enable_tidscan,omitempty"`
	// Завершение сессии при возникновении ошибки SQL (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	ExitOnError *string `json:"exit_on_error,omitempty"`
	// Максимальное количество элементов FROM, которые планировщик может объединять при оптимизации запросов (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	FromCollapseLimit *string `json:"from_collapse_limit,omitempty"`
	// Включение JIT-компиляции для ускорения выполнения запросов (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	Jit *string `json:"jit,omitempty"`
	// Режим использования кэша планов подготовленных запросов (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	PlanCacheMode *string `json:"plan_cache_mode,omitempty"`
	// Всегда заключать идентификаторы в кавычки при генерации SQL (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	QuoteAllIdentifiers *string `json:"quote_all_identifiers,omitempty"`
	// Использование стандартного поведения строковых литералов SQL (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	StandardConformingStrings *string `json:"standard_conforming_strings,omitempty"`
	// Максимальное время выполнения SQL-запроса перед автоматической отменой (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	StatementTimeout *string `json:"statement_timeout,omitempty"`
	// Часовой пояс сервера PostgreSQL по умолчанию (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	Timezone *string `json:"timezone,omitempty"`
	// Преобразование выражений вида `NULL = NULL` в проверку IS NULL (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	TransformNullEquals *string `json:"transform_null_equals,omitempty"`
	// Количество объектов, которые может блокировать одна транзакция (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	MaxLocksPerTransaction *string `json:"max_locks_per_transaction,omitempty"`
	// Лимит стоимости операций autovacuum перед приостановкой работы (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	AutovacuumVacuumCostLimit *string `json:"autovacuum_vacuum_cost_limit,omitempty"`
	// Максимальный интервал времени между автоматическими контрольными точками (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	CheckpointTimeout *string `json:"checkpoint_timeout,omitempty"`
	// Доля интервала checkpoint, за которую PostgreSQL распределяет запись данных (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	CheckpointCompletionTarget *string `json:"checkpoint_completion_target,omitempty"`
	// Включение сжатия WAL-записей для уменьшения объема журнала (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	WalCompression *string `json:"wal_compression,omitempty"`
	// Оценочная стоимость случайного чтения страницы для планировщика запросов (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	RandomPageCost *string `json:"random_page_cost,omitempty"`
	// Количество параллельных операций ввода-вывода, которые планировщик может учитывать (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	EffectiveIoConcurrency *string `json:"effective_io_concurrency,omitempty"`
	// Включение записи в журнал информации об ожидании блокировок дольше deadlock_timeout (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	LogLockWaits *string `json:"log_lock_waits,omitempty"`
	// Минимальный размер временных файлов, при котором они записываются в журнал (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	LogTempFiles *string `json:"log_temp_files,omitempty"`
	// Включение сбора статистики времени операций ввода-вывода (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	TrackIoTiming *string `json:"track_io_timing,omitempty"`
	// Максимальный объем памяти для операций обслуживания, таких как VACUUM и CREATE INDEX (`postgres` | `postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	MaintenanceWorkMem *string `json:"maintenance_work_mem,omitempty"`
	// Время ожидания неактивной сессии перед автоматическим завершением соединения (`postgres14` | `postgres15` | `postgres16` | `postgres17` | `postgres18`).
	IdleSessionTimeout *string `json:"idle_session_timeout,omitempty"`
	// Метод выполнения операций ввода-вывода PostgreSQL (`postgres18`).
	IoMethod *string `json:"io_method,omitempty"`
	// Количество фоновых процессов для выполнения операций ввода-вывода (`postgres18`).
	IoWorkers *string `json:"io_workers,omitempty"`
}

// NewPostgres instantiates a new Postgres object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostgres() *Postgres {
	this := Postgres{}
	return &this
}

// NewPostgresWithDefaults instantiates a new Postgres object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostgresWithDefaults() *Postgres {
	this := Postgres{}
	return &this
}

// GetMaxConnections returns the MaxConnections field value if set, zero value otherwise.
func (o *Postgres) GetMaxConnections() string {
	if o == nil || IsNil(o.MaxConnections) {
		var ret string
		return ret
	}
	return *o.MaxConnections
}

// GetMaxConnectionsOk returns a tuple with the MaxConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetMaxConnectionsOk() (*string, bool) {
	if o == nil || IsNil(o.MaxConnections) {
		return nil, false
	}
	return o.MaxConnections, true
}

// HasMaxConnections returns a boolean if a field has been set.
func (o *Postgres) HasMaxConnections() bool {
	if o != nil && !IsNil(o.MaxConnections) {
		return true
	}

	return false
}

// SetMaxConnections gets a reference to the given string and assigns it to the MaxConnections field.
func (o *Postgres) SetMaxConnections(v string) {
	o.MaxConnections = &v
}

// GetAutovacuumAnalyzeScaleFactor returns the AutovacuumAnalyzeScaleFactor field value if set, zero value otherwise.
func (o *Postgres) GetAutovacuumAnalyzeScaleFactor() string {
	if o == nil || IsNil(o.AutovacuumAnalyzeScaleFactor) {
		var ret string
		return ret
	}
	return *o.AutovacuumAnalyzeScaleFactor
}

// GetAutovacuumAnalyzeScaleFactorOk returns a tuple with the AutovacuumAnalyzeScaleFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetAutovacuumAnalyzeScaleFactorOk() (*string, bool) {
	if o == nil || IsNil(o.AutovacuumAnalyzeScaleFactor) {
		return nil, false
	}
	return o.AutovacuumAnalyzeScaleFactor, true
}

// HasAutovacuumAnalyzeScaleFactor returns a boolean if a field has been set.
func (o *Postgres) HasAutovacuumAnalyzeScaleFactor() bool {
	if o != nil && !IsNil(o.AutovacuumAnalyzeScaleFactor) {
		return true
	}

	return false
}

// SetAutovacuumAnalyzeScaleFactor gets a reference to the given string and assigns it to the AutovacuumAnalyzeScaleFactor field.
func (o *Postgres) SetAutovacuumAnalyzeScaleFactor(v string) {
	o.AutovacuumAnalyzeScaleFactor = &v
}

// GetAutovacuumMaxWorkers returns the AutovacuumMaxWorkers field value if set, zero value otherwise.
func (o *Postgres) GetAutovacuumMaxWorkers() string {
	if o == nil || IsNil(o.AutovacuumMaxWorkers) {
		var ret string
		return ret
	}
	return *o.AutovacuumMaxWorkers
}

// GetAutovacuumMaxWorkersOk returns a tuple with the AutovacuumMaxWorkers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetAutovacuumMaxWorkersOk() (*string, bool) {
	if o == nil || IsNil(o.AutovacuumMaxWorkers) {
		return nil, false
	}
	return o.AutovacuumMaxWorkers, true
}

// HasAutovacuumMaxWorkers returns a boolean if a field has been set.
func (o *Postgres) HasAutovacuumMaxWorkers() bool {
	if o != nil && !IsNil(o.AutovacuumMaxWorkers) {
		return true
	}

	return false
}

// SetAutovacuumMaxWorkers gets a reference to the given string and assigns it to the AutovacuumMaxWorkers field.
func (o *Postgres) SetAutovacuumMaxWorkers(v string) {
	o.AutovacuumMaxWorkers = &v
}

// GetAutovacuumNaptime returns the AutovacuumNaptime field value if set, zero value otherwise.
func (o *Postgres) GetAutovacuumNaptime() string {
	if o == nil || IsNil(o.AutovacuumNaptime) {
		var ret string
		return ret
	}
	return *o.AutovacuumNaptime
}

// GetAutovacuumNaptimeOk returns a tuple with the AutovacuumNaptime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetAutovacuumNaptimeOk() (*string, bool) {
	if o == nil || IsNil(o.AutovacuumNaptime) {
		return nil, false
	}
	return o.AutovacuumNaptime, true
}

// HasAutovacuumNaptime returns a boolean if a field has been set.
func (o *Postgres) HasAutovacuumNaptime() bool {
	if o != nil && !IsNil(o.AutovacuumNaptime) {
		return true
	}

	return false
}

// SetAutovacuumNaptime gets a reference to the given string and assigns it to the AutovacuumNaptime field.
func (o *Postgres) SetAutovacuumNaptime(v string) {
	o.AutovacuumNaptime = &v
}

// GetAutovacuumVacuumInsertScaleFactor returns the AutovacuumVacuumInsertScaleFactor field value if set, zero value otherwise.
func (o *Postgres) GetAutovacuumVacuumInsertScaleFactor() string {
	if o == nil || IsNil(o.AutovacuumVacuumInsertScaleFactor) {
		var ret string
		return ret
	}
	return *o.AutovacuumVacuumInsertScaleFactor
}

// GetAutovacuumVacuumInsertScaleFactorOk returns a tuple with the AutovacuumVacuumInsertScaleFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetAutovacuumVacuumInsertScaleFactorOk() (*string, bool) {
	if o == nil || IsNil(o.AutovacuumVacuumInsertScaleFactor) {
		return nil, false
	}
	return o.AutovacuumVacuumInsertScaleFactor, true
}

// HasAutovacuumVacuumInsertScaleFactor returns a boolean if a field has been set.
func (o *Postgres) HasAutovacuumVacuumInsertScaleFactor() bool {
	if o != nil && !IsNil(o.AutovacuumVacuumInsertScaleFactor) {
		return true
	}

	return false
}

// SetAutovacuumVacuumInsertScaleFactor gets a reference to the given string and assigns it to the AutovacuumVacuumInsertScaleFactor field.
func (o *Postgres) SetAutovacuumVacuumInsertScaleFactor(v string) {
	o.AutovacuumVacuumInsertScaleFactor = &v
}

// GetAutovacuumVacuumScaleFactor returns the AutovacuumVacuumScaleFactor field value if set, zero value otherwise.
func (o *Postgres) GetAutovacuumVacuumScaleFactor() string {
	if o == nil || IsNil(o.AutovacuumVacuumScaleFactor) {
		var ret string
		return ret
	}
	return *o.AutovacuumVacuumScaleFactor
}

// GetAutovacuumVacuumScaleFactorOk returns a tuple with the AutovacuumVacuumScaleFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetAutovacuumVacuumScaleFactorOk() (*string, bool) {
	if o == nil || IsNil(o.AutovacuumVacuumScaleFactor) {
		return nil, false
	}
	return o.AutovacuumVacuumScaleFactor, true
}

// HasAutovacuumVacuumScaleFactor returns a boolean if a field has been set.
func (o *Postgres) HasAutovacuumVacuumScaleFactor() bool {
	if o != nil && !IsNil(o.AutovacuumVacuumScaleFactor) {
		return true
	}

	return false
}

// SetAutovacuumVacuumScaleFactor gets a reference to the given string and assigns it to the AutovacuumVacuumScaleFactor field.
func (o *Postgres) SetAutovacuumVacuumScaleFactor(v string) {
	o.AutovacuumVacuumScaleFactor = &v
}

// GetAutovacuumWorkMem returns the AutovacuumWorkMem field value if set, zero value otherwise.
func (o *Postgres) GetAutovacuumWorkMem() string {
	if o == nil || IsNil(o.AutovacuumWorkMem) {
		var ret string
		return ret
	}
	return *o.AutovacuumWorkMem
}

// GetAutovacuumWorkMemOk returns a tuple with the AutovacuumWorkMem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetAutovacuumWorkMemOk() (*string, bool) {
	if o == nil || IsNil(o.AutovacuumWorkMem) {
		return nil, false
	}
	return o.AutovacuumWorkMem, true
}

// HasAutovacuumWorkMem returns a boolean if a field has been set.
func (o *Postgres) HasAutovacuumWorkMem() bool {
	if o != nil && !IsNil(o.AutovacuumWorkMem) {
		return true
	}

	return false
}

// SetAutovacuumWorkMem gets a reference to the given string and assigns it to the AutovacuumWorkMem field.
func (o *Postgres) SetAutovacuumWorkMem(v string) {
	o.AutovacuumWorkMem = &v
}

// GetBgwriterDelay returns the BgwriterDelay field value if set, zero value otherwise.
func (o *Postgres) GetBgwriterDelay() string {
	if o == nil || IsNil(o.BgwriterDelay) {
		var ret string
		return ret
	}
	return *o.BgwriterDelay
}

// GetBgwriterDelayOk returns a tuple with the BgwriterDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetBgwriterDelayOk() (*string, bool) {
	if o == nil || IsNil(o.BgwriterDelay) {
		return nil, false
	}
	return o.BgwriterDelay, true
}

// HasBgwriterDelay returns a boolean if a field has been set.
func (o *Postgres) HasBgwriterDelay() bool {
	if o != nil && !IsNil(o.BgwriterDelay) {
		return true
	}

	return false
}

// SetBgwriterDelay gets a reference to the given string and assigns it to the BgwriterDelay field.
func (o *Postgres) SetBgwriterDelay(v string) {
	o.BgwriterDelay = &v
}

// GetBgwriterLruMaxpages returns the BgwriterLruMaxpages field value if set, zero value otherwise.
func (o *Postgres) GetBgwriterLruMaxpages() string {
	if o == nil || IsNil(o.BgwriterLruMaxpages) {
		var ret string
		return ret
	}
	return *o.BgwriterLruMaxpages
}

// GetBgwriterLruMaxpagesOk returns a tuple with the BgwriterLruMaxpages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetBgwriterLruMaxpagesOk() (*string, bool) {
	if o == nil || IsNil(o.BgwriterLruMaxpages) {
		return nil, false
	}
	return o.BgwriterLruMaxpages, true
}

// HasBgwriterLruMaxpages returns a boolean if a field has been set.
func (o *Postgres) HasBgwriterLruMaxpages() bool {
	if o != nil && !IsNil(o.BgwriterLruMaxpages) {
		return true
	}

	return false
}

// SetBgwriterLruMaxpages gets a reference to the given string and assigns it to the BgwriterLruMaxpages field.
func (o *Postgres) SetBgwriterLruMaxpages(v string) {
	o.BgwriterLruMaxpages = &v
}

// GetDeadlockTimeout returns the DeadlockTimeout field value if set, zero value otherwise.
func (o *Postgres) GetDeadlockTimeout() string {
	if o == nil || IsNil(o.DeadlockTimeout) {
		var ret string
		return ret
	}
	return *o.DeadlockTimeout
}

// GetDeadlockTimeoutOk returns a tuple with the DeadlockTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetDeadlockTimeoutOk() (*string, bool) {
	if o == nil || IsNil(o.DeadlockTimeout) {
		return nil, false
	}
	return o.DeadlockTimeout, true
}

// HasDeadlockTimeout returns a boolean if a field has been set.
func (o *Postgres) HasDeadlockTimeout() bool {
	if o != nil && !IsNil(o.DeadlockTimeout) {
		return true
	}

	return false
}

// SetDeadlockTimeout gets a reference to the given string and assigns it to the DeadlockTimeout field.
func (o *Postgres) SetDeadlockTimeout(v string) {
	o.DeadlockTimeout = &v
}

// GetGinPendingListLimit returns the GinPendingListLimit field value if set, zero value otherwise.
func (o *Postgres) GetGinPendingListLimit() string {
	if o == nil || IsNil(o.GinPendingListLimit) {
		var ret string
		return ret
	}
	return *o.GinPendingListLimit
}

// GetGinPendingListLimitOk returns a tuple with the GinPendingListLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetGinPendingListLimitOk() (*string, bool) {
	if o == nil || IsNil(o.GinPendingListLimit) {
		return nil, false
	}
	return o.GinPendingListLimit, true
}

// HasGinPendingListLimit returns a boolean if a field has been set.
func (o *Postgres) HasGinPendingListLimit() bool {
	if o != nil && !IsNil(o.GinPendingListLimit) {
		return true
	}

	return false
}

// SetGinPendingListLimit gets a reference to the given string and assigns it to the GinPendingListLimit field.
func (o *Postgres) SetGinPendingListLimit(v string) {
	o.GinPendingListLimit = &v
}

// GetIdleInTransactionSessionTimeout returns the IdleInTransactionSessionTimeout field value if set, zero value otherwise.
func (o *Postgres) GetIdleInTransactionSessionTimeout() string {
	if o == nil || IsNil(o.IdleInTransactionSessionTimeout) {
		var ret string
		return ret
	}
	return *o.IdleInTransactionSessionTimeout
}

// GetIdleInTransactionSessionTimeoutOk returns a tuple with the IdleInTransactionSessionTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetIdleInTransactionSessionTimeoutOk() (*string, bool) {
	if o == nil || IsNil(o.IdleInTransactionSessionTimeout) {
		return nil, false
	}
	return o.IdleInTransactionSessionTimeout, true
}

// HasIdleInTransactionSessionTimeout returns a boolean if a field has been set.
func (o *Postgres) HasIdleInTransactionSessionTimeout() bool {
	if o != nil && !IsNil(o.IdleInTransactionSessionTimeout) {
		return true
	}

	return false
}

// SetIdleInTransactionSessionTimeout gets a reference to the given string and assigns it to the IdleInTransactionSessionTimeout field.
func (o *Postgres) SetIdleInTransactionSessionTimeout(v string) {
	o.IdleInTransactionSessionTimeout = &v
}

// GetJoinCollapseLimit returns the JoinCollapseLimit field value if set, zero value otherwise.
func (o *Postgres) GetJoinCollapseLimit() string {
	if o == nil || IsNil(o.JoinCollapseLimit) {
		var ret string
		return ret
	}
	return *o.JoinCollapseLimit
}

// GetJoinCollapseLimitOk returns a tuple with the JoinCollapseLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetJoinCollapseLimitOk() (*string, bool) {
	if o == nil || IsNil(o.JoinCollapseLimit) {
		return nil, false
	}
	return o.JoinCollapseLimit, true
}

// HasJoinCollapseLimit returns a boolean if a field has been set.
func (o *Postgres) HasJoinCollapseLimit() bool {
	if o != nil && !IsNil(o.JoinCollapseLimit) {
		return true
	}

	return false
}

// SetJoinCollapseLimit gets a reference to the given string and assigns it to the JoinCollapseLimit field.
func (o *Postgres) SetJoinCollapseLimit(v string) {
	o.JoinCollapseLimit = &v
}

// GetLockTimeout returns the LockTimeout field value if set, zero value otherwise.
func (o *Postgres) GetLockTimeout() string {
	if o == nil || IsNil(o.LockTimeout) {
		var ret string
		return ret
	}
	return *o.LockTimeout
}

// GetLockTimeoutOk returns a tuple with the LockTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetLockTimeoutOk() (*string, bool) {
	if o == nil || IsNil(o.LockTimeout) {
		return nil, false
	}
	return o.LockTimeout, true
}

// HasLockTimeout returns a boolean if a field has been set.
func (o *Postgres) HasLockTimeout() bool {
	if o != nil && !IsNil(o.LockTimeout) {
		return true
	}

	return false
}

// SetLockTimeout gets a reference to the given string and assigns it to the LockTimeout field.
func (o *Postgres) SetLockTimeout(v string) {
	o.LockTimeout = &v
}

// GetMaxPreparedTransactions returns the MaxPreparedTransactions field value if set, zero value otherwise.
func (o *Postgres) GetMaxPreparedTransactions() string {
	if o == nil || IsNil(o.MaxPreparedTransactions) {
		var ret string
		return ret
	}
	return *o.MaxPreparedTransactions
}

// GetMaxPreparedTransactionsOk returns a tuple with the MaxPreparedTransactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetMaxPreparedTransactionsOk() (*string, bool) {
	if o == nil || IsNil(o.MaxPreparedTransactions) {
		return nil, false
	}
	return o.MaxPreparedTransactions, true
}

// HasMaxPreparedTransactions returns a boolean if a field has been set.
func (o *Postgres) HasMaxPreparedTransactions() bool {
	if o != nil && !IsNil(o.MaxPreparedTransactions) {
		return true
	}

	return false
}

// SetMaxPreparedTransactions gets a reference to the given string and assigns it to the MaxPreparedTransactions field.
func (o *Postgres) SetMaxPreparedTransactions(v string) {
	o.MaxPreparedTransactions = &v
}

// GetSharedBuffers returns the SharedBuffers field value if set, zero value otherwise.
func (o *Postgres) GetSharedBuffers() string {
	if o == nil || IsNil(o.SharedBuffers) {
		var ret string
		return ret
	}
	return *o.SharedBuffers
}

// GetSharedBuffersOk returns a tuple with the SharedBuffers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetSharedBuffersOk() (*string, bool) {
	if o == nil || IsNil(o.SharedBuffers) {
		return nil, false
	}
	return o.SharedBuffers, true
}

// HasSharedBuffers returns a boolean if a field has been set.
func (o *Postgres) HasSharedBuffers() bool {
	if o != nil && !IsNil(o.SharedBuffers) {
		return true
	}

	return false
}

// SetSharedBuffers gets a reference to the given string and assigns it to the SharedBuffers field.
func (o *Postgres) SetSharedBuffers(v string) {
	o.SharedBuffers = &v
}

// GetLogMinDurationStatement returns the LogMinDurationStatement field value if set, zero value otherwise.
func (o *Postgres) GetLogMinDurationStatement() string {
	if o == nil || IsNil(o.LogMinDurationStatement) {
		var ret string
		return ret
	}
	return *o.LogMinDurationStatement
}

// GetLogMinDurationStatementOk returns a tuple with the LogMinDurationStatement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetLogMinDurationStatementOk() (*string, bool) {
	if o == nil || IsNil(o.LogMinDurationStatement) {
		return nil, false
	}
	return o.LogMinDurationStatement, true
}

// HasLogMinDurationStatement returns a boolean if a field has been set.
func (o *Postgres) HasLogMinDurationStatement() bool {
	if o != nil && !IsNil(o.LogMinDurationStatement) {
		return true
	}

	return false
}

// SetLogMinDurationStatement gets a reference to the given string and assigns it to the LogMinDurationStatement field.
func (o *Postgres) SetLogMinDurationStatement(v string) {
	o.LogMinDurationStatement = &v
}

// GetWalBuffers returns the WalBuffers field value if set, zero value otherwise.
func (o *Postgres) GetWalBuffers() string {
	if o == nil || IsNil(o.WalBuffers) {
		var ret string
		return ret
	}
	return *o.WalBuffers
}

// GetWalBuffersOk returns a tuple with the WalBuffers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetWalBuffersOk() (*string, bool) {
	if o == nil || IsNil(o.WalBuffers) {
		return nil, false
	}
	return o.WalBuffers, true
}

// HasWalBuffers returns a boolean if a field has been set.
func (o *Postgres) HasWalBuffers() bool {
	if o != nil && !IsNil(o.WalBuffers) {
		return true
	}

	return false
}

// SetWalBuffers gets a reference to the given string and assigns it to the WalBuffers field.
func (o *Postgres) SetWalBuffers(v string) {
	o.WalBuffers = &v
}

// GetTempBuffers returns the TempBuffers field value if set, zero value otherwise.
func (o *Postgres) GetTempBuffers() string {
	if o == nil || IsNil(o.TempBuffers) {
		var ret string
		return ret
	}
	return *o.TempBuffers
}

// GetTempBuffersOk returns a tuple with the TempBuffers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetTempBuffersOk() (*string, bool) {
	if o == nil || IsNil(o.TempBuffers) {
		return nil, false
	}
	return o.TempBuffers, true
}

// HasTempBuffers returns a boolean if a field has been set.
func (o *Postgres) HasTempBuffers() bool {
	if o != nil && !IsNil(o.TempBuffers) {
		return true
	}

	return false
}

// SetTempBuffers gets a reference to the given string and assigns it to the TempBuffers field.
func (o *Postgres) SetTempBuffers(v string) {
	o.TempBuffers = &v
}

// GetWorkMem returns the WorkMem field value if set, zero value otherwise.
func (o *Postgres) GetWorkMem() string {
	if o == nil || IsNil(o.WorkMem) {
		var ret string
		return ret
	}
	return *o.WorkMem
}

// GetWorkMemOk returns a tuple with the WorkMem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetWorkMemOk() (*string, bool) {
	if o == nil || IsNil(o.WorkMem) {
		return nil, false
	}
	return o.WorkMem, true
}

// HasWorkMem returns a boolean if a field has been set.
func (o *Postgres) HasWorkMem() bool {
	if o != nil && !IsNil(o.WorkMem) {
		return true
	}

	return false
}

// SetWorkMem gets a reference to the given string and assigns it to the WorkMem field.
func (o *Postgres) SetWorkMem(v string) {
	o.WorkMem = &v
}

// GetDefaultTransactionIsolation returns the DefaultTransactionIsolation field value if set, zero value otherwise.
func (o *Postgres) GetDefaultTransactionIsolation() string {
	if o == nil || IsNil(o.DefaultTransactionIsolation) {
		var ret string
		return ret
	}
	return *o.DefaultTransactionIsolation
}

// GetDefaultTransactionIsolationOk returns a tuple with the DefaultTransactionIsolation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetDefaultTransactionIsolationOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultTransactionIsolation) {
		return nil, false
	}
	return o.DefaultTransactionIsolation, true
}

// HasDefaultTransactionIsolation returns a boolean if a field has been set.
func (o *Postgres) HasDefaultTransactionIsolation() bool {
	if o != nil && !IsNil(o.DefaultTransactionIsolation) {
		return true
	}

	return false
}

// SetDefaultTransactionIsolation gets a reference to the given string and assigns it to the DefaultTransactionIsolation field.
func (o *Postgres) SetDefaultTransactionIsolation(v string) {
	o.DefaultTransactionIsolation = &v
}

// GetEffectiveCacheSize returns the EffectiveCacheSize field value if set, zero value otherwise.
func (o *Postgres) GetEffectiveCacheSize() string {
	if o == nil || IsNil(o.EffectiveCacheSize) {
		var ret string
		return ret
	}
	return *o.EffectiveCacheSize
}

// GetEffectiveCacheSizeOk returns a tuple with the EffectiveCacheSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetEffectiveCacheSizeOk() (*string, bool) {
	if o == nil || IsNil(o.EffectiveCacheSize) {
		return nil, false
	}
	return o.EffectiveCacheSize, true
}

// HasEffectiveCacheSize returns a boolean if a field has been set.
func (o *Postgres) HasEffectiveCacheSize() bool {
	if o != nil && !IsNil(o.EffectiveCacheSize) {
		return true
	}

	return false
}

// SetEffectiveCacheSize gets a reference to the given string and assigns it to the EffectiveCacheSize field.
func (o *Postgres) SetEffectiveCacheSize(v string) {
	o.EffectiveCacheSize = &v
}

// GetMaxWalSize returns the MaxWalSize field value if set, zero value otherwise.
func (o *Postgres) GetMaxWalSize() string {
	if o == nil || IsNil(o.MaxWalSize) {
		var ret string
		return ret
	}
	return *o.MaxWalSize
}

// GetMaxWalSizeOk returns a tuple with the MaxWalSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetMaxWalSizeOk() (*string, bool) {
	if o == nil || IsNil(o.MaxWalSize) {
		return nil, false
	}
	return o.MaxWalSize, true
}

// HasMaxWalSize returns a boolean if a field has been set.
func (o *Postgres) HasMaxWalSize() bool {
	if o != nil && !IsNil(o.MaxWalSize) {
		return true
	}

	return false
}

// SetMaxWalSize gets a reference to the given string and assigns it to the MaxWalSize field.
func (o *Postgres) SetMaxWalSize(v string) {
	o.MaxWalSize = &v
}

// GetMinWalSize returns the MinWalSize field value if set, zero value otherwise.
func (o *Postgres) GetMinWalSize() string {
	if o == nil || IsNil(o.MinWalSize) {
		var ret string
		return ret
	}
	return *o.MinWalSize
}

// GetMinWalSizeOk returns a tuple with the MinWalSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetMinWalSizeOk() (*string, bool) {
	if o == nil || IsNil(o.MinWalSize) {
		return nil, false
	}
	return o.MinWalSize, true
}

// HasMinWalSize returns a boolean if a field has been set.
func (o *Postgres) HasMinWalSize() bool {
	if o != nil && !IsNil(o.MinWalSize) {
		return true
	}

	return false
}

// SetMinWalSize gets a reference to the given string and assigns it to the MinWalSize field.
func (o *Postgres) SetMinWalSize(v string) {
	o.MinWalSize = &v
}

// GetWalLevel returns the WalLevel field value if set, zero value otherwise.
func (o *Postgres) GetWalLevel() string {
	if o == nil || IsNil(o.WalLevel) {
		var ret string
		return ret
	}
	return *o.WalLevel
}

// GetWalLevelOk returns a tuple with the WalLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetWalLevelOk() (*string, bool) {
	if o == nil || IsNil(o.WalLevel) {
		return nil, false
	}
	return o.WalLevel, true
}

// HasWalLevel returns a boolean if a field has been set.
func (o *Postgres) HasWalLevel() bool {
	if o != nil && !IsNil(o.WalLevel) {
		return true
	}

	return false
}

// SetWalLevel gets a reference to the given string and assigns it to the WalLevel field.
func (o *Postgres) SetWalLevel(v string) {
	o.WalLevel = &v
}

// GetMaxReplicationSlots returns the MaxReplicationSlots field value if set, zero value otherwise.
func (o *Postgres) GetMaxReplicationSlots() string {
	if o == nil || IsNil(o.MaxReplicationSlots) {
		var ret string
		return ret
	}
	return *o.MaxReplicationSlots
}

// GetMaxReplicationSlotsOk returns a tuple with the MaxReplicationSlots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetMaxReplicationSlotsOk() (*string, bool) {
	if o == nil || IsNil(o.MaxReplicationSlots) {
		return nil, false
	}
	return o.MaxReplicationSlots, true
}

// HasMaxReplicationSlots returns a boolean if a field has been set.
func (o *Postgres) HasMaxReplicationSlots() bool {
	if o != nil && !IsNil(o.MaxReplicationSlots) {
		return true
	}

	return false
}

// SetMaxReplicationSlots gets a reference to the given string and assigns it to the MaxReplicationSlots field.
func (o *Postgres) SetMaxReplicationSlots(v string) {
	o.MaxReplicationSlots = &v
}

// GetMaxWalSenders returns the MaxWalSenders field value if set, zero value otherwise.
func (o *Postgres) GetMaxWalSenders() string {
	if o == nil || IsNil(o.MaxWalSenders) {
		var ret string
		return ret
	}
	return *o.MaxWalSenders
}

// GetMaxWalSendersOk returns a tuple with the MaxWalSenders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetMaxWalSendersOk() (*string, bool) {
	if o == nil || IsNil(o.MaxWalSenders) {
		return nil, false
	}
	return o.MaxWalSenders, true
}

// HasMaxWalSenders returns a boolean if a field has been set.
func (o *Postgres) HasMaxWalSenders() bool {
	if o != nil && !IsNil(o.MaxWalSenders) {
		return true
	}

	return false
}

// SetMaxWalSenders gets a reference to the given string and assigns it to the MaxWalSenders field.
func (o *Postgres) SetMaxWalSenders(v string) {
	o.MaxWalSenders = &v
}

// GetMaxWorkerProcesses returns the MaxWorkerProcesses field value if set, zero value otherwise.
func (o *Postgres) GetMaxWorkerProcesses() string {
	if o == nil || IsNil(o.MaxWorkerProcesses) {
		var ret string
		return ret
	}
	return *o.MaxWorkerProcesses
}

// GetMaxWorkerProcessesOk returns a tuple with the MaxWorkerProcesses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetMaxWorkerProcessesOk() (*string, bool) {
	if o == nil || IsNil(o.MaxWorkerProcesses) {
		return nil, false
	}
	return o.MaxWorkerProcesses, true
}

// HasMaxWorkerProcesses returns a boolean if a field has been set.
func (o *Postgres) HasMaxWorkerProcesses() bool {
	if o != nil && !IsNil(o.MaxWorkerProcesses) {
		return true
	}

	return false
}

// SetMaxWorkerProcesses gets a reference to the given string and assigns it to the MaxWorkerProcesses field.
func (o *Postgres) SetMaxWorkerProcesses(v string) {
	o.MaxWorkerProcesses = &v
}

// GetMaxLogicalReplicationWorkers returns the MaxLogicalReplicationWorkers field value if set, zero value otherwise.
func (o *Postgres) GetMaxLogicalReplicationWorkers() string {
	if o == nil || IsNil(o.MaxLogicalReplicationWorkers) {
		var ret string
		return ret
	}
	return *o.MaxLogicalReplicationWorkers
}

// GetMaxLogicalReplicationWorkersOk returns a tuple with the MaxLogicalReplicationWorkers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetMaxLogicalReplicationWorkersOk() (*string, bool) {
	if o == nil || IsNil(o.MaxLogicalReplicationWorkers) {
		return nil, false
	}
	return o.MaxLogicalReplicationWorkers, true
}

// HasMaxLogicalReplicationWorkers returns a boolean if a field has been set.
func (o *Postgres) HasMaxLogicalReplicationWorkers() bool {
	if o != nil && !IsNil(o.MaxLogicalReplicationWorkers) {
		return true
	}

	return false
}

// SetMaxLogicalReplicationWorkers gets a reference to the given string and assigns it to the MaxLogicalReplicationWorkers field.
func (o *Postgres) SetMaxLogicalReplicationWorkers(v string) {
	o.MaxLogicalReplicationWorkers = &v
}

// GetMaxParallelMaintenanceWorkers returns the MaxParallelMaintenanceWorkers field value if set, zero value otherwise.
func (o *Postgres) GetMaxParallelMaintenanceWorkers() string {
	if o == nil || IsNil(o.MaxParallelMaintenanceWorkers) {
		var ret string
		return ret
	}
	return *o.MaxParallelMaintenanceWorkers
}

// GetMaxParallelMaintenanceWorkersOk returns a tuple with the MaxParallelMaintenanceWorkers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetMaxParallelMaintenanceWorkersOk() (*string, bool) {
	if o == nil || IsNil(o.MaxParallelMaintenanceWorkers) {
		return nil, false
	}
	return o.MaxParallelMaintenanceWorkers, true
}

// HasMaxParallelMaintenanceWorkers returns a boolean if a field has been set.
func (o *Postgres) HasMaxParallelMaintenanceWorkers() bool {
	if o != nil && !IsNil(o.MaxParallelMaintenanceWorkers) {
		return true
	}

	return false
}

// SetMaxParallelMaintenanceWorkers gets a reference to the given string and assigns it to the MaxParallelMaintenanceWorkers field.
func (o *Postgres) SetMaxParallelMaintenanceWorkers(v string) {
	o.MaxParallelMaintenanceWorkers = &v
}

// GetMaxParallelWorkers returns the MaxParallelWorkers field value if set, zero value otherwise.
func (o *Postgres) GetMaxParallelWorkers() string {
	if o == nil || IsNil(o.MaxParallelWorkers) {
		var ret string
		return ret
	}
	return *o.MaxParallelWorkers
}

// GetMaxParallelWorkersOk returns a tuple with the MaxParallelWorkers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetMaxParallelWorkersOk() (*string, bool) {
	if o == nil || IsNil(o.MaxParallelWorkers) {
		return nil, false
	}
	return o.MaxParallelWorkers, true
}

// HasMaxParallelWorkers returns a boolean if a field has been set.
func (o *Postgres) HasMaxParallelWorkers() bool {
	if o != nil && !IsNil(o.MaxParallelWorkers) {
		return true
	}

	return false
}

// SetMaxParallelWorkers gets a reference to the given string and assigns it to the MaxParallelWorkers field.
func (o *Postgres) SetMaxParallelWorkers(v string) {
	o.MaxParallelWorkers = &v
}

// GetMaxParallelWorkersPerGather returns the MaxParallelWorkersPerGather field value if set, zero value otherwise.
func (o *Postgres) GetMaxParallelWorkersPerGather() string {
	if o == nil || IsNil(o.MaxParallelWorkersPerGather) {
		var ret string
		return ret
	}
	return *o.MaxParallelWorkersPerGather
}

// GetMaxParallelWorkersPerGatherOk returns a tuple with the MaxParallelWorkersPerGather field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetMaxParallelWorkersPerGatherOk() (*string, bool) {
	if o == nil || IsNil(o.MaxParallelWorkersPerGather) {
		return nil, false
	}
	return o.MaxParallelWorkersPerGather, true
}

// HasMaxParallelWorkersPerGather returns a boolean if a field has been set.
func (o *Postgres) HasMaxParallelWorkersPerGather() bool {
	if o != nil && !IsNil(o.MaxParallelWorkersPerGather) {
		return true
	}

	return false
}

// SetMaxParallelWorkersPerGather gets a reference to the given string and assigns it to the MaxParallelWorkersPerGather field.
func (o *Postgres) SetMaxParallelWorkersPerGather(v string) {
	o.MaxParallelWorkersPerGather = &v
}

// GetArrayNulls returns the ArrayNulls field value if set, zero value otherwise.
func (o *Postgres) GetArrayNulls() string {
	if o == nil || IsNil(o.ArrayNulls) {
		var ret string
		return ret
	}
	return *o.ArrayNulls
}

// GetArrayNullsOk returns a tuple with the ArrayNulls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetArrayNullsOk() (*string, bool) {
	if o == nil || IsNil(o.ArrayNulls) {
		return nil, false
	}
	return o.ArrayNulls, true
}

// HasArrayNulls returns a boolean if a field has been set.
func (o *Postgres) HasArrayNulls() bool {
	if o != nil && !IsNil(o.ArrayNulls) {
		return true
	}

	return false
}

// SetArrayNulls gets a reference to the given string and assigns it to the ArrayNulls field.
func (o *Postgres) SetArrayNulls(v string) {
	o.ArrayNulls = &v
}

// GetBackendFlushAfter returns the BackendFlushAfter field value if set, zero value otherwise.
func (o *Postgres) GetBackendFlushAfter() string {
	if o == nil || IsNil(o.BackendFlushAfter) {
		var ret string
		return ret
	}
	return *o.BackendFlushAfter
}

// GetBackendFlushAfterOk returns a tuple with the BackendFlushAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetBackendFlushAfterOk() (*string, bool) {
	if o == nil || IsNil(o.BackendFlushAfter) {
		return nil, false
	}
	return o.BackendFlushAfter, true
}

// HasBackendFlushAfter returns a boolean if a field has been set.
func (o *Postgres) HasBackendFlushAfter() bool {
	if o != nil && !IsNil(o.BackendFlushAfter) {
		return true
	}

	return false
}

// SetBackendFlushAfter gets a reference to the given string and assigns it to the BackendFlushAfter field.
func (o *Postgres) SetBackendFlushAfter(v string) {
	o.BackendFlushAfter = &v
}

// GetBackslashQuote returns the BackslashQuote field value if set, zero value otherwise.
func (o *Postgres) GetBackslashQuote() string {
	if o == nil || IsNil(o.BackslashQuote) {
		var ret string
		return ret
	}
	return *o.BackslashQuote
}

// GetBackslashQuoteOk returns a tuple with the BackslashQuote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetBackslashQuoteOk() (*string, bool) {
	if o == nil || IsNil(o.BackslashQuote) {
		return nil, false
	}
	return o.BackslashQuote, true
}

// HasBackslashQuote returns a boolean if a field has been set.
func (o *Postgres) HasBackslashQuote() bool {
	if o != nil && !IsNil(o.BackslashQuote) {
		return true
	}

	return false
}

// SetBackslashQuote gets a reference to the given string and assigns it to the BackslashQuote field.
func (o *Postgres) SetBackslashQuote(v string) {
	o.BackslashQuote = &v
}

// GetBgwriterFlushAfter returns the BgwriterFlushAfter field value if set, zero value otherwise.
func (o *Postgres) GetBgwriterFlushAfter() string {
	if o == nil || IsNil(o.BgwriterFlushAfter) {
		var ret string
		return ret
	}
	return *o.BgwriterFlushAfter
}

// GetBgwriterFlushAfterOk returns a tuple with the BgwriterFlushAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetBgwriterFlushAfterOk() (*string, bool) {
	if o == nil || IsNil(o.BgwriterFlushAfter) {
		return nil, false
	}
	return o.BgwriterFlushAfter, true
}

// HasBgwriterFlushAfter returns a boolean if a field has been set.
func (o *Postgres) HasBgwriterFlushAfter() bool {
	if o != nil && !IsNil(o.BgwriterFlushAfter) {
		return true
	}

	return false
}

// SetBgwriterFlushAfter gets a reference to the given string and assigns it to the BgwriterFlushAfter field.
func (o *Postgres) SetBgwriterFlushAfter(v string) {
	o.BgwriterFlushAfter = &v
}

// GetBgwriterLruMultiplier returns the BgwriterLruMultiplier field value if set, zero value otherwise.
func (o *Postgres) GetBgwriterLruMultiplier() string {
	if o == nil || IsNil(o.BgwriterLruMultiplier) {
		var ret string
		return ret
	}
	return *o.BgwriterLruMultiplier
}

// GetBgwriterLruMultiplierOk returns a tuple with the BgwriterLruMultiplier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetBgwriterLruMultiplierOk() (*string, bool) {
	if o == nil || IsNil(o.BgwriterLruMultiplier) {
		return nil, false
	}
	return o.BgwriterLruMultiplier, true
}

// HasBgwriterLruMultiplier returns a boolean if a field has been set.
func (o *Postgres) HasBgwriterLruMultiplier() bool {
	if o != nil && !IsNil(o.BgwriterLruMultiplier) {
		return true
	}

	return false
}

// SetBgwriterLruMultiplier gets a reference to the given string and assigns it to the BgwriterLruMultiplier field.
func (o *Postgres) SetBgwriterLruMultiplier(v string) {
	o.BgwriterLruMultiplier = &v
}

// GetDefaultTransactionReadOnly returns the DefaultTransactionReadOnly field value if set, zero value otherwise.
func (o *Postgres) GetDefaultTransactionReadOnly() string {
	if o == nil || IsNil(o.DefaultTransactionReadOnly) {
		var ret string
		return ret
	}
	return *o.DefaultTransactionReadOnly
}

// GetDefaultTransactionReadOnlyOk returns a tuple with the DefaultTransactionReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetDefaultTransactionReadOnlyOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultTransactionReadOnly) {
		return nil, false
	}
	return o.DefaultTransactionReadOnly, true
}

// HasDefaultTransactionReadOnly returns a boolean if a field has been set.
func (o *Postgres) HasDefaultTransactionReadOnly() bool {
	if o != nil && !IsNil(o.DefaultTransactionReadOnly) {
		return true
	}

	return false
}

// SetDefaultTransactionReadOnly gets a reference to the given string and assigns it to the DefaultTransactionReadOnly field.
func (o *Postgres) SetDefaultTransactionReadOnly(v string) {
	o.DefaultTransactionReadOnly = &v
}

// GetEnableHashagg returns the EnableHashagg field value if set, zero value otherwise.
func (o *Postgres) GetEnableHashagg() string {
	if o == nil || IsNil(o.EnableHashagg) {
		var ret string
		return ret
	}
	return *o.EnableHashagg
}

// GetEnableHashaggOk returns a tuple with the EnableHashagg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetEnableHashaggOk() (*string, bool) {
	if o == nil || IsNil(o.EnableHashagg) {
		return nil, false
	}
	return o.EnableHashagg, true
}

// HasEnableHashagg returns a boolean if a field has been set.
func (o *Postgres) HasEnableHashagg() bool {
	if o != nil && !IsNil(o.EnableHashagg) {
		return true
	}

	return false
}

// SetEnableHashagg gets a reference to the given string and assigns it to the EnableHashagg field.
func (o *Postgres) SetEnableHashagg(v string) {
	o.EnableHashagg = &v
}

// GetEnableHashjoin returns the EnableHashjoin field value if set, zero value otherwise.
func (o *Postgres) GetEnableHashjoin() string {
	if o == nil || IsNil(o.EnableHashjoin) {
		var ret string
		return ret
	}
	return *o.EnableHashjoin
}

// GetEnableHashjoinOk returns a tuple with the EnableHashjoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetEnableHashjoinOk() (*string, bool) {
	if o == nil || IsNil(o.EnableHashjoin) {
		return nil, false
	}
	return o.EnableHashjoin, true
}

// HasEnableHashjoin returns a boolean if a field has been set.
func (o *Postgres) HasEnableHashjoin() bool {
	if o != nil && !IsNil(o.EnableHashjoin) {
		return true
	}

	return false
}

// SetEnableHashjoin gets a reference to the given string and assigns it to the EnableHashjoin field.
func (o *Postgres) SetEnableHashjoin(v string) {
	o.EnableHashjoin = &v
}

// GetEnableIncrementalSort returns the EnableIncrementalSort field value if set, zero value otherwise.
func (o *Postgres) GetEnableIncrementalSort() string {
	if o == nil || IsNil(o.EnableIncrementalSort) {
		var ret string
		return ret
	}
	return *o.EnableIncrementalSort
}

// GetEnableIncrementalSortOk returns a tuple with the EnableIncrementalSort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetEnableIncrementalSortOk() (*string, bool) {
	if o == nil || IsNil(o.EnableIncrementalSort) {
		return nil, false
	}
	return o.EnableIncrementalSort, true
}

// HasEnableIncrementalSort returns a boolean if a field has been set.
func (o *Postgres) HasEnableIncrementalSort() bool {
	if o != nil && !IsNil(o.EnableIncrementalSort) {
		return true
	}

	return false
}

// SetEnableIncrementalSort gets a reference to the given string and assigns it to the EnableIncrementalSort field.
func (o *Postgres) SetEnableIncrementalSort(v string) {
	o.EnableIncrementalSort = &v
}

// GetEnableIndexscan returns the EnableIndexscan field value if set, zero value otherwise.
func (o *Postgres) GetEnableIndexscan() string {
	if o == nil || IsNil(o.EnableIndexscan) {
		var ret string
		return ret
	}
	return *o.EnableIndexscan
}

// GetEnableIndexscanOk returns a tuple with the EnableIndexscan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetEnableIndexscanOk() (*string, bool) {
	if o == nil || IsNil(o.EnableIndexscan) {
		return nil, false
	}
	return o.EnableIndexscan, true
}

// HasEnableIndexscan returns a boolean if a field has been set.
func (o *Postgres) HasEnableIndexscan() bool {
	if o != nil && !IsNil(o.EnableIndexscan) {
		return true
	}

	return false
}

// SetEnableIndexscan gets a reference to the given string and assigns it to the EnableIndexscan field.
func (o *Postgres) SetEnableIndexscan(v string) {
	o.EnableIndexscan = &v
}

// GetEnableIndexonlyscan returns the EnableIndexonlyscan field value if set, zero value otherwise.
func (o *Postgres) GetEnableIndexonlyscan() string {
	if o == nil || IsNil(o.EnableIndexonlyscan) {
		var ret string
		return ret
	}
	return *o.EnableIndexonlyscan
}

// GetEnableIndexonlyscanOk returns a tuple with the EnableIndexonlyscan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetEnableIndexonlyscanOk() (*string, bool) {
	if o == nil || IsNil(o.EnableIndexonlyscan) {
		return nil, false
	}
	return o.EnableIndexonlyscan, true
}

// HasEnableIndexonlyscan returns a boolean if a field has been set.
func (o *Postgres) HasEnableIndexonlyscan() bool {
	if o != nil && !IsNil(o.EnableIndexonlyscan) {
		return true
	}

	return false
}

// SetEnableIndexonlyscan gets a reference to the given string and assigns it to the EnableIndexonlyscan field.
func (o *Postgres) SetEnableIndexonlyscan(v string) {
	o.EnableIndexonlyscan = &v
}

// GetEnableMaterial returns the EnableMaterial field value if set, zero value otherwise.
func (o *Postgres) GetEnableMaterial() string {
	if o == nil || IsNil(o.EnableMaterial) {
		var ret string
		return ret
	}
	return *o.EnableMaterial
}

// GetEnableMaterialOk returns a tuple with the EnableMaterial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetEnableMaterialOk() (*string, bool) {
	if o == nil || IsNil(o.EnableMaterial) {
		return nil, false
	}
	return o.EnableMaterial, true
}

// HasEnableMaterial returns a boolean if a field has been set.
func (o *Postgres) HasEnableMaterial() bool {
	if o != nil && !IsNil(o.EnableMaterial) {
		return true
	}

	return false
}

// SetEnableMaterial gets a reference to the given string and assigns it to the EnableMaterial field.
func (o *Postgres) SetEnableMaterial(v string) {
	o.EnableMaterial = &v
}

// GetEnableMemoize returns the EnableMemoize field value if set, zero value otherwise.
func (o *Postgres) GetEnableMemoize() string {
	if o == nil || IsNil(o.EnableMemoize) {
		var ret string
		return ret
	}
	return *o.EnableMemoize
}

// GetEnableMemoizeOk returns a tuple with the EnableMemoize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetEnableMemoizeOk() (*string, bool) {
	if o == nil || IsNil(o.EnableMemoize) {
		return nil, false
	}
	return o.EnableMemoize, true
}

// HasEnableMemoize returns a boolean if a field has been set.
func (o *Postgres) HasEnableMemoize() bool {
	if o != nil && !IsNil(o.EnableMemoize) {
		return true
	}

	return false
}

// SetEnableMemoize gets a reference to the given string and assigns it to the EnableMemoize field.
func (o *Postgres) SetEnableMemoize(v string) {
	o.EnableMemoize = &v
}

// GetEnableMergejoin returns the EnableMergejoin field value if set, zero value otherwise.
func (o *Postgres) GetEnableMergejoin() string {
	if o == nil || IsNil(o.EnableMergejoin) {
		var ret string
		return ret
	}
	return *o.EnableMergejoin
}

// GetEnableMergejoinOk returns a tuple with the EnableMergejoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetEnableMergejoinOk() (*string, bool) {
	if o == nil || IsNil(o.EnableMergejoin) {
		return nil, false
	}
	return o.EnableMergejoin, true
}

// HasEnableMergejoin returns a boolean if a field has been set.
func (o *Postgres) HasEnableMergejoin() bool {
	if o != nil && !IsNil(o.EnableMergejoin) {
		return true
	}

	return false
}

// SetEnableMergejoin gets a reference to the given string and assigns it to the EnableMergejoin field.
func (o *Postgres) SetEnableMergejoin(v string) {
	o.EnableMergejoin = &v
}

// GetEnableParallelAppend returns the EnableParallelAppend field value if set, zero value otherwise.
func (o *Postgres) GetEnableParallelAppend() string {
	if o == nil || IsNil(o.EnableParallelAppend) {
		var ret string
		return ret
	}
	return *o.EnableParallelAppend
}

// GetEnableParallelAppendOk returns a tuple with the EnableParallelAppend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetEnableParallelAppendOk() (*string, bool) {
	if o == nil || IsNil(o.EnableParallelAppend) {
		return nil, false
	}
	return o.EnableParallelAppend, true
}

// HasEnableParallelAppend returns a boolean if a field has been set.
func (o *Postgres) HasEnableParallelAppend() bool {
	if o != nil && !IsNil(o.EnableParallelAppend) {
		return true
	}

	return false
}

// SetEnableParallelAppend gets a reference to the given string and assigns it to the EnableParallelAppend field.
func (o *Postgres) SetEnableParallelAppend(v string) {
	o.EnableParallelAppend = &v
}

// GetEnableParallelHash returns the EnableParallelHash field value if set, zero value otherwise.
func (o *Postgres) GetEnableParallelHash() string {
	if o == nil || IsNil(o.EnableParallelHash) {
		var ret string
		return ret
	}
	return *o.EnableParallelHash
}

// GetEnableParallelHashOk returns a tuple with the EnableParallelHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetEnableParallelHashOk() (*string, bool) {
	if o == nil || IsNil(o.EnableParallelHash) {
		return nil, false
	}
	return o.EnableParallelHash, true
}

// HasEnableParallelHash returns a boolean if a field has been set.
func (o *Postgres) HasEnableParallelHash() bool {
	if o != nil && !IsNil(o.EnableParallelHash) {
		return true
	}

	return false
}

// SetEnableParallelHash gets a reference to the given string and assigns it to the EnableParallelHash field.
func (o *Postgres) SetEnableParallelHash(v string) {
	o.EnableParallelHash = &v
}

// GetEnablePartitionPruning returns the EnablePartitionPruning field value if set, zero value otherwise.
func (o *Postgres) GetEnablePartitionPruning() string {
	if o == nil || IsNil(o.EnablePartitionPruning) {
		var ret string
		return ret
	}
	return *o.EnablePartitionPruning
}

// GetEnablePartitionPruningOk returns a tuple with the EnablePartitionPruning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetEnablePartitionPruningOk() (*string, bool) {
	if o == nil || IsNil(o.EnablePartitionPruning) {
		return nil, false
	}
	return o.EnablePartitionPruning, true
}

// HasEnablePartitionPruning returns a boolean if a field has been set.
func (o *Postgres) HasEnablePartitionPruning() bool {
	if o != nil && !IsNil(o.EnablePartitionPruning) {
		return true
	}

	return false
}

// SetEnablePartitionPruning gets a reference to the given string and assigns it to the EnablePartitionPruning field.
func (o *Postgres) SetEnablePartitionPruning(v string) {
	o.EnablePartitionPruning = &v
}

// GetEnablePartitionwiseJoin returns the EnablePartitionwiseJoin field value if set, zero value otherwise.
func (o *Postgres) GetEnablePartitionwiseJoin() string {
	if o == nil || IsNil(o.EnablePartitionwiseJoin) {
		var ret string
		return ret
	}
	return *o.EnablePartitionwiseJoin
}

// GetEnablePartitionwiseJoinOk returns a tuple with the EnablePartitionwiseJoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetEnablePartitionwiseJoinOk() (*string, bool) {
	if o == nil || IsNil(o.EnablePartitionwiseJoin) {
		return nil, false
	}
	return o.EnablePartitionwiseJoin, true
}

// HasEnablePartitionwiseJoin returns a boolean if a field has been set.
func (o *Postgres) HasEnablePartitionwiseJoin() bool {
	if o != nil && !IsNil(o.EnablePartitionwiseJoin) {
		return true
	}

	return false
}

// SetEnablePartitionwiseJoin gets a reference to the given string and assigns it to the EnablePartitionwiseJoin field.
func (o *Postgres) SetEnablePartitionwiseJoin(v string) {
	o.EnablePartitionwiseJoin = &v
}

// GetEnablePartitionwiseAggregate returns the EnablePartitionwiseAggregate field value if set, zero value otherwise.
func (o *Postgres) GetEnablePartitionwiseAggregate() string {
	if o == nil || IsNil(o.EnablePartitionwiseAggregate) {
		var ret string
		return ret
	}
	return *o.EnablePartitionwiseAggregate
}

// GetEnablePartitionwiseAggregateOk returns a tuple with the EnablePartitionwiseAggregate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetEnablePartitionwiseAggregateOk() (*string, bool) {
	if o == nil || IsNil(o.EnablePartitionwiseAggregate) {
		return nil, false
	}
	return o.EnablePartitionwiseAggregate, true
}

// HasEnablePartitionwiseAggregate returns a boolean if a field has been set.
func (o *Postgres) HasEnablePartitionwiseAggregate() bool {
	if o != nil && !IsNil(o.EnablePartitionwiseAggregate) {
		return true
	}

	return false
}

// SetEnablePartitionwiseAggregate gets a reference to the given string and assigns it to the EnablePartitionwiseAggregate field.
func (o *Postgres) SetEnablePartitionwiseAggregate(v string) {
	o.EnablePartitionwiseAggregate = &v
}

// GetEnableSeqscan returns the EnableSeqscan field value if set, zero value otherwise.
func (o *Postgres) GetEnableSeqscan() string {
	if o == nil || IsNil(o.EnableSeqscan) {
		var ret string
		return ret
	}
	return *o.EnableSeqscan
}

// GetEnableSeqscanOk returns a tuple with the EnableSeqscan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetEnableSeqscanOk() (*string, bool) {
	if o == nil || IsNil(o.EnableSeqscan) {
		return nil, false
	}
	return o.EnableSeqscan, true
}

// HasEnableSeqscan returns a boolean if a field has been set.
func (o *Postgres) HasEnableSeqscan() bool {
	if o != nil && !IsNil(o.EnableSeqscan) {
		return true
	}

	return false
}

// SetEnableSeqscan gets a reference to the given string and assigns it to the EnableSeqscan field.
func (o *Postgres) SetEnableSeqscan(v string) {
	o.EnableSeqscan = &v
}

// GetEnableSort returns the EnableSort field value if set, zero value otherwise.
func (o *Postgres) GetEnableSort() string {
	if o == nil || IsNil(o.EnableSort) {
		var ret string
		return ret
	}
	return *o.EnableSort
}

// GetEnableSortOk returns a tuple with the EnableSort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetEnableSortOk() (*string, bool) {
	if o == nil || IsNil(o.EnableSort) {
		return nil, false
	}
	return o.EnableSort, true
}

// HasEnableSort returns a boolean if a field has been set.
func (o *Postgres) HasEnableSort() bool {
	if o != nil && !IsNil(o.EnableSort) {
		return true
	}

	return false
}

// SetEnableSort gets a reference to the given string and assigns it to the EnableSort field.
func (o *Postgres) SetEnableSort(v string) {
	o.EnableSort = &v
}

// GetEnableTidscan returns the EnableTidscan field value if set, zero value otherwise.
func (o *Postgres) GetEnableTidscan() string {
	if o == nil || IsNil(o.EnableTidscan) {
		var ret string
		return ret
	}
	return *o.EnableTidscan
}

// GetEnableTidscanOk returns a tuple with the EnableTidscan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetEnableTidscanOk() (*string, bool) {
	if o == nil || IsNil(o.EnableTidscan) {
		return nil, false
	}
	return o.EnableTidscan, true
}

// HasEnableTidscan returns a boolean if a field has been set.
func (o *Postgres) HasEnableTidscan() bool {
	if o != nil && !IsNil(o.EnableTidscan) {
		return true
	}

	return false
}

// SetEnableTidscan gets a reference to the given string and assigns it to the EnableTidscan field.
func (o *Postgres) SetEnableTidscan(v string) {
	o.EnableTidscan = &v
}

// GetExitOnError returns the ExitOnError field value if set, zero value otherwise.
func (o *Postgres) GetExitOnError() string {
	if o == nil || IsNil(o.ExitOnError) {
		var ret string
		return ret
	}
	return *o.ExitOnError
}

// GetExitOnErrorOk returns a tuple with the ExitOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetExitOnErrorOk() (*string, bool) {
	if o == nil || IsNil(o.ExitOnError) {
		return nil, false
	}
	return o.ExitOnError, true
}

// HasExitOnError returns a boolean if a field has been set.
func (o *Postgres) HasExitOnError() bool {
	if o != nil && !IsNil(o.ExitOnError) {
		return true
	}

	return false
}

// SetExitOnError gets a reference to the given string and assigns it to the ExitOnError field.
func (o *Postgres) SetExitOnError(v string) {
	o.ExitOnError = &v
}

// GetFromCollapseLimit returns the FromCollapseLimit field value if set, zero value otherwise.
func (o *Postgres) GetFromCollapseLimit() string {
	if o == nil || IsNil(o.FromCollapseLimit) {
		var ret string
		return ret
	}
	return *o.FromCollapseLimit
}

// GetFromCollapseLimitOk returns a tuple with the FromCollapseLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetFromCollapseLimitOk() (*string, bool) {
	if o == nil || IsNil(o.FromCollapseLimit) {
		return nil, false
	}
	return o.FromCollapseLimit, true
}

// HasFromCollapseLimit returns a boolean if a field has been set.
func (o *Postgres) HasFromCollapseLimit() bool {
	if o != nil && !IsNil(o.FromCollapseLimit) {
		return true
	}

	return false
}

// SetFromCollapseLimit gets a reference to the given string and assigns it to the FromCollapseLimit field.
func (o *Postgres) SetFromCollapseLimit(v string) {
	o.FromCollapseLimit = &v
}

// GetJit returns the Jit field value if set, zero value otherwise.
func (o *Postgres) GetJit() string {
	if o == nil || IsNil(o.Jit) {
		var ret string
		return ret
	}
	return *o.Jit
}

// GetJitOk returns a tuple with the Jit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetJitOk() (*string, bool) {
	if o == nil || IsNil(o.Jit) {
		return nil, false
	}
	return o.Jit, true
}

// HasJit returns a boolean if a field has been set.
func (o *Postgres) HasJit() bool {
	if o != nil && !IsNil(o.Jit) {
		return true
	}

	return false
}

// SetJit gets a reference to the given string and assigns it to the Jit field.
func (o *Postgres) SetJit(v string) {
	o.Jit = &v
}

// GetPlanCacheMode returns the PlanCacheMode field value if set, zero value otherwise.
func (o *Postgres) GetPlanCacheMode() string {
	if o == nil || IsNil(o.PlanCacheMode) {
		var ret string
		return ret
	}
	return *o.PlanCacheMode
}

// GetPlanCacheModeOk returns a tuple with the PlanCacheMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetPlanCacheModeOk() (*string, bool) {
	if o == nil || IsNil(o.PlanCacheMode) {
		return nil, false
	}
	return o.PlanCacheMode, true
}

// HasPlanCacheMode returns a boolean if a field has been set.
func (o *Postgres) HasPlanCacheMode() bool {
	if o != nil && !IsNil(o.PlanCacheMode) {
		return true
	}

	return false
}

// SetPlanCacheMode gets a reference to the given string and assigns it to the PlanCacheMode field.
func (o *Postgres) SetPlanCacheMode(v string) {
	o.PlanCacheMode = &v
}

// GetQuoteAllIdentifiers returns the QuoteAllIdentifiers field value if set, zero value otherwise.
func (o *Postgres) GetQuoteAllIdentifiers() string {
	if o == nil || IsNil(o.QuoteAllIdentifiers) {
		var ret string
		return ret
	}
	return *o.QuoteAllIdentifiers
}

// GetQuoteAllIdentifiersOk returns a tuple with the QuoteAllIdentifiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetQuoteAllIdentifiersOk() (*string, bool) {
	if o == nil || IsNil(o.QuoteAllIdentifiers) {
		return nil, false
	}
	return o.QuoteAllIdentifiers, true
}

// HasQuoteAllIdentifiers returns a boolean if a field has been set.
func (o *Postgres) HasQuoteAllIdentifiers() bool {
	if o != nil && !IsNil(o.QuoteAllIdentifiers) {
		return true
	}

	return false
}

// SetQuoteAllIdentifiers gets a reference to the given string and assigns it to the QuoteAllIdentifiers field.
func (o *Postgres) SetQuoteAllIdentifiers(v string) {
	o.QuoteAllIdentifiers = &v
}

// GetStandardConformingStrings returns the StandardConformingStrings field value if set, zero value otherwise.
func (o *Postgres) GetStandardConformingStrings() string {
	if o == nil || IsNil(o.StandardConformingStrings) {
		var ret string
		return ret
	}
	return *o.StandardConformingStrings
}

// GetStandardConformingStringsOk returns a tuple with the StandardConformingStrings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetStandardConformingStringsOk() (*string, bool) {
	if o == nil || IsNil(o.StandardConformingStrings) {
		return nil, false
	}
	return o.StandardConformingStrings, true
}

// HasStandardConformingStrings returns a boolean if a field has been set.
func (o *Postgres) HasStandardConformingStrings() bool {
	if o != nil && !IsNil(o.StandardConformingStrings) {
		return true
	}

	return false
}

// SetStandardConformingStrings gets a reference to the given string and assigns it to the StandardConformingStrings field.
func (o *Postgres) SetStandardConformingStrings(v string) {
	o.StandardConformingStrings = &v
}

// GetStatementTimeout returns the StatementTimeout field value if set, zero value otherwise.
func (o *Postgres) GetStatementTimeout() string {
	if o == nil || IsNil(o.StatementTimeout) {
		var ret string
		return ret
	}
	return *o.StatementTimeout
}

// GetStatementTimeoutOk returns a tuple with the StatementTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetStatementTimeoutOk() (*string, bool) {
	if o == nil || IsNil(o.StatementTimeout) {
		return nil, false
	}
	return o.StatementTimeout, true
}

// HasStatementTimeout returns a boolean if a field has been set.
func (o *Postgres) HasStatementTimeout() bool {
	if o != nil && !IsNil(o.StatementTimeout) {
		return true
	}

	return false
}

// SetStatementTimeout gets a reference to the given string and assigns it to the StatementTimeout field.
func (o *Postgres) SetStatementTimeout(v string) {
	o.StatementTimeout = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *Postgres) GetTimezone() string {
	if o == nil || IsNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *Postgres) HasTimezone() bool {
	if o != nil && !IsNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *Postgres) SetTimezone(v string) {
	o.Timezone = &v
}

// GetTransformNullEquals returns the TransformNullEquals field value if set, zero value otherwise.
func (o *Postgres) GetTransformNullEquals() string {
	if o == nil || IsNil(o.TransformNullEquals) {
		var ret string
		return ret
	}
	return *o.TransformNullEquals
}

// GetTransformNullEqualsOk returns a tuple with the TransformNullEquals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetTransformNullEqualsOk() (*string, bool) {
	if o == nil || IsNil(o.TransformNullEquals) {
		return nil, false
	}
	return o.TransformNullEquals, true
}

// HasTransformNullEquals returns a boolean if a field has been set.
func (o *Postgres) HasTransformNullEquals() bool {
	if o != nil && !IsNil(o.TransformNullEquals) {
		return true
	}

	return false
}

// SetTransformNullEquals gets a reference to the given string and assigns it to the TransformNullEquals field.
func (o *Postgres) SetTransformNullEquals(v string) {
	o.TransformNullEquals = &v
}

// GetMaxLocksPerTransaction returns the MaxLocksPerTransaction field value if set, zero value otherwise.
func (o *Postgres) GetMaxLocksPerTransaction() string {
	if o == nil || IsNil(o.MaxLocksPerTransaction) {
		var ret string
		return ret
	}
	return *o.MaxLocksPerTransaction
}

// GetMaxLocksPerTransactionOk returns a tuple with the MaxLocksPerTransaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetMaxLocksPerTransactionOk() (*string, bool) {
	if o == nil || IsNil(o.MaxLocksPerTransaction) {
		return nil, false
	}
	return o.MaxLocksPerTransaction, true
}

// HasMaxLocksPerTransaction returns a boolean if a field has been set.
func (o *Postgres) HasMaxLocksPerTransaction() bool {
	if o != nil && !IsNil(o.MaxLocksPerTransaction) {
		return true
	}

	return false
}

// SetMaxLocksPerTransaction gets a reference to the given string and assigns it to the MaxLocksPerTransaction field.
func (o *Postgres) SetMaxLocksPerTransaction(v string) {
	o.MaxLocksPerTransaction = &v
}

// GetAutovacuumVacuumCostLimit returns the AutovacuumVacuumCostLimit field value if set, zero value otherwise.
func (o *Postgres) GetAutovacuumVacuumCostLimit() string {
	if o == nil || IsNil(o.AutovacuumVacuumCostLimit) {
		var ret string
		return ret
	}
	return *o.AutovacuumVacuumCostLimit
}

// GetAutovacuumVacuumCostLimitOk returns a tuple with the AutovacuumVacuumCostLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetAutovacuumVacuumCostLimitOk() (*string, bool) {
	if o == nil || IsNil(o.AutovacuumVacuumCostLimit) {
		return nil, false
	}
	return o.AutovacuumVacuumCostLimit, true
}

// HasAutovacuumVacuumCostLimit returns a boolean if a field has been set.
func (o *Postgres) HasAutovacuumVacuumCostLimit() bool {
	if o != nil && !IsNil(o.AutovacuumVacuumCostLimit) {
		return true
	}

	return false
}

// SetAutovacuumVacuumCostLimit gets a reference to the given string and assigns it to the AutovacuumVacuumCostLimit field.
func (o *Postgres) SetAutovacuumVacuumCostLimit(v string) {
	o.AutovacuumVacuumCostLimit = &v
}

// GetCheckpointTimeout returns the CheckpointTimeout field value if set, zero value otherwise.
func (o *Postgres) GetCheckpointTimeout() string {
	if o == nil || IsNil(o.CheckpointTimeout) {
		var ret string
		return ret
	}
	return *o.CheckpointTimeout
}

// GetCheckpointTimeoutOk returns a tuple with the CheckpointTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetCheckpointTimeoutOk() (*string, bool) {
	if o == nil || IsNil(o.CheckpointTimeout) {
		return nil, false
	}
	return o.CheckpointTimeout, true
}

// HasCheckpointTimeout returns a boolean if a field has been set.
func (o *Postgres) HasCheckpointTimeout() bool {
	if o != nil && !IsNil(o.CheckpointTimeout) {
		return true
	}

	return false
}

// SetCheckpointTimeout gets a reference to the given string and assigns it to the CheckpointTimeout field.
func (o *Postgres) SetCheckpointTimeout(v string) {
	o.CheckpointTimeout = &v
}

// GetCheckpointCompletionTarget returns the CheckpointCompletionTarget field value if set, zero value otherwise.
func (o *Postgres) GetCheckpointCompletionTarget() string {
	if o == nil || IsNil(o.CheckpointCompletionTarget) {
		var ret string
		return ret
	}
	return *o.CheckpointCompletionTarget
}

// GetCheckpointCompletionTargetOk returns a tuple with the CheckpointCompletionTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetCheckpointCompletionTargetOk() (*string, bool) {
	if o == nil || IsNil(o.CheckpointCompletionTarget) {
		return nil, false
	}
	return o.CheckpointCompletionTarget, true
}

// HasCheckpointCompletionTarget returns a boolean if a field has been set.
func (o *Postgres) HasCheckpointCompletionTarget() bool {
	if o != nil && !IsNil(o.CheckpointCompletionTarget) {
		return true
	}

	return false
}

// SetCheckpointCompletionTarget gets a reference to the given string and assigns it to the CheckpointCompletionTarget field.
func (o *Postgres) SetCheckpointCompletionTarget(v string) {
	o.CheckpointCompletionTarget = &v
}

// GetWalCompression returns the WalCompression field value if set, zero value otherwise.
func (o *Postgres) GetWalCompression() string {
	if o == nil || IsNil(o.WalCompression) {
		var ret string
		return ret
	}
	return *o.WalCompression
}

// GetWalCompressionOk returns a tuple with the WalCompression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetWalCompressionOk() (*string, bool) {
	if o == nil || IsNil(o.WalCompression) {
		return nil, false
	}
	return o.WalCompression, true
}

// HasWalCompression returns a boolean if a field has been set.
func (o *Postgres) HasWalCompression() bool {
	if o != nil && !IsNil(o.WalCompression) {
		return true
	}

	return false
}

// SetWalCompression gets a reference to the given string and assigns it to the WalCompression field.
func (o *Postgres) SetWalCompression(v string) {
	o.WalCompression = &v
}

// GetRandomPageCost returns the RandomPageCost field value if set, zero value otherwise.
func (o *Postgres) GetRandomPageCost() string {
	if o == nil || IsNil(o.RandomPageCost) {
		var ret string
		return ret
	}
	return *o.RandomPageCost
}

// GetRandomPageCostOk returns a tuple with the RandomPageCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetRandomPageCostOk() (*string, bool) {
	if o == nil || IsNil(o.RandomPageCost) {
		return nil, false
	}
	return o.RandomPageCost, true
}

// HasRandomPageCost returns a boolean if a field has been set.
func (o *Postgres) HasRandomPageCost() bool {
	if o != nil && !IsNil(o.RandomPageCost) {
		return true
	}

	return false
}

// SetRandomPageCost gets a reference to the given string and assigns it to the RandomPageCost field.
func (o *Postgres) SetRandomPageCost(v string) {
	o.RandomPageCost = &v
}

// GetEffectiveIoConcurrency returns the EffectiveIoConcurrency field value if set, zero value otherwise.
func (o *Postgres) GetEffectiveIoConcurrency() string {
	if o == nil || IsNil(o.EffectiveIoConcurrency) {
		var ret string
		return ret
	}
	return *o.EffectiveIoConcurrency
}

// GetEffectiveIoConcurrencyOk returns a tuple with the EffectiveIoConcurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetEffectiveIoConcurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.EffectiveIoConcurrency) {
		return nil, false
	}
	return o.EffectiveIoConcurrency, true
}

// HasEffectiveIoConcurrency returns a boolean if a field has been set.
func (o *Postgres) HasEffectiveIoConcurrency() bool {
	if o != nil && !IsNil(o.EffectiveIoConcurrency) {
		return true
	}

	return false
}

// SetEffectiveIoConcurrency gets a reference to the given string and assigns it to the EffectiveIoConcurrency field.
func (o *Postgres) SetEffectiveIoConcurrency(v string) {
	o.EffectiveIoConcurrency = &v
}

// GetLogLockWaits returns the LogLockWaits field value if set, zero value otherwise.
func (o *Postgres) GetLogLockWaits() string {
	if o == nil || IsNil(o.LogLockWaits) {
		var ret string
		return ret
	}
	return *o.LogLockWaits
}

// GetLogLockWaitsOk returns a tuple with the LogLockWaits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetLogLockWaitsOk() (*string, bool) {
	if o == nil || IsNil(o.LogLockWaits) {
		return nil, false
	}
	return o.LogLockWaits, true
}

// HasLogLockWaits returns a boolean if a field has been set.
func (o *Postgres) HasLogLockWaits() bool {
	if o != nil && !IsNil(o.LogLockWaits) {
		return true
	}

	return false
}

// SetLogLockWaits gets a reference to the given string and assigns it to the LogLockWaits field.
func (o *Postgres) SetLogLockWaits(v string) {
	o.LogLockWaits = &v
}

// GetLogTempFiles returns the LogTempFiles field value if set, zero value otherwise.
func (o *Postgres) GetLogTempFiles() string {
	if o == nil || IsNil(o.LogTempFiles) {
		var ret string
		return ret
	}
	return *o.LogTempFiles
}

// GetLogTempFilesOk returns a tuple with the LogTempFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetLogTempFilesOk() (*string, bool) {
	if o == nil || IsNil(o.LogTempFiles) {
		return nil, false
	}
	return o.LogTempFiles, true
}

// HasLogTempFiles returns a boolean if a field has been set.
func (o *Postgres) HasLogTempFiles() bool {
	if o != nil && !IsNil(o.LogTempFiles) {
		return true
	}

	return false
}

// SetLogTempFiles gets a reference to the given string and assigns it to the LogTempFiles field.
func (o *Postgres) SetLogTempFiles(v string) {
	o.LogTempFiles = &v
}

// GetTrackIoTiming returns the TrackIoTiming field value if set, zero value otherwise.
func (o *Postgres) GetTrackIoTiming() string {
	if o == nil || IsNil(o.TrackIoTiming) {
		var ret string
		return ret
	}
	return *o.TrackIoTiming
}

// GetTrackIoTimingOk returns a tuple with the TrackIoTiming field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetTrackIoTimingOk() (*string, bool) {
	if o == nil || IsNil(o.TrackIoTiming) {
		return nil, false
	}
	return o.TrackIoTiming, true
}

// HasTrackIoTiming returns a boolean if a field has been set.
func (o *Postgres) HasTrackIoTiming() bool {
	if o != nil && !IsNil(o.TrackIoTiming) {
		return true
	}

	return false
}

// SetTrackIoTiming gets a reference to the given string and assigns it to the TrackIoTiming field.
func (o *Postgres) SetTrackIoTiming(v string) {
	o.TrackIoTiming = &v
}

// GetMaintenanceWorkMem returns the MaintenanceWorkMem field value if set, zero value otherwise.
func (o *Postgres) GetMaintenanceWorkMem() string {
	if o == nil || IsNil(o.MaintenanceWorkMem) {
		var ret string
		return ret
	}
	return *o.MaintenanceWorkMem
}

// GetMaintenanceWorkMemOk returns a tuple with the MaintenanceWorkMem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetMaintenanceWorkMemOk() (*string, bool) {
	if o == nil || IsNil(o.MaintenanceWorkMem) {
		return nil, false
	}
	return o.MaintenanceWorkMem, true
}

// HasMaintenanceWorkMem returns a boolean if a field has been set.
func (o *Postgres) HasMaintenanceWorkMem() bool {
	if o != nil && !IsNil(o.MaintenanceWorkMem) {
		return true
	}

	return false
}

// SetMaintenanceWorkMem gets a reference to the given string and assigns it to the MaintenanceWorkMem field.
func (o *Postgres) SetMaintenanceWorkMem(v string) {
	o.MaintenanceWorkMem = &v
}

// GetIdleSessionTimeout returns the IdleSessionTimeout field value if set, zero value otherwise.
func (o *Postgres) GetIdleSessionTimeout() string {
	if o == nil || IsNil(o.IdleSessionTimeout) {
		var ret string
		return ret
	}
	return *o.IdleSessionTimeout
}

// GetIdleSessionTimeoutOk returns a tuple with the IdleSessionTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetIdleSessionTimeoutOk() (*string, bool) {
	if o == nil || IsNil(o.IdleSessionTimeout) {
		return nil, false
	}
	return o.IdleSessionTimeout, true
}

// HasIdleSessionTimeout returns a boolean if a field has been set.
func (o *Postgres) HasIdleSessionTimeout() bool {
	if o != nil && !IsNil(o.IdleSessionTimeout) {
		return true
	}

	return false
}

// SetIdleSessionTimeout gets a reference to the given string and assigns it to the IdleSessionTimeout field.
func (o *Postgres) SetIdleSessionTimeout(v string) {
	o.IdleSessionTimeout = &v
}

// GetIoMethod returns the IoMethod field value if set, zero value otherwise.
func (o *Postgres) GetIoMethod() string {
	if o == nil || IsNil(o.IoMethod) {
		var ret string
		return ret
	}
	return *o.IoMethod
}

// GetIoMethodOk returns a tuple with the IoMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetIoMethodOk() (*string, bool) {
	if o == nil || IsNil(o.IoMethod) {
		return nil, false
	}
	return o.IoMethod, true
}

// HasIoMethod returns a boolean if a field has been set.
func (o *Postgres) HasIoMethod() bool {
	if o != nil && !IsNil(o.IoMethod) {
		return true
	}

	return false
}

// SetIoMethod gets a reference to the given string and assigns it to the IoMethod field.
func (o *Postgres) SetIoMethod(v string) {
	o.IoMethod = &v
}

// GetIoWorkers returns the IoWorkers field value if set, zero value otherwise.
func (o *Postgres) GetIoWorkers() string {
	if o == nil || IsNil(o.IoWorkers) {
		var ret string
		return ret
	}
	return *o.IoWorkers
}

// GetIoWorkersOk returns a tuple with the IoWorkers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Postgres) GetIoWorkersOk() (*string, bool) {
	if o == nil || IsNil(o.IoWorkers) {
		return nil, false
	}
	return o.IoWorkers, true
}

// HasIoWorkers returns a boolean if a field has been set.
func (o *Postgres) HasIoWorkers() bool {
	if o != nil && !IsNil(o.IoWorkers) {
		return true
	}

	return false
}

// SetIoWorkers gets a reference to the given string and assigns it to the IoWorkers field.
func (o *Postgres) SetIoWorkers(v string) {
	o.IoWorkers = &v
}

func (o Postgres) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Postgres) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxConnections) {
		toSerialize["max_connections"] = o.MaxConnections
	}
	if !IsNil(o.AutovacuumAnalyzeScaleFactor) {
		toSerialize["autovacuum_analyze_scale_factor"] = o.AutovacuumAnalyzeScaleFactor
	}
	if !IsNil(o.AutovacuumMaxWorkers) {
		toSerialize["autovacuum_max_workers"] = o.AutovacuumMaxWorkers
	}
	if !IsNil(o.AutovacuumNaptime) {
		toSerialize["autovacuum_naptime"] = o.AutovacuumNaptime
	}
	if !IsNil(o.AutovacuumVacuumInsertScaleFactor) {
		toSerialize["autovacuum_vacuum_insert_scale_factor"] = o.AutovacuumVacuumInsertScaleFactor
	}
	if !IsNil(o.AutovacuumVacuumScaleFactor) {
		toSerialize["autovacuum_vacuum_scale_factor"] = o.AutovacuumVacuumScaleFactor
	}
	if !IsNil(o.AutovacuumWorkMem) {
		toSerialize["autovacuum_work_mem"] = o.AutovacuumWorkMem
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
	if !IsNil(o.JoinCollapseLimit) {
		toSerialize["join_collapse_limit"] = o.JoinCollapseLimit
	}
	if !IsNil(o.LockTimeout) {
		toSerialize["lock_timeout"] = o.LockTimeout
	}
	if !IsNil(o.MaxPreparedTransactions) {
		toSerialize["max_prepared_transactions"] = o.MaxPreparedTransactions
	}
	if !IsNil(o.SharedBuffers) {
		toSerialize["shared_buffers"] = o.SharedBuffers
	}
	if !IsNil(o.LogMinDurationStatement) {
		toSerialize["log_min_duration_statement"] = o.LogMinDurationStatement
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
	if !IsNil(o.DefaultTransactionIsolation) {
		toSerialize["default_transaction_isolation"] = o.DefaultTransactionIsolation
	}
	if !IsNil(o.EffectiveCacheSize) {
		toSerialize["effective_cache_size"] = o.EffectiveCacheSize
	}
	if !IsNil(o.MaxWalSize) {
		toSerialize["max_wal_size"] = o.MaxWalSize
	}
	if !IsNil(o.MinWalSize) {
		toSerialize["min_wal_size"] = o.MinWalSize
	}
	if !IsNil(o.WalLevel) {
		toSerialize["wal_level"] = o.WalLevel
	}
	if !IsNil(o.MaxReplicationSlots) {
		toSerialize["max_replication_slots"] = o.MaxReplicationSlots
	}
	if !IsNil(o.MaxWalSenders) {
		toSerialize["max_wal_senders"] = o.MaxWalSenders
	}
	if !IsNil(o.MaxWorkerProcesses) {
		toSerialize["max_worker_processes"] = o.MaxWorkerProcesses
	}
	if !IsNil(o.MaxLogicalReplicationWorkers) {
		toSerialize["max_logical_replication_workers"] = o.MaxLogicalReplicationWorkers
	}
	if !IsNil(o.MaxParallelMaintenanceWorkers) {
		toSerialize["max_parallel_maintenance_workers"] = o.MaxParallelMaintenanceWorkers
	}
	if !IsNil(o.MaxParallelWorkers) {
		toSerialize["max_parallel_workers"] = o.MaxParallelWorkers
	}
	if !IsNil(o.MaxParallelWorkersPerGather) {
		toSerialize["max_parallel_workers_per_gather"] = o.MaxParallelWorkersPerGather
	}
	if !IsNil(o.ArrayNulls) {
		toSerialize["array_nulls"] = o.ArrayNulls
	}
	if !IsNil(o.BackendFlushAfter) {
		toSerialize["backend_flush_after"] = o.BackendFlushAfter
	}
	if !IsNil(o.BackslashQuote) {
		toSerialize["backslash_quote"] = o.BackslashQuote
	}
	if !IsNil(o.BgwriterFlushAfter) {
		toSerialize["bgwriter_flush_after"] = o.BgwriterFlushAfter
	}
	if !IsNil(o.BgwriterLruMultiplier) {
		toSerialize["bgwriter_lru_multiplier"] = o.BgwriterLruMultiplier
	}
	if !IsNil(o.DefaultTransactionReadOnly) {
		toSerialize["default_transaction_read_only"] = o.DefaultTransactionReadOnly
	}
	if !IsNil(o.EnableHashagg) {
		toSerialize["enable_hashagg"] = o.EnableHashagg
	}
	if !IsNil(o.EnableHashjoin) {
		toSerialize["enable_hashjoin"] = o.EnableHashjoin
	}
	if !IsNil(o.EnableIncrementalSort) {
		toSerialize["enable_incremental_sort"] = o.EnableIncrementalSort
	}
	if !IsNil(o.EnableIndexscan) {
		toSerialize["enable_indexscan"] = o.EnableIndexscan
	}
	if !IsNil(o.EnableIndexonlyscan) {
		toSerialize["enable_indexonlyscan"] = o.EnableIndexonlyscan
	}
	if !IsNil(o.EnableMaterial) {
		toSerialize["enable_material"] = o.EnableMaterial
	}
	if !IsNil(o.EnableMemoize) {
		toSerialize["enable_memoize"] = o.EnableMemoize
	}
	if !IsNil(o.EnableMergejoin) {
		toSerialize["enable_mergejoin"] = o.EnableMergejoin
	}
	if !IsNil(o.EnableParallelAppend) {
		toSerialize["enable_parallel_append"] = o.EnableParallelAppend
	}
	if !IsNil(o.EnableParallelHash) {
		toSerialize["enable_parallel_hash"] = o.EnableParallelHash
	}
	if !IsNil(o.EnablePartitionPruning) {
		toSerialize["enable_partition_pruning"] = o.EnablePartitionPruning
	}
	if !IsNil(o.EnablePartitionwiseJoin) {
		toSerialize["enable_partitionwise_join"] = o.EnablePartitionwiseJoin
	}
	if !IsNil(o.EnablePartitionwiseAggregate) {
		toSerialize["enable_partitionwise_aggregate"] = o.EnablePartitionwiseAggregate
	}
	if !IsNil(o.EnableSeqscan) {
		toSerialize["enable_seqscan"] = o.EnableSeqscan
	}
	if !IsNil(o.EnableSort) {
		toSerialize["enable_sort"] = o.EnableSort
	}
	if !IsNil(o.EnableTidscan) {
		toSerialize["enable_tidscan"] = o.EnableTidscan
	}
	if !IsNil(o.ExitOnError) {
		toSerialize["exit_on_error"] = o.ExitOnError
	}
	if !IsNil(o.FromCollapseLimit) {
		toSerialize["from_collapse_limit"] = o.FromCollapseLimit
	}
	if !IsNil(o.Jit) {
		toSerialize["jit"] = o.Jit
	}
	if !IsNil(o.PlanCacheMode) {
		toSerialize["plan_cache_mode"] = o.PlanCacheMode
	}
	if !IsNil(o.QuoteAllIdentifiers) {
		toSerialize["quote_all_identifiers"] = o.QuoteAllIdentifiers
	}
	if !IsNil(o.StandardConformingStrings) {
		toSerialize["standard_conforming_strings"] = o.StandardConformingStrings
	}
	if !IsNil(o.StatementTimeout) {
		toSerialize["statement_timeout"] = o.StatementTimeout
	}
	if !IsNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}
	if !IsNil(o.TransformNullEquals) {
		toSerialize["transform_null_equals"] = o.TransformNullEquals
	}
	if !IsNil(o.MaxLocksPerTransaction) {
		toSerialize["max_locks_per_transaction"] = o.MaxLocksPerTransaction
	}
	if !IsNil(o.AutovacuumVacuumCostLimit) {
		toSerialize["autovacuum_vacuum_cost_limit"] = o.AutovacuumVacuumCostLimit
	}
	if !IsNil(o.CheckpointTimeout) {
		toSerialize["checkpoint_timeout"] = o.CheckpointTimeout
	}
	if !IsNil(o.CheckpointCompletionTarget) {
		toSerialize["checkpoint_completion_target"] = o.CheckpointCompletionTarget
	}
	if !IsNil(o.WalCompression) {
		toSerialize["wal_compression"] = o.WalCompression
	}
	if !IsNil(o.RandomPageCost) {
		toSerialize["random_page_cost"] = o.RandomPageCost
	}
	if !IsNil(o.EffectiveIoConcurrency) {
		toSerialize["effective_io_concurrency"] = o.EffectiveIoConcurrency
	}
	if !IsNil(o.LogLockWaits) {
		toSerialize["log_lock_waits"] = o.LogLockWaits
	}
	if !IsNil(o.LogTempFiles) {
		toSerialize["log_temp_files"] = o.LogTempFiles
	}
	if !IsNil(o.TrackIoTiming) {
		toSerialize["track_io_timing"] = o.TrackIoTiming
	}
	if !IsNil(o.MaintenanceWorkMem) {
		toSerialize["maintenance_work_mem"] = o.MaintenanceWorkMem
	}
	if !IsNil(o.IdleSessionTimeout) {
		toSerialize["idle_session_timeout"] = o.IdleSessionTimeout
	}
	if !IsNil(o.IoMethod) {
		toSerialize["io_method"] = o.IoMethod
	}
	if !IsNil(o.IoWorkers) {
		toSerialize["io_workers"] = o.IoWorkers
	}
	return toSerialize, nil
}

type NullablePostgres struct {
	value *Postgres
	isSet bool
}

func (v NullablePostgres) Get() *Postgres {
	return v.value
}

func (v *NullablePostgres) Set(val *Postgres) {
	v.value = val
	v.isSet = true
}

func (v NullablePostgres) IsSet() bool {
	return v.isSet
}

func (v *NullablePostgres) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostgres(val *Postgres) *NullablePostgres {
	return &NullablePostgres{value: val, isSet: true}
}

func (v NullablePostgres) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostgres) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


