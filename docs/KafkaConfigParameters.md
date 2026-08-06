# KafkaConfigParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Partitions** | Pointer to **string** | Количество партиций топика. | [optional] 
**CleanupPolicy** | Pointer to **string** | Политика очистки старых сегментов лога: &#x60;delete&#x60; — удалять, &#x60;compact&#x60; — уплотнять. | [optional] 
**CompressionType** | Pointer to **string** | Тип сжатия сообщений в топике. | [optional] 
**DeleteRetentionMs** | Pointer to **string** | Время (в мс) хранения меток удаления для уплотняемых топиков. | [optional] 
**FileDeleteDelayMs** | Pointer to **string** | Задержка (в мс) перед удалением файла из файловой системы. | [optional] 
**FlushMessages** | Pointer to **string** | Количество сообщений, после которого данные принудительно сбрасываются на диск. | [optional] 
**FlushMs** | Pointer to **string** | Интервал (в мс), после которого данные принудительно сбрасываются на диск. | [optional] 
**IndexIntervalBytes** | Pointer to **string** | Интервал (в байтах), с которым Kafka добавляет запись в индекс смещений. | [optional] 
**MinCompactionLagMs** | Pointer to **string** | Минимальное время (в мс), в течение которого сообщение остается неуплотненным. | [optional] 
**MaxCompactionLagMs** | Pointer to **string** | Максимальное время (в мс), в течение которого сообщение может оставаться неуплотненным. | [optional] 
**MaxMessageBytes** | Pointer to **string** | Максимальный размер (в байтах) пакета сообщений. | [optional] 
**MessageFormatVersion** | Pointer to **string** | Версия формата сообщений, в котором Kafka добавляет сообщения в лог. | [optional] 
**MessageTimestampDifferenceMaxMs** | Pointer to **string** | Максимально допустимая разница (в мс) между временной меткой сообщения и временем его получения брокером. | [optional] 
**MessageDownconversionEnable** | Pointer to **string** | Понижение версии формата сообщений для старых клиентов. | [optional] 
**MessageTimestampType** | Pointer to **string** | Источник временной метки сообщения: &#x60;CreateTime&#x60; — время создания сообщения клиентом, &#x60;LogAppendTime&#x60; — время добавления сообщения в лог брокером. | [optional] 
**MinCleanableDirtyRatio** | Pointer to **string** | Доля неуплотненных данных в логе, при которой запускается уплотнение. | [optional] 
**MinInsyncReplicas** | Pointer to **string** | Минимальное количество синхронизированных реплик, необходимое для подтверждения записи. | [optional] 
**Preallocate** | Pointer to **string** | Предварительное выделение места на диске при создании нового сегмента лога. | [optional] 
**RetentionBytes** | Pointer to **string** | Максимальный размер (в байтах) партиции топика, после которого старые сегменты удаляются. &#x60;-1&#x60; — без ограничения. | [optional] 
**RetentionMs** | Pointer to **string** | Время (в мс) хранения сообщений в топике. &#x60;-1&#x60; — хранить бессрочно. | [optional] 
**SegmentBytes** | Pointer to **string** | Максимальный размер (в байтах) одного сегмента лога. | [optional] 
**SegmentIndexBytes** | Pointer to **string** | Максимальный размер (в байтах) индексного файла сегмента лога. | [optional] 
**SegmentJitterMs** | Pointer to **string** | Максимальное случайное отклонение (в мс) от времени ротации сегмента. | [optional] 
**SegmentMs** | Pointer to **string** | Период (в мс), после которого Kafka создает новый сегмент лога. | [optional] 
**UncleanLeaderElectionEnable** | Pointer to **string** | Возможность выбрать лидером партиции реплику, которая не входит в число синхронизированных. | [optional] 

## Methods

### NewKafkaConfigParameters

`func NewKafkaConfigParameters() *KafkaConfigParameters`

NewKafkaConfigParameters instantiates a new KafkaConfigParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaConfigParametersWithDefaults

`func NewKafkaConfigParametersWithDefaults() *KafkaConfigParameters`

NewKafkaConfigParametersWithDefaults instantiates a new KafkaConfigParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartitions

`func (o *KafkaConfigParameters) GetPartitions() string`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *KafkaConfigParameters) GetPartitionsOk() (*string, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *KafkaConfigParameters) SetPartitions(v string)`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *KafkaConfigParameters) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.

### GetCleanupPolicy

`func (o *KafkaConfigParameters) GetCleanupPolicy() string`

GetCleanupPolicy returns the CleanupPolicy field if non-nil, zero value otherwise.

### GetCleanupPolicyOk

`func (o *KafkaConfigParameters) GetCleanupPolicyOk() (*string, bool)`

GetCleanupPolicyOk returns a tuple with the CleanupPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanupPolicy

`func (o *KafkaConfigParameters) SetCleanupPolicy(v string)`

SetCleanupPolicy sets CleanupPolicy field to given value.

### HasCleanupPolicy

`func (o *KafkaConfigParameters) HasCleanupPolicy() bool`

HasCleanupPolicy returns a boolean if a field has been set.

### GetCompressionType

`func (o *KafkaConfigParameters) GetCompressionType() string`

GetCompressionType returns the CompressionType field if non-nil, zero value otherwise.

### GetCompressionTypeOk

`func (o *KafkaConfigParameters) GetCompressionTypeOk() (*string, bool)`

GetCompressionTypeOk returns a tuple with the CompressionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionType

`func (o *KafkaConfigParameters) SetCompressionType(v string)`

SetCompressionType sets CompressionType field to given value.

### HasCompressionType

`func (o *KafkaConfigParameters) HasCompressionType() bool`

HasCompressionType returns a boolean if a field has been set.

### GetDeleteRetentionMs

`func (o *KafkaConfigParameters) GetDeleteRetentionMs() string`

GetDeleteRetentionMs returns the DeleteRetentionMs field if non-nil, zero value otherwise.

### GetDeleteRetentionMsOk

`func (o *KafkaConfigParameters) GetDeleteRetentionMsOk() (*string, bool)`

GetDeleteRetentionMsOk returns a tuple with the DeleteRetentionMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteRetentionMs

`func (o *KafkaConfigParameters) SetDeleteRetentionMs(v string)`

SetDeleteRetentionMs sets DeleteRetentionMs field to given value.

### HasDeleteRetentionMs

`func (o *KafkaConfigParameters) HasDeleteRetentionMs() bool`

HasDeleteRetentionMs returns a boolean if a field has been set.

### GetFileDeleteDelayMs

`func (o *KafkaConfigParameters) GetFileDeleteDelayMs() string`

GetFileDeleteDelayMs returns the FileDeleteDelayMs field if non-nil, zero value otherwise.

### GetFileDeleteDelayMsOk

`func (o *KafkaConfigParameters) GetFileDeleteDelayMsOk() (*string, bool)`

GetFileDeleteDelayMsOk returns a tuple with the FileDeleteDelayMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileDeleteDelayMs

`func (o *KafkaConfigParameters) SetFileDeleteDelayMs(v string)`

SetFileDeleteDelayMs sets FileDeleteDelayMs field to given value.

### HasFileDeleteDelayMs

`func (o *KafkaConfigParameters) HasFileDeleteDelayMs() bool`

HasFileDeleteDelayMs returns a boolean if a field has been set.

### GetFlushMessages

`func (o *KafkaConfigParameters) GetFlushMessages() string`

GetFlushMessages returns the FlushMessages field if non-nil, zero value otherwise.

### GetFlushMessagesOk

`func (o *KafkaConfigParameters) GetFlushMessagesOk() (*string, bool)`

GetFlushMessagesOk returns a tuple with the FlushMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlushMessages

`func (o *KafkaConfigParameters) SetFlushMessages(v string)`

SetFlushMessages sets FlushMessages field to given value.

### HasFlushMessages

`func (o *KafkaConfigParameters) HasFlushMessages() bool`

HasFlushMessages returns a boolean if a field has been set.

### GetFlushMs

`func (o *KafkaConfigParameters) GetFlushMs() string`

GetFlushMs returns the FlushMs field if non-nil, zero value otherwise.

### GetFlushMsOk

`func (o *KafkaConfigParameters) GetFlushMsOk() (*string, bool)`

GetFlushMsOk returns a tuple with the FlushMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlushMs

`func (o *KafkaConfigParameters) SetFlushMs(v string)`

SetFlushMs sets FlushMs field to given value.

### HasFlushMs

`func (o *KafkaConfigParameters) HasFlushMs() bool`

HasFlushMs returns a boolean if a field has been set.

### GetIndexIntervalBytes

`func (o *KafkaConfigParameters) GetIndexIntervalBytes() string`

GetIndexIntervalBytes returns the IndexIntervalBytes field if non-nil, zero value otherwise.

### GetIndexIntervalBytesOk

`func (o *KafkaConfigParameters) GetIndexIntervalBytesOk() (*string, bool)`

GetIndexIntervalBytesOk returns a tuple with the IndexIntervalBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexIntervalBytes

`func (o *KafkaConfigParameters) SetIndexIntervalBytes(v string)`

SetIndexIntervalBytes sets IndexIntervalBytes field to given value.

### HasIndexIntervalBytes

`func (o *KafkaConfigParameters) HasIndexIntervalBytes() bool`

HasIndexIntervalBytes returns a boolean if a field has been set.

### GetMinCompactionLagMs

`func (o *KafkaConfigParameters) GetMinCompactionLagMs() string`

GetMinCompactionLagMs returns the MinCompactionLagMs field if non-nil, zero value otherwise.

### GetMinCompactionLagMsOk

`func (o *KafkaConfigParameters) GetMinCompactionLagMsOk() (*string, bool)`

GetMinCompactionLagMsOk returns a tuple with the MinCompactionLagMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCompactionLagMs

`func (o *KafkaConfigParameters) SetMinCompactionLagMs(v string)`

SetMinCompactionLagMs sets MinCompactionLagMs field to given value.

### HasMinCompactionLagMs

`func (o *KafkaConfigParameters) HasMinCompactionLagMs() bool`

HasMinCompactionLagMs returns a boolean if a field has been set.

### GetMaxCompactionLagMs

`func (o *KafkaConfigParameters) GetMaxCompactionLagMs() string`

GetMaxCompactionLagMs returns the MaxCompactionLagMs field if non-nil, zero value otherwise.

### GetMaxCompactionLagMsOk

`func (o *KafkaConfigParameters) GetMaxCompactionLagMsOk() (*string, bool)`

GetMaxCompactionLagMsOk returns a tuple with the MaxCompactionLagMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCompactionLagMs

`func (o *KafkaConfigParameters) SetMaxCompactionLagMs(v string)`

SetMaxCompactionLagMs sets MaxCompactionLagMs field to given value.

### HasMaxCompactionLagMs

`func (o *KafkaConfigParameters) HasMaxCompactionLagMs() bool`

HasMaxCompactionLagMs returns a boolean if a field has been set.

### GetMaxMessageBytes

`func (o *KafkaConfigParameters) GetMaxMessageBytes() string`

GetMaxMessageBytes returns the MaxMessageBytes field if non-nil, zero value otherwise.

### GetMaxMessageBytesOk

`func (o *KafkaConfigParameters) GetMaxMessageBytesOk() (*string, bool)`

GetMaxMessageBytesOk returns a tuple with the MaxMessageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMessageBytes

`func (o *KafkaConfigParameters) SetMaxMessageBytes(v string)`

SetMaxMessageBytes sets MaxMessageBytes field to given value.

### HasMaxMessageBytes

`func (o *KafkaConfigParameters) HasMaxMessageBytes() bool`

HasMaxMessageBytes returns a boolean if a field has been set.

### GetMessageFormatVersion

`func (o *KafkaConfigParameters) GetMessageFormatVersion() string`

GetMessageFormatVersion returns the MessageFormatVersion field if non-nil, zero value otherwise.

### GetMessageFormatVersionOk

`func (o *KafkaConfigParameters) GetMessageFormatVersionOk() (*string, bool)`

GetMessageFormatVersionOk returns a tuple with the MessageFormatVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageFormatVersion

`func (o *KafkaConfigParameters) SetMessageFormatVersion(v string)`

SetMessageFormatVersion sets MessageFormatVersion field to given value.

### HasMessageFormatVersion

`func (o *KafkaConfigParameters) HasMessageFormatVersion() bool`

HasMessageFormatVersion returns a boolean if a field has been set.

### GetMessageTimestampDifferenceMaxMs

`func (o *KafkaConfigParameters) GetMessageTimestampDifferenceMaxMs() string`

GetMessageTimestampDifferenceMaxMs returns the MessageTimestampDifferenceMaxMs field if non-nil, zero value otherwise.

### GetMessageTimestampDifferenceMaxMsOk

`func (o *KafkaConfigParameters) GetMessageTimestampDifferenceMaxMsOk() (*string, bool)`

GetMessageTimestampDifferenceMaxMsOk returns a tuple with the MessageTimestampDifferenceMaxMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTimestampDifferenceMaxMs

`func (o *KafkaConfigParameters) SetMessageTimestampDifferenceMaxMs(v string)`

SetMessageTimestampDifferenceMaxMs sets MessageTimestampDifferenceMaxMs field to given value.

### HasMessageTimestampDifferenceMaxMs

`func (o *KafkaConfigParameters) HasMessageTimestampDifferenceMaxMs() bool`

HasMessageTimestampDifferenceMaxMs returns a boolean if a field has been set.

### GetMessageDownconversionEnable

`func (o *KafkaConfigParameters) GetMessageDownconversionEnable() string`

GetMessageDownconversionEnable returns the MessageDownconversionEnable field if non-nil, zero value otherwise.

### GetMessageDownconversionEnableOk

`func (o *KafkaConfigParameters) GetMessageDownconversionEnableOk() (*string, bool)`

GetMessageDownconversionEnableOk returns a tuple with the MessageDownconversionEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageDownconversionEnable

`func (o *KafkaConfigParameters) SetMessageDownconversionEnable(v string)`

SetMessageDownconversionEnable sets MessageDownconversionEnable field to given value.

### HasMessageDownconversionEnable

`func (o *KafkaConfigParameters) HasMessageDownconversionEnable() bool`

HasMessageDownconversionEnable returns a boolean if a field has been set.

### GetMessageTimestampType

`func (o *KafkaConfigParameters) GetMessageTimestampType() string`

GetMessageTimestampType returns the MessageTimestampType field if non-nil, zero value otherwise.

### GetMessageTimestampTypeOk

`func (o *KafkaConfigParameters) GetMessageTimestampTypeOk() (*string, bool)`

GetMessageTimestampTypeOk returns a tuple with the MessageTimestampType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTimestampType

`func (o *KafkaConfigParameters) SetMessageTimestampType(v string)`

SetMessageTimestampType sets MessageTimestampType field to given value.

### HasMessageTimestampType

`func (o *KafkaConfigParameters) HasMessageTimestampType() bool`

HasMessageTimestampType returns a boolean if a field has been set.

### GetMinCleanableDirtyRatio

`func (o *KafkaConfigParameters) GetMinCleanableDirtyRatio() string`

GetMinCleanableDirtyRatio returns the MinCleanableDirtyRatio field if non-nil, zero value otherwise.

### GetMinCleanableDirtyRatioOk

`func (o *KafkaConfigParameters) GetMinCleanableDirtyRatioOk() (*string, bool)`

GetMinCleanableDirtyRatioOk returns a tuple with the MinCleanableDirtyRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCleanableDirtyRatio

`func (o *KafkaConfigParameters) SetMinCleanableDirtyRatio(v string)`

SetMinCleanableDirtyRatio sets MinCleanableDirtyRatio field to given value.

### HasMinCleanableDirtyRatio

`func (o *KafkaConfigParameters) HasMinCleanableDirtyRatio() bool`

HasMinCleanableDirtyRatio returns a boolean if a field has been set.

### GetMinInsyncReplicas

`func (o *KafkaConfigParameters) GetMinInsyncReplicas() string`

GetMinInsyncReplicas returns the MinInsyncReplicas field if non-nil, zero value otherwise.

### GetMinInsyncReplicasOk

`func (o *KafkaConfigParameters) GetMinInsyncReplicasOk() (*string, bool)`

GetMinInsyncReplicasOk returns a tuple with the MinInsyncReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinInsyncReplicas

`func (o *KafkaConfigParameters) SetMinInsyncReplicas(v string)`

SetMinInsyncReplicas sets MinInsyncReplicas field to given value.

### HasMinInsyncReplicas

`func (o *KafkaConfigParameters) HasMinInsyncReplicas() bool`

HasMinInsyncReplicas returns a boolean if a field has been set.

### GetPreallocate

`func (o *KafkaConfigParameters) GetPreallocate() string`

GetPreallocate returns the Preallocate field if non-nil, zero value otherwise.

### GetPreallocateOk

`func (o *KafkaConfigParameters) GetPreallocateOk() (*string, bool)`

GetPreallocateOk returns a tuple with the Preallocate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreallocate

`func (o *KafkaConfigParameters) SetPreallocate(v string)`

SetPreallocate sets Preallocate field to given value.

### HasPreallocate

`func (o *KafkaConfigParameters) HasPreallocate() bool`

HasPreallocate returns a boolean if a field has been set.

### GetRetentionBytes

`func (o *KafkaConfigParameters) GetRetentionBytes() string`

GetRetentionBytes returns the RetentionBytes field if non-nil, zero value otherwise.

### GetRetentionBytesOk

`func (o *KafkaConfigParameters) GetRetentionBytesOk() (*string, bool)`

GetRetentionBytesOk returns a tuple with the RetentionBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionBytes

`func (o *KafkaConfigParameters) SetRetentionBytes(v string)`

SetRetentionBytes sets RetentionBytes field to given value.

### HasRetentionBytes

`func (o *KafkaConfigParameters) HasRetentionBytes() bool`

HasRetentionBytes returns a boolean if a field has been set.

### GetRetentionMs

`func (o *KafkaConfigParameters) GetRetentionMs() string`

GetRetentionMs returns the RetentionMs field if non-nil, zero value otherwise.

### GetRetentionMsOk

`func (o *KafkaConfigParameters) GetRetentionMsOk() (*string, bool)`

GetRetentionMsOk returns a tuple with the RetentionMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionMs

`func (o *KafkaConfigParameters) SetRetentionMs(v string)`

SetRetentionMs sets RetentionMs field to given value.

### HasRetentionMs

`func (o *KafkaConfigParameters) HasRetentionMs() bool`

HasRetentionMs returns a boolean if a field has been set.

### GetSegmentBytes

`func (o *KafkaConfigParameters) GetSegmentBytes() string`

GetSegmentBytes returns the SegmentBytes field if non-nil, zero value otherwise.

### GetSegmentBytesOk

`func (o *KafkaConfigParameters) GetSegmentBytesOk() (*string, bool)`

GetSegmentBytesOk returns a tuple with the SegmentBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentBytes

`func (o *KafkaConfigParameters) SetSegmentBytes(v string)`

SetSegmentBytes sets SegmentBytes field to given value.

### HasSegmentBytes

`func (o *KafkaConfigParameters) HasSegmentBytes() bool`

HasSegmentBytes returns a boolean if a field has been set.

### GetSegmentIndexBytes

`func (o *KafkaConfigParameters) GetSegmentIndexBytes() string`

GetSegmentIndexBytes returns the SegmentIndexBytes field if non-nil, zero value otherwise.

### GetSegmentIndexBytesOk

`func (o *KafkaConfigParameters) GetSegmentIndexBytesOk() (*string, bool)`

GetSegmentIndexBytesOk returns a tuple with the SegmentIndexBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentIndexBytes

`func (o *KafkaConfigParameters) SetSegmentIndexBytes(v string)`

SetSegmentIndexBytes sets SegmentIndexBytes field to given value.

### HasSegmentIndexBytes

`func (o *KafkaConfigParameters) HasSegmentIndexBytes() bool`

HasSegmentIndexBytes returns a boolean if a field has been set.

### GetSegmentJitterMs

`func (o *KafkaConfigParameters) GetSegmentJitterMs() string`

GetSegmentJitterMs returns the SegmentJitterMs field if non-nil, zero value otherwise.

### GetSegmentJitterMsOk

`func (o *KafkaConfigParameters) GetSegmentJitterMsOk() (*string, bool)`

GetSegmentJitterMsOk returns a tuple with the SegmentJitterMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentJitterMs

`func (o *KafkaConfigParameters) SetSegmentJitterMs(v string)`

SetSegmentJitterMs sets SegmentJitterMs field to given value.

### HasSegmentJitterMs

`func (o *KafkaConfigParameters) HasSegmentJitterMs() bool`

HasSegmentJitterMs returns a boolean if a field has been set.

### GetSegmentMs

`func (o *KafkaConfigParameters) GetSegmentMs() string`

GetSegmentMs returns the SegmentMs field if non-nil, zero value otherwise.

### GetSegmentMsOk

`func (o *KafkaConfigParameters) GetSegmentMsOk() (*string, bool)`

GetSegmentMsOk returns a tuple with the SegmentMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentMs

`func (o *KafkaConfigParameters) SetSegmentMs(v string)`

SetSegmentMs sets SegmentMs field to given value.

### HasSegmentMs

`func (o *KafkaConfigParameters) HasSegmentMs() bool`

HasSegmentMs returns a boolean if a field has been set.

### GetUncleanLeaderElectionEnable

`func (o *KafkaConfigParameters) GetUncleanLeaderElectionEnable() string`

GetUncleanLeaderElectionEnable returns the UncleanLeaderElectionEnable field if non-nil, zero value otherwise.

### GetUncleanLeaderElectionEnableOk

`func (o *KafkaConfigParameters) GetUncleanLeaderElectionEnableOk() (*string, bool)`

GetUncleanLeaderElectionEnableOk returns a tuple with the UncleanLeaderElectionEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncleanLeaderElectionEnable

`func (o *KafkaConfigParameters) SetUncleanLeaderElectionEnable(v string)`

SetUncleanLeaderElectionEnable sets UncleanLeaderElectionEnable field to given value.

### HasUncleanLeaderElectionEnable

`func (o *KafkaConfigParameters) HasUncleanLeaderElectionEnable() bool`

HasUncleanLeaderElectionEnable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


