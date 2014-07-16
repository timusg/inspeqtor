
/*
*/
package parser

const numNTSymbols = 10
type(
	gotoTable [numStates]gotoRow
	gotoRow	[numNTSymbols] int
)

var gotoTab = gotoTable{
	gotoRow{ // S0
		
		-1, // S'
		1, // Check
		2, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // Inittype
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S1
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // Inittype
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S2
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		5, // Checktype
		-1, // Action
		6, // Inittype
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S3
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // Inittype
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S4
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // Inittype
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S5
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // Inittype
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S6
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		16, // Checktype
		-1, // Action
		-1, // Inittype
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S7
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // Inittype
		-1, // IntAmount
		-1, // HumanAmount
		18, // Rule
		17, // RuleList
		

	},
	gotoRow{ // S8
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // Inittype
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S9
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // Inittype
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S10
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // Inittype
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S11
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // Inittype
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S12
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // Inittype
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S13
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // Inittype
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S14
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // Inittype
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S15
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // Inittype
		-1, // IntAmount
		-1, // HumanAmount
		18, // Rule
		20, // RuleList
		

	},
	gotoRow{ // S16
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // Inittype
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S17
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // Inittype
		-1, // IntAmount
		-1, // HumanAmount
		22, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S18
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // Inittype
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S19
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // Inittype
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S20
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // Inittype
		-1, // IntAmount
		-1, // HumanAmount
		22, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S21
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // Inittype
		-1, // IntAmount
		-1, // HumanAmount
		18, // Rule
		24, // RuleList
		

	},
	gotoRow{ // S22
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // Inittype
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S23
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // Inittype
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S24
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // Inittype
		-1, // IntAmount
		-1, // HumanAmount
		22, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S25
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // Inittype
		-1, // IntAmount
		27, // HumanAmount
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S26
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // Inittype
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S27
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // Inittype
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S28
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // Inittype
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S29
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		31, // Action
		-1, // Inittype
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S30
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // Inittype
		34, // IntAmount
		-1, // HumanAmount
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S31
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // Inittype
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S32
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // Inittype
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S33
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // Inittype
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S34
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // Inittype
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S35
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // Inittype
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S36
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // Inittype
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S37
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		38, // Action
		-1, // Inittype
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S38
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // Inittype
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Rule
		-1, // RuleList
		

	},
	
}
