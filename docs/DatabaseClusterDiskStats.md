# DatabaseClusterDiskStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | **float32** | Размер (в Кб) диска кластера базы данных. | 
**Used** | **float32** | Размер (в Кб) использованного пространства диска кластера базы данных. | 

## Methods

### NewDatabaseClusterDiskStats

`func NewDatabaseClusterDiskStats(size float32, used float32, ) *DatabaseClusterDiskStats`

NewDatabaseClusterDiskStats instantiates a new DatabaseClusterDiskStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseClusterDiskStatsWithDefaults

`func NewDatabaseClusterDiskStatsWithDefaults() *DatabaseClusterDiskStats`

NewDatabaseClusterDiskStatsWithDefaults instantiates a new DatabaseClusterDiskStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *DatabaseClusterDiskStats) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DatabaseClusterDiskStats) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DatabaseClusterDiskStats) SetSize(v float32)`

SetSize sets Size field to given value.


### GetUsed

`func (o *DatabaseClusterDiskStats) GetUsed() float32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *DatabaseClusterDiskStats) GetUsedOk() (*float32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *DatabaseClusterDiskStats) SetUsed(v float32)`

SetUsed sets Used field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


