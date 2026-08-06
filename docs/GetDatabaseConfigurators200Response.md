# GetDatabaseConfigurators200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**DatabaseConfigurators** | [**[]DatabaseConfigurator**](DatabaseConfigurator.md) |  | 

## Methods

### NewGetDatabaseConfigurators200Response

`func NewGetDatabaseConfigurators200Response(meta Meta, databaseConfigurators []DatabaseConfigurator, ) *GetDatabaseConfigurators200Response`

NewGetDatabaseConfigurators200Response instantiates a new GetDatabaseConfigurators200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDatabaseConfigurators200ResponseWithDefaults

`func NewGetDatabaseConfigurators200ResponseWithDefaults() *GetDatabaseConfigurators200Response`

NewGetDatabaseConfigurators200ResponseWithDefaults instantiates a new GetDatabaseConfigurators200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetDatabaseConfigurators200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetDatabaseConfigurators200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetDatabaseConfigurators200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetDatabaseConfigurators

`func (o *GetDatabaseConfigurators200Response) GetDatabaseConfigurators() []DatabaseConfigurator`

GetDatabaseConfigurators returns the DatabaseConfigurators field if non-nil, zero value otherwise.

### GetDatabaseConfiguratorsOk

`func (o *GetDatabaseConfigurators200Response) GetDatabaseConfiguratorsOk() (*[]DatabaseConfigurator, bool)`

GetDatabaseConfiguratorsOk returns a tuple with the DatabaseConfigurators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseConfigurators

`func (o *GetDatabaseConfigurators200Response) SetDatabaseConfigurators(v []DatabaseConfigurator)`

SetDatabaseConfigurators sets DatabaseConfigurators field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


