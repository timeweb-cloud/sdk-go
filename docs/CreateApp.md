# CreateApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderId** | **string** | ID провайдера. | 
**Type** | **string** | Тип приложения. | 
**RepositoryId** | **string** | ID репозитория. | 
**BuildCmd** | **string** | Команда сборки приложения. | 
**Envs** | Pointer to **map[string]interface{}** | Переменные окружения приложения. Объект с ключами и значениями типа string. | [optional] 
**BranchName** | **string** | Название ветки репозитория из которой необходимо собрать приложение. | 
**IsAutoDeploy** | **bool** | Автоматический деплой. | 
**CommitSha** | **string** | Хэш коммита из которого необходимо собрать приложение. | 
**Name** | **string** | Имя приложения. | 
**Comment** | **string** | Комментарий к приложению. | 
**PresetId** | **float32** | ID тарифа. | 
**EnvVersion** | Pointer to **string** | Версия окружения. | [optional] 
**Framework** | [**Frameworks**](Frameworks.md) |  | 
**IndexDir** | Pointer to **string** | Путь к директории с индексным файлом. Обязателен для приложений &#x60;type: frontend&#x60;. Не используется для приложений &#x60;type: backend&#x60;. Значение всегда должно начинаться с &#x60;/&#x60;. | [optional] 
**RunCmd** | Pointer to **string** | Команда для запуска приложения. Обязательна для приложений &#x60;type: backend&#x60;. Не используется для приложений &#x60;type: frontend&#x60;. | [optional] 

## Methods

### NewCreateApp

`func NewCreateApp(providerId string, type_ string, repositoryId string, buildCmd string, branchName string, isAutoDeploy bool, commitSha string, name string, comment string, presetId float32, framework Frameworks, ) *CreateApp`

NewCreateApp instantiates a new CreateApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAppWithDefaults

`func NewCreateAppWithDefaults() *CreateApp`

NewCreateAppWithDefaults instantiates a new CreateApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderId

`func (o *CreateApp) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *CreateApp) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *CreateApp) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.


### GetType

`func (o *CreateApp) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateApp) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateApp) SetType(v string)`

SetType sets Type field to given value.


### GetRepositoryId

`func (o *CreateApp) GetRepositoryId() string`

GetRepositoryId returns the RepositoryId field if non-nil, zero value otherwise.

### GetRepositoryIdOk

`func (o *CreateApp) GetRepositoryIdOk() (*string, bool)`

GetRepositoryIdOk returns a tuple with the RepositoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryId

`func (o *CreateApp) SetRepositoryId(v string)`

SetRepositoryId sets RepositoryId field to given value.


### GetBuildCmd

`func (o *CreateApp) GetBuildCmd() string`

GetBuildCmd returns the BuildCmd field if non-nil, zero value otherwise.

### GetBuildCmdOk

`func (o *CreateApp) GetBuildCmdOk() (*string, bool)`

GetBuildCmdOk returns a tuple with the BuildCmd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildCmd

`func (o *CreateApp) SetBuildCmd(v string)`

SetBuildCmd sets BuildCmd field to given value.


### GetEnvs

`func (o *CreateApp) GetEnvs() map[string]interface{}`

GetEnvs returns the Envs field if non-nil, zero value otherwise.

### GetEnvsOk

`func (o *CreateApp) GetEnvsOk() (*map[string]interface{}, bool)`

GetEnvsOk returns a tuple with the Envs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvs

`func (o *CreateApp) SetEnvs(v map[string]interface{})`

SetEnvs sets Envs field to given value.

### HasEnvs

`func (o *CreateApp) HasEnvs() bool`

HasEnvs returns a boolean if a field has been set.

### GetBranchName

`func (o *CreateApp) GetBranchName() string`

GetBranchName returns the BranchName field if non-nil, zero value otherwise.

### GetBranchNameOk

`func (o *CreateApp) GetBranchNameOk() (*string, bool)`

GetBranchNameOk returns a tuple with the BranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchName

`func (o *CreateApp) SetBranchName(v string)`

SetBranchName sets BranchName field to given value.


### GetIsAutoDeploy

`func (o *CreateApp) GetIsAutoDeploy() bool`

GetIsAutoDeploy returns the IsAutoDeploy field if non-nil, zero value otherwise.

### GetIsAutoDeployOk

`func (o *CreateApp) GetIsAutoDeployOk() (*bool, bool)`

GetIsAutoDeployOk returns a tuple with the IsAutoDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoDeploy

`func (o *CreateApp) SetIsAutoDeploy(v bool)`

SetIsAutoDeploy sets IsAutoDeploy field to given value.


### GetCommitSha

`func (o *CreateApp) GetCommitSha() string`

GetCommitSha returns the CommitSha field if non-nil, zero value otherwise.

### GetCommitShaOk

`func (o *CreateApp) GetCommitShaOk() (*string, bool)`

GetCommitShaOk returns a tuple with the CommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitSha

`func (o *CreateApp) SetCommitSha(v string)`

SetCommitSha sets CommitSha field to given value.


### GetName

`func (o *CreateApp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateApp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateApp) SetName(v string)`

SetName sets Name field to given value.


### GetComment

`func (o *CreateApp) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateApp) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateApp) SetComment(v string)`

SetComment sets Comment field to given value.


### GetPresetId

`func (o *CreateApp) GetPresetId() float32`

GetPresetId returns the PresetId field if non-nil, zero value otherwise.

### GetPresetIdOk

`func (o *CreateApp) GetPresetIdOk() (*float32, bool)`

GetPresetIdOk returns a tuple with the PresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetId

`func (o *CreateApp) SetPresetId(v float32)`

SetPresetId sets PresetId field to given value.


### GetEnvVersion

`func (o *CreateApp) GetEnvVersion() string`

GetEnvVersion returns the EnvVersion field if non-nil, zero value otherwise.

### GetEnvVersionOk

`func (o *CreateApp) GetEnvVersionOk() (*string, bool)`

GetEnvVersionOk returns a tuple with the EnvVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVersion

`func (o *CreateApp) SetEnvVersion(v string)`

SetEnvVersion sets EnvVersion field to given value.

### HasEnvVersion

`func (o *CreateApp) HasEnvVersion() bool`

HasEnvVersion returns a boolean if a field has been set.

### GetFramework

`func (o *CreateApp) GetFramework() Frameworks`

GetFramework returns the Framework field if non-nil, zero value otherwise.

### GetFrameworkOk

`func (o *CreateApp) GetFrameworkOk() (*Frameworks, bool)`

GetFrameworkOk returns a tuple with the Framework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramework

`func (o *CreateApp) SetFramework(v Frameworks)`

SetFramework sets Framework field to given value.


### GetIndexDir

`func (o *CreateApp) GetIndexDir() string`

GetIndexDir returns the IndexDir field if non-nil, zero value otherwise.

### GetIndexDirOk

`func (o *CreateApp) GetIndexDirOk() (*string, bool)`

GetIndexDirOk returns a tuple with the IndexDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexDir

`func (o *CreateApp) SetIndexDir(v string)`

SetIndexDir sets IndexDir field to given value.

### HasIndexDir

`func (o *CreateApp) HasIndexDir() bool`

HasIndexDir returns a boolean if a field has been set.

### GetRunCmd

`func (o *CreateApp) GetRunCmd() string`

GetRunCmd returns the RunCmd field if non-nil, zero value otherwise.

### GetRunCmdOk

`func (o *CreateApp) GetRunCmdOk() (*string, bool)`

GetRunCmdOk returns a tuple with the RunCmd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunCmd

`func (o *CreateApp) SetRunCmd(v string)`

SetRunCmd sets RunCmd field to given value.

### HasRunCmd

`func (o *CreateApp) HasRunCmd() bool`

HasRunCmd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


