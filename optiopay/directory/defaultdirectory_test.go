package directory

import (
	"testing"

	"github.com/n340r/go-notes/optiopay/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// TestSuite is our test suite
type TestSuite struct {
	suite.Suite
}

// createEmployees creates employee for every letter in alphabet
func createEmployees() map[rune]*model.Employee {
	org := map[rune]*model.Employee{}
	for _, r := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		org[r] = &model.Employee{
			Name:      string(r),
			Employees: []*model.Employee{},
		}
	}
	return org
}

// Test1 represents the following case:
//
// topManager = Node A
// Employee 1 = Node E
// Employee 2 = Node I
//                A
//             /     \
//            B       C
//          /   \   /   \
//         D     E F     G
//       /   \
//      H     I
//
// Lowest Common Manager for H and E must be B
// Lowest Common Manager for H and A must be A
// Lowest Common Manager for G and I must be A

func (suite *TestSuite) Test1() {
	employees := createEmployees()

	employees['A'].AddEmployees(employees['B'], employees['C'])
	employees['B'].AddEmployees(employees['D'], employees['E'])
	employees['C'].AddEmployees(employees['F'], employees['G'])
	employees['D'].AddEmployees(employees['H'], employees['I'])

	directory := NewDefaultDirectory()

	lcm := directory.GetLowestCommonManager(employees['A'], employees['E'], employees['H'])
	assert.Equal(suite.T(), lcm, employees['B'])

	lcm = directory.GetLowestCommonManager(employees['A'], employees['G'], employees['I'])
	assert.Equal(suite.T(), lcm, employees['A'])

	lcm = directory.GetLowestCommonManager(employees['A'], employees['A'], employees['H'])
	assert.Equal(suite.T(), lcm, employees['A'])
}

// TestLCM runs our test suite
func TestLCM(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
