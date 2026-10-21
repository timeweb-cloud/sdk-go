# ChatsCsv

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Csv** | **string** | Содержимое CSV-файла: по строке на чат и на каждое его сообщение (ID диалога, ID агента, ID сообщения, время, автор, текст, оценка, время ответа в секундах) | 
**FileName** | **string** | Рекомендуемое имя файла (без расширения) | 

## Methods

### NewChatsCsv

`func NewChatsCsv(csv string, fileName string, ) *ChatsCsv`

NewChatsCsv instantiates a new ChatsCsv object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatsCsvWithDefaults

`func NewChatsCsvWithDefaults() *ChatsCsv`

NewChatsCsvWithDefaults instantiates a new ChatsCsv object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCsv

`func (o *ChatsCsv) GetCsv() string`

GetCsv returns the Csv field if non-nil, zero value otherwise.

### GetCsvOk

`func (o *ChatsCsv) GetCsvOk() (*string, bool)`

GetCsvOk returns a tuple with the Csv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsv

`func (o *ChatsCsv) SetCsv(v string)`

SetCsv sets Csv field to given value.


### GetFileName

`func (o *ChatsCsv) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *ChatsCsv) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *ChatsCsv) SetFileName(v string)`

SetFileName sets FileName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


