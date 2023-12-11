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
	"github.com/gopcua/opcua/ua"
)

var (
	endpoint = flag.String("endpoint", "opc.tcp://127.0.0.1:4840", "OPC UA Endpoint URL")
	certfile = flag.String("cert", "cert.pem", "Path to certificate file")
	keyfile  = flag.String("key", "key.pem", "Path to PEM Private Key file")
	gencert  = flag.Bool("gen-cert", false, "Generate a new certificate")
)

func main() {
	flag.BoolVar(&debug.Enable, "debug", false, "enable debug logging")
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

	s := server.New(*endpoint, opts...)

	mrw := NewMapNamespace("MyTestNamespace")

	num := 42

	mrw.Data["Tag1"] = 123.4
	mrw.Data["Tag2"] = 42
	mrw.Data["Tag3.Tag4"] = "some string"
	mrw.Data["Tag5"] = true
	mrw.Data["Tag6"] = time.Now()

	// some background process updating the map
	go func() {
		for {
			num++
			mrw.Mu.Lock()
			mrw.Data["Tag2"] = num
			mrw.Mu.Unlock()
			time.Sleep(time.Second)
		}
	}()

	// register our custom read handler.
	s.RegisterHandler(id.ReadRequest_Encoding_DefaultBinary, mrw.CustomRead)
	s.RegisterHandler(id.WriteRequest_Encoding_DefaultBinary, mrw.CustomWrite)
	//s.RegisterHandler(id.BrowseRequest_Encoding_DefaultBinary, mrw.CustomBrowse)

	s.RegisterHandler(id.CreateSubscriptionRequest_Encoding_DefaultBinary, mrw.CreateSubscription)
	s.RegisterHandler(id.PublishRequest_Encoding_DefaultBinary, mrw.Publish)
	s.RegisterHandler(id.CreateMonitoredItemsRequest_Encoding_DefaultBinary, mrw.CreateMonitoredItems)

	mrw_id := s.AddNamespace(&mrw)
	log.Printf("map namespace added at index %d", mrw_id)

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

// when first attempting to read a tag it tries to create a subscription request.
//
//    0000   f0 d4 15 92 21 f1 be 80 ad fa 06 f6 08 00 45 00   ....!.........E.
//    0010   00 88 9d 80 40 00 40 06 16 b6 c0 a8 02 a7 c0 a8   ....@.@.........
//    0020   02 42 a2 74 12 e8 eb e2 00 65 9f 69 33 be 80 18   .B.t.....e.i3...
//    0030   01 f5 55 0c 00 00 01 01 08 0a ed ff 66 61 14 ce   ..U.........fa..
//    0040   09 00 4d 53 47 46 54 00 00 00 75 0e 1f 24 0d 31   ..MSGFT...u..$.1
//    0050   f8 39 0e 00 00 00 0e 00 00 00 01 00 13 03 02 00   .9..............
//    0060   00 f2 d0 a9 5f 10 60 f7 b5 7f 22 da 01 0c 00 00   ...._.`...".....
//    0070   00 00 00 00 00 ff ff ff ff 60 ea 00 00 00 00 00   .........`......
//    0080   00 00 00 00 00 40 7f 40 b0 04 00 00 78 00 00 00   .....@.@....x...
//    0090   ff ff 00 00 01 00                                 ......
