# Users API Design

## Routes

* GET /api/users? - Search user records
* POST /api/users - Create a new user record
* GET /api/users/:id - Get a single user record
* PUT /api/users/:id - Update a single user record
* DELETE /api/users/:id - Delete a single user record

## Payload Structure

### Properties

* ID - String. Reading resources only.
* Version - String. Reading resources only.
* Created - Date/Time. Reading resources only.
* Updated - Date/Time. Reading resources only.
* Name - String
* Email - String. Optional.
* Logins - Array
  * Provider - String
  * Provider ID - String
  * Display Name - String

### POJO + JSON Schema

#### Retrieve
##### Headers
* Content-Type: application/json; charset=UTF-8
* ETag: "F4FA9B81-5E6A-4300-ACEA-B878D9A325E3"
* Link: </schemas/user.json>; rel="describedBy"; type="application/schema+json"

##### Body
```json
{
    "id": "FC14606D-77D0-42EE-9C65-793B8BFF599A",
    "created": "2018-11-09T14:12:00Z",
    "updated": "2018-11-09T14:12:00Z",
    "name": "Graham",
    "email": "graham@grahamcox.co.uk",
    "logins": [
        {
            "provider": "google",
            "providerId": "123454321",
            "displayName": "graham@grahamcox.co.uk"
        }
    ]
}
```

### HAL

#### Retrieve
##### Headers
* Content-Type: application/hal+json; charset=UTF-8
* ETag: "F4FA9B81-5E6A-4300-ACEA-B878D9A325E3"

##### Body
```json
{
    "_links": {
        "self": {
            "href": "/api/users/FC14606D-77D0-42EE-9C65-793B8BFF599A"
        },
        "worlds:clients": {
            "href": "/api/clients?user_id=FC14606D-77D0-42EE-9C65-793B8BFF599A"
        },
        "curies": [
            {
                "name": "worlds",
                "href": "/rels/worlds/{rel}",
                "templated": true
            }
        ]
    },
    "created": "2018-11-09T14:12:00Z",
    "updated": "2018-11-09T14:12:00Z",
    "name": "Graham",
    "email": "graham@grahamcox.co.uk",
    "logins": [
        {
            "provider": "google",
            "providerId": "123454321",
            "displayName": "graham@grahamcox.co.uk"
        }
    ]
}
```

### Siren
#### Retrieve
##### Headers
* Content-Type: application/vnd.siren+json; charset=UTF-8
* ETag: "F4FA9B81-5E6A-4300-ACEA-B878D9A325E3"

##### Payload
```json
{
    "class": [
        "user"
    ],
    "properties": {
        "created": "2018-11-09T14:12:00Z",
        "updated": "2018-11-09T14:12:00Z",
        "name": "Graham",
        "email": "graham@grahamcox.co.uk",
        "logins": [
            {
                "provider": "google",
                "providerId": "123454321",
                "displayName": "graham@grahamcox.co.uk"
            }
        ]
    },
    "actions": [
        {
            "name": "update",
            "title": "Update User",
            "method": "POST",
            "href": "/api/users/FC14606D-77D0-42EE-9C65-793B8BFF599A",
            "type": "application/json",
            "fields": [
                {"name": "name", "type": "text"},
                {"name": "email", "type": "text"},
                {"name": "logins", "type": "?????"} // TODO
            ]
        }
    ],
    "links": [
        {
            "rel": "self",
            "href": "/api/users/FC14606D-77D0-42EE-9C65-793B8BFF599A"
        }
    ],
    "entities": [
        {
            "class": [
                "client",
                "collection"
            ],
            "rel": "/rels/users/clients",
            "href": "/api/clients?user_id=FC14606D-77D0-42EE-9C65-793B8BFF599A"
        }
    ]
}
```

### JSON-API
#### Retrieve
##### Headers
* Content-Type: application/vnd.api+json; charset=UTF-8
* ETag: "F4FA9B81-5E6A-4300-ACEA-B878D9A325E3"

##### Payload
```json
{
    "links": {
        "self": "/api/users/FC14606D-77D0-42EE-9C65-793B8BFF599A"
    },
    "data": {
        "type": "user",
        "id": "FC14606D-77D0-42EE-9C65-793B8BFF599A",
        "attributes": {
            "name": "Graham",
            "email": "graham@grahamcox.co.uk",
            "logins": [
                {
                    "provider": "google",
                    "providerId": "123454321",
                    "displayName": "graham@grahamcox.co.uk"
                }
            ]
        },
        "meta": {
            "created": "2018-11-09T14:12:00Z",
            "updated": "2018-11-09T14:12:00Z"
        },
        "relationships": {
            "clients": {
                "links": {
                    "self": "/api/users/FC14606D-77D0-42EE-9C65-793B8BFF599A/relationships/clients",
                    "related": "/api/clients?user_id=FC14606D-77D0-42EE-9C65-793B8BFF599A"
                }
            }
        }
    }
}
```

### JSON-LD Hydra
#### Retrieve
##### Headers
* Link: </hydra/vocab>; rel="http://www.w3.org/ns/hydra/core#apiDocumentation"
* Content-Type: application/vnd.api+json; charset=UTF-8
* ETag: "F4FA9B81-5E6A-4300-ACEA-B878D9A325E3"

##### Payload
```json
{
    "@context": "https://api.myjson.com/bins/8dx8u",
    "@id": "http://myjson.com/api/users/FC14606D-77D0-42EE-9C65-793B8BFF599A",
    "@type": [
        "Person",
        "User"
    ],
    "created": "2018-11-09T14:12:00Z",
    "updated": "2018-11-09T14:12:00Z",
    "name": "Graham",
    "email": "graham@grahamcox.co.uk",
    "logins": [
        {
            "provider": "google",
            "providerId": "123454321",
            "providerName": "graham@grahamcox.co.uk"
        }
    ]
}
```

#### Context
##### Payload
```json
{
      "@context": {
        "Person": "http://schema.org/Person",
        "User": "/hydra/types/User",
        "created": "http://schema.org/dateCreated",
        "updated": "http://schema.org/dateModified",
        "name": "http://schema.org/name",
        "email": "http://schema.org/email",
        "logins": {
            "@id": "/hydra/logins",
            "@type": "/hydra/types/Login",
            "@container": "@set"
        },
        "provider": "//hydra/provider",
        "providerId": "//hydra/providerId",
        "providerName": "//hydra/providerName"
    }
}
```
