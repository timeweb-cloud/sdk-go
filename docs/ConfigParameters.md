# ConfigParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoIncrementIncrement** | Pointer to **string** | Интервал между значениями столбцов с атрибутом &#x60;AUTO_INCREMENT&#x60; (&#x60;mysql5&#x60; | &#x60;mysql&#x60;). | [optional] 
**AutoIncrementOffset** | Pointer to **string** | Начальное значение для столбцов с атрибутом &#x60;AUTO_INCREMENT&#x60; (&#x60;mysql5&#x60; | &#x60;mysql&#x60;). | [optional] 
**InnodbIoCapacity** | Pointer to **string** | Количество операций ввода-вывода в секунду &#x60;IOPS&#x60; (&#x60;mysql5&#x60; | &#x60;mysql&#x60;). | [optional] 
**InnodbPurgeThreads** | Pointer to **string** | Количество потоков ввода-вывода, используемых для операций очистки (&#x60;mysql5&#x60; | &#x60;mysql&#x60;). | [optional] 
**InnodbReadIoThreads** | Pointer to **string** | Количество потоков ввода-вывода, используемых для операций чтения (&#x60;mysql5&#x60; | &#x60;mysql&#x60;). | [optional] 
**InnodbThreadConcurrency** | Pointer to **string** | Максимальное число потоков, которые могут исполняться (&#x60;mysql5&#x60; | &#x60;mysql&#x60;). | [optional] 
**InnodbWriteIoThreads** | Pointer to **string** | Количество потоков ввода-вывода, используемых для операций записи (&#x60;mysql5&#x60; | &#x60;mysql&#x60;). | [optional] 
**JoinBufferSize** | Pointer to **string** | Минимальный размер буфера (&#x60;mysql5&#x60; | &#x60;mysql&#x60;). | [optional] 
**MaxAllowedPacket** | Pointer to **string** | Максимальный размер одного пакета, строки или параметра, отправляемого функцией &#x60;mysql_stmt_send_long_data()&#x60; (&#x60;mysql5&#x60; | &#x60;mysql&#x60;). | [optional] 
**MaxHeapTableSize** | Pointer to **string** | Максимальный размер пользовательских MEMORY-таблиц (&#x60;mysql5&#x60; | &#x60;mysql&#x60;). | [optional] 
**AutovacuumAnalyzeScaleFactor** | Pointer to **string** | Доля измененных или удаленных записей в таблице, при которой процесс автоочистки выполнит команду &#x60;ANALYZE&#x60; (&#x60;postgres&#x60; | &#x60;postgres14&#x60;| &#x60;postgres15&#x60;). | [optional] 
**BgwriterDelay** | Pointer to **string** | Задержка между запусками процесса фоновой записи (&#x60;postgres&#x60; | &#x60;postgres14&#x60;| &#x60;postgres15&#x60;). | [optional] 
**BgwriterLruMaxpages** | Pointer to **string** | Максимальное число элементов буферного кеша (&#x60;postgres&#x60; | &#x60;postgres14&#x60;| &#x60;postgres15&#x60;). | [optional] 
**DeadlockTimeout** | Pointer to **string** | Время ожидания, по истечении которого будет выполняться проверка состояния перекрестной блокировки (&#x60;postgres&#x60; | &#x60;postgres14&#x60;| &#x60;postgres15&#x60;). | [optional] 
**GinPendingListLimit** | Pointer to **string** | Максимальный размер очереди записей индекса &#x60;GIN&#x60; (&#x60;postgres&#x60; | &#x60;postgres14&#x60;| &#x60;postgres15&#x60;). | [optional] 
**IdleInTransactionSessionTimeout** | Pointer to **string** | Время простоя открытой транзакции, при превышении которого будет завершена сессия с этой транзакцией (&#x60;postgres&#x60; | &#x60;postgres14&#x60;| &#x60;postgres15&#x60;). | [optional] 
**IdleSessionTimeout** | Pointer to **string** | Время простоя не открытой транзакции, при превышении которого будет завершена сессия с этой транзакцией (&#x60;postgres&#x60; | &#x60;postgres14&#x60;| &#x60;postgres15&#x60;). | [optional] 
**JoinCollapseLimit** | Pointer to **string** | Значение количества элементов в списке &#x60;FROM&#x60; при превышении которого, планировщик будет переносить в список явные инструкции &#x60;JOIN&#x60; (&#x60;postgres&#x60; | &#x60;postgres14&#x60;| &#x60;postgres15&#x60;). | [optional] 
**LockTimeout** | Pointer to **string** | Время ожидания освобождения блокировки (&#x60;postgres&#x60; | &#x60;postgres14&#x60;| &#x60;postgres15&#x60;). | [optional] 
**MaxPreparedTransactions** | Pointer to **string** | Максимальное число транзакций, которые могут одновременно находиться в подготовленном состоянии (&#x60;postgres&#x60; | &#x60;postgres14&#x60;| &#x60;postgres15&#x60;). | [optional] 
**MaxConnections** | Pointer to **string** | Допустимое количество соединений (&#x60;postgres&#x60; | &#x60;postgres14&#x60;| &#x60;postgres15&#x60; | &#x60;mysql&#x60;). | [optional] 
**SharedBuffers** | Pointer to **string** | Устанавливает количество буферов общей памяти, используемых сервером (&#x60;postgres&#x60; | &#x60;postgres14&#x60;| &#x60;postgres15&#x60;). | [optional] 
**WalBuffers** | Pointer to **string** | Устанавливает количество буферов дисковых страниц в общей памяти для WAL (&#x60;postgres&#x60; | &#x60;postgres14&#x60;| &#x60;postgres15&#x60;). | [optional] 
**TempBuffers** | Pointer to **string** | Устанавливает максимальное количество временных буферов, используемых каждой сессией (&#x60;postgres&#x60; | &#x60;postgres14&#x60;| &#x60;postgres15&#x60;). | [optional] 
**WorkMem** | Pointer to **string** | Устанавливает максимальное количество памяти, используемое для рабочих пространств запросов (&#x60;postgres&#x60; | &#x60;postgres14&#x60;| &#x60;postgres15&#x60;). | [optional] 
**SqlMode** | Pointer to **string** | Устанавливает режим SQL. Можно задать несколько режимов, разделяя их запятой. (&#x60;mysql&#x60;). | [optional] 
**QueryCacheType** | Pointer to **string** | Параметр включает или отключает работу MySQL Query Cache (&#x60;mysql&#x60;). | [optional] 
**QueryCacheSize** | Pointer to **string** | Размер в байтах, доступный для кэша запросов (&#x60;mysql&#x60;). | [optional] 

## Methods

### NewConfigParameters

`func NewConfigParameters() *ConfigParameters`

NewConfigParameters instantiates a new ConfigParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigParametersWithDefaults

`func NewConfigParametersWithDefaults() *ConfigParameters`

NewConfigParametersWithDefaults instantiates a new ConfigParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoIncrementIncrement

`func (o *ConfigParameters) GetAutoIncrementIncrement() string`

GetAutoIncrementIncrement returns the AutoIncrementIncrement field if non-nil, zero value otherwise.

### GetAutoIncrementIncrementOk

`func (o *ConfigParameters) GetAutoIncrementIncrementOk() (*string, bool)`

GetAutoIncrementIncrementOk returns a tuple with the AutoIncrementIncrement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoIncrementIncrement

`func (o *ConfigParameters) SetAutoIncrementIncrement(v string)`

SetAutoIncrementIncrement sets AutoIncrementIncrement field to given value.

### HasAutoIncrementIncrement

`func (o *ConfigParameters) HasAutoIncrementIncrement() bool`

HasAutoIncrementIncrement returns a boolean if a field has been set.

### GetAutoIncrementOffset

`func (o *ConfigParameters) GetAutoIncrementOffset() string`

GetAutoIncrementOffset returns the AutoIncrementOffset field if non-nil, zero value otherwise.

### GetAutoIncrementOffsetOk

`func (o *ConfigParameters) GetAutoIncrementOffsetOk() (*string, bool)`

GetAutoIncrementOffsetOk returns a tuple with the AutoIncrementOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoIncrementOffset

`func (o *ConfigParameters) SetAutoIncrementOffset(v string)`

SetAutoIncrementOffset sets AutoIncrementOffset field to given value.

### HasAutoIncrementOffset

`func (o *ConfigParameters) HasAutoIncrementOffset() bool`

HasAutoIncrementOffset returns a boolean if a field has been set.

### GetInnodbIoCapacity

`func (o *ConfigParameters) GetInnodbIoCapacity() string`

GetInnodbIoCapacity returns the InnodbIoCapacity field if non-nil, zero value otherwise.

### GetInnodbIoCapacityOk

`func (o *ConfigParameters) GetInnodbIoCapacityOk() (*string, bool)`

GetInnodbIoCapacityOk returns a tuple with the InnodbIoCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnodbIoCapacity

`func (o *ConfigParameters) SetInnodbIoCapacity(v string)`

SetInnodbIoCapacity sets InnodbIoCapacity field to given value.

### HasInnodbIoCapacity

`func (o *ConfigParameters) HasInnodbIoCapacity() bool`

HasInnodbIoCapacity returns a boolean if a field has been set.

### GetInnodbPurgeThreads

`func (o *ConfigParameters) GetInnodbPurgeThreads() string`

GetInnodbPurgeThreads returns the InnodbPurgeThreads field if non-nil, zero value otherwise.

### GetInnodbPurgeThreadsOk

`func (o *ConfigParameters) GetInnodbPurgeThreadsOk() (*string, bool)`

GetInnodbPurgeThreadsOk returns a tuple with the InnodbPurgeThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnodbPurgeThreads

`func (o *ConfigParameters) SetInnodbPurgeThreads(v string)`

SetInnodbPurgeThreads sets InnodbPurgeThreads field to given value.

### HasInnodbPurgeThreads

`func (o *ConfigParameters) HasInnodbPurgeThreads() bool`

HasInnodbPurgeThreads returns a boolean if a field has been set.

### GetInnodbReadIoThreads

`func (o *ConfigParameters) GetInnodbReadIoThreads() string`

GetInnodbReadIoThreads returns the InnodbReadIoThreads field if non-nil, zero value otherwise.

### GetInnodbReadIoThreadsOk

`func (o *ConfigParameters) GetInnodbReadIoThreadsOk() (*string, bool)`

GetInnodbReadIoThreadsOk returns a tuple with the InnodbReadIoThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnodbReadIoThreads

`func (o *ConfigParameters) SetInnodbReadIoThreads(v string)`

SetInnodbReadIoThreads sets InnodbReadIoThreads field to given value.

### HasInnodbReadIoThreads

`func (o *ConfigParameters) HasInnodbReadIoThreads() bool`

HasInnodbReadIoThreads returns a boolean if a field has been set.

### GetInnodbThreadConcurrency

`func (o *ConfigParameters) GetInnodbThreadConcurrency() string`

GetInnodbThreadConcurrency returns the InnodbThreadConcurrency field if non-nil, zero value otherwise.

### GetInnodbThreadConcurrencyOk

`func (o *ConfigParameters) GetInnodbThreadConcurrencyOk() (*string, bool)`

GetInnodbThreadConcurrencyOk returns a tuple with the InnodbThreadConcurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnodbThreadConcurrency

`func (o *ConfigParameters) SetInnodbThreadConcurrency(v string)`

SetInnodbThreadConcurrency sets InnodbThreadConcurrency field to given value.

### HasInnodbThreadConcurrency

`func (o *ConfigParameters) HasInnodbThreadConcurrency() bool`

HasInnodbThreadConcurrency returns a boolean if a field has been set.

### GetInnodbWriteIoThreads

`func (o *ConfigParameters) GetInnodbWriteIoThreads() string`

GetInnodbWriteIoThreads returns the InnodbWriteIoThreads field if non-nil, zero value otherwise.

### GetInnodbWriteIoThreadsOk

`func (o *ConfigParameters) GetInnodbWriteIoThreadsOk() (*string, bool)`

GetInnodbWriteIoThreadsOk returns a tuple with the InnodbWriteIoThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnodbWriteIoThreads

`func (o *ConfigParameters) SetInnodbWriteIoThreads(v string)`

SetInnodbWriteIoThreads sets InnodbWriteIoThreads field to given value.

### HasInnodbWriteIoThreads

`func (o *ConfigParameters) HasInnodbWriteIoThreads() bool`

HasInnodbWriteIoThreads returns a boolean if a field has been set.

### GetJoinBufferSize

`func (o *ConfigParameters) GetJoinBufferSize() string`

GetJoinBufferSize returns the JoinBufferSize field if non-nil, zero value otherwise.

### GetJoinBufferSizeOk

`func (o *ConfigParameters) GetJoinBufferSizeOk() (*string, bool)`

GetJoinBufferSizeOk returns a tuple with the JoinBufferSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinBufferSize

`func (o *ConfigParameters) SetJoinBufferSize(v string)`

SetJoinBufferSize sets JoinBufferSize field to given value.

### HasJoinBufferSize

`func (o *ConfigParameters) HasJoinBufferSize() bool`

HasJoinBufferSize returns a boolean if a field has been set.

### GetMaxAllowedPacket

`func (o *ConfigParameters) GetMaxAllowedPacket() string`

GetMaxAllowedPacket returns the MaxAllowedPacket field if non-nil, zero value otherwise.

### GetMaxAllowedPacketOk

`func (o *ConfigParameters) GetMaxAllowedPacketOk() (*string, bool)`

GetMaxAllowedPacketOk returns a tuple with the MaxAllowedPacket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAllowedPacket

`func (o *ConfigParameters) SetMaxAllowedPacket(v string)`

SetMaxAllowedPacket sets MaxAllowedPacket field to given value.

### HasMaxAllowedPacket

`func (o *ConfigParameters) HasMaxAllowedPacket() bool`

HasMaxAllowedPacket returns a boolean if a field has been set.

### GetMaxHeapTableSize

`func (o *ConfigParameters) GetMaxHeapTableSize() string`

GetMaxHeapTableSize returns the MaxHeapTableSize field if non-nil, zero value otherwise.

### GetMaxHeapTableSizeOk

`func (o *ConfigParameters) GetMaxHeapTableSizeOk() (*string, bool)`

GetMaxHeapTableSizeOk returns a tuple with the MaxHeapTableSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHeapTableSize

`func (o *ConfigParameters) SetMaxHeapTableSize(v string)`

SetMaxHeapTableSize sets MaxHeapTableSize field to given value.

### HasMaxHeapTableSize

`func (o *ConfigParameters) HasMaxHeapTableSize() bool`

HasMaxHeapTableSize returns a boolean if a field has been set.

### GetAutovacuumAnalyzeScaleFactor

`func (o *ConfigParameters) GetAutovacuumAnalyzeScaleFactor() string`

GetAutovacuumAnalyzeScaleFactor returns the AutovacuumAnalyzeScaleFactor field if non-nil, zero value otherwise.

### GetAutovacuumAnalyzeScaleFactorOk

`func (o *ConfigParameters) GetAutovacuumAnalyzeScaleFactorOk() (*string, bool)`

GetAutovacuumAnalyzeScaleFactorOk returns a tuple with the AutovacuumAnalyzeScaleFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutovacuumAnalyzeScaleFactor

`func (o *ConfigParameters) SetAutovacuumAnalyzeScaleFactor(v string)`

SetAutovacuumAnalyzeScaleFactor sets AutovacuumAnalyzeScaleFactor field to given value.

### HasAutovacuumAnalyzeScaleFactor

`func (o *ConfigParameters) HasAutovacuumAnalyzeScaleFactor() bool`

HasAutovacuumAnalyzeScaleFactor returns a boolean if a field has been set.

### GetBgwriterDelay

`func (o *ConfigParameters) GetBgwriterDelay() string`

GetBgwriterDelay returns the BgwriterDelay field if non-nil, zero value otherwise.

### GetBgwriterDelayOk

`func (o *ConfigParameters) GetBgwriterDelayOk() (*string, bool)`

GetBgwriterDelayOk returns a tuple with the BgwriterDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgwriterDelay

`func (o *ConfigParameters) SetBgwriterDelay(v string)`

SetBgwriterDelay sets BgwriterDelay field to given value.

### HasBgwriterDelay

`func (o *ConfigParameters) HasBgwriterDelay() bool`

HasBgwriterDelay returns a boolean if a field has been set.

### GetBgwriterLruMaxpages

`func (o *ConfigParameters) GetBgwriterLruMaxpages() string`

GetBgwriterLruMaxpages returns the BgwriterLruMaxpages field if non-nil, zero value otherwise.

### GetBgwriterLruMaxpagesOk

`func (o *ConfigParameters) GetBgwriterLruMaxpagesOk() (*string, bool)`

GetBgwriterLruMaxpagesOk returns a tuple with the BgwriterLruMaxpages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgwriterLruMaxpages

`func (o *ConfigParameters) SetBgwriterLruMaxpages(v string)`

SetBgwriterLruMaxpages sets BgwriterLruMaxpages field to given value.

### HasBgwriterLruMaxpages

`func (o *ConfigParameters) HasBgwriterLruMaxpages() bool`

HasBgwriterLruMaxpages returns a boolean if a field has been set.

### GetDeadlockTimeout

`func (o *ConfigParameters) GetDeadlockTimeout() string`

GetDeadlockTimeout returns the DeadlockTimeout field if non-nil, zero value otherwise.

### GetDeadlockTimeoutOk

`func (o *ConfigParameters) GetDeadlockTimeoutOk() (*string, bool)`

GetDeadlockTimeoutOk returns a tuple with the DeadlockTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadlockTimeout

`func (o *ConfigParameters) SetDeadlockTimeout(v string)`

SetDeadlockTimeout sets DeadlockTimeout field to given value.

### HasDeadlockTimeout

`func (o *ConfigParameters) HasDeadlockTimeout() bool`

HasDeadlockTimeout returns a boolean if a field has been set.

### GetGinPendingListLimit

`func (o *ConfigParameters) GetGinPendingListLimit() string`

GetGinPendingListLimit returns the GinPendingListLimit field if non-nil, zero value otherwise.

### GetGinPendingListLimitOk

`func (o *ConfigParameters) GetGinPendingListLimitOk() (*string, bool)`

GetGinPendingListLimitOk returns a tuple with the GinPendingListLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGinPendingListLimit

`func (o *ConfigParameters) SetGinPendingListLimit(v string)`

SetGinPendingListLimit sets GinPendingListLimit field to given value.

### HasGinPendingListLimit

`func (o *ConfigParameters) HasGinPendingListLimit() bool`

HasGinPendingListLimit returns a boolean if a field has been set.

### GetIdleInTransactionSessionTimeout

`func (o *ConfigParameters) GetIdleInTransactionSessionTimeout() string`

GetIdleInTransactionSessionTimeout returns the IdleInTransactionSessionTimeout field if non-nil, zero value otherwise.

### GetIdleInTransactionSessionTimeoutOk

`func (o *ConfigParameters) GetIdleInTransactionSessionTimeoutOk() (*string, bool)`

GetIdleInTransactionSessionTimeoutOk returns a tuple with the IdleInTransactionSessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleInTransactionSessionTimeout

`func (o *ConfigParameters) SetIdleInTransactionSessionTimeout(v string)`

SetIdleInTransactionSessionTimeout sets IdleInTransactionSessionTimeout field to given value.

### HasIdleInTransactionSessionTimeout

`func (o *ConfigParameters) HasIdleInTransactionSessionTimeout() bool`

HasIdleInTransactionSessionTimeout returns a boolean if a field has been set.

### GetIdleSessionTimeout

`func (o *ConfigParameters) GetIdleSessionTimeout() string`

GetIdleSessionTimeout returns the IdleSessionTimeout field if non-nil, zero value otherwise.

### GetIdleSessionTimeoutOk

`func (o *ConfigParameters) GetIdleSessionTimeoutOk() (*string, bool)`

GetIdleSessionTimeoutOk returns a tuple with the IdleSessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleSessionTimeout

`func (o *ConfigParameters) SetIdleSessionTimeout(v string)`

SetIdleSessionTimeout sets IdleSessionTimeout field to given value.

### HasIdleSessionTimeout

`func (o *ConfigParameters) HasIdleSessionTimeout() bool`

HasIdleSessionTimeout returns a boolean if a field has been set.

### GetJoinCollapseLimit

`func (o *ConfigParameters) GetJoinCollapseLimit() string`

GetJoinCollapseLimit returns the JoinCollapseLimit field if non-nil, zero value otherwise.

### GetJoinCollapseLimitOk

`func (o *ConfigParameters) GetJoinCollapseLimitOk() (*string, bool)`

GetJoinCollapseLimitOk returns a tuple with the JoinCollapseLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinCollapseLimit

`func (o *ConfigParameters) SetJoinCollapseLimit(v string)`

SetJoinCollapseLimit sets JoinCollapseLimit field to given value.

### HasJoinCollapseLimit

`func (o *ConfigParameters) HasJoinCollapseLimit() bool`

HasJoinCollapseLimit returns a boolean if a field has been set.

### GetLockTimeout

`func (o *ConfigParameters) GetLockTimeout() string`

GetLockTimeout returns the LockTimeout field if non-nil, zero value otherwise.

### GetLockTimeoutOk

`func (o *ConfigParameters) GetLockTimeoutOk() (*string, bool)`

GetLockTimeoutOk returns a tuple with the LockTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockTimeout

`func (o *ConfigParameters) SetLockTimeout(v string)`

SetLockTimeout sets LockTimeout field to given value.

### HasLockTimeout

`func (o *ConfigParameters) HasLockTimeout() bool`

HasLockTimeout returns a boolean if a field has been set.

### GetMaxPreparedTransactions

`func (o *ConfigParameters) GetMaxPreparedTransactions() string`

GetMaxPreparedTransactions returns the MaxPreparedTransactions field if non-nil, zero value otherwise.

### GetMaxPreparedTransactionsOk

`func (o *ConfigParameters) GetMaxPreparedTransactionsOk() (*string, bool)`

GetMaxPreparedTransactionsOk returns a tuple with the MaxPreparedTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPreparedTransactions

`func (o *ConfigParameters) SetMaxPreparedTransactions(v string)`

SetMaxPreparedTransactions sets MaxPreparedTransactions field to given value.

### HasMaxPreparedTransactions

`func (o *ConfigParameters) HasMaxPreparedTransactions() bool`

HasMaxPreparedTransactions returns a boolean if a field has been set.

### GetMaxConnections

`func (o *ConfigParameters) GetMaxConnections() string`

GetMaxConnections returns the MaxConnections field if non-nil, zero value otherwise.

### GetMaxConnectionsOk

`func (o *ConfigParameters) GetMaxConnectionsOk() (*string, bool)`

GetMaxConnectionsOk returns a tuple with the MaxConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnections

`func (o *ConfigParameters) SetMaxConnections(v string)`

SetMaxConnections sets MaxConnections field to given value.

### HasMaxConnections

`func (o *ConfigParameters) HasMaxConnections() bool`

HasMaxConnections returns a boolean if a field has been set.

### GetSharedBuffers

`func (o *ConfigParameters) GetSharedBuffers() string`

GetSharedBuffers returns the SharedBuffers field if non-nil, zero value otherwise.

### GetSharedBuffersOk

`func (o *ConfigParameters) GetSharedBuffersOk() (*string, bool)`

GetSharedBuffersOk returns a tuple with the SharedBuffers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedBuffers

`func (o *ConfigParameters) SetSharedBuffers(v string)`

SetSharedBuffers sets SharedBuffers field to given value.

### HasSharedBuffers

`func (o *ConfigParameters) HasSharedBuffers() bool`

HasSharedBuffers returns a boolean if a field has been set.

### GetWalBuffers

`func (o *ConfigParameters) GetWalBuffers() string`

GetWalBuffers returns the WalBuffers field if non-nil, zero value otherwise.

### GetWalBuffersOk

`func (o *ConfigParameters) GetWalBuffersOk() (*string, bool)`

GetWalBuffersOk returns a tuple with the WalBuffers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalBuffers

`func (o *ConfigParameters) SetWalBuffers(v string)`

SetWalBuffers sets WalBuffers field to given value.

### HasWalBuffers

`func (o *ConfigParameters) HasWalBuffers() bool`

HasWalBuffers returns a boolean if a field has been set.

### GetTempBuffers

`func (o *ConfigParameters) GetTempBuffers() string`

GetTempBuffers returns the TempBuffers field if non-nil, zero value otherwise.

### GetTempBuffersOk

`func (o *ConfigParameters) GetTempBuffersOk() (*string, bool)`

GetTempBuffersOk returns a tuple with the TempBuffers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempBuffers

`func (o *ConfigParameters) SetTempBuffers(v string)`

SetTempBuffers sets TempBuffers field to given value.

### HasTempBuffers

`func (o *ConfigParameters) HasTempBuffers() bool`

HasTempBuffers returns a boolean if a field has been set.

### GetWorkMem

`func (o *ConfigParameters) GetWorkMem() string`

GetWorkMem returns the WorkMem field if non-nil, zero value otherwise.

### GetWorkMemOk

`func (o *ConfigParameters) GetWorkMemOk() (*string, bool)`

GetWorkMemOk returns a tuple with the WorkMem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkMem

`func (o *ConfigParameters) SetWorkMem(v string)`

SetWorkMem sets WorkMem field to given value.

### HasWorkMem

`func (o *ConfigParameters) HasWorkMem() bool`

HasWorkMem returns a boolean if a field has been set.

### GetSqlMode

`func (o *ConfigParameters) GetSqlMode() string`

GetSqlMode returns the SqlMode field if non-nil, zero value otherwise.

### GetSqlModeOk

`func (o *ConfigParameters) GetSqlModeOk() (*string, bool)`

GetSqlModeOk returns a tuple with the SqlMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlMode

`func (o *ConfigParameters) SetSqlMode(v string)`

SetSqlMode sets SqlMode field to given value.

### HasSqlMode

`func (o *ConfigParameters) HasSqlMode() bool`

HasSqlMode returns a boolean if a field has been set.

### GetQueryCacheType

`func (o *ConfigParameters) GetQueryCacheType() string`

GetQueryCacheType returns the QueryCacheType field if non-nil, zero value otherwise.

### GetQueryCacheTypeOk

`func (o *ConfigParameters) GetQueryCacheTypeOk() (*string, bool)`

GetQueryCacheTypeOk returns a tuple with the QueryCacheType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryCacheType

`func (o *ConfigParameters) SetQueryCacheType(v string)`

SetQueryCacheType sets QueryCacheType field to given value.

### HasQueryCacheType

`func (o *ConfigParameters) HasQueryCacheType() bool`

HasQueryCacheType returns a boolean if a field has been set.

### GetQueryCacheSize

`func (o *ConfigParameters) GetQueryCacheSize() string`

GetQueryCacheSize returns the QueryCacheSize field if non-nil, zero value otherwise.

### GetQueryCacheSizeOk

`func (o *ConfigParameters) GetQueryCacheSizeOk() (*string, bool)`

GetQueryCacheSizeOk returns a tuple with the QueryCacheSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryCacheSize

`func (o *ConfigParameters) SetQueryCacheSize(v string)`

SetQueryCacheSize sets QueryCacheSize field to given value.

### HasQueryCacheSize

`func (o *ConfigParameters) HasQueryCacheSize() bool`

HasQueryCacheSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


