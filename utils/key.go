package utils

func Encode(b []byte ,key uint8)[]byte  {
	for k,v:=range b{
		u := key+ v
		b[k]=u
	}

	return b
}

func Decode(b []byte,key uint8)[]byte{
	for k,v:=range b{
		u:=v-key
		b[k]=u
	}
	return b
}
