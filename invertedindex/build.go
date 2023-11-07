package invertedindex


import (
	"bufio"
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strings"
)


// Builds inverted index with txt docs in a given path
func BuildInvertedIndex (path string)(ii InvertedIndex){
	ii = InvertedIndex{}
	ii.Token2docs = make(map[string][]int)

	// Walk through the directory and its subdirectories
	err := filepath.Walk(path, func(	filePath string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}

		// Check if the current item is a file and ends with ".txt"
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".txt") {
			// Read the file
			file, err := os.Open(filePath)
			if err != nil {
				fmt.Println(err)
				return err
			}
			defer file.Close()

			// Read the content of the file
			scanner := bufio.NewScanner(file)
			docID := len(ii.Docs)
			ii.Docs = append(ii.Docs, filePath) // Store the document path
			tokensInCurrentDoc := make(map[string]struct{})

			for scanner.Scan() {
				line := scanner.Text()
				// Tokenize the line (split by space)
				tokens := strings.Fields(line)
				for _, token := range tokens {
					// Store the token and the document ID in the inverted index
					if _, exists := tokensInCurrentDoc[token]; !exists {
						// Store the token and the document ID in the inverted index
						ii.Token2docs[token] = append(ii.Token2docs[token], docID)
						// Mark the token as processed for the current document
						tokensInCurrentDoc[token] = struct{}{}
					}
				}
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error building inverted index:", err)
	}

	return ii
}


func SaveInvertedIndex(ii InvertedIndex, filePath string) error {
	jsonData, err := json.Marshal(ii)
	if err != nil {
		fmt.Println("Error marshaling data:", err)
		return err
	}

	err = ioutil.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}

	fmt.Println("Inverted index saved to", filePath)
	return nil
}

// Load the inverted index from a JSON file
func LoadInvertedIndex(filePath string) (InvertedIndex, error) {
	jsonData, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return InvertedIndex{}, err
	}

	var ii InvertedIndex
	err = json.Unmarshal(jsonData, &ii)
	if err != nil {
		fmt.Println("Error unmarshaling data:", err)
		return InvertedIndex{}, err
	}

	fmt.Println("Inverted index loaded from", filePath)
	return ii, nil
}




