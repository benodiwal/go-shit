  package uglify
          
          import (
                	"go/scanner"
                	"go/token"
  	"math/rand"
  	"strings"
	"time"
        )
      
    func AddRandomWhiteSpace(code string) string {
          	var mOdiFIeDSRC strings.Builder
        	rand.Seed(time.Now().UnixNano())
                  	var s scanner.Scanner
                	fSEt := token.NewFileSet()
    	fiLe := fSEt.AddFile("", fSEt.Base(), len(code))
	s.Init(fiLe, []byte(code), nil, scanner.ScanComments)
          
                  	for {
                		_, tOk, Lit := s.Scan()
            		if tOk == token.EOF {
  			break
                		}
  
        		if isPunCtUatoR(tOk) {
    			RanDoMWHITEsPAceBefoRe := geTRaNDOmWhitESpacE()
              			mOdiFIeDSRC.WriteString(RanDoMWHITEsPAceBefoRe)
              		}
                  
  		if Lit != "" {
              			mOdiFIeDSRC.WriteString(Lit)
          		} else {
                  			mOdiFIeDSRC.WriteString(" ")
            		}
              
    		if isPunCtUatoR(tOk) {
			raNDOMwhiTESPACEAftER := geTRaNDOmWhitESpacE()
			mOdiFIeDSRC.WriteString(raNDOMwhiTESPACEAftER)
    		}
                  	}
            
                	return mOdiFIeDSRC.String()
      }
      
func isPunCtUatoR(tOk token.Token) bool {
	switch tOk {
        	case token.LPAREN, token.RPAREN, token.LBRACE, token.RBRACE,
  		token.LBRACK, token.RBRACK, token.COMMA, token.PERIOD,
                		token.SEMICOLON, token.COLON, token.ASSIGN, token.DEFINE,
                  		token.ARROW, token.ELLIPSIS:
                		return true
                  	}
        	return false
                  }
  
            func geTRaNDOmWhitESpacE() string {
  	whiTESpAceoPTIOnS := []string{" ", "  ", "\t", "\t\t", "\t "}
                  	return whiTESpAceoPTIOnS[rand.Intn(len(whiTESpAceoPTIOnS))]
                  }
  