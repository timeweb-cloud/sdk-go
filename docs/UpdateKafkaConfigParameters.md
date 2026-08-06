# UpdateKafkaConfigParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Partitions** | Pointer to **int32** | Количество партиций топика. Количество партиций нельзя уменьшить: если передать значение меньше текущего, останется текущее. | [optional] 
**CleanupPolicy** | Pointer to **string** | Политика очистки старых сегментов лога: &#x60;delete&#x60; — удалять, &#x60;compact&#x60; — уплотнять. | [optional] 
**CompressionType** | Pointer to **string** | Тип сжатия сообщений в топике. | [optional] 
**DeleteRetentionMs** | Pointer to **int64** | Время (в мс) хранения меток удаления для уплотняемых топиков. Максимальное значение — 9223372036854775807. | [optional] 
**FileDeleteDelayMs** | Pointer to **int64** | Задержка (в мс) перед удалением файла из файловой системы. Максимальное значение — 9223372036854775807. | [optional] 
**FlushMessages** | Pointer to **int64** | Количество сообщений, после которого данные принудительно сбрасываются на диск. Максимальное значение — 9223372036854775807. | [optional] 
**FlushMs** | Pointer to **int64** | Интервал (в мс), после которого данные принудительно сбрасываются на диск. Максимальное значение — 9223372036854775807. | [optional] 
**IndexIntervalBytes** | Pointer to **int32** | Интервал (в байтах), с которым Kafka добавляет запись в индекс смещений. | [optional] 
**MinCompactionLagMs** | Pointer to **int64** | Минимальное время (в мс), в течение которого сообщение остается неуплотненным. Максимальное значение — 9223372036854775807. | [optional] 
**MaxCompactionLagMs** | Pointer to **int64** | Максимальное время (в мс), в течение которого сообщение может оставаться неуплотненным. Максимальное значение — 9223372036854775807. | [optional] 
**MaxMessageBytes** | Pointer to **int32** | Максимальный размер (в байтах) пакета сообщений. | [optional] 
**MessageFormatVersion** | Pointer to **string** | Версия формата сообщений, в котором Kafka добавляет сообщения в лог. | [optional] 
**MessageTimestampDifferenceMaxMs** | Pointer to **int64** | Максимально допустимая разница (в мс) между временной меткой сообщения и временем его получения брокером. Максимальное значение — 9223372036854775807. | [optional] 
**MessageDownconversionEnable** | Pointer to **string** | Понижение версии формата сообщений для старых клиентов. | [optional] 
**MessageTimestampType** | Pointer to **string** | Источник временной метки сообщения: &#x60;CreateTime&#x60; — время создания сообщения клиентом, &#x60;LogAppendTime&#x60; — время добавления сообщения в лог брокером. | [optional] 
**MinCleanableDirtyRatio** | Pointer to **float32** | Доля неуплотненных данных в логе, при которой запускается уплотнение. | [optional] 
**MinInsyncReplicas** | Pointer to **int32** | Минимальное количество синхронизированных реплик, необходимое для подтверждения записи. | [optional] 
**Preallocate** | Pointer to **string** | Предварительное выделение места на диске при создании нового сегмента лога. | [optional] 
**RetentionBytes** | Pointer to **int64** | Максимальный размер (в байтах) партиции топика, после которого старые сегменты удаляются. &#x60;-1&#x60; — без ограничения. Максимальное значение — 9223372036854775807. | [optional] 
**RetentionMs** | Pointer to **int64** | Время (в мс) хранения сообщений в топике. &#x60;-1&#x60; — хранить бессрочно. Максимальное значение — 9223372036854775807. | [optional] 
**SegmentBytes** | Pointer to **int32** | Максимальный размер (в байтах) одного сегмента лога. | [optional] 
**SegmentIndexBytes** | Pointer to **int32** | Максимальный размер (в байтах) индексного файла сегмента лога. | [optional] 
**SegmentJitterMs** | Pointer to **int64** | Максимальное случайное отклонение (в мс) от времени ротации сегмента. Максимальное значение — 9223372036854775807. | [optional] 
**SegmentMs** | Pointer to **int64** | Период (в мс), после которого Kafka создает новый сегмент лога. Максимальное значение — 9223372036854775807. | [optional] 
**UncleanLeaderElectionEnable** | Pointer to **string** | Возможность выбрать лидером партиции реплику, которая не входит в число синхронизированных. | [optional] 

## Methods

### NewUpdateKafkaConfigParameters

`func NewUpdateKafkaConfigParameters() *UpdateKafkaConfigParameters`

NewUpdateKafkaConfigParameters instantiates a new UpdateKafkaConfigParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateKafkaConfigParametersWithDefaults

`func NewUpdateKafkaConfigParametersWithDefaults() *UpdateKafkaConfigParameters`

NewUpdateKafkaConfigParametersWithDefaults instantiates a new UpdateKafkaConfigParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartitions

`func (o *UpdateKafkaConfigParameters) GetPartitions() int32`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *UpdateKafkaConfigParameters) GetPartitionsOk() (*int32, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *UpdateKafkaConfigParameters) SetPartitions(v int32)`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *UpdateKafkaConfigParameters) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.

### GetCleanupPolicy

`func (o *UpdateKafkaConfigParameters) GetCleanupPolicy() string`

GetCleanupPolicy returns the CleanupPolicy field if non-nil, zero value otherwise.

### GetCleanupPolicyOk

`func (o *UpdateKafkaConfigParameters) GetCleanupPolicyOk() (*string, bool)`

GetCleanupPolicyOk returns a tuple with the CleanupPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanupPolicy

`func (o *UpdateKafkaConfigParameters) SetCleanupPolicy(v string)`

SetCleanupPolicy sets CleanupPolicy field to given value.

### HasCleanupPolicy

`func (o *UpdateKafkaConfigParameters) HasCleanupPolicy() bool`

HasCleanupPolicy returns a boolean if a field has been set.

### GetCompressionType

`func (o *UpdateKafkaConfigParameters) GetCompressionType() string`

GetCompressionType returns the CompressionType field if non-nil, zero value otherwise.

### GetCompressionTypeOk

`func (o *UpdateKafkaConfigParameters) GetCompressionTypeOk() (*string, bool)`

GetCompressionTypeOk returns a tuple with the CompressionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionType

`func (o *UpdateKafkaConfigParameters) SetCompressionType(v string)`

SetCompressionType sets CompressionType field to given value.

### HasCompressionType

`func (o *UpdateKafkaConfigParameters) HasCompressionType() bool`

HasCompressionType returns a boolean if a field has been set.

### GetDeleteRetentionMs

`func (o *UpdateKafkaConfigParameters) GetDeleteRetentionMs() int64`

GetDeleteRetentionMs returns the DeleteRetentionMs field if non-nil, zero value otherwise.

### GetDeleteRetentionMsOk

`func (o *UpdateKafkaConfigParameters) GetDeleteRetentionMsOk() (*int64, bool)`

GetDeleteRetentionMsOk returns a tuple with the DeleteRetentionMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteRetentionMs

`func (o *UpdateKafkaConfigParameters) SetDeleteRetentionMs(v int64)`

SetDeleteRetentionMs sets DeleteRetentionMs field to given value.

### HasDeleteRetentionMs

`func (o *UpdateKafkaConfigParameters) HasDeleteRetentionMs() bool`

HasDeleteRetentionMs returns a boolean if a field has been set.

### GetFileDeleteDelayMs

`func (o *UpdateKafkaConfigParameters) GetFileDeleteDelayMs() int64`

GetFileDeleteDelayMs returns the FileDeleteDelayMs field if non-nil, zero value otherwise.

### GetFileDeleteDelayMsOk

`func (o *UpdateKafkaConfigParameters) GetFileDeleteDelayMsOk() (*int64, bool)`

GetFileDeleteDelayMsOk returns a tuple with the FileDeleteDelayMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileDeleteDelayMs

`func (o *UpdateKafkaConfigParameters) SetFileDeleteDelayMs(v int64)`

SetFileDeleteDelayMs sets FileDeleteDelayMs field to given value.

### HasFileDeleteDelayMs

`func (o *UpdateKafkaConfigParameters) HasFileDeleteDelayMs() bool`

HasFileDeleteDelayMs returns a boolean if a field has been set.

### GetFlushMessages

`func (o *UpdateKafkaConfigParameters) GetFlushMessages() int64`

GetFlushMessages returns the FlushMessages field if non-nil, zero value otherwise.

### GetFlushMessagesOk

`func (o *UpdateKafkaConfigParameters) GetFlushMessagesOk() (*int64, bool)`

GetFlushMessagesOk returns a tuple with the FlushMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlushMessages

`func (o *UpdateKafkaConfigParameters) SetFlushMessages(v int64)`

SetFlushMessages sets FlushMessages field to given value.

### HasFlushMessages

`func (o *UpdateKafkaConfigParameters) HasFlushMessages() bool`

HasFlushMessages returns a boolean if a field has been set.

### GetFlushMs

`func (o *UpdateKafkaConfigParameters) GetFlushMs() int64`

GetFlushMs returns the FlushMs field if non-nil, zero value otherwise.

### GetFlushMsOk

`func (o *UpdateKafkaConfigParameters) GetFlushMsOk() (*int64, bool)`

GetFlushMsOk returns a tuple with the FlushMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlushMs

`func (o *UpdateKafkaConfigParameters) SetFlushMs(v int64)`

SetFlushMs sets FlushMs field to given value.

### HasFlushMs

`func (o *UpdateKafkaConfigParameters) HasFlushMs() bool`

HasFlushMs returns a boolean if a field has been set.

### GetIndexIntervalBytes

`func (o *UpdateKafkaConfigParameters) GetIndexIntervalBytes() int32`

GetIndexIntervalBytes returns the IndexIntervalBytes field if non-nil, zero value otherwise.

### GetIndexIntervalBytesOk

`func (o *UpdateKafkaConfigParameters) GetIndexIntervalBytesOk() (*int32, bool)`

GetIndexIntervalBytesOk returns a tuple with the IndexIntervalBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexIntervalBytes

`func (o *UpdateKafkaConfigParameters) SetIndexIntervalBytes(v int32)`

SetIndexIntervalBytes sets IndexIntervalBytes field to given value.

### HasIndexIntervalBytes

`func (o *UpdateKafkaConfigParameters) HasIndexIntervalBytes() bool`

HasIndexIntervalBytes returns a boolean if a field has been set.

### GetMinCompactionLagMs

`func (o *UpdateKafkaConfigParameters) GetMinCompactionLagMs() int64`

GetMinCompactionLagMs returns the MinCompactionLagMs field if non-nil, zero value otherwise.

### GetMinCompactionLagMsOk

`func (o *UpdateKafkaConfigParameters) GetMinCompactionLagMsOk() (*int64, bool)`

GetMinCompactionLagMsOk returns a tuple with the MinCompactionLagMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCompactionLagMs

`func (o *UpdateKafkaConfigParameters) SetMinCompactionLagMs(v int64)`

SetMinCompactionLagMs sets MinCompactionLagMs field to given value.

### HasMinCompactionLagMs

`func (o *UpdateKafkaConfigParameters) HasMinCompactionLagMs() bool`

HasMinCompactionLagMs returns a boolean if a field has been set.

### GetMaxCompactionLagMs

`func (o *UpdateKafkaConfigParameters) GetMaxCompactionLagMs() int64`

GetMaxCompactionLagMs returns the MaxCompactionLagMs field if non-nil, zero value otherwise.

### GetMaxCompactionLagMsOk

`func (o *UpdateKafkaConfigParameters) GetMaxCompactionLagMsOk() (*int64, bool)`

GetMaxCompactionLagMsOk returns a tuple with the MaxCompactionLagMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCompactionLagMs

`func (o *UpdateKafkaConfigParameters) SetMaxCompactionLagMs(v int64)`

SetMaxCompactionLagMs sets MaxCompactionLagMs field to given value.

### HasMaxCompactionLagMs

`func (o *UpdateKafkaConfigParameters) HasMaxCompactionLagMs() bool`

HasMaxCompactionLagMs returns a boolean if a field has been set.

### GetMaxMessageBytes

`func (o *UpdateKafkaConfigParameters) GetMaxMessageBytes() int32`

GetMaxMessageBytes returns the MaxMessageBytes field if non-nil, zero value otherwise.

### GetMaxMessageBytesOk

`func (o *UpdateKafkaConfigParameters) GetMaxMessageBytesOk() (*int32, bool)`

GetMaxMessageBytesOk returns a tuple with the MaxMessageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMessageBytes

`func (o *UpdateKafkaConfigParameters) SetMaxMessageBytes(v int32)`

SetMaxMessageBytes sets MaxMessageBytes field to given value.

### HasMaxMessageBytes

`func (o *UpdateKafkaConfigParameters) HasMaxMessageBytes() bool`

HasMaxMessageBytes returns a boolean if a field has been set.

### GetMessageFormatVersion

`func (o *UpdateKafkaConfigParameters) GetMessageFormatVersion() string`

GetMessageFormatVersion returns the MessageFormatVersion field if non-nil, zero value otherwise.

### GetMessageFormatVersionOk

`func (o *UpdateKafkaConfigParameters) GetMessageFormatVersionOk() (*string, bool)`

GetMessageFormatVersionOk returns a tuple with the MessageFormatVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageFormatVersion

`func (o *UpdateKafkaConfigParameters) SetMessageFormatVersion(v string)`

SetMessageFormatVersion sets MessageFormatVersion field to given value.

### HasMessageFormatVersion

`func (o *UpdateKafkaConfigParameters) HasMessageFormatVersion() bool`

HasMessageFormatVersion returns a boolean if a field has been set.

### GetMessageTimestampDifferenceMaxMs

`func (o *UpdateKafkaConfigParameters) GetMessageTimestampDifferenceMaxMs() int64`

GetMessageTimestampDifferenceMaxMs returns the MessageTimestampDifferenceMaxMs field if non-nil, zero value otherwise.

### GetMessageTimestampDifferenceMaxMsOk

`func (o *UpdateKafkaConfigParameters) GetMessageTimestampDifferenceMaxMsOk() (*int64, bool)`

GetMessageTimestampDifferenceMaxMsOk returns a tuple with the MessageTimestampDifferenceMaxMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTimestampDifferenceMaxMs

`func (o *UpdateKafkaConfigParameters) SetMessageTimestampDifferenceMaxMs(v int64)`

SetMessageTimestampDifferenceMaxMs sets MessageTimestampDifferenceMaxMs field to given value.

### HasMessageTimestampDifferenceMaxMs

`func (o *UpdateKafkaConfigParameters) HasMessageTimestampDifferenceMaxMs() bool`

HasMessageTimestampDifferenceMaxMs returns a boolean if a field has been set.

### GetMessageDownconversionEnable

`func (o *UpdateKafkaConfigParameters) GetMessageDownconversionEnable() string`

GetMessageDownconversionEnable returns the MessageDownconversionEnable field if non-nil, zero value otherwise.

### GetMessageDownconversionEnableOk

`func (o *UpdateKafkaConfigParameters) GetMessageDownconversionEnableOk() (*string, bool)`

GetMessageDownconversionEnableOk returns a tuple with the MessageDownconversionEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageDownconversionEnable

`func (o *UpdateKafkaConfigParameters) SetMessageDownconversionEnable(v string)`

SetMessageDownconversionEnable sets MessageDownconversionEnable field to given value.

### HasMessageDownconversionEnable

`func (o *UpdateKafkaConfigParameters) HasMessageDownconversionEnable() bool`

HasMessageDownconversionEnable returns a boolean if a field has been set.

### GetMessageTimestampType

`func (o *UpdateKafkaConfigParameters) GetMessageTimestampType() string`

GetMessageTimestampType returns the MessageTimestampType field if non-nil, zero value otherwise.

### GetMessageTimestampTypeOk

`func (o *UpdateKafkaConfigParameters) GetMessageTimestampTypeOk() (*string, bool)`

GetMessageTimestampTypeOk returns a tuple with the MessageTimestampType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTimestampType

`func (o *UpdateKafkaConfigParameters) SetMessageTimestampType(v string)`

SetMessageTimestampType sets MessageTimestampType field to given value.

### HasMessageTimestampType

`func (o *UpdateKafkaConfigParameters) HasMessageTimestampType() bool`

HasMessageTimestampType returns a boolean if a field has been set.

### GetMinCleanableDirtyRatio

`func (o *UpdateKafkaConfigParameters) GetMinCleanableDirtyRatio() float32`

GetMinCleanableDirtyRatio returns the MinCleanableDirtyRatio field if non-nil, zero value otherwise.

### GetMinCleanableDirtyRatioOk

`func (o *UpdateKafkaConfigParameters) GetMinCleanableDirtyRatioOk() (*float32, bool)`

GetMinCleanableDirtyRatioOk returns a tuple with the MinCleanableDirtyRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCleanableDirtyRatio

`func (o *UpdateKafkaConfigParameters) SetMinCleanableDirtyRatio(v float32)`

SetMinCleanableDirtyRatio sets MinCleanableDirtyRatio field to given value.

### HasMinCleanableDirtyRatio

`func (o *UpdateKafkaConfigParameters) HasMinCleanableDirtyRatio() bool`

HasMinCleanableDirtyRatio returns a boolean if a field has been set.

### GetMinInsyncReplicas

`func (o *UpdateKafkaConfigParameters) GetMinInsyncReplicas() int32`

GetMinInsyncReplicas returns the MinInsyncReplicas field if non-nil, zero value otherwise.

### GetMinInsyncReplicasOk

`func (o *UpdateKafkaConfigParameters) GetMinInsyncReplicasOk() (*int32, bool)`

GetMinInsyncReplicasOk returns a tuple with the MinInsyncReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinInsyncReplicas

`func (o *UpdateKafkaConfigParameters) SetMinInsyncReplicas(v int32)`

SetMinInsyncReplicas sets MinInsyncReplicas field to given value.

### HasMinInsyncReplicas

`func (o *UpdateKafkaConfigParameters) HasMinInsyncReplicas() bool`

HasMinInsyncReplicas returns a boolean if a field has been set.

### GetPreallocate

`func (o *UpdateKafkaConfigParameters) GetPreallocate() string`

GetPreallocate returns the Preallocate field if non-nil, zero value otherwise.

### GetPreallocateOk

`func (o *UpdateKafkaConfigParameters) GetPreallocateOk() (*string, bool)`

GetPreallocateOk returns a tuple with the Preallocate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreallocate

`func (o *UpdateKafkaConfigParameters) SetPreallocate(v string)`

SetPreallocate sets Preallocate field to given value.

### HasPreallocate

`func (o *UpdateKafkaConfigParameters) HasPreallocate() bool`

HasPreallocate returns a boolean if a field has been set.

### GetRetentionBytes

`func (o *UpdateKafkaConfigParameters) GetRetentionBytes() int64`

GetRetentionBytes returns the RetentionBytes field if non-nil, zero value otherwise.

### GetRetentionBytesOk

`func (o *UpdateKafkaConfigParameters) GetRetentionBytesOk() (*int64, bool)`

GetRetentionBytesOk returns a tuple with the RetentionBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionBytes

`func (o *UpdateKafkaConfigParameters) SetRetentionBytes(v int64)`

SetRetentionBytes sets RetentionBytes field to given value.

### HasRetentionBytes

`func (o *UpdateKafkaConfigParameters) HasRetentionBytes() bool`

HasRetentionBytes returns a boolean if a field has been set.

### GetRetentionMs

`func (o *UpdateKafkaConfigParameters) GetRetentionMs() int64`

GetRetentionMs returns the RetentionMs field if non-nil, zero value otherwise.

### GetRetentionMsOk

`func (o *UpdateKafkaConfigParameters) GetRetentionMsOk() (*int64, bool)`

GetRetentionMsOk returns a tuple with the RetentionMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionMs

`func (o *UpdateKafkaConfigParameters) SetRetentionMs(v int64)`

SetRetentionMs sets RetentionMs field to given value.

### HasRetentionMs

`func (o *UpdateKafkaConfigParameters) HasRetentionMs() bool`

HasRetentionMs returns a boolean if a field has been set.

### GetSegmentBytes

`func (o *UpdateKafkaConfigParameters) GetSegmentBytes() int32`

GetSegmentBytes returns the SegmentBytes field if non-nil, zero value otherwise.

### GetSegmentBytesOk

`func (o *UpdateKafkaConfigParameters) GetSegmentBytesOk() (*int32, bool)`

GetSegmentBytesOk returns a tuple with the SegmentBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentBytes

`func (o *UpdateKafkaConfigParameters) SetSegmentBytes(v int32)`

SetSegmentBytes sets SegmentBytes field to given value.

### HasSegmentBytes

`func (o *UpdateKafkaConfigParameters) HasSegmentBytes() bool`

HasSegmentBytes returns a boolean if a field has been set.

### GetSegmentIndexBytes

`func (o *UpdateKafkaConfigParameters) GetSegmentIndexBytes() int32`

GetSegmentIndexBytes returns the SegmentIndexBytes field if non-nil, zero value otherwise.

### GetSegmentIndexBytesOk

`func (o *UpdateKafkaConfigParameters) GetSegmentIndexBytesOk() (*int32, bool)`

GetSegmentIndexBytesOk returns a tuple with the SegmentIndexBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentIndexBytes

`func (o *UpdateKafkaConfigParameters) SetSegmentIndexBytes(v int32)`

SetSegmentIndexBytes sets SegmentIndexBytes field to given value.

### HasSegmentIndexBytes

`func (o *UpdateKafkaConfigParameters) HasSegmentIndexBytes() bool`

HasSegmentIndexBytes returns a boolean if a field has been set.

### GetSegmentJitterMs

`func (o *UpdateKafkaConfigParameters) GetSegmentJitterMs() int64`

GetSegmentJitterMs returns the SegmentJitterMs field if non-nil, zero value otherwise.

### GetSegmentJitterMsOk

`func (o *UpdateKafkaConfigParameters) GetSegmentJitterMsOk() (*int64, bool)`

GetSegmentJitterMsOk returns a tuple with the SegmentJitterMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentJitterMs

`func (o *UpdateKafkaConfigParameters) SetSegmentJitterMs(v int64)`

SetSegmentJitterMs sets SegmentJitterMs field to given value.

### HasSegmentJitterMs

`func (o *UpdateKafkaConfigParameters) HasSegmentJitterMs() bool`

HasSegmentJitterMs returns a boolean if a field has been set.

### GetSegmentMs

`func (o *UpdateKafkaConfigParameters) GetSegmentMs() int64`

GetSegmentMs returns the SegmentMs field if non-nil, zero value otherwise.

### GetSegmentMsOk

`func (o *UpdateKafkaConfigParameters) GetSegmentMsOk() (*int64, bool)`

GetSegmentMsOk returns a tuple with the SegmentMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentMs

`func (o *UpdateKafkaConfigParameters) SetSegmentMs(v int64)`

SetSegmentMs sets SegmentMs field to given value.

### HasSegmentMs

`func (o *UpdateKafkaConfigParameters) HasSegmentMs() bool`

HasSegmentMs returns a boolean if a field has been set.

### GetUncleanLeaderElectionEnable

`func (o *UpdateKafkaConfigParameters) GetUncleanLeaderElectionEnable() string`

GetUncleanLeaderElectionEnable returns the UncleanLeaderElectionEnable field if non-nil, zero value otherwise.

### GetUncleanLeaderElectionEnableOk

`func (o *UpdateKafkaConfigParameters) GetUncleanLeaderElectionEnableOk() (*string, bool)`

GetUncleanLeaderElectionEnableOk returns a tuple with the UncleanLeaderElectionEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncleanLeaderElectionEnable

`func (o *UpdateKafkaConfigParameters) SetUncleanLeaderElectionEnable(v string)`

SetUncleanLeaderElectionEnable sets UncleanLeaderElectionEnable field to given value.

### HasUncleanLeaderElectionEnable

`func (o *UpdateKafkaConfigParameters) HasUncleanLeaderElectionEnable() bool`

HasUncleanLeaderElectionEnable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


