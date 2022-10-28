package main

import (
	"fmt"
	"math"
)

func isEnoughChangeForTickets(moneyForTickets []int) string {

	const ticketCost = 25
	isEnoughChange := true
	dollarsBillsMap := map[int]int{
		100: 0,
		50:  0,
		25:  0,
	}
	for _, moneyForTicket := range moneyForTickets {
		change := moneyForTicket - ticketCost
		if change > 0 {
			for _, denomination := range []int{100, 50, 25} {
				billsCount := dollarsBillsMap[denomination]
				if change >= denomination && billsCount > 0 {
					changeBillsCount := int(math.Floor(float64(change) / float64(denomination)))
					if changeBillsCount >= billsCount {
						dollarsBillsMap[denomination] = 0
						change -= billsCount * denomination
					} else {
						dollarsBillsMap[denomination] -= changeBillsCount
						change -= changeBillsCount * denomination
					}
				}
			}
		}

		if change > 0 {
			isEnoughChange = false
			break
		}

		dollarsBillsMap[moneyForTicket]++
	}

	if isEnoughChange == false {
		return "NO"
	}

	return "YES"
}

func doTestIsEnoughChangeForTickets(moneyForTickets []int, expectedResult bool) {
	isEnoughChangeForTicketsResult := isEnoughChangeForTickets(moneyForTickets)
	expectedStringResult := "NO"
	if true == expectedResult {
		expectedStringResult = "YES"
	}
	if isEnoughChangeForTicketsResult != expectedStringResult {
		fmt.Printf("TEST NOT PASSED. moneyForTickets: %v, expectedResult: %s, isEnoughChangeForTicketsResult: %s\n", moneyForTickets, expectedStringResult, isEnoughChangeForTicketsResult)

		return
	}

	fmt.Printf("test passed. moneyForTickets: %v, expectedResult: %s, isEnoughChangeForTicketsResult: %s\n", moneyForTickets, expectedStringResult, isEnoughChangeForTicketsResult)
}

func main() {
	doTestIsEnoughChangeForTickets(
		[]int{25, 25, 25, 25, 50, 100, 50},
		true,
	)
	doTestIsEnoughChangeForTickets(
		[]int{25, 25, 50},
		true,
	)
	doTestIsEnoughChangeForTickets(
		[]int{25, 100},
		false,
	)
	doTestIsEnoughChangeForTickets(
		[]int{50},
		false,
	)
	doTestIsEnoughChangeForTickets(
		[]int{25, 25, 50, 50, 100},
		false,
	)
	doTestIsEnoughChangeForTickets(
		[]int{},
		true,
	)
	doTestIsEnoughChangeForTickets(
		[]int{25, 25, 50, 50, 25, 100, 25, 25, 50, 50},
		true,
	)
	doTestIsEnoughChangeForTickets(
		[]int{25, 25, 50, 50, 25, 100, 25, 25, 50, 50, 50},
		false,
	)
	doTestIsEnoughChangeForTickets(
		[]int{25, 25, 25, 25, 50, 100, 50},
		true,
	)
	doTestIsEnoughChangeForTickets(
		[]int{25, 25, 25, 25, 50, 100, 50, 25, 100, 25, 50},
		true,
	)
}
