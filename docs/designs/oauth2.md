# OAuth2 Service Design

## Flows

### OIDC - Authorization Code Flow

**Todo - Future Work**

### OIDC - Implicit Flow

**Todo - Future Work**

### OIDC - Hybrid Flow

**Not supported**

### OAuth2 - Authorization Code Grant

**Not supported**

See "OIDC - Authorization Code Flow"

### OAuth2 - Implicit Grant

**Not supported**

See "OIDC - Implicit Flow"

### OAuth2 - Resource Owner Password Credentials Grant

**Not supported**

### OAuth2 - Client Credentials Grant

Client makes a call to the Token Endpoint, providing:

* grant_type: client_credentials

Client receives a response containing:

* access_token
* token_type
* expires_in

## Supported Endpoints

### GET /oauth2/authorize

#### OIDC - Authorization Code Flow

**Not defined yet.**

#### OIDC - Implicit Flow

**Not defined yet.**

### POST /oauth2/token

Authorization header *must* be a HTTP Basic Authorization containing the Client ID and Client Secret.

#### OIDC - Authorization Code Flow

**Not defined yet.**

#### OAuth2 - Client Credentials Grant

Request Payload has:

* grant_type: client_credentials

Server will resolve the Client Details from the Authorization header. These then resolve to a User that owns the Client. An Access Token is then generated for this User and returned to the HTTP Client. No Refresh Token will be generated.
