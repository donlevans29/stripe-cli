// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package fixtures

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// FS statically implements the virtual filesystem provided to vfsgen.
var FS = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		"/charge.captured.json": &vfsgen۰CompressedFileInfo{
			name:             "charge.captured.json",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 399,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8f\xc1\x6a\xc3\x30\x10\x44\xef\xfa\x8a\x65\xe9\xd1\x60\xe7\xaa\x5f\x29\xc5\x2c\xf2\xa6\x36\x8d\x2c\xa1\x5d\x99\x96\xa0\x7f\x2f\x92\x93\xd4\xb4\x34\x27\x49\x33\xf3\x18\xcd\xd5\x00\xe0\xe8\x59\x09\x2d\xd4\x07\x00\x2a\xfb\x78\x21\xe5\x71\xe3\x24\x4b\x58\xd1\xc2\x60\x00\x4a\x57\xb3\xe7\xe5\x53\x73\x62\x41\x0b\xaf\x2d\xbe\x43\x00\xb8\x92\x67\xb4\x80\x6e\xa6\xf4\xce\xd8\xdd\xf5\x48\x3a\x57\xbd\xdf\x4e\xfd\xee\xc9\x8f\xe9\x59\xe7\x30\x55\x3b\x06\xd1\x23\x94\xc8\xcb\xe3\x4f\x4d\x93\x90\x93\x6b\x15\x1a\x3e\xc6\x6d\x11\x7a\xe4\x01\x90\x7c\xc8\xab\xa2\x85\xd3\x30\x1c\x64\x97\x53\xe2\xd5\x7d\x55\x2c\xcb\x74\x24\x1c\xc5\x3a\x05\x2d\x9c\xe9\x22\x7c\x33\x4a\x3b\x4b\xf7\xcf\xb8\x1b\xf3\x6c\x5d\xff\x72\xdd\x6f\x76\x99\x4a\xff\x87\xf8\x35\xd9\xdc\x4b\xdf\x4c\x31\xdf\x01\x00\x00\xff\xff\xc3\xf1\x99\xe6\x8f\x01\x00\x00"),
		},
		"/charge.disputed.created.json": &vfsgen۰CompressedFileInfo{
			name:             "charge.disputed.created.json",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 278,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8e\x4d\x4e\xc4\x30\x0c\x85\xf7\x39\x85\xe5\xf5\x48\xd3\xd9\x66\xcd\x86\x33\x20\x54\x59\x19\x43\x23\xc8\x0f\xb6\x33\xa2\x42\xbd\x3b\x4a\x4b\x4b\x37\x91\xf2\x7d\xef\xd9\xfe\x71\x00\x38\x26\x36\x42\x0f\xfd\x03\x80\xc6\xa9\x7e\x92\xf1\xf8\x60\xd1\x58\x32\x7a\x18\x1c\xc0\x72\xe9\xd9\xb7\xf8\x6d\x4d\x58\xd1\xc3\xcb\x1a\xdf\x4a\x00\x98\x29\x31\x7a\xc0\x30\x91\xbc\x33\x5e\x76\x5e\xc9\xa6\xce\xaf\x8f\xdb\x75\x73\xfa\x2f\x13\xdb\x54\xee\x5d\xd7\xa2\x76\x2e\x09\x25\x3d\x6e\x5a\x99\x96\x26\x61\x5d\x61\xe5\x63\x0c\xc2\x64\xfc\x14\xb5\x36\xe3\xe7\xfc\xd5\xa2\xcc\x47\x1f\x00\x29\x95\x96\x0d\x3d\xdc\x86\xe1\x84\x43\x13\xe1\x1c\xe6\x3e\xa6\xe9\x1d\xff\xcc\xe2\xf6\xf7\xd5\x2d\xee\x37\x00\x00\xff\xff\x10\x83\xb1\x14\x16\x01\x00\x00"),
		},
		"/charge.failed.json": &vfsgen۰CompressedFileInfo{
			name:             "charge.failed.json",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 272,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8e\xc1\x8e\x83\x20\x14\x45\xf7\x7c\xc5\xcd\x5b\x9b\xa8\x5b\xd6\xf3\x17\x93\x89\x21\xf8\xa6\x9a\x0a\x18\x78\x98\x36\x8d\xff\xde\xa0\xd5\xba\x21\xe1\x9c\x7b\x2f\xbc\x14\x40\x9d\x63\x31\xa4\x51\x2e\x00\x09\xbb\x79\x32\xc2\xdd\xc2\x31\x8d\xc1\x93\x46\xa3\x80\xb5\x2a\xd9\xff\xf1\x21\x39\x72\x22\x8d\xdf\x2d\xbe\x97\x00\xf2\xc6\x31\x69\x90\x1d\x4c\xbc\x31\x55\x07\x9f\x8d\x0c\x85\xd7\x4b\x5b\xef\x2e\x7d\xa5\x63\x19\x42\x5f\xf4\x1c\x92\x5c\x4b\xd1\xb8\x74\xfe\x69\x63\x29\xe4\x68\xb7\x27\x24\xdc\xbb\x7d\xea\x87\xed\x34\x7a\xee\xcf\x26\x40\xc6\x85\xec\x85\x34\xda\xa6\xb9\x60\x9b\x63\x64\x6f\x9f\x65\x20\xa7\x9e\x3e\x66\x55\xc7\xf9\xa7\x56\xf5\x0e\x00\x00\xff\xff\x6f\xef\xe2\x0e\x10\x01\x00\x00"),
		},
		"/charge.refunded.json": &vfsgen۰CompressedFileInfo{
			name:             "charge.refunded.json",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 411,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x90\xc1\x6a\xc3\x30\x10\x44\xef\xfa\x8a\x61\xe9\xd1\x10\xe7\xaa\x5f\x29\xc5\x08\x7b\x53\x9b\x56\x92\xd1\xae\x4c\x4b\xd0\xbf\x17\x59\x69\x1a\x8a\x4f\x39\x49\x9a\x37\x83\x76\xf6\x6a\x00\x1a\x3c\xab\x23\x8b\xfa\x00\x48\xd9\xaf\x9f\x4e\x79\xd8\x38\xc9\x12\x03\x59\xf4\x06\x28\x5d\xf5\x5e\x96\x2f\xcd\x89\x85\x2c\x5e\x77\x7b\x0b\x01\x14\x9c\x67\xb2\xa0\x71\x76\xe9\x9d\xa9\xfb\xd5\x57\xa7\x73\xd5\x4f\xdb\xf9\xd4\x98\xfc\x41\xcf\x3a\xc7\xa9\xe2\x35\x8a\x3e\x86\x92\xf3\x72\x9f\x69\xd7\x24\xe6\x34\xee\x5f\x68\xfc\x18\xb6\x45\xdc\xdd\x0f\x90\xf3\x31\x07\x25\x8b\x73\xdf\x3f\xc8\x63\x4e\x89\xc3\xf8\x5d\x63\x59\x26\xba\x91\xb2\x9f\xa5\x3b\xae\x90\xf8\x92\xc3\x74\x5c\xa1\xb1\x27\x2b\xdc\x76\x63\x41\x2f\xd7\x76\xb7\xcb\x54\xfe\x0d\x65\x80\x37\x53\xcc\x4f\x00\x00\x00\xff\xff\xbf\x7c\xf8\x9e\x9b\x01\x00\x00"),
		},
		"/charge.succeeded.json": &vfsgen۰CompressedFileInfo{
			name:             "charge.succeeded.json",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 262,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8e\xc1\x0a\x83\x30\x0c\x86\xef\x7d\x8a\x90\xb3\xa0\x5e\xfb\x2a\x63\x48\xa8\xd9\x94\xad\xad\x34\xa9\x6c\x0c\xdf\x7d\x54\xa7\xf3\x12\xc8\xf7\xfd\x3f\xc9\xc7\x00\x60\xe7\x59\x09\x2d\x94\x05\x00\x95\xfd\xf4\x24\xe5\x6e\xe6\x24\x63\x0c\x68\xa1\x31\x00\x4b\x55\xb2\xb7\xf1\xa5\x39\xb1\xa0\x85\xcb\x1a\xdf\x4a\x00\x18\xc8\x33\x5a\x40\x37\x50\xba\x33\x56\x3b\x9f\x48\x87\xc2\xeb\xb9\xad\x37\x27\x7f\xe9\x59\x87\xd8\x17\x3d\x45\xd1\x73\x29\x91\x97\xe3\xa7\x95\x49\xcc\xc9\xad\x27\x34\x3e\xba\x79\x14\x3a\xf2\x00\x48\x3e\xe6\xa0\x68\xa1\x6d\x9a\x13\x76\x39\x25\x0e\xee\x5d\x6a\x59\x7a\xfc\x99\xc5\xec\xf3\x6a\x16\xf3\x0d\x00\x00\xff\xff\x9b\xdc\x5e\x6a\x06\x01\x00\x00"),
		},
		"/checkout.session.completed.json": &vfsgen۰CompressedFileInfo{
			name:             "checkout.session.completed.json",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 1259,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x92\x4d\x6e\xdb\x30\x10\x85\xf7\x3a\x05\x41\x64\xe9\x46\x4e\x81\x6e\xb4\xea\x3d\x8a\x80\xa0\xa9\xb1\x35\x30\xff\xca\x19\x19\x31\x0c\xdd\xbd\x20\x15\x3b\xb4\xea\xfe\x28\xde\x98\xe4\x7c\x8f\x1c\xcd\x7b\x97\x46\x08\xa9\x1c\xb0\x96\x9d\xc8\x1b\x21\x24\x83\x8b\x56\x33\xa8\x13\x24\xc2\xe0\x65\x27\xb6\x8d\x10\xd3\x26\xb3\x7b\x7c\xe3\x31\x01\xc9\x4e\xfc\x28\xf8\x2c\x12\x42\x7a\xed\x40\x76\x42\x9a\x01\xcc\x31\x8c\xac\x08\xa8\xc8\x37\x57\x22\x6a\x1e\x32\xd1\x9e\x5e\xda\x2b\xd5\xbe\x53\xf4\x81\x39\xe0\x21\xf4\x19\x8c\x81\xb8\x96\x27\xed\xe8\xd6\x67\x39\xa3\xd1\x18\x20\x52\x63\xb2\x59\x30\x30\x47\xea\xda\x36\xff\xef\xd0\x3f\x87\x74\x68\xef\x2e\x11\x42\x1a\xed\x0d\xd8\x35\x8a\xa8\xcf\x0e\x3c\xab\xb9\x31\xc5\xe7\x38\x7f\xbf\x34\x3a\xf5\xf2\xb5\x22\x2d\x7a\x50\xc8\xe0\x3e\xe6\x33\xff\x2e\xd5\xba\x9a\x15\x7f\xa1\x01\x53\xfd\x58\x29\xf7\x40\x26\x61\xe4\x79\xf8\xd2\x04\xb7\x0f\x89\xf5\xce\x82\x30\x81\x39\x78\xf1\x07\xa1\x76\x61\xf4\x9c\x35\x2f\xdf\xb6\xdb\x65\xd5\x8c\x29\x81\x37\xe7\x5c\x1f\xa9\x5f\x96\x7f\x8e\xda\x33\x72\x2e\x7f\xad\x2a\xd3\x6d\xfd\xda\xd4\x27\xd3\xe6\xb1\xff\xd7\x69\x45\x7d\x80\xc7\xde\xd7\xc4\x43\xdf\x0f\xf0\x2f\xdb\xe7\xd0\x28\x2c\xf8\xd3\x65\x99\xb9\x0e\xfb\x49\xae\xea\xf6\xfd\xf1\xbf\xf6\x3b\x33\x9f\x4c\x6a\x4e\x4d\x31\x33\x67\xe6\x2e\x8e\xa9\xbf\x23\x33\x1b\x8e\x50\x9c\xe7\x70\x54\x27\x24\x2d\x6f\xd5\xa9\x92\xee\xd0\x5a\xf4\x07\xd5\x03\x6b\xb4\xb4\xbc\x05\x9c\xc6\x92\x71\xe2\x84\x11\xbe\xc3\x9b\x76\xd1\xc2\xb3\x09\xae\xba\x6f\xb5\xa7\xca\x04\xbf\xc7\xe4\xfe\xc3\xdb\xf6\xe9\x52\xef\xb3\x2b\xed\x6f\xea\x35\x43\x5c\xb8\x55\xbc\xbf\x3f\x7b\xe0\x7c\x93\x93\x3b\x35\xbf\x02\x00\x00\xff\xff\x59\xd3\xe5\x2b\xeb\x04\x00\x00"),
		},
		"/customer.created.json": &vfsgen۰CompressedFileInfo{
			name:             "customer.created.json",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 160,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\xcb\x41\xaa\xc3\x30\x0c\x84\xe1\xbd\x4f\x31\xcc\x3a\x90\xf7\xb6\xbe\x4a\x29\xc1\xb4\x2a\x09\x54\xb1\xb1\x94\x50\x28\xbe\x7b\x71\x4a\xe8\x52\xfa\xfe\x79\x07\x80\x93\x8a\x27\x46\xf4\x03\xa0\x8b\x96\x67\x72\x99\x76\xa9\xb6\xe4\x95\x11\x7f\x01\x68\x43\x6f\x1f\xcb\xcb\xb7\x2a\xc6\x88\xcb\x91\x7f\x47\x00\xd7\xa4\xc2\x08\xde\x36\xf3\xac\x52\x39\x9c\x52\x92\xcf\x5d\xc6\xfd\x7f\x3c\xd5\x7e\xac\xe2\x73\xbe\xf7\xa0\x64\x73\x1e\xef\x16\x80\x6b\x68\xe1\x13\x00\x00\xff\xff\xc9\xf8\x11\x48\xa0\x00\x00\x00"),
		},
		"/customer.deleted.json": &vfsgen۰CompressedFileInfo{
			name:             "customer.deleted.json",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 278,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x8d\x41\x0a\xc2\x30\x10\x45\xf7\x39\xc5\x67\x70\x59\xa8\x6e\x73\x15\x91\x12\xcc\x48\x03\x4d\x13\x92\x69\x11\x4a\xee\x2e\x69\x0d\x6e\x44\x77\x33\xf3\xfe\xfc\xb7\x29\x80\x06\xcf\x62\x48\xa3\x2e\x00\x09\xfb\x38\x19\xe1\x61\xe5\x94\x5d\x98\x49\xe3\xac\x80\xd2\xd5\xec\xc3\x3d\x65\x49\x9c\x49\xe3\xba\xc7\x8f\x27\x80\x66\xe3\x99\x34\xe8\xbe\x64\x09\x9e\x13\x75\x8d\x44\x23\x63\x25\xfd\x7a\xe9\x1b\xcd\x1f\xec\x59\xc6\x60\x6b\x20\x86\x2c\xb4\x9f\x4b\xf7\xbb\x7c\xb0\x3c\xb1\xb0\xfd\x23\xe9\x4f\x5b\x9b\xb5\xb3\xe5\x9b\xf3\x28\x7a\x5b\x15\x70\x53\x45\xbd\x02\x00\x00\xff\xff\x19\x49\x0c\x5a\x16\x01\x00\x00"),
		},
		"/customer.source.created.json": &vfsgen۰CompressedFileInfo{
			name:             "customer.source.created.json",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 339,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x90\x4d\x0a\xc2\x30\x10\x85\xf7\x39\xc5\x30\xb8\x2c\x44\xb7\xb9\x8a\x48\x08\xed\x48\x83\xa6\x29\x99\x69\x11\x4a\xee\x2e\x69\x1b\x75\xa3\x6e\x42\xf2\xbe\xf7\x03\x59\x14\x00\xda\x40\xe2\xd0\x40\x79\x00\xa0\x50\x18\xef\x4e\xc8\xce\x94\xd8\xc7\x01\x0d\x1c\x15\x40\x6e\x8a\xf7\xea\x1f\x32\x25\x62\x34\x70\x5e\xed\x5b\x08\x00\x07\x17\x08\x0d\x60\x3b\xb1\xc4\x40\x09\x9b\x4a\x46\x27\x7d\x21\x7a\x3e\xe9\x4a\xf9\x8d\x03\x49\x1f\xbb\x62\x18\x23\x0b\xae\x72\x6e\x7e\x97\x5b\x8e\x53\x6a\xe9\xcf\x86\x3e\x2c\xf5\x6e\x7c\x97\xf5\x16\xfa\x3e\xfd\xd1\x96\x5c\xe0\xd7\x97\xac\xda\xbe\x68\x00\x25\xde\xec\xec\xd9\xe1\x0e\xb3\xaa\xe7\x45\x65\xf5\x0c\x00\x00\xff\xff\x36\x44\x63\x6b\x53\x01\x00\x00"),
		},
		"/customer.source.updated.json": &vfsgen۰CompressedFileInfo{
			name:             "customer.source.updated.json",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 574,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x91\xc1\x4e\xc3\x30\x0c\x86\xef\x79\x0a\xcb\xe2\x58\xa9\x70\xcd\xab\x20\x14\x99\xc5\xd3\x2a\xc8\x5c\xc5\x4e\x85\x34\xe5\xdd\x51\xba\x74\x94\x03\x20\x76\x6b\xff\xcf\xfe\x7f\x3b\xbe\x38\x00\x0c\x89\x8d\xd0\x43\xfb\x01\x40\xe3\x34\xbf\x93\x71\x58\x38\xeb\x24\x67\xf4\xf0\xe8\x00\xea\xd0\x6a\x8f\xd3\x87\x95\xcc\x8a\x1e\x9e\xd7\xf2\x6b\x13\x00\x9e\x29\x31\x7a\xc0\x43\x51\x93\xc4\x19\x87\x8d\xcc\x64\xa7\x46\xc6\xe5\x69\xdc\xa8\x7e\xe1\xc4\x76\x92\xd8\x0a\x66\x51\xc3\x55\xae\xc3\xef\xe6\x41\xa5\xe4\x03\xff\x91\x31\x3e\x5c\xb6\x6f\x3f\xc5\x3a\x5e\x9b\x7e\x8e\xde\xb9\x65\x4a\x7a\x7b\x92\x55\xeb\x89\x1e\xd0\xe4\x2d\x2c\x93\x12\x76\x58\xff\x33\x72\x28\x73\x24\xe3\x78\xdf\xe8\x3b\xb9\xfb\x35\x7a\xdf\x42\xed\xea\x91\x76\x97\xef\xfa\x51\xa4\x19\xbc\x52\xc6\x9b\x5c\xbf\xaf\xea\x00\x5e\x5c\x75\x9f\x01\x00\x00\xff\xff\x0b\x34\xd5\x59\x3e\x02\x00\x00"),
		},
		"/customer.subscription.created.json": &vfsgen۰CompressedFileInfo{
			name:             "customer.subscription.created.json",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 801,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x92\xcf\x6a\xc4\x20\x10\xc6\xef\xfb\x14\xc3\xd0\xe3\xc2\xa6\x3d\xfa\x2a\xa5\x04\x6b\x2c\x91\xc6\x3f\x38\x63\xe8\xb2\xf8\xee\xc5\x6c\x6c\x4d\x4a\x0e\x65\x6f\x8e\xbf\xf9\x9c\xcf\x4f\x6f\x27\x00\x00\xec\xad\x66\x89\x02\xee\x25\x00\xb2\xb6\x61\x92\xac\xfb\x59\x47\x32\xde\xa1\x80\x6e\x61\xf9\x7c\x57\x7c\x98\x2f\x4e\x51\x13\x0a\x78\x5d\x45\x55\x0c\x80\x4e\x5a\x8d\x02\x50\x25\x62\x6f\x75\xc4\xf3\x2f\x0b\x92\xc7\xc2\x2e\xf3\xf3\xa5\x72\x6a\x1b\xac\xe6\xd1\x0f\xa5\x25\x78\xe2\xad\x34\x4a\x4b\x8d\xcf\x65\x97\x7c\x8a\x6a\x19\xc7\xfe\xb3\x9f\x0d\x49\xfc\xc1\x79\x5d\xe5\xf3\xb1\xc9\x30\x49\x77\x64\xb0\xb0\x87\xcc\xa9\x14\xa3\x76\xea\x5a\x14\x89\x86\x46\x00\x80\xc6\xb1\x8e\xb3\x9c\x0a\xb4\xde\xf1\xb8\xc5\xd2\xfa\xe4\x18\x05\xbc\x74\x5d\xb7\x21\x21\xfa\x21\x29\xde\x0d\x6b\xee\x64\xaf\xb5\xa5\xe1\xf9\x5f\xb1\x50\x7a\x27\x15\x4d\xe0\xf2\xfc\x07\xf1\xb4\x3d\x0f\xc6\xb4\x7e\x14\x01\xf8\x74\xab\x95\x30\x43\xde\x25\xc6\xda\xb6\x7f\x6e\xef\xbe\x0e\x29\x4f\xba\x1c\x55\x56\xcb\x31\x9b\x96\xdc\x54\x6f\x7f\x53\x39\xd5\xfd\x7c\xfa\x0e\x00\x00\xff\xff\x94\x09\xf3\x6f\x21\x03\x00\x00"),
		},
		"/customer.subscription.deleted.json": &vfsgen۰CompressedFileInfo{
			name:             "customer.subscription.deleted.json",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 851,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x52\xcd\x6e\xf3\x20\x10\xbc\xfb\x29\x56\xab\x1c\x23\xc5\xdf\x77\xe4\x55\xaa\xc8\xa2\x66\xab\xa0\x9a\x1f\xc1\x62\x35\x8a\x78\xf7\x0a\x1c\x3b\x44\xf2\x21\x8d\x2f\x66\x99\x61\x99\x99\xe5\xd6\x01\xe0\x60\x88\x25\x0a\x28\x05\x00\x32\x19\x3f\x49\xa6\x61\xa6\x10\xb5\xb3\x28\xa0\xef\x00\xf2\xb1\x70\xbf\xf4\x0f\xa7\x40\x11\x05\x7c\x54\xfa\x72\x08\x00\xad\x34\x84\x02\x70\x4c\x91\x9d\xa1\x80\xc7\x15\xf1\x92\x2f\x05\x39\xcd\xff\x4e\x2b\x1a\x1f\xb0\x21\xbe\x38\x55\x08\xde\x45\x6e\x8f\x05\x69\xe2\xa6\xab\xee\x45\x97\xc2\x58\xaf\x61\xf7\x3d\xcc\x3a\x4a\xbc\x83\xb9\xfe\xf3\x71\x5f\x94\x9f\xa4\xdd\x17\x54\x90\x37\xc5\x8c\x29\x04\xb2\xe3\xb5\xb0\x53\x54\x1b\x19\x00\xb5\x65\x0a\xb3\x9c\x0a\x64\x9c\xe5\x4b\x0b\x4a\xe3\x92\x65\x14\xf0\xbf\xef\xfb\x66\xdf\x07\xa7\xd2\xc8\x4f\x97\x34\x1e\xcc\x75\x25\x6c\x68\x7e\xc9\x7c\x4c\x9f\x71\x0c\xda\x73\x19\xe6\x6e\x08\x2d\xe3\xed\x30\xee\x63\x17\x80\x87\xdb\x5a\x09\xad\xf2\x53\x2e\x4c\xe6\xf1\x76\x96\xaf\x35\x0b\xf7\x51\xd5\x26\x65\x55\x1b\x34\x84\xbc\xad\xcf\x7f\xf6\x3e\x28\x9a\x88\x49\xbd\x90\xc1\xe9\x70\x6b\xeb\x27\x1b\x4d\x2c\x4b\xc3\x45\x5f\xd1\x71\xee\x72\xf7\x1b\x00\x00\xff\xff\xa6\xf2\x2a\xfa\x53\x03\x00\x00"),
		},
		"/customer.subscription.updated.json": &vfsgen۰CompressedFileInfo{
			name:             "customer.subscription.updated.json",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 931,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x53\x4d\x6b\xc3\x30\x0c\xbd\xe7\x57\x08\xd1\x63\xa1\xd9\x8e\xfe\x2b\x63\x04\x37\x56\x69\x58\xfd\x81\x2d\x87\x95\xe2\xff\x3e\x94\x26\xa9\x0b\x1d\xb4\xcd\x25\x96\xde\xb3\xac\xf7\x64\x5f\x1a\x00\xec\x2c\xb1\x46\x05\x12\x00\x20\x93\x0d\x27\xcd\xd4\x8d\x14\xd3\xe0\x1d\x2a\x68\x1b\x80\xb2\x15\xee\x61\xf8\xe5\x1c\x29\xa1\x82\xaf\x89\x7e\xdd\x04\x80\x4e\x5b\x42\x05\xd8\xe7\xc4\xde\x52\xc4\xed\x82\x04\xcd\x47\x41\x76\xe3\xc7\x6e\x41\xd3\x0d\xb6\xc4\x47\x6f\x84\x10\x7c\xe2\x7a\x5b\xd4\x36\xad\x7d\x4d\xb9\xe4\x73\xec\xa7\x63\xd8\xff\x74\xe3\x90\x34\xce\x60\x99\xfe\x65\xfb\xb8\xa9\x70\xd2\xee\x71\x43\x82\xbc\xd9\x4c\x9f\x63\x24\xd7\x9f\x85\x9d\x93\x59\xc9\x00\x38\x38\xa6\x38\xea\x93\x40\xd6\x3b\x3e\xd6\xa0\xb6\x3e\x3b\x46\x05\x9f\x6d\xdb\x56\xf9\x10\xbd\xc9\x3d\xdf\x1d\x52\x69\xb0\xe7\x85\xb0\xa2\xe5\x29\xf1\x29\xef\x53\x1f\x87\xc0\x32\xcc\x87\x26\xd4\x8c\xb7\xcd\x98\xc7\xae\x00\x37\x97\x25\x52\x83\x29\x77\xbe\x30\xd9\xdb\xdd\xb9\x7e\xb5\x58\x98\x47\x35\x15\x91\xd5\x54\xa0\x22\x94\x75\xfd\xfd\xb2\xf6\x2e\x07\xa3\x99\xcc\x13\x1e\xec\x36\x97\x3a\xbe\x93\xf1\x92\x2d\xf2\xb4\x8c\xae\x9e\xd7\x9c\x3f\x78\x2f\x05\xf6\x3a\xfe\x3b\xcd\x46\x34\x96\xe6\x2f\x00\x00\xff\xff\x13\xcd\xf0\xab\xa3\x03\x00\x00"),
		},
		"/customer.updated.json": &vfsgen۰CompressedFileInfo{
			name:             "customer.updated.json",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 358,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x90\xcd\x0a\xc2\x30\x10\x84\xef\x79\x8a\x65\xf1\x58\xa8\x5e\xf3\x2a\x22\x65\x35\x5b\x5a\x30\x4d\x48\xb6\x45\x28\x79\x77\x49\x35\x35\x17\xf5\x92\x9f\x99\x6f\x67\x61\x56\x05\x80\x9d\x65\x21\xd4\x90\x3f\x00\x28\x6c\xfd\x9d\x84\xbb\x85\x43\x1c\xdd\x84\x1a\x8e\x0a\x20\x35\x99\xed\xc7\x87\xcc\x81\x23\x6a\x38\x6f\xf8\x6b\x08\x00\x27\xb2\x8c\x1a\xf0\x36\x47\x71\x96\x03\x36\xc5\xf1\x24\x43\x76\xda\xe5\xd4\x16\x37\x7e\x6c\xcb\x32\x38\x93\x01\xef\xa2\xe0\x26\xa7\xe6\x77\x78\x37\x7b\x43\xc2\xe6\xcf\x92\xf6\xb0\x96\xb7\x1e\x4d\xfa\xba\xb3\x4a\x09\x64\xe3\xde\x45\x61\xc9\x50\xd5\xd0\x5b\xef\x9d\xcb\x01\x57\x0a\xb8\xcb\x49\xd5\x77\x3e\x2f\x2a\xa9\x67\x00\x00\x00\xff\xff\x7b\x9d\x95\x3c\x66\x01\x00\x00"),
		},
		"/invoice.created.json": &vfsgen۰CompressedFileInfo{
			name:             "invoice.created.json",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 585,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x90\xd1\x4a\xc4\x30\x10\x45\xdf\xf3\x15\x97\xc1\xc7\x85\xad\x3e\xe6\x57\x44\x4a\x48\x47\x36\x68\x92\x92\x99\x14\x65\xc9\xbf\x4b\xaa\x5d\x2b\xac\x82\xfb\x14\x26\xf7\x0c\x39\x37\x67\x03\xd0\x18\x59\x1d\x59\xf4\x01\x20\xe5\x38\xbf\x3a\xe5\x71\xe1\x22\x21\x27\xb2\x18\x0c\xd0\x0e\x9d\x7d\x0e\x6f\x5a\x0b\x0b\x59\x3c\xae\xf8\xe7\x12\x40\xc9\x45\x26\x0b\xf2\x55\x34\x47\x2e\x74\xd8\x92\xd9\xe9\xa9\x27\xc7\xe5\xfe\xb8\xa5\xf2\x1d\x47\xd6\x53\x9e\x3a\x30\x67\xd1\xfd\x5a\x71\x51\x2e\x5e\xeb\x9d\xe4\x5a\xfc\xfa\x8c\xe6\x97\x71\x09\xe2\xe8\x2b\x6c\xeb\xd9\x0e\xd7\xa5\x42\x5a\x72\xf0\x1c\x94\xe3\x75\xaf\x1d\x70\xa3\x9a\xaf\xa5\x70\xf2\xef\x9d\xae\x32\x5d\x60\xec\xbe\xc4\x82\xee\xce\xdb\x64\xc3\xd4\xf6\x94\x8b\xb9\x26\x25\x8b\x87\x61\x18\xfe\xd3\xea\xcf\x46\x37\xb7\xf9\x55\xf9\xa7\x9a\x01\x9e\x4c\x33\x1f\x01\x00\x00\xff\xff\xa4\xd1\x46\x8f\x49\x02\x00\x00"),
		},
		"/invoice.finalized.json": &vfsgen۰CompressedFileInfo{
			name:             "invoice.finalized.json",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 709,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x90\xc1\x6e\xf3\x20\x0c\xc7\xef\x3c\x85\x65\xf5\x18\x29\xf9\xbe\x23\xaf\x32\x4d\x11\x22\xae\x8a\x56\x20\xc2\x26\xda\x16\xf1\xee\x13\x59\xe9\x72\x68\x2b\xb5\x27\x30\xff\x9f\xe5\x9f\x59\x15\x00\x8e\x9e\xc4\xa0\x86\x5a\x00\xa0\x90\x9f\xcf\x46\x68\x5c\x28\xb1\x8b\x01\x35\x0c\x0a\xa0\x74\x95\x3d\xba\x4f\xc9\x89\x18\x35\xbc\x6d\xf8\x6f\x13\x00\x06\xe3\x09\x35\xa0\xcd\x2c\xd1\x53\xc2\xae\x25\xb3\x91\x53\x4d\xfa\xe5\x5f\xdf\x52\xfe\x8b\x3d\xc9\x29\x4e\x15\x98\x23\xcb\xbe\x2d\x19\xcf\x57\xaf\xed\x8d\x63\x4e\x76\x1b\x23\xf1\x63\x5c\x1c\x1b\xbc\x84\x65\x3b\x4b\x77\x5b\xca\x85\x25\x3a\x4b\x4e\xc8\xdf\xf6\xda\x01\x2f\xaa\xd9\x9c\x12\x05\xfb\x55\xe9\xcc\xd3\x15\x86\xdd\x97\x68\xc0\xc3\xda\x2a\xed\xa6\xb2\xa7\x8c\x8f\x39\x08\x6a\xf8\x3f\x0c\xc3\x33\x5b\x3d\xdc\xe8\xe5\x6d\xee\x2a\x3f\xa3\x36\x1e\x5d\x30\x67\xf7\x4d\xd3\x63\xc9\xfe\xb0\x5e\xae\x75\x44\xdf\xba\xee\xba\xab\x66\xf0\xae\x8a\xfa\x09\x00\x00\xff\xff\x6d\xbf\x92\xc1\xc5\x02\x00\x00"),
		},
		"/invoice.payment_failed.json": &vfsgen۰CompressedFileInfo{
			name:             "invoice.payment_failed.json",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 712,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x90\xc1\x6a\x84\x30\x10\x86\xef\x3e\xc5\x30\xec\x51\xd0\xf6\x98\x6b\xa1\x2f\x51\x8a\x0c\x71\x5a\x43\x37\x89\x24\x13\xe9\x22\x79\xf7\x12\xab\x5b\x5b\xb6\x0b\xbb\x27\x1d\xff\x6f\x66\xbe\x71\xae\x00\xb0\xb3\x2c\x84\x0a\x4a\x01\x80\xc2\x76\x3c\x92\x70\x37\x71\x88\xc6\x3b\x54\xd0\x56\x00\xb9\x2e\xec\x9b\xf9\x94\x14\x38\xa2\x82\x97\x05\xff\x6e\x02\x40\x47\x96\x51\x01\xea\x14\xc5\x5b\x0e\x58\x6f\xc9\x48\x32\x94\xa4\x99\x1e\x9a\x2d\x8d\x3f\xb1\x65\x19\x7c\x5f\x80\xd1\x47\xd9\xb7\x05\xb2\xf1\xec\xb5\x7c\x8b\x3e\x05\xbd\xac\x11\xff\xd1\xe9\x81\xc2\x3b\x3f\xad\x23\x9f\xc9\x1c\x71\x45\xf3\xf2\xcc\xf5\x65\x45\xe3\x26\x6f\x34\x1b\x61\x7b\xd9\x72\x07\xdc\x29\xaa\x53\x08\xec\xf4\xa9\xd0\x29\xf6\x67\x18\x76\x3f\x48\x01\x1e\xe6\xad\x52\xa6\xcf\x7b\x8a\xac\x4f\x4e\x50\xc1\x63\xdb\xb6\xb7\x5c\x75\xf5\xa2\xbb\xaf\xf9\x57\xf9\x16\xb5\x6e\xa4\xd3\x75\xbd\xe6\x30\xaf\xaf\x65\x78\xf3\x8b\xff\x23\x5c\x6d\x6b\x5f\xab\x5c\x7d\x05\x00\x00\xff\xff\x15\x9b\x00\x1a\xc8\x02\x00\x00"),
		},
		"/invoice.payment_succeeded.json": &vfsgen۰CompressedFileInfo{
			name:             "invoice.payment_succeeded.json",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 698,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x90\xc1\x6a\xc3\x30\x0c\x86\xef\x7e\x0a\x21\x7a\x0c\x24\xdb\xd1\xaf\x32\x46\x30\x8e\x46\xcd\x66\x3b\x58\x72\x58\x09\x7e\xf7\xe1\xac\xe9\xbc\xd1\x15\xda\x53\x22\xff\x9f\xac\x4f\x5e\x15\x00\x8e\x9e\xc4\xa0\x86\x5a\x00\xa0\x90\x9f\x3f\x8c\xd0\xb8\x50\x62\x17\x03\x6a\x18\x14\x40\xe9\x2a\xfb\xe6\x3e\x25\x27\x62\xd4\xf0\xb2\xe1\xdf\x4d\x00\x18\x8c\x27\xd4\x80\x36\xb3\x44\x4f\x09\xbb\x3d\x99\x8d\x1c\x6b\xd2\x2f\x4f\xfd\x9e\xf2\x4f\xec\x49\x8e\x71\xaa\xc0\x1c\x59\xda\xb6\x64\x3c\x5f\xbc\xb6\x33\x8e\x39\xd9\x6d\x8c\xc4\xf7\x71\x71\x6c\xf0\x1c\x96\xed\x5b\xba\xeb\x52\x2e\x2c\xd1\x59\x72\x42\xfe\xba\x57\x03\x3c\xa8\x66\x73\x4a\x14\xec\xa9\xd2\x99\xa7\x0b\x0c\xcd\x93\x68\xc0\xc3\xba\x57\xda\x4d\xa5\xa5\x8c\x8f\x39\x08\x6a\x78\x1e\x86\xe1\x9e\xad\x6e\x6e\xf4\xf0\x36\xff\x2a\xdf\xa3\x36\xce\xe6\x74\x5b\xaf\x3f\xac\xe7\xdf\x7a\x79\xff\x8b\xff\x23\xac\xf6\xb1\xaf\xaa\xa8\xaf\x00\x00\x00\xff\xff\x55\x80\x9d\x12\xba\x02\x00\x00"),
		},
		"/invoice.updated.json": &vfsgen۰CompressedFileInfo{
			name:             "invoice.updated.json",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 724,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x92\x41\x6b\xc4\x20\x10\x85\xef\xfe\x8a\x61\xd8\x63\x20\x69\x8f\xfe\x95\x52\x82\x8d\xb3\xac\x50\x35\xe8\x18\x5a\x82\xff\xbd\x98\x26\xa9\x0b\xbb\x4b\xb3\xb7\xf8\xde\xe7\xe4\x3d\x99\x59\x00\x60\x6f\x89\x15\x4a\x28\x07\x00\x64\xb2\xe3\xa7\x62\xea\x27\x0a\xd1\x78\x87\x12\x3a\x01\x90\x9b\xc2\x9e\xcd\x17\xa7\x40\x11\x25\xbc\x2d\xf8\xef\x25\x00\x74\xca\x12\x4a\xc0\x21\x45\xf6\x96\x02\x36\x9b\x33\x2a\xbe\x14\xa7\x9d\x5e\xda\xcd\x8d\x7f\xb6\x25\xbe\x78\x5d\x80\xd1\x47\xc6\x45\xce\xcd\xed\xe1\xc6\x4d\xde\x0c\x64\x98\xec\xed\xf9\x15\x70\xff\x17\xd5\xcd\xa0\x6c\xdc\xab\x2f\xda\x90\x42\x20\x37\x7c\x17\x3a\x45\xbd\xc3\x50\x55\x93\x80\xa7\x79\x3b\x49\xa3\x73\x4d\x29\xeb\x93\x63\x94\xf0\xda\x75\xdd\x2a\xe7\xff\xb4\x7a\xd8\xe8\xe9\x36\x77\x23\x1f\x89\xd6\xa7\x51\x2b\x26\xfd\x38\x62\x7b\x9a\xd7\xcf\xab\x37\x39\x14\xb8\xec\xa2\x56\xd5\x3e\xae\xfa\xd9\xfb\x32\xe0\x43\x05\xdc\xe5\x7c\x5d\x41\x00\xbc\x8b\x2c\x7e\x02\x00\x00\xff\xff\x06\x47\x1e\x10\xd4\x02\x00\x00"),
		},
		"/payment_intent.amount_capturable_updated.json": &vfsgen۰CompressedFileInfo{
			name:             "payment_intent.amount_capturable_updated.json",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 751,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x92\xc1\x6a\xc3\x30\x10\x44\xef\xfe\x8a\x65\xe9\x31\xd4\x8e\xe9\x49\xbf\x52\x8a\x51\x9d\x0d\x11\x44\x2b\xb1\x5a\x85\x98\xe0\x7f\x2f\x8a\xd3\xc4\x6a\x4b\x0e\xc5\x07\x23\xcd\x9b\xd1\x32\xd2\xa5\x01\xc0\xc1\x93\x5a\x34\x50\x16\x00\xa8\xe4\xe3\xd1\x2a\x0d\x27\x92\xe4\x02\xa3\x81\xae\x01\x98\x37\x85\xdd\xbb\xb3\x66\xa1\x84\x06\xde\xaf\xf8\x62\x02\x40\xb6\x9e\xd0\x00\x46\x3b\x79\x62\x2d\xa1\x87\xb0\xc3\xcd\xb7\x1e\xad\x1e\x8a\xde\x9e\xb6\x6d\xcd\xa4\x07\x74\x33\x95\x98\x90\x74\x6d\x16\xeb\xd3\x7d\xc6\x65\xce\x29\x5e\x0f\x1c\xad\x3c\x8e\x81\xdb\x7a\x4d\x96\xe1\xb2\xff\x24\x29\xf4\x5b\x5f\x7f\x2b\x27\x00\xd2\x39\x0e\x3e\xf0\x32\xe9\xf6\x0f\x71\x22\x7b\x8d\xe9\xbb\xbe\xab\xd5\xf1\x34\xde\xf2\xf1\xbe\x3d\x37\xeb\xff\xbc\x79\xde\x98\x63\x25\xd6\xe7\x8d\x2d\xcc\x3f\x1b\xb3\x3e\x64\x56\x34\xd0\x77\x5d\x57\x35\x16\xcb\xa5\x0e\x8f\x2c\x6f\x39\xdb\x63\xd5\x6a\xe0\xbd\x13\x8f\x06\x54\x32\xfd\x16\xac\xba\xc0\xcf\x13\xb2\x08\xf1\x38\x15\x35\xa7\xea\xca\x7e\xbc\x19\x03\xf8\x72\xa9\xf7\x8c\xdb\xcd\x6b\x87\x90\x66\xe1\x21\xcb\xb1\xd0\x07\xd5\x98\x4c\xdb\x26\x15\x17\xe9\x75\x0c\x1e\xeb\xe6\x1b\x80\x8f\x66\x6e\xbe\x02\x00\x00\xff\xff\xf6\x13\xe9\xbc\xef\x02\x00\x00"),
		},
		"/payment_intent.canceled.json": &vfsgen۰CompressedFileInfo{
			name:             "payment_intent.canceled.json",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 473,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x51\xcd\x6a\xc4\x20\x10\xbe\xfb\x14\xc3\xd0\x63\x20\x69\x8f\xbe\x4a\x29\x32\xd5\x29\x2b\x44\x4d\x75\x5c\x1a\x16\xdf\xbd\xc4\xb4\x5d\xd2\x53\xd8\x93\xf8\xfd\xf1\x7d\xcc\x4d\x01\xa0\x09\x2c\x84\x1a\xb6\x0f\x00\x0a\x87\x65\x26\x61\x73\xe5\x5c\x7c\x8a\xa8\x61\x52\x00\x6d\xd8\xb4\x1f\xfe\x4b\x6a\xe6\x82\x1a\x5e\xbb\x7c\x37\x01\x60\xa4\xc0\xa8\x01\x17\x5a\x03\x47\x31\x3e\x0a\x47\xc1\xe1\x97\x5f\x48\x2e\x1b\x3f\x5e\x9f\xc7\xa3\xa6\xdc\x45\x81\xe5\x92\x5c\x8f\x49\xe5\x60\xce\x14\xca\x5f\xc7\x8e\x51\x48\x35\x0a\x6a\x78\x99\xa6\x69\xb8\xe3\xb6\xe6\xcc\xd1\xae\x5b\x4a\x2d\x0e\x7f\x98\xd6\xdf\x36\x9c\x69\x6d\x2c\x45\xcb\x33\xbb\x53\xf5\xc7\xa7\xdb\x11\xd1\xde\xb5\x71\x8f\x78\x6c\xda\xee\x9d\x49\x7c\x8a\x26\x33\x95\x7e\x05\xcc\xfc\x59\xb9\x08\x3b\xf3\xbe\x1a\x5b\x8b\xa4\xc0\xf9\xdf\x3e\x05\xf0\xa6\x9a\xfa\x0e\x00\x00\xff\xff\xba\xb3\xa7\x14\xd9\x01\x00\x00"),
		},
		"/payment_intent.created.json": &vfsgen۰CompressedFileInfo{
			name:             "payment_intent.created.json",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 291,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8f\xcf\xca\x83\x30\x10\xc4\xef\x79\x8a\x65\xcf\x82\xf9\xbe\x63\x5e\xa5\x48\x58\x74\x8b\x42\xf3\x87\x64\x23\x95\xe2\xbb\x97\x68\xad\xf5\x12\xc8\xcc\x6f\x66\x92\x97\x02\x40\xeb\x58\x08\x0d\xd4\x0b\x00\x0a\xbb\xf8\x20\x61\x3b\x73\xca\x53\xf0\x68\x40\x2b\x80\xb5\xa9\xec\x7d\x7a\x4a\x49\x9c\xd1\xc0\x6d\xc3\xf7\x10\x00\x7a\x72\x8c\x06\x30\xd2\xe2\xd8\x8b\x9d\xbc\xb0\x17\x6c\x0e\x3f\x92\x8c\xd5\x6f\xe7\xbf\xf6\xca\xe4\x13\x72\x2c\x63\x18\xb6\x9a\x90\x2f\xe1\x44\x2e\x7f\xdf\xb8\x69\xe4\x42\xf1\x82\x06\xfe\xb5\xd6\xcd\xa9\xf7\x25\x25\xf6\xfd\x52\x5b\x4a\x1e\xf0\xc7\x3a\x76\xf7\x19\x2b\x4b\xdc\x7f\x82\x3d\xa5\x01\xbb\x0f\xb8\xaa\xe3\xec\xd4\xaa\xde\x01\x00\x00\xff\xff\xf4\x28\x40\xb3\x23\x01\x00\x00"),
		},
		"/payment_intent.payment_failed.json": &vfsgen۰CompressedFileInfo{
			name:             "payment_intent.payment_failed.json",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 674,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x52\xcd\x6a\xc3\x30\x0c\xbe\xfb\x29\x84\xd8\x31\x50\x37\xf4\xe4\x57\x19\x25\x78\xa9\x4a\x03\xb3\x1d\x6c\x39\x34\x94\xbc\xfb\x50\xd2\xb5\xf6\xd8\x7a\x58\x2e\xc1\xfe\x7e\xf4\xf1\xc9\x37\x05\x80\x9d\x23\xb6\x68\x40\x0e\x00\xc8\xe4\xc6\x4f\xcb\xd4\x4d\x14\xd3\x10\x3c\x1a\xd0\x0a\x60\x69\x84\x7b\x1e\xae\x9c\x23\x25\x34\xf0\xbe\xd2\x37\x11\x00\x7a\xeb\x08\x0d\xe0\x68\x67\x47\x9e\xc5\xf4\x12\x4e\xd8\x7c\xe3\xa3\xe5\x8b\xe0\xbb\x69\xbf\xab\x39\xe9\x49\xba\x8b\xc4\x26\x24\x2e\xc5\xd1\xba\xf4\xc8\xb8\xe5\x9c\xc7\x75\x60\x6f\xe3\x73\x0c\xdc\xcf\x25\x53\xc2\x65\xf7\x41\x51\xd8\x07\x5d\x7d\x6d\xa1\x04\x40\xba\x8e\x9d\x0b\x7e\x4b\xba\xff\x05\x9c\xc9\xae\x36\xad\x6e\x75\x8d\xf6\x53\xbf\xfa\xb7\x07\x7c\x5c\x2f\xaa\xfc\x2f\xcd\xeb\xc6\x06\xcf\xe4\xf9\x75\x63\x1b\xe7\x9f\x8d\x59\x17\xb2\x67\x34\xd0\x6a\xad\xcb\xc6\x82\x3f\x0f\xd1\x89\x09\xc7\x4c\x55\x99\x39\x46\xf2\xfd\x2c\x58\x4e\x55\xcf\x3f\x16\x6d\x00\xdf\x6e\xf5\x9d\x19\x4e\xcb\xdf\x8a\x4e\x16\xb8\x3e\xa4\x6d\x65\xc7\xba\x2c\x05\x70\x54\x8b\xfa\x0a\x00\x00\xff\xff\xa2\x70\xa1\x29\xa2\x02\x00\x00"),
		},
		"/payment_intent.succeeded.json": &vfsgen۰CompressedFileInfo{
			name:             "payment_intent.succeeded.json",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 674,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x52\x4d\x6a\xf3\x30\x10\xdd\xeb\x14\xc3\xf0\x2d\x0d\x51\xcc\xb7\xd2\x55\x4a\x30\xaa\x33\x21\x86\x4a\x32\xd2\xc8\xc4\x04\xdf\xbd\x8c\xed\x26\x52\x69\xb3\x28\x5e\x18\xe9\xfd\xcc\xe3\x8d\xee\x0a\x00\x3b\x47\x6c\xd1\x80\x1c\x00\x90\xc9\x8d\x1f\x96\xa9\x9b\x28\xa6\x21\x78\x34\xa0\x15\xc0\xd2\x08\xf7\x32\xdc\x38\x47\x4a\x68\xe0\x6d\xa5\x6f\x22\x00\xf4\xd6\x11\x1a\xc0\xd1\xce\x8e\x3c\x8b\xe9\x35\x9c\xb1\xf9\xc2\x47\xcb\x57\xc1\x0f\xd3\xf1\x50\x73\xd2\x93\xb4\x8b\xc4\x26\x24\x2e\xc5\xd1\xba\xf4\xc8\xb8\xe5\x9c\xc7\x75\x60\x6f\xe3\x73\x0c\xec\xe7\x92\x29\xe1\xb2\x7b\xa7\x28\xec\xff\x6d\xfd\x15\x4a\x00\xa4\xdb\xd8\xb9\xe0\xb7\xa4\xc7\x1f\xc0\x99\xec\x6a\xd3\xea\x56\xd7\x68\x3f\xf5\xbb\x3f\x3e\xae\x17\x55\xfe\x97\xe6\x75\x63\x83\x67\xf2\xfc\xba\xb1\x8d\xf3\xc7\xc6\xac\x0b\xd9\x33\x1a\x68\xb5\xd6\x65\x63\xc1\x5f\x86\xe8\xc4\x84\x63\xa6\xaa\xcc\x1c\x23\xf9\x7e\x16\x2c\xa7\xaa\xe7\x6f\x8b\x36\x80\xff\xee\xf5\x9d\x19\xce\xcb\xef\x8a\x4e\x16\xb8\x3e\xa4\x6d\x65\xa7\xba\x2c\x05\x70\x52\x8b\xfa\x0c\x00\x00\xff\xff\x3b\xe9\xa9\x40\xa2\x02\x00\x00"),
		},
		"/payment_method.attached.json": &vfsgen۰CompressedFileInfo{
			name:             "payment_method.attached.json",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 606,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x51\xcd\x6a\x84\x30\x10\xbe\xe7\x29\x86\xa1\x47\x41\x57\x7a\xca\xab\x94\x12\xa6\xee\x14\x85\xc6\x84\x64\x94\x95\x25\xef\x5e\xa2\xee\xae\x16\x0b\x92\x43\x98\xf9\x7e\x66\xf2\xe5\xae\x00\xd0\x58\x16\x42\x0d\xb9\x00\x40\x61\xeb\x7f\x48\xd8\x8c\x1c\x62\xe7\x7a\xd4\x50\x29\x80\x54\x64\xee\x77\x77\x93\x21\x70\x44\x0d\x1f\x33\x7d\x11\x01\x60\x4f\x96\x51\x03\x36\x43\x14\x67\x39\x60\xf1\x40\x3c\x49\x9b\x91\x72\xbc\x94\x0f\x34\xbe\x60\xcb\xd2\xba\x6b\x26\x78\x17\x05\xe7\x76\x2a\x8e\xcd\x3d\x4d\x96\x7b\x31\xab\xe6\x70\xc4\x9e\xf3\xff\xa0\x8d\x38\x90\x8d\xcf\x00\x96\x10\x26\xbf\xbc\x86\xc2\x6b\x0c\xac\xf5\x96\x99\x97\x1b\xec\x17\x87\xcc\x7e\xaf\xf7\x67\xa3\x04\x40\xbe\x79\x63\x5d\xbf\x6c\x7a\x39\x00\x27\xa6\xd9\xa6\xae\xea\x6a\x8f\x36\x63\xb3\xfa\xe3\xb3\x9d\xd4\xf6\x3e\x97\x98\x21\x11\x6a\xda\x53\xc1\x95\x6f\xf7\x7d\x47\x77\xd7\x54\xfe\x35\x38\xfc\x3d\x05\xf0\xa9\x92\xfa\x0d\x00\x00\xff\xff\x36\x69\xa0\x73\x5e\x02\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/charge.captured.json"].(os.FileInfo),
		fs["/charge.disputed.created.json"].(os.FileInfo),
		fs["/charge.failed.json"].(os.FileInfo),
		fs["/charge.refunded.json"].(os.FileInfo),
		fs["/charge.succeeded.json"].(os.FileInfo),
		fs["/checkout.session.completed.json"].(os.FileInfo),
		fs["/customer.created.json"].(os.FileInfo),
		fs["/customer.deleted.json"].(os.FileInfo),
		fs["/customer.source.created.json"].(os.FileInfo),
		fs["/customer.source.updated.json"].(os.FileInfo),
		fs["/customer.subscription.created.json"].(os.FileInfo),
		fs["/customer.subscription.deleted.json"].(os.FileInfo),
		fs["/customer.subscription.updated.json"].(os.FileInfo),
		fs["/customer.updated.json"].(os.FileInfo),
		fs["/invoice.created.json"].(os.FileInfo),
		fs["/invoice.finalized.json"].(os.FileInfo),
		fs["/invoice.payment_failed.json"].(os.FileInfo),
		fs["/invoice.payment_succeeded.json"].(os.FileInfo),
		fs["/invoice.updated.json"].(os.FileInfo),
		fs["/payment_intent.amount_capturable_updated.json"].(os.FileInfo),
		fs["/payment_intent.canceled.json"].(os.FileInfo),
		fs["/payment_intent.created.json"].(os.FileInfo),
		fs["/payment_intent.payment_failed.json"].(os.FileInfo),
		fs["/payment_intent.succeeded.json"].(os.FileInfo),
		fs["/payment_method.attached.json"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
