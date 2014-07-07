#!/bin/bash
set -e

# a helper script to run tests

if ! which nsqd >/dev/null; then
    echo "missing nsqd binary" && exit 1
fi

if ! which nsqlookupd >/dev/null; then
    echo "missing nsqlookupd binary" && exit 1
fi

# run nsqlookupd
LOOKUP_LOGFILE=$(mktemp -t nsqlookupd.XXXXXXX)
echo "starting nsqlookupd"
echo "  logging to $LOOKUP_LOGFILE"
nsqlookupd >$LOOKUP_LOGFILE 2>&1 &
LOOKUPD_PID=$!

cat >/tmp/cert.pem <<EOF
-----BEGIN CERTIFICATE-----
MIIEbjCCA1agAwIBAgIJAK6x7y6AwBmLMA0GCSqGSIb3DQEBBQUAMIGAMQswCQYD
VQQGEwJVUzERMA8GA1UECBMITmV3IFlvcmsxFjAUBgNVBAcTDU5ldyBZb3JrIENp
dHkxDDAKBgNVBAoTA05TUTETMBEGA1UEAxMKdGVzdC5sb2NhbDEjMCEGCSqGSIb3
DQEJARYUbXJlaWZlcnNvbkBnbWFpbC5jb20wHhcNMTMwNjI4MDA0MzQ4WhcNMTYw
NDE3MDA0MzQ4WjCBgDELMAkGA1UEBhMCVVMxETAPBgNVBAgTCE5ldyBZb3JrMRYw
FAYDVQQHEw1OZXcgWW9yayBDaXR5MQwwCgYDVQQKEwNOU1ExEzARBgNVBAMTCnRl
c3QubG9jYWwxIzAhBgkqhkiG9w0BCQEWFG1yZWlmZXJzb25AZ21haWwuY29tMIIB
IjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAnX0KB+svwy+yHU2qggz/EaGg
craKShagKo+9M9y5HLM852ngk5c+t+tJJbx3N954Wr1FXBuGIv1ltU05rU4zhvBS
25tVP1UIEnT5pBt2TeetLkl199Y7fxh1hKmnwJMG3fy3VZdNXEndBombXMmtXpQY
shuEJHKeUNDbQKz5X+GjEdkTPO/HY/VMHsxS23pbSimQozMg3hvLIdgv0aS3QECz
ydZBgTPThy3uDtHIuCpxCwXd/vDF68ATlYgo3h3lh2vxNwM/pjklIUhzMh4XaKQF
7m3/0KbtUcXfy0QHueeuMr11E9MAFNyRN4xf9Fk1yB97KJ3PJBTC5WD/m1nW+QID
AQABo4HoMIHlMB0GA1UdDgQWBBR3HMBws4lmYYSIgwoZsfW+bbgaMjCBtQYDVR0j
BIGtMIGqgBR3HMBws4lmYYSIgwoZsfW+bbgaMqGBhqSBgzCBgDELMAkGA1UEBhMC
VVMxETAPBgNVBAgTCE5ldyBZb3JrMRYwFAYDVQQHEw1OZXcgWW9yayBDaXR5MQww
CgYDVQQKEwNOU1ExEzARBgNVBAMTCnRlc3QubG9jYWwxIzAhBgkqhkiG9w0BCQEW
FG1yZWlmZXJzb25AZ21haWwuY29tggkArrHvLoDAGYswDAYDVR0TBAUwAwEB/zAN
BgkqhkiG9w0BAQUFAAOCAQEANOYTbanW2iyV1v4oYpcM/y3TWcQKzSME8D2SGFZb
dbMYU81hH3TTlQdvyeh3FAcdjhKE8Xi/RfNNjEslTBscdKXePGpZg6eXRNJzPP5K
KZPf5u6tcpAeUOKrMqbGwbE+h2QixxG1EoVQtE421szsU2P7nHRTdHzKFRnOerfl
Phm3NocR0P40Rv7WKdxpOvqc+XKf0onTruoVYoPWGpwcLixCG0zu4ZQ23/L/Dy18
4u70Hbq6O/6kq9FBFaDNp3IhiEdu2Cq6ZplU6bL9XDF27KIEErHwtuqBHVlMG+zB
oH/k9vZvwH7OwAjHdKp+1yeZFLYC8K5hjFIHqcdwpZCNIg==
-----END CERTIFICATE-----
EOF

cat >/tmp/key.pem <<EOF
-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEAnX0KB+svwy+yHU2qggz/EaGgcraKShagKo+9M9y5HLM852ng
k5c+t+tJJbx3N954Wr1FXBuGIv1ltU05rU4zhvBS25tVP1UIEnT5pBt2TeetLkl1
99Y7fxh1hKmnwJMG3fy3VZdNXEndBombXMmtXpQYshuEJHKeUNDbQKz5X+GjEdkT
PO/HY/VMHsxS23pbSimQozMg3hvLIdgv0aS3QECzydZBgTPThy3uDtHIuCpxCwXd
/vDF68ATlYgo3h3lh2vxNwM/pjklIUhzMh4XaKQF7m3/0KbtUcXfy0QHueeuMr11
E9MAFNyRN4xf9Fk1yB97KJ3PJBTC5WD/m1nW+QIDAQABAoIBACvtfKbIywG+hAf4
ad7skRjx5DcbA2e29+XnQfb9UgTXWd2SgrmoLi5OypBkCTzkKN3mfTo70yZfV8dC
Sxwz+9tfnTz0DssjhKThS+CiaFVCkeOfSfBfKSlCQUVHrSrh18CDhP+yvDlJwQTZ
zSQMfPcsh9bmJe2kqtQP7ZgUp1o+vaB8Sju8YYrO6FllxbdLRGm4pfvvrHIRRmXa
oVHn0ei0JpwoTY9kHYht4LNeJnbP/MCWdmcuv3Gnel7jAlhaKab5aNIGr0Xe7aIQ
iX6mpZ0/Rnt8o/XcTOg8l3ruIdVuySX6SYn08JMnfFkXdNYRVhoV1tC5ElWkaZLf
hPmj2yECgYEAyts0R0b8cZ6HTAyuLm3ilw0s0v0/MM9ZtaqMRilr2WEtAhF0GpHG
TzmGnii0WcTNXD7NTsNcECR/0ZpXPRleMczsL2Juwd4FkQ37h7hdKPseJNrfyHRg
VolOFBX9H14C3wMB9cwdsG4Egw7fE27WCoreEquHgwFxl1zBrXKH088CgYEAxr8w
BKZs0bF7LRrFT5pH8hpMLYHMYk8ZIOfgmEGVBKDQCOERPR9a9kqUss7wl/98LVNK
RnFlyWD6Z0/QcQsLL4LjBeZJ25qEMc6JXm9VGAzhXA1ZkUofVoYCnG+f6KUn8CuJ
/AcV2ZDFsEP10IiQG0hKsceXiwFEvEr8306tMrcCgYBLgnscSR0xAeyk71dq6vZc
ecgEpcX+2kAvclOSzlpZ6WVCjtKkDT0/Qk+M0eQIQkybGLl9pxS+4Yc+s2/jy2yX
pwsHvGE0AvwZeZX2eDcdSRR4bYy9ZixyKdwJeAHnyivRbaIuJ5Opl9pQGpoI9snv
1K9DTdw8dK4exKVHdgl/WwKBgDkmLsuXg4EEtPOyV/xc08VVNIR9Z2T5c7NXmeiO
KyiKiWeUOF3ID2L07S9BfENozq9F3PzGjMtMXJSqibiHwW6nB1rh7mj8VHjx9+Q0
xVZGFeNfX1r84mgB3uxW2LeQDhzsmB/lda37CC14TU3qhu2hawEV8IijE73FHlOk
Dv+fAoGAI4/XO5o5tNn5Djo8gHmGMCbinUE9+VySxl7wd7PK8w2VSofO88ofixDk
NX94yBYhg5WZcLdPm45RyUnq+WVQYz9IKUrdxLFTH+wxyzUqZCW7jgXCvWV+071q
vqm9C+kndq+18/1VKuCSGWnF7Ay4lbsgPXY2s4VKRxcb3QpZSPU=
-----END RSA PRIVATE KEY-----
EOF

# run nsqd configured to use our lookupd above
rm -f *.dat
NSQD_LOGFILE=$(mktemp -t nsqlookupd.XXXXXXX)
echo "starting nsqd --data-path=/tmp --lookupd-tcp-address=127.0.0.1:4160 --tls-cert=/tmp/cert.pem --tls-key=/tmp/key.pem"
echo "  logging to $NSQD_LOGFILE"
nsqd --data-path=/tmp --lookupd-tcp-address=127.0.0.1:4160 --tls-cert=/tmp/cert.pem --tls-key=/tmp/key.pem >$NSQD_LOGFILE 2>&1 &
NSQD_PID=$!

sleep 0.3

cleanup() {
    echo "killing nsqd PID $NSQD_PID"
    kill -s TERM $NSQD_PID || cat $NSQD_LOGFILE
    echo "killing nsqlookupd PID $LOOKUPD_PID"
    kill -s TERM $LOOKUPD_PID || cat $LOOKUP_LOGFILE
}
trap cleanup INT TERM EXIT

go test -v -timeout 15s
