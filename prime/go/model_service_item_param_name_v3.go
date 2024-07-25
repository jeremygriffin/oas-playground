/*
MilMove Prime V3 API

The Prime V3 API is a RESTful API that enables the Prime contractor to request information about upcoming moves, update the details and status of those moves, and make payment requests. It uses Mutual TLS for authentication procedures.  All endpoints are located at `/prime/v3/`. 

API version: 0.0.1
Contact: milmove-developers@caci.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// ServiceItemParamNameV3 the model 'ServiceItemParamNameV3'
type ServiceItemParamNameV3 string

// List of ServiceItemParamName
const (
	ACTUAL_PICKUP_DATE ServiceItemParamNameV3 = "ActualPickupDate"
	CONTRACT_CODE ServiceItemParamNameV3 = "ContractCode"
	CONTRACT_YEAR_NAME ServiceItemParamNameV3 = "ContractYearName"
	CUBIC_FEET_BILLED ServiceItemParamNameV3 = "CubicFeetBilled"
	CUBIC_FEET_CRATING ServiceItemParamNameV3 = "CubicFeetCrating"
	DIMENSION_HEIGHT ServiceItemParamNameV3 = "DimensionHeight"
	DIMENSION_LENGTH ServiceItemParamNameV3 = "DimensionLength"
	DIMENSION_WIDTH ServiceItemParamNameV3 = "DimensionWidth"
	DISTANCE_ZIP ServiceItemParamNameV3 = "DistanceZip"
	DISTANCE_ZIP_SIT_DEST ServiceItemParamNameV3 = "DistanceZipSITDest"
	DISTANCE_ZIP_SIT_ORIGIN ServiceItemParamNameV3 = "DistanceZipSITOrigin"
	EIA_FUEL_PRICE ServiceItemParamNameV3 = "EIAFuelPrice"
	ESCALATION_COMPOUNDED ServiceItemParamNameV3 = "EscalationCompounded"
	FSC_MULTIPLIER ServiceItemParamNameV3 = "FSCMultiplier"
	FSC_PRICE_DIFFERENCE_IN_CENTS ServiceItemParamNameV3 = "FSCPriceDifferenceInCents"
	FSC_WEIGHT_BASED_DISTANCE_MULTIPLIER ServiceItemParamNameV3 = "FSCWeightBasedDistanceMultiplier"
	IS_PEAK ServiceItemParamNameV3 = "IsPeak"
	MARKET_DEST ServiceItemParamNameV3 = "MarketDest"
	MARKET_ORIGIN ServiceItemParamNameV3 = "MarketOrigin"
	MTO_AVAILABLE_TO_PRIME_AT ServiceItemParamNameV3 = "MTOAvailableToPrimeAt"
	NTS_PACKING_FACTOR ServiceItemParamNameV3 = "NTSPackingFactor"
	NUMBER_DAYS_SIT ServiceItemParamNameV3 = "NumberDaysSIT"
	PRICE_AREA_DEST ServiceItemParamNameV3 = "PriceAreaDest"
	PRICE_AREA_INTL_DEST ServiceItemParamNameV3 = "PriceAreaIntlDest"
	PRICE_AREA_INTL_ORIGIN ServiceItemParamNameV3 = "PriceAreaIntlOrigin"
	PRICE_AREA_ORIGIN ServiceItemParamNameV3 = "PriceAreaOrigin"
	PRICE_RATE_OR_FACTOR ServiceItemParamNameV3 = "PriceRateOrFactor"
	PSI_LINEHAUL_DOM ServiceItemParamNameV3 = "PSI_LinehaulDom"
	PSI_LINEHAUL_DOM_PRICE ServiceItemParamNameV3 = "PSI_LinehaulDomPrice"
	PSI_LINEHAUL_SHORT ServiceItemParamNameV3 = "PSI_LinehaulShort"
	PSI_LINEHAUL_SHORT_PRICE ServiceItemParamNameV3 = "PSI_LinehaulShortPrice"
	PSI_PRICE_DOM_DEST ServiceItemParamNameV3 = "PSI_PriceDomDest"
	PSI_PRICE_DOM_DEST_PRICE ServiceItemParamNameV3 = "PSI_PriceDomDestPrice"
	PSI_PRICE_DOM_ORIGIN ServiceItemParamNameV3 = "PSI_PriceDomOrigin"
	PSI_PRICE_DOM_ORIGIN_PRICE ServiceItemParamNameV3 = "PSI_PriceDomOriginPrice"
	PSI_SHIPPING_LINEHAUL_INTL_CO ServiceItemParamNameV3 = "PSI_ShippingLinehaulIntlCO"
	PSI_SHIPPING_LINEHAUL_INTL_CO_PRICE ServiceItemParamNameV3 = "PSI_ShippingLinehaulIntlCOPrice"
	PSI_SHIPPING_LINEHAUL_INTL_OC ServiceItemParamNameV3 = "PSI_ShippingLinehaulIntlOC"
	PSI_SHIPPING_LINEHAUL_INTL_OC_PRICE ServiceItemParamNameV3 = "PSI_ShippingLinehaulIntlOCPrice"
	PSI_SHIPPING_LINEHAUL_INTL_OO ServiceItemParamNameV3 = "PSI_ShippingLinehaulIntlOO"
	PSI_SHIPPING_LINEHAUL_INTL_OO_PRICE ServiceItemParamNameV3 = "PSI_ShippingLinehaulIntlOOPrice"
	RATE_AREA_NON_STD_DEST ServiceItemParamNameV3 = "RateAreaNonStdDest"
	RATE_AREA_NON_STD_ORIGIN ServiceItemParamNameV3 = "RateAreaNonStdOrigin"
	REFERENCE_DATE ServiceItemParamNameV3 = "ReferenceDate"
	REQUESTED_PICKUP_DATE ServiceItemParamNameV3 = "RequestedPickupDate"
	SERVICE_AREA_DEST ServiceItemParamNameV3 = "ServiceAreaDest"
	SERVICE_AREA_ORIGIN ServiceItemParamNameV3 = "ServiceAreaOrigin"
	SERVICES_SCHEDULE_DEST ServiceItemParamNameV3 = "ServicesScheduleDest"
	SERVICES_SCHEDULE_ORIGIN ServiceItemParamNameV3 = "ServicesScheduleOrigin"
	SIT_PAYMENT_REQUEST_END ServiceItemParamNameV3 = "SITPaymentRequestEnd"
	SIT_PAYMENT_REQUEST_START ServiceItemParamNameV3 = "SITPaymentRequestStart"
	SIT_SCHEDULE_DEST ServiceItemParamNameV3 = "SITScheduleDest"
	SIT_SCHEDULE_ORIGIN ServiceItemParamNameV3 = "SITScheduleOrigin"
	SIT_SERVICE_AREA_DEST ServiceItemParamNameV3 = "SITServiceAreaDest"
	SIT_SERVICE_AREA_ORIGIN ServiceItemParamNameV3 = "SITServiceAreaOrigin"
	WEIGHT_ADJUSTED ServiceItemParamNameV3 = "WeightAdjusted"
	WEIGHT_BILLED ServiceItemParamNameV3 = "WeightBilled"
	WEIGHT_ESTIMATED ServiceItemParamNameV3 = "WeightEstimated"
	WEIGHT_ORIGINAL ServiceItemParamNameV3 = "WeightOriginal"
	WEIGHT_REWEIGH ServiceItemParamNameV3 = "WeightReweigh"
	ZIP_DEST_ADDRESS ServiceItemParamNameV3 = "ZipDestAddress"
	ZIP_PICKUP_ADDRESS ServiceItemParamNameV3 = "ZipPickupAddress"
	ZIP_SIT_DEST_HHG_FINAL_ADDRESS ServiceItemParamNameV3 = "ZipSITDestHHGFinalAddress"
	ZIP_SIT_DEST_HHG_ORIGINAL_ADDRESS ServiceItemParamNameV3 = "ZipSITDestHHGOriginalAddress"
	ZIP_SIT_ORIGIN_HHG_ACTUAL_ADDRESS ServiceItemParamNameV3 = "ZipSITOriginHHGActualAddress"
	ZIP_SIT_ORIGIN_HHG_ORIGINAL_ADDRESS ServiceItemParamNameV3 = "ZipSITOriginHHGOriginalAddress"
	STANDALONE_CRATE ServiceItemParamNameV3 = "StandaloneCrate"
	STANDALONE_CRATE_CAP ServiceItemParamNameV3 = "StandaloneCrateCap"
	UNCAPPED_REQUEST_TOTAL ServiceItemParamNameV3 = "UncappedRequestTotal"
)

// All allowed values of ServiceItemParamNameV3 enum
var AllowedServiceItemParamNameV3EnumValues = []ServiceItemParamNameV3{
	"ActualPickupDate",
	"ContractCode",
	"ContractYearName",
	"CubicFeetBilled",
	"CubicFeetCrating",
	"DimensionHeight",
	"DimensionLength",
	"DimensionWidth",
	"DistanceZip",
	"DistanceZipSITDest",
	"DistanceZipSITOrigin",
	"EIAFuelPrice",
	"EscalationCompounded",
	"FSCMultiplier",
	"FSCPriceDifferenceInCents",
	"FSCWeightBasedDistanceMultiplier",
	"IsPeak",
	"MarketDest",
	"MarketOrigin",
	"MTOAvailableToPrimeAt",
	"NTSPackingFactor",
	"NumberDaysSIT",
	"PriceAreaDest",
	"PriceAreaIntlDest",
	"PriceAreaIntlOrigin",
	"PriceAreaOrigin",
	"PriceRateOrFactor",
	"PSI_LinehaulDom",
	"PSI_LinehaulDomPrice",
	"PSI_LinehaulShort",
	"PSI_LinehaulShortPrice",
	"PSI_PriceDomDest",
	"PSI_PriceDomDestPrice",
	"PSI_PriceDomOrigin",
	"PSI_PriceDomOriginPrice",
	"PSI_ShippingLinehaulIntlCO",
	"PSI_ShippingLinehaulIntlCOPrice",
	"PSI_ShippingLinehaulIntlOC",
	"PSI_ShippingLinehaulIntlOCPrice",
	"PSI_ShippingLinehaulIntlOO",
	"PSI_ShippingLinehaulIntlOOPrice",
	"RateAreaNonStdDest",
	"RateAreaNonStdOrigin",
	"ReferenceDate",
	"RequestedPickupDate",
	"ServiceAreaDest",
	"ServiceAreaOrigin",
	"ServicesScheduleDest",
	"ServicesScheduleOrigin",
	"SITPaymentRequestEnd",
	"SITPaymentRequestStart",
	"SITScheduleDest",
	"SITScheduleOrigin",
	"SITServiceAreaDest",
	"SITServiceAreaOrigin",
	"WeightAdjusted",
	"WeightBilled",
	"WeightEstimated",
	"WeightOriginal",
	"WeightReweigh",
	"ZipDestAddress",
	"ZipPickupAddress",
	"ZipSITDestHHGFinalAddress",
	"ZipSITDestHHGOriginalAddress",
	"ZipSITOriginHHGActualAddress",
	"ZipSITOriginHHGOriginalAddress",
	"StandaloneCrate",
	"StandaloneCrateCap",
	"UncappedRequestTotal",
}

func (v *ServiceItemParamNameV3) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ServiceItemParamNameV3(value)
	for _, existing := range AllowedServiceItemParamNameV3EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ServiceItemParamNameV3", value)
}

// NewServiceItemParamNameV3FromValue returns a pointer to a valid ServiceItemParamNameV3
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewServiceItemParamNameV3FromValue(v string) (*ServiceItemParamNameV3, error) {
	ev := ServiceItemParamNameV3(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ServiceItemParamNameV3: valid values are %v", v, AllowedServiceItemParamNameV3EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ServiceItemParamNameV3) IsValid() bool {
	for _, existing := range AllowedServiceItemParamNameV3EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceItemParamName value
func (v ServiceItemParamNameV3) Ptr() *ServiceItemParamNameV3 {
	return &v
}

type NullableServiceItemParamNameV3 struct {
	value *ServiceItemParamNameV3
	isSet bool
}

func (v NullableServiceItemParamNameV3) Get() *ServiceItemParamNameV3 {
	return v.value
}

func (v *NullableServiceItemParamNameV3) Set(val *ServiceItemParamNameV3) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceItemParamNameV3) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceItemParamNameV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceItemParamNameV3(val *ServiceItemParamNameV3) *NullableServiceItemParamNameV3 {
	return &NullableServiceItemParamNameV3{value: val, isSet: true}
}

func (v NullableServiceItemParamNameV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceItemParamNameV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

