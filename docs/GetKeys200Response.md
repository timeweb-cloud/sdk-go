# GetKeys200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**SshKeys** | [**[]SshKey**](SshKey.md) |  | 

## Methods

### NewGetKeys200Response

`func NewGetKeys200Response(meta Meta, sshKeys []SshKey, ) *GetKeys200Response`

NewGetKeys200Response instantiates a new GetKeys200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetKeys200ResponseWithDefaults

`func NewGetKeys200ResponseWithDefaults() *GetKeys200Response`

NewGetKeys200ResponseWithDefaults instantiates a new GetKeys200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetKeys200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetKeys200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetKeys200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetSshKeys

`func (o *GetKeys200Response) GetSshKeys() []SshKey`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *GetKeys200Response) GetSshKeysOk() (*[]SshKey, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *GetKeys200Response) SetSshKeys(v []SshKey)`

SetSshKeys sets SshKeys field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


