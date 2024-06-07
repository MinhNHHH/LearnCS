package functv

import (
	"bytes"
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

type PrettyPrinter struct {
	indent int
	buffer bytes.Buffer
}

func NewPrettyPrinter() *PrettyPrinter {
	return &PrettyPrinter{}
}

func (pp *PrettyPrinter) Print(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		pp.startElement(n)
	case html.TextNode:
		pp.printText(n)
	case html.CommentNode:
		pp.printComment(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		pp.Print(c)
	}
	if n.Type == html.ElementNode {
		pp.endElement(n)
	}
}

func (pp *PrettyPrinter) printComment(n *html.Node) {
	pp.buffer.WriteString(strings.Repeat(" ", pp.indent))
	pp.buffer.WriteString(fmt.Sprintf("<!-- %s -->\n", n.Data))
}

func (pp *PrettyPrinter) printText(n *html.Node) {
	if strings.TrimSpace(n.Data) != "" {
		pp.buffer.WriteString(strings.Repeat(" ", pp.indent))
		pp.buffer.WriteString(n.Data + "\n")
	}
}

func (pp *PrettyPrinter) startElement(n *html.Node) {
	pp.buffer.WriteString(strings.Repeat(" ", pp.indent))
	pp.buffer.WriteString("<" + n.Data + " ")
	for _, node := range n.Attr {
		pp.buffer.WriteString(fmt.Sprintf("%s='%s'", node.Key, node.Val))
	}
	if n.FirstChild == nil {
		pp.buffer.WriteString("/>\n")
	} else {
		pp.buffer.WriteString(">\n")
		pp.indent += 2
	}
}

func (pp *PrettyPrinter) endElement(n *html.Node) {
	if n.FirstChild != nil {
		pp.indent -= 2
		pp.buffer.WriteString(strings.Repeat(" ", pp.indent))
		pp.buffer.WriteString(fmt.Sprintf("</%s>\n", n.Data))
	}
}
func (pp *PrettyPrinter) String() string {
	return pp.buffer.String()
}
