# GetServerStatistics200ResponseNetworkTrafficInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoggedAt** | **time.Time** | Дата события в формате ISO 8061 | 
**Incoming** | **float32** | Входящий трафик сети в Мб/с | 
**Outgoing** | **float32** | Исходящий трафик сети в Мб/с | 

## Methods

### NewGetServerStatistics200ResponseNetworkTrafficInner

`func NewGetServerStatistics200ResponseNetworkTrafficInner(loggedAt time.Time, incoming float32, outgoing float32, ) *GetServerStatistics200ResponseNetworkTrafficInner`

NewGetServerStatistics200ResponseNetworkTrafficInner instantiates a new GetServerStatistics200ResponseNetworkTrafficInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServerStatistics200ResponseNetworkTrafficInnerWithDefaults

`func NewGetServerStatistics200ResponseNetworkTrafficInnerWithDefaults() *GetServerStatistics200ResponseNetworkTrafficInner`

NewGetServerStatistics200ResponseNetworkTrafficInnerWithDefaults instantiates a new GetServerStatistics200ResponseNetworkTrafficInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoggedAt

`func (o *GetServerStatistics200ResponseNetworkTrafficInner) GetLoggedAt() time.Time`

GetLoggedAt returns the LoggedAt field if non-nil, zero value otherwise.

### GetLoggedAtOk

`func (o *GetServerStatistics200ResponseNetworkTrafficInner) GetLoggedAtOk() (*time.Time, bool)`

GetLoggedAtOk returns a tuple with the LoggedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggedAt

`func (o *GetServerStatistics200ResponseNetworkTrafficInner) SetLoggedAt(v time.Time)`

SetLoggedAt sets LoggedAt field to given value.


### GetIncoming

`func (o *GetServerStatistics200ResponseNetworkTrafficInner) GetIncoming() float32`

GetIncoming returns the Incoming field if non-nil, zero value otherwise.

### GetIncomingOk

`func (o *GetServerStatistics200ResponseNetworkTrafficInner) GetIncomingOk() (*float32, bool)`

GetIncomingOk returns a tuple with the Incoming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncoming

`func (o *GetServerStatistics200ResponseNetworkTrafficInner) SetIncoming(v float32)`

SetIncoming sets Incoming field to given value.


### GetOutgoing

`func (o *GetServerStatistics200ResponseNetworkTrafficInner) GetOutgoing() float32`

GetOutgoing returns the Outgoing field if non-nil, zero value otherwise.

### GetOutgoingOk

`func (o *GetServerStatistics200ResponseNetworkTrafficInner) GetOutgoingOk() (*float32, bool)`

GetOutgoingOk returns a tuple with the Outgoing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoing

`func (o *GetServerStatistics200ResponseNetworkTrafficInner) SetOutgoing(v float32)`

SetOutgoing sets Outgoing field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


