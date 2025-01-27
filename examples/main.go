package main

import (
	"context"
	"fmt"
	gopipedrive "github.com/whatcrm/go-pipedrive"
)

func main() {
	clientID := "3901a5b464e88c6f"
	clientSecret := "c0b1303ae63c8f334b9fdc285f5a7cb815bfaa8c"

	redirectURI := "https://api.pipedrive.com/api/v2/activities"
	tokenR := "13820726:22706498:23829d97784ab30ce17f68cab26d3692a0271eef"

	client, err := gopipedrive.NewClient(clientID, clientSecret, redirectURI)
	if err != nil {
		fmt.Println("error: ", err)
	}
	//client.Token = "v1u:AQIBAHj-LzTNK2yuuuaLqifzhWb9crUNKTpk4FlQ9rjnXqp_6AE-mSRsNEc-qt-JZF1QvUyHAAAAfjB8BgkqhkiG9w0BBwagbzBtAgEAMGgGCSqGSIb3DQEHATAeBglghkgBZQMEAS4wEQQMGqcEIUgSSD15bcXPAgEQgDtLA6W-OJ_ytVMrWzkVH7tyBFy3C5ZLb1hgRXukifjFKHrPSN0XT5mLYmrNgQAPNro5MhWeqqjIaVI6zA:SC_C4md7GxgWi1HRkdLnG323HkZ5NmS2V-4GTVVFTUnlPOVi9kkP-v0VC5DjRN_ECMhvv5aW-kOtR57R6NQYEBp2l-yTvmh2F17jbdPyCEccU0Y_0Uh-M4N3NcJA2gknYpa_UDNrVPqtEk5t8UNnYczX14Cyi1A3N2_gA_wJRXT2QaMzsu9ntOTAVLEB4HtFVPz31nKkEo_JXRgu7uo996WcShHxp22ULH36NxLoV1qyPv7U1sSounQePbqlCj930OMeV_03cFRB7DXnQ2cAC7DSVLJUtiMXR1L1879JzybZ_WVytsnmEsjbs_EfsKqoFGlgkREWW15LO803L3h-qqqCxaVcknYU9YKcztB2OG500wUdibaf8IjUduJdbpsomc_wiGUKamZctcsyfGz26oLmlfFJCUdSWsaot6oGOks121-b6qxtKg"

	ctx := context.Background()

	res, err := client.RefreshToken(ctx, tokenR)
	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println("result: ", res)
}
