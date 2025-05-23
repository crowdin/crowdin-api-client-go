<p align="center">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="https://support.crowdin.com/assets/logos/symbol/png/crowdin-symbol-cWhite.png">
    <source media="(prefers-color-scheme: light)" srcset="https://support.crowdin.com/assets/logos/symbol/png/crowdin-symbol-cDark.png">
    <img width="150" height="150" src="https://support.crowdin.com/assets/logos/symbol/png/crowdin-symbol-cDark.png">
  </picture>
</p>

# Crowdin Go client

The Crowdin Go client is a lightweight interface to the Crowdin API. It provides common services for making API requests.

Our API is a full-featured RESTful API that helps you to integrate localization into your development process. The endpoints that we use allow you to easily make calls to retrieve information and to execute actions needed.

<div align="center">

[**`Developer portal`**](https://developer.crowdin.com/) &nbsp;|&nbsp;
[**`Crowdin API`**](https://developer.crowdin.com/api/v2/) &nbsp;|&nbsp;
[**`Crowdin Enterprise API`**](https://developer.crowdin.com/enterprise/api/v2/)

[![Go Reference](https://pkg.go.dev/badge/github.com/crowdin/crowdin-api-client-go/crowdin.svg)](https://pkg.go.dev/github.com/crowdin/crowdin-api-client-go/crowdin)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/crowdin/crowdin-api-client-go)](https://github.com/crowdin/crowdin-api-client-go/blob/main/go.mod)
[![tests](https://github.com/crowdin/crowdin-api-client-go/actions/workflows/basic.yml/badge.svg)](https://github.com/crowdin/crowdin-api-client-go/actions/workflows/basic.yml)
[![codecov](https://codecov.io/gh/crowdin/crowdin-api-client-go/graph/badge.svg?token=BC055K8EOG)](https://codecov.io/gh/crowdin/crowdin-api-client-go)
[![GitHub contributors](https://img.shields.io/github/contributors/crowdin/crowdin-api-client-go?cacheSeconds=3600)](https://github.com/crowdin/crowdin-api-client-go/graphs/contributors)
[![License](https://img.shields.io/github/license/crowdin/crowdin-api-client-go?cacheSeconds=3600)](https://github.com/crowdin/crowdin-api-client-go/blob/main/LICENSE)

</div>

## Installation

```bash
go get github.com/crowdin/crowdin-api-client-go
```

## Quick Start

Create a new Crowdin client, then use the exposed services to access different parts of the Crowdin API.  
You can generate Personal Access Token in your Crowdin Account Settings.

```go
import "github.com/crowdin/crowdin-api-client-go/crowdin"

client, err := crowdin.NewClient(
    os.Getenv("CROWDIN_ACCESS_TOKEN"),
    crowdin.WithOrganization("organization-name"), // optional for Crowdin Enterprise
)
```

For example, to create a new project:

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/crowdin/crowdin-api-client-go/crowdin"
	"github.com/crowdin/crowdin-api-client-go/crowdin/model"
)

func main() {
    client, err := crowdin.NewClient(os.Getenv("CROWDIN_ACCESS_TOKEN"))
    if err != nil {
        log.Fatalf("Error creating client: %s", err)
    }

    ctx := context.Background()
    request := &model.ProjectsAddRequest{
        Name: "My Project",
		SourceLanguageID: "en",
		TargetLanguageIDs: []string{"uk", "de"},

    }
    project, _, err := client.Projects.Add(ctx, request)
    if err != nil {
        log.Fatalf("Error creating project: %s", err)
    }

    fmt.Printf("Project: %+v\n", project)
}

```

Some API methods have optional parameters that can be passed.  
For example, to list all projects for a specific user:

```go

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/crowdin/crowdin-api-client-go/crowdin"
	"github.com/crowdin/crowdin-api-client-go/crowdin/model"
)

func main() {
    client, err := crowdin.NewClient(os.Getenv("CROWDIN_ACCESS_TOKEN"))
    if err != nil {
        log.Fatalf("Error creating client: %s", err)
    }

    // list all projects for a specific user
    opts := &model.ProjectsListOptions{UserID: 1}
    projects, _, err := client.Projects.List(context.Background(), opts)
    if err != nil {
        log.Fatalf("Error getting projects: %s", err)
    }

    fmt.Printf("Projects: %+v\n", projects)
}
```

### Error Handling

In case of an error, the client returns an error object. This can either be a generic error with an error message and a code, or a validation error that additionally contains validation error codes.

To detect this condition of error, you can use a type assertion:

```go
res, _, err := client.SourceStrings.Add(ctx, 1, nil)
if err != nil {
    if validationErr, ok := err.(*model.ValidationErrorResponse); ok {
        fmt.Printf("Validation error: %v\n", validationErr)
    } else 
        fmt.Printf("Error: %v\n", err)
    }
}
```

### HTTP Request Timeout

To set a timeout for HTTP requests, you can pass a custom HTTP client with a timeout to the client.  
You can also use the context package, you can pass cancellation signals and deadlines to various services of the client to handle a request. 

```go
client, err := crowdin.NewClient(
    os.Getenv("CROWDIN_ACCESS_TOKEN"),
    crowdin.WithHTTPClient(&http.Client{Timeout: 10 * time.Second}),
)
```


## GraphQL API

This library also provides possibility to use [GraphQL API](https://support.crowdin.com/developer/graphql-api/).

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/crowdin/crowdin-api-client-go/crowdin"
)

func main() {
	client, _ := crowdin.NewClient(token)

	// Create a new GraphQL query request.
	query := client.GraphQL.NewRequest(`
		query Demo( $projectLimit: Int!, $withTranslations: Boolean! ) {
			viewer {
				projects( first: $projectLimit ) {
					edges {
						node {
							id
							name
							description

							translations(first: 20, languageId: "en") @include(if: $withTranslations) {
								edges {
									node {
										... on PlainStringTranslation {
											text
										}
									}
								}
							}
						}
					}
					totalCount
				}
			}
		}
	`)

	// Set the variables to be used in the query.
	query.Var("projectLimit", 2)
	query.Var("withTranslations", true)

	// Capture the response in a map.
	// Can also be a struct with the same fields as the response.
	var resp map[string]any

	// Send the query to the server.
	if err := client.GraphQL.Query(context.Background(), query, &resp); err != nil {
		log.Fatalf("graphql: %v", err)
	}

	fmt.Printf("%+v\n", resp)
}
```

## Seeking Assistance

If you find any problems or would like to suggest a feature, please read the [How can I contribute](/CONTRIBUTING.md#how-can-i-contribute) section in our contributing guidelines.

## Contributing

If you want to contribute please read the [Contributing](/CONTRIBUTING.md) guidelines.

## License

<pre>
The Crowdin Go client is licensed under the MIT License.
See the LICENSE file distributed with this work for additional
information regarding copyright ownership.

Except as contained in the LICENSE file, the name(s) of the above copyright
holders shall not be used in advertising or otherwise to promote the sale,
use or other dealings in this Software without prior written authorization.
</pre>
