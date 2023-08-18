# DeleteIPsFromAllowedListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ips** | **[]string** | Список удаляемых из списка разрешенных IP-адресов. | 

## Methods

### NewDeleteIPsFromAllowedListRequest

`func NewDeleteIPsFromAllowedListRequest(ips []string, ) *DeleteIPsFromAllowedListRequest`

NewDeleteIPsFromAllowedListRequest instantiates a new DeleteIPsFromAllowedListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteIPsFromAllowedListRequestWithDefaults

`func NewDeleteIPsFromAllowedListRequestWithDefaults() *DeleteIPsFromAllowedListRequest`

NewDeleteIPsFromAllowedListRequestWithDefaults instantiates a new DeleteIPsFromAllowedListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIps

`func (o *DeleteIPsFromAllowedListRequest) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *DeleteIPsFromAllowedListRequest) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *DeleteIPsFromAllowedListRequest) SetIps(v []string)`

SetIps sets Ips field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


