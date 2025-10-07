# ClusterEdit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Новое название кластера | [optional] 
**Description** | Pointer to **string** | Новое описание кластера | [optional] 
**OidcProvider** | Pointer to [**ClusterEditOidcProvider**](ClusterEditOidcProvider.md) |  | [optional] 

## Methods

### NewClusterEdit

`func NewClusterEdit() *ClusterEdit`

NewClusterEdit instantiates a new ClusterEdit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterEditWithDefaults

`func NewClusterEditWithDefaults() *ClusterEdit`

NewClusterEditWithDefaults instantiates a new ClusterEdit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClusterEdit) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterEdit) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterEdit) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterEdit) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ClusterEdit) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClusterEdit) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClusterEdit) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClusterEdit) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOidcProvider

`func (o *ClusterEdit) GetOidcProvider() ClusterEditOidcProvider`

GetOidcProvider returns the OidcProvider field if non-nil, zero value otherwise.

### GetOidcProviderOk

`func (o *ClusterEdit) GetOidcProviderOk() (*ClusterEditOidcProvider, bool)`

GetOidcProviderOk returns a tuple with the OidcProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcProvider

`func (o *ClusterEdit) SetOidcProvider(v ClusterEditOidcProvider)`

SetOidcProvider sets OidcProvider field to given value.

### HasOidcProvider

`func (o *ClusterEdit) HasOidcProvider() bool`

HasOidcProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


