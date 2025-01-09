# UpdeteSettings

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

### NewUpdeteSettings

`func NewUpdeteSettings() *UpdeteSettings`

NewUpdeteSettings instantiates a new UpdeteSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdeteSettingsWithDefaults

`func NewUpdeteSettingsWithDefaults() *UpdeteSettings`

NewUpdeteSettingsWithDefaults instantiates a new UpdeteSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAutoDeploy

`func (o *UpdeteSettings) GetIsAutoDeploy() bool`

GetIsAutoDeploy returns the IsAutoDeploy field if non-nil, zero value otherwise.

### GetIsAutoDeployOk

`func (o *UpdeteSettings) GetIsAutoDeployOk() (*bool, bool)`

GetIsAutoDeployOk returns a tuple with the IsAutoDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoDeploy

`func (o *UpdeteSettings) SetIsAutoDeploy(v bool)`

SetIsAutoDeploy sets IsAutoDeploy field to given value.

### HasIsAutoDeploy

`func (o *UpdeteSettings) HasIsAutoDeploy() bool`

HasIsAutoDeploy returns a boolean if a field has been set.

### GetBuildCmd

`func (o *UpdeteSettings) GetBuildCmd() string`

GetBuildCmd returns the BuildCmd field if non-nil, zero value otherwise.

### GetBuildCmdOk

`func (o *UpdeteSettings) GetBuildCmdOk() (*string, bool)`

GetBuildCmdOk returns a tuple with the BuildCmd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildCmd

`func (o *UpdeteSettings) SetBuildCmd(v string)`

SetBuildCmd sets BuildCmd field to given value.

### HasBuildCmd

`func (o *UpdeteSettings) HasBuildCmd() bool`

HasBuildCmd returns a boolean if a field has been set.

### GetEnvs

`func (o *UpdeteSettings) GetEnvs() map[string]interface{}`

GetEnvs returns the Envs field if non-nil, zero value otherwise.

### GetEnvsOk

`func (o *UpdeteSettings) GetEnvsOk() (*map[string]interface{}, bool)`

GetEnvsOk returns a tuple with the Envs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvs

`func (o *UpdeteSettings) SetEnvs(v map[string]interface{})`

SetEnvs sets Envs field to given value.

### HasEnvs

`func (o *UpdeteSettings) HasEnvs() bool`

HasEnvs returns a boolean if a field has been set.

### GetBranchName

`func (o *UpdeteSettings) GetBranchName() string`

GetBranchName returns the BranchName field if non-nil, zero value otherwise.

### GetBranchNameOk

`func (o *UpdeteSettings) GetBranchNameOk() (*string, bool)`

GetBranchNameOk returns a tuple with the BranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchName

`func (o *UpdeteSettings) SetBranchName(v string)`

SetBranchName sets BranchName field to given value.

### HasBranchName

`func (o *UpdeteSettings) HasBranchName() bool`

HasBranchName returns a boolean if a field has been set.

### GetCommitSha

`func (o *UpdeteSettings) GetCommitSha() string`

GetCommitSha returns the CommitSha field if non-nil, zero value otherwise.

### GetCommitShaOk

`func (o *UpdeteSettings) GetCommitShaOk() (*string, bool)`

GetCommitShaOk returns a tuple with the CommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitSha

`func (o *UpdeteSettings) SetCommitSha(v string)`

SetCommitSha sets CommitSha field to given value.

### HasCommitSha

`func (o *UpdeteSettings) HasCommitSha() bool`

HasCommitSha returns a boolean if a field has been set.

### GetEnvVersion

`func (o *UpdeteSettings) GetEnvVersion() string`

GetEnvVersion returns the EnvVersion field if non-nil, zero value otherwise.

### GetEnvVersionOk

`func (o *UpdeteSettings) GetEnvVersionOk() (*string, bool)`

GetEnvVersionOk returns a tuple with the EnvVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVersion

`func (o *UpdeteSettings) SetEnvVersion(v string)`

SetEnvVersion sets EnvVersion field to given value.

### HasEnvVersion

`func (o *UpdeteSettings) HasEnvVersion() bool`

HasEnvVersion returns a boolean if a field has been set.

### GetIndexDir

`func (o *UpdeteSettings) GetIndexDir() string`

GetIndexDir returns the IndexDir field if non-nil, zero value otherwise.

### GetIndexDirOk

`func (o *UpdeteSettings) GetIndexDirOk() (*string, bool)`

GetIndexDirOk returns a tuple with the IndexDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexDir

`func (o *UpdeteSettings) SetIndexDir(v string)`

SetIndexDir sets IndexDir field to given value.

### HasIndexDir

`func (o *UpdeteSettings) HasIndexDir() bool`

HasIndexDir returns a boolean if a field has been set.

### GetRunCmd

`func (o *UpdeteSettings) GetRunCmd() string`

GetRunCmd returns the RunCmd field if non-nil, zero value otherwise.

### GetRunCmdOk

`func (o *UpdeteSettings) GetRunCmdOk() (*string, bool)`

GetRunCmdOk returns a tuple with the RunCmd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunCmd

`func (o *UpdeteSettings) SetRunCmd(v string)`

SetRunCmd sets RunCmd field to given value.

### HasRunCmd

`func (o *UpdeteSettings) HasRunCmd() bool`

HasRunCmd returns a boolean if a field has been set.

### GetFramework

`func (o *UpdeteSettings) GetFramework() Frameworks`

GetFramework returns the Framework field if non-nil, zero value otherwise.

### GetFrameworkOk

`func (o *UpdeteSettings) GetFrameworkOk() (*Frameworks, bool)`

GetFrameworkOk returns a tuple with the Framework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramework

`func (o *UpdeteSettings) SetFramework(v Frameworks)`

SetFramework sets Framework field to given value.

### HasFramework

`func (o *UpdeteSettings) HasFramework() bool`

HasFramework returns a boolean if a field has been set.

### GetName

`func (o *UpdeteSettings) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdeteSettings) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdeteSettings) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdeteSettings) HasName() bool`

HasName returns a boolean if a field has been set.

### GetComment

`func (o *UpdeteSettings) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *UpdeteSettings) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *UpdeteSettings) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *UpdeteSettings) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetPresetId

`func (o *UpdeteSettings) GetPresetId() float32`

GetPresetId returns the PresetId field if non-nil, zero value otherwise.

### GetPresetIdOk

`func (o *UpdeteSettings) GetPresetIdOk() (*float32, bool)`

GetPresetIdOk returns a tuple with the PresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetId

`func (o *UpdeteSettings) SetPresetId(v float32)`

SetPresetId sets PresetId field to given value.

### HasPresetId

`func (o *UpdeteSettings) HasPresetId() bool`

HasPresetId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


