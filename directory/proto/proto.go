// Package proto contains the definitions shared between RPC directory server and client,
// one pair for each remote call.
// TODO: Maybe move to gprc?
package proto

import "upspin.googlesource.com/upspin.git/upspin"

type LookupRequest struct {
	Name upspin.PathName
}

type LookupResponse struct {
	Entry *upspin.DirEntry
}

type PutRequest struct {
	Entry *upspin.DirEntry
}

type PutResponse struct {
}

type MakeDirectoryRequest struct {
	Name upspin.PathName
}

type MakeDirectoryResponse struct {
	Location upspin.Location
}

type GlobRequest struct {
	Pattern string
}

type GlobResponse struct {
	Entries []*upspin.DirEntry
}

type DeleteRequest struct {
	Name upspin.PathName
}

type DeleteResponse struct {
}

type WhichAccessRequest struct {
	Name upspin.PathName
}

type WhichAccessResponse struct {
	Name upspin.PathName
}