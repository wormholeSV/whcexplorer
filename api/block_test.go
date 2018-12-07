package api

import (
	"fmt"
	"github.com/copernet/whcexplorer/model"
	"testing"
)

func TestFeeRates(t *testing.T)  {

	summary,err :=model.GetFeeRateSummary(1266634)
	fmt.Println(err)
	fmt.Println(summary)
}
