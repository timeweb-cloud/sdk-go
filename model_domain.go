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

// checks if the Domain type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Domain{}

// Domain Домен
type Domain struct {
	// Допустимые периоды продления домена.
	AllowedBuyPeriods []DomainAllowedBuyPeriodsInner `json:"allowed_buy_periods"`
	// Количество дней, оставшихся до конца срока регистрации домена.
	DaysLeft float32 `json:"days_left"`
	// Статус домена.
	DomainStatus string `json:"domain_status"`
	// Дата окончания срока регистрации домена, для доменов без срока окончания регистрации будет приходить 0000-00-00.
	Expiration string `json:"expiration"`
	// Полное имя домена.
	Fqdn string `json:"fqdn"`
	// ID домена.
	Id float32 `json:"id"`
	// Это логическое значение, которое показывает, включено ли автопродление домена.
	IsAutoprolongEnabled NullableBool `json:"is_autoprolong_enabled"`
	// Это логическое значение, которое показывает, является ли домен премиальным.
	IsPremium bool `json:"is_premium"`
	// Это логическое значение, которое показывает, можно ли сейчас продлить домен.
	IsProlongAllowed bool `json:"is_prolong_allowed"`
	// Это логическое значение, которое показывает, является ли домен техническим.
	IsTechnical bool `json:"is_technical"`
	// Это логическое значение, которое показывает, включено ли скрытие данных администратора домена для whois. Если приходит null, значит для данной зоны эта услуга не доступна.
	IsWhoisPrivacyEnabled NullableBool `json:"is_whois_privacy_enabled"`
	// Привязанный к домену IP-адрес.
	LinkedIp NullableString `json:"linked_ip"`
	// До какого числа оплачен домен.
	PaidTill NullableString `json:"paid_till"`
	// ID администратора, на которого зарегистрирован домен.
	PersonId NullableFloat32 `json:"person_id"`
	// Стоимость премиального домена.
	PremiumProlongCost NullableFloat32 `json:"premium_prolong_cost"`
	// ID регистратора домена.
	Provider NullableString `json:"provider"`
	// Статус заявки на продление/регистрацию/трансфер домена.
	RequestStatus NullableString `json:"request_status"`
	// Список поддоменов.
	Subdomains []Subdomain `json:"subdomains"`
	// ID доменной зоны.
	TldId NullableFloat32 `json:"tld_id"`
}

// NewDomain instantiates a new Domain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomain(allowedBuyPeriods []DomainAllowedBuyPeriodsInner, daysLeft float32, domainStatus string, expiration string, fqdn string, id float32, isAutoprolongEnabled NullableBool, isPremium bool, isProlongAllowed bool, isTechnical bool, isWhoisPrivacyEnabled NullableBool, linkedIp NullableString, paidTill NullableString, personId NullableFloat32, premiumProlongCost NullableFloat32, provider NullableString, requestStatus NullableString, subdomains []Subdomain, tldId NullableFloat32) *Domain {
	this := Domain{}
	this.AllowedBuyPeriods = allowedBuyPeriods
	this.DaysLeft = daysLeft
	this.DomainStatus = domainStatus
	this.Expiration = expiration
	this.Fqdn = fqdn
	this.Id = id
	this.IsAutoprolongEnabled = isAutoprolongEnabled
	this.IsPremium = isPremium
	this.IsProlongAllowed = isProlongAllowed
	this.IsTechnical = isTechnical
	this.IsWhoisPrivacyEnabled = isWhoisPrivacyEnabled
	this.LinkedIp = linkedIp
	this.PaidTill = paidTill
	this.PersonId = personId
	this.PremiumProlongCost = premiumProlongCost
	this.Provider = provider
	this.RequestStatus = requestStatus
	this.Subdomains = subdomains
	this.TldId = tldId
	return &this
}

// NewDomainWithDefaults instantiates a new Domain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainWithDefaults() *Domain {
	this := Domain{}
	return &this
}

// GetAllowedBuyPeriods returns the AllowedBuyPeriods field value
func (o *Domain) GetAllowedBuyPeriods() []DomainAllowedBuyPeriodsInner {
	if o == nil {
		var ret []DomainAllowedBuyPeriodsInner
		return ret
	}

	return o.AllowedBuyPeriods
}

// GetAllowedBuyPeriodsOk returns a tuple with the AllowedBuyPeriods field value
// and a boolean to check if the value has been set.
func (o *Domain) GetAllowedBuyPeriodsOk() ([]DomainAllowedBuyPeriodsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowedBuyPeriods, true
}

// SetAllowedBuyPeriods sets field value
func (o *Domain) SetAllowedBuyPeriods(v []DomainAllowedBuyPeriodsInner) {
	o.AllowedBuyPeriods = v
}

// GetDaysLeft returns the DaysLeft field value
func (o *Domain) GetDaysLeft() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.DaysLeft
}

// GetDaysLeftOk returns a tuple with the DaysLeft field value
// and a boolean to check if the value has been set.
func (o *Domain) GetDaysLeftOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DaysLeft, true
}

// SetDaysLeft sets field value
func (o *Domain) SetDaysLeft(v float32) {
	o.DaysLeft = v
}

// GetDomainStatus returns the DomainStatus field value
func (o *Domain) GetDomainStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DomainStatus
}

// GetDomainStatusOk returns a tuple with the DomainStatus field value
// and a boolean to check if the value has been set.
func (o *Domain) GetDomainStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DomainStatus, true
}

// SetDomainStatus sets field value
func (o *Domain) SetDomainStatus(v string) {
	o.DomainStatus = v
}

// GetExpiration returns the Expiration field value
func (o *Domain) GetExpiration() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value
// and a boolean to check if the value has been set.
func (o *Domain) GetExpirationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expiration, true
}

// SetExpiration sets field value
func (o *Domain) SetExpiration(v string) {
	o.Expiration = v
}

// GetFqdn returns the Fqdn field value
func (o *Domain) GetFqdn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value
// and a boolean to check if the value has been set.
func (o *Domain) GetFqdnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fqdn, true
}

// SetFqdn sets field value
func (o *Domain) SetFqdn(v string) {
	o.Fqdn = v
}

// GetId returns the Id field value
func (o *Domain) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Domain) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Domain) SetId(v float32) {
	o.Id = v
}

// GetIsAutoprolongEnabled returns the IsAutoprolongEnabled field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *Domain) GetIsAutoprolongEnabled() bool {
	if o == nil || o.IsAutoprolongEnabled.Get() == nil {
		var ret bool
		return ret
	}

	return *o.IsAutoprolongEnabled.Get()
}

// GetIsAutoprolongEnabledOk returns a tuple with the IsAutoprolongEnabled field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Domain) GetIsAutoprolongEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsAutoprolongEnabled.Get(), o.IsAutoprolongEnabled.IsSet()
}

// SetIsAutoprolongEnabled sets field value
func (o *Domain) SetIsAutoprolongEnabled(v bool) {
	o.IsAutoprolongEnabled.Set(&v)
}

// GetIsPremium returns the IsPremium field value
func (o *Domain) GetIsPremium() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPremium
}

// GetIsPremiumOk returns a tuple with the IsPremium field value
// and a boolean to check if the value has been set.
func (o *Domain) GetIsPremiumOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPremium, true
}

// SetIsPremium sets field value
func (o *Domain) SetIsPremium(v bool) {
	o.IsPremium = v
}

// GetIsProlongAllowed returns the IsProlongAllowed field value
func (o *Domain) GetIsProlongAllowed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsProlongAllowed
}

// GetIsProlongAllowedOk returns a tuple with the IsProlongAllowed field value
// and a boolean to check if the value has been set.
func (o *Domain) GetIsProlongAllowedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsProlongAllowed, true
}

// SetIsProlongAllowed sets field value
func (o *Domain) SetIsProlongAllowed(v bool) {
	o.IsProlongAllowed = v
}

// GetIsTechnical returns the IsTechnical field value
func (o *Domain) GetIsTechnical() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsTechnical
}

// GetIsTechnicalOk returns a tuple with the IsTechnical field value
// and a boolean to check if the value has been set.
func (o *Domain) GetIsTechnicalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsTechnical, true
}

// SetIsTechnical sets field value
func (o *Domain) SetIsTechnical(v bool) {
	o.IsTechnical = v
}

// GetIsWhoisPrivacyEnabled returns the IsWhoisPrivacyEnabled field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *Domain) GetIsWhoisPrivacyEnabled() bool {
	if o == nil || o.IsWhoisPrivacyEnabled.Get() == nil {
		var ret bool
		return ret
	}

	return *o.IsWhoisPrivacyEnabled.Get()
}

// GetIsWhoisPrivacyEnabledOk returns a tuple with the IsWhoisPrivacyEnabled field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Domain) GetIsWhoisPrivacyEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsWhoisPrivacyEnabled.Get(), o.IsWhoisPrivacyEnabled.IsSet()
}

// SetIsWhoisPrivacyEnabled sets field value
func (o *Domain) SetIsWhoisPrivacyEnabled(v bool) {
	o.IsWhoisPrivacyEnabled.Set(&v)
}

// GetLinkedIp returns the LinkedIp field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Domain) GetLinkedIp() string {
	if o == nil || o.LinkedIp.Get() == nil {
		var ret string
		return ret
	}

	return *o.LinkedIp.Get()
}

// GetLinkedIpOk returns a tuple with the LinkedIp field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Domain) GetLinkedIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LinkedIp.Get(), o.LinkedIp.IsSet()
}

// SetLinkedIp sets field value
func (o *Domain) SetLinkedIp(v string) {
	o.LinkedIp.Set(&v)
}

// GetPaidTill returns the PaidTill field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Domain) GetPaidTill() string {
	if o == nil || o.PaidTill.Get() == nil {
		var ret string
		return ret
	}

	return *o.PaidTill.Get()
}

// GetPaidTillOk returns a tuple with the PaidTill field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Domain) GetPaidTillOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaidTill.Get(), o.PaidTill.IsSet()
}

// SetPaidTill sets field value
func (o *Domain) SetPaidTill(v string) {
	o.PaidTill.Set(&v)
}

// GetPersonId returns the PersonId field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *Domain) GetPersonId() float32 {
	if o == nil || o.PersonId.Get() == nil {
		var ret float32
		return ret
	}

	return *o.PersonId.Get()
}

// GetPersonIdOk returns a tuple with the PersonId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Domain) GetPersonIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PersonId.Get(), o.PersonId.IsSet()
}

// SetPersonId sets field value
func (o *Domain) SetPersonId(v float32) {
	o.PersonId.Set(&v)
}

// GetPremiumProlongCost returns the PremiumProlongCost field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *Domain) GetPremiumProlongCost() float32 {
	if o == nil || o.PremiumProlongCost.Get() == nil {
		var ret float32
		return ret
	}

	return *o.PremiumProlongCost.Get()
}

// GetPremiumProlongCostOk returns a tuple with the PremiumProlongCost field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Domain) GetPremiumProlongCostOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PremiumProlongCost.Get(), o.PremiumProlongCost.IsSet()
}

// SetPremiumProlongCost sets field value
func (o *Domain) SetPremiumProlongCost(v float32) {
	o.PremiumProlongCost.Set(&v)
}

// GetProvider returns the Provider field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Domain) GetProvider() string {
	if o == nil || o.Provider.Get() == nil {
		var ret string
		return ret
	}

	return *o.Provider.Get()
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Domain) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Provider.Get(), o.Provider.IsSet()
}

// SetProvider sets field value
func (o *Domain) SetProvider(v string) {
	o.Provider.Set(&v)
}

// GetRequestStatus returns the RequestStatus field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Domain) GetRequestStatus() string {
	if o == nil || o.RequestStatus.Get() == nil {
		var ret string
		return ret
	}

	return *o.RequestStatus.Get()
}

// GetRequestStatusOk returns a tuple with the RequestStatus field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Domain) GetRequestStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestStatus.Get(), o.RequestStatus.IsSet()
}

// SetRequestStatus sets field value
func (o *Domain) SetRequestStatus(v string) {
	o.RequestStatus.Set(&v)
}

// GetSubdomains returns the Subdomains field value
func (o *Domain) GetSubdomains() []Subdomain {
	if o == nil {
		var ret []Subdomain
		return ret
	}

	return o.Subdomains
}

// GetSubdomainsOk returns a tuple with the Subdomains field value
// and a boolean to check if the value has been set.
func (o *Domain) GetSubdomainsOk() ([]Subdomain, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subdomains, true
}

// SetSubdomains sets field value
func (o *Domain) SetSubdomains(v []Subdomain) {
	o.Subdomains = v
}

// GetTldId returns the TldId field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *Domain) GetTldId() float32 {
	if o == nil || o.TldId.Get() == nil {
		var ret float32
		return ret
	}

	return *o.TldId.Get()
}

// GetTldIdOk returns a tuple with the TldId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Domain) GetTldIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TldId.Get(), o.TldId.IsSet()
}

// SetTldId sets field value
func (o *Domain) SetTldId(v float32) {
	o.TldId.Set(&v)
}

func (o Domain) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Domain) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["allowed_buy_periods"] = o.AllowedBuyPeriods
	toSerialize["days_left"] = o.DaysLeft
	toSerialize["domain_status"] = o.DomainStatus
	toSerialize["expiration"] = o.Expiration
	toSerialize["fqdn"] = o.Fqdn
	toSerialize["id"] = o.Id
	toSerialize["is_autoprolong_enabled"] = o.IsAutoprolongEnabled.Get()
	toSerialize["is_premium"] = o.IsPremium
	toSerialize["is_prolong_allowed"] = o.IsProlongAllowed
	toSerialize["is_technical"] = o.IsTechnical
	toSerialize["is_whois_privacy_enabled"] = o.IsWhoisPrivacyEnabled.Get()
	toSerialize["linked_ip"] = o.LinkedIp.Get()
	toSerialize["paid_till"] = o.PaidTill.Get()
	toSerialize["person_id"] = o.PersonId.Get()
	toSerialize["premium_prolong_cost"] = o.PremiumProlongCost.Get()
	toSerialize["provider"] = o.Provider.Get()
	toSerialize["request_status"] = o.RequestStatus.Get()
	toSerialize["subdomains"] = o.Subdomains
	toSerialize["tld_id"] = o.TldId.Get()
	return toSerialize, nil
}

type NullableDomain struct {
	value *Domain
	isSet bool
}

func (v NullableDomain) Get() *Domain {
	return v.value
}

func (v *NullableDomain) Set(val *Domain) {
	v.value = val
	v.isSet = true
}

func (v NullableDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomain(val *Domain) *NullableDomain {
	return &NullableDomain{value: val, isSet: true}
}

func (v NullableDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


