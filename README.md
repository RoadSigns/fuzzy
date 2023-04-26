# Fuzzy
> Fuzzy is a package that provides algorithms that are used to see how similar strings are.

One to two paragraph statement about your product and what it does.

## Installation

```sh
go get github.com/roadsigns/fuzzy
```

## Algorithms available

### Match
Match checks to see if the two strings are a partial match and if so return true, if they aren't then it returns a false. Match is case-sensitive.

```go
import "github.com/roadsigns/fuzzy"

fuzzy.Match("wtrmln", "watermelon")
```

### Levenshtein
Levenshtein allows us to calculate the edit distance between the source and the target.
Additional information for this algorithm can be found https://en.wikipedia.org/wiki/Levenshtein_distance

```go
import "github.com/roadsigns/fuzzy"

fuzzy.Levenshtien("kitten", "sitting")
```

### Jaro
The Jaro similarity algorithm is a measure of the similarity between two strings. It is commonly used to calculate the similarity between two strings of text.
Additional information for this algorithm can be found https://en.wikipedia.org/wiki/Jaro%E2%80%93Winkler_distance#Jaro_Similarity

```go
import "github.com/roadsigns/fuzzy"

fuzzy.Jaro("gorilla", "guerrilla")
```

### Soundex
The Soundex algorithm is a phonetic algorithm for indexing words by their pronunciation.
Additional information for this algorithm can be found https://en.wikipedia.org/wiki/Soundex
```go
import "github.com/roadsigns/fuzzy"

fuzzy.Soundex("Smith")  // S530
fuzzy.Soundex("Smythe") // S530
```

### Sørensen–Dice coefficient
The Sørensen–Dice coefficient algorithm is a similarity coefficient that is used to compare the similarity of two strings.
Additional information for this algorithm can be found https://en.wikipedia.org/wiki/S%C3%B8rensen%E2%80%93Dice_coefficient
```go
import "github.com/roadsigns/fuzzy"

fuzzy.SorensenDice("elon musk", "colon musk")
```  

### Q-Gram
The Q-Gram algorithm is a similarity algorithm that It does this by dividing each string into substrings of length Q and then comparing the substrings to determine how many of them are the same in both strings.
Additional information for this algorithm can be found https://en.wikipedia.org/wiki/N-gram
```go
import "github.com/roadsigns/fuzzy"

fuzzy.QGram("John Jones", "Jon Jones", 3)
```