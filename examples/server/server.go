// Copyright 2018-2020 opcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package main

import (
	"context"
	"crypto/rsa"
	"crypto/tls"
	"flag"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gopcua/opcua/debug"
	"github.com/gopcua/opcua/id"
	"github.com/gopcua/opcua/server"
	"github.com/gopcua/opcua/server/attrs"
	"github.com/gopcua/opcua/ua"
)

var (
	endpoint = flag.String("endpoint", "opc.tcp://localhost:4840", "OPC UA Endpoint URL")
	certfile = flag.String("cert", "cert.pem", "Path to certificate file")
	keyfile  = flag.String("key", "key.pem", "Path to PEM Private Key file")
	gencert  = flag.Bool("gen-cert", false, "Generate a new certificate")
)

func main() {
	flag.BoolVar(&debug.Enable, "debug", true, "enable debug logging")
	flag.Parse()
	log.SetFlags(0)

	var opts []server.Option

	opts = append(opts,
		server.EnableSecurity("None", ua.MessageSecurityModeNone),
		server.EnableSecurity("Basic128Rsa15", ua.MessageSecurityModeSign),
		server.EnableSecurity("Basic128Rsa15", ua.MessageSecurityModeSignAndEncrypt),
		server.EnableSecurity("Basic256", ua.MessageSecurityModeSign),
		server.EnableSecurity("Basic256", ua.MessageSecurityModeSignAndEncrypt),
		server.EnableSecurity("Basic256Sha256", ua.MessageSecurityModeSignAndEncrypt),
		server.EnableSecurity("Basic256Sha256", ua.MessageSecurityModeSign),
		server.EnableSecurity("Aes128_Sha256_RsaOaep", ua.MessageSecurityModeSign),
		server.EnableSecurity("Aes128_Sha256_RsaOaep", ua.MessageSecurityModeSignAndEncrypt),
		server.EnableSecurity("Aes256_Sha256_RsaPss", ua.MessageSecurityModeSign),
		server.EnableSecurity("Aes256_Sha256_RsaPss", ua.MessageSecurityModeSignAndEncrypt),
	)

	opts = append(opts,
		server.EnableAuthMode(ua.UserTokenTypeAnonymous),
		server.EnableAuthMode(ua.UserTokenTypeUserName),
		server.EnableAuthMode(ua.UserTokenTypeCertificate),
		//		server.EnableAuthWithoutEncryption(), // Dangerous and not recommended, shown for illustration only
	)

	var cert []byte
	if *gencert || (*certfile != "" && *keyfile != "") {
		debug.Printf("Loading cert/key from %s/%s", *certfile, *keyfile)
		c, err := tls.LoadX509KeyPair(*certfile, *keyfile)
		if err != nil {
			log.Printf("Failed to load certificate: %s", err)
		} else {
			pk, ok := c.PrivateKey.(*rsa.PrivateKey)
			if !ok {
				log.Fatalf("Invalid private key")
			}
			cert = c.Certificate[0]
			opts = append(opts, server.PrivateKey(pk), server.Certificate(cert))
		}
	}

	n1 := server.NewNode(
		ua.NewStringNodeID(0, "MyValueNode"),
		map[ua.AttributeID]*ua.Variant{
			ua.AttributeIDBrowseName: ua.MustVariant(attrs.BrowseName("MyValueNode")),
			ua.AttributeIDValue:      ua.MustVariant(123.45),
		},
		nil,
		func() *ua.Variant { return ua.MustVariant("MyValue") },
	)

	s := server.New(*endpoint, opts...)
	objects := s.AddressSpace().Objects()
	objects.AddObject(n1)
	objects.AddObject(server.NewNode(
		ua.NewNumericNodeID(0, 2259),
		map[ua.AttributeID]*ua.Variant{
			ua.AttributeIDBrowseName: ua.MustVariant(attrs.BrowseName("Dunno")),
			ua.AttributeIDValue:      ua.MustVariant(123.45),
		},
		nil,
		func() *ua.Variant { return ua.MustVariant("Dunno") },
	))

	s.AddressSpace().AddNode(server.NewNode(
		ua.NewStringNodeID(0, "TestNodeName"),
		map[ua.AttributeID]*ua.Variant{
			ua.AttributeIDBrowseName: ua.MustVariant(attrs.BrowseName("TestNodeName")),
		},
		nil,
		func() *ua.Variant { return ua.MustVariant("TestValue") },
	))

	mrw := MapReadWriter{}
	mrw.Data = make(map[string]any)
	mrw.Data["s=Tag1"] = 123.4
	mrw.Data["s=Tag2"] = 42
	mrw.Data["s=Tag3.Tag4"] = "some string"
	mrw.Data["s=Tag5"] = true
	mrw.Data["s=Tag6"] = time.Now()

	// register our custom read handler.
	s.RegisterHandler(id.ReadRequest_Encoding_DefaultBinary, mrw.CustomRead)
	s.RegisterHandler(id.WriteRequest_Encoding_DefaultBinary, mrw.CustomWrite)

	if err := s.Start(context.Background()); err != nil {
		log.Fatalf("Error starting server, exiting: %s", err)
	}
	defer s.Close()

	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, os.Interrupt)
	defer signal.Stop(sigch)

	log.Printf("Press CTRL-C to exit")
	<-sigch
	log.Printf("Shutting down the server...")
}
