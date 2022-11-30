package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	sha256ChecksunProvided = "bb0997287d646189e850e2d96f7775a92cec04c38c756167bafe96a2be63fa78"
	filename               = "/home/rodrigo-oliveira/Downloads/JetBrains.Rider-2022.2.4.tar.gz"
	sha256Function         = "sha256"
)

type testCase struct {
	caseDescription string
	fileToCheck     string
	providedHash    string
	expectedReturn  bool
	expectedError   error
}

func TestMain(t *testing.T) {
	testCases := []testCase{
		{
			caseDescription: "When sha256 hash function got valid parameters",
			fileToCheck:     filename,
			providedHash:    sha256ChecksunProvided,
			expectedReturn:  true,
			expectedError:   nil,
		},
	}

	for _, test := range testCases {
		t.Run(test.caseDescription, func(t *testing.T) {
			result, err := checkSha256(test.fileToCheck, test.providedHash)

			assert.Equal(t, test.expectedError, err, nil)
			assert.Equal(t, test.expectedReturn, result)
		})

	}
}
