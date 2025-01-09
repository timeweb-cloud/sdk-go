# Deploy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | **string** | ID приложения. | 
**CommitSha** | **string** | Хэш коммита. | 
**Id** | **string** | ID. | 
**StartedAt** | **time.Time** | Время запуска деплоя. | 
**EndedAt** | **NullableTime** | Время окончания деплоя. Определено для завершенных деплоев | 
**Status** | [**DeployStatus**](DeployStatus.md) |  | 
**CommitMsg** | **string** | Сообщение коммита. | 

## Methods

### NewDeploy

`func NewDeploy(appId string, commitSha string, id string, startedAt time.Time, endedAt NullableTime, status DeployStatus, commitMsg string, ) *Deploy`

NewDeploy instantiates a new Deploy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployWithDefaults

`func NewDeployWithDefaults() *Deploy`

NewDeployWithDefaults instantiates a new Deploy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *Deploy) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *Deploy) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *Deploy) SetAppId(v string)`

SetAppId sets AppId field to given value.


### GetCommitSha

`func (o *Deploy) GetCommitSha() string`

GetCommitSha returns the CommitSha field if non-nil, zero value otherwise.

### GetCommitShaOk

`func (o *Deploy) GetCommitShaOk() (*string, bool)`

GetCommitShaOk returns a tuple with the CommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitSha

`func (o *Deploy) SetCommitSha(v string)`

SetCommitSha sets CommitSha field to given value.


### GetId

`func (o *Deploy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Deploy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Deploy) SetId(v string)`

SetId sets Id field to given value.


### GetStartedAt

`func (o *Deploy) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *Deploy) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *Deploy) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetEndedAt

`func (o *Deploy) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *Deploy) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *Deploy) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.


### SetEndedAtNil

`func (o *Deploy) SetEndedAtNil(b bool)`

 SetEndedAtNil sets the value for EndedAt to be an explicit nil

### UnsetEndedAt
`func (o *Deploy) UnsetEndedAt()`

UnsetEndedAt ensures that no value is present for EndedAt, not even an explicit nil
### GetStatus

`func (o *Deploy) GetStatus() DeployStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Deploy) GetStatusOk() (*DeployStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Deploy) SetStatus(v DeployStatus)`

SetStatus sets Status field to given value.


### GetCommitMsg

`func (o *Deploy) GetCommitMsg() string`

GetCommitMsg returns the CommitMsg field if non-nil, zero value otherwise.

### GetCommitMsgOk

`func (o *Deploy) GetCommitMsgOk() (*string, bool)`

GetCommitMsgOk returns a tuple with the CommitMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMsg

`func (o *Deploy) SetCommitMsg(v string)`

SetCommitMsg sets CommitMsg field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


