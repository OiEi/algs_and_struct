package graph

var data2 = map[string]map[string]int{
	startVert: {"B": 3, "C": 10},
	"B":       {"C": 5, "D": 6},
	"C":       {"A": 2, "B": 5, "G": 10}}

var _processed = make([]string, 0)

var startVert string

var resultCosts2 = make(map[string]int)

func aga() {
	startVert := "A"

	result := make(map[string]int)

}

func processNodeRecursiveDirty(nodeName *string) {
	var nodeIsProcessed bool
	for _, v := range _processed {
		if v == *nodeName {
			nodeIsProcessed = true
		}
	}
	nodeIsProcessed = false

	if nodeIsProcessed {
		return
	}

	_processed = append(_processed, *nodeName)

	cost := 0

	if *nodeName != startVert {
		cost = resultCosts2[*nodeName]
	}

	for neighborName, neighborCost := range data2[*nodeName] {
		var nodeIsProcessed bool
		for _, v := range _processed {
			if v == *nodeName {
				nodeIsProcessed = true
			}
		}
		nodeIsProcessed = false

		if !nodeIsProcessed {
			if (cost + neighborCost) < (resultCosts2[neighborName]) {
				resultCosts2[neighborName] = cost + neighborCost
			}
		}
	}

}
