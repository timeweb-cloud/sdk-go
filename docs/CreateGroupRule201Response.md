# CreateGroupRule201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 
**Rule** | [**FirewallRuleOutAPI**](FirewallRuleOutAPI.md) |  | 

## Methods

### NewCreateGroupRule201Response

`func NewCreateGroupRule201Response(responseId string, rule FirewallRuleOutAPI, ) *CreateGroupRule201Response`

NewCreateGroupRule201Response instantiates a new CreateGroupRule201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGroupRule201ResponseWithDefaults

`func NewCreateGroupRule201ResponseWithDefaults() *CreateGroupRule201Response`

NewCreateGroupRule201ResponseWithDefaults instantiates a new CreateGroupRule201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *CreateGroupRule201Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *CreateGroupRule201Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *CreateGroupRule201Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.


### GetRule

`func (o *CreateGroupRule201Response) GetRule() FirewallRuleOutAPI`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *CreateGroupRule201Response) GetRuleOk() (*FirewallRuleOutAPI, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *CreateGroupRule201Response) SetRule(v FirewallRuleOutAPI)`

SetRule sets Rule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


