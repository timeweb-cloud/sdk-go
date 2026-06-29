# Mysql

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JoinBufferSize** | Pointer to **string** | Размер буфера, используемого при соединениях таблиц без индексов (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60;). | [optional] 
**MaxConnections** | Pointer to **string** | Максимальное количество одновременных подключений к серверу (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60; | &#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**SortBufferSize** | Pointer to **string** | Размер буфера сортировки для операций ORDER BY и GROUP BY (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60;). | [optional] 
**ThreadCacheSize** | Pointer to **string** | Количество потоков, которые сервер сохраняет для повторного использования (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60;). | [optional] 
**InnodbBufferPoolSize** | Pointer to **string** | Размер буферного пула InnoDB для хранения данных и индексов в памяти (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60;). | [optional] 
**AutoIncrementIncrement** | Pointer to **string** | Интервал между значениями столбцов с атрибутом &#x60;AUTO_INCREMENT&#x60; (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60;). | [optional] 
**AutoIncrementOffset** | Pointer to **string** | Начальное значение для столбцов с атрибутом &#x60;AUTO_INCREMENT&#x60; (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60;). | [optional] 
**InnodbIoCapacity** | Pointer to **string** | Количество операций ввода-вывода в секунду &#x60;IOPS&#x60;, используемых InnoDB (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60;). | [optional] 
**InnodbPurgeThreads** | Pointer to **string** | Количество потоков, используемых для фоновой очистки undo-записей InnoDB (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60;). | [optional] 
**InnodbReadIoThreads** | Pointer to **string** | Количество потоков ввода-вывода для операций чтения InnoDB (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60;). | [optional] 
**InnodbThreadConcurrency** | Pointer to **string** | Ограничение количества одновременно выполняющихся потоков InnoDB (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60;). | [optional] 
**InnodbWriteIoThreads** | Pointer to **string** | Количество потоков ввода-вывода для операций записи InnoDB (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60;). | [optional] 
**InnodbLogFileSize** | Pointer to **string** | Размер файла журнала транзакций InnoDB redo log (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60;). | [optional] 
**MaxAllowedPacket** | Pointer to **string** | Максимальный размер пакета данных, который может передаваться между клиентом и сервером (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60;). | [optional] 
**MaxHeapTableSize** | Pointer to **string** | Максимальный размер таблиц типа MEMORY (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60;). | [optional] 
**SqlMode** | Pointer to **string** | Режим работы SQL сервера, определяющий поведение обработки запросов (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60;). | [optional] 
**QueryCacheType** | Pointer to **string** | Тип кэша запросов (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60;). | [optional] 
**QueryCacheSize** | Pointer to **string** | Объем памяти, выделяемый для кэширования результатов запросов (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60;). | [optional] 
**InnodbFlushLogAtTrxCommit** | Pointer to **string** | Режим записи журнала InnoDB при фиксации транзакций (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60;). | [optional] 
**TransactionIsolation** | Pointer to **string** | Уровень изоляции транзакций по умолчанию (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60;). | [optional] 
**LongQueryTime** | Pointer to **string** | Время выполнения запроса, после которого он считается долгим и может попасть в slow query log (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60;). | [optional] 
**TmpTableSize** | Pointer to **string** | Максимальный размер временных таблиц в памяти (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60;). | [optional] 
**TableOpenCache** | Pointer to **string** | Количество открытых таблиц, которые сервер может хранить в кэше (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60;). | [optional] 
**TableOpenCacheInstances** | Pointer to **string** | Количество экземпляров кэша открытых таблиц для снижения конкуренции между потоками (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60;). | [optional] 
**InnodbFlushMethod** | Pointer to **string** | Метод выполнения операций записи и синхронизации файлов InnoDB (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60;). | [optional] 
**InnodbStrictMode** | Pointer to **string** | Включение строгой проверки операций InnoDB (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60;). | [optional] 
**SlowQueryLog** | Pointer to **string** | Включение журнала медленных запросов (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60;). | [optional] 
**BinlogCacheSize** | Pointer to **string** | Размер кэша бинарного журнала для транзакций (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60;). | [optional] 
**BinlogGroupCommitSyncDelay** | Pointer to **string** | Задержка синхронизации групповой фиксации бинарного журнала в микросекундах (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60;). | [optional] 
**BinlogRowImage** | Pointer to **string** | Количество информации, записываемой в бинарный журнал при row-based репликации (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60;). | [optional] 
**BinlogRowsQueryLogEvents** | Pointer to **string** | Включение записи SQL-запросов в бинарный журнал при row-based репликации (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60;). | [optional] 
**CharacterSetServer** | Pointer to **string** | Кодировка по умолчанию для сервера MySQL (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60;). | [optional] 
**ExplicitDefaultsForTimestamp** | Pointer to **string** | Определяет автоматическое поведение TIMESTAMP без явных значений по умолчанию (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60;). | [optional] 
**GroupConcatMaxLen** | Pointer to **string** | Максимальная длина результата функции GROUP_CONCAT (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60;). | [optional] 
**InnodbAdaptiveHashIndex** | Pointer to **string** | Включение или отключение адаптивного хэш-индекса InnoDB для ускорения поиска по индексам (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60;). | [optional] 
**InnodbLockWaitTimeout** | Pointer to **string** | Время ожидания блокировки InnoDB перед завершением транзакции с ошибкой (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60;). | [optional] 
**InnodbNumaInterleave** | Pointer to **string** | Включение распределения памяти InnoDB между NUMA-узлами (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60;). | [optional] 
**NetReadTimeout** | Pointer to **string** | Время ожидания данных от клиента при чтении сетевого соединения (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60;). | [optional] 
**NetWriteTimeout** | Pointer to **string** | Время ожидания записи данных клиенту через сетевое соединение (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60;). | [optional] 
**RegexpTimeLimit** | Pointer to **string** | Максимальное время выполнения регулярных выражений (&#x60;mysql&#x60; | &#x60;mysql8_4&#x60;). | [optional] 
**SyncBinlog** | Pointer to **string** | Количество операций записи бинарного журнала перед принудительной синхронизацией на диск (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60;). | [optional] 
**TableDefinitionCache** | Pointer to **string** | Количество определений таблиц, хранящихся в кэше (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60;). | [optional] 
**LogBinTrustFunctionCreators** | Pointer to **string** | Разрешение создания хранимых функций без проверки бинарной регистрации (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60;). | [optional] 
**SkipNameResolve** | Pointer to **string** | Отключение DNS-разрешения имен клиентов при подключении к серверу (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60;). | [optional] 
**InnodbRedoLogCapacity** | Pointer to **string** | Общий размер redo log InnoDB для хранения журнала восстановления (&#x60;mysql8_4&#x60;). | [optional] 
**WaitTimeout** | Pointer to **string** | Время ожидания неактивного клиентского соединения перед закрытием (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60;). | [optional] 
**InteractiveTimeout** | Pointer to **string** | Время ожидания неактивного интерактивного соединения перед закрытием (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60;). | [optional] 
**DefaultTimeZone** | Pointer to **string** | Часовой пояс сервера MySQL по умолчанию (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60;). | [optional] 
**PxcStrictMode** | Pointer to **string** | Режим строгой проверки операций в Percona XtraDB Cluster (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60;). | [optional] 

## Methods

### NewMysql

`func NewMysql() *Mysql`

NewMysql instantiates a new Mysql object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlWithDefaults

`func NewMysqlWithDefaults() *Mysql`

NewMysqlWithDefaults instantiates a new Mysql object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJoinBufferSize

`func (o *Mysql) GetJoinBufferSize() string`

GetJoinBufferSize returns the JoinBufferSize field if non-nil, zero value otherwise.

### GetJoinBufferSizeOk

`func (o *Mysql) GetJoinBufferSizeOk() (*string, bool)`

GetJoinBufferSizeOk returns a tuple with the JoinBufferSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinBufferSize

`func (o *Mysql) SetJoinBufferSize(v string)`

SetJoinBufferSize sets JoinBufferSize field to given value.

### HasJoinBufferSize

`func (o *Mysql) HasJoinBufferSize() bool`

HasJoinBufferSize returns a boolean if a field has been set.

### GetMaxConnections

`func (o *Mysql) GetMaxConnections() string`

GetMaxConnections returns the MaxConnections field if non-nil, zero value otherwise.

### GetMaxConnectionsOk

`func (o *Mysql) GetMaxConnectionsOk() (*string, bool)`

GetMaxConnectionsOk returns a tuple with the MaxConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnections

`func (o *Mysql) SetMaxConnections(v string)`

SetMaxConnections sets MaxConnections field to given value.

### HasMaxConnections

`func (o *Mysql) HasMaxConnections() bool`

HasMaxConnections returns a boolean if a field has been set.

### GetSortBufferSize

`func (o *Mysql) GetSortBufferSize() string`

GetSortBufferSize returns the SortBufferSize field if non-nil, zero value otherwise.

### GetSortBufferSizeOk

`func (o *Mysql) GetSortBufferSizeOk() (*string, bool)`

GetSortBufferSizeOk returns a tuple with the SortBufferSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBufferSize

`func (o *Mysql) SetSortBufferSize(v string)`

SetSortBufferSize sets SortBufferSize field to given value.

### HasSortBufferSize

`func (o *Mysql) HasSortBufferSize() bool`

HasSortBufferSize returns a boolean if a field has been set.

### GetThreadCacheSize

`func (o *Mysql) GetThreadCacheSize() string`

GetThreadCacheSize returns the ThreadCacheSize field if non-nil, zero value otherwise.

### GetThreadCacheSizeOk

`func (o *Mysql) GetThreadCacheSizeOk() (*string, bool)`

GetThreadCacheSizeOk returns a tuple with the ThreadCacheSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadCacheSize

`func (o *Mysql) SetThreadCacheSize(v string)`

SetThreadCacheSize sets ThreadCacheSize field to given value.

### HasThreadCacheSize

`func (o *Mysql) HasThreadCacheSize() bool`

HasThreadCacheSize returns a boolean if a field has been set.

### GetInnodbBufferPoolSize

`func (o *Mysql) GetInnodbBufferPoolSize() string`

GetInnodbBufferPoolSize returns the InnodbBufferPoolSize field if non-nil, zero value otherwise.

### GetInnodbBufferPoolSizeOk

`func (o *Mysql) GetInnodbBufferPoolSizeOk() (*string, bool)`

GetInnodbBufferPoolSizeOk returns a tuple with the InnodbBufferPoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnodbBufferPoolSize

`func (o *Mysql) SetInnodbBufferPoolSize(v string)`

SetInnodbBufferPoolSize sets InnodbBufferPoolSize field to given value.

### HasInnodbBufferPoolSize

`func (o *Mysql) HasInnodbBufferPoolSize() bool`

HasInnodbBufferPoolSize returns a boolean if a field has been set.

### GetAutoIncrementIncrement

`func (o *Mysql) GetAutoIncrementIncrement() string`

GetAutoIncrementIncrement returns the AutoIncrementIncrement field if non-nil, zero value otherwise.

### GetAutoIncrementIncrementOk

`func (o *Mysql) GetAutoIncrementIncrementOk() (*string, bool)`

GetAutoIncrementIncrementOk returns a tuple with the AutoIncrementIncrement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoIncrementIncrement

`func (o *Mysql) SetAutoIncrementIncrement(v string)`

SetAutoIncrementIncrement sets AutoIncrementIncrement field to given value.

### HasAutoIncrementIncrement

`func (o *Mysql) HasAutoIncrementIncrement() bool`

HasAutoIncrementIncrement returns a boolean if a field has been set.

### GetAutoIncrementOffset

`func (o *Mysql) GetAutoIncrementOffset() string`

GetAutoIncrementOffset returns the AutoIncrementOffset field if non-nil, zero value otherwise.

### GetAutoIncrementOffsetOk

`func (o *Mysql) GetAutoIncrementOffsetOk() (*string, bool)`

GetAutoIncrementOffsetOk returns a tuple with the AutoIncrementOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoIncrementOffset

`func (o *Mysql) SetAutoIncrementOffset(v string)`

SetAutoIncrementOffset sets AutoIncrementOffset field to given value.

### HasAutoIncrementOffset

`func (o *Mysql) HasAutoIncrementOffset() bool`

HasAutoIncrementOffset returns a boolean if a field has been set.

### GetInnodbIoCapacity

`func (o *Mysql) GetInnodbIoCapacity() string`

GetInnodbIoCapacity returns the InnodbIoCapacity field if non-nil, zero value otherwise.

### GetInnodbIoCapacityOk

`func (o *Mysql) GetInnodbIoCapacityOk() (*string, bool)`

GetInnodbIoCapacityOk returns a tuple with the InnodbIoCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnodbIoCapacity

`func (o *Mysql) SetInnodbIoCapacity(v string)`

SetInnodbIoCapacity sets InnodbIoCapacity field to given value.

### HasInnodbIoCapacity

`func (o *Mysql) HasInnodbIoCapacity() bool`

HasInnodbIoCapacity returns a boolean if a field has been set.

### GetInnodbPurgeThreads

`func (o *Mysql) GetInnodbPurgeThreads() string`

GetInnodbPurgeThreads returns the InnodbPurgeThreads field if non-nil, zero value otherwise.

### GetInnodbPurgeThreadsOk

`func (o *Mysql) GetInnodbPurgeThreadsOk() (*string, bool)`

GetInnodbPurgeThreadsOk returns a tuple with the InnodbPurgeThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnodbPurgeThreads

`func (o *Mysql) SetInnodbPurgeThreads(v string)`

SetInnodbPurgeThreads sets InnodbPurgeThreads field to given value.

### HasInnodbPurgeThreads

`func (o *Mysql) HasInnodbPurgeThreads() bool`

HasInnodbPurgeThreads returns a boolean if a field has been set.

### GetInnodbReadIoThreads

`func (o *Mysql) GetInnodbReadIoThreads() string`

GetInnodbReadIoThreads returns the InnodbReadIoThreads field if non-nil, zero value otherwise.

### GetInnodbReadIoThreadsOk

`func (o *Mysql) GetInnodbReadIoThreadsOk() (*string, bool)`

GetInnodbReadIoThreadsOk returns a tuple with the InnodbReadIoThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnodbReadIoThreads

`func (o *Mysql) SetInnodbReadIoThreads(v string)`

SetInnodbReadIoThreads sets InnodbReadIoThreads field to given value.

### HasInnodbReadIoThreads

`func (o *Mysql) HasInnodbReadIoThreads() bool`

HasInnodbReadIoThreads returns a boolean if a field has been set.

### GetInnodbThreadConcurrency

`func (o *Mysql) GetInnodbThreadConcurrency() string`

GetInnodbThreadConcurrency returns the InnodbThreadConcurrency field if non-nil, zero value otherwise.

### GetInnodbThreadConcurrencyOk

`func (o *Mysql) GetInnodbThreadConcurrencyOk() (*string, bool)`

GetInnodbThreadConcurrencyOk returns a tuple with the InnodbThreadConcurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnodbThreadConcurrency

`func (o *Mysql) SetInnodbThreadConcurrency(v string)`

SetInnodbThreadConcurrency sets InnodbThreadConcurrency field to given value.

### HasInnodbThreadConcurrency

`func (o *Mysql) HasInnodbThreadConcurrency() bool`

HasInnodbThreadConcurrency returns a boolean if a field has been set.

### GetInnodbWriteIoThreads

`func (o *Mysql) GetInnodbWriteIoThreads() string`

GetInnodbWriteIoThreads returns the InnodbWriteIoThreads field if non-nil, zero value otherwise.

### GetInnodbWriteIoThreadsOk

`func (o *Mysql) GetInnodbWriteIoThreadsOk() (*string, bool)`

GetInnodbWriteIoThreadsOk returns a tuple with the InnodbWriteIoThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnodbWriteIoThreads

`func (o *Mysql) SetInnodbWriteIoThreads(v string)`

SetInnodbWriteIoThreads sets InnodbWriteIoThreads field to given value.

### HasInnodbWriteIoThreads

`func (o *Mysql) HasInnodbWriteIoThreads() bool`

HasInnodbWriteIoThreads returns a boolean if a field has been set.

### GetInnodbLogFileSize

`func (o *Mysql) GetInnodbLogFileSize() string`

GetInnodbLogFileSize returns the InnodbLogFileSize field if non-nil, zero value otherwise.

### GetInnodbLogFileSizeOk

`func (o *Mysql) GetInnodbLogFileSizeOk() (*string, bool)`

GetInnodbLogFileSizeOk returns a tuple with the InnodbLogFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnodbLogFileSize

`func (o *Mysql) SetInnodbLogFileSize(v string)`

SetInnodbLogFileSize sets InnodbLogFileSize field to given value.

### HasInnodbLogFileSize

`func (o *Mysql) HasInnodbLogFileSize() bool`

HasInnodbLogFileSize returns a boolean if a field has been set.

### GetMaxAllowedPacket

`func (o *Mysql) GetMaxAllowedPacket() string`

GetMaxAllowedPacket returns the MaxAllowedPacket field if non-nil, zero value otherwise.

### GetMaxAllowedPacketOk

`func (o *Mysql) GetMaxAllowedPacketOk() (*string, bool)`

GetMaxAllowedPacketOk returns a tuple with the MaxAllowedPacket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAllowedPacket

`func (o *Mysql) SetMaxAllowedPacket(v string)`

SetMaxAllowedPacket sets MaxAllowedPacket field to given value.

### HasMaxAllowedPacket

`func (o *Mysql) HasMaxAllowedPacket() bool`

HasMaxAllowedPacket returns a boolean if a field has been set.

### GetMaxHeapTableSize

`func (o *Mysql) GetMaxHeapTableSize() string`

GetMaxHeapTableSize returns the MaxHeapTableSize field if non-nil, zero value otherwise.

### GetMaxHeapTableSizeOk

`func (o *Mysql) GetMaxHeapTableSizeOk() (*string, bool)`

GetMaxHeapTableSizeOk returns a tuple with the MaxHeapTableSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHeapTableSize

`func (o *Mysql) SetMaxHeapTableSize(v string)`

SetMaxHeapTableSize sets MaxHeapTableSize field to given value.

### HasMaxHeapTableSize

`func (o *Mysql) HasMaxHeapTableSize() bool`

HasMaxHeapTableSize returns a boolean if a field has been set.

### GetSqlMode

`func (o *Mysql) GetSqlMode() string`

GetSqlMode returns the SqlMode field if non-nil, zero value otherwise.

### GetSqlModeOk

`func (o *Mysql) GetSqlModeOk() (*string, bool)`

GetSqlModeOk returns a tuple with the SqlMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlMode

`func (o *Mysql) SetSqlMode(v string)`

SetSqlMode sets SqlMode field to given value.

### HasSqlMode

`func (o *Mysql) HasSqlMode() bool`

HasSqlMode returns a boolean if a field has been set.

### GetQueryCacheType

`func (o *Mysql) GetQueryCacheType() string`

GetQueryCacheType returns the QueryCacheType field if non-nil, zero value otherwise.

### GetQueryCacheTypeOk

`func (o *Mysql) GetQueryCacheTypeOk() (*string, bool)`

GetQueryCacheTypeOk returns a tuple with the QueryCacheType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryCacheType

`func (o *Mysql) SetQueryCacheType(v string)`

SetQueryCacheType sets QueryCacheType field to given value.

### HasQueryCacheType

`func (o *Mysql) HasQueryCacheType() bool`

HasQueryCacheType returns a boolean if a field has been set.

### GetQueryCacheSize

`func (o *Mysql) GetQueryCacheSize() string`

GetQueryCacheSize returns the QueryCacheSize field if non-nil, zero value otherwise.

### GetQueryCacheSizeOk

`func (o *Mysql) GetQueryCacheSizeOk() (*string, bool)`

GetQueryCacheSizeOk returns a tuple with the QueryCacheSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryCacheSize

`func (o *Mysql) SetQueryCacheSize(v string)`

SetQueryCacheSize sets QueryCacheSize field to given value.

### HasQueryCacheSize

`func (o *Mysql) HasQueryCacheSize() bool`

HasQueryCacheSize returns a boolean if a field has been set.

### GetInnodbFlushLogAtTrxCommit

`func (o *Mysql) GetInnodbFlushLogAtTrxCommit() string`

GetInnodbFlushLogAtTrxCommit returns the InnodbFlushLogAtTrxCommit field if non-nil, zero value otherwise.

### GetInnodbFlushLogAtTrxCommitOk

`func (o *Mysql) GetInnodbFlushLogAtTrxCommitOk() (*string, bool)`

GetInnodbFlushLogAtTrxCommitOk returns a tuple with the InnodbFlushLogAtTrxCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnodbFlushLogAtTrxCommit

`func (o *Mysql) SetInnodbFlushLogAtTrxCommit(v string)`

SetInnodbFlushLogAtTrxCommit sets InnodbFlushLogAtTrxCommit field to given value.

### HasInnodbFlushLogAtTrxCommit

`func (o *Mysql) HasInnodbFlushLogAtTrxCommit() bool`

HasInnodbFlushLogAtTrxCommit returns a boolean if a field has been set.

### GetTransactionIsolation

`func (o *Mysql) GetTransactionIsolation() string`

GetTransactionIsolation returns the TransactionIsolation field if non-nil, zero value otherwise.

### GetTransactionIsolationOk

`func (o *Mysql) GetTransactionIsolationOk() (*string, bool)`

GetTransactionIsolationOk returns a tuple with the TransactionIsolation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIsolation

`func (o *Mysql) SetTransactionIsolation(v string)`

SetTransactionIsolation sets TransactionIsolation field to given value.

### HasTransactionIsolation

`func (o *Mysql) HasTransactionIsolation() bool`

HasTransactionIsolation returns a boolean if a field has been set.

### GetLongQueryTime

`func (o *Mysql) GetLongQueryTime() string`

GetLongQueryTime returns the LongQueryTime field if non-nil, zero value otherwise.

### GetLongQueryTimeOk

`func (o *Mysql) GetLongQueryTimeOk() (*string, bool)`

GetLongQueryTimeOk returns a tuple with the LongQueryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongQueryTime

`func (o *Mysql) SetLongQueryTime(v string)`

SetLongQueryTime sets LongQueryTime field to given value.

### HasLongQueryTime

`func (o *Mysql) HasLongQueryTime() bool`

HasLongQueryTime returns a boolean if a field has been set.

### GetTmpTableSize

`func (o *Mysql) GetTmpTableSize() string`

GetTmpTableSize returns the TmpTableSize field if non-nil, zero value otherwise.

### GetTmpTableSizeOk

`func (o *Mysql) GetTmpTableSizeOk() (*string, bool)`

GetTmpTableSizeOk returns a tuple with the TmpTableSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmpTableSize

`func (o *Mysql) SetTmpTableSize(v string)`

SetTmpTableSize sets TmpTableSize field to given value.

### HasTmpTableSize

`func (o *Mysql) HasTmpTableSize() bool`

HasTmpTableSize returns a boolean if a field has been set.

### GetTableOpenCache

`func (o *Mysql) GetTableOpenCache() string`

GetTableOpenCache returns the TableOpenCache field if non-nil, zero value otherwise.

### GetTableOpenCacheOk

`func (o *Mysql) GetTableOpenCacheOk() (*string, bool)`

GetTableOpenCacheOk returns a tuple with the TableOpenCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableOpenCache

`func (o *Mysql) SetTableOpenCache(v string)`

SetTableOpenCache sets TableOpenCache field to given value.

### HasTableOpenCache

`func (o *Mysql) HasTableOpenCache() bool`

HasTableOpenCache returns a boolean if a field has been set.

### GetTableOpenCacheInstances

`func (o *Mysql) GetTableOpenCacheInstances() string`

GetTableOpenCacheInstances returns the TableOpenCacheInstances field if non-nil, zero value otherwise.

### GetTableOpenCacheInstancesOk

`func (o *Mysql) GetTableOpenCacheInstancesOk() (*string, bool)`

GetTableOpenCacheInstancesOk returns a tuple with the TableOpenCacheInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableOpenCacheInstances

`func (o *Mysql) SetTableOpenCacheInstances(v string)`

SetTableOpenCacheInstances sets TableOpenCacheInstances field to given value.

### HasTableOpenCacheInstances

`func (o *Mysql) HasTableOpenCacheInstances() bool`

HasTableOpenCacheInstances returns a boolean if a field has been set.

### GetInnodbFlushMethod

`func (o *Mysql) GetInnodbFlushMethod() string`

GetInnodbFlushMethod returns the InnodbFlushMethod field if non-nil, zero value otherwise.

### GetInnodbFlushMethodOk

`func (o *Mysql) GetInnodbFlushMethodOk() (*string, bool)`

GetInnodbFlushMethodOk returns a tuple with the InnodbFlushMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnodbFlushMethod

`func (o *Mysql) SetInnodbFlushMethod(v string)`

SetInnodbFlushMethod sets InnodbFlushMethod field to given value.

### HasInnodbFlushMethod

`func (o *Mysql) HasInnodbFlushMethod() bool`

HasInnodbFlushMethod returns a boolean if a field has been set.

### GetInnodbStrictMode

`func (o *Mysql) GetInnodbStrictMode() string`

GetInnodbStrictMode returns the InnodbStrictMode field if non-nil, zero value otherwise.

### GetInnodbStrictModeOk

`func (o *Mysql) GetInnodbStrictModeOk() (*string, bool)`

GetInnodbStrictModeOk returns a tuple with the InnodbStrictMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnodbStrictMode

`func (o *Mysql) SetInnodbStrictMode(v string)`

SetInnodbStrictMode sets InnodbStrictMode field to given value.

### HasInnodbStrictMode

`func (o *Mysql) HasInnodbStrictMode() bool`

HasInnodbStrictMode returns a boolean if a field has been set.

### GetSlowQueryLog

`func (o *Mysql) GetSlowQueryLog() string`

GetSlowQueryLog returns the SlowQueryLog field if non-nil, zero value otherwise.

### GetSlowQueryLogOk

`func (o *Mysql) GetSlowQueryLogOk() (*string, bool)`

GetSlowQueryLogOk returns a tuple with the SlowQueryLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlowQueryLog

`func (o *Mysql) SetSlowQueryLog(v string)`

SetSlowQueryLog sets SlowQueryLog field to given value.

### HasSlowQueryLog

`func (o *Mysql) HasSlowQueryLog() bool`

HasSlowQueryLog returns a boolean if a field has been set.

### GetBinlogCacheSize

`func (o *Mysql) GetBinlogCacheSize() string`

GetBinlogCacheSize returns the BinlogCacheSize field if non-nil, zero value otherwise.

### GetBinlogCacheSizeOk

`func (o *Mysql) GetBinlogCacheSizeOk() (*string, bool)`

GetBinlogCacheSizeOk returns a tuple with the BinlogCacheSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinlogCacheSize

`func (o *Mysql) SetBinlogCacheSize(v string)`

SetBinlogCacheSize sets BinlogCacheSize field to given value.

### HasBinlogCacheSize

`func (o *Mysql) HasBinlogCacheSize() bool`

HasBinlogCacheSize returns a boolean if a field has been set.

### GetBinlogGroupCommitSyncDelay

`func (o *Mysql) GetBinlogGroupCommitSyncDelay() string`

GetBinlogGroupCommitSyncDelay returns the BinlogGroupCommitSyncDelay field if non-nil, zero value otherwise.

### GetBinlogGroupCommitSyncDelayOk

`func (o *Mysql) GetBinlogGroupCommitSyncDelayOk() (*string, bool)`

GetBinlogGroupCommitSyncDelayOk returns a tuple with the BinlogGroupCommitSyncDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinlogGroupCommitSyncDelay

`func (o *Mysql) SetBinlogGroupCommitSyncDelay(v string)`

SetBinlogGroupCommitSyncDelay sets BinlogGroupCommitSyncDelay field to given value.

### HasBinlogGroupCommitSyncDelay

`func (o *Mysql) HasBinlogGroupCommitSyncDelay() bool`

HasBinlogGroupCommitSyncDelay returns a boolean if a field has been set.

### GetBinlogRowImage

`func (o *Mysql) GetBinlogRowImage() string`

GetBinlogRowImage returns the BinlogRowImage field if non-nil, zero value otherwise.

### GetBinlogRowImageOk

`func (o *Mysql) GetBinlogRowImageOk() (*string, bool)`

GetBinlogRowImageOk returns a tuple with the BinlogRowImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinlogRowImage

`func (o *Mysql) SetBinlogRowImage(v string)`

SetBinlogRowImage sets BinlogRowImage field to given value.

### HasBinlogRowImage

`func (o *Mysql) HasBinlogRowImage() bool`

HasBinlogRowImage returns a boolean if a field has been set.

### GetBinlogRowsQueryLogEvents

`func (o *Mysql) GetBinlogRowsQueryLogEvents() string`

GetBinlogRowsQueryLogEvents returns the BinlogRowsQueryLogEvents field if non-nil, zero value otherwise.

### GetBinlogRowsQueryLogEventsOk

`func (o *Mysql) GetBinlogRowsQueryLogEventsOk() (*string, bool)`

GetBinlogRowsQueryLogEventsOk returns a tuple with the BinlogRowsQueryLogEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinlogRowsQueryLogEvents

`func (o *Mysql) SetBinlogRowsQueryLogEvents(v string)`

SetBinlogRowsQueryLogEvents sets BinlogRowsQueryLogEvents field to given value.

### HasBinlogRowsQueryLogEvents

`func (o *Mysql) HasBinlogRowsQueryLogEvents() bool`

HasBinlogRowsQueryLogEvents returns a boolean if a field has been set.

### GetCharacterSetServer

`func (o *Mysql) GetCharacterSetServer() string`

GetCharacterSetServer returns the CharacterSetServer field if non-nil, zero value otherwise.

### GetCharacterSetServerOk

`func (o *Mysql) GetCharacterSetServerOk() (*string, bool)`

GetCharacterSetServerOk returns a tuple with the CharacterSetServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacterSetServer

`func (o *Mysql) SetCharacterSetServer(v string)`

SetCharacterSetServer sets CharacterSetServer field to given value.

### HasCharacterSetServer

`func (o *Mysql) HasCharacterSetServer() bool`

HasCharacterSetServer returns a boolean if a field has been set.

### GetExplicitDefaultsForTimestamp

`func (o *Mysql) GetExplicitDefaultsForTimestamp() string`

GetExplicitDefaultsForTimestamp returns the ExplicitDefaultsForTimestamp field if non-nil, zero value otherwise.

### GetExplicitDefaultsForTimestampOk

`func (o *Mysql) GetExplicitDefaultsForTimestampOk() (*string, bool)`

GetExplicitDefaultsForTimestampOk returns a tuple with the ExplicitDefaultsForTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitDefaultsForTimestamp

`func (o *Mysql) SetExplicitDefaultsForTimestamp(v string)`

SetExplicitDefaultsForTimestamp sets ExplicitDefaultsForTimestamp field to given value.

### HasExplicitDefaultsForTimestamp

`func (o *Mysql) HasExplicitDefaultsForTimestamp() bool`

HasExplicitDefaultsForTimestamp returns a boolean if a field has been set.

### GetGroupConcatMaxLen

`func (o *Mysql) GetGroupConcatMaxLen() string`

GetGroupConcatMaxLen returns the GroupConcatMaxLen field if non-nil, zero value otherwise.

### GetGroupConcatMaxLenOk

`func (o *Mysql) GetGroupConcatMaxLenOk() (*string, bool)`

GetGroupConcatMaxLenOk returns a tuple with the GroupConcatMaxLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupConcatMaxLen

`func (o *Mysql) SetGroupConcatMaxLen(v string)`

SetGroupConcatMaxLen sets GroupConcatMaxLen field to given value.

### HasGroupConcatMaxLen

`func (o *Mysql) HasGroupConcatMaxLen() bool`

HasGroupConcatMaxLen returns a boolean if a field has been set.

### GetInnodbAdaptiveHashIndex

`func (o *Mysql) GetInnodbAdaptiveHashIndex() string`

GetInnodbAdaptiveHashIndex returns the InnodbAdaptiveHashIndex field if non-nil, zero value otherwise.

### GetInnodbAdaptiveHashIndexOk

`func (o *Mysql) GetInnodbAdaptiveHashIndexOk() (*string, bool)`

GetInnodbAdaptiveHashIndexOk returns a tuple with the InnodbAdaptiveHashIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnodbAdaptiveHashIndex

`func (o *Mysql) SetInnodbAdaptiveHashIndex(v string)`

SetInnodbAdaptiveHashIndex sets InnodbAdaptiveHashIndex field to given value.

### HasInnodbAdaptiveHashIndex

`func (o *Mysql) HasInnodbAdaptiveHashIndex() bool`

HasInnodbAdaptiveHashIndex returns a boolean if a field has been set.

### GetInnodbLockWaitTimeout

`func (o *Mysql) GetInnodbLockWaitTimeout() string`

GetInnodbLockWaitTimeout returns the InnodbLockWaitTimeout field if non-nil, zero value otherwise.

### GetInnodbLockWaitTimeoutOk

`func (o *Mysql) GetInnodbLockWaitTimeoutOk() (*string, bool)`

GetInnodbLockWaitTimeoutOk returns a tuple with the InnodbLockWaitTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnodbLockWaitTimeout

`func (o *Mysql) SetInnodbLockWaitTimeout(v string)`

SetInnodbLockWaitTimeout sets InnodbLockWaitTimeout field to given value.

### HasInnodbLockWaitTimeout

`func (o *Mysql) HasInnodbLockWaitTimeout() bool`

HasInnodbLockWaitTimeout returns a boolean if a field has been set.

### GetInnodbNumaInterleave

`func (o *Mysql) GetInnodbNumaInterleave() string`

GetInnodbNumaInterleave returns the InnodbNumaInterleave field if non-nil, zero value otherwise.

### GetInnodbNumaInterleaveOk

`func (o *Mysql) GetInnodbNumaInterleaveOk() (*string, bool)`

GetInnodbNumaInterleaveOk returns a tuple with the InnodbNumaInterleave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnodbNumaInterleave

`func (o *Mysql) SetInnodbNumaInterleave(v string)`

SetInnodbNumaInterleave sets InnodbNumaInterleave field to given value.

### HasInnodbNumaInterleave

`func (o *Mysql) HasInnodbNumaInterleave() bool`

HasInnodbNumaInterleave returns a boolean if a field has been set.

### GetNetReadTimeout

`func (o *Mysql) GetNetReadTimeout() string`

GetNetReadTimeout returns the NetReadTimeout field if non-nil, zero value otherwise.

### GetNetReadTimeoutOk

`func (o *Mysql) GetNetReadTimeoutOk() (*string, bool)`

GetNetReadTimeoutOk returns a tuple with the NetReadTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetReadTimeout

`func (o *Mysql) SetNetReadTimeout(v string)`

SetNetReadTimeout sets NetReadTimeout field to given value.

### HasNetReadTimeout

`func (o *Mysql) HasNetReadTimeout() bool`

HasNetReadTimeout returns a boolean if a field has been set.

### GetNetWriteTimeout

`func (o *Mysql) GetNetWriteTimeout() string`

GetNetWriteTimeout returns the NetWriteTimeout field if non-nil, zero value otherwise.

### GetNetWriteTimeoutOk

`func (o *Mysql) GetNetWriteTimeoutOk() (*string, bool)`

GetNetWriteTimeoutOk returns a tuple with the NetWriteTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetWriteTimeout

`func (o *Mysql) SetNetWriteTimeout(v string)`

SetNetWriteTimeout sets NetWriteTimeout field to given value.

### HasNetWriteTimeout

`func (o *Mysql) HasNetWriteTimeout() bool`

HasNetWriteTimeout returns a boolean if a field has been set.

### GetRegexpTimeLimit

`func (o *Mysql) GetRegexpTimeLimit() string`

GetRegexpTimeLimit returns the RegexpTimeLimit field if non-nil, zero value otherwise.

### GetRegexpTimeLimitOk

`func (o *Mysql) GetRegexpTimeLimitOk() (*string, bool)`

GetRegexpTimeLimitOk returns a tuple with the RegexpTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegexpTimeLimit

`func (o *Mysql) SetRegexpTimeLimit(v string)`

SetRegexpTimeLimit sets RegexpTimeLimit field to given value.

### HasRegexpTimeLimit

`func (o *Mysql) HasRegexpTimeLimit() bool`

HasRegexpTimeLimit returns a boolean if a field has been set.

### GetSyncBinlog

`func (o *Mysql) GetSyncBinlog() string`

GetSyncBinlog returns the SyncBinlog field if non-nil, zero value otherwise.

### GetSyncBinlogOk

`func (o *Mysql) GetSyncBinlogOk() (*string, bool)`

GetSyncBinlogOk returns a tuple with the SyncBinlog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncBinlog

`func (o *Mysql) SetSyncBinlog(v string)`

SetSyncBinlog sets SyncBinlog field to given value.

### HasSyncBinlog

`func (o *Mysql) HasSyncBinlog() bool`

HasSyncBinlog returns a boolean if a field has been set.

### GetTableDefinitionCache

`func (o *Mysql) GetTableDefinitionCache() string`

GetTableDefinitionCache returns the TableDefinitionCache field if non-nil, zero value otherwise.

### GetTableDefinitionCacheOk

`func (o *Mysql) GetTableDefinitionCacheOk() (*string, bool)`

GetTableDefinitionCacheOk returns a tuple with the TableDefinitionCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableDefinitionCache

`func (o *Mysql) SetTableDefinitionCache(v string)`

SetTableDefinitionCache sets TableDefinitionCache field to given value.

### HasTableDefinitionCache

`func (o *Mysql) HasTableDefinitionCache() bool`

HasTableDefinitionCache returns a boolean if a field has been set.

### GetLogBinTrustFunctionCreators

`func (o *Mysql) GetLogBinTrustFunctionCreators() string`

GetLogBinTrustFunctionCreators returns the LogBinTrustFunctionCreators field if non-nil, zero value otherwise.

### GetLogBinTrustFunctionCreatorsOk

`func (o *Mysql) GetLogBinTrustFunctionCreatorsOk() (*string, bool)`

GetLogBinTrustFunctionCreatorsOk returns a tuple with the LogBinTrustFunctionCreators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogBinTrustFunctionCreators

`func (o *Mysql) SetLogBinTrustFunctionCreators(v string)`

SetLogBinTrustFunctionCreators sets LogBinTrustFunctionCreators field to given value.

### HasLogBinTrustFunctionCreators

`func (o *Mysql) HasLogBinTrustFunctionCreators() bool`

HasLogBinTrustFunctionCreators returns a boolean if a field has been set.

### GetSkipNameResolve

`func (o *Mysql) GetSkipNameResolve() string`

GetSkipNameResolve returns the SkipNameResolve field if non-nil, zero value otherwise.

### GetSkipNameResolveOk

`func (o *Mysql) GetSkipNameResolveOk() (*string, bool)`

GetSkipNameResolveOk returns a tuple with the SkipNameResolve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipNameResolve

`func (o *Mysql) SetSkipNameResolve(v string)`

SetSkipNameResolve sets SkipNameResolve field to given value.

### HasSkipNameResolve

`func (o *Mysql) HasSkipNameResolve() bool`

HasSkipNameResolve returns a boolean if a field has been set.

### GetInnodbRedoLogCapacity

`func (o *Mysql) GetInnodbRedoLogCapacity() string`

GetInnodbRedoLogCapacity returns the InnodbRedoLogCapacity field if non-nil, zero value otherwise.

### GetInnodbRedoLogCapacityOk

`func (o *Mysql) GetInnodbRedoLogCapacityOk() (*string, bool)`

GetInnodbRedoLogCapacityOk returns a tuple with the InnodbRedoLogCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnodbRedoLogCapacity

`func (o *Mysql) SetInnodbRedoLogCapacity(v string)`

SetInnodbRedoLogCapacity sets InnodbRedoLogCapacity field to given value.

### HasInnodbRedoLogCapacity

`func (o *Mysql) HasInnodbRedoLogCapacity() bool`

HasInnodbRedoLogCapacity returns a boolean if a field has been set.

### GetWaitTimeout

`func (o *Mysql) GetWaitTimeout() string`

GetWaitTimeout returns the WaitTimeout field if non-nil, zero value otherwise.

### GetWaitTimeoutOk

`func (o *Mysql) GetWaitTimeoutOk() (*string, bool)`

GetWaitTimeoutOk returns a tuple with the WaitTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitTimeout

`func (o *Mysql) SetWaitTimeout(v string)`

SetWaitTimeout sets WaitTimeout field to given value.

### HasWaitTimeout

`func (o *Mysql) HasWaitTimeout() bool`

HasWaitTimeout returns a boolean if a field has been set.

### GetInteractiveTimeout

`func (o *Mysql) GetInteractiveTimeout() string`

GetInteractiveTimeout returns the InteractiveTimeout field if non-nil, zero value otherwise.

### GetInteractiveTimeoutOk

`func (o *Mysql) GetInteractiveTimeoutOk() (*string, bool)`

GetInteractiveTimeoutOk returns a tuple with the InteractiveTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractiveTimeout

`func (o *Mysql) SetInteractiveTimeout(v string)`

SetInteractiveTimeout sets InteractiveTimeout field to given value.

### HasInteractiveTimeout

`func (o *Mysql) HasInteractiveTimeout() bool`

HasInteractiveTimeout returns a boolean if a field has been set.

### GetDefaultTimeZone

`func (o *Mysql) GetDefaultTimeZone() string`

GetDefaultTimeZone returns the DefaultTimeZone field if non-nil, zero value otherwise.

### GetDefaultTimeZoneOk

`func (o *Mysql) GetDefaultTimeZoneOk() (*string, bool)`

GetDefaultTimeZoneOk returns a tuple with the DefaultTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTimeZone

`func (o *Mysql) SetDefaultTimeZone(v string)`

SetDefaultTimeZone sets DefaultTimeZone field to given value.

### HasDefaultTimeZone

`func (o *Mysql) HasDefaultTimeZone() bool`

HasDefaultTimeZone returns a boolean if a field has been set.

### GetPxcStrictMode

`func (o *Mysql) GetPxcStrictMode() string`

GetPxcStrictMode returns the PxcStrictMode field if non-nil, zero value otherwise.

### GetPxcStrictModeOk

`func (o *Mysql) GetPxcStrictModeOk() (*string, bool)`

GetPxcStrictModeOk returns a tuple with the PxcStrictMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPxcStrictMode

`func (o *Mysql) SetPxcStrictMode(v string)`

SetPxcStrictMode sets PxcStrictMode field to given value.

### HasPxcStrictMode

`func (o *Mysql) HasPxcStrictMode() bool`

HasPxcStrictMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


