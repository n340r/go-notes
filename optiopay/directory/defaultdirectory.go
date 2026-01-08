package directory

import "github.com/n340r/go-notes/optiopay/model"

// ??? can we replace struct{} with any ?
type DefaultDirectory struct{}

// ??? i do not understand this syntax and the purpose of it
var _ Directory = (*DefaultDirectory)(nil)

func NewDefaultDirectory() DefaultDirectory {
	return DefaultDirectory{}
}

func (d DefaultDirectory) GetLowestCommonManager(top, employee1, employee2 *model.Employee) *model.Employee {
	return getLowestCommonManager(top, employee1, employee2).lowestCommonManager
}

type lcm struct {
	lowestCommonManager *model.Employee
	foundEmployees      int
}

func getLowestCommonManager(manager, employee1, employee2 *model.Employee) lcm {
	//count found employees
	foundEmployees := 0 // ??? do i get it right that this thing is saved/embedded/closured in a recursive descend context ? am i thinking right ? is this the right way of putting it ?
	var lcmEmployee *model.Employee

	// decend
	for _, child := range manager.Employees {
		childLCM := getLowestCommonManager(child, employee1, employee2)
		// each child should tell us { howNameEmployeesFound, isLCMknown}
		// and if one child found LCM, bubble it up immediately
		if childLCM.lowestCommonManager != nil {
			return childLCM
		}

		//otherwise accumulate found employees
		foundEmployees += childLCM.foundEmployees
	}

	// If I myself am one of the targets, count it
	if manager == employee1 || manager == employee2 {
		// if yes, then we increment
		foundEmployees++
	}

	// a node sees that it is a LCM
	if foundEmployees == 2 {
		lcmEmployee = manager
	}

	return lcm{
		lowestCommonManager: lcmEmployee, // ??? what do i put there since we are on a leaf and do not know what manager brought us to that leaf ?
		foundEmployees:      foundEmployees,
	}

}
