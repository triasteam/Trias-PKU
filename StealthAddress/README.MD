# BIP63 Stealth Address
## done
original plan
receiver address:(a, A) (b, B)
sender : (r, R) R=rG, P=H(rA)G+B
msg in block: R, P

scan methods: P'=(H(aR)+b)G=H(arG)G+bG=H(rA)G+B=P

# test
go run main.go

the cost time of generation and verify Stealth Address

## todo
improve the efficiency of original plan

