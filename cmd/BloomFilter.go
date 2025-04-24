package cmd

import (
	"bufio"
	"fmt"
	"hash"
	"hash/fnv"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

type BloomFilter struct {
	size     int
	hashes   []hash.Hash64
	bitArray []bool
}

func NewBloomFilter(size int, hashCount int) *BloomFilter {
	bitArray := make([]bool, size)
	hashes := make([]hash.Hash64, hashCount)

	for i := 0; i < hashCount; i++ {
		hashes[i] = fnv.New64()
	}

	return &BloomFilter{
		size:     size,
		bitArray: bitArray,
		hashes:   hashes,
	}
}

func (bf *BloomFilter) AddFile(filePath string) {
	file, err := os.Open(filePath)

	var words []string

	if err != nil {
		log.Fatal(err)
	}

	Scanner := bufio.NewScanner(file)
	Scanner.Split(bufio.ScanWords)

	for Scanner.Scan() {
		words = append(words, Scanner.Text())
	}

	for _, word := range words {
		for _, hashFunc := range bf.hashes {
			hashFunc.Reset()
			hashFunc.Write([]byte(word))
			index := hashFunc.Sum64() % uint64(bf.size)
			bf.bitArray[index] = true
		}
	}

	if err := Scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func (bf *BloomFilter) Check(word string) bool {
	for _, hashFunc := range bf.hashes {
		hashFunc.Reset()
		hashFunc.Write([]byte(word))
		index := hashFunc.Sum64() % uint64(bf.size)
		if !bf.bitArray[index] {
			fmt.Printf("%s is DEFINITELY NOT in the set", word)
			return false
		}
	}
	fmt.Printf("%s is PROBABLY in the set\n", word)
	return true
}

// BloomFilterCmd represents the BloomFilter command
var BloomFilterCmd = &cobra.Command{
	Use:   "bloomfilter",
	Short: "Bloom filter implementation",
	Long: `Probabilistic data structure for membership testing.
Example: bloomfilter 1000 3 data.txt hello`,
	Args: cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("\nBloomFilter called\n")

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

		fmt.Println("\nBloomFilter end")
	},
}

func init() {
	rootCmd.AddCommand(BloomFilterCmd)
}
