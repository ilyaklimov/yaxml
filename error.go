package yaxml

import ()

type Error struct {
    Code 		int 	`xml:"code,attr" json:"code"`
    Message 	string 	`xml:",chardata" json:"message"`
}