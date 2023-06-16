// Copyright 2018-2019 opcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"go/format"
	"io"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/gopcua/opcua/id"
	"github.com/gopcua/opcua/server/refs"
	"github.com/gopcua/opcua/ua"
)

func main() {
	in := flag.String("in", "../schema/Opc.Ua.PredefinedNodes.xml", "XML of predefined nodes")
	out := flag.String("out", "nodes_gen.go", "generated file")
	pkg := flag.String("pkg", "server", "package name")
	flag.Parse()

	log.SetFlags(0)

	f, err := os.Open(*in)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var nodes []*Node
	d := xml.NewDecoder(f)
	for {
		tok, err := d.Token()
		if tok == nil || err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		switch ty := tok.(type) {
		case xml.StartElement:
			if ty.Name.Space != "http://opcfoundation.org/UA/" {
				continue
			}
			n := new(Node)
			if err := d.DecodeElement(n, &ty); err != nil {
				log.Fatal(err)
			}
			n.Type = ty.Name.Local
			// fmt.Println(n.NodeID.Identifier.String())
			// fmt.Printf("%#v\n", n)
			nodes = append(nodes, n)
		}
	}

	m := map[string]*Node{}
	for _, n := range nodes {
		// trim the _xx suffix from the node class
		n.NodeClass = n.NodeClass[0:strings.Index(n.NodeClass, "_")] // DataType_32 -> DataType

		m[n.NodeID.Identifier.String()] = n
	}

	// create HasSubtype refs
	for _, n := range m {
		sid := n.SuperTypeID.Identifier.String()
		if sid == "" {
			continue
		}
		ref := refs.HasSubtype(&ua.ExpandedNodeID{NodeID: n.NodeID.Identifier})
		m[sid].Refs = append(m[sid].Refs, ref)
	}

	data := map[string]interface{}{
		"Package": *pkg,
		"Nodes":   nodes,
	}

	var b bytes.Buffer
	if err := tmpl.Execute(&b, data); err != nil {
		log.Fatal(err)
	}
	src, err := format.Source(b.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile(*out, src, 0644); err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %s/%s", *pkg, *out)
}

type Node struct {
	Type      string `xml:"-"`
	Xmlns     string `xml:",attr"`
	NodeClass string
	NodeID    struct {
		Identifier *ua.NodeID
	} `xml:"NodeId"`
	BrowseName struct {
		NamespaceIndex string
		Name           string
	}
	ReferenceTypeID struct {
		Identifier *ua.NodeID
	} `xml:"ReferenceTypeId"`
	TypeDefinitionID struct {
		Identifier *ua.NodeID
	} `xml:"TypeDefinitionId"`
	SuperTypeID struct {
		Identifier *ua.NodeID
	} `xml:"SuperTypeId"`
	InverseName *struct {
		Locale string
		Text   string
	}
	IsAbstract bool

	Refs []*ua.ReferenceDescription
}

var funcs = template.FuncMap{
	"idname": id.Name,
}

var tmpl = template.Must(template.New("").Funcs(funcs).Parse(`// Generated code. DO NOT EDIT
 // Copyright 2018-2023 opcua authors. All rights reserved.
 // Use of this source code is governed by a MIT-style license that can be
 // found in the LICENSE file.
 package {{.Package}}

 import (
	"github.com/gopcua/opcua/id"
	"github.com/gopcua/opcua/ua"
 )

 func PredefinedNodes() []Node{
 	return []Node{
 {{- range .Nodes }}
 		&node{
 			{{- with .NodeID.Identifier }}
 			id: ua.NewNumericNodeID({{.Namespace}}, {{.IntID}}),
 			{{- end}}
 			attr: map[ua.AttributeID]*ua.Variant{
 				ua.AttributeIDNodeClass: ua.MustVariant(ua.NodeClass{{.NodeClass}}),
 				ua.AttributeIDBrowseName: ua.MustVariant(&ua.QualifiedName{Name:"{{.BrowseName.Name}}"}),
 				ua.AttributeIDDisplayName: ua.MustVariant(&ua.LocalizedText{Text:"{{.BrowseName.Name}}"}),
 				{{- with .InverseName }}
 				ua.AttributeIDInverseName: ua.MustVariant(&ua.LocalizedText{Locale:"{{.Locale}}", Text:"{{.Text}}"}),
 				{{- end}}
 			},
			{{- if .Refs }}
			refs: []*ua.ReferenceDescription{
			{{- range .Refs }}
				{
					ReferenceTypeID: ua.NewNumericNodeID(0, id.{{idname .ReferenceTypeID.IntID}}),
					TypeDefinition: ua.NewNumericExpandedNodeID(0, {{.TypeDefinition.NodeID.IntID}}),
					IsForward: {{.IsForward}},
				},
 			{{- end }}
			},
			{{- end }}
 		},
 {{- end }}
 	}
 }
 `))
