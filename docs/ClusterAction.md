# ClusterAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Действие над кластером базы данных: &#x60;reboot&#x60; — перезагрузить, &#x60;shutdown&#x60; — выключить, &#x60;start&#x60; — включить. | 

## Methods

### NewClusterAction

`func NewClusterAction(action string, ) *ClusterAction`

NewClusterAction instantiates a new ClusterAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterActionWithDefaults

`func NewClusterActionWithDefaults() *ClusterAction`

NewClusterActionWithDefaults instantiates a new ClusterAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ClusterAction) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ClusterAction) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ClusterAction) SetAction(v string)`

SetAction sets Action field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


