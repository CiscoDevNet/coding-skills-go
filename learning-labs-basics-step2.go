//  THIS SAMPLE APPLICATION AND INFORMATION IS PROVIDED "AS IS" WITHOUT WARRANTY
//  OF ANY KIND BY CISCO, EITHER EXPRESSED OR IMPLIED, INCLUDING BUT NOT LIMITED
//  TO THE IMPLIED WARRANTIES OF MERCHANTABILITY FITNESS FOR A PARTICULAR
//  PURPOSE, NONINFRINGEMENT, SATISFACTORY QUALITY OR ARISING FROM A COURSE OF
//  DEALING, LAW, USAGE, OR TRADE PRACTICE. CISCO TAKES NO RESPONSIBILITY
//  REGARDING ITS USAGE IN AN APPLICATION, AND IT IS PRESENTED ONLY AS AN
//  EXAMPLE. THE SAMPLE CODE HAS NOT BEEN THOROUGHLY TESTED AND IS PROVIDED AS AN
//  EXAMPLE ONLY, THEREFORE CISCO DOES NOT GUARANTEE OR MAKE ANY REPRESENTATIONS
//  REGARDING ITS RELIABILITY, SERVICEABILITY, OR FUNCTION. IN NO EVENT DOES
//  CISCO WARRANT THAT THE SOFTWARE IS ERROR FREE OR THAT CUSTOMER WILL BE ABLE
//  TO OPERATE THE SOFTWARE WITHOUT PROBLEMS OR INTERRUPTIONS. NOR DOES CISCO
//  WARRANT THAT THE SOFTWARE OR ANY EQUIPMENT ON WHICH THE SOFTWARE IS USED WILL
//  BE FREE OF VULNERABILITY TO INTRUSION OR ATTACK. THIS SAMPLE APPLICATION IS
//  NOT SUPPORTED BY CISCO IN ANY MANNER. CISCO DOES NOT ASSUME ANY LIABILITY
//  ARISING FROM THE USE OF THE APPLICATION. FURTHERMORE, IN NO EVENT SHALL CISCO
//  OR ITS SUPPLIERS BE LIABLE FOR ANY INCIDENTAL OR CONSEQUENTIAL DAMAGES, LOST
//  PROFITS, OR LOST DATA, OR ANY OTHER INDIRECT DAMAGES EVEN IF CISCO OR ITS
//  SUPPLIERS HAVE BEEN INFORMED OF THE POSSIBILITY THEREOF.-->

package main

import(
	"github.com/CiscoDevNet/coding-skills-go/controller"
	"github.com/antonholmquist/jason"
	"log"
)

func main() {

	// Use the same library and method we defined before
	body := controller.GetHosts("https://sandboxapic.cisco.com")

	// Use the JSON library, jason, to do the heavy lifting of creating go objects
	v, _ := jason.NewObjectFromBytes(body)

	// Since the result is an array of objects, we need to use GetObjectArray()
	r, _ := v.GetObjectArray("response")

	// Loop over the objects
	for _, obj := range r {

		// Use the jason library's GetString() method to get the right elements of the
		// returned JSON.
		//
		// Example JSON
		// {
		//   "connectedInterfaceId": "30bb14c1-8fb6-45c4-8f6d-5b845a7f448c",
		//   "connectedInterfaceName": "GigabitEthernet2/0/2",
		//   "connectedNetworkDeviceId": "7895a45f-47aa-42ee-9d06-c66d3b784594",
		//   "connectedNetworkDeviceIpAddress": "40.0.2.18",
		//   "hostIp": "40.0.5.12",
		//   "hostMac": "00:50:56:8A:27:A3",
		//   "hostType": "WIRED",
		//   "id": "8f41bef8-698c-4701-af14-471e910ed9ff",
		//   "lastUpdated": "September 29, 2014 1:54:13 PM PDT",
		//   "numUpdates": 1,
		//   "source": 200,
		//   "userStatus": "Active",
		//   "vlanId": "1"
		// }

		iFace, err := obj.GetString("connectedInterfaceId")
		deviceType, err := obj.GetString("hostType")

		log.Println("id: " + iFace + " type: " + deviceType)

		if (err != nil) {
			log.Println(err)
		}
	}

}
