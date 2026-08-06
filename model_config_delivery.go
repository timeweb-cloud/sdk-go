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

// checks if the ConfigDelivery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigDelivery{}

// ConfigDelivery Настройки раздачи контента
type ConfigDelivery struct {
	// Раздавать контент по HTTP/3 (QUIC).
	Http3 *bool `json:"http3,omitempty"`
	// Сжимать ответы алгоритмом gzip.
	Gzip *bool `json:"gzip,omitempty"`
	// Режим раздачи больших файлов.
	LargeFiles *bool `json:"large_files,omitempty"`
	// Размер слайса в мегабайтах — файл забирается с источника и кэшируется частями такого размера. `null` — не использовать слайсинг.
	SliceSize NullableInt32 `json:"slice_size,omitempty"`
	// Оптимизировать изображения на лету.
	ImageOptimization *bool `json:"image_optimization,omitempty"`
	Packaging NullableConfigDeliveryPackaging `json:"packaging,omitempty"`
}

// NewConfigDelivery instantiates a new ConfigDelivery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigDelivery() *ConfigDelivery {
	this := ConfigDelivery{}
	return &this
}

// NewConfigDeliveryWithDefaults instantiates a new ConfigDelivery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigDeliveryWithDefaults() *ConfigDelivery {
	this := ConfigDelivery{}
	return &this
}

// GetHttp3 returns the Http3 field value if set, zero value otherwise.
func (o *ConfigDelivery) GetHttp3() bool {
	if o == nil || IsNil(o.Http3) {
		var ret bool
		return ret
	}
	return *o.Http3
}

// GetHttp3Ok returns a tuple with the Http3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigDelivery) GetHttp3Ok() (*bool, bool) {
	if o == nil || IsNil(o.Http3) {
		return nil, false
	}
	return o.Http3, true
}

// HasHttp3 returns a boolean if a field has been set.
func (o *ConfigDelivery) HasHttp3() bool {
	if o != nil && !IsNil(o.Http3) {
		return true
	}

	return false
}

// SetHttp3 gets a reference to the given bool and assigns it to the Http3 field.
func (o *ConfigDelivery) SetHttp3(v bool) {
	o.Http3 = &v
}

// GetGzip returns the Gzip field value if set, zero value otherwise.
func (o *ConfigDelivery) GetGzip() bool {
	if o == nil || IsNil(o.Gzip) {
		var ret bool
		return ret
	}
	return *o.Gzip
}

// GetGzipOk returns a tuple with the Gzip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigDelivery) GetGzipOk() (*bool, bool) {
	if o == nil || IsNil(o.Gzip) {
		return nil, false
	}
	return o.Gzip, true
}

// HasGzip returns a boolean if a field has been set.
func (o *ConfigDelivery) HasGzip() bool {
	if o != nil && !IsNil(o.Gzip) {
		return true
	}

	return false
}

// SetGzip gets a reference to the given bool and assigns it to the Gzip field.
func (o *ConfigDelivery) SetGzip(v bool) {
	o.Gzip = &v
}

// GetLargeFiles returns the LargeFiles field value if set, zero value otherwise.
func (o *ConfigDelivery) GetLargeFiles() bool {
	if o == nil || IsNil(o.LargeFiles) {
		var ret bool
		return ret
	}
	return *o.LargeFiles
}

// GetLargeFilesOk returns a tuple with the LargeFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigDelivery) GetLargeFilesOk() (*bool, bool) {
	if o == nil || IsNil(o.LargeFiles) {
		return nil, false
	}
	return o.LargeFiles, true
}

// HasLargeFiles returns a boolean if a field has been set.
func (o *ConfigDelivery) HasLargeFiles() bool {
	if o != nil && !IsNil(o.LargeFiles) {
		return true
	}

	return false
}

// SetLargeFiles gets a reference to the given bool and assigns it to the LargeFiles field.
func (o *ConfigDelivery) SetLargeFiles(v bool) {
	o.LargeFiles = &v
}

// GetSliceSize returns the SliceSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigDelivery) GetSliceSize() int32 {
	if o == nil || IsNil(o.SliceSize.Get()) {
		var ret int32
		return ret
	}
	return *o.SliceSize.Get()
}

// GetSliceSizeOk returns a tuple with the SliceSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigDelivery) GetSliceSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SliceSize.Get(), o.SliceSize.IsSet()
}

// HasSliceSize returns a boolean if a field has been set.
func (o *ConfigDelivery) HasSliceSize() bool {
	if o != nil && o.SliceSize.IsSet() {
		return true
	}

	return false
}

// SetSliceSize gets a reference to the given NullableInt32 and assigns it to the SliceSize field.
func (o *ConfigDelivery) SetSliceSize(v int32) {
	o.SliceSize.Set(&v)
}
// SetSliceSizeNil sets the value for SliceSize to be an explicit nil
func (o *ConfigDelivery) SetSliceSizeNil() {
	o.SliceSize.Set(nil)
}

// UnsetSliceSize ensures that no value is present for SliceSize, not even an explicit nil
func (o *ConfigDelivery) UnsetSliceSize() {
	o.SliceSize.Unset()
}

// GetImageOptimization returns the ImageOptimization field value if set, zero value otherwise.
func (o *ConfigDelivery) GetImageOptimization() bool {
	if o == nil || IsNil(o.ImageOptimization) {
		var ret bool
		return ret
	}
	return *o.ImageOptimization
}

// GetImageOptimizationOk returns a tuple with the ImageOptimization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigDelivery) GetImageOptimizationOk() (*bool, bool) {
	if o == nil || IsNil(o.ImageOptimization) {
		return nil, false
	}
	return o.ImageOptimization, true
}

// HasImageOptimization returns a boolean if a field has been set.
func (o *ConfigDelivery) HasImageOptimization() bool {
	if o != nil && !IsNil(o.ImageOptimization) {
		return true
	}

	return false
}

// SetImageOptimization gets a reference to the given bool and assigns it to the ImageOptimization field.
func (o *ConfigDelivery) SetImageOptimization(v bool) {
	o.ImageOptimization = &v
}

// GetPackaging returns the Packaging field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigDelivery) GetPackaging() ConfigDeliveryPackaging {
	if o == nil || IsNil(o.Packaging.Get()) {
		var ret ConfigDeliveryPackaging
		return ret
	}
	return *o.Packaging.Get()
}

// GetPackagingOk returns a tuple with the Packaging field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigDelivery) GetPackagingOk() (*ConfigDeliveryPackaging, bool) {
	if o == nil {
		return nil, false
	}
	return o.Packaging.Get(), o.Packaging.IsSet()
}

// HasPackaging returns a boolean if a field has been set.
func (o *ConfigDelivery) HasPackaging() bool {
	if o != nil && o.Packaging.IsSet() {
		return true
	}

	return false
}

// SetPackaging gets a reference to the given NullableConfigDeliveryPackaging and assigns it to the Packaging field.
func (o *ConfigDelivery) SetPackaging(v ConfigDeliveryPackaging) {
	o.Packaging.Set(&v)
}
// SetPackagingNil sets the value for Packaging to be an explicit nil
func (o *ConfigDelivery) SetPackagingNil() {
	o.Packaging.Set(nil)
}

// UnsetPackaging ensures that no value is present for Packaging, not even an explicit nil
func (o *ConfigDelivery) UnsetPackaging() {
	o.Packaging.Unset()
}

func (o ConfigDelivery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigDelivery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Http3) {
		toSerialize["http3"] = o.Http3
	}
	if !IsNil(o.Gzip) {
		toSerialize["gzip"] = o.Gzip
	}
	if !IsNil(o.LargeFiles) {
		toSerialize["large_files"] = o.LargeFiles
	}
	if o.SliceSize.IsSet() {
		toSerialize["slice_size"] = o.SliceSize.Get()
	}
	if !IsNil(o.ImageOptimization) {
		toSerialize["image_optimization"] = o.ImageOptimization
	}
	if o.Packaging.IsSet() {
		toSerialize["packaging"] = o.Packaging.Get()
	}
	return toSerialize, nil
}

type NullableConfigDelivery struct {
	value *ConfigDelivery
	isSet bool
}

func (v NullableConfigDelivery) Get() *ConfigDelivery {
	return v.value
}

func (v *NullableConfigDelivery) Set(val *ConfigDelivery) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigDelivery) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigDelivery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigDelivery(val *ConfigDelivery) *NullableConfigDelivery {
	return &NullableConfigDelivery{value: val, isSet: true}
}

func (v NullableConfigDelivery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigDelivery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


