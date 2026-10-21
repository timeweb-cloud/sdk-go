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

// checks if the ToolAuthDataParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ToolAuthDataParams{}

// ToolAuthDataParams Параметры авторизации. Набор полей зависит от `type`: для `none` — пустой объект, для `bearer` — `token`, для `custom_headers` — `headers`, для `oauth2` — поля OAuth2-клиента.
type ToolAuthDataParams struct {
	// Токен для заголовка `Authorization: Bearer` (для `bearer`). В ответах заменён на `********`
	Token *string `json:"token,omitempty"`
	// HTTP-заголовки, которые передаются MCP-серверу (для `custom_headers`). В ответах значения заменены на `********`
	Headers *map[string]string `json:"headers,omitempty"`
	// ID OAuth2-клиента (для `oauth2`). Не нужен, если сервер поддерживает динамическую регистрацию клиентов
	ClientId *string `json:"client_id,omitempty"`
	// Секрет OAuth2-клиента (для `oauth2`). В ответах заменён на `********`
	ClientSecret *string `json:"client_secret,omitempty"`
	// Адрес страницы авторизации провайдера (для `oauth2`). Если не задан, определяется по метаданным сервера
	AuthorizationUrl *string `json:"authorization_url,omitempty"`
	// Адрес получения токенов провайдера (для `oauth2`). Если не задан, определяется по метаданным сервера
	TokenUrl *string `json:"token_url,omitempty"`
	// Адрес динамической регистрации клиента (для `oauth2`)
	RegistrationUrl *string `json:"registration_url,omitempty"`
	// Запрашиваемые права (для `oauth2`). Если не заданы, берутся из метаданных сервера
	Scopes []string `json:"scopes,omitempty"`
}

// NewToolAuthDataParams instantiates a new ToolAuthDataParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToolAuthDataParams() *ToolAuthDataParams {
	this := ToolAuthDataParams{}
	return &this
}

// NewToolAuthDataParamsWithDefaults instantiates a new ToolAuthDataParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewToolAuthDataParamsWithDefaults() *ToolAuthDataParams {
	this := ToolAuthDataParams{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *ToolAuthDataParams) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolAuthDataParams) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *ToolAuthDataParams) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *ToolAuthDataParams) SetToken(v string) {
	o.Token = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *ToolAuthDataParams) GetHeaders() map[string]string {
	if o == nil || IsNil(o.Headers) {
		var ret map[string]string
		return ret
	}
	return *o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolAuthDataParams) GetHeadersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *ToolAuthDataParams) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string]string and assigns it to the Headers field.
func (o *ToolAuthDataParams) SetHeaders(v map[string]string) {
	o.Headers = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ToolAuthDataParams) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolAuthDataParams) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ToolAuthDataParams) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ToolAuthDataParams) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *ToolAuthDataParams) GetClientSecret() string {
	if o == nil || IsNil(o.ClientSecret) {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolAuthDataParams) GetClientSecretOk() (*string, bool) {
	if o == nil || IsNil(o.ClientSecret) {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *ToolAuthDataParams) HasClientSecret() bool {
	if o != nil && !IsNil(o.ClientSecret) {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *ToolAuthDataParams) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetAuthorizationUrl returns the AuthorizationUrl field value if set, zero value otherwise.
func (o *ToolAuthDataParams) GetAuthorizationUrl() string {
	if o == nil || IsNil(o.AuthorizationUrl) {
		var ret string
		return ret
	}
	return *o.AuthorizationUrl
}

// GetAuthorizationUrlOk returns a tuple with the AuthorizationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolAuthDataParams) GetAuthorizationUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorizationUrl) {
		return nil, false
	}
	return o.AuthorizationUrl, true
}

// HasAuthorizationUrl returns a boolean if a field has been set.
func (o *ToolAuthDataParams) HasAuthorizationUrl() bool {
	if o != nil && !IsNil(o.AuthorizationUrl) {
		return true
	}

	return false
}

// SetAuthorizationUrl gets a reference to the given string and assigns it to the AuthorizationUrl field.
func (o *ToolAuthDataParams) SetAuthorizationUrl(v string) {
	o.AuthorizationUrl = &v
}

// GetTokenUrl returns the TokenUrl field value if set, zero value otherwise.
func (o *ToolAuthDataParams) GetTokenUrl() string {
	if o == nil || IsNil(o.TokenUrl) {
		var ret string
		return ret
	}
	return *o.TokenUrl
}

// GetTokenUrlOk returns a tuple with the TokenUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolAuthDataParams) GetTokenUrlOk() (*string, bool) {
	if o == nil || IsNil(o.TokenUrl) {
		return nil, false
	}
	return o.TokenUrl, true
}

// HasTokenUrl returns a boolean if a field has been set.
func (o *ToolAuthDataParams) HasTokenUrl() bool {
	if o != nil && !IsNil(o.TokenUrl) {
		return true
	}

	return false
}

// SetTokenUrl gets a reference to the given string and assigns it to the TokenUrl field.
func (o *ToolAuthDataParams) SetTokenUrl(v string) {
	o.TokenUrl = &v
}

// GetRegistrationUrl returns the RegistrationUrl field value if set, zero value otherwise.
func (o *ToolAuthDataParams) GetRegistrationUrl() string {
	if o == nil || IsNil(o.RegistrationUrl) {
		var ret string
		return ret
	}
	return *o.RegistrationUrl
}

// GetRegistrationUrlOk returns a tuple with the RegistrationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolAuthDataParams) GetRegistrationUrlOk() (*string, bool) {
	if o == nil || IsNil(o.RegistrationUrl) {
		return nil, false
	}
	return o.RegistrationUrl, true
}

// HasRegistrationUrl returns a boolean if a field has been set.
func (o *ToolAuthDataParams) HasRegistrationUrl() bool {
	if o != nil && !IsNil(o.RegistrationUrl) {
		return true
	}

	return false
}

// SetRegistrationUrl gets a reference to the given string and assigns it to the RegistrationUrl field.
func (o *ToolAuthDataParams) SetRegistrationUrl(v string) {
	o.RegistrationUrl = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *ToolAuthDataParams) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolAuthDataParams) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *ToolAuthDataParams) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *ToolAuthDataParams) SetScopes(v []string) {
	o.Scopes = v
}

func (o ToolAuthDataParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ToolAuthDataParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	if !IsNil(o.ClientId) {
		toSerialize["client_id"] = o.ClientId
	}
	if !IsNil(o.ClientSecret) {
		toSerialize["client_secret"] = o.ClientSecret
	}
	if !IsNil(o.AuthorizationUrl) {
		toSerialize["authorization_url"] = o.AuthorizationUrl
	}
	if !IsNil(o.TokenUrl) {
		toSerialize["token_url"] = o.TokenUrl
	}
	if !IsNil(o.RegistrationUrl) {
		toSerialize["registration_url"] = o.RegistrationUrl
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	return toSerialize, nil
}

type NullableToolAuthDataParams struct {
	value *ToolAuthDataParams
	isSet bool
}

func (v NullableToolAuthDataParams) Get() *ToolAuthDataParams {
	return v.value
}

func (v *NullableToolAuthDataParams) Set(val *ToolAuthDataParams) {
	v.value = val
	v.isSet = true
}

func (v NullableToolAuthDataParams) IsSet() bool {
	return v.isSet
}

func (v *NullableToolAuthDataParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToolAuthDataParams(val *ToolAuthDataParams) *NullableToolAuthDataParams {
	return &NullableToolAuthDataParams{value: val, isSet: true}
}

func (v NullableToolAuthDataParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToolAuthDataParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


