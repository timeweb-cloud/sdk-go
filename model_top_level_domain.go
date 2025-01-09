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

// checks if the TopLevelDomain type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TopLevelDomain{}

// TopLevelDomain Доменная зона.
type TopLevelDomain struct {
	// Список доступных периодов для регистрации/продления доменов со стоимостью.
	AllowedBuyPeriods []TopLevelDomainAllowedBuyPeriodsInner `json:"allowed_buy_periods"`
	// Количество дней до истечение срока регистрации, когда можно продлять домен.
	EarlyRenewPeriod NullableFloat32 `json:"early_renew_period"`
	// Количество дней, которые действует льготный период когда вы ещё можете продлить домен, после окончания его регистрации
	GracePeriod float32 `json:"grace_period"`
	// ID доменной зоны.
	Id float32 `json:"id"`
	// Это логическое значение, которое показывает, опубликована ли доменная зона.
	IsPublished bool `json:"is_published"`
	// Это логическое значение, которое показывает, зарегистрирована ли доменная зона.
	IsRegistered bool `json:"is_registered"`
	// Это логическое значение, которое показывает, включено ли по умолчанию скрытие данных администратора для доменной зоны.
	IsWhoisPrivacyDefaultEnabled bool `json:"is_whois_privacy_default_enabled"`
	// Это логическое значение, которое показывает, доступно ли управление скрытием данных администратора для доменной зоны.
	IsWhoisPrivacyEnabled bool `json:"is_whois_privacy_enabled"`
	// Имя доменной зоны.
	Name string `json:"name"`
	// Цена регистрации домена
	Price float32 `json:"price"`
	// Цена продления домена.
	ProlongPrice float32 `json:"prolong_price"`
	// Регистратор доменной зоны.
	Registrar string `json:"registrar"`
	// Цена услуги трансфера домена.
	Transfer float32 `json:"transfer"`
	// Цена услуги скрытия данных администратора для доменной зоны.
	WhoisPrivacyPrice float32 `json:"whois_privacy_price"`
}

// NewTopLevelDomain instantiates a new TopLevelDomain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTopLevelDomain(allowedBuyPeriods []TopLevelDomainAllowedBuyPeriodsInner, earlyRenewPeriod NullableFloat32, gracePeriod float32, id float32, isPublished bool, isRegistered bool, isWhoisPrivacyDefaultEnabled bool, isWhoisPrivacyEnabled bool, name string, price float32, prolongPrice float32, registrar string, transfer float32, whoisPrivacyPrice float32) *TopLevelDomain {
	this := TopLevelDomain{}
	this.AllowedBuyPeriods = allowedBuyPeriods
	this.EarlyRenewPeriod = earlyRenewPeriod
	this.GracePeriod = gracePeriod
	this.Id = id
	this.IsPublished = isPublished
	this.IsRegistered = isRegistered
	this.IsWhoisPrivacyDefaultEnabled = isWhoisPrivacyDefaultEnabled
	this.IsWhoisPrivacyEnabled = isWhoisPrivacyEnabled
	this.Name = name
	this.Price = price
	this.ProlongPrice = prolongPrice
	this.Registrar = registrar
	this.Transfer = transfer
	this.WhoisPrivacyPrice = whoisPrivacyPrice
	return &this
}

// NewTopLevelDomainWithDefaults instantiates a new TopLevelDomain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTopLevelDomainWithDefaults() *TopLevelDomain {
	this := TopLevelDomain{}
	return &this
}

// GetAllowedBuyPeriods returns the AllowedBuyPeriods field value
func (o *TopLevelDomain) GetAllowedBuyPeriods() []TopLevelDomainAllowedBuyPeriodsInner {
	if o == nil {
		var ret []TopLevelDomainAllowedBuyPeriodsInner
		return ret
	}

	return o.AllowedBuyPeriods
}

// GetAllowedBuyPeriodsOk returns a tuple with the AllowedBuyPeriods field value
// and a boolean to check if the value has been set.
func (o *TopLevelDomain) GetAllowedBuyPeriodsOk() ([]TopLevelDomainAllowedBuyPeriodsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowedBuyPeriods, true
}

// SetAllowedBuyPeriods sets field value
func (o *TopLevelDomain) SetAllowedBuyPeriods(v []TopLevelDomainAllowedBuyPeriodsInner) {
	o.AllowedBuyPeriods = v
}

// GetEarlyRenewPeriod returns the EarlyRenewPeriod field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *TopLevelDomain) GetEarlyRenewPeriod() float32 {
	if o == nil || o.EarlyRenewPeriod.Get() == nil {
		var ret float32
		return ret
	}

	return *o.EarlyRenewPeriod.Get()
}

// GetEarlyRenewPeriodOk returns a tuple with the EarlyRenewPeriod field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TopLevelDomain) GetEarlyRenewPeriodOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.EarlyRenewPeriod.Get(), o.EarlyRenewPeriod.IsSet()
}

// SetEarlyRenewPeriod sets field value
func (o *TopLevelDomain) SetEarlyRenewPeriod(v float32) {
	o.EarlyRenewPeriod.Set(&v)
}

// GetGracePeriod returns the GracePeriod field value
func (o *TopLevelDomain) GetGracePeriod() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.GracePeriod
}

// GetGracePeriodOk returns a tuple with the GracePeriod field value
// and a boolean to check if the value has been set.
func (o *TopLevelDomain) GetGracePeriodOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GracePeriod, true
}

// SetGracePeriod sets field value
func (o *TopLevelDomain) SetGracePeriod(v float32) {
	o.GracePeriod = v
}

// GetId returns the Id field value
func (o *TopLevelDomain) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TopLevelDomain) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TopLevelDomain) SetId(v float32) {
	o.Id = v
}

// GetIsPublished returns the IsPublished field value
func (o *TopLevelDomain) GetIsPublished() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPublished
}

// GetIsPublishedOk returns a tuple with the IsPublished field value
// and a boolean to check if the value has been set.
func (o *TopLevelDomain) GetIsPublishedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPublished, true
}

// SetIsPublished sets field value
func (o *TopLevelDomain) SetIsPublished(v bool) {
	o.IsPublished = v
}

// GetIsRegistered returns the IsRegistered field value
func (o *TopLevelDomain) GetIsRegistered() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsRegistered
}

// GetIsRegisteredOk returns a tuple with the IsRegistered field value
// and a boolean to check if the value has been set.
func (o *TopLevelDomain) GetIsRegisteredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsRegistered, true
}

// SetIsRegistered sets field value
func (o *TopLevelDomain) SetIsRegistered(v bool) {
	o.IsRegistered = v
}

// GetIsWhoisPrivacyDefaultEnabled returns the IsWhoisPrivacyDefaultEnabled field value
func (o *TopLevelDomain) GetIsWhoisPrivacyDefaultEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsWhoisPrivacyDefaultEnabled
}

// GetIsWhoisPrivacyDefaultEnabledOk returns a tuple with the IsWhoisPrivacyDefaultEnabled field value
// and a boolean to check if the value has been set.
func (o *TopLevelDomain) GetIsWhoisPrivacyDefaultEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsWhoisPrivacyDefaultEnabled, true
}

// SetIsWhoisPrivacyDefaultEnabled sets field value
func (o *TopLevelDomain) SetIsWhoisPrivacyDefaultEnabled(v bool) {
	o.IsWhoisPrivacyDefaultEnabled = v
}

// GetIsWhoisPrivacyEnabled returns the IsWhoisPrivacyEnabled field value
func (o *TopLevelDomain) GetIsWhoisPrivacyEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsWhoisPrivacyEnabled
}

// GetIsWhoisPrivacyEnabledOk returns a tuple with the IsWhoisPrivacyEnabled field value
// and a boolean to check if the value has been set.
func (o *TopLevelDomain) GetIsWhoisPrivacyEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsWhoisPrivacyEnabled, true
}

// SetIsWhoisPrivacyEnabled sets field value
func (o *TopLevelDomain) SetIsWhoisPrivacyEnabled(v bool) {
	o.IsWhoisPrivacyEnabled = v
}

// GetName returns the Name field value
func (o *TopLevelDomain) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TopLevelDomain) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TopLevelDomain) SetName(v string) {
	o.Name = v
}

// GetPrice returns the Price field value
func (o *TopLevelDomain) GetPrice() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *TopLevelDomain) GetPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *TopLevelDomain) SetPrice(v float32) {
	o.Price = v
}

// GetProlongPrice returns the ProlongPrice field value
func (o *TopLevelDomain) GetProlongPrice() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ProlongPrice
}

// GetProlongPriceOk returns a tuple with the ProlongPrice field value
// and a boolean to check if the value has been set.
func (o *TopLevelDomain) GetProlongPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProlongPrice, true
}

// SetProlongPrice sets field value
func (o *TopLevelDomain) SetProlongPrice(v float32) {
	o.ProlongPrice = v
}

// GetRegistrar returns the Registrar field value
func (o *TopLevelDomain) GetRegistrar() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Registrar
}

// GetRegistrarOk returns a tuple with the Registrar field value
// and a boolean to check if the value has been set.
func (o *TopLevelDomain) GetRegistrarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Registrar, true
}

// SetRegistrar sets field value
func (o *TopLevelDomain) SetRegistrar(v string) {
	o.Registrar = v
}

// GetTransfer returns the Transfer field value
func (o *TopLevelDomain) GetTransfer() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Transfer
}

// GetTransferOk returns a tuple with the Transfer field value
// and a boolean to check if the value has been set.
func (o *TopLevelDomain) GetTransferOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transfer, true
}

// SetTransfer sets field value
func (o *TopLevelDomain) SetTransfer(v float32) {
	o.Transfer = v
}

// GetWhoisPrivacyPrice returns the WhoisPrivacyPrice field value
func (o *TopLevelDomain) GetWhoisPrivacyPrice() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.WhoisPrivacyPrice
}

// GetWhoisPrivacyPriceOk returns a tuple with the WhoisPrivacyPrice field value
// and a boolean to check if the value has been set.
func (o *TopLevelDomain) GetWhoisPrivacyPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WhoisPrivacyPrice, true
}

// SetWhoisPrivacyPrice sets field value
func (o *TopLevelDomain) SetWhoisPrivacyPrice(v float32) {
	o.WhoisPrivacyPrice = v
}

func (o TopLevelDomain) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TopLevelDomain) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["allowed_buy_periods"] = o.AllowedBuyPeriods
	toSerialize["early_renew_period"] = o.EarlyRenewPeriod.Get()
	toSerialize["grace_period"] = o.GracePeriod
	toSerialize["id"] = o.Id
	toSerialize["is_published"] = o.IsPublished
	toSerialize["is_registered"] = o.IsRegistered
	toSerialize["is_whois_privacy_default_enabled"] = o.IsWhoisPrivacyDefaultEnabled
	toSerialize["is_whois_privacy_enabled"] = o.IsWhoisPrivacyEnabled
	toSerialize["name"] = o.Name
	toSerialize["price"] = o.Price
	toSerialize["prolong_price"] = o.ProlongPrice
	toSerialize["registrar"] = o.Registrar
	toSerialize["transfer"] = o.Transfer
	toSerialize["whois_privacy_price"] = o.WhoisPrivacyPrice
	return toSerialize, nil
}

type NullableTopLevelDomain struct {
	value *TopLevelDomain
	isSet bool
}

func (v NullableTopLevelDomain) Get() *TopLevelDomain {
	return v.value
}

func (v *NullableTopLevelDomain) Set(val *TopLevelDomain) {
	v.value = val
	v.isSet = true
}

func (v NullableTopLevelDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableTopLevelDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopLevelDomain(val *TopLevelDomain) *NullableTopLevelDomain {
	return &NullableTopLevelDomain{value: val, isSet: true}
}

func (v NullableTopLevelDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopLevelDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


