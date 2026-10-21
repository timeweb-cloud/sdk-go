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

// checks if the AiGatewayBatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiGatewayBatch{}

// AiGatewayBatch Батч
type AiGatewayBatch struct {
	// ID батча
	Id *string `json:"id,omitempty"`
	// Тип объекта
	Object *string `json:"object,omitempty"`
	// Эндпоинт запросов батча
	Endpoint *string `json:"endpoint,omitempty"`
	// Модель батча. На этапе `validating` может быть `null`
	Model NullableString `json:"model,omitempty"`
	// ID входного файла
	InputFileId *string `json:"input_file_id,omitempty"`
	// ID файла результатов — когда батч выполнен
	OutputFileId NullableString `json:"output_file_id,omitempty"`
	// ID файла с ошибками отдельных запросов, если они были
	ErrorFileId NullableString `json:"error_file_id,omitempty"`
	// Срок выполнения
	CompletionWindow *string `json:"completion_window,omitempty"`
	// Статус батча. `cancelling` — отмена выполняется, может занять несколько минут
	Status *string `json:"status,omitempty"`
	// Время создания, Unix-время в секундах
	CreatedAt *int32 `json:"created_at,omitempty"`
	// Время завершения, Unix-время в секундах
	CompletedAt NullableInt32 `json:"completed_at,omitempty"`
	// Крайний срок выполнения, Unix-время в секундах
	ExpiresAt NullableInt32 `json:"expires_at,omitempty"`
	RequestCounts *AiGatewayBatchRequestCounts `json:"request_counts,omitempty"`
	// Метки батча из запроса
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Расход токенов батча
	Usage map[string]interface{} `json:"usage,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AiGatewayBatch AiGatewayBatch

// NewAiGatewayBatch instantiates a new AiGatewayBatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiGatewayBatch() *AiGatewayBatch {
	this := AiGatewayBatch{}
	return &this
}

// NewAiGatewayBatchWithDefaults instantiates a new AiGatewayBatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiGatewayBatchWithDefaults() *AiGatewayBatch {
	this := AiGatewayBatch{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AiGatewayBatch) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayBatch) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AiGatewayBatch) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AiGatewayBatch) SetId(v string) {
	o.Id = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *AiGatewayBatch) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayBatch) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *AiGatewayBatch) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *AiGatewayBatch) SetObject(v string) {
	o.Object = &v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *AiGatewayBatch) GetEndpoint() string {
	if o == nil || IsNil(o.Endpoint) {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayBatch) GetEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.Endpoint) {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *AiGatewayBatch) HasEndpoint() bool {
	if o != nil && !IsNil(o.Endpoint) {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *AiGatewayBatch) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetModel returns the Model field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiGatewayBatch) GetModel() string {
	if o == nil || IsNil(o.Model.Get()) {
		var ret string
		return ret
	}
	return *o.Model.Get()
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiGatewayBatch) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Model.Get(), o.Model.IsSet()
}

// HasModel returns a boolean if a field has been set.
func (o *AiGatewayBatch) HasModel() bool {
	if o != nil && o.Model.IsSet() {
		return true
	}

	return false
}

// SetModel gets a reference to the given NullableString and assigns it to the Model field.
func (o *AiGatewayBatch) SetModel(v string) {
	o.Model.Set(&v)
}
// SetModelNil sets the value for Model to be an explicit nil
func (o *AiGatewayBatch) SetModelNil() {
	o.Model.Set(nil)
}

// UnsetModel ensures that no value is present for Model, not even an explicit nil
func (o *AiGatewayBatch) UnsetModel() {
	o.Model.Unset()
}

// GetInputFileId returns the InputFileId field value if set, zero value otherwise.
func (o *AiGatewayBatch) GetInputFileId() string {
	if o == nil || IsNil(o.InputFileId) {
		var ret string
		return ret
	}
	return *o.InputFileId
}

// GetInputFileIdOk returns a tuple with the InputFileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayBatch) GetInputFileIdOk() (*string, bool) {
	if o == nil || IsNil(o.InputFileId) {
		return nil, false
	}
	return o.InputFileId, true
}

// HasInputFileId returns a boolean if a field has been set.
func (o *AiGatewayBatch) HasInputFileId() bool {
	if o != nil && !IsNil(o.InputFileId) {
		return true
	}

	return false
}

// SetInputFileId gets a reference to the given string and assigns it to the InputFileId field.
func (o *AiGatewayBatch) SetInputFileId(v string) {
	o.InputFileId = &v
}

// GetOutputFileId returns the OutputFileId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiGatewayBatch) GetOutputFileId() string {
	if o == nil || IsNil(o.OutputFileId.Get()) {
		var ret string
		return ret
	}
	return *o.OutputFileId.Get()
}

// GetOutputFileIdOk returns a tuple with the OutputFileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiGatewayBatch) GetOutputFileIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OutputFileId.Get(), o.OutputFileId.IsSet()
}

// HasOutputFileId returns a boolean if a field has been set.
func (o *AiGatewayBatch) HasOutputFileId() bool {
	if o != nil && o.OutputFileId.IsSet() {
		return true
	}

	return false
}

// SetOutputFileId gets a reference to the given NullableString and assigns it to the OutputFileId field.
func (o *AiGatewayBatch) SetOutputFileId(v string) {
	o.OutputFileId.Set(&v)
}
// SetOutputFileIdNil sets the value for OutputFileId to be an explicit nil
func (o *AiGatewayBatch) SetOutputFileIdNil() {
	o.OutputFileId.Set(nil)
}

// UnsetOutputFileId ensures that no value is present for OutputFileId, not even an explicit nil
func (o *AiGatewayBatch) UnsetOutputFileId() {
	o.OutputFileId.Unset()
}

// GetErrorFileId returns the ErrorFileId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiGatewayBatch) GetErrorFileId() string {
	if o == nil || IsNil(o.ErrorFileId.Get()) {
		var ret string
		return ret
	}
	return *o.ErrorFileId.Get()
}

// GetErrorFileIdOk returns a tuple with the ErrorFileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiGatewayBatch) GetErrorFileIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorFileId.Get(), o.ErrorFileId.IsSet()
}

// HasErrorFileId returns a boolean if a field has been set.
func (o *AiGatewayBatch) HasErrorFileId() bool {
	if o != nil && o.ErrorFileId.IsSet() {
		return true
	}

	return false
}

// SetErrorFileId gets a reference to the given NullableString and assigns it to the ErrorFileId field.
func (o *AiGatewayBatch) SetErrorFileId(v string) {
	o.ErrorFileId.Set(&v)
}
// SetErrorFileIdNil sets the value for ErrorFileId to be an explicit nil
func (o *AiGatewayBatch) SetErrorFileIdNil() {
	o.ErrorFileId.Set(nil)
}

// UnsetErrorFileId ensures that no value is present for ErrorFileId, not even an explicit nil
func (o *AiGatewayBatch) UnsetErrorFileId() {
	o.ErrorFileId.Unset()
}

// GetCompletionWindow returns the CompletionWindow field value if set, zero value otherwise.
func (o *AiGatewayBatch) GetCompletionWindow() string {
	if o == nil || IsNil(o.CompletionWindow) {
		var ret string
		return ret
	}
	return *o.CompletionWindow
}

// GetCompletionWindowOk returns a tuple with the CompletionWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayBatch) GetCompletionWindowOk() (*string, bool) {
	if o == nil || IsNil(o.CompletionWindow) {
		return nil, false
	}
	return o.CompletionWindow, true
}

// HasCompletionWindow returns a boolean if a field has been set.
func (o *AiGatewayBatch) HasCompletionWindow() bool {
	if o != nil && !IsNil(o.CompletionWindow) {
		return true
	}

	return false
}

// SetCompletionWindow gets a reference to the given string and assigns it to the CompletionWindow field.
func (o *AiGatewayBatch) SetCompletionWindow(v string) {
	o.CompletionWindow = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AiGatewayBatch) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayBatch) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AiGatewayBatch) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AiGatewayBatch) SetStatus(v string) {
	o.Status = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AiGatewayBatch) GetCreatedAt() int32 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayBatch) GetCreatedAtOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AiGatewayBatch) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *AiGatewayBatch) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiGatewayBatch) GetCompletedAt() int32 {
	if o == nil || IsNil(o.CompletedAt.Get()) {
		var ret int32
		return ret
	}
	return *o.CompletedAt.Get()
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiGatewayBatch) GetCompletedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompletedAt.Get(), o.CompletedAt.IsSet()
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *AiGatewayBatch) HasCompletedAt() bool {
	if o != nil && o.CompletedAt.IsSet() {
		return true
	}

	return false
}

// SetCompletedAt gets a reference to the given NullableInt32 and assigns it to the CompletedAt field.
func (o *AiGatewayBatch) SetCompletedAt(v int32) {
	o.CompletedAt.Set(&v)
}
// SetCompletedAtNil sets the value for CompletedAt to be an explicit nil
func (o *AiGatewayBatch) SetCompletedAtNil() {
	o.CompletedAt.Set(nil)
}

// UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
func (o *AiGatewayBatch) UnsetCompletedAt() {
	o.CompletedAt.Unset()
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiGatewayBatch) GetExpiresAt() int32 {
	if o == nil || IsNil(o.ExpiresAt.Get()) {
		var ret int32
		return ret
	}
	return *o.ExpiresAt.Get()
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiGatewayBatch) GetExpiresAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiresAt.Get(), o.ExpiresAt.IsSet()
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *AiGatewayBatch) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt.IsSet() {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given NullableInt32 and assigns it to the ExpiresAt field.
func (o *AiGatewayBatch) SetExpiresAt(v int32) {
	o.ExpiresAt.Set(&v)
}
// SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil
func (o *AiGatewayBatch) SetExpiresAtNil() {
	o.ExpiresAt.Set(nil)
}

// UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
func (o *AiGatewayBatch) UnsetExpiresAt() {
	o.ExpiresAt.Unset()
}

// GetRequestCounts returns the RequestCounts field value if set, zero value otherwise.
func (o *AiGatewayBatch) GetRequestCounts() AiGatewayBatchRequestCounts {
	if o == nil || IsNil(o.RequestCounts) {
		var ret AiGatewayBatchRequestCounts
		return ret
	}
	return *o.RequestCounts
}

// GetRequestCountsOk returns a tuple with the RequestCounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayBatch) GetRequestCountsOk() (*AiGatewayBatchRequestCounts, bool) {
	if o == nil || IsNil(o.RequestCounts) {
		return nil, false
	}
	return o.RequestCounts, true
}

// HasRequestCounts returns a boolean if a field has been set.
func (o *AiGatewayBatch) HasRequestCounts() bool {
	if o != nil && !IsNil(o.RequestCounts) {
		return true
	}

	return false
}

// SetRequestCounts gets a reference to the given AiGatewayBatchRequestCounts and assigns it to the RequestCounts field.
func (o *AiGatewayBatch) SetRequestCounts(v AiGatewayBatchRequestCounts) {
	o.RequestCounts = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiGatewayBatch) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiGatewayBatch) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *AiGatewayBatch) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *AiGatewayBatch) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetUsage returns the Usage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiGatewayBatch) GetUsage() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiGatewayBatch) GetUsageOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Usage) {
		return map[string]interface{}{}, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *AiGatewayBatch) HasUsage() bool {
	if o != nil && IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given map[string]interface{} and assigns it to the Usage field.
func (o *AiGatewayBatch) SetUsage(v map[string]interface{}) {
	o.Usage = v
}

func (o AiGatewayBatch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiGatewayBatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	if !IsNil(o.Endpoint) {
		toSerialize["endpoint"] = o.Endpoint
	}
	if o.Model.IsSet() {
		toSerialize["model"] = o.Model.Get()
	}
	if !IsNil(o.InputFileId) {
		toSerialize["input_file_id"] = o.InputFileId
	}
	if o.OutputFileId.IsSet() {
		toSerialize["output_file_id"] = o.OutputFileId.Get()
	}
	if o.ErrorFileId.IsSet() {
		toSerialize["error_file_id"] = o.ErrorFileId.Get()
	}
	if !IsNil(o.CompletionWindow) {
		toSerialize["completion_window"] = o.CompletionWindow
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.CompletedAt.IsSet() {
		toSerialize["completed_at"] = o.CompletedAt.Get()
	}
	if o.ExpiresAt.IsSet() {
		toSerialize["expires_at"] = o.ExpiresAt.Get()
	}
	if !IsNil(o.RequestCounts) {
		toSerialize["request_counts"] = o.RequestCounts
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Usage != nil {
		toSerialize["usage"] = o.Usage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AiGatewayBatch) UnmarshalJSON(bytes []byte) (err error) {
	varAiGatewayBatch := _AiGatewayBatch{}

	if err = json.Unmarshal(bytes, &varAiGatewayBatch); err == nil {
		*o = AiGatewayBatch(varAiGatewayBatch)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object")
		delete(additionalProperties, "endpoint")
		delete(additionalProperties, "model")
		delete(additionalProperties, "input_file_id")
		delete(additionalProperties, "output_file_id")
		delete(additionalProperties, "error_file_id")
		delete(additionalProperties, "completion_window")
		delete(additionalProperties, "status")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "completed_at")
		delete(additionalProperties, "expires_at")
		delete(additionalProperties, "request_counts")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "usage")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAiGatewayBatch struct {
	value *AiGatewayBatch
	isSet bool
}

func (v NullableAiGatewayBatch) Get() *AiGatewayBatch {
	return v.value
}

func (v *NullableAiGatewayBatch) Set(val *AiGatewayBatch) {
	v.value = val
	v.isSet = true
}

func (v NullableAiGatewayBatch) IsSet() bool {
	return v.isSet
}

func (v *NullableAiGatewayBatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiGatewayBatch(val *AiGatewayBatch) *NullableAiGatewayBatch {
	return &NullableAiGatewayBatch{value: val, isSet: true}
}

func (v NullableAiGatewayBatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiGatewayBatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


