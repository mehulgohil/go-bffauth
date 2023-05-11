# go-bffauth
BFF Auth Pattern implemented in Go 

## Introduction

In the realm of modern web development, creating secure and efficient user authentication flows is paramount.
With the rise of microservices and distributed architectures,
implementing a robust authentication strategy becomes even more crucial.

The Backend for Frontend (BFF)
authentication pattern is a design approach
that focuses
on providing a dedicated backend
to handle and address all authentication requirements and challenges of the frontend application. 
(SPA)

## History

SPAs typically use token-based authentication, such as JSON Web Tokens (JWTs), to authenticate and authorize users.
However, managing tokens securely in SPA can be challenging.
Storing tokens in client-side storage (e.g., local storage or cookies)
can expose them to potential attacks like cross-site request forgery
(CSRF).
Developers must implement stringent security measures to protect tokens and prevent unauthorized access or misuse

The BFF pattern solves this problem by introducing an intermediate layer—the Backend for Frontend.
This layer acts as a proxy between the front-end client and the main backend services,
handling authentication-related concerns and providing a specialized authentication interface.

## Flow Diagram

The following diagram illustrates how this pattern works in detail:

![sequenceDiagram.png](.attachments%2FsequenceDiagram.png)

1. When the frontend needs to authenticate the user, it calls an API endpoint (/login) on the BFF to start the login handshake.
2. The BFF uses OAuth2 Authorization Code Flow to connect with Auth0 to authenticate and authorize the user and get the id and access tokens.
3. The backend stores the user's tokens in a cache.
4. An encrypted cookie is issued for the frontend representing the user authentication session.
5. When the frontend needs to call an external API, it passes the encrypted cookie to the BFF together with the URL and data to invoke the API.
6. The BFF retrieves the access token from the cache and makes a call to the backend API including that token on the authorization header.
7. When the external API returns a response to the BFF, this one forwards that response back to the frontend.