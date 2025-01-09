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

// checks if the Mailbox type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Mailbox{}

// Mailbox Почтовый ящик
type Mailbox struct {
	AutoReply MailboxAutoReply `json:"auto_reply"`
	SpamFilter MailboxSpamFilter `json:"spam_filter"`
	ForwardingIncoming MailboxForwardingIncoming `json:"forwarding_incoming"`
	ForwardingOutgoing MailboxForwardingOutgoing `json:"forwarding_outgoing"`
	// Комментарий к почтовому ящику
	Comment string `json:"comment"`
	// Домен почты
	Fqdn string `json:"fqdn"`
	// Название почтового ящика
	Mailbox string `json:"mailbox"`
	// Пароль почтового ящика
	Password string `json:"password"`
	// Использованное место на почтовом ящике (в Мб)
	UsageSpace float32 `json:"usage_space"`
	// Доступен ли Webmail
	IsWebmail bool `json:"is_webmail"`
	// IDN домен почтового ящика
	IdnName string `json:"idn_name"`
	// Есть ли доступ через dovecot
	IsDovecot bool `json:"is_dovecot"`
}

// NewMailbox instantiates a new Mailbox object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMailbox(autoReply MailboxAutoReply, spamFilter MailboxSpamFilter, forwardingIncoming MailboxForwardingIncoming, forwardingOutgoing MailboxForwardingOutgoing, comment string, fqdn string, mailbox string, password string, usageSpace float32, isWebmail bool, idnName string, isDovecot bool) *Mailbox {
	this := Mailbox{}
	this.AutoReply = autoReply
	this.SpamFilter = spamFilter
	this.ForwardingIncoming = forwardingIncoming
	this.ForwardingOutgoing = forwardingOutgoing
	this.Comment = comment
	this.Fqdn = fqdn
	this.Mailbox = mailbox
	this.Password = password
	this.UsageSpace = usageSpace
	this.IsWebmail = isWebmail
	this.IdnName = idnName
	this.IsDovecot = isDovecot
	return &this
}

// NewMailboxWithDefaults instantiates a new Mailbox object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMailboxWithDefaults() *Mailbox {
	this := Mailbox{}
	return &this
}

// GetAutoReply returns the AutoReply field value
func (o *Mailbox) GetAutoReply() MailboxAutoReply {
	if o == nil {
		var ret MailboxAutoReply
		return ret
	}

	return o.AutoReply
}

// GetAutoReplyOk returns a tuple with the AutoReply field value
// and a boolean to check if the value has been set.
func (o *Mailbox) GetAutoReplyOk() (*MailboxAutoReply, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoReply, true
}

// SetAutoReply sets field value
func (o *Mailbox) SetAutoReply(v MailboxAutoReply) {
	o.AutoReply = v
}

// GetSpamFilter returns the SpamFilter field value
func (o *Mailbox) GetSpamFilter() MailboxSpamFilter {
	if o == nil {
		var ret MailboxSpamFilter
		return ret
	}

	return o.SpamFilter
}

// GetSpamFilterOk returns a tuple with the SpamFilter field value
// and a boolean to check if the value has been set.
func (o *Mailbox) GetSpamFilterOk() (*MailboxSpamFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpamFilter, true
}

// SetSpamFilter sets field value
func (o *Mailbox) SetSpamFilter(v MailboxSpamFilter) {
	o.SpamFilter = v
}

// GetForwardingIncoming returns the ForwardingIncoming field value
func (o *Mailbox) GetForwardingIncoming() MailboxForwardingIncoming {
	if o == nil {
		var ret MailboxForwardingIncoming
		return ret
	}

	return o.ForwardingIncoming
}

// GetForwardingIncomingOk returns a tuple with the ForwardingIncoming field value
// and a boolean to check if the value has been set.
func (o *Mailbox) GetForwardingIncomingOk() (*MailboxForwardingIncoming, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ForwardingIncoming, true
}

// SetForwardingIncoming sets field value
func (o *Mailbox) SetForwardingIncoming(v MailboxForwardingIncoming) {
	o.ForwardingIncoming = v
}

// GetForwardingOutgoing returns the ForwardingOutgoing field value
func (o *Mailbox) GetForwardingOutgoing() MailboxForwardingOutgoing {
	if o == nil {
		var ret MailboxForwardingOutgoing
		return ret
	}

	return o.ForwardingOutgoing
}

// GetForwardingOutgoingOk returns a tuple with the ForwardingOutgoing field value
// and a boolean to check if the value has been set.
func (o *Mailbox) GetForwardingOutgoingOk() (*MailboxForwardingOutgoing, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ForwardingOutgoing, true
}

// SetForwardingOutgoing sets field value
func (o *Mailbox) SetForwardingOutgoing(v MailboxForwardingOutgoing) {
	o.ForwardingOutgoing = v
}

// GetComment returns the Comment field value
func (o *Mailbox) GetComment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Comment
}

// GetCommentOk returns a tuple with the Comment field value
// and a boolean to check if the value has been set.
func (o *Mailbox) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comment, true
}

// SetComment sets field value
func (o *Mailbox) SetComment(v string) {
	o.Comment = v
}

// GetFqdn returns the Fqdn field value
func (o *Mailbox) GetFqdn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value
// and a boolean to check if the value has been set.
func (o *Mailbox) GetFqdnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fqdn, true
}

// SetFqdn sets field value
func (o *Mailbox) SetFqdn(v string) {
	o.Fqdn = v
}

// GetMailbox returns the Mailbox field value
func (o *Mailbox) GetMailbox() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mailbox
}

// GetMailboxOk returns a tuple with the Mailbox field value
// and a boolean to check if the value has been set.
func (o *Mailbox) GetMailboxOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mailbox, true
}

// SetMailbox sets field value
func (o *Mailbox) SetMailbox(v string) {
	o.Mailbox = v
}

// GetPassword returns the Password field value
func (o *Mailbox) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *Mailbox) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *Mailbox) SetPassword(v string) {
	o.Password = v
}

// GetUsageSpace returns the UsageSpace field value
func (o *Mailbox) GetUsageSpace() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.UsageSpace
}

// GetUsageSpaceOk returns a tuple with the UsageSpace field value
// and a boolean to check if the value has been set.
func (o *Mailbox) GetUsageSpaceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsageSpace, true
}

// SetUsageSpace sets field value
func (o *Mailbox) SetUsageSpace(v float32) {
	o.UsageSpace = v
}

// GetIsWebmail returns the IsWebmail field value
func (o *Mailbox) GetIsWebmail() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsWebmail
}

// GetIsWebmailOk returns a tuple with the IsWebmail field value
// and a boolean to check if the value has been set.
func (o *Mailbox) GetIsWebmailOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsWebmail, true
}

// SetIsWebmail sets field value
func (o *Mailbox) SetIsWebmail(v bool) {
	o.IsWebmail = v
}

// GetIdnName returns the IdnName field value
func (o *Mailbox) GetIdnName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdnName
}

// GetIdnNameOk returns a tuple with the IdnName field value
// and a boolean to check if the value has been set.
func (o *Mailbox) GetIdnNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdnName, true
}

// SetIdnName sets field value
func (o *Mailbox) SetIdnName(v string) {
	o.IdnName = v
}

// GetIsDovecot returns the IsDovecot field value
func (o *Mailbox) GetIsDovecot() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDovecot
}

// GetIsDovecotOk returns a tuple with the IsDovecot field value
// and a boolean to check if the value has been set.
func (o *Mailbox) GetIsDovecotOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDovecot, true
}

// SetIsDovecot sets field value
func (o *Mailbox) SetIsDovecot(v bool) {
	o.IsDovecot = v
}

func (o Mailbox) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Mailbox) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["auto_reply"] = o.AutoReply
	toSerialize["spam_filter"] = o.SpamFilter
	toSerialize["forwarding_incoming"] = o.ForwardingIncoming
	toSerialize["forwarding_outgoing"] = o.ForwardingOutgoing
	toSerialize["comment"] = o.Comment
	toSerialize["fqdn"] = o.Fqdn
	toSerialize["mailbox"] = o.Mailbox
	toSerialize["password"] = o.Password
	toSerialize["usage_space"] = o.UsageSpace
	toSerialize["is_webmail"] = o.IsWebmail
	toSerialize["idn_name"] = o.IdnName
	toSerialize["is_dovecot"] = o.IsDovecot
	return toSerialize, nil
}

type NullableMailbox struct {
	value *Mailbox
	isSet bool
}

func (v NullableMailbox) Get() *Mailbox {
	return v.value
}

func (v *NullableMailbox) Set(val *Mailbox) {
	v.value = val
	v.isSet = true
}

func (v NullableMailbox) IsSet() bool {
	return v.isSet
}

func (v *NullableMailbox) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMailbox(val *Mailbox) *NullableMailbox {
	return &NullableMailbox{value: val, isSet: true}
}

func (v NullableMailbox) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMailbox) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


