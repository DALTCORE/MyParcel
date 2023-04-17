# MyParcel Go Package [WIP, not working yet!]
MyParcel is a Go package that provides functionality for sending and tracking letters and parcels. It interacts with the MyParcel API to perform operations such as creating shipments and tracking packages.

## Installation
To install the MyParcel Go package, use the following command:

```
go get github.com/daltcore/myparcel
```

## Usage
To use the MyParcel Go package, import it into your Go project:

```
import "github.com/daltcore/myparcel"
```

Then, create a new MyParcel client:

```
client := myparcel.NewClient("<your API key>")
````

You can then use the client to perform various operations, such as creating a shipment:

```
shipment := myparcel.Shipment{
    Carrier:    "postnl",
    Name:       "John Doe",
    Street:     "Main Street 1",
    PostalCode: "1000 AA",
    City:       "Amsterdam",
    Country:    "NL",
}


result, err := client.CreateShipment(shipment)
if err != nil {
    // Handle error
}

fmt.Println(result.ID)
```
You can also use the client to track a package:

```
result, err := client.TrackPackage("<your shipment ID>")
if err != nil {
    // Handle error
}

fmt.Println(result.Status)
```

For more information about the available operations and data structures, see the GoDoc documentation.

## Features
The MyParcel Go package includes the following features:

- Create shipments
- Track packages
- Retrieve shipment labels
- Retrieve shipment status updates

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