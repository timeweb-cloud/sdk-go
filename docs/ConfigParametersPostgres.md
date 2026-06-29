# ConfigParametersPostgres

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxConnections** | Pointer to **string** | Максимальное количество одновременных подключений к серверу (&#x60;mysql5&#x60; | &#x60;mysql&#x60; | &#x60;mysql8_4&#x60; | &#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**AutovacuumAnalyzeScaleFactor** | Pointer to **string** | Доля изменения строк таблицы перед запуском автоматического анализа (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**AutovacuumMaxWorkers** | Pointer to **string** | Максимальное количество процессов autovacuum, которые могут работать одновременно (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**AutovacuumNaptime** | Pointer to **string** | Интервал между запусками процессов autovacuum (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**AutovacuumVacuumInsertScaleFactor** | Pointer to **string** | Доля вставленных строк перед запуском vacuum для таблиц с большим количеством вставок (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**AutovacuumVacuumScaleFactor** | Pointer to **string** | Доля измененных или удаленных строк перед запуском autovacuum (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**AutovacuumWorkMem** | Pointer to **string** | Объем памяти, используемый одним процессом autovacuum (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**BgwriterDelay** | Pointer to **string** | Интервал между циклами фонового процесса записи страниц (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**BgwriterLruMaxpages** | Pointer to **string** | Максимальное количество страниц, записываемых background writer за один цикл (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**DeadlockTimeout** | Pointer to **string** | Время ожидания блокировки перед проверкой взаимной блокировки (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**GinPendingListLimit** | Pointer to **string** | Максимальный размер списка ожидающих вставок индекса GIN (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**IdleInTransactionSessionTimeout** | Pointer to **string** | Время ожидания неактивной транзакционной сессии перед завершением соединения (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**JoinCollapseLimit** | Pointer to **string** | Максимальное количество таблиц в JOIN, которые планировщик может переупорядочить (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**LockTimeout** | Pointer to **string** | Максимальное время ожидания блокировки перед отменой запроса (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**MaxPreparedTransactions** | Pointer to **string** | Максимальное количество подготовленных транзакций, которые могут существовать одновременно (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**SharedBuffers** | Pointer to **string** | Размер общей памяти, используемой PostgreSQL для буферного кэша (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**LogMinDurationStatement** | Pointer to **string** | Минимальное время выполнения запроса, после которого он записывается в журнал (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**WalBuffers** | Pointer to **string** | Размер памяти, используемой для буферизации WAL-записей (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**TempBuffers** | Pointer to **string** | Максимальный объем памяти для временных таблиц каждой сессии (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**WorkMem** | Pointer to **string** | Объем памяти, используемый одной операцией сортировки или хеширования (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**DefaultTransactionIsolation** | Pointer to **string** | Уровень изоляции транзакций по умолчанию (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**EffectiveCacheSize** | Pointer to **string** | Оценка объема дискового кэша, доступного планировщику запросов (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**MaxWalSize** | Pointer to **string** | Максимальный размер WAL перед запуском контрольной точки (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**MinWalSize** | Pointer to **string** | Минимальный размер WAL, который сохраняется между контрольными точками (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**WalLevel** | Pointer to **string** | Уровень детализации записи WAL для восстановления и репликации (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**MaxReplicationSlots** | Pointer to **string** | Максимальное количество слотов репликации, которые могут быть созданы (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**MaxWalSenders** | Pointer to **string** | Максимальное количество процессов отправки WAL для репликации (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**MaxWorkerProcesses** | Pointer to **string** | Максимальное количество фоновых процессов PostgreSQL (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**MaxLogicalReplicationWorkers** | Pointer to **string** | Максимальное количество процессов логической репликации (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**MaxParallelMaintenanceWorkers** | Pointer to **string** | Максимальное количество параллельных процессов для операций обслуживания (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**MaxParallelWorkers** | Pointer to **string** | Максимальное количество параллельных рабочих процессов для запросов (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**MaxParallelWorkersPerGather** | Pointer to **string** | Максимальное количество параллельных рабочих процессов на один Gather-узел (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**ArrayNulls** | Pointer to **string** | Разрешение использования NULL в массивах PostgreSQL (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**BackendFlushAfter** | Pointer to **string** | Количество страниц, после записи которых выполняется принудительная очистка данных на диск серверным процессом (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**BackslashQuote** | Pointer to **string** | Управление использованием обратного слеша в строковых литералах (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**BgwriterFlushAfter** | Pointer to **string** | Количество страниц, после которого background writer выполняет очистку данных на диск (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**BgwriterLruMultiplier** | Pointer to **string** | Множитель количества страниц, которые background writer пытается очистить (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**DefaultTransactionReadOnly** | Pointer to **string** | Определяет режим транзакций только для чтения по умолчанию (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**EnableHashagg** | Pointer to **string** | Разрешение использования Hash Aggregate планировщиком запросов (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**EnableHashjoin** | Pointer to **string** | Разрешение использования Hash Join планировщиком запросов (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**EnableIncrementalSort** | Pointer to **string** | Разрешение использования инкрементальной сортировки планировщиком (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**EnableIndexscan** | Pointer to **string** | Разрешение использования обычного индексного сканирования (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**EnableIndexonlyscan** | Pointer to **string** | Разрешение использования index-only scan (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**EnableMaterial** | Pointer to **string** | Разрешение использования материализации промежуточных результатов запросов (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**EnableMemoize** | Pointer to **string** | Разрешение использования Memoize узлов планировщиком запросов (&#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**EnableMergejoin** | Pointer to **string** | Разрешение использования Merge Join планировщиком запросов (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**EnableParallelAppend** | Pointer to **string** | Разрешение использования параллельного Append для запросов (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**EnableParallelHash** | Pointer to **string** | Разрешение использования параллельных Hash операций (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**EnablePartitionPruning** | Pointer to **string** | Разрешение удаления ненужных разделов таблицы при планировании запроса (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**EnablePartitionwiseJoin** | Pointer to **string** | Разрешение выполнения соединений между секционированными таблицами с учетом секций (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**EnablePartitionwiseAggregate** | Pointer to **string** | Разрешение выполнения агрегатных операций отдельно для секций таблиц (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**EnableSeqscan** | Pointer to **string** | Разрешение использования последовательного сканирования таблиц планировщиком запросов (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**EnableSort** | Pointer to **string** | Разрешение использования операций сортировки планировщиком запросов (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**EnableTidscan** | Pointer to **string** | Разрешение использования TID Scan для поиска строк по физическим идентификаторам (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**ExitOnError** | Pointer to **string** | Завершение сессии при возникновении ошибки SQL (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**FromCollapseLimit** | Pointer to **string** | Максимальное количество элементов FROM, которые планировщик может объединять при оптимизации запросов (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**Jit** | Pointer to **string** | Включение JIT-компиляции для ускорения выполнения запросов (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**PlanCacheMode** | Pointer to **string** | Режим использования кэша планов подготовленных запросов (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**QuoteAllIdentifiers** | Pointer to **string** | Всегда заключать идентификаторы в кавычки при генерации SQL (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**StandardConformingStrings** | Pointer to **string** | Использование стандартного поведения строковых литералов SQL (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**StatementTimeout** | Pointer to **string** | Максимальное время выполнения SQL-запроса перед автоматической отменой (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**Timezone** | Pointer to **string** | Часовой пояс сервера PostgreSQL по умолчанию (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**TransformNullEquals** | Pointer to **string** | Преобразование выражений вида &#x60;NULL &#x3D; NULL&#x60; в проверку IS NULL (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**MaxLocksPerTransaction** | Pointer to **string** | Количество объектов, которые может блокировать одна транзакция (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**AutovacuumVacuumCostLimit** | Pointer to **string** | Лимит стоимости операций autovacuum перед приостановкой работы (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**CheckpointTimeout** | Pointer to **string** | Максимальный интервал времени между автоматическими контрольными точками (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**CheckpointCompletionTarget** | Pointer to **string** | Доля интервала checkpoint, за которую PostgreSQL распределяет запись данных (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**WalCompression** | Pointer to **string** | Включение сжатия WAL-записей для уменьшения объема журнала (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**RandomPageCost** | Pointer to **string** | Оценочная стоимость случайного чтения страницы для планировщика запросов (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**EffectiveIoConcurrency** | Pointer to **string** | Количество параллельных операций ввода-вывода, которые планировщик может учитывать (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**LogLockWaits** | Pointer to **string** | Включение записи в журнал информации об ожидании блокировок дольше deadlock_timeout (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**LogTempFiles** | Pointer to **string** | Минимальный размер временных файлов, при котором они записываются в журнал (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**TrackIoTiming** | Pointer to **string** | Включение сбора статистики времени операций ввода-вывода (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**MaintenanceWorkMem** | Pointer to **string** | Максимальный объем памяти для операций обслуживания, таких как VACUUM и CREATE INDEX (&#x60;postgres&#x60; | &#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**IdleSessionTimeout** | Pointer to **string** | Время ожидания неактивной сессии перед автоматическим завершением соединения (&#x60;postgres14&#x60; | &#x60;postgres15&#x60; | &#x60;postgres16&#x60; | &#x60;postgres17&#x60; | &#x60;postgres18&#x60;). | [optional] 
**IoMethod** | Pointer to **string** | Метод выполнения операций ввода-вывода PostgreSQL (&#x60;postgres18&#x60;). | [optional] 
**IoWorkers** | Pointer to **string** | Количество фоновых процессов для выполнения операций ввода-вывода (&#x60;postgres18&#x60;). | [optional] 

## Methods

### NewConfigParametersPostgres

`func NewConfigParametersPostgres() *ConfigParametersPostgres`

NewConfigParametersPostgres instantiates a new ConfigParametersPostgres object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigParametersPostgresWithDefaults

`func NewConfigParametersPostgresWithDefaults() *ConfigParametersPostgres`

NewConfigParametersPostgresWithDefaults instantiates a new ConfigParametersPostgres object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxConnections

`func (o *ConfigParametersPostgres) GetMaxConnections() string`

GetMaxConnections returns the MaxConnections field if non-nil, zero value otherwise.

### GetMaxConnectionsOk

`func (o *ConfigParametersPostgres) GetMaxConnectionsOk() (*string, bool)`

GetMaxConnectionsOk returns a tuple with the MaxConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnections

`func (o *ConfigParametersPostgres) SetMaxConnections(v string)`

SetMaxConnections sets MaxConnections field to given value.

### HasMaxConnections

`func (o *ConfigParametersPostgres) HasMaxConnections() bool`

HasMaxConnections returns a boolean if a field has been set.

### GetAutovacuumAnalyzeScaleFactor

`func (o *ConfigParametersPostgres) GetAutovacuumAnalyzeScaleFactor() string`

GetAutovacuumAnalyzeScaleFactor returns the AutovacuumAnalyzeScaleFactor field if non-nil, zero value otherwise.

### GetAutovacuumAnalyzeScaleFactorOk

`func (o *ConfigParametersPostgres) GetAutovacuumAnalyzeScaleFactorOk() (*string, bool)`

GetAutovacuumAnalyzeScaleFactorOk returns a tuple with the AutovacuumAnalyzeScaleFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutovacuumAnalyzeScaleFactor

`func (o *ConfigParametersPostgres) SetAutovacuumAnalyzeScaleFactor(v string)`

SetAutovacuumAnalyzeScaleFactor sets AutovacuumAnalyzeScaleFactor field to given value.

### HasAutovacuumAnalyzeScaleFactor

`func (o *ConfigParametersPostgres) HasAutovacuumAnalyzeScaleFactor() bool`

HasAutovacuumAnalyzeScaleFactor returns a boolean if a field has been set.

### GetAutovacuumMaxWorkers

`func (o *ConfigParametersPostgres) GetAutovacuumMaxWorkers() string`

GetAutovacuumMaxWorkers returns the AutovacuumMaxWorkers field if non-nil, zero value otherwise.

### GetAutovacuumMaxWorkersOk

`func (o *ConfigParametersPostgres) GetAutovacuumMaxWorkersOk() (*string, bool)`

GetAutovacuumMaxWorkersOk returns a tuple with the AutovacuumMaxWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutovacuumMaxWorkers

`func (o *ConfigParametersPostgres) SetAutovacuumMaxWorkers(v string)`

SetAutovacuumMaxWorkers sets AutovacuumMaxWorkers field to given value.

### HasAutovacuumMaxWorkers

`func (o *ConfigParametersPostgres) HasAutovacuumMaxWorkers() bool`

HasAutovacuumMaxWorkers returns a boolean if a field has been set.

### GetAutovacuumNaptime

`func (o *ConfigParametersPostgres) GetAutovacuumNaptime() string`

GetAutovacuumNaptime returns the AutovacuumNaptime field if non-nil, zero value otherwise.

### GetAutovacuumNaptimeOk

`func (o *ConfigParametersPostgres) GetAutovacuumNaptimeOk() (*string, bool)`

GetAutovacuumNaptimeOk returns a tuple with the AutovacuumNaptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutovacuumNaptime

`func (o *ConfigParametersPostgres) SetAutovacuumNaptime(v string)`

SetAutovacuumNaptime sets AutovacuumNaptime field to given value.

### HasAutovacuumNaptime

`func (o *ConfigParametersPostgres) HasAutovacuumNaptime() bool`

HasAutovacuumNaptime returns a boolean if a field has been set.

### GetAutovacuumVacuumInsertScaleFactor

`func (o *ConfigParametersPostgres) GetAutovacuumVacuumInsertScaleFactor() string`

GetAutovacuumVacuumInsertScaleFactor returns the AutovacuumVacuumInsertScaleFactor field if non-nil, zero value otherwise.

### GetAutovacuumVacuumInsertScaleFactorOk

`func (o *ConfigParametersPostgres) GetAutovacuumVacuumInsertScaleFactorOk() (*string, bool)`

GetAutovacuumVacuumInsertScaleFactorOk returns a tuple with the AutovacuumVacuumInsertScaleFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutovacuumVacuumInsertScaleFactor

`func (o *ConfigParametersPostgres) SetAutovacuumVacuumInsertScaleFactor(v string)`

SetAutovacuumVacuumInsertScaleFactor sets AutovacuumVacuumInsertScaleFactor field to given value.

### HasAutovacuumVacuumInsertScaleFactor

`func (o *ConfigParametersPostgres) HasAutovacuumVacuumInsertScaleFactor() bool`

HasAutovacuumVacuumInsertScaleFactor returns a boolean if a field has been set.

### GetAutovacuumVacuumScaleFactor

`func (o *ConfigParametersPostgres) GetAutovacuumVacuumScaleFactor() string`

GetAutovacuumVacuumScaleFactor returns the AutovacuumVacuumScaleFactor field if non-nil, zero value otherwise.

### GetAutovacuumVacuumScaleFactorOk

`func (o *ConfigParametersPostgres) GetAutovacuumVacuumScaleFactorOk() (*string, bool)`

GetAutovacuumVacuumScaleFactorOk returns a tuple with the AutovacuumVacuumScaleFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutovacuumVacuumScaleFactor

`func (o *ConfigParametersPostgres) SetAutovacuumVacuumScaleFactor(v string)`

SetAutovacuumVacuumScaleFactor sets AutovacuumVacuumScaleFactor field to given value.

### HasAutovacuumVacuumScaleFactor

`func (o *ConfigParametersPostgres) HasAutovacuumVacuumScaleFactor() bool`

HasAutovacuumVacuumScaleFactor returns a boolean if a field has been set.

### GetAutovacuumWorkMem

`func (o *ConfigParametersPostgres) GetAutovacuumWorkMem() string`

GetAutovacuumWorkMem returns the AutovacuumWorkMem field if non-nil, zero value otherwise.

### GetAutovacuumWorkMemOk

`func (o *ConfigParametersPostgres) GetAutovacuumWorkMemOk() (*string, bool)`

GetAutovacuumWorkMemOk returns a tuple with the AutovacuumWorkMem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutovacuumWorkMem

`func (o *ConfigParametersPostgres) SetAutovacuumWorkMem(v string)`

SetAutovacuumWorkMem sets AutovacuumWorkMem field to given value.

### HasAutovacuumWorkMem

`func (o *ConfigParametersPostgres) HasAutovacuumWorkMem() bool`

HasAutovacuumWorkMem returns a boolean if a field has been set.

### GetBgwriterDelay

`func (o *ConfigParametersPostgres) GetBgwriterDelay() string`

GetBgwriterDelay returns the BgwriterDelay field if non-nil, zero value otherwise.

### GetBgwriterDelayOk

`func (o *ConfigParametersPostgres) GetBgwriterDelayOk() (*string, bool)`

GetBgwriterDelayOk returns a tuple with the BgwriterDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgwriterDelay

`func (o *ConfigParametersPostgres) SetBgwriterDelay(v string)`

SetBgwriterDelay sets BgwriterDelay field to given value.

### HasBgwriterDelay

`func (o *ConfigParametersPostgres) HasBgwriterDelay() bool`

HasBgwriterDelay returns a boolean if a field has been set.

### GetBgwriterLruMaxpages

`func (o *ConfigParametersPostgres) GetBgwriterLruMaxpages() string`

GetBgwriterLruMaxpages returns the BgwriterLruMaxpages field if non-nil, zero value otherwise.

### GetBgwriterLruMaxpagesOk

`func (o *ConfigParametersPostgres) GetBgwriterLruMaxpagesOk() (*string, bool)`

GetBgwriterLruMaxpagesOk returns a tuple with the BgwriterLruMaxpages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgwriterLruMaxpages

`func (o *ConfigParametersPostgres) SetBgwriterLruMaxpages(v string)`

SetBgwriterLruMaxpages sets BgwriterLruMaxpages field to given value.

### HasBgwriterLruMaxpages

`func (o *ConfigParametersPostgres) HasBgwriterLruMaxpages() bool`

HasBgwriterLruMaxpages returns a boolean if a field has been set.

### GetDeadlockTimeout

`func (o *ConfigParametersPostgres) GetDeadlockTimeout() string`

GetDeadlockTimeout returns the DeadlockTimeout field if non-nil, zero value otherwise.

### GetDeadlockTimeoutOk

`func (o *ConfigParametersPostgres) GetDeadlockTimeoutOk() (*string, bool)`

GetDeadlockTimeoutOk returns a tuple with the DeadlockTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadlockTimeout

`func (o *ConfigParametersPostgres) SetDeadlockTimeout(v string)`

SetDeadlockTimeout sets DeadlockTimeout field to given value.

### HasDeadlockTimeout

`func (o *ConfigParametersPostgres) HasDeadlockTimeout() bool`

HasDeadlockTimeout returns a boolean if a field has been set.

### GetGinPendingListLimit

`func (o *ConfigParametersPostgres) GetGinPendingListLimit() string`

GetGinPendingListLimit returns the GinPendingListLimit field if non-nil, zero value otherwise.

### GetGinPendingListLimitOk

`func (o *ConfigParametersPostgres) GetGinPendingListLimitOk() (*string, bool)`

GetGinPendingListLimitOk returns a tuple with the GinPendingListLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGinPendingListLimit

`func (o *ConfigParametersPostgres) SetGinPendingListLimit(v string)`

SetGinPendingListLimit sets GinPendingListLimit field to given value.

### HasGinPendingListLimit

`func (o *ConfigParametersPostgres) HasGinPendingListLimit() bool`

HasGinPendingListLimit returns a boolean if a field has been set.

### GetIdleInTransactionSessionTimeout

`func (o *ConfigParametersPostgres) GetIdleInTransactionSessionTimeout() string`

GetIdleInTransactionSessionTimeout returns the IdleInTransactionSessionTimeout field if non-nil, zero value otherwise.

### GetIdleInTransactionSessionTimeoutOk

`func (o *ConfigParametersPostgres) GetIdleInTransactionSessionTimeoutOk() (*string, bool)`

GetIdleInTransactionSessionTimeoutOk returns a tuple with the IdleInTransactionSessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleInTransactionSessionTimeout

`func (o *ConfigParametersPostgres) SetIdleInTransactionSessionTimeout(v string)`

SetIdleInTransactionSessionTimeout sets IdleInTransactionSessionTimeout field to given value.

### HasIdleInTransactionSessionTimeout

`func (o *ConfigParametersPostgres) HasIdleInTransactionSessionTimeout() bool`

HasIdleInTransactionSessionTimeout returns a boolean if a field has been set.

### GetJoinCollapseLimit

`func (o *ConfigParametersPostgres) GetJoinCollapseLimit() string`

GetJoinCollapseLimit returns the JoinCollapseLimit field if non-nil, zero value otherwise.

### GetJoinCollapseLimitOk

`func (o *ConfigParametersPostgres) GetJoinCollapseLimitOk() (*string, bool)`

GetJoinCollapseLimitOk returns a tuple with the JoinCollapseLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinCollapseLimit

`func (o *ConfigParametersPostgres) SetJoinCollapseLimit(v string)`

SetJoinCollapseLimit sets JoinCollapseLimit field to given value.

### HasJoinCollapseLimit

`func (o *ConfigParametersPostgres) HasJoinCollapseLimit() bool`

HasJoinCollapseLimit returns a boolean if a field has been set.

### GetLockTimeout

`func (o *ConfigParametersPostgres) GetLockTimeout() string`

GetLockTimeout returns the LockTimeout field if non-nil, zero value otherwise.

### GetLockTimeoutOk

`func (o *ConfigParametersPostgres) GetLockTimeoutOk() (*string, bool)`

GetLockTimeoutOk returns a tuple with the LockTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockTimeout

`func (o *ConfigParametersPostgres) SetLockTimeout(v string)`

SetLockTimeout sets LockTimeout field to given value.

### HasLockTimeout

`func (o *ConfigParametersPostgres) HasLockTimeout() bool`

HasLockTimeout returns a boolean if a field has been set.

### GetMaxPreparedTransactions

`func (o *ConfigParametersPostgres) GetMaxPreparedTransactions() string`

GetMaxPreparedTransactions returns the MaxPreparedTransactions field if non-nil, zero value otherwise.

### GetMaxPreparedTransactionsOk

`func (o *ConfigParametersPostgres) GetMaxPreparedTransactionsOk() (*string, bool)`

GetMaxPreparedTransactionsOk returns a tuple with the MaxPreparedTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPreparedTransactions

`func (o *ConfigParametersPostgres) SetMaxPreparedTransactions(v string)`

SetMaxPreparedTransactions sets MaxPreparedTransactions field to given value.

### HasMaxPreparedTransactions

`func (o *ConfigParametersPostgres) HasMaxPreparedTransactions() bool`

HasMaxPreparedTransactions returns a boolean if a field has been set.

### GetSharedBuffers

`func (o *ConfigParametersPostgres) GetSharedBuffers() string`

GetSharedBuffers returns the SharedBuffers field if non-nil, zero value otherwise.

### GetSharedBuffersOk

`func (o *ConfigParametersPostgres) GetSharedBuffersOk() (*string, bool)`

GetSharedBuffersOk returns a tuple with the SharedBuffers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedBuffers

`func (o *ConfigParametersPostgres) SetSharedBuffers(v string)`

SetSharedBuffers sets SharedBuffers field to given value.

### HasSharedBuffers

`func (o *ConfigParametersPostgres) HasSharedBuffers() bool`

HasSharedBuffers returns a boolean if a field has been set.

### GetLogMinDurationStatement

`func (o *ConfigParametersPostgres) GetLogMinDurationStatement() string`

GetLogMinDurationStatement returns the LogMinDurationStatement field if non-nil, zero value otherwise.

### GetLogMinDurationStatementOk

`func (o *ConfigParametersPostgres) GetLogMinDurationStatementOk() (*string, bool)`

GetLogMinDurationStatementOk returns a tuple with the LogMinDurationStatement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogMinDurationStatement

`func (o *ConfigParametersPostgres) SetLogMinDurationStatement(v string)`

SetLogMinDurationStatement sets LogMinDurationStatement field to given value.

### HasLogMinDurationStatement

`func (o *ConfigParametersPostgres) HasLogMinDurationStatement() bool`

HasLogMinDurationStatement returns a boolean if a field has been set.

### GetWalBuffers

`func (o *ConfigParametersPostgres) GetWalBuffers() string`

GetWalBuffers returns the WalBuffers field if non-nil, zero value otherwise.

### GetWalBuffersOk

`func (o *ConfigParametersPostgres) GetWalBuffersOk() (*string, bool)`

GetWalBuffersOk returns a tuple with the WalBuffers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalBuffers

`func (o *ConfigParametersPostgres) SetWalBuffers(v string)`

SetWalBuffers sets WalBuffers field to given value.

### HasWalBuffers

`func (o *ConfigParametersPostgres) HasWalBuffers() bool`

HasWalBuffers returns a boolean if a field has been set.

### GetTempBuffers

`func (o *ConfigParametersPostgres) GetTempBuffers() string`

GetTempBuffers returns the TempBuffers field if non-nil, zero value otherwise.

### GetTempBuffersOk

`func (o *ConfigParametersPostgres) GetTempBuffersOk() (*string, bool)`

GetTempBuffersOk returns a tuple with the TempBuffers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempBuffers

`func (o *ConfigParametersPostgres) SetTempBuffers(v string)`

SetTempBuffers sets TempBuffers field to given value.

### HasTempBuffers

`func (o *ConfigParametersPostgres) HasTempBuffers() bool`

HasTempBuffers returns a boolean if a field has been set.

### GetWorkMem

`func (o *ConfigParametersPostgres) GetWorkMem() string`

GetWorkMem returns the WorkMem field if non-nil, zero value otherwise.

### GetWorkMemOk

`func (o *ConfigParametersPostgres) GetWorkMemOk() (*string, bool)`

GetWorkMemOk returns a tuple with the WorkMem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkMem

`func (o *ConfigParametersPostgres) SetWorkMem(v string)`

SetWorkMem sets WorkMem field to given value.

### HasWorkMem

`func (o *ConfigParametersPostgres) HasWorkMem() bool`

HasWorkMem returns a boolean if a field has been set.

### GetDefaultTransactionIsolation

`func (o *ConfigParametersPostgres) GetDefaultTransactionIsolation() string`

GetDefaultTransactionIsolation returns the DefaultTransactionIsolation field if non-nil, zero value otherwise.

### GetDefaultTransactionIsolationOk

`func (o *ConfigParametersPostgres) GetDefaultTransactionIsolationOk() (*string, bool)`

GetDefaultTransactionIsolationOk returns a tuple with the DefaultTransactionIsolation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTransactionIsolation

`func (o *ConfigParametersPostgres) SetDefaultTransactionIsolation(v string)`

SetDefaultTransactionIsolation sets DefaultTransactionIsolation field to given value.

### HasDefaultTransactionIsolation

`func (o *ConfigParametersPostgres) HasDefaultTransactionIsolation() bool`

HasDefaultTransactionIsolation returns a boolean if a field has been set.

### GetEffectiveCacheSize

`func (o *ConfigParametersPostgres) GetEffectiveCacheSize() string`

GetEffectiveCacheSize returns the EffectiveCacheSize field if non-nil, zero value otherwise.

### GetEffectiveCacheSizeOk

`func (o *ConfigParametersPostgres) GetEffectiveCacheSizeOk() (*string, bool)`

GetEffectiveCacheSizeOk returns a tuple with the EffectiveCacheSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveCacheSize

`func (o *ConfigParametersPostgres) SetEffectiveCacheSize(v string)`

SetEffectiveCacheSize sets EffectiveCacheSize field to given value.

### HasEffectiveCacheSize

`func (o *ConfigParametersPostgres) HasEffectiveCacheSize() bool`

HasEffectiveCacheSize returns a boolean if a field has been set.

### GetMaxWalSize

`func (o *ConfigParametersPostgres) GetMaxWalSize() string`

GetMaxWalSize returns the MaxWalSize field if non-nil, zero value otherwise.

### GetMaxWalSizeOk

`func (o *ConfigParametersPostgres) GetMaxWalSizeOk() (*string, bool)`

GetMaxWalSizeOk returns a tuple with the MaxWalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWalSize

`func (o *ConfigParametersPostgres) SetMaxWalSize(v string)`

SetMaxWalSize sets MaxWalSize field to given value.

### HasMaxWalSize

`func (o *ConfigParametersPostgres) HasMaxWalSize() bool`

HasMaxWalSize returns a boolean if a field has been set.

### GetMinWalSize

`func (o *ConfigParametersPostgres) GetMinWalSize() string`

GetMinWalSize returns the MinWalSize field if non-nil, zero value otherwise.

### GetMinWalSizeOk

`func (o *ConfigParametersPostgres) GetMinWalSizeOk() (*string, bool)`

GetMinWalSizeOk returns a tuple with the MinWalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinWalSize

`func (o *ConfigParametersPostgres) SetMinWalSize(v string)`

SetMinWalSize sets MinWalSize field to given value.

### HasMinWalSize

`func (o *ConfigParametersPostgres) HasMinWalSize() bool`

HasMinWalSize returns a boolean if a field has been set.

### GetWalLevel

`func (o *ConfigParametersPostgres) GetWalLevel() string`

GetWalLevel returns the WalLevel field if non-nil, zero value otherwise.

### GetWalLevelOk

`func (o *ConfigParametersPostgres) GetWalLevelOk() (*string, bool)`

GetWalLevelOk returns a tuple with the WalLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalLevel

`func (o *ConfigParametersPostgres) SetWalLevel(v string)`

SetWalLevel sets WalLevel field to given value.

### HasWalLevel

`func (o *ConfigParametersPostgres) HasWalLevel() bool`

HasWalLevel returns a boolean if a field has been set.

### GetMaxReplicationSlots

`func (o *ConfigParametersPostgres) GetMaxReplicationSlots() string`

GetMaxReplicationSlots returns the MaxReplicationSlots field if non-nil, zero value otherwise.

### GetMaxReplicationSlotsOk

`func (o *ConfigParametersPostgres) GetMaxReplicationSlotsOk() (*string, bool)`

GetMaxReplicationSlotsOk returns a tuple with the MaxReplicationSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReplicationSlots

`func (o *ConfigParametersPostgres) SetMaxReplicationSlots(v string)`

SetMaxReplicationSlots sets MaxReplicationSlots field to given value.

### HasMaxReplicationSlots

`func (o *ConfigParametersPostgres) HasMaxReplicationSlots() bool`

HasMaxReplicationSlots returns a boolean if a field has been set.

### GetMaxWalSenders

`func (o *ConfigParametersPostgres) GetMaxWalSenders() string`

GetMaxWalSenders returns the MaxWalSenders field if non-nil, zero value otherwise.

### GetMaxWalSendersOk

`func (o *ConfigParametersPostgres) GetMaxWalSendersOk() (*string, bool)`

GetMaxWalSendersOk returns a tuple with the MaxWalSenders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWalSenders

`func (o *ConfigParametersPostgres) SetMaxWalSenders(v string)`

SetMaxWalSenders sets MaxWalSenders field to given value.

### HasMaxWalSenders

`func (o *ConfigParametersPostgres) HasMaxWalSenders() bool`

HasMaxWalSenders returns a boolean if a field has been set.

### GetMaxWorkerProcesses

`func (o *ConfigParametersPostgres) GetMaxWorkerProcesses() string`

GetMaxWorkerProcesses returns the MaxWorkerProcesses field if non-nil, zero value otherwise.

### GetMaxWorkerProcessesOk

`func (o *ConfigParametersPostgres) GetMaxWorkerProcessesOk() (*string, bool)`

GetMaxWorkerProcessesOk returns a tuple with the MaxWorkerProcesses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWorkerProcesses

`func (o *ConfigParametersPostgres) SetMaxWorkerProcesses(v string)`

SetMaxWorkerProcesses sets MaxWorkerProcesses field to given value.

### HasMaxWorkerProcesses

`func (o *ConfigParametersPostgres) HasMaxWorkerProcesses() bool`

HasMaxWorkerProcesses returns a boolean if a field has been set.

### GetMaxLogicalReplicationWorkers

`func (o *ConfigParametersPostgres) GetMaxLogicalReplicationWorkers() string`

GetMaxLogicalReplicationWorkers returns the MaxLogicalReplicationWorkers field if non-nil, zero value otherwise.

### GetMaxLogicalReplicationWorkersOk

`func (o *ConfigParametersPostgres) GetMaxLogicalReplicationWorkersOk() (*string, bool)`

GetMaxLogicalReplicationWorkersOk returns a tuple with the MaxLogicalReplicationWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLogicalReplicationWorkers

`func (o *ConfigParametersPostgres) SetMaxLogicalReplicationWorkers(v string)`

SetMaxLogicalReplicationWorkers sets MaxLogicalReplicationWorkers field to given value.

### HasMaxLogicalReplicationWorkers

`func (o *ConfigParametersPostgres) HasMaxLogicalReplicationWorkers() bool`

HasMaxLogicalReplicationWorkers returns a boolean if a field has been set.

### GetMaxParallelMaintenanceWorkers

`func (o *ConfigParametersPostgres) GetMaxParallelMaintenanceWorkers() string`

GetMaxParallelMaintenanceWorkers returns the MaxParallelMaintenanceWorkers field if non-nil, zero value otherwise.

### GetMaxParallelMaintenanceWorkersOk

`func (o *ConfigParametersPostgres) GetMaxParallelMaintenanceWorkersOk() (*string, bool)`

GetMaxParallelMaintenanceWorkersOk returns a tuple with the MaxParallelMaintenanceWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxParallelMaintenanceWorkers

`func (o *ConfigParametersPostgres) SetMaxParallelMaintenanceWorkers(v string)`

SetMaxParallelMaintenanceWorkers sets MaxParallelMaintenanceWorkers field to given value.

### HasMaxParallelMaintenanceWorkers

`func (o *ConfigParametersPostgres) HasMaxParallelMaintenanceWorkers() bool`

HasMaxParallelMaintenanceWorkers returns a boolean if a field has been set.

### GetMaxParallelWorkers

`func (o *ConfigParametersPostgres) GetMaxParallelWorkers() string`

GetMaxParallelWorkers returns the MaxParallelWorkers field if non-nil, zero value otherwise.

### GetMaxParallelWorkersOk

`func (o *ConfigParametersPostgres) GetMaxParallelWorkersOk() (*string, bool)`

GetMaxParallelWorkersOk returns a tuple with the MaxParallelWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxParallelWorkers

`func (o *ConfigParametersPostgres) SetMaxParallelWorkers(v string)`

SetMaxParallelWorkers sets MaxParallelWorkers field to given value.

### HasMaxParallelWorkers

`func (o *ConfigParametersPostgres) HasMaxParallelWorkers() bool`

HasMaxParallelWorkers returns a boolean if a field has been set.

### GetMaxParallelWorkersPerGather

`func (o *ConfigParametersPostgres) GetMaxParallelWorkersPerGather() string`

GetMaxParallelWorkersPerGather returns the MaxParallelWorkersPerGather field if non-nil, zero value otherwise.

### GetMaxParallelWorkersPerGatherOk

`func (o *ConfigParametersPostgres) GetMaxParallelWorkersPerGatherOk() (*string, bool)`

GetMaxParallelWorkersPerGatherOk returns a tuple with the MaxParallelWorkersPerGather field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxParallelWorkersPerGather

`func (o *ConfigParametersPostgres) SetMaxParallelWorkersPerGather(v string)`

SetMaxParallelWorkersPerGather sets MaxParallelWorkersPerGather field to given value.

### HasMaxParallelWorkersPerGather

`func (o *ConfigParametersPostgres) HasMaxParallelWorkersPerGather() bool`

HasMaxParallelWorkersPerGather returns a boolean if a field has been set.

### GetArrayNulls

`func (o *ConfigParametersPostgres) GetArrayNulls() string`

GetArrayNulls returns the ArrayNulls field if non-nil, zero value otherwise.

### GetArrayNullsOk

`func (o *ConfigParametersPostgres) GetArrayNullsOk() (*string, bool)`

GetArrayNullsOk returns a tuple with the ArrayNulls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayNulls

`func (o *ConfigParametersPostgres) SetArrayNulls(v string)`

SetArrayNulls sets ArrayNulls field to given value.

### HasArrayNulls

`func (o *ConfigParametersPostgres) HasArrayNulls() bool`

HasArrayNulls returns a boolean if a field has been set.

### GetBackendFlushAfter

`func (o *ConfigParametersPostgres) GetBackendFlushAfter() string`

GetBackendFlushAfter returns the BackendFlushAfter field if non-nil, zero value otherwise.

### GetBackendFlushAfterOk

`func (o *ConfigParametersPostgres) GetBackendFlushAfterOk() (*string, bool)`

GetBackendFlushAfterOk returns a tuple with the BackendFlushAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendFlushAfter

`func (o *ConfigParametersPostgres) SetBackendFlushAfter(v string)`

SetBackendFlushAfter sets BackendFlushAfter field to given value.

### HasBackendFlushAfter

`func (o *ConfigParametersPostgres) HasBackendFlushAfter() bool`

HasBackendFlushAfter returns a boolean if a field has been set.

### GetBackslashQuote

`func (o *ConfigParametersPostgres) GetBackslashQuote() string`

GetBackslashQuote returns the BackslashQuote field if non-nil, zero value otherwise.

### GetBackslashQuoteOk

`func (o *ConfigParametersPostgres) GetBackslashQuoteOk() (*string, bool)`

GetBackslashQuoteOk returns a tuple with the BackslashQuote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackslashQuote

`func (o *ConfigParametersPostgres) SetBackslashQuote(v string)`

SetBackslashQuote sets BackslashQuote field to given value.

### HasBackslashQuote

`func (o *ConfigParametersPostgres) HasBackslashQuote() bool`

HasBackslashQuote returns a boolean if a field has been set.

### GetBgwriterFlushAfter

`func (o *ConfigParametersPostgres) GetBgwriterFlushAfter() string`

GetBgwriterFlushAfter returns the BgwriterFlushAfter field if non-nil, zero value otherwise.

### GetBgwriterFlushAfterOk

`func (o *ConfigParametersPostgres) GetBgwriterFlushAfterOk() (*string, bool)`

GetBgwriterFlushAfterOk returns a tuple with the BgwriterFlushAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgwriterFlushAfter

`func (o *ConfigParametersPostgres) SetBgwriterFlushAfter(v string)`

SetBgwriterFlushAfter sets BgwriterFlushAfter field to given value.

### HasBgwriterFlushAfter

`func (o *ConfigParametersPostgres) HasBgwriterFlushAfter() bool`

HasBgwriterFlushAfter returns a boolean if a field has been set.

### GetBgwriterLruMultiplier

`func (o *ConfigParametersPostgres) GetBgwriterLruMultiplier() string`

GetBgwriterLruMultiplier returns the BgwriterLruMultiplier field if non-nil, zero value otherwise.

### GetBgwriterLruMultiplierOk

`func (o *ConfigParametersPostgres) GetBgwriterLruMultiplierOk() (*string, bool)`

GetBgwriterLruMultiplierOk returns a tuple with the BgwriterLruMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgwriterLruMultiplier

`func (o *ConfigParametersPostgres) SetBgwriterLruMultiplier(v string)`

SetBgwriterLruMultiplier sets BgwriterLruMultiplier field to given value.

### HasBgwriterLruMultiplier

`func (o *ConfigParametersPostgres) HasBgwriterLruMultiplier() bool`

HasBgwriterLruMultiplier returns a boolean if a field has been set.

### GetDefaultTransactionReadOnly

`func (o *ConfigParametersPostgres) GetDefaultTransactionReadOnly() string`

GetDefaultTransactionReadOnly returns the DefaultTransactionReadOnly field if non-nil, zero value otherwise.

### GetDefaultTransactionReadOnlyOk

`func (o *ConfigParametersPostgres) GetDefaultTransactionReadOnlyOk() (*string, bool)`

GetDefaultTransactionReadOnlyOk returns a tuple with the DefaultTransactionReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTransactionReadOnly

`func (o *ConfigParametersPostgres) SetDefaultTransactionReadOnly(v string)`

SetDefaultTransactionReadOnly sets DefaultTransactionReadOnly field to given value.

### HasDefaultTransactionReadOnly

`func (o *ConfigParametersPostgres) HasDefaultTransactionReadOnly() bool`

HasDefaultTransactionReadOnly returns a boolean if a field has been set.

### GetEnableHashagg

`func (o *ConfigParametersPostgres) GetEnableHashagg() string`

GetEnableHashagg returns the EnableHashagg field if non-nil, zero value otherwise.

### GetEnableHashaggOk

`func (o *ConfigParametersPostgres) GetEnableHashaggOk() (*string, bool)`

GetEnableHashaggOk returns a tuple with the EnableHashagg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableHashagg

`func (o *ConfigParametersPostgres) SetEnableHashagg(v string)`

SetEnableHashagg sets EnableHashagg field to given value.

### HasEnableHashagg

`func (o *ConfigParametersPostgres) HasEnableHashagg() bool`

HasEnableHashagg returns a boolean if a field has been set.

### GetEnableHashjoin

`func (o *ConfigParametersPostgres) GetEnableHashjoin() string`

GetEnableHashjoin returns the EnableHashjoin field if non-nil, zero value otherwise.

### GetEnableHashjoinOk

`func (o *ConfigParametersPostgres) GetEnableHashjoinOk() (*string, bool)`

GetEnableHashjoinOk returns a tuple with the EnableHashjoin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableHashjoin

`func (o *ConfigParametersPostgres) SetEnableHashjoin(v string)`

SetEnableHashjoin sets EnableHashjoin field to given value.

### HasEnableHashjoin

`func (o *ConfigParametersPostgres) HasEnableHashjoin() bool`

HasEnableHashjoin returns a boolean if a field has been set.

### GetEnableIncrementalSort

`func (o *ConfigParametersPostgres) GetEnableIncrementalSort() string`

GetEnableIncrementalSort returns the EnableIncrementalSort field if non-nil, zero value otherwise.

### GetEnableIncrementalSortOk

`func (o *ConfigParametersPostgres) GetEnableIncrementalSortOk() (*string, bool)`

GetEnableIncrementalSortOk returns a tuple with the EnableIncrementalSort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIncrementalSort

`func (o *ConfigParametersPostgres) SetEnableIncrementalSort(v string)`

SetEnableIncrementalSort sets EnableIncrementalSort field to given value.

### HasEnableIncrementalSort

`func (o *ConfigParametersPostgres) HasEnableIncrementalSort() bool`

HasEnableIncrementalSort returns a boolean if a field has been set.

### GetEnableIndexscan

`func (o *ConfigParametersPostgres) GetEnableIndexscan() string`

GetEnableIndexscan returns the EnableIndexscan field if non-nil, zero value otherwise.

### GetEnableIndexscanOk

`func (o *ConfigParametersPostgres) GetEnableIndexscanOk() (*string, bool)`

GetEnableIndexscanOk returns a tuple with the EnableIndexscan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIndexscan

`func (o *ConfigParametersPostgres) SetEnableIndexscan(v string)`

SetEnableIndexscan sets EnableIndexscan field to given value.

### HasEnableIndexscan

`func (o *ConfigParametersPostgres) HasEnableIndexscan() bool`

HasEnableIndexscan returns a boolean if a field has been set.

### GetEnableIndexonlyscan

`func (o *ConfigParametersPostgres) GetEnableIndexonlyscan() string`

GetEnableIndexonlyscan returns the EnableIndexonlyscan field if non-nil, zero value otherwise.

### GetEnableIndexonlyscanOk

`func (o *ConfigParametersPostgres) GetEnableIndexonlyscanOk() (*string, bool)`

GetEnableIndexonlyscanOk returns a tuple with the EnableIndexonlyscan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIndexonlyscan

`func (o *ConfigParametersPostgres) SetEnableIndexonlyscan(v string)`

SetEnableIndexonlyscan sets EnableIndexonlyscan field to given value.

### HasEnableIndexonlyscan

`func (o *ConfigParametersPostgres) HasEnableIndexonlyscan() bool`

HasEnableIndexonlyscan returns a boolean if a field has been set.

### GetEnableMaterial

`func (o *ConfigParametersPostgres) GetEnableMaterial() string`

GetEnableMaterial returns the EnableMaterial field if non-nil, zero value otherwise.

### GetEnableMaterialOk

`func (o *ConfigParametersPostgres) GetEnableMaterialOk() (*string, bool)`

GetEnableMaterialOk returns a tuple with the EnableMaterial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMaterial

`func (o *ConfigParametersPostgres) SetEnableMaterial(v string)`

SetEnableMaterial sets EnableMaterial field to given value.

### HasEnableMaterial

`func (o *ConfigParametersPostgres) HasEnableMaterial() bool`

HasEnableMaterial returns a boolean if a field has been set.

### GetEnableMemoize

`func (o *ConfigParametersPostgres) GetEnableMemoize() string`

GetEnableMemoize returns the EnableMemoize field if non-nil, zero value otherwise.

### GetEnableMemoizeOk

`func (o *ConfigParametersPostgres) GetEnableMemoizeOk() (*string, bool)`

GetEnableMemoizeOk returns a tuple with the EnableMemoize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMemoize

`func (o *ConfigParametersPostgres) SetEnableMemoize(v string)`

SetEnableMemoize sets EnableMemoize field to given value.

### HasEnableMemoize

`func (o *ConfigParametersPostgres) HasEnableMemoize() bool`

HasEnableMemoize returns a boolean if a field has been set.

### GetEnableMergejoin

`func (o *ConfigParametersPostgres) GetEnableMergejoin() string`

GetEnableMergejoin returns the EnableMergejoin field if non-nil, zero value otherwise.

### GetEnableMergejoinOk

`func (o *ConfigParametersPostgres) GetEnableMergejoinOk() (*string, bool)`

GetEnableMergejoinOk returns a tuple with the EnableMergejoin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMergejoin

`func (o *ConfigParametersPostgres) SetEnableMergejoin(v string)`

SetEnableMergejoin sets EnableMergejoin field to given value.

### HasEnableMergejoin

`func (o *ConfigParametersPostgres) HasEnableMergejoin() bool`

HasEnableMergejoin returns a boolean if a field has been set.

### GetEnableParallelAppend

`func (o *ConfigParametersPostgres) GetEnableParallelAppend() string`

GetEnableParallelAppend returns the EnableParallelAppend field if non-nil, zero value otherwise.

### GetEnableParallelAppendOk

`func (o *ConfigParametersPostgres) GetEnableParallelAppendOk() (*string, bool)`

GetEnableParallelAppendOk returns a tuple with the EnableParallelAppend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableParallelAppend

`func (o *ConfigParametersPostgres) SetEnableParallelAppend(v string)`

SetEnableParallelAppend sets EnableParallelAppend field to given value.

### HasEnableParallelAppend

`func (o *ConfigParametersPostgres) HasEnableParallelAppend() bool`

HasEnableParallelAppend returns a boolean if a field has been set.

### GetEnableParallelHash

`func (o *ConfigParametersPostgres) GetEnableParallelHash() string`

GetEnableParallelHash returns the EnableParallelHash field if non-nil, zero value otherwise.

### GetEnableParallelHashOk

`func (o *ConfigParametersPostgres) GetEnableParallelHashOk() (*string, bool)`

GetEnableParallelHashOk returns a tuple with the EnableParallelHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableParallelHash

`func (o *ConfigParametersPostgres) SetEnableParallelHash(v string)`

SetEnableParallelHash sets EnableParallelHash field to given value.

### HasEnableParallelHash

`func (o *ConfigParametersPostgres) HasEnableParallelHash() bool`

HasEnableParallelHash returns a boolean if a field has been set.

### GetEnablePartitionPruning

`func (o *ConfigParametersPostgres) GetEnablePartitionPruning() string`

GetEnablePartitionPruning returns the EnablePartitionPruning field if non-nil, zero value otherwise.

### GetEnablePartitionPruningOk

`func (o *ConfigParametersPostgres) GetEnablePartitionPruningOk() (*string, bool)`

GetEnablePartitionPruningOk returns a tuple with the EnablePartitionPruning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePartitionPruning

`func (o *ConfigParametersPostgres) SetEnablePartitionPruning(v string)`

SetEnablePartitionPruning sets EnablePartitionPruning field to given value.

### HasEnablePartitionPruning

`func (o *ConfigParametersPostgres) HasEnablePartitionPruning() bool`

HasEnablePartitionPruning returns a boolean if a field has been set.

### GetEnablePartitionwiseJoin

`func (o *ConfigParametersPostgres) GetEnablePartitionwiseJoin() string`

GetEnablePartitionwiseJoin returns the EnablePartitionwiseJoin field if non-nil, zero value otherwise.

### GetEnablePartitionwiseJoinOk

`func (o *ConfigParametersPostgres) GetEnablePartitionwiseJoinOk() (*string, bool)`

GetEnablePartitionwiseJoinOk returns a tuple with the EnablePartitionwiseJoin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePartitionwiseJoin

`func (o *ConfigParametersPostgres) SetEnablePartitionwiseJoin(v string)`

SetEnablePartitionwiseJoin sets EnablePartitionwiseJoin field to given value.

### HasEnablePartitionwiseJoin

`func (o *ConfigParametersPostgres) HasEnablePartitionwiseJoin() bool`

HasEnablePartitionwiseJoin returns a boolean if a field has been set.

### GetEnablePartitionwiseAggregate

`func (o *ConfigParametersPostgres) GetEnablePartitionwiseAggregate() string`

GetEnablePartitionwiseAggregate returns the EnablePartitionwiseAggregate field if non-nil, zero value otherwise.

### GetEnablePartitionwiseAggregateOk

`func (o *ConfigParametersPostgres) GetEnablePartitionwiseAggregateOk() (*string, bool)`

GetEnablePartitionwiseAggregateOk returns a tuple with the EnablePartitionwiseAggregate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePartitionwiseAggregate

`func (o *ConfigParametersPostgres) SetEnablePartitionwiseAggregate(v string)`

SetEnablePartitionwiseAggregate sets EnablePartitionwiseAggregate field to given value.

### HasEnablePartitionwiseAggregate

`func (o *ConfigParametersPostgres) HasEnablePartitionwiseAggregate() bool`

HasEnablePartitionwiseAggregate returns a boolean if a field has been set.

### GetEnableSeqscan

`func (o *ConfigParametersPostgres) GetEnableSeqscan() string`

GetEnableSeqscan returns the EnableSeqscan field if non-nil, zero value otherwise.

### GetEnableSeqscanOk

`func (o *ConfigParametersPostgres) GetEnableSeqscanOk() (*string, bool)`

GetEnableSeqscanOk returns a tuple with the EnableSeqscan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSeqscan

`func (o *ConfigParametersPostgres) SetEnableSeqscan(v string)`

SetEnableSeqscan sets EnableSeqscan field to given value.

### HasEnableSeqscan

`func (o *ConfigParametersPostgres) HasEnableSeqscan() bool`

HasEnableSeqscan returns a boolean if a field has been set.

### GetEnableSort

`func (o *ConfigParametersPostgres) GetEnableSort() string`

GetEnableSort returns the EnableSort field if non-nil, zero value otherwise.

### GetEnableSortOk

`func (o *ConfigParametersPostgres) GetEnableSortOk() (*string, bool)`

GetEnableSortOk returns a tuple with the EnableSort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSort

`func (o *ConfigParametersPostgres) SetEnableSort(v string)`

SetEnableSort sets EnableSort field to given value.

### HasEnableSort

`func (o *ConfigParametersPostgres) HasEnableSort() bool`

HasEnableSort returns a boolean if a field has been set.

### GetEnableTidscan

`func (o *ConfigParametersPostgres) GetEnableTidscan() string`

GetEnableTidscan returns the EnableTidscan field if non-nil, zero value otherwise.

### GetEnableTidscanOk

`func (o *ConfigParametersPostgres) GetEnableTidscanOk() (*string, bool)`

GetEnableTidscanOk returns a tuple with the EnableTidscan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTidscan

`func (o *ConfigParametersPostgres) SetEnableTidscan(v string)`

SetEnableTidscan sets EnableTidscan field to given value.

### HasEnableTidscan

`func (o *ConfigParametersPostgres) HasEnableTidscan() bool`

HasEnableTidscan returns a boolean if a field has been set.

### GetExitOnError

`func (o *ConfigParametersPostgres) GetExitOnError() string`

GetExitOnError returns the ExitOnError field if non-nil, zero value otherwise.

### GetExitOnErrorOk

`func (o *ConfigParametersPostgres) GetExitOnErrorOk() (*string, bool)`

GetExitOnErrorOk returns a tuple with the ExitOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitOnError

`func (o *ConfigParametersPostgres) SetExitOnError(v string)`

SetExitOnError sets ExitOnError field to given value.

### HasExitOnError

`func (o *ConfigParametersPostgres) HasExitOnError() bool`

HasExitOnError returns a boolean if a field has been set.

### GetFromCollapseLimit

`func (o *ConfigParametersPostgres) GetFromCollapseLimit() string`

GetFromCollapseLimit returns the FromCollapseLimit field if non-nil, zero value otherwise.

### GetFromCollapseLimitOk

`func (o *ConfigParametersPostgres) GetFromCollapseLimitOk() (*string, bool)`

GetFromCollapseLimitOk returns a tuple with the FromCollapseLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromCollapseLimit

`func (o *ConfigParametersPostgres) SetFromCollapseLimit(v string)`

SetFromCollapseLimit sets FromCollapseLimit field to given value.

### HasFromCollapseLimit

`func (o *ConfigParametersPostgres) HasFromCollapseLimit() bool`

HasFromCollapseLimit returns a boolean if a field has been set.

### GetJit

`func (o *ConfigParametersPostgres) GetJit() string`

GetJit returns the Jit field if non-nil, zero value otherwise.

### GetJitOk

`func (o *ConfigParametersPostgres) GetJitOk() (*string, bool)`

GetJitOk returns a tuple with the Jit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJit

`func (o *ConfigParametersPostgres) SetJit(v string)`

SetJit sets Jit field to given value.

### HasJit

`func (o *ConfigParametersPostgres) HasJit() bool`

HasJit returns a boolean if a field has been set.

### GetPlanCacheMode

`func (o *ConfigParametersPostgres) GetPlanCacheMode() string`

GetPlanCacheMode returns the PlanCacheMode field if non-nil, zero value otherwise.

### GetPlanCacheModeOk

`func (o *ConfigParametersPostgres) GetPlanCacheModeOk() (*string, bool)`

GetPlanCacheModeOk returns a tuple with the PlanCacheMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanCacheMode

`func (o *ConfigParametersPostgres) SetPlanCacheMode(v string)`

SetPlanCacheMode sets PlanCacheMode field to given value.

### HasPlanCacheMode

`func (o *ConfigParametersPostgres) HasPlanCacheMode() bool`

HasPlanCacheMode returns a boolean if a field has been set.

### GetQuoteAllIdentifiers

`func (o *ConfigParametersPostgres) GetQuoteAllIdentifiers() string`

GetQuoteAllIdentifiers returns the QuoteAllIdentifiers field if non-nil, zero value otherwise.

### GetQuoteAllIdentifiersOk

`func (o *ConfigParametersPostgres) GetQuoteAllIdentifiersOk() (*string, bool)`

GetQuoteAllIdentifiersOk returns a tuple with the QuoteAllIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAllIdentifiers

`func (o *ConfigParametersPostgres) SetQuoteAllIdentifiers(v string)`

SetQuoteAllIdentifiers sets QuoteAllIdentifiers field to given value.

### HasQuoteAllIdentifiers

`func (o *ConfigParametersPostgres) HasQuoteAllIdentifiers() bool`

HasQuoteAllIdentifiers returns a boolean if a field has been set.

### GetStandardConformingStrings

`func (o *ConfigParametersPostgres) GetStandardConformingStrings() string`

GetStandardConformingStrings returns the StandardConformingStrings field if non-nil, zero value otherwise.

### GetStandardConformingStringsOk

`func (o *ConfigParametersPostgres) GetStandardConformingStringsOk() (*string, bool)`

GetStandardConformingStringsOk returns a tuple with the StandardConformingStrings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardConformingStrings

`func (o *ConfigParametersPostgres) SetStandardConformingStrings(v string)`

SetStandardConformingStrings sets StandardConformingStrings field to given value.

### HasStandardConformingStrings

`func (o *ConfigParametersPostgres) HasStandardConformingStrings() bool`

HasStandardConformingStrings returns a boolean if a field has been set.

### GetStatementTimeout

`func (o *ConfigParametersPostgres) GetStatementTimeout() string`

GetStatementTimeout returns the StatementTimeout field if non-nil, zero value otherwise.

### GetStatementTimeoutOk

`func (o *ConfigParametersPostgres) GetStatementTimeoutOk() (*string, bool)`

GetStatementTimeoutOk returns a tuple with the StatementTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementTimeout

`func (o *ConfigParametersPostgres) SetStatementTimeout(v string)`

SetStatementTimeout sets StatementTimeout field to given value.

### HasStatementTimeout

`func (o *ConfigParametersPostgres) HasStatementTimeout() bool`

HasStatementTimeout returns a boolean if a field has been set.

### GetTimezone

`func (o *ConfigParametersPostgres) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ConfigParametersPostgres) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ConfigParametersPostgres) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ConfigParametersPostgres) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetTransformNullEquals

`func (o *ConfigParametersPostgres) GetTransformNullEquals() string`

GetTransformNullEquals returns the TransformNullEquals field if non-nil, zero value otherwise.

### GetTransformNullEqualsOk

`func (o *ConfigParametersPostgres) GetTransformNullEqualsOk() (*string, bool)`

GetTransformNullEqualsOk returns a tuple with the TransformNullEquals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformNullEquals

`func (o *ConfigParametersPostgres) SetTransformNullEquals(v string)`

SetTransformNullEquals sets TransformNullEquals field to given value.

### HasTransformNullEquals

`func (o *ConfigParametersPostgres) HasTransformNullEquals() bool`

HasTransformNullEquals returns a boolean if a field has been set.

### GetMaxLocksPerTransaction

`func (o *ConfigParametersPostgres) GetMaxLocksPerTransaction() string`

GetMaxLocksPerTransaction returns the MaxLocksPerTransaction field if non-nil, zero value otherwise.

### GetMaxLocksPerTransactionOk

`func (o *ConfigParametersPostgres) GetMaxLocksPerTransactionOk() (*string, bool)`

GetMaxLocksPerTransactionOk returns a tuple with the MaxLocksPerTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLocksPerTransaction

`func (o *ConfigParametersPostgres) SetMaxLocksPerTransaction(v string)`

SetMaxLocksPerTransaction sets MaxLocksPerTransaction field to given value.

### HasMaxLocksPerTransaction

`func (o *ConfigParametersPostgres) HasMaxLocksPerTransaction() bool`

HasMaxLocksPerTransaction returns a boolean if a field has been set.

### GetAutovacuumVacuumCostLimit

`func (o *ConfigParametersPostgres) GetAutovacuumVacuumCostLimit() string`

GetAutovacuumVacuumCostLimit returns the AutovacuumVacuumCostLimit field if non-nil, zero value otherwise.

### GetAutovacuumVacuumCostLimitOk

`func (o *ConfigParametersPostgres) GetAutovacuumVacuumCostLimitOk() (*string, bool)`

GetAutovacuumVacuumCostLimitOk returns a tuple with the AutovacuumVacuumCostLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutovacuumVacuumCostLimit

`func (o *ConfigParametersPostgres) SetAutovacuumVacuumCostLimit(v string)`

SetAutovacuumVacuumCostLimit sets AutovacuumVacuumCostLimit field to given value.

### HasAutovacuumVacuumCostLimit

`func (o *ConfigParametersPostgres) HasAutovacuumVacuumCostLimit() bool`

HasAutovacuumVacuumCostLimit returns a boolean if a field has been set.

### GetCheckpointTimeout

`func (o *ConfigParametersPostgres) GetCheckpointTimeout() string`

GetCheckpointTimeout returns the CheckpointTimeout field if non-nil, zero value otherwise.

### GetCheckpointTimeoutOk

`func (o *ConfigParametersPostgres) GetCheckpointTimeoutOk() (*string, bool)`

GetCheckpointTimeoutOk returns a tuple with the CheckpointTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckpointTimeout

`func (o *ConfigParametersPostgres) SetCheckpointTimeout(v string)`

SetCheckpointTimeout sets CheckpointTimeout field to given value.

### HasCheckpointTimeout

`func (o *ConfigParametersPostgres) HasCheckpointTimeout() bool`

HasCheckpointTimeout returns a boolean if a field has been set.

### GetCheckpointCompletionTarget

`func (o *ConfigParametersPostgres) GetCheckpointCompletionTarget() string`

GetCheckpointCompletionTarget returns the CheckpointCompletionTarget field if non-nil, zero value otherwise.

### GetCheckpointCompletionTargetOk

`func (o *ConfigParametersPostgres) GetCheckpointCompletionTargetOk() (*string, bool)`

GetCheckpointCompletionTargetOk returns a tuple with the CheckpointCompletionTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckpointCompletionTarget

`func (o *ConfigParametersPostgres) SetCheckpointCompletionTarget(v string)`

SetCheckpointCompletionTarget sets CheckpointCompletionTarget field to given value.

### HasCheckpointCompletionTarget

`func (o *ConfigParametersPostgres) HasCheckpointCompletionTarget() bool`

HasCheckpointCompletionTarget returns a boolean if a field has been set.

### GetWalCompression

`func (o *ConfigParametersPostgres) GetWalCompression() string`

GetWalCompression returns the WalCompression field if non-nil, zero value otherwise.

### GetWalCompressionOk

`func (o *ConfigParametersPostgres) GetWalCompressionOk() (*string, bool)`

GetWalCompressionOk returns a tuple with the WalCompression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalCompression

`func (o *ConfigParametersPostgres) SetWalCompression(v string)`

SetWalCompression sets WalCompression field to given value.

### HasWalCompression

`func (o *ConfigParametersPostgres) HasWalCompression() bool`

HasWalCompression returns a boolean if a field has been set.

### GetRandomPageCost

`func (o *ConfigParametersPostgres) GetRandomPageCost() string`

GetRandomPageCost returns the RandomPageCost field if non-nil, zero value otherwise.

### GetRandomPageCostOk

`func (o *ConfigParametersPostgres) GetRandomPageCostOk() (*string, bool)`

GetRandomPageCostOk returns a tuple with the RandomPageCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomPageCost

`func (o *ConfigParametersPostgres) SetRandomPageCost(v string)`

SetRandomPageCost sets RandomPageCost field to given value.

### HasRandomPageCost

`func (o *ConfigParametersPostgres) HasRandomPageCost() bool`

HasRandomPageCost returns a boolean if a field has been set.

### GetEffectiveIoConcurrency

`func (o *ConfigParametersPostgres) GetEffectiveIoConcurrency() string`

GetEffectiveIoConcurrency returns the EffectiveIoConcurrency field if non-nil, zero value otherwise.

### GetEffectiveIoConcurrencyOk

`func (o *ConfigParametersPostgres) GetEffectiveIoConcurrencyOk() (*string, bool)`

GetEffectiveIoConcurrencyOk returns a tuple with the EffectiveIoConcurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveIoConcurrency

`func (o *ConfigParametersPostgres) SetEffectiveIoConcurrency(v string)`

SetEffectiveIoConcurrency sets EffectiveIoConcurrency field to given value.

### HasEffectiveIoConcurrency

`func (o *ConfigParametersPostgres) HasEffectiveIoConcurrency() bool`

HasEffectiveIoConcurrency returns a boolean if a field has been set.

### GetLogLockWaits

`func (o *ConfigParametersPostgres) GetLogLockWaits() string`

GetLogLockWaits returns the LogLockWaits field if non-nil, zero value otherwise.

### GetLogLockWaitsOk

`func (o *ConfigParametersPostgres) GetLogLockWaitsOk() (*string, bool)`

GetLogLockWaitsOk returns a tuple with the LogLockWaits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLockWaits

`func (o *ConfigParametersPostgres) SetLogLockWaits(v string)`

SetLogLockWaits sets LogLockWaits field to given value.

### HasLogLockWaits

`func (o *ConfigParametersPostgres) HasLogLockWaits() bool`

HasLogLockWaits returns a boolean if a field has been set.

### GetLogTempFiles

`func (o *ConfigParametersPostgres) GetLogTempFiles() string`

GetLogTempFiles returns the LogTempFiles field if non-nil, zero value otherwise.

### GetLogTempFilesOk

`func (o *ConfigParametersPostgres) GetLogTempFilesOk() (*string, bool)`

GetLogTempFilesOk returns a tuple with the LogTempFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogTempFiles

`func (o *ConfigParametersPostgres) SetLogTempFiles(v string)`

SetLogTempFiles sets LogTempFiles field to given value.

### HasLogTempFiles

`func (o *ConfigParametersPostgres) HasLogTempFiles() bool`

HasLogTempFiles returns a boolean if a field has been set.

### GetTrackIoTiming

`func (o *ConfigParametersPostgres) GetTrackIoTiming() string`

GetTrackIoTiming returns the TrackIoTiming field if non-nil, zero value otherwise.

### GetTrackIoTimingOk

`func (o *ConfigParametersPostgres) GetTrackIoTimingOk() (*string, bool)`

GetTrackIoTimingOk returns a tuple with the TrackIoTiming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackIoTiming

`func (o *ConfigParametersPostgres) SetTrackIoTiming(v string)`

SetTrackIoTiming sets TrackIoTiming field to given value.

### HasTrackIoTiming

`func (o *ConfigParametersPostgres) HasTrackIoTiming() bool`

HasTrackIoTiming returns a boolean if a field has been set.

### GetMaintenanceWorkMem

`func (o *ConfigParametersPostgres) GetMaintenanceWorkMem() string`

GetMaintenanceWorkMem returns the MaintenanceWorkMem field if non-nil, zero value otherwise.

### GetMaintenanceWorkMemOk

`func (o *ConfigParametersPostgres) GetMaintenanceWorkMemOk() (*string, bool)`

GetMaintenanceWorkMemOk returns a tuple with the MaintenanceWorkMem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceWorkMem

`func (o *ConfigParametersPostgres) SetMaintenanceWorkMem(v string)`

SetMaintenanceWorkMem sets MaintenanceWorkMem field to given value.

### HasMaintenanceWorkMem

`func (o *ConfigParametersPostgres) HasMaintenanceWorkMem() bool`

HasMaintenanceWorkMem returns a boolean if a field has been set.

### GetIdleSessionTimeout

`func (o *ConfigParametersPostgres) GetIdleSessionTimeout() string`

GetIdleSessionTimeout returns the IdleSessionTimeout field if non-nil, zero value otherwise.

### GetIdleSessionTimeoutOk

`func (o *ConfigParametersPostgres) GetIdleSessionTimeoutOk() (*string, bool)`

GetIdleSessionTimeoutOk returns a tuple with the IdleSessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleSessionTimeout

`func (o *ConfigParametersPostgres) SetIdleSessionTimeout(v string)`

SetIdleSessionTimeout sets IdleSessionTimeout field to given value.

### HasIdleSessionTimeout

`func (o *ConfigParametersPostgres) HasIdleSessionTimeout() bool`

HasIdleSessionTimeout returns a boolean if a field has been set.

### GetIoMethod

`func (o *ConfigParametersPostgres) GetIoMethod() string`

GetIoMethod returns the IoMethod field if non-nil, zero value otherwise.

### GetIoMethodOk

`func (o *ConfigParametersPostgres) GetIoMethodOk() (*string, bool)`

GetIoMethodOk returns a tuple with the IoMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoMethod

`func (o *ConfigParametersPostgres) SetIoMethod(v string)`

SetIoMethod sets IoMethod field to given value.

### HasIoMethod

`func (o *ConfigParametersPostgres) HasIoMethod() bool`

HasIoMethod returns a boolean if a field has been set.

### GetIoWorkers

`func (o *ConfigParametersPostgres) GetIoWorkers() string`

GetIoWorkers returns the IoWorkers field if non-nil, zero value otherwise.

### GetIoWorkersOk

`func (o *ConfigParametersPostgres) GetIoWorkersOk() (*string, bool)`

GetIoWorkersOk returns a tuple with the IoWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoWorkers

`func (o *ConfigParametersPostgres) SetIoWorkers(v string)`

SetIoWorkers sets IoWorkers field to given value.

### HasIoWorkers

`func (o *ConfigParametersPostgres) HasIoWorkers() bool`

HasIoWorkers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


