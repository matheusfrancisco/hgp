package main


// Bitset represents a 2D bitset with a given length
type Bitset struct {
	length        int
	numbersBitset [][]uint64
}

func NewBitset(rows, lengthPerRow int) *Bitset {
	if rows <= 0 || lengthPerRow <= 0 {
		panic("rows and lengthPerRow must be greater than 0")
	}
	// Calculate the number of uint64s needed for each row
	//this should  be used to  make sure that the number
	// of elements is a multiple of 64
	numberSizeOfElements := (lengthPerRow + 63) / 64

	// Initialize the 2D slice
	numbersBitset := make([][]uint64, rows)
	for i := 0; i < rows; i++ {
		numbersBitset[i] = make([]uint64, numberSizeOfElements)
	}

	return &Bitset{
		length:        lengthPerRow,
		numbersBitset: numbersBitset,
	}
}

func makeBitsetAndPopulate(infoFile *FileType) *Bitset {
	bitset := NewBitset(91, infoFile.NumberOfLines)

	for i, item := range infoFile.Items {
		bitset.SetBits(i, item)
	}

	return bitset
}
func (b *Bitset) SetBits(index int, numbers []int) {
	for _, number := range numbers {
		indexToSet := int(index / 64)
		offset := uint(index % 64)
		mask := uint64(1) << offset
		b.SetBit(number, indexToSet, mask)
	}
}

func (b *Bitset) SetBit(row int, index int, mask uint64) {
	b.numbersBitset[row][index] |= mask
}

func (b *Bitset) IsSet(row int, index int, mask uint64) bool {
  return (b.numbersBitset[row][index] & mask) != 0
}
