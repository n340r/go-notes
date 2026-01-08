package directory

import "github.com/n340r/go-notes/optiopay/model"

type Directory interface {
	GetLowestCommonManager(top, employee1, employee2 *model.Employee) *model.Employee
}
