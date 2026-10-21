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

// checks if the AiGatewayResponseCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiGatewayResponseCreateRequest{}

// AiGatewayResponseCreateRequest Запрос Responses API в формате OpenAI. Остальные параметры OpenAI передаются модели как есть
type AiGatewayResponseCreateRequest struct {
	// ID модели из `GET /v1/models`
	Model string `json:"model"`
	// Текст или массив входных элементов (сообщения, результаты вызова функций и т. п.)
	Input string `json:"input"`
	// Системные инструкции
	Instructions *string `json:"instructions,omitempty"`
	// ID сохранённого ответа, чтобы продолжить диалог
	PreviousResponseId *string `json:"previous_response_id,omitempty"`
	// Отправлять ответ потоком типизированных событий Server-Sent Events
	Stream *bool `json:"stream,omitempty"`
	// Выполнить запрос в фоне: результат получайте через `GET /v1/responses/{response_id}`
	Background *bool `json:"background,omitempty"`
	// Сохранить ответ, чтобы получить его позже или продолжить диалог
	Store *bool `json:"store,omitempty"`
	// Максимум выходных токенов
	MaxOutputTokens *int32 `json:"max_output_tokens,omitempty"`
	// Инструменты, которые может вызвать модель
	Tools []map[string]interface{} `json:"tools,omitempty"`
	Reasoning *AiGatewayResponseCreateRequestReasoning `json:"reasoning,omitempty"`
	// Настройки формата текста ответа
	Text map[string]interface{} `json:"text,omitempty"`
	// Температура выборки
	Temperature *float32 `json:"temperature,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AiGatewayResponseCreateRequest AiGatewayResponseCreateRequest

// NewAiGatewayResponseCreateRequest instantiates a new AiGatewayResponseCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiGatewayResponseCreateRequest(model string, input string) *AiGatewayResponseCreateRequest {
	this := AiGatewayResponseCreateRequest{}
	this.Model = model
	this.Input = input
	var stream bool = false
	this.Stream = &stream
	return &this
}

// NewAiGatewayResponseCreateRequestWithDefaults instantiates a new AiGatewayResponseCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiGatewayResponseCreateRequestWithDefaults() *AiGatewayResponseCreateRequest {
	this := AiGatewayResponseCreateRequest{}
	var stream bool = false
	this.Stream = &stream
	return &this
}

// GetModel returns the Model field value
func (o *AiGatewayResponseCreateRequest) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *AiGatewayResponseCreateRequest) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *AiGatewayResponseCreateRequest) SetModel(v string) {
	o.Model = v
}

// GetInput returns the Input field value
func (o *AiGatewayResponseCreateRequest) GetInput() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Input
}

// GetInputOk returns a tuple with the Input field value
// and a boolean to check if the value has been set.
func (o *AiGatewayResponseCreateRequest) GetInputOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Input, true
}

// SetInput sets field value
func (o *AiGatewayResponseCreateRequest) SetInput(v string) {
	o.Input = v
}

// GetInstructions returns the Instructions field value if set, zero value otherwise.
func (o *AiGatewayResponseCreateRequest) GetInstructions() string {
	if o == nil || IsNil(o.Instructions) {
		var ret string
		return ret
	}
	return *o.Instructions
}

// GetInstructionsOk returns a tuple with the Instructions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayResponseCreateRequest) GetInstructionsOk() (*string, bool) {
	if o == nil || IsNil(o.Instructions) {
		return nil, false
	}
	return o.Instructions, true
}

// HasInstructions returns a boolean if a field has been set.
func (o *AiGatewayResponseCreateRequest) HasInstructions() bool {
	if o != nil && !IsNil(o.Instructions) {
		return true
	}

	return false
}

// SetInstructions gets a reference to the given string and assigns it to the Instructions field.
func (o *AiGatewayResponseCreateRequest) SetInstructions(v string) {
	o.Instructions = &v
}

// GetPreviousResponseId returns the PreviousResponseId field value if set, zero value otherwise.
func (o *AiGatewayResponseCreateRequest) GetPreviousResponseId() string {
	if o == nil || IsNil(o.PreviousResponseId) {
		var ret string
		return ret
	}
	return *o.PreviousResponseId
}

// GetPreviousResponseIdOk returns a tuple with the PreviousResponseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayResponseCreateRequest) GetPreviousResponseIdOk() (*string, bool) {
	if o == nil || IsNil(o.PreviousResponseId) {
		return nil, false
	}
	return o.PreviousResponseId, true
}

// HasPreviousResponseId returns a boolean if a field has been set.
func (o *AiGatewayResponseCreateRequest) HasPreviousResponseId() bool {
	if o != nil && !IsNil(o.PreviousResponseId) {
		return true
	}

	return false
}

// SetPreviousResponseId gets a reference to the given string and assigns it to the PreviousResponseId field.
func (o *AiGatewayResponseCreateRequest) SetPreviousResponseId(v string) {
	o.PreviousResponseId = &v
}

// GetStream returns the Stream field value if set, zero value otherwise.
func (o *AiGatewayResponseCreateRequest) GetStream() bool {
	if o == nil || IsNil(o.Stream) {
		var ret bool
		return ret
	}
	return *o.Stream
}

// GetStreamOk returns a tuple with the Stream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayResponseCreateRequest) GetStreamOk() (*bool, bool) {
	if o == nil || IsNil(o.Stream) {
		return nil, false
	}
	return o.Stream, true
}

// HasStream returns a boolean if a field has been set.
func (o *AiGatewayResponseCreateRequest) HasStream() bool {
	if o != nil && !IsNil(o.Stream) {
		return true
	}

	return false
}

// SetStream gets a reference to the given bool and assigns it to the Stream field.
func (o *AiGatewayResponseCreateRequest) SetStream(v bool) {
	o.Stream = &v
}

// GetBackground returns the Background field value if set, zero value otherwise.
func (o *AiGatewayResponseCreateRequest) GetBackground() bool {
	if o == nil || IsNil(o.Background) {
		var ret bool
		return ret
	}
	return *o.Background
}

// GetBackgroundOk returns a tuple with the Background field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayResponseCreateRequest) GetBackgroundOk() (*bool, bool) {
	if o == nil || IsNil(o.Background) {
		return nil, false
	}
	return o.Background, true
}

// HasBackground returns a boolean if a field has been set.
func (o *AiGatewayResponseCreateRequest) HasBackground() bool {
	if o != nil && !IsNil(o.Background) {
		return true
	}

	return false
}

// SetBackground gets a reference to the given bool and assigns it to the Background field.
func (o *AiGatewayResponseCreateRequest) SetBackground(v bool) {
	o.Background = &v
}

// GetStore returns the Store field value if set, zero value otherwise.
func (o *AiGatewayResponseCreateRequest) GetStore() bool {
	if o == nil || IsNil(o.Store) {
		var ret bool
		return ret
	}
	return *o.Store
}

// GetStoreOk returns a tuple with the Store field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayResponseCreateRequest) GetStoreOk() (*bool, bool) {
	if o == nil || IsNil(o.Store) {
		return nil, false
	}
	return o.Store, true
}

// HasStore returns a boolean if a field has been set.
func (o *AiGatewayResponseCreateRequest) HasStore() bool {
	if o != nil && !IsNil(o.Store) {
		return true
	}

	return false
}

// SetStore gets a reference to the given bool and assigns it to the Store field.
func (o *AiGatewayResponseCreateRequest) SetStore(v bool) {
	o.Store = &v
}

// GetMaxOutputTokens returns the MaxOutputTokens field value if set, zero value otherwise.
func (o *AiGatewayResponseCreateRequest) GetMaxOutputTokens() int32 {
	if o == nil || IsNil(o.MaxOutputTokens) {
		var ret int32
		return ret
	}
	return *o.MaxOutputTokens
}

// GetMaxOutputTokensOk returns a tuple with the MaxOutputTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayResponseCreateRequest) GetMaxOutputTokensOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxOutputTokens) {
		return nil, false
	}
	return o.MaxOutputTokens, true
}

// HasMaxOutputTokens returns a boolean if a field has been set.
func (o *AiGatewayResponseCreateRequest) HasMaxOutputTokens() bool {
	if o != nil && !IsNil(o.MaxOutputTokens) {
		return true
	}

	return false
}

// SetMaxOutputTokens gets a reference to the given int32 and assigns it to the MaxOutputTokens field.
func (o *AiGatewayResponseCreateRequest) SetMaxOutputTokens(v int32) {
	o.MaxOutputTokens = &v
}

// GetTools returns the Tools field value if set, zero value otherwise.
func (o *AiGatewayResponseCreateRequest) GetTools() []map[string]interface{} {
	if o == nil || IsNil(o.Tools) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Tools
}

// GetToolsOk returns a tuple with the Tools field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayResponseCreateRequest) GetToolsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Tools) {
		return nil, false
	}
	return o.Tools, true
}

// HasTools returns a boolean if a field has been set.
func (o *AiGatewayResponseCreateRequest) HasTools() bool {
	if o != nil && !IsNil(o.Tools) {
		return true
	}

	return false
}

// SetTools gets a reference to the given []map[string]interface{} and assigns it to the Tools field.
func (o *AiGatewayResponseCreateRequest) SetTools(v []map[string]interface{}) {
	o.Tools = v
}

// GetReasoning returns the Reasoning field value if set, zero value otherwise.
func (o *AiGatewayResponseCreateRequest) GetReasoning() AiGatewayResponseCreateRequestReasoning {
	if o == nil || IsNil(o.Reasoning) {
		var ret AiGatewayResponseCreateRequestReasoning
		return ret
	}
	return *o.Reasoning
}

// GetReasoningOk returns a tuple with the Reasoning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayResponseCreateRequest) GetReasoningOk() (*AiGatewayResponseCreateRequestReasoning, bool) {
	if o == nil || IsNil(o.Reasoning) {
		return nil, false
	}
	return o.Reasoning, true
}

// HasReasoning returns a boolean if a field has been set.
func (o *AiGatewayResponseCreateRequest) HasReasoning() bool {
	if o != nil && !IsNil(o.Reasoning) {
		return true
	}

	return false
}

// SetReasoning gets a reference to the given AiGatewayResponseCreateRequestReasoning and assigns it to the Reasoning field.
func (o *AiGatewayResponseCreateRequest) SetReasoning(v AiGatewayResponseCreateRequestReasoning) {
	o.Reasoning = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *AiGatewayResponseCreateRequest) GetText() map[string]interface{} {
	if o == nil || IsNil(o.Text) {
		var ret map[string]interface{}
		return ret
	}
	return o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayResponseCreateRequest) GetTextOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Text) {
		return map[string]interface{}{}, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *AiGatewayResponseCreateRequest) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given map[string]interface{} and assigns it to the Text field.
func (o *AiGatewayResponseCreateRequest) SetText(v map[string]interface{}) {
	o.Text = v
}

// GetTemperature returns the Temperature field value if set, zero value otherwise.
func (o *AiGatewayResponseCreateRequest) GetTemperature() float32 {
	if o == nil || IsNil(o.Temperature) {
		var ret float32
		return ret
	}
	return *o.Temperature
}

// GetTemperatureOk returns a tuple with the Temperature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayResponseCreateRequest) GetTemperatureOk() (*float32, bool) {
	if o == nil || IsNil(o.Temperature) {
		return nil, false
	}
	return o.Temperature, true
}

// HasTemperature returns a boolean if a field has been set.
func (o *AiGatewayResponseCreateRequest) HasTemperature() bool {
	if o != nil && !IsNil(o.Temperature) {
		return true
	}

	return false
}

// SetTemperature gets a reference to the given float32 and assigns it to the Temperature field.
func (o *AiGatewayResponseCreateRequest) SetTemperature(v float32) {
	o.Temperature = &v
}

func (o AiGatewayResponseCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiGatewayResponseCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["model"] = o.Model
	toSerialize["input"] = o.Input
	if !IsNil(o.Instructions) {
		toSerialize["instructions"] = o.Instructions
	}
	if !IsNil(o.PreviousResponseId) {
		toSerialize["previous_response_id"] = o.PreviousResponseId
	}
	if !IsNil(o.Stream) {
		toSerialize["stream"] = o.Stream
	}
	if !IsNil(o.Background) {
		toSerialize["background"] = o.Background
	}
	if !IsNil(o.Store) {
		toSerialize["store"] = o.Store
	}
	if !IsNil(o.MaxOutputTokens) {
		toSerialize["max_output_tokens"] = o.MaxOutputTokens
	}
	if !IsNil(o.Tools) {
		toSerialize["tools"] = o.Tools
	}
	if !IsNil(o.Reasoning) {
		toSerialize["reasoning"] = o.Reasoning
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.Temperature) {
		toSerialize["temperature"] = o.Temperature
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AiGatewayResponseCreateRequest) UnmarshalJSON(bytes []byte) (err error) {
	varAiGatewayResponseCreateRequest := _AiGatewayResponseCreateRequest{}

	if err = json.Unmarshal(bytes, &varAiGatewayResponseCreateRequest); err == nil {
		*o = AiGatewayResponseCreateRequest(varAiGatewayResponseCreateRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "model")
		delete(additionalProperties, "input")
		delete(additionalProperties, "instructions")
		delete(additionalProperties, "previous_response_id")
		delete(additionalProperties, "stream")
		delete(additionalProperties, "background")
		delete(additionalProperties, "store")
		delete(additionalProperties, "max_output_tokens")
		delete(additionalProperties, "tools")
		delete(additionalProperties, "reasoning")
		delete(additionalProperties, "text")
		delete(additionalProperties, "temperature")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAiGatewayResponseCreateRequest struct {
	value *AiGatewayResponseCreateRequest
	isSet bool
}

func (v NullableAiGatewayResponseCreateRequest) Get() *AiGatewayResponseCreateRequest {
	return v.value
}

func (v *NullableAiGatewayResponseCreateRequest) Set(val *AiGatewayResponseCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAiGatewayResponseCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAiGatewayResponseCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiGatewayResponseCreateRequest(val *AiGatewayResponseCreateRequest) *NullableAiGatewayResponseCreateRequest {
	return &NullableAiGatewayResponseCreateRequest{value: val, isSet: true}
}

func (v NullableAiGatewayResponseCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiGatewayResponseCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


