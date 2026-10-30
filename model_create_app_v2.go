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

// checks if the CreateAppV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAppV2{}

// CreateAppV2 Параметры создания приложения. Для Git-репозитория обязательны provider_id, repository_id, branch_name, commit_sha, build_cmd, is_auto_deploy и framework. Для frontend также обязательны index_dir и preset_id. Для backend из Git обязателен run_cmd. Для backend из готового Docker-образа передайте docker_image и docker_image_tag; Git-параметры не требуются. Для любого backend передайте либо preset_id, либо configuration. Эти поля нельзя передавать одновременно.
type CreateAppV2 struct {
	// ID провайдера. Обязателен при создании из Git-репозитория. Не требуется для готового Docker-образа.
	ProviderId *string `json:"provider_id,omitempty"`
	// Тип приложения: frontend — статический сайт, backend — серверное приложение или Docker-образ.
	Type string `json:"type"`
	// ID репозитория. Обязателен при создании из Git-репозитория. Не требуется для готового Docker-образа.
	RepositoryId *string `json:"repository_id,omitempty"`
	// Команда сборки приложения из Git-репозитория. Обязательна для Git; можно передать пустую строку, если сборка не нужна.
	BuildCmd *string `json:"build_cmd,omitempty"`
	// Переменные окружения приложения. Ключи и значения — строки. Длина ключа — не более 255 символов, значения — не более 10240. По умолчанию пустой объект.
	Envs *map[string]string `json:"envs,omitempty"`
	// Ветка Git-репозитория, из которой необходимо собрать приложение. Обязательна при создании из Git.
	BranchName *string `json:"branch_name,omitempty"`
	// Автоматически запускать деплой при изменении Git-репозитория. Поле обязательно при создании из Git. Для провайдера git и готового Docker-образа автодеплой выключен.
	IsAutoDeploy *bool `json:"is_auto_deploy,omitempty"`
	// Хэш коммита для первого деплоя. Обязателен при создании из Git.
	CommitSha *string `json:"commit_sha,omitempty"`
	// Имя приложения, от 1 до 255 символов. Допустимы латинские и русские буквы, цифры, пробелы и символы - . ( ) , & # [ ] _.
	Name string `json:"name"`
	// Комментарий к приложению. Обязательное поле; можно передать пустую строку. Допустимы те же символы, что и в имени приложения.
	Comment string `json:"comment"`
	// ID тарифа приложения. Обязателен для frontend. Для backend передайте либо preset_id, либо configuration. Нельзя передавать вместе с configuration. Доступные тарифы: GET /api/v1/presets/apps.
	PresetId *int32 `json:"preset_id,omitempty"`
	// Версия окружения для сборки и запуска, например версия Node.js. Можно передать null.
	EnvVersion NullableString `json:"env_version,omitempty"`
	// Фреймворк приложения, строка с учётом регистра. Обязателен при создании из Git. Для frontend: `angular`, `astro`, `ember`, `next.js`, `nuxt`, `preact`, `react`, `remix`, `solid`, `solid-start`, `svelte`, `sveltekit`, `tanstack-start`, `vue`, `static-nobuild`, `another`. Для backend: `django`, `express`, `phoenix`, `Spring`, `laravel`, `beego`, `fastapi`, `ASP.NET Core`, `hapi`, `celery`, `flask`, `gin`, `docker`, `docker-compose`, `fastify`, `nest`, `symfony`, `yii`, `next.js`, `nuxt`, `astro`, `tanstack-start`, `solid-start`, `ktor`, `axum`, `actix`, `tornado`, `aiohttp`, `fiber`, `echo`, `Micronaut`, `Quarkus`, `remix`, `sveltekit`. Для готового Docker-образа не требуется.
	Framework *string `json:"framework,omitempty"`
	// Путь к каталогу с собранными статическими файлами, например /dist. Обязателен для frontend. Не используется для backend.
	IndexDir *string `json:"index_dir,omitempty"`
	// Команда запуска backend-приложения из Git. Обязательна для backend из Git. Не требуется для готового Docker-образа и frontend.
	RunCmd *string `json:"run_cmd,omitempty"`
	// Системные пакеты для установки в окружение приложения. Допустимы строчные латинские буквы, цифры и символы . + - =. По умолчанию пустой массив.
	SystemDependencies []string `json:"system_dependencies,omitempty"`
	// ID проекта, в который необходимо добавить приложение. Проект должен принадлежать аккаунту.
	ProjectId *int32 `json:"project_id,omitempty"`
	Configuration *CreateAppV2Configuration `json:"configuration,omitempty"`
	// Язык или способ развёртывания. Обязателен в v2. Для frontend: `javascript`. Для backend: `node`, `elixir`, `go`, `php`, `python`, `ruby`, `docker`, `docker-compose`, `dotnet`, `java`, `kotlin`, `rust`, `bun`, `deno`, `custom-image`. Для SSR-приложений из Git используйте node, для готового Docker-образа — custom-image.
	Language string `json:"language"`
	// Каталог приложения в Git-репозитории, относительно которого выполняется сборка. Непустой путь начинается с /, не содержит .. или двойной начальный /. Допустимы латинские буквы, цифры, . - _ /. Можно передать пустую строку.
	BuildEntryPath *string `json:"build_entry_path,omitempty"`
	Network *CreateAppV2Network `json:"network,omitempty"`
	// HTTP-путь для проверки работоспособности backend-приложения.
	HealthCheck *string `json:"health_check,omitempty"`
	// ID глобальных переменных окружения аккаунта, которые необходимо связать с приложением.
	GlobalEnvIds []float32 `json:"global_env_ids,omitempty"`
	// Имя готового Docker-образа для backend, без тега. Передавайте вместе с docker_image_tag. При указании обоих полей Git-параметры не требуются.
	DockerImage *string `json:"docker_image,omitempty"`
	// Тег готового Docker-образа. Передавайте вместе с docker_image.
	DockerImageTag *string `json:"docker_image_tag,omitempty"`
	// Добавить доверенные корневые сертификаты Минцифры в окружение backend-приложения. Применяется в российском регионе.
	IsRussianTrustedCa *bool `json:"is_russian_trusted_ca,omitempty"`
	// Добавить доверенные корневые сертификаты ТЦИ в окружение backend-приложения. Применяется в российском регионе.
	IsTciTrustedCa *bool `json:"is_tci_trusted_ca,omitempty"`
}

// NewCreateAppV2 instantiates a new CreateAppV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAppV2(type_ string, name string, comment string, language string) *CreateAppV2 {
	this := CreateAppV2{}
	this.Type = type_
	this.Name = name
	this.Comment = comment
	this.Language = language
	var isRussianTrustedCa bool = false
	this.IsRussianTrustedCa = &isRussianTrustedCa
	var isTciTrustedCa bool = false
	this.IsTciTrustedCa = &isTciTrustedCa
	return &this
}

// NewCreateAppV2WithDefaults instantiates a new CreateAppV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAppV2WithDefaults() *CreateAppV2 {
	this := CreateAppV2{}
	var isRussianTrustedCa bool = false
	this.IsRussianTrustedCa = &isRussianTrustedCa
	var isTciTrustedCa bool = false
	this.IsTciTrustedCa = &isTciTrustedCa
	return &this
}

// GetProviderId returns the ProviderId field value if set, zero value otherwise.
func (o *CreateAppV2) GetProviderId() string {
	if o == nil || IsNil(o.ProviderId) {
		var ret string
		return ret
	}
	return *o.ProviderId
}

// GetProviderIdOk returns a tuple with the ProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAppV2) GetProviderIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderId) {
		return nil, false
	}
	return o.ProviderId, true
}

// HasProviderId returns a boolean if a field has been set.
func (o *CreateAppV2) HasProviderId() bool {
	if o != nil && !IsNil(o.ProviderId) {
		return true
	}

	return false
}

// SetProviderId gets a reference to the given string and assigns it to the ProviderId field.
func (o *CreateAppV2) SetProviderId(v string) {
	o.ProviderId = &v
}

// GetType returns the Type field value
func (o *CreateAppV2) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateAppV2) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateAppV2) SetType(v string) {
	o.Type = v
}

// GetRepositoryId returns the RepositoryId field value if set, zero value otherwise.
func (o *CreateAppV2) GetRepositoryId() string {
	if o == nil || IsNil(o.RepositoryId) {
		var ret string
		return ret
	}
	return *o.RepositoryId
}

// GetRepositoryIdOk returns a tuple with the RepositoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAppV2) GetRepositoryIdOk() (*string, bool) {
	if o == nil || IsNil(o.RepositoryId) {
		return nil, false
	}
	return o.RepositoryId, true
}

// HasRepositoryId returns a boolean if a field has been set.
func (o *CreateAppV2) HasRepositoryId() bool {
	if o != nil && !IsNil(o.RepositoryId) {
		return true
	}

	return false
}

// SetRepositoryId gets a reference to the given string and assigns it to the RepositoryId field.
func (o *CreateAppV2) SetRepositoryId(v string) {
	o.RepositoryId = &v
}

// GetBuildCmd returns the BuildCmd field value if set, zero value otherwise.
func (o *CreateAppV2) GetBuildCmd() string {
	if o == nil || IsNil(o.BuildCmd) {
		var ret string
		return ret
	}
	return *o.BuildCmd
}

// GetBuildCmdOk returns a tuple with the BuildCmd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAppV2) GetBuildCmdOk() (*string, bool) {
	if o == nil || IsNil(o.BuildCmd) {
		return nil, false
	}
	return o.BuildCmd, true
}

// HasBuildCmd returns a boolean if a field has been set.
func (o *CreateAppV2) HasBuildCmd() bool {
	if o != nil && !IsNil(o.BuildCmd) {
		return true
	}

	return false
}

// SetBuildCmd gets a reference to the given string and assigns it to the BuildCmd field.
func (o *CreateAppV2) SetBuildCmd(v string) {
	o.BuildCmd = &v
}

// GetEnvs returns the Envs field value if set, zero value otherwise.
func (o *CreateAppV2) GetEnvs() map[string]string {
	if o == nil || IsNil(o.Envs) {
		var ret map[string]string
		return ret
	}
	return *o.Envs
}

// GetEnvsOk returns a tuple with the Envs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAppV2) GetEnvsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Envs) {
		return nil, false
	}
	return o.Envs, true
}

// HasEnvs returns a boolean if a field has been set.
func (o *CreateAppV2) HasEnvs() bool {
	if o != nil && !IsNil(o.Envs) {
		return true
	}

	return false
}

// SetEnvs gets a reference to the given map[string]string and assigns it to the Envs field.
func (o *CreateAppV2) SetEnvs(v map[string]string) {
	o.Envs = &v
}

// GetBranchName returns the BranchName field value if set, zero value otherwise.
func (o *CreateAppV2) GetBranchName() string {
	if o == nil || IsNil(o.BranchName) {
		var ret string
		return ret
	}
	return *o.BranchName
}

// GetBranchNameOk returns a tuple with the BranchName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAppV2) GetBranchNameOk() (*string, bool) {
	if o == nil || IsNil(o.BranchName) {
		return nil, false
	}
	return o.BranchName, true
}

// HasBranchName returns a boolean if a field has been set.
func (o *CreateAppV2) HasBranchName() bool {
	if o != nil && !IsNil(o.BranchName) {
		return true
	}

	return false
}

// SetBranchName gets a reference to the given string and assigns it to the BranchName field.
func (o *CreateAppV2) SetBranchName(v string) {
	o.BranchName = &v
}

// GetIsAutoDeploy returns the IsAutoDeploy field value if set, zero value otherwise.
func (o *CreateAppV2) GetIsAutoDeploy() bool {
	if o == nil || IsNil(o.IsAutoDeploy) {
		var ret bool
		return ret
	}
	return *o.IsAutoDeploy
}

// GetIsAutoDeployOk returns a tuple with the IsAutoDeploy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAppV2) GetIsAutoDeployOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAutoDeploy) {
		return nil, false
	}
	return o.IsAutoDeploy, true
}

// HasIsAutoDeploy returns a boolean if a field has been set.
func (o *CreateAppV2) HasIsAutoDeploy() bool {
	if o != nil && !IsNil(o.IsAutoDeploy) {
		return true
	}

	return false
}

// SetIsAutoDeploy gets a reference to the given bool and assigns it to the IsAutoDeploy field.
func (o *CreateAppV2) SetIsAutoDeploy(v bool) {
	o.IsAutoDeploy = &v
}

// GetCommitSha returns the CommitSha field value if set, zero value otherwise.
func (o *CreateAppV2) GetCommitSha() string {
	if o == nil || IsNil(o.CommitSha) {
		var ret string
		return ret
	}
	return *o.CommitSha
}

// GetCommitShaOk returns a tuple with the CommitSha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAppV2) GetCommitShaOk() (*string, bool) {
	if o == nil || IsNil(o.CommitSha) {
		return nil, false
	}
	return o.CommitSha, true
}

// HasCommitSha returns a boolean if a field has been set.
func (o *CreateAppV2) HasCommitSha() bool {
	if o != nil && !IsNil(o.CommitSha) {
		return true
	}

	return false
}

// SetCommitSha gets a reference to the given string and assigns it to the CommitSha field.
func (o *CreateAppV2) SetCommitSha(v string) {
	o.CommitSha = &v
}

// GetName returns the Name field value
func (o *CreateAppV2) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateAppV2) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateAppV2) SetName(v string) {
	o.Name = v
}

// GetComment returns the Comment field value
func (o *CreateAppV2) GetComment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Comment
}

// GetCommentOk returns a tuple with the Comment field value
// and a boolean to check if the value has been set.
func (o *CreateAppV2) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comment, true
}

// SetComment sets field value
func (o *CreateAppV2) SetComment(v string) {
	o.Comment = v
}

// GetPresetId returns the PresetId field value if set, zero value otherwise.
func (o *CreateAppV2) GetPresetId() int32 {
	if o == nil || IsNil(o.PresetId) {
		var ret int32
		return ret
	}
	return *o.PresetId
}

// GetPresetIdOk returns a tuple with the PresetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAppV2) GetPresetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.PresetId) {
		return nil, false
	}
	return o.PresetId, true
}

// HasPresetId returns a boolean if a field has been set.
func (o *CreateAppV2) HasPresetId() bool {
	if o != nil && !IsNil(o.PresetId) {
		return true
	}

	return false
}

// SetPresetId gets a reference to the given int32 and assigns it to the PresetId field.
func (o *CreateAppV2) SetPresetId(v int32) {
	o.PresetId = &v
}

// GetEnvVersion returns the EnvVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAppV2) GetEnvVersion() string {
	if o == nil || IsNil(o.EnvVersion.Get()) {
		var ret string
		return ret
	}
	return *o.EnvVersion.Get()
}

// GetEnvVersionOk returns a tuple with the EnvVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAppV2) GetEnvVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnvVersion.Get(), o.EnvVersion.IsSet()
}

// HasEnvVersion returns a boolean if a field has been set.
func (o *CreateAppV2) HasEnvVersion() bool {
	if o != nil && o.EnvVersion.IsSet() {
		return true
	}

	return false
}

// SetEnvVersion gets a reference to the given NullableString and assigns it to the EnvVersion field.
func (o *CreateAppV2) SetEnvVersion(v string) {
	o.EnvVersion.Set(&v)
}
// SetEnvVersionNil sets the value for EnvVersion to be an explicit nil
func (o *CreateAppV2) SetEnvVersionNil() {
	o.EnvVersion.Set(nil)
}

// UnsetEnvVersion ensures that no value is present for EnvVersion, not even an explicit nil
func (o *CreateAppV2) UnsetEnvVersion() {
	o.EnvVersion.Unset()
}

// GetFramework returns the Framework field value if set, zero value otherwise.
func (o *CreateAppV2) GetFramework() string {
	if o == nil || IsNil(o.Framework) {
		var ret string
		return ret
	}
	return *o.Framework
}

// GetFrameworkOk returns a tuple with the Framework field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAppV2) GetFrameworkOk() (*string, bool) {
	if o == nil || IsNil(o.Framework) {
		return nil, false
	}
	return o.Framework, true
}

// HasFramework returns a boolean if a field has been set.
func (o *CreateAppV2) HasFramework() bool {
	if o != nil && !IsNil(o.Framework) {
		return true
	}

	return false
}

// SetFramework gets a reference to the given string and assigns it to the Framework field.
func (o *CreateAppV2) SetFramework(v string) {
	o.Framework = &v
}

// GetIndexDir returns the IndexDir field value if set, zero value otherwise.
func (o *CreateAppV2) GetIndexDir() string {
	if o == nil || IsNil(o.IndexDir) {
		var ret string
		return ret
	}
	return *o.IndexDir
}

// GetIndexDirOk returns a tuple with the IndexDir field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAppV2) GetIndexDirOk() (*string, bool) {
	if o == nil || IsNil(o.IndexDir) {
		return nil, false
	}
	return o.IndexDir, true
}

// HasIndexDir returns a boolean if a field has been set.
func (o *CreateAppV2) HasIndexDir() bool {
	if o != nil && !IsNil(o.IndexDir) {
		return true
	}

	return false
}

// SetIndexDir gets a reference to the given string and assigns it to the IndexDir field.
func (o *CreateAppV2) SetIndexDir(v string) {
	o.IndexDir = &v
}

// GetRunCmd returns the RunCmd field value if set, zero value otherwise.
func (o *CreateAppV2) GetRunCmd() string {
	if o == nil || IsNil(o.RunCmd) {
		var ret string
		return ret
	}
	return *o.RunCmd
}

// GetRunCmdOk returns a tuple with the RunCmd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAppV2) GetRunCmdOk() (*string, bool) {
	if o == nil || IsNil(o.RunCmd) {
		return nil, false
	}
	return o.RunCmd, true
}

// HasRunCmd returns a boolean if a field has been set.
func (o *CreateAppV2) HasRunCmd() bool {
	if o != nil && !IsNil(o.RunCmd) {
		return true
	}

	return false
}

// SetRunCmd gets a reference to the given string and assigns it to the RunCmd field.
func (o *CreateAppV2) SetRunCmd(v string) {
	o.RunCmd = &v
}

// GetSystemDependencies returns the SystemDependencies field value if set, zero value otherwise.
func (o *CreateAppV2) GetSystemDependencies() []string {
	if o == nil || IsNil(o.SystemDependencies) {
		var ret []string
		return ret
	}
	return o.SystemDependencies
}

// GetSystemDependenciesOk returns a tuple with the SystemDependencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAppV2) GetSystemDependenciesOk() ([]string, bool) {
	if o == nil || IsNil(o.SystemDependencies) {
		return nil, false
	}
	return o.SystemDependencies, true
}

// HasSystemDependencies returns a boolean if a field has been set.
func (o *CreateAppV2) HasSystemDependencies() bool {
	if o != nil && !IsNil(o.SystemDependencies) {
		return true
	}

	return false
}

// SetSystemDependencies gets a reference to the given []string and assigns it to the SystemDependencies field.
func (o *CreateAppV2) SetSystemDependencies(v []string) {
	o.SystemDependencies = v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *CreateAppV2) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAppV2) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *CreateAppV2) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *CreateAppV2) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *CreateAppV2) GetConfiguration() CreateAppV2Configuration {
	if o == nil || IsNil(o.Configuration) {
		var ret CreateAppV2Configuration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAppV2) GetConfigurationOk() (*CreateAppV2Configuration, bool) {
	if o == nil || IsNil(o.Configuration) {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *CreateAppV2) HasConfiguration() bool {
	if o != nil && !IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given CreateAppV2Configuration and assigns it to the Configuration field.
func (o *CreateAppV2) SetConfiguration(v CreateAppV2Configuration) {
	o.Configuration = &v
}

// GetLanguage returns the Language field value
func (o *CreateAppV2) GetLanguage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *CreateAppV2) GetLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value
func (o *CreateAppV2) SetLanguage(v string) {
	o.Language = v
}

// GetBuildEntryPath returns the BuildEntryPath field value if set, zero value otherwise.
func (o *CreateAppV2) GetBuildEntryPath() string {
	if o == nil || IsNil(o.BuildEntryPath) {
		var ret string
		return ret
	}
	return *o.BuildEntryPath
}

// GetBuildEntryPathOk returns a tuple with the BuildEntryPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAppV2) GetBuildEntryPathOk() (*string, bool) {
	if o == nil || IsNil(o.BuildEntryPath) {
		return nil, false
	}
	return o.BuildEntryPath, true
}

// HasBuildEntryPath returns a boolean if a field has been set.
func (o *CreateAppV2) HasBuildEntryPath() bool {
	if o != nil && !IsNil(o.BuildEntryPath) {
		return true
	}

	return false
}

// SetBuildEntryPath gets a reference to the given string and assigns it to the BuildEntryPath field.
func (o *CreateAppV2) SetBuildEntryPath(v string) {
	o.BuildEntryPath = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *CreateAppV2) GetNetwork() CreateAppV2Network {
	if o == nil || IsNil(o.Network) {
		var ret CreateAppV2Network
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAppV2) GetNetworkOk() (*CreateAppV2Network, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *CreateAppV2) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given CreateAppV2Network and assigns it to the Network field.
func (o *CreateAppV2) SetNetwork(v CreateAppV2Network) {
	o.Network = &v
}

// GetHealthCheck returns the HealthCheck field value if set, zero value otherwise.
func (o *CreateAppV2) GetHealthCheck() string {
	if o == nil || IsNil(o.HealthCheck) {
		var ret string
		return ret
	}
	return *o.HealthCheck
}

// GetHealthCheckOk returns a tuple with the HealthCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAppV2) GetHealthCheckOk() (*string, bool) {
	if o == nil || IsNil(o.HealthCheck) {
		return nil, false
	}
	return o.HealthCheck, true
}

// HasHealthCheck returns a boolean if a field has been set.
func (o *CreateAppV2) HasHealthCheck() bool {
	if o != nil && !IsNil(o.HealthCheck) {
		return true
	}

	return false
}

// SetHealthCheck gets a reference to the given string and assigns it to the HealthCheck field.
func (o *CreateAppV2) SetHealthCheck(v string) {
	o.HealthCheck = &v
}

// GetGlobalEnvIds returns the GlobalEnvIds field value if set, zero value otherwise.
func (o *CreateAppV2) GetGlobalEnvIds() []float32 {
	if o == nil || IsNil(o.GlobalEnvIds) {
		var ret []float32
		return ret
	}
	return o.GlobalEnvIds
}

// GetGlobalEnvIdsOk returns a tuple with the GlobalEnvIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAppV2) GetGlobalEnvIdsOk() ([]float32, bool) {
	if o == nil || IsNil(o.GlobalEnvIds) {
		return nil, false
	}
	return o.GlobalEnvIds, true
}

// HasGlobalEnvIds returns a boolean if a field has been set.
func (o *CreateAppV2) HasGlobalEnvIds() bool {
	if o != nil && !IsNil(o.GlobalEnvIds) {
		return true
	}

	return false
}

// SetGlobalEnvIds gets a reference to the given []float32 and assigns it to the GlobalEnvIds field.
func (o *CreateAppV2) SetGlobalEnvIds(v []float32) {
	o.GlobalEnvIds = v
}

// GetDockerImage returns the DockerImage field value if set, zero value otherwise.
func (o *CreateAppV2) GetDockerImage() string {
	if o == nil || IsNil(o.DockerImage) {
		var ret string
		return ret
	}
	return *o.DockerImage
}

// GetDockerImageOk returns a tuple with the DockerImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAppV2) GetDockerImageOk() (*string, bool) {
	if o == nil || IsNil(o.DockerImage) {
		return nil, false
	}
	return o.DockerImage, true
}

// HasDockerImage returns a boolean if a field has been set.
func (o *CreateAppV2) HasDockerImage() bool {
	if o != nil && !IsNil(o.DockerImage) {
		return true
	}

	return false
}

// SetDockerImage gets a reference to the given string and assigns it to the DockerImage field.
func (o *CreateAppV2) SetDockerImage(v string) {
	o.DockerImage = &v
}

// GetDockerImageTag returns the DockerImageTag field value if set, zero value otherwise.
func (o *CreateAppV2) GetDockerImageTag() string {
	if o == nil || IsNil(o.DockerImageTag) {
		var ret string
		return ret
	}
	return *o.DockerImageTag
}

// GetDockerImageTagOk returns a tuple with the DockerImageTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAppV2) GetDockerImageTagOk() (*string, bool) {
	if o == nil || IsNil(o.DockerImageTag) {
		return nil, false
	}
	return o.DockerImageTag, true
}

// HasDockerImageTag returns a boolean if a field has been set.
func (o *CreateAppV2) HasDockerImageTag() bool {
	if o != nil && !IsNil(o.DockerImageTag) {
		return true
	}

	return false
}

// SetDockerImageTag gets a reference to the given string and assigns it to the DockerImageTag field.
func (o *CreateAppV2) SetDockerImageTag(v string) {
	o.DockerImageTag = &v
}

// GetIsRussianTrustedCa returns the IsRussianTrustedCa field value if set, zero value otherwise.
func (o *CreateAppV2) GetIsRussianTrustedCa() bool {
	if o == nil || IsNil(o.IsRussianTrustedCa) {
		var ret bool
		return ret
	}
	return *o.IsRussianTrustedCa
}

// GetIsRussianTrustedCaOk returns a tuple with the IsRussianTrustedCa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAppV2) GetIsRussianTrustedCaOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRussianTrustedCa) {
		return nil, false
	}
	return o.IsRussianTrustedCa, true
}

// HasIsRussianTrustedCa returns a boolean if a field has been set.
func (o *CreateAppV2) HasIsRussianTrustedCa() bool {
	if o != nil && !IsNil(o.IsRussianTrustedCa) {
		return true
	}

	return false
}

// SetIsRussianTrustedCa gets a reference to the given bool and assigns it to the IsRussianTrustedCa field.
func (o *CreateAppV2) SetIsRussianTrustedCa(v bool) {
	o.IsRussianTrustedCa = &v
}

// GetIsTciTrustedCa returns the IsTciTrustedCa field value if set, zero value otherwise.
func (o *CreateAppV2) GetIsTciTrustedCa() bool {
	if o == nil || IsNil(o.IsTciTrustedCa) {
		var ret bool
		return ret
	}
	return *o.IsTciTrustedCa
}

// GetIsTciTrustedCaOk returns a tuple with the IsTciTrustedCa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAppV2) GetIsTciTrustedCaOk() (*bool, bool) {
	if o == nil || IsNil(o.IsTciTrustedCa) {
		return nil, false
	}
	return o.IsTciTrustedCa, true
}

// HasIsTciTrustedCa returns a boolean if a field has been set.
func (o *CreateAppV2) HasIsTciTrustedCa() bool {
	if o != nil && !IsNil(o.IsTciTrustedCa) {
		return true
	}

	return false
}

// SetIsTciTrustedCa gets a reference to the given bool and assigns it to the IsTciTrustedCa field.
func (o *CreateAppV2) SetIsTciTrustedCa(v bool) {
	o.IsTciTrustedCa = &v
}

func (o CreateAppV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAppV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProviderId) {
		toSerialize["provider_id"] = o.ProviderId
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.RepositoryId) {
		toSerialize["repository_id"] = o.RepositoryId
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
	if !IsNil(o.IsAutoDeploy) {
		toSerialize["is_auto_deploy"] = o.IsAutoDeploy
	}
	if !IsNil(o.CommitSha) {
		toSerialize["commit_sha"] = o.CommitSha
	}
	toSerialize["name"] = o.Name
	toSerialize["comment"] = o.Comment
	if !IsNil(o.PresetId) {
		toSerialize["preset_id"] = o.PresetId
	}
	if o.EnvVersion.IsSet() {
		toSerialize["env_version"] = o.EnvVersion.Get()
	}
	if !IsNil(o.Framework) {
		toSerialize["framework"] = o.Framework
	}
	if !IsNil(o.IndexDir) {
		toSerialize["index_dir"] = o.IndexDir
	}
	if !IsNil(o.RunCmd) {
		toSerialize["run_cmd"] = o.RunCmd
	}
	if !IsNil(o.SystemDependencies) {
		toSerialize["system_dependencies"] = o.SystemDependencies
	}
	if !IsNil(o.ProjectId) {
		toSerialize["project_id"] = o.ProjectId
	}
	if !IsNil(o.Configuration) {
		toSerialize["configuration"] = o.Configuration
	}
	toSerialize["language"] = o.Language
	if !IsNil(o.BuildEntryPath) {
		toSerialize["build_entry_path"] = o.BuildEntryPath
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.HealthCheck) {
		toSerialize["health_check"] = o.HealthCheck
	}
	if !IsNil(o.GlobalEnvIds) {
		toSerialize["global_env_ids"] = o.GlobalEnvIds
	}
	if !IsNil(o.DockerImage) {
		toSerialize["docker_image"] = o.DockerImage
	}
	if !IsNil(o.DockerImageTag) {
		toSerialize["docker_image_tag"] = o.DockerImageTag
	}
	if !IsNil(o.IsRussianTrustedCa) {
		toSerialize["is_russian_trusted_ca"] = o.IsRussianTrustedCa
	}
	if !IsNil(o.IsTciTrustedCa) {
		toSerialize["is_tci_trusted_ca"] = o.IsTciTrustedCa
	}
	return toSerialize, nil
}

type NullableCreateAppV2 struct {
	value *CreateAppV2
	isSet bool
}

func (v NullableCreateAppV2) Get() *CreateAppV2 {
	return v.value
}

func (v *NullableCreateAppV2) Set(val *CreateAppV2) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAppV2) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAppV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAppV2(val *CreateAppV2) *NullableCreateAppV2 {
	return &NullableCreateAppV2{value: val, isSet: true}
}

func (v NullableCreateAppV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAppV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


