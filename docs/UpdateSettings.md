# UpdateSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAutoDeploy** | Pointer to **bool** | Автоматический деплой. | [optional] 
**BuildCmd** | Pointer to **string** | Команда сборки приложения. | [optional] 
**Envs** | Pointer to **map[string]interface{}** | Переменные окружения приложения. Объект с ключами и значениями типа string. | [optional] 
**BranchName** | Pointer to **string** | Название ветки репозитория из которой необходимо собрать приложение. | [optional] 
**CommitSha** | Pointer to **string** | Хэш коммита. | [optional] 
**EnvVersion** | Pointer to **string** | Версия окружения. | [optional] 
**IndexDir** | Pointer to **string** | Путь к директории с индексным файлом. Используется для приложений &#x60;type: frontend&#x60;. Не используется для приложений &#x60;type: backend&#x60;. Значение всегда должно начинаться с &#x60;/&#x60;. | [optional] 
**RunCmd** | Pointer to **string** | Команда для запуска приложения. Используется для приложений &#x60;type: backend&#x60;. Не используется для приложений &#x60;type: frontend&#x60;. | [optional] 
**Framework** | Pointer to [**Frameworks**](Frameworks.md) |  | [optional] 
**Name** | Pointer to **string** | Имя приложения. | [optional] 
**Comment** | Pointer to **string** | Комментарий к приложению. | [optional] 
**PresetId** | Pointer to **float32** | ID тарифа. | [optional] 

## Methods

### NewUpdateSettings

`func NewUpdateSettings() *UpdateSettings`

NewUpdateSettings instantiates a new UpdateSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSettingsWithDefaults

`func NewUpdateSettingsWithDefaults() *UpdateSettings`

NewUpdateSettingsWithDefaults instantiates a new UpdateSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAutoDeploy

`func (o *UpdateSettings) GetIsAutoDeploy() bool`

GetIsAutoDeploy returns the IsAutoDeploy field if non-nil, zero value otherwise.

### GetIsAutoDeployOk

`func (o *UpdateSettings) GetIsAutoDeployOk() (*bool, bool)`

GetIsAutoDeployOk returns a tuple with the IsAutoDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoDeploy

`func (o *UpdateSettings) SetIsAutoDeploy(v bool)`

SetIsAutoDeploy sets IsAutoDeploy field to given value.

### HasIsAutoDeploy

`func (o *UpdateSettings) HasIsAutoDeploy() bool`

HasIsAutoDeploy returns a boolean if a field has been set.

### GetBuildCmd

`func (o *UpdateSettings) GetBuildCmd() string`

GetBuildCmd returns the BuildCmd field if non-nil, zero value otherwise.

### GetBuildCmdOk

`func (o *UpdateSettings) GetBuildCmdOk() (*string, bool)`

GetBuildCmdOk returns a tuple with the BuildCmd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildCmd

`func (o *UpdateSettings) SetBuildCmd(v string)`

SetBuildCmd sets BuildCmd field to given value.

### HasBuildCmd

`func (o *UpdateSettings) HasBuildCmd() bool`

HasBuildCmd returns a boolean if a field has been set.

### GetEnvs

`func (o *UpdateSettings) GetEnvs() map[string]interface{}`

GetEnvs returns the Envs field if non-nil, zero value otherwise.

### GetEnvsOk

`func (o *UpdateSettings) GetEnvsOk() (*map[string]interface{}, bool)`

GetEnvsOk returns a tuple with the Envs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvs

`func (o *UpdateSettings) SetEnvs(v map[string]interface{})`

SetEnvs sets Envs field to given value.

### HasEnvs

`func (o *UpdateSettings) HasEnvs() bool`

HasEnvs returns a boolean if a field has been set.

### GetBranchName

`func (o *UpdateSettings) GetBranchName() string`

GetBranchName returns the BranchName field if non-nil, zero value otherwise.

### GetBranchNameOk

`func (o *UpdateSettings) GetBranchNameOk() (*string, bool)`

GetBranchNameOk returns a tuple with the BranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchName

`func (o *UpdateSettings) SetBranchName(v string)`

SetBranchName sets BranchName field to given value.

### HasBranchName

`func (o *UpdateSettings) HasBranchName() bool`

HasBranchName returns a boolean if a field has been set.

### GetCommitSha

`func (o *UpdateSettings) GetCommitSha() string`

GetCommitSha returns the CommitSha field if non-nil, zero value otherwise.

### GetCommitShaOk

`func (o *UpdateSettings) GetCommitShaOk() (*string, bool)`

GetCommitShaOk returns a tuple with the CommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitSha

`func (o *UpdateSettings) SetCommitSha(v string)`

SetCommitSha sets CommitSha field to given value.

### HasCommitSha

`func (o *UpdateSettings) HasCommitSha() bool`

HasCommitSha returns a boolean if a field has been set.

### GetEnvVersion

`func (o *UpdateSettings) GetEnvVersion() string`

GetEnvVersion returns the EnvVersion field if non-nil, zero value otherwise.

### GetEnvVersionOk

`func (o *UpdateSettings) GetEnvVersionOk() (*string, bool)`

GetEnvVersionOk returns a tuple with the EnvVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVersion

`func (o *UpdateSettings) SetEnvVersion(v string)`

SetEnvVersion sets EnvVersion field to given value.

### HasEnvVersion

`func (o *UpdateSettings) HasEnvVersion() bool`

HasEnvVersion returns a boolean if a field has been set.

### GetIndexDir

`func (o *UpdateSettings) GetIndexDir() string`

GetIndexDir returns the IndexDir field if non-nil, zero value otherwise.

### GetIndexDirOk

`func (o *UpdateSettings) GetIndexDirOk() (*string, bool)`

GetIndexDirOk returns a tuple with the IndexDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexDir

`func (o *UpdateSettings) SetIndexDir(v string)`

SetIndexDir sets IndexDir field to given value.

### HasIndexDir

`func (o *UpdateSettings) HasIndexDir() bool`

HasIndexDir returns a boolean if a field has been set.

### GetRunCmd

`func (o *UpdateSettings) GetRunCmd() string`

GetRunCmd returns the RunCmd field if non-nil, zero value otherwise.

### GetRunCmdOk

`func (o *UpdateSettings) GetRunCmdOk() (*string, bool)`

GetRunCmdOk returns a tuple with the RunCmd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunCmd

`func (o *UpdateSettings) SetRunCmd(v string)`

SetRunCmd sets RunCmd field to given value.

### HasRunCmd

`func (o *UpdateSettings) HasRunCmd() bool`

HasRunCmd returns a boolean if a field has been set.

### GetFramework

`func (o *UpdateSettings) GetFramework() Frameworks`

GetFramework returns the Framework field if non-nil, zero value otherwise.

### GetFrameworkOk

`func (o *UpdateSettings) GetFrameworkOk() (*Frameworks, bool)`

GetFrameworkOk returns a tuple with the Framework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramework

`func (o *UpdateSettings) SetFramework(v Frameworks)`

SetFramework sets Framework field to given value.

### HasFramework

`func (o *UpdateSettings) HasFramework() bool`

HasFramework returns a boolean if a field has been set.

### GetName

`func (o *UpdateSettings) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateSettings) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateSettings) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateSettings) HasName() bool`

HasName returns a boolean if a field has been set.

### GetComment

`func (o *UpdateSettings) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *UpdateSettings) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *UpdateSettings) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *UpdateSettings) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetPresetId

`func (o *UpdateSettings) GetPresetId() float32`

GetPresetId returns the PresetId field if non-nil, zero value otherwise.

### GetPresetIdOk

`func (o *UpdateSettings) GetPresetIdOk() (*float32, bool)`

GetPresetIdOk returns a tuple with the PresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetId

`func (o *UpdateSettings) SetPresetId(v float32)`

SetPresetId sets PresetId field to given value.

### HasPresetId

`func (o *UpdateSettings) HasPresetId() bool`

HasPresetId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


