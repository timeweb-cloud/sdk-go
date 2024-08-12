# GetDeploySettings200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultDeploySettings** | Pointer to [**[]DeploySettingsInner**](DeploySettingsInner.md) |  | [optional] 
**Meta** | Pointer to [**Meta**](Meta.md) |  | [optional] 

## Methods

### NewGetDeploySettings200Response

`func NewGetDeploySettings200Response() *GetDeploySettings200Response`

NewGetDeploySettings200Response instantiates a new GetDeploySettings200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeploySettings200ResponseWithDefaults

`func NewGetDeploySettings200ResponseWithDefaults() *GetDeploySettings200Response`

NewGetDeploySettings200ResponseWithDefaults instantiates a new GetDeploySettings200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultDeploySettings

`func (o *GetDeploySettings200Response) GetDefaultDeploySettings() []DeploySettingsInner`

GetDefaultDeploySettings returns the DefaultDeploySettings field if non-nil, zero value otherwise.

### GetDefaultDeploySettingsOk

`func (o *GetDeploySettings200Response) GetDefaultDeploySettingsOk() (*[]DeploySettingsInner, bool)`

GetDefaultDeploySettingsOk returns a tuple with the DefaultDeploySettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDeploySettings

`func (o *GetDeploySettings200Response) SetDefaultDeploySettings(v []DeploySettingsInner)`

SetDefaultDeploySettings sets DefaultDeploySettings field to given value.

### HasDefaultDeploySettings

`func (o *GetDeploySettings200Response) HasDefaultDeploySettings() bool`

HasDefaultDeploySettings returns a boolean if a field has been set.

### GetMeta

`func (o *GetDeploySettings200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetDeploySettings200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetDeploySettings200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetDeploySettings200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


