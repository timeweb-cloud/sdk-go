# GetPersons200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**Persons** | [**[]Person**](Person.md) |  | 

## Methods

### NewGetPersons200Response

`func NewGetPersons200Response(meta Meta, persons []Person, ) *GetPersons200Response`

NewGetPersons200Response instantiates a new GetPersons200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPersons200ResponseWithDefaults

`func NewGetPersons200ResponseWithDefaults() *GetPersons200Response`

NewGetPersons200ResponseWithDefaults instantiates a new GetPersons200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetPersons200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetPersons200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetPersons200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetPersons

`func (o *GetPersons200Response) GetPersons() []Person`

GetPersons returns the Persons field if non-nil, zero value otherwise.

### GetPersonsOk

`func (o *GetPersons200Response) GetPersonsOk() (*[]Person, bool)`

GetPersonsOk returns a tuple with the Persons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersons

`func (o *GetPersons200Response) SetPersons(v []Person)`

SetPersons sets Persons field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


