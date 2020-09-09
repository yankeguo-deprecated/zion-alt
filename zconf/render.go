package zconf

import (
	"bytes"
	"encoding/json"
	"errors"
	"net"
	"os"
	"os/user"
	"strconv"
	"strings"
	"text/template"
)

var templateFuncs = map[string]interface{}{
	// built-in functions
	"netResolveIPAddr":    net.ResolveIPAddr,
	"osHostname":          os.Hostname,
	"osUserCacheDir":      os.UserCacheDir,
	"osUserConfigDir":     os.UserConfigDir,
	"osUserHomeDir":       os.UserHomeDir,
	"osGetegid":           os.Getegid,
	"osGetenv":            os.Getenv,
	"osGeteuid":           os.Geteuid,
	"osGetgid":            os.Getgid,
	"osGetgroups":         os.Getgroups,
	"osGetpagesize":       os.Getpagesize,
	"osGetpid":            os.Getpid,
	"osGetppid":           os.Getppid,
	"osGetuid":            os.Getuid,
	"osGetwd":             os.Getwd,
	"osTempDir":           os.TempDir,
	"osUserLookupGroup":   user.LookupGroup,
	"osUserLookupGroupId": user.LookupGroupId,
	"osUserCurrent":       user.Current,
	"osUserLookup":        user.Lookup,
	"osUserLookupId":      user.LookupId,
	"stringsContains":     strings.Contains,
	"stringsFields":       strings.Fields,
	"stringsIndex":        strings.Index,
	"stringsLastIndex":    strings.LastIndex,
	"stringsHasPrefix":    strings.HasPrefix,
	"stringsHasSuffix":    strings.HasSuffix,
	"stringsRepeat":       strings.Repeat,
	"stringsReplaceAll":   strings.ReplaceAll,
	"stringsSplit":        strings.Split,
	"stringsToLower":      strings.ToLower,
	"stringsToUpper":      strings.ToUpper,
	"stringsTrimPrefix":   strings.TrimPrefix,
	"stringsTrimSpace":    strings.TrimSpace,
	"stringsTrimSuffix":   strings.TrimSuffix,
	"strconvParseBool":    strconv.ParseBool,
	"strconvParseInt":     strconv.ParseInt,
	"strconvParseUint":    strconv.ParseUint,
	"strconvFormatBool":   strconv.FormatBool,
	"strconvFormatInt":    strconv.FormatInt,
	"strconvFormatUint":   strconv.FormatUint,
	"strconvAtoi":         strconv.Atoi,
	"strconvItoa":         strconv.Itoa,
	"jsonMarshal": func(v interface{}) (s string, err error) {
		var buf []byte
		if buf, err = json.Marshal(v); err != nil {
			return
		}
		s = string(buf)
		return
	},

	// extra functions
	"stringsFromBytes": func(buf []byte) string {
		return string(buf)
	},
	"osHostnameSequentialSuffix": func() (id int64, err error) {
		var hostname string
		if hostname = os.Getenv("HOSTNAME"); hostname == "" {
			if hostname, err = os.Hostname(); err != nil {
				return
			}
		}
		splits := strings.Split(hostname, "-")
		if len(splits) < 2 {
			err = errors.New("no sequence suffix found in os.Hostname()")
			return
		}
		id, err = strconv.ParseInt(splits[len(splits)-1], 10, 64)
		return
	},
}

func environ() map[string]string {
	out, envs := make(map[string]string), os.Environ()
	for _, kv := range envs {
		kvs := strings.SplitN(kv, "=", 2)
		if len(kvs) == 2 {
			out[strings.TrimSpace(kvs[0])] = strings.TrimSpace(kvs[1])
		}
	}
	return out
}

func Render(tpl []byte) (out []byte, err error) {
	tmpl := template.New("__main__").Funcs(templateFuncs).Option("missingkey=zero")
	if tmpl, err = tmpl.Parse(string(tpl)); err != nil {
		return
	}
	buf := &bytes.Buffer{}
	if err = tmpl.Execute(buf, map[string]interface{}{
		"Env": environ(),
	}); err != nil {
		return
	}
	out = buf.Bytes()
	return
}
