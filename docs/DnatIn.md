# DnatIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | **string** | Протокол | [default to "tcp_udp"]
**Local** | [**DnatInLocal**](DnatInLocal.md) |  | 
**Public** | [**DnatInPublic**](DnatInPublic.md) |  | 

## Methods

### NewDnatIn

`func NewDnatIn(protocol string, local DnatInLocal, public DnatInPublic, ) *DnatIn`

NewDnatIn instantiates a new DnatIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnatInWithDefaults

`func NewDnatInWithDefaults() *DnatIn`

NewDnatInWithDefaults instantiates a new DnatIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *DnatIn) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *DnatIn) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *DnatIn) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetLocal

`func (o *DnatIn) GetLocal() DnatInLocal`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *DnatIn) GetLocalOk() (*DnatInLocal, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *DnatIn) SetLocal(v DnatInLocal)`

SetLocal sets Local field to given value.


### GetPublic

`func (o *DnatIn) GetPublic() DnatInPublic`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *DnatIn) GetPublicOk() (*DnatInPublic, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *DnatIn) SetPublic(v DnatInPublic)`

SetPublic sets Public field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


