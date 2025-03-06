package log

import stlog "log"

var log *stlog.Logger

type fileLog string //把日志写到文件系统

func (fl fileLog) Write(date []byte) (int, error) {
	f, err := os.OpenFile(string(fl), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		return 0, err
	}
	defer f.colse()
	return f.Write(data)
}

func Run(destination string) {
	log = stlog.New(fileLog(destination), "go", stlog.LstdFlags)
}

func RegisterHandlers() {
	http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			msg, err := ioutil.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			write(string(msg))
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	})
}
func write(message string) {
	log.Printf("%v\n", message)
}
