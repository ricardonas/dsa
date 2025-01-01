func longestConsecutive(nums []int) int {
	// Convert nums array into unique hash table (set) for fast lookup O(1)
    numsSet := make(map[int]struct{})
	for _, n := range nums {
		numsSet[n] = struct{}{}
	}

    // Number of the longest sequence built.
	longest := 0

    // Loop through each number to check sequences.
	for n := range numsSet {
        // Verify if the current number does not have a number before.
		if _, ok := numsSet[n-1]; !ok {
			seqLen := 1
            // Loop through until there's a next sequence.
			for {	
				if _, ok = numsSet[n+seqLen]; ok {
					seqLen++
					continue
				}
                // Update the longest if it less than current built sequence.
				longest = max(seqLen, longest)
				break
			}
		}
	}

	return longest
}
