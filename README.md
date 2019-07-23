![Hex.pm](https://img.shields.io/hexpm/l/plug.svg)
# APIs and Clients for PA and IPAM operator
Scheme, typing, client packages and Kubernetes-like API objects for PA and IPAM operator.

## Usage
```go
import (
	blendedv1 "github.com/inwinstack/blended/apis/inwinstack/v1"
	blendedset "github.com/inwinstack/blended/generated/clientset/versioned"
)
```
Create a new client, then use the exposed services to access different parts of the Kubernetes custom API.
