# GetDatabaseDefaultParameters200ResponseConfigParams

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
**QueryCacheType** | Pointer to **string** | Тип кэша запросов (&#x60;mysql5&#x60;). | [optional] 
**QueryCacheSize** | Pointer to **string** | Объем памяти, выделяемый для кэширования результатов запросов (&#x60;mysql5&#x60;). | [optional] 
**QueryCacheLimit** | Pointer to **string** | Максимальный размер результата запроса, который может быть закэширован (&#x60;mysql5&#x60;). | [optional] 
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
**ClientOutputBufferLimitNormal** | Pointer to **string** | Ограничение буфера вывода для обычных клиентских подключений. Формат: &#x60;hard-limit soft-limit soft-seconds&#x60; (&#x60;valkey&#x60; | &#x60;valkey7&#x60; | &#x60;valkey8_1&#x60; | &#x60;valkey9_1&#x60;). | [optional] 
**NotifyKeyspaceEvents** | Pointer to **string** | Настройка уведомлений о событиях пространства ключей (&#x60;valkey&#x60; | &#x60;valkey7&#x60; | &#x60;valkey8_1&#x60; | &#x60;valkey9_1&#x60;). | [optional] 
**ClientOutputBufferLimitPubsub** | Pointer to **string** | Ограничение буфера вывода для клиентов pub/sub. Формат: &#x60;hard-limit soft-limit soft-seconds&#x60; (&#x60;valkey&#x60; | &#x60;valkey7&#x60; | &#x60;valkey8_1&#x60; | &#x60;valkey9_1&#x60;). | [optional] 
**Databases** | Pointer to **string** | Количество логических баз данных на сервере (&#x60;valkey&#x60; | &#x60;valkey7&#x60; | &#x60;valkey8_1&#x60; | &#x60;valkey9_1&#x60;). | [optional] 
**Timeout** | Pointer to **string** | Время ожидания в секундах перед закрытием неактивного клиентского соединения. &#x60;0&#x60; — отключено (&#x60;valkey&#x60; | &#x60;valkey7&#x60; | &#x60;valkey8_1&#x60; | &#x60;valkey9_1&#x60;). | [optional] 
**MaxmemoryPolicy** | Pointer to **string** | Политика вытеснения ключей при достижении лимита памяти (&#x60;valkey&#x60; | &#x60;valkey7&#x60; | &#x60;valkey8_1&#x60; | &#x60;valkey9_1&#x60;). | [optional] 
**SlowlogLogSlowerThan** | Pointer to **string** | Минимальное время выполнения команды в микросекундах для записи в журнал медленных команд (&#x60;valkey&#x60; | &#x60;valkey7&#x60; | &#x60;valkey8_1&#x60; | &#x60;valkey9_1&#x60;). | [optional] 
**SlowlogMaxLen** | Pointer to **string** | Максимальное количество записей, хранящихся в журнале медленных команд (&#x60;valkey&#x60; | &#x60;valkey7&#x60; | &#x60;valkey8_1&#x60; | &#x60;valkey9_1&#x60;). | [optional] 
**Save** | Pointer to **string** | Условие создания снимка RDB на диск. Формат: &#x60;seconds changes&#x60; — сохранение выполняется, если за указанное время было сделано не менее указанного количества изменений (&#x60;valkey&#x60; | &#x60;valkey7&#x60; | &#x60;valkey8_1&#x60; | &#x60;valkey9_1&#x60;). | [optional] 
**Appendonly** | Pointer to **string** | Включение режима AOF (Append Only File) для персистентного хранения данных (&#x60;valkey&#x60; | &#x60;valkey7&#x60; | &#x60;valkey8_1&#x60; | &#x60;valkey9_1&#x60;). | [optional] 
**Appendfsync** | Pointer to **string** | Режим синхронизации AOF-файла с диском: &#x60;always&#x60; — при каждой записи, &#x60;everysec&#x60; — раз в секунду, &#x60;no&#x60; — управление передаётся ОС (&#x60;valkey&#x60; | &#x60;valkey7&#x60; | &#x60;valkey8_1&#x60; | &#x60;valkey9_1&#x60;). | [optional] 
**TcpKeepalive** | Pointer to **string** | Интервал проверки активности TCP-соединения в секундах. &#x60;0&#x60; — отключено (&#x60;valkey&#x60; | &#x60;valkey7&#x60; | &#x60;valkey8_1&#x60; | &#x60;valkey9_1&#x60;). | [optional] 

## Methods

### NewGetDatabaseDefaultParameters200ResponseConfigParams

`func NewGetDatabaseDefaultParameters200ResponseConfigParams() *GetDatabaseDefaultParameters200ResponseConfigParams`

NewGetDatabaseDefaultParameters200ResponseConfigParams instantiates a new GetDatabaseDefaultParameters200ResponseConfigParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDatabaseDefaultParameters200ResponseConfigParamsWithDefaults

`func NewGetDatabaseDefaultParameters200ResponseConfigParamsWithDefaults() *GetDatabaseDefaultParameters200ResponseConfigParams`

NewGetDatabaseDefaultParameters200ResponseConfigParamsWithDefaults instantiates a new GetDatabaseDefaultParameters200ResponseConfigParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJoinBufferSize

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetJoinBufferSize() string`

GetJoinBufferSize returns the JoinBufferSize field if non-nil, zero value otherwise.

### GetJoinBufferSizeOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetJoinBufferSizeOk() (*string, bool)`

GetJoinBufferSizeOk returns a tuple with the JoinBufferSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinBufferSize

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetJoinBufferSize(v string)`

SetJoinBufferSize sets JoinBufferSize field to given value.

### HasJoinBufferSize

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasJoinBufferSize() bool`

HasJoinBufferSize returns a boolean if a field has been set.

### GetMaxConnections

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetMaxConnections() string`

GetMaxConnections returns the MaxConnections field if non-nil, zero value otherwise.

### GetMaxConnectionsOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetMaxConnectionsOk() (*string, bool)`

GetMaxConnectionsOk returns a tuple with the MaxConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnections

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetMaxConnections(v string)`

SetMaxConnections sets MaxConnections field to given value.

### HasMaxConnections

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasMaxConnections() bool`

HasMaxConnections returns a boolean if a field has been set.

### GetSortBufferSize

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetSortBufferSize() string`

GetSortBufferSize returns the SortBufferSize field if non-nil, zero value otherwise.

### GetSortBufferSizeOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetSortBufferSizeOk() (*string, bool)`

GetSortBufferSizeOk returns a tuple with the SortBufferSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBufferSize

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetSortBufferSize(v string)`

SetSortBufferSize sets SortBufferSize field to given value.

### HasSortBufferSize

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasSortBufferSize() bool`

HasSortBufferSize returns a boolean if a field has been set.

### GetThreadCacheSize

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetThreadCacheSize() string`

GetThreadCacheSize returns the ThreadCacheSize field if non-nil, zero value otherwise.

### GetThreadCacheSizeOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetThreadCacheSizeOk() (*string, bool)`

GetThreadCacheSizeOk returns a tuple with the ThreadCacheSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadCacheSize

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetThreadCacheSize(v string)`

SetThreadCacheSize sets ThreadCacheSize field to given value.

### HasThreadCacheSize

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasThreadCacheSize() bool`

HasThreadCacheSize returns a boolean if a field has been set.

### GetInnodbBufferPoolSize

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetInnodbBufferPoolSize() string`

GetInnodbBufferPoolSize returns the InnodbBufferPoolSize field if non-nil, zero value otherwise.

### GetInnodbBufferPoolSizeOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetInnodbBufferPoolSizeOk() (*string, bool)`

GetInnodbBufferPoolSizeOk returns a tuple with the InnodbBufferPoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnodbBufferPoolSize

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetInnodbBufferPoolSize(v string)`

SetInnodbBufferPoolSize sets InnodbBufferPoolSize field to given value.

### HasInnodbBufferPoolSize

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasInnodbBufferPoolSize() bool`

HasInnodbBufferPoolSize returns a boolean if a field has been set.

### GetAutoIncrementIncrement

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetAutoIncrementIncrement() string`

GetAutoIncrementIncrement returns the AutoIncrementIncrement field if non-nil, zero value otherwise.

### GetAutoIncrementIncrementOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetAutoIncrementIncrementOk() (*string, bool)`

GetAutoIncrementIncrementOk returns a tuple with the AutoIncrementIncrement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoIncrementIncrement

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetAutoIncrementIncrement(v string)`

SetAutoIncrementIncrement sets AutoIncrementIncrement field to given value.

### HasAutoIncrementIncrement

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasAutoIncrementIncrement() bool`

HasAutoIncrementIncrement returns a boolean if a field has been set.

### GetAutoIncrementOffset

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetAutoIncrementOffset() string`

GetAutoIncrementOffset returns the AutoIncrementOffset field if non-nil, zero value otherwise.

### GetAutoIncrementOffsetOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetAutoIncrementOffsetOk() (*string, bool)`

GetAutoIncrementOffsetOk returns a tuple with the AutoIncrementOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoIncrementOffset

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetAutoIncrementOffset(v string)`

SetAutoIncrementOffset sets AutoIncrementOffset field to given value.

### HasAutoIncrementOffset

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasAutoIncrementOffset() bool`

HasAutoIncrementOffset returns a boolean if a field has been set.

### GetInnodbIoCapacity

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetInnodbIoCapacity() string`

GetInnodbIoCapacity returns the InnodbIoCapacity field if non-nil, zero value otherwise.

### GetInnodbIoCapacityOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetInnodbIoCapacityOk() (*string, bool)`

GetInnodbIoCapacityOk returns a tuple with the InnodbIoCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnodbIoCapacity

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetInnodbIoCapacity(v string)`

SetInnodbIoCapacity sets InnodbIoCapacity field to given value.

### HasInnodbIoCapacity

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasInnodbIoCapacity() bool`

HasInnodbIoCapacity returns a boolean if a field has been set.

### GetInnodbPurgeThreads

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetInnodbPurgeThreads() string`

GetInnodbPurgeThreads returns the InnodbPurgeThreads field if non-nil, zero value otherwise.

### GetInnodbPurgeThreadsOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetInnodbPurgeThreadsOk() (*string, bool)`

GetInnodbPurgeThreadsOk returns a tuple with the InnodbPurgeThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnodbPurgeThreads

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetInnodbPurgeThreads(v string)`

SetInnodbPurgeThreads sets InnodbPurgeThreads field to given value.

### HasInnodbPurgeThreads

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasInnodbPurgeThreads() bool`

HasInnodbPurgeThreads returns a boolean if a field has been set.

### GetInnodbReadIoThreads

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetInnodbReadIoThreads() string`

GetInnodbReadIoThreads returns the InnodbReadIoThreads field if non-nil, zero value otherwise.

### GetInnodbReadIoThreadsOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetInnodbReadIoThreadsOk() (*string, bool)`

GetInnodbReadIoThreadsOk returns a tuple with the InnodbReadIoThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnodbReadIoThreads

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetInnodbReadIoThreads(v string)`

SetInnodbReadIoThreads sets InnodbReadIoThreads field to given value.

### HasInnodbReadIoThreads

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasInnodbReadIoThreads() bool`

HasInnodbReadIoThreads returns a boolean if a field has been set.

### GetInnodbThreadConcurrency

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetInnodbThreadConcurrency() string`

GetInnodbThreadConcurrency returns the InnodbThreadConcurrency field if non-nil, zero value otherwise.

### GetInnodbThreadConcurrencyOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetInnodbThreadConcurrencyOk() (*string, bool)`

GetInnodbThreadConcurrencyOk returns a tuple with the InnodbThreadConcurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnodbThreadConcurrency

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetInnodbThreadConcurrency(v string)`

SetInnodbThreadConcurrency sets InnodbThreadConcurrency field to given value.

### HasInnodbThreadConcurrency

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasInnodbThreadConcurrency() bool`

HasInnodbThreadConcurrency returns a boolean if a field has been set.

### GetInnodbWriteIoThreads

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetInnodbWriteIoThreads() string`

GetInnodbWriteIoThreads returns the InnodbWriteIoThreads field if non-nil, zero value otherwise.

### GetInnodbWriteIoThreadsOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetInnodbWriteIoThreadsOk() (*string, bool)`

GetInnodbWriteIoThreadsOk returns a tuple with the InnodbWriteIoThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnodbWriteIoThreads

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetInnodbWriteIoThreads(v string)`

SetInnodbWriteIoThreads sets InnodbWriteIoThreads field to given value.

### HasInnodbWriteIoThreads

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasInnodbWriteIoThreads() bool`

HasInnodbWriteIoThreads returns a boolean if a field has been set.

### GetInnodbLogFileSize

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetInnodbLogFileSize() string`

GetInnodbLogFileSize returns the InnodbLogFileSize field if non-nil, zero value otherwise.

### GetInnodbLogFileSizeOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetInnodbLogFileSizeOk() (*string, bool)`

GetInnodbLogFileSizeOk returns a tuple with the InnodbLogFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnodbLogFileSize

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetInnodbLogFileSize(v string)`

SetInnodbLogFileSize sets InnodbLogFileSize field to given value.

### HasInnodbLogFileSize

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasInnodbLogFileSize() bool`

HasInnodbLogFileSize returns a boolean if a field has been set.

### GetMaxAllowedPacket

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetMaxAllowedPacket() string`

GetMaxAllowedPacket returns the MaxAllowedPacket field if non-nil, zero value otherwise.

### GetMaxAllowedPacketOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetMaxAllowedPacketOk() (*string, bool)`

GetMaxAllowedPacketOk returns a tuple with the MaxAllowedPacket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAllowedPacket

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetMaxAllowedPacket(v string)`

SetMaxAllowedPacket sets MaxAllowedPacket field to given value.

### HasMaxAllowedPacket

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasMaxAllowedPacket() bool`

HasMaxAllowedPacket returns a boolean if a field has been set.

### GetMaxHeapTableSize

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetMaxHeapTableSize() string`

GetMaxHeapTableSize returns the MaxHeapTableSize field if non-nil, zero value otherwise.

### GetMaxHeapTableSizeOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetMaxHeapTableSizeOk() (*string, bool)`

GetMaxHeapTableSizeOk returns a tuple with the MaxHeapTableSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHeapTableSize

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetMaxHeapTableSize(v string)`

SetMaxHeapTableSize sets MaxHeapTableSize field to given value.

### HasMaxHeapTableSize

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasMaxHeapTableSize() bool`

HasMaxHeapTableSize returns a boolean if a field has been set.

### GetSqlMode

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetSqlMode() string`

GetSqlMode returns the SqlMode field if non-nil, zero value otherwise.

### GetSqlModeOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetSqlModeOk() (*string, bool)`

GetSqlModeOk returns a tuple with the SqlMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlMode

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetSqlMode(v string)`

SetSqlMode sets SqlMode field to given value.

### HasSqlMode

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasSqlMode() bool`

HasSqlMode returns a boolean if a field has been set.

### GetQueryCacheType

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetQueryCacheType() string`

GetQueryCacheType returns the QueryCacheType field if non-nil, zero value otherwise.

### GetQueryCacheTypeOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetQueryCacheTypeOk() (*string, bool)`

GetQueryCacheTypeOk returns a tuple with the QueryCacheType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryCacheType

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetQueryCacheType(v string)`

SetQueryCacheType sets QueryCacheType field to given value.

### HasQueryCacheType

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasQueryCacheType() bool`

HasQueryCacheType returns a boolean if a field has been set.

### GetQueryCacheSize

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetQueryCacheSize() string`

GetQueryCacheSize returns the QueryCacheSize field if non-nil, zero value otherwise.

### GetQueryCacheSizeOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetQueryCacheSizeOk() (*string, bool)`

GetQueryCacheSizeOk returns a tuple with the QueryCacheSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryCacheSize

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetQueryCacheSize(v string)`

SetQueryCacheSize sets QueryCacheSize field to given value.

### HasQueryCacheSize

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasQueryCacheSize() bool`

HasQueryCacheSize returns a boolean if a field has been set.

### GetQueryCacheLimit

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetQueryCacheLimit() string`

GetQueryCacheLimit returns the QueryCacheLimit field if non-nil, zero value otherwise.

### GetQueryCacheLimitOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetQueryCacheLimitOk() (*string, bool)`

GetQueryCacheLimitOk returns a tuple with the QueryCacheLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryCacheLimit

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetQueryCacheLimit(v string)`

SetQueryCacheLimit sets QueryCacheLimit field to given value.

### HasQueryCacheLimit

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasQueryCacheLimit() bool`

HasQueryCacheLimit returns a boolean if a field has been set.

### GetInnodbFlushLogAtTrxCommit

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetInnodbFlushLogAtTrxCommit() string`

GetInnodbFlushLogAtTrxCommit returns the InnodbFlushLogAtTrxCommit field if non-nil, zero value otherwise.

### GetInnodbFlushLogAtTrxCommitOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetInnodbFlushLogAtTrxCommitOk() (*string, bool)`

GetInnodbFlushLogAtTrxCommitOk returns a tuple with the InnodbFlushLogAtTrxCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnodbFlushLogAtTrxCommit

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetInnodbFlushLogAtTrxCommit(v string)`

SetInnodbFlushLogAtTrxCommit sets InnodbFlushLogAtTrxCommit field to given value.

### HasInnodbFlushLogAtTrxCommit

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasInnodbFlushLogAtTrxCommit() bool`

HasInnodbFlushLogAtTrxCommit returns a boolean if a field has been set.

### GetTransactionIsolation

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetTransactionIsolation() string`

GetTransactionIsolation returns the TransactionIsolation field if non-nil, zero value otherwise.

### GetTransactionIsolationOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetTransactionIsolationOk() (*string, bool)`

GetTransactionIsolationOk returns a tuple with the TransactionIsolation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIsolation

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetTransactionIsolation(v string)`

SetTransactionIsolation sets TransactionIsolation field to given value.

### HasTransactionIsolation

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasTransactionIsolation() bool`

HasTransactionIsolation returns a boolean if a field has been set.

### GetLongQueryTime

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetLongQueryTime() string`

GetLongQueryTime returns the LongQueryTime field if non-nil, zero value otherwise.

### GetLongQueryTimeOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetLongQueryTimeOk() (*string, bool)`

GetLongQueryTimeOk returns a tuple with the LongQueryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongQueryTime

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetLongQueryTime(v string)`

SetLongQueryTime sets LongQueryTime field to given value.

### HasLongQueryTime

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasLongQueryTime() bool`

HasLongQueryTime returns a boolean if a field has been set.

### GetTmpTableSize

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetTmpTableSize() string`

GetTmpTableSize returns the TmpTableSize field if non-nil, zero value otherwise.

### GetTmpTableSizeOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetTmpTableSizeOk() (*string, bool)`

GetTmpTableSizeOk returns a tuple with the TmpTableSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmpTableSize

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetTmpTableSize(v string)`

SetTmpTableSize sets TmpTableSize field to given value.

### HasTmpTableSize

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasTmpTableSize() bool`

HasTmpTableSize returns a boolean if a field has been set.

### GetTableOpenCache

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetTableOpenCache() string`

GetTableOpenCache returns the TableOpenCache field if non-nil, zero value otherwise.

### GetTableOpenCacheOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetTableOpenCacheOk() (*string, bool)`

GetTableOpenCacheOk returns a tuple with the TableOpenCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableOpenCache

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetTableOpenCache(v string)`

SetTableOpenCache sets TableOpenCache field to given value.

### HasTableOpenCache

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasTableOpenCache() bool`

HasTableOpenCache returns a boolean if a field has been set.

### GetTableOpenCacheInstances

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetTableOpenCacheInstances() string`

GetTableOpenCacheInstances returns the TableOpenCacheInstances field if non-nil, zero value otherwise.

### GetTableOpenCacheInstancesOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetTableOpenCacheInstancesOk() (*string, bool)`

GetTableOpenCacheInstancesOk returns a tuple with the TableOpenCacheInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableOpenCacheInstances

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetTableOpenCacheInstances(v string)`

SetTableOpenCacheInstances sets TableOpenCacheInstances field to given value.

### HasTableOpenCacheInstances

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasTableOpenCacheInstances() bool`

HasTableOpenCacheInstances returns a boolean if a field has been set.

### GetInnodbFlushMethod

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetInnodbFlushMethod() string`

GetInnodbFlushMethod returns the InnodbFlushMethod field if non-nil, zero value otherwise.

### GetInnodbFlushMethodOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetInnodbFlushMethodOk() (*string, bool)`

GetInnodbFlushMethodOk returns a tuple with the InnodbFlushMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnodbFlushMethod

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetInnodbFlushMethod(v string)`

SetInnodbFlushMethod sets InnodbFlushMethod field to given value.

### HasInnodbFlushMethod

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasInnodbFlushMethod() bool`

HasInnodbFlushMethod returns a boolean if a field has been set.

### GetInnodbStrictMode

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetInnodbStrictMode() string`

GetInnodbStrictMode returns the InnodbStrictMode field if non-nil, zero value otherwise.

### GetInnodbStrictModeOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetInnodbStrictModeOk() (*string, bool)`

GetInnodbStrictModeOk returns a tuple with the InnodbStrictMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnodbStrictMode

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetInnodbStrictMode(v string)`

SetInnodbStrictMode sets InnodbStrictMode field to given value.

### HasInnodbStrictMode

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasInnodbStrictMode() bool`

HasInnodbStrictMode returns a boolean if a field has been set.

### GetSlowQueryLog

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetSlowQueryLog() string`

GetSlowQueryLog returns the SlowQueryLog field if non-nil, zero value otherwise.

### GetSlowQueryLogOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetSlowQueryLogOk() (*string, bool)`

GetSlowQueryLogOk returns a tuple with the SlowQueryLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlowQueryLog

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetSlowQueryLog(v string)`

SetSlowQueryLog sets SlowQueryLog field to given value.

### HasSlowQueryLog

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasSlowQueryLog() bool`

HasSlowQueryLog returns a boolean if a field has been set.

### GetBinlogCacheSize

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetBinlogCacheSize() string`

GetBinlogCacheSize returns the BinlogCacheSize field if non-nil, zero value otherwise.

### GetBinlogCacheSizeOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetBinlogCacheSizeOk() (*string, bool)`

GetBinlogCacheSizeOk returns a tuple with the BinlogCacheSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinlogCacheSize

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetBinlogCacheSize(v string)`

SetBinlogCacheSize sets BinlogCacheSize field to given value.

### HasBinlogCacheSize

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasBinlogCacheSize() bool`

HasBinlogCacheSize returns a boolean if a field has been set.

### GetBinlogGroupCommitSyncDelay

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetBinlogGroupCommitSyncDelay() string`

GetBinlogGroupCommitSyncDelay returns the BinlogGroupCommitSyncDelay field if non-nil, zero value otherwise.

### GetBinlogGroupCommitSyncDelayOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetBinlogGroupCommitSyncDelayOk() (*string, bool)`

GetBinlogGroupCommitSyncDelayOk returns a tuple with the BinlogGroupCommitSyncDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinlogGroupCommitSyncDelay

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetBinlogGroupCommitSyncDelay(v string)`

SetBinlogGroupCommitSyncDelay sets BinlogGroupCommitSyncDelay field to given value.

### HasBinlogGroupCommitSyncDelay

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasBinlogGroupCommitSyncDelay() bool`

HasBinlogGroupCommitSyncDelay returns a boolean if a field has been set.

### GetBinlogRowImage

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetBinlogRowImage() string`

GetBinlogRowImage returns the BinlogRowImage field if non-nil, zero value otherwise.

### GetBinlogRowImageOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetBinlogRowImageOk() (*string, bool)`

GetBinlogRowImageOk returns a tuple with the BinlogRowImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinlogRowImage

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetBinlogRowImage(v string)`

SetBinlogRowImage sets BinlogRowImage field to given value.

### HasBinlogRowImage

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasBinlogRowImage() bool`

HasBinlogRowImage returns a boolean if a field has been set.

### GetBinlogRowsQueryLogEvents

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetBinlogRowsQueryLogEvents() string`

GetBinlogRowsQueryLogEvents returns the BinlogRowsQueryLogEvents field if non-nil, zero value otherwise.

### GetBinlogRowsQueryLogEventsOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetBinlogRowsQueryLogEventsOk() (*string, bool)`

GetBinlogRowsQueryLogEventsOk returns a tuple with the BinlogRowsQueryLogEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinlogRowsQueryLogEvents

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetBinlogRowsQueryLogEvents(v string)`

SetBinlogRowsQueryLogEvents sets BinlogRowsQueryLogEvents field to given value.

### HasBinlogRowsQueryLogEvents

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasBinlogRowsQueryLogEvents() bool`

HasBinlogRowsQueryLogEvents returns a boolean if a field has been set.

### GetCharacterSetServer

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetCharacterSetServer() string`

GetCharacterSetServer returns the CharacterSetServer field if non-nil, zero value otherwise.

### GetCharacterSetServerOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetCharacterSetServerOk() (*string, bool)`

GetCharacterSetServerOk returns a tuple with the CharacterSetServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacterSetServer

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetCharacterSetServer(v string)`

SetCharacterSetServer sets CharacterSetServer field to given value.

### HasCharacterSetServer

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasCharacterSetServer() bool`

HasCharacterSetServer returns a boolean if a field has been set.

### GetExplicitDefaultsForTimestamp

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetExplicitDefaultsForTimestamp() string`

GetExplicitDefaultsForTimestamp returns the ExplicitDefaultsForTimestamp field if non-nil, zero value otherwise.

### GetExplicitDefaultsForTimestampOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetExplicitDefaultsForTimestampOk() (*string, bool)`

GetExplicitDefaultsForTimestampOk returns a tuple with the ExplicitDefaultsForTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitDefaultsForTimestamp

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetExplicitDefaultsForTimestamp(v string)`

SetExplicitDefaultsForTimestamp sets ExplicitDefaultsForTimestamp field to given value.

### HasExplicitDefaultsForTimestamp

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasExplicitDefaultsForTimestamp() bool`

HasExplicitDefaultsForTimestamp returns a boolean if a field has been set.

### GetGroupConcatMaxLen

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetGroupConcatMaxLen() string`

GetGroupConcatMaxLen returns the GroupConcatMaxLen field if non-nil, zero value otherwise.

### GetGroupConcatMaxLenOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetGroupConcatMaxLenOk() (*string, bool)`

GetGroupConcatMaxLenOk returns a tuple with the GroupConcatMaxLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupConcatMaxLen

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetGroupConcatMaxLen(v string)`

SetGroupConcatMaxLen sets GroupConcatMaxLen field to given value.

### HasGroupConcatMaxLen

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasGroupConcatMaxLen() bool`

HasGroupConcatMaxLen returns a boolean if a field has been set.

### GetInnodbAdaptiveHashIndex

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetInnodbAdaptiveHashIndex() string`

GetInnodbAdaptiveHashIndex returns the InnodbAdaptiveHashIndex field if non-nil, zero value otherwise.

### GetInnodbAdaptiveHashIndexOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetInnodbAdaptiveHashIndexOk() (*string, bool)`

GetInnodbAdaptiveHashIndexOk returns a tuple with the InnodbAdaptiveHashIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnodbAdaptiveHashIndex

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetInnodbAdaptiveHashIndex(v string)`

SetInnodbAdaptiveHashIndex sets InnodbAdaptiveHashIndex field to given value.

### HasInnodbAdaptiveHashIndex

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasInnodbAdaptiveHashIndex() bool`

HasInnodbAdaptiveHashIndex returns a boolean if a field has been set.

### GetInnodbLockWaitTimeout

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetInnodbLockWaitTimeout() string`

GetInnodbLockWaitTimeout returns the InnodbLockWaitTimeout field if non-nil, zero value otherwise.

### GetInnodbLockWaitTimeoutOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetInnodbLockWaitTimeoutOk() (*string, bool)`

GetInnodbLockWaitTimeoutOk returns a tuple with the InnodbLockWaitTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnodbLockWaitTimeout

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetInnodbLockWaitTimeout(v string)`

SetInnodbLockWaitTimeout sets InnodbLockWaitTimeout field to given value.

### HasInnodbLockWaitTimeout

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasInnodbLockWaitTimeout() bool`

HasInnodbLockWaitTimeout returns a boolean if a field has been set.

### GetInnodbNumaInterleave

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetInnodbNumaInterleave() string`

GetInnodbNumaInterleave returns the InnodbNumaInterleave field if non-nil, zero value otherwise.

### GetInnodbNumaInterleaveOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetInnodbNumaInterleaveOk() (*string, bool)`

GetInnodbNumaInterleaveOk returns a tuple with the InnodbNumaInterleave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnodbNumaInterleave

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetInnodbNumaInterleave(v string)`

SetInnodbNumaInterleave sets InnodbNumaInterleave field to given value.

### HasInnodbNumaInterleave

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasInnodbNumaInterleave() bool`

HasInnodbNumaInterleave returns a boolean if a field has been set.

### GetNetReadTimeout

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetNetReadTimeout() string`

GetNetReadTimeout returns the NetReadTimeout field if non-nil, zero value otherwise.

### GetNetReadTimeoutOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetNetReadTimeoutOk() (*string, bool)`

GetNetReadTimeoutOk returns a tuple with the NetReadTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetReadTimeout

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetNetReadTimeout(v string)`

SetNetReadTimeout sets NetReadTimeout field to given value.

### HasNetReadTimeout

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasNetReadTimeout() bool`

HasNetReadTimeout returns a boolean if a field has been set.

### GetNetWriteTimeout

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetNetWriteTimeout() string`

GetNetWriteTimeout returns the NetWriteTimeout field if non-nil, zero value otherwise.

### GetNetWriteTimeoutOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetNetWriteTimeoutOk() (*string, bool)`

GetNetWriteTimeoutOk returns a tuple with the NetWriteTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetWriteTimeout

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetNetWriteTimeout(v string)`

SetNetWriteTimeout sets NetWriteTimeout field to given value.

### HasNetWriteTimeout

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasNetWriteTimeout() bool`

HasNetWriteTimeout returns a boolean if a field has been set.

### GetRegexpTimeLimit

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetRegexpTimeLimit() string`

GetRegexpTimeLimit returns the RegexpTimeLimit field if non-nil, zero value otherwise.

### GetRegexpTimeLimitOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetRegexpTimeLimitOk() (*string, bool)`

GetRegexpTimeLimitOk returns a tuple with the RegexpTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegexpTimeLimit

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetRegexpTimeLimit(v string)`

SetRegexpTimeLimit sets RegexpTimeLimit field to given value.

### HasRegexpTimeLimit

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasRegexpTimeLimit() bool`

HasRegexpTimeLimit returns a boolean if a field has been set.

### GetSyncBinlog

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetSyncBinlog() string`

GetSyncBinlog returns the SyncBinlog field if non-nil, zero value otherwise.

### GetSyncBinlogOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetSyncBinlogOk() (*string, bool)`

GetSyncBinlogOk returns a tuple with the SyncBinlog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncBinlog

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetSyncBinlog(v string)`

SetSyncBinlog sets SyncBinlog field to given value.

### HasSyncBinlog

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasSyncBinlog() bool`

HasSyncBinlog returns a boolean if a field has been set.

### GetTableDefinitionCache

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetTableDefinitionCache() string`

GetTableDefinitionCache returns the TableDefinitionCache field if non-nil, zero value otherwise.

### GetTableDefinitionCacheOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetTableDefinitionCacheOk() (*string, bool)`

GetTableDefinitionCacheOk returns a tuple with the TableDefinitionCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableDefinitionCache

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetTableDefinitionCache(v string)`

SetTableDefinitionCache sets TableDefinitionCache field to given value.

### HasTableDefinitionCache

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasTableDefinitionCache() bool`

HasTableDefinitionCache returns a boolean if a field has been set.

### GetLogBinTrustFunctionCreators

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetLogBinTrustFunctionCreators() string`

GetLogBinTrustFunctionCreators returns the LogBinTrustFunctionCreators field if non-nil, zero value otherwise.

### GetLogBinTrustFunctionCreatorsOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetLogBinTrustFunctionCreatorsOk() (*string, bool)`

GetLogBinTrustFunctionCreatorsOk returns a tuple with the LogBinTrustFunctionCreators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogBinTrustFunctionCreators

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetLogBinTrustFunctionCreators(v string)`

SetLogBinTrustFunctionCreators sets LogBinTrustFunctionCreators field to given value.

### HasLogBinTrustFunctionCreators

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasLogBinTrustFunctionCreators() bool`

HasLogBinTrustFunctionCreators returns a boolean if a field has been set.

### GetSkipNameResolve

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetSkipNameResolve() string`

GetSkipNameResolve returns the SkipNameResolve field if non-nil, zero value otherwise.

### GetSkipNameResolveOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetSkipNameResolveOk() (*string, bool)`

GetSkipNameResolveOk returns a tuple with the SkipNameResolve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipNameResolve

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetSkipNameResolve(v string)`

SetSkipNameResolve sets SkipNameResolve field to given value.

### HasSkipNameResolve

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasSkipNameResolve() bool`

HasSkipNameResolve returns a boolean if a field has been set.

### GetInnodbRedoLogCapacity

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetInnodbRedoLogCapacity() string`

GetInnodbRedoLogCapacity returns the InnodbRedoLogCapacity field if non-nil, zero value otherwise.

### GetInnodbRedoLogCapacityOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetInnodbRedoLogCapacityOk() (*string, bool)`

GetInnodbRedoLogCapacityOk returns a tuple with the InnodbRedoLogCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnodbRedoLogCapacity

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetInnodbRedoLogCapacity(v string)`

SetInnodbRedoLogCapacity sets InnodbRedoLogCapacity field to given value.

### HasInnodbRedoLogCapacity

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasInnodbRedoLogCapacity() bool`

HasInnodbRedoLogCapacity returns a boolean if a field has been set.

### GetWaitTimeout

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetWaitTimeout() string`

GetWaitTimeout returns the WaitTimeout field if non-nil, zero value otherwise.

### GetWaitTimeoutOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetWaitTimeoutOk() (*string, bool)`

GetWaitTimeoutOk returns a tuple with the WaitTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitTimeout

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetWaitTimeout(v string)`

SetWaitTimeout sets WaitTimeout field to given value.

### HasWaitTimeout

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasWaitTimeout() bool`

HasWaitTimeout returns a boolean if a field has been set.

### GetInteractiveTimeout

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetInteractiveTimeout() string`

GetInteractiveTimeout returns the InteractiveTimeout field if non-nil, zero value otherwise.

### GetInteractiveTimeoutOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetInteractiveTimeoutOk() (*string, bool)`

GetInteractiveTimeoutOk returns a tuple with the InteractiveTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractiveTimeout

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetInteractiveTimeout(v string)`

SetInteractiveTimeout sets InteractiveTimeout field to given value.

### HasInteractiveTimeout

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasInteractiveTimeout() bool`

HasInteractiveTimeout returns a boolean if a field has been set.

### GetDefaultTimeZone

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetDefaultTimeZone() string`

GetDefaultTimeZone returns the DefaultTimeZone field if non-nil, zero value otherwise.

### GetDefaultTimeZoneOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetDefaultTimeZoneOk() (*string, bool)`

GetDefaultTimeZoneOk returns a tuple with the DefaultTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTimeZone

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetDefaultTimeZone(v string)`

SetDefaultTimeZone sets DefaultTimeZone field to given value.

### HasDefaultTimeZone

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasDefaultTimeZone() bool`

HasDefaultTimeZone returns a boolean if a field has been set.

### GetPxcStrictMode

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetPxcStrictMode() string`

GetPxcStrictMode returns the PxcStrictMode field if non-nil, zero value otherwise.

### GetPxcStrictModeOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetPxcStrictModeOk() (*string, bool)`

GetPxcStrictModeOk returns a tuple with the PxcStrictMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPxcStrictMode

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetPxcStrictMode(v string)`

SetPxcStrictMode sets PxcStrictMode field to given value.

### HasPxcStrictMode

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasPxcStrictMode() bool`

HasPxcStrictMode returns a boolean if a field has been set.

### GetAutovacuumAnalyzeScaleFactor

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetAutovacuumAnalyzeScaleFactor() string`

GetAutovacuumAnalyzeScaleFactor returns the AutovacuumAnalyzeScaleFactor field if non-nil, zero value otherwise.

### GetAutovacuumAnalyzeScaleFactorOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetAutovacuumAnalyzeScaleFactorOk() (*string, bool)`

GetAutovacuumAnalyzeScaleFactorOk returns a tuple with the AutovacuumAnalyzeScaleFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutovacuumAnalyzeScaleFactor

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetAutovacuumAnalyzeScaleFactor(v string)`

SetAutovacuumAnalyzeScaleFactor sets AutovacuumAnalyzeScaleFactor field to given value.

### HasAutovacuumAnalyzeScaleFactor

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasAutovacuumAnalyzeScaleFactor() bool`

HasAutovacuumAnalyzeScaleFactor returns a boolean if a field has been set.

### GetAutovacuumMaxWorkers

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetAutovacuumMaxWorkers() string`

GetAutovacuumMaxWorkers returns the AutovacuumMaxWorkers field if non-nil, zero value otherwise.

### GetAutovacuumMaxWorkersOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetAutovacuumMaxWorkersOk() (*string, bool)`

GetAutovacuumMaxWorkersOk returns a tuple with the AutovacuumMaxWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutovacuumMaxWorkers

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetAutovacuumMaxWorkers(v string)`

SetAutovacuumMaxWorkers sets AutovacuumMaxWorkers field to given value.

### HasAutovacuumMaxWorkers

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasAutovacuumMaxWorkers() bool`

HasAutovacuumMaxWorkers returns a boolean if a field has been set.

### GetAutovacuumNaptime

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetAutovacuumNaptime() string`

GetAutovacuumNaptime returns the AutovacuumNaptime field if non-nil, zero value otherwise.

### GetAutovacuumNaptimeOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetAutovacuumNaptimeOk() (*string, bool)`

GetAutovacuumNaptimeOk returns a tuple with the AutovacuumNaptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutovacuumNaptime

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetAutovacuumNaptime(v string)`

SetAutovacuumNaptime sets AutovacuumNaptime field to given value.

### HasAutovacuumNaptime

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasAutovacuumNaptime() bool`

HasAutovacuumNaptime returns a boolean if a field has been set.

### GetAutovacuumVacuumInsertScaleFactor

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetAutovacuumVacuumInsertScaleFactor() string`

GetAutovacuumVacuumInsertScaleFactor returns the AutovacuumVacuumInsertScaleFactor field if non-nil, zero value otherwise.

### GetAutovacuumVacuumInsertScaleFactorOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetAutovacuumVacuumInsertScaleFactorOk() (*string, bool)`

GetAutovacuumVacuumInsertScaleFactorOk returns a tuple with the AutovacuumVacuumInsertScaleFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutovacuumVacuumInsertScaleFactor

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetAutovacuumVacuumInsertScaleFactor(v string)`

SetAutovacuumVacuumInsertScaleFactor sets AutovacuumVacuumInsertScaleFactor field to given value.

### HasAutovacuumVacuumInsertScaleFactor

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasAutovacuumVacuumInsertScaleFactor() bool`

HasAutovacuumVacuumInsertScaleFactor returns a boolean if a field has been set.

### GetAutovacuumVacuumScaleFactor

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetAutovacuumVacuumScaleFactor() string`

GetAutovacuumVacuumScaleFactor returns the AutovacuumVacuumScaleFactor field if non-nil, zero value otherwise.

### GetAutovacuumVacuumScaleFactorOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetAutovacuumVacuumScaleFactorOk() (*string, bool)`

GetAutovacuumVacuumScaleFactorOk returns a tuple with the AutovacuumVacuumScaleFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutovacuumVacuumScaleFactor

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetAutovacuumVacuumScaleFactor(v string)`

SetAutovacuumVacuumScaleFactor sets AutovacuumVacuumScaleFactor field to given value.

### HasAutovacuumVacuumScaleFactor

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasAutovacuumVacuumScaleFactor() bool`

HasAutovacuumVacuumScaleFactor returns a boolean if a field has been set.

### GetAutovacuumWorkMem

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetAutovacuumWorkMem() string`

GetAutovacuumWorkMem returns the AutovacuumWorkMem field if non-nil, zero value otherwise.

### GetAutovacuumWorkMemOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetAutovacuumWorkMemOk() (*string, bool)`

GetAutovacuumWorkMemOk returns a tuple with the AutovacuumWorkMem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutovacuumWorkMem

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetAutovacuumWorkMem(v string)`

SetAutovacuumWorkMem sets AutovacuumWorkMem field to given value.

### HasAutovacuumWorkMem

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasAutovacuumWorkMem() bool`

HasAutovacuumWorkMem returns a boolean if a field has been set.

### GetBgwriterDelay

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetBgwriterDelay() string`

GetBgwriterDelay returns the BgwriterDelay field if non-nil, zero value otherwise.

### GetBgwriterDelayOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetBgwriterDelayOk() (*string, bool)`

GetBgwriterDelayOk returns a tuple with the BgwriterDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgwriterDelay

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetBgwriterDelay(v string)`

SetBgwriterDelay sets BgwriterDelay field to given value.

### HasBgwriterDelay

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasBgwriterDelay() bool`

HasBgwriterDelay returns a boolean if a field has been set.

### GetBgwriterLruMaxpages

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetBgwriterLruMaxpages() string`

GetBgwriterLruMaxpages returns the BgwriterLruMaxpages field if non-nil, zero value otherwise.

### GetBgwriterLruMaxpagesOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetBgwriterLruMaxpagesOk() (*string, bool)`

GetBgwriterLruMaxpagesOk returns a tuple with the BgwriterLruMaxpages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgwriterLruMaxpages

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetBgwriterLruMaxpages(v string)`

SetBgwriterLruMaxpages sets BgwriterLruMaxpages field to given value.

### HasBgwriterLruMaxpages

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasBgwriterLruMaxpages() bool`

HasBgwriterLruMaxpages returns a boolean if a field has been set.

### GetDeadlockTimeout

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetDeadlockTimeout() string`

GetDeadlockTimeout returns the DeadlockTimeout field if non-nil, zero value otherwise.

### GetDeadlockTimeoutOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetDeadlockTimeoutOk() (*string, bool)`

GetDeadlockTimeoutOk returns a tuple with the DeadlockTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadlockTimeout

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetDeadlockTimeout(v string)`

SetDeadlockTimeout sets DeadlockTimeout field to given value.

### HasDeadlockTimeout

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasDeadlockTimeout() bool`

HasDeadlockTimeout returns a boolean if a field has been set.

### GetGinPendingListLimit

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetGinPendingListLimit() string`

GetGinPendingListLimit returns the GinPendingListLimit field if non-nil, zero value otherwise.

### GetGinPendingListLimitOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetGinPendingListLimitOk() (*string, bool)`

GetGinPendingListLimitOk returns a tuple with the GinPendingListLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGinPendingListLimit

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetGinPendingListLimit(v string)`

SetGinPendingListLimit sets GinPendingListLimit field to given value.

### HasGinPendingListLimit

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasGinPendingListLimit() bool`

HasGinPendingListLimit returns a boolean if a field has been set.

### GetIdleInTransactionSessionTimeout

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetIdleInTransactionSessionTimeout() string`

GetIdleInTransactionSessionTimeout returns the IdleInTransactionSessionTimeout field if non-nil, zero value otherwise.

### GetIdleInTransactionSessionTimeoutOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetIdleInTransactionSessionTimeoutOk() (*string, bool)`

GetIdleInTransactionSessionTimeoutOk returns a tuple with the IdleInTransactionSessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleInTransactionSessionTimeout

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetIdleInTransactionSessionTimeout(v string)`

SetIdleInTransactionSessionTimeout sets IdleInTransactionSessionTimeout field to given value.

### HasIdleInTransactionSessionTimeout

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasIdleInTransactionSessionTimeout() bool`

HasIdleInTransactionSessionTimeout returns a boolean if a field has been set.

### GetJoinCollapseLimit

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetJoinCollapseLimit() string`

GetJoinCollapseLimit returns the JoinCollapseLimit field if non-nil, zero value otherwise.

### GetJoinCollapseLimitOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetJoinCollapseLimitOk() (*string, bool)`

GetJoinCollapseLimitOk returns a tuple with the JoinCollapseLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinCollapseLimit

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetJoinCollapseLimit(v string)`

SetJoinCollapseLimit sets JoinCollapseLimit field to given value.

### HasJoinCollapseLimit

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasJoinCollapseLimit() bool`

HasJoinCollapseLimit returns a boolean if a field has been set.

### GetLockTimeout

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetLockTimeout() string`

GetLockTimeout returns the LockTimeout field if non-nil, zero value otherwise.

### GetLockTimeoutOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetLockTimeoutOk() (*string, bool)`

GetLockTimeoutOk returns a tuple with the LockTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockTimeout

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetLockTimeout(v string)`

SetLockTimeout sets LockTimeout field to given value.

### HasLockTimeout

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasLockTimeout() bool`

HasLockTimeout returns a boolean if a field has been set.

### GetMaxPreparedTransactions

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetMaxPreparedTransactions() string`

GetMaxPreparedTransactions returns the MaxPreparedTransactions field if non-nil, zero value otherwise.

### GetMaxPreparedTransactionsOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetMaxPreparedTransactionsOk() (*string, bool)`

GetMaxPreparedTransactionsOk returns a tuple with the MaxPreparedTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPreparedTransactions

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetMaxPreparedTransactions(v string)`

SetMaxPreparedTransactions sets MaxPreparedTransactions field to given value.

### HasMaxPreparedTransactions

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasMaxPreparedTransactions() bool`

HasMaxPreparedTransactions returns a boolean if a field has been set.

### GetSharedBuffers

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetSharedBuffers() string`

GetSharedBuffers returns the SharedBuffers field if non-nil, zero value otherwise.

### GetSharedBuffersOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetSharedBuffersOk() (*string, bool)`

GetSharedBuffersOk returns a tuple with the SharedBuffers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedBuffers

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetSharedBuffers(v string)`

SetSharedBuffers sets SharedBuffers field to given value.

### HasSharedBuffers

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasSharedBuffers() bool`

HasSharedBuffers returns a boolean if a field has been set.

### GetLogMinDurationStatement

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetLogMinDurationStatement() string`

GetLogMinDurationStatement returns the LogMinDurationStatement field if non-nil, zero value otherwise.

### GetLogMinDurationStatementOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetLogMinDurationStatementOk() (*string, bool)`

GetLogMinDurationStatementOk returns a tuple with the LogMinDurationStatement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogMinDurationStatement

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetLogMinDurationStatement(v string)`

SetLogMinDurationStatement sets LogMinDurationStatement field to given value.

### HasLogMinDurationStatement

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasLogMinDurationStatement() bool`

HasLogMinDurationStatement returns a boolean if a field has been set.

### GetWalBuffers

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetWalBuffers() string`

GetWalBuffers returns the WalBuffers field if non-nil, zero value otherwise.

### GetWalBuffersOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetWalBuffersOk() (*string, bool)`

GetWalBuffersOk returns a tuple with the WalBuffers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalBuffers

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetWalBuffers(v string)`

SetWalBuffers sets WalBuffers field to given value.

### HasWalBuffers

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasWalBuffers() bool`

HasWalBuffers returns a boolean if a field has been set.

### GetTempBuffers

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetTempBuffers() string`

GetTempBuffers returns the TempBuffers field if non-nil, zero value otherwise.

### GetTempBuffersOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetTempBuffersOk() (*string, bool)`

GetTempBuffersOk returns a tuple with the TempBuffers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempBuffers

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetTempBuffers(v string)`

SetTempBuffers sets TempBuffers field to given value.

### HasTempBuffers

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasTempBuffers() bool`

HasTempBuffers returns a boolean if a field has been set.

### GetWorkMem

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetWorkMem() string`

GetWorkMem returns the WorkMem field if non-nil, zero value otherwise.

### GetWorkMemOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetWorkMemOk() (*string, bool)`

GetWorkMemOk returns a tuple with the WorkMem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkMem

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetWorkMem(v string)`

SetWorkMem sets WorkMem field to given value.

### HasWorkMem

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasWorkMem() bool`

HasWorkMem returns a boolean if a field has been set.

### GetDefaultTransactionIsolation

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetDefaultTransactionIsolation() string`

GetDefaultTransactionIsolation returns the DefaultTransactionIsolation field if non-nil, zero value otherwise.

### GetDefaultTransactionIsolationOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetDefaultTransactionIsolationOk() (*string, bool)`

GetDefaultTransactionIsolationOk returns a tuple with the DefaultTransactionIsolation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTransactionIsolation

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetDefaultTransactionIsolation(v string)`

SetDefaultTransactionIsolation sets DefaultTransactionIsolation field to given value.

### HasDefaultTransactionIsolation

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasDefaultTransactionIsolation() bool`

HasDefaultTransactionIsolation returns a boolean if a field has been set.

### GetEffectiveCacheSize

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetEffectiveCacheSize() string`

GetEffectiveCacheSize returns the EffectiveCacheSize field if non-nil, zero value otherwise.

### GetEffectiveCacheSizeOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetEffectiveCacheSizeOk() (*string, bool)`

GetEffectiveCacheSizeOk returns a tuple with the EffectiveCacheSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveCacheSize

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetEffectiveCacheSize(v string)`

SetEffectiveCacheSize sets EffectiveCacheSize field to given value.

### HasEffectiveCacheSize

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasEffectiveCacheSize() bool`

HasEffectiveCacheSize returns a boolean if a field has been set.

### GetMaxWalSize

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetMaxWalSize() string`

GetMaxWalSize returns the MaxWalSize field if non-nil, zero value otherwise.

### GetMaxWalSizeOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetMaxWalSizeOk() (*string, bool)`

GetMaxWalSizeOk returns a tuple with the MaxWalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWalSize

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetMaxWalSize(v string)`

SetMaxWalSize sets MaxWalSize field to given value.

### HasMaxWalSize

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasMaxWalSize() bool`

HasMaxWalSize returns a boolean if a field has been set.

### GetMinWalSize

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetMinWalSize() string`

GetMinWalSize returns the MinWalSize field if non-nil, zero value otherwise.

### GetMinWalSizeOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetMinWalSizeOk() (*string, bool)`

GetMinWalSizeOk returns a tuple with the MinWalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinWalSize

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetMinWalSize(v string)`

SetMinWalSize sets MinWalSize field to given value.

### HasMinWalSize

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasMinWalSize() bool`

HasMinWalSize returns a boolean if a field has been set.

### GetWalLevel

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetWalLevel() string`

GetWalLevel returns the WalLevel field if non-nil, zero value otherwise.

### GetWalLevelOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetWalLevelOk() (*string, bool)`

GetWalLevelOk returns a tuple with the WalLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalLevel

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetWalLevel(v string)`

SetWalLevel sets WalLevel field to given value.

### HasWalLevel

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasWalLevel() bool`

HasWalLevel returns a boolean if a field has been set.

### GetMaxReplicationSlots

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetMaxReplicationSlots() string`

GetMaxReplicationSlots returns the MaxReplicationSlots field if non-nil, zero value otherwise.

### GetMaxReplicationSlotsOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetMaxReplicationSlotsOk() (*string, bool)`

GetMaxReplicationSlotsOk returns a tuple with the MaxReplicationSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReplicationSlots

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetMaxReplicationSlots(v string)`

SetMaxReplicationSlots sets MaxReplicationSlots field to given value.

### HasMaxReplicationSlots

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasMaxReplicationSlots() bool`

HasMaxReplicationSlots returns a boolean if a field has been set.

### GetMaxWalSenders

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetMaxWalSenders() string`

GetMaxWalSenders returns the MaxWalSenders field if non-nil, zero value otherwise.

### GetMaxWalSendersOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetMaxWalSendersOk() (*string, bool)`

GetMaxWalSendersOk returns a tuple with the MaxWalSenders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWalSenders

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetMaxWalSenders(v string)`

SetMaxWalSenders sets MaxWalSenders field to given value.

### HasMaxWalSenders

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasMaxWalSenders() bool`

HasMaxWalSenders returns a boolean if a field has been set.

### GetMaxWorkerProcesses

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetMaxWorkerProcesses() string`

GetMaxWorkerProcesses returns the MaxWorkerProcesses field if non-nil, zero value otherwise.

### GetMaxWorkerProcessesOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetMaxWorkerProcessesOk() (*string, bool)`

GetMaxWorkerProcessesOk returns a tuple with the MaxWorkerProcesses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWorkerProcesses

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetMaxWorkerProcesses(v string)`

SetMaxWorkerProcesses sets MaxWorkerProcesses field to given value.

### HasMaxWorkerProcesses

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasMaxWorkerProcesses() bool`

HasMaxWorkerProcesses returns a boolean if a field has been set.

### GetMaxLogicalReplicationWorkers

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetMaxLogicalReplicationWorkers() string`

GetMaxLogicalReplicationWorkers returns the MaxLogicalReplicationWorkers field if non-nil, zero value otherwise.

### GetMaxLogicalReplicationWorkersOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetMaxLogicalReplicationWorkersOk() (*string, bool)`

GetMaxLogicalReplicationWorkersOk returns a tuple with the MaxLogicalReplicationWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLogicalReplicationWorkers

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetMaxLogicalReplicationWorkers(v string)`

SetMaxLogicalReplicationWorkers sets MaxLogicalReplicationWorkers field to given value.

### HasMaxLogicalReplicationWorkers

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasMaxLogicalReplicationWorkers() bool`

HasMaxLogicalReplicationWorkers returns a boolean if a field has been set.

### GetMaxParallelMaintenanceWorkers

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetMaxParallelMaintenanceWorkers() string`

GetMaxParallelMaintenanceWorkers returns the MaxParallelMaintenanceWorkers field if non-nil, zero value otherwise.

### GetMaxParallelMaintenanceWorkersOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetMaxParallelMaintenanceWorkersOk() (*string, bool)`

GetMaxParallelMaintenanceWorkersOk returns a tuple with the MaxParallelMaintenanceWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxParallelMaintenanceWorkers

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetMaxParallelMaintenanceWorkers(v string)`

SetMaxParallelMaintenanceWorkers sets MaxParallelMaintenanceWorkers field to given value.

### HasMaxParallelMaintenanceWorkers

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasMaxParallelMaintenanceWorkers() bool`

HasMaxParallelMaintenanceWorkers returns a boolean if a field has been set.

### GetMaxParallelWorkers

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetMaxParallelWorkers() string`

GetMaxParallelWorkers returns the MaxParallelWorkers field if non-nil, zero value otherwise.

### GetMaxParallelWorkersOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetMaxParallelWorkersOk() (*string, bool)`

GetMaxParallelWorkersOk returns a tuple with the MaxParallelWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxParallelWorkers

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetMaxParallelWorkers(v string)`

SetMaxParallelWorkers sets MaxParallelWorkers field to given value.

### HasMaxParallelWorkers

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasMaxParallelWorkers() bool`

HasMaxParallelWorkers returns a boolean if a field has been set.

### GetMaxParallelWorkersPerGather

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetMaxParallelWorkersPerGather() string`

GetMaxParallelWorkersPerGather returns the MaxParallelWorkersPerGather field if non-nil, zero value otherwise.

### GetMaxParallelWorkersPerGatherOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetMaxParallelWorkersPerGatherOk() (*string, bool)`

GetMaxParallelWorkersPerGatherOk returns a tuple with the MaxParallelWorkersPerGather field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxParallelWorkersPerGather

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetMaxParallelWorkersPerGather(v string)`

SetMaxParallelWorkersPerGather sets MaxParallelWorkersPerGather field to given value.

### HasMaxParallelWorkersPerGather

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasMaxParallelWorkersPerGather() bool`

HasMaxParallelWorkersPerGather returns a boolean if a field has been set.

### GetArrayNulls

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetArrayNulls() string`

GetArrayNulls returns the ArrayNulls field if non-nil, zero value otherwise.

### GetArrayNullsOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetArrayNullsOk() (*string, bool)`

GetArrayNullsOk returns a tuple with the ArrayNulls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayNulls

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetArrayNulls(v string)`

SetArrayNulls sets ArrayNulls field to given value.

### HasArrayNulls

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasArrayNulls() bool`

HasArrayNulls returns a boolean if a field has been set.

### GetBackendFlushAfter

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetBackendFlushAfter() string`

GetBackendFlushAfter returns the BackendFlushAfter field if non-nil, zero value otherwise.

### GetBackendFlushAfterOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetBackendFlushAfterOk() (*string, bool)`

GetBackendFlushAfterOk returns a tuple with the BackendFlushAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendFlushAfter

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetBackendFlushAfter(v string)`

SetBackendFlushAfter sets BackendFlushAfter field to given value.

### HasBackendFlushAfter

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasBackendFlushAfter() bool`

HasBackendFlushAfter returns a boolean if a field has been set.

### GetBackslashQuote

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetBackslashQuote() string`

GetBackslashQuote returns the BackslashQuote field if non-nil, zero value otherwise.

### GetBackslashQuoteOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetBackslashQuoteOk() (*string, bool)`

GetBackslashQuoteOk returns a tuple with the BackslashQuote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackslashQuote

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetBackslashQuote(v string)`

SetBackslashQuote sets BackslashQuote field to given value.

### HasBackslashQuote

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasBackslashQuote() bool`

HasBackslashQuote returns a boolean if a field has been set.

### GetBgwriterFlushAfter

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetBgwriterFlushAfter() string`

GetBgwriterFlushAfter returns the BgwriterFlushAfter field if non-nil, zero value otherwise.

### GetBgwriterFlushAfterOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetBgwriterFlushAfterOk() (*string, bool)`

GetBgwriterFlushAfterOk returns a tuple with the BgwriterFlushAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgwriterFlushAfter

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetBgwriterFlushAfter(v string)`

SetBgwriterFlushAfter sets BgwriterFlushAfter field to given value.

### HasBgwriterFlushAfter

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasBgwriterFlushAfter() bool`

HasBgwriterFlushAfter returns a boolean if a field has been set.

### GetBgwriterLruMultiplier

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetBgwriterLruMultiplier() string`

GetBgwriterLruMultiplier returns the BgwriterLruMultiplier field if non-nil, zero value otherwise.

### GetBgwriterLruMultiplierOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetBgwriterLruMultiplierOk() (*string, bool)`

GetBgwriterLruMultiplierOk returns a tuple with the BgwriterLruMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgwriterLruMultiplier

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetBgwriterLruMultiplier(v string)`

SetBgwriterLruMultiplier sets BgwriterLruMultiplier field to given value.

### HasBgwriterLruMultiplier

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasBgwriterLruMultiplier() bool`

HasBgwriterLruMultiplier returns a boolean if a field has been set.

### GetDefaultTransactionReadOnly

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetDefaultTransactionReadOnly() string`

GetDefaultTransactionReadOnly returns the DefaultTransactionReadOnly field if non-nil, zero value otherwise.

### GetDefaultTransactionReadOnlyOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetDefaultTransactionReadOnlyOk() (*string, bool)`

GetDefaultTransactionReadOnlyOk returns a tuple with the DefaultTransactionReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTransactionReadOnly

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetDefaultTransactionReadOnly(v string)`

SetDefaultTransactionReadOnly sets DefaultTransactionReadOnly field to given value.

### HasDefaultTransactionReadOnly

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasDefaultTransactionReadOnly() bool`

HasDefaultTransactionReadOnly returns a boolean if a field has been set.

### GetEnableHashagg

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetEnableHashagg() string`

GetEnableHashagg returns the EnableHashagg field if non-nil, zero value otherwise.

### GetEnableHashaggOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetEnableHashaggOk() (*string, bool)`

GetEnableHashaggOk returns a tuple with the EnableHashagg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableHashagg

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetEnableHashagg(v string)`

SetEnableHashagg sets EnableHashagg field to given value.

### HasEnableHashagg

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasEnableHashagg() bool`

HasEnableHashagg returns a boolean if a field has been set.

### GetEnableHashjoin

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetEnableHashjoin() string`

GetEnableHashjoin returns the EnableHashjoin field if non-nil, zero value otherwise.

### GetEnableHashjoinOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetEnableHashjoinOk() (*string, bool)`

GetEnableHashjoinOk returns a tuple with the EnableHashjoin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableHashjoin

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetEnableHashjoin(v string)`

SetEnableHashjoin sets EnableHashjoin field to given value.

### HasEnableHashjoin

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasEnableHashjoin() bool`

HasEnableHashjoin returns a boolean if a field has been set.

### GetEnableIncrementalSort

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetEnableIncrementalSort() string`

GetEnableIncrementalSort returns the EnableIncrementalSort field if non-nil, zero value otherwise.

### GetEnableIncrementalSortOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetEnableIncrementalSortOk() (*string, bool)`

GetEnableIncrementalSortOk returns a tuple with the EnableIncrementalSort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIncrementalSort

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetEnableIncrementalSort(v string)`

SetEnableIncrementalSort sets EnableIncrementalSort field to given value.

### HasEnableIncrementalSort

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasEnableIncrementalSort() bool`

HasEnableIncrementalSort returns a boolean if a field has been set.

### GetEnableIndexscan

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetEnableIndexscan() string`

GetEnableIndexscan returns the EnableIndexscan field if non-nil, zero value otherwise.

### GetEnableIndexscanOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetEnableIndexscanOk() (*string, bool)`

GetEnableIndexscanOk returns a tuple with the EnableIndexscan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIndexscan

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetEnableIndexscan(v string)`

SetEnableIndexscan sets EnableIndexscan field to given value.

### HasEnableIndexscan

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasEnableIndexscan() bool`

HasEnableIndexscan returns a boolean if a field has been set.

### GetEnableIndexonlyscan

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetEnableIndexonlyscan() string`

GetEnableIndexonlyscan returns the EnableIndexonlyscan field if non-nil, zero value otherwise.

### GetEnableIndexonlyscanOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetEnableIndexonlyscanOk() (*string, bool)`

GetEnableIndexonlyscanOk returns a tuple with the EnableIndexonlyscan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIndexonlyscan

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetEnableIndexonlyscan(v string)`

SetEnableIndexonlyscan sets EnableIndexonlyscan field to given value.

### HasEnableIndexonlyscan

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasEnableIndexonlyscan() bool`

HasEnableIndexonlyscan returns a boolean if a field has been set.

### GetEnableMaterial

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetEnableMaterial() string`

GetEnableMaterial returns the EnableMaterial field if non-nil, zero value otherwise.

### GetEnableMaterialOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetEnableMaterialOk() (*string, bool)`

GetEnableMaterialOk returns a tuple with the EnableMaterial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMaterial

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetEnableMaterial(v string)`

SetEnableMaterial sets EnableMaterial field to given value.

### HasEnableMaterial

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasEnableMaterial() bool`

HasEnableMaterial returns a boolean if a field has been set.

### GetEnableMemoize

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetEnableMemoize() string`

GetEnableMemoize returns the EnableMemoize field if non-nil, zero value otherwise.

### GetEnableMemoizeOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetEnableMemoizeOk() (*string, bool)`

GetEnableMemoizeOk returns a tuple with the EnableMemoize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMemoize

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetEnableMemoize(v string)`

SetEnableMemoize sets EnableMemoize field to given value.

### HasEnableMemoize

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasEnableMemoize() bool`

HasEnableMemoize returns a boolean if a field has been set.

### GetEnableMergejoin

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetEnableMergejoin() string`

GetEnableMergejoin returns the EnableMergejoin field if non-nil, zero value otherwise.

### GetEnableMergejoinOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetEnableMergejoinOk() (*string, bool)`

GetEnableMergejoinOk returns a tuple with the EnableMergejoin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMergejoin

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetEnableMergejoin(v string)`

SetEnableMergejoin sets EnableMergejoin field to given value.

### HasEnableMergejoin

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasEnableMergejoin() bool`

HasEnableMergejoin returns a boolean if a field has been set.

### GetEnableParallelAppend

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetEnableParallelAppend() string`

GetEnableParallelAppend returns the EnableParallelAppend field if non-nil, zero value otherwise.

### GetEnableParallelAppendOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetEnableParallelAppendOk() (*string, bool)`

GetEnableParallelAppendOk returns a tuple with the EnableParallelAppend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableParallelAppend

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetEnableParallelAppend(v string)`

SetEnableParallelAppend sets EnableParallelAppend field to given value.

### HasEnableParallelAppend

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasEnableParallelAppend() bool`

HasEnableParallelAppend returns a boolean if a field has been set.

### GetEnableParallelHash

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetEnableParallelHash() string`

GetEnableParallelHash returns the EnableParallelHash field if non-nil, zero value otherwise.

### GetEnableParallelHashOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetEnableParallelHashOk() (*string, bool)`

GetEnableParallelHashOk returns a tuple with the EnableParallelHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableParallelHash

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetEnableParallelHash(v string)`

SetEnableParallelHash sets EnableParallelHash field to given value.

### HasEnableParallelHash

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasEnableParallelHash() bool`

HasEnableParallelHash returns a boolean if a field has been set.

### GetEnablePartitionPruning

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetEnablePartitionPruning() string`

GetEnablePartitionPruning returns the EnablePartitionPruning field if non-nil, zero value otherwise.

### GetEnablePartitionPruningOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetEnablePartitionPruningOk() (*string, bool)`

GetEnablePartitionPruningOk returns a tuple with the EnablePartitionPruning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePartitionPruning

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetEnablePartitionPruning(v string)`

SetEnablePartitionPruning sets EnablePartitionPruning field to given value.

### HasEnablePartitionPruning

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasEnablePartitionPruning() bool`

HasEnablePartitionPruning returns a boolean if a field has been set.

### GetEnablePartitionwiseJoin

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetEnablePartitionwiseJoin() string`

GetEnablePartitionwiseJoin returns the EnablePartitionwiseJoin field if non-nil, zero value otherwise.

### GetEnablePartitionwiseJoinOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetEnablePartitionwiseJoinOk() (*string, bool)`

GetEnablePartitionwiseJoinOk returns a tuple with the EnablePartitionwiseJoin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePartitionwiseJoin

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetEnablePartitionwiseJoin(v string)`

SetEnablePartitionwiseJoin sets EnablePartitionwiseJoin field to given value.

### HasEnablePartitionwiseJoin

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasEnablePartitionwiseJoin() bool`

HasEnablePartitionwiseJoin returns a boolean if a field has been set.

### GetEnablePartitionwiseAggregate

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetEnablePartitionwiseAggregate() string`

GetEnablePartitionwiseAggregate returns the EnablePartitionwiseAggregate field if non-nil, zero value otherwise.

### GetEnablePartitionwiseAggregateOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetEnablePartitionwiseAggregateOk() (*string, bool)`

GetEnablePartitionwiseAggregateOk returns a tuple with the EnablePartitionwiseAggregate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePartitionwiseAggregate

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetEnablePartitionwiseAggregate(v string)`

SetEnablePartitionwiseAggregate sets EnablePartitionwiseAggregate field to given value.

### HasEnablePartitionwiseAggregate

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasEnablePartitionwiseAggregate() bool`

HasEnablePartitionwiseAggregate returns a boolean if a field has been set.

### GetEnableSeqscan

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetEnableSeqscan() string`

GetEnableSeqscan returns the EnableSeqscan field if non-nil, zero value otherwise.

### GetEnableSeqscanOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetEnableSeqscanOk() (*string, bool)`

GetEnableSeqscanOk returns a tuple with the EnableSeqscan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSeqscan

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetEnableSeqscan(v string)`

SetEnableSeqscan sets EnableSeqscan field to given value.

### HasEnableSeqscan

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasEnableSeqscan() bool`

HasEnableSeqscan returns a boolean if a field has been set.

### GetEnableSort

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetEnableSort() string`

GetEnableSort returns the EnableSort field if non-nil, zero value otherwise.

### GetEnableSortOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetEnableSortOk() (*string, bool)`

GetEnableSortOk returns a tuple with the EnableSort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSort

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetEnableSort(v string)`

SetEnableSort sets EnableSort field to given value.

### HasEnableSort

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasEnableSort() bool`

HasEnableSort returns a boolean if a field has been set.

### GetEnableTidscan

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetEnableTidscan() string`

GetEnableTidscan returns the EnableTidscan field if non-nil, zero value otherwise.

### GetEnableTidscanOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetEnableTidscanOk() (*string, bool)`

GetEnableTidscanOk returns a tuple with the EnableTidscan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTidscan

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetEnableTidscan(v string)`

SetEnableTidscan sets EnableTidscan field to given value.

### HasEnableTidscan

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasEnableTidscan() bool`

HasEnableTidscan returns a boolean if a field has been set.

### GetExitOnError

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetExitOnError() string`

GetExitOnError returns the ExitOnError field if non-nil, zero value otherwise.

### GetExitOnErrorOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetExitOnErrorOk() (*string, bool)`

GetExitOnErrorOk returns a tuple with the ExitOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitOnError

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetExitOnError(v string)`

SetExitOnError sets ExitOnError field to given value.

### HasExitOnError

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasExitOnError() bool`

HasExitOnError returns a boolean if a field has been set.

### GetFromCollapseLimit

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetFromCollapseLimit() string`

GetFromCollapseLimit returns the FromCollapseLimit field if non-nil, zero value otherwise.

### GetFromCollapseLimitOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetFromCollapseLimitOk() (*string, bool)`

GetFromCollapseLimitOk returns a tuple with the FromCollapseLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromCollapseLimit

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetFromCollapseLimit(v string)`

SetFromCollapseLimit sets FromCollapseLimit field to given value.

### HasFromCollapseLimit

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasFromCollapseLimit() bool`

HasFromCollapseLimit returns a boolean if a field has been set.

### GetJit

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetJit() string`

GetJit returns the Jit field if non-nil, zero value otherwise.

### GetJitOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetJitOk() (*string, bool)`

GetJitOk returns a tuple with the Jit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJit

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetJit(v string)`

SetJit sets Jit field to given value.

### HasJit

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasJit() bool`

HasJit returns a boolean if a field has been set.

### GetPlanCacheMode

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetPlanCacheMode() string`

GetPlanCacheMode returns the PlanCacheMode field if non-nil, zero value otherwise.

### GetPlanCacheModeOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetPlanCacheModeOk() (*string, bool)`

GetPlanCacheModeOk returns a tuple with the PlanCacheMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanCacheMode

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetPlanCacheMode(v string)`

SetPlanCacheMode sets PlanCacheMode field to given value.

### HasPlanCacheMode

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasPlanCacheMode() bool`

HasPlanCacheMode returns a boolean if a field has been set.

### GetQuoteAllIdentifiers

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetQuoteAllIdentifiers() string`

GetQuoteAllIdentifiers returns the QuoteAllIdentifiers field if non-nil, zero value otherwise.

### GetQuoteAllIdentifiersOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetQuoteAllIdentifiersOk() (*string, bool)`

GetQuoteAllIdentifiersOk returns a tuple with the QuoteAllIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAllIdentifiers

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetQuoteAllIdentifiers(v string)`

SetQuoteAllIdentifiers sets QuoteAllIdentifiers field to given value.

### HasQuoteAllIdentifiers

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasQuoteAllIdentifiers() bool`

HasQuoteAllIdentifiers returns a boolean if a field has been set.

### GetStandardConformingStrings

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetStandardConformingStrings() string`

GetStandardConformingStrings returns the StandardConformingStrings field if non-nil, zero value otherwise.

### GetStandardConformingStringsOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetStandardConformingStringsOk() (*string, bool)`

GetStandardConformingStringsOk returns a tuple with the StandardConformingStrings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardConformingStrings

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetStandardConformingStrings(v string)`

SetStandardConformingStrings sets StandardConformingStrings field to given value.

### HasStandardConformingStrings

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasStandardConformingStrings() bool`

HasStandardConformingStrings returns a boolean if a field has been set.

### GetStatementTimeout

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetStatementTimeout() string`

GetStatementTimeout returns the StatementTimeout field if non-nil, zero value otherwise.

### GetStatementTimeoutOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetStatementTimeoutOk() (*string, bool)`

GetStatementTimeoutOk returns a tuple with the StatementTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementTimeout

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetStatementTimeout(v string)`

SetStatementTimeout sets StatementTimeout field to given value.

### HasStatementTimeout

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasStatementTimeout() bool`

HasStatementTimeout returns a boolean if a field has been set.

### GetTimezone

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetTransformNullEquals

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetTransformNullEquals() string`

GetTransformNullEquals returns the TransformNullEquals field if non-nil, zero value otherwise.

### GetTransformNullEqualsOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetTransformNullEqualsOk() (*string, bool)`

GetTransformNullEqualsOk returns a tuple with the TransformNullEquals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformNullEquals

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetTransformNullEquals(v string)`

SetTransformNullEquals sets TransformNullEquals field to given value.

### HasTransformNullEquals

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasTransformNullEquals() bool`

HasTransformNullEquals returns a boolean if a field has been set.

### GetMaxLocksPerTransaction

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetMaxLocksPerTransaction() string`

GetMaxLocksPerTransaction returns the MaxLocksPerTransaction field if non-nil, zero value otherwise.

### GetMaxLocksPerTransactionOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetMaxLocksPerTransactionOk() (*string, bool)`

GetMaxLocksPerTransactionOk returns a tuple with the MaxLocksPerTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLocksPerTransaction

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetMaxLocksPerTransaction(v string)`

SetMaxLocksPerTransaction sets MaxLocksPerTransaction field to given value.

### HasMaxLocksPerTransaction

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasMaxLocksPerTransaction() bool`

HasMaxLocksPerTransaction returns a boolean if a field has been set.

### GetAutovacuumVacuumCostLimit

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetAutovacuumVacuumCostLimit() string`

GetAutovacuumVacuumCostLimit returns the AutovacuumVacuumCostLimit field if non-nil, zero value otherwise.

### GetAutovacuumVacuumCostLimitOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetAutovacuumVacuumCostLimitOk() (*string, bool)`

GetAutovacuumVacuumCostLimitOk returns a tuple with the AutovacuumVacuumCostLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutovacuumVacuumCostLimit

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetAutovacuumVacuumCostLimit(v string)`

SetAutovacuumVacuumCostLimit sets AutovacuumVacuumCostLimit field to given value.

### HasAutovacuumVacuumCostLimit

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasAutovacuumVacuumCostLimit() bool`

HasAutovacuumVacuumCostLimit returns a boolean if a field has been set.

### GetCheckpointTimeout

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetCheckpointTimeout() string`

GetCheckpointTimeout returns the CheckpointTimeout field if non-nil, zero value otherwise.

### GetCheckpointTimeoutOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetCheckpointTimeoutOk() (*string, bool)`

GetCheckpointTimeoutOk returns a tuple with the CheckpointTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckpointTimeout

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetCheckpointTimeout(v string)`

SetCheckpointTimeout sets CheckpointTimeout field to given value.

### HasCheckpointTimeout

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasCheckpointTimeout() bool`

HasCheckpointTimeout returns a boolean if a field has been set.

### GetCheckpointCompletionTarget

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetCheckpointCompletionTarget() string`

GetCheckpointCompletionTarget returns the CheckpointCompletionTarget field if non-nil, zero value otherwise.

### GetCheckpointCompletionTargetOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetCheckpointCompletionTargetOk() (*string, bool)`

GetCheckpointCompletionTargetOk returns a tuple with the CheckpointCompletionTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckpointCompletionTarget

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetCheckpointCompletionTarget(v string)`

SetCheckpointCompletionTarget sets CheckpointCompletionTarget field to given value.

### HasCheckpointCompletionTarget

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasCheckpointCompletionTarget() bool`

HasCheckpointCompletionTarget returns a boolean if a field has been set.

### GetWalCompression

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetWalCompression() string`

GetWalCompression returns the WalCompression field if non-nil, zero value otherwise.

### GetWalCompressionOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetWalCompressionOk() (*string, bool)`

GetWalCompressionOk returns a tuple with the WalCompression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalCompression

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetWalCompression(v string)`

SetWalCompression sets WalCompression field to given value.

### HasWalCompression

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasWalCompression() bool`

HasWalCompression returns a boolean if a field has been set.

### GetRandomPageCost

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetRandomPageCost() string`

GetRandomPageCost returns the RandomPageCost field if non-nil, zero value otherwise.

### GetRandomPageCostOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetRandomPageCostOk() (*string, bool)`

GetRandomPageCostOk returns a tuple with the RandomPageCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomPageCost

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetRandomPageCost(v string)`

SetRandomPageCost sets RandomPageCost field to given value.

### HasRandomPageCost

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasRandomPageCost() bool`

HasRandomPageCost returns a boolean if a field has been set.

### GetEffectiveIoConcurrency

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetEffectiveIoConcurrency() string`

GetEffectiveIoConcurrency returns the EffectiveIoConcurrency field if non-nil, zero value otherwise.

### GetEffectiveIoConcurrencyOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetEffectiveIoConcurrencyOk() (*string, bool)`

GetEffectiveIoConcurrencyOk returns a tuple with the EffectiveIoConcurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveIoConcurrency

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetEffectiveIoConcurrency(v string)`

SetEffectiveIoConcurrency sets EffectiveIoConcurrency field to given value.

### HasEffectiveIoConcurrency

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasEffectiveIoConcurrency() bool`

HasEffectiveIoConcurrency returns a boolean if a field has been set.

### GetLogLockWaits

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetLogLockWaits() string`

GetLogLockWaits returns the LogLockWaits field if non-nil, zero value otherwise.

### GetLogLockWaitsOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetLogLockWaitsOk() (*string, bool)`

GetLogLockWaitsOk returns a tuple with the LogLockWaits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLockWaits

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetLogLockWaits(v string)`

SetLogLockWaits sets LogLockWaits field to given value.

### HasLogLockWaits

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasLogLockWaits() bool`

HasLogLockWaits returns a boolean if a field has been set.

### GetLogTempFiles

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetLogTempFiles() string`

GetLogTempFiles returns the LogTempFiles field if non-nil, zero value otherwise.

### GetLogTempFilesOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetLogTempFilesOk() (*string, bool)`

GetLogTempFilesOk returns a tuple with the LogTempFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogTempFiles

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetLogTempFiles(v string)`

SetLogTempFiles sets LogTempFiles field to given value.

### HasLogTempFiles

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasLogTempFiles() bool`

HasLogTempFiles returns a boolean if a field has been set.

### GetTrackIoTiming

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetTrackIoTiming() string`

GetTrackIoTiming returns the TrackIoTiming field if non-nil, zero value otherwise.

### GetTrackIoTimingOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetTrackIoTimingOk() (*string, bool)`

GetTrackIoTimingOk returns a tuple with the TrackIoTiming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackIoTiming

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetTrackIoTiming(v string)`

SetTrackIoTiming sets TrackIoTiming field to given value.

### HasTrackIoTiming

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasTrackIoTiming() bool`

HasTrackIoTiming returns a boolean if a field has been set.

### GetMaintenanceWorkMem

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetMaintenanceWorkMem() string`

GetMaintenanceWorkMem returns the MaintenanceWorkMem field if non-nil, zero value otherwise.

### GetMaintenanceWorkMemOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetMaintenanceWorkMemOk() (*string, bool)`

GetMaintenanceWorkMemOk returns a tuple with the MaintenanceWorkMem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceWorkMem

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetMaintenanceWorkMem(v string)`

SetMaintenanceWorkMem sets MaintenanceWorkMem field to given value.

### HasMaintenanceWorkMem

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasMaintenanceWorkMem() bool`

HasMaintenanceWorkMem returns a boolean if a field has been set.

### GetIdleSessionTimeout

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetIdleSessionTimeout() string`

GetIdleSessionTimeout returns the IdleSessionTimeout field if non-nil, zero value otherwise.

### GetIdleSessionTimeoutOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetIdleSessionTimeoutOk() (*string, bool)`

GetIdleSessionTimeoutOk returns a tuple with the IdleSessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleSessionTimeout

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetIdleSessionTimeout(v string)`

SetIdleSessionTimeout sets IdleSessionTimeout field to given value.

### HasIdleSessionTimeout

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasIdleSessionTimeout() bool`

HasIdleSessionTimeout returns a boolean if a field has been set.

### GetIoMethod

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetIoMethod() string`

GetIoMethod returns the IoMethod field if non-nil, zero value otherwise.

### GetIoMethodOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetIoMethodOk() (*string, bool)`

GetIoMethodOk returns a tuple with the IoMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoMethod

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetIoMethod(v string)`

SetIoMethod sets IoMethod field to given value.

### HasIoMethod

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasIoMethod() bool`

HasIoMethod returns a boolean if a field has been set.

### GetIoWorkers

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetIoWorkers() string`

GetIoWorkers returns the IoWorkers field if non-nil, zero value otherwise.

### GetIoWorkersOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetIoWorkersOk() (*string, bool)`

GetIoWorkersOk returns a tuple with the IoWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoWorkers

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetIoWorkers(v string)`

SetIoWorkers sets IoWorkers field to given value.

### HasIoWorkers

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasIoWorkers() bool`

HasIoWorkers returns a boolean if a field has been set.

### GetClientOutputBufferLimitNormal

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetClientOutputBufferLimitNormal() string`

GetClientOutputBufferLimitNormal returns the ClientOutputBufferLimitNormal field if non-nil, zero value otherwise.

### GetClientOutputBufferLimitNormalOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetClientOutputBufferLimitNormalOk() (*string, bool)`

GetClientOutputBufferLimitNormalOk returns a tuple with the ClientOutputBufferLimitNormal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOutputBufferLimitNormal

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetClientOutputBufferLimitNormal(v string)`

SetClientOutputBufferLimitNormal sets ClientOutputBufferLimitNormal field to given value.

### HasClientOutputBufferLimitNormal

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasClientOutputBufferLimitNormal() bool`

HasClientOutputBufferLimitNormal returns a boolean if a field has been set.

### GetNotifyKeyspaceEvents

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetNotifyKeyspaceEvents() string`

GetNotifyKeyspaceEvents returns the NotifyKeyspaceEvents field if non-nil, zero value otherwise.

### GetNotifyKeyspaceEventsOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetNotifyKeyspaceEventsOk() (*string, bool)`

GetNotifyKeyspaceEventsOk returns a tuple with the NotifyKeyspaceEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyKeyspaceEvents

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetNotifyKeyspaceEvents(v string)`

SetNotifyKeyspaceEvents sets NotifyKeyspaceEvents field to given value.

### HasNotifyKeyspaceEvents

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasNotifyKeyspaceEvents() bool`

HasNotifyKeyspaceEvents returns a boolean if a field has been set.

### GetClientOutputBufferLimitPubsub

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetClientOutputBufferLimitPubsub() string`

GetClientOutputBufferLimitPubsub returns the ClientOutputBufferLimitPubsub field if non-nil, zero value otherwise.

### GetClientOutputBufferLimitPubsubOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetClientOutputBufferLimitPubsubOk() (*string, bool)`

GetClientOutputBufferLimitPubsubOk returns a tuple with the ClientOutputBufferLimitPubsub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOutputBufferLimitPubsub

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetClientOutputBufferLimitPubsub(v string)`

SetClientOutputBufferLimitPubsub sets ClientOutputBufferLimitPubsub field to given value.

### HasClientOutputBufferLimitPubsub

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasClientOutputBufferLimitPubsub() bool`

HasClientOutputBufferLimitPubsub returns a boolean if a field has been set.

### GetDatabases

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetDatabases() string`

GetDatabases returns the Databases field if non-nil, zero value otherwise.

### GetDatabasesOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetDatabasesOk() (*string, bool)`

GetDatabasesOk returns a tuple with the Databases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabases

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetDatabases(v string)`

SetDatabases sets Databases field to given value.

### HasDatabases

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasDatabases() bool`

HasDatabases returns a boolean if a field has been set.

### GetTimeout

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetMaxmemoryPolicy

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetMaxmemoryPolicy() string`

GetMaxmemoryPolicy returns the MaxmemoryPolicy field if non-nil, zero value otherwise.

### GetMaxmemoryPolicyOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetMaxmemoryPolicyOk() (*string, bool)`

GetMaxmemoryPolicyOk returns a tuple with the MaxmemoryPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxmemoryPolicy

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetMaxmemoryPolicy(v string)`

SetMaxmemoryPolicy sets MaxmemoryPolicy field to given value.

### HasMaxmemoryPolicy

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasMaxmemoryPolicy() bool`

HasMaxmemoryPolicy returns a boolean if a field has been set.

### GetSlowlogLogSlowerThan

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetSlowlogLogSlowerThan() string`

GetSlowlogLogSlowerThan returns the SlowlogLogSlowerThan field if non-nil, zero value otherwise.

### GetSlowlogLogSlowerThanOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetSlowlogLogSlowerThanOk() (*string, bool)`

GetSlowlogLogSlowerThanOk returns a tuple with the SlowlogLogSlowerThan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlowlogLogSlowerThan

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetSlowlogLogSlowerThan(v string)`

SetSlowlogLogSlowerThan sets SlowlogLogSlowerThan field to given value.

### HasSlowlogLogSlowerThan

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasSlowlogLogSlowerThan() bool`

HasSlowlogLogSlowerThan returns a boolean if a field has been set.

### GetSlowlogMaxLen

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetSlowlogMaxLen() string`

GetSlowlogMaxLen returns the SlowlogMaxLen field if non-nil, zero value otherwise.

### GetSlowlogMaxLenOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetSlowlogMaxLenOk() (*string, bool)`

GetSlowlogMaxLenOk returns a tuple with the SlowlogMaxLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlowlogMaxLen

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetSlowlogMaxLen(v string)`

SetSlowlogMaxLen sets SlowlogMaxLen field to given value.

### HasSlowlogMaxLen

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasSlowlogMaxLen() bool`

HasSlowlogMaxLen returns a boolean if a field has been set.

### GetSave

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetSave() string`

GetSave returns the Save field if non-nil, zero value otherwise.

### GetSaveOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetSaveOk() (*string, bool)`

GetSaveOk returns a tuple with the Save field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSave

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetSave(v string)`

SetSave sets Save field to given value.

### HasSave

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasSave() bool`

HasSave returns a boolean if a field has been set.

### GetAppendonly

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetAppendonly() string`

GetAppendonly returns the Appendonly field if non-nil, zero value otherwise.

### GetAppendonlyOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetAppendonlyOk() (*string, bool)`

GetAppendonlyOk returns a tuple with the Appendonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppendonly

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetAppendonly(v string)`

SetAppendonly sets Appendonly field to given value.

### HasAppendonly

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasAppendonly() bool`

HasAppendonly returns a boolean if a field has been set.

### GetAppendfsync

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetAppendfsync() string`

GetAppendfsync returns the Appendfsync field if non-nil, zero value otherwise.

### GetAppendfsyncOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetAppendfsyncOk() (*string, bool)`

GetAppendfsyncOk returns a tuple with the Appendfsync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppendfsync

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetAppendfsync(v string)`

SetAppendfsync sets Appendfsync field to given value.

### HasAppendfsync

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasAppendfsync() bool`

HasAppendfsync returns a boolean if a field has been set.

### GetTcpKeepalive

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetTcpKeepalive() string`

GetTcpKeepalive returns the TcpKeepalive field if non-nil, zero value otherwise.

### GetTcpKeepaliveOk

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) GetTcpKeepaliveOk() (*string, bool)`

GetTcpKeepaliveOk returns a tuple with the TcpKeepalive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpKeepalive

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) SetTcpKeepalive(v string)`

SetTcpKeepalive sets TcpKeepalive field to given value.

### HasTcpKeepalive

`func (o *GetDatabaseDefaultParameters200ResponseConfigParams) HasTcpKeepalive() bool`

HasTcpKeepalive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


