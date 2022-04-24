package matcher

import (
	"strings"
	"sync"
)

type MatchResult struct {
	Prefix string
	length int
	sync   *sync.RWMutex
}

func Matcher(prefixDataStore []string, inputString string) MatchResult {
	wg := &sync.WaitGroup{}
	output := MatchResult{Prefix: "", length: 0, sync: &sync.RWMutex{}}
	for _, prefix := range prefixDataStore {
		//empty prefix OR if prefix length is greater than input string then there is no need to match it
		if len(prefix) == 0 || len(prefix) > len(inputString) {
			continue
		} else if prefix[0] != inputString[0] { //if first letter doesn't match no need to compare rest of it
			continue
		} else { //match and modify output concurrenctly
			wg.Add(1)
			go matchAndModifyOutput(prefix, inputString, &output, wg)
		}
	}
	wg.Wait()
	return output
}

func matchAndModifyOutput(prefix, inputString string, output *MatchResult, wg *sync.WaitGroup) {
	defer wg.Done()
	output.sync.RLock()
	cond1 := strings.HasPrefix(inputString, prefix)
	cond2 := len(prefix) > output.length
	output.sync.RUnlock()
	if cond1 && cond2 {
		modifyOutput(prefix, output)
	}
}

func modifyOutput(prefix string, output *MatchResult) {
	output.sync.Lock()
	defer output.sync.Unlock()
	output.Prefix, output.length = prefix, len(prefix)
}
