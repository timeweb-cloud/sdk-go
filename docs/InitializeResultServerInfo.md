# InitializeResultServerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Имя сервера: &#x60;timeweb-cloud-mcp&#x60; на &#x60;/api/v1/mcp&#x60;, &#x60;timeweb-cloud-mcp-search&#x60; на &#x60;/api/v1/mcp/search&#x60; | 
**Version** | **string** | Версия сервера | 

## Methods

### NewInitializeResultServerInfo

`func NewInitializeResultServerInfo(name string, version string, ) *InitializeResultServerInfo`

NewInitializeResultServerInfo instantiates a new InitializeResultServerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitializeResultServerInfoWithDefaults

`func NewInitializeResultServerInfoWithDefaults() *InitializeResultServerInfo`

NewInitializeResultServerInfoWithDefaults instantiates a new InitializeResultServerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InitializeResultServerInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InitializeResultServerInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InitializeResultServerInfo) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *InitializeResultServerInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *InitializeResultServerInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *InitializeResultServerInfo) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


