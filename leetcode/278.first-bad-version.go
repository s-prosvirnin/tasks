package leetcode

// https://leetcode.com/problems/first-bad-version

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	pos := search(1, n+1)

	return pos
}

func search(startPos int, endPos int) int {
	posInterval := endPos - startPos

	if posInterval == 0 {
		return startPos
	}

	if posInterval == 1 {
		isBadVersion := isBadVersion(startPos)
		if isBadVersion == true {
			return startPos
		}
		return startPos + 1
	}

	middlePos := int((posInterval)/2) + startPos
	isBadVersion := isBadVersion(middlePos)
	var searchPos int
	if isBadVersion == false {
		searchPos = search(middlePos, endPos)
	} else {
		searchPos = search(startPos, middlePos)
	}

	return searchPos
}
