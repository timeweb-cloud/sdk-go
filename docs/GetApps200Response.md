# GetApps200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**Apps** | [**[]App**](App.md) |  | 

## Methods

### NewGetApps200Response

`func NewGetApps200Response(meta Meta, apps []App, ) *GetApps200Response`

NewGetApps200Response instantiates a new GetApps200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApps200ResponseWithDefaults

`func NewGetApps200ResponseWithDefaults() *GetApps200Response`

NewGetApps200ResponseWithDefaults instantiates a new GetApps200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetApps200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetApps200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetApps200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetApps

`func (o *GetApps200Response) GetApps() []App`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *GetApps200Response) GetAppsOk() (*[]App, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *GetApps200Response) SetApps(v []App)`

SetApps sets Apps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


