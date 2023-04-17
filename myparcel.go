package myparcel

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	// Carriers
	CarrierPostnl     = 1 // (PostNL)
	CarrierBpost      = 2 // (bpost. Only available on Sendmyparcel.be)
	CarrierCheapCargo = 3 // (CheapCargo/pallets)
	CarrierDpd        = 4 // (DPD. Only available on Sendmyparcel.be)
	CarrierInstabox   = 5 // (Instabox. Only available on MyParcel.nl)
	CarrierUps        = 8 // (UPS. Only available on MyParcel.nl)

	// Package types
	ParcelPackage = 1 // This is the standard package type used for NL, EU and Global shipments. It supports a variety of additional options such as insurance, xl format etc. We will look at these options in more detail later. This package is most commonly used when creating shipments.
	ParcelMailbox = 2 // This package type is only available on MyParcel.nl and Flespakket for NL shipment that fit into a mailbox. It does not support additional options. Note: If you still make the request with additional options, bear in mind that you need to pay more than is necessary!
	ParcelLetter  = 3 // This package type is available on MyParcel.nl for NL, EU and Global shipments. The label for this shipment is unpaid meaning that you will need to pay the postal office/courier to sent this letter/package. Therefore, it does not support additional options.
	ParcelStamp   = 4 // This package type is only available on MyParcel.nl for NL shipments and does not support any additional options. Its price is calculated using the package weight. Note: This shipment will appear on your invoice on shipment_status 2 (pending - registered) as opposed to all other package types, which won't appear on your invoice until shipment status 3 (enroute - handed to carrier).

	// Delivery types
	DeliveryMorning  = 1 // Morning delivery
	DeliveryStandard = 2 // Standard delivery
	DeliveryEvening  = 3 // Evening delivery
	DeliveryPickup   = 4 // Pickup point delivery
)

// JSONTime is a custom time type that can be marshalled to JSON.
type JSONTime time.Time

// MarshalJSON returns the JSON encoding of time specified in the format
func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))

	if t.IsZero() {
		return []byte("null"), nil
	}

	return []byte(stamp), nil
}

// IsZero returns true if the time is zero.
func (t JSONTime) IsZero() bool {
	return time.Time(t).IsZero()
}

// Client is the MyParcel client.
type Client struct {
	apiBaseURL string
	ApiKey     string
	httpClient *http.Client
	UserAgent  string
}

// Shipment struct
type ShipmentStruct struct {
	Recipient           RecipientStruct `json:"recipient"`            // Required: Yes. The recipient address.
	ReferenceIdentifier string          `json:"reference_identifier"` // Required: No. Arbitrary reference indentifier to identify this shipment.
	Options             OptionsStruct   `json:"options"`              // Required: Yes. The shipment options.
	Carrier             int             `json:"carrier"`              // Required: Yes. The carrier that will deliver the package.
}

// Insurance struct
type RecipientStruct struct {
	Cc         string `json:"cc"`               // Required: yes. The address country code.
	Region     string `json:"region,omitempty"` // Required: no. The region, department, state or province of the address.
	City       string `json:"city"`             // Required: yes. The address city.
	Street     string `json:"street"`           // Required: yes. The address street name. When shipping to an international destination, you may include street number in this field.
	Number     string `json:"number"`           // Required: yes for domestic shipments in NL and BE. Street number.
	PostalCode string `json:"postal_code"`      // Required: yes for NL and EU destinations except for IE. The address postal code.
	Person     string `json:"person"`           // Required: yes. The person at this address. Up to 40 characters long.
	Phone      string `json:"phone,omitempty"`  // Required: no. The address phone.
	Email      string `json:"email,omitempty"`  // Required: no The address email.
}

// Options struct
type OptionsStruct struct {
	PackageType      int             `json:"package_type"`             // Required: yes. The package type. For international shipment only package type 1 (package) is allowed.
	OnlyRecipient    int             `json:"only_recipient,omitempty"` // Required: No. Deliver the package to the recipient only.
	DeliveryType     int             `json:"delivery_type,omitempty"`  // Required: Yes if delivery_date has been specified. The delivery type for the package.
	DeliveryDate     JSONTime        `json:"delivery_date,omitempty"`  // Required: Yes if delivery type has been specified. The delivery date time for this shipment.
	Signature        int             `json:"signature,omitempty"`      // Required: No. Package must be signed for.
	Return           int             `json:"return,omitempty"`         // Required: No. Return the package if the recipient is not home.
	Insurance        InsuranceStruct `json:"insurance"`                // Required: No. Insurance price for the package.
	LargeFormat      int             `json:"large_format"`             // Required: No. Large format package.
	LabelDescription string          `json:"label_description"`        // Required: No. This description will appear on the shipment label. Note: This will be overridden for return shipment by the following: Retour – 3SMYPAMYPAXXXXXX
	AgeCheck         int             `json:"age_check"`                // Required: No. The recipient must sign for the package and must be at least 18 years old.
}

// Insurance struct
type InsuranceStruct struct {
	Amount   int    `json:"amount"`   // Required: yes. The amount is without decimal separators (in cents).
	Currency string `json:"currency"` // Required: yes. The insurance currency code. Must be one of the following: EUR.
}

// ShipmentRequest to io.Reader
func (s ShipmentRequest) toReader() (io.Reader, error) {
	// convert to json
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	// return the reader
	return bytes.NewReader(b), nil
}

// ShipmentResponse struct
type ShipmentResponse struct {
	Data struct {
		Ids []struct {
			ID                  int    `json:"id"`
			ReferenceIdentifier string `json:"reference_identifier"`
		} `json:"ids"`
	} `json:"data"`
}

// ShipmentRequest struct with data and with multiple shipments
type ShipmentRequest struct {
	Data struct {
		Shipments []ShipmentStruct `json:"shipments"`
	} `json:"data"`
}

// NewClient returns a new MyParcel client.
// The API key is required to use the MyParcel API.
func NewClient(apiKey string) *Client {

	// base64 encode the api key
	apiKey = base64.StdEncoding.EncodeToString([]byte(apiKey))

	// return the client
	return &Client{
		apiBaseURL: "https://api.myparcel.nl",
		httpClient: &http.Client{},
		UserAgent:  "MyParcelGoClient/0.0.1",
		ApiKey:     apiKey,
	}
}

// CreateShipment creates a new shipment.
func (c *Client) CreateShipment(shipment ShipmentStruct) (string, error) {

	// create the request
	request := ShipmentRequest{
		Data: struct {
			Shipments []ShipmentStruct `json:"shipments"`
		}{
			Shipments: []ShipmentStruct{shipment},
		},
	}

	// convert the request to io.Reader
	reader, err := request.toReader()
	if err != nil {
		return "", err
	}

	// create the http request
	req, err := http.NewRequest("POST", c.apiBaseURL+"/shipments", reader)
	if err != nil {
		return "", err
	}

	// set the headers
	req.Header.Set("Content-Type", "application/vnd.shipment+json;version=1.1;charset=utf-8")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.ApiKey)
	req.Header.Set("User-Agent", c.UserAgent)

	// send the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// check the response status code
	if resp.StatusCode != 201 {
		// echo the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		return "", fmt.Errorf("MyParcel API returned status code %d: %s", resp.StatusCode, body)
	}

	return "", nil
}
