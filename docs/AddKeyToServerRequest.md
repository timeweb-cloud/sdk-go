# AddKeyToServerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SshKeyIds** | **[]float32** | Массив ID SSH-ключей | 

## Methods

### NewAddKeyToServerRequest

`func NewAddKeyToServerRequest(sshKeyIds []float32, ) *AddKeyToServerRequest`

NewAddKeyToServerRequest instantiates a new AddKeyToServerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddKeyToServerRequestWithDefaults

`func NewAddKeyToServerRequestWithDefaults() *AddKeyToServerRequest`

NewAddKeyToServerRequestWithDefaults instantiates a new AddKeyToServerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSshKeyIds

`func (o *AddKeyToServerRequest) GetSshKeyIds() []float32`

GetSshKeyIds returns the SshKeyIds field if non-nil, zero value otherwise.

### GetSshKeyIdsOk

`func (o *AddKeyToServerRequest) GetSshKeyIdsOk() (*[]float32, bool)`

GetSshKeyIdsOk returns a tuple with the SshKeyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyIds

`func (o *AddKeyToServerRequest) SetSshKeyIds(v []float32)`

SetSshKeyIds sets SshKeyIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


