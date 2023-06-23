package server

import (
	"crypto/rand"
	"log"
	"time"

	"github.com/gopcua/opcua/debug"
	"github.com/gopcua/opcua/ua"
	"github.com/gopcua/opcua/uasc"
)

const (
	sessionTimeoutMin     = 100            // 100ms
	sessionTimeoutMax     = 30 * 60 * 1000 // 30 minutes
	sessionTimeoutDefault = 60 * 1000      // 60s

	sessionNonceLength = 32
)

// SessionService implements the Session Service Set.
//
// https://reference.opcfoundation.org/Core/Part4/v105/docs/5.6
type SessionService struct {
	srv *Server
}

// https://reference.opcfoundation.org/Core/Part4/v105/docs/5.6.2
func (s *SessionService) CreateSessionRequest(sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.CreateSessionRequest)
	if !ok {
		debug.Printf("handleCreateSessionRequest: Expected *ua.CreateSessionRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	// New session
	sess := s.srv.sb.NewSession()

	// Ensure session timeout is reasonable
	sess.cfg.sessionTimeout = time.Duration(req.RequestedSessionTimeout) * time.Millisecond
	if sess.cfg.sessionTimeout > sessionTimeoutMax || sess.cfg.sessionTimeout < sessionTimeoutMin {
		sess.cfg.sessionTimeout = sessionTimeoutDefault
	}

	nonce := make([]byte, sessionNonceLength)
	if _, err := rand.Read(nonce); err != nil {
		log.Printf("error creating session nonce")
		return nil, ua.StatusBadInternalError
	}
	sess.serverNonce = nonce
	sess.remoteCertificate = req.ClientCertificate

	sig, alg, err := sc.NewSessionSignature(req.ClientCertificate, req.ClientNonce)
	if err != nil {
		log.Printf("error creating session signature")
		return nil, ua.StatusBadInternalError
	}

	response := &ua.CreateSessionResponse{
		ResponseHeader:        responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
		SessionID:             sess.ID,
		AuthenticationToken:   sess.AuthTokenID,
		RevisedSessionTimeout: float64(sess.cfg.sessionTimeout / time.Millisecond),
		MaxRequestMessageSize: 0, // Not used
		ServerSignature: &ua.SignatureData{
			Signature: sig,
			Algorithm: alg,
		},
		ServerCertificate: s.srv.cfg.certificate,
		ServerNonce:       nonce,
		ServerEndpoints:   s.srv.endpoints,
	}

	return response, nil
}

// https://reference.opcfoundation.org/Core/Part4/v105/docs/5.6.3
func (s *SessionService) ActivateSessionRequest(sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.ActivateSessionRequest)
	if !ok {
		debug.Printf("handleActivateSessionRequest: Expected *ua.ActivateSessionRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	sess := s.srv.sb.Session(req.RequestHeader.AuthenticationToken)
	if sess == nil {
		return nil, ua.StatusBadSessionIDInvalid
	}

	err := sc.VerifySessionSignature(sess.remoteCertificate, sess.serverNonce, req.ClientSignature.Signature)
	if err != nil {
		debug.Printf("error verifying session signature with nonce: %s", err)
		return nil, ua.StatusBadSecurityChecksFailed
	}

	nonce := make([]byte, sessionNonceLength)
	if _, err := rand.Read(nonce); err != nil {
		log.Printf("error creating session nonce")
		return nil, ua.StatusBadInternalError
	}
	sess.serverNonce = nonce

	response := &ua.ActivateSessionResponse{
		ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
		ServerNonce:    nonce,
		// Results:         []ua.StatusCode{},
		// DiagnosticInfos: []*ua.DiagnosticInfo{},
	}

	return response, nil
}

// https://reference.opcfoundation.org/Core/Part4/v105/docs/5.6.4
func (s *SessionService) CloseSessionRequest(sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.CloseSessionRequest)
	if !ok {
		debug.Printf("handleCloseSessionRequest: Expected *ua.CloseSessionRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	err := s.srv.sb.Close(req.RequestHeader.AuthenticationToken)
	if err != nil {
		return nil, ua.StatusBadSessionIDInvalid
	}

	//TODO: deal with 'delete subscriptions' field in request
	response := &ua.CloseSessionResponse{
		ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	}

	return response, nil
}

// https://reference.opcfoundation.org/Core/Part4/v105/docs/5.6.5
func (s *SessionService) CancelRequest(sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.CancelRequest)
	if !ok {
		debug.Printf("handleCancelRequest: Expected *ua.CancelRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusBadServiceUnsupported)}
	// response := &ua.CancelResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}
