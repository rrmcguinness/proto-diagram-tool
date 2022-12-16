package proto

import (
	"fmt"
)

const (
	PatternRelationshipInheritance = "%s --|> %s" // inheritance
	PatternRelationshipComposition = "%s --* %s"  // Strong life-cycle dependency (Parent / Child Relationship)
	PatternRelationshipAggregation = "%s --o %s"  // Weak life-cycle dependency (Associated relationship, but independent)
	PatternRelationshipAssociation = "%s --> %s"  // A is associated to B
	PatternRelationshipLinkSolid   = "%s -- %s"   // A is linked to B
	PatternRelationshipLinkDashed  = "%s .. %s"   // A is linked loosely to B
	PatternRelationshipDependency  = "%s ..> %s"  // A depends on B
	PatternRelationshipRealization = "%s ..|> %s" // A realizes B

	PatternType           = "%s%s %s"          // Visibility, Type, Name
	PatternFunctionVoid   = "%s%s(%s)"         //Visibility, Function Name, Arguments
	PatternFunctionReturn = "%s %s"            // Visibility, Function name, Arguments, Return Type
	PatternList           = "%sList~%s~ %s"    // Visibility, Type, Name
	PatternMap            = "%sMap~%s, %s~ %s" // Visibility, KeyType, ValueType, Name

	Colon  = ":"
	Period = "."
	Space  = " "

	VisibilityEmpty     = ""
	VisibilityPublic    = "+"
	VisibilityPrivate   = "-"
	VisibilityProtected = "#"
	VisibilityPackage   = "~"

	AbstractSuffix = "*"
	StaticSuffix   = "$"

	DiagramClass = "classDiagram"
)

func FormatType(visibility string, typeName string, name string) string {
	return fmt.Sprintf(PatternType, visibility, typeName, name)
}

func FormatFunctionVoid(vis string, fName string, arguments map[string]string) string {
	var argString = ""
	for k, v := range arguments {
		argString += FormatType(VisibilityEmpty, v, k)
	}
	return fmt.Sprintf(PatternFunctionVoid, vis, fName, argString)
}

func FormatFunctionReturn(vis string, fName string, arguments map[string]string, rType string) string {
	return fmt.Sprintf(PatternFunctionReturn, FormatFunctionVoid(vis, fName, arguments), rType)
}

func FormatGenericList(vis string, iType string, name string) string {
	return fmt.Sprintf(PatternList, vis, iType, name)
}

func FormatGenericMap(vis string, kType string, vType string, name string) string {
	return fmt.Sprintf(PatternMap, vis, kType, vType, name)
}
