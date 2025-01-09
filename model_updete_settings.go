/*
Timeweb Cloud API

# Введение API Timeweb Cloud позволяет вам управлять ресурсами в облаке программным способом с использованием обычных HTTP-запросов.  Множество функций, которые доступны в панели управления Timeweb Cloud, также доступны через API, что позволяет вам автоматизировать ваши собственные сценарии.  В этой документации сперва будет описан общий дизайн и принципы работы API, а после этого конкретные конечные точки. Также будут приведены примеры запросов к ним.   ## Запросы Запросы должны выполняться по протоколу `HTTPS`, чтобы гарантировать шифрование транзакций. Поддерживаются следующие методы запроса: |Метод|Применение| |--- |--- | |GET|Извлекает данные о коллекциях и отдельных ресурсах.| |POST|Для коллекций создает новый ресурс этого типа. Также используется для выполнения действий с конкретным ресурсом.| |PUT|Обновляет существующий ресурс.| |PATCH|Некоторые ресурсы поддерживают частичное обновление, то есть обновление только части атрибутов ресурса, в этом случае вместо метода PUT будет использован PATCH.| |DELETE|Удаляет ресурс.|  Методы `POST`, `PUT` и `PATCH` могут включать объект в тело запроса с типом содержимого `application/json`.  ### Параметры в запросах Некоторые коллекции поддерживают пагинацию, поиск или сортировку в запросах. В параметрах запроса требуется передать: - `limit` — обозначает количество записей, которое необходимо вернуть  - `offset` — указывает на смещение, относительно начала списка  - `search` — позволяет указать набор символов для поиска  - `sort` — можно задать правило сортировки коллекции  ## Ответы Запросы вернут один из следующих кодов состояния ответа HTTP:  |Статус|Описание| |--- |--- | |200 OK|Действие с ресурсом было выполнено успешно.| |201 Created|Ресурс был успешно создан. При этом ресурс может быть как уже готовым к использованию, так и находиться в процессе запуска.| |204 No Content|Действие с ресурсом было выполнено успешно, и ответ не содержит дополнительной информации в теле.| |400 Bad Request|Был отправлен неверный запрос, например, в нем отсутствуют обязательные параметры и т. д. Тело ответа будет содержать дополнительную информацию об ошибке.| |401 Unauthorized|Ошибка аутентификации.| |403 Forbidden|Аутентификация прошла успешно, но недостаточно прав для выполнения действия.| |404 Not Found|Запрашиваемый ресурс не найден.| |409 Conflict|Запрос конфликтует с текущим состоянием.| |423 Locked|Ресурс из запроса заблокирован от применения к нему указанного метода.| |429 Too Many Requests|Был достигнут лимит по количеству запросов в единицу времени.| |500 Internal Server Error|При выполнении запроса произошла какая-то внутренняя ошибка. Чтобы решить эту проблему, лучше всего создать тикет в панели управления.|  ### Структура успешного ответа Все конечные точки будут возвращать данные в формате `JSON`. Ответы на `GET`-запросы будут иметь на верхнем уровне следующую структуру атрибутов:  |Название поля|Тип|Описание| |--- |--- |--- | |[entity_name]|object, object[], string[], number[], boolean|Динамическое поле, которое будет меняться в зависимости от запрашиваемого ресурса и будет содержать все атрибуты, необходимые для описания этого ресурса. Например, при запросе списка баз данных будет возвращаться поле `dbs`, а при запросе конкретного облачного сервера `server`. Для некоторых конечных точек в ответе может возвращаться сразу несколько ресурсов.| |meta|object|Опционально. Объект, который содержит вспомогательную информацию о ресурсе. Чаще всего будет встречаться при запросе коллекций и содержать поле `total`, которое будет указывать на количество элементов в коллекции.| |response_id|string|Опционально. В большинстве случаев в ответе будет содержаться ID ответа в формате UUIDv4, который однозначно указывает на ваш запрос внутри нашей системы. Если вам потребуется задать вопрос нашей поддержке, приложите к вопросу этот ID— так мы сможем найти ответ на него намного быстрее. Также вы можете использовать этот ID, чтобы убедиться, что это новый ответ на запрос и результат не был получен из кэша.|  Пример запроса на получение списка SSH-ключей: ```     HTTP/2.0 200 OK     {       \"ssh_keys\":[           {             \"body\":\"ssh-rsa AAAAB3NzaC1sdfghjkOAsBwWhs= example@device.local\",             \"created_at\":\"2021-09-15T19:52:27Z\",             \"expired_at\":null,             \"id\":5297,             \"is_default\":false,             \"name\":\"example@device.local\",             \"used_at\":null,             \"used_by\":[]           }       ],       \"meta\":{           \"total\":1       },       \"response_id\":\"94608d15-8672-4eed-8ab6-28bd6fa3cdf7\"     } ```  ### Структура ответа с ошибкой |Название поля|Тип|Описание| |--- |--- |--- | |status_code|number|Короткий числовой идентификатор ошибки.| |error_code|string|Короткий текстовый идентификатор ошибки, который уточняет числовой идентификатор и удобен для программной обработки. Самый простой пример — это код `not_found` для ошибки 404.| |message|string, string[]|Опционально. В большинстве случаев в ответе будет содержаться человекочитаемое подробное описание ошибки или ошибок, которые помогут понять, что нужно исправить.| |response_id|string|Опционально. В большинстве случае в ответе будет содержаться ID ответа в формате UUIDv4, который однозначно указывает на ваш запрос внутри нашей системы. Если вам потребуется задать вопрос нашей поддержке, приложите к вопросу этот ID — так мы сможем найти ответ на него намного быстрее.|  Пример: ```     HTTP/2.0 403 Forbidden     {       \"status_code\": 403,       \"error_code\":  \"forbidden\",       \"message\":     \"You do not have access for the attempted action\",       \"response_id\": \"94608d15-8672-4eed-8ab6-28bd6fa3cdf7\"     } ```  ## Статусы ресурсов Важно учесть, что при создании большинства ресурсов внутри платформы вам будет сразу возвращен ответ от сервера со статусом `200 OK` или `201 Created` и ID созданного ресурса в теле ответа, но при этом этот ресурс может быть ещё в *состоянии запуска*.  Для того чтобы понять, в каком состоянии сейчас находится ваш ресурс, мы добавили поле `status` в ответ на получение информации о ресурсе.  Список статусов будет отличаться в зависимости от типа ресурса. Увидеть поддерживаемый список статусов вы сможете в описании каждого конкретного ресурса.     ## Ограничение скорости запросов (Rate Limiting) Чтобы обеспечить стабильность для всех пользователей, Timeweb Cloud защищает API от всплесков входящего трафика, анализируя количество запросов c каждого аккаунта к каждой конечной точке.  Если ваше приложение отправляет более 20 запросов в секунду на одну конечную точку, то для этого запроса API может вернуть код состояния HTTP `429 Too Many Requests`.   ## Аутентификация Доступ к API осуществляется с помощью JWT-токена. Токенами можно управлять внутри панели управления Timeweb Cloud в разделе *API и Terraform*.  Токен необходимо передавать в заголовке каждого запроса в формате: ```   Authorization: Bearer $TIMEWEB_CLOUD_TOKEN ```  ## Формат примеров API Примеры в этой документации описаны с помощью `curl`, HTTP-клиента командной строки. На компьютерах `Linux` и `macOS` обычно по умолчанию установлен `curl`, и он доступен для загрузки на всех популярных платформах, включая `Windows`.  Каждый пример разделен на несколько строк символом `\\`, который совместим с `bash`. Типичный пример выглядит так: ```   curl -X PATCH      -H \"Content-Type: application/json\"      -H \"Authorization: Bearer $TIMEWEB_CLOUD_TOKEN\"      -d '{\"name\":\"Cute Corvus\",\"comment\":\"Development Server\"}'      \"https://api.timeweb.cloud/api/v1/dedicated/1051\" ``` - Параметр `-X` задает метод запроса. Для согласованности метод будет указан во всех примерах, даже если он явно не требуется для методов `GET`. - Строки `-H` задают требуемые HTTP-заголовки. - Примеры, для которых требуется объект JSON в теле запроса, передают требуемые данные через параметр `-d`.  Чтобы использовать приведенные примеры, не подставляя каждый раз в них свой токен, вы можете добавить токен один раз в переменные окружения в вашей консоли. Например, на `Linux` это можно сделать с помощью команды:  ``` TIMEWEB_CLOUD_TOKEN=\"token\" ```  После этого токен будет автоматически подставляться в ваши запросы.  Обратите внимание, что все значения в этой документации являются примерами. Не полагайтесь на IDы операционных систем, тарифов и т.д., используемые в примерах. Используйте соответствующую конечную точку для получения значений перед созданием ресурсов.   ## Версионирование API построено согласно принципам [семантического версионирования](https://semver.org/lang/ru). Это значит, что мы гарантируем обратную совместимость всех изменений в пределах одной мажорной версии.  Мажорная версия каждой конечной точки обозначается в пути запроса, например, запрос `/api/v1/servers` указывает, что этот метод имеет версию 1.

API version: 1.0.0
Contact: info@timeweb.cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UpdeteSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdeteSettings{}

// UpdeteSettings struct for UpdeteSettings
type UpdeteSettings struct {
	// Автоматический деплой.
	IsAutoDeploy *bool `json:"is_auto_deploy,omitempty"`
	// Команда сборки приложения.
	BuildCmd *string `json:"build_cmd,omitempty"`
	// Переменные окружения приложения. Объект с ключами и значениями типа string.
	Envs map[string]interface{} `json:"envs,omitempty"`
	// Название ветки репозитория из которой необходимо собрать приложение.
	BranchName *string `json:"branch_name,omitempty"`
	// Хэш коммита.
	CommitSha *string `json:"commit_sha,omitempty"`
	// Версия окружения.
	EnvVersion *string `json:"env_version,omitempty"`
	// Путь к директории с индексным файлом. Используется для приложений `type: frontend`. Не используется для приложений `type: backend`. Значение всегда должно начинаться с `/`.
	IndexDir *string `json:"index_dir,omitempty"`
	// Команда для запуска приложения. Используется для приложений `type: backend`. Не используется для приложений `type: frontend`.
	RunCmd *string `json:"run_cmd,omitempty"`
	Framework *Frameworks `json:"framework,omitempty"`
	// Имя приложения.
	Name *string `json:"name,omitempty"`
	// Комментарий к приложению.
	Comment *string `json:"comment,omitempty"`
	// ID тарифа.
	PresetId *float32 `json:"preset_id,omitempty"`
}

// NewUpdeteSettings instantiates a new UpdeteSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdeteSettings() *UpdeteSettings {
	this := UpdeteSettings{}
	return &this
}

// NewUpdeteSettingsWithDefaults instantiates a new UpdeteSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdeteSettingsWithDefaults() *UpdeteSettings {
	this := UpdeteSettings{}
	return &this
}

// GetIsAutoDeploy returns the IsAutoDeploy field value if set, zero value otherwise.
func (o *UpdeteSettings) GetIsAutoDeploy() bool {
	if o == nil || IsNil(o.IsAutoDeploy) {
		var ret bool
		return ret
	}
	return *o.IsAutoDeploy
}

// GetIsAutoDeployOk returns a tuple with the IsAutoDeploy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdeteSettings) GetIsAutoDeployOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAutoDeploy) {
		return nil, false
	}
	return o.IsAutoDeploy, true
}

// HasIsAutoDeploy returns a boolean if a field has been set.
func (o *UpdeteSettings) HasIsAutoDeploy() bool {
	if o != nil && !IsNil(o.IsAutoDeploy) {
		return true
	}

	return false
}

// SetIsAutoDeploy gets a reference to the given bool and assigns it to the IsAutoDeploy field.
func (o *UpdeteSettings) SetIsAutoDeploy(v bool) {
	o.IsAutoDeploy = &v
}

// GetBuildCmd returns the BuildCmd field value if set, zero value otherwise.
func (o *UpdeteSettings) GetBuildCmd() string {
	if o == nil || IsNil(o.BuildCmd) {
		var ret string
		return ret
	}
	return *o.BuildCmd
}

// GetBuildCmdOk returns a tuple with the BuildCmd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdeteSettings) GetBuildCmdOk() (*string, bool) {
	if o == nil || IsNil(o.BuildCmd) {
		return nil, false
	}
	return o.BuildCmd, true
}

// HasBuildCmd returns a boolean if a field has been set.
func (o *UpdeteSettings) HasBuildCmd() bool {
	if o != nil && !IsNil(o.BuildCmd) {
		return true
	}

	return false
}

// SetBuildCmd gets a reference to the given string and assigns it to the BuildCmd field.
func (o *UpdeteSettings) SetBuildCmd(v string) {
	o.BuildCmd = &v
}

// GetEnvs returns the Envs field value if set, zero value otherwise.
func (o *UpdeteSettings) GetEnvs() map[string]interface{} {
	if o == nil || IsNil(o.Envs) {
		var ret map[string]interface{}
		return ret
	}
	return o.Envs
}

// GetEnvsOk returns a tuple with the Envs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdeteSettings) GetEnvsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Envs) {
		return map[string]interface{}{}, false
	}
	return o.Envs, true
}

// HasEnvs returns a boolean if a field has been set.
func (o *UpdeteSettings) HasEnvs() bool {
	if o != nil && !IsNil(o.Envs) {
		return true
	}

	return false
}

// SetEnvs gets a reference to the given map[string]interface{} and assigns it to the Envs field.
func (o *UpdeteSettings) SetEnvs(v map[string]interface{}) {
	o.Envs = v
}

// GetBranchName returns the BranchName field value if set, zero value otherwise.
func (o *UpdeteSettings) GetBranchName() string {
	if o == nil || IsNil(o.BranchName) {
		var ret string
		return ret
	}
	return *o.BranchName
}

// GetBranchNameOk returns a tuple with the BranchName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdeteSettings) GetBranchNameOk() (*string, bool) {
	if o == nil || IsNil(o.BranchName) {
		return nil, false
	}
	return o.BranchName, true
}

// HasBranchName returns a boolean if a field has been set.
func (o *UpdeteSettings) HasBranchName() bool {
	if o != nil && !IsNil(o.BranchName) {
		return true
	}

	return false
}

// SetBranchName gets a reference to the given string and assigns it to the BranchName field.
func (o *UpdeteSettings) SetBranchName(v string) {
	o.BranchName = &v
}

// GetCommitSha returns the CommitSha field value if set, zero value otherwise.
func (o *UpdeteSettings) GetCommitSha() string {
	if o == nil || IsNil(o.CommitSha) {
		var ret string
		return ret
	}
	return *o.CommitSha
}

// GetCommitShaOk returns a tuple with the CommitSha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdeteSettings) GetCommitShaOk() (*string, bool) {
	if o == nil || IsNil(o.CommitSha) {
		return nil, false
	}
	return o.CommitSha, true
}

// HasCommitSha returns a boolean if a field has been set.
func (o *UpdeteSettings) HasCommitSha() bool {
	if o != nil && !IsNil(o.CommitSha) {
		return true
	}

	return false
}

// SetCommitSha gets a reference to the given string and assigns it to the CommitSha field.
func (o *UpdeteSettings) SetCommitSha(v string) {
	o.CommitSha = &v
}

// GetEnvVersion returns the EnvVersion field value if set, zero value otherwise.
func (o *UpdeteSettings) GetEnvVersion() string {
	if o == nil || IsNil(o.EnvVersion) {
		var ret string
		return ret
	}
	return *o.EnvVersion
}

// GetEnvVersionOk returns a tuple with the EnvVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdeteSettings) GetEnvVersionOk() (*string, bool) {
	if o == nil || IsNil(o.EnvVersion) {
		return nil, false
	}
	return o.EnvVersion, true
}

// HasEnvVersion returns a boolean if a field has been set.
func (o *UpdeteSettings) HasEnvVersion() bool {
	if o != nil && !IsNil(o.EnvVersion) {
		return true
	}

	return false
}

// SetEnvVersion gets a reference to the given string and assigns it to the EnvVersion field.
func (o *UpdeteSettings) SetEnvVersion(v string) {
	o.EnvVersion = &v
}

// GetIndexDir returns the IndexDir field value if set, zero value otherwise.
func (o *UpdeteSettings) GetIndexDir() string {
	if o == nil || IsNil(o.IndexDir) {
		var ret string
		return ret
	}
	return *o.IndexDir
}

// GetIndexDirOk returns a tuple with the IndexDir field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdeteSettings) GetIndexDirOk() (*string, bool) {
	if o == nil || IsNil(o.IndexDir) {
		return nil, false
	}
	return o.IndexDir, true
}

// HasIndexDir returns a boolean if a field has been set.
func (o *UpdeteSettings) HasIndexDir() bool {
	if o != nil && !IsNil(o.IndexDir) {
		return true
	}

	return false
}

// SetIndexDir gets a reference to the given string and assigns it to the IndexDir field.
func (o *UpdeteSettings) SetIndexDir(v string) {
	o.IndexDir = &v
}

// GetRunCmd returns the RunCmd field value if set, zero value otherwise.
func (o *UpdeteSettings) GetRunCmd() string {
	if o == nil || IsNil(o.RunCmd) {
		var ret string
		return ret
	}
	return *o.RunCmd
}

// GetRunCmdOk returns a tuple with the RunCmd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdeteSettings) GetRunCmdOk() (*string, bool) {
	if o == nil || IsNil(o.RunCmd) {
		return nil, false
	}
	return o.RunCmd, true
}

// HasRunCmd returns a boolean if a field has been set.
func (o *UpdeteSettings) HasRunCmd() bool {
	if o != nil && !IsNil(o.RunCmd) {
		return true
	}

	return false
}

// SetRunCmd gets a reference to the given string and assigns it to the RunCmd field.
func (o *UpdeteSettings) SetRunCmd(v string) {
	o.RunCmd = &v
}

// GetFramework returns the Framework field value if set, zero value otherwise.
func (o *UpdeteSettings) GetFramework() Frameworks {
	if o == nil || IsNil(o.Framework) {
		var ret Frameworks
		return ret
	}
	return *o.Framework
}

// GetFrameworkOk returns a tuple with the Framework field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdeteSettings) GetFrameworkOk() (*Frameworks, bool) {
	if o == nil || IsNil(o.Framework) {
		return nil, false
	}
	return o.Framework, true
}

// HasFramework returns a boolean if a field has been set.
func (o *UpdeteSettings) HasFramework() bool {
	if o != nil && !IsNil(o.Framework) {
		return true
	}

	return false
}

// SetFramework gets a reference to the given Frameworks and assigns it to the Framework field.
func (o *UpdeteSettings) SetFramework(v Frameworks) {
	o.Framework = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdeteSettings) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdeteSettings) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdeteSettings) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdeteSettings) SetName(v string) {
	o.Name = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *UpdeteSettings) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdeteSettings) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *UpdeteSettings) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *UpdeteSettings) SetComment(v string) {
	o.Comment = &v
}

// GetPresetId returns the PresetId field value if set, zero value otherwise.
func (o *UpdeteSettings) GetPresetId() float32 {
	if o == nil || IsNil(o.PresetId) {
		var ret float32
		return ret
	}
	return *o.PresetId
}

// GetPresetIdOk returns a tuple with the PresetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdeteSettings) GetPresetIdOk() (*float32, bool) {
	if o == nil || IsNil(o.PresetId) {
		return nil, false
	}
	return o.PresetId, true
}

// HasPresetId returns a boolean if a field has been set.
func (o *UpdeteSettings) HasPresetId() bool {
	if o != nil && !IsNil(o.PresetId) {
		return true
	}

	return false
}

// SetPresetId gets a reference to the given float32 and assigns it to the PresetId field.
func (o *UpdeteSettings) SetPresetId(v float32) {
	o.PresetId = &v
}

func (o UpdeteSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdeteSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsAutoDeploy) {
		toSerialize["is_auto_deploy"] = o.IsAutoDeploy
	}
	if !IsNil(o.BuildCmd) {
		toSerialize["build_cmd"] = o.BuildCmd
	}
	if !IsNil(o.Envs) {
		toSerialize["envs"] = o.Envs
	}
	if !IsNil(o.BranchName) {
		toSerialize["branch_name"] = o.BranchName
	}
	if !IsNil(o.CommitSha) {
		toSerialize["commit_sha"] = o.CommitSha
	}
	if !IsNil(o.EnvVersion) {
		toSerialize["env_version"] = o.EnvVersion
	}
	if !IsNil(o.IndexDir) {
		toSerialize["index_dir"] = o.IndexDir
	}
	if !IsNil(o.RunCmd) {
		toSerialize["run_cmd"] = o.RunCmd
	}
	if !IsNil(o.Framework) {
		toSerialize["framework"] = o.Framework
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.PresetId) {
		toSerialize["preset_id"] = o.PresetId
	}
	return toSerialize, nil
}

type NullableUpdeteSettings struct {
	value *UpdeteSettings
	isSet bool
}

func (v NullableUpdeteSettings) Get() *UpdeteSettings {
	return v.value
}

func (v *NullableUpdeteSettings) Set(val *UpdeteSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdeteSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdeteSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdeteSettings(val *UpdeteSettings) *NullableUpdeteSettings {
	return &NullableUpdeteSettings{value: val, isSet: true}
}

func (v NullableUpdeteSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdeteSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


