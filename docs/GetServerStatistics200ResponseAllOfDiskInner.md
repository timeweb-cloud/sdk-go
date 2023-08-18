# GetServerStatistics200ResponseAllOfDiskInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoggedAt** | **time.Time** | Дата события в формате ISO 8061 | 
**Write** | **float32** | Нагрузка на запись диска в Мб/с | 
**Read** | **float32** | Нагрузка на чтение диска в Мб/с | 

## Methods

### NewGetServerStatistics200ResponseAllOfDiskInner

`func NewGetServerStatistics200ResponseAllOfDiskInner(loggedAt time.Time, write float32, read float32, ) *GetServerStatistics200ResponseAllOfDiskInner`

NewGetServerStatistics200ResponseAllOfDiskInner instantiates a new GetServerStatistics200ResponseAllOfDiskInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServerStatistics200ResponseAllOfDiskInnerWithDefaults

`func NewGetServerStatistics200ResponseAllOfDiskInnerWithDefaults() *GetServerStatistics200ResponseAllOfDiskInner`

NewGetServerStatistics200ResponseAllOfDiskInnerWithDefaults instantiates a new GetServerStatistics200ResponseAllOfDiskInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoggedAt

`func (o *GetServerStatistics200ResponseAllOfDiskInner) GetLoggedAt() time.Time`

GetLoggedAt returns the LoggedAt field if non-nil, zero value otherwise.

### GetLoggedAtOk

`func (o *GetServerStatistics200ResponseAllOfDiskInner) GetLoggedAtOk() (*time.Time, bool)`

GetLoggedAtOk returns a tuple with the LoggedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggedAt

`func (o *GetServerStatistics200ResponseAllOfDiskInner) SetLoggedAt(v time.Time)`

SetLoggedAt sets LoggedAt field to given value.


### GetWrite

`func (o *GetServerStatistics200ResponseAllOfDiskInner) GetWrite() float32`

GetWrite returns the Write field if non-nil, zero value otherwise.

### GetWriteOk

`func (o *GetServerStatistics200ResponseAllOfDiskInner) GetWriteOk() (*float32, bool)`

GetWriteOk returns a tuple with the Write field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrite

`func (o *GetServerStatistics200ResponseAllOfDiskInner) SetWrite(v float32)`

SetWrite sets Write field to given value.


### GetRead

`func (o *GetServerStatistics200ResponseAllOfDiskInner) GetRead() float32`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *GetServerStatistics200ResponseAllOfDiskInner) GetReadOk() (*float32, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *GetServerStatistics200ResponseAllOfDiskInner) SetRead(v float32)`

SetRead sets Read field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


