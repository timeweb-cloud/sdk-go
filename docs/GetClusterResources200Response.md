# GetClusterResources200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 
**Resources** | [**Resources**](Resources.md) |  | 

## Methods

### NewGetClusterResources200Response

`func NewGetClusterResources200Response(responseId string, resources Resources, ) *GetClusterResources200Response`

NewGetClusterResources200Response instantiates a new GetClusterResources200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetClusterResources200ResponseWithDefaults

`func NewGetClusterResources200ResponseWithDefaults() *GetClusterResources200Response`

NewGetClusterResources200ResponseWithDefaults instantiates a new GetClusterResources200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *GetClusterResources200Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *GetClusterResources200Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *GetClusterResources200Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.


### GetResources

`func (o *GetClusterResources200Response) GetResources() Resources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *GetClusterResources200Response) GetResourcesOk() (*Resources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *GetClusterResources200Response) SetResources(v Resources)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


