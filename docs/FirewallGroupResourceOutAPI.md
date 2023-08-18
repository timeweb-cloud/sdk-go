# FirewallGroupResourceOutAPI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | resource id | 
**Type** | [**ResourceType**](ResourceType.md) |  | 

## Methods

### NewFirewallGroupResourceOutAPI

`func NewFirewallGroupResourceOutAPI(id int32, type_ ResourceType, ) *FirewallGroupResourceOutAPI`

NewFirewallGroupResourceOutAPI instantiates a new FirewallGroupResourceOutAPI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallGroupResourceOutAPIWithDefaults

`func NewFirewallGroupResourceOutAPIWithDefaults() *FirewallGroupResourceOutAPI`

NewFirewallGroupResourceOutAPIWithDefaults instantiates a new FirewallGroupResourceOutAPI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FirewallGroupResourceOutAPI) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FirewallGroupResourceOutAPI) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FirewallGroupResourceOutAPI) SetId(v int32)`

SetId sets Id field to given value.


### GetType

`func (o *FirewallGroupResourceOutAPI) GetType() ResourceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FirewallGroupResourceOutAPI) GetTypeOk() (*ResourceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FirewallGroupResourceOutAPI) SetType(v ResourceType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


