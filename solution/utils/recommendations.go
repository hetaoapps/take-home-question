package utils

import (
	"strings"
	"unicode"
)

// NormalizeText function to normalize the input text
func NormalizeText(text string) string {
	// Convert to lowercase
	text = strings.ToLower(text)

	// Remove punctuation and special characters
	var builder strings.Builder
	for _, char := range text {
		if unicode.IsLetter(char) || unicode.IsSpace(char) {
			builder.WriteRune(char)
		}
	}

	// Remove extra spaces
	normalizedText := strings.Join(strings.Fields(builder.String()), " ")
	return normalizedText
}

// Calculate Jaccard Similarity between two strings
func JaccardSimilarity(str1, str2 string) float64 {
	set1 := make(map[string]struct{})
	set2 := make(map[string]struct{})

	words1 := strings.Fields(str1)
	words2 := strings.Fields(str2)

	for _, word := range words1 {
		set1[word] = struct{}{}
	}

	for _, word := range words2 {
		set2[word] = struct{}{}
	}

	intersection := 0
	for word := range set1 {
		if _, exists := set2[word]; exists {
			intersection++
		}
	}

	union := len(set1) + len(set2) - intersection

	if union == 0 {
		return 0.0
	}

	return float64(intersection) / float64(union)
}

func getSimilarity(query1, query2 string) float64 {
	normalizedQuery1 := NormalizeText(query1)
	normalizedQuery2 := NormalizeText(query2)

	similarity := JaccardSimilarity(normalizedQuery1, normalizedQuery2)
	return similarity
}
