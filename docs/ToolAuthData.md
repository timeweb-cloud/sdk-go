# ToolAuthData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Способ авторизации | 
**Data** | Pointer to [**ToolAuthDataParams**](ToolAuthDataParams.md) |  | [optional] 

## Methods

### NewToolAuthData

`func NewToolAuthData(type_ string, ) *ToolAuthData`

NewToolAuthData instantiates a new ToolAuthData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolAuthDataWithDefaults

`func NewToolAuthDataWithDefaults() *ToolAuthData`

NewToolAuthDataWithDefaults instantiates a new ToolAuthData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ToolAuthData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ToolAuthData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ToolAuthData) SetType(v string)`

SetType sets Type field to given value.


### GetData

`func (o *ToolAuthData) GetData() ToolAuthDataParams`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ToolAuthData) GetDataOk() (*ToolAuthDataParams, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ToolAuthData) SetData(v ToolAuthDataParams)`

SetData sets Data field to given value.

### HasData

`func (o *ToolAuthData) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


