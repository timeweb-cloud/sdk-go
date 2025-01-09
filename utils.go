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
	"reflect"
	"time"
)

// PtrBool is a helper routine that returns a pointer to given boolean value.
func PtrBool(v bool) *bool { return &v }

// PtrInt is a helper routine that returns a pointer to given integer value.
func PtrInt(v int) *int { return &v }

// PtrInt32 is a helper routine that returns a pointer to given integer value.
func PtrInt32(v int32) *int32 { return &v }

// PtrInt64 is a helper routine that returns a pointer to given integer value.
func PtrInt64(v int64) *int64 { return &v }

// PtrFloat32 is a helper routine that returns a pointer to given float value.
func PtrFloat32(v float32) *float32 { return &v }

// PtrFloat64 is a helper routine that returns a pointer to given float value.
func PtrFloat64(v float64) *float64 { return &v }

// PtrString is a helper routine that returns a pointer to given string value.
func PtrString(v string) *string { return &v }

// PtrTime is helper routine that returns a pointer to given Time value.
func PtrTime(v time.Time) *time.Time { return &v }

type NullableBool struct {
	value *bool
	isSet bool
}

func (v NullableBool) Get() *bool {
	return v.value
}

func (v *NullableBool) Set(val *bool) {
	v.value = val
	v.isSet = true
}

func (v NullableBool) IsSet() bool {
	return v.isSet
}

func (v *NullableBool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBool(val *bool) *NullableBool {
	return &NullableBool{value: val, isSet: true}
}

func (v NullableBool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type NullableInt struct {
	value *int
	isSet bool
}

func (v NullableInt) Get() *int {
	return v.value
}

func (v *NullableInt) Set(val *int) {
	v.value = val
	v.isSet = true
}

func (v NullableInt) IsSet() bool {
	return v.isSet
}

func (v *NullableInt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInt(val *int) *NullableInt {
	return &NullableInt{value: val, isSet: true}
}

func (v NullableInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type NullableInt32 struct {
	value *int32
	isSet bool
}

func (v NullableInt32) Get() *int32 {
	return v.value
}

func (v *NullableInt32) Set(val *int32) {
	v.value = val
	v.isSet = true
}

func (v NullableInt32) IsSet() bool {
	return v.isSet
}

func (v *NullableInt32) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInt32(val *int32) *NullableInt32 {
	return &NullableInt32{value: val, isSet: true}
}

func (v NullableInt32) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInt32) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type NullableInt64 struct {
	value *int64
	isSet bool
}

func (v NullableInt64) Get() *int64 {
	return v.value
}

func (v *NullableInt64) Set(val *int64) {
	v.value = val
	v.isSet = true
}

func (v NullableInt64) IsSet() bool {
	return v.isSet
}

func (v *NullableInt64) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInt64(val *int64) *NullableInt64 {
	return &NullableInt64{value: val, isSet: true}
}

func (v NullableInt64) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInt64) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type NullableFloat32 struct {
	value *float32
	isSet bool
}

func (v NullableFloat32) Get() *float32 {
	return v.value
}

func (v *NullableFloat32) Set(val *float32) {
	v.value = val
	v.isSet = true
}

func (v NullableFloat32) IsSet() bool {
	return v.isSet
}

func (v *NullableFloat32) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFloat32(val *float32) *NullableFloat32 {
	return &NullableFloat32{value: val, isSet: true}
}

func (v NullableFloat32) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFloat32) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type NullableFloat64 struct {
	value *float64
	isSet bool
}

func (v NullableFloat64) Get() *float64 {
	return v.value
}

func (v *NullableFloat64) Set(val *float64) {
	v.value = val
	v.isSet = true
}

func (v NullableFloat64) IsSet() bool {
	return v.isSet
}

func (v *NullableFloat64) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFloat64(val *float64) *NullableFloat64 {
	return &NullableFloat64{value: val, isSet: true}
}

func (v NullableFloat64) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFloat64) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type NullableString struct {
	value *string
	isSet bool
}

func (v NullableString) Get() *string {
	return v.value
}

func (v *NullableString) Set(val *string) {
	v.value = val
	v.isSet = true
}

func (v NullableString) IsSet() bool {
	return v.isSet
}

func (v *NullableString) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableString(val *string) *NullableString {
	return &NullableString{value: val, isSet: true}
}

func (v NullableString) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableString) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type NullableTime struct {
	value *time.Time
	isSet bool
}

func (v NullableTime) Get() *time.Time {
	return v.value
}

func (v *NullableTime) Set(val *time.Time) {
	v.value = val
	v.isSet = true
}

func (v NullableTime) IsSet() bool {
	return v.isSet
}

func (v *NullableTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTime(val *time.Time) *NullableTime {
	return &NullableTime{value: val, isSet: true}
}

func (v NullableTime) MarshalJSON() ([]byte, error) {
	return v.value.MarshalJSON()
}

func (v *NullableTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

// IsNil checks if an input is nil
func IsNil(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	case reflect.Array:
		return reflect.ValueOf(i).IsZero()
	}
	return false
}

type MappedNullable interface {
	ToMap() (map[string]interface{}, error)
}
