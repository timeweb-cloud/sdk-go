# DbDiskStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | **float32** | Размер (в Кб) диска базы данных. | 
**Used** | **float32** | Размер (в Кб) использованного пространства диска базы данных. | 

## Methods

### NewDbDiskStats

`func NewDbDiskStats(size float32, used float32, ) *DbDiskStats`

NewDbDiskStats instantiates a new DbDiskStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbDiskStatsWithDefaults

`func NewDbDiskStatsWithDefaults() *DbDiskStats`

NewDbDiskStatsWithDefaults instantiates a new DbDiskStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *DbDiskStats) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DbDiskStats) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DbDiskStats) SetSize(v float32)`

SetSize sets Size field to given value.


### GetUsed

`func (o *DbDiskStats) GetUsed() float32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *DbDiskStats) GetUsedOk() (*float32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *DbDiskStats) SetUsed(v float32)`

SetUsed sets Used field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


