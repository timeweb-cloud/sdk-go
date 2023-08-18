# DatabaseClusterNetworksInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Тип сети. | 
**Ips** | [**[]DatabaseClusterNetworksInnerIpsInner**](DatabaseClusterNetworksInnerIpsInner.md) | Список IP-адресов сети. | 

## Methods

### NewDatabaseClusterNetworksInner

`func NewDatabaseClusterNetworksInner(type_ string, ips []DatabaseClusterNetworksInnerIpsInner, ) *DatabaseClusterNetworksInner`

NewDatabaseClusterNetworksInner instantiates a new DatabaseClusterNetworksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseClusterNetworksInnerWithDefaults

`func NewDatabaseClusterNetworksInnerWithDefaults() *DatabaseClusterNetworksInner`

NewDatabaseClusterNetworksInnerWithDefaults instantiates a new DatabaseClusterNetworksInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DatabaseClusterNetworksInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DatabaseClusterNetworksInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DatabaseClusterNetworksInner) SetType(v string)`

SetType sets Type field to given value.


### GetIps

`func (o *DatabaseClusterNetworksInner) GetIps() []DatabaseClusterNetworksInnerIpsInner`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *DatabaseClusterNetworksInner) GetIpsOk() (*[]DatabaseClusterNetworksInnerIpsInner, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *DatabaseClusterNetworksInner) SetIps(v []DatabaseClusterNetworksInnerIpsInner)`

SetIps sets Ips field to given value.


### SetIpsNil

`func (o *DatabaseClusterNetworksInner) SetIpsNil(b bool)`

 SetIpsNil sets the value for Ips to be an explicit nil

### UnsetIps
`func (o *DatabaseClusterNetworksInner) UnsetIps()`

UnsetIps ensures that no value is present for Ips, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


