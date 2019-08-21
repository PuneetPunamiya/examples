// Code generated by goa v2.0.4, DO NOT EDIT.
//
// HTTP request path constructors for the secured_service service.
//
// Command:
// $ goa gen goa.design/examples/security/design -o
// $(GOPATH)/src/goa.design/examples/security

package client

// SigninSecuredServicePath returns the URL path to the secured_service service signin HTTP endpoint.
func SigninSecuredServicePath() string {
	return "/signin"
}

// SecureSecuredServicePath returns the URL path to the secured_service service secure HTTP endpoint.
func SecureSecuredServicePath() string {
	return "/secure"
}

// DoublySecureSecuredServicePath returns the URL path to the secured_service service doubly_secure HTTP endpoint.
func DoublySecureSecuredServicePath() string {
	return "/secure"
}

// AlsoDoublySecureSecuredServicePath returns the URL path to the secured_service service also_doubly_secure HTTP endpoint.
func AlsoDoublySecureSecuredServicePath() string {
	return "/secure"
}
