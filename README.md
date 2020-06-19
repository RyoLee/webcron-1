# webcron
------------
## ONLY FOR K8S ENV WITH INGRESS ##

1.load db

$ mysql -u username -p -D webcron < install.sql

2.mount your app.conf on /go/src/github.com/RyoLee/webcron/conf 

3.add to your ingress

"nginx.ingress.kubernetes.io/rewrite-target": "/$2"

and

"path": "/webcron(/|$)(.*)"

4.check

http://your-ingress-svc/webcron

user：admin

passwd：admin888
