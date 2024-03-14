package package202310

/*
src:
	https://leetcode.com/problems/count-vowels-permutation/description/

desc:
	Given an integer n, your task is to count how many strings of length n can be formed under the following rules:
		Each character is a lower case vowel ('a', 'e', 'i', 'o', 'u')
		Each vowel 'a' may only be followed by an 'e'.
		Each vowel 'e' may only be followed by an 'a' or an 'i'.
		Each vowel 'i' may not be followed by another 'i'.
		Each vowel 'o' may only be followed by an 'i' or a 'u'.
		Each vowel 'u' may only be followed by an 'a'.
	Since the answer may be too large, return it modulo 10^9 + 7.

	Example 1:
		Input: n = 1
		Output: 5
		Explanation: All possible strings are: "a", "e", "i" , "o" and "u".
	Example 2:
		Input: n = 2
		Output: 10
		Explanation: All possible strings are: "ae", "ea", "ei", "ia", "ie", "io", "iu", "oi", "ou" and "ua".
	Example 3:
		Input: n = 5
		Output: 68
	Constraints:
		1 <= n <= 2 * 10^4
*/

// CountVowelPermutation is Recursive soloution.
func CountVowelPermutation(n int) int {
	a, e, i, o, u := countVowelPermutation(n)
	return (a + e + i + o + u) % (1e9 + 7)
}

func countVowelPermutation(n int) (int, int, int, int, int) {
	const MOD = 1e9 + 7
	if n != 1 {
		a, e, i, o, u := countVowelPermutation(n - 1)
		return (e + i + u) % MOD, (a + i) % MOD, (e + o) % MOD, i % MOD, (i + o) % MOD
	}
	return 1, 1, 1, 1, 1
}

// CountVowelPermutationDynamic is provided Dynamic soloution
func CountVowelPermutationDynamic(n int) int {
	const MOD = 1e9 + 7
	a, e, i, o, u := 1, 1, 1, 1, 1

	for x := 1; x < n; x++ {
		a, e, i, o, u = (e)%MOD, (a+i)%MOD, (a+e+o+u)%MOD, (i+u)%MOD, (a)%MOD
	}
	return (a + e + i + o + u) % MOD
}

// CountVowelPermutationDynamic2 is Another Dynamic soloution (NOT MINE):
func CountVowelPermutationDynamic2(n int) int {
	const MOD = 1e9 + 7

	a, e, i, o, u := 1, 1, 1, 1, 1

	for j := 1; j < n; j++ {
		aNext := e
		eNext := (a + i) % MOD
		iNext := (a + e + o + u) % MOD
		oNext := (i + u) % MOD
		uNext := a

		a, e, i, o, u = aNext, eNext, iNext, oNext, uNext
	}

	return (a + e + i + o + u) % MOD
}
