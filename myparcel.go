package myparcel

import (
	"net/http"
	"sync"
	"time"
)

const (
	// Carriers
	CARRIER_POSTNL      = 1 // (PostNL)
	CARRIER_BPOST       = 2 // (bpost. Only available on Sendmyparcel.be)
	CARRIER_CHEAP_CARGO = 3 // (CheapCargo/pallets)
	CARRIER_DPD         = 4 // (DPD. Only available on Sendmyparcel.be)
	CARRIER_INSTABOX    = 5 // (Instabox. Only available on MyParcel.nl)
	CARRIER_UPS         = 8 // (UPS. Only available on MyParcel.nl)

	// Package types
	PARCEL_PACKAGE = 1 // This is the standard package type used for NL, EU and Global shipments. It supports a variety of additional options such as insurance, xl format etc. We will look at these options in more detail later. This package is most commonly used when creating shipments.
	PARCEL_MAILBOX = 2 // This package type is only available on MyParcel.nl and Flespakket for NL shipment that fit into a mailbox. It does not support additional options. Note: If you still make the request with additional options, bear in mind that you need to pay more than is necessary!
	PARCEL_LETTER  = 3 // This package type is available on MyParcel.nl for NL, EU and Global shipments. The label for this shipment is unpaid meaning that you will need to pay the postal office/courier to sent this letter/package. Therefore, it does not support additional options.
	PARCEL_STAMP   = 4 // This package type is only available on MyParcel.nl for NL shipments and does not support any additional options. Its price is calculated using the package weight. Note: This shipment will appear on your invoice on shipment_status 2 (pending - registered) as opposed to all other package types, which won't appear on your invoice until shipment status 3 (enroute - handed to carrier).

	// Delivery types
	DELIVERY_MORNING  = 1 // Morning delivery
	DELIVERY_STANDARD = 2 // Standard delivery
	DELIVERY_EVENING  = 3 // Evening delivery
	DELIVERY_PICKUP   = 4 // Pickup point delivery
)

// Client is the MyParcel client.
type Client struct {
	apiBaseURL string
	ApiKey     string
	httpClient *http.Client
	sync.Mutex
}

// Shipment struct
type Shipment struct {
	ReferenceIdentifier string `json:"reference_identifier"` // Required: No. Arbitrary reference indentifier to identify this shipment.
	Recipient           struct {
		Cc         string `json:"cc"`          // Required: yes. The address country code.
		Region     string `json:"region"`      // Required: no. The region, department, state or province of the address.
		City       string `json:"city"`        // Required: yes. The address city.
		Street     string `json:"street"`      // Required: yes. The address street name. When shipping to an international destination, you may include street number in this field.
		Number     string `json:"number"`      // Required: yes for domestic shipments in NL and BE. Street number.
		PostalCode string `json:"postal_code"` // Required: yes for NL and EU destinations except for IE. The address postal code.
		Person     string `json:"person"`      // Required: yes. The person at this address. Up to 40 characters long.
		Phone      string `json:"phone"`       // Required: no. The address phone.
		Email      string `json:"email"`       // Required: no The address email.
	} `json:"recipient"` // Required: Yes for application/vnd.shipment_recipients+json Use this field when you want to create multiple shipments for multiple addresses.
	Options struct {
		PackageType   int       `json:"package_type"`            // Required: yes. The package type. For international shipment only package type 1 (package) is allowed.
		OnlyRecipient int       `json:"only_recipient"`          // Required: No. Deliver the package to the recipient only.
		DeliveryType  int       `json:"delivery_type" omitempty` // Required: Yes if delivery_date has been specified. The delivery type for the package.
		DeliveryDate  time.Time `json:"delivery_date" omitempty` // Required: Yes if delivery type has been specified. The delivery date time for this shipment.
		Signature     int       `json:"signature"`               // Required: No. Package must be signed for.
		Return        int       `json:"return"`                  // Required: No. Return the package if the recipient is not home.
		Insurance     struct {
			Amount   int    `json:"amount"`   // Required: yes. The amount is without decimal separators (in cents).
			Currency string `json:"currency"` // Required: yes. The insurance currency code. Must be one of the following: EUR.
		} `json:"insurance"` // Required: No. Insurance price for the package.
		LargeFormat      int    `json:"large_format"`      // Required: No. Large format package.
		LabelDescription string `json:"label_description"` // Required: No. This description will appear on the shipment label. Note: This will be overridden for return shipment by the following: Retour – 3SMYPAMYPAXXXXXX
		AgeCheck         int    `json:"age_check"`         // Required: No. The recipient must sign for the package and must be at least 18 years old.
	} `json:"options"` // Required: Yes. The shipment options.
	Carrier int `json:"carrier"` // Required: Yes. The carrier that will deliver the package.
}

// NewClient returns a new MyParcel client.
// The API key is required to use the MyParcel API.
func NewClient(apiKey string) *Client {
	return &Client{
		apiBaseURL: "https://api.myparcel.nl",
		httpClient: &http.Client{},
		ApiKey:     apiKey,
	}
}

// CreateShipment creates a new shipment.
func (c *Client) CreateShipment(shipment Shipment) (string, error) {

}
