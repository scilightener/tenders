package tender

import "strings"

type serviceType string

const (
	ServiceTypeConstruction serviceType = "CONSTRUCTION"
	ServiceTypeDelivery     serviceType = "DELIVERY"
	ServiceTypeManufacture  serviceType = "MANUFACTURE"
)

func ServiceTypeFromString(s string) (serviceType, error) {
	up := strings.ToUpper(s)
	switch up {
	case string(ServiceTypeConstruction):
		return ServiceTypeConstruction, nil
	case string(ServiceTypeDelivery):
		return ServiceTypeDelivery, nil
	case string(ServiceTypeManufacture):
		return ServiceTypeManufacture, nil
	default:
		return "", ErrInvalidServiceType(s)
	}
}
