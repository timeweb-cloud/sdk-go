# ConfigParametersValkey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientOutputBufferLimitNormal** | Pointer to **string** | Ограничение буфера вывода для обычных клиентских подключений. Формат: &#x60;hard-limit soft-limit soft-seconds&#x60; (&#x60;valkey&#x60; | &#x60;valkey7&#x60; | &#x60;valkey8_1&#x60; | &#x60;valkey9_1&#x60;). | [optional] 
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

### NewConfigParametersValkey

`func NewConfigParametersValkey() *ConfigParametersValkey`

NewConfigParametersValkey instantiates a new ConfigParametersValkey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigParametersValkeyWithDefaults

`func NewConfigParametersValkeyWithDefaults() *ConfigParametersValkey`

NewConfigParametersValkeyWithDefaults instantiates a new ConfigParametersValkey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientOutputBufferLimitNormal

`func (o *ConfigParametersValkey) GetClientOutputBufferLimitNormal() string`

GetClientOutputBufferLimitNormal returns the ClientOutputBufferLimitNormal field if non-nil, zero value otherwise.

### GetClientOutputBufferLimitNormalOk

`func (o *ConfigParametersValkey) GetClientOutputBufferLimitNormalOk() (*string, bool)`

GetClientOutputBufferLimitNormalOk returns a tuple with the ClientOutputBufferLimitNormal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOutputBufferLimitNormal

`func (o *ConfigParametersValkey) SetClientOutputBufferLimitNormal(v string)`

SetClientOutputBufferLimitNormal sets ClientOutputBufferLimitNormal field to given value.

### HasClientOutputBufferLimitNormal

`func (o *ConfigParametersValkey) HasClientOutputBufferLimitNormal() bool`

HasClientOutputBufferLimitNormal returns a boolean if a field has been set.

### GetClientOutputBufferLimitPubsub

`func (o *ConfigParametersValkey) GetClientOutputBufferLimitPubsub() string`

GetClientOutputBufferLimitPubsub returns the ClientOutputBufferLimitPubsub field if non-nil, zero value otherwise.

### GetClientOutputBufferLimitPubsubOk

`func (o *ConfigParametersValkey) GetClientOutputBufferLimitPubsubOk() (*string, bool)`

GetClientOutputBufferLimitPubsubOk returns a tuple with the ClientOutputBufferLimitPubsub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOutputBufferLimitPubsub

`func (o *ConfigParametersValkey) SetClientOutputBufferLimitPubsub(v string)`

SetClientOutputBufferLimitPubsub sets ClientOutputBufferLimitPubsub field to given value.

### HasClientOutputBufferLimitPubsub

`func (o *ConfigParametersValkey) HasClientOutputBufferLimitPubsub() bool`

HasClientOutputBufferLimitPubsub returns a boolean if a field has been set.

### GetDatabases

`func (o *ConfigParametersValkey) GetDatabases() string`

GetDatabases returns the Databases field if non-nil, zero value otherwise.

### GetDatabasesOk

`func (o *ConfigParametersValkey) GetDatabasesOk() (*string, bool)`

GetDatabasesOk returns a tuple with the Databases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabases

`func (o *ConfigParametersValkey) SetDatabases(v string)`

SetDatabases sets Databases field to given value.

### HasDatabases

`func (o *ConfigParametersValkey) HasDatabases() bool`

HasDatabases returns a boolean if a field has been set.

### GetTimeout

`func (o *ConfigParametersValkey) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ConfigParametersValkey) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ConfigParametersValkey) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ConfigParametersValkey) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetMaxmemoryPolicy

`func (o *ConfigParametersValkey) GetMaxmemoryPolicy() string`

GetMaxmemoryPolicy returns the MaxmemoryPolicy field if non-nil, zero value otherwise.

### GetMaxmemoryPolicyOk

`func (o *ConfigParametersValkey) GetMaxmemoryPolicyOk() (*string, bool)`

GetMaxmemoryPolicyOk returns a tuple with the MaxmemoryPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxmemoryPolicy

`func (o *ConfigParametersValkey) SetMaxmemoryPolicy(v string)`

SetMaxmemoryPolicy sets MaxmemoryPolicy field to given value.

### HasMaxmemoryPolicy

`func (o *ConfigParametersValkey) HasMaxmemoryPolicy() bool`

HasMaxmemoryPolicy returns a boolean if a field has been set.

### GetSlowlogLogSlowerThan

`func (o *ConfigParametersValkey) GetSlowlogLogSlowerThan() string`

GetSlowlogLogSlowerThan returns the SlowlogLogSlowerThan field if non-nil, zero value otherwise.

### GetSlowlogLogSlowerThanOk

`func (o *ConfigParametersValkey) GetSlowlogLogSlowerThanOk() (*string, bool)`

GetSlowlogLogSlowerThanOk returns a tuple with the SlowlogLogSlowerThan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlowlogLogSlowerThan

`func (o *ConfigParametersValkey) SetSlowlogLogSlowerThan(v string)`

SetSlowlogLogSlowerThan sets SlowlogLogSlowerThan field to given value.

### HasSlowlogLogSlowerThan

`func (o *ConfigParametersValkey) HasSlowlogLogSlowerThan() bool`

HasSlowlogLogSlowerThan returns a boolean if a field has been set.

### GetSlowlogMaxLen

`func (o *ConfigParametersValkey) GetSlowlogMaxLen() string`

GetSlowlogMaxLen returns the SlowlogMaxLen field if non-nil, zero value otherwise.

### GetSlowlogMaxLenOk

`func (o *ConfigParametersValkey) GetSlowlogMaxLenOk() (*string, bool)`

GetSlowlogMaxLenOk returns a tuple with the SlowlogMaxLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlowlogMaxLen

`func (o *ConfigParametersValkey) SetSlowlogMaxLen(v string)`

SetSlowlogMaxLen sets SlowlogMaxLen field to given value.

### HasSlowlogMaxLen

`func (o *ConfigParametersValkey) HasSlowlogMaxLen() bool`

HasSlowlogMaxLen returns a boolean if a field has been set.

### GetSave

`func (o *ConfigParametersValkey) GetSave() string`

GetSave returns the Save field if non-nil, zero value otherwise.

### GetSaveOk

`func (o *ConfigParametersValkey) GetSaveOk() (*string, bool)`

GetSaveOk returns a tuple with the Save field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSave

`func (o *ConfigParametersValkey) SetSave(v string)`

SetSave sets Save field to given value.

### HasSave

`func (o *ConfigParametersValkey) HasSave() bool`

HasSave returns a boolean if a field has been set.

### GetAppendonly

`func (o *ConfigParametersValkey) GetAppendonly() string`

GetAppendonly returns the Appendonly field if non-nil, zero value otherwise.

### GetAppendonlyOk

`func (o *ConfigParametersValkey) GetAppendonlyOk() (*string, bool)`

GetAppendonlyOk returns a tuple with the Appendonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppendonly

`func (o *ConfigParametersValkey) SetAppendonly(v string)`

SetAppendonly sets Appendonly field to given value.

### HasAppendonly

`func (o *ConfigParametersValkey) HasAppendonly() bool`

HasAppendonly returns a boolean if a field has been set.

### GetAppendfsync

`func (o *ConfigParametersValkey) GetAppendfsync() string`

GetAppendfsync returns the Appendfsync field if non-nil, zero value otherwise.

### GetAppendfsyncOk

`func (o *ConfigParametersValkey) GetAppendfsyncOk() (*string, bool)`

GetAppendfsyncOk returns a tuple with the Appendfsync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppendfsync

`func (o *ConfigParametersValkey) SetAppendfsync(v string)`

SetAppendfsync sets Appendfsync field to given value.

### HasAppendfsync

`func (o *ConfigParametersValkey) HasAppendfsync() bool`

HasAppendfsync returns a boolean if a field has been set.

### GetTcpKeepalive

`func (o *ConfigParametersValkey) GetTcpKeepalive() string`

GetTcpKeepalive returns the TcpKeepalive field if non-nil, zero value otherwise.

### GetTcpKeepaliveOk

`func (o *ConfigParametersValkey) GetTcpKeepaliveOk() (*string, bool)`

GetTcpKeepaliveOk returns a tuple with the TcpKeepalive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpKeepalive

`func (o *ConfigParametersValkey) SetTcpKeepalive(v string)`

SetTcpKeepalive sets TcpKeepalive field to given value.

### HasTcpKeepalive

`func (o *ConfigParametersValkey) HasTcpKeepalive() bool`

HasTcpKeepalive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


