# ProjectResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | ID для каждого ресурса проекта. Автоматически генерируется при создании. | 
**CreatedAt** | **string** | Значение времени, указанное в комбинированном формате даты и времени ISO8601, которое представляет, когда был создан ресурс. | 
**ResourceId** | **float32** | ID ресурса проекта (сервера, хранилища, кластера, балансировщика, базы данных или выделенного сервера). | 
**Project** | [**Project**](Project.md) |  | 
**Type** | **string** | Тип ресурса проекта | 

## Methods

### NewProjectResource

`func NewProjectResource(id float32, createdAt string, resourceId float32, project Project, type_ string, ) *ProjectResource`

NewProjectResource instantiates a new ProjectResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectResourceWithDefaults

`func NewProjectResourceWithDefaults() *ProjectResource`

NewProjectResourceWithDefaults instantiates a new ProjectResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectResource) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectResource) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectResource) SetId(v float32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ProjectResource) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProjectResource) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProjectResource) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetResourceId

`func (o *ProjectResource) GetResourceId() float32`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ProjectResource) GetResourceIdOk() (*float32, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ProjectResource) SetResourceId(v float32)`

SetResourceId sets ResourceId field to given value.


### GetProject

`func (o *ProjectResource) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ProjectResource) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ProjectResource) SetProject(v Project)`

SetProject sets Project field to given value.


### GetType

`func (o *ProjectResource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProjectResource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProjectResource) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


