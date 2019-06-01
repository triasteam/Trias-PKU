# BIP63 Stealth Address
## done
original plan   
receiver address:(a, A) (b, B)   
sender : (r, R) R=rG, P=H(rA)G+B   
msg in block: R, P   

scan methods: P'=(H(aR)+b)G=H(arG)G+bG=H(rA)G+B=P

## improve
expose E'=E(aR) with R,P together

scan methods: 
if E(aR) == E'
    P'==P
else
    exit

so reduced validation steps by half.

# test
go run main.go

the cost time of generation and verify Stealth Address   
    Origin Generate Cost Time 138.774237ms
    Origin Generate and Verify Cost Time 234.081016ms
    Improve Generate Cost Time 130.661866ms
    Phase 01 Cost Time 73.712819ms
    Phase 02 Cost Time 1.342067ms
    Phase 03 Cost Time 31.253998ms
so reduce cost time from 105 to 75

## todo
improve the efficiency of original plan

