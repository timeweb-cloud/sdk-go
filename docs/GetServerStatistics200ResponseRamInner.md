# GetServerStatistics200ResponseRamInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoggedAt** | **time.Time** | Дата события в формате ISO 8061 | 
**Total** | **float32** | Общее количество оперативной памяти в Мб | 
**Used** | **float32** | Количество потревляемой оперативной памяти в Мб | 
**UsedCached** | **float32** | Количество закешированной оперативной памяти в Мб | 
**Available** | **float32** | Количество доступной оперативной памяти в Мб | 

## Methods

### NewGetServerStatistics200ResponseRamInner

`func NewGetServerStatistics200ResponseRamInner(loggedAt time.Time, total float32, used float32, usedCached float32, available float32, ) *GetServerStatistics200ResponseRamInner`

NewGetServerStatistics200ResponseRamInner instantiates a new GetServerStatistics200ResponseRamInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServerStatistics200ResponseRamInnerWithDefaults

`func NewGetServerStatistics200ResponseRamInnerWithDefaults() *GetServerStatistics200ResponseRamInner`

NewGetServerStatistics200ResponseRamInnerWithDefaults instantiates a new GetServerStatistics200ResponseRamInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoggedAt

`func (o *GetServerStatistics200ResponseRamInner) GetLoggedAt() time.Time`

GetLoggedAt returns the LoggedAt field if non-nil, zero value otherwise.

### GetLoggedAtOk

`func (o *GetServerStatistics200ResponseRamInner) GetLoggedAtOk() (*time.Time, bool)`

GetLoggedAtOk returns a tuple with the LoggedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggedAt

`func (o *GetServerStatistics200ResponseRamInner) SetLoggedAt(v time.Time)`

SetLoggedAt sets LoggedAt field to given value.


### GetTotal

`func (o *GetServerStatistics200ResponseRamInner) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetServerStatistics200ResponseRamInner) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetServerStatistics200ResponseRamInner) SetTotal(v float32)`

SetTotal sets Total field to given value.


### GetUsed

`func (o *GetServerStatistics200ResponseRamInner) GetUsed() float32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *GetServerStatistics200ResponseRamInner) GetUsedOk() (*float32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *GetServerStatistics200ResponseRamInner) SetUsed(v float32)`

SetUsed sets Used field to given value.


### GetUsedCached

`func (o *GetServerStatistics200ResponseRamInner) GetUsedCached() float32`

GetUsedCached returns the UsedCached field if non-nil, zero value otherwise.

### GetUsedCachedOk

`func (o *GetServerStatistics200ResponseRamInner) GetUsedCachedOk() (*float32, bool)`

GetUsedCachedOk returns a tuple with the UsedCached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCached

`func (o *GetServerStatistics200ResponseRamInner) SetUsedCached(v float32)`

SetUsedCached sets UsedCached field to given value.


### GetAvailable

`func (o *GetServerStatistics200ResponseRamInner) GetAvailable() float32`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *GetServerStatistics200ResponseRamInner) GetAvailableOk() (*float32, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *GetServerStatistics200ResponseRamInner) SetAvailable(v float32)`

SetAvailable sets Available field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


