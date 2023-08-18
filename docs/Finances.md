# Finances

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balance** | **float32** | Баланс аккаунта. | 
**Currency** | **string** | Валюта, которая используется на аккаунте. | 
**DiscountEndDateAt** | **NullableString** | Значение времени, указанное в комбинированном формате даты и времени ISO8601, которое представляет, когда заканчивается скидка для аккаунта. | 
**DiscountPercent** | **float32** | Процент скидки для аккаунта. | 
**HourlyCost** | **float32** | Стоимость услуг на аккаунте в час. | 
**HourlyFee** | **float32** | Абонентская плата в час (с учетом скидок). | 
**MonthlyCost** | **float32** | Стоимость услуг на аккаунте в месяц. | 
**MonthlyFee** | **float32** | Абонентская плата в месяц (с учетом скидок). | 
**TotalPaid** | **float32** | Общая сумма платежей на аккаунте. | 
**HoursLeft** | **NullableFloat32** | Сколько часов работы услуг оплачено на аккаунте. | 
**AutopayCardInfo** | **NullableString** | Информация о карте, используемой для автоплатежей. | 

## Methods

### NewFinances

`func NewFinances(balance float32, currency string, discountEndDateAt NullableString, discountPercent float32, hourlyCost float32, hourlyFee float32, monthlyCost float32, monthlyFee float32, totalPaid float32, hoursLeft NullableFloat32, autopayCardInfo NullableString, ) *Finances`

NewFinances instantiates a new Finances object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinancesWithDefaults

`func NewFinancesWithDefaults() *Finances`

NewFinancesWithDefaults instantiates a new Finances object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalance

`func (o *Finances) GetBalance() float32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *Finances) GetBalanceOk() (*float32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *Finances) SetBalance(v float32)`

SetBalance sets Balance field to given value.


### GetCurrency

`func (o *Finances) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Finances) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Finances) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDiscountEndDateAt

`func (o *Finances) GetDiscountEndDateAt() string`

GetDiscountEndDateAt returns the DiscountEndDateAt field if non-nil, zero value otherwise.

### GetDiscountEndDateAtOk

`func (o *Finances) GetDiscountEndDateAtOk() (*string, bool)`

GetDiscountEndDateAtOk returns a tuple with the DiscountEndDateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountEndDateAt

`func (o *Finances) SetDiscountEndDateAt(v string)`

SetDiscountEndDateAt sets DiscountEndDateAt field to given value.


### SetDiscountEndDateAtNil

`func (o *Finances) SetDiscountEndDateAtNil(b bool)`

 SetDiscountEndDateAtNil sets the value for DiscountEndDateAt to be an explicit nil

### UnsetDiscountEndDateAt
`func (o *Finances) UnsetDiscountEndDateAt()`

UnsetDiscountEndDateAt ensures that no value is present for DiscountEndDateAt, not even an explicit nil
### GetDiscountPercent

`func (o *Finances) GetDiscountPercent() float32`

GetDiscountPercent returns the DiscountPercent field if non-nil, zero value otherwise.

### GetDiscountPercentOk

`func (o *Finances) GetDiscountPercentOk() (*float32, bool)`

GetDiscountPercentOk returns a tuple with the DiscountPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPercent

`func (o *Finances) SetDiscountPercent(v float32)`

SetDiscountPercent sets DiscountPercent field to given value.


### GetHourlyCost

`func (o *Finances) GetHourlyCost() float32`

GetHourlyCost returns the HourlyCost field if non-nil, zero value otherwise.

### GetHourlyCostOk

`func (o *Finances) GetHourlyCostOk() (*float32, bool)`

GetHourlyCostOk returns a tuple with the HourlyCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyCost

`func (o *Finances) SetHourlyCost(v float32)`

SetHourlyCost sets HourlyCost field to given value.


### GetHourlyFee

`func (o *Finances) GetHourlyFee() float32`

GetHourlyFee returns the HourlyFee field if non-nil, zero value otherwise.

### GetHourlyFeeOk

`func (o *Finances) GetHourlyFeeOk() (*float32, bool)`

GetHourlyFeeOk returns a tuple with the HourlyFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyFee

`func (o *Finances) SetHourlyFee(v float32)`

SetHourlyFee sets HourlyFee field to given value.


### GetMonthlyCost

`func (o *Finances) GetMonthlyCost() float32`

GetMonthlyCost returns the MonthlyCost field if non-nil, zero value otherwise.

### GetMonthlyCostOk

`func (o *Finances) GetMonthlyCostOk() (*float32, bool)`

GetMonthlyCostOk returns a tuple with the MonthlyCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyCost

`func (o *Finances) SetMonthlyCost(v float32)`

SetMonthlyCost sets MonthlyCost field to given value.


### GetMonthlyFee

`func (o *Finances) GetMonthlyFee() float32`

GetMonthlyFee returns the MonthlyFee field if non-nil, zero value otherwise.

### GetMonthlyFeeOk

`func (o *Finances) GetMonthlyFeeOk() (*float32, bool)`

GetMonthlyFeeOk returns a tuple with the MonthlyFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyFee

`func (o *Finances) SetMonthlyFee(v float32)`

SetMonthlyFee sets MonthlyFee field to given value.


### GetTotalPaid

`func (o *Finances) GetTotalPaid() float32`

GetTotalPaid returns the TotalPaid field if non-nil, zero value otherwise.

### GetTotalPaidOk

`func (o *Finances) GetTotalPaidOk() (*float32, bool)`

GetTotalPaidOk returns a tuple with the TotalPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPaid

`func (o *Finances) SetTotalPaid(v float32)`

SetTotalPaid sets TotalPaid field to given value.


### GetHoursLeft

`func (o *Finances) GetHoursLeft() float32`

GetHoursLeft returns the HoursLeft field if non-nil, zero value otherwise.

### GetHoursLeftOk

`func (o *Finances) GetHoursLeftOk() (*float32, bool)`

GetHoursLeftOk returns a tuple with the HoursLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoursLeft

`func (o *Finances) SetHoursLeft(v float32)`

SetHoursLeft sets HoursLeft field to given value.


### SetHoursLeftNil

`func (o *Finances) SetHoursLeftNil(b bool)`

 SetHoursLeftNil sets the value for HoursLeft to be an explicit nil

### UnsetHoursLeft
`func (o *Finances) UnsetHoursLeft()`

UnsetHoursLeft ensures that no value is present for HoursLeft, not even an explicit nil
### GetAutopayCardInfo

`func (o *Finances) GetAutopayCardInfo() string`

GetAutopayCardInfo returns the AutopayCardInfo field if non-nil, zero value otherwise.

### GetAutopayCardInfoOk

`func (o *Finances) GetAutopayCardInfoOk() (*string, bool)`

GetAutopayCardInfoOk returns a tuple with the AutopayCardInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutopayCardInfo

`func (o *Finances) SetAutopayCardInfo(v string)`

SetAutopayCardInfo sets AutopayCardInfo field to given value.


### SetAutopayCardInfoNil

`func (o *Finances) SetAutopayCardInfoNil(b bool)`

 SetAutopayCardInfoNil sets the value for AutopayCardInfo to be an explicit nil

### UnsetAutopayCardInfo
`func (o *Finances) UnsetAutopayCardInfo()`

UnsetAutopayCardInfo ensures that no value is present for AutopayCardInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


