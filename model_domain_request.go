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
	"time"
)

// checks if the DomainRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DomainRequest{}

// DomainRequest Заявка на продление/регистрацию/трансфер домена.
type DomainRequest struct {
	// ID пользователя
	AccountId string `json:"account_id"`
	// Код авторизации для переноса домена.
	AuthCode NullableString `json:"auth_code"`
	// Дата создания заявки.
	Date time.Time `json:"date"`
	// Идентификационный номер бандла, в который входит данная заявка (null - если заявка не входит ни в один бандл).
	DomainBundleId NullableString `json:"domain_bundle_id"`
	// Код ошибки трансфера домена.
	ErrorCodeTransfer NullableString `json:"error_code_transfer"`
	// Полное имя домена.
	Fqdn string `json:"fqdn"`
	// ID группы доменных зон.
	GroupId float32 `json:"group_id"`
	// ID заявки.
	Id float32 `json:"id"`
	// Это логическое значение, которое показывает включена ли услуга \"Антиспам\" для домена
	IsAntispamEnabled bool `json:"is_antispam_enabled"`
	// Это логическое значение, которое показывает, включено ли автопродление домена.
	IsAutoprolongEnabled bool `json:"is_autoprolong_enabled"`
	// Это логическое значение, которое показывает, включено ли скрытие данных администратора домена для whois. Опция недоступна для доменов в зонах .ru и .рф.
	IsWhoisPrivacyEnabled bool `json:"is_whois_privacy_enabled"`
	// Информационное сообщение о заявке.
	Message NullableString `json:"message"`
	// Источник (способ) оплаты заявки.
	MoneySource NullableString `json:"money_source"`
	Period DomainPaymentPeriod `json:"period"`
	// Идентификационный номер персоны для заявки на регистрацию.
	PersonId float32 `json:"person_id"`
	Prime DomainPrimeType `json:"prime"`
	// Количество дней до конца регистрации домена, за которые мы уведомим о необходимости продления.
	SoonExpire float32 `json:"soon_expire"`
	// Это значение используется для сортировки доменных зон в панели управления.
	SortOrder float32 `json:"sort_order"`
	// Тип заявки.
	Type string `json:"type"`
}

// NewDomainRequest instantiates a new DomainRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainRequest(accountId string, authCode NullableString, date time.Time, domainBundleId NullableString, errorCodeTransfer NullableString, fqdn string, groupId float32, id float32, isAntispamEnabled bool, isAutoprolongEnabled bool, isWhoisPrivacyEnabled bool, message NullableString, moneySource NullableString, period DomainPaymentPeriod, personId float32, prime DomainPrimeType, soonExpire float32, sortOrder float32, type_ string) *DomainRequest {
	this := DomainRequest{}
	this.AccountId = accountId
	this.AuthCode = authCode
	this.Date = date
	this.DomainBundleId = domainBundleId
	this.ErrorCodeTransfer = errorCodeTransfer
	this.Fqdn = fqdn
	this.GroupId = groupId
	this.Id = id
	this.IsAntispamEnabled = isAntispamEnabled
	this.IsAutoprolongEnabled = isAutoprolongEnabled
	this.IsWhoisPrivacyEnabled = isWhoisPrivacyEnabled
	this.Message = message
	this.MoneySource = moneySource
	this.Period = period
	this.PersonId = personId
	this.Prime = prime
	this.SoonExpire = soonExpire
	this.SortOrder = sortOrder
	this.Type = type_
	return &this
}

// NewDomainRequestWithDefaults instantiates a new DomainRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainRequestWithDefaults() *DomainRequest {
	this := DomainRequest{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *DomainRequest) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *DomainRequest) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *DomainRequest) SetAccountId(v string) {
	o.AccountId = v
}

// GetAuthCode returns the AuthCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DomainRequest) GetAuthCode() string {
	if o == nil || o.AuthCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.AuthCode.Get()
}

// GetAuthCodeOk returns a tuple with the AuthCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DomainRequest) GetAuthCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthCode.Get(), o.AuthCode.IsSet()
}

// SetAuthCode sets field value
func (o *DomainRequest) SetAuthCode(v string) {
	o.AuthCode.Set(&v)
}

// GetDate returns the Date field value
func (o *DomainRequest) GetDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *DomainRequest) GetDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *DomainRequest) SetDate(v time.Time) {
	o.Date = v
}

// GetDomainBundleId returns the DomainBundleId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DomainRequest) GetDomainBundleId() string {
	if o == nil || o.DomainBundleId.Get() == nil {
		var ret string
		return ret
	}

	return *o.DomainBundleId.Get()
}

// GetDomainBundleIdOk returns a tuple with the DomainBundleId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DomainRequest) GetDomainBundleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DomainBundleId.Get(), o.DomainBundleId.IsSet()
}

// SetDomainBundleId sets field value
func (o *DomainRequest) SetDomainBundleId(v string) {
	o.DomainBundleId.Set(&v)
}

// GetErrorCodeTransfer returns the ErrorCodeTransfer field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DomainRequest) GetErrorCodeTransfer() string {
	if o == nil || o.ErrorCodeTransfer.Get() == nil {
		var ret string
		return ret
	}

	return *o.ErrorCodeTransfer.Get()
}

// GetErrorCodeTransferOk returns a tuple with the ErrorCodeTransfer field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DomainRequest) GetErrorCodeTransferOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorCodeTransfer.Get(), o.ErrorCodeTransfer.IsSet()
}

// SetErrorCodeTransfer sets field value
func (o *DomainRequest) SetErrorCodeTransfer(v string) {
	o.ErrorCodeTransfer.Set(&v)
}

// GetFqdn returns the Fqdn field value
func (o *DomainRequest) GetFqdn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value
// and a boolean to check if the value has been set.
func (o *DomainRequest) GetFqdnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fqdn, true
}

// SetFqdn sets field value
func (o *DomainRequest) SetFqdn(v string) {
	o.Fqdn = v
}

// GetGroupId returns the GroupId field value
func (o *DomainRequest) GetGroupId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *DomainRequest) GetGroupIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *DomainRequest) SetGroupId(v float32) {
	o.GroupId = v
}

// GetId returns the Id field value
func (o *DomainRequest) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DomainRequest) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DomainRequest) SetId(v float32) {
	o.Id = v
}

// GetIsAntispamEnabled returns the IsAntispamEnabled field value
func (o *DomainRequest) GetIsAntispamEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsAntispamEnabled
}

// GetIsAntispamEnabledOk returns a tuple with the IsAntispamEnabled field value
// and a boolean to check if the value has been set.
func (o *DomainRequest) GetIsAntispamEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsAntispamEnabled, true
}

// SetIsAntispamEnabled sets field value
func (o *DomainRequest) SetIsAntispamEnabled(v bool) {
	o.IsAntispamEnabled = v
}

// GetIsAutoprolongEnabled returns the IsAutoprolongEnabled field value
func (o *DomainRequest) GetIsAutoprolongEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsAutoprolongEnabled
}

// GetIsAutoprolongEnabledOk returns a tuple with the IsAutoprolongEnabled field value
// and a boolean to check if the value has been set.
func (o *DomainRequest) GetIsAutoprolongEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsAutoprolongEnabled, true
}

// SetIsAutoprolongEnabled sets field value
func (o *DomainRequest) SetIsAutoprolongEnabled(v bool) {
	o.IsAutoprolongEnabled = v
}

// GetIsWhoisPrivacyEnabled returns the IsWhoisPrivacyEnabled field value
func (o *DomainRequest) GetIsWhoisPrivacyEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsWhoisPrivacyEnabled
}

// GetIsWhoisPrivacyEnabledOk returns a tuple with the IsWhoisPrivacyEnabled field value
// and a boolean to check if the value has been set.
func (o *DomainRequest) GetIsWhoisPrivacyEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsWhoisPrivacyEnabled, true
}

// SetIsWhoisPrivacyEnabled sets field value
func (o *DomainRequest) SetIsWhoisPrivacyEnabled(v bool) {
	o.IsWhoisPrivacyEnabled = v
}

// GetMessage returns the Message field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DomainRequest) GetMessage() string {
	if o == nil || o.Message.Get() == nil {
		var ret string
		return ret
	}

	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DomainRequest) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// SetMessage sets field value
func (o *DomainRequest) SetMessage(v string) {
	o.Message.Set(&v)
}

// GetMoneySource returns the MoneySource field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DomainRequest) GetMoneySource() string {
	if o == nil || o.MoneySource.Get() == nil {
		var ret string
		return ret
	}

	return *o.MoneySource.Get()
}

// GetMoneySourceOk returns a tuple with the MoneySource field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DomainRequest) GetMoneySourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MoneySource.Get(), o.MoneySource.IsSet()
}

// SetMoneySource sets field value
func (o *DomainRequest) SetMoneySource(v string) {
	o.MoneySource.Set(&v)
}

// GetPeriod returns the Period field value
func (o *DomainRequest) GetPeriod() DomainPaymentPeriod {
	if o == nil {
		var ret DomainPaymentPeriod
		return ret
	}

	return o.Period
}

// GetPeriodOk returns a tuple with the Period field value
// and a boolean to check if the value has been set.
func (o *DomainRequest) GetPeriodOk() (*DomainPaymentPeriod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Period, true
}

// SetPeriod sets field value
func (o *DomainRequest) SetPeriod(v DomainPaymentPeriod) {
	o.Period = v
}

// GetPersonId returns the PersonId field value
func (o *DomainRequest) GetPersonId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PersonId
}

// GetPersonIdOk returns a tuple with the PersonId field value
// and a boolean to check if the value has been set.
func (o *DomainRequest) GetPersonIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PersonId, true
}

// SetPersonId sets field value
func (o *DomainRequest) SetPersonId(v float32) {
	o.PersonId = v
}

// GetPrime returns the Prime field value
func (o *DomainRequest) GetPrime() DomainPrimeType {
	if o == nil {
		var ret DomainPrimeType
		return ret
	}

	return o.Prime
}

// GetPrimeOk returns a tuple with the Prime field value
// and a boolean to check if the value has been set.
func (o *DomainRequest) GetPrimeOk() (*DomainPrimeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prime, true
}

// SetPrime sets field value
func (o *DomainRequest) SetPrime(v DomainPrimeType) {
	o.Prime = v
}

// GetSoonExpire returns the SoonExpire field value
func (o *DomainRequest) GetSoonExpire() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SoonExpire
}

// GetSoonExpireOk returns a tuple with the SoonExpire field value
// and a boolean to check if the value has been set.
func (o *DomainRequest) GetSoonExpireOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SoonExpire, true
}

// SetSoonExpire sets field value
func (o *DomainRequest) SetSoonExpire(v float32) {
	o.SoonExpire = v
}

// GetSortOrder returns the SortOrder field value
func (o *DomainRequest) GetSortOrder() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value
// and a boolean to check if the value has been set.
func (o *DomainRequest) GetSortOrderOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SortOrder, true
}

// SetSortOrder sets field value
func (o *DomainRequest) SetSortOrder(v float32) {
	o.SortOrder = v
}

// GetType returns the Type field value
func (o *DomainRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DomainRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DomainRequest) SetType(v string) {
	o.Type = v
}

func (o DomainRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DomainRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_id"] = o.AccountId
	toSerialize["auth_code"] = o.AuthCode.Get()
	toSerialize["date"] = o.Date
	toSerialize["domain_bundle_id"] = o.DomainBundleId.Get()
	toSerialize["error_code_transfer"] = o.ErrorCodeTransfer.Get()
	toSerialize["fqdn"] = o.Fqdn
	toSerialize["group_id"] = o.GroupId
	toSerialize["id"] = o.Id
	toSerialize["is_antispam_enabled"] = o.IsAntispamEnabled
	toSerialize["is_autoprolong_enabled"] = o.IsAutoprolongEnabled
	toSerialize["is_whois_privacy_enabled"] = o.IsWhoisPrivacyEnabled
	toSerialize["message"] = o.Message.Get()
	toSerialize["money_source"] = o.MoneySource.Get()
	toSerialize["period"] = o.Period
	toSerialize["person_id"] = o.PersonId
	toSerialize["prime"] = o.Prime
	toSerialize["soon_expire"] = o.SoonExpire
	toSerialize["sort_order"] = o.SortOrder
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableDomainRequest struct {
	value *DomainRequest
	isSet bool
}

func (v NullableDomainRequest) Get() *DomainRequest {
	return v.value
}

func (v *NullableDomainRequest) Set(val *DomainRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainRequest(val *DomainRequest) *NullableDomainRequest {
	return &NullableDomainRequest{value: val, isSet: true}
}

func (v NullableDomainRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


