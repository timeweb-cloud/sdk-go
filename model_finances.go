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

// checks if the Finances type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Finances{}

// Finances Платежная информация
type Finances struct {
	// Баланс аккаунта.
	Balance float32 `json:"balance"`
	// Валюта, которая используется на аккаунте.
	Currency string `json:"currency"`
	// Значение времени, указанное в комбинированном формате даты и времени ISO8601, которое представляет, когда заканчивается скидка для аккаунта.
	DiscountEndDateAt NullableString `json:"discount_end_date_at"`
	// Процент скидки для аккаунта.
	DiscountPercent float32 `json:"discount_percent"`
	// Стоимость услуг на аккаунте в час.
	HourlyCost float32 `json:"hourly_cost"`
	// Абонентская плата в час (с учетом скидок).
	HourlyFee float32 `json:"hourly_fee"`
	// Стоимость услуг на аккаунте в месяц.
	MonthlyCost float32 `json:"monthly_cost"`
	// Абонентская плата в месяц (с учетом скидок).
	MonthlyFee float32 `json:"monthly_fee"`
	// Общая сумма платежей на аккаунте.
	TotalPaid float32 `json:"total_paid"`
	// Сколько часов работы услуг оплачено на аккаунте.
	HoursLeft NullableFloat32 `json:"hours_left"`
	// Информация о карте, используемой для автоплатежей.
	AutopayCardInfo NullableString `json:"autopay_card_info"`
}

// NewFinances instantiates a new Finances object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinances(balance float32, currency string, discountEndDateAt NullableString, discountPercent float32, hourlyCost float32, hourlyFee float32, monthlyCost float32, monthlyFee float32, totalPaid float32, hoursLeft NullableFloat32, autopayCardInfo NullableString) *Finances {
	this := Finances{}
	this.Balance = balance
	this.Currency = currency
	this.DiscountEndDateAt = discountEndDateAt
	this.DiscountPercent = discountPercent
	this.HourlyCost = hourlyCost
	this.HourlyFee = hourlyFee
	this.MonthlyCost = monthlyCost
	this.MonthlyFee = monthlyFee
	this.TotalPaid = totalPaid
	this.HoursLeft = hoursLeft
	this.AutopayCardInfo = autopayCardInfo
	return &this
}

// NewFinancesWithDefaults instantiates a new Finances object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinancesWithDefaults() *Finances {
	this := Finances{}
	return &this
}

// GetBalance returns the Balance field value
func (o *Finances) GetBalance() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value
// and a boolean to check if the value has been set.
func (o *Finances) GetBalanceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Balance, true
}

// SetBalance sets field value
func (o *Finances) SetBalance(v float32) {
	o.Balance = v
}

// GetCurrency returns the Currency field value
func (o *Finances) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *Finances) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *Finances) SetCurrency(v string) {
	o.Currency = v
}

// GetDiscountEndDateAt returns the DiscountEndDateAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Finances) GetDiscountEndDateAt() string {
	if o == nil || o.DiscountEndDateAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.DiscountEndDateAt.Get()
}

// GetDiscountEndDateAtOk returns a tuple with the DiscountEndDateAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Finances) GetDiscountEndDateAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiscountEndDateAt.Get(), o.DiscountEndDateAt.IsSet()
}

// SetDiscountEndDateAt sets field value
func (o *Finances) SetDiscountEndDateAt(v string) {
	o.DiscountEndDateAt.Set(&v)
}

// GetDiscountPercent returns the DiscountPercent field value
func (o *Finances) GetDiscountPercent() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.DiscountPercent
}

// GetDiscountPercentOk returns a tuple with the DiscountPercent field value
// and a boolean to check if the value has been set.
func (o *Finances) GetDiscountPercentOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiscountPercent, true
}

// SetDiscountPercent sets field value
func (o *Finances) SetDiscountPercent(v float32) {
	o.DiscountPercent = v
}

// GetHourlyCost returns the HourlyCost field value
func (o *Finances) GetHourlyCost() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.HourlyCost
}

// GetHourlyCostOk returns a tuple with the HourlyCost field value
// and a boolean to check if the value has been set.
func (o *Finances) GetHourlyCostOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HourlyCost, true
}

// SetHourlyCost sets field value
func (o *Finances) SetHourlyCost(v float32) {
	o.HourlyCost = v
}

// GetHourlyFee returns the HourlyFee field value
func (o *Finances) GetHourlyFee() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.HourlyFee
}

// GetHourlyFeeOk returns a tuple with the HourlyFee field value
// and a boolean to check if the value has been set.
func (o *Finances) GetHourlyFeeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HourlyFee, true
}

// SetHourlyFee sets field value
func (o *Finances) SetHourlyFee(v float32) {
	o.HourlyFee = v
}

// GetMonthlyCost returns the MonthlyCost field value
func (o *Finances) GetMonthlyCost() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MonthlyCost
}

// GetMonthlyCostOk returns a tuple with the MonthlyCost field value
// and a boolean to check if the value has been set.
func (o *Finances) GetMonthlyCostOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonthlyCost, true
}

// SetMonthlyCost sets field value
func (o *Finances) SetMonthlyCost(v float32) {
	o.MonthlyCost = v
}

// GetMonthlyFee returns the MonthlyFee field value
func (o *Finances) GetMonthlyFee() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MonthlyFee
}

// GetMonthlyFeeOk returns a tuple with the MonthlyFee field value
// and a boolean to check if the value has been set.
func (o *Finances) GetMonthlyFeeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonthlyFee, true
}

// SetMonthlyFee sets field value
func (o *Finances) SetMonthlyFee(v float32) {
	o.MonthlyFee = v
}

// GetTotalPaid returns the TotalPaid field value
func (o *Finances) GetTotalPaid() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalPaid
}

// GetTotalPaidOk returns a tuple with the TotalPaid field value
// and a boolean to check if the value has been set.
func (o *Finances) GetTotalPaidOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalPaid, true
}

// SetTotalPaid sets field value
func (o *Finances) SetTotalPaid(v float32) {
	o.TotalPaid = v
}

// GetHoursLeft returns the HoursLeft field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *Finances) GetHoursLeft() float32 {
	if o == nil || o.HoursLeft.Get() == nil {
		var ret float32
		return ret
	}

	return *o.HoursLeft.Get()
}

// GetHoursLeftOk returns a tuple with the HoursLeft field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Finances) GetHoursLeftOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.HoursLeft.Get(), o.HoursLeft.IsSet()
}

// SetHoursLeft sets field value
func (o *Finances) SetHoursLeft(v float32) {
	o.HoursLeft.Set(&v)
}

// GetAutopayCardInfo returns the AutopayCardInfo field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Finances) GetAutopayCardInfo() string {
	if o == nil || o.AutopayCardInfo.Get() == nil {
		var ret string
		return ret
	}

	return *o.AutopayCardInfo.Get()
}

// GetAutopayCardInfoOk returns a tuple with the AutopayCardInfo field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Finances) GetAutopayCardInfoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutopayCardInfo.Get(), o.AutopayCardInfo.IsSet()
}

// SetAutopayCardInfo sets field value
func (o *Finances) SetAutopayCardInfo(v string) {
	o.AutopayCardInfo.Set(&v)
}

func (o Finances) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Finances) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["balance"] = o.Balance
	toSerialize["currency"] = o.Currency
	toSerialize["discount_end_date_at"] = o.DiscountEndDateAt.Get()
	toSerialize["discount_percent"] = o.DiscountPercent
	toSerialize["hourly_cost"] = o.HourlyCost
	toSerialize["hourly_fee"] = o.HourlyFee
	toSerialize["monthly_cost"] = o.MonthlyCost
	toSerialize["monthly_fee"] = o.MonthlyFee
	toSerialize["total_paid"] = o.TotalPaid
	toSerialize["hours_left"] = o.HoursLeft.Get()
	toSerialize["autopay_card_info"] = o.AutopayCardInfo.Get()
	return toSerialize, nil
}

type NullableFinances struct {
	value *Finances
	isSet bool
}

func (v NullableFinances) Get() *Finances {
	return v.value
}

func (v *NullableFinances) Set(val *Finances) {
	v.value = val
	v.isSet = true
}

func (v NullableFinances) IsSet() bool {
	return v.isSet
}

func (v *NullableFinances) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinances(val *Finances) *NullableFinances {
	return &NullableFinances{value: val, isSet: true}
}

func (v NullableFinances) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinances) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


