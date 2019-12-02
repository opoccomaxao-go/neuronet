package Core

import (
	"fmt"
	"math/rand"
	"time"
)

type IOVector []NetDataType

type Sample struct {
	In  IOVector
	Out IOVector
}

func CreateSample(input IOVector, output IOVector) Sample {
	return Sample{In: input, Out: output}
}

func CreateIOVector(array ...NetDataType) IOVector {
	return array
}

func CreateIOVectorByLength(length int) IOVector {
	return make(IOVector, length)
}

func (s *Sample) ToString() string {
	return fmt.Sprintf("f(%v)->%f", s.In, s.Out)
}

type SampleArray []Sample

func (s SampleArray) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
}
