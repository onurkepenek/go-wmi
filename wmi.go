package gowmi

/*
#cgo CFLAGS: -g -Wall -I/usr/include
#cgo LDFLAGS: -lc -lm -L./ -lwmiclient  -lsamba-static -L/usr/lib/x86_64-linux-gnu/heimdal/ -lgssapi  -lgnutls -lgnutls-openssl  -lpopt -lkrb5 -lasn1 -lcom_err -ldl -lhdb
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include "openvas_wmi_interface.h"

static char** makeCharArray(int size) {
        return calloc(sizeof(char*), size);
}

static void setArrayString(char **a, char *s, int n) {
        a[n] = s;
}

static void freeCharArray(char **a, int size) {
        int i;
        for (i = 0; i < size; i++)
                free(a[i]);
        free(a);
}
*/
import "C"

import (
	"fmt"
	"strings"
)

func Query(host string, user string, pass string, namespace string, query string) ([]map[string]interface{}, error) {

	var args_list [5]string

	args_list[0] = "wmic"
	args_list[1] = "-U"
	args_list[2] = fmt.Sprintf("%s%%%s", user, pass)
	args_list[3] = fmt.Sprintf("//%s%s", host, "[sign]")
	args_list[4] = namespace

	cargs := C.makeCharArray(C.int(len(args_list)))
	defer C.freeCharArray(cargs, C.int(len(args_list)))
	for i, s := range args_list {
		C.setArrayString(cargs, C.CString(s), C.int(i))
	}

	wmi_t := C.wmi_connect(C.int(len(args_list)), cargs)
	if wmi_t == nil {
		return nil, fmt.Errorf("unable to connect server")
	}
	defer C.wmi_close(wmi_t)
	var outval *C.char
	ret := C.wmi_query(wmi_t, C.CString(query), &outval)

	if ret != 0 {
		return nil, fmt.Errorf("wmi query error")
	}

	if outval == nil {
		return nil, fmt.Errorf("wmi query error")
	}
	//fmt.Println(C.GoString(outval))

	var response []map[string]interface{}
	res := C.GoString(outval)
	lines := strings.Split(res, "\n")

	var header []string
	for i, line := range lines {

		line = strings.TrimSpace(line)
		if i == 0 {
			header = strings.Split(line, "|/|")
			continue
		}
		item := make(map[string]interface{})
		fields := strings.Split(line, "|/|")
		if len(fields) < 1 {
			continue
		}

		if len(header) != len(fields) {
			continue
		}

		for j, field := range fields {

			if len(field) < 1 {
				continue
			}

			if field == "(null)" {
				item[header[j]] = nil
			} else if strings.Contains(field, ";") {
				item[header[j]] = strings.Split(field, ";")
			} else {
				item[header[j]] = field
			}

		}

		if len(item) > 0 {
			response = append(response, item)
		}

	}

	return response, nil
}
