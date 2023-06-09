# MyParcel Go Package [WIP, partially working]
MyParcel is a Go package that provides functionality for sending and tracking letters and parcels. It interacts with the MyParcel API to perform operations such as creating shipments and tracking packages.

## Installation
To install the MyParcel Go package, use the following command:

```bash
go get github.com/daltcore/myparcel
```

## Usage
To use the MyParcel Go package, import it into your Go project:

```go
import "github.com/daltcore/myparcel"
```

Then, create a new MyParcel client:

```go
client := myparcel.NewClient("<your API key>")
````

You can then use the client to perform various operations, such as creating a shipment:

```go
shipment := myparcel.ShipmentStruct{
    ReferenceIdentifier: "123456789",
    Recipient: myparcel.RecipientStruct{
        Street:     "Testlane",
        Number:     "1",
        PostalCode: "1234 AB",
        City:       "Amsterdam",
        Cc:         "NL",
        Person:     "D. Duck",
    },
    Options: myparcel.OptionsStruct{
        PackageType: myparcel.ParcelPackage,
    },
    Carrier: myparcel.CarrierPostnl,
}

id, err := client.CreateShipment(shipment)
if err != nil {
    panic(err)
}

fmt.Println(id)
```
You can also use the client to track a package:

```go
shipmentResponse, err := client.GetShipment(12345)
	if err != nil {
		panic(err)
	}

fmt.Printf("%+v\n", shipmentResponse)
```

For more information about the available operations and data structures, see the GoDoc documentation.

## Features
The MyParcel Go package includes the following features:

- Create shipments
- Get shipments
- Track packages

## Contributing
If you would like to contribute to the MyParcel Go package, please follow these steps:

- Fork this repository
- Create a new branch for your feature: `git checkout -b my-new-feature`
- Make your changes and commit them: `git commit -am 'Add some feature'`
- Push to the branch: `git push origin my-new-feature`
- Create a new pull request

Please ensure that your code follows the project's coding style and that all tests pass before submitting a pull request.

## License
This project is licensed under the MIT License. See the LICENSE file for details.