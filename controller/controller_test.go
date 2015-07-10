package controller

import(
	"github.com/CiscoDevNet/coding-skills-go/controller"
	"testing"
	"fmt"
)

func TestController(t *testing.T) {

	body := controller.GetHosts("https://sandboxapic.cisco.com")

	fmt.Println(body)

	if body == nil {
		t.Errorf("error")
	}

}
