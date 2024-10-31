package main

// num: an integer determining how many words to modify
// i: an integer representing the position to start modifying
// words: a slice of strings where each word may be modified
// f: takes a string and returns a string, applied to each world
// starter: starting point of the loop

func NumModify(num int, i int, words []string, f func(string) string) {
	for starter := i - num; starter < i && starter >= 0; starter++ {
		words[starter] = f(words[starter])
	}
}
 
                  //////////  example ////////
									
// words := []string{"apple", "banana", "cherry", "date", "fig"}
// NumModify(2, 4, words, strings.ToUpper)


// i - num : 4 - 2 which gives us 2
//  starter < i : starter < 4 which means it will consider indeces 2 and 3 
   /// Index 2: "cherry"
   /// Index 3: "date"

// starter >= 0 : to avoid out of bounds 
// 	words[starter] = f(words[starter]) 
   /// so words[2] is "cherry"
	 /// f(words[2]) is strings.ToUpper("cherry"), which results in "CHERRY"
	 /// assignment words[2] = "CHERRY" updates the slice
	 /// this happens for index[3] as well
/// words[starter] = f(words[starter]) : in the end the modified string is returned back to words
