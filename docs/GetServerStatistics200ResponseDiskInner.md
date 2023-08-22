# GetServerStatistics200ResponseDiskInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoggedAt** | **time.Time** | Дата события в формате ISO 8061 | 
**Write** | **float32** | Нагрузка на запись диска в Мб/с | 
**Read** | **float32** | Нагрузка на чтение диска в Мб/с | 

## Methods

### NewGetServerStatistics200ResponseDiskInner

`func NewGetServerStatistics200ResponseDiskInner(loggedAt time.Time, write float32, read float32, ) *GetServerStatistics200ResponseDiskInner`

NewGetServerStatistics200ResponseDiskInner instantiates a new GetServerStatistics200ResponseDiskInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServerStatistics200ResponseDiskInnerWithDefaults

`func NewGetServerStatistics200ResponseDiskInnerWithDefaults() *GetServerStatistics200ResponseDiskInner`

NewGetServerStatistics200ResponseDiskInnerWithDefaults instantiates a new GetServerStatistics200ResponseDiskInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoggedAt

`func (o *GetServerStatistics200ResponseDiskInner) GetLoggedAt() time.Time`

GetLoggedAt returns the LoggedAt field if non-nil, zero value otherwise.

### GetLoggedAtOk

`func (o *GetServerStatistics200ResponseDiskInner) GetLoggedAtOk() (*time.Time, bool)`

GetLoggedAtOk returns a tuple with the LoggedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggedAt

`func (o *GetServerStatistics200ResponseDiskInner) SetLoggedAt(v time.Time)`

SetLoggedAt sets LoggedAt field to given value.


### GetWrite

`func (o *GetServerStatistics200ResponseDiskInner) GetWrite() float32`

GetWrite returns the Write field if non-nil, zero value otherwise.

### GetWriteOk

`func (o *GetServerStatistics200ResponseDiskInner) GetWriteOk() (*float32, bool)`

GetWriteOk returns a tuple with the Write field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrite

`func (o *GetServerStatistics200ResponseDiskInner) SetWrite(v float32)`

SetWrite sets Write field to given value.


### GetRead

`func (o *GetServerStatistics200ResponseDiskInner) GetRead() float32`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *GetServerStatistics200ResponseDiskInner) GetReadOk() (*float32, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *GetServerStatistics200ResponseDiskInner) SetRead(v float32)`

SetRead sets Read field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


