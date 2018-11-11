![Hex.pm](https://img.shields.io/hexpm/l/plug.svg)
# APIs and Clients for PA and IPAM operator
Scheme, typing, client packages and Kubernetes-like API objects for PA and IPAM operator.

## Usage
```go
import (
	inwinv1 "github.com/inwinstack/blended/apis/inwinstack/v1"
	inwinclientset "github.com/inwinstack/blended/client/clientset/versioned/typed/inwinstack/v1"
)
```
Create a new client, then use the exposed services to access different parts of the Kubernetes custom API.
