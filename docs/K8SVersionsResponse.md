# K8SVersionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | Pointer to **string** | ID запроса | [optional] 
**Meta** | [**SchemasMeta**](SchemasMeta.md) |  | 
**K8sVersions** | **[]string** | Массив версий k8s | 

## Methods

### NewK8SVersionsResponse

`func NewK8SVersionsResponse(meta SchemasMeta, k8sVersions []string, ) *K8SVersionsResponse`

NewK8SVersionsResponse instantiates a new K8SVersionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK8SVersionsResponseWithDefaults

`func NewK8SVersionsResponseWithDefaults() *K8SVersionsResponse`

NewK8SVersionsResponseWithDefaults instantiates a new K8SVersionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *K8SVersionsResponse) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *K8SVersionsResponse) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *K8SVersionsResponse) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.

### HasResponseId

`func (o *K8SVersionsResponse) HasResponseId() bool`

HasResponseId returns a boolean if a field has been set.

### GetMeta

`func (o *K8SVersionsResponse) GetMeta() SchemasMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *K8SVersionsResponse) GetMetaOk() (*SchemasMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *K8SVersionsResponse) SetMeta(v SchemasMeta)`

SetMeta sets Meta field to given value.


### GetK8sVersions

`func (o *K8SVersionsResponse) GetK8sVersions() []string`

GetK8sVersions returns the K8sVersions field if non-nil, zero value otherwise.

### GetK8sVersionsOk

`func (o *K8SVersionsResponse) GetK8sVersionsOk() (*[]string, bool)`

GetK8sVersionsOk returns a tuple with the K8sVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sVersions

`func (o *K8SVersionsResponse) SetK8sVersions(v []string)`

SetK8sVersions sets K8sVersions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


