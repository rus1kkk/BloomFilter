package cmd

import (
	"bufio"
	"fmt"
	"hash"
	"hash/fnv"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

type BloomFilter struct {
	size        int
	hashes      []hash.Hash64
	bitArray    []bool
	numElements int
}

func NewBloomFilter(size int, hashCount int) *BloomFilter {
	bitArray := make([]bool, size)
	hashes := make([]hash.Hash64, hashCount)

	for i := 0; i < hashCount; i++ {
		hashes[i] = fnv.New64()
	}

	return &BloomFilter{
		size:        size,
		bitArray:    bitArray,
		hashes:      hashes,
		numElements: 0,
	}
}

func (bf *BloomFilter) AddFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var words []string
	collisionMap := make(map[int]map[uint64][]string)
	for i := range bf.hashes {
		collisionMap[i] = make(map[uint64][]string)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	for _, word := range words {
		for hashFuncIndex, hashFunc := range bf.hashes {
			hashFunc.Reset()
			hashFunc.Write([]byte(word))
			hash := hashFunc.Sum64()
			index := hash % uint64(bf.size)

			collisionMap[hashFuncIndex][hash] = append(collisionMap[hashFuncIndex][hash], word)
			bf.bitArray[index] = true
			bf.numElements++
		}
	}

	//print collisions
	for hashFuncIndex, hashToWords := range collisionMap {
		for hash, words := range hashToWords {
			if len(words) > 1 {
				fmt.Printf("Collision in hash%d (%x):\n", hashFuncIndex+1, hash)
				for _, word := range words {
					fmt.Println(" -", word)
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func (bf *BloomFilter) Check(word string) bool {
	for _, hashFunc := range bf.hashes {
		hashFunc.Reset()
		hashFunc.Write([]byte(word))
		index := hashFunc.Sum64() % uint64(bf.size)
		if !bf.bitArray[index] {
			fmt.Printf("%s is DEFINITELY NOT in the set \n", word)
			return false
		}
	}
	fmt.Printf("%s is PROBABLY in the set\n", word)
	return true
}

func (bf *BloomFilter) FalsePositiveProbability() float64 {
	if bf.numElements == 0 {
		return 0.0
	}
	return math.Pow(1-math.Exp(-float64(len(bf.hashes)*bf.numElements)/float64(bf.size)), float64(len(bf.hashes)))
}

// BloomFilterCmd represents the BloomFilter command
var BloomFilterCmd = &cobra.Command{
	Use:   "bloomfilter",
	Short: "Bloom filter implementation",
	Long: `Probabilistic data structure for membership testing.
Example: bloomfilter 1000 3 data.txt hello`,
	Args: cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("BloomFilter called")

		size, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("Invalid size: %v\n", err)
			panic(err)
		}
		hashCount, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Printf("Invalid hash count: %v\n", err)
			panic(err)
		}
		filePath := args[2]
		checkWord := args[3]

		BloomFilter := NewBloomFilter(size, hashCount)
		BloomFilter.AddFile(filePath)
		BloomFilter.Check(checkWord)
		fmt.Print("The probability of a false positive = ")
		fmt.Println(BloomFilter.FalsePositiveProbability())

		fmt.Println("BloomFilter end")
	},
}

func init() {
	rootCmd.AddCommand(BloomFilterCmd)
}
