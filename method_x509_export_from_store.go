package kalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include "cpp/KalkanCrypt.h"
//
// int X509ExportCertificateFromStore(char *alias, int flag, char *outCert, int *outCertLength) {
//     return kc_funcs->X509ExportCertificateFromStore(alias, flag, outCert, outCertLength);
// }
import "C"
import "unsafe"

// X509ExportCertificateFromStore экспортирует сертификата из хранилища
func (cli *Client) X509ExportCertificateFromStore() (string, error) {
	flag := 1
	outCertLength := 32768

	cert := C.malloc((C.ulong)(C.sizeof_char * outCertLength))
	defer C.free(cert)
	alias := C.CString("")
	defer C.free(unsafe.Pointer(alias))

	rc := (int)(C.X509ExportCertificateFromStore(
		alias,
		(C.int)(flag),
		(*C.char)(cert),
		(*C.int)(unsafe.Pointer(&outCertLength)),
	))
	resultCert := C.GoString((*C.char)(cert))

	return resultCert, cli.returnErr(rc)
}
