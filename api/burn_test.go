package api

import (
	"fmt"
	"github.com/copernet/whcexplorer/model"
	"testing"
)

func TestCountBurn(t *testing.T) {

	cnt, err := model.GetBurnCount()
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(cnt)
}

func TestGetBurnList(t *testing.T) {

	list, err := model.GetBurnList(5, 2)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(list)

	block := model.GetLastBlock()
	for _, info := range list {
		generateProcess(info, block)
	}
	fmt.Println(list)
}

func TestSummary(t *testing.T) {
	summary, err := model.GetBurnSummary()
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(summary)
}

func TestCase(t *testing.T) {

	te := 10
	switch te {
	case 10:
		fmt.Println("I am 10")
	case 11:
		fmt.Println("I am 11")
	case 12:
		fmt.Println("I am 12")

	}

}
