package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAppParameters(t *testing.T) {
	params := AppParameters()
	t.Run("Parameters check", func(t *testing.T) {
		require.NotNil(t, params)
		require.IsType(t, Params{}, params)
	})
}

func logFilenameTests() []struct {
	descr             string
	filename          string
	expectedDataCount int
	expectedError     bool
} {
	return []struct {
		descr             string
		filename          string
		expectedDataCount int
		expectedError     bool
	}{
		{
			descr:             "Filename is empty",
			filename:          "",
			expectedDataCount: 0,
			expectedError:     true,
		},
		{
			descr:             "Filename is wrong",
			filename:          "sadfsadfa.sss",
			expectedDataCount: 0,
			expectedError:     true,
		},
		{
			descr:             "Filename is correct",
			filename:          "journal.log",
			expectedDataCount: 14,
			expectedError:     false,
		},
	}
}

func TestLogData(t *testing.T) {
	tests := logFilenameTests()

	for _, td := range tests {
		td := td
		t.Run(td.descr, func(t *testing.T) {
			logData, err := LogData(td.filename)
			require.Equal(t, td.expectedDataCount, len(logData))
			if td.expectedError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func logStatsTests() []struct {
	descr           string
	level           string
	expectedTotal   string
	expectedByLevel string
} {
	return []struct {
		descr           string
		level           string
		expectedTotal   string
		expectedByLevel string
	}{
		{
			descr:           "Level info",
			level:           "info",
			expectedTotal:   "Total lines number: 14",
			expectedByLevel: "Lines number with level 'info': 8",
		},
		{
			descr:           "Level warning",
			level:           "warning",
			expectedTotal:   "Total lines number: 14",
			expectedByLevel: "Lines number with level 'warning': 4",
		},
		{
			descr:           "Level error",
			level:           "error",
			expectedTotal:   "Total lines number: 14",
			expectedByLevel: "Lines number with level 'error': 2",
		},
		{
			descr:           "Empty level",
			level:           "",
			expectedTotal:   "Total lines number: 14",
			expectedByLevel: "Lines number with level '': 0",
		},
		{
			descr:           "Wrong level",
			level:           "asdfsafd",
			expectedTotal:   "Total lines number: 14",
			expectedByLevel: "Lines number with level 'asdfsafd': 0",
		},
	}
}

func TestLogStats(t *testing.T) {
	filename := "journal.log"

	tests := logStatsTests()
	for _, td := range tests {
		td := td
		t.Run(td.descr, func(t *testing.T) {
			logData, _ := LogData(filename)
			stats := LogStats(logData, td.level)
			require.Equal(t, 2, len(stats))
			require.Equal(t, td.expectedTotal, stats[0])
			require.Equal(t, td.expectedByLevel, stats[1])
		})
	}
}

func TestWriteLog(t *testing.T) {
	filename := "journal.log"
	outputFilename := "out.txt"
	os.Remove(outputFilename)

	t.Run("Output file must be written", func(t *testing.T) {
		require.NoFileExists(t, outputFilename)
		logData, _ := LogData(filename)
		stats := LogStats(logData, "info")
		WriteStats(stats, outputFilename)
		require.FileExists(t, outputFilename)
	})
}
