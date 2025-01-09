# ServerLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | ID диска. | 
**LoggedAt** | **time.Time** | Дата лога. | 
**Event** | **string** | Событие сервера. | 

## Methods

### NewServerLog

`func NewServerLog(id float32, loggedAt time.Time, event string, ) *ServerLog`

NewServerLog instantiates a new ServerLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerLogWithDefaults

`func NewServerLogWithDefaults() *ServerLog`

NewServerLogWithDefaults instantiates a new ServerLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServerLog) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerLog) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerLog) SetId(v float32)`

SetId sets Id field to given value.


### GetLoggedAt

`func (o *ServerLog) GetLoggedAt() time.Time`

GetLoggedAt returns the LoggedAt field if non-nil, zero value otherwise.

### GetLoggedAtOk

`func (o *ServerLog) GetLoggedAtOk() (*time.Time, bool)`

GetLoggedAtOk returns a tuple with the LoggedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggedAt

`func (o *ServerLog) SetLoggedAt(v time.Time)`

SetLoggedAt sets LoggedAt field to given value.


### GetEvent

`func (o *ServerLog) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *ServerLog) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *ServerLog) SetEvent(v string)`

SetEvent sets Event field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


