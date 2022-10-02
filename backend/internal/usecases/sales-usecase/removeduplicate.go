package sales_usecase

import "github.com/beto-ouverney/go-affiliates/backend/internal/entities"

// typeSet for Generics in removeDuplicates
type typeSet interface {
	entities.Producer | entities.Product | entities.Affiliate
}

// removeDuplicate removes duplicate strings or struct from a slice
func removeDuplicate[T typeSet](sliceList []T) []T {
	allKeys := make(map[T]bool)
	var list []T
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
