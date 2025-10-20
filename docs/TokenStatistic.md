# TokenStatistic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | **time.Time** | Время точки статистики | 
**IngoingTokens** | **float32** | Количество входящих токенов (промпты) | 
**OutgoingTokens** | **float32** | Количество исходящих токенов (ответы) | 

## Methods

### NewTokenStatistic

`func NewTokenStatistic(time time.Time, ingoingTokens float32, outgoingTokens float32, ) *TokenStatistic`

NewTokenStatistic instantiates a new TokenStatistic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenStatisticWithDefaults

`func NewTokenStatisticWithDefaults() *TokenStatistic`

NewTokenStatisticWithDefaults instantiates a new TokenStatistic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *TokenStatistic) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *TokenStatistic) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *TokenStatistic) SetTime(v time.Time)`

SetTime sets Time field to given value.


### GetIngoingTokens

`func (o *TokenStatistic) GetIngoingTokens() float32`

GetIngoingTokens returns the IngoingTokens field if non-nil, zero value otherwise.

### GetIngoingTokensOk

`func (o *TokenStatistic) GetIngoingTokensOk() (*float32, bool)`

GetIngoingTokensOk returns a tuple with the IngoingTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngoingTokens

`func (o *TokenStatistic) SetIngoingTokens(v float32)`

SetIngoingTokens sets IngoingTokens field to given value.


### GetOutgoingTokens

`func (o *TokenStatistic) GetOutgoingTokens() float32`

GetOutgoingTokens returns the OutgoingTokens field if non-nil, zero value otherwise.

### GetOutgoingTokensOk

`func (o *TokenStatistic) GetOutgoingTokensOk() (*float32, bool)`

GetOutgoingTokensOk returns a tuple with the OutgoingTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingTokens

`func (o *TokenStatistic) SetOutgoingTokens(v float32)`

SetOutgoingTokens sets OutgoingTokens field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


