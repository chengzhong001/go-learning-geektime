package pipefilter

import (
	"errors"
	"strings"
)

var ErrSplitFilterWrongFormat = errors.New("input data should be string")

//定义error变量名必须是以Err开头

type SplitFilter struct {
	delimiter string
}

func NewSplitFilter(delimiter string) *SplitFilter {
	return &SplitFilter{delimiter}
}

func (sf *SplitFilter) Process(data Request) (Response, error) {
	str, ok := data.(string)
	if !ok {
		return nil, ErrSplitFilterWrongFormat
	}
	parts := strings.Split(str, sf.delimiter)
	return parts, nil
}
